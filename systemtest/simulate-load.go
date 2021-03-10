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
	"path"
	"strings"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/util"
)

// We need this file and this empty func to avoid
// "no buildable Go source files" on travis
func main() {

}

func removeSystemtestJSON(t *testing.T, qSrc *qst.QuestionnaireT) {

	clientPth := strings.Replace(qSrc.FilePath1(), "systemtest.json", "systemtest_src.json", -1)

	t.Logf("Deleting %v and 'systemtest_src.json' ", qSrc.FilePath1())

	err := cloudio.Delete(clientPth)
	if err != nil && !cloudio.IsNotExist(err) {
		t.Logf("Could not delete %v: %v", clientPth, err)
	}
	err = cloudio.Delete(qSrc.FilePath1())
	if err != nil && !cloudio.IsNotExist(err) {
		t.Logf("Could not delete %v: %v", qSrc.FilePath1(), err)
	}

}

func clientPageToServer(t *testing.T, clQ *qst.QuestionnaireT, idxPage int, urlMain string, sessCook *http.Cookie) {

	ctr.Reset()

	// values from ctr.IncrementStr() are stored into clQ (client questionnaire)
	// and POSTed to the server

	vals := url.Values{}
	for i1, p := range clQ.Pages {
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
				clQ.Pages[i1].Groups[i2].Inputs[i3].Response = ctr.GetLastStr()
				log.Printf("Input %12v set to value %2v ", inp.Name, ctr.GetLastStr())
			}
		}
	}
	tmp := strings.Replace(vals.Encode(), "submitBtn=next&token="+lgn.FormToken(), "...", -1)
	t.Logf("POST requesting %v?%v", urlMain, util.UpTo(tmp, 60))
	t1 := time.Now()
	resp, err := util.Request("POST", urlMain, vals, []*http.Cookie{sessCook})
	if err != nil {
		t.Errorf("error requesting %v: %v", urlMain, err)
	}
	t2 := time.Now()
	dur := t2.Sub(t1).Nanoseconds() / 1000 / 1000
	t.Logf("request roundtrip %9v ms", dur)

	t3 := time.Now().Add(-15 * time.Millisecond).Truncate(time.Second)
	clQ.Pages[idxPage].Finished = t3

	_ = resp
	// respStr := string(resp)

}

// FillQuestAndComparesServerResult loads a questionnaire from a template.
// It then fills the questionnaire page by page.
// It then compares the local copy of the questionnaire data
// with the "server" version.
func FillQuestAndComparesServerResult(t *testing.T, qSrc *qst.QuestionnaireT, urlMain string, sessCook *http.Cookie) {

	var clQ = &qst.QuestionnaireT{}
	var err error
	pthBase := path.Join(qst.BasePath(), qSrc.Survey.Type+".json")

	// creating client quest from scratch
	aBase, err := qst.Load1(pthBase)
	if err != nil {
		t.Fatalf("base questionnaire loading from file failed: %v", err)
	}
	err = aBase.Validate()
	if err != nil {
		t.Fatalf("base questionnaire validation caused error: %v", err)
	}
	clQ, err = aBase.Split()
	if err != nil {
		t.Fatalf("Client questionnaire splitting failed: %v", err)
	}
	clQ.Survey = qSrc.Survey
	for i := 0; i < len(clQ.Pages); i++ {
		clQ.Pages[i].Finished = time.Now()
	}
	clQ.UserID = "systemtest"
	clQ.RemoteIP = "127.0.0.1:12345"
	clQ.CurrPage = 777 // hopeless to mimic every server side setting
	/*
		err = clQ.Validate()
		if err != nil {
			t.Fatalf("Client questionnaire validation caused error: %v", err)
		}
	*/
	t.Logf("Client questionnaire loaded from file; %v pages", len(clQ.Pages))

	// After UserID have been set => deletion possible
	removeSystemtestJSON(t, clQ)
	defer removeSystemtestJSON(t, clQ)

	//
	// Doing load
	for idx := range clQ.Pages {
		clientPageToServer(t, clQ, idx, urlMain, sessCook)
	}

	clientPth := strings.Replace(clQ.FilePath1(), "systemtest.json", "systemtest_src.json", -1)

	//
	// Comparing client questionnaire to server questionnaire
	clQ.Split()
	clQ.Save1(clientPth)
	srvQ, err := qst.Load1(clQ.FilePath1()) // new from template
	if err != nil {
		t.Fatalf("Loading questionnaire from file caused error: %v", err)
	}

	t.Logf("clientQst saved to %v  - server %v", clientPth, srvQ.FilePath1())

	equal, err := clQ.Compare(srvQ, false)
	if !equal {
		t.Fatalf(
			"%22s - questionnaires are unequal: %v \n\t%v\n\t%v",
			clQ.Survey.String(), err, clQ.FilePath1(), srvQ.FilePath1(),
		)
	} else {
		t.Logf("==================================================")
		t.Logf("clientQst and srvQst are EQUAL for %v", clQ.Survey.String())
	}

}

// SimulateLoad logs in as 'systemtest'
// and performs some requests.
func SimulateLoad(t *testing.T, qSrc *qst.QuestionnaireT, loginURI, mobile string) {

	port := cfg.Get().BindSocket
	host := fmt.Sprintf("http://localhost:%v", port)

	urlLogin := host + loginURI
	// t.Logf("url login  %v", urlLogin)

	urlMain := host + cfg.Pref()
	// t.Logf("url main   %v", urlMain)

	//
	// Login and save session cookie
	var sessCook *http.Cookie
	{
		t.Logf("\nGetting cookie for %v mobile %v", qSrc.Survey.String(), mobile)
		t.Logf("================================")
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
