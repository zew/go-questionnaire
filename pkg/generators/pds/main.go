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
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Private Debt Survey"}
	// q.Variations = 1

	// page 0
	{
		page := q.AddPage()
		page.SuppressInProgressbar = true

		page.Label = trl.S{"de": "Begrüßung"}
		page.Short = trl.S{"de": "Begrüßung"}

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

	// page1
	{
		page := q.AddPage()

		page.Section = trl.S{"de": "Section 1"}
		page.Label = trl.S{"de": "Label long"}
		page.Short = trl.S{"de": "Label<br>short"}
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
			sentimentSingleRow(qst.WrapPageT(page), "xx2", lblMain, 1)
		}

		{
			lblMain := trl.S{
				"de": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
				"en": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
			}
			sentimentSingleRow(qst.WrapPageT(page), "xx3", lblMain, 2)
		}

	} // page1

	// page2
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

	} // page2

	// page3
	{
		page := q.AddPage()

		page.Section = trl.S{"de": "Section 2"}
		page.Label = trl.S{"de": "Label long"}
		page.Short = trl.S{"de": "Label<br>short"}
		page.WidthMax("42rem")

		lowMidUpperSum(qst.WrapPageT(page), "xx")
		lowMidUpperSum(qst.WrapPageT(page), "yy")

	} // page3

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
