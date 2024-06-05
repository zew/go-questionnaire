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
	// gr.ColWidths("2.6fr  1.7fr  1.7fr 1.4fr")
	gr.ColWidths("2.1fr  1.7fr  1.7fr 1.4fr")

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
			inp.Min = -40
			inp.Min = -100
			inp.Max = 50
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
		"de": "Sonderfragen<br>Zinserwartungen",
		"en": "Special questions<br>Interest rates",
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
					<i>jährliche Inflationsrate in Deutschland</i> 
						sowie den 
					<i>USA</i>
						(durchschnittliche jährliche Veränderung des VPI in Prozent) 
					<i>für 2024 bzw. im Zeitraum 2024-2027 </i>
					erwarte ich bei den folgenden Werten:
				`,
			"en": `
				I expect the following values 
					<i>for the annual inflation rate in Germany</i> 
				and the 
					<i>USA</i>
				(annual average change of the CPI, in percent) 
					<i>for 2024 respectively the period 2024-2027</i>:
				`,
		}.Outline("4.")

		headers := []trl.S{
			{
				"de": `Punktprognose  in Prozent für das Jahr 2024`,
				"en": `Point forecast in percent for the year 2024`,
			},
			{
				"de": `Punktprognose  in Prozent für den Zeitraum 2024-2027`,
				"en": `Point forecast in percent for the period   2024-2027`,
			},
			{
				"de": `keine<br>Angabe`,
				"en": `no estimate`,
			},
		}

		rows := []trl.S{
			{
				"de": `Inflation,<br> Deutschland`,
				"en": `Inflation rate,<br> Germany`,
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
					<i>jährliche Wachstumsrate des realen Bruttoinlandprodukts in Deutschland</i> 
				sowie den 
					<i>USA für 2024 bzw. im Zeitraum 2024-2027</i>
				erwarte ich bei den folgenden Werten:			
			`,
			"en": `
				I expect the following values for the 
					<i>annual real GDP growth rate in Germany</i>
				 and the 
				 	<i>USA</i> 
					<i>for 2024 respectively the period 2024-2027</i>:			
			
			`,
		}.Outline("5.")

		headers := []trl.S{
			{
				"de": `Punktprognose  in Prozent für das Jahr 2024`,
				"en": `Point forecast in percent for the year 2024`,
			},
			{
				"de": `Punktprognose  in Prozent für den Zeitraum 2024-2027`,
				"en": `Point forecast in percent for the period   2024-2027`,
			},
			{
				"de": `keine<br>Angabe`,
				"en": `no estimate`,
			},
		}

		rows := []trl.S{
			{
				"de": `BIP-Wachstumsrate,<br> Deutschland`,
				"en": `Real GDP growth rate, G<br>ermany`,
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

	return nil
}
