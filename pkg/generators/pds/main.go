package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create questionnaire
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"en"} // governs default language code

	q.Survey.Org = trl.S{
		"en": "ZEW",
		"de": "ZEW",
	}
	q.Survey.Name = trl.S{
		"en": "Private Debt Survey",
		"de": "Private Debt Survey",
	}
	// q.Variations = 1

	// page0
	{
		page := q.AddPage()
		page.SuppressInProgressbar = true

		page.Label = trl.S{
			"en": "Greeting",
			"de": "Begrüßung",
		}
		page.Short = trl.S{
			"en": "Greeting",
			"de": "Begrüßung",
		}

		page.ValidationFuncName = ""
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./welcome-page.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page1
	{
		page := q.AddPage()
		// page.SuppressInProgressbar = true
		page.ValidationFuncName = "pdsPage1"

		page.Label = trl.S{
			"en": "Identification and asset classes",
			"de": "Identification and asset classes",
		}
		page.Short = trl.S{
			"en": "Asset classes,<br>tranches",
			"de": "Asset classes,<br>tranches",
		}
		page.CounterProgress = "-"

		page.WidthMax("42rem")

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "identification"
				inp.MaxChars = 24
				inp.Placeholder = trl.S{
					"en": "name of manager",
					"de": "Name Manager",
				}
				inp.Label = trl.S{
					"en": "Identification",
					"de": "Identifikation",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
			}
		}

		// gr2
		{
			lblMain := trl.S{
				"de": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
				"en": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
			}
			radiosSingleRow(
				qst.WrapPageT(page),
				"teamsize",
				lblMain,
				mCh2,
			)
		}

		if false {
			// gr3
			{
				lblMain := trl.S{
					"en": `Which asset classes do you invest in?`,
					"de": `Wählen Sie Ihre Assetklassen.`,
				}
				checkBoxRow(
					qst.WrapPageT(page),
					lblMain,
					assetClassesInputs,
					assetClassesLabels,
				)
			}

			// gr4
			{
				lblMain := trl.S{
					"en": `Your  strategies`,
					"de": `Ihre Strategie`,
				}
				checkBoxRow(
					qst.WrapPageT(page),
					lblMain,
					trancheTypeNamesAC1,
					allLbls["ac1-tranche-types"],
				)
			}
		}

		//
		// gr5
		{
			lblMain := trl.S{
				"en": `Which asset classes do you invest in?`,
				"de": `Wählen Sie Ihre Assetklassen.`,
			}
			checkBoxCascade(
				qst.WrapPageT(page),
				lblMain,
				assetClassesInputs,
				assetClassesLabels,
			)

		}

	}

	// page10
	{
		page := q.AddPage()
		// page.Section = trl.S{
		// 	"en": "Section 1",
		// 	"de": "Section 1",
		// }
		page.Label = trl.S{
			"en": "1. Portfolio changes (past 3 months)",
			"de": "1. Portfolio changes (past 3 months)",
		}
		page.Short = trl.S{
			"en": "Portfolio changes",
			"de": "Portfolio changes",
		}
		page.CounterProgress = "1"

		page.WidthMax("42rem")
		page.WidthMax("64rem")

		restrictedTextMultiCols(qst.WrapPageT(page), rT1)

		lblDuration := trl.S{
			"en": "Average time to close a deal in weeks",
			"de": "Durchschnittl. Zeit bis Abschluss in Wochen",
		}.Outline("1.2")
		slidersPctRowLabelsLeft(
			qst.WrapPageT(page),
			"closing_time",
			lblDuration,
			suffixWeeks,
			"3",
		)

		chapter3(
			qst.WrapPageT(page),
			"closing_time2",
			trl.S{
				"en": "Alternative visualisation using radios",
				"de": "Alternative visualisation using radios",
			}.Outline("1.2"),
			mCh5,
		)

		restrictedTextMultiCols(qst.WrapPageT(page), rT2)

		restrictedTextMultiCols(qst.WrapPageT(page), rT3)

		restrictedTextMultiCols(qst.WrapPageT(page), rT4)

		shareESG := trl.S{
			"en": `<bb>Share ESG KPIs</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
			"de": `<bb>Share ESG KPIs</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
		}.Outline("1.6")
		slidersPctRowLabelsLeft(
			qst.WrapPageT(page),
			"esg",
			shareESG,
			suffixPercent,
			"2",
		)

		shareESGRatch := trl.S{
			"en": `<bb>Share ESG ratchets</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
			"de": `<bb>Share ESG ratchets</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
		}.Outline("1.7")
		slidersPctRowLabelsLeft(
			qst.WrapPageT(page),
			"esgratch",
			shareESGRatch,
			suffixPercent,
			"2",
		)

		share15Degree := trl.S{
			"en": `<bb>Share 1.5°C target</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target? </span>`,
			"de": `<bb>Share 1.5°C target</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target? </span>`,
		}.Outline("1.8")
		slidersPctRowLabelsLeft(
			qst.WrapPageT(page),
			"esg15degrees",
			share15Degree,
			suffixPercent,
			"2",
		)

	}

	// page20
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "2. Overall (existing) Portfolio",
			"de": "2. Overall (existing) Portfolio",
		}
		page.Short = trl.S{
			"en": "Portfolio base",
			"de": "Portfolio base",
		}
		page.CounterProgress = "2"

		page.WidthMax("64rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": "Assets under Management",
					"de": "Assets under Management",
				}.Outline("2.1")
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

		page21Types := []string{
			"restricted-text-million",
			"restricted-text-million",
			"restricted-text-million",
			"restricted-text-million",
		}
		page21Inputs := []string{
			"portfolio_value",
			"capital_called",
			"capital_repaid",
			"capital_reserve",
		}
		page21Lbls := []trl.S{
			{
				"en": "Fair market value of current portfolio in mn €",
				"de": "Fair market value of current portfolio in mn €",
			},
			{
				"en": "Capital called from investor in mn €",
				"de": "Capital called from investor in mn €",
			},
			{
				"en": "Repaid capital either reinvested or distributed to investor in mn €",
				"de": "Repaid capital either reinvested or distributed to investor in mn €",
			},
			{
				"en": "Dry powder in mn €",
				"de": "Dry powder in mn €",
			},
		}

		for i := 0; i < len(page21Lbls); i++ {
			rn := rune(97 + i)
			page21Lbls[i] = page21Lbls[i].Outline(fmt.Sprintf("&nbsp;&nbsp;%c.)", rn))
		}

		for idx1, inpName := range page21Inputs {

			if page21Types[idx1] == "restricted-text-million" {
				restrTextRowLabelsTop(
					qst.WrapPageT(page),
					inpName,
					page21Lbls[idx1],
					rTSingleRowMill,
				)
			}

			if page21Types[idx1] == "range-pct" {
				slidersPctRowLabelsTop(
					qst.WrapPageT(page),
					inpName,
					page21Lbls[idx1],
					suffixPercent,
				)
			}

			if page21Types[idx1] == "radios1-4" {
				chapter3(
					qst.WrapPageT(page),
					inpName,
					page21Lbls[idx1],
					mCh2a,
				)
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": "Portfolio Composition",
					"de": "Portfolio Composition",
				}.Outline("2.2")
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

		restrictedTextMultiCols(qst.WrapPageT(page), r221)
		restrictedTextMultiCols(qst.WrapPageT(page), r222)

	}

	// page23
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "2.3 Portfolio risk",
			"de": "2.3 Portfolio risk",
		}
		page.Short = trl.S{
			"en": "Portfolio risk",
			"de": "Portfolio risk",
		}
		page.CounterProgress = "2.3"

		page.WidthMax("64rem")

		page23Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",
			"range-pct",
			"restricted-text",
			"range-pct",

			"range-pct",
			"range-pct",
			"range-pct",
		}
		page23Inputs := []string{
			"covenants_per_credit",
			"share_covenant_holiday",
			"share_covenant_reset",
			"share_covenant_breach",
			"share_loan_defaults",
			"share_default_recovered",

			"share_esg_kpis",
			"share_esg_ratchets",
			"share_esg_15degrees",
		}
		page23Lbls := []trl.S{
			{
				"en": "Average Number of Financial Covenants per credit",
				"de": "Average Number of Financial Covenants per credit",
			},
			{
				"en": "What is the share of portfolio (at fair market value) with a covenant holiday?",
				"de": "What is the share of portfolio (at fair market value) with a covenant holiday?",
			},
			{
				"en": "What is the share of portfolio (at fair market value) with a covenant reset?",
				"de": "What is the share of portfolio (at fair market value) with a covenant reset?",
			},
			{
				"en": "What is the share of portfolio (at fair market value) with a covenant breach?",
				"de": "What is the share of portfolio (at fair market value) with a covenant breach?",
			},
			{
				"en": "Share of defaulted loans (measured at cost/principal amount of loan)",
				"de": "Share of defaulted loans (measured at cost/principal amount of loan)",
			},
			{
				"en": "If you had a default in the past. What was the recovery rate (share of principal)?",
				"de": "If you had a default in the past. What was the recovery rate (share of principal)?",
			},
			// esg
			{
				"en": "What is the share of portfolio (at fair market value) with explicit ESG targets in the credit documentation?",
				"de": "What is the share of portfolio (at fair market value) with explicit ESG targets in the credit documentation?",
			},
			{
				"en": "What is the share of portfolio (at fair market value) with ESG ratchets?",
				"de": "What is the share of portfolio (at fair market value) with ESG ratchets?",
			},
			{
				"en": "What is the share of portfolio (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?",
				"de": "What is the share of portfolio (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?",
			},
		}

		for i := 0; i < len(page23Lbls); i++ {
			// page23Lbls[i] = page23Lbls[i].Outline(fmt.Sprintf("2.3.%v", i+1))
			rn := rune(97 + i)
			page23Lbls[i] = page23Lbls[i].Outline(fmt.Sprintf("&nbsp;&nbsp;%c.)", rn))
		}

		for idx1, inpName := range page23Inputs {

			if page23Types[idx1] == "restricted-text" {
				restrTextRowLabelsTop(
					qst.WrapPageT(page),
					inpName,
					page23Lbls[idx1],
					rTSingleRowPercent,
				)
			}

			if page23Types[idx1] == "range-pct" {
				slidersPctRowLabelsTop(
					qst.WrapPageT(page),
					inpName,
					page23Lbls[idx1],
					suffixPercent,
				)
			}

			if page23Types[idx1] == "radios1-4" {
				chapter3(
					qst.WrapPageT(page),
					inpName,
					page23Lbls[idx1],
					mCh2a,
				)
			}

		}

	}

	// page3
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "3. Index Questions",
			"de": "3. Index Questions",
		}
		page.Short = trl.S{
			"en": "Indizes",
			"de": "Indizes",
		}
		page.CounterProgress = "3"

		page.WidthMax("64rem")

		page5Inputs := []string{
			"financing_situation_pricing",
			"deal_quality",
			"deal_documentation",
			"deal_amount",
		}
		page5Lbls := []trl.S{
			{
				"en": "Financing situation/pricing",
				"de": "Financing situation/pricing",
			},
			{
				"en": "Assess the change in deal quality with respect to the risk return profile",
				"de": "Assess the change in deal quality with respect to the risk return profile",
			},
			{
				"en": "Assess the quality of deal documentation (covenant strength, enforcement rights, etc.)",
				"de": "Assess the quality of deal documentation (covenant strength, enforcement rights, etc.)",
			},
			{
				"en": "Do you observe more deals, same amount of deals or less deals",
				"de": "Do you observe more deals, same amount of deals or less deals",
			},
		}

		for i := 0; i < len(page5Lbls); i++ {
			page5Lbls[i] = page5Lbls[i].Outline(fmt.Sprintf("3.%v", i+1))
		}

		for idx1, inpName := range page5Inputs {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 0
				gr.BottomVSpacers = 1
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = page5Lbls[idx1].Bold()
					inp.Label = page5Lbls[idx1]

					inp.ColSpan = 1
					// inp.ColSpanLabel = 1
				}
			}

			chapter3(
				qst.WrapPageT(page),
				inpName+"_past3m",
				trl.S{
					"en": "<i>Last</i> 3&nbsp;months",
					"de": "<i>Last</i> 3&nbsp;months",
				}.Outline("a.)"),
				mCh4,
			)
			chapter3(
				qst.WrapPageT(page),
				inpName+"_next3m",
				trl.S{
					"en": "<i>Next</i> 3&nbsp;months",
					"de": "<i>Next</i> 3&nbsp;months",
				}.Outline("b.)"),
				mCh4,
			)

		}

	}

	// page4
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "4. Qualitative questions",
			"de": "4. Qualitative questions",
		}
		page.Short = trl.S{
			"en": "Quality",
			"de": "Quality",
		}
		page.CounterProgress = "4"

		page.WidthMax("42rem")

		// matrix1
		{
			inps := []string{
				"business_cycle",
				"interest_rates",
				"inflation_deflation",
				"demographics",
				"supply_chains",
				"health_issues",
				"regulatory_environment",
				"other",
			}

			lbls := map[string]string{
				"business_cycle":         "Business cycle",
				"interest_rates":         "Interest rates",
				"inflation_deflation":    "Inflation/deflation",
				"demographics":           "Demographics",
				"supply_chains":          "Supply chains",
				"health_issues":          "Health issues",
				"regulatory_environment": "Regulatory environment",
				"other":                  "Other",
			}

			lblMain := trl.S{
				"en": `What do you think are the main risks for your investment strategy over the next 3 months? Please choose three.`,
				"de": `What do you think are the main risks for your investment strategy over the next 3 months? Please choose three.`,
			}.Outline("4.1")
			prio3Matrix(qst.WrapPageT(page), "risks", lblMain, inps, lbls, true)
		}

		//
		// matrix2
		{
			inps := []string{
				"energy",
				"materials",
				"industrials",
				"consumer_discretionary",
				"consumer_staples",
				"health_care",
				"financials",
				"information_technology",
				"communication_services",
				"utilities",
				"real_estate",
			}

			lbls := map[string]string{
				"energy":                 "Energy",
				"materials":              "Materials",
				"industrials":            "Industrials",
				"consumer_discretionary": "Consumer Discretionary",
				"consumer_staples":       "Consumer Staples",
				"health_care":            "Health Care",
				"financials":             "Financials",
				"information_technology": "Information Technology",
				"communication_services": "Communication Services",
				"utilities":              "Utilities",
				"real_estate":            "Real Estate",
			}

			lblMain := trl.S{
				"en": `What GICS sectors provides the most attractive 
					investment opportunities in the next three months? 
					Please rank the top three.`,
				"de": `What GICS sectors provides the most attractive 
					investment opportunities in the next three months? 
					Please rank the top three.`,
			}.Outline("4.2")
			prio3Matrix(qst.WrapPageT(page), "gicsprio", lblMain, inps, lbls, false)
		}

		{
			esgImportance1 := trl.S{
				"en": `How important are ESG considerations 
					to core principal in your investment process?
			`,
				"de": `How important are ESG considerations 
					to core principal in your investment process?
			`,
			}.Outline("4.3")

			// rejected at meeting 2022-11-14
			// rangePercentage(qst.WrapPageT(page), "esg", esgImportance1, "importance2")

			radiosSingleRow(
				qst.WrapPageT(page),
				"esg_importance1",
				esgImportance1,
				mCh3,
			)
		}

		//
		// matrix3
		{
			inps := []string{
				"availability",
				"quality",
				"performance",
				"greenwashing",
				"regulation",
				"opportunities",
				"other",
			}

			lbls := map[string]string{
				"availability":  "ESG data availability",
				"quality":       "ESG data quality",
				"performance":   "Concerns about performance/sacrificing returns",
				"greenwashing":  "Concerns about greenwashing",
				"regulation":    "Complex regulatory landscape",
				"opportunities": "Lack of suitable investments",
				"other":         "Other",
			}

			lblMain := trl.S{
				"en": `What is the biggest challenge related to the implementation of ESG into your investment strategy?`,
				"de": `What is the biggest challenge related to the implementation of ESG into your investment strategy?`,
			}.Outline("4.4")
			prio3Matrix(qst.WrapPageT(page), "esg_challenge", lblMain, inps, lbls, true)
		}

		{

			var inps = []string{
				"poverty",
				"hunger",
				"health",
				"education",
				"gender_eq",
				"water",
				"energy",
				"work",
				"industry",
				"inequality",
				"communities",
				"responsible",
				"climate",
				"life_water",
				"life_land",
				"peace",
				"partnership",
			}
			var lbls = []trl.S{
				{
					"de": "No poverty",
					"en": "No poverty",
				},
				{
					"de": "Zero hunger",
					"en": "Zero hunger",
				},
				{
					"de": "Good health and well-being",
					"en": "Good health and well-being",
				},
				{
					"de": "Quality education",
					"en": "Quality education",
				},
				{
					"de": "Gender equality",
					"en": "Gender equality",
				},
				{
					"de": "Clean water and sanitation",
					"en": "Clean water and sanitation",
				},
				{
					"de": "Affordable and clean energy",
					"en": "Affordable and clean energy",
				},
				{
					"de": "Decent work and economic growth",
					"en": "Decent work and economic growth",
				},
				{
					"de": "Industry innovation and infrastructure",
					"en": "Industry innovation and infrastructure",
				},
				{
					"de": "Reduce inequality",
					"en": "Reduce inequality",
				},
				{
					"de": "Sustainable cities and communities",
					"en": "Sustainable cities and communities",
				},
				{
					"de": "Responsible consumption and production",
					"en": "Responsible consumption and production",
				},
				{
					"de": "Climate action",
					"en": "Climate action",
				},
				{
					"de": "Life below water",
					"en": "Life below water",
				},
				{
					"de": "Life on land",
					"en": "Life on land",
				},
				{
					"de": "Peace, justice and strong institutions",
					"en": "Peace, justice and strong institutions",
				},
				{
					"de": "Partnership for the goals",
					"en": "Partnership for the goals",
				},
			}

			unSDG := trl.S{
				"en": `What UN SDGs are supported by your investment strategy?`,
				"de": `What UN SDGs are supported by your investment strategy?`,
			}.Outline("4.5")
			checkBoxColumn(qst.WrapPageT(page), unSDG, 2, inps, lbls)
		}

	}

	// page-slider-demo
	{
		page := q.AddPage()

		page.Label = trl.S{"en": "Slider"}
		page.Short = trl.S{"en": "Slider"}
		page.CounterProgress = "-"
		page.WidthMax("42rem")

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 11
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Name = "range01"
				inp.Type = "range"
				inp.DynamicFuncParamset = "1"

				inp.Min = 0
				inp.Max = 100
				inp.Step = 10
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.Width = "90%"

				// inp.Label = trl.S{
				// 	"de": "Normal Slider",
				// 	"en": "Normal Slider",
				// }
				inp.Suffix = trl.S{
					"de": "unit",
					"en": "unit",
				}

				inp.ColSpan = 4
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 8
			}
			/* 			{
			   				inp := gr.AddInput()
			   				inp.Name = "range01_display"
			   				inp.Type = "text"
			   				inp.MaxChars = 8
			   				inp.ColSpan = 2
			   				inp.ColSpanLabel = 0
			   				inp.ColSpanControl = 1

			   				inp.Style = css.NewStylesResponsive(inp.Style)
			   				inp.Style.Desktop.StyleBox.Position = "relative"
			   				inp.Style.Desktop.StyleBox.Top = "0.58rem"
			   				inp.Style.Desktop.StyleBox.Left = "0.58rem"
			   			}
			   			{
			   				inp := gr.AddInput()
			   				inp.Name = "range01_noanswer"
			   				inp.Type = "radio"
			   				inp.ColSpan = 2
			   				inp.Label = trl.S{
			   					"de": "nicht verfügb.",
			   					"en": "no answer",
			   				}
			   				inp.ValueRadio = "xx"
			   				inp.ColSpanLabel = 1
			   				inp.ColSpanControl = 1

			   				// inp.ControlTop()
			   				// inp.ControlBottom()

			   				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
			   				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)

			   				inp.StyleCtl.Desktop.StyleGridItem.Col = "auto/1"
			   				inp.StyleLbl.Desktop.StyleGridItem.Col = "auto/1"
			   			}
			*/
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "range"

				inp.JSBlockStrings = map[string]string{}
				inp.JSBlockStrings["inputName"] = "range01" // as above
			}

		}

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.WidthMax("85%")
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./slider/inner-1.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.WidthMax("85%")
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./slider/inner-2.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	//
	//
	if false {
		gn := "A"
		page := q.AddPage()
		page.GeneratorFuncName = "pds01"
		page.Label = trl.S{
			"en": fmt.Sprintf("dyn page label %v", gn),
			"de": fmt.Sprintf("dyn page label %v", gn),
		}
		page.Short = trl.S{
			"en": fmt.Sprintf("dyn %v", gn),
			"de": fmt.Sprintf("dyn %v", gn),
		}
		page.WidthMax("42rem")
		page.NoNavigation = false
	}
	if false {
		gn := "B"
		page := q.AddPage()
		page.GeneratorFuncName = "pds01"
		page.Label = trl.S{
			"en": fmt.Sprintf("dyn page label %v", gn),
			"de": fmt.Sprintf("dyn page label %v", gn),
		}
		page.Short = trl.S{
			"en": fmt.Sprintf("dyn %v", gn),
			"de": fmt.Sprintf("dyn %v", gn),
		}
		page.WidthMax("42rem")
		page.NoNavigation = false
	}

	//
	//
	// Report of results
	{
		p := q.AddPage()
		p.NoNavigation = true
		p.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}
		{
			// gr := p.AddGroup()
			// gr.Cols = 1
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.DynamicFunc = "RepsonseStatistics"
			// }
		}
	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
