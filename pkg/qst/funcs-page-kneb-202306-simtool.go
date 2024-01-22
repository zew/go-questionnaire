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

					Im Folgenden stellen wir Ihnen Fragen, 
					die Sie mit Hilfe der interaktiven Graphik beantworten können.
					<br>
					<br>

					Frau Blau möchte über die Projektlaufzeit jedes Jahr 
					100&nbsp;Bäume pflanzen. 
					
					Sie wählt einen Anteil von 60% an Baumart&nbsp;2 aus. 
					
					Wie hoch ist ihr prognostizierter Ertrag  in den  
					<i><u>besten 5 von 100&nbsp;Fällen</u></i>?	
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
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 4
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

						Im Folgenden stellen wir Ihnen Fragen, 
						die Sie mit Hilfe der interaktiven Graphik beantworten können.
						<br>
						<br>
	

						Frau Blau möchte über einen Zeitraum von 20&nbsp;Jahren 
						einen monatlichen Sparbetrag von 100&nbsp;Euro anlegen. 
						
						Sie wählt einen Aktienanteil von 60%.

						Wie hoch ist ihr prognostiziertes 
						Vermögen in den <i><u>besten 5 von 100&nbsp;Fällen</u></i>?
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
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 4
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
			gr.BottomVSpacers = btmSpacers
			gr.Cols = 4
			var radioValues = []string{
				"0pct",
				"10pct",
				"20pct",
				"30pct",
				"40pct",
				"50pct",
				"60pct",
				"70pct",
				"80pct",
				"90pct",
				"100pct",
			}
			var labels = []trl.S{
				{
					"de": "0% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "10% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "20% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "30% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "40% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "50% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "60% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "70% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "80% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "90% Anteil&nbsp;2",
					"en": "todo",
				},
				{
					"de": "100% Anteil&nbsp;2",
					"en": "todo",
				},
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `

						Herr Lila kann jedes Jahr 100&nbsp;Bäume pflanzen. 
						
						Er möchte am Ende des Projekts einen Ertrag 
						von ca. 42.000&nbsp;Euro (<u>im Durchschnitt</u>) erzielen.

						<br>
						Welchen Anteil an Baumart&nbsp;2 
						muss sein Waldstück mindestens haben, damit ihm das gelingt? 

					`,
					"en": `todo`,
				}.OutlineHid("C25.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}

			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Validator = "must;kneb_qc25_nf"
				rad.Validator = "must"
				rad.Name = "qc25_share_nf"

				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols / 3
				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
			}

		} else {

			gr := page.AddGroup()
			gr.BottomVSpacers = btmSpacers
			gr.Cols = 4
			var radioValues = []string{
				"0pct",
				"10pct",
				"20pct",
				"30pct",
				"40pct",
				"50pct",
				"60pct",
				"70pct",
				"80pct",
				"90pct",
				"100pct",
			}
			var labels = []trl.S{
				{
					"de": "0% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "10% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "20% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "30% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "40% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "50% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "60% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "70% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "80% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "90% Aktien&shy;anteil",
					"en": "todo",
				},
				{
					"de": "100% Aktien&shy;anteil",
					"en": "todo",
				},
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Herr Lila kann im Monat 100 Euro zur Seite legen. 
						
						Wenn er nach 20 Jahren ca. 34.000 Euro angespart haben möchte (<u>im Durchschnitt</u>), 
						welchen Aktienanteil sollte sein Portfolio mindestens haben?
					`,
					"en": `todo`,
				}.OutlineHid("C25.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}

			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Validator = "must;kneb_qc25_ff"
				rad.Validator = "must"
				rad.Name = "qc25_share_ff"

				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols / 3
				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
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
						Frau Gelb möchte während des Projekts einen Ertrag von 40.000&nbsp;Euro erzielen. 
												
						Sie ist bereit einen Anteil von 50% an Baumart&nbsp;2 in ihrem Wald zu akzeptieren.
			
						Wie viele Bäume muss sie jährlich pflanzen, 
						damit ihr dies 
						<i><u>in den schlechtesten 5 von 100&nbsp;Fällen</u></i> 
						gelingen kann?
					
					`,
					"en": `todo`,
				}.OutlineHid("C26.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must;kneb_qc26_nf"
				inp.Validator = "must"
				inp.Name = "qc26_area_nf"

				inp.Min = 0
				// 40 Hektar
				inp.Max = 700
				inp.Step = 10
				inp.Placeholder = trl.S{"de": "#", "en": "#"}
				inp.MaxChars = 6
				inp.Suffix = trl.S{
					"de": `Bäume`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 4
			}
			/*
				{
					inp := gr.AddInput()
					inp.ColSpanControl = 1
					inp.Type = "javascript-block"
					inp.Name = "knebVisiblePrev" // js filename
				}
			*/
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
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Frau Gelb möchte in 20 Jahren ein Vermögen von 40.000 Euro aufbauen. 

						Sie ist bereit einen Aktienanteil von 50% in ihrem Portfolio zu akzeptieren.

						Wie viele Euro muss sie monatlich sparen, 
						damit ihr dies 
						<i><u>in den schlechtesten 5 von 100&nbsp;Fällen</u></i> 
						gelingen kann?
		
					`,
					"en": `todo`,
				}.OutlineHid("C26.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must;kneb_qc26_ff"
				inp.Validator = "must"
				inp.Name = "qc26_area_ff"

				inp.Min = 0
				// 310€
				inp.Max = 200 * 1000
				// inp.Step = 10
				inp.Placeholder = trl.S{"de": "#", "en": "#"}
				inp.MaxChars = 6
				inp.Suffix = trl.S{
					"de": `€`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 4
			}
			/*
				{
					inp := gr.AddInput()
					inp.ColSpanControl = 1
					inp.Type = "javascript-block"
					inp.Name = "knebVisiblePrev" // js filename
				}
			*/
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
