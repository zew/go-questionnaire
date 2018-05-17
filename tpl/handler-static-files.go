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

	"github.com/zew/go-questionaire/cfg"
)

func StaticDownloadH(w http.ResponseWriter, r *http.Request) {
	pth := r.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref())
	pth = strings.Trim(pth, "/")
	m := mime.TypeByExtension(filepath.Ext(pth))
	if m != "" {
		w.Header().Set("Content-Type", m)
	}
	fpth := filepath.Join(".", "static", pth)
	// bts, _ := ioutil.ReadFile(fpth)
	// w.Write(bts)
	f, err := os.Open(fpth)
	if err != nil {
		s := fmt.Sprintf("StaticDownloadH: Could not open %v: %v", fpth, err)
		log.Printf(s)
		w.Write([]byte(s))
		return
	}
	defer f.Close()

	fInfo, err := f.Stat()
	if err != nil {
		s := fmt.Sprintf("StaticDownloadH: Could not get fInfo of %v: %v", fpth, err)
		log.Printf(s)
		w.Write([]byte(s))
		return
	}
	contentLength := fInfo.Size()
	w.Header().Set("Content-Length", fmt.Sprintf("%v", contentLength))

	_, err = io.Copy(w, f) // most memory efficient
	if err != nil {
		s := fmt.Sprintf("StaticDownloadH: Could not copy file stream into response writer %v: %v", fpth, err)
		log.Printf(s)
		w.Write([]byte(s))
		return
	}
}