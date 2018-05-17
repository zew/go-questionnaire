package qst

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func GenerateExample2() *QuestionaireT {

	quest := QuestionaireT{}

	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	page1 := newPage()

	i1a := inputT{}
	i1a.Type = "textblock"
	i1a.Label = transMapT{
		"de": "gut",
		"en": "good",
	}
	i1b := inputT{}
	i1b.Type = "textblock"
	i1b.Label = transMapT{
		"de": "normal",
		"en": "normal",
	}
	i1c := inputT{}
	i1c.Type = "textblock"
	i1c.Label = transMapT{
		"de": "schlecht",
		"en": "bad",
	}
	i1d := inputT{}
	i1d.Type = "textblock"
	i1d.Label = transMapT{
		"de": "keine Ang.",
		"en": "no answer",
	}

	i2 := inputT{}
	i2.Type = "radiogroup"
	i2.Name = "y0_euro"
	i2.Label = transMapT{
		"de": "Eurozone",
		"en": "eurozone",
	}
	i2.Radios = []radioT{
		radioT{},
		radioT{},
		radioT{},
		radioT{},
	}
	i3 := i2
	i3.Name = "y0_germany"
	i3.Label = transMapT{
		"de": "Deutschland",
		"en": "Germany",
	}

	gr1 := groupT{}
	gr1.Cols = 4
	gr1.Label = transMapT{
		"de": "1.",
		"en": "1.",
	}
	gr1.Desc = transMapT{
		"de": "Die gesamtwirtschaftliche Situation beurteilen wir als",
		"en": "We assess the overall market situation as",
	}
	gr1.Members = append(gr1.Members, i1a, i1b, i1c, i1d, i2, i3)
	page1.Elements = append(page1.Elements, gr1)

	quest.Pages = append(quest.Pages, page1)

	page2 := newPage()
	quest.Pages = append(quest.Pages, page2)

	bts, err := json.MarshalIndent(quest, "", "  ")
	if err != nil {
		log.Fatalf("Marshal default questionaire: %v", err)
	}
	err = ioutil.WriteFile("questionaire-example2.json", bts, 0644)
	if err != nil {
		log.Fatalf("Could not write example-quest2.json: %v", err)
	}
	return &quest
}
