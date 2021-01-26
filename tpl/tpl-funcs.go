package tpl

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/handler"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/util"
)

// general template funcs
var fcByKey = func(k string) handler.Info {
	return handler.Infos().ByKey(k)
}
var fcURLByKey = func(k string) string {
	return cfg.Pref(handler.Infos().URLByKey(k))
}

// fcNav renders the nav core;
// it is called by the nav template via {{nav .Req .S.BookMarkUri}};
// param SBookMarkURI is app specific
var fcNav = func(r *http.Request, q *qst.QuestionnaireT, SBookMarkURI string) template.HTML {

	bts := &strings.Builder{}

	w := httptest.NewRecorder()
	isAdmin := false
	l, isLogin, err := lgn.LoggedInCheck(w, r, "admin")
	if isLogin && err == nil {
		isAdmin = true
	}

	//
	lc := ""
	if q != nil {
		if q.LangCode != "" {
			lc = q.LangCode
		} else {
			lc = q.LangCodes[0]
		}
	} else {
		sess := sessx.New(w, r)
		lcSess := sess.EffectiveStr("lang_code")
		if lcSess != "" {
			lc = lcSess
		} else {
			lc = cfg.Get().LangCodes[0]
		}
	}

	//
	langCodes := cfg.Get().LangCodes
	if q != nil {
		langCodes = q.LangCodes
	}

	//
	// nav items dynamic
	root := handler.Tree(lc)
	prev := "root"
	reqURL := cfg.TrimPrefix(r.URL.Path) + "?"
	type insertT struct {
		handler.Info
		asChild bool
	}

	// nav language
	{
		inserts := []insertT{
			{
				handler.Info{Title: cfg.Get().Mp["language"].Tr(lc), Keys: []string{"language"}},
				false,
			},
		}
		for i, lpLC := range langCodes {
			title := cfg.Get().Mp["lang_"+lpLC].Tr(lc)
			nodeURL := reqURL + "&lang_code=" + lpLC
			// log.Printf("url is '%v'", nodeURL)
			asChild := false // default is sibling
			if i == 0 {
				asChild = true
			}
			ins := insertT{
				handler.Info{Title: title, Keys: []string{"lang_" + lpLC}, Urls: []string{nodeURL}},
				asChild,
			}
			if lc == lpLC {
				ins.Urls = []string{}
				ins.Active = true
			}
			inserts = append(inserts, ins)
		}
		for i := 0; i < len(inserts); i++ {
			ok := root.AppendAfterByKey(prev, &inserts[i].Info, inserts[i].asChild)
			if !ok {
				log.Printf("appending %vth node %v after %v asChild %v failed", i, inserts[i].Title, prev, inserts[i].asChild)
				break
			}
			prev = inserts[i].Keys[0]
		}
	}

	// nav questionnaire pages
	prev = "language"

	if q != nil && len(q.Pages) > 0 {
		inserts := []insertT{
			{
				handler.Info{Title: cfg.Get().Mp["questionnaire"].Tr(lc), Keys: []string{"quest-pages"}},
				false,
			},
		}
		cntr := 0
		for i, lpP := range q.Pages {
			if lpP.NoNavigation {
				continue
			}
			cntr++
			title := lpP.Short.Tr(lc)
			nodeURL := cfg.Pref("?&page=" + fmt.Sprint(i))
			// log.Printf("url is '%v'", nodeURL)
			asChild := false // default is sibling
			if cntr == 1 {
				asChild = true
			}
			ins := insertT{
				handler.Info{Title: title, Keys: []string{nodeURL}, Urls: []string{nodeURL}},
				asChild,
			}
			if i == q.CurrPage {
				ins.Urls = []string{}
				ins.Active = true
			}
			inserts = append(inserts, ins)
		}
		for i := 0; i < len(inserts); i++ {
			ok := root.AppendAfterByKey(prev, &inserts[i].Info, inserts[i].asChild)
			if !ok {
				log.Printf("appending %vth node %v after %v asChild %v failed", i, inserts[i].Title, prev, inserts[i].asChild)
				break
			}
			prev = inserts[i].Keys[0]
		}

	}

	//
	// change logout title
	if isLogin {
		nd := root.ByKey("logout")
		if nd != nil {
			nd.Node.Title = cfg.Get().Mp["logout"].Tr(lc) + fmt.Sprintf(" (%v)", l.User) // append user name
			ok := root.SetByKey("logout", &nd.Node)
			if !ok {
				log.Printf("could not replace node 'logout'")
			}
		}
	}

	root.NavHTML(bts, r, isLogin, isAdmin, 0) // the dynamic part
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

// staticTplFuncs cannot refer to funcs with nested in-package funcs - precise reason obscure
var staticTplFuncs = template.FuncMap{
	// "toHTML":          func(arg string) template.HTML { return template.HTML("Nogo - gosec violation") },
	"toHTML":          func(arg string) template.HTML { return template.HTML(arg) },
	"formToken":       func() template.HTMLAttr { return template.HTMLAttr(lgn.FormToken()) },
	"cfg":             func() *cfg.ConfigT { return cfg.Get() }, // access to config
	"byKey":           fcByKey,                                  // template usage {{ index (   byKey "landing-page" ).Urls  0  }} - no prefix applied yet
	"urlByKey":        fcURLByKey,                               // template usage {{        urlByKey "landing-page"            }} - prefix already applied
	"nav":             fcNav,                                    // template usage {{ nav .Req  }}
	"lgn":             fcLogin,                                  // template usage {{ lgn .Req  }}
	"executeTemplate": fcExecBundledTemplate,                    // template usage {{ executeTemplate "myTpl" . }} or {{ executeTemplate .DynTemplate .}}

	//
	//
	// Type conversion stuff (static)
	"toJs":       func(arg string) template.JS { return template.JS(arg) },       // JavaScript expression
	"toJsStr":    func(arg string) template.JSStr { return template.JSStr(arg) }, // JavaScript string - *automatic quotation*
	"toURL":      func(arg string) template.URL { return template.URL(arg) },
	"toHtml":     func(arg string) template.HTML { return template.HTML(arg) },
	"toHtmlAttr": func(arg string) template.HTMLAttr { return template.HTMLAttr(arg) },
	"toCss":      func(arg string) template.CSS { return template.CSS(arg) },

	// Advanced conversions (still static)
	"toStr":    func(v interface{}) string { return fmt.Sprintf("%v", v) },
	"title":    strings.Title,
	"humanize": func(arg float64) template.HTML { return template.HTML(util.HumanizeFloat(arg)) },

	// Dynamic info about links
	"linkByKey": func(arg string) handler.Info { return handler.Infos().ByKey(arg) },
	"linksForNavigation": func() handler.InfosT {
		ret := handler.InfosT{}
		infos := []handler.Info(*handler.Infos())
		for _, l := range infos {
			ret = append(ret, l)
		}
		return ret
	},

	// Algebra
	"addInt": func(a, summand int) int {
		return a + summand
	},
	"max": func(a, b int) int {
		if a > b {
			return a
		}
		return b
	},

	// exists checks whether the struct 'data'
	// has a field or method name.
	// Usage {{if exists . "Q"}} ... {{end}}
	// 'bad' but inevitable in *general purpose* layout templates.
	// stackoverflow.com/questions/44675087/
	//
	// data might be a simple map - this is checked first
	"exists": func(data interface{}, name string) bool {
		mp, ok := data.(map[string]interface{})
		if ok {
			_, ok = mp[name]
			return ok
		}

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
