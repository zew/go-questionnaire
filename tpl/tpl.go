// Package tpl parses bundles of related templates and keeps them in a map;
// bundles can be master-layouts with several content-templates;
// http requests dont need clones or specific template funcs;
// executeTemplate(bundle, dynamicName, data) replaces the static
// {{template  constantName}}.
// Parsing bundles into parsedBundles should be done at application init time,
// to avoid mutexing the parsedBundles map.
// The func Parse() is exposed for this purpose - for bootstrapping.
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
	"github.com/zew/util"
)

var staticTplFuncs = template.FuncMap{
	"toHtml": func(arg string) template.HTML { return template.HTML(arg) },
	"prefix": cfg.Pref,
	"cfg":    func() *cfg.ConfigT { return cfg.Get() },
	"addint": func(i1, i2 int) int { return i1 + i2 },
	"executeTemplate": func(tplBundle *template.Template, name string, data interface{}) (ret template.HTML, err error) {
		buf := bytes.NewBuffer([]byte{})
		err = tplBundle.ExecuteTemplate(buf, name, data)
		if err != nil {
			log.Printf("callTemplate erred: %v", err)
			return
		}
		ret = template.HTML(buf.String())
		return
	},
}

// StaticFuncMap returns the static funcs, every template should have.
func StaticFuncMap() template.FuncMap {
	return staticTplFuncs
}

// A parsed bundle of coherent templates.
// A base to clone from.
type bundleT struct {
	IsParsed bool
	*template.Template
}

// coherent templates organized by a string, the bundle type
// for instance main.html or style.css
var parsedBundles = map[string]*bundleT{}

// Parse prepares a bunch of templates (bundle) for execution.
// bundles are grouped by name:
//      bundle could be docs.html =>  for docs.html, docs_*.html and _*.html
// This requires consistent naming.
//
// Or bundles can be grouped by extension.
//      bundle could be style.css =>  for style.css, style_*.css and _*.css
//
// Finally: main.html bundles all html templates.
//     bundle would be main.html =>  for main.html, *.html
// The parsed bundle will be *big* and will be used for *many* request.
// But this does *not* matter, since it is a pointer and since we dont clone it.
//
// Our bundles are completely threadsafe - but are still able
// to *dynamically*  xecute sub templates by calling executeTemplate.
//
// The price is a bulky struct with template data.
// see main.TplDataT
func (bt *bundleT) Parse(bundle string) *template.Template {

	if !bt.IsParsed {

		// An independent map for template funcs - because dynamic funcs might be added later - and we dont want interdependencies between requests
		// Template cloning would obliterate this
		mp := template.FuncMap{}
		for key, fc := range staticTplFuncs {
			mp[key] = fc
		}

		tplBase := template.New(bundle)
		tplBase = tplBase.Funcs(mp)

		var err error
		bt.Template, err = tplBase.ParseFiles(filepath.Join(".", "templates", bundle)) // i.e. main.html
		if err != nil {
			log.Fatal(err)
		}

		// We want to keep the cost of the cloning operations (code below) minimal.
		// => Keep bundles as small as possible
		ext := filepath.Ext(bundle) // i.e. .html
		mask := strings.TrimSuffix(bundle, ext) + "_*" + ext
		if bundle == "main.html" {
			mask = "*" + ext // main.html pulls in all *.html templates
		}
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

// Get the named template bundle.
func Get(w http.ResponseWriter, r *http.Request, bundle string) *template.Template {

	if _, ok := parsedBundles[bundle]; !ok {
		panic(fmt.Sprintf("Template bundle %v must be prepared on initialization.", bundle))
	}
	if !cfg.Get().IsProduction {
		bt := &bundleT{}
		bt.Parse(bundle)
		parsedBundles[bundle] = bt
	}

	tpl := parsedBundles[bundle].Template
	return tpl
}

// GetStatic is the same as Get()
func GetStatic(w http.ResponseWriter, r *http.Request, bundle string) *template.Template {
	return Get(w, r, bundle)
}

// Parse is meant for bootstrapping the application.
// It fills parsedBundles map.
func Parse(bundles ...string) {
	parsedBundles = map[string]*bundleT{}
	for _, bundle := range bundles {
		bt := &bundleT{}
		bt.Parse(bundle)
		parsedBundles[bundle] = bt
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
	for bundle := range parsedBundles {
		bt := &bundleT{}
		bt.Parse(bundle)
		parsedBundles[bundle] = bt
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("templates reloaded"))
	w.Write([]byte(util.IndentedDump(parsedBundles)))
}
