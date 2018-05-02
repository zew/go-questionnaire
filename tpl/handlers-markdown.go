package tpl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"

	"github.com/zew/questionaire/cfg"
)

// We want the markdown files editable locally with locally working links and images.
// We also want the markdown files served by the application.
// And we want - github style - a README.md served from the app root.
func CreateAndRegisterHandlerForDocs(mux1 *http.ServeMux) {

	fragm := "/doc" // significant url path fragment

	argFunc := func(w http.ResponseWriter, r *http.Request) {

		// Relay any non html and md request to static file handler
		ext := strings.ToLower(path.Ext(r.URL.Path))
		if ext != ".html" && ext != ".md" && ext != "" {
			StaticDownloadH(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		pth := r.URL.Path
		pth = strings.TrimPrefix(pth, cfg.Pref(fragm))
		pth = strings.Trim(pth, "/")

		log.Printf("doc pth %v => %v (prefix %v)", r.URL.Path, pth, cfg.Pref(fragm))

		if pth == "" {
			pth = "index.html" // quasi "index.html"
		}
		if strings.HasSuffix(pth, ".html") {
			pth = strings.Replace(pth, ".html", ".md", -1)
		}
		fpth := filepath.Join("static", fragm, pth)
		if strings.HasSuffix(fpth, "README") || strings.HasSuffix(fpth, "README.md") {
			fpth = filepath.Join(".", pth) // Readme is read directly from the app root
		}
		// log.Printf("doc fpth %v", fpth)
		bts, err := ioutil.ReadFile(fpth)
		if err != nil {
			str := fmt.Sprintf("File %v was not found.", pth)
			_, err = w.Write([]byte(str))
			if err != nil {
				log.Printf("%v", err)
			}
			return
		}

		// Rewrite URLs to be served by application
		// i.e. ./my-image.jpg to  ./[AppPrefix]/doc/
		dest := "(" + cfg.Pref() + "/doc/"
		bts = bytes.Replace(bts, []byte("(./"), []byte(dest), -1)

		// Useful for links back to application:
		// Rewrite Links with {{AppPrefix}} to application url prefix
		bts = bytes.Replace(bts, []byte("/{{AppPrefix}}"), []byte(cfg.Pref()), -1)

		bts = append(bts, []byte("\n\nRendered by russross/blackfriday")...) // Inconspicuous rendering marker

		// Render markdown
		output := blackfriday.MarkdownCommon(bts)
		_, err = w.Write(output)
		if err != nil {
			log.Printf("%v", err)
		}
	}

	log.Printf("registering docs handler %-30v -%v- %T \n", cfg.Pref(fragm), argFunc, argFunc)
	mux1.HandleFunc(cfg.Pref(fragm), argFunc)
	mux1.HandleFunc(cfg.PrefWTS(fragm), argFunc) // make sure /taxkit/doc/...  also serves argFunc, see config.Pref()

}
