package fmtest

import fmtt "fmt"

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202408(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 8
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "Sonderfrage zur US Präsidentschaftswahl",
		"en": "Special questions: US presidential election",
	}
	page.Short = trl.S{
		"de": "Sonderfrage<br>US Präsidentschaftswahl",
		"en": "Special questions:<br>US presidential election",
	}
	page.WidthMax("48rem")

	indent := css.NewStylesResponsive(nil)
	indent.Desktop.StyleBox.Padding = "0  0.8rem  0  1.4rem"
	indent.Mobile.StyleBox.Padding = " 0  0.8rem  0  0.6rem"

	indent.Desktop.StyleBox.Margin = "0  0  0.95rem  0"

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
					Wir möchten Sie gerne zu Ihrer persönlichen Einschätzung der Wahrscheinlichkeit einer Wahl von Donald Trump als nächster Präsident der USA befragen und wie bestimmte Ereignisse Ihre Erwartung beeinflusst haben. Bitte geben Sie Ihre Antworten auf einer Skala von 0% (absolut unwahrscheinlich) bis 100% (absolut sicher) an.

					<br>
					<br>

					<b>
					Wie hoch war Ihrer Meinung nach die Wahrscheinlichkeit mit der Donald Trump als nächster US-Präsident gewählt wird…
					</b>
				`,
				"en": `
					We would like to ask for your personal assessment of the probability of Donald Trump being elected as the next US president and how certain events affected your expectation. Please state your answers on a scale from 0% (no chance) to 100% (absolutely certain).


					<br>
					<br>

					<b>
					In your opinion, how likely was it that Donald Trump would be elected as the next US president…
					</b>


				`,
			}.Outline("4.")
			inp.ColSpanLabel = 1
			inp.ColSpan = gr.Cols
		}

		{
			lbls := []trl.S{
				{
					"de": `…vor dem TV-Duell zwischen Trump und Biden am 27.&nbsp;Juni&nbsp;2024?`,
					"en": `…before the TV debate between Trump and Biden on 27&nbsp;June&nbsp;2024?`,
				},
				{
					"de": `…nach dem TV-Duell zwischen Trump und Biden am 27.&nbsp;Juni&nbsp;2024 aber vor dem Attentat auf Trump am 13.&nbsp;Juli&nbsp;2024?`,
					"en": `…after the TV debate between Trump and Biden on 27&nbsp;June&nbsp;2024 but before the assassination attempt on Trump on 13&nbsp;July&nbsp;2024?`,
				},
				{
					"de": `…nach dem Attentat auf Trump am 13.&nbsp;Juli&nbsp;2024 aber bevor sich Biden am 21.&nbsp;Juli&nbsp;2024 aus dem Präsidentschaftsrennen zurückgezogen hat?`,
					"en": `…after the assassination attempt on Trump on 13&nbsp;July&nbsp;2024 but before Biden decided to step down from the presidential race on 21&nbsp;July&nbsp;2024? `,
				},
				{
					"de": `…nachdem sich Biden am 21.&nbsp;Juli&nbsp;2024 aus dem Präsidentschaftsrennen zurückgezogen hat?`,
					"en": `…after Biden decided to step down from the presidential race on 21&nbsp;July&nbsp;2024?`,
				},
			}

			for idx, lbl := range lbls {
				inp := gr.AddInput()
				inp.Type = "number"

				if false {
					inp.Label = lbl.Outline(fmtt.Sprintf("%v.", idx+1))
				}
				inp.Label = lbl

				inp.Name = fmtt.Sprintf("sp_trump_%v", idx+1)

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1

				inp.Min = 0
				inp.Max = 100
				inp.Step = 1
				inp.MaxChars = 5
				inp.Suffix = trl.S{
					"de": "Prozent",
					"en": "percent",
				}
				inp.Placeholder = trl.S{
					"de": "00",
					"en": "00",
				}

				inp.Style = indent

			}

		}
	}

	return nil

}
