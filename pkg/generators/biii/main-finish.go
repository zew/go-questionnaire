package biii

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func finish(q *qst.QuestionnaireT) {
	//
	// Finish questionnaire?  - one before last page
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Abschluss<br><br>", "en": "Finish"}
		page.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "q44_dsgvo"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 6
				inp.Validator = "must"
				inp.Label = trl.S{
					"de": `
						<b>Einwilligungserklärung gemäß DSGVO</b>

						<br>

						Die Antworten dieser Online-Umfrage werden von uns streng vertraulich, 
						DSGVO-konform behandelt und nur in anonymer bzw. aggregierter Form benutzt.

						<br>

						Im <a href="/doc/site-imprint.md" >Impressum</a> finden Sie umfangreiche Angaben zum Datenschutz.

						<br>

						Hiermit willige ich ein, dass meine gesammelten Daten 
						für die Marktstudie Impact Investing in Deutschland 2022 
						der Bundesinitiative Impact Investing verwendet werden.


						<br>

						Um den Daten bestmöglich wissenschaftlich auswerten, 
						validieren und ggf. bereinigen zu können, wären wir dankbar, 
						wenn Sie uns den Namen Ihrer Organisation nennen könnten. 
						
						Es gilt selbstverständlich weiterhin, dass die Daten streng vertraulich, 
						DSGVO-konform behandelt und nur in anonymer bzw. aggregierter 
						Form benutzt werden.

					`,
				}

				inp.ControlFirst()
				inp.ControlTop()
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				// inp.Type = "hidden"
				inp.Name = "q45_org_name"
				inp.MaxChars = 27
				inp.Label = trl.S{"de": "Name Ihrer Organisation&nbsp;&nbsp;"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0 0 0 6rem"
			}
		}

		// gr0
		{
			labels := []trl.S{
				{"de": `
					Ich erkläre mich einverstanden, 
					dass meine angegebenen personenbezogenen Daten 
					genutzt werden 
					und zu Auswertungszwecken an die European Venture Philanthropy Association (EVPA) 
					weitergeleitet 
					
					und in anonymer bzw. aggregierter Form für eine europaweite 
					Studie zu Impact Investing von EVPA 
					und Global Steering Group for Impact Investments (GSG) verwendet werden.
					
				`},

				{"de": `
					Ich erkläre mich einverstanden, 
					dass meine angegebenen personenbezogenen Daten 
					<i>anonymisiert</i> 
					zu Auswertungszwecken 
					an die European Venture Philanthropy Association (EVPA) weitergeleitet 
					und in anonymer bzw. aggregierter Form für eine europaweite 
					Studie zu Impact Investing von EVPA und Global Steering Group 
					for Impact Investments (GSG) verwendet werden.				
				`},

				{"de": "Meine Daten sollen nicht an die EVPA weitergeleitet werden."},
			}
			radioValues := []string{
				"evpa_yes",
				"evpa_anonymous",
				"evpa_not",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>Weitergabe meiner Daten an die EVPA</b> 
					<br>
					<br>
					
					Mit Ihrer Einwilligung würden wir mit den im Rahmen dieser Erhebung 
					gesammelten Daten gerne auch zu einer europaweiten 
					Studie zu Impact Investing von EVPA und GSG beitragen. 
					
					Dies wäre ein wichtiger Schritt hin zu einem besseren Verständnis 
					des Impact Investing Marktes auf europäischer Ebene 
					sowie zu einer Vergleichbarkeit nationaler Entwicklungen 
					und Trends im Impact Investing.					
					<br>
					<br>
					
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q45"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
				rad.ControlTop()

				rad.Validator = "mustRadioGroup"

			}
		}

		// // gr1
		// {
		// 	gr := page.AddGroup()
		// 	gr.Cols = 1
		// 	gr.BottomVSpacers = 1
		// 	{
		// 		inp := gr.AddInput()
		// 		inp.Type = "dyn-composite"
		// 		inp.DynamicFuncParamset = ""
		// 		inp.DynamicFunc = fmt.Sprintf("QuestForOrg__%v__%v", 0, 0)

		// 		inp.ColSpanControl = 1
		// 	}
		// }

		// gr2
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `Fragebogen abschließen um die Daten final zu speichern.`,
					"en": ``,
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

		// pge.ExampleSixColumnsLabelRight()

	}

	// Report of results
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "", "en": ""}
		page.NoNavigation = true
		page.SuppressProgressbar = true
		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{"de": `	
					<span style="font-size: 130%">
					Vielen Dank für Ihre Mithilfe!
					<br>
					<br>
					Wir werden die Ergebnisse der Erhebung in der 2022 Bundesinitiative
					Impact Investing (BIII) Markstudie aufnehmen und mit Ihnen teilen!
					</span>
				`}
			}

			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.ColSpanControl = 1
			// 	inp.DynamicFunc = "PersonalLink"
			// }
		}
	}

}
