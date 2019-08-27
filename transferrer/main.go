// Package transferrer fetches completed questionnaires
// from /transferrer-endpoint as JSON via http request.
package main

import (
	"bytes"
	"compress/gzip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
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
	BindSocket    int
	URLPathPrefix string

	AdminLogin string // Some admin account of the remote machine
	Pass       string

	SurveyType string
	WaveID     string

	DownloadDir string
}

// Example returns a minimal configuration, to be extended or adapted
func Example() RemoteConnConfigT {
	r := RemoteConnConfigT{}
	r.RemoteHost = "https://survey2.zew.de"
	r.RemoteHost = "https://www.peu2018.eu"

	r.BindSocket = 443
	r.URLPathPrefix = "taxdb"

	r.AdminLogin = "login"
	r.Pass = "secret"

	r.SurveyType = "fmt"
	r.SurveyType = "peu2018"

	r.WaveID = qst.NewSurvey(r.SurveyType).WaveID()
	r.WaveID = "2018-08"

	r.DownloadDir = path.Join(qst.BasePath(), "downloaded")

	return r
}

// Save writes a JSON file
func (r *RemoteConnConfigT) Save(fn string) {
	byts, err := json.MarshalIndent(r, " ", "\t")
	util.BubbleUp(err)
	err = ioutil.WriteFile(fn, byts, 0644)
	util.BubbleUp(err)
}

// Load loads a JSON file
func Load(fn string) (r *RemoteConnConfigT) {
	file, err := util.LoadConfigFile(fn)
	if err != nil {
		log.Fatalf("Could not load config file %v: %v", fn, err)
	}
	log.Printf("Found config file: %v", fn)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Err closing config file %v: %v", fn, err)
		}
		log.Printf("Closed config file: %v", fn)
	}()
	decoder := json.NewDecoder(file)
	tempCfg := RemoteConnConfigT{}
	err = decoder.Decode(&tempCfg)
	if err != nil {
		log.Fatal(err)
	}
	return &tempCfg
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// For database files, static files and templates relative paths to work,
	// as if running from main app dir:
	err := os.Chdir("..")
	if err != nil {
		log.Fatalf("Error - cannot 'cd' to main app dir: %v", err)
	}
	// if doChDirUp {
	// }

	// We need config and logins
	// for main app at least initialized
	// See below.
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
	}

	// logins data is directly read from file
	// It only contains the remote salt
	// required to create form request tokens
	lgn.LgnsPath = path.Join("/transferrer", "logins-remote-salt.json")
	{
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

		cloudio.MarshalWriteFile(lgn.Example(), path.Join("/transferrer", "logins-example.json"))

	}

	return

	// The actual config for *this* app:
	fl := util.NewFlags()
	fl.Add(
		util.FlagT{
			Long:       "remote_file",
			Short:      "rmt",
			DefaultVal: "transferrer-remote.json",
			Desc:       "JSON file containing connection to remote host",
		},
	)
	fl.Gen()

	var c2 RemoteConnConfigT
	c2 = Example()
	c2.Save("transferrer-remote-example.json")
	rmt := fl.ByKey("rmt").Val
	c2 = *(Load(rmt))

	// make cfg.Pref() work properly:
	cfg.Get().URLPathPrefix = c2.URLPathPrefix

	host := fmt.Sprintf("%v:%v", c2.RemoteHost, c2.BindSocket)
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
		vals.Set("username", c2.AdminLogin)
		vals.Set("password", c2.Pass)
		vals.Set("token", lgn.FormToken())
		req, err := http.NewRequest("POST", urlReq, bytes.NewBufferString(vals.Encode())) // <-- URL-encoded payload
		if err != nil {
			log.Printf("error creating request for %v: %v", urlReq, err)
			return
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := getClient().Do(req)
		if err != nil {
			log.Printf("error requesting cookie from %v: %v; %v", urlReq, err, resp)
		}
		defer resp.Body.Close()
		for _, v := range resp.Cookies() {
			if v.Name == "session" {
				sessCook = v
			}
		}
		respBytes, _ := ioutil.ReadAll(resp.Body)
		mustHave := fmt.Sprintf("Logged in as %v", c2.AdminLogin)
		if !strings.Contains(string(respBytes), mustHave) {
			log.Fatalf("Login response must contain '%v'\n%v\n%v\n\n%v", mustHave, urlReq, vals, string(respBytes))
		}

		log.Printf("Cookie is %+v \ngleaned from %v", sessCook, req.URL)
		if sessCook == nil {
			log.Printf("we need a session cookie to continue")
			return
		}

		if !strings.Contains(string(respBytes), "Logged in as "+c2.AdminLogin) {
			log.Printf("Response must contain 'Logged in as %v' \n\n%v", c2.AdminLogin, string(respBytes))
			return
		}

	}

	//
	//
	// Post values and check the response
	{
		log.Printf(" ")
		log.Printf("Transferrer endpoint")
		log.Printf("==================")
		urlReq := urlMain

		vals := url.Values{}
		vals.Set("survey_id", c2.SurveyType)
		vals.Set("wave_id", c2.WaveID)
		vals.Set("fetch_all", "1")
		method := "POST"
		log.Printf("%v requesting %v?%v", method, urlReq, vals.Encode())
		req, err := http.NewRequest(method, urlReq, bytes.NewBufferString(vals.Encode())) // <-- URL-encoded payload
		if err != nil {
			log.Printf("error creating request %v: %v", urlReq, err)
			return
		}
		// strangely, the json *response* is empty, if we omit this:
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		for _, v := range []*http.Cookie{sessCook} {
			req.AddCookie(v)
		}

		resp, err := getClient().Do(req)
		if err != nil {
			log.Printf("error requesting %v: %v", urlReq, err)
			return
		}

		defer resp.Body.Close()
		var rdr1 io.ReadCloser
		switch resp.Header.Get("Content-Encoding") {
		case "gzip":
			// the server actually sent compressed data
			rdr1, err = gzip.NewReader(resp.Body)
			if err != nil {
				log.Printf("could not read the response as gzip: %v", err)
				return
			}
			defer rdr1.Close()
		default:
			rdr1 = resp.Body
		}

		// Check response status
		rsc := resp.StatusCode
		if rsc != http.StatusOK && rsc != http.StatusTemporaryRedirect && rsc != http.StatusSeeOther {
			log.Printf("bad response %q ", resp.Status)
			return
		}

		dir := path.Join(c2.DownloadDir, c2.SurveyType, c2.WaveID)
		dirEmpty := path.Join(dir, "empty")
		err = os.MkdirAll(dirEmpty, 0755)
		if err != nil {
			log.Printf("Could not create path 2 %v", dir)
			return
		}

		qs := []*qst.QuestionnaireT{}
		// err = json.Unmarshal(respBytes, &qs)
		dec := json.NewDecoder(rdr1)
		err = dec.Decode(&qs)
		if err != nil {
			log.Printf("Unmarshal %v", err)
			return
		}

		//
		//
		//
		maxPages := 0
		for _, q := range qs {
			if maxPages < len(q.Pages) {
				maxPages = len(q.Pages)
			}
		}

		allKeys := [][]string{}
		allVals := [][]string{}
		staticCols := []string{"user_id", "lang_code"}
		for iPg := 0; iPg < maxPages; iPg++ {
			staticCols = append(staticCols, fmt.Sprintf("page_%v", iPg+1))
		}

		for i, q := range qs {

			md5Want := q.MD5

			pth2 := path.Join(dir, q.UserID)
			err := q.Save1(pth2)
			if err != nil {
				log.Printf("%3v: Error saving %v: %v", i, pth2, err)
				continue
			}

			if q.MD5 != md5Want {
				log.Printf("%3v: MD5 does not match: %v\nwnt %v\ngot %v", i, pth2, md5Want, q.MD5)
				continue
			}

			realEntries, _, _ := q.Statistics()
			if realEntries == 0 {
				log.Printf("%3v: %v. No answers given, skipping.", i, pth2)
				pth2a := pth2 + ".json"
				pthEmpty := path.Join(dirEmpty, q.UserID+".json")
				err := os.Rename(pth2a, pthEmpty)
				if err != nil {
					log.Printf("%3v: Error moving %v to %v - %v", i, pth2a, pthEmpty, err)
				}
				continue
			} else {
				pthEmpty := path.Join(dirEmpty, q.UserID+".json")
				if _, err := os.Stat(pthEmpty); err == nil {
					err := os.Remove(pthEmpty)
					if err != nil {
						log.Printf("%3v: Error removing previously empty %v - %v", i, pthEmpty, err)
					}
				}
			}

			// Prepare columns...
			finishes, ks, vs := q.KeysValues()

			ks = append(staticCols, ks...)
			allKeys = append(allKeys, ks)

			//
			prepend := []string{qs[i].UserID, qs[i].LangCode}
			for iPg := 0; iPg < maxPages; iPg++ {
				if iPg < len(finishes) {
					prepend = append(prepend, finishes[iPg])
				} else {
					prepend = append(prepend, "n.a.") // response had less than max pages - not finishing time
				}
			}
			vs = append(prepend, vs...)
			allVals = append(allVals, vs)
		}

		allKeysSuperset := Superset(allKeys)

		allKeysSSMap := map[string]int{}
		for idx, v := range allKeysSuperset {
			allKeysSSMap[v] = idx
		}
		valsBySuperset := [][]string{}

		log.Printf("%v keys superset; %v", len(allKeysSuperset), util.IndentedDump(allKeysSuperset))
		log.Printf("%v map  keys    ; %v", len(allKeysSSMap), util.IndentedDump(allKeysSSMap))
		// log.Printf("%v", util.IndentedDump(allVals))

		// Collect values...
		for i1 := 0; i1 < len(allVals); i1++ {
			keys := allKeys[i1]
			vals := allVals[i1]
			valsBySuperset = append(valsBySuperset, make([]string, len(allKeysSuperset)))
			for i2 := 0; i2 < len(keys); i2++ {
				v := vals[i2]
				k := keys[i2]
				pos := allKeysSSMap[k]
				valsBySuperset[i1][pos] = v
			}
		}

		// Data into CSV matrix...
		var wtr = new(bytes.Buffer)
		csvWtr := csv.NewWriter(wtr)
		csvWtr.Comma = ';'
		if err := csvWtr.Write(allKeysSuperset); err != nil {
			log.Printf("error writing header line to csv: %v", err)
		}
		for _, record := range valsBySuperset {
			if err := csvWtr.Write(record); err != nil {
				log.Printf("error writing record to csv: %v", err)
			}
		}

		// Write any buffered data to the underlying writer (standard output).
		csvWtr.Flush()
		if err := csvWtr.Error(); err != nil {
			log.Printf("error flushing csv to response writer: %v", err)
		}
		err = ioutil.WriteFile("online-responses.csv", wtr.Bytes(), 0644)
		if err != nil {
			log.Printf("Could not write file: %v", err)
		}
		log.Printf("Regular finish. %v questionnaire(s) processed", len(qs))

	}

}
