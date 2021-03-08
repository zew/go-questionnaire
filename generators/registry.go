package generators

import (
	"bytes"
	myfmt "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"path"
	"strings"
	"time"

	"github.com/go-playground/form"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/generators/fmt"
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/tpl"
)

type genT func(params []qst.ParamT) (*qst.QuestionnaireT, error)

var gens = map[string]genT{
	"fmt": fmt.Create,
	"pat": pat.Create,
	// "flit":    flit.Create,
	// "example": example.Create,
	// "peu2018": peu2018.Create,
	// "mul":     mul.Create,
	// "euref":   euref.Create,
	// "lt2020":  lt2020.Create,
}

// Get returns all questionnaire generators
func Get() map[string]genT {
	return gens
}

func get() []string {
	ret := []string{}
	gens := Get()
	for key := range gens {
		ret = append(ret, key)
	}
	return ret
}

type frmT struct {
	Type     string `json:"type"`
	Year     int    `json:"year"`
	Month    int    `json:"month"`
	Deadline string `json:"deadline"`
	// Params    []qst.ParamT `json:"params"`
	ParamKeys []string `json:"param_keys,omitempty"`
	ParamVals []string `json:"param_vals,omitempty"`
	Submit    string   `json:"submit,omitempty"`
}

// GenerateQuestionnaireTemplates generates a questionnaire for a bespoke survey
func GenerateQuestionnaireTemplates(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s := qst.NewSurvey("fmt") // type is modified later
	errStr := ""
	if r.Method == "POST" {
		// myfmt.Fprint(w, "is POST<br>\n")
		frm := frmT{}
		dec := form.NewDecoder()
		dec.SetTagName("json") // recognizes and ignores ,omitempty
		err := dec.Decode(&frm, r.Form)
		if err != nil {
			errStr += myfmt.Sprint(err.Error() + "<br>\n")
		}

		// myfmt.Fprint(w, "<pre>"+util.IndentedDump(frm)+"</pre><br>\n")

		s.Type = frm.Type
		s.Year = frm.Year
		s.Month = time.Month(frm.Month)
		t, err := time.Parse("02.01.2006 15:04", frm.Deadline)
		if err != nil {
			errStr += myfmt.Sprint(err.Error() + "<br>\n")
		}
		wavePeriod := time.Date(s.Year, s.Month, 1, 0, 0, 0, 0, cfg.Get().Loc)
		if t.Sub(wavePeriod) > (30*24)*time.Hour ||
			t.Sub(wavePeriod) < -(10*24)*time.Hour {
			errStr += myfmt.Sprint("Should the deadline not be close to the Year-Month?<br>\n")
		}

		s.Deadline = t

		newParams := []qst.ParamT{}
		for i := 0; i < len(frm.ParamKeys); i++ {
			p := qst.ParamT{}
			p.Name = frm.ParamKeys[i]
			p.Val = frm.ParamVals[i]
			newParams = append(newParams, p)
		}
		s.Params = newParams
		// myfmt.Fprint(w, "<pre>"+util.IndentedDump(s)+"</pre><br>\n")

	}

	html := s.HTMLForm(get(), errStr)
	myfmt.Fprintf(w, html)
	myfmt.Fprintf(w, "<br>")
	//
	for key, fnc := range Get() {

		if key != s.Type {
			continue
		}

		q, err := fnc(s.Params)
		if err != nil {
			myfmt.Fprintf(w, "Error creating %v: %v<br>\n", key, err)
			return
		}

		tr1, tr2 := q.Survey.Org, q.Survey.Name // save orig values
		q.Survey = s
		q.Survey.Org, q.Survey.Name = tr1, tr2

		fn := path.Join(qst.BasePath(), key+".json")
		err = q.Save1(fn)
		if err != nil {
			myfmt.Fprintf(w, "Error saving %v: %v<br>\n", fn, err)
			return
		}
		myfmt.Fprintf(w, "%v generated<br>\n", key)

		//
		// create empty styles-quest-[surveytype].css"
		// if it does not yet exist
		fcCreate := func(desktopOrMobile string) (bool, error) {
			pth := path.Join(".", "templates", desktopOrMobile+q.Survey.Type+".css")
			_, err := cloudio.ReadFile(pth)
			if err != nil {
				if cloudio.IsNotExist(err) {
					rdr := &bytes.Buffer{}
					err := cloudio.WriteFile(pth, rdr, 0755)
					if err != nil {
						return false, myfmt.Errorf("Could not create %v: %v <br>\n", pth, err)
					}
					myfmt.Fprintf(w, "Done creating template %v<br>\n", pth)
					return true, nil
				} else {
					return false, myfmt.Errorf("Other error while checking for %v: %v <br>\n", pth, err)
				}
			}
			return false, nil
		}

		// add to parsed templates
		for _, bt := range []string{"styles-quest-"} {
			ok, err := fcCreate(bt)
			if err != nil {
				myfmt.Fprintf(w, "Could not generate template %v for %v<br>\n", bt, err)
				continue
			}
			if ok {
				// parse new and previous templates
				dummyReq, err := http.NewRequest("GET", "", nil)
				if err != nil {
					log.Fatalf("failed to create request for pre-loading assets %v", err)
				}
				respRec := httptest.NewRecorder()
				tpl.TemplatesPreparse(respRec, dummyReq)
				log.Printf("\n%v", respRec.Body.String())

			}
		}

	}
}

// GenerateLandtagsVariations creates 16 questionnaire templates
func GenerateLandtagsVariations(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	key := "lt2020"

	for i := 0; i < 32; i++ {

		form := url.Values{}
		form.Add("type", key)
		form.Add("year", "2020")
		form.Add("month", "5")
		form.Add("deadline", "01.01.2030 00:00")
		form.Add("params[0].name", "varianten")
		form.Add("params[0].val", myfmt.Sprintf("%04b", i%16))
		form.Add("params[1].name", "aboveOrBelowMedian")
		if i < 16 {
			form.Add("params[1].val", "besseren")
		} else {
			form.Add("params[1].val", "schlechteren")
		}
		form.Add("Submit", "any")
		// myfmt.Fprint(w, "<pre>"+util.IndentedDump(form)+"</pre><br>\n")

		var resp *http.Response
		var err error

		if true {
			req, err := http.NewRequest(
				"POST",
				"https://localhost:8083"+cfg.PrefTS("generate-questionnaire-templates"),
				bytes.NewBufferString(form.Encode()),
			)
			if err != nil {
				myfmt.Fprintf(w, "Request creation error %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
			client := http.DefaultClient
			resp, err = client.Do(req)
			if err != nil {
				myfmt.Fprintf(w, "Request execution error %v", err)
				return
			}
		} else {
			resp, err = http.PostForm(
				"https://localhost:8083"+cfg.PrefTS("generate-questionnaire-templates"),
				form,
			)
			if err != nil {
				myfmt.Fprintf(w, "Request execution error %v", err)
				return
			}
		}

		defer resp.Body.Close()
		respBts, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			myfmt.Fprintf(w, "Error reading response body %v", err)
			return
		}

		myfmt.Fprintf(w, "%s\n", respBts)

		fn := path.Join(qst.BasePath(), key+".json")
		qst, err := qst.Load1(fn)
		if err != nil {
			myfmt.Fprintf(w, "Error re-loading qst for %v: %v", fn, err)
			return
		}

		fnNew := strings.ReplaceAll(fn, ".json", myfmt.Sprintf("-%02v.json", i))
		qst.Save1(fnNew)

		myfmt.Fprintf(w, "Iter %v - stop; resp status %v<br><br>\n", i, resp.Status)
		myfmt.Fprintf(w, "<hr>\n")

	}

}
