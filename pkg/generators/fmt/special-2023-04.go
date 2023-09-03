package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func numberInputWithNoAnswer(
	q *qst.QuestionnaireT,
	page *qst.WrappedPageT,
	inps []string,
	main trl.S,
	headers []trl.S,
	rows []trl.S,
) {

	gr := page.AddGroup()
	gr.BottomVSpacers = 4
	gr.Cols = 3
	// gr.Style = css.NewStylesResponsive(gr.Style)
	// gr.Style.Mobile.StyleGridContainer.GapRow = "0.2rem"
	gr.ColWidths("4.6fr    1.7fr 1.7fr")

	// main label
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = main
		inp.Style = css.NewStylesResponsive(inp.Style)
		inp.Style.Desktop.StyleBox.Position = "relative"
		inp.Style.Desktop.StyleBox.Top = "0.5rem"
	}

	// first row: labels
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = 1
		inp.Label = trl.S{
			"de": "&nbsp;",
			"en": "&nbsp;",
		}
	}
	// first row - cols 2-3
	for _, lbl := range headers {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = 1
		inp.Label = lbl
		inp.Style = css.ItemCenteredMCA(inp.Style)
		inp.Style = css.ItemEndCA(inp.Style)
		inp.Style.Desktop.StyleBox.Position = "relative"
		inp.Style.Desktop.StyleBox.Top = "0.3rem"
	}

	//

	//
	//
	// second to fourth row: inputs
	for i, row := range rows {

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = row
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_pf", inps[i])
			inp.Suffix = trl.S{"de": "%", "en": "%"}
			inp.ColSpan = 1
			inp.ColSpanControl = 3
			inp.Min = -40
			inp.Min = -100
			inp.Max = 50
			inp.Max = 1000
			inp.Step = 0.01
			inp.MaxChars = 5
			inp.Style = css.ItemCenteredMCA(inp.Style)
			inp.Style = css.ItemEndCA(inp.Style)
		}

		{
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.ColSpan = 1
			inp.Name = fmt.Sprintf("%v__noanswer", inps[i])
			inp.ColSpanControl = 1
			inp.Style = css.ItemEndCA(inp.Style)
			// inp.Style = css.ItemStartMA(inp.Style)
		}

	}

}

func special202304(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2023 && q.Survey.Month == 4
	if !cond {
		return nil
	}

	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Spezialfragen: Inflation, BIP-Wachstum und DAX-Renditen",
			"en": "Special questions: Inflation, GDP growth and DAX returns",
		}
		page.Short = trl.S{
			"de": "Inflation, BIP <br>und DAX-Renditen",
			"en": "Inflation, GDP <br>and DAX returns",
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
				"de": `Auf Sicht von 12 bzw. 36&nbsp;Monaten, was sind Ihre Prognosen für die jährliche Inflationsrate in Deutschland?`,
				"en": `Looking ahead 
							<i>12</i> 
								and 
							<i>36</i>&nbsp;months, what are your forecasts for the 
							<i>annual inflation rate</i>
								for 
							<i>Germany</i>?`,
			}.Outline("3.")

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
							(HVPI April 2024 vs. HVPI April 2023)
							`,
					"en": `Inflation rate in <i>12&nbsp;months</i> 
							(HCPI April 2024 vs. HCPI April 2023)
							`,
				}, {
					"de": `Durchschnittliche jährliche Inflationsrate in Deutschland auf
							Sicht von <i>36&nbsp;Monaten</i>
							(HVPI April 2026 vs. HVPI April 2023)
							`,
					"en": `<i><u>Average</u></i> inflation rate over the next <i>36&nbsp;months</i> 
							(HCPI April 2026 vs. HCPI April 2023)
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
				"de": `Auf Sicht von 12&nbsp;Monaten, was sind Ihre Prognosen für die Wachstumsrate 
						des realen Bruttoinlandprodukts in Deutschland bzw. die Rendite des DAX?`,
				"en": `Looking ahead
							<i>12</i>&nbsp;months, what are your forecasts for the  
							<i>annual growth rate of German real GDP</i>
								and the 
							<i>return</i>
								of the
							<i>DAX</i>?`,
			}.Outline("4.")

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
							(BIP Q1 2024 vs. BIP Q1 2023)
							`,
					"en": `German real GDP growth in 12&nbsp;months 
							(GDP Q1 2024 vs. GDP Q1 2023)`,
				}, {
					"de": `DAX-Rendite über die nächsten 12&nbsp;Monate
							(DAX April 2024 vs. April 2023)		
							`,
					"en": `Return of the DAX over the next 12&nbsp;months 
							(DAX April 2024 vs. April 2023)`,
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

	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Spezialfragen: Spannungen im Bankensystem und deren wirtschaftliche Folgen",
			"en": "Special question: tensions in the banking system and their economic implications",
		}
		page.Short = trl.S{
			"de": "Spannungen im<br>Bankensystem",
			"en": "Tensions in the <br>banking system",
		}
		page.WidthMax("42rem")

		colTemplate, _, _ := colTemplateWithFreeRow()

		{
			gb := qst.NewGridBuilderRadios(
				colTemplate,
				improvedDeterioratedPlusMinus6(),
				[]string{
					"bt6m_german_gdp",
					"bt6m_german_inf",
					"bt6m_euro_gdp",
					"bt6m_euro_inf",
					"bt6m_ecb_rate",

					"bt6m_us_gdp",
					"bt6m_us_inf",
					"bt6m_fed_rate",
				},
				radioVals6,
				[]trl.S{
					{
						"de": `BIP Deutschland`,
						"en": `German GDP`,
					},
					{
						"de": `Inflation Deutschland`,
						"en": `German Inflation`,
					},
					{
						"de": `BIP Euroraum`,
						"en": `Euro area GDP`,
					},
					{
						"de": `Inflation Euroraum`,
						"en": `Euro area Inflation`,
					},
					{
						"de": `Leitzins der EZB`,
						"en": `ECB's main refinancing rate`,
					},

					{
						"de": `BIP USA`,
						"en": `US GDP`,
					},
					{
						"de": `Inflation USA`,
						"en": `US inflation`,
					},
					{
						"de": `Leitzins der FED`,
						"en": `Fed funds rate`,
					},
				},
			)
			gb.MainLabel = trl.S{
				"de": `Mit Blick auf die nächsten <i>6 Monate</i>, was glauben Sie, wie sich die jüngsten Spannungen im Bankensystem 
						auf die folgenden wirtschaftlichen Variablen auswirken werden? (++ = starker Anstieg, + = leichter Anstieg, 0 = keine Auswirkungen, - = leichter Rückgang, -- = starker Rückgang)
	
			`,
				"en": `Looking ahead <i>6 months</i>, what do you think will be the overall impact of the recent tensions in the banking sector on the following economic variables?`,
			}.Outline("5.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 4
		}

		{
			gb := qst.NewGridBuilderRadios(
				colTemplate,
				improvedDeterioratedPlusMinus6(),
				[]string{
					"bt24m_german_gdp",
					"bt24m_german_inf",
					"bt24m_euro_gdp",
					"bt24m_euro_inf",
					"bt24m_ecb_rate",

					"bt24m_us_gdp",
					"bt24m_us_inf",
					"bt24m_fed_rate",
				},
				radioVals6,
				[]trl.S{
					{
						"de": `BIP Deutschland`,
						"en": `German GDP`,
					},
					{
						"de": `Inflation Deutschland`,
						"en": `German Inflation`,
					},
					{
						"de": `BIP Euroraum`,
						"en": `Euro area GDP`,
					},
					{
						"de": `Inflation Euroraum`,
						"en": `Euro area Inflation`,
					},
					{
						"de": `Leitzins der EZB`,
						"en": `ECB's main refinancing rate`,
					},

					{
						"de": `BIP USA`,
						"en": `US GDP`,
					},
					{
						"de": `Inflation USA`,
						"en": `US inflation`,
					},
					{
						"de": `Leitzins der FED`,
						"en": `Fed funds rate`,
					},
				},
			)
			gb.MainLabel = trl.S{
				"de": `Mit Blick auf die nächsten <i>2 Jahre</i>, was glauben Sie, wie sich die jüngsten Spannungen im Bankensystem 
						auf die folgenden wirtschaftlichen Variablen auswirken werden? (++ = starker Anstieg, + = leichter Anstieg, 0 = keine Auswirkungen, - = leichter Rückgang, -- = starker Rückgang)`,
				"en": `Looking ahead <i>2 years</i>, what do you think will be the overall impact of the recent tensions in the banking sector on the following economic variables?`,
			}.Outline("6.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 3
		}
	}

	return nil
}
