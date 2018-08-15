package min

import (
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

// Create creates an minimal example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionaireT, error) {
	q := qst.QuestionaireT{}
	q.Survey = qst.NewSurvey("min")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	q.LangCode = "de"
	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Beispielumfrage", "en": "Example survey"}

	for i1 := 0; i1 < 3; i1++ {
		page := q.AddPage()
		_ = page
	}
	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	return &q, nil
}
