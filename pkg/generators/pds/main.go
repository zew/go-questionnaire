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
			gr.BottomVSpacers = 1
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
			"en": "Manager Characteristics (latest available data)",
			"de": "Manager Characteristics (latest available data)",
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
				"de": `Strategie`,
				"en": `Strategy`,
			}
			multipleChoiceSingleRow(qst.WrapPageT(page), "strategy", lblMain, mCh1)
		}
		// gr2
		{
			lblMain := trl.S{
				"de": `Teamgröße`,
				"en": `Team size`,
			}
			mode2mod := mCh2
			mode2mod.GroupBottomSpacers = 3
			multipleChoiceSingleRow(qst.WrapPageT(page), "teamsize", lblMain, mode2mod)
		}

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
					"en": `A common explanation instead of the repeated one below each block. Should we group by question type (deals, time, volume) - or by strategy (senior, unit...) `,
					"de": `A common explanation instead of the repeated one below each block. Should we group by question type (deals, time, volume) - or by strategy (senior, unit...) `,
				}
			}
		}
		for idx1, trancheTypeName := range trancheTypeNames {
			{
				trancheTypeLbl := allLbls["tranche-types"][idx1]
				restrictedText(qst.WrapPageT(page), trancheTypeName, trancheTypeLbl, rT1)
			}

			{
				trancheTypeLbl := allLbls["tranche-types"][idx1]
				restrictedText(qst.WrapPageT(page), trancheTypeName, trancheTypeLbl, rT2)
			}
		}
	}

	// page2
	{
		page := q.AddPage()

		page.Label = trl.S{"de": "Label long"}
		page.Short = trl.S{"de": "Label<br>short"}

		page.Label = trl.S{
			"en": "Priority matrix and multi choice",
			"de": "Priority matrix and multi choice",
		}
		page.Short = trl.S{
			"en": "PM and MC",
			"de": "PM and MC",
		}

		page.WidthMax("42rem")

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
				"de": `What GICS sectors provides the most attractive 
					investment opportunities in the next three months? 
					Please rank the top three.`,
			}
			prio3Matrix(qst.WrapPageT(page), "xx", lblMain, inps, lbls)
		}

		{
			inps := []string{
				"row1",
				"row2",
				"row3",
				"utilities",
				"real_estate",
			}

			lbls := map[string]string{
				"row1":        "row1",
				"row2":        "row2",
				"row3":        "row3",
				"utilities":   "Utilities 2",
				"real_estate": "Real Estate 2",
			}

			lblMain := trl.S{
				"de": `Another priority type question.`,
				"en": `Another priority type question.`,
			}
			prio3Matrix(qst.WrapPageT(page), "xx1", lblMain, inps, lbls)
		}

		{
			lblMain := trl.S{
				"de": `How do you expect the quality of deals in terms of the risk-return profile change in Q1 2023?`,
				"en": `How do you expect the quality of deals in terms of the risk-return profile change in Q1 2023?`,
			}
			multipleChoiceSingleRow(qst.WrapPageT(page), "xx2", lblMain, mChExample1)
		}

		{
			lblMain := trl.S{
				"de": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
				"en": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
			}
			multipleChoiceSingleRow(qst.WrapPageT(page), "xx3", lblMain, mChExample2)
		}

	} // page2

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

		lbl := trl.S{
			"de": `Please state the volume (in million Euro) of deals closed in Q4 2022.`,
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022.`,
		}
		restrictedText(qst.WrapPageT(page), "xx", lbl, rT2)
		restrictedText(qst.WrapPageT(page), "yy", lbl, rT2)

	} // page4

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
