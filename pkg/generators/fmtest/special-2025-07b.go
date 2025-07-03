package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202507b(q *qst.QuestionnaireT, page *qst.WrappedPageT) error {

	cond := false
	cond = cond || q.Survey.Year == 2025 && q.Survey.Month == 7
	if !cond {
		return nil
	}

	// qs2b
	{
		lblMain := trl.S{
			"de": `
					Wie schätzen Sie den potentiellen Einfluss von zukünftigen Ölpreisentwicklungen auf den Hauptrefinanzierungssatz der EZB in den nächsten 6 bis 12 Monaten ein?
				`,
			"en": `
					How do you assess the potential impact of future oil price developments on the ECB’s main refinancing rate over the next 6 to 12 months? Oil price developments are most likely to…
				`,
		}.Outline("3b.")

		lblsRows := []trl.S{
			{
				"de": "die Wahrscheinlichkeit von Zinserhöhungen erhöhen (++)",
				"en": "increase the likelihood of rate hikes (++)",
			},
			{
				"de": "die Wahrscheinlichkeit von Zinssenkungen verringern (+)",
				"en": "decrease the likelihood of rate cuts (+)",
			},
			{
				"de": "keinen Einfluss haben",
				"en": "have no impact",
			},
			{
				"de": "die Wahrscheinlichkeit von Zinserhöhungen verringern (-)",
				"en": "decrease the likelihood of rate hikes (-)",
			},
			{
				"de": "die Wahrscheinlichkeit von Zinssenkungen erhöhen (--)",
				"en": "increase the likelihood of rate cuts (--)",
			},
		}

		const col1Width = 1

		gr := page.AddGroup()
		gr.Cols = col1Width
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.7rem"
		gr.Style.Desktop.StyleBox.Width = "80%"
		gr.Style.Mobile.StyleBox.Width = "100%"

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = lblMain
		}

		for i1 := 0; i1 < len(lblsRows); i1++ {

			{
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = "qs3b"
				inp.ValueRadio = fmt.Sprint(i1 + 1)

				inp.ColSpan = col1Width

				inp.Label = lblsRows[i1]

				// inp.ControlTopNudge()
				inp.LabelPadRight()

				inp.ControlFirst()
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 6

			}

		}

	}

	return nil
}
