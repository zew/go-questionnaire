// Package cfg implements a configuration database,
// loaded from a json file.
// Filename must be given as command line argument
// or environment variable.
// Access to the config data is made in threadsafe manner.
package cfg

import (
	"encoding/json"
	"io"
	"log"
	"path"
	"path/filepath"
	"time"

	"github.com/zew/util"

	"github.com/zew/go-questionnaire/trl"
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

	AppName       string `json:"app_name"`       // with case, i.e. 'My App'
	URLPathPrefix string `json:"urlpath_prefix"` // lower case - no slashes, i.e. 'myapp'
	AppMnemonic   string `json:"app_mnemonic"`   // For differentiation of static dirs - when URLPathPrefix is empty; imagine multiple instances

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

	LocationName   string         `json:"location,omitempty"` // i.e. "Europe/Berlin", see Go\lib\time\zoneinfo.zip
	Loc            *time.Location `json:"-"`                  // Initialized during load
	SessionTimeout int            `json:"session_timeout"`    // hours until the session is lost
	FormTimeout    int            `json:"form_timeout"`       // hours until a form post is rejected

	AppInstanceID int64    `json:"app_instance_id,omitempty"` // append to URLs of cached static jpg, js and css files - change to trigger reload
	LangCodes     []string `json:"lang_codes"`

	CSSVars     cssVars            `json:"css_vars"`      // global CSS variables
	CSSVarsSite map[string]cssVars `json:"css_vars_site"` // site specific overwrites of css_vars

	CPUProfile string `json:"cpu_profile"` // the filename to write to

	AllowSkipForward bool `json:"allow_skip_forward"` // skipping back always allowed, skipping forward is configurable

	// available language codes for the application, first element is default
	// multi language strings for the application
	Mp trl.Map `json:"translation_map"`

	// Profiles are sets of attributes, selected by the `p` parameter at login, containing key-values which are copied into the logged in user's attributes
	Profiles map[string]map[string]string

	AnonymousSurveyID string              `json:"anonymous_survey_id,omitempty"` // anonymous login - redirect / forward url
	DirectLoginRanges []directLoginRangeT `json:"direct_login_ranges,omitempty"` // user id to language preselection for direct login
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
		log.Printf("Your location name must be valid, i.e. 'Europe/Berlin', compare Go\\lib\\time\\zoneinfo.zip: %v", err)
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
	}

	//
	cfgS = &tempCfg // replace pointer in one go - should be threadsafe
	dmp := util.IndentedDump(cfgS)
	if len(dmp) > 700 {
		dmp = dmp[:700]
	}
	log.Printf("\n%s\n...config loaded", dmp)
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
//    mux.HandleFunc(appcfg.Pref(urlPath),   argFunc)   // Claim "/path"
//    mux.HandleFunc(appcfg.PrefTS(urlPath), argFunc)   // Claim "/path/"
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

// Pref for templates: cfg.Pref
func (c *ConfigT) Pref(pth ...string) string {
	return Pref(pth...)
}

// PrefTS is like Prefix(); TS stands for (with) trailing slash;
// useful for registering handlers
// so that /p1/p2/  also serves /p1/p2
func PrefTS(pth ...string) string {
	p := Pref(pth...)
	return p + "/"
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
		CSSVars: cssVars{
			{Key: "logo-text", Val: "ZEW"},
			{IsURL: true, Key: "img-bg", Val: "/img/ui/bg-bw-bland.jpg"},
			{IsURL: true, Key: "img-logo-icon", Val: "/img/ui/icon-forschung-zew-prim.svg"},
			{IsURL: true, Key: "img-loggedin-icon", Val: "/img/ui/logged-in-icon-zew.svg"},
			{Key: "nav-height", Val: "8vh"},
			{Key: "nav-rest-height", Val: "calc(100vh - var(--nav-height))", Desc: "we can calc() the remainder"},
			{Key: "nav-bar-position", Val: "relative", Desc: "fixed or relative"},
			{Key: "content-top", Val: "0", Desc: "fixed navbar => content-top = var(--nav-height); otherwise 0"},
			{Key: "bg", Colorname: "white", Desc: "main background f <body>"},
			{Key: "fg", Colorname: "black", Desc: "main foreground"},
			{Key: "input-bg", Colorname: "white", Desc: "input+select background"},
			{Key: "input-fg", Colorname: "black", Desc: "input+select foreground"},
			{Key: "valid", R: 233, G: 255, B: 233, Alpha: .999, Desc: "ok, valid"}, // slight hue of input-bg, otherwise too annoying for empty inputs
			{Key: "err", Colorname: "lightcoral", Desc: "errors and alerts"},
			{Key: "pri", R: 000, G: 105, B: 180, Alpha: .999, Desc: "primary color"},
			{Key: "pri-hov", R: 002, G: 134, B: 228, Desc: "hover   - slightly lighter"},
			{Key: "pri-vis", R: 000, G: 071, B: 122, Desc: "visited - slightly darker"},
			{Key: "sec", R: 228, G: 223, B: 206, Alpha: 1.0},
			{Key: "sec-drk1", R: 219, G: 216, B: 194, Desc: "darker, for menu 3"},
			{Key: "sec-drk2", R: 190, G: 187, B: 170, Desc: "darker, for borders"},
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
			"site1": {
				{Key: "logo-text", Val: "4WALLS"},
				{IsURL: true, Key: "img-bg", Val: "none"},
				{IsURL: true, Key: "img-logo-icon", Val: "/img/ui/4walls-logo-3.png"},
				{IsURL: true, Key: "img-loggedin-icon", Val: "/img/ui/logged-in-icon-4walls.svg"},
				{Key: "bg", R: 12, G: 12, B: 12, Desc: "main background f <body>"},
				{Key: "fg", R: 224, G: 224, B: 224, Alpha: .999, Desc: "main foreground"},
				{Key: "input-bg", R: 224, G: 224, B: 224, Desc: "input+select background"},
				{Key: "input-fg", R: 12, G: 12, B: 12, Desc: "input+select foreground"},
				{Key: "pri", R: 216, G: 29, B: 160, Alpha: .999, Desc: "primary color"},
				{Key: "pri-hov", R: 250, G: 50, B: 200, Desc: "hover   - slightly lighter"},
				{Key: "pri-vis", R: 166, G: 12, B: 120, Desc: "visited - slightly darker"},

				{Key: "pri", R: 247, G: 19, B: 78, Alpha: .999, Desc: "primary color"},
				{Key: "pri-hov", R: 255, G: 45, B: 100, Desc: "hover   - slightly lighter"},
				{Key: "pri-vis", R: 200, G: 9, B: 90, Desc: "visited - slightly darker"},

				{Key: "sec", R: 48, G: 48, B: 48, Alpha: 1.0},
				{Key: "sec-drk1", R: 32, G: 32, B: 32, Desc: "darker, for menu 3"},
				{Key: "sec-drk2", R: 1, G: 1, B: 1, Desc: "darker, for borders"},
			}, "zew": {
				{Key: "dummy", R: 48, G: 48, B: 48, Alpha: 1.0},
			},
		},
		AppInstanceID: time.Now().Unix(),
		LangCodes:     []string{"de", "en", "es", "fr", "it", "pl"},
		CPUProfile:    "", // the filename, i.e. cpu.pprof

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
		},
		Mp: trl.Map{
			"page": {
				"en": "Page",
				"de": "Seite",
				"es": "Página",
				"fr": "Page",
				"it": "Pagina",
				"pl": "Strona",
			},
			"start": {
				"en": "Start",
				"de": "Start",
				"es": "Inicia",
				"fr": "Commencer",
				"it": "Inizia",
				"pl": "Uruchomić",
			},
			"next": {
				"en": "Next",
				"de": "Weiter",
				"es": "Continuar",
				"fr": "Continuer",
				"it": "Continuare",
				"pl": "Kontynuować",
			},
			"end": {
				"en": "End",
				"de": "Ende",
				"es": "Fin",
				"fr": "Fin",
				"it": "Fine",
				"pl": "Końcu",
			},
			"app_label_h1": {
				"en": "My Org",
				"de": "Meine Organisation",
				"es": "Mi organización",
				"fr": "Mon organisation",
				"it": "La mia organizzazion",
				"pl": "Moja organizacja",
			},
			"app_label": {
				"en": "My Example App", // yes, repeat of AppName
				"de": "Meine Beispiel Anwendung",
				"es": "Mi aplicación de ejemplo",
				"fr": "Mon exemple d'application",
				"it": "La mia App esempio",
				"pl": "Moja Przykładowa aplikacja",
			},

			"correct_errors": {
				"de": "Bitte korrigieren Sie die unten angezeigten Fehler.",
				"en": "Please correct the errors displayed below.",
				"es": "Por favor corrija los errores que aparecen a continuación",
				"fr": "Veuillez corriger les erreurs affichées ci-dessous",
				"it": "Per piacere correga gli errori sottostanti.",
				"pl": "Popraw błędy wyświetlane poniżej",
			},
			"imprint": {
				"de": "Impressum",
				"en": "Imprint",
				"es": "Empreinte",
				"fr": "Mentions légales", //"Impresión" ahv,
				"it": "Impronta",
				"pl": "Nadrukiem",
			},
			"login_by_hash_failed": {
				"de": "Anmeldung via Hash gescheitert.\nBitte nutzen Sie den übermittelten Link um sich anzumelden.\nWenn der Link in zwei Zeilen geteilt wurde, verbinden Sie die Zeilen wieder.",
				"en": "Login by hash failed.\nPlease use the provided link to login.\nIf the link was split into two lines, reconnect them.",
				"es": "Error al iniciar sesión por hash.\nPor favor, utilice el enlace proporcionado para iniciar sesión.\nSi el enlace se dividió en dos líneas, vuelva a conectarlas.",
				"fr": "Login par hachage a échoué.\nVeuillez utiliser le lien fourni pour vous connecter.\nSi le lien a été divisé en deux lignes, reconnectez-les.",
				"it": "Il login non è andato a buon fine.\nPer piacere si utilizzi il link fornitovi per effettuare il login.\nSe il link è spezzato in due, le due parti devono essere riconnesse.",
				"pl": "Logowanie przez hash nie powiodło się. \nProszę użyć przesłanego linku, aby się zarejestrować. \nJeśli łącze zostało podzielone na dwa wiersze, Połącz ponownie wiersze.",
			},
			"finish_questionnaire": {
				"de": "Fragebogen abschließen",
				"en": "Finish this survey",
				"es": "Terminé esta encuesta",
				"fr": "Terminer ce sondage", // "Terminé cette enquête" ahv,
				"it": "Finire questo sondaggio",
				"pl": "Zakończyłem tę ankietę",
			},
			"finished_by_participant": {
				"de": "Sie haben den Fragebogen bereits abgeschlossen (%v).",
				"en": "You already finished this survey wave at %v",
				"es": "Usted ya terminó esta ola de encuestas en %v",
				"fr": "Vous avez déjà terminé cette vague de sondage à %v",
				"it": "Lei ha già completato questo questionario (%v)",
				"pl": "Już skończyłeś tę falę pomiarową na %v",
			},
			"deadline_exceeded": {
				"de": "Diese Umfrage wurde am %v beendet.",
				"en": "Current survey was closed at %v.",
				"es": "La encuesta actual se cerró en %v",
				"fr": "L'enquête en cours a été clôturée à %v",
				"it": "Questo questionario è stato chiuso il %v.",
				"pl": "Aktualna Ankieta została zamknięta w %v",
			},
			"percentage_answered": {
				"de": "Sie haben %v von %v Fragen beantwortet: %2.1f Prozent.  <br>\n",
				"en": "You answered %v out of %v questions: %2.1f percent.  <br>\n",
				"es": "Usted contestó %v de %v preguntas: %2.1f por ciento. <br>\n",
				"fr": "Vous avez répondu %v sur %v questions: %2.1f pour cent. <br>\n",
				"it": "Lei ha risposto a %v domande su %v: %2.1f per cento.  <br>\n",
				"pl": "Odpowiedziałeś %v na %v pytania: %2.1f procent. <br>\n",
			},
			"survey_ending": {
				"de": "Umfrage endet am %v. <br>\nVeröffentlichung am %v.  <br>\n",
				"en": "Survey will finish at %v. <br>\nPublication will be at %v.<br>\n",
				"es": "La encuesta terminará en %v.\nPublicación será en %v. <br>\n",
				"fr": "L'enquête se terminera à %v.\nPublication sera à %v. <br>\n",
				"it": "Il sondaggio verrà concluso il %v. <br>\nLa pubblicazione avverrà il %v.<br>\n",
				"pl": "Ankieta zakończy się w %v.\nPublikacja będzie %v. <br>\n",
			},
			"review_by_personal_link": {
				"de": "Sie können ihre Daten jederzeit über Ihren persönlichen Link prüfen/ändern. <br>\n<a href='/?submitBtn=prev'>Zurück</a><br>\n",
				"en": "You may review or change your data using your personal link. <br>\n<a href='/?submitBtn=prev'>Back</a><br>\n",
				"es": "Usted puede revisar o cambiar sus datos usando su enlace personal.<br>\n<a href='/?submitBtn=prev'>Atrás</a><br>\n",
				"fr": "Vous pouvez consulter ou modifier vos données à l'aide de votre lien personnel.<br>\n<a href='/?submitBtn=prev'>Précédent</a><br>\n",
				"it": "Può rivedere o modificare i suoi dati usando il Suo link personale. <br>\n<a href='/?submitBtn=prev'>Indietro</a><br>\n",
				"pl": "Dane można przejrzeć lub zmienić przy użyciu osobistego łącza.<br>\n<a href='/?submitBtn=prev'>Wstecz</a><br>\n",
			},
			"not_a_number": {
				"de": "'%v' keine Zahl",
				"en": "'%v' not a number",
				"es": "'%v' no es un número",
				"fr": "'%v' pas un certain nombre",
				"it": "'%v' non è un numero",
				"pl": "'%v' nie liczba",
			},
			"too_big": {
				"de": "Max %.0f",
				"en": "max %.0f",
				"es": "Máximo %.0f",
				"fr": "Max %.0f",
				"it": "Massimo %.0f",
				"pl": "Maksymalna %.0f",
			},
			"too_small": {
				"de": "Min %.0f",
				"en": "min %.0f",
				"es": "Mínimo %.0f",
				"fr": "Min %.0f",
				"it": "Minimo %.0f",
				"pl": "Minimalne %.0f",
			},
			"must_one_option": {
				"de": "Bitte eine Option wählen",
				"en": "Please choose one option",
				"es": "Por favor elija una opción",
				"fr": "Veuillez choisir une option",
				"it": "Si prega di selezionare una opzione",
				"pl": "Proszę wybrać jedną z opcji",
			},
			"yes": {
				"de": "Ja",
				"en": "Yes",
				"es": "Sí",
				"fr": "Oui",
				"it": "Sì",
				"pl": "Tak",
			},
			"no": {
				"de": "Nein",
				"en": "No",
				"es": "No",
				"fr": "Non",
				"it": "No",
				"pl": "Nie",
			},
		},
	}
	return ex
}
