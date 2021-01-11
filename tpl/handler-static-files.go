package tpl

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

// contained checks if a path breaks of the current dir;
// i.e.
//    ./content/site-x/../../../etc/passwd
// does break out
func contained(pth string) bool {

	if strings.HasSuffix(pth, "README.md") {
		return true
	}

	cntr := 0
	dirs := strings.Split(pth, "/")

	for _, dir := range dirs {
		if dir == "." || dir == "" {
			continue
		}
		if dir == ".." {
			cntr--
		} else {
			cntr++
		}
		if cntr < 0 {
			return false
		}
	}
	return true

}

// StaticDownloadH serves static files.
// It guesses the Content-Type header.
// It writes a Content-Length header.
// It serves the file chunk-wise without
// consuming only a buffer of memory.
func StaticDownloadH(w http.ResponseWriter, r *http.Request) {
	pth := r.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref())
	pth = strings.Trim(pth, "/")
	if !contained(pth) {
		s := fmt.Sprintf("no breaking out from doc dir: %v", pth)
		log.Print(s)
		fmt.Fprint(w, s)
		return
	}

	m := mime.TypeByExtension(filepath.Ext(pth))
	if m != "" {
		w.Header().Set("Content-Type", m)
		// log.Printf("Found Content-Type %-22v for %v", m, pth)
	}
	fpth := filepath.Join(".", "static", pth) // this enforces only local files

	/* #nosec */
	f, err := os.Open(fpth)
	if err != nil {
		s := fmt.Sprintf("StaticDownloadH: Could not open %v: %v", fpth, err)
		log.Printf(s)
		fmt.Fprint(w, s)
		w.WriteHeader(http.StatusNotFound) // otherwise - browser CSS files are retried eternally
		return
	}
	defer f.Close()

	fInfo, err := f.Stat()
	if err != nil {
		s := fmt.Sprintf("StaticDownloadH: Could not get fInfo of %v: %v", fpth, err)
		log.Printf(s)
		fmt.Fprint(w, s)
		w.WriteHeader(http.StatusNotFound) // otherwise - browser CSS files are retried eternally
		return
	}
	contentLength := fInfo.Size()
	w.Header().Set("Content-Length", fmt.Sprintf("%v", contentLength))

	// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
	w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*120))

	_, err = io.Copy(w, f) // most memory efficient
	if err != nil {
		s := fmt.Sprintf("StaticDownloadH: Could not copy file stream into response writer %v: %v", fpth, err)
		log.Printf(s)
		fmt.Fprint(w, s)
		return
	}
}
