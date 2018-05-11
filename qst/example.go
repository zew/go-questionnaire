package qst

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// GenerateExample creates an example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func GenerateExample() *QuestionaireT {

	quest := QuestionaireT{}

	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	for i1 := 0; i1 < 3; i1++ {
		page := newPage()
		for i2 := 0; i2 < i1+1; i2++ {

			inp := newInput()
			if i2 == 1 {
				inp.Type = "checkbox"
			}

			gr := groupT{}
			gr.Title = map[string]string{"de": fmt.Sprintf("Gruppe-%v-Titel", i2), "en": fmt.Sprintf("Group-%v-Title", i2)}
			gr.Desc = map[string]string{"de": fmt.Sprintf("Gruppe-%v-Beschreibung", i2), "en": fmt.Sprintf("Group-%v-Description", i2)}
			gr.Members = append(gr.Members, inp)
			page.Elements = append(page.Elements, gr)

			if i1 == 0 {
				inp1 := inputT{}
				inp1.Type = "checkboxgroup"
				inp1.Name = "foo"
				inp1.Vals = []string{"Fanta", "Miranda", "Cholera"}
				inp1.Title = map[string]string{"de": "Brause", "en": "Softdrink"}
				inp1.Desc = map[string]string{"de": "", "en": ""}

				inp2 := inputT{}
				inp2.Type = "radiogroup"
				inp2.Name = "bar"
				inp2.Vals = []string{"Logic", "Basic", "Plastic"}
				inp2.Title = map[string]string{"de": "Programmiersprache", "en": "Programming language"}
				inp2.Desc = map[string]string{"de": "", "en": ""}

				gr := groupT{}
				gr.Cols = 2
				gr.Title = map[string]string{"de": "Gruppe-Pasta", "en": "Group-Pasta"}
				gr.Desc = map[string]string{"de": "", "en": ""}
				gr.Members = append(gr.Members, inp1, inp2)

				page.Elements = nil
				page.Elements = append(page.Elements, gr)
			}

			if i1 == 1 {

				inp1 := inputT{}
				inp1.Type = "checkbox"
				inp1.Name = "pizz"
				inp1.Title = map[string]string{"de": "Fladenbrot", "en": "Pizza"}
				inp1.Desc = map[string]string{"de": "würzig belegtes Fladenbrot", "en": "a traditional Italian dish"}

				inp2 := inputT{}
				inp2.Type = "checkbox"
				inp2.Name = "o-oil"
				inp2.Title = map[string]string{"de": "Olivenöl", "en": "olive oil"}
				inp2.Desc = map[string]string{"de": "ungesättigte Fettsäuren", "en": "digestable fatty acids"}

				inp3 := inputT{}
				inp3.Type = "checkbox"
				inp3.Name = "pastatype1"
				inp3.Title = map[string]string{"de": "Röhrennudeln", "en": "Cannelloni"}
				inp3.Desc = map[string]string{"de": "große, dicke Röhrennudeln", "en": "a cylindrical type of pasta"}

				inp4 := inputT{}
				inp4.Type = "checkbox"
				inp4.Name = "pastatype2"
				inp4.Title = map[string]string{"de": "Bindfäden", "en": "Spahetti"}
				inp4.Desc = map[string]string{"de": "Form von Teigwaren und Nudeln", "en": "long, thin, solid, cylindrical"}

				gr := groupT{}
				gr.Cols = 2
				gr.Title = map[string]string{"de": "Gruppe-Pasta", "en": "Group-Pasta"}
				gr.Desc = map[string]string{"de": "", "en": ""}
				gr.Members = append(gr.Members, inp1, inp2, inp3, inp4)

				page.Elements = nil
				page.Elements = append(page.Elements, gr)
			}

		}
		quest.Pages = append(quest.Pages, page)
	}

	bts, err := json.MarshalIndent(quest, "", "  ")
	if err != nil {
		log.Fatalf("Marshal default questionaire: %v", err)
	}
	err = ioutil.WriteFile("questionaire-example.json", bts, 0644)
	if err != nil {
		log.Fatalf("Could not write example-quest.json: %v", err)
	}
	return &quest
}
