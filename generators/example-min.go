package generators

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/zew/go-questionaire/qst"
)

// MinimalExample creates an example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func MinimalExample() *qst.QuestionaireT {
	quest := qst.QuestionaireT{}
	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	for i1 := 0; i1 < 3; i1++ {
		page := quest.AddPage()
		_ = page
	}

	bts, err := json.MarshalIndent(quest, "", "  ")
	if err != nil {
		log.Fatalf("Marshal questionaire-min-example.json: %v", err)
	}
	err = ioutil.WriteFile("questionaire-min-example.json", bts, 0644)
	if err != nil {
		log.Fatalf("Could not write questionaire-example.json: %v", err)
	}
	return &quest
}
