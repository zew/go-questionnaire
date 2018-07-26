package peu2018

import (
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

// Create creates an minimal example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionaireT, error) {
	quest := qst.QuestionaireT{}
	quest.Survey = qst.NewSurvey("eup")
	quest.Survey.Params = params
	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"
	quest.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	quest.Survey.Name = trl.S{"de": "Umfrage Europäische Parlamente", "en": "European Parliaments Survey"}

	for i1 := 0; i1 < 3; i1++ {
		page := quest.AddPage()
		_ = page
	}
	(&quest).Hyphenize()
	return &quest, nil
}
