package qst

import (
	"fmt"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/trl"
)

// surveyT stores the interval components of a questionaire wave.
// For quarterly intervals, it needs to be extended
type surveyT struct {
	Type string `json:"type,omitempty"` // The type identifier, i.e. "fmt" or "cep"
	Org  trl.S  `json:"org,omitempty"`  // organization, i.e. Unicef
	Name trl.S  `json:"name,omitempty"` // full name, i.e. programming languages survey

	Year  int        `json:"year,omitempty"`  // The wave year
	Month time.Month `json:"month,omitempty"` // The wave month, 1-based, int

	Deadline time.Time `json:"deadline,omitempty"` // No more responses accepted
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
func (s *surveyT) HTMLForm(vals []string) string {

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

			<input type="submit" name="submit"   value="Submit" accesskey="s"  /> <br>
		</form>
	`

	if s == nil {
		*s = NewSurvey("fmt")
	}
	dd := dropDown(vals, s.Type)

	ret = fmt.Sprintf(ret, dd, s.Year, fmt.Sprintf("%02d", int(s.Month)+0), s.Deadline.Format("02.01.2006 15:04"))
	return ret

}
