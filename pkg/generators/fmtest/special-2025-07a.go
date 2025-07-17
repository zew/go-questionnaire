package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202507a(q *qst.QuestionnaireT, page *qst.WrappedPageT) error {

	cond := false
	cond = cond || q.Survey.Year == 2025 && q.Survey.Month == 7
	if !cond {
		return nil
	}

	// qs2b
	{
		lblMain := trl.S{
			"de": `
					Was ist Ihre Prognose für den <i>durchschnittlichen Brent Rohölpreis (USD pro Barrel)</i> für die folgenden Perioden? Bitte geben Sie Punktschätzungen zusammen mit einem zentralen 90% Konfidenzintervall an.
				`,
			"en": `
					What is your forecast for the <i>average Brent crude oil price (USD per barrel)</i> for the following periods? Please provide point estimates along with a central 90% confidence interval.
				`,
		}.Outline("2b.")

		lblsCols := []trl.S{
			{
				"de": "&nbsp;",
				"en": "&nbsp;",
			},
			{
				"de": "Punktprognose",
				"en": "Point forecast",
			},
			{
				"de": "90% Konfidenz&shy;intervall",
				"en": "90% confidence interval",
			},
		}

		inpNames := []string{
			"qs2b_brent_2025q3",
			"qs2b_brent_2025q4",
			"qs2b_brent_2026",
		}

		lblsRows := []trl.S{
			{
				"de": "Q3 2025",
				"en": "Q3 2025",
			},
			{
				"de": "Q4 2025",
				"en": "Q4 2025",
			},
			{
				"de": "2026",
				"en": "2026",
			},
		}

		const col1Width = 1
		const col2Width = 1
		const col34Width = 1

		gr := page.AddGroup()
		gr.Cols = col1Width + col2Width + 2*col34Width
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

		// header row with column labels
		for i1 := 0; i1 < len(lblsCols); i1++ {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				if i1 == 0 {
					inp.ColSpan = col1Width
				}
				if i1 == 1 {
					inp.ColSpan = col2Width
				}
				if i1 == 2 {
					inp.ColSpan = 2 * col34Width
				}
				inp.Label = lblsCols[i1]
				if i1 == 2 {
					inp.LabelCenter()
				}
				inp.LabelBottom()
			}
		}

		for i1 := 0; i1 < len(inpNames); i1++ {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = col1Width
				inp.Label = lblsRows[i1]

				inp.ColSpanLabel = 1
				inp.LabelRight()
				inp.LabelPadRight()

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v%v", inpNames[i1], "_pf")

				inp.Min = 0
				inp.Max = 200
				inp.Step = 1
				inp.MaxChars = 5

				inp.Placeholder = trl.S{
					"de": "$",
					"en": "$",
				}

				inp.ColSpan = col2Width
				// inp.ControlTopNudge()

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v%v", inpNames[i1], "_lb")
				inp.Label = trl.S{
					"de": "zwischen",
					"en": "between",
				}

				inp.Min = 0
				inp.Max = 200
				inp.Step = 1
				inp.MaxChars = 5

				inp.Placeholder = trl.S{
					"de": "$",
					"en": "$",
				}

				inp.ColSpan = col34Width
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1

				// inp.ControlCenter()
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v%v", inpNames[i1], "_ub")
				inp.Label = trl.S{
					"de": "und",
					"en": "and",
				}
				inp.Placeholder = trl.S{
					"de": "$",
					"en": "$",
				}

				inp.Suffix = trl.S{"de": "&nbsp;USD", "en": "&nbsp;USD"}

				inp.Min = 0
				inp.Max = 200
				inp.Step = 1
				inp.MaxChars = 5

				inp.ColSpan = col34Width
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1

				// inp.ControlCenter()
			}

		}

	}

	return nil
}
