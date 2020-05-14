package generators

import (
	myfmt "fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/generators/euref"
	"github.com/zew/go-questionnaire/generators/example"
	"github.com/zew/go-questionnaire/generators/flit"
	"github.com/zew/go-questionnaire/generators/fmt"
	"github.com/zew/go-questionnaire/generators/lt2020"
	"github.com/zew/go-questionnaire/generators/mul"
	"github.com/zew/go-questionnaire/generators/peu2018"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/util"
)

type genT func(params []qst.ParamT) (*qst.QuestionnaireT, error)

var gens = map[string]genT{
	"fmt":     fmt.Create,
	"flit":    flit.Create,
	"example": example.Create,
	"peu2018": peu2018.Create,
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

// SurveyGenerate generates a questionnaire for a bespoke survey
func SurveyGenerate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s := qst.NewSurvey("fmt") // type is modified later
	errStr := ""
	if r.Method == "POST" {
		frm := struct {
			Type     string       `json:"type"`
			Year     int          `json:"year"`
			Month    int          `json:"month"`
			Deadline string       `json:"deadline"`
			Params   []qst.ParamT `json:"params"`
			Submit   string       `json:"submit"`
		}{}
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
		fcCreate := func(desktopOrMobile string) {
			pth := filepath.Join(".", "templates", desktopOrMobile+q.Survey.Type+".css")
			if ok, _ := util.FileExists(pth); !ok {
				err := ioutil.WriteFile(pth, []byte{}, 0755)
				if err != nil {
					log.Fatalf("Could not create %v: %v", pth, err)
				}
				log.Printf("done creating file %v", pth)
			}
		}
		fcCreate("main_desktop_")
		fcCreate("main_mobile_")

	}

}
