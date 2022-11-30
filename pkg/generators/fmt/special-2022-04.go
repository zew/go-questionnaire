package fmt

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202204(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 4
	if !cond {
		return nil
	}

	//
	//
	//
	{
		page := q.AddPage()
		// page.NavigationCondition = "GermanOnly"
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage zu den Auswirkungen der neuen geopolitischen Situation, die durch Russlands Krieg gegen die Ukraine herbeigeführt wurde",
			"en": "Special question on the new geopolitical situation resulting from Russia’s war against the Ukraine:",
		}
		page.Short = trl.S{
			"de": "Krieg gegen<br>Ukraine",
			"en": "War against<br>Ukraine",
		}
		page.WidthMax("48rem")

		rowLabelsTimeHorizon1 := []trl.S{
			{
				"de": "Für 2022",
				"en": "For 2022",
			},
			{
				"de": "Für 2023",
				"en": "For 2023",
			},
			{
				"de": "Für 2024",
				"en": "For 2024",
			},
			{
				"de": "Auf Sicht von 5&nbsp;Jahren",
				"en": "In the next 5 years",
			},
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				raiseDecrease6b(),
				[]string{
					"war_on_gdp_2022",
					"war_on_gdp_2023",
					"war_on_gdp_2024",
					"war_on_gdp_5yrs",
				},
				radioVals6,
				rowLabelsTimeHorizon1,
			)

			gb.MainLabel = trl.S{
				"de": `
				<b>1.</b>
				Die neue geopolitische Situation hat sich folgendermaßen auf meine Prognosen des realen Bruttoinlandsproduktes (BIP) von Deutschland ausgewirkt:
			`,
				"en": `
				<b>1.</b>
				The new geopolitical situation had the following effects on my forecasts of real gross domestic product (GDP) in Germany:
			`}

			gr := page.AddGrid(gb)
			_ = gr
		}

		//
		//
		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				improvedDeteriorated6b(),
				[]string{
					"war_sec_banks",
					"war_sec_insur",
					"war_sec_cars",
					"war_sec_chemi",
					"war_sec_steel",
					"war_sec_elect",
					"war_sec_mecha",
					"war_sec_consu",
					"war_sec_const",
					"war_sec_utili",
					"war_sec_servi",
					"war_sec_telec",
					"war_sec_infor",
				},
				radioVals6,
				rowLabelsSectors,
			)
			gb.MainLabel = trl.S{
				"de": `<b>2.</b> 
				Auf Sicht von 5 Jahren hat sich die neue geopolitische Situation folgendermaßen auf meine Prognosen der Ertragslage deutscher Unternehmen in den folgenden Branchen ausgewirkt:`,
				"en": `<b>2.</b> 
				The new geopolitical situation had the following effects on my forecasts of the profit situation of German companies in the following industries over the next 5 years.
				`,
			}
			gr := page.AddGrid(gb)
			_ = gr
		}

		//
		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				raiseDecrease6b(),
				[]string{
					"war_govt_bonds_2022",
					"war_govt_bonds_2023",
					"war_govt_bonds_2024",
					"war_govt_bonds_5yrs",
				},
				radioVals6,
				rowLabelsTimeHorizon1,
			)

			gb.MainLabel = trl.S{
				"de": `
				<b>3.</b>
				Die neue geopolitische Situation hat sich folgendermaßen auf meine Prognosen der Zinsen von 10-jährigen Bundesanleihen ausgewirkt:

				`,
				"en": `
				<b>3.</b>
				The new geopolitical situation had the following effects on my forecasts of interest rates on the 10-year German Bund:
				`,
			}

			gr := page.AddGrid(gb)
			_ = gr
		}

	}

	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfrage: Klimawandel",
			"en": "Special: Climate change",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Klimawandel",
			"en": "Special:<br>Climate change",
		}
		page.WidthMax("36rem")
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				// labelsImproveDeteriorateSectoral(),
				labelsStronglyPositiveStronglyNegativeInfluence(),
				[]string{
					"climatech_sec_banks",
					"climatech_sec_insur",
					"climatech_sec_cars",
					"climatech_sec_chemi",
					"climatech_sec_steel",
					"climatech_sec_elect",
					"climatech_sec_mecha",
					"climatech_sec_consu",
					"climatech_sec_const",
					"climatech_sec_utili",
					"climatech_sec_servi",
					"climatech_sec_telec",
					"climatech_sec_infor",
				},
				radioVals6,
				rowLabelsSectors,
			)
			gb.MainLabel = trl.S{
				"de": `<b>3.</b> 
					<u>Die folgende Sonderfrage dient als Teil einer weiterführenden wissenschaftlichen Arbeit.</u>
					<br>
					Ein möglicher negativer Nebeneffekt der neuen geopolitischen Lage, 
					die von Russlands Krieg gegen die Ukraine herbeigeführt wurde, 
					ist, dass die Klimapolitik an Priorität verlieren könnte. 
					Was glauben Sie, wie wirkt sich der Klimawandel 
					im Allgemeinen auf die langfristigen Ertragschancen 
					von Unternehmen aus den folgenden Sektoren aus?
				`,
				"en": `<b>3.</b> 
					<u>The following special question serves as part of a continuing scientific work.</u>
					<br>

					 A possible negative side effect of the new geopolitical situation brought about by Russia's war against Ukraine is that climate policy may receive a lower priority. 
					 How will climate change in general affect the long-term earnings prospects of firms belonging to the following sectors?

				`,
			}
			gr := page.AddGrid(gb)
			_ = gr
		}

	}
	return nil

}
