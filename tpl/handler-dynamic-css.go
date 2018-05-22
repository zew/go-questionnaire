package tpl

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zew/go-questionaire/cfg"
)

var baseCSS *template.Template // The parsed css templates (base), to clone from;

// Serving a dynamic CSS is useful mainly,
// when you have several instances of your application,
// differentiated by a few colors.
func ServeDynCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/css")
	cssFileName := filepath.Base(r.URL.Path) //  "/css/design.css"  => design.css
	t := GetStatic(w, r, cssFileName)
	err := t.ExecuteTemplate(w, cssFileName, cfg.Get().Css)
	if err != nil {
		log.Printf("Error executing template %v: %v", cssFileName, err)
	}
}
