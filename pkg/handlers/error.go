// Package handlers contains handler funcs; which are
// not stored in the main package,
// because the systemtests must access these handler funcs.
package handlers

import (
	"log"
	"net/http"
	"strings"

	"github.com/zew/go-questionnaire/pkg/qst"

	"github.com/zew/go-questionnaire/pkg/tpl"
)

func errorH(w http.ResponseWriter, r *http.Request, msg string) {

	shorter := msg
	if len(shorter) > 100 {
		shorter = shorter[:100]
	}
	log.Print(shorter)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w2 := &strings.Builder{}
	tpl.ExecContent(w2, r, msg, "error.html")

	mp := map[string]interface{}{
		"HTMLTitle": "Error page",
		"Q":         &qst.QuestionnaireT{LangCode: "en"}, // just the lang code for the outer layout template
		"Content":   w2.String(),
	}
	// tpl.Exec(w, r, mp, "layout.html")
	tpl.RenderStack(r, w, []string{"layout.html", "documentation.html"}, mp)

}
