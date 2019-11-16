// Package bootstrap provides identical initialization
// to main(), but also to system tests;
// initialization is gocloud-enabled so that config files
// can be loaded from an app engine bucket.
package bootstrap

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/util"
)

// Config loads configuration and logins according to flags or env vars.
func Config() {

	fl := util.NewFlags()
	fl.Add(
		util.FlagT{
			Long:       "config_file",
			Short:      "cfg",
			DefaultVal: "config.json",
			Desc:       "JSON file containing config data",
		},
	)
	fl.Add(
		util.FlagT{
			Long:       "logins_file",
			Short:      "lgn",
			DefaultVal: "logins.json",
			Desc:       "JSON file containing logins data",
		},
	)
	fl.Gen()

	{
		cfg.CfgPath = fl.ByKey("cfg").Val

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
		log.Printf("Opened reader to cloud config %v", fileName)
		cfg.Load(r)

		cloudio.MarshalWriteFile(cfg.Example(), "config-example.json")
	}

	{
		lgn.LgnsPath = fl.ByKey("lgn").Val
		fileName := lgn.LgnsPath
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
		log.Printf("Opened reader to cloud config %v", fileName)
		lgn.Load(r)

		cloudio.MarshalWriteFile(lgn.Example(), "logins-example.json")

	}

	//
	//
	tpls := []string{
		"main_desktop.html", "main_mobile.html",
		"main_desktop1.css", "main_desktop2.css",
		"main_mobile.css", "main_mobile_menu_without_js.css",
	}

	err := filepath.Walk(filepath.Join(".", "templates"), func(path string, f os.FileInfo, err error) error {
		base := filepath.Base(path)
		if strings.HasPrefix(base, "main_desktop_") ||
			strings.HasPrefix(base, "main_mobile_") {
			log.Printf("Adding %v", base)
			tpls = append(tpls, base)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walinkg templates: %v", err)
	}

	tpl.Parse(tpls...)

	sessx.Mgr().Lifetime = time.Duration(cfg.Get().SessionTimeout) * time.Hour // default is 24 hours
	// sessx.Mgr().Secure(true)            // true breaks session persistence in excel-db - but not in go-countdown - it leads to sesson breakdown on iphone safari mobile, maybe because appengine is http with TLS outside

}
