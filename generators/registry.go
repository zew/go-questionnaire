package generators

import (
	myfmt "fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/generators/fmt"
	"github.com/zew/go-questionaire/generators/min"
	"github.com/zew/go-questionaire/generators/peu2018"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/qst"
)

type genT func(params []qst.ParamT) (*qst.QuestionaireT, error)

var gens = map[string]genT{
	"fmt":     fmt.Create,
	"min":     min.Create,
	"peu2018": peu2018.Create,
}

// Get returns all questionaire generators
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

// SurveyGenerate generates a questionaire for a bespoke survey
func SurveyGenerate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if cfg.Get().IsProduction {
		l, isLoggedIn, err := lgn.LoggedInCheck(w, r)
		if err != nil {
			myfmt.Fprintf(w, "Login error %v", err)
			return
		}
		if !isLoggedIn {
			myfmt.Fprintf(w, "Not logged in")
			return
		}
		if !l.HasRole("admin") {
			myfmt.Fprintf(w, "admin login required")
			return
		}
	}

	s := qst.NewSurvey("fmt") // type is modified later
	if r.Method == "POST" {
		frm := struct {
			Type     string `json:"type"`
			Year     int    `json:"year"`
			Month    int    `json:"month"`
			Deadline string `json:"deadline"`

			Params []qst.ParamT `json:"params"`

			Submit string `json:"submit"`
		}{}
		dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
		err := dec.Decode(r.Form, &frm)
		if err != nil {
			myfmt.Fprint(w, err.Error()+"<br>\n")
		}
		// myfmt.Fprint(w, "<pre>"+util.IndentedDump(frm)+"</pre><br>\n")
		s.Type = frm.Type
		s.Year = frm.Year
		s.Month = time.Month(frm.Month)
		t, err := time.Parse("02.01.2006 15:04", frm.Deadline)
		if err != nil {
			myfmt.Fprint(w, err.Error()+"<br>\n")
		}
		wavePeriod := time.Date(s.Year, s.Month, 1, 0, 0, 0, 0, cfg.Get().Loc)
		if t.Sub(wavePeriod) > (30*24)*time.Hour ||
			t.Sub(wavePeriod) < -(10*24)*time.Hour {
			myfmt.Fprint(w, "Should the deadline not be close to the Year-Month?<br>\n")
		}

		s.Deadline = t

		s.Params = frm.Params
		// myfmt.Fprint(w, "<pre>"+util.IndentedDump(s)+"</pre><br>\n")

	}
	html := s.HTMLForm(get())
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

		fn := filepath.Join(qst.BasePath(), key+".json")
		err = q.Save1(fn)
		if err != nil {
			myfmt.Fprintf(w, "Error saving %v: %v<br>\n", fn, err)
			return
		}
		myfmt.Fprintf(w, "%v generated<br>\n", key)

	}

}
