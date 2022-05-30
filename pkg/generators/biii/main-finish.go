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
		// page.Label = trl.S{"de": "", "en": ""}
		page.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("36rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = "q44_dsgvo"
				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6
				rad.Validator = "must"
				rad.Label = trl.S{
					"de": `
						<b>Einwilligungserklärung gemäß DSGVO</b>

						<br>

						Die Antworten dieser Online-Umfrage werden von uns streng vertraulich, 
						DSGVO-konform behandelt und nur in anonymer bzw. aggregierter Form benutzt.

						<br>

						Hiermit willige ich ein, dass meine gesammelten Daten 
						für die BIII Impact Investment Studie 2022 verwendet werden. 
					`,
				}
				rad.ControlFirst()
			}
		}

		// gr0
		{
			labels := []trl.S{
				{"de": "Ich erkläre mich einverstanden, dass meine angegebenen personenbezogenen Daten genutzt werden und zu Auswertungszwecken an die European Venture Philanthropy Association (EVPA) weitergeleitet werden dürfen. "},
				{"de": "Meine Daten sollen anonymisiert verwendet werden."},
			}
			radioValues := []string{
				"evpa",
				"anonymous",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>Weitergabe meiner Daten</b> <br><br>"}
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

				rad.Validator = "mustRadioGroup"

				if idx == 0 {
					{
						inp := gr.AddInput()
						inp.Type = "text"
						inp.Name = "q45_org_name"
						inp.MaxChars = 27
						inp.Label = trl.S{"de": "Name Ihrer Organisation"}
						inp.ColSpan = gr.Cols
						inp.ColSpanLabel = 2
						inp.ColSpanControl = 3
						inp.Style = css.NewStylesResponsive(inp.Style)
						inp.Style.Desktop.StyleBox.Margin = "0 0 0 6rem"
					}
					{
						inp := gr.AddInput()
						inp.Type = "textarea"
						inp.Name = "q45_comment"
						inp.MaxChars = 150
						inp.Label = trl.S{"de": "Kommentar"}
						inp.ColSpan = gr.Cols
						inp.ColSpanLabel = 2
						inp.ColSpanControl = 3
						inp.Placeholder = trl.S{"de": "Kommentarfeld"}
						inp.Style = css.NewStylesResponsive(inp.Style)
						inp.Style.Desktop.StyleBox.Margin = "0 0 0 6rem"
					}
				}

			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
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
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PersonalLink"
			}
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.DynamicFunc = "RenderStaticContent"
			// 	inp.DynamicFuncParamset   = "site-imprint.md"
			// 	inp.ColSpan = 1
			// 	inp.ColSpanLabel = 1
			// }
		}
	}

}
