// Package cfg implements a configuration database,
// loaded from a json file.
// Filename must be given as command line argument
// or environment variable.
// Access to the config data is made in threadsafe manner.
package cfg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sort"
	"strconv"
	"sync"
	"time"

	"github.com/zew/util"

	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/trl"
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

	CSS map[string]string `json:"css"` // differentiate multiple instances by color and stuff - without duplicating entire css files

	AllowSkipForward bool `json:"allow_skip_forward"` // skipping back always allowed, skipping forward is configurable

	// available language codes for the application, first element is default
	LangCodes []string `json:"lang_codes"`
	// multi language strings for the application
	Mp               trl.Map        `json:"translation_map"`
	UserIDToLanguage map[int]string `json:"user_id_to_language,omitempty"` // user id to language preselection for direct login
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
	dmp := util.IndentedDump(cfgS)
	if len(dmp) > 700 {
		dmp = dmp[:700]
	}
	log.Printf("config loaded 1\n%s", dmp)
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

// UserLangCode returns the language code for a user ID.
// Its taken from ConfigT.UserIDToLanguage
//
// user_id_to_language: {
// 	      1:   fr,
// 	   1305:   en,
//  }
//
func (c *ConfigT) UserLangCode(userIDStr string) (ret string, err error) {
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return
	}
	idxUnsorted := make([]int, 0, len(c.UserIDToLanguage))
	for id := range c.UserIDToLanguage {
		idxUnsorted = append(idxUnsorted, id)
	}
	sort.Ints(idxUnsorted)
	for _, key := range idxUnsorted {
		if key > userID {
			return
		}
		lc := c.UserIDToLanguage[key]
		log.Printf("UserID %5v: key %4v for UserIDToLanguage[key] %v", userID, key, lc)
		ret = lc
	}
	err = fmt.Errorf("No language code found for %v", userID)
	return
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
func Example() *ConfigT {
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

		LangCodes: []string{"de", "en", "es", "fr", "it", "pl"},
		UserIDToLanguage: map[int]string{
			1:        "fr",
			578:      "de",
			1305:     "en", // Belgians
			1326:     "en", // bg
			1343:     "en", // cy
			1349:     "en", // cz
			1370:     "de",
			1466:     "en", // dk
			1479:     "en", // ee
			1484:     "es",
			1538:     "en", // fi
			1551:     "fr",
			1625:     "en",
			1690:     "en", // gr
			1711:     "en", // hr
			1722:     "en", // hu
			1743:     "en", // ie
			1754:     "it",
			1828:     "en", // lt
			1834:     "en", // lv
			1835:     "en", // lt
			1840:     "en", // lu
			1845:     "en", // lv
			1852:     "en", // mt
			1858:     "en", // nl
			1884:     "pl",
			1935:     "en", // pt
			1956:     "en", // ro
			1988:     "en", // se
			2008:     "en", // si
			2016:     "en", // sk
			2030:     "en",
			2037:     "fr",
			2384:     "it",
			10000000: "it", // needed for closure
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
				"fr": "Impresión",
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
				"fr": "Terminé cette enquête",
				"it": "Finito questo sondaggio",
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
	ex.Save("config-example.json")
	return ex
}
