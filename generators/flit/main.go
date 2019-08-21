package flit

import (
	"fmt"

	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {
	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("flit")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"de": "Deutsch"}
	q.LangCodesOrder = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Financial Literacy Test"}

	// page 1
	{
		page := q.AddPage()
		gr := page.AddGroup()
		gr.Cols = 4

		{
			inp := gr.AddInput()
			inp.Name = fmt.Sprintf("first_name")
			inp.Type = "text"
			inp.Label = trl.S{"de": "Vorname"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			inp.MaxChars = 10
		}
		{
			inp := gr.AddInput()
			inp.Name = fmt.Sprintf("first_name2")
			inp.Type = "text"
			inp.Label = trl.S{"de": "Vorname2"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			inp.MaxChars = 10
		}
	}

	// page 2
	{
		page := q.AddPage()
		gr := page.AddGroup()
		gr.Cols = 4
		inp := gr.AddInput()
		inp.Name = fmt.Sprintf("last_name")
		inp.Type = "text"
		inp.Label = trl.S{"de": "Familienname"}
		inp.ColSpanLabel = 2
		inp.ColSpanControl = 2
		inp.MaxChars = 20
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
