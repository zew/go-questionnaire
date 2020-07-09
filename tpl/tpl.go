// Package tpl parses and bundles templates related by `/templates/bundles.json`
// and stores them into a map `cache`;
// i.e. bundles of master-layouts with content-templates;
// http requests dont need clones for specific template funcs;
// function executeTemplate(...dynamicName...) replaces {{template  constantName}}.
// Parsing and bundling is done at application init time,
// avoiding mutexing the `cache` map.
package tpl

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"path"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/handler"
	"github.com/zew/go-questionnaire/lgn"
)

// general template funcs
var fcByKey = func(k string) handler.Info {
	return handler.Infos().ByKey(k)
}
var fcURLByKey = func(k string) string {
	return cfg.Pref(handler.Infos().URLByKey(k))
}

// fcNav renders the nav core; it is called by the nav template
var fcNav = func(r *http.Request) template.HTML {
	bts := &bytes.Buffer{}
	w := httptest.NewRecorder()
	isAdmin := false
	_, isLogin, err := lgn.LoggedInCheck(w, r, "admin")
	if isLogin && err == nil {
		isAdmin = true
	}
	handler.Tree().NavHTML(bts, r, isLogin, isAdmin, 0) // the dynamic part
	return template.HTML(bts.String())
}

// fcLogin returns the login username or empty string
var fcLogin = func(r *http.Request) template.HTML {
	w := httptest.NewRecorder()
	l, isLoggedIn, _ := lgn.LoggedInCheck(w, r)
	if !isLoggedIn {
		return template.HTML("")
	}
	return template.HTML(l.User)
}

// fcExecBundledTemplate is used in staticTplFuncs;
// thus cannot refer to funcs with nested in-package funcs - precise reason obscure
func fcExecBundledTemplate(tName string, mp map[string]interface{}) (template.HTML, error) {
	w := bytes.NewBuffer([]byte{})
	var err error
	err = cache[tName].ExecuteTemplate(w, tName, mp)
	if err != nil {
		err = errors.Wrap(err, fmt.Sprintf("fcExecBundledTemplate erred: %v", err))
		log.Print(err)
		return template.HTML(err.Error()), err
	}
	return template.HTML(w.String()), nil
}

// staticTplFuncs cannot refer to funcs with nested in-package funcs - precise reason obscure
var staticTplFuncs = template.FuncMap{
	"toHTML":          func(arg string) template.HTML { return template.HTML("Nogo - gosec violation") },
	"formToken":       func() template.HTMLAttr { return template.HTMLAttr(lgn.FormToken()) },
	"cfg":             func() *cfg.ConfigT { return cfg.Get() }, // access to config
	"byKey":           fcByKey,                                  // template usage {{ index (   byKey "landing-page" ).Urls  0  }} - no prefix applied yet
	"urlByKey":        fcURLByKey,                               // template usage {{        urlByKey "landing-page"            }} - prefix already applied
	"nav":             fcNav,                                    // template usage {{ nav .Req  }}
	"lgn":             fcLogin,                                  // template usage {{ lgn .Req  }}
	"executeTemplate": fcExecBundledTemplate,                    // template usage {{ executeTemplate "myTpl" . }} or {{ executeTemplate .DynTemplate .}}
	"addint":          func(i1, i2 int) int { return i1 + i2 },
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

// general template with general template funcs for cloning;
// no sync - see tpl()
var cache = map[string]*template.Template{}

// bundles of templates
// overriden in init() / TemplatesPreparse()
// loaded from bundles.json
var bundles = map[string][]string{
	"main-desktop.html": {
		"nav-css-2020.html",
		"example-01.html",
		"example-02.html",
	},
	"main-content.html": {
		"main-content-header.html",
		"main-content-sidebar.html",
	},
}

// bundle appends a parsed template base with another parsed template b
func bundle(base *template.Template, b string) error {

	pth := path.Join(".", "templates", b) // not filepath; cloudio always has forward slash
	bCnt, err := cloudio.ReadFile(pth)
	if err != nil {
		msg := fmt.Sprintf("cannot open bundle template %v: %v", pth, err)
		return errors.Wrap(err, msg)
	}

	// from now on without extension
	b = strings.TrimSuffix(b, path.Ext(b))

	tB := template.New(b)
	tB = tB.Funcs(staticTplFuncs)
	tB2, err := tB.Parse(string(bCnt))
	if err != nil {
		msg := fmt.Sprintf("parsing failed for bundle template %v: %v", pth, err)
		return errors.Wrap(err, msg)
	}

	// adding the *parsed* template
	// callable via   {{ template b . }}
	base, err = base.AddParseTree(b, tB2.Tree)
	if err != nil {
		msg := fmt.Sprintf("failure adding parse tree of bundle template %v: %v", pth, err)
		return errors.Wrap(err, msg)
	}

	return nil

}

// adding a closure over the current request;
// to access the current URL for instance;
// however this is incompatible with using cached pre-parsed templates;
// we need to add dynamic stuff as params instead;
// see nav(*http.Request) as example.
func obsoleteAddDynamicFuncs(t *template.Template, r *http.Request) {
}

// tpl returns a parsed template by name
// either pre-parsed from cache, or freshly parsed;
// called for every template at app *init* time;
// thus no sync mutex is required
func tpl(r *http.Request, tName string) (*template.Template, error) {

	tDerived, ok := cache[tName] // template from cache...
	if !ok || !cfg.Get().IsProduction {

		// or parse it anew
		pth := path.Join(".", "templates", tName) // not filepath; cloudio always has forward slash
		cnts, err := cloudio.ReadFile(pth)
		if err != nil {
			msg := fmt.Sprintf("cannot open template %v: %v", pth, err)
			return nil, errors.Wrap(err, msg)
		}
		// tDerived, err = base.ParseFiles(pth)

		base := template.New(tName)
		base = base.Funcs(staticTplFuncs)

		tDerived, err = base.Parse(string(cnts))
		if err != nil {
			msg := fmt.Sprintf("parsing failed for %v: %v", pth, err)
			return nil, errors.Wrap(err, msg)
		}

		// bundling templates together
		for _, bdl := range bundles[tName] {
			err = bundle(tDerived, bdl)
			if err != nil {
				msg := fmt.Sprintf("bundling failed for %v:\n %v", bdl, err)
				return nil, errors.Wrap(err, msg)
			}
			// log.Printf("\ttemplate %v bundled with %v:\n\t%v", tName, bdl, base.DefinedTemplates())
			// log.Printf("\ttemplate %v bundled with %v", tName, bdl)
		}

		// cache[tName] = tDerived // caching only once in preparseTemplates() to avoid contention
		// log.Printf("  freshly parsed  - template %-30v", tName)
	} else {
		// log.Printf("  from cache      - template %-30v", tName)
	}

	// funcs can only be added *before* parsing
	obsoleteAddDynamicFuncs(tDerived, r)

	return tDerived, nil
}

// Exec template with map of contents into writer w;
// cnt as io.Writer would be more efficient?
func Exec(w io.Writer, r *http.Request, mp map[string]interface{}, tName string) {

	t, err := tpl(r, tName)
	if err != nil {
		log.Printf("parsing template %q error: %v", tName, err)
		fmt.Fprintf(w, "parsing template %q error: %v", tName, err)
		return
	}

	//
	if mp != nil {
		// mp["Cfg"] = cfg.Get()  // cfg is made accessible via funcMap
		mp["Req"] = r
		if _, ok := mp["HTMLTitle"]; !ok {
			mp["HTMLTitle"] = "html title"
		}

		// string to template.HTML
		wrapThem := []string{"Content", "Navigation"}
		for _, val := range wrapThem {
			if _, ok := mp[val]; ok {
				cnv, ok1 := mp[val].(string)
				if !ok1 {
					mp[val] = template.HTML(fmt.Sprintf("key %v must be string, in order to be converted to template.HTML", val))
				} else {
					mp[val] = template.HTML(cnv)
				}
			}
		}
	}

	err = t.ExecuteTemplate(w, tName, mp)
	if err != nil {
		log.Printf("template execution error: %v", err)
		fmt.Fprintf(w, "template execution error: %v", err)
	}
	// fmt.Fprintf(w, "\nUA: %v", r.Header.Get("User-Agent"))

}

// ExecContent is a simplified version of Exec()
// with only one content element
func ExecContent(w io.Writer, r *http.Request, cnt, tName string) {
	mp := map[string]interface{}{}
	mp["Content"] = cnt
	Exec(w, r, mp, tName)
}
