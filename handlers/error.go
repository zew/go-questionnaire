// Package handlers contains handler funcs; which are
// not stored in the main package,
// because the systemtests must access these handler funcs.
package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zew/go-questionnaire/qst"

	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/tpl"
)

func errorH(w http.ResponseWriter, r *http.Request, msg string) {

	sess := sessx.New(w, r)

	shorter := msg
	if len(shorter) > 100 {
		shorter = shorter[:100]
	}
	log.Print(shorter)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tplBundle := tpl.Get(w, r, "main-desktop.html")

	ts := &tpl.StackT{"error.html", "non-existent.html"}

	d := tplDataExtT{
		Q: &qst.QuestionnaireT{LangCode: "en"}, // just setting the lang code for the outer layout template
	}

	d.TplDataT = tpl.TplDataT{
		TplBundle: tplBundle,
		TS:        ts,
		Sess:      &sess,
		Cnt:       msg,
	}

	err := tplBundle.Execute(w, d)
	if err != nil {
		s := fmt.Sprintf("Executing template caused: %v", err)
		log.Print(s)
		w.Write([]byte(s))
		return
	}
}
