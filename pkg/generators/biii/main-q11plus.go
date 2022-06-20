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
				{"de": `Wirkungsorientierte Finanzierungsinstrumente (SOC (z.B. SIB / DIB) ) 
						<div style="font-size:80%; line-height: 100%; margin-top: 0.3rem; margin-left: 1rem;">
							SOC: Social Outcomes Contracting; 
							<br>
							SIB: Social Impact Bond; 
							<br>
							DIB: Development Impact Bond
						</div>
					
				`},
				{"de": "Grüne Anleihen (Green Bonds)  "},

				{"de": "Schwellen- und Entwicklungsländer (Emerging markets)  "},
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
					
					Bitte tragen Sie soweit möglich ungefähre Impact Investitionsvolumina nach Anlageklassen/Instrumenten (in Prozent) ein. (Investitionsvolumina können mehrfach zugeordnet werden)
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
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("44rem")

		page.ValidationFuncName = "biiiQ12"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		{
			mainLbl := trl.S{
				"de": fmt.Sprintf(`
						<b>12. </b> &nbsp;	
						
						Welche Themenfelder decken Ihre Investitionen ab? 
						<br>
						<b>(Wählen Sie die Top Fünf (maximal). Optional mit Investmentvolumina)</b>
					
						<!--
						<br>
						<br>
						(Wählen Sie bis zu fünf)
						
						-->

					`),
			}
			gr := page.AddBiiiPrio(mainLbl, q12Labels, q12inputNames, map[int]bool{16: true}, 1)
			_ = gr
			// gr.WidthMax("38rem")

		}

	}

	// page 7
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - p6"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true

		page.WidthMax("34rem")

		page.ValidationFuncName = "biiiQ13"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		{
			mainLbl := trl.S{
				"de": fmt.Sprintf(`
						<b>13. </b> &nbsp;	
						Auf die Erreichung welcher Sustainable Development Goals (SDGs)/ Ziele für nachhaltige Entwicklung der UN arbeiten Sie mit Ihren Investitionen hin?					
						<br>
						<b>(Wählen Sie die Top Fünf (maximal))</b>

						<!--
						<br>
						(Wählen Sie bis zu fünf)
						-->
					`),
			}
			gr := page.AddBiiiPrio(mainLbl, q13Labels, q13inputNames, map[int]bool{}, 0)
			_ = gr
			// gr.WidthMax("38rem")
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
				{"de": "Komplexität des Geschäftsmodells "},
				{"de": "Länder- und Währungsrisiken "},
				{"de": "Liquiditäts- und Exit-Risiko"},
				{"de": "Makroökonomische Risiken"},
				{"de": "Andere, bitte nennen"},
			}
			subName := []string{
				"competence_lack",
				"complexity",
				"country_currency",
				"liquidy_liquidation",
				"macro",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					<b>15.</b> &nbsp;	
					Wo liegen bei einem Impact Investment die größten finanziellen Risiken?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q15_%v", subName[idx])

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q15_other_label"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}
			// for idx, label := range labels {
			// 	rad := gr.AddInput()
			// 	rad.Type = "checkbox"
			// 	rad.Name = fmt.Sprintf("q15_%v", subName[idx])

			// 	rad.ColSpan = 1
			// 	rad.ColSpanLabel = 1
			// 	rad.ColSpanControl = 6

			// 	rad.Label = label

			// 	rad.Style = css.NewStylesResponsive(rad.Style)
			// 	// rad.Style.Desktop.StyleBox.Margin = "0 0 0 2.4rem"

			// 	rad.ControlFirst()
			// }
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
				inp.Label = trl.S{
					"de": `
					<b>16.</b> &nbsp; 
					Überwachen Sie folgende Risiken in Bezug auf Ihre Impact Investments?
					<br>
					(Mehrfachauswahl möglich)
				`}
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
			subName := []string{
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
					<br>
					(Mehrfachauswahl möglich)

				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q17_%v", subName[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.Style = css.NewStylesResponsive(rad.Style)
				// rad.Style.Desktop.StyleBox.Margin = "0 0 0 2.4rem"

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
				{"de": "Gewinnorientierte Organisationen mit sozialem Auftrag ohne Ausschüttungs&shy;sperre"},
				{"de": "Gewinnorientierte Organisationen mit sozialem Auftrag mit Ausschüttungs&shy;sperre"},
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
				inp.Label = trl.S{
					"de": `<b>19.</b> &nbsp;	
					In welche Art von Organisation(en) investieren Sie?
					<br>
					(Mehrfachauswahl möglich)
				`}
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
		page.Short = trl.S{"de": "Messung"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir messen den Impact unserer Investments nicht "},
				{"de": "Wir messen den negativen Impact (Praktiken, die der Gesellschaft und/oder der Umwelt schaden)"},
				{"de": "Wir messen den positiven Impact anhand klar definierter KPIs (Key Performance Indicators)"},
				{"de": "Wir haben Zielvereinbarungen nach bestimmten Indikatoren und kontrollieren deren Einhaltung"},
				{"de": "Wir messen sowohl den Unternehmens- als auch den Investoren-Impact"},
				{"de": "Wir messen unseren zusätzlichen Beitrag (Additionalität) im Vergleich zu einem Base-Line-Szenario"},
				{"de": "Wir erstellen eine Gesamtbilanz unseres negativen und positiven Impacts (net-impact)"},
				{"de": "Andere, bitte nennen"},
			}
			subName := []string{
				"mo_measure",
				"neg_externalities",
				"kpis",
				"contracted_objectives",
				"impact_comp_investor",
				"additionality",
				"total_score",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `

					<p style='font-size: 130%;font-weight: bold'>
						Impact Messung und Management (IMM)
					</p>

					<p style='font-size: 110%; text-align: justify'>
					Im folgenden Teil wollen wir Daten zu den am Markt verwendeten Messungs- und Managementstrategien erheben.
					</p>

					<br>


					<b>20.</b> &nbsp;	
					Wie messen Sie den Impact Ihrer Impact Investments?? 
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q20_%v", subName[idx])

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q20_other_label"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

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
					Wie managen Sie den Impact Ihrer Impact Investments?
					<br>
					(Mehrfachauswahl möglich)
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
					<br>
					(Mehrfachauswahl möglich)
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
				{"de": "Andere, bitte nennen"},
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
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>23.</b> &nbsp;	
					Welche der folgenden Rahmenwerke nutzen Sie für das IMM?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q23_%v", subName[idx])

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q23_other_label"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

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
					Lassen Sie Ihren sozialen und/oder ökologischen Impact durch eine externe Prüfung verifizieren? 
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

	}

	// page 12
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"

		page.SuppressInProgressbar = true

		page.WidthMax("42rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3,
				q25Columns,
				[]string{"q25performance", "q25impact"},
				q25RadioVals,
				[]trl.S{
					{"de": `Finanzielle Performance`},
					{"de": `Impact Performance`},
				},
			)
			gb.MainLabel = trl.S{
				"de": `
					<b>25.</b> &nbsp;	Wie bewerten Sie die Performance Ihres Impact-Portfolios?
				`,
			}
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 4
		}

	}

	// page 13
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "Integrität<br>Entwicklung"}
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

					<p style='font-size: 130%; font-weight: bold;'>
						Integrität und Regulierung
					</p>
					Im folgenden Teil wollen wir Daten zu den aktuellen regulatorischen Rahmenbedingungen und der Integrität des Marktes erheben.
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
						Wie bewerten Sie die aktuellen rechtlichen Rahmenbedingungen für Impact Investing in Deutschland und international?
					`),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 2
			}

		}
		{
			gr := page.AddGroup()
			gr.Cols = 7
			gr.BottomVSpacers = 4

			inp := gr.AddInput()
			inp.Type = "textarea"

			inp.Name = "q26_comment"
			inp.MaxChars = 150
			inp.Placeholder = trl.S{
				"de": "Optionaler Kommentar ",
			}
			inp.Label = trl.S{
				"de": "Optionaler Kommentar ",
			}
			// inp.Label = label

			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 4
			inp.ColSpanControl = 5
			// inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)

		}

		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				q27columns,
				[]string{
					"q27_sfrd",
					"q27_mifid2",
					"q27_eu_taxonomy",
					"q27_opfim",
				},
				radioVals5,
				[]trl.S{
					{"de": `SFRD`},
					{"de": `MiFID II`},
					{"de": `EU Taxonomy`},
					{"de": `Operating Principles for Impact Management`},
				},
			)
			gb.MainLabel = trl.S{
				"de": `
					<b>27.</b> &nbsp;
					Wie schätzen Sie die regulatorischen Anforderungen der folgenden Richtlinien für die Praxis des Impact Investings ein? 
				`,
			}
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 1
		}

		{
			colsBelow1 := append([]float32{1.0}, columnTemplate5...)
			colsBelow1 = []float32{
				1.4, 2.2, //   3.0, 1,  |  4.6 separated to two cols
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
			}
			colsBelow2 := []float32{}
			for i := 0; i < len(colsBelow1); i += 2 {
				colsBelow2 = append(colsBelow2, colsBelow1[i]+colsBelow1[i+1])
			}

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
			} else {
				log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q27_other_label"
				inp.MaxChars = 17
				inp.ColSpan = 1
				inp.ColSpanLabel = 2.4
				inp.ColSpanControl = 4
				inp.Label = trl.S{
					"de": "Andere",
					"en": "Other",
				}
			}
			for idx := 0; idx < len(oneToFiveVolume); idx++ {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q27" + "_other"
				rad.ValueRadio = fmt.Sprint(idx + 1)
				rad.ColSpan = 1
				rad.ColSpanLabel = colsBelow1[2*(idx+1)]
				rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1
			}
		}

		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				oneToFiveEfficiency,
				[]string{
					"q28_definitions",
					"q28_guidelines",
					"q28_transparency",
					"q28_objectives",
					"q28_measurments",
					"q28_data",
					"q28_duedilligence",
				},
				radioVals5,
				[]trl.S{
					{"de": "Einheitliche Definition von Impact Investing "},
					{"de": "Verpflichtende Impact Messungs- und Reporting-Bestimmungen "},
					{"de": "Mehr Wissen und Transparenz "},
					{"de": "Klare Impact Ziele in Strategie und Entscheidungsprozessen"},
					{"de": "Entwicklung von Messinstrumenten über den gesamten Impact-Messungs&shy;prozess "},
					{"de": "Effektive Datenerhebung, -speicherung und -validierung"},
					{"de": "Obligatorische Due Diligence für Impact"},
				},
			)
			gb.MainLabel = trl.S{
				"de": `
					<b>28.</b> &nbsp;
					Wie effektiv sind Ihrer Meinung nach die folgenden Maßnahmen, um Impact Washing zu verhindern? 
				`,
			}
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 1
		}

		{
			colsBelow1 := append([]float32{1.0}, columnTemplate5...)
			colsBelow1 = []float32{
				1.4, 2.2, //   3.0, 1,  |  4.6 separated to two cols
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
			}
			colsBelow2 := []float32{}
			for i := 0; i < len(colsBelow1); i += 2 {
				colsBelow2 = append(colsBelow2, colsBelow1[i]+colsBelow1[i+1])
			}

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
			} else {
				log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q28_other_label"
				inp.MaxChars = 17
				inp.ColSpan = 1
				inp.ColSpanLabel = 2.4
				inp.ColSpanControl = 4
				inp.Label = trl.S{
					"de": "Andere",
					"en": "Other",
				}
			}
			for idx := 0; idx < len(oneToFiveVolume); idx++ {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q28" + "_other"
				rad.ValueRadio = fmt.Sprint(idx + 1)
				rad.ColSpan = 1
				rad.ColSpanLabel = colsBelow1[2*(idx+1)]
				rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1
			}
		}

	}

	// page 14
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr 0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<p style='font-size: 130%; font-weight: bold'>
						Entwicklung des Impact Investing-Marktes
					</p>

					<br>
				`}
				inp.ColSpan = gr.Cols
			}
		}

		// gr1
		{
			labels := []trl.S{
				{"de": "Wir waren vor 2&nbsp;Jahren noch nicht am Markt "},
				{"de": "Wir sind in 2&nbsp;Jahren gewachsen - um "},
				{"de": "Unsere Anlagesumme ist in den letzten 2&nbsp;Jahren gleichgeblieben"},
				{"de": "Wir sind in 2&nbsp;Jahren geschrumpft - um "},
			}
			radioValues := []string{
				"younger_than2",
				"growth",
				"unchanged",
				"shrinking",
			}
			gr := page.AddGroup()
			gr.Cols = 5
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>29.</b> &nbsp;	
					Wie groß war das Wachstum Ihrer Anlagesumme 
					in Impact Investing in den letzten zwei Jahren?

				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q29"
				rad.ValueRadio = radioValues[idx]

				if idx%2 == 1 {
					rad.ColSpan = 3
				} else {
					rad.ColSpan = 3
				}
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()

				if idx%2 == 1 {
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("q29_%v_pct", radioValues[idx])
					// inp.Label = label
					inp.ColSpan = 2
					// inp.ColSpanLabel = 16
					inp.ColSpanControl = 7
					inp.Min = 0
					inp.Max = 100
					inp.Step = 0.1
					inp.MaxChars = 5
					inp.Suffix = trl.S{"de": "%"}
					inp.Placeholder = trl.S{"de": "00"}
					inp.LabelPadRight()

					inp.Style = css.NewStylesResponsive(inp.Style)
					inp.Style.Desktop.StyleBox.Position = "relative"
					inp.Style.Desktop.StyleBox.Left = "-3rem"
				}
			}
		}

		// gr 2
		{
			{
				gb := qst.NewGridBuilderRadios(
					columnTemplate5,
					q30columns,
					[]string{
						"q30_management",
						"q30_measurement",
						"q30_methods",
						"q30_defintions",

						"q30_benchmarks",
						"q30_certificates",
						"q30_labels",
						"q30_kpis",

						"q30_data",
						"q30_ratings",
						"q30_best_pract",
						"q30_reporting",

						"q30_legal_frame",
					},
					radioVals5,
					[]trl.S{
						{"de": "Impact Management"},
						{"de": "Impact Messung "},
						{"de": "Methoden zur Impact Messung "},
						{"de": "Impact Definition "},

						{"de": "Impact Benchmarks"},
						{"de": "Impact Zertifizierung "},
						{"de": "Impact Labels"},
						{"de": "Katalog von Kriterien und KPIs"},

						{"de": "Impact Datenverfügbarkeit "},
						{"de": "Impact Ratings"},
						{"de": "Best Practices"},
						{"de": "Berichterstattung "},

						{"de": "Gesetzlicher Rahmen"},
					},
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
						<b>30. </b> &nbsp;	
						Welchen Fortschritt hat es in letzten drei Jahren in den folgenden Bereichen gegeben?
					`),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 3
			}

		}

	}

	// page 15
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				q31columns,
				[]string{
					"q31_regulation",
					"q31_transparency",
					"q31_intermediaries",
					"q31_capacities",

					// "q31_bestpract",
					"q31_methodoloy",
					// "q31_integration",
					"q31_quality",

					"q31_differentiation",
					"q31_demand",
					"q31_incentives",
					"q31_successstories",

					"q31_certifications",
				},
				radioVals5,
				[]trl.S{
					{"de": "Regulatorischer Rahmen "},
					{"de": "Verbesserte Informationslage und Markttransparenz "},
					{"de": "Entwicklung dynamischer intermediärer Strukturen"},
					{"de": "Impact-Managementkapazitäten der Investees"},

					// {"de": "Fähigkeit der Investoren zur Implementierung von Best Practices"},
					{"de": "Entwicklung einer standardisierten Methodik zur Impact-Messung und -Steuerung"},
					// {"de": "Integration von Impact Management und Messung in alle Investitionsprozesse "},
					{"de": "Qualitativ hochwertige Angebote von Impact Investment Produkten "},

					{"de": "Ausdifferenzierung von Impact-Anspruchsniveaus in Asset Klassen "},
					{"de": "Hohe Nachfrage durch Investees / Sozialunternehmen"},
					{"de": "Staatliche Anreize zur Unterstützung des Impact Investing"},
					{"de": "Verbreitung von Erfolgsgeschichten "},

					{"de": "Entwicklung von Impact Labels und Zertifizierungen"},
				},
			)
			gb.MainLabel = trl.S{
				"de": `
					<b>31.</b> &nbsp;
					Für wie relevant erachten Sie die folgenden Rahmenbedingungen für die weitere Entwicklung und das Wachstum des Impact Investment Sektors? 
				`,
			}
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 1
		}
		{
			colsBelow1 := append([]float32{1.0}, columnTemplate5...)
			colsBelow1 = []float32{
				1.4, 2.2, //   3.0, 1,  |  4.6 separated to two cols
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
			}
			colsBelow2 := []float32{}
			for i := 0; i < len(colsBelow1); i += 2 {
				colsBelow2 = append(colsBelow2, colsBelow1[i]+colsBelow1[i+1])
			}

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
			} else {
				log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q31_other_label"
				inp.MaxChars = 17
				inp.ColSpan = 1
				inp.ColSpanLabel = 2.4
				inp.ColSpanControl = 4
				inp.Label = trl.S{
					"de": "Andere",
					"en": "Other",
				}
			}
			for idx := 0; idx < len(oneToFiveVolume); idx++ {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q31" + "_other"
				rad.ValueRadio = fmt.Sprint(idx + 1)
				rad.ColSpan = 1
				rad.ColSpanLabel = colsBelow1[2*(idx+1)]
				rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1
			}
		}

	}

	// page 16
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		page.ValidationFuncName = "biiiQ32"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		{

			q32inputNames := []string{
				"q32_method_mgt_meas",
				"q32_definition",
				"q32_certifications",
				"q32_education",

				"q32_data",
				"q32_legislation",
				"q32_bestpract",
				"q32_reporting",

				"q32_assetclass",
				"q32_productdesign",
				"q32_other",
			}
			q32Labels := []trl.S{
				{"de": "Impact Methodologie, Management und Messung "},
				{"de": "Harmonisierte Impact Definition "},
				{"de": "Impact Labels, Ratings und Zertifizierungen "},
				{"de": "Bildung und Weiterbildung"},

				{"de": "Impact Datenverfügbarkeit und Kriterienkatalog "},
				{"de": "Gesetzgebung Impact  "},
				{"de": "Benchmarks und Best Practices "},
				{"de": "Berichterstattung "},

				{"de": "Impact Investing als Anlageklasse  "},
				{"de": "Produktgestaltung"},
				{"de": "Andere, bitte nennen"},
			}

			mainLbl := trl.S{
				"de": fmt.Sprintf(`
						<b>32.</b> &nbsp;
						In welchen Bereichen sind die größten Fortschritte notwendig?
						<br>
						<b>
						(Wählen Sie die Top Fünf. Optional mit Priorisierung (1=größter Fortschritt))
						</b>
					`),
			}
			gr := page.AddBiiiPrio(mainLbl, q32Labels, q32inputNames, map[int]bool{10: true}, 2)
			_ = gr
			// gr.WidthMax("38rem")
		}

	}

	// page 17
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("44rem")

		page.ValidationFuncName = "biiiQ33"

		{

			q33inputNames := []string{
				"q33_paris",
				"q33_leisure",
				"q33_education",
				"q33_work",

				"q33_research",
				"q33_health",
				"q33_social_services",
				"q33_env_protection",

				"q33_oceans",
				"q33_wash",
				"q33_agriculture",
				"q33_energy",

				"q33_residential",
				"q33_it",
				"q33_production",
				"q33_urban_dev",

				"q33_financial_access",
				"q33_other",
			}
			q33Labels := []trl.S{
				{"de": "Paris-Aligned oder Net Zero"},
				{"de": "Kultur und Freizeit (Kultur, Kunst, Sport, sonstige Freizeitgestaltung und soziale Vereine)"},
				{"de": "Bildung (Grundschule, Sekundarschule, Hochschule, Sonstiges)"},
				{"de": "Arbeitsmarktintegration"},

				{"de": "Forschung"},
				{"de": "Gesundheit (Krankenhäuser, Rehabilitation, Pflegeheime, psychische Gesundheit / Krisenintervention)"},
				{"de": "Soziale Dienste (Notfall, Hilfe, Einkommensunterstützung / Unterhalt)"},
				{"de": "Umweltschutz (Forstwirtschaft, Land, Abfall, Luft, biologische Vielfalt und Ökosysteme,"},

				{"de": "Meere und Küstengebiete)"},
				{"de": "WASH (Wasser, Sanitärversorgung und Hygiene)"},
				{"de": "Landwirtschaft"},
				{"de": "Energie (Zugang zu Energie, erneuerbare Energie)"},

				{"de": "Wohnen"},
				{"de": "IT / Technologien"},
				{"de": "Fertigung / Produktion"},
				{"de": "Stadterneuerung / Territoriale Entwicklung"},

				{"de": "Finanzielle Eingliederung und Zugang zu Finanzmitteln (d.h. Mikrofinanzierung, Mikroversicherungen, Finanz Bildungsdienstleistungen, Bankwesen)"},
				{"de": "Andere, bitte nennen"},
			}

			mainLbl := trl.S{
				"de": fmt.Sprintf(`

					<b>33.</b> &nbsp;

					In welchen Themenfeldern sehen Sie 
						a) den <b>größten Bedarf an Kapital</b>, 
						b) das <b>größte Potential</b> 
					für Impact Investments? 
					<br>
					<b>
					(Wählen Sie die Top Fünf. Optional mit Priorisierung (1=größter Bedarf/ Potential))
					</b>



					<table style="margin-left: -0.4rem;margin-bottom: -0.45rem;">
						<tr>

							<td style="width:43%%">

								&nbsp;
							</td>
							<td style="width:28%%">
								a) größten Bedarf an Kapital/
							</td>
							<td style="width:28%%">
								b) größtes Potential 
							</td>


						</tr>
					</table>



				`),
			}
			gr := page.AddBiiiPrio2Cols(mainLbl, q33Labels, q33inputNames, map[int]bool{17: true})
			_ = gr
		}

	}

	// page 18
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		page.ValidationFuncName = "biiiQ34"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		{
			q34inputNames := []string{
				"q34_private_equity",
				"q34_public_equity",
				"q34_private_external",
				"q34_public_external",

				"q34_real_estate",
				"q34_money_deposits",
				"q34_soc",
				"q34_green_bonds",

				"q34_emerging_markets",
				"q34_microfinance",
				"q34_commodities",
				"q34_slb",

				"q34_hybrid",
				"q34_other",
			}
			q34Labels := []trl.S{
				{"de": "Privates (nicht börsengehandeltes) Beteiligungskapital (Private equity)"},
				{"de": "Börsengehandeltes Beteiligungskapital / Aktien (Public equity)"},
				{"de": "Private (nicht börsengehandeltes) Fremdkapital/ Anleihen (Private debt)"},
				{"de": "Börsengehandeltes Fremdkapital / Anleihen (Public debt)  "},

				{"de": "Immobilien (Real estate)  "},
				{"de": "Einlagen oder Zahlungsmitteläquivalente / Geldwerte (Deposits oder cash equivalents / monetary assets) "},
				{"de": `Wirkungsorientierte Finanzierungsinstrumente (SOC (z.B. SIB / DIB) )
							<div style="font-size:80%; line-height: 100%; margin-top: 0.3rem; margin-left: 1rem;">
								SOC: Social Outcomes Contracting;
								<br>
								SIB: Social Impact Bond;
								<br>
								DIB: Development Impact Bond
							</div>

						`},
				{"de": "Grüne Anleihen (Green Bonds)  "},

				{"de": "Schwellen- und Entwicklungsländer (Emerging markets)  "},
				{"de": "Mikrofinanzierung (Microfinance)  "},
				{"de": "Rohstoffe (Commodities)  "},
				{"de": "Sustainability-Linked Bonds (SLBs)"},

				{"de": "Hybride Finanz&shy;instrumente (Hybrid financial instruments), bitte nennen"},
				{"de": "Andere, bitte nennen"},
			}

			mainLbl := trl.S{
				"de": fmt.Sprintf(`
						<b>34.</b> &nbsp;
						In welchen Asset Klassen erwarten Sie eine besonders dynamische Entwicklung?
						<br>
						<b>
						(Wählen Sie die Top Fünf. Optional mit Priorisierung (1=dynamischste Entwicklung))
						</b>
					`),
			}
			gr := page.AddBiiiPrio(mainLbl, q34Labels, q34inputNames, map[int]bool{12: true, 13: true}, 2)
			_ = gr
		}
	}

	// page 19
	{
		page := q.AddPage()
		page.Short = trl.S{"de": ""}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3a,
				q33bColumns,
				[]string{
					"q35_europe",
					"q35_east_asia",
					"q35_north_america",
					"q35_near_east",

					"q35_south_asia",
					"q35_latin_america",
					"q35_sub_sahara",
					"q35_central_asia",
				},
				radioVals3,
				[]trl.S{
					{"de": "Europa"},
					{"de": "Ostasien und Pazifik"},
					{"de": "Nordamerika"},
					{"de": "Naher Osten / Nordafrika"},

					{"de": "Südasien"},
					{"de": "Lateinamerika / Karibik"},
					{"de": "Sub-Sahara Afrika"},
					{"de": "Zentralasien "},
				},
			)
			gb.MainLabel = trl.S{
				"de": `
					<b>35.</b> &nbsp;
					In welchen Regionen sehen Sie das größte Potenzial für effektive Impact Investments?
				`,
			}
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 3
		}

		// gr1
		{
			labels := []trl.S{
				{"de": "Impact Investing ist zu einem relevanten Segment des Kapitalmarkts (> 10 % aller Anlagen) geworden"},
				{"de": "Impact Investing wächst vor allem in der Entwicklungszusammenarbeit"},
				{"de": "Impact Investing wächst vor allem wegen der Implementierung von Retail-Lösungen"},
				{"de": "Impact Investing wächst vor allem durch Investments im Entwicklungsbereich"},

				{"de": "VC nimmt die führende Rolle ein"},
				{"de": "Regulatorische Begrenzungen dämpfen die Dynamik des Impact Investing-Marktes"},
				{"de": "Impact Investing stagniert"},
				{"de": "Andere, bitte nennen"},
			}
			subName := []string{
				"greater10pct",
				"development_coop",
				"retail",
				"development",

				"leading_vc",
				"regulations_dampen",
				"stagnation",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `<b>36.</b> &nbsp;	
                    Wo sehen Sie den Impact Investing Markt in Deutschland in drei Jahren?
					<br>
					(Mehrfachauswahl möglich)
                `}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q36_%v", subName[idx])

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q36_other_label"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}

		}

	}

	// page 20
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "Über Sie"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("38rem")

		page.ValidationFuncName = "biiiPage19"

		// gr0
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{"de": `

					<p style='font-size: 130%; font-weight: bold;'>
						Über Sie und Ihre Organisation 
					</p>
				`}
			inp.ColSpan = gr.Cols
		}

		// gr1
		{
			labels := []trl.S{
				{"de": "<18"},
				{"de": "18-39"},
				{"de": "40-69"},
				{"de": "70+"},
			}
			radioValues := []string{
				"under18",
				"betw18to39",
				"betw40to69",
				"over69",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>37.</b> &nbsp;	
					Wie alt sind Sie?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q37"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
			}
		}

		// gr2
		{
			labels := []trl.S{
				{"de": "Männlich"},
				{"de": "Weiblich"},
				{"de": "Divers"},
			}
			radioValues := []string{
				"male",
				"female",
				"diverse",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>38.</b> &nbsp;	
					Sind sie…?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q38"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
			}
		}

		//
		//
		//
		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 6
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>39.</b> &nbsp;	
					In welchen Regionen/ Ländern investieren Sie Ihr Kapital?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}

			labels := []trl.S{
				{"de": "Deutschland"},
				{"de": "Österreich"},
				{"de": "Schweiz"},
				{"de": "Westeuropa"},
				{"de": "Mittelosteuropa"},
				{"de": "Weltweit"},
			}
			mainNames := []string{
				"germany",
				"austria",
				"switzerland",
				"western_europe",
				"middle_east_eur",
				"worldwide",
			}

			subLabels := map[int][]trl.S{
				3: {
					{"de": "Andorra"},
					{"de": "Belgien"},
					{"de": "Dänemark"},
					{"de": "Finnland"},
					{"de": "Frankreich"},
					{"de": "Griechenland"},
					{"de": "Großbritannien"},
					{"de": "Island"},
					{"de": "Irland"},
					{"de": "Italien"},
					{"de": "Liechtenstein"},
					{"de": "Luxemburg"},
					{"de": "Malta"},
					{"de": "Monaco"},
					{"de": "Niederlanden"},
					{"de": "Norwegen"},
					{"de": "Portugal"},
					{"de": "Schweden"},
					{"de": "Spanien "},
				},
				4: {
					{"de": "Albanien"},
					{"de": "Armenien"},
					{"de": "Aserbaidschan"},
					{"de": "Belarus"},
					{"de": "Bosnien und Herzegowina"},
					{"de": "Estland"},
					{"de": "Georgien"},
					{"de": "Kosovo"},
					{"de": "Kroatien"},
					{"de": "Lettland"},
					{"de": "Litauen"},
					{"de": "Montenegro"},
					{"de": "Nordmazedonien"},
					{"de": "Polen"},
					{"de": "Republik Moldau"},
					{"de": "Rumänien"},
					{"de": "Serbien"},
					{"de": "Slowakei"},
					{"de": "Tschechische Republik"},
					{"de": "Ukraine"},
					{"de": "Ungarn"},
					{"de": "Russland"},
				},
				5: {
					{"de": "Lateinamerika / Karibik"},
					{"de": "Nordamerika"},
					{"de": "Naher Osten / Nordafrika"},
					{"de": "Ostasien und Pazifik"},
					{"de": "Sub-Sahara Afrika"},
					{"de": "Südasien"},
					{"de": "Zentralasien "}},
			}

			subNames := map[int][]string{
				3: {
					"andorra",
					"belgien",
					"denmark",
					"finnland",
					"frankreich",
					"griechenland",
					"grossbritannien",
					"island",
					"irland",
					"italien",
					"liechtenstein",
					"luxemburg",
					"malta",
					"monaco",
					"niederlanden",
					"norwegen",
					"portugal",
					"schweden",
					"spanien",
				},
				4: {
					"albanien",
					"armenien",
					"aserbaidschan",
					"belarus",
					"bosnien_herzeg",
					"estland",
					"georgien",
					"kosovo",
					"kroatien",
					"lettland",
					"litauen",
					"montenegro",
					"nordmazedonien",
					"polen",
					"moldau",
					"rumaenien",
					"serbien",
					"slowakei",
					"tschech_rep",
					"ukraine",
					"ungarn",
					"russland",
				},
				5: {
					"lateinam_karibik",
					"nordamerika",
					"nahost_nordafr",
					"ostasien_paz",
					"subsahara",
					"suedasien",
					"zentralasien",
				},
			}

			for i1, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q39_%v", mainNames[i1])

				rad.ColSpan = gr.Cols
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 11
				rad.Label = label
				rad.ControlFirst()

				//
				if _, ok := subLabels[i1]; ok {
					labels := subLabels[i1]
					names := subNames[i1]
					for i2, label := range labels {
						rad := gr.AddInput()
						rad.Type = "checkbox"
						rad.Name = fmt.Sprintf("q39_sub%v_%v", i1, names[i2])
						rad.ColSpan = 2

						rad.ColSpanLabel = 2
						rad.ColSpanControl = 6

						rad.Label = label
						rad.ControlFirst()

						// log.Printf("%v - %v", names[i2], label)

						rad.Style = css.NewStylesResponsive(rad.Style)
						rad.Style.Desktop.StyleBox.Position = "relative"
						rad.Style.Desktop.StyleBox.Left = "3.2rem"
					}
				} // idx==0

			} // range labels q39
		}

		//
		//
		// gr5
		{
			labels := []trl.S{
				{"de": "Deutschland"},
				{"de": "Österreich"},
				{"de": "Schweiz"},
				{"de": "Andere, bitte nennen"},
			}
			subName := []string{
				"germany",
				"austria",
				"switzerland",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>40.</b> &nbsp;	
					Wo sitzt Ihr Managementteam?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q40_%v", subName[idx])

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q40_other_label"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}

		}

		//
		//
		// gr6
		{
			labels := []trl.S{
				{"de": "Deutschland"},
				{"de": "Österreich"},
				{"de": "Schweiz"},
				{"de": "Andere, bitte nennen"},
			}
			radioValues := []string{
				"germany",
				"austria",
				"switzerland",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			gr.BottomVSpacers = 4
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>41.</b> &nbsp;	
					In welchem Land ist der Hauptsitz Ihrer Organisation?
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q41"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q41_other"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}

		}

	}

}
