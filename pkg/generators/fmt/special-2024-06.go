package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func twoNumbersNoAnswer(
	page *qst.WrappedPageT,
	main trl.S,
	inps []string, // input base for each row
	colSuffixes []string, // input suffixes
	colHeaders []trl.S,
	rowLabels []trl.S,
) {

	gr := page.AddGroup()
	gr.BottomVSpacers = 4
	gr.Cols = 4
	gr.ColWidths("1.9fr  1.7fr  1.7fr 1.4fr")

	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleGridContainer.GapColumn = "1.2rem"
	// gr.Style.Mobile.StyleGridContainer.GapRow = "0.2rem"

	// main label
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = main
		inp.Style = css.NewStylesResponsive(inp.Style)
		inp.Style.Desktop.StyleBox.Position = "relative"
		inp.Style.Desktop.StyleBox.Top = "0.55rem"
	}

	// first row - empty first cell
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = 1
		inp.Label = trl.S{
			"de": "&nbsp;",
			"en": "&nbsp;",
		}
	}

	// first row - cols 2-4
	for _, lbl := range colHeaders {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = 1
		inp.Label = lbl
		inp.Style = css.ItemCenteredMCA(inp.Style)
		inp.Style = css.ItemEndCA(inp.Style)
		inp.Style.Desktop.StyleBox.Position = "relative"
		inp.Style.Desktop.StyleBox.Top = "0.85rem"

		inp.Style.Mobile.StyleText.FontSize = 90
		inp.Style.Mobile.StyleText.LineHeight = 100

	}

	//
	//
	//
	// second to fourth row: inputs
	for i, row := range rowLabels {

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = row
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v", inps[i], colSuffixes[0])
			inp.Suffix = trl.S{"de": "%", "en": "%"}
			inp.ColSpan = 1
			inp.ColSpanControl = 3
			inp.Min = -100
			inp.Max = 1000
			inp.Step = 0.01
			inp.MaxChars = 5
			inp.Style = css.ItemCenteredMCA(inp.Style)
			inp.Style = css.ItemEndCA(inp.Style)
		}

		// different suffix
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v", inps[i], colSuffixes[1])
			inp.Suffix = trl.S{"de": "%", "en": "%"}
			inp.ColSpan = 1
			inp.ColSpanControl = 3
			inp.Min = -100
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

func labelNumberX2(
	page *qst.WrappedPageT,
	main trl.S,
	inps []string, // input base for each row
	colHeaders []trl.S,
	rowLabels1 []trl.S,
	rowLabels2 []trl.S,
) {

	gr := page.AddGroup()
	gr.BottomVSpacers = 3
	gr.Cols = 4
	gr.ColWidths("2.1fr  1.9fr  2.1fr   1.9fr")

	// main label
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = main
	}

	// first row
	for _, lbl := range colHeaders {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = 2
		inp.Label = lbl.Bold()
		// inp.Style = css.ItemCenteredMCA(inp.Style)
		inp.Style = css.ItemEndCA(inp.Style)
		inp.Style.Desktop.StyleBox.Position = "relative"
		inp.Style.Desktop.StyleBox.Top = "0.15rem"
		inp.Style.Desktop.StyleBox.Padding = "0 1.5rem 0 4.5rem"
		inp.Style.Mobile.StyleBox.Padding = "0  1.5rem 0 0"

		inp.Style.Mobile.StyleText.FontSize = 90
		inp.Style.Mobile.StyleText.LineHeight = 100
	}

	//
	//
	// second to fourth row: inputs
	for i, row := range rowLabels1 {

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = row
			inp.Style = css.ItemEndMA(inp.Style)
		}

		if i != 0 {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("rexp_ecb_%v", inps[i])
			inp.Suffix = trl.S{"de": "Basispunkte", "en": "basis points"}
			inp.ColSpan = 1
			inp.ColSpanControl = 3
			inp.Min = -200
			inp.Max = 200
			inp.Step = 1
			inp.MaxChars = 5
			inp.Style = css.ItemCenteredMCA(inp.Style)
			inp.Style = css.ItemEndCA(inp.Style)
		} else {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{"de": "&nbsp;", "en": "&nbsp;"}
			}

		}

		// different suffix
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = rowLabels2[i]
			inp.Style = css.ItemEndMA(inp.Style)
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("rexp_fed_%v", inps[i])
			inp.Suffix = trl.S{"de": "Basispunkte", "en": "basis points"}
			inp.ColSpan = 1
			inp.ColSpanControl = 3
			inp.Min = -200
			inp.Max = 200
			inp.Step = 1
			inp.MaxChars = 5
			inp.Style = css.ItemCenteredMCA(inp.Style)
			inp.Style = css.ItemEndCA(inp.Style)
		}

	}

}

func special202406b(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 6
	if !cond {
		return nil
	}

	page := q.AddPage()
	page.Label = trl.S{
		"de": "Sonderfragen zu Zinserwartungen",
		"en": "Special questions: Interest rate expectations",
	}
	page.Short = trl.S{
		"de": "Sonderfragen:<br>Zinserwartungen",
		"en": "Special questions:<br>Interest rates",
	}

	page.WidthMax("42rem")

	//
	//
	//
	{

		inps := []string{
			"inf_ger",
			"inf_us",
		}
		suffixes := []string{
			// "2024",
			"0yr",
			"3yrs",
		}

		lblMain := trl.S{
			"de": `Die 
						<i>jährliche Inflationsrate im Euroraum</i> 
					(durchschnittliche jährliche Veränderung des HVPI in Prozent) 
					sowie den 
						<i>USA</i>
					(durchschnittliche jährliche Veränderung des PCE in Prozent) 
						<i>für 2024 bzw. im Zeitraum 2024-2026 </i>
					erwarte ich bei den folgenden Werten:
				`,
			"en": `
					I expect the following values 
						<i>for the annual inflation rate in the euro area</i> 
					(annual average change of the HICP, in percent) 
					and the 
						<i>USA</i>
					(annual average change of the PCE, in percent) 
						<i>for 2024 respectively the period 2024-2026</i>:
				`,
		}.Outline("3.")

		headers := []trl.S{
			{
				"de": `Punktprognose  in Prozent für das Jahr 2024`,
				"en": `Point forecast in percent for the year 2024`,
			},
			{
				"de": `Punktprognose  in Prozent für den Zeitraum 2024-2026`,
				"en": `Point forecast in percent for the period   2024-2026`,
			},
			{
				"de": `keine<br>Angabe`,
				"en": `no estimate`,
			},
		}

		rows := []trl.S{
			{
				"de": `Inflation,<br> Eurozone`,
				"en": `Inflation rate,<br> euro area`,
			}, {
				"de": `Inflation,<br> USA`,
				"en": `Inflation rate,<br> USA `,
			},
		}

		twoNumbersNoAnswer(
			qst.WrapPageT(page),
			lblMain,
			inps,
			suffixes,
			headers,
			rows,
		)
	}

	//
	//
	//
	{

		inps := []string{
			"gdp_growth_ger",
			"gdp_growth_us",
		}
		suffixes := []string{
			"0yr",
			"3yrs",
		}
		lblMain := trl.S{
			"de": `
				Die 
					<i>jährliche Wachstumsrate des realen Bruttoinlandprodukts im Euroraum</i> 
				sowie den 
					<i>USA für 2024 bzw. im Zeitraum 2024-2026</i>
				erwarte ich bei den folgenden Werten:			
			`,
			"en": `
				I expect the following values for the 
					<i>annual real GDP growth rate in the euro area</i>
				 and the 
				 	<i>USA</i> 
					<i>for 2024 respectively the period 2024-2026</i>:			
			
			`,
		}.Outline("4.")

		headers := []trl.S{
			{
				"de": `Punktprognose  in Prozent für das Jahr 2024`,
				"en": `Point forecast in percent for the year 2024`,
			},
			{
				"de": `Punktprognose  in Prozent für den Zeitraum 2024-2026`,
				"en": `Point forecast in percent for the period   2024-2026`,
			},
			{
				"de": `keine<br>Angabe`,
				"en": `no estimate`,
			},
		}

		rows := []trl.S{
			{
				"de": `BIP-Wachstumsrate,<br> Eurozone`,
				"en": `Real GDP growth rate, <br>euro area`,
			},
			{
				"de": `BIP-Wachstumsrate,<br> USA `,
				"en": `Real GDP growth rate,<br> USA  `,
			},
		}

		twoNumbersNoAnswer(
			qst.WrapPageT(page),
			lblMain,
			inps,
			suffixes,
			headers,
			rows,
		)
	}

	{

		inps := []string{
			"2024_6",
			"2024_7",
			"2024_9",
			"2024_10",
			"2024_12",
			"2025_1",
			"2025_3",
			"2025_4",
			"2025_6",
			"2025_7",
		}

		lblMain := trl.S{
			"de": `
				Wir möchten Sie zu Ihren Erwartungen über zukünftige Zinsentscheidungen der Europäischen Zentralbank (EZB) und des Federal Open Market Commitee (FOMC) des Federal Reserve System befragen. Geben Sie hierzu Ihre Erwartungen bezüglich der 
					<i>Zinsschritte (in Basispunkten, negative Werte für Zinssenkungen)</i> 
				bei den nachfolgenden Treffen der Komitees an: 

				<br>
				<small>Hinweis: Derzeit liegt der Leitzins der EZB bei 4,25% und die Federal Funds Rate in den USA bei 5,25-5,50%. Für die FOMC-Treffen im Jahr 2025 sind noch keine genauen Termine bekannt.</small>
			`,
			"en": `
				We would now like to ask you about your expectations on future interest rate decisions by the European Central Bank (ECB) and the Federal Open Market Commitee (FOMC) of the Federal Reserve System. Please state your expectations on 
					<i>interest rate movements (in basis points, negative values for interest rate cuts)</i> 
				after the following meetings of the commitees: 

				<br>
				<small>Hint: The Main Refinancing Operations Rate of the ECB currently stands at 4,25% and the Federal Funds Rate in the USA stands at 5,25-5,50%. The precise dates for the FOMC meetings in 2025 are unknown at this point.</small>
			
			`,
		}.Outline("5.")

		headers := []trl.S{
			{
				"de": `Änderung des Leitzins (EZB)`,
				"en": `Changes of <br>the Main Refinancing Operations Rate (ECB)`,
			},
			{
				"de": `Änderung der Federal Funds Rate (FOMC)`,
				"en": `Changes of <br>the Federal Funds Rate (FOMC)`,
			},
		}

		rowsLeft := []trl.S{
			{
				// "de": "6. Juni 2024",
				// "en": "June 6, 2024",
				"de": "&nbsp;",
				"en": "&nbsp;",
			},
			{
				"de": "18. Juli 2024",
				"en": "July 18, 2024",
			},
			{
				"de": "12. September 2024",
				"en": "September 12, 2024",
			},
			{
				"de": "17. Oktober 2024",
				"en": "October 17, 2024",
			},
			{
				"de": "12. Dezember 2024",
				"en": "December 12, 2024",
			},
			{
				"de": "1. Januar 2025",
				"en": "January 1, 2025",
			},
			{
				"de": "6. März 2025",
				"en": "March 6, 2025",
			},
			{
				"de": "17. April 2025",
				"en": "April 17, 2025",
			},
			{
				"de": "5. Juni 2025",
				"en": "June 5, 2025",
			},
			{
				"de": "24. Juli 2025",
				"en": "Juli 24, 2025",
			},
		}
		rowsRight := []trl.S{
			{
				"de": "12. Juni 2024",
				"en": "June 12, 2024",
			},
			{
				"de": "31. Juli 2024",
				"en": "July 31, 2024",
			},
			{
				"de": "18. September 2024",
				"en": "September 18, 2024",
			},
			{
				"de": "7. November 2024",
				"en": "November 7, 2024",
			},
			{
				"de": "18. Dezember 2024",
				"en": "December 18, 2024",
			},
			{
				"de": "Januar 2025",
				"en": "January 2025",
			},
			{
				"de": "März 2025",
				"en": "March 2025",
			},
			{
				"de": "April/Mai 2025",
				"en": "April/May 2025",
			},
			{
				"de": "Juni 2025",
				"en": "June 2025",
			},
			{
				"de": "Juli 2025",
				"en": "July 2025",
			},
		}

		labelNumberX2(
			qst.WrapPageT(page),
			lblMain,
			inps,
			headers,
			rowsLeft,
			rowsRight,
		)
	}

	//
	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapColumn = "0rem"
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.4rem"
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "dow_exp_6m"

			inp.Label = trl.S{
				"de": `Den <i>Dow Jones (USA)</i> erwarte ich in 6 Monaten bei `,
				"en": `Six months ahead, I expect the <i>Dow Jones (USA)</i> to stand at`,
			}.Outline("6.")
			inp.Suffix = trl.S{"de": "Punkten.", "en": "points."}
			inp.Placeholder = trl.S{
				"de": "00000",
				"en": "00000",
			}
			inp.Min = 1000
			inp.Max = 90000
			inp.Step = 0
			inp.MaxChars = 6

			inp.ColSpan = 6
			inp.ColSpanLabel = 4
			inp.ColSpanControl = 2
		}
	}

	return nil
}
