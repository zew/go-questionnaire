package tpl

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/zew/go-questionnaire/cfg"
)

// ServeDynCSS is useful mainly,
// when you have several instances of your application,
// differentiated by a few colors.
func ServeDynCSS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/css")

	cssFileName := filepath.Base(r.URL.Path) //  "/css/design.css"  => design.css
	t := GetStatic(w, r, cssFileName)
	err := t.ExecuteTemplate(w, cssFileName, cfg.Get().CSS)
	if err != nil {
		log.Printf("Error executing template %v: %v", cssFileName, err)
	}
}
