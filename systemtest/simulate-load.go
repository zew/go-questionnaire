// Package systemtest contains system tests;
// ../main_test.go contains a detailed coverage test.
// However, test is run from the app dir one above.
// Working dir will be initially /go-questionnaire/systemtest,
// but we will step up one dir in the code below.
package systemtest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/util"
)

// We need this file and this empty func to avoid
// "no buildable Go source files" on travis
func main() {

}

func clientPageToServer(t *testing.T, q *qst.QuestionnaireT, idxPage int, urlMain string, sessCook *http.Cookie) {

	ctr.Reset()

	vals := url.Values{}
	for i1, p := range q.Pages {
		if i1 != idxPage {
			continue
		}
		vals.Set("token", lgn.FormToken())
		vals.Set("submitBtn", "next")
		for i2, grp := range p.Groups {
			for i3, inp := range grp.Inputs {
				if inp.IsLayout() {
					continue
				}
				vals.Set(inp.Name, ctr.IncrementStr())
				q.Pages[i1].Groups[i2].Inputs[i3].Response = ctr.GetLastStr()
				log.Printf("Input %12v set to value %2v ", inp.Name, ctr.GetLastStr())
			}
		}
	}
	t.Logf("POST requesting %v?\n%v", urlMain, strings.Replace(vals.Encode(), "submitBtn=next&token="+lgn.FormToken(), "...", -1))
	t1 := time.Now()
	resp, err := util.Request("POST", urlMain, vals, []*http.Cookie{sessCook})
	if err != nil {
		t.Errorf("error requesting %v: %v", urlMain, err)
	}
	t2 := time.Now()
	dur := t2.Sub(t1).Nanoseconds() / 1000 / 1000
	t.Logf("%9v ms request roundtrip", dur)
	t3 := time.Now().Add(-15 * time.Millisecond).Truncate(time.Second)

	q.Pages[idxPage].Finished = t3

	_ = resp
	// respStr := string(resp)

}

// FillQuestAndComparesServerResult loads a questionnaire from at template.
// It then fills the questionnaire page by page.
// It then compares the local copy of the questionnaire data
// with the "server" version.
func FillQuestAndComparesServerResult(t *testing.T, qSrc *qst.QuestionnaireT, urlMain string, sessCook *http.Cookie) {

	var clientQuest = &qst.QuestionnaireT{}
	var err error
	clientQuest, err = qst.Load1(qSrc.FilePath1(qSrc.Survey.Type + ".json")) // new from template
	if err != nil {
		t.Fatalf("Loading client questionnaire from file caused error: %v", err)
	}
	clientQuest.Survey = qSrc.Survey
	clientQuest.UserID = "systemtest"
	clientQuest.RemoteIP = "127.0.0.1:12345"
	clientQuest.CurrPage = 777 // hopeless to mimic every server side setting
	err = clientQuest.Validate()
	if err != nil {
		t.Fatalf("Client questionnaire validation caused error: %v", err)
	}
	t.Logf("Client questionnaire loaded from file; %v pages", len(clientQuest.Pages))

	for idx := range clientQuest.Pages {
		clientPageToServer(t, clientQuest, idx, urlMain, sessCook)
	}

	clientPth := filepath.Join(clientQuest.Survey.Type, clientQuest.Survey.WaveID(), "systemtest_src")
	clientQuest.Save1(clientQuest.FilePath1(clientPth))

	// Comparing client questionnaire to server questionnaire
	serverPth := strings.Replace(clientPth, "systemtest_src", "systemtest", -1)
	serverQuest, err := qst.Load1(clientQuest.FilePath1(serverPth)) // new from template
	if err != nil {
		t.Fatalf("Loading questionnaire from file caused error: %v", err)
	}
	equal, err := clientQuest.Compare(serverQuest)
	if !equal {
		t.Logf("Delete older versions of systemtest.json")
		t.Fatalf("Questionnaires are unequal: %v", err)
	}

}

// SimulateLoad logs in as 'systemtest'
// and performs some requests.
func SimulateLoad(t *testing.T, qSrc *qst.QuestionnaireT,
	loginURI, mobile string) {

	port := cfg.Get().BindSocket
	host := fmt.Sprintf("http://localhost:%v", port)

	urlLogin := host + loginURI
	t.Logf("url login  %v", urlLogin)

	urlMain := host + cfg.Pref()
	t.Logf("url main   %v", urlMain)

	//
	// Login and save session cookie
	var sessCook *http.Cookie
	{
		t.Logf(" ")
		t.Logf("Getting cookie")
		t.Logf("==================")
		urlReq := urlLogin

		vals := url.Values{}
		// vals.Set("username", "systemtest")
		vals.Set("mobile", mobile)
		req, err := http.NewRequest("GET", urlReq, bytes.NewBufferString(vals.Encode())) // <-- URL-encoded payload
		if err != nil {
			t.Errorf("error creating request for %v: %v", urlReq, err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		client := util.HttpClient()
		resp, err := client.Do(req)
		if err != nil {
			t.Errorf("error requesting cookie from %v: %v; %v", urlReq, err, resp)
		}
		defer resp.Body.Close()
		for _, v := range resp.Cookies() {
			if v.Name == "session" {
				sessCook = v
			}
		}

		t.Logf("Cookie is %+v \ngleaned from %v; server log *above* shows result", sessCook, req.URL)
		if sessCook == nil {
			t.Fatal("we need a session cookie to continue")
		}

		respBytes, _ := ioutil.ReadAll(resp.Body)
		mustNotHave := fmt.Sprintf("Login by hash failed")
		if strings.Contains(string(respBytes), mustNotHave) {
			t.Fatalf("Response must not contain '%v' \n\n%v", mustNotHave, string(respBytes))
		} else {
			t.Logf("Webpage reports: Login successful")
		}

	}

	ctr.Reset()

	if qSrc.Survey.Type == "fmt" {
		fmtSpecialTest(t, urlMain, sessCook)
	}

	FillQuestAndComparesServerResult(t, qSrc, urlMain, sessCook)

}
