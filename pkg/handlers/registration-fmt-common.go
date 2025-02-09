package handlers

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/lgn"
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

type validatorHeadliner interface {
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

// RegistrationsFMTDownload1 returns the CSV files
func RegistrationsFMTDownload1(w http.ResponseWriter, r *http.Request) {

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

	bts, err := os.ReadFile(fp)
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

// test a connection with *short* timeout
// => preventing http response from being blocked
func isPortOpen(hostPort string, shortTO time.Duration) error {
	address := hostPort
	conn, err := net.DialTimeout("tcp", address, shortTO)
	if err != nil {
		return err
	}
	defer conn.Close()
	return nil
}

func isPrivateIP(ip net.IP) bool {
	privateRanges := []net.IPNet{
		{IP: net.IPv4(10, 0, 0, 0), Mask: net.CIDRMask(8, 32)},
		{IP: net.IPv4(172, 16, 0, 0), Mask: net.CIDRMask(12, 32)},
		{IP: net.IPv4(192, 168, 0, 0), Mask: net.CIDRMask(16, 32)},
		// zew specific
		{IP: net.IPv4(193, 196, 0, 0), Mask: net.CIDRMask(16, 32)},
	}

	for _, r := range privateRanges {
		if r.Contains(ip) {
			return true
		}
	}
	return false
}

func RegistrationsFMTDownload2(w http.ResponseWriter, r *http.Request) {

	remoteIP, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	ip := net.ParseIP(remoteIP)
	if ip == nil || !isPrivateIP(ip) {
		http.Error(w, "Access denied", http.StatusForbidden)
		log.Printf("IP %v blocked from accessing registration CSV", remoteIP)
		return
	}

	fn := "registration-fmt-de.csv"
	fd, _ := mustDir(fn)
	pth := filepath.Join(fd, fn)

	// open f
	f, err := os.Open(pth)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer f.Close()

	// file info to set headers
	stat, err := f.Stat()
	if err != nil {
		http.Error(w, "Could not get file info", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream") // or use "text/plain"
	w.Header().Set("Content-Length", fmt.Sprintf("%d", stat.Size()))
	w.Header().Set("Content-Disposition", "attachment; filename="+stat.Name())

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, "Failed to send file", http.StatusInternalServerError)
	}

	// fmt.Fprintf(w, "Welcome Internal User! Your IP: %s\n", remoteIP)

}
