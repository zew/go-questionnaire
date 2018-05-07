package tpl

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/zew/go-questionaire/cfg"
)

var baseCSS *template.Template // The parsed css templates (base), to clone from

// Serving a dynamic CSS is useful mainly,
// when you have several instances of your application,
// differentiated by a few colors.
func ServeDynCss(w http.ResponseWriter, r *http.Request) {

	if !cfg.Get().IsProduction || baseCSS == nil {
		var err error
		baseCSS, err = baseCSS.ParseGlob("./templates/*.css")
		if err != nil {
			log.Fatal(err)
		}
	}

	t, err := baseCSS.Clone()
	if err != nil {
		log.Fatal(err)
	}

	t = t.Funcs(StaticFuncMap())

	cssFileName := filepath.Base(r.URL.Path) //  "/css/design.css"  => design.css

	w.Header().Set("Content-Type", "text/css")
	err = t.ExecuteTemplate(w, cssFileName, cfg.Get().Css)
	if err != nil {
		log.Printf("Error executing template %v: %v", cssFileName, err)
	}
}
