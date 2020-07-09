package tpl

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/pkg/errors"
	"github.com/russross/blackfriday"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
)

var sfx = strings.HasSuffix // alias for function
var pfx = strings.HasPrefix

type staticPrefixT string // significant url path fragment

var packageDocPrefix = staticPrefixT("/doc/") // application singleton

// MarkDownFromFile handles markdown rendering;
// files reside inside ./app-bucket/content;
// files are further specified by /[site]/[lang]/subPth
// subPth is a partial path plus filename
func MarkDownFromFile(subPth, site, lang string) (string, error) {

	var (
		bts []byte
		err error
	)

	// special file path: README.md is read directly from the app root via classic ioutil
	if strings.HasSuffix(subPth, "README.md") {
		bts, err = ioutil.ReadFile("./README.md")
		if err != nil {
			s := fmt.Sprintf("MarkdownH: cannot open README.md in app root: %v", err)
			log.Printf(s)
			return "", errors.Wrap(err, s)
		}
		// rewrite links in README.MD from app root
		//    ./app-bucket/content/somedir/my-img.png
		// to
		//          /urlprefix/doc/somedir/my-img.png
		//                    /doc/somedir/my-img.png  (without prefix)
		{
			needle := []byte("./app-bucket/content/")
			subst := []byte(cfg.PrefTS(string(packageDocPrefix)))
			bts = bytes.Replace(bts, needle, subst, -1)
		}

	} else {
		pth := path.Join(".", "content", site, lang, subPth) // not filepath; cloudio always has forward slash
		bts, err = cloudio.ReadFile(pth)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				bts, err = cloudio.ReadFile(path.Join(".", "content", site, subPth))
				if errors.Is(err, os.ErrNotExist) {
					bts, err = cloudio.ReadFile(path.Join(".", "content", subPth))
				}
			}
		}
		if err != nil {
			s := fmt.Sprintf("MarkdownH: cannot open markdown %v or upwards: %v", pth, err)
			log.Printf(s)
			return "", errors.Wrap(err, s)
		}

		{
			// static and dynamic link back
			needle1 := []byte("(./../../../../../../README.md")
			needle2 := []byte("(./../../../../../README.md")
			needle3 := []byte("(./../../../../README.md")
			subst := []byte("(" + cfg.Pref("/doc/README.md"))
			bts = bytes.Replace(bts, needle1, subst, -1)
			bts = bytes.Replace(bts, needle2, subst, -1)
			bts = bytes.Replace(bts, needle3, subst, -1)
		}

		{
			// relative links between static files dont work, if browser url has no trailing slash;
			// rewrite
			//                   ./linux-instructions.md
			// to
			//     ./urlprefix/doc/linux-instructions.md
			needle := []byte("(./")
			subst := []byte("(" + cfg.PrefTS("/doc/"))
			bts = bytes.Replace(bts, needle, subst, -1)
		}

	}

	// rewrite Links from static content to back application:
	//     {{AppPrefix}}
	// to
	//     /urlprefix/
	bts = bytes.Replace(bts, []byte("/{{AppPrefix}}"), []byte(cfg.Pref()), -1)

	// Since blackfriday version 1.52, bullet lists only work for UNIX line breaks
	// bts = bytes.ReplaceAll(bts, []byte("\r\n"), []byte("\n"))

	// Render markdown
	output := string(blackfriday.MarkdownCommon(bts))
	output = "<div class='markdown'>" + output + "</div>"

	// output += "<br>\n<br>\n<br>\n<p style='font-size: 75%;'>\nRendered by russross/blackfriday</p>\n" // Inconspicuous rendering marker

	return output, nil

}

// ServeHTTP serves everything under the file directory fragm (for instance /doc/).
// We want the markdown files editable locally with locally working links and images.
// We want the markdown files served by the application.
// We want the markdown files served at github.com and git.zew.de.
//
// We want README.md served from the app root.
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
	siteName := "site1"
	if q, ok, _ := qst.FromSession(w, r); ok {
		siteName = q.Survey.Type
	}

	if isMarkdown {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		output, err := MarkDownFromFile(pth, siteName, langCode)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		ExecContent(w, r, output, "main-desktop.html")
	} else {
		m := mime.TypeByExtension(ext)
		if m != "" {
			w.Header().Set("Content-Type", m)
		}
		// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
		w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*120))
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
			log.Printf(s)
			return
		}
		fmt.Fprint(w, string(bts))
	}

}

// NewDocServer maps docPrefix to ./app-bucket/content;
// for instance
//             /doc/
//   /urlprefix/doc/
// serves files from
//  ./app-bucket/content
//
// Markdown files are converted to HTML;
// needs session to differentiate files by language setting
func NewDocServer(docPrefix string) {

	if !strings.HasPrefix(docPrefix, "/") {
		docPrefix = "/" + docPrefix
	}
	if !strings.HasSuffix(docPrefix, "/") {
		docPrefix = docPrefix + "/"
	}

	packageDocPrefix = staticPrefixT(docPrefix)

	// mux.Handle(cfg.Pref(docPrefix), &packageDocPrefix) // now via ServeMarkdown()
	// mux.Handle(cfg.PrefTS(docPrefix), &packageDocPrefix) // now via ServeMarkdown()

}

// ServeDoc serves markdown and other content in app-prefix/doc/
func ServeDoc(w http.ResponseWriter, r *http.Request) {
	packageDocPrefix.ServeHTTP(w, r)
}
