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
					<b>jährliche Inflationsrate in Deutschland</b> (durchschnittliche jährliche Veränderung des HICP in Prozent) 
					bzw. die durchschnittliche jährliche 
					<b>Wachstumsrate des realen Bruttoinlandprodukts</b> in Deutschland  

					Sicht von <b>zwölf Monaten</b> bzw. <b>drei Jahren</b>   

					zwischen den folgenden Werten liegen:
				`,
				"en": `
					todo
				`,
			}.Outline("1.")

			inpNames := []string{
				"inflation_12m",
				"inflation_36m",
				"growth_12m",
				"growth_36m",
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

			checkBoxCascade(
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
					Mit einer Wahrscheinlichkeit von 90 Prozent wird die durchschnittliche jährliche <b>Rendite des DAX</b> auf Sicht von 
					<b>zwölf Monaten</b> bzw. 
					<b>drei Jahren</b> zwischen den folgenden Werten liegen:
					<br>
					Hinweis: Im Zeitraum 2011-2021 betrug die jährliche DAX-Rendite im Durchschnitt 8,9 Prozent mit einer Standardabweichung von 14,7 Prozent.
				`,
				"en": `
					todo
				`,
			}.Outline("2.")

			inpNames := []string{
				"dax_12m",
				"dax_36m",
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

			checkBoxCascade(
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
					"inf_narrative_a",
					"inf_narrative_b",
					"inf_narrative_c",
					"inf_narrative_d",
					"inf_narrative_e",
					"inf_narrative_f",
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
						Wie beurteilen Sie die folgenden Aussagen zum Zusammenhang zwischen der <b>Inflationsentwicklung</b>
						 und dem 
						<b>Rendite-Risiko-Profil</b> des DAX 
						 in 
						<b>2023</b>?
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
						<b>Inflationsentwicklung</b> und dem 
						<b>Rendite-Risiko-Profil des DAX</b> in 
						<b>2023</b> am besten widerspiegelt: 
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
				inp.Name = "inf_to_dax"
				inp.Label = lbls[i]
				inp.ValueRadio = radioVals[i]
				inp.ColSpan = 1

				inp.ColSpanLabel = 1
				inp.ColSpanControl = 9

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "justify"

				inp.ControlFirst()
			}

		}

	}

	return nil
}
