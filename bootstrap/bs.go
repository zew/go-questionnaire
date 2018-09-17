// Package bootstrap provides identical initialization
// to main(), but also to system tests
package bootstrap

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/tpl"
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

	cfg.CfgPath = (*fl)[0].Val
	cfg.Load()

	lgn.LgnsPath = (*fl)[1].Val
	lgn.Load()

	tpls := []string{
		"main_desktop.html", "main_mobile.html",
		"main_desktop1.css", "main_desktop2.css",
		"main_mobile.css", "mobile_menu_without_js.css",
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

}
