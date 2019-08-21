package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
)

// ConfigReloadH can be everywhere but in package cfg.
// Package cfg must remain free of iocloud
var ConfigReloadH = func(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	fileName := cfg.CfgPath
	r, bucketClose, err := cloudio.Open(fileName)
	if err != nil {
		log.Fatalf("Error opening writer to %v: %v", fileName, err)
	}
	defer func() {
		err := r.Close()
		if err != nil {
			log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
		}
	}()
	defer func() {
		err := bucketClose()
		if err != nil {
			log.Printf("Error closing bucket of writer to %v: %v", fileName, err)
		}
	}()
	log.Printf("Opened reader to cloud config %v", cfg.CfgPath)
	cfg.Load(r)

	fmt.Fprintf(w, "cfg reloaded")

}
