package qst

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// GenerateExample creates an example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func GenerateMinimalExample() *QuestionaireT {
	quest := QuestionaireT{}
	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	for i1 := 0; i1 < 3; i1++ {
		page := newPage()
		quest.Pages = append(quest.Pages, page)
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

func exampleNineLabelledRadios() groupT {
	inp1 := inputT{}
	inp1.Type = "radiogroup"
	inp1.Name = "lemonade"
	inp1.Label = map[string]string{"de": "Brause", "en": "Softdrink"}
	inp1.Desc = map[string]string{"de": "links - zentriert - rechts", "en": "left - centered - right"}
	inp1.HAlignLabel = HLeft
	inp1.HAlignControl = HLeft
	inp1.Radios = []radioT{
		radioT{Val: "fanta", Label: transMapT{"de": "Fanta", "en": "Fanta"}, HAlign: HLeft},
		radioT{Val: "miranda", Label: transMapT{"de": "Miranda", "en": "Baccara"}, HAlign: HCenter},
		radioT{Val: "cholera", Label: transMapT{"de": "Cholera", "en": "Scabies"}, HAlign: HRight},
	}

	inp2 := inputT{}
	inp2.Type = "radiogroup"
	inp2.Name = "language1"
	inp2.Label = map[string]string{"de": "Programmiersprache", "en": "Programming language"}
	inp2.Desc = map[string]string{"de": "rechts - zentriert - links", "en": "right - centered - left"}
	inp2.HAlignLabel = HLeft
	inp2.HAlignControl = HLeft
	inp2.Radios = []radioT{
		radioT{Label: transMapT{"de": "Logik", "en": "Reasoning"}, HAlign: HRight},
		radioT{Label: transMapT{"de": "Basic", "en": "Basics"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Plastik", "en": "Plastics"}, HAlign: HLeft},
	}

	inp3 := inputT{}
	inp3.Type = "radiogroup"
	inp3.Name = "language2"
	inp3.Label = map[string]string{"de": "Programmiersprache", "en": "Programming language"}
	inp3.Desc = map[string]string{"de": "alle zentriert", "en": "all centered"}
	inp3.HAlignLabel = HLeft
	inp3.HAlignControl = HLeft
	inp3.Radios = []radioT{
		radioT{Label: transMapT{"de": "Logik", "en": "Reasoning"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Basic", "en": "Basics"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Plastik", "en": "Plastics"}, HAlign: HCenter},
	}

	gr := groupT{}
	gr.Cols = 4
	gr.Label = map[string]string{"de": "Diverses", "en": "Miscellaneous"}
	gr.Desc = map[string]string{"de": "", "en": ""}
	gr.Inputs = append(gr.Inputs, inp1, inp2, inp3)
	return gr
}

func exampleSixColumnsLabelRight() groupT {
	inp3 := inputT{}
	inp3.Type = "radiogroup"
	inp3.Name = "layoutlang"
	inp3.Label = map[string]string{"de": "Layouting Archtitektur", "en": "Layout architecture"}
	inp3.Desc = map[string]string{"de": "<br>Label rechts, Rest zentriert", "en": "<br>label right, rest centered"}
	inp3.HAlignLabel = HRight
	inp3.HAlignControl = HLeft
	inp3.Radios = []radioT{
		radioT{Label: transMapT{"de": "HTML", "en": "HTML"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Winword", "en": "Microsoft Word"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Tec", "en": "Tec - Latec"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Markdown", "en": "Markdown"}, HAlign: HCenter},
		radioT{Label: transMapT{"de": "Breakdown", "en": "Breakdown"}, HAlign: HCenter},
	}

	gr := groupT{}
	gr.Cols = 6
	gr.Label = map[string]string{"de": "Fünf mit Label", "en": "Five with label"}
	gr.Desc = map[string]string{"de": "", "en": ""}
	gr.Inputs = append(gr.Inputs, inp3)
	return gr
}

func exampleFourCheckboxesPasta() groupT {

	inp1 := inputT{}
	inp1.Type = "checkbox"
	inp1.Name = "pizz"
	inp1.Label = map[string]string{"de": "Fladenbrot", "en": "Pizza"}
	inp1.Desc = map[string]string{"de": "würzig belegtes Fladenbrot", "en": "a traditional Italian dish"}
	inp1.HAlignLabel = HLeft
	inp1.HAlignControl = HLeft

	inp2 := inputT{}
	inp2.Type = "checkbox"
	inp2.Name = "o-oil"
	inp2.Label = map[string]string{"de": "Olivenöl", "en": "Olive oil"}
	inp2.Desc = map[string]string{"de": "ungesättigte Fettsäuren", "en": "digestable fatty acids"}
	inp2.HAlignLabel = HLeft
	inp2.HAlignControl = HLeft

	inp3 := inputT{}
	inp3.Type = "checkbox"
	inp3.Name = "pastatype1"
	inp3.Label = map[string]string{"de": "Röhrennudeln", "en": "Cannelloni"}
	inp3.Desc = map[string]string{"de": "große, dicke Röhrennudeln", "en": "a cylindrical type of pasta"}
	inp3.HAlignLabel = HLeft
	inp3.HAlignControl = HCenter

	inp4 := inputT{}
	inp4.Type = "checkbox"
	inp4.Name = "pastatype2"
	inp4.Label = map[string]string{"de": "Bindfäden", "en": "Spahetti"}
	inp4.Desc = map[string]string{"de": "Form von Teigwaren und Nudeln", "en": "long, thin, solid, cylindrical"}
	inp4.HAlignLabel = HLeft
	inp4.HAlignControl = HCenter

	gr := groupT{}
	gr.Cols = 4
	gr.Label = map[string]string{"de": "Pasta", "en": "Pasta"}
	gr.Desc = map[string]string{"de": "links links - zentriert zentriert", "en": "left left - centered centered"}
	checkboxes := []inputT{inp1, inp2, inp3, inp4}
	gr.Inputs = append(gr.Inputs, checkboxes...)
	return gr
}

func exampleFinlandMatrixNoLabels() groupT {

	names := []string{"fintv_1", "fintv_2"}
	tm := []transMapT{
		{"de": "gut", "en": "good"},
		{"de": "normal", "en": "normal"},
		{"de": "schlecht", "en": "bad"},
		{"de": "keine Ang.", "en": "no answer"},
	}
	gr := radioMatrix(tm, names, nil)
	gr.Cols = 4 //  necessary, otherwise no vspacers
	gr.Label = transMapT{
		"de": "712.",
		"en": "712.",
	}
	gr.Desc = transMapT{
		"de": "Finnisches Fernsehen ist ",
		"en": "Finnish TV is",
	}
	return *gr
}
