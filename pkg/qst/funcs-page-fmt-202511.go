package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func fmt202511_0(q *QuestionnaireT, page *pageT) error {
	return fmt202511(q, page, 0)
}
func fmt202511_1(q *QuestionnaireT, page *pageT) error {
	return fmt202511(q, page, 1)
}

func fmt202511(q *QuestionnaireT, page *pageT, instance int) error {

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
					stellen wir Ihnen drei Szenarien vor.
					<br>
					<br>

					Im ersten Szenario bitten wir Sie, einen Wert in der Graphik <i>abzulesen</i>. 
					Sie müssen keinen der Werte anpassen. 
					
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
					Szenario 1: Frau Blau möchte über die Projektlaufzeit 
					jedes Jahr 100&nbsp;Bäume pflanzen. 
					Sie wählt einen Anteil von 50% an Baumart&nbsp;2 aus.

					Wie hoch ist ihr prognostizierter Ertrag 
					<u><b>in den besten 5 von 100&nbsp;Fällen</b></u>?
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
						stellen wir Ihnen drei Szenarien vor.
						<br>
						<br>
	
						Im ersten Szenario bitten wir Sie, einen Wert in der Graphik <i>abzulesen</i>. 
						Sie müssen keinen der Werte anpassen. 	
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
						Szenario 1: Frau Blau möchte über einen Zeitraum von 20&nbsp;Jahren 
						einen monatlichen Sparbetrag von 100&nbsp;Euro anlegen. 
						Sie wählt einen Aktienanteil von 50%.

						Wie hoch ist ihr prognostiziertes Vermögen  
						<u><b>in den besten 5 von 100&nbsp;Fällen</b></u>?
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
