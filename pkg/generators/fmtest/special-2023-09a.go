package fmtest

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202309a(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2023 && q.Survey.Month == 9
	if !cond {
		return nil
	}

	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfragen zu Inflation, BIP-Wachstum und DAX-Renditen",
			"en": "Special questions: Inflation, GDP growth and DAX returns",
		}
		page.Short = trl.S{
			"de": "Sonderfragen:<br>Inflation, BIP & DAX-Renditen",
			"en": "Special questions:<br>Inflation, GDP & DAX returns",
		}

		page.WidthMax("42rem")

		{

			inps := []string{
				"inf_12m",
				"inf_36m",
			}

			_ = trl.S{
				"de": `xx`,
				"en": `xx`,
			}
			lblMain := trl.S{
				"de": `Auf Sicht von 
						<i>12 bzw. 36&nbsp;Monaten</i>, 
						was sind Ihre Prognosen für die 
						<i>jährliche Inflationsrate</i> 
						  in 
						<i>Deutschland</i>
					?`,
				"en": `Looking ahead 
						<i>12</i> 
							and 
						<i>36</i>&nbsp;months, what are your forecasts for the 
						<i>annual inflation rate</i>
							for 
						<i>Germany</i>
					?`,
			}.Outline("1.")

			headers := []trl.S{
				{
					"de": `Punktprognose in Prozent`,
					"en": `Point forecast in percent`,
				}, {
					"de": `keine<br>Angabe`,
					"en": `no estimate`,
				},
			}

			rows := []trl.S{
				{
					"de": `Inflationsrate auf Sicht von <i>12&nbsp;Monaten</i>
							(HVPI September 2024 vs. HVPI September 2023)
							`,
					"en": `Inflation rate in <i>12&nbsp;months</i> 
							(HICP September 2024 vs. HICP September 2023)
							`,
				}, {
					"de": `Durchschnittliche jährliche Inflationsrate in Deutschland auf
							Sicht von <i>36&nbsp;Monaten</i>
							(HVPI September 2026 vs. HVPI September 2023)
							`,
					"en": `<i><u>Average</u></i> inflation rate over the next <i>36&nbsp;months</i> 
							(HICP September 2026 vs. HICP September 2023)
							`,
				},
			}

			numberInputWithNoAnswer(
				q,
				qst.WrapPageT(page),
				inps,
				lblMain,
				headers,
				rows,
			)
		}

		{

			inps := []string{
				"gdp_growth_12m",
				"dax_return_12m",
			}
			lblMain := trl.S{
				"de": `Auf Sicht von 
							<i>12&nbsp;Monaten</i>, was sind Ihre Prognosen für die 
							<i>Wachstumsrate des realen Bruttoinlandprodukts in Deutschland</i> 
							   bzw. die 
							<i>Rendite des DAX</i>
						?`,
				"en": `Looking ahead
							<i>12</i>&nbsp;months, what are your forecasts for the  
							<i>annual growth rate of German real GDP</i>
								and the 
							<i>return</i>
								of the
							<i>DAX</i>
						?`,
			}.Outline("2.")

			headers := []trl.S{
				{
					"de": `Punktprognose in Prozent`,
					"en": `Point forecast in percent`,
				}, {
					"de": `keine<br>Angabe`,
					"en": `no estimate`,
				},
			}

			rows := []trl.S{
				{
					"de": `BIP-Wachstumsrate in Deutschland auf Sicht von 12&nbsp;Monaten
							(BIP Q3 2024 vs. BIP Q3 2023)
							`,
					"en": `German real GDP growth in 12&nbsp;months 
							(GDP Q3 2024 vs. GDP Q3 2023)`,
				}, {
					"de": `DAX-Rendite über die nächsten 12&nbsp;Monate
							(DAX September 2024 vs. DAX September 2023)
							`,
					"en": `Return of the DAX over the next 12&nbsp;months 
							(DAX September 2024 vs. DAX September 2023)`,
				},
			}

			numberInputWithNoAnswer(
				q,
				qst.WrapPageT(page),
				inps,
				lblMain,
				headers,
				rows,
			)
		}
	}

	return nil
}
