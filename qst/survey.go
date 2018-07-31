package qst

import (
	"fmt"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/trl"
)

// ParamT contains changing parameters to a questionaire
type ParamT struct {
	Name string `json:"name,omitempty"` // i.e. main_refinance_rate_ecb
	Val  string `json:"val,omitempty"`  // i.e. "01.02.2018: 3.2%"
	// Challenge string `json:"challenge,omitempty"` // i.e.  Set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%` as in `main refinance rate of the ECB (01.02.2018: 3.2%)`"
}

// Param returns the value of a questionaire param
func (s *surveyT) Param(name string) (string, error) {
	for _, p := range s.Params {
		if p.Name == name {
			return p.Val, nil
		}
	}
	return "?", fmt.Errorf("Param %q not found", name)
}

// surveyT stores the interval components of a questionaire wave.
// For quarterly intervals, it needs to be extended
type surveyT struct {
	Type string `json:"type,omitempty"` // The type identifier, i.e. "fmt" or "cep"
	Org  trl.S  `json:"org,omitempty"`  // organization, i.e. Unicef
	Name trl.S  `json:"name,omitempty"` // full name, i.e. programming languages survey

	Year  int        `json:"year,omitempty"`  // The wave year
	Month time.Month `json:"month,omitempty"` // The wave month, 1-based, int

	Deadline time.Time `json:"deadline,omitempty"` // No more responses accepted

	Params []ParamT `json:"params,omitempty"` // I.e. NASDAQ at the time begin of the wave - being used in the wording of the questions
}

// NewSurvey returns a survey based on current time
func NewSurvey(tp string) surveyT {

	if tp == "" {
		panic("survey must have a type")
	}
	s := surveyT{Type: tp}

	t := time.Now()
	if t.Day() > 20 {
		t = t.AddDate(0, 1, 0)
	}
	s.Year = t.Year()
	s.Month = t.Month()

	s.Deadline = time.Date(s.Year, s.Month, 28, 23, 59, 59, 0, cfg.Get().Loc) // This is arbitrary

	s.Params = []ParamT{}

	return s
}

// String is the default identifier
func (s surveyT) String() string {
	return s.Type + "-" + s.WaveID()
}

// WaveID returns the year-month in standard format yyyy-mm
func (s surveyT) WaveID() string {
	// Notice the month +1
	// It is necessary, even though the spec says 'January = 1'
	t := time.Date(s.Year, s.Month+1, 0, 0, 0, 0, 0, cfg.Get().Loc)
	return t.Format("2006-01")
}

// Label is a pretty identifier
func (s surveyT) Label() string {
	if s.Year == 0 {
		return ""
	}
	// Notice the month +1
	// It is necessary, even though the spec says 'January = 1'
	t := time.Date(s.Year, s.Month+1, 0, 0, 0, 0, 0, cfg.Get().Loc)
	// return t.Format("January 2006") // yields English month names.
	return t.Format("2006-01")
}

func dropDown(vals []string, selected string) string {

	opts := ""
	for _, val := range vals {
		isSelected := ""
		if val == selected {
			isSelected = "selected"
		}
		opts += fmt.Sprintf("\t\t<option value='%v' %v >%v</option>\n", val, isSelected, val)
	}

	str := `
	<select name="type">
		%v
	</select>`
	str = fmt.Sprintf(str, opts)
	return str
}

// HTMLForm renders an HTML edit form
// for survey data
func (s *surveyT) HTMLForm(questTypes []string) string {

	ret := `
		<style>
			.survey-edit-form span {
				display: inline-block;
				min-width: 140px;
			}
		</style>

		<form method="POST" class="survey-edit-form" >
		
			<span>Type     </span> %v <br>

			<span>Year     </span><input type="text" name="year"      value="%v"  /> <br>
	 		<span>Month    </span><input type="text" name="month"     value="%v"  /> <br>
			<span>Deadline </span><input type="text" name="deadline"  value="%v" placeholder="dd.mm.yyyy hh:mm" /> <br>

			%v

			<input type="submit" name="submit" id="submit"  value="Submit" accesskey="s"  /> <br>
		</form>

		<script> document.getElementById('submit').focus(); </script>

			%v
		`

	if s == nil {
		*s = NewSurvey("fmt")
	}
	dd := dropDown(questTypes, s.Type)

	kv := ""

	nP := 2
	if len(s.Params) < nP {
		for i := len(s.Params); i < nP; i++ {
			s.Params = append(s.Params, ParamT{})
		}
	}
	for i := 0; i < nP; i++ {
		kv += fmt.Sprintf(
			`
			<span>Param%v    </span>
			<input type="text" name="Params[%v].Name" placeholder="name%v"  value="%v" /> 
			<input type="text" name="Params[%v].Val"  placeholder="val%v"   value="%v" />
			<br>
		`,
			i,
			i, i, s.Params[i].Name,
			i, i, s.Params[i].Val,
		)
	}

	link := fmt.Sprintf(
		"<a href='%v/generate-hashes?survey_id=%v&wave_id=%v' target='_blank' >Generate hash logins</a><br>",
		cfg.Pref(), s.Type, s.WaveID(),
	)

	ret = fmt.Sprintf(ret,
		dd,
		s.Year, fmt.Sprintf("%02d", int(s.Month)+0), s.Deadline.Format("02.01.2006 15:04"),
		kv,
		link,
	)
	return ret

}
