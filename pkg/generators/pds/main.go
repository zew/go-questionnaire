package pds

import (
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

	q.Survey.Org = trl.S{"en": "ZEW"}
	q.Survey.Name = trl.S{"en": "Private Debt Survey"}
	// q.Variations = 1

	// page 0
	{
		page := q.AddPage()

		page.Label = trl.S{"en": "Chart"}
		page.Short = trl.S{"en": "Chart"}
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

		page.Label = trl.S{"en": "Slider"}
		page.Short = trl.S{"en": "Slider"}
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

	} // page1

	// page2
	{
		page := q.AddPage()

		page.Section = trl.S{"en": "Sociodemographics"}
		page.Label = trl.S{"en": "Age, origin, experience"}
		page.Short = trl.S{"en": "Sociodemo-<br>graphics"}
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q01_age"

				inp.Label = trl.S{"en": "How old are you?"}
				inp.MaxChars = 4
				inp.Step = 1
				inp.Min = 15
				inp.Max = 150
				inp.Validator = "inRange100"

				inp.ColSpan = 4
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
				// inp.Suffix = trl.S{"en": "years"}
				inp.Suffix = trl.S{"en": "&nbsp; years"}
			}

		}

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
