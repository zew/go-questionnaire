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
	"io"
	"net/http"
	"net/url"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/util"
)

// We need this file and this empty func to avoid
// "no buildable Go source files" on travis
// Keep this empty func to avoid
// "no buildable Go source files" on travis
func main() {

}

//                 type    page    inpName  value
var presets = map[string]map[int]map[string]string{
	"pat1": {
		3: {
			"q_found_compr_a": "est_2",
			"q_found_compr_b": "est_c",
		},
		8: {
			"q_tpref_compr_a": "3",
			"q_tpref_compr_b": "7",
		},
		9: {
			"q2_seq1_row1_rad": "1",
			"q2_seq1_row2_rad": "1",
			"q2_seq1_row3_rad": "1",

			"q2_seq2_row1_rad": "1",
			"q2_seq2_row2_rad": "2",
			"q2_seq2_row3_rad": "2",
		},
		10: {
			"q4a_opt1": "3",
			"q4a_opt2": "3",
			"q4a_opt3": "4",

			"q4b_opt1": "3",
			"q4b_opt2": "3",
			"q4b_opt3": "4",
		},
		13: {
			"q17": "citizenshipyes",
		},
	},
	"pat2": {
		1: {
			"q17": "citizenshipyes",
		},
		7: {
			"q_found_compr_a": "est_2",
			"q_found_compr_b": "est_c",
		},
		16: {
			"part2_q1_q1": "3",
			"part2_q1_q2": "3",
			"part2_q1_q3": "4",

			"part2_q2_q1": "3",
			"part2_q2_q2": "3",
			"part2_q2_q3": "4",

			"part2_q3_q1": "3",
			"part2_q3_q2": "3",
			"part2_q3_q3": "4",
		},
		17: {
			"part2_q4_q1": "3",
			"part2_q4_q2": "3",
			"part2_q4_q3": "4",

			"part2_q5_q1": "3",
			"part2_q5_q2": "3",
			"part2_q5_q3": "4",

			"part2_q6_q1": "3",
			"part2_q6_q2": "3",
			"part2_q6_q3": "4",
		},
	},

	"pat3": {
		1: {
			"q17": "citizenshipyes",
		},
		5: {
			"q_tpref_compr_a": "3",
			"q_tpref_compr_b": "7",
		},
		6: {
			"q2_seq1_row1_rad": "1",
			"q2_seq1_row2_rad": "1",
			"q2_seq1_row3_rad": "1",

			"q2_seq2_row1_rad": "1",
			"q2_seq2_row2_rad": "2",
			"q2_seq2_row3_rad": "2",
		},
		7: {
			"q4a_opt1": "3",
			"q4a_opt2": "3",
			"q4a_opt3": "4",

			"q4b_opt1": "3",
			"q4b_opt2": "3",
			"q4b_opt3": "4",
		},
		13: {
			"pop3_part2_q1_1": "7",
			"pop3_part2_q1_2": "1",
			"pop3_part2_q1_3": "1",
			"pop3_part2_q1_4": "1",

			"pop3_part2_q2_1": "0",
			"pop3_part2_q2_2": "9",
			"pop3_part2_q2_3": "1",
			"pop3_part2_q2_4": "0",

			"pop3_part2_q3_1": "0",
			"pop3_part2_q3_2": "9",
			"pop3_part2_q3_3": "1",
			"pop3_part2_q3_4": "0",
		},

		14: {
			"pop3_part2_q4_1": "7",
			"pop3_part2_q4_2": "1",
			"pop3_part2_q4_3": "1",
			"pop3_part2_q4_4": "1",

			"pop3_part2_q5_1": "0",
			"pop3_part2_q5_2": "9",
			"pop3_part2_q5_3": "1",
			"pop3_part2_q5_4": "0",

			"pop3_part2_q6_1": "0",
			"pop3_part2_q6_2": "9",
			"pop3_part2_q6_3": "1",
			"pop3_part2_q6_4": "0",
		},
	},
	"biii": {
		0: {
			"q1_role": "investor",
			"q2":      "private_investor",
			"q3":      "esg",
		},
		1: {
			"q4":  "now",
			"q4a": "10yrs",
		},
		2: {
			"q5":              "all",
			"q6_impact":       "1000",
			"q6_other":        "2000",
			"q6_conventional": "111111000",
		},
	},
}

var skipPages = map[string]map[int]interface{}{
	"biii": {
		5: nil,
	},
}

func getPreset(qType string, pageIdx int, inpName string) (string, bool) {
	if _, ok1 := presets[qType]; ok1 {
		if _, ok2 := presets[qType][pageIdx]; ok2 {
			if vl, ok3 := presets[qType][pageIdx][inpName]; ok3 {
				return vl, true
			}
		}
	}
	return "", false
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

// copySystemtestJSONForInspection copies files for inspection
func copySystemtestJSONForInspection(t *testing.T, q *qst.QuestionnaireT) {
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

// clientPageToServer doing POST requests to a single page of the test server
func clientPageToServer(t *testing.T, clQ *qst.QuestionnaireT, idxPage int,
	urlMain string, sessCook *http.Cookie) {

	ctr.Reset()

	// values from ctr.IncrementStr() are stored into clQ (client questionnaire)
	// and POSTed to the server

	// loop fills clQ and concatenates the URL params for the POST request below
	lpCntr := 0
	vals := url.Values{}
	for i1, p := range clQ.Pages {

		if i1 != idxPage {
			continue
		}

		clQ.Pages[idxPage].Finished = time.Now()

		vals.Set("token", lgn.FormToken())
		vals.Set("submitBtn", "next")

		// condense radio inputs
		// map distinct for weeding out multiple radios with same name
		condensed := map[string]string{}

		for i2, grp := range p.Groups {
			for i3, inp := range grp.Inputs {
				if inp.IsLayout() {
					continue
				}
				val, ok := "", false
				if _, visitedBefore := condensed[inp.Name]; visitedBefore {
					val = condensed[inp.Name]
				} else {
					val, ok = getPreset(clQ.Survey.Type, i1, inp.Name)
					if !ok {
						val = ctr.IncrementStr()
						if val == "9" { // preventing inRange10
							ctr.Reset()
						}
					}
					condensed[inp.Name] = val
				}

				// condensed[inp.Name] = val
				vals.Set(inp.Name, val)
				clQ.Pages[i1].Groups[i2].Inputs[i3].Response = val
				lpCntr++
				if lpCntr < 3 || lpCntr%10 == 0 {
					t.Logf("Input %12v set to value %2v ", inp.Name, val)
				}
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
// clQ  is a fake user response file - recording the data requested to the test server - "client questionnaire"
//
func FillQuestAndComparesServerResult(t *testing.T, qSrc *qst.QuestionnaireT, urlMain string, sessCook *http.Cookie) {

	var clQ = &qst.QuestionnaireT{} // see func description
	var err error

	pthBase := path.Join(qst.BasePath(), qSrc.Survey.Filename()+".json")

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
	defer copySystemtestJSONForInspection(t, clQ) // resolved inversely

	//
	// Doing load
	for idx := range clQ.Pages {
		if _, ok := skipPages[clQ.Survey.Type][idx]; ok {
			t.Logf("	Survey %v - skipping page %v", clQ.Survey.Type, idx)
			continue
		}
		clientPageToServer(t, clQ, idx, urlMain, sessCook)
	}
	clQ.CurrPage = len(clQ.Pages) - 2 // last page does not get requested

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

		respBytes, _ := io.ReadAll(resp.Body)
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
