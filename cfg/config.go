// Package cfg implements a configuration database,
// loaded from a json file.
// Filename must be given as command line argument
// or environment variable.
// Access to the config data is made in threadsafe manner.
package cfg

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"github.com/zew/util"

	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/trl"
)

// ConfigT holds the application config
type ConfigT struct {
	sync.Mutex

	IsProduction bool `json:"is_production"` // true => templates are not recompiled

	AppName       string `json:"app_name"`       // with case, i.e. 'My App'
	URLPathPrefix string `json:"urlpath_prefix"` // lower case - no slashes, i.e. 'myapp'
	AppMnemonic   string `json:"app_mnemonic"`   // For differentiation of static dirs - when URLPathPrefix is empty; imagine multiple instances

	BindHost               string `json:"bind_host"`
	BindSocket             int    `json:"bind_socket"`
	BindSocketFallbackHTTP int    `json:"bind_socket_fallback_http"`
	BindSocketTests        int    `json:"bind_socket_tests,omitempty"` // another port for running test server
	TLS                    bool   `json:"tls"`
	TLS13                  bool   `json:"tls13"`                   // ultra safe - but excludes internet explorer 11
	ReadTimeOut            int    `json:"http_read_time_out"`      // for large requests
	WriteTimeOut           int    `json:"http_write_time_out"`     // for *responding* large files over slow networks, i.e. videos, set to 30 or 60 secs
	MaxPostSize            int64  `json:"max_post_size,omitempty"` // request body size limit, against DOS attacks, limits file uploads

	LocationName   string         `json:"location,omitempty"` // i.e. "Europe/Berlin", see Go\lib\time\zoneinfo.zip
	Loc            *time.Location `json:"-"`                  // Initialized during load
	SessionTimeout int            `json:"session_timeout"`    // hours until the session is lost
	FormTimeout    int            `json:"form_timeout"`       // hours until a form post is rejected

	AllowSkipForward bool `json:"allow_skip_forward"` // skipping back always allowed, skipping forward is configurable

	CSS map[string]string `json:"css"` // differentiate multiple instances by color and stuff - without duplicating entire css files

	// available language codes for the application, first element is default
	LangCodes []string `json:"lang_codes"`
	// multi language strings for the application
	Mp trl.Map `json:"translation_map"`
}

// CfgPath is obtained by ENV variable or command line flag in main package.
// Being set from the main package.
// Holds the relative path and filename to look for; could be ".cfg/config.json".
// Relative to the app main dir.
var CfgPath = path.Join(".", "config.json")

var cfgS *ConfigT // package variable 'singleton' - needs to be an allocated struct - to hold pointer receiver-re-assignment

// Get provides access to the app configuration
func Get() *ConfigT {
	// Same as lgn.Get().
	// No lock needed here.
	// Since in load(), we simply exchange one pointer by another at the end of loading.
	// c.Lock()
	// defer c.Unlock()
	return cfgS
}

// SwitchToTestConfig is used to run systemtests on a different port without TLS.
func SwitchToTestConfig() {
	cfgS.BindSocket = cfgS.BindSocketTests
	cfgS.TLS = false // certificate not valid for localhost
	log.Printf("Testing config: Port %v, TLS %v", cfgS.BindSocket, cfgS.TLS)
}

// Load reads from a JSON file.
// To avoid concurrent access problems:
// No method to ConfigT, no pointer receiver.
// We could *copy* at the end of method  *c = *newCfg,
// but onerous.
// Instead:
// cfgS = &tempCfg
//
// Contains some validations.
func Load() {
	// c.Lock()
	// defer c.Unlock()
	file, err := util.LoadConfigFile(CfgPath)
	if err != nil {
		log.Fatalf("Could not load config file: %v", err)
	}
	log.Printf("Found config file: %v", CfgPath)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Err closing config file: %v", err)
		}
		log.Printf("Closed config file: %v", CfgPath)
	}()

	decoder := json.NewDecoder(file)
	tempCfg := ConfigT{}
	err = decoder.Decode(&tempCfg)
	if err != nil {
		log.Fatal(err)
	}

	if tempCfg.AppName == "" {
		log.Fatal("Config underspecified; at least app_name should be set")
	}
	if tempCfg.AppMnemonic == "" {
		log.Fatal("Config underspecified; at least app_mnemonic should be set")
	}
	if len(tempCfg.LangCodes) < 1 {
		log.Fatal("You must specify at least one language code such as 'en' or 'de'  in your configuration.")
	}
	trl.LangCodes = tempCfg.LangCodes // trl.LangCodes is a redundant copy of cfg.LangCodes - but keeps the packages separate

	tempCfg.Loc, err = time.LoadLocation(tempCfg.LocationName)
	if err != nil {
		log.Fatalf("Your location name must be valid, i.e. 'Europe/Berlin', compare Go\\lib\\time\\zoneinfo.zip: %v", err)
	}

	//
	cfgS = &tempCfg // replace pointer in one go - should be threadsafe
	log.Printf("config loaded 1\n%s", util.IndentedDump(cfgS))
}

// LoadH is a convenience handler - to reload the config via http request
func LoadH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// We cannot use lgn.LoggedInCheck()
	// causing circular dependency
	// Therefore we need implementing it here
	sess := sessx.New(w, r)
	type loginTypeTemp struct {
		User  string            `json:"user"`
		Roles map[string]string `json:"roles"` // i.e. admin: true , gender: female, height: 188
	}
	l := &loginTypeTemp{}
	loggedIn, err := sess.EffectiveObj("login", l)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !loggedIn {
		http.Error(w, "admin login required for this function", http.StatusInternalServerError)
		return
	}
	if _, ok := l.Roles["admin"]; !ok {
		http.Error(w, "admin login required for this function", http.StatusInternalServerError)
		return
	}

	Load()
	w.Write([]byte("cfg reloaded"))
}

// Save stores the config to a JSON file
func (c *ConfigT) Save(fn ...string) {

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(c, firstColLeftMostPrefix, "\t")
	util.BubbleUp(err)

	saveDir := path.Dir(CfgPath)
	err = os.Chmod(saveDir, 0755)
	util.BubbleUp(err)

	configFile := path.Base(CfgPath)
	if len(fn) > 0 {
		configFile = fn[0]
	}
	savePath := path.Join(saveDir, configFile)
	err = ioutil.WriteFile(savePath, byts, 0644)
	util.BubbleUp(err)

	log.Printf("Saved config file to %v", savePath)
}

// Pref prefixes a URL path with an application dir prefix.
// Any URL Path is prefixed with the URLPathPrefix, if URLPathPrefix is set.
// Prevents unnecessary slashes.
// No trailing slash
// Routes with trailing "/" such as "/path/"
// get a redirect "/path" => "/path/" if "/path" is not registered yet.
// This behavior of func server.go - (mux *ServeMux) Handle(...) is nasty
// since it depends on the ORDER of registrations.
//
// Best strategy might be
//    mux.HandleFunc(appcfg.Pref(urlPath), argFunc)      // Claim "/path"
//    mux.HandleFunc(appcfg.PrefWTS(urlPath), argFunc)   // Claim "/path/"
// Notice the order - other way around would block "/path" with a redirect handler
func Pref(pth ...string) string {

	if cfgS.URLPathPrefix != "" {
		if len(pth) > 0 {
			return path.Join("/", cfgS.URLPathPrefix, pth[0])
		}
		return path.Join("/", cfgS.URLPathPrefix)
	}

	// No URLPathPrefix
	if len(pth) > 0 {
		return path.Join("/", pth[0])
	}
	return ""

}

// PrefWTS is like Prefix(); WTS stands for with trailing slash
func PrefWTS(pth ...string) string {
	p := Pref(pth...)
	return p + "/"
}

// Example writes a minimal configuration to file, to be extended or adapted
func Example() {
	ex := &ConfigT{
		IsProduction:           true,
		AppName:                "My Example App",
		URLPathPrefix:          "exmpl",
		AppMnemonic:            "exmpl",
		BindHost:               "0.0.0.0",
		BindSocket:             8081,
		BindSocketFallbackHTTP: 8082,
		BindSocketTests:        8181,
		TLS:                    false,
		TLS13:                  false,
		ReadTimeOut:            5,
		WriteTimeOut:           30,
		MaxPostSize:            int64(2 << 20), // 2 MB
		LocationName:           "Europe/Berlin",
		SessionTimeout:         2,
		FormTimeout:            2,

		LangCodes: []string{"de", "en"},
		Mp: trl.Map{
			"page": {
				"en": "Page",
				"de": "Seite",
			},
			"app_label_h1": {
				"en": "My Org",
				"de": "Meine Organisation",
			},
			"app_label": {
				"en": "My Example App", // yes, repeat of AppName
				"de": "Meine Beispiel Anwendung",
			},

			"correct_errors": {
				"de": "Bitte korrigieren Sie die unten angezeigten Fehler.",
				"en": "Please correct the errors displayed below.",
			},
			"imprint": {
				"de": "Impressum",
				"en": "Imprint",
			},
			"login_by_hash_failed": {
				"de": "Anmeldung via Hash gescheitert.\nBitte nutzen Sie den übermittelten Link um sich anzumelden.\nWenn der Link in zwei Zeilen geteilt wurde, verbinden Sie die Zeilen wieder.",
				"en": "Login by hash failed.\nPlease use the provided link to login.\nIf the link was split into two lines, reconnect them.",
			},
			"finished_by_participant": {
				"de": "Sie haben den Fragebogen bereits abgeschlossen (%v).",
				"en": "You already finished this survey wave at %v",
			},
			"deadline_exceeded": {
				"de": "Diese Umfrage wurde am %v beendet.",
				"en": "Current survey was closed at %v.",
			},
			"percentage_answered": {
				"de": "Sie haben %v von %v Fragen beantwortet: %2.1f Prozent.  <br>\n",
				"en": "You answered %v out of %v questions: %2.1f percent.  <br>\n",
			},
			"survey_ending": {
				"de": "Umfrage endet am %v. <br>\nVeröffentlichung am %v.  <br>\n",
				"en": "Survey will finish at %v. <br>\nPublication will be at %v.<br>\n",
			},
			"review_by_personal_link": {
				"de": "Sie können ihre Daten jederzeit über Ihren persönlichen Link prüfen/ändern. <br>\n<a href='/?submitBtn=prev'>Zurück</a><br>\n",
				"en": "You may review or change your data using your personal link. <br>\n<a href='/?submitBtn=prev'>Back</a><br>\n",
			},
			"not_a_number": {
				"de": "'%v' keine Zahl",
				"en": "'%v' not a number",
			},
			"too_big": {
				"de": "Max %.0f",
				"en": "max %.0f",
			},
			"too_small": {
				"de": "Min %.0f",
				"en": "min %.0f",
			},
			"must_one_option": {
				"de": "Bitte eine Option wählen",
				"en": "Please choose one option",
			},
		},
	}
	ex.Save("config-example.json")
}
