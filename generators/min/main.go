package min

import (
	"github.com/zew/go-questionaire/qst"
)

// Create creates an minimal example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create() *qst.QuestionaireT {
	quest := qst.QuestionaireT{}
	quest.Survey = qst.NewSurvey("min")
	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	for i1 := 0; i1 < 3; i1++ {
		page := quest.AddPage()
		_ = page
	}
	return &quest
}
