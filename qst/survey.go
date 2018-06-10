package qst

import (
	"time"

	"github.com/zew/go-questionaire/cfg"
)

// surveyT stores the interval components of a questionaire wave.
// For quarterly intervals, it needs to be extended
type surveyT struct {
	Type string `json:"type,omitempty"` // The type identifier, i.e. "fmt" or "cep"

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
	return t.Format("January 2006")
}
