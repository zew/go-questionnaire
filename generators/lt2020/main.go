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

	/*

	 */

	ctr.Reset()
	ctrPages := ctr.New()

	// qst.RadioVali = "mustRadioGroup"
	qst.HeaderClass = "go-quest-opt-label"
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("lt2020")
	q.Survey.Params = params
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Landtagsumfrage"}

	// groupOrdinal := "[groupID]"

	vars, err := q.Survey.Param("varianten")
	if err != nil {
		return &q, err
	}
	if len(vars) != 4 {
		return &q, fmt.Errorf("varianten must have four '0's and '1's - is %v", vars)
	}
	varianten := []bool{true, true, true, true}
	for idx, v := range vars {
		if v != '0' && v != '1' {
			return &q, fmt.Errorf("varianten must consist of '0's and '1's - is %v", vars)
		}
		varianten[idx] = (v == '1')
	}

	// besseren/schlechteren
	aob, err := q.Survey.Param("aboveOrBelowMedian")
	if err != nil {
		return &q, err
	}

	//
	// Page 1
	{

		p := q.AddPage()
		p.NoNavigation = true
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Begrüßung",
		}
		p.Label = trl.S{
			"de": "",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Begrüßung",
		}

		{
			// greeting
			gr := p.AddGroup()
			gr.Cols = 2 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": "",
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
				Sehr geehrte Frau Landtagsabgeordnete, <br>
				Sehr geehrter Herr Landtagsabgeordneter,<br>
				<br>
				gemeinsam mit der Universität Mannheim untersucht das ZEW – Leibniz-Zentrum für Europäische Wirtschaftsforschung im Rahmen einer Umfrage, wie Landtagsabgeordnete über die im Grundgesetz verankerte „Schuldenbremse“ und den Bildungsföderalismus in Deutschland denken. Dabei ist uns auch an Ihrer Einschätzung sehr gelegen.
				Durch Ihre Beteiligung unterstützen Sie ein wichtiges wissenschaftliches Forschungsprojekt zur deutschen Finanzpolitik, das von der Deutschen Forschungsgemeinschaft (DFG) gefördert wird. Wir sind Ihnen dafür sehr dankbar.<br>
				<br>
				Mit freundlichen Grüßen<br>
				Prof. Dr. Friedrich Heinemann<br>
				
<!--
				<input type='hidden' name='variant'          value='%v' />
				<input type='hidden' name='besserschlechter' value='%v' />
-->
				`, vars, aob),
			}

			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "variant"
				inp.Response = vars
			}
			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "besserschlechter"
				inp.Response = aob
			}

			{
				gr := p.AddGroup()
				gr.Cols = 1
				gr.Width = 99
				{
					inp := gr.AddInput()
					inp.Type = "button"
					inp.Name = "submitBtn"
					inp.Response = "1"
					inp.Label = trl.S{
						"de": "Umfrage starten",
					}
					inp.AccessKey = "n"
					inp.ColSpanControl = 1
					inp.HAlignControl = qst.HRight
				}
			}

		}
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
			"de": "Teil 1",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Verschuldung 1",
		}

		{
			// p1q1
			names1stMatrix := []string{"einhaltung2020"}
			emptyRowLabels := []trl.S{}
			// ausge&shy;schlossen  - manual override of hyphenization
			gr := p.AddRadioMatrixGroup(labels9("ausge&shy;schlossen", "", "sicher"), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 1
			}
			gr.Desc = trl.S{
				"de": `Erinnern Sie sich bitte an die Zeit zu Anfang 2020, also vor Ausbruch der Corona-Krise.<br>
				Für wie wahrscheinlich hielten Sie es zu diesem Zeitpunkt, 
				dass Ihr Bundesland die Vorgaben der grundgesetzlichen Schuldenbremse einhalten 
				und ab 2020 einen (konjunkturbereinigt) ausgeglichenen Haushalt aufweisen würde?`,
			}

		}
	}

	//
	// Page 3
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
			"de": "Verschuldung 2",
		}

		{
			// p1q2a
			gr := p.AddGroup()
			gr.Cols = 2 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 2
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`Die Corona-Krise führt voraussichtlich zu einem starken Rückgang der wirtschaftlichen Aktivität 
				in diesem Jahr und damit verbunden zu einer Verschlechterung der öffentlichen Haushaltslage.<br>
				
				<b>a)</b> Welches Haushaltsdefizit erwarten Sie für Ihr Bundesland in 2020?
				`),
			}

			{
				inp := gr.AddInput()
				inp.Name = "bundesland_defizit_2020"
				inp.Type = "number"
				inp.Step = 0.1
				inp.MaxChars = 6
				// inp.Validator = "inRange1000"
				inp.ColSpanLabel = 0
				inp.Desc = trl.S{
					"de": ``,
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
				"de": fmt.Sprintf(`<b>b)</b> Welches wirtschaftliche Wachstum (BIP) 
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
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 3
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`Für wie erstrebenswert erachten Sie es, dass Ihr Bundesland wieder einen ausgeglichenen Haushalt 
					gemäß Vorgaben der Schuldenbremse vorlegt, wenn die Corona-Krise vorbei ist?<br>`),
			}
		}

	}

	//
	// Page 4
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zur grundgesetzlichen Schuldenbremse",
		}
		p.Label = trl.S{
			"de": "Teil 3",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Verschuldung 3",
		}

		// group1
		{
			gr := p.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 4
			}
			gr.Desc = trl.S{
				"de": `Vor Ausbruch der Corona-Krise wurde vermehrt diskutiert, 
					ob die grundgesetzliche Schuldenbremse angemessen ist.<br>
				Welche Position haben Sie zu Anfang 2020 unterstützt, also vor Ausbruch der Corona-Krise? <br>
					(Mehrfachantworten möglich)
					`,
			}

			if !varianten[1] {

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

			} else {

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
				//
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
					inp.Name = "sb_verschaerfung_1"
					inp.Type = "checkbox"
					inp.ColSpanLabel = 4
					inp.ColSpanControl = 1
					inp.HAlignControl = qst.HCenter
					inp.Desc = trl.S{
						"de": `Verschärfung der Schuldenbremse und Verringerung des Verschuldungsspielraums.`,
					}
				}

			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_sonstiges_1"
				inp.Type = "textarea"
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.MaxChars = 155
				inp.CSSLabel = "vertical-align-top-two-rows"
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
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 5
			}
			gr.Desc = trl.S{
				"de": `Welche Position in der Diskussion zur grundgesetzlichen Schuldenbremse würden 
				Sie heute am ehesten unterstützen?<br>
					(Mehrfachantworten möglich)
					`,
			}

			if !varianten[1] {

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

			} else {

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
					inp.Name = "sb_verschaerfung_2"
					inp.Type = "checkbox"
					inp.ColSpanLabel = 4
					inp.ColSpanControl = 1
					inp.HAlignControl = qst.HCenter
					inp.Desc = trl.S{
						"de": `Verschärfung der Schuldenbremse und Verringerung des Verschuldungsspielraums.`,
					}
				}

			}

			{
				inp := gr.AddInput()
				inp.Name = "sb_sonstiges_2"
				inp.Type = "textarea"
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.MaxChars = 155
				inp.CSSLabel = "vertical-align-top-two-rows"
				inp.Desc = trl.S{
					"de": `Sonstiges`,
				}
			}

		}

	}

	//
	// Page 5
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
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 6
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`Für wie wichtig halten Sie es, dass Schülerleistungen zwischen den Bundesländern vergleichbar sind?`),
			}

		}

		if !varianten[0] {
			{
				gr := p.AddGroup()
				gr.Cols = 2 // necessary, otherwise no vspacers
				gr.OddRowsColoring = false

				gr.Label = trl.S{
					"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 7
				}
				gr.Desc = trl.S{
					"de": `Was schätzen Sie, welcher Anteil der Bürger*innen Ihres Bundeslandes hält es für „sehr“ oder „eher“ wichtig, 
					dass Schülerleistungen zwischen den Bundesländern vergleichbar sind?`,
				}

				// p3q7
				{
					inp := gr.AddInput()
					inp.Name = "leistung_vergleichbarkeit_buerger"
					inp.Type = "number"
					inp.MaxChars = 4
					inp.Step = 1
					inp.ColSpanLabel = 0
					inp.ColSpanControl = 2
					// inp.Validator = "inRange100"
					inp.Suffix = trl.S{
						"de": `Prozent`,
					}
				}

			}
		}

	}

	//
	// Page 6
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
			gr.BottomVSpacers = 1
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %va: <br>", ctrPages.Increment()), // Frage 8a
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
				Was <b>schätzen</b> Sie, welcher Anteil der öffentlichen Finanzierung von Schulen 
				(allgemeinbildend und beruflich) kommt derzeit von den verschiedenen staatlichen Ebenen in Deutschland?
				`),
			}

			// p4q8a
			{
				inp := gr.AddInput()
				inp.Name = "tb1"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`Gemeinden`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb2"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`Bundesländer`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb3"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`Bund`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb4"
				inp.Type = "textblock"
				inp.Desc = trl.S{
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
				// inp.Validator = "inRange100"
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
				// inp.Validator = "inRange100"
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
				// inp.Validator = "inRange100"
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

			{
				inp := gr.AddInput()
				inp.Name = "tb1000"
				inp.Type = "textblock"
				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
				<br>
				<div style='display: inline-block; margin: 6px 0; font-size: 95%%'>Bitte achten Sie darauf, 
				dass Ihre Angaben insgesamt 100%% ergeben.</div>
				`),
				}
			}

		}

	}

	//
	// Page 7
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 3",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 3",
		}

		{
			gr := p.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.BottomVSpacers = 1
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %vb: <br>", ctrPages.GetLast()), // Frage 8b
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
				Und welcher Anteil der öffentlichen Finanzierung für Schulen (allgemeinbildend und beruflich) 
				<b>sollte</b> Ihrer Meinung nach von den verschiedenen staatlichen Ebenen kommen?
				
				`),
			}

			// p4q8b
			{
				inp := gr.AddInput()
				inp.Name = "tb6"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`Gemeinden`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb7"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`Bundesländer`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb8"
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`Bund`),
				}
			}
			{
				inp := gr.AddInput()
				inp.Name = "tb9"
				inp.Type = "textblock"
				inp.Desc = trl.S{
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
				// inp.Validator = "inRange100"
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
				// inp.Validator = "inRange100"
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
				// inp.Validator = "inRange100"
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
	// Page 8
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 4",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 4",
		}

		{
			gr := p.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			gr.BottomVSpacers = 1
			gr.OddRowsColoring = false

			// p4q9
			{
				inp := gr.AddInput()
				inp.Name = "mathematik_rang"
				inp.Type = "number"
				inp.MaxChars = 2
				inp.Step = 1
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				// inp.Validator = "inRange20"
				inp.Label = trl.S{
					"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 9
				}
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
					Eine aktuelle Bildungsstudie hat die Mathematikleistungen von Schüler*innen der 9. Jahrgangsstufe
					in den 16 deutschen Bundesländern verglichen. <br> 
					Was schätzen Sie, welchen Platz haben die Schüler*innen Ihres Bundeslandes belegt? <br>
					(1 ist der beste Platz, 16 der schlechteste Platz.)				
				
					`),
				}
				inp.Suffix = trl.S{
					"de": fmt.Sprintf(`-ter Platz`),
				}
			}
		}

		{
			gr := p.AddGroup()
			gr.Cols = 4 // necessary, otherwise no vspacers
			gr.BottomVSpacers = 1
			gr.OddRowsColoring = false

			// p4q10
			{
				inp := gr.AddInput()
				inp.Name = "mathematik_rang_buerger"
				inp.Type = "number"
				inp.MaxChars = 2
				inp.Step = 1
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				// inp.Validator = "inRange20"
				inp.Label = trl.S{
					"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 10
				}
				inp.Desc = trl.S{
					"de": fmt.Sprintf(`
						Was schätzen Sie, was denken die Bürger*innen Ihres Bundeslandes, welchen Platz die Schüler*innen
						ihres Bundeslandes bei der Bildungsstudie belegt haben? <br>
						(1 ist der beste Platz, 16 der schlechteste Platz.)				
					`),
				}
				inp.Suffix = trl.S{
					"de": fmt.Sprintf(`-ter Platz`),
				}
			}

		}

	}

	//
	// Page 9
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 5",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 5",
		}

		{
			names1stMatrix := []string{"schuelervergleichstest"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsFiverDafuerDagegen(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			// gr.BottomVSpacers = 1
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 11
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
				Sind Sie für oder gegen folgenden Reformvorschlag: <br>
				Es werden deutschlandweit einheitliche Schülervergleichstests in Mathematik und Deutsch in allen Schulformen eingeführt, 
				die ab der 5. Klasse alle zwei Jahre regelmäßig stattfinden. 
				Die Durchschnittsergebnisse pro Bundesland werden veröffentlicht, um die Schülerleistungen der Bundesländer miteinander zu vergleichen.<br>
				<br>
				Ich bin… 
				
				`),
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 2 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			// p3q7
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 12
			}
			gr.Desc = trl.S{
				"de": `Was schätzen Sie, welcher Anteil der Bürger*innen Ihres Bundeslandes ist „sehr“ oder „eher“ für 
					den vorher genannten Reformvorschlag, deutschlandweit einheitliche Schülervergleichstests einzuführen?`,
			}
			{
				inp := gr.AddInput()
				inp.Name = "schuelervergleichstest_buerger"
				inp.Type = "number"
				inp.MaxChars = 4
				inp.Step = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 2
				// inp.Validator = "inRange100"
				inp.Suffix = trl.S{
					"de": `Prozent`,
				}
			}

		}

	}

	//
	// Page 10
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 6",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 6",
		}

		{
			names1stMatrix := []string{"vergleichstest_regelmaessig"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsFiverDafuerDagegen(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 5 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.BottomVSpacers = 1
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 13
			}

			/*
				{
					// Dynamic header for following headless radio matrix
					gr := p.AddGroup()
					gr.BottomVSpacers = 0
					inp := gr.AddInput()
					inp.Type = "dynamic"
					inp.DynamicFunc = "HasEuroQuestion"
				}
			*/

			treatment := fmt.Sprintf(`
				<br>
				<div style='border: 1px solid grey; marging 12px 0; padding: 8px;'>
					In einer aktuellen Bildungsstudie sind die Mathematikleistungen der Schüler*innen 
					Ihres Bundeslandes in der <b>%v</b> Hälfte aller Bundesländer. 
				</div>
				`, aob)

			if !varianten[2] {
				treatment = ""
			}

			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
					Jetzt würden wir gerne noch einmal Ihre Meinung zu regelmäßigen Vergleichstests erfahren.<br>

					%v

					<br>
					Sind Sie für oder gegen vorher genannten Reformvorschlag, 
					deutschlandweit einheitliche Schülervergleichstests einzuführen?<br>
					<br>
					 Ich bin… 
				
				`, treatment),
			}

		}

	}

	//
	// Page 11
	{

		p := q.AddPage()
		p.NoNavigation = false
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "Fragen zum Bildungsföderalismus",
		}
		p.Label = trl.S{
			"de": "Teil 7",
		}
		p.Desc = trl.S{
			"de": "",
		}
		p.Short = trl.S{
			"de": "Bildung 7",
		}

		{

			rowLabels := []trl.S{}
			if !varianten[3] {
				rowLabels = []trl.S{
					{"de": "… aller befragter Bürger*innen in Deutschland."},
					{"de": "… aller befragter Bürger*innen in Ihrem Bundesland."},

					{"de": "… aller befragter Wähler*innen Ihrer Partei in Deutschland."},
					{"de": "… aller befragter Wähler*innen Ihrer Partei in Ihrem Bundesland."},

					{"de": "Ich möchte keine Informationen erhalten."},
				}

			} else {
				rowLabels = []trl.S{
					{"de": "… aller befragter Wähler*innen Ihrer Partei in Deutschland."},
					{"de": "… aller befragter Wähler*innen Ihrer Partei in Ihrem Bundesland."},

					{"de": "… aller befragter Bürger*innen in Deutschland."},
					{"de": "… aller befragter Bürger*innen in Ihrem Bundesland."},

					{"de": "Ich möchte keine Informationen erhalten."},
				}

			}

			gr := p.AddRadioGroupVertical("info_ueber_andere", rowLabels, 1)
			gr.Cols = 1 // necessary, otherwise no vspacers
			gr.OddRowsColoring = false
			gr.BottomVSpacers = 1
			gr.Label = trl.S{
				"de": fmt.Sprintf("Frage %v: <br>", ctrPages.Increment()), // Frage 14
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf(`
					Parallel zu unserer Landtagsumfrage befragen wir aktuell auch die deutsche Bevölkerung zu den gleichen Themen.<br>
					Im Folgenden bieten wir Ihnen an, dass Sie nach Abschluss der Umfragen erfahren, 
					 wie andere Befragte zum Reformvorschlag, deutschlandweit einheitliche Schülervergleichstests einzuführen, stehen.
					Wir senden Ihnen die gewählten Informationen per E-Mail zu. <br>
					<br>
					Welche der folgenden Informationen möchten Sie erhalten?<br>
					(Bitte wählen Sie <b>eine</b> der folgenden Optionen)<br>
					<br>
					Die durchschnittliche Zustimmung…
				
				`),
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 1
			gr.Width = 99
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = "11"
				inp.Label = trl.S{
					"de": "Umfrage abschließen",
				}
				inp.AccessKey = "n"
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HRight
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
