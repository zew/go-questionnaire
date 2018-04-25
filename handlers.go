package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/zew/questionaire/cfg"
	"github.com/zew/questionaire/sessx"
	"github.com/zew/questionaire/tpl"

	"github.com/pkg/errors"
)

// Template Data
type TplDataT struct {
	TemplateName string
	HtmlTitle    string
	CntBefore    interface{}
	CntAfter     interface{}
	Q            *tQuestionaire
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

	helper := func(err error, msg string) {
		err = errors.Wrap(err, msg)
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sess := sessx.New(w, r, sessionManager)

	var q = &tQuestionaire{}
	ok, err := sess.EffectiveParamObj("questionaire", q)
	if err != nil {
		helper(err, "Reading questionaire from session caused error")
		return
	}
	if ok {
		log.Printf("Questionaire loaded from session; %v pages", len(q.Pages))
	} else {
		q, err = LoadQuestionaire("questionaire.json")
		if err != nil {
			helper(err, "Loading questionaire from file caused error")
			return
		}
		err = q.Validate()
		if err != nil {
			helper(err, "Questionaire validation caused error")
			return
		}
		log.Printf("Questionaire loaded from file; %v pages", len(q.Pages))
	}

	//
	prevPage, ok, err := sess.EffectiveParamInt("curr_page")
	if err != nil {
		helper(err, "Reading request parameter caused error")
		return
	}
	submit := sess.EffectiveParam("submit")
	log.Printf("submit is '%v'", submit)

	currPage := q.CurrPage
	if submit == "prev" {
		currPage = q.Prev()
	}
	if submit == "next" {
		currPage = q.Next()
	}

	if err == nil && ok {
		q.CurrPage = currPage
	}

	// Todo: Parse POST request and put values into q
	if q.Pages[prevPage].Finished.IsZero() {
		q.Pages[prevPage].Finished = time.Now()
	}
	// q.Pages[1].Elements[0].Response += " aa"
	for i := 0; i < len(q.Pages[prevPage].Elements); i++ {
		key := q.Pages[prevPage].Elements[i].Name
		val := sess.EffectiveParam(key)
		log.Printf("(Page#%2v) Setting '%v' to '%v'", prevPage, key, val)
		ok, _ := sess.Exists(key)
		if ok {
			q.Pages[prevPage].Elements[i].Response = val
		}
	}

	//
	if ok := sess.EffectiveParamIsSet("lang_code"); !ok {
		sess.PutString("lang_code", q.LangCode)
	}

	//
	err = sess.PutObject("questionaire", q)
	if err != nil {
		helper(err, "Putting questionaire into session caused error")
		return
	}

	content := strings.Replace("<a href='[urlprefix]/'>Home</a>", "[urlprefix]", cfg.Val("urlPrefix"), -1)
	content = "<br>" + content + "<br>"

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.Get(w, r, sessionManager).Execute(
		w,
		TplDataT{
			HtmlTitle:    cfg.Get().AppName,
			CntBefore:    content,
			CntAfter:     content,
			TemplateName: "ct01.html",
			Q:            q,
		},
	)
	if err != nil {
		helper(err, "Executing template caused error")
		return
	}

}
