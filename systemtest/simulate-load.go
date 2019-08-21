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
	"os"
	"path"
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

func removeSystemtestJSON(t *testing.T) {
	err := filepath.Walk(filepath.Join(".", "responses"), func(path string, f os.FileInfo, err error) error {
		base := filepath.Base(path)
		if base == "systemtest.json" || base == "systemtest_src.json" {
			log.Printf("Removing %v", path)
			os.Remove(path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
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

// FillQuestAndComparesServerResult loads a questionnaire from a template.
// It then fills the questionnaire page by page.
// It then compares the local copy of the questionnaire data
// with the "server" version.
func FillQuestAndComparesServerResult(t *testing.T, qSrc *qst.QuestionnaireT, urlMain string, sessCook *http.Cookie) {

	var clientQuest = &qst.QuestionnaireT{}
	var err error
	pthBase := path.Join(qst.BasePath(), qSrc.Survey.Type+".json")
	clientQuest, err = qst.Load1(pthBase)
	if err != nil {
		t.Fatalf("Client questionnaire loading from file failed: %v", err)
	}
	clientQuest, err = clientQuest.Split()
	if err != nil {
		t.Fatalf("Client questionnaire splitting failed: %v", err)
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

	//
	// Doing load
	for idx := range clientQuest.Pages {
		clientPageToServer(t, clientQuest, idx, urlMain, sessCook)
	}

	//
	// Comparing client questionnaire to server questionnaire
	clientPth := strings.Replace(clientQuest.FilePath1(), "systemtest.json", "systemtest_src.json", -1)
	clientQuest.Split()
	clientQuest.Save1(clientPth)
	serverQuest, err := qst.Load1(clientQuest.FilePath1()) // new from template
	if err != nil {
		t.Fatalf("Loading questionnaire from file caused error: %v", err)
	}

	t.Logf("clientQuest and serverQuest saved: \n\t%v \n\t%v  ", clientPth, serverQuest.FilePath1())

	equal, err := clientQuest.Compare(serverQuest, false)
	if !equal {

		clientPthFail := strings.Replace(clientPth, ".json", "_FAIL.json", -1)
		os.Remove(clientPthFail)
		os.Rename(clientPth, clientPthFail)
		serverPthFail := strings.Replace(serverQuest.FilePath1(), ".json", "_FAIL.json", -1)
		os.Remove(serverPthFail)
		os.Rename(serverQuest.FilePath1(), serverPthFail)

		t.Logf("Delete older versions of systemtest.json")
		t.Fatalf("Questionnaires are unequal: %v", err)
	} else {
		t.Logf("clientQuest and serverQuest EQUAL")
		t.Logf("=================================")
		t.Logf("   ")
	}

}

// SimulateLoad logs in as 'systemtest'
// and performs some requests.
func SimulateLoad(t *testing.T, qSrc *qst.QuestionnaireT,
	loginURI, mobile string) {

	removeSystemtestJSON(t)
	defer removeSystemtestJSON(t)

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
			for k, v := range resp.Cookies() {
				t.Logf("\tfound cookie\t%v %v", k, v)
			}
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
