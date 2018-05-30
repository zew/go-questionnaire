package qst

import (
	"time"

	"github.com/zew/go-questionaire/cfg"
)

// WaveID_T stores the interval components of a questionaire wave.
// For quarterly intervals, it needs to be extended
type WaveID_T struct {
	Year  int
	Month time.Month // 1-based, int

	Deadline time.Time
}

// NewWaveID returns wave ID based on current time
func NewWaveID() WaveID_T {
	t := time.Now()
	if t.Day() > 20 {
		t = t.AddDate(0, 1, 0)
	}
	w := WaveID_T{}
	w.Year = t.Year()
	w.Month = t.Month()

	w.Deadline = time.Date(w.Year, w.Month, 28, 23, 59, 59, 0, cfg.Get().Loc)
	return w
}

// String is the default identifier
func (w WaveID_T) String() string {
	// Notice the month +1
	// It is necessary, even though the spec says 'January = 1'
	t := time.Date(w.Year, w.Month+1, 0, 0, 0, 0, 0, cfg.Get().Loc)
	return t.Format("2006-01")
}

// Label is a pretty identifier
func (w WaveID_T) Label() string {
	// Notice the month +1
	// It is necessary, even though the spec says 'January = 1'
	t := time.Date(w.Year, w.Month+1, 0, 0, 0, 0, 0, cfg.Get().Loc)
	return t.Format("January 2006")
}
