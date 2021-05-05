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
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
)

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

// general template with general template funcs for cloning;
// no sync - see tpl()
var cache = map[string]*template.Template{}

// bundles of templates
// overridden in init() / TemplatesPreparse()
// loaded from bundles.json
var bundles = map[string][]string{
	"layout.html": {
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
func bundle(base *template.Template, extend string) error {

	pth := path.Join(".", "templates", extend) // not filepath; cloudio always has forward slash
	bCnt, err := cloudio.ReadFile(pth)
	if err != nil {
		msg := fmt.Sprintf("cannot open bundle template %v: %v", pth, err)
		return errors.Wrap(err, msg)
	}

	// either everything with extension - or everything without
	//   we do the latter
	// extend = strings.TrimSuffix(extend, path.Ext(extend))

	tB := template.New(extend)
	tB = tB.Funcs(staticTplFuncs)
	tB2, err := tB.Parse(string(bCnt))
	if err != nil {
		msg := fmt.Sprintf("parsing failed for bundle template %v: %v", pth, err)
		return errors.Wrap(err, msg)
	}

	// adding the *parsed* template
	// callable via   {{ template b . }}
	base, err = base.AddParseTree(extend, tB2.Tree)
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

// SiteCore returns only the non-numerical part of the site name;
// for instance 'pat' from 'pat0';
// used for sharing the same CSS files and CSS settings among multiple questionnaires
func SiteCore(site string) (string, string) {
	firstDigit := false
	cr := strings.Builder{}
	vr := strings.Builder{}
	for _, rn := range site {
		if !firstDigit &&
			rn >= 48 &&
			rn <= 57 {
			firstDigit = true
		}
		if firstDigit {
			vr.WriteRune(rn)
		} else {
			cr.WriteRune(rn)
		}
	}
	return cr.String(), vr.String()
}

// Get returns a parsed bundle of templates by name
// either pre-parsed from cache, or freshly parsed;
// called for every template at app *init* time;
// thus no sync mutex is required
func Get(tName string) (*template.Template, error) {

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
	// obsoleteAddDynamicFuncs(tDerived, r)

	return tDerived, nil
}

/*
Exec template with map of contents into writer w;
cnt as io.Writer would be more efficient?

Automatically created keys
		Req
		Sess
		L
		LangCode
		Site

		CSSSite

		HTMLTitle
		LogoTitle


Keys expected to be supplied by caller
		Content

*/
func Exec(w io.Writer, r *http.Request, mp map[string]interface{}, tName string) {

	t, err := Get(tName)
	if err != nil {
		log.Printf("parsing template %q error: %v", tName, err)
		fmt.Fprintf(w, "parsing template %q error: %v", tName, err)
		return
	}

	// check for associated templates - not applicable
	if t.Lookup(tName) == nil && t.Name() != tName {
		log.Printf("template bundle %q does not contain %q: %v", tName, tName, err)
		fmt.Fprintf(w, "template bundle %q does not contain %q: %v", tName, tName, err)
		return
	}

	//
	if mp == nil {
		mp = map[string]interface{}{}
	}

	// set certain keys automatically

	// mp["Cfg"] = cfg.Get()  // cfg is made accessible via funcMap

	mp["Req"] = r

	if _, ok := mp["Sess"]; !ok {
		mp["Sess"] = sessx.New(w, r)
	}

	if _, ok := mp["L"]; !ok {
		l, _, err := lgn.LoggedInCheck(w, r)
		if err != nil {
			log.Printf("Login by hash error: %v", err)
			fmt.Fprintf(w, "Login by hash error: %v", err)
		}
		// if !isLoggedIn { // valid condition
		mp["L"] = l
	}

	//
	if _, ok := mp["LangCode"]; !ok {
		mp["LangCode"] = cfg.Get().LangCodes[0]
	}

	// Site - either explicit - or derived from questionnaire type
	if _, ok := mp["Site"]; !ok {

		if qIntf, ok2 := mp["Q"]; ok2 {
			q, ok3 := qIntf.(*qst.QuestionnaireT)
			if ok3 && q.Survey.Type != "" {
				mp["Site"] = q.Survey.Type
			} else {
				mp["Site"] = "no-site-specified-a"
			}
		} else {
			mp["Site"] = "no-site-specified-b"
		}

	}

	site := mp["Site"].(string)
	core, _ := SiteCore(site)
	mp["SiteCore"], _ = SiteCore(site)

	// mp["CSSSite"] must be of type cfg.[]cssVar; we only check for existence
	if _, ok := mp["CSSSite"]; !ok {
		mp["CSSSite"] = cfg.Get().CSSVarsSite[core]
	} else {
		mp["CSSSite"] = cfg.Get().CSSVars
	}

	if _, ok := mp["HTMLTitle"]; !ok {
		if site, okSite := mp["Site"]; !okSite {
			mp["HTMLTitle"] =
				cfg.Get().MpSite[site.(string)]["app_org"].Tr(mp["LangCode"].(string)) +
					" " +
					cfg.Get().MpSite[site.(string)]["app_label"].Tr(mp["LangCode"].(string))

		} else {
			mp["HTMLTitle"] =
				cfg.Get().Mp["app_org"].Tr(mp["LangCode"].(string)) +
					" " +
					cfg.Get().Mp["app_label"].Tr(mp["LangCode"].(string))
		}
	}

	if _, ok := mp["LogoTitle"]; !ok {
		mp["LogoTitle"] = mp["HTMLTitle"]
	}

	if _, ok := mp["Content"]; !ok {
		mp["Content"] = "<p style='color:red;' >Warning: no content supplied</p>\n"
	}

	// string to template.HTML
	wrapThem := []string{"Content", "Navigation"}
	for _, val := range wrapThem {
		if _, ok := mp[val]; ok {
			cnv, ok1 := mp[val].(string)
			if !ok1 {
				mp[val] = template.HTML(fmt.Sprintf("key %v must be string, to be converted to template.HTML", val))
			} else {
				mp[val] = template.HTML(cnv)
			}
		}
	}

	//
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
