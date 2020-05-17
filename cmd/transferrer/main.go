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
	BindSocket    string
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

	r.DownloadDir = "../../app-bucket/dl"

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
	return &tempCfg
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// we must change to main app dir,
	// so that referring to ./app-bucket works
	err := os.Chdir("../..")
	if err != nil {
		log.Fatalf("Error - cannot 'cd' to main app dir: %v", err)
	}

	// We need config and logins
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

	//
	// The actual config for *this* app:
	fl := util.NewFlags()
	fl.Add(
		util.FlagT{
			Long:       "remote_file",
			Short:      "rmt",
			DefaultVal: path.Join("/transferrer", "remote.json"),
			Desc:       "JSON file containing connection to remote host",
		},
	)
	fl.Gen()

	var c2 RemoteConnConfigT
	c2 = Example()
	cloudio.MarshalWriteFile(&c2, path.Join("/transferrer", "remote-example.json"))

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
		c2 = *(LoadRemote(r))
	}

	// make cfg.Pref() work properly:
	cfg.Get().URLPathPrefix = c2.URLPathPrefix

	host := fmt.Sprintf("%v:%v", c2.RemoteHost, c2.BindSocket)
	if c2.BindSocket == "" {
		host = c2.RemoteHost
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
		vals.Set("username", c2.AdminLogin)
		vals.Set("password", c2.Pass)
		vals.Set("token", lgn.FormToken())
		req, err := http.NewRequest(
			"POST",
			urlReq,
			strings.NewReader(vals.Encode()), // <-- URL-encoded payload
		)
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
			log.Fatalf(
				"Login response must contain '%v'\n%v\n%v\n\n%v",
				mustHave, urlReq, vals, string(respBytes),
			)
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

		dirFull := path.Join(c2.DownloadDir, c2.SurveyType, c2.WaveID)
		dirEmpty := path.Join(dirFull, "empty")

		qs := []*qst.QuestionnaireT{}
		dec := json.NewDecoder(rdr1)
		err = dec.Decode(&qs)
		if err != nil {
			log.Printf("Unmarshal %v", err)
			return
		}
		log.Printf("Unmarshalled %v questionnaires from responese stream", len(qs))

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

		//
		//
		// Process questionnaires
		for i, q := range qs {

			serverSideMD5 := q.MD5

			pthFull := path.Join(dirFull, q.UserID)
			err := q.Save1(pthFull)
			if err != nil {
				log.Printf("%3v: Error saving %v: %v", i, pthFull, err)
				continue
			}

			//
			if q.MD5 != serverSideMD5 {
				// log.Printf("%3v: MD5 does not match: %v\nwnt %v\ngot %v", i, pth2, md5BeforeSave, q.MD5)
				log.Printf("%3v: Server side and new client side MD5 hashes do not match %v - %v", i, q.Survey.String(), pthFull)
			}

			//
			//
			// Delete empty questionnaires and save them elsewhere
			pthEmpty := path.Join(dirEmpty, q.UserID+".json")
			err = cloudio.Delete(pthEmpty)
			if err != nil && !cloudio.IsNotExist(err) {
				log.Printf("%3v: Error removing previously empty %v - %v", i, pthEmpty, err)
			}
			realEntries, _, _ := q.Statistics()
			if realEntries == 0 {
				log.Printf("%3v: %v. No answers given, skipping, deleting, moving to %v.", i, pthFull, pthEmpty)
				err = cloudio.Delete(pthFull)
				if err != nil && !cloudio.IsNotExist(err) {
					log.Printf("%3v: Error removing empty %v - %v", i, pthFull, err)
				}

				err := q.Save1(pthEmpty)
				if err != nil {
					log.Printf("%3v: Error saving  to empty %v: %v", i, pthEmpty, err)
				}
				continue
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

		// log.Printf("%v keys superset; %v", len(allKeysSuperset), util.IndentedDump(allKeysSuperset))
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
		fn := fmt.Sprintf("/dl/online-responses-%v-%v.csv", c2.SurveyType, c2.WaveID)
		err = cloudio.WriteFile(fn, wtr, 0644)
		if err != nil {
			log.Printf("Could not write file: %v", err)
		}
		log.Printf("Regular finish. %v questionnaire(s) processed", len(qs))

	}

}
