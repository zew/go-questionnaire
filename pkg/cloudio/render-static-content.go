package cloudio

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/russross/blackfriday/v2"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// RenderStaticContent writes the content of subPth into w;
// *.md files are rendered to HTML; *.html files only get URLs rewriting;
// static files reside in ./app-bucket/content;
// files may be differentiated by /[site]/[lang]/subPth
// subPth is a partial path plus filename
func RenderStaticContent(w io.Writer, subPth, site, lang, packageDocPrefix string) error {

	var (
		bts []byte
		err error
	)

	// special file path: README.md is read directly from the app root via classic ioutil
	if strings.HasSuffix(subPth, "README.md") {
		bts, err = os.ReadFile("./README.md")
		if err != nil {
			s := fmt.Sprintf("MarkdownH: cannot open README.md in app root: %v", err)
			log.Print(s)
			return fmt.Errorf(s+" %w", err)
		}
		// rewrite links in README.MD from app root
		//    ./app-bucket/content/somedir/my-img.png
		// to
		//          /urlprefix/doc/somedir/my-img.png
		//                    /doc/somedir/my-img.png  (without prefix)
		{
			needle := []byte("./app-bucket/content/")
			subst := []byte(cfg.PrefTS(packageDocPrefix))
			bts = bytes.Replace(bts, needle, subst, -1)
		}

	} else {

		pths := []string{
			path.Join(".", "content", site, lang, subPth),
			path.Join(".", "content", site, subPth),
			path.Join(".", "content", subPth),
		}

		var lpErr error
		for _, pth := range pths {
			bts, lpErr = ReadFile(pth)
			if lpErr == nil {
				lenRaw := float64(len(bts)) / 1024
				log.Printf("MarkdownH: found %v - size %4.3f kB - subPath %s", pth, lenRaw, subPth)
				break
			}
			if errors.Is(lpErr, os.ErrNotExist) {
				continue
			}
		}
		if lpErr != nil {
			errDecorated := fmt.Errorf("MarkdownH: cannot open markdown \n\t%w  \n\t%v", lpErr, pths)
			log.Print(errDecorated)
			return errDecorated
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
		// log.Printf("  bts repl README:      %2.4f kB", float32(len(bts))/1024)

	}

	// rewrite Links from static content to back application:
	//     {{AppPrefix}}
	// to
	//     /urlprefix/
	bts = bytes.Replace(bts, []byte("/{{AppPrefix}}"), []byte(cfg.Pref()), -1)

	fmt.Fprint(w, "\n\t<div class='markdown'>\n")

	ext := path.Ext(subPth)
	w1 := &strings.Builder{}
	if ext == ".html" {
		// no conversion
	} else {
		// since blackfriday version 1.52,
		// 	conversion only works for UNIX line breaks
		if false {
			bts = bytes.ReplaceAll(bts, []byte("\r\n"), []byte("\n"))
		}

		// log.Printf("  markdown:  %2.4f kB", float32(len(bts))/1024)
		bts = blackfriday.Run(bts) // render markdown
		// log.Printf("  html:      %2.4f kB", float32(len(bts))/1024)
	}
	fmt.Fprint(w1, string(bts))

	hp := trl.HyphenizeHTML(w1.String())

	fmt.Fprint(w, hp)
	fmt.Fprintf(w, "\n\t</div>  <!-- markdown  %2.4f kB -->\n", float32(len(hp))/1024)

	// output += "<br>\n<br>\n<br>\n<p style='font-size: 75%;'>\nRendered by russross/blackfriday</p>\n" // Inconspicuous rendering marker

	return nil

}
