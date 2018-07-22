// Package bootstrap provides identical initialization
// for main(), but also to various integration tests
package bootstrap

import (
	"io/ioutil"
	"log"
	"path/filepath"

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

	//
	//
	// Create an empty site.css if it does not exist
	pth := filepath.Join(".", "templates", "site.css")
	if ok, _ := util.FileExists(pth); !ok {
		err := ioutil.WriteFile(pth, []byte{}, 0755)
		if err != nil {
			log.Fatalf("Could not create %v: %v", pth, err)
		}
		log.Printf("done creating file %v", pth)
	}

	tpls := []string{"main.html", "design.css", "site.css", "mobile.html", "mobile.css"}
	tpl.Parse(tpls...)

}
