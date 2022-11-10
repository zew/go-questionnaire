package pds

import (
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

	// page 0
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

	// page 0a
	{
		page := q.AddPage()
		// page.SuppressInProgressbar = true

		page.Label = trl.S{
			"en": "Manager Characteristics",
			"de": "Manager Characteristics",
		}
		page.Short = trl.S{
			"en": "Manager",
			"de": "Manager",
		}

		page.ValidationFuncName = ""
		page.WidthMax("42rem")

		// gr0
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

		// gr1
		{
			lblMain := trl.S{
				"en": `Strategy -  shouldn't this be checkboxes? See 'asset classes' below`,
				"de": `Strategie - shouldn't this be checkboxes? See 'asset classes' below`,
			}
			multipleChoiceSingleRow(qst.WrapPageT(page), "strategy", lblMain, mCh1)
		}
		// gr2
		{
			lblMain := trl.S{
				"de": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
				"en": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
			}

			mode2mod := mCh2
			mode2mod.GroupBottomSpacers = 3
			multipleChoiceSingleRow(qst.WrapPageT(page), "teamsize", lblMain, mode2mod)
		}

		// gr3
		assetClass(qst.WrapPageT(page))

	} // page0a

	// page1
	{
		page := q.AddPage()
		// page.Section = trl.S{
		// 	"en": "Section 1",
		// 	"de": "Section 1",
		// }
		page.Label = trl.S{
			"en": "1. Portfolio Changes (past 3 Months)",
			"de": "1. Portfolio Changes (past 3 Months)",
		}
		page.Short = trl.S{
			"en": "Portfolio Changes",
			"de": "Portfolio Changes",
		}

		page.WidthMax("42rem")

		{
			gr := page.AddGroup()
			gr.Cols = 2
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 2
				inp.Label = trl.S{
					"en": `Should we group by question type (deals, time, volume) - 
						or as now by strategy (senior, unit...)?
						<br>
						Also: One page per strategy
						<br>
						Also: This repeats for each 'asses class' ?
					`,
					"de": `Should we group by question type (deals, time, volume) - 
						or as now by strategy (senior, unit...)?
						<br>
						Also: One page per strategy
						<br>
						Also: This repeats for each 'asses class' ?
					`,
				}
			}
		}
		for idx1, trancheTypeName := range trancheTypeNames {

			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 0
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.Label = allLbls["tranche-types"][idx1].Bold()
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
					inp.StyleLbl.Desktop.StyleText.FontSize = 115
				}
			}

			// trancheTypeLbl := allLbls["tranche-types"][idx1]
			lblNumber := trl.S{
				"en": "Total number of new deals",
				"de": "Gesamtzahl neue Abschlüsse",
			}

			restrictedText(qst.WrapPageT(page), trancheTypeName, lblNumber, rT1)
			lblDuration := trl.S{
				"en": "Average time to close a deal in weeks",
				"de": "Durchschnittl. Zeit bis Abschluss in Wochen",
			}
			rangeClosingTime(qst.WrapPageT(page), trancheTypeName, lblDuration)

			volBySegment := trl.S{
				"en": "Total volume of new deals by segment",
				"de": "Gesamtvolumen neuer Abschlüsse nach Marktsegment",
			}
			restrictedText(qst.WrapPageT(page), trancheTypeName, volBySegment, rT2)

			volByRegion := trl.S{
				"en": "Total volume of new deals by region",
				"de": "Gesamtvolumen neuer Abschlüsse nach Region",
			}
			restrictedText(qst.WrapPageT(page), trancheTypeName, volByRegion, rT3)

			volBySector := trl.S{
				"en": "Total volume of new deals by sector",
				"de": "Gesamtvolumen neuer Abschlüsse nach Sektor",
			}
			restrictedText(qst.WrapPageT(page), trancheTypeName, volBySector, rT4)

			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 1
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.Label = trl.S{
						"en": "Should next three add up to 100%?",
						"de": "Should next three add up to 100%?",
					}
				}
			}
			shareESG := trl.S{
				"en": `<b>Share ESG KPIs</b> <br>What is the share of new deals (at Fair Market Value) with explicit ESG targets in the credit documentation?`,
				"de": `<b>Share ESG KPIs</b> <br>What is the share of new deals (at Fair Market Value) with explicit ESG targets in the credit documentation?`,
			}
			rangePercentage(qst.WrapPageT(page), trancheTypeName, shareESG, "esg")

			shareESGRatch := trl.S{
				"en": `<b>Share ESG ratchets</b> <br> What is the share of new deals (at Fair Market Value) with ESG ratchets?`,
				"de": `<b>Share ESG ratchets</b> <br> What is the share of new deals (at Fair Market Value) with ESG ratchets?`,
			}
			rangePercentage(qst.WrapPageT(page), trancheTypeName, shareESGRatch, "esgratch")

			share15Degree := trl.S{
				"en": `<b>Share 1.5°C target</b> <br> What is the share of new deals (at Fair Market Value) where the creditor explicitly states a strategy to add to the 1.5°C target?`,
				"de": `<b>Share 1.5°C target</b> <br> What is the share of new deals (at Fair Market Value) where the creditor explicitly states a strategy to add to the 1.5°C target?`,
			}
			rangePercentage(qst.WrapPageT(page), trancheTypeName, share15Degree, "esg15degrees")

		}
	}

	// page3
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

	} // page3

	// page2
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

	} // page2

	// page4
	{
		page := q.AddPage()

		// page.Section = trl.S{
		// 	"en": "Section 2",
		// 	"de": "Section 2",
		// }
		page.Label = trl.S{"de": "Label long"}
		page.Short = trl.S{"de": "Label<br>short"}
		page.WidthMax("42rem")

		{
			lblMain := trl.S{
				"de": `How do you expect the quality of deals in terms of the risk-return profile change in Q1 2023?`,
				"en": `How do you expect the quality of deals in terms of the risk-return profile change in Q1 2023?`,
			}
			multipleChoiceSingleRow(qst.WrapPageT(page), "xx2", lblMain, mChExample1)
		}

	} // page4

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
