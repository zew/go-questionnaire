package generators

import (
	"fmt"
	"log"
	"time"

	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

func nextWaveID() string {
	t := time.Now()
	if t.Day() > 20 {
		t = t.AddDate(0, 1, 0)
	}
	return t.Format("2006-01")
}

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

func labelsGoodBad() []trl.S {

	tm := []trl.S{
		{
			"de": "gut",
			"en": "good",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "schlecht",
			"en": "bad",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

func labelsImproveDeteriorate() []trl.S {

	tm := []trl.S{
		{
			"de": "verbessern",
			"en": "improve",
		},
		{
			"de": "nicht verändern",
			"en": "not change",
		},
		{
			"de": "verschlechtern",
			"en": "deteriorate",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	return tm

}

// FMT creates a JSON file for a financial markets survey
func FMT() *qst.QuestionaireT {

	quest := qst.QuestionaireT{}
	quest.WaveID = qst.NewWaveID()

	// This should be put into a separate, admin protected http handler
	for i := 99 * 1000; i > 99*1000-10; i-- {
		checkStr := fmt.Sprintf("u=%v&wave_id=%v&h=%v", i, quest.WaveID, lgn.Get().Salt)
		hsh := lgn.Md5Str([]byte(checkStr))
		log.Printf("%8v - %8v - %v", i, quest.WaveID, hsh)
	}

	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	// Page 1
	{
		page := quest.AddPage()
		page.Section = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		page.Label = trl.S{"de": "Status und Ausblick", "en": "Status and outlook"}

		//
		//
		labels123Matrix := []trl.S{
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "Deutschland",
				"en": "Germany",
			},
			{
				"de": "USA",
				"en": "US",
			},
			{
				"de": "Weltwirtschaft",
				"en": "Global economy",
			},
		}

		// gr1
		{
			names1stMatrix := []string{
				"y0_ez",
				"y0_deu",
				"y0_usa",
				"y0_glob",
			}
			gr := page.AddRadioMatrixGroup(labelsGoodBad(), names1stMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{
				"de": "1.",
				"en": "1.",
			}
			gr.Desc = trl.S{
				"de": "Die gesamtwirtschaftliche Situation beurteilen wir als",
				"en": "We assess the overall economic situation as",
			}

		}

		//
		// gr2
		{
			names2stMatrix := []string{
				"y_ez",
				"y_deu",
				"y_usa",
				"y_glob",
			}
			gr := page.AddRadioMatrixGroup(labelsImproveDeteriorate(), names2stMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{
				"de": "2a.",
				"en": "2a.",
			}
			gr.Desc = trl.S{
				"de": "Die gesamtwirtschaftliche Situation wird sich mittelfristig (<b>6</b> Mo.)",
				"en": "The overall economic situation medium term (<b>6</b> months) will",
			}

		}

		//
		// gr3
		{
			names3rdMatrix := []string{
				"y24_ez",
				"y24_deu",
				"y24_usa",
				"y24_glob",
			}

			gr := page.AddRadioMatrixGroup(labelsImproveDeteriorate(), names3rdMatrix, labels123Matrix)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{
				"de": "2b.",
				"en": "2b.",
			}
			gr.Desc = trl.S{
				"de": "Die gesamtwirtschaftliche Situation wird sich langfristig (<b>24</b> Mo.)",
				"en": "The overall economic situation long term (<b>24</b> months) will",
			}

		}

	}

	// page 2
	{
		page := quest.AddPage()
		page.Label = trl.S{"de": "Wachstum", "en": "Growth"}

		{
			gr := page.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "3a.", "en": "3a."}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "y_q_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das <b>deutsche</b> BIP Wachstum in %v (real, saisonbereinigt, nicht annualisiert):", nextQ()),
					"en": fmt.Sprintf("Our estimate for the <b>German</b> GDP growth in %v (real, seasonally adjusted, non annualized):", nextQ()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "y_y_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das BIP Wachstum für das Jahr %v (real, saisonbereinigt):", nextY()),
					"en": fmt.Sprintf("Our estimate for the GDP growth in %v (real, seasonally adjusted):", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "3b.", "en": "3b."}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "yshr_q_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Die Wahrscheinlichkeit eines negativen Wachstums des <b>deutschen</b> BIP in %v liegt bei:", nextQ()),
					"en": fmt.Sprintf("The probability of negative growth for the <b>German</b> GDP in %v is:", nextQ()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "yshr_y_deu"
				inp.MaxChars = 4
				inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Die Wahrscheinlichkeit einer Rezession in Deutschland (mind. 2&nbsp;Quartale neg. Wachstum) bis Q4 %v liegt bei:", nextY()),
					"en": fmt.Sprintf("The probability of a recession in Germany (at least 2&nbsp;quarters neg. growth) until Q4 %v is:", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

		}

	}

	//
	// page 3 - inflation
	{
		page := quest.AddPage()
		page.Label = trl.S{"de": "Inflation", "en": "Inflation"}
		gr := page.AddGroup()
		gr.Cols = 2 // necessary, otherwise no vspacers
		gr.Label = trl.S{"de": "4.", "en": "4."}
	}

	//
	// page 4 - Credit situation
	{
		page := quest.AddPage()
		page.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
		page.Label = trl.S{"de": "Markt", "en": "market"}

		gr := page.AddGroup()
		gr.Cols = 2 // necessary, otherwise no vspacers
		gr.Label = trl.S{"de": "5.", "en": "5."}
	}

	{
		page := quest.AddPage()
		page.Section = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Label = trl.S{"de": "", "en": ""}

		{
			gr := page.AddGroup()
			gr.Cols = 1 // necessary, otherwise no vspacers
			gr.Label = trl.S{"de": "", "en": ""}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Desc = trl.S{
					"de": "Wollen Sie uns noch etwas mitteilen?",
					"en": "Any remarks or advice for us?",
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "remark"
				inp.MaxChars = 300
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			// gr.Desc = trl.S{"de": "Abschluss", "en": "Finish"}
			{

				inp := gr.AddInput()
				inp.Type = "radiogroup"
				inp.Name = "finished"

				inp.Label = trl.S{"de": "Abschluss", "en": "Finish"}
				inp.Desc = trl.S{"de": "", "en": ""}

				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1 // ignored for radiogroup

				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "\nFragebogen ist abgeschlossen \nund kann nicht mehr geöffnet werden.",
						"en": "\nQuestionaire is finished.\nNo more edits.",
					}
				}
				{
					rad := inp.AddRadio()
					rad.HAlign = qst.HLeft
					rad.HAlign = qst.HCenter
					rad.Label = trl.S{
						"de": "Zugang bleibt bestehen.  \nDaten können in weiteren Sitzungen \ngeändert/ergänzt werden.",
						"en": "Leave questionaire open. \nData  can be changed/completed     \nin later sessions.",
					}
				}
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "Durch Klicken auf 'OK' erhalten Sie eine Zusammenfassung Ihrer Antworten.",
					"en": "By Clicking 'OK' you receive a summary of your answers.",
				}
				inp.ColSpanLabel = 2
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submit_final_ok"
				inp.Response = "OK"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "OK",
					"en": "OK",
				}
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
			}
		}

		// page.ExampleSixColumnsLabelRight()

	}

	// quest.ClosingTime = time.Now()

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
