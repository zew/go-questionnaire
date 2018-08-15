package peu2018

import (
	"fmt"
	"math/rand"

	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

// Create creates an minimal example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionaireT, error) {
	q := qst.QuestionaireT{}
	q.Survey = qst.NewSurvey("eup")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	q.LangCode = "de"
	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Umfrage Europäische Parlamente", "en": "European Parliaments Survey"}

	for i1 := 0; i1 < 4; i1++ {
		p := q.AddPage()
		nP := rand.Intn(4) + 2
		for i2 := 0; i2 < nP; i2++ {
			gr := p.AddGroup()
			gr.Cols = 1
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{"de": fmt.Sprintf("tb %v", i2)}
		}
	}
	q.Variations = 8

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	return &q, nil
}
