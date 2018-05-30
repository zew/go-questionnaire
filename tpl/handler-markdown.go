package tpl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/sessx"
)

// CreateAndRegisterHandlerForDocs handles markdown rendering.
// We want the markdown files editable locally with locally working links and images.
// We also want the markdown files served by the application.
// And we want - github style - a README.md served from the app root.
// URL should have *.html extension, not *.md.
func CreateAndRegisterHandlerForDocs(mux1 *http.ServeMux) {

	fragm := "/doc" // significant url path fragment

	argFunc := func(w http.ResponseWriter, r *http.Request) {

		// Relay any non html and md request to static file handler
		ext := strings.ToLower(path.Ext(r.URL.Path))
		if ext != ".html" &&
			ext != ".md" &&
			!strings.HasSuffix(strings.ToLower(r.URL.Path), "readme") &&
			!strings.HasSuffix(strings.ToLower(r.URL.Path), "/doc/") &&
			!strings.HasSuffix(strings.ToLower(r.URL.Path), "doc") &&
			true {
			StaticDownloadH(w, r)
			return
		}

		// Requests ending on .md or readme are suffixed with .html and redirected
		if ext == ".md" || strings.HasSuffix(strings.ToLower(r.URL.Path), "readme") {
			newPth := r.URL.Path
			newPth = strings.TrimSuffix(newPth, path.Ext(r.URL.Path))
			newPth += ".html"
			red := newPth + url.QueryEscape(r.URL.RawQuery)
			red = newPth + "?" + r.URL.RawQuery
			log.Printf("Redirecting %20v to %20v", r.URL.Path, red)
			http.Redirect(w, r, red, http.StatusTemporaryRedirect)
			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		pth := r.URL.Path
		pth = strings.TrimPrefix(pth, cfg.Pref(fragm))
		pth = strings.Trim(pth, "/")
		if strings.Contains(pth, "../") {
			w.Write([]byte("no breaking out from doc dir"))
			return
		}
		// log.Printf("doc pth %v => %v (prefix %v)", r.URL.Path, pth, cfg.Pref(fragm))

		if pth == "" {
			pth = "index.html" // default file index.md assumed to exist in ./static/fragm
		}

		// Whereas the URL path end with .html,
		// the file path ends with *.md
		fpth := filepath.Join("static", fragm, pth)

		// Special file path: Readme is read directly from the app root
		if strings.HasSuffix(strings.ToLower(fpth), "readme.html") {
			fpth = filepath.Join(".", pth)
		}

		// Should always be the case ...
		if strings.HasSuffix(fpth, ".html") {
			fpth = strings.Replace(fpth, ".html", ".md", -1)
		}

		{
			s := fmt.Sprintf("doc path/fpath: %20v; %-20v", pth, fpth)
			log.Printf(s)
			// w.Write([]byte(s))
		}

		langCode := cfg.Get().LangCodes[0]
		sess := sessx.New(w, r)
		if ok := sess.EffectiveIsSet("lang_code"); ok {
			langCode = sess.EffectiveStr("lang_code")
		}

		fpthLangage := filepath.Join(filepath.Dir(fpth), langCode, filepath.Base(fpth))

		bts, err := ioutil.ReadFile(fpthLangage)
		if err != nil {
			if os.IsNotExist(err) {
				// bts, err = ioutil.ReadFile(strings.TrimSuffix(fpth, ".md") + ".MD")
				bts, err = ioutil.ReadFile(fpth)
			}
		}
		if err != nil {
			s := fmt.Sprintf("MarkdownH: File %v was not found.", fpth)
			log.Printf(s)
			w.Write([]byte(s))
			return
		}

		// Rewrite source file URLs to be served by application
		// i.e. ./my-image.jpg to  ./[AppPrefix]/doc/
		dest := "(" + cfg.Pref() + "/doc/"
		bts = bytes.Replace(bts, []byte("(./"), []byte(dest), -1)

		// Useful for links back to application:
		// Rewrite Links with {{AppPrefix}} to application url prefix
		bts = bytes.Replace(bts, []byte("/{{AppPrefix}}"), []byte(cfg.Pref()), -1)

		// Render markdown
		output := string(blackfriday.MarkdownCommon(bts))
		output += "<br>\n<br>\n<br>\n<p style='font-size: 75%;'>\nRendered by russross/blackfriday</p>\n" // Inconspicuous rendering marker

		tplBundle := Get(w, r, "main.html")
		ts := &StackT{"markdown.html"}
		err = tplBundle.Execute(
			w,
			TplDataT{
				TplBundle: tplBundle,
				TS:        ts,
				Sess:      &sess,
				Cnt:       output,
			},
		)
		if err != nil {
			s := fmt.Sprintf("Executing template caused: %v", err)
			log.Print(s)
			w.Write([]byte(s))
			return
		}

	}

	log.Printf("registering docs handler %-30v 'funcVar' %T \n", cfg.Pref(fragm), argFunc)
	mux1.HandleFunc(cfg.Pref(fragm), argFunc)
	mux1.HandleFunc(cfg.PrefWTS(fragm), argFunc) // make sure /taxkit/doc/...  also serves argFunc, see config.Pref()

}
