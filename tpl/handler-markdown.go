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
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
)

// MarkDownFromFile handles markdown rendering;
// directory "doc" is a hardcoded assumption.
func MarkDownFromFile(fpth, surveyType, langCode string) (string, error) {

	// bts, err = ioutil.ReadFile(strings.TrimSuffix(fpth, ".md") + ".MD")

	fpthSurveyLang := filepath.Join(filepath.Dir(fpth), surveyType, langCode, filepath.Base(fpth))
	fpthLang := filepath.Join(filepath.Dir(fpth), langCode, filepath.Base(fpth))

	bts, err := ioutil.ReadFile(fpthSurveyLang)
	if err != nil {
		if os.IsNotExist(err) {
			bts, err = ioutil.ReadFile(fpthLang)
			if os.IsNotExist(err) {
				bts, err = ioutil.ReadFile(fpth)
			}
		}
	}

	if err != nil {
		s := fmt.Sprintf("MarkdownH: Found neither %v nor %v nor %v.", fpthSurveyLang, fpthLang, fpth)
		log.Printf(s)
		return "", fmt.Errorf(s)
	}

	//
	// Rewrite source file URLs to be served by application
	//

	// readme.md from root
	// ./static/img/doc/my-img.png
	//  /survey/img/doc/my-img.png
	//         /img/doc/my-img.png  (without prefix)
	{
		needle := []byte("(./static/")
		subst := []byte("(" + cfg.PrefWTS())
		bts = bytes.Replace(bts, needle, subst, -1)
	}

	// relative urls from ./static/doc/
	{
		//     ./../../README.md
		// /survey/doc/README.md
		needle := []byte("(./../../")
		subst := []byte("(" + cfg.PrefWTS("/doc/"))
		bts = bytes.Replace(bts, needle, subst, -1)
	}
	{
		//    ./../img/doc/zew-footer.png
		// /survey/img/doc/zew-footer.png
		needle := []byte("(./../img/")
		subst := []byte("(" + cfg.PrefWTS("/img/"))
		bts = bytes.Replace(bts, needle, subst, -1)
	}
	{
		//            ./linux-instructions.md
		// ./survey/doc/linux-instructions.md
		needle := []byte("(./")
		subst := []byte("(" + cfg.PrefWTS("/doc/"))
		bts = bytes.Replace(bts, needle, subst, -1)
	}

	// Useful for links back to application:
	// Rewrite Links with {{AppPrefix}} to application url prefix
	bts = bytes.Replace(bts, []byte("/{{AppPrefix}}"), []byte(cfg.Pref()), -1)

	// Render markdown
	output := string(blackfriday.MarkdownCommon(bts))
	// output += "<br>\n<br>\n<br>\n<p style='font-size: 75%;'>\nRendered by russross/blackfriday</p>\n" // Inconspicuous rendering marker

	return output, nil

}

type staticPrefixT string // significant url path fragment

// ServeHTTP serves everything under the file directory fragm (for instance /doc/).
// We want      the markdown files editable locally with locally working links and images.
// We also want the markdown files served by the application.
//
// And we want - github style - a README.md served from the app root.
// URL should have *.html extension, not *.md.
//
//
// We want files separated by survey type and language.
// We link
//     /doc/site-imprint.md
// In the directory static, we will search
//     /doc/fmt/en/site-imprint.md
//     /doc/en/site-imprint.md
//     /doc/site-imprint.md
func (fragm *staticPrefixT) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fragWTS := string(*fragm)
	frag := strings.TrimSuffix(fragWTS, "/")

	lcP := strings.ToLower(r.URL.Path) // lower case path

	ext := path.Ext(r.URL.Path) // file extension
	ext = strings.ToLower(ext)

	sfx := strings.HasSuffix // alias for function
	pfx := strings.HasPrefix

	startsWithPath := pfx(lcP, cfg.Pref(fragWTS))

	byExtension := ext == ".html" || ext == ".md"
	pureReadme := sfx(lcP, "readme") // readme.html and readme.md and index.html are covered line above
	endsWithPath := sfx(lcP, fragWTS) || sfx(lcP, frag)

	isMarkup := startsWithPath && (byExtension || pureReadme || endsWithPath)

	// log.Printf("path %q  ext %q", lcP, ext)
	// log.Printf("isMarkup %v => (startsWithPath && (byExtension || pureReadme || endsWithPath)  =>  %v && (%v || %v || %v))", isMarkup, startsWithPath, byExtension, pureReadme, endsWithPath)

	// Relay any non html or non md request to static file handler
	if !isMarkup {
		log.Printf("Markdown handler handing off to staticDownload(): \n%v", lcP)
		StaticDownloadH(w, r)
		return
	}

	// log.Printf("Markdown handler dealing with: \n%v", lcP)

	// Requests ending on .md or readme are suffixed with .html and redirected
	if ext == ".md" || strings.HasSuffix(strings.ToLower(r.URL.Path), "readme") {
		newPth := r.URL.Path
		newPth = strings.TrimSuffix(newPth, path.Ext(r.URL.Path))
		newPth += ".html"
		// red := newPth + url.QueryEscape(r.URL.RawQuery)
		red := newPth + "?" + r.URL.RawQuery
		log.Printf("Redirecting %20v to %20v", r.URL.Path, red)
		http.Redirect(w, r, red, http.StatusSeeOther)
		return
	}

	pth := r.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref(fragWTS))
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
	fpth := filepath.Join("static", fragWTS, pth)

	// Special file path: Readme is read directly from the app root
	if strings.HasSuffix(strings.ToLower(fpth), "readme.html") {
		fpth = filepath.Join(".", pth)
	}

	// Should always be the case ...
	if strings.HasSuffix(fpth, ".html") {
		fpth = strings.Replace(fpth, ".html", ".md", -1)
	}

	s := fmt.Sprintf("doc path/fpath: %20v; %-20v", pth, fpth)
	log.Printf(s)
	// w.Write([]byte(s))

	langCode := cfg.Get().LangCodes[0]
	sess := sessx.New(w, r)
	if ok := sess.EffectiveIsSet("lang_code"); ok {
		langCode = sess.EffectiveStr("lang_code")
	}

	// survey name
	var q = &qst.QuestionnaireT{}

	surveyType := ""
	if ok, _ := sess.EffectiveObj("questionnaire", q); ok {
		surveyType = q.Survey.Type
	}

	output, err := MarkDownFromFile(fpth, surveyType, langCode)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tplBundle := Get(w, r, "main_desktop.html")
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
