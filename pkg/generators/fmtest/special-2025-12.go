package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202512(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2025 && q.Survey.Month == 12
	if !cond {
		return nil
	}

	// page 1
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "&nbsp;",
			"en": "&nbsp;",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			"de": "Klima-<br>erwartungen",
			"en": "Climate<br>Expectations",
		}
		page.WidthMax("56rem")

		//
		{
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Label = trl.S{
					"de": `
						Im kommenden Abschnitt würden wir gerne mehr über Ihre Erwartungen zum Klimawandel erfahren. Wenn in den Fragen von einer „klimaneutralen Wirtschaft“ die Rede ist, dann ist damit eine Wirtschaft gemeint, die eine Netto-Null-Bilanz bei den Treibhausgasemissionen erreicht hat. Das heißt, alle vom Menschen verursachten Emissionen werden durch die Entfernung aus der Atmosphäre ausgeglichen, sodass keine Auswirkungen auf das Klimasystem entstehen.
				`,
					"en": `
						In the next section, we would like to learn more about your expectations regarding climate change. When the questions refer to a “climate-neutral economy”, this should be read as an economy that has achieved a net-zero greenhouse gas emissions balance, meaning all human-caused emissions are balanced by removals from the atmosphere, ensuring no impact on the climate system. 

					`,
				}
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplate6b,
					labelsUnlikely(),
					[]string{"ssq1a", "ssq1b", "ssq1c"},
					[]string{"very_unlikely", "unlikely", "neutral", "likely", "very_likely", "no_answer"},
					[]trl.S{
						{
							"de": "In den nächsten 25&nbsp;Jahren wird sich der Lebens&shy;standard vieler Menschen <i>weltweit</i> aufgrund des Klimawandels verschlechtern.",
							"en": "Over the next 25&nbsp;years, the standard of living for many people around the world will decline due to climate change.",
						},
						{
							"de": "In den nächsten 25&nbsp;Jahren wird sich der Lebens&shy;standard vieler Menschen <i>in Deutschland</i> aufgrund des Klimawandels verschlechtern.",
							"en": "Over the next 25&nbsp;years, the standard of living for many people in Germany will decline due to climate change.",
						},
						{
							"de": "Bis zum Ende dieses Jahrhunderts wird das Erdsystem einen kritischen Kipppunkt erreichen, der zu irreversiblen Umweltveränderungen führen wird. ",
							"en": "By the end of this century, the Earth system is going to reach a critical tipping point leading to irreversible environmental changes.",
						},
					},
				)

				gb.MainLabel = trl.S{
					"de": `
						Angenommen, die globale Klimapolitik bleibt so, wie sie heute ist, und es werden keine zusätzlichen Maßnahmen zur Bekämpfung des Klimawandels ergriffen:
						<br>
						<br>
						Wie wahrscheinlich schätzen Sie unter diesen Bedingungen die folgenden Szenarien ein?
						<br>
						<br>
				`,
					"en": `
						Assuming that global climate policies stay the way they are today, and no additional measures are taken to tackle climate change:
						<br>
						<br>
						How likely do you find the following scenarios under these conditions?
						<br>
						<br>
				`,
				}.Outline("1.")
				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 4

				gr.Style = css.NewStylesResponsive(gr.Style)
				gr.Style.Desktop.StyleGridContainer.GapRow = "0.9rem"
				gr.Style.Desktop.StyleGridContainer.GapColumn = "0.6rem"
				gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

				gr.Class = "grid-2025-12-ssq1"
				//
				// rad.StyleLbl.Desktop.StyleText.AlignHorizontal = "justify"
			}

		}

		qst.ChangeHistoryJS(q, page)

	}

	// page 2

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			"de": "Klima-<br>erwartungen",
			"en": "Climate<br>Expectations",
		}
		page.WidthMax("56rem")
		page.SuppressInProgressbar = true
		// page.NoNavigation = true

		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `
						Das Pariser Klimaabkommen zielt darauf ab, den „Anstieg der durchschnittlichen Erdtemperatur deutlich unter 2 °C über dem vorindustriellen Niveau“ zu halten.

				`,
					"en": `
						The Paris Agreement aims to hold “the increase in the global average temperature to well below 2 °C above pre-industrial levels”. 					
					`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `
						Angenommen, die globale Klimapolitik bleibt so, wie sie heute ist, und es werden keine zusätzlichen Maßnahmen zur Bekämpfung des Klimawandels ergriffen:
						<br>
						<br>
						Wie hoch schätzen Sie unter diesen Bedingungen den Anstieg der durchschnittlichen Erdtemperatur bis zum Ende dieses Jahrhunderts ein?
						<br>
						<br>
				`,
					"en": `
						Assuming that global climate policies stay the way they are today, and no additional measures are taken to tackle climate change:
						<br>
						<br>
						What would be your expectations for the global average temperature rise by the end of this century under these conditions? 
						<br>
						<br>
				`,
				}.Outline("2.")
			}
		}
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.WidthMax("40rem")
			gr.BottomVSpacers = 2

			lbls := []trl.S{
				{
					"de": "Kein signifikanter Anstieg ",
					"en": "No significant rise",
				},
				{
					"de": "Etwa 1,5 °C",
					"en": "About 1.5 °C",
				},
				{
					"de": "Etwa 2 °C",
					"en": "About 2 °C",
				},
				{
					"de": "Etwa 3 °C",
					"en": "About 3 °C",
				},
				{
					"de": "Etwa 4 °C",
					"en": "About 4 °C",
				},
				{
					"de": "Mehr als 4 °C",
					"en": "More than 4 °C",
				},
			}

			for i := 0; i < len(lbls); i++ {
				{
					inp := gr.AddInput()
					inp.Type = "radio"
					inp.Name = "ssq2"
					inp.ValueRadio = fmt.Sprintf("%v", i+1)
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 6
					inp.Label = lbls[i]
					inp.ControlFirst()
				}
			}

		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5a,
				labelsCertainty(),
				[]string{"ssq2"},
				[]string{"very_uncertain", "uncertain", "neutral", "certain", "very_certain", "no_answer"},
				nil,
			)

			gb.MainLabel = trl.S{
				"de": `
					Auf einer Skala von 1 (überhaupt nicht sicher) bis 5 (sehr sicher), wie sicher sind Sie sich hinsichtlich Ihrer Einschätzung in der vorherigen Frage?
				`,
				"en": `
					On a scale from 1 (not at all confident) to 5 (very confident), how confident are you about your assessment in the previous question?
				`,
			}.Outline("3.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 2

		}

	}

	return nil

}
