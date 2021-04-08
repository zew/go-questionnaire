package handlers

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
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

func mustDir(fn string) string {
	// dir := cfg.Pref(filepath.Join("static", "registrations"))
	dir := filepath.Join(".", "static", "fmt-registrations")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		log.Panicf("Could not create dir %v; %v", dir, err)
	}

	fp := filepath.Join(dir, fn)

	fInfo, err := os.Stat(fp)
	_ = fInfo
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
	}

	return dir
}
