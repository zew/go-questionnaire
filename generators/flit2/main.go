package flit2

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates questionnaire
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	qst.RadioVali = "mustRadioGroup"
	qst.CSSLabelHeader = "special-line-height-higher"
	qst.CSSLabelRow = "special-input-margin-vertical special-line-height-higher"

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("flit2")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"en": "English"}
	q.LangCodesOrder = []string{"en"} // governs default language code

	q.Survey.Org = trl.S{"en": "ZEW"}
	q.Survey.Name = trl.S{"en": "Financial Literacy Test"}

	for i1 := 0; i1 < 3; i1++ {
		page := q.AddPage()
		gr := page.AddGroup()
		gr.Cols = 3
		inp := gr.AddInput()
		inp.Name = fmt.Sprintf("name%v", i1)
		inp.Type = "text"
		inp.Label = trl.S{"de": "Vorname", "en": "first name"}
		inp.MaxChars = 10
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
