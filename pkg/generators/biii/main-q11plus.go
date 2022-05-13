package biii

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func page4Quest11(q *qst.QuestionnaireT) {

	// page 4
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p4"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("42rem")

		//
		//
		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 7

			labels := []trl.S{
				{"de": "Privates (nicht börsengehandeltes) Beteiligungskapital (Private equity)"},
				{"de": "Börsengehandeltes Beteiligungskapital / Aktien (Public equity)"},
				{"de": "Privates (nicht börsengehandeltes) Fremdkapital/ Anleihen (Private debt)  "},
				{"de": "Börsengehandeltes Fremdkapital / Anleihen (Public debt)  "},
				{"de": "Immobilien (Real estate)  "},
				{"de": "Einlagen oder Zahlungsmitteläquivalente / Geldwerte (Deposits or cash equivalents / monetary assets)  "},
				{"de": "Soziale Infrastruktur Finanzierung (SOC (z.B. SIB / DIB) )  "},
				{"de": "Grüne Anleihen (Green Bonds)  "},
				{"de": "Schwellenländer(markt) (Emerging markets)  "},
				{"de": "Mikrofinanzierung (Microfinance)  "},
				{"de": "Rohstoffe (Commodities)  "},
				{"de": "Sustainability-Linked Bonds (SLBs)  "},
			}
			subName := []string{
				"private_equity",
				"public_equity",
				"private_external",
				"public_external",
				"real_estate",
				"money_deposits",
				"soc",
				"green_bonds",
				"emerging_markets",
				"microfinance",
				"commodities",
				"slb",
				// "hybrid",
			}
			//
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<br>
					<b>11.</b> &nbsp;	
					
					Bitte tragen Sie soweit möglich, ungefähre anteilige Impact Investitionsvolumina nach Anlageklassen/Instrumenten (in Prozent) ein. 					
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q11_%v", subName[idx])
				inp.Label = label

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 6
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				// inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Suffix = trl.S{"de": "%"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.LabelPadRight()

			}

			// q11 - row 10
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q11_hybrid_explain"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": "Hybride Finanzinstrumente (Hybrid financial instruments), bitte nennen"}
				inp.ColSpan = 5
				inp.ColSpanLabel = 1.9
				inp.ColSpanControl = 2.1
				inp.ControlTop()
				inp.LabelPadRight()
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q11_hybrid")
				inp.ColSpan = 2
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "%"}
				inp.Placeholder = trl.S{"de": "00"}
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleBox.Padding = "0 0 0 0.3rem"
			}

			// q11 - row 11
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q11_other_explain"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": "Andere, bitte nennen"}
				inp.ColSpan = 5
				inp.ColSpanLabel = 1.9
				inp.ColSpanControl = 2.1
				inp.ControlTop()
				inp.LabelPadRight()
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q11_other")
				inp.ColSpan = 2
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "%"}
				inp.Placeholder = trl.S{"de": "00"}
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleBox.Padding = "0 0 0 0.3rem"
			}

		}
	}

	// page 5
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p5"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("48rem")

		//
		//
		//
		// gr0, gr1
		var columnTemplateLocal = []float32{
			3.6, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
		}
		// additional row below each block
		colsBelow1 := append([]float32{1.0}, columnTemplateLocal...)
		colsBelow1 = []float32{
			// 1.4, 2.2, //   3.0, 1,  |  4.6 separated to two cols
			1.38, 2.1, //   3.0, 1,  |  4.6 separated to two cols
			0.0, 1, //     3.0, 1,  |  4.6 separated to two cols
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
		}
		colsBelow2 := []float32{}
		for i := 0; i < len(colsBelow1); i += 2 {
			colsBelow2 = append(colsBelow2, colsBelow1[i]+colsBelow1[i+1])
		}

		// for idx, assCl := range inputNamesAssetClassesChangeQ3 {
		{

			inpNames := []string{}
			for _, nm := range q12inputNames {
				inpNames = append(inpNames, "q12_"+nm)
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					oneToFiveNumbers,
					inpNames,
					radioVals5,
					q12Labels,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
						<b>12. </b> &nbsp;	
						
						Welche Themenfelder decken Ihre Investitionen ab? Bitte tragen Sie soweit möglich, anteilige Impact Investitionsvolumina ein.

						<!--
						<br>
						<br>
						(Mehrfachauswahl in der Reihenfolge der investierten Volumina. 1 bis 5, 1= höchstes Volumen)
						-->

					`),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 1
			}

			{

				gr := page.AddGroup()
				gr.Cols = 7
				gr.BottomVSpacers = 4
				stl := ""
				for colIdx := 0; colIdx < len(colsBelow2); colIdx++ {
					stl = fmt.Sprintf(
						"%v   %vfr ",
						stl,
						colsBelow2[colIdx],
					)
				}
				gr.Style = css.NewStylesResponsive(gr.Style)
				if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
					gr.Style.Desktop.StyleBox.Display = "grid"
					gr.Style.Desktop.StyleGridContainer.TemplateColumns = stl
					// log.Printf("fmt special 2021-09: grid template - %v", stl)
				} else {
					log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
				}

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q12__other_explain"
					inp.MaxChars = 17
					inp.ColSpan = 1
					inp.ColSpanLabel = 2.4
					inp.ColSpanControl = 4
					// inp.Placeholder = trl.S{"de": "Andere: Welche?", "en": "Other: Which?"}
					inp.Label = trl.S{
						"de": "Andere",
						"en": "Other",
					}

					// inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
					// inp.StyleCtl.Desktop.StyleBox.WidthMax = "14.0rem"
					// inp.StyleCtl.Mobile.StyleBox.WidthMax = "4.0rem"

				}

				//
				for idx := 0; idx < len(oneToFiveNumbers); idx++ {
					rad := gr.AddInput()
					rad.Type = "radio"

					rad.Name = "q12" + "__other"
					rad.ValueRadio = fmt.Sprint(idx + 1)

					rad.ColSpan = 1
					rad.ColSpanLabel = colsBelow1[2*(idx+1)]
					rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1
				}

			}

		}

	} // page4
}
