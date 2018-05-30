// Package transferrer fetches completed questionaires
// from /transferrer-endpoint as JSON via http request.
package main

import (
	"bytes"
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

	"github.com/zew/go-questionaire/bootstrap"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/util"
)

const (
	remoteHost = "https://survey2.zew.de"
	user       = "pbu"
	pass       = "pb165205"
)

var downloadDir = filepath.Join(".", "responses", "downloaded")

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	bootstrap.Config()

	port := cfg.Get().BindSocket

	host := fmt.Sprintf("%v:%v", remoteHost, port)

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
		vals.Set("username", user)
		vals.Set("password", pass)
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

		log.Printf("Cookie is %+v \ngleaned from %v", sessCook, req.URL)
		if sessCook == nil {
			log.Printf("we need a session cookie to continue")
			return
		}

		respBytes, _ := ioutil.ReadAll(resp.Body)
		if strings.Contains(string(respBytes), "logged in as "+user) {
			log.Printf("Response must contain 'logged in as %v' \n\n%v", user, string(respBytes))
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

		waveID := qst.NewWaveID().String()

		vals := url.Values{}
		vals.Set("wave_id", waveID)
		log.Printf("POST requesting %v?%v", urlReq, vals.Encode())
		resp, err := util.Request("POST", urlReq, vals, []*http.Cookie{sessCook})
		if err != nil {
			log.Printf("error requesting %v: %v", urlReq, err)
			return
		}

		err = os.MkdirAll(filepath.Join(downloadDir, waveID), 0755)
		if err != nil {
			log.Printf("Could not create path 2 %v", filepath.Join(downloadDir, waveID))
			return
		}

		// respStr := string(resp)
		qs := []*qst.QuestionaireT{}
		err = json.Unmarshal(resp, &qs)
		if err != nil {
			log.Printf("Unmarshal %v", err)
			return
		}
		log.Printf("%v questionaire(s) unmarshalled", len(qs))

		for i, q := range qs {

			pth := filepath.Join(downloadDir, waveID, q.UserID)
			if !strings.HasSuffix(pth, ".json") {
				pth += ".json"
			}

			md5Want := q.MD5

			err := q.Save(pth)
			if err != nil {
				log.Printf("%3v: Error saving %v: %v", i, pth, err)
				return
			}

			if q.MD5 != md5Want {
				log.Printf("%3v: MD5 does not match: %v\nwnt %v\ngot %v", i, pth, md5Want, q.MD5)
				return
			}
		}

	}

}
