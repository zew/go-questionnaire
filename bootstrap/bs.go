// Package bootstrap provides identical initialization
// for main(), but also to various integration tests
package bootstrap

import (
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/util"
)

// Loading configuration and logins according to flags or env vars.
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

}
