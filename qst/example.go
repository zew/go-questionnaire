package qst

import (
	"log"
)

func labelsGoodBad() []transMapT {

	tm := []transMapT{
		transMapT{
			"de": "gut",
			"en": "good",
		},
		transMapT{
			"de": "normal",
			"en": "normal",
		},
		transMapT{
			"de": "schlecht",
			"en": "bad",
		},
		transMapT{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsImproveDeteriorate() []transMapT {

	tm := []transMapT{
		transMapT{
			"de": "verbessern",
			"en": "improve",
		},
		transMapT{
			"de": "nicht verändern",
			"en": "not change",
		},
		transMapT{
			"de": "verschlechtern",
			"en": "deteriorate",
		},
		transMapT{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func radioMatrix(headerLabels []transMapT, inpNames []string, rowLabels []transMapT) *groupT {

	grp := groupT{}

	// Header row - first column - empty cell
	if len(rowLabels) > 0 {
		if len(rowLabels) != len(inpNames) { // consistence check
			panic("radioMatrix(): if row labels exist, they should exist for *all* rows")
		}
		inp := inputT{}
		inp.Type = "textblock"
		inp.Label = transMapT{
			"de": "",
			"en": "",
		}
		grp.Inputs = append(grp.Inputs, inp)
	}

	// Header row - next columns
	for _, lbl := range headerLabels {
		inp := inputT{}
		inp.Type = "textblock"
		inp.HAlignLabel = HCenter
		inp.Desc = lbl // for instance transMapT{"de": "gut", "en": "good"}
		grp.Inputs = append(grp.Inputs, inp)
	}

	//
	for i, name := range inpNames {
		inp := inputT{}
		inp.Type = "radiogroup"
		inp.Name = name // "y0_euro"
		if len(rowLabels) > 0 {
			// for instance transMapT{"de": "Euroraum", "en": "euro area"}
			// inp.Label = rowLabels[i]
			inp.Desc = rowLabels[i]
		}
		for i2 := 0; i2 < len(headerLabels); i2++ {
			rad := radioT{}
			rad.HAlign = HCenter
			inp.Radios = append(inp.Radios, rad)
		}
		grp.Inputs = append(grp.Inputs, inp)
	}

	return &grp
}

func GenerateExample2() *QuestionaireT {

	quest := QuestionaireT{}

	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	page1 := newPage()

	//
	//
	names1stMatrix := []string{
		"y0_ez",
		"y0_deu",
		"y0_usa",
		"y0_glob",
	}
	labels123Matrix := []transMapT{
		transMapT{
			"de": "Euroraum",
			"en": "Euro area",
		},
		transMapT{
			"de": "Deutschland",
			"en": "Germany",
		},
		transMapT{
			"de": "USA",
			"en": "US",
		},
		transMapT{
			"de": "Weltwirtschaft",
			"en": "Global economy",
		},
	}
	gr1 := radioMatrix(labelsGoodBad(), names1stMatrix, labels123Matrix)
	gr1.Cols = 5 // necessary, otherwise no vspacers
	gr1.Label = transMapT{
		"de": "1.",
		"en": "1.",
	}
	gr1.Desc = transMapT{
		"de": "Die gesamtwirtschaftliche Situation beurteilen wir als",
		"en": "We assess the overall economic situation as",
	}
	page1.Groups = append(page1.Groups, *gr1)

	//
	//
	names2stMatrix := []string{
		"y_ez",
		"y_deu",
		"y_usa",
		"y_glob",
	}
	gr2 := radioMatrix(labelsImproveDeteriorate(), names2stMatrix, labels123Matrix)
	gr2.Cols = 5 // necessary, otherwise no vspacers
	gr2.Label = transMapT{
		"de": "2a.",
		"en": "2a.",
	}
	gr2.Desc = transMapT{
		"de": "Die gesamtwirtschaftliche Situation wird sich mittelfristig (<b>6</b> Mo.)",
		"en": "The overall economic situation medium term (<b>6</b> months) will",
	}
	page1.Groups = append(page1.Groups, *gr2)

	//
	//
	names3rdMatrix := []string{
		"y24_ez",
		"y24_deu",
		"y24_usa",
		"y24_glob",
	}

	gr3 := radioMatrix(labelsImproveDeteriorate(), names3rdMatrix, labels123Matrix)
	gr3.Cols = 5 // necessary, otherwise no vspacers
	gr3.Label = transMapT{
		"de": "2b.",
		"en": "2b.",
	}
	gr3.Desc = transMapT{
		"de": "Die gesamtwirtschaftliche Situation wird sich langfristig (<b>24</b> Mo.)",
		"en": "The overall economic situation long term (<b>24</b> months) will",
	}

	page1.Groups = append(page1.Groups, *gr3)
	quest.Pages = append(quest.Pages, page1)

	page2 := newPage()
	page1.Groups = append(page1.Groups, exampleFourCheckboxesPasta())
	page1.Groups = append(page1.Groups, exampleNineLabelledRadios())
	page1.Groups = append(page1.Groups, exampleSixColumnsLabelRight())
	page1.Groups = append(page1.Groups, exampleFinlandMatrixNoLabels())

	quest.Pages = append(quest.Pages, page2)

	err := quest.Validate()
	if err != nil {
		log.Fatalf("Error validating questionaire: %v", err)
	}
	err = quest.Save("questionaire-example.json")
	if err != nil {
		log.Fatalf("Error saving questionaire-example.json: %v", err)
	}
	err = quest.Save("questionaire.json")
	if err != nil {
		log.Fatalf("Error saving questionaire.json: %v", err)
	}

	return &quest
}
