package generators

import (
	myfmt "fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/generators/fmt"
	"github.com/zew/go-questionaire/generators/min"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/qst"
)

type genT func() *qst.QuestionaireT

var gens = map[string]genT{
	"fmt": fmt.Create,
	"min": min.Create,
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

	s := qst.NewSurvey("fmt") // type is modified later
	if r.Method == "POST" {
		frm := struct {
			Type     string `json:"type"`
			Year     int    `json:"year"`
			Month    int    `json:"month"`
			Deadline string `json:"deadline"`
			Submit   string `json:"submit"`
		}{}
		dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
		err := dec.Decode(r.Form, &frm)
		if err != nil {
			io.WriteString(w, err.Error()+"<br>\n")
		}
		s.Type = frm.Type
		s.Year = frm.Year
		s.Month = time.Month(frm.Month)
		t, err := time.Parse("02.01.2006 15:04", frm.Deadline)
		if err != nil {
			io.WriteString(w, err.Error()+"<br>\n")
		}
		wavePeriod := time.Date(s.Year, s.Month, 1, 0, 0, 0, 0, cfg.Get().Loc)
		if t.Sub(wavePeriod) > (30*24)*time.Hour ||
			t.Sub(wavePeriod) < -(10*24)*time.Hour {
			io.WriteString(w, "Should the deadline not be close to the Year-Month?<br>\n")
		}

		s.Deadline = t
		// io.WriteString(w, util.IndentedDump(frm)+"<br>\n")
	}
	html := s.HTMLForm(get())
	io.WriteString(w, html)

	//
	for key, fnc := range Get() {

		if key != s.Type {
			continue
		}

		q := fnc()
		tr1, tr2 := q.Survey.Org, q.Survey.Name // save orig values
		q.Survey = s
		q.Survey.Org, q.Survey.Name = tr1, tr2

		fn := filepath.Join(qst.BasePath(), key+".json")
		err := q.Save1(fn)
		if err != nil {
			log.Fatalf("Error saving %v: %v", fn, err)
		}
		io.WriteString(w, myfmt.Sprintf("%v generated<br>\n", key))

	}

}
