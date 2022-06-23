package biii

import (
	"fmt"
	"math"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create creates a JSON file for a financial markets survey
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "BIII", "en": "BIII"}

	q.Survey.Name = trl.S{
		// "de": "Erhebung 2022 - Marktstudie Impact Investment",
		"de": "Marktstudie Impact Investing in Deutschland 2022 der Bundesinitiative Impact Investing",
	}

	// page -1
	{
		page := q.AddPage()

		page.Short = trl.S{"de": "Start"}
		page.Label = trl.S{"de": ""}
		page.WidthMax("42rem")

		page.SuppressProgressbar = true
		page.SuppressInProgressbar = true

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "welcome.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page 0
	{
		page := q.AddPage()

		page.Short = trl.S{"de": "Start"}
		page.Label = trl.S{"de": ""}
		page.WidthMax("42rem")

		page.SuppressProgressbar = true
		page.SuppressInProgressbar = true

		page.ValidationFuncName = "biiiPage0"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		// gr0
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

					<!--
					<p style='text-align: justify; font-size: 130%'>
						Willkommen zur Marktstudie der <a target='_blank' href='https://bundesinitiative-impact-investing.de/'>Bundesinitiative Impact Investing (BIII)</a>
						und der <a target='_blank' href='https://www.wiso.uni-hamburg.de/fachbereich-sozoek/professuren/busch/04-team/busch-timo.html'>Universität Hamburg</a>
						 im Auftrag der AIR GmbH
						  –  Online Umsetzung durch das <a  target='_blank' href='https://zew.de/'>ZEW Mannheim</a>
					</p>
					-->

					<p style='text-align: justify;'>

					Wir nutzen bewusst eine breite Definition von Impact Investments, die das Verständnis des Global Impact Investing Networks (GIIN) widerspiegelt.

					 Demnach sind Impact Investments "Investitionen, die mit der Absicht getätigt werden,
					 neben einer finanziellen Rendite auch eine positive,
					 messbare soziale und ökologische Wirkung zu erzielen" (GIIN, 2017).

					 Auch in akademischen Studien wird ähnlich argumentiert.

					 Busch et al. (2021) bezeichnen Impact bezogene Investments
					 als "Investitionen, die sich auf Veränderungen in der realen Welt
					 im Hinblick auf die Lösung sozialer Probleme und/oder die Milderung
					 ökologischer Schäden konzentrieren".

					 Ob entsprechende Investments aktiv und zusätzlich zu sozialen
					 und ökologischen Lösungen und Veränderungen beitragen, 
					 wird im Rahmen der Erhebung separat ermittelt.

					Wir würden Sie bitten, 
					den Online-Fragebogen <i>vollständig</i> auszufüllen 
					und zu beenden. 
					</p>
					`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "page-0-data-protection.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PermaLink"
			}

		}

		// gr1
		{
			var radioValues = []string{"investor", "investee", "assetmgr", "passiveparticipant"}
			var labels = []trl.S{
				{"de": "Investor<br>(asset owner)"},
				{"de": "Investee (Kapitalempfänger)"},
				{"de": "Intermediär<br>(Vermögensverwalter, Asset Manager)"},
				{"de": "Ein anderer (passiver) Marktteilnehmer (z.B. Berater, ...)"},
			}

			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>1.</b> &nbsp;	Sind Sie…?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q01"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.ColSpan = gr.Cols
				rad.Label = label
				rad.ControlFirst()
			}
		}

		// gr2
		{
			labels := []trl.S{
				{"de": "Privatinvestor"},
				{"de": "Business Angel"},
				{"de": "VC / PE Impact Fondsmanager"},
				{"de": "Privates Finanzinstitut (einschließlich traditioneller Banken und ethischer Banken)"},
				{"de": "Versicherungsgesellschaft oder Pensionsfonds"},
				{"de": "Mikrofinanzinstitution"},
				{"de": "Crowdfunding-Plattform"},
				{"de": "Stiftung"},
				{"de": "Family Office"},
				{"de": "Investmentfondsmanager eines börsennotierten Unternehmens"},
				{"de": "Entwicklungsfinanzierungsagentur oder -einrichtung"},
				{"de": "Öffentlicher Finanzierungsfonds oder -einrichtung"},
				{"de": "Inkubator oder Accelerator"},
				{"de": "Andere, bitte nennen"},
			}
			radioValues := []string{
				"private_investor",
				"business_angel",
				"impact_fund_mgr",
				"private_bank",
				"insurance_fund",
				"micro_institution",
				"crowdfunding_plattform",
				"foundation",
				"family_office",
				"mgr_listed_comp",
				"development_agency",
				"public_fund",
				"incubator",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					<b>2.</b> &nbsp;	
					Als welche Art von Organisation ordnen Sie sich ein?
					`,
				}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q02"
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
					inp.Name = "q02_other"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}
		}

		// gr3
		{
			labels := []trl.S{
				{"de": "Ausschlüsse oder Negative Screening"},
				{"de": "Best-in-Class"},
				{"de": "ESG (ökologische, soziale oder Governance) Integration"},
				{"de": "Norms-based Screening (z.B. UN's Global Compact, OECD Guidelines for Multinational Enterprises, …)"},
				{"de": "Thematische Funds oder Themenbezogene Produkte (z.B. Klima, Menschenrechte, Gesundheit, ...)"},
				{"de": "Impact Investments"},
				{"de": "Immobilienfonds mit Nachhaltigkeitsfokus"},
				{"de": "Wir tätigen keine Investments"},
				{"de": "Andere, bitte nennen"},
			}
			radioValues := []string{
				"exclusions",
				"best_in_class",
				"esg",
				"norms_based",
				"theme_funds",
				"impact_investments",
				"real_estate",
				"no_investing",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					<b>3.</b> &nbsp;
					Welchen Nachhaltigkeits-/Impact-Fokus haben Sie bei Ihrer Investment-Strategie bzw. welche Produktgestaltung nutzen Sie?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q03_%v", radioValues[idx])

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
					inp.Name = "q03_other_label"
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

	// page 1
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Ihre grundsätzliche Position zu Impact Investing"}
		page.Short = trl.S{"de": "Grundposition"}
		page.WidthMax("38rem")

		page.SuppressProgressbar = true

		page.ValidationFuncName = "biiiPage1"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		// gr0
		{
			labels := []trl.S{
				{"de": "Gegenwärtig"},
				{"de": "In Planung"},
				{"de": "Möglicherweise in der Zukunft"},
				{"de": "Nein"},
			}
			radioValues := []string{
				"now",
				"in_planning",
				"in_future",
				"no",
			}
			gr := page.AddGroup()
			gr.Cols = 5
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>4.</b> &nbsp;	Arbeitet Ihre Organisation mit Impact Investments?"}
				inp.ColSpan = gr.Cols
				inp.ColSpan = 4

			}

			// composit validation
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Validator = "biii_branch1"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.Position = "relative"
				inp.Style.Desktop.Top = "7rem"
				inp.Style.Desktop.Left = "-6rem"
			}

			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q04"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
				// rad.Validator = "must"
				// rad.ErrMsg = "must_central_biii"

				//
				if idx == 0 {
					labels := []trl.S{
						{"de": "seit >10 Jahren "},
						{"de": "seit >5 Jahren"},
						{"de": "seit >3 Jahren"},
						{"de": "seit <=3 Jahren"},
						{"de": "erst seit kurzem"},
					}
					radioValues := []string{
						"10yrs",
						"5yrs",
						"3yrs",
						"lessthan3",
						"recently",
					}
					// gr := page.AddGroup()
					// gr.Cols = 1
					for idx, label := range labels {
						rad := gr.AddInput()
						rad.Type = "radio"
						rad.Name = "q04a"
						rad.ValueRadio = radioValues[idx]

						rad.ColSpan = gr.Cols
						rad.ColSpanLabel = 1
						rad.ColSpanControl = 6

						rad.Label = label

						rad.ControlFirst()

						rad.Style = css.NewStylesResponsive(rad.Style)
						rad.Style.Desktop.StyleBox.Margin = "0 0 0 3.2rem"

					}
				} // idx==0

			} // range labels

		}
	}

	//
	//
	// branch "now"
	//
	// page 2
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "Basisparameter"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir arbeiten nur mit Impact Investments"},
				{"de": "Impact Investing ist Teil unserer SRI / ESG Aktivitäten"},
				{"de": "Impact Investing ist ein unabhängiger Bereich neben unseren konventionellen und SRI / ESG Aktivitäten"},
			}
			radioValues := []string{
				"all",
				"partly",
				"separate_dept",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>5.</b> &nbsp;	Welchen Platz haben Impact Investments in Ihrer Organisation?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q05"
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
		// gr1
		{
			labels := []trl.S{
				{"de": "Impact Investments am 31/12/2021"},
				{"de": "Andere Investments mit ESG Bezug am 31/12/2021"},
				{"de": "Konventionelle Investments am 31/12/2021"},
			}
			names := []string{
				"impact",
				"other",
				"conventional",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>6.</b> &nbsp;	Welches Investitionsvolumen (Assets under Management, Kreditsumme, investiertes Kapital) verzeichnet Ihre Organisation <i>insgesamt</i>?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q06_%v", names[idx])
				inp.Label = label

				inp.ColSpan = 1
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 3
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0 0 0 2.5rem"

				if idx == 0 {
					labelsSub := []trl.S{
						{"de": "Direkte Investments"},
						{"de": "Indirekte Investments bzw. über Intermediäre (Fonds, Asset Manager,…)"},
					}
					namesSub := []string{
						"direkt",
						"indirekt",
					}
					for idx2, labelSub := range labelsSub {

						inp := gr.AddInput()
						inp.Type = "number"
						inp.Name = fmt.Sprintf("q06a_%v", namesSub[idx2])
						inp.Label = labelSub

						inp.ColSpan = gr.Cols
						inp.ColSpanLabel = 18
						inp.ColSpanControl = 11
						inp.Min = 0
						inp.Max = 100
						inp.Step = 0.1
						inp.MaxChars = 5
						inp.Suffix = trl.S{"de": "% Anteil"}
						inp.Placeholder = trl.S{"de": "00"}

						inp.ControlTop()
						inp.Style = css.NewStylesResponsive(inp.Style)
						inp.Style.Desktop.StyleBox.Margin = "0 0 0 6.4rem"

						inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
						inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

					}
				}

			}
		}

	}

	// page 3
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - 2"}
		page.Label = trl.S{"de": "Ihre Impact Investing Ansätze"}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `Im folgenden Teil wollen wir einen Überblick über Ihr Verständnis und Ihre Umsetzung von Impact Investing erfassen.`}
				inp.ColSpan = gr.Cols
			}
		}

		// gr1
		{
			labels := []trl.S{
				{"de": "Werte und ethische Überzeugungen"},
				{"de": "Kundennachfrage"},
				{"de": "Wirtschaftliche Motive (Impact Investing ist ein wichtiges neues Geschäftsfeld)"},
				{"de": "Minimierung und Management von Risiken"},
				{"de": "Lösung drängender gesellschaftlicher und/oder ökologischer Probleme"},
				{"de": "Minimierung des sozialen und/und ökologischen Schadens unserer Investments"},
				{"de": "Andere, bitte nennen&nbsp;&nbsp;"},
			}
			subName := []string{
				"ethics",
				"demand",
				"business_growth",
				"risk_reduction",
				"ecology",
				"damage_control",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					<b>7.</b> &nbsp;
					Was sind die Beweggründe dafür, dass Sie im Impact Investing tätig sind/ wurden?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q07_%v", subName[idx])

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
					inp.Name = "q07_other_label"
					inp.MaxChars = 20
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}
		}
		// 	for idx, label := range labels {
		// 		rad := gr.AddInput()
		// 		rad.Type = "checkbox"
		// 		rad.Name = fmt.Sprintf("q07_%v", subName[idx])

		// 		rad.ColSpan = 1
		// 		rad.ColSpanLabel = 1
		// 		rad.ColSpanControl = 6

		// 		rad.Label = label

		// 		rad.ControlFirst()
		// 	}
		// }

		// gr2
		{
			labels := []trl.S{
				{"de": "Über marktübliche risikoadjustierte Renditen"},
				{"de": "Marktübliche risikoadjustierte Renditen"},
				{"de": "Unter marktübliche risikoadjustierte Renditen"},
				{"de": "Negative Renditen"},
			}
			radioValues := []string{
				"over_market_avg",
				"equal_market_avg",
				"below_market_avg",
				"negative",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>8.</b> &nbsp;	Welche finanziellen Ziele verfolgen Sie mit <u>Ihren Impact Investments</u>?"}
				inp.ColSpan = gr.Cols
			}
			for idx, labl := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q08"
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
				{"de": "Signalling durch Berichterstattung über Impact Investments"},
				{"de": "Technische Unterstützung, Beratung, Vernetzung etc."},
				{"de": "Aktive Mitwirkung durch einen Sitz im Aufsichtsrat"},
				{"de": "Stimmrechtsausübung oder Proxy Voting"},
				{"de": "Aktiver Dialog mit Unternehmen"},
				{"de": "Bereitstellung von Kapital zu günstigen Konditionen (concessionary capital)"},
				{"de": "Unterstützung zur Entwicklung neuer Märkte"},
				{"de": "Weitere, bitte nennen&nbsp;&nbsp;"},
			}
			subName := []string{
				"reporting",
				"tech_support",
				"board_member",
				"proxy_voting",
				"dialogue",
				"capital_provision",
				"market_development",
				"other",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					<b>9.</b> &nbsp;
					Welche Einflussmöglichkeiten nutzen Sie als Impact Investor?
					<br>
					(Mehrfachauswahl möglich)
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q09_%v", subName[idx])

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
					inp.Name = "q09_other_label"
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

	// page 4
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "II Now - 3"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIINow"
		page.SuppressInProgressbar = true

		page.WidthMax("42rem")

		page.ValidationFuncName = "biiiPage5"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		//
		//
		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 35
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `

					<br>

					<p style='text-align: justify;'>
					<b>10.</b> &nbsp;

					Welche Anlagestile verfolgen Sie mit <u>Ihren Impact Investments</u>?

					<br>
					&nbsp;
					<br>

					Die Grundlage für die folgende Frage bildet die GIIN (2017) Definition:
					"Impact Investments sind Investitionen, die mit der Absicht getätigt werden,
					neben einer finanziellen Rendite auch eine positive,
					messbare soziale und ökologische Wirkung zu erzielen" (GIIN, 2017).

					Tragen Sie bitte ausgehend von dieser breiten Definition
					die entsprechenden Investitionsvolumina ein.

					<b>
						Wichtig dabei ist, dass jedes Volumen <i>nur einmalig eingetragen</i> wird, 
						sodass die Summe der entsprechenden Investitionsvolumina (A, B, C, D, E) 
						den Gesamtbetrag Ihrer Impact Investments darstellt. 
						
						Der „% Anteil“ bezieht sich auf das zugehörige Einzelvolumen.
					</b>

					</p>
					<br>
				`}
				inp.ColSpan = gr.Cols
			}

			// 10a
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10a")
				inp.Label = trl.S{"de": `
					<b>A)</b> &nbsp; Kapitalerhöhungen / -zuführungen (z.B. IPO, PE Investment, Kredite etc.),
					die zur Generierung eines zusätzlichen Impacts führen
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.ControlTop()
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10a_pct")
				inp.Label = trl.S{"de": `
					Der hierdurch erzeugte, realweltliche Impact
					(sozial und/oder ökologisch) wird gemessen und dokumentiert (Outcome Ebene)
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.ControlTop()
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "-0.5rem 0 1.5rem 2.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}

			// 10b
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10b")
				inp.Label = trl.S{"de": `
					<b>B)</b> &nbsp;
					Andere (über A hinausgehende) Investitionen in Investees mit klaren Impact Zielen
					 (z.B. Unternehmen bei denen Impact den Kern des Geschäftsmodells
						bildet oder Unternehmen mit Science-based Targets)
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.ControlTop()
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10b_pct")
				inp.Label = trl.S{"de": `
					Der hierdurch erzeugte, realweltliche Impact (sozial und/oder ökologisch)
					wird gemessen und dokumentiert (Outcome Ebene)
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.ControlTop()
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "-0.5rem 0 1.5rem 2.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>C)</b> &nbsp;
					Alle weiteren (über A und B hinaus gehenden) Investitionen und Aktien
				`}
				inp.ColSpan = gr.Cols
			}

			// 10c1, 10c2
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10c1")
				inp.Label = trl.S{"de": `
					<!-- <b>C1)</b> &nbsp;  -->
					Für die eine strategische Engagement Policy existiert (z.B. aktiver Dialog mit den Investees)
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.ControlTop()
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10c1_pct")
				inp.Label = trl.S{"de": `
					Investitionen für die der transformative Impact der Engagement Policy dokumentiert wird
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.ControlTop()
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "-0.5rem 0 0.2rem 2.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10c2")
				inp.Label = trl.S{"de": `
					<!-- <b>C2)</b> &nbsp;  -->
					Für die eine strategische Voting Policy (z.B. Ausübung von Stimmrechten) existiert
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.ControlTop()
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10c2_pct")
				inp.Label = trl.S{"de": `
					Aktien für die der transformative Impact der Voting Policy dokumentiert wird
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.ControlTop()
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "-0.5rem 0 1.5rem 2.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}

			// 10d
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10d")
				inp.Label = trl.S{"de": `
					<b>D)</b> &nbsp;
					Alle weiteren (über A, B und C hinaus gehenden) ESG gemanagten Investitionen (z.B. Ausschlüsse / Best-in-class Ansätze / ESG Integration etc)
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.ControlTop()
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10d_pct")
				inp.Label = trl.S{"de": `
					Investitionen für die die Impact Performance 
					mit einem Benchmark-Vergleich oder anhand von SDG Beiträgen dargestellt 
					wird (Output Ebene)
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.ControlTop()
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "-0.5rem 0 1.5rem 2.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}

			// 10e
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q10e_other"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": `
					<b>E)</b> &nbsp;
					Weitere (Bitte nennen)
				`}
				inp.ColSpan = 5*5 - 1
				inp.ColSpanLabel = 1.9
				inp.ColSpanControl = 2.1

				inp.ControlTop()
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0 "

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10e")

				inp.ColSpan = 2 * 5
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 2
				inp.Min = 0
				inp.Max = math.MaxFloat64
				inp.Step = 1000
				inp.Step = 0.01
				inp.MaxChars = 9
				inp.Suffix = trl.S{"de": "Mio €"}
				inp.Placeholder = trl.S{"de": "0,00"}

				inp.ControlTop()
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleBox.Padding = "0 0 0 0.9rem"

			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("q10e_pct")
				inp.Label = trl.S{"de": `
					Impact relevante Informationen werden gemessen und dokumentiert
				`}

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 16
				inp.ColSpanControl = 7
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1
				inp.MaxChars = 5
				inp.Suffix = trl.S{"de": "% Anteil"}
				inp.Placeholder = trl.S{"de": "00"}

				inp.ControlTop()
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "-0.5rem 0 1.5rem 2.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.8rem 0 0"

			}

		}

	}

	page4Quest11(&q)

	later(&q)

	finish(&q)

	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}
	return &q, nil

}
