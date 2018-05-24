package qst

import (
	"fmt"
	"log"
	"time"
)

func nextQ() string {
	t := time.Now()
	m := t.Month() // 1 - january
	y := t.Year()
	qNow := int(m/3) + 1
	qNext := qNow + 1
	if qNext > 4 {
		qNext = 1
		y++
	}
	return fmt.Sprintf("Q%v %v", qNext, y)
}
func nextY() string {
	t := time.Now()
	y := t.Year()
	y++
	return fmt.Sprintf("%v", y)
}

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

	// Page 1
	{
		page := newPage()
		page.Label = transMapT{"de": "Status und Ausblick", "en": "Status and outlook"}

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

		page.Groups = append(page.Groups, *gr1, *gr2, *gr3)
		quest.Pages = append(quest.Pages, page)
	}

	// page 2
	{
		page := newPage()
		page.Label = transMapT{"de": "Wachstum", "en": "Growth"}

		grp1 := groupT{}
		grp1.Cols = 2 // necessary, otherwise no vspacers
		grp1.Label = transMapT{"de": "3a.", "en": "3a."}
		{
			inp := inputT{}
			inp.Type = "text"
			inp.Name = "y_q_deu"
			inp.Desc = transMapT{
				"de": fmt.Sprintf("Unsere Prognose für das <b>deutsche</b> BIP Wachstum in %v (real, saisonbereinigt, nicht annualisiert):", nextQ()),
				"en": fmt.Sprintf("Our estimate for the <b>German</b> GDP growth in %v (real, seasonally adjusted, non annualized):", nextQ()),
			}
			inp.HAlignLabel = HLeft
			inp.HAlignControl = HCenter
			inp.Validator = "inRange20"

			grp1.Inputs = append(grp1.Inputs, inp)
		}

		{
			inp := inputT{}
			inp.Type = "text"
			inp.Name = "y_y_deu"
			inp.Desc = transMapT{
				"de": fmt.Sprintf("Unsere Prognose für das BIP Wachstum für das Jahr %v (real, saisonbereinigt):", nextY()),
				"en": fmt.Sprintf("Our estimate for the GDP growth in %v (real, seasonally adjusted):", nextY()),
			}
			inp.HAlignLabel = HLeft
			inp.HAlignControl = HCenter
			inp.Validator = "inRange20"

			grp1.Inputs = append(grp1.Inputs, inp)
		}

		grp2 := groupT{}
		grp2.Cols = 2 // necessary, otherwise no vspacers
		grp2.Label = transMapT{"de": "3b.", "en": "3b."}

		{
			inp := inputT{}
			inp.Type = "text"
			inp.Name = "yshr_q_deu"
			inp.Desc = transMapT{
				"de": fmt.Sprintf("Die Wahrscheinlichkeit eines negativen Wachstums des <b>deutschen</b> BIP in %v liegt bei:", nextQ()),
				"en": fmt.Sprintf("The probability of negative growth for the <b>German</b> GDP in %v is:", nextQ()),
			}
			inp.HAlignLabel = HLeft
			inp.HAlignControl = HCenter
			inp.Validator = "inRange100"

			grp2.Inputs = append(grp2.Inputs, inp)
		}

		{
			inp := inputT{}
			inp.Type = "text"
			inp.Name = "yshr_y_deu"
			inp.Desc = transMapT{
				"de": fmt.Sprintf("Die Wahrscheinlichkeit einer Rezession in Deutschland (mind. 2&nbsp;Quartale neg. Wachstum) bis Q4 %v liegt bei:", nextY()),
				"en": fmt.Sprintf("The probability of a recession in Germany (at least 2&nbsp;quarters neg. growth) until Q4 %v is:", nextY()),
			}
			inp.HAlignLabel = HLeft
			inp.HAlignControl = HCenter
			inp.Validator = "inRange100"

			grp2.Inputs = append(grp2.Inputs, inp)
		}

		page.Groups = append(page.Groups, grp1, grp2)
		quest.Pages = append(quest.Pages, page)
	}

	// page 3 - inflation
	{
		page := newPage()
		page.Label = transMapT{"de": "Inflation", "en": "Inflation"}

		grp1 := groupT{}
		grp1.Cols = 2 // necessary, otherwise no vspacers
		grp1.Label = transMapT{"de": "4.", "en": "4."}

		page.Groups = append(page.Groups, grp1)
		quest.Pages = append(quest.Pages, page)
	}

	//
	// Page test
	{
		page := newPage()
		page.Groups = append(page.Groups, exampleFourCheckboxesPasta())
		page.Groups = append(page.Groups, exampleNineLabelledRadios())
		page.Groups = append(page.Groups, exampleSixColumnsLabelRight())
		page.Groups = append(page.Groups, exampleFinlandMatrixNoLabels())
		quest.Pages = append(quest.Pages, page)
	}

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
