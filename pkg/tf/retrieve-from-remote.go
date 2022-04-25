// Package tf - transferrer - fetches completed questionnaires
// from /transferrer-endpoint as gzipped JSON via http(s) request;
// downloads and CSVs are stored to ./app-bucket/responses/downloaded;
//
// multiple configs are required in ./app-bucket/transferrer;
// config-autogen.json      is mostly a dummy to satisfy bootstrap;
// logins-remote-salt.json  is needed to login remotely;
// remote-fmt.json or remote-fmt-localhost.json contain
// 			https POST request data
// 			destination survey
// 			login
package tf

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/pbberlin/flags"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/util"
)

func getClient() *http.Client {
	client := util.HttpClient()
	client = &http.Client{}
	log.Printf("client timeout is %v", client.Timeout)
	return client
}

// RemoteConnConfigT is on top of
// of the ordinary config
type RemoteConnConfigT struct {
	RemoteHost    string
	BindSocket    string
	URLPathPrefix string

	AdminLogin string // Some admin account of the remote machine
	Pass       string

	SurveyType string
	WaveID     string // special value "current" is evaluated to current year

	DownloadDir string
	MinUserID   int // constrain range of UserIDs being processed, exclude test user data entry
	MaxUserID   int // see MinUserID
}

// Example returns a minimal configuration, to be extended or adapted
func Example() RemoteConnConfigT {
	r := RemoteConnConfigT{}
	r.RemoteHost = "https://www.peu2018.eu"
	r.RemoteHost = "https://financial-literacy-test.appspot.com"
	r.RemoteHost = "https://survey2.zew.de"

	r.BindSocket = "443"
	r.URLPathPrefix = "survey"
	r.URLPathPrefix = ""

	r.AdminLogin = "transferrer"
	r.Pass = "Spark!sh32"

	r.SurveyType = "fmt"
	r.SurveyType = "flit"
	r.SurveyType = "lt2020"

	r.WaveID = qst.NewSurvey(r.SurveyType).WaveID()
	r.WaveID = "2020-05"

	r.DownloadDir = "responses/downloaded"

	return r
}

// LoadRemote reads from an io.Reader
// to avoid cyclical deps.
func LoadRemote(r io.Reader) *RemoteConnConfigT {

	log.Printf("Loading remote config...")

	decoder := json.NewDecoder(r)
	tempCfg := RemoteConnConfigT{}
	err := decoder.Decode(&tempCfg)
	if err != nil {
		log.Fatalf("error decoding into RemoteConnConfigT: %v", err)
	}

	if tempCfg.WaveID == "current" {
		tNow := time.Now()
		tempCfg.WaveID = fmt.Sprintf("%v%02d", tNow.Year(), int(tNow.Month()))
	}

	return &tempCfg
}

// RetrieveFromServer requests the JSONified questionnaires
// from the survey server endpoint; decompresses the GZIPPed
// response and parses the bytes into a slice of questionnaires
func RetrieveFromServer() (
	[]*qst.QuestionnaireT,
	*RemoteConnConfigT,
	error,
) {

	// we need config and logins
	// for main app at least initialized
	{
		//
		// We take a config;
		// save it to file and then activate it from file.
		cf := &cfg.ConfigT{}
		cf.AppName = "Transferrer for Go Questionnaire - http client"
		cf.AppMnemonic = "tf"
		cf.LangCodes = []string{"en"}
		cf.Loc = time.FixedZone("UTC", 1*3600) // cf.Loc is needed below
		// cf.URLPathPrefix is needed for cfg.Pref() properly working
		// It is set later from transferrer config

		pthAutogen := path.Join("/transferrer", "config-autogen.json")
		cloudio.MarshalWriteFile(&cf, pthAutogen)
		cfg.CfgPath = pthAutogen

		fileName := cfg.CfgPath
		r, bucketClose, err := cloudio.Open(fileName)
		if err != nil {
			log.Fatalf("Error opening writer to %v: %v", fileName, err)
		}
		defer func() {
			if r != nil {
				err := r.Close()
				if err != nil {
					log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
				}
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
	}

	//
	//
	// logins data is directly read from file;
	// it only contains the remote salt
	// required to create form request tokens
	lgn.LgnsPath = path.Join("/transferrer", "logins-remote-salt.json")
	{
		fileName := lgn.LgnsPath
		r, bucketClose, err := cloudio.Open(fileName)
		if err != nil {
			log.Fatalf("Error opening writer to %v: %v", fileName, err)
		}
		defer func() {
			if r != nil {
				err := r.Close()
				if err != nil {
					log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
				}
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

		cloudio.MarshalWriteFile(lgn.Example(), path.Join("/transferrer", "logins-example.json"))

	}

	//
	//
	//
	// the actual config for *this* app:
	fl := flags.New()
	fl.Add(
		flags.FlagT{
			Long:       "remote_file",
			Short:      "rmt",
			DefaultVal: path.Join("/transferrer", "remote.json"),
			Desc:       "JSON file containing connection to remote host",
		},
	)
	fl.Gen()
	var cfgRem RemoteConnConfigT
	cfgRem = Example()
	cloudio.MarshalWriteFile(&cfgRem, path.Join("/transferrer", "example-remote.json"))
	{
		rmt := fl.ByKey("rmt").Val
		fileName := rmt
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
		cfgRem = *(LoadRemote(r))
	}

	//
	//
	//
	//
	//

	// make cfg.Pref() work properly:
	cfg.Get().URLPathPrefix = cfgRem.URLPathPrefix

	host := fmt.Sprintf("%v:%v", cfgRem.RemoteHost, cfgRem.BindSocket)
	if cfgRem.BindSocket == "" {
		host = cfgRem.RemoteHost
	}
	log.Printf("Remote host is: %v", host)

	defer func() {
		log.Printf("  ")
		log.Printf("  ")
		log.Printf("================")
		log.Printf("Login         via   %v%v%v", host, cfg.Pref(), "login-primitive")
		log.Printf("Check results via   %v%v%v", host, cfg.Pref(), "transferrer-endpoint?...")
	}()

	urlLogin := host + cfg.Pref("/login-primitive")
	log.Printf("url import %v", urlLogin)

	urlMain := host + cfg.Pref("/transferrer-endpoint")
	log.Printf("url main   %v", urlMain)

	//
	// Login and save session cookie
	var sessCook *http.Cookie
	{
		log.Printf(" ")
		log.Printf("Getting cookie")
		log.Printf("==================")
		urlReq := urlLogin

		vals := url.Values{}
		vals.Set("username", cfgRem.AdminLogin)
		vals.Set("password", cfgRem.Pass)
		vals.Set("token", lgn.FormToken())
		req, err := http.NewRequest(
			"POST",
			urlReq,
			strings.NewReader(vals.Encode()), // <-- URL-encoded payload
		)
		if err != nil {
			return nil, &cfgRem, fmt.Errorf("error creating request for %v: %w", urlReq, err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := getClient().Do(req)
		if err != nil {
			log.Printf("error requesting cookie from %v: %v; %v", urlReq, err, resp)
		}

		if resp == nil {
			return nil, &cfgRem, fmt.Errorf("response is nil - from %v", urlReq)
		}

		defer resp.Body.Close()

		for _, v := range resp.Cookies() {
			if v.Name == "session" {
				sessCook = v
			}
		}
		respBytes, _ := io.ReadAll(resp.Body)
		mustHave := fmt.Sprintf("Logged in as %v", cfgRem.AdminLogin)
		if !strings.Contains(string(respBytes), mustHave) {
			log.Fatalf(
				"Login response must contain '%v'\n%v\n%v\n\n%v",
				mustHave, urlReq, vals, string(respBytes),
			)
		}

		log.Printf("Cookie is %+v \ngleaned from %v", sessCook, req.URL)
		if sessCook == nil {
			return nil, &cfgRem, fmt.Errorf("we need a session cookie to continue")
		}

		if !strings.Contains(string(respBytes), "Logged in as "+cfgRem.AdminLogin) {
			return nil, &cfgRem, fmt.Errorf("Response must contain 'Logged in as %v' \n\n%v", cfgRem.AdminLogin, string(respBytes))
		}

	}

	//
	//
	// Post values and check the response
	log.Printf(" ")
	log.Printf("Transferrer endpoint")
	log.Printf("==================")
	urlReq := urlMain

	vals := url.Values{}
	vals.Set("survey_id", cfgRem.SurveyType)
	vals.Set("wave_id", cfgRem.WaveID)
	vals.Set("fetch_all", "1")
	method := "POST"
	log.Printf("%v requesting %v?%v", method, urlReq, vals.Encode())
	req, err := http.NewRequest(method, urlReq, bytes.NewBufferString(vals.Encode())) // <-- URL-encoded payload
	if err != nil {
		return nil, &cfgRem, fmt.Errorf("error creating request %v: %w", urlReq, err)
	}
	// strangely, the json *response* is empty, if we omit this:
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	for _, v := range []*http.Cookie{sessCook} {
		req.AddCookie(v)
	}

	resp, err := getClient().Do(req)
	if err != nil {
		return nil, &cfgRem, fmt.Errorf("error requesting %v: %w", urlReq, err)
	}

	defer resp.Body.Close()
	var rdr1 io.ReadCloser

	log.Printf("Content encoding is -%v-", resp.Header.Get("Content-Encoding"))

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":

		if false {
			// a hack - to spy into the response
			// if the http download does not work...
			bts, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, &cfgRem, fmt.Errorf("could not read all response %w", err)
			}
			log.Printf("response is %v bytes", len(bts))
			if len(bts) < 15000 {
				log.Printf("response is %s", bts)
			}
			btsRdr := bytes.NewReader(bts)
			rdr1, err = gzip.NewReader(btsRdr)
			if err != nil {
				return nil, &cfgRem, fmt.Errorf("could not read the response as gzip #1: %w", err)
			}
		}

		rdr1, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, &cfgRem, fmt.Errorf("could not read the response as gzip #2: %w", err)
		}
		defer rdr1.Close()
	default:
		rdr1 = resp.Body
	}

	if false {
		// a hack - to load from file
		// if the http download would not work...
		fr, err := os.Open("./app-bucket/dl/transferrer-endpoint.json")
		if err != nil {
			log.Printf("shortcut file not present; %v", err)
		} else {
			rdr1 = fr
		}
	}

	// Check response status
	rsc := resp.StatusCode
	if rsc != http.StatusOK && rsc != http.StatusTemporaryRedirect && rsc != http.StatusSeeOther {
		return nil, &cfgRem, fmt.Errorf("bad response %v - %q ", resp.StatusCode, resp.Status)
	}

	dec := json.NewDecoder(rdr1)
	qs := []*qst.QuestionnaireT{}
	err = dec.Decode(&qs)
	if err != nil {
		log.Printf("Unmarshal failed: %v", err)
		fn := "tmp-transferrer-endpoint-response-error.html"
		bts, _ := io.ReadAll(rdr1)
		os.WriteFile(fn, bts, 0777)
		return nil, &cfgRem, fmt.Errorf("Response written to %v", fn)

	}
	log.Printf("Unmarshalled %v questionnaires from responese stream", len(qs))
	log.Printf("====================================================")

	return qs, &cfgRem, nil

}
