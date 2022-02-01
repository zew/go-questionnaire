package fmt

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202202b(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2022 || q.Survey.Month != 2 {
		return nil
	}

	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage: Inflation, Prognosetreiber und Geldpolitik 3",
			"en": "Special: Inflation, forecast drivers and monetary policy 3",
		}
		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik 3",
			"en": "Inflation,<br>Mon. Policy 3",
		}
		// page.WidthMax("42rem")
		page.WidthMax("42rem")

		//
		//
		//
		// gr3 ... gr10
		var columnTemplateLocal = []float32{
			0.2, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				labelsStrongIncreaseStrongDecrease2(),
				[]string{"infl_euro_area_2025_2030"},
				radioVals6,
				[]trl.S{
					{
						"de": " &nbsp; ",
						"en": " &nbsp; ",
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
					<b>4.</b>  &nbsp;
						Gegenüber dem Zeitraum 2022-2024 wird die durchschnittliche Inflationsrate im Euroraum im 
						<b>Zeitraum 2025-2030</b>
						
					`,

				"en": `
					<b>4.</b>  &nbsp;
					
					`,
			}

			gr := page.AddGrid(gb)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.6rem"
			// gr.BottomVSpacers = 1
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				labelsStrongIncreaseStrongDecrease2(),
				[]string{"ezb_rate_2025_2030"},
				radioVals6,
				[]trl.S{
					{
						"de": " &nbsp; ",
						"en": " &nbsp; ",
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
					<b>5.</b>  &nbsp;
						Gegenüber dem Zeitraum 2022-2024 wird der durchschnittliche Hauptrefinanzierungssatz der EZB im 
						<b>Zeitraum 2025-2030</b>
						
					`,

				"en": `
					<b>5.</b>  &nbsp;
					
					`,
			}

			gr := page.AddGrid(gb)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.6rem"
			// gr.BottomVSpacers = 1
		}

	} // page

	//
	//
	return nil
}
