package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/zew/questionaire/cfg"
	"github.com/zew/questionaire/sessx"
	"github.com/zew/questionaire/tpl"
)

// Template Data
type TplDataT struct {
	TemplateName string
	HtmlTitle    string
	Cnt          interface{}
	P            tPage
}

func staticDownloadH(w http.ResponseWriter, r *http.Request) {
	internalSubDir := r.URL.Path
	internalSubDir = strings.TrimPrefix(internalSubDir, cfg.Val("urlPrefix"))
	bts, _ := ioutil.ReadFile("./static" + internalSubDir)
	w.Write(bts)
}

func serveCss(w http.ResponseWriter, r *http.Request) {
	base := filepath.Base(r.URL.Path) //  "/css/design.css"  => design.css
	t := template.New(base)
	t = t.Funcs(tpl.StaticFuncMap())
	var err error
	t, err = t.ParseFiles(filepath.Join(".", "templates", base))
	if err != nil {
		log.Printf("error parsing template %v: %v", base, err)
	}
	w.Header().Set("Content-Type", "text/css")
	err = t.Execute(w, cfg.Get().Css)
	if err != nil {
		log.Printf("error executing template %v: %v", base, err)
	}
}

func sessionPut(w http.ResponseWriter, r *http.Request) {
	sess := sessx.New(w, r, sessionManager)
	sess.PutString("session-test-key", "session-test-value")
	w.Write([]byte("session[session-test-key] set to session-test-value"))
}

func sessionGet(w http.ResponseWriter, r *http.Request) {
	sess := sessx.New(w, r, sessionManager)
	val1 := sess.EffectiveParam("session-test-key")
	cnt1 := fmt.Sprintf("session-test-key is %v<br>\n", val1)
	w.Write([]byte(cnt1))
	val2 := sess.EffectiveParam("request-test-key")
	cnt2 := fmt.Sprintf("request-test-key is %v<br>\n", val2)
	w.Write([]byte(cnt2))
}

func mainH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	sess := sessx.New(w, r, sessionManager)
	login := sess.EffectiveParam("login")

	content := ``
	if login != "" {
		content += fmt.Sprintf("Username is %v<br>\n", login)
	}
	content = strings.Replace(content, "[urlprefix]", cfg.Val("urlPrefix"), -1)

	quest := generateExample()

	err := tpl.Get(w, r, sessionManager).Execute(
		w,
		TplDataT{
			HtmlTitle:    cfg.Get().AppName,
			Cnt:          content,
			P:            quest.Pages[0],
			TemplateName: "ct01.html",
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
