package joined

import (
	"io"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionaire/generators"
	"github.com/zew/go-questionaire/qst"
)

func init() {
	qst.Generators = map[string]interface{}{}
	gens := generators.Get()
	for key := range gens {
		qst.Generators[key] = nil
	}
	qst.Generators["more"] = nil
}

// SurveyGenerate generates a questionaire for a bespoke survey
func SurveyGenerate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s := qst.NewSurvey("fmt")
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
		s.Deadline = t
		// io.WriteString(w, util.IndentedDump(frm)+"<br>\n")
	}
	html := s.HTMLForm()
	io.WriteString(w, html)

	//
	for key, fnc := range generators.Get() {

		q := fnc()
		q.Survey = s

		fn := filepath.Join(qst.BasePath(), key+".json")
		err := q.Save1(fn)
		if err != nil {
			log.Fatalf("Error saving %v: %v", fn, err)
		}

	}

}
