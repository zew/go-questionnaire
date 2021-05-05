package tpl

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/zew/go-questionnaire/cfg"
)

/*ServeDynCSS can serve CSS files for different sites;
the url path specifies the key to a CSSVarsSite entry;
i.e.    /css/site-1/design.css

Currently all CSS vars are set in the main template layout.html;
therefore CSS files can be aggressively cached.

Access from CSS would be
	{{ cfg.CSSVarsSite.site-1.HTML }}
	{{  (.ByKey "sec-drk2" ).RGBA    }}

Thus currently we dont need to serve CSS files as golang templates,
but it costs nothing since templates are preparsed at application init,
and we retain the possibility to use templating dynamics in future.

*/
func ServeDynCSS(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/css")
	// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
	w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*120))

	dir := path.Dir(r.URL.Path) //  /css/site-1/design.css  => /css/site-1/
	siteName := path.Base(dir)  //  /css/site-1/            => site-1

	cssFileName := path.Base(r.URL.Path) //  /css/site-1/design.css  => design.css
	t, err := Get(cssFileName)
	if err != nil {
		log.Printf("Error retrieving CSS template %v site %q: %v", cssFileName, siteName, err)
		log.Printf("\t CSS-referrer %v", r.Referer())
		return
	}

	effectiveSiteVars, ok := cfg.Get().CSSVarsSite[siteName]
	if !ok {
		if siteName == cfg.Get().AppMnemonic {
			// markdown - without any survey being set
		} else {
			log.Printf("CSS template: %v site %q does not exist in cfg.CSSVarsSite", cssFileName, siteName)
		}
		effectiveSiteVars = cfg.Get().CSSVars // defaults
	}

	err = t.ExecuteTemplate(w, cssFileName, effectiveSiteVars)
	if err != nil {
		log.Printf("Error executing CSS template %v site %q: %v", cssFileName, siteName, err)
	}

	// log.Printf("Success executing CSS template %v site %q", cssFileName, siteName)
}
