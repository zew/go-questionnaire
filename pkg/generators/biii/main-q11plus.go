package biii

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func page4Quest11(q *qst.QuestionnaireT) {

	// page 5
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p4"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
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

	//
	//
	// page 6
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p5"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
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

	}

	// page 7
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p6"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true

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

	// page 8
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p7"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

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
				rad.Name = fmt.Sprintf("q16_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				// rad.Style.Desktop.StyleBox.Margin = "0 0 0 2.4rem"

				rad.ControlFirst()
			}
		}

	}

	// page 9
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p8"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

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

	}

	// page 10
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p9"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir messen den Impact unserer Investments nicht "},
				{"de": "Wir messen negative Externalitäten"},
				{"de": "Wir messen den positiven Impact anhand klar definierter KPIs (Key Performance Indicators)"},
				{"de": "Wir haben Zielvereinbarungen nach bestimmten Indikatoren und Controlling "},
				{"de": "Wir messen sowohl den Unternehmens- als auch den Investoren-Impact"},
				{"de": "Wir messen unseren zusätzlichen Beitrag (Additionalität) im Vergleich zu einem Base-line-Szenario"},
				{"de": "Wir erstellen eine Gesamtbilanz unseres negativen und positiven Impacts (net-impact)"},
			}
			subName := []string{
				"mo_measure",
				"neg_externalities",
				"kpis",
				"contracted_objectives",
				"impact_comp_investor",
				"additionality",
				"total_score",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `

					<p style='font-size: 130%'>
						Impact Messung und Management (IMM)
					</p>

					<p style='font-size: 110%; text-align: justify'>
					Im folgenden Teil wollen wir einen Überblick über die am Markt verwendeten Messungs- und Managementstrategien erheben. 
					</p>

					<br>


					<b>20.</b> &nbsp;	
					Wie messen Sie den Impact? 
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q20_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				rad.ControlFirst()
			}
		}

		// gr1
		{
			labels := []trl.S{
				{"de": "Wir managen den Impact unserer Investments nicht"},
				{"de": "Wir haben eine klare, auf Wirkungsziele ausgerichtete Investmentstrategie "},
				{"de": "Wir haben eine Theory of Change für unsere Impact Investing-Aktivitäten"},
				{"de": "Es gibt einen klaren Impact Management Prozess "},
				{"de": "Wir nutzen eine Materiality Matrix (Wesentlichkeits Matrix) "},
				{"de": "Die Wirkung des Investments wird über den gesamten Investmentprozess hinweg (von der Auswahl bis zum Reporting) als zentrales Kriterium berücksichtigt"},
				{"de": "Wir haben ein eigenständiges und regelmäßiges Impact Reporting"},
				{"de": "Die Erreichung der Impact-Ziele wird bei der Managementvergütung berücksichtigt"},
			}
			subName := []string{
				"none",
				"strategy",
				"theory_of_change",
				"impact_mgt_process",

				"materiality_matrix",
				"whole_lifecycle",
				"own_reporting",
				"coupled_salaries",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>21.</b> &nbsp;	
					Wie managen Sie den Impact Ihrer Investments?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q21_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				rad.ControlFirst()
			}
		}

	}

	// page 11
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p10"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIIMeasure"
		page.SuppressInProgressbar = true

		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir nutzen öffentlich verfügbare Daten der Investees (Jahresbericht, Website, Einzelreports, Investorenkommunikation etc.)"},
				{"de": "Wir identifizieren den Impact anhand der Unternehmensziele der Investees "},
				{"de": "Wir verwenden die Daten, die von den Investees erhoben wurden"},
				{"de": "Wir erheben eigene Daten zum Investee"},

				{"de": "Wir greifen auf die Daten externer Datenanbieter zurück "},
				{"de": "Wir führen eigene empirische Untersuchungen zum Impact des Investments durch"},
				{"de": "Wir greifen auf unabhängige empirische Untersuchungen zurück "},
				{"de": "Wir nutzen Labels und Zertifizierungen "},
			}
			subName := []string{
				"public_sources",
				"investee_objectives",
				"investee_data",
				"own_data",

				"third_party",
				"own_inquiries",
				"independent_inquiries",
				"certificates",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>22.</b> &nbsp;	
					Woher kommen die Daten für die Messung des Impacts?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q22_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				rad.ControlFirst()
			}
		}

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir nutzen keine anerkannten Rahmenwerke zur Impact Messung"},
				{"de": "Wir nutzen eigene Metriken und Indikatoren "},
				{"de": "Principles for Responsible Investment (PRI) "},
				{"de": "Operating Principles for Impact Management "},

				{"de": "SDG Impact Standards "},
				{"de": "EVPA five-steps process"},
				{"de": "Impact Management Project (IMP) 5 dimensions of impact "},
				{"de": "SVI Principles of Social Value and SROI "},

				{"de": "GIIN Compass "},
				{"de": "Theory of Change (ToC) "},
				{"de": "GRI "},
				{"de": "BLab assessment (B Corp) "},

				{"de": "IRIS+ "},
				{"de": "EU Taxonomie"},
			}
			subName := []string{
				"none",
				"own",
				"pri",
				"op_for_im",

				"sdg",
				"evpa",
				"five_dims",
				"svi",

				"giin",
				"toc",
				"gri",
				"blab",

				"irisplus",
				"eu_taxonomy",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>23.</b> &nbsp;	
					Welche der folgenden Rahmenwerke nutzen Sie für das IMM?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q23_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				rad.ControlFirst()
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q23other"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": "Andere, bitte nennen"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0 0 3.4rem"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "1.2rem 0 0 0"
			}
		}

		// gr2
		{
			labels := []trl.S{
				{"de": "Ja, unser Impact Messungs- und Managementsystem wird durch eine externe Prüfung verifiziert"},
				{"de": "Ja, unsere Impact Performance wird durch eine externe Prüfung verifiziert"},
				{"de": "Nein, aber wir ziehen es in Betracht"},
				{"de": "Keiner der oben genannten Punkte"},
			}
			radioValues := []string{
				"impact_and_mgt",
				"performance",
				"considering",
				"none",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `<b>24.</b> &nbsp;	
					Lassen Sie Ihren sozialen oder ökologischen Impact durch eine externe Prüfung verifizieren? 
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, labl := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q24"
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
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": `
						<b>25.</b> &nbsp;	Wie bewerten Sie die Performance Ihres Impact-Portfolios?
					`,
				}
			}
		}

		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3,
				q25Columns,
				[]string{"q25performance"},
				q25RadioVals,
				[]trl.S{
					{"de": `
						Finanzielle Performance
					`},
				},
			)
			// gb.MainLabel = ...
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 2

		}
		// gr5
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3,
				q25Columns,
				[]string{"q25impact"},
				q25RadioVals,
				[]trl.S{
					{"de": `
						Impact Performance
					`},
				},
			)
			// gb.MainLabel = ...
			gr := page.AddGrid(gb)
			_ = gr
		}

	}

	// page 12
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p9"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("42rem")

		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{"de": `

					<p style='font-size: 130%'>
						Integrität und Regulierung
					</p>
				`}
			inp.ColSpan = gr.Cols
		}

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
			for _, nm := range q26inputNames {
				inpNames = append(inpNames, "q26_"+nm)
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					oneToFiveEfficiency,
					inpNames,
					radioVals5,
					q26Labels,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
						<b>26. </b> &nbsp;	
						Wie bewerten Sie die akftuellen rechtlichen Rahmenbedingungen für Impact Investing in Deutschland und international?
					`),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 3
			}

		}

	}

}
