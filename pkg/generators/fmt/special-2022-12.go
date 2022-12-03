package fmt

import (
	"fmt"

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
			"en": "Special questions December 2022",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022",
			"en": "Special<br>End of 2022",
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
						Lieber Finanzmarktexperte, liebe Finanzmarktexpertin,<br>
						zum Jahresende 2022 möchten wir Sie im Rahmen unseres Sonderfragenteils ausführlicher als sonst 
						zu Ihrem Ausblick für die deutsche Wirtschaft, 
						die Inflationsentwicklung in Deutschland und die Entwicklung des DAX befragen.
						</p>
						<p style=''>
						Die Ergebnisse werden wir Ihnen in unserem Finanzmarktreport in besonders ausführlicher Form Verfügung stellen.
						</p>
						<p style=''>
						Vielen Dank für Ihre Teilnahme. 
						</p>
					`,

					"en": `
						<p style=''>
						Dear expert,<br>
						</p>
						<p style=''>
						as we approach the end of 2022, 
						we would like to ask you in more detail than usual about your outlook for the German economy, 
						the development of inflation in Germany 
						and the development of the DAX as part of our special question section.				
						</p>
						<p style=''>
						Thank you very much for participating.
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
			"en": "Special questions December 2022 - 2",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 2",
			"en": "Special<br>End of 2022 - 2",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("54rem")

		lblsColsQ1Q2 := []trl.S{
			{
				"de": "&nbsp;",
				"en": "&nbsp;",
			},
			{
				"de": "Untergrenze des 90-Prozent-Konfidenzintervalls",
				"en": "lower bound of the 90 percent confidence interval",
			},
			{
				"de": "Obergrenze des 90-Prozent-Konfidenzintervalls",
				"en": "upper bound of the 90 percent confidence interval",
			},
			{
				"de": "keine Angabe",
				"en": "no estimate",
			},
		}

		// q1
		{
			lblMain := trl.S{
				"de": `
					Mit einer Wahrscheinlichkeit von 90 Prozent werden die jährliche 
					<i>Inflationsrate in Deutschland</i> 
					(durchschnittliche jährliche Veränderung des HICP in Prozent) bzw. die jährliche 
					<i>Wachstumsrate des realen Bruttoinlandprodukts in Deutschland 2023</i> 
					bzw. 
					<i>im Zeitraum 2023&#8209;2025</i> 
					zwischen den folgenden Werten liegen:
				`,
				"en": `
					With a probability of 90 per cent, the annual 
					<i>inflation rate in Germany</i> 
					(annual average change of the HICP, in percent) and the annual growth rate of real 
					<i>German GDP</i> 
					for the year 
					<i>2023</i> 
					and the 
					<i>years between 2023 and 2025</i> 
					will lie between the following values:
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
					"en": "Inflation rate, Germany, 2023",
				},
				{
					"de": "Durchschn. Inflationsrate in Deutschland pro Jahr, 2023&#8209;2025",
					"en": "Avg. annual inflation rate, Germany per year, 2023&#8209;2025",
				},
				{
					"de": "BIP-Wachstumsrate in Deutschland, 2023",
					"en": "Growth rate of annual real German GDP, 2023",
				},
				{
					"de": "Durchschn. BIP-Wachstumsrate in Deutschland pro Jahr, 2023&#8209;2025",
					"en": "Avg. growth rate of annual real German GDP per year, 2023&#8209;2025",
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
					Mit einer Wahrscheinlichkeit von 90 Prozent wird die jährliche 
					<i>Rendite des DAX 2023</i> 
					bzw. 
					<i>im Zeitraum 2023&#8209;2025</i> 
					zwischen den folgenden Werten liegen:
					<br><br>
					Hinweis: Im Zeitraum 2011-2021 betrug die jährliche DAX-Rendite im Durchschnitt 8,9 Prozent mit einer Standardabweichung von 14,7 Prozent.
				`,
				"en": `
					With a probability of 90 percent, the 
					<i>annual return of the DAX</i> 
					for the year 
					<i>2023</i> 
					and the 
					<i>years between 2023 and 2025</i> 
					will lie between the following values:
					<br><br>
					Note: Between 2011 and 2021, the average annual DAX return was 8.9 percent with a standard deviation of 14.7 percent.
				`,
			}.Outline("2.")

			inpNames := []string{
				"qs2_dax_12m",
				"qs2_dax_36m",
			}

			lblsRows := []trl.S{
				{
					"de": "DAX-Rendite, 2023",
					"en": "DAX return, 2023",
				},
				{
					"de": "Durchschn. DAX-Rendite pro Jahr, 2023&#8209;2025",
					"en": "Average DAX return per year, 2023&#8209;2025",
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

		// gr3 - q3a dynamic
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.DynamicFuncParamset = ""
				inp.DynamicFunc = fmt.Sprintf("Special202212Q3__%v__%v", 0, 0)

				inp.ColSpanControl = 1
			}
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
					"en": "agree",
				},
				{
					"de": "stimme weder zu noch lehne ab",
					"en": "undecided",
				},
				{
					"de": "stimme nicht zu",
					"en": "disagree",
				},
				{
					"de": "stimme überhaupt nicht zu",
					"en": "strongly disagree",
				},

				{
					"de": "keine<br>Angabe",
					"en": "no answer",
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
					"qs3a_inf_narrative_a_static",
					"qs3a_inf_narrative_b_static",
					"qs3a_inf_narrative_c_static",
					"qs3a_inf_narrative_d_static",
					"qs3a_inf_narrative_e_static",
					"qs3a_inf_narrative_f_static",
				},
				radioVals6,
				[]trl.S{
					{
						"de": `Eine Entspannung bei der Inflationsentwicklung, eine weniger restriktive Geldpolitik der EZB und nachlassende Rezessionsrisiken wirken sich 
								<i>positiv</i> 
								auf das Rendite-Risiko-Profil in 2023 aus.`,
						"en": `An easing in the development of inflation development, a less restrictive monetary stance by the ECB and diminishing recession risks have a 
								<i>positive</i> 
								impact on the return-risk-profile in 2023.`,
					},
					{
						"de": `Den DAX-Konzernen gelingt es auch weiterhin, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen unverändert bleiben oder sogar steigen, was sich 
								<i>positiv</i> 
								auf das Rendite-Risiko-Profil des DAX in 2023 auswirkt. `,
						"en": `DAX companies will continue to succeed in passing on their rising costs to their customers. The profit margins of DAX companies will therefore remain unchanged or even increase, which has a 
								<i>positive</i> 
								impact on the return-risk-profile of the DAX in 2023.`,
					},
					{
						"de": `Die Entwicklung der Inflation spielt für das Rendite-Risiko-Profil des DAX in 2023 
								<i>keine Roll</i>e
								.`,
						"en": `The development of inflation does 
								<i>not impact</i> 
								the return-risk-profile of the DAX.`,
					},
					{
						"de": `	<i>Positive</i> 
								und 
								<i>negative</i> 
								Effekte der Inflation gleichen sich aus. Die Entwicklung der Inflation ist daher insgesamt 
								<i>neutral</i> 
								für das Rendite-Risiko-Profil des DAX in 2023.`,
						"en": `
								<i>Positive</i> 
								and 
								<i>negative</i> 
								effects of inflation cancel each other out. Overall, the development of inflation is 
								<i>neutral</i> 
								for the return-risk-profile of the DAX in 2023.`,
					},
					{
						"de": `Den DAX-Konzernen gelingt es nicht, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen fallen, was sich 
								<i>negativ</i> 
								auf das Rendite-Risiko-Profil des DAX in 2023 auswirkt.`,
						"en": `DAX companies will not to succeed in passing on their rising costs to their customers. The profit margins of DAX companies will therefore decrease, which has a 
								<i>negative</i> 
								impact on the return-risk-profile of the DAX in 2023.`,
					},
					{
						"de": `Anhaltend hohe Inflationsraten, weitere Zinserhöhungen durch die EZB und zunehmende Rezessionsrisiken wirken sich 
								<i>negativ</i> 
								auf das Rendite-Risiko-Profil des DAX in 2023 aus.
						`,
						"en": `Persistently high inflation rates, further interest rate hikes by the ECB and increasing recession risks will have a 
								<i>negative</i> 
								impact on the return-risk-profile of the DAX in 2023.`,
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
						Wie beurteilen Sie die folgenden Aussagen zum Zusammenhang zwischen der 
						<i>Inflationsentwicklung</i>
						und dem 
						<i>Rendite-Risiko-Profil des DAX</i>
						in 
						<i>2023</i>
						?
					`,
				"en": `
						Do you agree or disagree with the following statements about the relationship between the 
						<i>developments of inflation </i>
						and the 
						<i>return-risk-profile of the DAX</i>
						in 
						<i>2023</i>
						?
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
					"de": `Eine Entspannung bei der Inflationsentwicklung, eine weniger restriktive Geldpolitik der EZB und nachlassende Rezessionsrisiken wirken sich 
							<i>positiv</i> 
							auf das Rendite-Risiko-Profil in 2023 aus.`,
					"en": `An easing in the development of inflation development, a less restrictive monetary stance by the ECB and diminishing recession risks have a 
							<i>positive</i> 
							impact on the return-risk-profile in 2023.`,
				},
				{
					"de": `Den DAX-Konzernen gelingt es auch weiterhin, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen unverändert bleiben oder sogar steigen, was sich 
							<i>positiv</i> 
							auf das Rendite-Risiko-Profil des DAX in 2023 auswirkt. `,
					"en": `DAX companies will continue to succeed in passing on their rising costs to their customers. The profit margins of DAX companies will therefore remain unchanged or even increase, which has a 
							<i>positive</i> 
							impact on the return-risk-profile of the DAX in 2023.`,
				},
				{
					"de": `Die aktuelle Entwicklung der Inflation spielt für das Rendite-Risiko-Profil des DAX in 2023 
							<i>keine Roll</i>e
							.`,
					"en": `The development of inflation does 
							<i>not impact</i> 
							the return-risk-profile of the DAX.`,
				},
				{
					"de": `
							<i>Positive</i> 
							und 
							<i>negative</i> 
							Effekte der Inflation gleichen sich aus. Die aktuelle Entwicklung der Inflation ist daher insgesamt 
							<i>neutral</i> 
							für das Rendite-Risiko-Profil des DAX in 2023.`,
					"en": `
							<i>Positive</i> 
							and 
							<i>negative</i> 
							effects of inflation cancel each other out. Overall, the development of inflation is 
							<i>neutral</i> 
							for the return-risk-profile of the DAX in 2023.`,
				},
				{
					"de": `Den DAX-Konzernen gelingt es nicht, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen fallen, was sich 
							<i>negativ</i> 
							auf das Rendite-Risiko-Profil des DAX in 2023 auswirkt.`,
					"en": `DAX companies will not to succeed in passing on their rising costs to their customers. The profit margins of DAX companies will therefore decrease, which has a 
							<i>negative</i> 
							impact on the return-risk-profile of the DAX in 2023.`,
				},
				{
					"de": `Anhaltend hohe Inflationsraten, weitere Zinserhöhungen durch die EZB und zunehmende Rezessionsrisiken wirken sich 
							<i>negativ</i> 
							auf das Rendite-Risiko-Profil des DAX in 2023 aus.`,
					"en": `Persistently high inflation rates, further interest rate hikes by the ECB and increasing recession risks will have a 
							<i>negative</i> 
							impact on the return-risk-profile of the DAX in 2023.`,
				},
				{
					"de": `Keine Antwort`,
					"en": `No answer`,
				},
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Bitte wählen Sie nun aus den folgenden Aussagen diejenige aus, die ihrer Meinung nach den Zusammenhang zwischen der 
						<i>Inflationsentwicklung </i>
						und dem 
						<i>Rendite-Risiko-Profil des DAX </i>
						in 
						<i>2023 </i>
						am besten widerspiegelt:
					`,
					"en": `
						From the following statements, please select the one that, in your opinion, best reflects the relationship between the 
						<i>development of inflation </i>
						and the 
						<i>risk-return profile of the DAX</i>
						 in 
						 <i>2023</i>
						 :

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
			"en": "Special questions December 2022 - 3",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 3",
			"en": "Special<br>End of 2022 - 3",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("46rem")
		page.WidthMax("54rem")

		//
		// gr1 - q4a
		{

			colLblsQ4 := []trl.S{
				{
					"de": "Meinen eigenen Analysen",
					"en": "My own analyses",
				},
				{
					"de": "Analysen von Experten/-innen aus meinem Unternehmen",
					"en": "Analyses by experts in my company",
				},
				{
					"de": "Analysen aus externen Quellen",
					"en": "Analyses from external sources",
				},

				{
					"de": "keine<br>Angabe",
					"en": "no answer",
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
					"qs4a_growth",
					"qs4a_inf",
					"qs4a_dax",
				},
				radioVals4,
				[]trl.S{
					{
						"de": `Wirtschaftswachstum Deutschland`,
						"en": `GDP growth, Germany`,
					},
					{
						"de": `Inflation in Deutschland`,
						"en": `Inflation, Germany`,
					},
					{
						"de": `Entwicklung des DAX`,
						"en": `Developments of the DAX`,
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
						Meine Einschätzungen mit Blick auf die folgenden Bereiche beruhen hauptsächlich auf
					`,
				"en": `
						My expectations with respect to the following areas are mainly based on
					`,
			}.Outline("4a.")

			gr := page.AddGrid(gb)
			_ = gr
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
		}

		//
		// gr2 - q4b
		{

			colLbls4b := []trl.S{
				{
					"de": "nicht relevant",
					"en": "not relevant",
				},
				{
					"de": "leicht relevant",
					"en": "slightly relevant",
				},
				{
					"de": "stark relevant",
					"en": "highly relevant",
				},

				{
					"de": "keine<br>Angabe",
					"en": "no answer",
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
				colLbls4b,
				[]string{
					"qs4b_relevance",
				},
				radioVals4,
				[]trl.S{
					{
						"de": `Wie relevant sind die Prognosen der Bundesbank für Ihre eigenen Inflationsprognosen für Deutschland?`,
						"en": `How relevant are the inflation forecasts of Bundesbank for your own inflation forecasts for Germany?`,
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
					Bundesbankpräsident Joachim Nagel äußert sich regelmäßig zum Inflationsausblick für Deutschland. Im November 2022 äußerte er sich folgendermaßen: "Auch im kommenden Jahr dürfte die Inflationsrate in Deutschland hoch bleiben. Ich halte es für wahrscheinlich, dass im Jahresdurchschnitt 2023 eine sieben vor dem Komma stehen wird".
						`,
				"en": `
					Bundesbank president Joachim Nagel regularly comments on the inflation outlook for Germany. In November 2022, he commented as follows: "The inflation rate in Germany is likely to remain high in the coming year. I believe it is likely that the annual average for 2023 will have a seven before the decimal point."
					`,
			}.Outline("4b.")

			gr := page.AddGrid(gb)
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "1.2rem"
			_ = gr
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1
		}

		//
		// gr3 - q4c
		{

			colLbls4c := []trl.S{
				{
					"de": "ja",
					"en": "yes",
				},
				{
					"de": "nein",
					"en": "no",
				},
				{
					"de": "keine<br>Angabe",
					"en": "no answer",
				},
			}

			var columnTemplateLocal = []float32{
				5.0, 1,
				0.0, 1,
				0.5, 1,
			}

			lbl1 := trl.S{
				"de": `
					War Ihnen die Aussage von Bundesbankpräsident Joachim Nagel bereits bekannt?
						`,
				"en": `
					Were you aware of this statement by Bundesbank president Joachim Nagel?
					`,
			}.Outline("4c.")

			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				colLbls4c,
				[]string{
					"qs4c_known",
				},
				radioVals4,
				[]trl.S{
					lbl1,
				},
			)

			// gb.MainLabel =

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
			"en": "Special questions December 2022 - 4",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 4",
			"en": "Special<br>End of 2022 - 4",
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
					"de": `Die Wachstumsrate des 
							<i>realen Bruttoinlandprodukts</i> 
							in 
							<i>Deutschland</i> 
							in 
							<i>2023</i> 
							erwarte ich bei `,
					"en": `I expect the growth rate of annual real 
							<i>German GDP</i> 
							in 
							<i>2023</i> 
							to come in at `,
				}.Outline("5a.")
				inp.LabelPadRight()
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs5a_growth"
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
					"de": `Die Wahrscheinlichkeit einer 
							<i>Rezession in Deutschland in 2023</i> 
							beträgt`,
					"en": `The probability of a 
							<i>recession</i> 
							in 
							<i>Germany</i> 
							in 
							<i>2023</i> 
							is `,
				}.Outline("5b.")
				inp.LabelPadRight()
				inp.Suffix = trl.S{"de": "%", "en": "pct"}

				inp.Name = "qs5b_recession"
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
					"de": `Die 
							<i>jährliche Inflationsrate</i> 
							in Deutschland (durchschnittliche jährliche Veränderung des HICP in Prozent) 
							<i>für 2023</i> 
							erwarte ich bei 
							`,
					"en": `My forecast for the
							<i>annual inflation rate in Germany</i>
							(annual average change of the HICP, in percent)
							<i>in 2023</i>
							is
							`,
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
					"de": `Mit Blick auf das Jahr 2023, was sind Ihre Prognosen für die 
							<i>Rendite</i>
							des 
							<i>DAX</i>
							bzw. die 
							<i>Volatilität</i>
							des 
							<i>DAX</i>
							(Standardabweichung der jährlichen DAX-Renditen)?`,
					"en": `For the year 2023, what are your forecasts for the 
							<i>return of the DAX</i>
							and 
							<i>volatility</i>
							of the 
							<i>DAX</i>
							(standard deviation of the annual DAX returns)?`,
				}.Outline("7.")
			}

			lblsCols := []trl.S{
				{
					"de": "Punktprognose in Prozent",
					"en": "point forecast in percent",
				},
				{
					"de": "keine Angabe",
					"en": "no estimate",
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
				"qs7_dax_return_12m",
				"qs7_dax_vola_12m",
			}

			lblsRows := []trl.S{
				{
					"de": "DAX Rendite, 2023",
					"en": "DAX return, 2023",
				},
				{
					"de": "DAX Volatilität, 2023",
					"en": "DAX volatility, 2023",
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
					"en": "My forecast for the <i>ECB&#39;s main refinancing rate at the end of 2023</i> is  ",
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
						Nehmen Sie an, dass Sie am 01.01.2023 1 Million Euro 
						<i>über die nächsten zwölf Monat</i>e
						 in ein Portfolio bestehend aus dem 
						 <i>DAX</i> 
						 und einer 
						 <i>risikolosen Anlage</i> 
						 mit jährlicher Verzinsung von 2 Prozent anlegen. Wie groß wäre der Anteil, den Sie persönlich in der aktuellen Situation in den 
						 <i>DAX</i> 
						 investieren würden?
						<br>
						Anteil DAX:
					`,

					"en": `
						Assume that on January 1, 2023 you invest 1 million euros 
						<i>over the next twelve months</i> 
						in a portfolio consisting of the 
						<i>DAX</i> 
						and a 
						<i>risk-free investment</i> 
						with an annual interest rate of 2 percent. What proportion would you personally invest in the 
						<i>DAX</i> 
						in the current situation?
						<br>
						Share DAX:
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
