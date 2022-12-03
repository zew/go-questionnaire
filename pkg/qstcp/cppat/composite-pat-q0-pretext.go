package cppat

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	qstif "github.com/zew/go-questionnaire/pkg/qstif"
)

var q1Pretext = []string{
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von drei Personen (%v) am besten eingestuft und von zwei weiteren am schlechtesten. Stiftung B wird von zwei Personen am besten eingestuft und von dreien mittel, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) mittel eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von zwei Personen am besten eingestuft und von dreien mittel, und so weiter.</i></p>",
	// "<p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) am besten eingestuft und von drei weiteren mittel. Stiftung B wird von zwei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) am besten eingestuft und von drei weiteren mittel. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) mittel eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von vier Personen (%v) am besten eingestuft und von einer weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von fünf Personen (%v) mittel eingestuft. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von vier Personen (%v) am besten eingestuft und von einer Person am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von fünf Personen (%v) mittel eingestuft. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von drei Personen (%v) am besten eingestuft und von zwei weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) am besten eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von fünf Personen (%v) mittel eingestuft. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</i></p>",
	// "<p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) mittel    eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</i></p>",
	"   <p><i>Erläuterung: In <span>Entscheidung 1</span> wird Stiftung A von zwei Personen (%v) am besten eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter</i></p>",
}

// PoliticalFoundationsPretext returns one of 16
// introductions to PoliticalFoundations question series
func PoliticalFoundationsPretext(q qstif.Q, seq0to5, paramSetIdx int, preflight bool) (string, []string, error) {

	zeroTo15 := q.Version()

	imgTag := fmt.Sprintf(
		`<img src='%v' class='q1-pretext-img' >`,
		// since fully dynamic - this works across localhost and survey2.zew.de
		cfg.Pref("/img/pat/person.png"),
	)

	return fmt.Sprintf(q1Pretext[zeroTo15], imgTag), []string{}, nil

}
