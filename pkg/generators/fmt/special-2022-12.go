package fmt

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202212(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 12
	if !cond {
		return nil
	}

	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderbefragung zum Jahresende 2022",
			"en": "Special end of year 2022",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022",
			"en": "Special<br>end of 2022",
		}
		page.WidthMax("46rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
				<p style=''>
					Todo Einleitungstext
				</p>
					`,

					"en": `
				<p style=''>
					Todo Einleitungstext
				</p>
					`,
				}
				inp.ColSpanLabel = 1
			}

		}

	}
	//
	//
	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderbefragung zum Jahresende 2022 - 2",
			"en": "Special end of year 2022 - 2",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 2",
			"en": "Special<br>end of 2022 - 2",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("46rem")
		page.WidthMax("54rem")

		lblsColsQ1Q2 := []trl.S{
			{
				"de": "&nbsp;",
				"en": "&nbsp;",
			},
			{
				"de": "Untergrenze des 90-Prozent-Konfidenzintervalls",
				"en": "Untergrenze des 90-Prozent-Konfidenzintervalls",
			},
			{
				"de": "Obergrenze des 90-Prozent-Konfidenzintervalls",
				"en": "Obergrenze des 90-Prozent-Konfidenzintervalls",
			},
			{
				"de": "keine Angabe",
				"en": "keine Angabe",
			},
		}

		// q1
		{
			lblMain := trl.S{
				"de": `
					Mit einer Wahrscheinlichkeit von 90 Prozent werden die durchschnittliche 
					<i>jährliche Inflationsrate in Deutschland</i> (durchschnittliche jährliche Veränderung des HICP in Prozent) 
					bzw. die durchschnittliche jährliche 
					<i>Wachstumsrate des realen Bruttoinlandprodukts</i> in Deutschland  

					Sicht von <i>zwölf Monaten</i> bzw. <i>drei Jahren</i>   

					zwischen den folgenden Werten liegen:
				`,
				"en": `
					todo
				`,
			}.Outline("1.")

			inpNames := []string{
				"qs1_inflation_12m",
				"qs1_inflation_36m",
				"qs1_growth_12m",
				"qs1_growth_36m",
			}

			lblsRows := []trl.S{
				{
					"de": "Inflationsrate in Deutschland, 2023",
					"en": "Inflationsrate in Deutschland, 2023",
				},
				{
					"de": "Durchschn. Inflationsrate in Deutschland, 2023-2025",
					"en": "Durchschn. Inflationsrate in Deutschland, 2023-2025",
				},
				{
					"de": "BIP-Wachstumsrate in Deutschland, 2023",
					"en": "BIP-Wachstumsrate in Deutschland, 2023",
				},
				{
					"de": "Durchschn. BIP-Wachstumsrate in Deutschland, 2023-2025",
					"en": "Durchschn. BIP-Wachstumsrate in Deutschland, 2023-2025",
				},
			}

			matrixOfPercentageInputs(
				qst.WrapPageT(page),
				lblMain,
				lblsColsQ1Q2,
				inpNames,
				lblsRows,
			)
		}

		// q2
		{
			lblMain := trl.S{
				"de": `
					Mit einer Wahrscheinlichkeit von 90 Prozent wird die durchschnittliche jährliche <i>Rendite des DAX</i> auf Sicht von 
					<i>zwölf Monaten</i> bzw. 
					<i>drei Jahren</i> zwischen den folgenden Werten liegen:
					<br>
					Hinweis: Im Zeitraum 2011-2021 betrug die jährliche DAX-Rendite im Durchschnitt 8,9 Prozent mit einer Standardabweichung von 14,7 Prozent.
				`,
				"en": `
					todo
				`,
			}.Outline("2.")

			inpNames := []string{
				"qs2_dax_12m",
				"qs2_dax_36m",
			}

			lblsRows := []trl.S{
				{
					"de": "DAX-Rendite, auf Sicht von 12&nbsp;Monaten",
					"en": "DAX-Rendite, auf Sicht von 12&nbsp;Monaten",
				},
				{
					"de": "DAX-Rendite, auf Sicht von 3&nbsp;Jahren",
					"en": "DAX-Rendite, auf Sicht von 3&nbsp;Jahren",
				},
			}

			matrixOfPercentageInputs(
				qst.WrapPageT(page),
				lblMain,
				lblsColsQ1Q2,
				inpNames,
				lblsRows,
			)

		}

		//
		// gr3 - q3a
		{

			colLblsQ3a := []trl.S{
				{
					"de": "stimme voll zu",
					"en": "strongly agree",
				},
				{
					"de": "stimme zu",
					"en": "Agree",
				},
				{
					"de": "stimme weder zu noch lehne ab",
					"en": "Undecided",
				},
				{
					"de": "stimme nicht zu",
					"en": "Disagree",
				},
				{
					"de": "stimme überhaupt nicht zu",
					"en": "strongly disagree",
				},

				{
					"de": "keine<br>Angabe",
					"en": "No answer",
				},
			}

			var columnTemplateLocal = []float32{
				4.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				colLblsQ3a,
				[]string{
					"qs3a_inf_narrative_a",
					"qs3a_inf_narrative_b",
					"qs3a_inf_narrative_c",
					"qs3a_inf_narrative_d",
					"qs3a_inf_narrative_e",
					"qs3a_inf_narrative_f",
				},
				radioVals6,
				[]trl.S{
					{
						"de": `Eine Entspannung bei der Inflationsentwicklung, ein vorsichtigeres Vorgehen der EZB und nachlassende Rezessionsrisiken wirken sich positiv auf das Rendite-Risiko-Profil in 2023 aus.`,
						"en": `todo`,
					},
					{
						"de": `Den DAX-Konzernen gelingt es, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen unverändert bleiben oder sogar steigen, was sich positiv auf  das Rendite-Risiko-Profil des DAX in 2023 auswirkt. `,
						"en": `todo`,
					},
					{
						"de": `Die aktuelle Entwicklung der Inflation spielt für das Rendite-Risiko-Profil des DAX in 2023 
									<i>keine Rolle</i>.`,
						"en": `todo`,
					},
					{
						"de": `Positive und negative Effekte der Inflation gleichen sich aus. 
									Die aktuelle Entwicklung der Inflation ist daher insgesamt 
									<i>neutral</i> für das Rendite-Risiko-Profil des DAX in 2023.`,
						"en": `todo`,
					},
					{
						"de": `Den DAX-Konzernen gelingt es nicht, ihre steigenden Kosten an ihre Kunden weiterzugeben. 
									Die Gewinn-Margen der DAX-Konzerne werden deswegen fallen, 
									was sich <i>negativ</i> auf  das Rendite-Risiko-Profil des DAX in 2023 auswirkt.`,
						"en": `todo`,
					},
					{
						"de": `Anhaltend hohe Inflationsraten, 
								weitere Zinserhöhungen durch die EZB und zunehmende Rezessionsrisiken wirken sich
								 negativ auf das Rendite-Risiko-Profil des DAX in 2023 aus.`,
						"en": `todo`,
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
						Wie beurteilen Sie die folgenden Aussagen zum Zusammenhang zwischen der <i>Inflationsentwicklung</i>
						 und dem 
						<i>Rendite-Risiko-Profil</i> des DAX 
						 in 
						<i>2023</i>?
					`,
				"en": `
						todo
					`,
			}.Outline("3a.")

			gr := page.AddGrid(gb)
			_ = gr
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
		}

		//
		// gr4 - q3b
		{

			gr := page.AddGroup()
			gr.Cols = 1

			radioVals := []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
				"no_answer",
			}

			lbls := []trl.S{
				{
					"de": `Eine Entspannung bei der Inflationsentwicklung, ein vorsichtigeres Vorgehen der EZB und nachlassende Rezessionsrisiken wirken sich positiv auf das Rendite-Risiko-Profil in 2023 aus.`,
					"en": `todo`,
				},
				{
					"de": `Den DAX-Konzernen gelingt es, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen unverändert bleiben oder sogar steigen, was sich positiv auf  das Rendite-Risiko-Profil des DAX in 2023 auswirkt. `,
					"en": `todo`,
				},
				{
					"de": `Die aktuelle Entwicklung der Inflation spielt für das Rendite-Risiko-Profil des DAX in 2023 keine Rolle.`,
					"en": `todo`,
				},
				{
					"de": `Positive und negative Effekte der Inflation gleichen sich aus. Die aktuelle Entwicklung der Inflation ist daher insgesamt neutral für das Rendite-Risiko-Profil des DAX in 2023.`,
					"en": `todo`,
				},
				{
					"de": `Den DAX-Konzernen gelingt es nicht, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen fallen, was sich negativ auf  das Rendite-Risiko-Profil des DAX in 2023 auswirkt.`,
					"en": `todo`,
				},
				{
					"de": `Anhaltend hohe Inflationsraten, weitere Zinserhöhungen durch die EZB und zunehmende Rezessionsrisiken wirken sich negativ auf das Rendite-Risiko-Profil des DAX in 2023 aus.`,
					"en": `todo`,
				},
				{
					"de": `Keine Antwort`,
					"en": `todo`,
				},
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Bitte wählen Sie nun aus den folgenden Aussagen diejenige aus, 
						die ihrer Meinung nach den Zusammenhang zwischen 
						der 
						<i>Inflationsentwicklung</i> und dem 
						<i>Rendite-Risiko-Profil des DAX</i> in 
						<i>2023</i> am besten widerspiegelt: 
					`,
					"en": `
						todo
					`,
				}.Outline("3b.")
				inp.ColSpan = gr.Cols
			}

			for i := 0; i < len(radioVals); i++ {
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = "qs3b_inf_to_dax"
				inp.Label = lbls[i]
				inp.ValueRadio = radioVals[i]
				inp.ColSpan = 1

				inp.ColSpanLabel = 1
				inp.ColSpanControl = 9

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "justify"

				inp.ControlFirst()

				inp.ControlTop()
				inp.StyleCtl.Desktop.StyleBox.Margin = "0.12rem 0 0 0"

			}

		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderbefragung zum Jahresende 2022 - 3",
			"en": "Special end of year 2022 - 3",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 3",
			"en": "Special<br>end of 2022 - 3",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("46rem")
		page.WidthMax("54rem")

		//
		// gr1 - q4
		{

			colLblsQ4 := []trl.S{
				{
					"de": "meinen eigenen Analysen",
					"en": "todo",
				},
				{
					"de": "Analysen von Experten/-innen aus meinem Unternehmen",
					"en": "todo",
				},
				{
					"de": "Analysen aus externen Quellen",
					"en": "todo",
				},

				{
					"de": "keine<br>Angabe",
					"en": "No answer",
				},
			}

			var columnTemplateLocal = []float32{
				4.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				colLblsQ4,
				[]string{
					"qs4_growth",
					"qs4_inf",
					"qs4_dax",
				},
				radioVals4,
				[]trl.S{
					{
						"de": `Wirtschaftswachstum Deutschland`,
						"en": `todo`,
					},
					{
						"de": `Inflation in Deutschland`,
						"en": `todo`,
					},
					{
						"de": `Entwicklung des DAX`,
						"en": `todo`,
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
						Meine Einschätzungen mit Blick auf die folgenden Bereiche beruhen hauptsächlich auf
					`,
				"en": `
						todo
					`,
			}.Outline("4.")

			gr := page.AddGrid(gb)
			_ = gr
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
		}

		//
		// gr2 - q4a

		{

			colLbls4a := []trl.S{
				{
					"de": "nicht relevant",
					"en": "todo",
				},
				{
					"de": "leicht relevant",
					"en": "todo",
				},
				{
					"de": "stark relevant",
					"en": "todo",
				},

				{
					"de": "keine<br>Angabe",
					"en": "No answer",
				},
			}

			var columnTemplateLocal = []float32{
				5.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				colLbls4a,
				[]string{
					"qs4a_relevance",
				},
				radioVals4,
				[]trl.S{
					{
						"de": `
							Wie relevant sind die Prognosen der Bundesbank für Ihre eigenen Inflationsprognosen für Deutschland?				`,
						"en": `todo`,
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
							Bundesbankpräsident Joachim Nagel äußert sich regelmäßig zum Inflationsausblick für Deutschland. 
							Anfang November 2022 äußerte er sich folgendermaßen: 
							"Auch im kommenden Jahr dürfte die Inflationsrate in Deutschland hoch bleiben. 
							Ich halte es für wahrscheinlich, dass im Jahresdurchschnitt 2023 eine sieben vor dem Komma stehen wird".						
						`,
				"en": `
						todo
					`,
			}.Outline("4a.")

			gr := page.AddGrid(gb)
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "1.2rem"
			_ = gr
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderbefragung zum Jahresende 2022 - 4",
			"en": "Special end of year 2022 - 4",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 4",
			"en": "Special<br>end of 2022 - 4",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("46rem")
		// page.WidthMax("54rem")

		{
			gr := page.AddGroup()
			gr.Cols = 3 //x
			// gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "number"

				inp.Label = trl.S{
					"de": "Die Wachstumsrate des realen Bruttoinlandprodukts in Deutschland in 2023 erwarte ich bei ",
					"en": "todo",
				}.Outline("5a.")
				inp.LabelPadRight()
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs51_growth"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5

				inp.ColSpan = 3
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"

				inp.Label = trl.S{
					"de": "Die Wahrscheinlichkeit einer <i>Rezession in Deutschland in 2023</i> beträgt ",
					"en": "todo",
				}.Outline("5b.")
				inp.LabelPadRight()
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs52_recession"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5

				inp.ColSpan = 3
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 3 //x
			// gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "number"

				inp.Label = trl.S{
					"de": "Die jährliche <i>Inflationsrate in Deutschland</i> (durchschnittliche jährliche Veränderung des HICP in Prozent) <i>für 2023</i> erwarte ich bei",
					"en": "todo",
				}.Outline("6.")
				inp.LabelPadRight()
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs6_infl"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5

				inp.ColSpan = 3
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 4 //x

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `Auf Sicht von 
							<i>zwölf Monaten</i>, 
							was sind Ihre Prognosen für die jährliche Rendite bzw. die Volatilität des DAX (Standardabweichung der jährlichen DAX-Renditen)?`,
					"en": `todo`,
				}.Outline("7.")
			}

			lblsCols := []trl.S{
				{
					"de": "Punktprognose in Prozent",
					"en": "todo",
				},
				{
					"de": "keine Angabe",
					"en": "todo",
				},
			}
			for i := 0; i < len(lblsCols); i++ {
				inp := gr.AddInput()
				inp.Type = "textblock"
				if i == 0 {
					inp.Type = "label-as-input"
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 1
				} else {
					inp.LabelCenter()
				}
				inp.Label = lblsCols[i]

				inp.ColSpan = 1
				if i == 0 {
					inp.ColSpan = 3
				}
			}

			inpNames := []string{
				"qs7_inflation_12m",
				"qs7_inflation_36m",
			}

			lblsRows := []trl.S{
				{
					"de": "Durchschn. DAX-Rendite, auf Sicht von 12 Monaten",
					"en": "todo",
				},
				{
					"de": "Volatilität, auf Sicht von 12 Monaten",
					"en": "todo",
				},
			}

			for i := 0; i < len(inpNames); i++ {

				{
					inp := gr.AddInput()
					inp.Type = "number"

					inp.Label = lblsRows[i]
					inp.Suffix = trl.S{"de": "%", "en": "pct"}

					inp.Name = inpNames[i]
					inp.Min = 0
					inp.Max = 100
					inp.Step = 0.1
					inp.MaxChars = 5

					inp.ControlCenter()

					inp.ColSpan = 3
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 1

				}

				{
					inp := gr.AddInput()
					inp.Type = "checkbox"
					inp.Name = inpNames[i] + "_noanswer"
					inp.ColSpan = 1
					inp.ColSpanControl = 1
					inp.ControlTopNudge()

				}

			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 3 //x
			// gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "number"

				inp.Label = trl.S{
					"de": "Den <i>Hauptrefinanzierungssatz der EZB</i> erwarte ich <i>Ende 2023</i> bei ",
					"en": "todo",
				}.Outline("8.")
				inp.LabelPadRight()
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs8_i"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5

				inp.ColSpan = 3
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 3 //x
			// gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Label = trl.S{
					"de": `
						Nehmen Sie an, dass Sie 1 Million Euro 
						<i>über die nächsten zwölf Monate</i> in ein Portfolio bestehend aus dem 
						<i>DAX</i> und einer 
						<i>risikolosen Anlage</i> mit jährlicher Verzinsung von 2 Prozent anlegen. 
						Wie groß wäre der Anteil, den Sie in der aktuellen Situation in den 
						<i>DAX</i> investieren würden?
						<br>
						Anteil DAX
					`,

					"en": `
					Todo Einleitungstext
					`,
				}.Outline("9.")

				// inp.Type = "textblock"
				// inp.ColSpan = gr.Cols

				inp.Type = "number"
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs9_sharedax"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5

				inp.ColSpan = 3
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1

				inp.ControlBottom()
				inp.LabelPadRight()

			}

		}

	}

	return nil
}
