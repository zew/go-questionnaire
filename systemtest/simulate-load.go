// Package systemtest fires up
// an instance of the app,
// imports json files,
// and checks the sums in the database;
//
// Package systemtest contains system tests;
// ../main_test.go contains a detailed coverage test.
//
// Test is executed from the *app* dir one above.
//
// Working dir will be initially /go-questionnaire/systemtest,
// but will be stepped up one dir in code below.
//
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
// Keep this empty func to avoid
// "no buildable Go source files" on travis
func main() {

}

func copyHelper(t *testing.T, pth string) {
	bts, err := cloudio.ReadFile(pth)
	if err != nil && !cloudio.IsNotExist(err) {
		t.Logf("copyHelper: Could not read file %v: %v", pth, err)
	}
	pth += ".inspection"
	rdr := bytes.NewReader(bts)
	err = cloudio.WriteFile(pth, rdr, 0777)
	if err != nil && !cloudio.IsNotExist(err) {
		t.Logf("copyHelper: Could not write file %v: %v", pth, err)
	}
}
func deleteHelper(t *testing.T, pth string) {
	err := cloudio.Delete(pth)
	if err != nil && !cloudio.IsNotExist(err) {
		t.Logf("deleteHelper: Could not delete %v: %v", pth, err)
	}
}

// copy files for inspection
func copySystemtestJSON(t *testing.T, q *qst.QuestionnaireT) {
	clientPth := strings.Replace(q.FilePath1(), "systemtest.json", "systemtest_src.json", -1)
	copyHelper(t, clientPth)
	copyHelper(t, q.FilePath1())
}

// delete answers of user systemtest
func removeSystemtestJSON(t *testing.T, q *qst.QuestionnaireT) {
	clientPth := strings.Replace(q.FilePath1(), "systemtest.json", "systemtest_src.json", -1)
	t.Logf("Deleting %v and 'systemtest_src.json' ", q.FilePath1())
	deleteHelper(t, clientPth)
	deleteHelper(t, q.FilePath1())
}

func clientPageToServer(t *testing.T, clQ *qst.QuestionnaireT, idxPage int,
	urlMain string, sessCook *http.Cookie) {

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

		// condense radio inputs
		// map distinct for weeding out multiple radios with same name
		radioVal := map[string]string{}
		for i2, grp := range p.Groups {
			for i3, inp := range grp.Inputs {
				if inp.IsLayout() {
					continue
				}
				val := ""
				if _, visitedBefore := radioVal[inp.Name]; visitedBefore {
					val = radioVal[inp.Name]
				} else {
					val = ctr.IncrementStr()
					radioVal[inp.Name] = val
				}

				radioVal[inp.Name] = val
				vals.Set(inp.Name, val)
				clQ.Pages[i1].Groups[i2].Inputs[i3].Response = val
				log.Printf("Input %12v set to value %2v ", inp.Name, val)
				// }
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
//
// qSrc is the basic survey template file for iterating pages and inputs
// clQ  is a fake  user response file - recording the data requested to the test server
//
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
	clQ.Attrs = map[string]string{}
	clQ.Attrs["user-attribute-1"] = "user-val-1"
	// qSrc.Attrs are empty...
	for k, v := range qSrc.Attrs {
		clQ.Attrs[k] = v // qSrc.Attrs is empty
	}
	clQ.UserID = "systemtest"
	clQ.RemoteIP = "127.0.0.1:12345"
	clQ.UserAgent = "golang client"
	clQ.LangCode = qSrc.LangCodes[0]
	/*
		//
		err = clQ.Validate()
		if err != nil {
			t.Fatalf("Client questionnaire validation caused error: %v", err)
		}
	*/
	t.Logf("Client questionnaire loaded from file; %v pages", len(clQ.Pages))

	// After UserID have been set => deletion possible
	removeSystemtestJSON(t, clQ)
	defer removeSystemtestJSON(t, clQ)
	defer copySystemtestJSON(t, clQ) // resolved inversely

	//
	// Doing load
	for idx := range clQ.Pages {
		clientPageToServer(t, clQ, idx, urlMain, sessCook)
	}
	clQ.CurrPage = len(clQ.Pages) - 2                  // last page does not get requested
	clQ.Pages[len(clQ.Pages)-1].Finished = time.Time{} // last page finishing time is zero value

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
func SimulateLoad(t *testing.T, q *qst.QuestionnaireT, loginURI, mobile string) {

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
		t.Logf("\nGetting cookie for %v mobile %v", q.Survey.String(), mobile)
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

	if q.Survey.Type == "fmt" {
		fmtSpecialTest(t, urlMain, sessCook)
	}

	FillQuestAndComparesServerResult(t, q, urlMain, sessCook)

}
