package lt2020

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	// qst.RadioVali = "mustRadioGroup"
	qst.CSSLabelHeader = ""
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("lt2020")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"de": "Deutsch"}
	q.LangCodesOrder = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Landtagsumfrage"}

	groupOrdinal := "[groupID]"

	//
	// Page 1
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zur grundgesetzlichen Schuldenbremse",
		}
		p.Label = trl.S{
			"de": "Teil 1",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Schuldenbremse 1",
		}

		{
			// p1q1
			names1stMatrix := []string{"einhaltung2020"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labels9("ausgeschlossen", "", "sicher"), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"de": `Erinnern Sie sich bitte an die Zeit zu Anfang 2020, also vor Ausbruch der Corona-Krise.<br>
				Für wie wahrscheinlich hielten Sie es zu diesem Zeitpunkt, 
				dass Ihr Bundesland die Vorgaben der grundgesetzlichen Schuldenbremse einhalten 
				und ab 2020 einen (konjunkturbereinigt) ausgeglichenen Haushalt aufweisen würde?`,
			}
		}

		{
			{
				// p1q2a
				gr := p.AddGroup()
				gr.Cols = 2 // necessary, otherwise no vspacers
				gr.OddRowsColoring = false
				gr.Label = trl.S{
					"de": fmt.Sprintf(`Frage %v: <br>`, groupOrdinal),
				}
				gr.Desc = trl.S{
					"de": fmt.Sprintf(`Die Corona-Krise führt voraussichtlich zu einem starken Rückgang der wirtschaftlichen Aktivität 
				in diesem Jahr und damit verbunden zu einer Verschlechterung der öffentlichen Haushaltslage.`),
				}

				{
					inp := gr.AddInput()
					inp.Name = "bundesland_defizit_2020"
					inp.Type = "number"
					inp.Step = 0.1
					inp.MaxChars = 6
					inp.Validator = "inRange20"
					inp.Desc = trl.S{
						"de": `a) Welches Haushaltsdefizit erwarten Sie für Ihr Bundesland in 2020?`,
					}
					inp.Suffix = trl.S{
						"de": `Mrd. Euro`,
					}
				}

			}

			// p1q2b
			{
				names1stMatrix := []string{"bundesland_wachstum_2020"}
				emptyRowLabels := []trl.S{}
				gr := p.AddRadioMatrixGroup(labelsFiverPercentages(), names1stMatrix, emptyRowLabels, 1)
				gr.Cols = 5 // necessary, otherwise no vspacers
				gr.OddRowsColoring = false
				gr.Desc = trl.S{
					"de": fmt.Sprintf(`b) Welches wirtschaftliche Wachstum (Bruttoinlandsprodukt / BIP) 
					erwarten Sie für Ihr Bundesland in 2020?`),
				}
			}

			// p1q3
			{
				names1stMatrix := []string{"balanced_budget"}
				emptyRowLabels := []trl.S{}
				gr := p.AddRadioMatrixGroup(labels9("überhaupt nicht erstrebenswert", "", "sehr erstrebenswert"), names1stMatrix, emptyRowLabels, 1)
				gr.Cols = 9 // necessary, otherwise no vspacers
				gr.OddRowsColoring = false
				gr.Label = trl.S{
					"de": fmt.Sprintf(`Frage %v: <br>`, groupOrdinal),
				}
				gr.Desc = trl.S{
					"de": fmt.Sprintf(`Für wie erstrebenswert erachten Sie es, dass Ihr Bundesland wieder einen ausgeglichenen Haushalt 
					gemäß Vorgaben der Schuldenbremse vorlegt, wenn die Corona-Krise vorbei ist?<br>`),
				}
			}

		}

		/* 		{
		   			gr := p.AddGroup()
		   			gr.Cols = 1
		   			gr.Width = 99
		   			{
		   				inp := gr.AddInput()
		   				inp.Type = "button"
		   				inp.Name = "submitBtn"
		   				inp.Response = "1"
		   				inp.Label = trl.S{
		   					"de": "Weiter",
		   				}
		   				inp.AccessKey = "n"
		   				inp.ColSpanControl = 1
		   				inp.HAlignControl = qst.HRight
		   			}
		   		}
		*/
	}

	//
	// Page 2
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zur grundgesetzlichen Schuldenbremse",
		}
		p.Label = trl.S{
			"de": "Teil 2",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Schuldenbremse 2",
		}

		// group1
		{
			gr := p.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": `Frage 4: <br>`,
			}
			gr.Desc = trl.S{
				"de": `Vor Ausbruch der Corona-Krise wurde vermehrt diskutiert, 
					ob die grundgesetzliche Schuldenbremse angemessen ist.<br>
				Welche Position haben Sie zu Anfang 2020 unterstützt, also vor Ausbruch der Corona-Krise? <br>
					(Mehrfachantworten möglich)
					`,
			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_verschaerfung_1"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Verschärfung der Schuldenbremse und Verringerung des Verschuldungsspielraums.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_lockerung_1"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Lockerung der Schuldenbremse und generelle Erhöhung des Verschuldungsspielraums.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "schulden_f_infrastruktur_1"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Verschuldung für Investitionen in Infrastruktur zulassen.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "schulden_f_klima_1"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Verschuldung für Klimapolitik zulassen.`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "sb_unveraendert_1"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Schuldenbremse sollte unverändert bleiben.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_sonstiges_1"
				inp.Type = "text"
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.MaxChars = 40
				inp.Desc = trl.S{
					"de": `Sonstiges`,
				}
			}

		}
		// group2
		{
			gr := p.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": `Frage 5: <br> 
					`,
			}
			gr.Desc = trl.S{
				"de": `Welche Position in der Diskussion zur grundgesetzlichen Schuldenbremse würden 
				Sie heute am ehesten unterstützen?<br>
					(Mehrfachantworten möglich)
					`,
			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_verschaerfung_2"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Verschärfung der Schuldenbremse und Verringerung des Verschuldungsspielraums.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_lockerung_2"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Lockerung der Schuldenbremse und generelle Erhöhung des Verschuldungsspielraums.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "schulden_f_infrastruktur_2"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Verschuldung für Investitionen in Infrastruktur zulassen.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "schulden_f_klima_2"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Verschuldung für Klimapolitik zulassen.`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "sb_unveraendert_2"
				inp.Type = "checkbox"
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HCenter
				inp.Desc = trl.S{
					"de": `Schuldenbremse sollte unverändert bleiben.`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_sonstiges_2"
				inp.Type = "text"
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.MaxChars = 40
				inp.Desc = trl.S{
					"de": `Sonstiges`,
				}
			}

		}

		/* 		{
		   			gr := p.AddGroup()
		   			gr.Cols = 1
		   			gr.Width = 99
		   			{
		   				inp := gr.AddInput()
		   				inp.Type = "button"
		   				inp.Name = "submitBtn"
		   				inp.Response = "2"
		   				inp.Label = trl.S{
		   					"de": "Weiter",
		   				}
		   				inp.AccessKey = "n"
		   				inp.ColSpanControl = 1
		   				inp.HAlignControl = qst.HRight
		   			}
		   		}
		*/
	}

	//
	// Page 3
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 1",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 1",
		}

		{
			// p3q6
			names1stMatrix := []string{"leistung_vergleichbarkeit"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsFiverWichtig(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": fmt.Sprintf(`Frage 6: <br>`),
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`Für wie wichtig halten Sie es, dass Schülerleistungen zwischen den Bundesländern vergleichbar sind?`),
			}

			// p3q7
			{
				inp := gr.AddInput()
				inp.Name = "leistung_vergleichbarkeit_buerger"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Frage 7: <br>`),
				}
				inp.Desc = trl.S{
					"de": `Was schätzen Sie, welcher Anteil der Bürger*innen Ihres Bundeslandes hält es für „sehr“ oder „eher“ wichtig, 
					dass Schülerleistungen zwischen den Bundesländern vergleichbar sind?`,
				}
				inp.Suffix = trl.S{
					"de": `Prozent`,
				}
			}

		}

		/* 		{
		   			gr := p.AddGroup()
		   			gr.Cols = 1
		   			gr.Width = 99
		   			{
		   				inp := gr.AddInput()
		   				inp.Type = "button"
		   				inp.Name = "submitBtn"
		   				inp.Response = "3"
		   				inp.Label = trl.S{
		   					"de": "Weiter",
		   				}
		   				inp.AccessKey = "n"
		   				inp.ColSpanControl = 1
		   				inp.HAlignControl = qst.HRight
		   			}
		   		}
		*/
	}

	//
	// Page 4
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 2",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 2",
		}

		{
			gr := p.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": fmt.Sprintf(`Frage 8<br>`),
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
				<b>a)</b> <br>
				Was <b>schätzen</b> Sie, welcher Anteil der öffentlichen Finanzierung von Schulen 
				(allgemeinbildend und beruflich) kommt derzeit von den verschiedenen staatlichen Ebenen in Deutschland?
				<br>
				<div style='display: inline-block; margin: 6px 0; font-size: 85%%'>Bitte achten Sie darauf, 
				dass Ihre Angaben insgesamt 100%% ergeben.</div>
				`),
			}

			// p4q8a
			{
				inp := gr.AddInput()
				inp.Name = "tb1"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Gemeinden`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb2"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Bundesländer`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb3"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Bund`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb4"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Summe`),
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "bildungsfinanzierung_gemeinden_ist"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `%`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "bildungsfinanzierung_laender_ist"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `%`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "bildungsfinanzierung_bund_ist"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `%`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "tb5"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`100%%`),
				}
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				// "de": fmt.Sprintf(`Frage 8<br>`),
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
				<b>b)</b> <br>
				Und welcher Anteil der öffentlichen Finanzierung für Schulen (allgemeinbildend und beruflich) 
				<b>sollte</b> Ihrer Meinung nach von den verschiedenen staatlichen Ebenen kommen?
				
				`),
			}

			// p4q8b
			{
				inp := gr.AddInput()
				inp.Name = "tb6"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Gemeinden`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb7"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Bundesländer`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb8"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Bund`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb9"
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`Summe`),
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "bildungsfinanzierung_gemeinden_soll"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `%`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "bildungsfinanzierung_laender_soll"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `%`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "bildungsfinanzierung_bund_soll"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
				inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `%`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Name = "tb10"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`100%%`),
				}
			}

		}

	}

	//
	// Page Finish
	{
		p := q.AddPage()
		p.NoNavigation = true
		p.Label = trl.S{
			"de": "Vielen Dank",
		}

		{
			// Only one group => shuffling is no problem
			gr := p.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
				inp.Desc = trl.S{
					"de": "Danke für Ihre Teilnahme an unserer Umfrage.",
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
				inp.Desc = trl.S{
					"de": "<span style='font-size: 100%;'>Ihre Eingaben wurden gespeichert.</span>",
				}
			}

		}

	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
