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
	"strings"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
)

// RetrieveFromRemote requests the JSONified questionnaires
// from the survey server endpoint; decompresses the GZIPPed
// response and parses the bytes into a slice of questionnaires
func RetrieveFromRemote(cfgRem *RemoteConnConfigT) (
	[]*qst.QuestionnaireT,
	error,
) {

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
			return nil, fmt.Errorf("error creating request for %v: %w", urlReq, err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		resp, err := getClient().Do(req)
		if err != nil {
			log.Printf("error requesting cookie from %v: %v; %v", urlReq, err, resp)
		}

		if resp == nil {
			return nil, fmt.Errorf("response is nil - from %v", urlReq)
		}

		defer resp.Body.Close()

		for _, v := range resp.Cookies() {
			// if v.Name == "session" {
			if v.Name == "__Secure-go-quest-session" {
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
			return nil, fmt.Errorf("we need a session cookie to continue")
		}

		if !strings.Contains(string(respBytes), "Logged in as "+cfgRem.AdminLogin) {
			return nil, fmt.Errorf("Response must contain 'Logged in as %v' \n\n%v", cfgRem.AdminLogin, string(respBytes))
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
		return nil, fmt.Errorf("error creating request %v: %w", urlReq, err)
	}
	// strangely, the json *response* is empty, if we omit this:
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	for _, v := range []*http.Cookie{sessCook} {
		req.AddCookie(v)
	}

	resp, err := getClient().Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting %v: %w", urlReq, err)
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
				return nil, fmt.Errorf("could not read all response %w", err)
			}
			log.Printf("response is %v bytes", len(bts))
			if len(bts) < 15000 {
				log.Printf("response is %s", bts)
			}
			btsRdr := bytes.NewReader(bts)
			rdr1, err = gzip.NewReader(btsRdr)
			if err != nil {
				return nil, fmt.Errorf("could not read the response as gzip #1: %w", err)
			}
		}

		rdr1, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("could not read the response as gzip #2: %w", err)
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
		return nil, fmt.Errorf("bad response %v - %q ", resp.StatusCode, resp.Status)
	}

	dec := json.NewDecoder(rdr1)
	qs := []*qst.QuestionnaireT{}
	err = dec.Decode(&qs)
	if err != nil {
		log.Printf("Unmarshal failed: %v", err)
		fn := "tmp-transferrer-endpoint-response-error.html"
		bts, _ := io.ReadAll(rdr1)
		os.WriteFile(fn, bts, 0777)
		return nil, fmt.Errorf("Response written to %v", fn)

	}
	log.Printf("Unmarshalled %v questionnaires from responese stream", len(qs))
	log.Printf("====================================================")

	return qs, nil

}
