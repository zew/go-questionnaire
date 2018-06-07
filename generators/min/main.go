package min

import (
	"github.com/zew/go-questionaire/qst"
)

// Create creates an example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create() *qst.QuestionaireT {
	quest := qst.QuestionaireT{}
	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	for i1 := 0; i1 < 3; i1++ {
		page := quest.AddPage()
		_ = page
	}
	return &quest
}
