package min

import (
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {
	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("min")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	q.LangCodesOrder = []string{"en", "de"}
	q.LangCode = "de"
	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Beispielumfrage", "en": "Example survey"}

	for i1 := 0; i1 < 3; i1++ {
		page := q.AddPage()
		_ = page
	}
	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	return &q, nil
}
