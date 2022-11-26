// Package cfg implements a configuration database,
// loaded from a json file.
// Filename must be given as command line argument
// or environment variable.
// Access to the config data is made in threadsafe manner.
package cfg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pbberlin/dbg"

	"github.com/zew/go-questionnaire/pkg/trl"
)

// directLoginRangeT draws a connection from a survey
// to a group of user IDs.
// Hashed ids between Start and Stop are allowed to directly login.
// A login is created for the specified SurveyID/WaveID (which must exist)
// and with the attributes stored in Profile i.e. "lang_code": "en"
type directLoginRangeT struct {
	Start    int               `json:"start,omitempty"`
	Stop     int               `json:"stop,omitempty"`
	SurveyID string            `json:"survey_id,omitempty"`
	WaveID   string            `json:"wave_id,omitempty"`
	Profile  map[string]string `json:"profile,omitempty"`
}

// ConfigT holds the application config
type ConfigT struct {
	// sync.Mutex not needed since we swap the pointer

	IsProduction bool `json:"is_production"` // true => templates are not recompiled

	AppName       string `json:"app_name"`       // with case, i.e. 'Survey Server'; use localized trl.Map app_label, app_org
	URLPathPrefix string `json:"urlpath_prefix"` // lower case - no slashes, i.e. 'myapp'
	AppMnemonic   string `json:"app_mnemonic"`   // differentiation of static dirs - when URLPathPrefix is empty; imagine multiple instances

	LetsEncrypt bool   `json:"lets_encrypt"`
	HostName    string `json:"host_name"` // for ACME cert; i.e. survey2.zew.de

	BindHost               string `json:"bind_host"`                   // "0.0.0.0"
	BindSocket             int    `json:"bind_socket"`                 // 8081 or 80
	BindSocketFallbackHTTP int    `json:"bind_socket_fallback_http"`   // 8082
	BindSocketTests        int    `json:"bind_socket_tests,omitempty"` // another port for running test server, 8181
	TLS                    bool   `json:"tls"`
	TLS13                  bool   `json:"tls13"`                     // ultra safe - but excludes internet explorer 11
	ReadTimeOut            int    `json:"http_read_time_out"`        // limit large requests
	ReadHeaderTimeOut      int    `json:"http_header_read_time_out"` // limit request header time - then use per request restrictions r = r.WithContext(ctx) to limit - stackoverflow.com/questions/39946583
	WriteTimeOut           int    `json:"http_write_time_out"`       // for *responding* large files over slow networks, i.e. videos, set to 30 or 60 secs

	TimeOutUsual      int      `json:"time_out_usual,omitempty"`
	TimeOutExceptions []string `json:"time_out_exceptions,omitempty"`

	MaxPostSize int64 `json:"max_post_size,omitempty"` // request body size limit, against DOS attacks, limits file uploads

	// LocationName i.e. "Europe/Berlin", see Go\lib\time\zoneinfo.zip;
	// LocationName only serves for initializing Loc at application start
	// after that, application should use Loc
	LocationName   string         `json:"location,omitempty"`
	Loc            *time.Location `json:"-"`               // Initialized during load; seconds east of UTC
	SessionTimeout int            `json:"session_timeout"` // hours until the session is lost
	FormTimeout    int            `json:"form_timeout"`    // hours until a form post is rejected

	AppInstanceID int64    `json:"app_instance_id,omitempty"` // append to URLs of cached static jpg, js and css files - change to trigger reload
	LangCodes     []string `json:"lang_codes"`                // available language codes for the application, first element is default

	CPUProfile string `json:"cpu_profile"` // CPUProfile - output filename

	Mp     trl.Map     `json:"translations_generic"` // Mp     - multi language strings for entire application -       [key].Tr(lc)
	MpSite trl.MapSite `json:"translations_site"`    // MpSite - multi language strings for specific survey -    [site][key].Tr(lc)

	// keep this last - since it trashes diff view
	CSSVars     cssVars            `json:"css_vars"`      // global CSS variables - no localization
	CSSVarsSite map[string]cssVars `json:"css_vars_site"` // [site|Survey.Type] specific CSS - overwrites/appends global css_vars - no localization

	AnonymousSurveyID string                       `json:"anonymous_survey_id,omitempty"` // on anonymous login - name of the survey into which the login is effected; also for redirect URL to LoginByHashID handler
	Profiles          map[string]map[string]string `json:"profiles"`                      // Profiles are sets of attributes, selected by the `p` parameter at login, containing key-values which are copied into the logged in user's attributes
	DirectLoginRanges []directLoginRangeT          `json:"direct_login_ranges,omitempty"` // DirectLoginRanges - user id to language preselection for direct login

}

// CfgPath is obtained by ENV variable or command line flag in main package.
// Being set from the main package.
// Holds the relative path and filename to look for; could be "./cfg/config.json".
// Relative to the app main dir.
var CfgPath = filepath.Join(".", "config.json")

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

// Load reads from an io.Reader
// to avoid cyclical deps.
//
// To avoid concurrent access problems:
// No method to ConfigT, no pointer receiver.
// We could *copy* at the end of method  *c = *newCfg,
// but onerous.
// Instead:
// cfgS = &tempCfg
//
// Contains some validations.
func Load(r io.Reader) {

	decoder := json.NewDecoder(r)
	tempCfg := ConfigT{}
	err := decoder.Decode(&tempCfg)
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
		log.Printf("Your location name must be valid, i.e. 'Europe/Berlin', \ncompare Go\\lib\\time\\zoneinfo.zip: %v", err)
		// Under windows, we get errors; see golang.org/pkg/time/#LoadLocation
		// $GOROOT/lib/time/zoneinfo.zip
		//
		// Fallback:
		tempCfg.Loc = time.FixedZone("UTC_-2", -2*60*60)
	}

	tempCfg.AppInstanceID = time.Now().Unix()

	for key := range tempCfg.CSSVarsSite {
		tempCfg.CSSVarsSite[key] = Stack(tempCfg.CSSVars, tempCfg.CSSVarsSite[key])
		log.Printf("combined CSSVars base plus %-10s- computed; %v entries", key, len(tempCfg.CSSVarsSite[key]))
		// if key == "biii" {
		// 	dbg.Dump(tempCfg.CSSVarsSite[key])
		// }
	}

	basePlusCustom := trl.CoreTranslations()
	for key := range tempCfg.Mp {
		basePlusCustom[key] = tempCfg.Mp[key] // set anew or overwrite base
	}
	tempCfg.Mp = basePlusCustom

	//
	cfgS = &tempCfg // replace pointer in one go - should be threadsafe
	dmp := dbg.Dump2String(cfgS)
	if len(dmp) > 700 {
		dmp = dmp[:700]
	}
	{
		dmp := dbg.Dump2String(tempCfg.MpSite)
		log.Printf("\ntranslations_site: %s\n...config loaded", dmp)
	}

	log.Printf("\n%s\n...config loaded", dmp)
}

// LoadFakeConfigForTests makes bootstrap overhead superfluous
// by creating a
func LoadFakeConfigForTests() {
	adHocConfig := &ConfigT{
		AppName:     "Fake Config for Tests - App Name",
		AppMnemonic: "fake-config-for-tests-app-mnemonic",
		LangCodes:   []string{"en"},
		Loc:         time.FixedZone("UTC_-2", -2*60*60),
	}
	bts, err := json.Marshal(adHocConfig)
	if err != nil {
		log.Panicf("Could not create fake config for tests: %v", err)
	}
	Load(bytes.NewReader(bts))
}

// Pref prefixes a URL path with an application dir prefix.
// Any URL Path is prefixed with the URLPathPrefix, if URLPathPrefix is set.
//
// Prevents unnecessary slashes.
// No trailing slash
// Routes with trailing "/" such as "/path/"
// get a redirect "/path" => "/path/" if "/path" is not registered yet.
// This behavior of func server.go - (mux *ServeMux) Handle(...) is nasty
// since it depends on the ORDER of registrations.
//
// Best strategy might be
//
//	mux.HandleFunc(appcfg.Pref(urlPath),   argFunc)   // Claim "/path"
//	mux.HandleFunc(appcfg.PrefTS(urlPath), argFunc)   // Claim "/path/"
//
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

// TrimPrefix removes the prefix from a URL
func TrimPrefix(url string) (ret string) {
	ret = strings.TrimSuffix(url, "/")
	ret = strings.TrimSuffix(ret, cfgS.URLPathPrefix)
	return
}

// Pref for templates: cfg.Pref
func (c *ConfigT) Pref(pth ...string) string {
	return Pref(pth...)
}

// Tr for global translations in templates
// i.e. {{ cfg.Tr .Q.LangCode "correct_errors" }}
func (c *ConfigT) Tr(langCode, key string) string {
	return c.Mp[key].Tr(langCode)
}

// Val for site and language specific values in templates;
// function falls back to key "default";
// i.e. {{ cfg.Val .Site "en"      "app_label"}}
//
//	{{ cfg.Val .Site "default" "img_logo_icon"}}
func (c *ConfigT) Val(site, langCode, key string) string {
	// site key missing
	if _, ok := c.MpSite[site]; !ok {
		// lang key missing
		if _, ok := c.MpSite["default"][key][langCode]; !ok {
			return c.MpSite["default"][key]["default"]
		}
		return c.MpSite["default"][key].Tr(langCode)
	}
	// site ok - but lang key missing
	if _, ok := c.MpSite[site][key][langCode]; !ok {
		return c.MpSite[site][key]["default"]
	}
	return c.MpSite[site][key].Tr(langCode)
}

// PrefTS is like Prefix(); TS stands for (with) trailing slash;
// useful for registering handlers
// so that /p1/p2/  also serves /p1/p2
func PrefTS(pth ...string) string {
	p := Pref(pth...)
	return p + "/"
}

// AbsoluteLink creates a HTTP URL
func (c *ConfigT) AbsoluteLink() string {
	port := fmt.Sprintf(":%v", c.BindSocket)
	if c.BindSocket == 0 || c.BindSocket == 80 || c.BindSocket == 443 {
		port = ""
	}
	lnk := fmt.Sprintf("%v%v", c.HostName, port)

	if c.URLPathPrefix != "" {
		lnk = lnk + "/" + c.URLPathPrefix
	}

	if c.TLS {
		lnk = "https://" + lnk
	} else {
		lnk = "http://" + lnk
	}

	return lnk
}

// Example writes a minimal configuration to file, to be extended or adapted
func Example() *ConfigT {
	ex := &ConfigT{
		IsProduction:           true,
		AppName:                "My Example App Label",
		URLPathPrefix:          "exmpl",
		AppMnemonic:            "exmpl",
		HostName:               "survey2.zew.de",
		BindHost:               "0.0.0.0",
		BindSocket:             8081,
		BindSocketFallbackHTTP: 8082,
		BindSocketTests:        8181,
		TLS:                    false,
		TLS13:                  false,
		ReadTimeOut:            10,
		ReadHeaderTimeOut:      10,
		WriteTimeOut:           60,
		TimeOutUsual:           10,
		TimeOutExceptions:      []string{"transferrer-endpoint", "download/", "download-stream/"},
		MaxPostSize:            int64(2 << 20), // 2 MB
		LocationName:           "Europe/Berlin",
		SessionTimeout:         2,
		FormTimeout:            2,
		// each cfg load or reload updates this value
		// AppInstanceID:          time.Now().Unix(),
		LangCodes: []string{"de", "en", "es", "fr", "it", "pl"},
		CSSVars: cssVars{
			// {Key: "logo-text", Val: "ZEW"}, // use localized trl.Map app_label, app_org
			{IsURL: true, Key: "img-bg", Val: "/img/ui/bg-bw-bland.jpg"},
			{IsURL: true, Key: "img-loggedin-icon", Val: "/img/ui/logged-in-icon-zew.svg"},
			{Key: "nav-height", Val: "8vh"},
			{Key: "nav-rest-height", Val: "calc(100vh - var(--nav-height))", Desc: "we can calc() the remainder"},
			{Key: "nav-bar-position", Val: "relative", Desc: "fixed or relative"},
			{Key: "content-top", Val: "0", Desc: "fixed navbar => content-top = var(--nav-height); otherwise 0"},
			{Key: "bg", Colorname: "white", Desc: "main background f <body>"},
			{Key: "fg", Colorname: "black", Desc: "main foreground, font-color"},
			{Key: "input-bg", Colorname: "white", Desc: "input+select background"},
			{Key: "input-fg", Colorname: "black", Desc: "input+select foreground"},
			{Key: "inp-border", R: 136, G: 136, B: 136, Desc: "input border of radio and checkbox - most browsers"},
			{Key: "inp-focus", R: 0, G: 127, B: 255, Desc: "chrome and firefox radio focus color; not edge"},

			/* we dont want Alpha: .5 anymore
			   instead we use flexible alpha values as follows:
			   background-color:   rgba(var(--clr-pri), 0.5); */
			{Key: "has-alpha", R: 240, G: 240, B: 240, Alpha: .9, Desc: "has alpha - but takes away flexibility"},

			{Key: "err", Colorname: "darkred", Desc: "errors and alerts"},                  // foreground - with bg-invalid
			{Key: "bg-valid", R: 233, G: 255, B: 233, Desc: "ok, valid, input background"}, // slight hue of input-bg, otherwise too annoying for empty inputs
			{Key: "bg-invalid", R: 255, G: 240, B: 240, Desc: "input background"},          //
			{Key: "pri", R: 000, G: 105, B: 180, Desc: "primary color - fonts and icons"},
			{Key: "pri-hov", R: 002, G: 134, B: 228, Desc: "hover   - slightly lighter"},
			{Key: "pri-vis", R: 000, G: 071, B: 122, Desc: "visited - slightly darker"},
			{Key: "sec", R: 228, G: 223, B: 206, Desc: "secondary color - for backgrounds"},
			{Key: "sec-drk1", R: 219, G: 216, B: 194, Desc: "darker, for menu 3"},
			{Key: "sec-drk2", R: 190, G: 187, B: 170, Desc: "darker, for borders"},
			{Key: "sec-lgt1", R: 236, G: 232, B: 221, Desc: "ligher"},
			{Key: "sec-lgt2", R: 241, G: 239, B: 232, Desc: "ligher more"},

			{Key: "zew2-md", R: 207, G: 136, B: 135},
			{Key: "zew2-dk", R: 177, G: 29, B: 28},
			{Key: "zew3-md", R: 138, G: 187, B: 206},
			{Key: "zew3-dk", R: 22, G: 119, B: 158},
			{Key: "zew4-md", R: 202, G: 192, B: 156},
			{Key: "zew4-dk", R: 149, G: 129, B: 58},
			{Key: "zew5-md", R: 233, G: 206, B: 134},
			{Key: "zew5-dk", R: 211, G: 158, B: 13},
		},
		CSSVarsSite: map[string]cssVars{
			"4walls": {
				{IsURL: true, Key: "img-bg", Val: "none"},
				{IsURL: true, Key: "img-loggedin-icon", Val: "/img/ui/logged-in-icon-4walls.svg"},
				{Key: "bg", R: 12, G: 12, B: 12, Desc: "main background f <body>"},
				{Key: "fg", R: 224, G: 224, B: 224, Desc: "main foreground"},
				{Key: "input-bg", R: 224, G: 224, B: 224, Desc: "input+select background"},
				{Key: "input-fg", R: 12, G: 12, B: 12, Desc: "input+select foreground"},
				{Key: "pri", R: 216, G: 29, B: 160, Desc: "primary color - fonts and icons"},
				{Key: "pri-hov", R: 250, G: 50, B: 200, Desc: "hover   - slightly lighter"},
				{Key: "pri-vis", R: 166, G: 12, B: 120, Desc: "visited - slightly darker"},

				{Key: "pri", R: 247, G: 19, B: 78, Desc: "primary color - fonts and icons"},
				{Key: "pri-hov", R: 255, G: 45, B: 100, Desc: "hover   - slightly lighter"},
				{Key: "pri-vis", R: 200, G: 9, B: 90, Desc: "visited - slightly darker"},

				{Key: "sec", R: 48, G: 48, B: 48, Desc: "secondary color - for backgrounds"},
				{Key: "sec-drk1", R: 32, G: 32, B: 32, Desc: "darker, for menu 3"},
				{Key: "sec-drk2", R: 1, G: 1, B: 1, Desc: "darker, for borders"},
			},
			"fmt": {
				{IsURL: true, Key: "img-bg", Val: "none"},
			},
			"pat": {
				{IsURL: true, Key: "img-bg", Val: "none"},
			},
		},
		CPUProfile: "", // the filename, i.e. cpu.pprof

		Profiles: map[string]map[string]string{
			"fmt1": {
				"lang_code":               "de",
				"main_refinance_rate_ecb": "3.5",
			},
			"fmt2": {
				"lang_code":               "en",
				"main_refinance_rate_ecb": "3.5",
			},
		},
		AnonymousSurveyID: "4walls",
		DirectLoginRanges: []directLoginRangeT{
			//
			// user-range  matchings first
			// survey type matchings second
			//
			//
			// matching by user id within start...stop
			// if ranges repeat, first match wins
			{
				Start:    1000 + 0,
				Stop:     1000 + 29,
				SurveyID: "flit",
				WaveID:   "2019-09",
				Profile: map[string]string{
					"lang_code":               "de",
					"main_refinance_rate_ecb": "3.5",
				},
			},
			{
				Start:    1000 + 50,
				Stop:     1000 + 59,
				SurveyID: "flit",
				WaveID:   "2019-09",
				Profile: map[string]string{
					"lang_code":               "en",
					"main_refinance_rate_ecb": "3.5",
				},
			},
			// matching by survey type
			// unrestricted by user id in any range
			//
			// *every* match of SurveyID wins
			{
				// Start:    any
				// Stop:     any
				SurveyID: "pat1",
				WaveID:   "2021-05",
				Profile: map[string]string{
					"lang_code":               "en",
					"main_refinance_rate_ecb": "3.5",
				},
			},
			{
				// Start:    any
				// Stop:     any
				SurveyID: "pat2",
				WaveID:   "2021-05",
				Profile: map[string]string{
					"lang_code":               "en",
					"main_refinance_rate_ecb": "3.5",
				},
			},
		},
		//
		// first
		// second example translation *appends*    trl.CoreTranslations()
		Mp: trl.Map{
			// example translation *keeps*      trl.CoreTranslations() defaults
			//    no "app_org": {...},
			// example translation *overwrites* trl.CoreTranslations()
			"app_label": {
				"de": "Meine --spezielle-- Beispiel Anwendung",
				"en": "My --special-- example app",
				"es": "Mi aplicación de ejemplo",
				"fr": "Mon exemple d'application",
				"it": "La mia App esempio",
				"pl": "Moja Przykładowa aplikacja",
			},
			// example translation *appends*    trl.CoreTranslations()
			"app_dept": {
				"de": "Meine Abteilung",
				"en": "My Department",
			},
		},
		MpSite: trl.MapSite{
			"default": {
				"img_logo_icon": {
					"default": "/img/ui/icon-forschung-zew-prim.svg",
				},
				"img_logo_icon_mobile": {
					"default": "/img/ui/icon-forschung-zew-prim.svg",
				},
			},
			"4walls": {
				"app_label": {
					"en": "4walls",
				},
				"img_logo_icon": {
					"default": "/img/ui/4walls-logo-3.png",
				},
				"img_logo_icon_mobile": {
					"default": "/img/ui/4walls-logo-3.png",
				},
			},
			"fmt": {
				"app_label": {
					"de": "Finanzmarkttest",
					"en": "financial markets survey",
				},
				"img_logo_icon": {
					"default": "/img/ui/icon-forschung-zew-prim.svg",
				},
				"img_logo_icon_mobile": {
					"default": "/img/ui/icon-forschung-zew-prim.svg",
				},
			},
			"pat": {
				"app_label": {
					"de": "Entscheidungsprozesse in der Politik",
				},
				"img_logo_icon": {
					"default": "/img/pat/vier-uni-logos.png",
				},
				"img_logo_icon_mobile": {
					"default": "/img/pat/vier-uni-logos-mobile.png",
				},
			},
		},
	}
	return ex
}
