package fmt

import (
	"fmt"
	"time"

	"github.com/zew/go-questionaire/trl"
)

func nextWaveID() string {
	t := time.Now()
	if t.Day() > 20 {
		t = t.AddDate(0, 1, 0)
	}
	return t.Format("2006-01")
}

func nextQ() string {
	t := time.Now()
	m := t.Month() // 1 - january
	y := t.Year()
	qNow := int(m/3) + 1
	qNext := qNow + 1
	if qNext > 4 {
		qNext = 1
		y++
	}
	return fmt.Sprintf("Q%v %v", qNext, y)
}

func nextY() string {
	t := time.Now()
	y := t.Year()
	y++
	return fmt.Sprintf("%v", y)
}

func labelsGoodBad() []trl.S {

	tm := []trl.S{
		{
			"de": "gut",
			"en": "good",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "schlecht",
			"en": "bad",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsImproveDeteriorate() []trl.S {

	tm := []trl.S{
		{
			"de": "verbessern",
			"en": "improve",
		},
		{
			"de": "nicht verändern",
			"en": "not change",
		},
		{
			"de": "verschlechtern",
			"en": "deteriorate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}
