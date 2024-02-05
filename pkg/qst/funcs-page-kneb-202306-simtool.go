package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func kneb202306simtool0(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 0)
}
func kneb202306simtool1(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 1)
}
func kneb202306simtool2(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 2)
}
func kneb202306simtool3(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 3)
}
func kneb202306simtool4(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 4)
}

// Dimension-1
// The same page is shown in two variations depending on version/group index.
// The variations are neutral frame and financal frame (nf, ff).
// Distinction is only in the name JS config.
// Dimension-2
// The same page is shown several times (instance).
// Validations differ.
func kneb202306simtool(q *QuestionnaireT, page *pageT, instance int) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"en": "Experiment-Chart",
		"de": "Experiment chart",
	}
	page.Label = trl.S{
		"en": "",
		"de": "",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("58rem")

	//
	// kneb_simtool_inst_0
	// kneb_simtool_inst_4
	valFunc := fmt.Sprintf("kneb_simtool_inst_%v", instance)
	if instance != 0 && instance != 4 {
		valFunc = ""
	}

	btmSpacers := 0

	grIdx := q.Version() % 2

	if instance == 1 {
		if grIdx == 0 {
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = btmSpacers
			{

				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					Um Sie beim Kennenlernen der Graphik weiter zu unterstützen, 
					haben wir drei  Beispiele vorbereitet.
					<br>
					<br>

					In Beispiel 1 bitten wir Sie, einen Wert in der Graphik <i>abzulesen</i>. 
					<br>

					
					Sie müssen dafür keinen der Werte anpassen. 
					
					Bitte tragen Sie Ihre Antwort in das Antwortfeld ein.
				`,
					"en": `todo`,
				}.OutlineHid("C24.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must;kneb_qc24_nf"
				inp.Validator = "must"
				inp.Name = "qc24_nf_return"

				inp.Min = 0
				// 20.900 Tonnen
				inp.Max = 280 * 1000
				inp.MaxChars = 6
				inp.Suffix = trl.S{
					"de": `€`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 2
				inp.LabelPadRight()
				inp.ControlBottom()
				inp.Label = trl.S{
					"de": `		
					Beispiel 1: Frau Blau möchte über die Projektlaufzeit 
					jedes Jahr 100&nbsp;Bäume pflanzen. 
					Sie wählt einen Anteil von 50% an Baumart&nbsp;2 aus.

					Wie hoch ist ihr prognostizierter Ertrag 
					<u><b>in den besten 5 von 100 Fällen</b></u>?
					`,
					"en": `todo`,
				}

			}
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebQuiz" // js filename
			}

		} else {
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = btmSpacers
			{

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": `
						Um Sie beim Kennenlernen der Graphik weiter zu unterstützen, 
						haben wir drei  Beispiele vorbereitet.
						<br>
						<br>
	
						In Beispiel 1 bitten wir Sie, einen Wert in der Graphik <i>abzulesen</i>. 
						<br>
	
						
						Sie müssen dafür keinen der Werte anpassen. 
						
						Bitte tragen Sie Ihre Antwort in das Antwortfeld ein.
					`,
						"en": `todo`,
					}.OutlineHid("C24.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1

				}
				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Validator = "must;kneb_qc24_ff"
					inp.Validator = "must"
					inp.Name = "qc24_ff_return"

					inp.Min = 0
					// 104.700 Euro
					inp.Max = 100 * 1000 * 1000
					inp.MaxChars = 6
					inp.Suffix = trl.S{
						"de": `€`,
						"en": `todo`,
					}
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 4
					inp.ColSpanControl = 2
					inp.LabelPadRight()
					inp.ControlBottom()
					inp.Label = trl.S{
						"de": `					
						Beispiel 1: Frau Blau möchte über einen Zeitraum von 20&nbsp;Jahren 
						einen monatlichen Sparbetrag von 100&nbsp;Euro anlegen. 
						Sie wählt einen Aktienanteil von 50%.

						Wie hoch ist ihr prognostiziertes Vermögen  
						<u><b>in den besten 5 von 100 Fällen</b></u>?
						`,
						"en": `todo`,
					}
				}

			}
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebQuiz" // js filename
			}

		}
	}

	if instance == 2 {
		if grIdx == 0 {
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = btmSpacers
			{

				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					In Beispiel 2 bitten wir Sie den Anteil von Baumart&nbsp;2 anzupassen. 
					Lesen Sie anschließend den durchschnittlichen Ertrag aus der Graphik ab 
					und tragen Sie Ihre Antwort in das Antwortfeld ein.
					`,
					"en": `todo`,
				}.OutlineHid("C25.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must;kneb_qc24_nf"
				inp.Validator = "must"
				inp.Name = "qc25_nf_return"

				inp.Min = 0
				// 20.900 Tonnen
				inp.Max = 280 * 1000
				inp.MaxChars = 6
				inp.Suffix = trl.S{
					"de": `€`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 2
				inp.LabelPadRight()
				inp.ControlBottom()
				inp.Label = trl.S{
					"de": `		
					Beispiel 2: Herr Lila kann jedes Jahr 100 Bäume pflanzen. 
					Er wählt einen Anteil von 70% an Baumart&nbsp;2 aus.
					Wie hoch ist sein prognostizierter Ertrag am Ende des Projekts 
					<u><b>im Durchschnitt</b></u>?
					`,
					"en": `todo`,
				}

			}
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebQuiz" // js filename
			}

		} else {
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = btmSpacers
			{

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": `
						In Beispiel 2 bitten wir Sie den Aktienanteil anzupassen. 
						Lesen Sie anschließend den durchschnittlichen Ertrag 
						aus der Graphik ab und tragen Sie Ihre Antwort in das Antwortfeld ein.
						`,
						"en": `todo`,
					}.OutlineHid("C25.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1

				}
				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Validator = "must;kneb_qc24_ff"
					inp.Validator = "must"
					inp.Name = "qc25_ff_return"

					inp.Min = 0
					// 104.700 Euro
					inp.Max = 100 * 1000 * 1000
					inp.MaxChars = 6
					inp.Suffix = trl.S{
						"de": `€`,
						"en": `todo`,
					}
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 4
					inp.ColSpanControl = 2
					inp.LabelPadRight()
					inp.ControlBottom()
					inp.Label = trl.S{
						"de": `		
						Beispiel 2: Herr Lila kann im Monat 100&nbsp;Euro zur Seite legen. 
						Er wählt einen Aktienanteil von 70% aus.

						Wie hoch ist sein prognostiziertes Vermögen nach 20&nbsp;Jahren 
						<u><b>im Durchschnitt</b></u>?
						`,
						"en": `todo`,
					}
				}

			}
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebQuiz" // js filename
			}

		}
	}

	if instance == 3 {
		if grIdx == 0 {
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = btmSpacers
			{

				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					In Beispiel 3 bitten wir Sie die Anzahl der Bäume anzupassen. 
					Bitte achten Sie dabei auf die Veränderungen in der Graphik. 
					Tragen Sie Ihre Antwort in das Antwortfeld ein.
					`,
					"en": `todo`,
				}.OutlineHid("C26.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must;kneb_qc24_nf"
				inp.Validator = "must"
				inp.Name = "qc26_nf_return"

				inp.Min = 0
				// 20.900 Tonnen
				inp.Max = 280 * 1000
				inp.MaxChars = 6
				inp.Suffix = trl.S{
					"de": `Bäume`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 2
				inp.LabelPadRight()
				inp.ControlBottom()
				inp.Label = trl.S{
					"de": `		
					Beispiel 3: Frau Gelb ist bereit einen Anteil von 50% 
					an Baumart&nbsp;2 in ihrem Wald zu pflanzen. 
					Sie möchte während des Projekts 
					<u><b>in den schlechtesten 5 von 100 Fällen</b></u>
					einen Ertrag von 30.500&nbsp;Euro erzielen. 
					
					
					Wie viele Bäume muss sie jährlich pflanzen, 
					damit ihr dies 
					<u><b>in den schlechtesten 5 von 100 Fällen</b></u>
					gelingen kann?
					`,
					"en": `todo`,
				}

			}
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebQuiz" // js filename
			}

		} else {
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = btmSpacers
			{

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": `
						In Beispiel 3 bitten wir Sie den monatlichen Sparbetrag anzupassen. 
						Bitte achten Sie dabei auf die Veränderungen in der Graphik. 
						Tragen Sie Ihre Antwort in das Antwortfeld ein.
						`,
						"en": `todo`,
					}.OutlineHid("C26.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1

				}
				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Validator = "must;kneb_qc24_ff"
					inp.Validator = "must"
					inp.Name = "qc26_ff_return"

					inp.Min = 0
					// 104.700 Euro
					inp.Max = 100 * 1000 * 1000
					inp.MaxChars = 6
					inp.Suffix = trl.S{
						"de": `€`,
						"en": `todo`,
					}
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 4
					inp.ColSpanControl = 2
					inp.LabelPadRight()
					inp.ControlBottom()
					inp.Label = trl.S{
						"de": `
						Beispiel 3: Frau Gelb ist bereit einen Aktienanteil von 50% 
						in ihrem Portfolio zu akzeptieren. 
						Sie möchte in 20&nbsp;Jahren 
						<b><u>in den schlechtesten 5 von 100 Fällen</u></b> 
						ein Vermögen von 30.500&nbsp;Euro aufbauen. 
						
						Wie viel Euro muss sie monatlich sparen, 
						damit ihr dies 
						<b><u>in den schlechtesten 5 von 100 Fällen</u></b> 
						 gelingen kann?

						`,
						"en": `todo`,
					}
				}

			}
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebQuiz" // js filename
			}

		}
	}

	// gr1 - hidden inputs saving the values from the echart
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = btmSpacers
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			// inp.Name = "simtool_instance"
			// we render this as textblock, to circumvent the uniqueness of fieldnames.
			// it only serves as read-only indicator for the embedded echart javascript
			inp.Label = trl.S{
				"de": fmt.Sprintf(`
					<div style='visible: none; height: 0.1rem; font-size: 4px;'> 
						<input type='hidden' value='%v'  name='simtool_instance' id='simtool_instance' />
					</div> 
				`, instance),
				"en": `todo`,
			}
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Validator = valFunc
			inp.Name = fmt.Sprintf("sparbetrag_bg_%v", instance)
		}

		// inversely related
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Validator = valFunc
			inp.Name = fmt.Sprintf("share_safe_bg_%v", instance)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Validator = valFunc
			inp.Name = fmt.Sprintf("share_risky_bg_%v", instance)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Validator = valFunc
			inp.Name = fmt.Sprintf("sim_history_%v", instance)
		}
	}

	// gr2 - echart file embedding
	// group index determines nf or ff
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = btmSpacers
		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.DynamicFunc = "RenderStaticContent"
			inp.DynamicFuncParamset = fmt.Sprintf("./echart/inner-%d.html", grIdx)
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

		//
		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "knebRenameNext" // js filename

			s1 := trl.S{
				"de": "Weiter",
				"en": "todo",
			}
			if instance == 4 {
				s1 = trl.S{
					"de": "Werte speichern und weiter",
					"en": "todo",
				}
			}

			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}

			inp.JSBlockStrings = map[string]string{}
			// inp.JSBlockStrings["pageID"] = fmt.Sprintf("pg%02v", len(q.Pages)-1)

		}
	}

	return nil
}
