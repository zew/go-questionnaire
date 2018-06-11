package min

import (
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

// Create creates an minimal example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create() *qst.QuestionaireT {
	quest := qst.QuestionaireT{}
	quest.Survey = qst.NewSurvey("min")
	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"
	quest.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	quest.Survey.Name = trl.S{"de": "Beispielumfrage", "en": "Example survey"}

	for i1 := 0; i1 < 3; i1++ {
		page := quest.AddPage()
		_ = page
	}
	return &quest
}
