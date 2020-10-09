package generators

import (
	"bytes"
	myfmt "fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/generators/euref"
	"github.com/zew/go-questionnaire/generators/example"
	"github.com/zew/go-questionnaire/generators/flit"
	"github.com/zew/go-questionnaire/generators/flit2"
	"github.com/zew/go-questionnaire/generators/fmt"
	"github.com/zew/go-questionnaire/generators/lt2020"
	"github.com/zew/go-questionnaire/generators/mul"
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/generators/peu2018"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/util"
)

type genT func(params []qst.ParamT) (*qst.QuestionnaireT, error)

var gens = map[string]genT{
	"fmt":     fmt.Create,
	"flit":    flit.Create,
	"flit2":   flit2.Create,
	"example": example.Create,
	"peu2018": peu2018.Create,
	"pat":     pat.Create,
	"mul":     mul.Create,
	"euref":   euref.Create,
	"lt2020":  lt2020.Create,
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
	Type     string       `json:"type"`
	Year     int          `json:"year"`
	Month    int          `json:"month"`
	Deadline string       `json:"deadline"`
	Params   []qst.ParamT `json:"params"`
	Submit   string       `json:"submit"`
}

// GenerateQuestionnaireTemplates generates a questionnaire for a bespoke survey
func GenerateQuestionnaireTemplates(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s := qst.NewSurvey("fmt") // type is modified later
	errStr := ""
	if r.Method == "POST" {
		// myfmt.Fprint(w, "is POST<br>\n")
		frm := frmT{}
		dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
		err := dec.Decode(r.Form, &frm)
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
		s.Params = frm.Params
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
		// create empty main_desktop_[surveytype].css"
		// create empty main_mobile_[surveytype].css"
		// if it does not yet exist
		fcCreate := func(desktopOrMobile string) (bool, error) {
			pth := filepath.Join(".", "templates", desktopOrMobile+q.Survey.Type+".css")
			if ok, _ := util.FileExists(pth); !ok {
				err := ioutil.WriteFile(pth, []byte{}, 0755)
				if err != nil {
					return false, myfmt.Errorf("Could not create %v: %v", pth, err)
				}
				myfmt.Fprintf(w, "done creating template %v", pth)
				return true, nil
			}
			return false, nil
		}

		for _, bt := range []string{"main_desktop_", "main_mobile_"} {
			ok, err := fcCreate(bt)
			if err != nil {
				myfmt.Fprintf(w, "Could not generated template %v for %v<br>\n", bt, err)
				continue
			}
			if ok {
				tpl.ParseH(w, r)
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
