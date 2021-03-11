package systemtest

import (
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/util"
)

func fmtSpecialTest(t *testing.T, urlMain string, sessCook *http.Cookie) {
	//
	//
	// Post values and check the response
	{
		t.Logf(" ")
		t.Logf("Main view")
		t.Logf("==================")
		urlReq := urlMain

		{
			t.Logf("Goto first entry page ")
			t.Logf("==================")
			vals := url.Values{}
			vals.Set("token", lgn.FormToken())
			vals.Set("page", "1")
			t.Logf("POST requesting %v?%v", urlReq, vals.Encode())
			_, err := util.Request("POST", urlReq, vals, []*http.Cookie{sessCook})
			if err != nil {
				t.Errorf("error requesting %v: %v", urlReq, err)
			}
		}

		ctr.Reset() // first response should be '1' - since we check this later

		vals := url.Values{}
		vals.Set("y0_ez", ctr.IncrementStr()) // Don't forget to reset; otherwise depending on generate.FMT() the result is not deterministic
		vals.Set("y0_deu", ctr.IncrementStr())
		vals.Set("token", lgn.FormToken())
		t.Logf("POST requesting %v?%v", urlReq, vals.Encode())
		resp, err := util.Request("POST", urlReq, vals, []*http.Cookie{sessCook})
		if err != nil {
			t.Errorf("error requesting %v: %v", urlReq, err)
		}

		respStr := string(resp)

		// character distance -
		// must be large - since `name='y0_deu'` is found for val='1' ... val='2'
		// but only the second/third radio has value='2' checked="checked"
		scope := 400
		{
			needle1 := `name='y0_ez'`
			needle2 := `value='1' checked="checked"`
			pos1 := strings.Index(respStr, needle1)
			pos2 := strings.Index(respStr, needle2)
			t.Logf("Response should contain: %v ... %v \n%v %v => %v",
				needle1, needle2, pos1, pos2, pos2-pos1,
			)
			if pos1 < 1 || pos2 < 1 || (pos2-pos1) > scope {
				// ioutil.WriteFile("tmp-response-test.html", resp, 0777)
				t.Fatalf("fmt response contains 1 fail: %v %v  %v", pos1, pos2, pos2-pos1)
			}
		}
		{
			// needle := `name='y0_deu' id='y0_deu' title=' Deutschland' value='2' checked="checked"`
			needle1 := `name='y0_deu'`
			needle2 := `value='2' checked="checked"`
			pos1 := strings.Index(respStr, needle1)
			pos2 := strings.Index(respStr, needle2)
			t.Logf("Response should contain: %v ... %v \n%v %v => %v",
				needle1, needle2, pos1, pos2, pos2-pos1,
			)
			if pos1 < 1 || pos2 < 1 || (pos2-pos1) > scope {
				t.Fatalf("fmt response contains 2 fail: %v %v  %v", pos1, pos2, pos2-pos1)
			}
		}

	}

	{
		urlReq := urlMain
		t.Logf("Go back to page 0")
		t.Logf("==================")
		vals := url.Values{}
		vals.Set("token", lgn.FormToken())
		vals.Set("page", "0")
		t.Logf("POST requesting %v?%v", urlReq, vals.Encode())
		_, err := util.Request("POST", urlReq, vals, []*http.Cookie{sessCook})
		if err != nil {
			t.Errorf("error requesting %v: %v", urlReq, err)
		}
	}
}
