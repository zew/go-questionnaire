// Package tpl parses bundles of related templates and keeps them in a map;
// bundles can be master-layouts with several content-templates;
// http requests derive clones from these prepared bundles;
// complemented with request specific template funcs;
// executeTemplate(dynamicName, data) replaces the static
// template(constantName) func.
package tpl

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/util"
)

var staticTplFuncs = template.FuncMap{
	"toHtml": func(arg string) template.HTML { return template.HTML(arg) },
	"prefix": cfg.Pref,
	"cfg":    func() *cfg.ConfigT { return cfg.Get() },
	"addint": func(i1, i2 int) int { return i1 + i2 },
	// dummies, to make parsing work
	"executeTemplate": func(name string, data interface{}) (ret template.HTML, err error) { return },
	"sess":            func() (sess sessx.SessT, err error) { return },
}

// StaticFuncMap returns the static funcs, every template should have.
func StaticFuncMap() template.FuncMap {
	return staticTplFuncs
}

// A parsed bundle of coherent templates.
// A base to clone from.
type baseTplT struct {
	IsParsed bool
	*template.Template
}

// coherent templates organized by a string, the bundle type
// for instance main.html or style.css
var cloneBase = map[string]*baseTplT{}

// Parse prepares a base, from which to clone.
// bundle could be main.html =>  for main.html, main_*.html and _*.html
// bundle could be style.css =>  for style.css, style_*.css and _*.css
func (bt *baseTplT) Parse(bundle string) *template.Template {

	if !bt.IsParsed {

		// An independent map for template funcs - because dynamic funcs will be added later - and we dont want interdependencies between requests
		mp := template.FuncMap{}
		for key, fc := range staticTplFuncs {
			mp[key] = fc
		}

		var err error
		tplBase := template.New(bundle)
		tplBase = tplBase.Funcs(mp)

		bt.Template, err = tplBase.ParseFiles(filepath.Join(".", "templates", bundle)) // i.e. main.html
		if err != nil {
			log.Fatal(err)
		}

		// We want to keep the cost of the cloning operations (code below) minimal.
		// => Keep bundles as small as possible
		ext := filepath.Ext(bundle) // i.e. .html
		mask := strings.TrimSuffix(bundle, ext) + "_*" + ext
		additional1, err := bt.Template.ParseGlob(filepath.Join(".", "templates", mask)) // i.e. main_*.html
		if err != nil {
			log.Print(err)
		} else {
			bt.Template = additional1
		}

		helpers := "_*" + ext
		additional2, err := bt.Template.ParseGlob(filepath.Join(".", "templates", helpers)) // i.e. _dropdown.html
		if err != nil {
			log.Print(err)
		} else {
			bt.Template = additional2
		}

		bt.IsParsed = true
	}

	return bt.Template
}

// Get returns a clone of the template bundle.
// It adds bundle-specific dynamic template execution: {{executeTemplate "t"}}.
// It adds request specific session access.
// Thus, each request gets its own clone. Thus more expensive than GetStatic().
func Get(w http.ResponseWriter, r *http.Request, bundle string) *template.Template {

	if _, ok := cloneBase[bundle]; !ok {
		panic(fmt.Sprintf("Template bundle %v must be prepared on initialization.", bundle))
	}

	if !cfg.Get().IsProduction {
		bt := &baseTplT{}
		bt.Parse(bundle)
		cloneBase[bundle] = bt
	}

	tpl, err := cloneBase[bundle].Template.Clone()
	if err != nil {
		log.Fatal(err)
	}

	// Create independent func map; prevent conflicts with other request adding dynamic functions
	mp := template.FuncMap{}
	for key, fc := range staticTplFuncs {
		mp[key] = fc
	}
	// Add dynamic tpl funcs.
	//
	// Executing a named template from this particular clone instance.
	mp["executeTemplate"] = func(name string, data interface{}) (ret template.HTML, err error) {
		buf := bytes.NewBuffer([]byte{})
		err = tpl.ExecuteTemplate(buf, name, data)
		if err != nil {
			log.Printf("callTemplate erred: %v", err)
			return
		}
		ret = template.HTML(buf.String())
		return
	}
	// response writer + request specific closure
	// Use sess.EffectiveParam("name") => sess.EffectiveParam "session-test-key"
	mp["sess"] = func() (sess *sessx.SessT, err error) {
		sessVal := sessx.New(w, r)
		sess = &sessVal
		return
	}

	tpl = tpl.Funcs(mp)
	return tpl

}

// GetStatic is the same as Get()
// Without dynamic funcs.
// Without the need for extra cloning.
// The same template can be executed for all requests.
// Useful for serving dynamic CSS files with few app specific variations.
func GetStatic(w http.ResponseWriter, r *http.Request, bundle string) *template.Template {
	if _, ok := cloneBase[bundle]; !ok {
		panic(fmt.Sprintf("Template bundle %v must be prepared on initialization.", bundle))
	}
	if !cfg.Get().IsProduction {
		bt := &baseTplT{}
		bt.Parse(bundle)
		cloneBase[bundle] = bt
	}

	tpl, err := cloneBase[bundle].Template.Clone()
	if err != nil {
		log.Fatal(err)
	}
	tpl = tpl.Funcs(staticTplFuncs)
	return tpl
}

// Parse is meant for bootstrapping the application.
// It fills the cloneBase.
func Parse(bundles ...string) {
	cloneBase = map[string]*baseTplT{}
	for _, bundle := range bundles {
		bt := &baseTplT{}
		bt.Parse(bundle)
		cloneBase[bundle] = bt
	}
}

// ParseH is a convenience handler to parse all base templates anew.
func ParseH(w http.ResponseWriter, r *http.Request) {
	_, loggedIn, err := lgn.LoggedInCheck(w, r, "admin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !loggedIn {
		http.Error(w, "admin login required for this function", http.StatusInternalServerError)
		return
	}
	for bundle := range cloneBase {
		bt := &baseTplT{}
		bt.Parse(bundle)
		cloneBase[bundle] = bt
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("templates reloaded"))
	w.Write([]byte(util.IndentedDump(cloneBase)))
}
