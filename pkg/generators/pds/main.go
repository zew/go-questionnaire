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

	// page23
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "2.3 Portfolio Risk",
			"de": "2.3 Portfolio Risk",
		}
		page.Short = trl.S{
			"en": "Portfolio Risk",
			"de": "Portfolio Risk",
		}

		page.WidthMax("72rem")

		page23Inputs := []string{
			"xxx1",
			"xxx2",
		}
		page23Lbls := []trl.S{
			{
				"en": "Average Number of Financial Covenants per credit",
				"de": "Average Number of Financial Covenants per credit",
			},
			{
				"en": "What is the share of portfolio (at Fair Market Value) with a covenant holiday?",
				"de": "What is the share of portfolio (at Fair Market Value) with a covenant holiday?",
			},
			{
				"en": "What is the share of portfolio (at Fair Market Value) with a covenant reset?",
				"de": "What is the share of portfolio (at Fair Market Value) with a covenant reset?",
			},
			{
				"en": "What is the share of portfolio (at Fair Market Value) with a covenant breach?",
				"de": "What is the share of portfolio (at Fair Market Value) with a covenant breach?",
			},
			{
				"en": "Share of defaulted loans (measured at cost/principal amount of loan)",
				"de": "Share of defaulted loans (measured at cost/principal amount of loan)",
			},
			{
				"en": "If you had a default in the past. What was the recovery rate (share of principal)?",
				"de": "If you had a default in the past. What was the recovery rate (share of principal)?",
			},
			{
				"en": "What is the share of portfolio (at Fair Market Value) with explicit ESG targets in the credit documentation?",
				"de": "What is the share of portfolio (at Fair Market Value) with explicit ESG targets in the credit documentation?",
			},
			{
				"en": "What is the share of portfolio (at Fair Market Value) with ESG ratchets?",
				"de": "What is the share of portfolio (at Fair Market Value) with ESG ratchets?",
			},
			{
				"en": "What is the share of portfolio (at Fair Market Value) where the creditor explicitly states a strategy to add to the 1.5°C target?",
				"de": "What is the share of portfolio (at Fair Market Value) where the creditor explicitly states a strategy to add to the 1.5°C target?",
			},
		}

		for idx1, inpName := range page23Inputs {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = page23Lbls[idx1].Bold()

					inp.ColSpan = 1
					// inp.ColSpanLabel = 1
				}
			}

			chapter3(
				qst.WrapPageT(page),
				inpName,
				"past3m",
				trl.S{
					"en": "Last 3&nbsp;months",
					"de": "Last 3&nbsp;months",
				},
				mCh4,
			)
			chapter3(
				qst.WrapPageT(page),
				inpName,
				"next3m",
				trl.S{
					"en": "Next 3&nbsp;months",
					"de": "Next 3&nbsp;months",
				},
				mCh4,
			)

		}

	}

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
			"en": "Manager Characteristics",
			"de": "Manager Characteristics",
		}
		page.Short = trl.S{
			"en": "Manager",
			"de": "Manager",
		}

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

			mode2mod := mCh2
			mode2mod.GroupBottomSpacers = 3
			radiosSingleRow(qst.WrapPageT(page), "teamsize", lblMain, mode2mod)
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
					// chb01,
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
					// chb02,
				)
			}
		}

		//
		// gr5
		{
			lblMain := trl.S{
				"en": `
					<!-- Suggestion; instead of three different surveys; -->
					rows 2,3 not shown in first wave<br>
					Which asset classes do you invest in?`,
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

	// page2
	{
		page := q.AddPage()
		// page.Section = trl.S{
		// 	"en": "Section 1",
		// 	"de": "Section 1",
		// }
		page.Label = trl.S{
			"en": "1. Portfolio Changes (past 3 months)",
			"de": "1. Portfolio Changes (past 3 months)",
		}
		page.Short = trl.S{
			"en": "Portfolio Changes",
			"de": "Portfolio Changes",
		}

		page.WidthMax("42rem")
		page.WidthMax("66rem")

		restrictedTextMultiCols(qst.WrapPageT(page), rT1)

		lblDuration := trl.S{
			"en": "Average time to close a deal in weeks",
			"de": "Durchschnittl. Zeit bis Abschluss in Wochen",
		}
		slidersRow(qst.WrapPageT(page), "closing_time", lblDuration, suffixWeeks)
		if false {
			// old: a single range
			rangeClosingTime(qst.WrapPageT(page), trancheTypeNamesAC1[0], lblDuration)
		}

		restrictedTextMultiCols(qst.WrapPageT(page), rT2)

		restrictedTextMultiCols(qst.WrapPageT(page), rT3)

		restrictedTextMultiCols(qst.WrapPageT(page), rT4)

		shareESG := trl.S{
			"en": `<b>Share ESG KPIs</b> <br>What is the share of new deals (at Fair Market Value) with explicit ESG targets in the credit documentation?`,
			"de": `<b>Share ESG KPIs</b> <br>What is the share of new deals (at Fair Market Value) with explicit ESG targets in the credit documentation?`,
		}
		slidersRow(qst.WrapPageT(page), "esg", shareESG, suffixPercent)

		shareESGRatch := trl.S{
			"en": `<b>Share ESG ratchets</b> <br> What is the share of new deals (at Fair Market Value) with ESG ratchets?`,
			"de": `<b>Share ESG ratchets</b> <br> What is the share of new deals (at Fair Market Value) with ESG ratchets?`,
		}
		slidersRow(qst.WrapPageT(page), "esgratch", shareESGRatch, suffixPercent)

		share15Degree := trl.S{
			"en": `<b>Share 1.5°C target</b> <br> What is the share of new deals (at Fair Market Value) where the creditor explicitly states a strategy to add to the 1.5°C target?`,
			"de": `<b>Share 1.5°C target</b> <br> What is the share of new deals (at Fair Market Value) where the creditor explicitly states a strategy to add to the 1.5°C target?`,
		}
		slidersRow(qst.WrapPageT(page), "esg15degrees", share15Degree, suffixPercent)

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

		page.WidthMax("72rem")

		page5Inputs := []string{
			"financing_situation_pricing",
			"deal_quality",
			"deal_documentation",
			"deal_amount",
		}
		page5Lbls := []trl.S{
			{
				"en": "Financing Situation/Pricing",
				"de": "Financing Situation/Pricing",
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

		for idx1, inpName := range page5Inputs {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = page5Lbls[idx1].Bold()

					inp.ColSpan = 1
					// inp.ColSpanLabel = 1
				}
			}

			chapter3(
				qst.WrapPageT(page),
				inpName,
				"past3m",
				trl.S{
					"en": "Last 3&nbsp;months",
					"de": "Last 3&nbsp;months",
				},
				mCh4,
			)
			chapter3(
				qst.WrapPageT(page),
				inpName,
				"next3m",
				trl.S{
					"en": "Next 3&nbsp;months",
					"de": "Next 3&nbsp;months",
				},
				mCh4,
			)

		}

	}

	// page4
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "4. Qualitative Questions",
			"de": "4. Qualitative Questions",
		}
		page.Short = trl.S{
			"en": "Quality",
			"de": "Quality",
		}

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
				"business_cycle":         "Business Cycle",
				"interest_rates":         "Interest Rates",
				"inflation_deflation":    "Inflation/Deflation",
				"demographics":           "Demographics",
				"supply_chains":          "Supply Chains",
				"health_issues":          "Health Issues",
				"regulatory_environment": "Regulatory Environment",
				"other":                  "Other",
			}

			lblMain := trl.S{
				"en": `What do you think are the main risks for your investment strategy over the next 3 months? Please choose three.`,
				"de": `What do you think are the main risks for your investment strategy over the next 3 months? Please choose three.`,
			}
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
			}
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
			}

			// rejected at meeting 2022-11-14
			// rangePercentage(qst.WrapPageT(page), "esg", esgImportance1, "importance2")

			radiosSingleRow(qst.WrapPageT(page), "esg_importance1", esgImportance1, mCh3)
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
			}
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
					"de": "(1) No Poverty",
					"en": "(1) No Poverty",
				},
				{
					"de": "(2) Zero Hunger",
					"en": "(2) Zero Hunger",
				},
				{
					"de": "(3) Good Health and Well-Being",
					"en": "(3) Good Health and Well-Being",
				},
				{
					"de": "(4) Quality Education",
					"en": "(4) Quality Education",
				},
				{
					"de": "(5) Gender Equality",
					"en": "(5) Gender Equality",
				},
				{
					"de": "(6) Clean Water and Sanitation",
					"en": "(6) Clean Water and Sanitation",
				},
				{
					"de": "(7) Affordable and Clean Energy",
					"en": "(7) Affordable and Clean Energy",
				},
				{
					"de": "(8) Decent Work and Economic Growth",
					"en": "(8) Decent Work and Economic Growth",
				},
				{
					"de": "(9) Industry Innovation and Infrastructure",
					"en": "(9) Industry Innovation and Infrastructure",
				},
				{
					"de": "(10) Reduce Inequality",
					"en": "(10) Reduce Inequality",
				},
				{
					"de": "(11) Sustainable Cities and Communities",
					"en": "(11) Sustainable Cities and Communities",
				},
				{
					"de": "(12) Responsible Consumption and Production",
					"en": "(12) Responsible Consumption and Production",
				},
				{
					"de": "(13) Climate Action",
					"en": "(13) Climate Action",
				},
				{
					"de": "(14) Life below Water",
					"en": "(14) Life below Water",
				},
				{
					"de": "(15) Life on Land",
					"en": "(15) Life on Land",
				},
				{
					"de": "(16) Peace, Justice and strong Institutions",
					"en": "(16) Peace, Justice and strong Institutions",
				},
				{
					"de": "(17) Partnership for the Goals",
					"en": "(17) Partnership for the Goals",
				},
			}

			unSDG := trl.S{
				"en": `What UN SDGs are supported by your investment strategy?`,
				"de": `What UN SDGs are supported by your investment strategy?`,
			}
			checkBoxColumn(qst.WrapPageT(page), unSDG, 2, inps, lbls)
		}

	}

	// page-slider-demo
	{
		page := q.AddPage()

		page.Label = trl.S{"de": "Slider"}
		page.Short = trl.S{"de": "Slider"}
		page.WidthMax("42rem")

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

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 11
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Name = "range01"
				inp.Type = "range"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 5
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.Width = "90%"

				inp.Label = trl.S{
					"de": "Normal Slider",
					"en": "Normal Slider",
				}

				inp.ColSpan = 7
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 8
			}
			{
				inp := gr.AddInput()
				inp.Name = "range01_display"
				inp.Type = "text"
				inp.MaxChars = 8
				inp.ColSpan = 2
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Name = "range01_noanswer"
				inp.Type = "radio"
				inp.ColSpan = 2
				inp.Label = trl.S{
					"de": "nicht verfügb.",
					"en": "not available",
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

			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "range"

				inp.JSBlockStrings = map[string]string{}
				inp.JSBlockStrings["inputName"] = "range01" // as above
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
