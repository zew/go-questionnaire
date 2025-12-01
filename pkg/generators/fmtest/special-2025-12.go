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
				[]string{"ssq3"},
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

	// page 3

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

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6a,
				labelsImpact(),
				[]string{"ssq4a", "ssq4b"},
				[]string{"strong_neg", "neg", "neutral", "pos", "strong_pos", "no_answer"},
				[]trl.S{
					{
						"de": "In den nächsten 5 Jahren",
						"en": "In the next 5 years ",
					},
					{
						"de": "In den nächsten 25 Jahren",
						"en": "In the next 25 years",
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
					Angenommen, die globale Klimapolitik bleibt so, wie sie heute ist, und es werden keine zusätzlichen Maßnahmen zur Bekämpfung des Klimawandels ergriffen:
					<br>
					<br>
					Wie wird sich der Klimawandel unter diesen Bedingungen Ihrer Meinung nach auf das Wirtschaftswachstum in der EU auswirken?
					<br>
				`,
				"en": `
					Assuming that global climate policies stay the way they are today and no additional measures are taken to tackle climate change:
					<br>
					<br>
					How do you think climate change will impact economic growth in the EU under these conditions?
					<br>
				`,
			}.Outline("4.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 2

			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.9rem"
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
			gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

			// gr.Class = "grid-2025-12-ssq1"
			//
			// rad.StyleLbl.Desktop.StyleText.AlignHorizontal = "justify"
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5a,
				labelsStringency(),
				[]string{"ssq5"},
				[]string{"much_less_s", "less_s", "neutral", "more_s", "much_more_s", "no_answer"},
				nil,
			)
			gb.MainLabel = trl.S{
				"de": `
					Was erwarten Sie: Die globalen klimapolitischen Maßnahmen werden in den nächsten 10 Jahren im Vergleich zu heute…
				`,
				"en": `
					Over the next 10 years, compared to today, global climate policies will be…
				`,
			}.Outline("5.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 2

		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5a,
				labelsCertainty(),
				[]string{"ssq6"},
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
			}.Outline("6.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 2

		}

	}

	// page 4

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
		page.NavigationCondition = "fmt202512Include"

		{

			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "previousPageAnswer"
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			{

				gb := qst.NewGridBuilderRadios(
					columnTemplate6a,
					labelsImpact(),
					[]string{"ssq6aa", "ssq6ab"},
					[]string{"strong_neg", "neg", "neutral", "pos", "strong_pos", "no_answer"},
					[]trl.S{
						{
							"de": "In den nächsten 5 Jahren",
							"en": "In the next 5 years ",
						},
						{
							"de": "In den nächsten 25 Jahren",
							"en": "In the next 25 years",
						},
					},
				)

				gb.MainLabel = trl.S{
					"de": `
					Wie wird sich der Klimawandel unter Berücksichtigung Ihrer Erwartungen an die künftige Klimapolitik Ihrer Meinung nach auf das Wirtschaftswachstum in der EU auswirken?
				`,
					"en": `
					Considering your expectations of future climate policies, how do you think climate change will impact economic growth in the EU?
				`,
				}.Outline("6a.")
				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 2

				gr.Style = css.NewStylesResponsive(gr.Style)
				gr.Style.Desktop.StyleGridContainer.GapRow = "0.9rem"
				gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
				gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

				// gr.Class = "grid-2025-12-ssq1"
				//
				// rad.StyleLbl.Desktop.StyleText.AlignHorizontal = "justify"
			}
		}

	}

	// page 5
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
					Bitte geben Sie an, inwieweit Sie den folgenden Aussagen zustimmen.
				`,
					"en": `
					Please indicate to what extent you agree with the following statements.
				`,
				}.Outline("7.")
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 7
			gr.BottomVSpacers = 0

			// equal to below
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.Display = "grid"
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
			gr.Style.Mobile.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"

			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
			gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

			gr.Style.Desktop.StyleText.FontSize = 90

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
				inp.Label = trl.S{
					"de": "&nbsp;",
					"en": "&nbsp;",
				}
			}
			for i := 0; i < len(labelsAgree()); i++ {
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
				inp.Label = labelsAgree()[i]
				inp.LabelCenter()
				inp.LabelBottom()
			}
		}

		//
		//
		lblsRandomized := []trl.S{
			{
				"de": "Der Klimawandel stellt ein bedeutendes Problem für Volkswirtschaften und Finanzmärkte dar.",
				"en": "Climate change represents a significant issue for economies and financial markets.",
			},
			{
				"de": "Mit den richtigen Maßnahmen ist es möglich, bis 2050 eine klimaneutrale Wirtschaft zu erreichen.",
				"en": "With the right measures, it is possible to achieve a climate-neutral economy by 2050.",
			},
			{
				"de": "Die Wirtschaft kann klimaneutral werden und dabei weiterwachsen.",
				"en": "The economy can become climate-neutral while growing at the same time.",
			},
			{
				"de": "Solange keine geeignete Ersatztechnologie verfügbar ist, sollten weiterhin Investitionen in emissionsintensive Sektoren fließen.",
				"en": "As long as there is no suitable replacement technology, investment should still flow into emissions-intensive industries.",
			},
			{
				"de": "Die Bewältigung des Klimawandels erfordert, dass emissionsintensive Unternehmen über die nötigen Finanzmittel verfügen, um auf emissionsarme Technologien umzustellen.",
				"en": "Responding to climate change requires that emissions-intensive companies have the funding to transition to low-emission technologies.",
			},
			{
				"de": "Die Bewältigung des Klimawandels erfordert den Rückbau emissionsintensiver Sektoren und den Ausbau emissionsarmer Sektoren.",
				"en": "Responding to climate change requires shrinking emissions-intensive industries and growing low-emissions industries.",
			},
			{
				"de": "Technologische Innovation wird der entscheidende Faktor für das Erreichen einer klimaneutralen Wirtschaft sein.",
				"en": "Technological innovation will be the decisive determinant of achieving a climate-neutral economy. ",
			},
		}
		for i1 := 0; i1 < len(lblsRandomized); i1++ {
			gr := page.AddGroup()
			firstCol := float32(1)
			gr.Cols = firstCol + 6
			gr.RandomizationGroup = 2
			gr.BottomVSpacers = 0
			if i1 == (len(lblsRandomized) - 1) {
				gr.BottomVSpacers = 2 // bad, because of shuffling
				gr.BottomVSpacers = 0
			}

			// equal to above
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.Display = "grid"
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
			gr.Style.Mobile.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"

			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
			gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

			// distinct
			gr.Style.Desktop.StyleBox.Margin = "0 0 0.6rem" // bottom margin

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = firstCol
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
				inp.Label = lblsRandomized[i1]
			}
			for i2 := 0; i2 < 6; i2++ {
				{
					inp := gr.AddInput()
					inp.Type = "radio"
					inp.Name = fmt.Sprintf("ssq7%v", i1+1)
					inp.ValueRadio = fmt.Sprintf("%v", i2+1)
					inp.ColSpan = 1
					inp.ColSpanLabel = 0
					inp.ColSpanControl = 1
				}
			}
		}

		//
		placeHolder := trl.S{"de": "00", "en": "00"}

		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.WidthMax("48rem")
			gr.BottomVSpacers = 2

			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.Margin = "3rem 0 0" // top margin - because randomized previous block...

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `
						Für jeden Euro, der im Kontext des Klimawandels investiert wird: Wie viel sollte in Klimaschutz und wie viel in Klimaanpassung investiert werden?
					`,
					"en": `
						For every euro invested in the context of climate change, how much should be invested in climate mitigation versus climate adaptation?
					`,
				}.Outline("8.")
			}

			lbls := []trl.S{
				{
					"de": "Klimaschutz (d. h. Vermeidung oder Minderung von Treibhausgasemissionen)",
					"en": "Climate mitigation (i.e., preventing or reducing greenhouse gas emissions)",
				},
				{
					"de": "Klimaanpassung (d. h. Anpassung an die Auswirkungen des Klimawandels)",
					"en": "Climate adaptation (i.e., adjusting to the effects of climate change)",
				},
				{
					"de": "Es sollte kein Geld im Kontext des Klimawandels investiert werden.",
					"en": "No money should be invested in the context of climate change",
				},
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "ssq8a"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Min = 0
				inp.Max = 100
				inp.Step = 1
				inp.MaxChars = 4
				inp.Placeholder = placeHolder
				inp.Label = lbls[0]
				inp.ControlCenter()
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "ssq8b"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Min = 0
				inp.Max = 100
				inp.Step = 1
				inp.MaxChars = 4
				inp.Placeholder = placeHolder
				inp.Label = lbls[1]
				inp.ControlCenter()
			}
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "ssq8c"
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Label = lbls[2]
				inp.ControlRight()
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.WidthMax("48rem")
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `
						Für jeden Euro, der in grüne Technologien investiert wird: Wie viel sollte in neue grüne Technologien und wie viel in die Skalierung bestehender grüner Technologien investiert werden?
					`,
					"en": `
						For every euro invested in green technologies, how much should be invested in new green technologies versus scaling up existing green technologies?
					`,
				}.Outline("9.")
			}

			lbls := []trl.S{
				{
					"de": "Neue grüne Technologien",
					"en": "New green technologies",
				},
				{
					"de": "Skalierung bestehender grüner Technologien",
					"en": "Scaling up existing green technologies",
				},
				{
					"de": "Es sollte kein Geld in grüne Technologien investiert werden.",
					"en": "No money should be invested in green technologies",
				},
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "ssq9a"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Min = 0
				inp.Max = 100
				inp.Step = 1
				inp.MaxChars = 4
				inp.Placeholder = placeHolder
				inp.Label = lbls[0]
				inp.ControlCenter()
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "ssq9b"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Min = 0
				inp.Max = 100
				inp.Step = 1
				inp.MaxChars = 4
				inp.Placeholder = placeHolder
				inp.Label = lbls[1]
				inp.ControlCenter()
			}
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "ssq9c"
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Label = lbls[2]
				inp.ControlRight()
				inp.ControlCenter()
				inp.StyleCtl.Desktop.StyleBox.Margin = "2rem 0 0 0"
				inp.StyleCtl.Mobile.StyleBox.Margin = "0"
			}

		}

	}

	return nil

}
