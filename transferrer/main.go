// Package transferrer fetches completed questionnaires
// from /transferrer-endpoint as JSON via http request.
package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/bootstrap"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/util"
)

// AddtlConfigT is on top of
// of the remote config
type AddtlConfigT struct {
	RemoteHost string

	AdminLogin string // Some admin account of the remote machine
	Pass       string

	SurveyType string
	WaveID     string

	DownloadDir string
}

// Example returns a minimal configuration, to be extended or adapted
func Example() AddtlConfigT {
	r := AddtlConfigT{}
	r.RemoteHost = "https://survey2.zew.de"
	r.RemoteHost = "https://www.peu2018.eu"

	r.AdminLogin = "login"
	r.Pass = "secret"

	r.SurveyType = "fmt"
	r.SurveyType = "peu2018"

	r.WaveID = qst.NewSurvey(r.SurveyType).WaveID()
	r.WaveID = "2018-08"

	r.DownloadDir = filepath.Join(qst.BasePath(), "downloaded")

	return r
}

// Save writes a JSON file
func (r *AddtlConfigT) Save(fn string) {
	byts, err := json.MarshalIndent(r, " ", "\t")
	util.BubbleUp(err)
	err = ioutil.WriteFile(fn, byts, 0644)
	util.BubbleUp(err)
}

// Load loads a JSON file
func Load(fn string) (r *AddtlConfigT) {
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
	tempCfg := AddtlConfigT{}
	err = decoder.Decode(&tempCfg)
	if err != nil {
		log.Fatal(err)
	}
	return &tempCfg
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	bootstrap.Config()

	var c2 AddtlConfigT

	//
	c2 = Example()
	c2.Save("remote-additional-example.json")

	c2 = *(Load("remote-additional.json"))
	host := fmt.Sprintf("%v:%v", c2.RemoteHost, cfg.Get().BindSocket)

	defer func() {
		log.Printf("  ")
		log.Printf("  ")
		log.Printf("================")
		log.Printf("Login         via   %v%v", host, "/survey/login-primitive")
		log.Printf("Check results via   %v%v", host, "/survey/transferrer-endpoint?wave_id=2018-06")
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
		client := util.HttpClient()
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("error requesting cookie from %v: %v; %v", urlReq, err, resp)
			return
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
			log.Fatalf("Login response must contain '%v'\n\n%v", mustHave, string(respBytes))
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
		log.Printf("POST requesting %v?%v", urlReq, vals.Encode())
		resp, err := util.Request("POST", urlReq, vals, []*http.Cookie{sessCook})
		if err != nil {
			log.Printf("error requesting %v: %v", urlReq, err)
			return
		}

		dir := filepath.Join(c2.DownloadDir, c2.SurveyType, c2.WaveID)
		dirEmpty := filepath.Join(dir, "empty")
		err = os.MkdirAll(dirEmpty, 0755)
		if err != nil {
			log.Printf("Could not create path 2 %v", dir)
			return
		}

		// respStr := string(resp)
		qs := []*qst.QuestionnaireT{}
		err = json.Unmarshal(resp, &qs)
		if err != nil {
			log.Printf("Unmarshal %v", err)
			return
		}
		log.Printf("%v questionnaire(s) unmarshalled; %10.3f", len(qs), float64(len(resp)/(1<<10))/(1<<10))

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

			// "deadline": "2018-10-31T23:59:00Z"
			tToBeCorrected := time.Date(2018, 10, 31, 24, 59, 0, 0, cfg.Get().Loc)
			tInstead := time.Date(2019, 02, 28, 24, 59, 0, 0, cfg.Get().Loc)
			if q.Survey.Deadline.Equal(tToBeCorrected) {
				log.Printf("Correction needed: %v", q.Survey.Deadline)
				q.Survey.Deadline = tInstead
				pth2 := filepath.Join(dir, q.UserID)
				err := q.Save1(pth2)
				if err != nil {
					log.Printf("%3v: Error saving %v: %v", i, pth2, err)
					continue
				}
				continue
			} else {
				// log.Printf("Difference       : %v", q.Survey.Deadline.Sub(tToBeCorrected))
			}

			pth2 := filepath.Join(dir, q.UserID)
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
				pth3 := filepath.Join(dirEmpty, q.UserID+".json")
				err := os.Rename(pth2a, pth3)
				if err != nil {
					log.Printf("%3v: Error moving %v to %v - %v", i, pth2a, pth3, err)
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
		log.Printf("Regular finish. %v questionnaire(s) processed; %.3f MB", len(qs), float64(len(resp)/(1<<10))/(1<<10))

	}

}
