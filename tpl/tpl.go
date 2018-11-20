// Package tpl parses bundles of related templates and keeps them in a map;
// bundles can be master-layouts with several content-templates;
// http requests dont need clones for specific template funcs;
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
	"reflect"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/util"
)

var staticTplFuncs = template.FuncMap{
	"formToken": func() template.HTMLAttr { return template.HTMLAttr(lgn.FormToken()) },
	"toHtml":    func(arg string) template.HTML { return template.HTML(arg) },
	"prefix":    cfg.Pref,
	"cfg":       func() *cfg.ConfigT { return cfg.Get() },
	"addint":    func(i1, i2 int) int { return i1 + i2 },
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
	// exists checks whether the struct 'data'
	// has a field or method name.
	// Usage {{if exists . "Q"}} ... {{end}}
	// 'bad' but inevitable in *general purpose* layout templates.
	// stackoverflow.com/questions/44675087/
	"exists": func(data interface{}, name string) bool {
		v := reflect.ValueOf(data)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return false
		}
		return v.FieldByName(name).IsValid()
	},
}

// StaticFuncMap returns the static funcs every template should have.
func StaticFuncMap() template.FuncMap {
	return staticTplFuncs
}

// A parsed bundle of coherent templates.
// Notice the Template.Tree .
// Remember: A parsed template has DefinedTemplates().
// With executeTemplate(..., name,...) any of these defined templates can be executed.
// It is therefore helpful to imagine *template.Template as a bunch or bundle of templates.
// bundleT is also a base to clone from, though cloning should never be necessary.
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
// But this does not matter, since it is a pointer and since we dont clone it.
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
		if bundle == "main_desktop.html" || bundle == "main_mobile.html" {
			mask = "*" + ext // main.html pulls in all *.html templates
		}
		additional1, err := bt.Template.ParseGlob(filepath.Join(".", "templates", mask)) // i.e. main_*.html
		if err != nil {
			if !strings.Contains(err.Error(), "html/template: pattern matches no files") { // Yes - I would love to check with IsNotExist(err), but I cannot change the standard library
				log.Print(err)
			}
		} else {
			bt.Template = additional1
		}

		helpers := "_*" + ext
		additional2, err := bt.Template.ParseGlob(filepath.Join(".", "templates", helpers)) // i.e. _dropdown.html
		if err != nil {
			if !strings.Contains(err.Error(), "html/template: pattern matches no files") { // Yes - I would love to check with IsNotExist(err), but I cannot change the standard library
				log.Print(err)
			}
		} else {
			bt.Template = additional2
		}

		dt := bt.Template.DefinedTemplates()
		dt = strings.Replace(dt, "; defined templates are:", "", -1)
		dt = strings.Replace(dt, "\n", "", -1)
		// log.Printf("Bundle %-28v %v", bt.Template.Name(), dt)

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
		parsedBundles[bundle] = bt // this causes: "fatal error: concurrent map writes" in development mode
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
