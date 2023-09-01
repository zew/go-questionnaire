package fmt

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202309b(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2023 && q.Survey.Month == 9
	if !cond {
		return nil
	}

	//
	//
	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfragen zu Inflation, BIP-Wachstum und DAX-Renditen - Teil 2",
			"en": "Special questions: Inflation, GDP growth and DAX returns - part 2",
		}
		// page.Label = trl.S{
		// 	"de": " &nbsp; ",
		// 	"en": " &nbsp; ",
		// }
		page.Short = trl.S{
			"de": " ",
			"en": " ",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("56rem")

		rowLbls3and4 := []trl.S{
			{
				"de": `Eine Entspannung bei der Inflationsentwicklung, eine weniger restriktive Geldpolitik der EZB und nachlassende Rezessionsrisiken wirken sich
						<i>positiv</i>
						auf das Rendite-Risiko-Profil des DAX in 2023 aus.`,
				"en": `An easing in the development of inflation, a less restrictive monetary stance by the ECB and diminishing recession risks have a
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
		}

		//
		// gr1a q3 - header
		{

			colLblsQ3 := []trl.S{
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
				colLblsQ3,
				[]string{},
				radioVals6,
				[]trl.S{},
			)

			gb.MainLabel = trl.S{
				"de": `
						Wie beurteilen Sie die folgenden Aussagen zum Zusammenhang zwischen der 
						<i>Inflationsentwicklung</i>
						und dem 
						<i>Rendite-Risiko-Profil des DAX</i>
						auf Sicht von 
						<i>12&nbsp;Monaten</i>
						?
					`,
				"en": `
						Do you agree or disagree with the following statements about the relationship between the 
						<i>developments of inflation </i>
						and the 
						<i>return-risk-profile of the DAX</i>
						over the 
						<i>next 12&nbsp;months</i>
						?
					`,
			}.Outline("3.")

			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 0
			gr.BottomVSpacers = 1
		}

		//
		// gr1b q4 radio rows
		{

			colLblsQ4 := []trl.S{}

			var columnTemplateLocal = []float32{
				4.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			inpNames := []string{
				"qs3_inf_narrative_a",
				"qs3_inf_narrative_b",
				"qs3_inf_narrative_c",
				"qs3_inf_narrative_d",
				"qs3_inf_narrative_e",
				"qs3_inf_narrative_f",
			}

			for i := 0; i < len(inpNames); i++ {
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					colLblsQ4,
					[]string{inpNames[i]},
					radioVals6,
					[]trl.S{rowLbls3and4[i]},
				)
				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 1
				gr.RandomizationGroup = 1
				gr.RandomizationSeed = 0
			}

		}

		//
		//
		// gr2a q4 - main label
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 1
			gr.Cols = 1
			{
				//
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": `<br>`,
						"en": `<br>`,
					}
					inp.ColSpan = gr.Cols
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
						auf Sicht von 
						<i>12&nbsp;Monaten</i>
						am besten widerspiegelt:
					`,
						"en": `
						From the following statements, please select the one that, in your opinion, best reflects the relationship between the 
						<i>development of inflation </i>
						and the 
						<i>risk-return profile of the DAX</i>
						over the next  
						<i>12&nbsp;months</i>
						 :

					`,
					}.Outline("4.")
					inp.ColSpan = gr.Cols
				}

			}
		}

		//
		// q4b - six shuffled groups
		{

			radioVals := []string{
				"a",
				"b",
				"c",
				"d",
				"e",
				"f",
				"no_answer",
			}

			last := trl.S{
				"de": `Keine Antwort`,
				"en": `No answer`,
			}

			lbls := rowLbls3and4
			lbls = append(lbls, last)

			for i := 0; i < len(radioVals); i++ {

				gr := page.AddGroup()
				gr.Cols = 1
				gr.RandomizationGroup = 2
				gr.RandomizationSeed = 1
				if i == len(radioVals)-1 {
					// no shuffling of no-answer
					gr.RandomizationGroup = 0
				}

				gr.BottomVSpacers = 1
				if i == len(radioVals)-1 {
					gr.BottomVSpacers = 3
				}

				{
					inp := gr.AddInput()
					inp.Type = "radio"
					inp.Name = "qs4_inf_narrative"
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

		//
		// gr3 q5
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
					"de": `Bitte beschreiben Sie kurz in ganzen Sätzen, 
						wieso Sie sich für Ihre Auswahl in der vorangegangen Frage (Frage 4) 
						entschieden haben`,
					"en": `Please describe briefly in complete sentences 
						why you chose your selection in the previous question (question 4) `,
				}.Outline("5.")
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qs5_motivate_q4"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

	}

	return nil
}
