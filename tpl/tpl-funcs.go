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
	// langCodes := cfg.Get().LangCodes
	// if q != nil {
	// 	langCodes = q.LangCodesOrder
	// }

	//
	root := handler.Tree()

	// append user name to logout link
	if isLogin {
		nd := root.ByKey("logout")
		if nd != nil {
			nd.Node.Title += fmt.Sprintf(" (%v)", l.User)
			ok := root.SetByKey("logout", &nd.Node)
			if !ok {
				log.Printf("could not replace node 'logout'")
			}
		}
	}

	// localize imprint
	nd := root.ByKey("imprint")
	if nd != nil {
		nd.Node.Title = cfg.Get().Mp["imprint"][lc]
		ok := root.SetByKey("imprint", &nd.Node)
		if !ok {
			log.Printf("could not replace node 'imprint'")
		}
	}

	prev := "loginlogout"
	prev = "root"

	type insert struct {
		handler.Info
		asChild bool
	}

	url := r.URL.Path + "?"

	inserts := []insert{
		{
			handler.Info{Title: "Language", Keys: []string{"language"}},
			false,
		},
		{
			handler.Info{Title: "English", Keys: []string{"english"}, Urls: []string{url + "&lang_code=en"}},
			true,
		},
		{
			handler.Info{Title: "Deutsch", Keys: []string{"deutsch"}, Urls: []string{url + "&lang_code=de"}},
			false,
		},
	}
	for i := 0; i < len(inserts); i++ {
		ok := root.AppendAfterByKey(prev, &inserts[i].Info, inserts[i].asChild)
		if !ok {
			log.Printf("appending %vth node %v after %v asChild %v failed", i, inserts[i].Title, prev, inserts[i].asChild)
			break
		}
		prev = inserts[i].Keys[0]
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
