package tpl

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"errors"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/sessx"
)

var sfx = strings.HasSuffix // alias for function
var pfx = strings.HasPrefix

type staticPrefixT string // significant url path fragment

var packageDocPrefix = staticPrefixT("/doc/") // application singleton

// RenderStaticContent writes the content of subPth into w;
// *.md files are rendered to HTML; *.html files only get URLs rewriting;
// static files reside in ./app-bucket/content;
// files may be differentiated by /[site]/[lang]/subPth
// subPth is a partial path plus filename
func RenderStaticContent(w io.Writer, subPth, site, lang string) error {
	return cloudio.RenderStaticContent(w, subPth, site, lang, string(packageDocPrefix))
}

// ServeHTTP serves everything under the file directory fragm (for instance /doc/).
// It is an improved http.FileServer(...).
// We want the markdown files editable locally with locally working links and images.
// We want the markdown files served by the application.
// We want the markdown files served at github.com and git.zew.de.
//
// We want README.md served from the app root.
//
// Markdown is rendered to HTML.
// Markdown and HTML get URLs rewritten
// Image files and other content is just served with automatic content-type detection
// and aggressive caching
//
// We want files separated by survey type and language.
// We link
//
//	/doc/site-imprint.md
//
// In the directory static, we will search
//
//	/doc/fmt/en/site-imprint.md
//	/doc/en/site-imprint.md
//	/doc/site-imprint.md
func (fragm *staticPrefixT) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fragTS := string(*fragm)
	frag := strings.TrimSuffix(fragTS, "/")

	lcP := strings.ToLower(r.URL.Path) // lower case path
	ext := path.Ext(lcP)               // lower case file extension

	byExtension := ext == ".html" || ext == ".md"
	pureReadme := sfx(lcP, "readme") // readme.html and readme.md and index.html are covered line above
	endsWithPath := sfx(lcP, fragTS) || sfx(lcP, frag)
	isMarkdown := (byExtension || pureReadme || endsWithPath)

	pth := r.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref(fragTS))
	pth = strings.Trim(pth, "/")
	if !contained(pth) {
		s := fmt.Sprintf("no breaking out from doc dir: %v", pth)
		log.Print(s)
		fmt.Fprint(w, s)
		return
	}
	if pth == "" {
		pth = "index.md" // default file index.md assumed to exist in ./static/fragm
	}

	// log.Printf("isMarkdown => byExtension || pureReadme || endsWithPath      %v => %v || %v || %v", isMarkdown, byExtension, pureReadme, endsWithPath)
	// log.Printf("path %q - ext %q - bucket path %q", lcP, ext, pth)

	langCode := cfg.Get().LangCodes[0]
	sess := sessx.New(w, r)
	if ok := sess.EffectiveIsSet("lang_code"); ok {
		langCode = sess.EffectiveStr("lang_code")
	}

	// site name
	siteName := cfg.Get().AppMnemonic
	if q, ok, _ := qst.FromSession(w, r); ok {
		siteName, _ = SiteCore(q.Survey.Type)
		// log.Printf("Markdown handler: derived site from questionnaire in session: %v", siteName)
	}

	if isMarkdown {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		w1 := &strings.Builder{}

		l, _, err := lgn.LoggedInCheck(w, r)
		if err != nil {
			fmt.Fprintf(w1, "login_by_hash_failed 2: %v", "LoginByHash error.")
			log.Printf("Login by hash error 2: %v", err)
		}

		fmt.Fprintf(w1, "\n")
		fmt.Fprintf(w1, "\t<script> var userID='%v';    </script>\n", l.User)
		fmt.Fprintf(w1, "\t<script> var provider='%v';  </script>\n", l.Provider)

		err = RenderStaticContent(w1, pth, siteName, langCode)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		HTMLTitle := path.Base(pth)
		HTMLTitle = strings.TrimSuffix(HTMLTitle, path.Ext(HTMLTitle))
		HTMLTitle = strings.ReplaceAll(HTMLTitle, "-", " ")
		if len(HTMLTitle) > 0 {
			HTMLTitle = strings.Title(HTMLTitle[0:1]) + HTMLTitle[1:]
		}

		langCodes := []string{langCode}
		for _, lc := range []string{"en", "de"} {
			if lc != langCode {
				langCodes = append(langCodes, lc)
			}
		}

		mp := map[string]interface{}{
			"Site":      siteName,
			"HTMLTitle": HTMLTitle,
			"Content":   w1.String(),
			"Q": &qst.QuestionnaireT{
				Survey:    qst.NewSurvey(siteName),
				LangCodes: langCodes,
			},
		}

		// Exec(w, r, mp, "layout.html", "documentation.html")
		RenderStack(r, w, []string{"layout.html", "documentation.html"}, mp)

	} else { // neither *.md nor *.html ...

		m := mime.TypeByExtension(ext)
		if m != "" {
			w.Header().Set("Content-Type", m)
		}
		// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
		w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*72))
		bts, err := cloudio.ReadFile(path.Join(".", "content", siteName, langCode, pth))
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				bts, err = cloudio.ReadFile(path.Join(".", "content", siteName, pth))
				if errors.Is(err, os.ErrNotExist) {
					bts, err = cloudio.ReadFile(path.Join(".", "content", pth))
				}
			}
		}
		if err != nil {
			s := fmt.Sprintf("DocHandler cannot open non-markdown %v or upwards: %v", path.Join(".", "content", siteName, langCode, pth), err)
			log.Print(s)
			return
		}
		fmt.Fprint(w, string(bts))
	}

}

// NewDocServer maps docPrefix to ./app-bucket/content;
// for instance
//
//	          /doc/
//	/urlprefix/doc/
//
// serves files from
//
//	./app-bucket/content
//
// Markdown files are converted to HTML;
// needs session to differentiate files by language setting
//
// the actual handler is ServeDoc() below
func NewDocServer(docPrefix string) {

	if !strings.HasPrefix(docPrefix, "/") {
		docPrefix = "/" + docPrefix
	}
	if !strings.HasSuffix(docPrefix, "/") {
		docPrefix = docPrefix + "/"
	}

	packageDocPrefix = staticPrefixT(docPrefix)
}

// ServeDoc serves markdown and other content in app-prefix/doc/
func ServeDoc(w http.ResponseWriter, r *http.Request) {
	packageDocPrefix.ServeHTTP(w, r)
}
