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
	// Entscheidung 8
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
		for i := 1; i < 2; i++ {

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

// Part2Intro renders
func Part2Intro(q *qst.QuestionnaireT) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	{
		gr := page.AddGroup()
		gr.Cols = 1
		// gr.BottomVSpacers = 2

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<h3>Teil 2</h3>

				<p>
					In diesem Teil der Studie beantworten Sie sechs Fragen. 
					Nach der Erhebung werden 10&nbsp;% aller Teilnehmer zufällig ausgewählt. 
					Jeder ausgewählte Teilnehmer wird in Abhängigkeit der Genauigkeit 
					seiner Antworten eine Bonuszahlung von bis zu 50&nbsp;Norstat&nbsp;Coins erhalten (Wert: 5&nbsp;Euro).
				</p>

				<p style='padding-bottom: 0; padding-top: 0.5rem'>
					Es geht in diesem Teil wieder um die drei Gruppen aus dem letzten Teil:
				</p>

				<ul>
					<li>
						Gruppe 1: Eine repräsentative Gruppe deutscher Bürger.
					</li>
					<li>
						Gruppe 2: Eine repräsentative Gruppe deutscher Land- und Bundestagspolitiker.
					</li>
					<li>
						Gruppe 3: Eine Gruppe deutscher Bürger, die <i>keine Politiker</i> sind, 
						die aber die <i>gleichen demographischen Eigenschaften wie Politiker</i> haben. 
						Das heißt, Gruppe 3 besteht z. B. zu 70&nbsp;% aus Männern, 
						nur 3&nbsp;% der Mitglieder sind unter 30&nbsp;Jahre alt, 
						87&nbsp;% der Mitglieder haben einen Hochschulabschluss 
						und nur 17&nbsp;% sind alleinstehend.
					</li>
				<ul>

				<p>
					Wir bitten Sie zu schätzen, welche Stiftung die Mitglieder dieser Gruppen als 
					Empfänger der 30&nbsp;€ bestimmt haben, wenn sie bestimmte Präferenzen 
					der Mitglieder der Vorstudie gesehen haben. 
				</p>

				<p>
					Falls Sie eine Bonuszahlung erhalten, 
					werden wir eine der sechs Fragen zufällig auswählen 
					und Ihre Schätzung mit den echten Entscheidungen der Gruppenmitglieder abgleichen. 
					Ihre Bonuszahlung ist umso höher, 
					je genauer Ihre Einschätzung ist. 
					Bitte überlegen Sie sich Ihre Antworten daher sehr genau!
				</p>

				<br>

				<p style="font-size:86%">
					<b>*Erläuterung:</b>
					Falls Sie in der ausgewählten Frage eine 100 % richtige Antwort geben, 
					werden Sie 50&nbsp;Norstat&nbsp;coins erhalten. 
					Für jede Person, die Sie bei Ihren folgenden Schätzungen zu viel oder zu wenig angeben, 
					werden Sie 2.5&nbsp;Norstat&nbsp;coins verlieren. 
					Falls beispielsweise alle 10&nbsp;Gruppenmitglieder Stiftung&nbsp;C wählten, 
					Sie aber angeben, dass 5&nbsp;Gruppenmitglieder Stiftung&nbsp;B wählen, 
					und weitere 5 Stiftung&nbsp;C wählten, 
					dann haben Sie für Stiftung&nbsp;C fünf Gruppenmitglieder zu wenig angegeben, 
					und für Stiftung&nbsp;B fünf zu viel. 
					Entsprechend wird Ihre Bezahlung auf 50-2.5× 5 -2.5× 5=25&nbsp;Norstat&nbsp;coins gesenkt.
				</p>

				`,
			}
		}

	}
	return nil

}

// Part2Block1 renders
// blockStart is either 0 - or 3
func Part2Block12(q *qst.QuestionnaireT, blockStart int) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	page.ValidationFuncName = "pat2-add-to-10"
	page.ValidationFuncMsg = trl.S{"de": "Wollen Sie wirklich weiterfahren, ohne dass sich Ihre Eintraege auf 10 summieren?"}

	//
	//
	//
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
					Schätzen Sie zunächst für die folgende Präferenzkonstellation der fünf Personen:
				</p>
			`}
			if blockStart > 0 {
				inp.Desc = trl.S{"de": `
				<p>
					Schätzen Sie als nächstes für die folgende Präferenzkonstellation der fünf Personen:
				</p>
			`}
			}
		}
	}

	// loop over matrix questions
	// blockStart is either 0 or 3
	zeroOrOne := blockStart / 3
	for i := zeroOrOne; i < zeroOrOne+1; i++ {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsStatic__%v__%v", i, i)
			}

		}
	}

	questLabels := []string{
		"Was glauben Sie, wie haben sich die 10&nbsp;Politiker aus Gruppe&nbsp;1 entschieden?",
		"Was glauben Sie, wie haben sich die 10&nbsp;deutschen Bürger aus Gruppe&nbsp;2 entschieden?",
		`Was glauben Sie, wie haben sich die 10&nbsp;deutschen Bürger aus Gruppe&nbsp;3 entschieden 
		(demographische Eigenschaften der Politiker, 
			also 70 % Männer, 3 % unter 30 Jahre, halb so oft alleinstehend)? `,
	}

	for i1 := blockStart; i1 < blockStart+3; i1++ {

		gr := page.AddGroup()
		gr.Cols = 24
		gr.BottomVSpacers = 3
		lbls := []string{"A", "B", "C"}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 24
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
					<p>
						<b> Frage %v</b> <br>
						%v 
					</p>
					`, i1+1, questLabels[i1%3]),
			}
		}
		for i2 := 0; i2 < 3; i2++ {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("part2_q%v_q%v", i1+1, i2+1)
			inp.MaxChars = 2
			inp.Min = 0
			inp.Max = 10
			inp.ColSpan = 8
			inp.Label = trl.S{"de": fmt.Sprintf("von 10 wählten Stiftung&nbsp;%v", lbls[i2])}
			inp.Validator = "inRange10"
			inp.ControlFirst()
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 24
			inp.Desc = trl.S{
				"de": `
					<p style='font-size:90%'>
					Ihre Antworten müssen sich auf 10 summieren.	
					</p>
					`,
			}
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "center"
		}

	}

	if blockStart > 0 {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<br>
					<p>
					<b>
						Dies ist das Ende dieser Studie. 
						Wir bedanken uns ganz herzlich für Ihre Teilnahme. 
						Falls Sie zu den zufällig ausgewählten 10% gehören, 
						werden Sie Ihre Bonuszahlung wie versprochen in den nächsten Tagen erhalten. 
					</b>
					</p>
				`}
			}
		}
	}

	return nil
}
