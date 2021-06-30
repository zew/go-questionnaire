package pat

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// ComprehensionCheck1
//  pat1 - page2
//  pat2 - page6
func ComprehensionCheck1(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		// loop over matrix questions

		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
				<p>
					<!--  <b>Frage:</b> <br>  -->
					Die Präferenzen der Gruppenmitglieder*innen sind wie folgt dargestellt:
				</p>
				`,
				}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsComprehensionCheck__0__0")
			}

		}

		imgTag := fmt.Sprintf(
			`<img src='%v' class='q1-pretext-img' >`,
			// cfg.Pref("/img/pat/person.png"),
			"/img/pat/person.png", // hack - works only online
		)

		{
			gr := page.AddGroup()
			gr.Cols = 1
			// gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
				<p>
					In diesem Beispiel wird Stif­tung A von 
					einer Person (%v) am besten eingestuft 
					und von vier weiteren mittel. 
				</p>
					
				<p>
					Stif­tung B wird von einer Person mittel eingestuft 
					und von vier weiteren am schlech­testen. 
				</p>
					
				<p>
					Stiftung C wird von vier Personen am besten eingestuft 
					und von einer am schlechtesten. 
				</p>
				`, imgTag),
				}
			}
		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

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
				inp.Desc = trl.S{
					"de": `
				<p>
					<b>Verständnistest:</b> <br>
					Nehmen Sie an, 
					die Präferenzen der Gruppenmitglieder 
					sind wie folgt gegeben 
					(dies sind andere Präferenzen als auf der vorherigen Seite):
				</p>
				`,
				}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = fmt.Sprintf("PoliticalFoundationsComprehensionCheck__0__1")
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 8
			gr.BottomVSpacers = 3

			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpan = 8
				inp.DynamicFunc = "ErrorProxy"
				inp.Param = "q_found_compr_"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Position = "relative"
				inp.Style.Desktop.StyleBox.Top = "1.38rem"
				// inp.Style.Desktop.StyleBox.Border = "1px solid green"
				inp.Style.Desktop.StyleBox.HeightMin = "0.3rem"

				inp.Style.Mobile.StyleBox.HeightMin = "1.3rem"
				inp.Style.Mobile.StyleBox.Top = "2.51rem"

			}

			{

				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": `
					<p style="margin-bottom: 0.05rem">
						<b>1.</b> &nbsp; 
						Wieviele Leute stufen Stiftung A als mittel ein? 
					</p>
					`,
				}
			}

			comprehA := []string{
				"est_0:0",
				"est_1:1",
				"est_2:2",
				"est_3:3",
				"est_4:4",
				"est_5:5",
			}

			for idx, kv := range comprehA {
				sp := strings.Split(kv, ":")
				radVal := sp[0]
				lbl := trl.S{"de": "&nbsp;&nbsp;" + sp[1]}

				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q_found_compr_a"
				rad.ValueRadio = radVal
				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 1
				rad.Label = lbl
				// rad.ControlFirst()
				rad.LabelRight()

				if idx == 0 {
					// rad.Validator = "must;comprehensionPOP2"
					rad.Validator = "comprehensionPOP2"
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": `
					<p style="margin-bottom: 0.05rem">
						<b>2.</b>
						Welche Stiftung wird von drei Leuten als am besten eingestuft?
					</p>
					`,
				}
			}

			comprehB := []string{
				"est_a:A",
				"est_b:B",
				"est_c:C",
			}

			for idx, kv := range comprehB {
				sp := strings.Split(kv, ":")
				radVal := sp[0]
				lbl := trl.S{"de": "&nbsp;&nbsp;" + sp[1]}

				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q_found_compr_b"
				rad.ValueRadio = radVal
				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 1
				rad.Label = lbl
				// rad.ControlFirst()
				rad.LabelRight()

				if idx == 0 {
					// rad.Validator = "must;comprehensionPOP2"
					rad.Validator = "comprehensionPOP2"
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 8
				inp.Desc = trl.S{
					"de": `
					<p style="margin-top: 0.35rem">
						Sie haben drei Versuche. Nach dem dritten Fehlversuch wird die Umfrage automatisch beendet. Falls Sie denken, 
						dass sie die Anleitung verstanden haben, 
						aber trotzdem nicht weiterkommen, 
						senden Sie bitte eine 
						<a href="mailto:politik-umfrage@uni-koeln.de">Email</a>.

					</p>
					`,
				}
			}

		}

	}

	return nil
}

// ComprehensionCheck2 - single question
func ComprehensionCheck2(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		// loop over matrix questions

		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
				<p>
					<b>Verständnistest:</b> <br>
					<!-- Nehmen Sie an, jemand habe sich wie folgt entschieden: -->
					Nehmen Sie an, ein*e Teilnehmer*in wie Sie 
					habe die Optionen wie 
					folgt verfügbar und nicht verfügbar gemacht.
				</p>
				`,
				}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = fmt.Sprintf("TimePreferenceSelfComprehensionCheck__0__0")
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpan = 1
				inp.DynamicFunc = "ErrorProxy"
				inp.Param = "q_tpref_compr"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Position = "relative"
				inp.Style.Desktop.StyleBox.Top = "1.28rem"
				// inp.Style.Desktop.StyleBox.Border = "1px solid green"
				inp.Style.Desktop.StyleBox.HeightMin = "0.3rem"

				inp.Style.Mobile.StyleBox.HeightMin = "1.3rem"
				inp.Style.Mobile.StyleBox.Top = "2.38rem"
			}

			// q2
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q_tpref_compr_a"
				inp.MaxChars = 3
				inp.Min = 0
				inp.Max = 100
				inp.ColSpan = 1
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				// inp.Placeholder = trl.S{"de": "0-5"}
				inp.Label = trl.S{"de": "<b>1.</b> Was ist der höchste Betrag, den der*die zukünftige Teilnehmer*in durch seine*ihre Auswahl bei den verfügbaren Optionen sofort erhalten kann?"}
				inp.Suffix = trl.S{"de": "€"}
				// inp.Validator = "must"
				inp.Validator = "comprehensionPOP3"
				inp.LabelPadRight()
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q_tpref_compr_b"
				inp.MaxChars = 3
				inp.Min = 0
				inp.Max = 100
				inp.ColSpan = 1
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				// inp.Placeholder = trl.S{"de": "A,B oder C"}
				inp.Label = trl.S{"de": "<b>2.</b> Wenn der*die zukünftige Teilnehmer*in sich entscheidet, 3 € sofort zu erhalten, wieviel wird er*sie in 6 Monaten erhalten? "}
				inp.Suffix = trl.S{"de": "€"}
				// inp.Validator = "inRange1000"
				// inp.Validator = "must"
				inp.Validator = "comprehensionPOP3"
				inp.LabelPadRight()
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{
					"de": `
					<p style="margin-top: 0.35rem">
						Sie haben drei Versuche. Nach dem dritten Fehlversuch wird die Umfrage automatisch beendet. Falls Sie denken, 
						dass sie die Anleitung verstanden haben, 
						aber trotzdem nicht weiterkommen, 
						senden Sie bitte eine 
						<a href="mailto:politik-umfrage@uni-koeln.de">Email</a>.

					</p>
					`,
				}
			}

		}

	}

	return nil
}
