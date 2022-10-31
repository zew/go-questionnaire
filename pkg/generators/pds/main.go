package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func helperVolumesGroup(
	page *qst.WrappedPageT,
	name string,
) {

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_main", name)

			inp.Label = trl.S{
				"de": `Please state the volume (in million Euro) of deals closed in Q4 2022.`,
			}
			inp.Suffix = trl.S{"de": "Million Euro"}

			inp.MaxChars = 12
			inp.Step = 1
			inp.Min = 0
			inp.Max = 1000 * 1000
			// inp.Validator = "inRange100"

			inp.ColSpan = 1
			inp.ColSpanLabel = 3
			inp.ColSpanControl = 2
		}

	}

	// gr2
	{
		gr := page.AddGroup()
		gr.Cols = 2
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Margin = "0 0 0 2rem"
		gr.BottomVSpacers = 3

		{

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = trl.S{
				"de": `Please state the volume (in million Euro) of deals
				 closed in Q4 2022 by market segment: `,
			}

		}

		lbls := map[string]string{
			"low":   "Lower Mid-Market (0-15m € EBITDA)",
			"mid":   "Core Mid-Market (15-50m € EBITDA)",
			"upper": "Upper Mid-Market (>50m € EBITDA)",
		}

		for _, suffx := range []string{"low", "mid", "upper"} {

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v", name, suffx)

			inp.Label = trl.S{
				"de": fmt.Sprintf("- %v", lbls[suffx]),
			}
			inp.Suffix = trl.S{"de": "Million Euro"}

			inp.MaxChars = 10
			inp.Step = 1
			inp.Min = 0
			inp.Max = 1000 * 1000
			// inp.Validator = "inRange100"

			inp.ColSpan = 2
			inp.ColSpanLabel = 3
			inp.ColSpanControl = 2
		}

		inp := gr.AddInput()
		inp.ColSpanControl = 1
		inp.Type = "javascript-block"
		inp.Name = "lowMidUpperSum"

		el1 := trl.S{
			"de": "Sind Sie sicher?",
			"en": "Are you sure?",
		}
		inp.JSBlockTrls = map[string]trl.S{
			"msg": el1,
		}

		inp.JSBlockStrings = map[string]string{}
		for idx, suffx := range []string{"main", "low", "mid", "upper"} {
			inpNamePlaceholder := fmt.Sprintf("%v_%v", "inp", idx+1) // {{.inp_1}}, {{.inp_2}}, ...
			nm := fmt.Sprintf("%v_%v", name, suffx)
			inp.JSBlockStrings[inpNamePlaceholder] = nm
		}

	}

}

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

	// page 1
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
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Name = "slider01"
				inp.Type = "range"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 5
				inp.Style.Desktop.Width = "90%"

				inp.Label = trl.S{
					"de": "Normal Slider",
					"en": "Normal Slider",
				}

				inp.ColSpan = 1
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 8
			}
		}

	} // page1

	// page2
	{
		page := q.AddPage()

		page.Section = trl.S{"de": "Section 1"}
		page.Label = trl.S{"de": "Label long"}
		page.Short = trl.S{"de": "Label<br>short"}
		page.WidthMax("42rem")

		helperVolumesGroup(qst.WrapPageT(page), "xx")
		helperVolumesGroup(qst.WrapPageT(page), "yy")

	} // page2

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
