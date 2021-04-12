package handlers

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
)

var mtxFMT = sync.Mutex{}

// yearValid - either empty or within 1930 and 2050
func yearValid(s string) bool {
	if s == "" {
		return true // no number is ok
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false // not a number => not ok
	}
	if i < 1930 {
		return false
	}
	if i > 2050 {
		return false
	}
	return true
}

type ValidatorHeadliner interface {
	Validate() (map[string]string, bool)
	Headline() string
}

func mustDir(fn string) (string, int64) {
	// dir := cfg.Pref(filepath.Join("static", "registrations"))
	dir := filepath.Join(".", "static", "fmt-registrations")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Panicf("Could not create dir %v; %v", dir, err)
	}

	fp := filepath.Join(dir, fn)

	fSize := int64(0)
	fInfo, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("File %v not found - create it", fp)
			f, err := os.Create(fp)
			if err != nil {
				log.Panicf("could not create file %v; %v", fp, err)
			} else {
				log.Printf("File %v created", fp)
				f.Close()
			}
		}
	} else {
		fSize = fInfo.Size()
	}

	return dir, fSize
}

// RegistrationsFMTDownload returns the CSV files
func RegistrationsFMTDownload(w http.ResponseWriter, r *http.Request) {

	if cfg.Get().IsProduction {
		l, isLoggedIn, err := lgn.LoggedInCheck(w, r)
		if err != nil {
			fmt.Fprintf(w, "Login error %v\n", err)
			return
		}
		if !isLoggedIn {
			fmt.Fprintf(w, "Not logged in\n")
			return
		}
		if !l.HasRole("fmt-registration-csv-download") {
			fmt.Fprintf(w, "Login found, but must have role 'fmt-registration-csv-download' - and no init password\n")
			return
		}
	}

	lang := r.URL.Query().Get("lang")
	if lang != "de" && lang != "en" {
		fmt.Fprintf(w, "Append either ?lang='de' or ?lang='en' to select the language \n")
		return
	}

	fn := fmt.Sprintf("registration-fmt-%v.csv", lang)
	fd, _ := mustDir(fn)
	fp := filepath.Join(fd, fn)
	atfn := fmt.Sprintf("attachment; filename=%v", fn)

	bts, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Fprintf(w, "Could not read content of %v; %v\n", fp, err)
		return
	}

	w.Header().Set("Content-type", "application/octet-stream")
	w.Header().Set("Content-Disposition", atfn)
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	bRdr := bytes.NewReader(bts)
	io.Copy(w, bRdr)

}
