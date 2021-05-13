package pat2

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Part1Intro renders
func Part1Intro(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		//
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<h3>Teil 1</h3>

				<p>
				In diesem Teil der Studie treffen Sie acht Entscheidungen 
				(und beantworten einige Fragen). 
				Nach der Erhebung werden 10&nbsp;% aller Teilnehmer 
				zufällig ausgewählt. Von jedem ausgewählten Teilnehmer wird eine der acht Entscheidungen zufällig bestimmt und genau wie unten beschrieben umgesetzt (alle unten erwähnten Personen existieren wirklich und alle Auszahlungen werden wie beschrieben getätigt).				
				</p>

				<br>
				<br>
				`,
			}
		}

	}

	return nil
}

func Part1Entscheidung78TwoTimesThree(q *qst.QuestionnaireT, pageIdx int, inpName string) error {

	keyVals := []string{
		"pol_gr1:Ein Politiker aus Gruppe 1",
		"cit_gr2:Ein Bürger aus Gruppe 2",
		"cit_gr3:Ein Bürger aus Gruppe 3 (gleiche demographische Eigenschaften wie die Politiker)",
	}

	page := q.EditPage(pageIdx)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<br>
				<p>
					Wer soll entscheiden, ob Stiftung A, B oder C die 30 € erhält? 
				</p>
			`}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		for _, kv := range keyVals {
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q1"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.ColSpanLabel = 5 + 2
			rad.ColSpanControl = 1
			rad.Label = lbl

			rad.ControlFirst()
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<p>
					Wer soll möglichst nicht entscheiden, ob Stiftung A, B oder C die 30 € erhält? 
				</p>
			`}
		}

	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		for _, kv := range keyVals {
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q2"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.ColSpanLabel = 5 + 2
			rad.ColSpanControl = 1
			rad.Label = lbl

			rad.ControlFirst()
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
				<p style="font-size:86%">
				<b>Erläuterung:</b>
				Falls dieser Teil umgesetzt wird und somit bestimmt, 
				welche Stiftung 30 € erhält, 
				werden zufällig zwei der drei Gruppen ausgewählt, 
				die tatsächlich festlegen können, 
				welche Stiftung das Geld erhält. 
				Die dritte Gruppe wird die Entscheidung 
				definitiv nicht treffen. 
				Von den zwei Gruppen, die die Entscheidung treffen können, 
				wird jene die Entscheidung treffen, 
				die Sie gemäß Ihrer Antworten auf die letzten Fragen 
				als besser erachten.
				</p>
			`}
		}
	}

	return nil
}

func Part1Entscheidung78(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		//
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `

				<p>
					Zuletzt haben Sie entschieden, 
					wie die Präferenzen von fünf Personen 
					in eine gemeinsame Entscheidung zusammengefasst werden sollen. 
					
					Dadurch haben Sie festgelegt, welche politische Stiftung 
					eine Spende von 30&nbsp;€ erhält. 
				</p>

				<p>
					Wir haben drei verschiedenen Gruppen von Studienteilnehmern 
					die gleichen Fragen wie Ihnen gestellt. 
					Die drei Gruppen sind die Folgenden:  
				</p>


				<ul>
				<li>Gruppe 1: Eine repräsentative Gruppe deutscher Bürger.
				</li>

				<li>Gruppe 2: Eine repräsentative Gruppe deutscher Land- und Bundestagspolitiker.
				</li>

				<li>Gruppe 3: Eine Gruppe deutscher Bürger, 
				die <i>keine Politiker</i> sind, 
				die aber die <i>gleichen demographischen Eigenschaften wie Politiker</i> haben. 
				Das heißt, Gruppe&nbsp;3 besteht z. B. zu 70&nbsp;% aus Männern, 
				nur 3&nbsp;% der Mitglieder sind unter 30&nbsp;Jahre alt, 
				87&nbsp;% der Mitglieder haben einen Hochschulabschluss 
				und nur 17&nbsp;% sind alleinstehend.  
				</li>
				</ul>
				
				<p>
					Wir haben aus jeder der drei Gruppen jeweils eine Person 
					zufällig für Sie ausgewählt. 
					Falls dieser zweite Teil der Studie umgesetzt wird, 
					wird nicht Ihre eigene Entscheidung bestimmen, 
					welche Stiftung die 30&nbsp;€ erhält. 
					Stattdessen wird eine Entscheidung 
					eines der drei ausgewählten Gruppenmitglieder 
					ausschlaggebend dafür sein, welche Stiftung das Geld erhält.  
				</p>
				<p>
					Sie können aber entscheiden, welcher 
					Gruppe (also Gruppe 1, 2 oder 3) die Person angehören soll, 
					die die Präferenzen zu einer gemeinsamen 
					Entscheidung zusammenfasst und somit festlegt, 
					welche Stiftung das Geld erhält.
				</p>

				<br>
				<br>
				`,
			}
		}

	}

	//
	//
	//
	// Entscheidung 7
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<p><b>Entscheidung 7.</b></p>
					<p>
						In dieser Entscheidung sind die Präferenzen der 
						fünf Personen aus der Vorstudie wie folgt:
					</p>
				`}
			}
		}

		// loop over matrix questions
		// for i := 0; i < 3; i++ {
		for i := 0; i < 1; i++ {

			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1

				{
					inp := gr.AddInput()
					inp.Type = "dyn-composite"
					inp.ColSpanControl = 1
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsStatic__%v__%v", i, i)
				}

			}
		}

		pageIdx := len(q.Pages) - 1
		Part1Entscheidung78TwoTimesThree(q, pageIdx, "dec7")

	}

	//
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<p><b>Entscheidung 8.</b></p>
					<p>
						Nun sind die Präferenzen der fünf Personen aus der Vorstudie wie folgt:
					</p>
				`}
			}
		}

		// loop over matrix questions
		// for i := 0; i < 3; i++ {
		for i := 3; i < 4; i++ {

			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1

				{
					inp := gr.AddInput()
					inp.Type = "dyn-composite"
					inp.ColSpanControl = 1
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsStatic__%v__%v", i, i)
				}

			}
		}

		pageIdx := len(q.Pages) - 1
		Part1Entscheidung78TwoTimesThree(q, pageIdx, "dec8")

	}

	return nil
}
