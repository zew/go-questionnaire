package qst

import "fmt"

var q1Pretext = []string{
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von drei Personen (%v) am besten eingestuft und von zwei weiteren am schlechtesten. Stiftung B wird von zwei Personen am besten eingestuft und von dreien mittel, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von zwei Personen (%v) mittel eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von zwei Personen am besten eingestuft und von dreien mittel, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von zwei Personen (%v) am besten eingestuft und von drei weiteren mittel. Stiftung B wird von zwei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von zwei Personen (%v) mittel eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von vier Personen (%v) am besten eingestuft und von einer weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von fünf Personen (%v) mittel eingestuft. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von vier Personen (%v) am besten eingestuft und von einer Person am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von fünf Personen (%v) mittel eingestuft. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von einer Person (%v)  am besten eingestuft und von vier weiteren am schlechtesten. Stiftung B wird von vier Personen am besten eingestuft und von einer am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von drei Personen (%v) am besten eingestuft und von zwei weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von zwei Personen (%v) am besten eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von fünf Personen mittel eingestuft, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von fünf Personen (%v) mittel eingestuft. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</p>",
	"<p>In <b>Entscheidung 1</b> wird Stiftung A von zwei Personen (%v) mittel eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.</p>",
}

// PoliticalFoundationsPretext returns one of 16
// introductions to PoliticalFoundations question series
func PoliticalFoundationsPretext(q *QuestionnaireT, seq0to5, paramSetIdx int) (string, []string, error) {

	userID := 0
	if q != nil {
		userID = q.UserIDInt()
	}

	zeroTo15 := userID % 16

	imgTag := fmt.Sprintf(
		`<img src='%v' class='q1-pretext-img' >`,
		// works on survey2.zew.de - not locally
		// cfg.Pref("/img/pat/person.png"),
		"/img/pat/person.png",
	)

	return fmt.Sprintf(q1Pretext[zeroTo15], imgTag), []string{}, nil

}
