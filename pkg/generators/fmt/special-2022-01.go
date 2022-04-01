package fmt

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special_2022_01(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2022 || q.Survey.Month != 1 {
		return nil
	}

	{
		page := q.AddPage()
		// page.NavigationCondition = "GermanOnly"
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage zum Regierungsprogramm der neuen Bundesregierung",
			"en": "Special: Program of the new German federal government:",
		}
		page.Short = trl.S{
			"de": "Regierungs-<br>programm",
			"en": "New govt. program",
		}
		page.WidthMax("48rem")

		// gr1
		rowLabelsTimeHorizon := []trl.S{
			{
				"de": "Auf Sicht von 6&nbsp;Monaten",
				"en": "In the next 6 months",
			},
			{
				"de": "Für 2022",
				"en": "For 2022",
			},
			{
				"de": "Für 2023",
				"en": "For 2023",
			},
			{
				"de": "Auf Sicht von 5&nbsp;Jahren",
				"en": "In the next 5 years",
			},
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				raiseDecrease6b(),
				[]string{
					"new_govt_gdp_6m", "new_govt_gdp_2022", "new_govt_gdp_2023", "new_govt_gdp_5yrs",
				},
				radioVals6,
				rowLabelsTimeHorizon,
			)

			gb.MainLabel = trl.S{
				"de": `
				<b>1.</b>
				Das Regierungsprogramm der neuen Bundesregierung 
				hat sich folgendermaßen auf meine  
				Prognosen des realen Bruttoinlandsproduktes (BIP) ausgewirkt:
			`,
				"en": `
				<b>1.</b>
				The program of the new federal government had the following 
				effects on my forecasts of real gross domestic product (GDP):
			`}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		//
		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				improvedDeteriorated6b(),
				[]string{
					"govt_sec_banks",
					"govt_sec_insur",
					"govt_sec_cars",
					"govt_sec_chemi",
					"govt_sec_steel",
					"govt_sec_elect",
					"govt_sec_mecha",
					"govt_sec_consu",
					"govt_sec_const",
					"govt_sec_utili",
					"govt_sec_servi",
					"govt_sec_telec",
					"govt_sec_infor",
				},
				radioVals6,
				rowLabelsSectors,
			)
			gb.MainLabel = trl.S{
				"de": "<b>2.</b> Auf Sicht von 5 Jahren hat sich das Regierungsprogramm der neuen Bundesregierung folgendermaßen auf meine Prognosen der Ertragslage deutscher Unternehmen in den folgenden Branchen ausgewirkt:",
				"en": "<b>2.</b> Over the next 5 years, the program of the new federal government had the following effects on my forecasts of the profit situation of German companies in the following industries:",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				raiseDecrease6b(),
				[]string{
					"new_govt_bond_6m", "new_govt_bond_2022", "new_govt_bond_2023", "new_govt_bond_5yrs",
				},
				radioVals6,
				rowLabelsTimeHorizon,
			)

			gb.MainLabel = trl.S{
				"de": `
				<b>3.</b>
				Das Regierungsprogramm der neuen Bundesregierung 
				hat sich folgendermaßen auf meine Prognosen
				 der Zinsen auf 10-jährigen Bundesanleihen ausgewirkt:

			`,
				"en": `
				<b>3.</b>
				The program of the new federal government had the following effects on my forecasts of interest rates on the 10-year German Bund:
			`}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	} // special page 4

	return nil
}
