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
	} // page 4

	//
	//
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
					oneToFiveVolume,
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
				for idx := 0; idx < len(oneToFiveVolume); idx++ {
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

	} // page 5

	// page 6
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p6"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("48rem")

		//
		//
		//
		// gr0
		var columnTemplateLocal = []float32{
			3.6, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
		}

		// for idx, assCl := range inputNamesAssetClassesChangeQ3 {
		{

			inpNames := []string{}
			for _, nm := range q13inputNames {
				inpNames = append(inpNames, "q13_"+nm)
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					oneToFiveImportance,
					inpNames,
					radioVals5,
					q13Labels,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
						<b>13. </b> &nbsp;	
						
						Auf die Erreichung welcher Sustainable Development Goals (SDGs)/ Ziele für nachhaltige Entwicklung der UN arbeiten Sie mit Ihren Investitionen hin?					

						<!--
						<br>
						<br>
						(Mehrfachauswahl in der Reihfolge der Wichtigkeit möglich)
						-->

					`),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 3
			}

		}
	}

	// page 7
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p7"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("48rem")

		// gr1
		{
			labels := []trl.S{
				{"de": "Weniger als 1 Monat"},
				{"de": "1-6 Monate "},
				{"de": "6-12 Monate"},
				{"de": "< 2 Jahre"},
				{"de": "2-4 Jahre"},
				{"de": "4-6 Jahre"},
				{"de": "6-8 Jahre"},
				{"de": "8-10 Jahre"},
				{"de": "10+ Jahre"},
			}
			radioValues := []string{
				"under1month",
				"months1to6",
				"months6to12",
				"under2yrs",
				"years2to4",
				"years4to6",
				"years6to8",
				"years8to10",
				"over10years",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `<b>14.</b> &nbsp;	
					Wie lang ist die durchschnittliche Laufzeit Ihrer Investments?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, labl := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q14"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = labl

				rad.ControlFirst()
			}
		}

		// gr2
		{
			labels := []trl.S{
				{"de": "Unzureichende Managementkompetenz der Investees"},
				{"de": "In der Komplexität des Geschäftsmodells "},
				{"de": "Länder- und Währungsrisiken "},
				{"de": "Liquiditäts- und Ausstiegsrisiko "},
				{"de": "Makroökonomische Risiken"},
			}
			subName := []string{
				"competence_lack",
				"complexity",
				"country_currency",
				"liquidy_liquidation",
				"macro",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>15.</b> &nbsp;	Wo liegen bei einem Impact Investment die größten finanziellen Risiken?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q15_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				// rad.Style.Desktop.StyleBox.Margin = "0 0 0 2.4rem"

				rad.ControlFirst()
			}
		}

		// gr3
		{
			labels := []trl.S{
				{"de": "Das Risiko, dass ein negativer Impact erzeugt wird"},
				{"de": "Das Risiko, dass die zuvor festgelegten sozialen und ökologischen Impact-Ziele nicht erreicht werden"},
				{"de": `Das "mission drift" Risiko; d.h. finanzielle Aspekte verdrängen die ursprünglichen Impact-Ziele `},
				{"de": "Es erfolgt keine kontinuierliche Impact Evaluation"},
			}
			subName := []string{
				"negative_impact",
				"underachievement",
				"mission_drift",
				"no_evaluation",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>16.</b> &nbsp; Überwachen Sie:"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q15_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				// rad.Style.Desktop.StyleBox.Margin = "0 0 0 2.4rem"

				rad.ControlFirst()
			}
		}

	} // page 7

	// page 8
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p8"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("48rem")

		// gr1
		{
			labels := []trl.S{
				{"de": "Direktinvestitionen (von der Organisation selbst verwaltet, d.h. Ihre Organisation verwaltet Impact Investitionen)"},
				{"de": "Indirekte Investitionen (d.h. Ihre Organisation investiert oder vermittelt über Fonds / Programme Dritter, die von Dritten verwaltet werden und in Impact investieren)"},
				{"de": "Vertrieb von Impact-Investment-Fonds, die von anderen verwaltet werden (d.h. Ihre Organisation vertreibt Fonds, die in anderen Ländern oder von anderen Organisationen verwaltet werden) "},
			}
			radioValues := []string{
				"direct",
				"indirect",
				"fund",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `<b>17.</b> &nbsp;	
					Investieren Sie direkt oder indirekt? 
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, labl := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q17"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = labl

				rad.ControlFirst()
			}
		}

		// gr2
		{
			labels := []trl.S{
				{"de": "Investees, die zur Lösung spezifischer sozialer und/oder ökologischer Herausforderungen beitragen, die ansonsten benachteiligte Bevölkerungsgruppen und/oder den Planeten betreffen"},
				{"de": "Investees, die positive Auswirkungen für die Menschen und/oder den Planeten erzielen"},
				{"de": "Investees, die Tätigkeiten mit erheblichen negativen Auswirkungen auf die Menschen und/oder den Planeten ausschließen"},
				{"de": "Investees, die soziale und/oder ökologische Daten nutzen, um den finanziellen Wert mittel- und langfristig zu maximieren"},
			}
			radioValues := []string{
				"challengers",
				"people",
				"prevention",
				"data_driven",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `<b>18.</b> &nbsp;	
					In welche Art von Investees fließt der Großteil Ihrer Impact Investitionen? 
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, labl := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q18"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = labl

				rad.ControlFirst()
			}
		}

		// gr3
		{
			labels := []trl.S{
				{"de": "Traditionelle Organisationen mit beabsichtigtem (intentional) sozialen und/oder ökologischen Impact"},
				{"de": "Gewinnorientierte Organisationen mit sozialem Auftrag ohne Gewinnsperre"},
				{"de": "Gewinnorientierte Organisationen mit sozialem Auftrag mit Gewinnsperre"},
				{"de": "Gemeinnützige Organisationen mit kommerziellen Aktivitäten"},
				{"de": "Gemeinnützige Organisationen ohne kommerzielle Aktivitäten"},
				{"de": "Nicht anwendbar (N/A) "},
			}
			subName := []string{
				"traditonal",
				"for_profit_statute",
				"non_profit_statute",
				"charity_commercial",
				"charity_non_commercial",
				"not_applicable",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>19.</b> &nbsp;	Welche Art von Organisation(en) unterstützen Sie?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q19_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				// rad.Style.Desktop.StyleBox.Margin = "0 0 0 2.4rem"

				rad.ControlFirst()
			}
		}

	} // page 8
}
