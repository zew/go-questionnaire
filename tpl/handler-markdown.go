package tpl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/sessx"
)

// MarkDownFromFile handles markdown rendering.
func MarkDownFromFile(fpth, langCode string) (string, error) {

	fpthLangage := filepath.Join(filepath.Dir(fpth), langCode, filepath.Base(fpth))

	bts, err := ioutil.ReadFile(fpthLangage)
	if err != nil {
		if os.IsNotExist(err) {
			// bts, err = ioutil.ReadFile(strings.TrimSuffix(fpth, ".md") + ".MD")
			bts, err = ioutil.ReadFile(fpth)
		}
	}
	if err != nil {
		s := fmt.Sprintf("MarkdownH: Found neither %v nor %v.", fpthLangage, fpth)
		log.Printf(s)
		return "", fmt.Errorf(s)
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
	// output += "<br>\n<br>\n<br>\n<p style='font-size: 75%;'>\nRendered by russross/blackfriday</p>\n" // Inconspicuous rendering marker

	return output, nil

}

type staticPrefixT string // significant url path fragment

// ServeHTTP serves everything under the file directory fragm.
// We want      the markdown files editable locally with locally working links and images.
// We also want the markdown files served by the application.
//
// And we want - github style - a README.md served from the app root.
// URL should have *.html extension, not *.md.
func (fragm *staticPrefixT) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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
		// red := newPth + url.QueryEscape(r.URL.RawQuery)
		red := newPth + "?" + r.URL.RawQuery
		log.Printf("Redirecting %20v to %20v", r.URL.Path, red)
		http.Redirect(w, r, red, http.StatusTemporaryRedirect)
		return
	}

	pth := r.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref(string(*fragm)))
	pth = strings.Trim(pth, "/")
	if strings.Contains(pth, "../") {
		s := fmt.Sprintf("no breaking out from doc dir: %v", pth)
		log.Print(s)
		fmt.Fprint(w, s)
		return
	}

	if pth == "" {
		pth = "index.html" // default file index.md assumed to exist in ./static/fragm
	}

	// the URL  path ends with .html,
	// the file path ends with *.md
	fpth := filepath.Join("static", string(*fragm), pth)

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

	output, err := MarkDownFromFile(fpth, langCode)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

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
		fmt.Fprint(w, s)
		return
	}

}

// CreateAndRegisterHandlerForDocs makes everything under /static/doc/
// serveable as markdown.
func CreateAndRegisterHandlerForDocs(mux1 *http.ServeMux) {

	fragm := "/doc/"
	fragmH := staticPrefixT(fragm)
	mux1.Handle(cfg.Pref(fragm), &fragmH)
	mux1.Handle(cfg.PrefWTS(fragm), &fragmH) // make sure /taxkit/doc/...  also serves &fragmH, see config.Pref()
	log.Printf("registering docs handler %-30v 'funcVar' %T \n", cfg.Pref(fragm), &fragmH)

}
