package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202602QType1(page *qst.WrappedPageT, colLabels []trl.S, inputStem string, rowLblsRandomized []trl.S) {

	// colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	colTemplateStr := "7fr       1fr 1fr 1fr 1fr 1fr   1.4fr"
	styleRowFree := "  7fr       1fr 1fr 1fr 1fr 1fr   1.4fr"

	//
	{
		gr := page.AddGroup()
		gr.Cols = 7
		gr.BottomVSpacers = 0

		// equal to below
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.Display = "grid"
		// gr.Style.Desktop.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		// gr.Style.Mobile.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

		gr.Style.Desktop.StyleText.FontSize = 90

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": "&nbsp;",
				"en": "&nbsp;",
			}
		}
		for i := 0; i < len(colLabels); i++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = colLabels[i]
			inp.LabelCenter()
			inp.LabelBottom()
		}
	}

	//
	//
	for i1 := 0; i1 < len(rowLblsRandomized); i1++ {
		gr := page.AddGroup()
		firstCol := float32(1)
		gr.Cols = firstCol + 6
		gr.RandomizationGroup = 2
		gr.BottomVSpacers = 0
		if i1 == (len(rowLblsRandomized) - 1) {
			gr.BottomVSpacers = 2 // bad, because of shuffling
			gr.BottomVSpacers = 0
		}

		// equal to above
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.Display = "grid"
		// gr.Style.Desktop.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		// gr.Style.Mobile.StyleGridContainer.TemplateColumns =  "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

		// distinct
		gr.Style.Desktop.StyleBox.Margin = "0 0 0.6rem" // bottom margin

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = firstCol
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = rowLblsRandomized[i1]
		}
		for i2 := 0; i2 < 6; i2++ {
			{
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_%v", inputStem, i1+1)
				inp.ValueRadio = fmt.Sprintf("%v", i2+1)
				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}
	}

	//
	//
	//
	//
	// row free input
	{

		gr := page.AddGroup()
		gr.Cols = 7

		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = styleRowFree
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = styleRowFree
		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

		gr.BottomVSpacers = 4

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = fmt.Sprintf("%v_free", inputStem)
			// inp.MaxChars = 17
			inp.MaxChars = 25
			inp.ColSpan = 1
			inp.ColSpanLabel = 2.4
			inp.ColSpanLabel = 0.9
			inp.ColSpanControl = 4
			inp.Label = trl.S{
				"de": "Andere",
				"en": "Other",
			}
		}

		//
		for idx := 0; idx < len(colLabels); idx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = fmt.Sprintf("%v_free_val", inputStem)
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1

			rad.ColSpanLabel = 0
			rad.ColSpanControl = 1
		}

	}

}
func special202602QType2(page *qst.WrappedPageT, inputStem string, rowLbls []trl.S) {

	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3

		for i := 0; i < len(rowLbls); i++ {
			inp1 := gr.AddInput()
			inp1.Type = "radio"
			inp1.Name = inputStem
			inp1.ValueRadio = fmt.Sprintf("%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 9
			inp1.Label = rowLbls[i]

			inp1.ControlFirst()
			// inp.LabelRight()
			// inp.LabelCenter()
			// inp.LabelBottom()

			if i == 5 {
				inp1.ColSpan = 1
				inp1.ColSpanLabel = 2
				inp1.ColSpanControl = 5

				inp1.ColSpanLabel = 4
				inp1.ColSpanControl = 2.7

				inp2 := gr.AddInput()
				inp2.Type = "number"
				inp2.Name = inputStem + "_pfc"
				inp2.Min = 0
				inp2.Max = 1000
				inp2.Step = 0.1
				inp2.Step = 0.01
				inp2.MaxChars = 4

				inp2.ColSpan = gr.Cols - 1
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1
				// inp2.Label = rowLbls[i]
				// inp2.
				inp2.Suffix = trl.S{
					"de": "Prozentpunkte",
					"en": "todo",
				}
				// inp.LabelRight()
				// inp2.ControlFirst()

			}

		}

		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "radio-xor-number"
			s1 := trl.S{
				"de": "unused",
				"en": "unused",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}
			inp.JSBlockStrings = map[string]string{
				"inp1":    inputStem,
				"inp2":    inputStem + "_pfc",
				"radioOn": inputStem + "6",
			}
		}

	}
}

func special202602(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2026 && q.Survey.Month == 2
	if !cond {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("72rem")
	page.WidthMax("64rem")

	page.Label = trl.S{
		"de": "Sonderfragen: Kurz- und mittelfristiges Wirtschaftswachstum - extra",
		"en": "Special: Short- and Medium-Term Economic Growth - extra",
	}
	page.Short = trl.S{
		"de": "Wirtschafts-<br>wachstum - extra",
		"en": "Economic<br>Growth - extra",
	}
	// page.WidthMax("42rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Für ein Quartal (z.B. 3. Quartal 2026) geben Teilnehmende ihre Einschätzung zum Wirtschaftswachstum in mehreren FMT-Befragungswellen ab. Dabei sind nicht nur die Prognosen, sondern auch deren Revisionen von großer Bedeutung. Wir sprechen dabei von einer Revision, wenn die Wirtschaftswachstumsprognose für ein gegebenes Quartal von der Prognose aus der vorhergehenden Umfragewelle abweicht.
				`,
				"en": `
					todo
				`,
			}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
				Wenn Sie Ihre Wirtschaftswachstumsprognose für ein gegebenes Quartal revidieren: Wie wichtig sind <i>typischerweise</i> die folgenden Faktoren für Ihre Revision?
			`,
				"en": `
				todo
			`,
			}.Outline("3.")
		}
	}

	//
	//
	colLabelsSsq3 := []trl.S{
		{
			"de": "sehr wichtig",
			"en": "very important",
		},
		{
			"de": "eher wichtig",
			"en": "important",
		},
		{
			"de": "weder noch",
			"en": "undecided",
		},
		{
			"de": "eher unwichtig",
			"en": "rather unimportant",
		},
		{
			"de": "sehr unwichtig",
			"en": "very unimportant",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}
	lblsSsq3 := []trl.S{
		{
			"de": `Neue öffentliche Wirtschaftsdaten (z.B. die Inflationsrate oder Industrieproduktion)`,
			"en": "todo",
		},
		{
			"de": `Neue nichtöffentliche Daten (z.B. Verkaufszahlen Ihres Unternehmens oder andere Interna)`,
			"en": "todo",
		},
		{
			"de": `Neue wirtschaftspolitische Maßnahmen (z.B. eine Leitzinsänderung)`,
			"en": "todo",
		},
		{
			"de": `Änderungen der durchschnittlichen Prognose aller Befragten (der Konsensusprognose)`,
			"en": "todo",
		},
	}
	special202602QType1(qst.WrapPageT(page), colLabelsSsq3, "ssq3", lblsSsq3)

	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
				Wenn Sie Ihre Wirtschaftswachstumsprognose für ein gegebenes Quartal nicht revidieren: Welche der folgenden Gründe für Ihre Entscheidung, keine Revision vorzunehmen, treffen <i>typischerweise</i> zu?
			`,
				"en": `
				todo
			`,
			}.Outline("4.")
		}
	}
	colLabelsSsq4 := []trl.S{
		{
			"de": "trifft zu",
			"en": "applies",
		},
		{
			"de": "trifft eher zu",
			"en": "rather applies",
		},
		{
			"de": "weder noch",
			"en": "undecided",
		},
		{
			"de": "trifft eher nicht zu",
			"en": "rather not applies",
		},
		{
			"de": "trifft nicht zu",
			"en": "not applies",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}
	lblsSsq4 := []trl.S{
		{
			"de": `Der Zeitaufwand bzw. die Kosten der Prognoseanpassung sind zu hoch, wenn es nur geringfügige Änderungen in den Wirtschaftsaussichten gibt`,
			"en": "todo",
		},
		{
			"de": `Es gibt keine neuen Informationen, die eine Änderung der Wirtschaftsaussichten nahelegen`,
			"en": "todo",
		},
		{
			"de": `Meine Prognose passe ich nur dann an, wenn sie entscheidungsrelevant ist`,
			"en": "todo",
		},
		{
			"de": `Meine Prognose ist glaubwürdiger, wenn sie seltener angepasst wird`,
			"en": "todo",
		},
		{
			"de": `Es gibt keine Änderung der durchschnittlichen Prognose aller Befragten (der Konsensusprognose)`,
			"en": "todo",
		},
	}
	special202602QType1(qst.WrapPageT(page), colLabelsSsq4, "ssq4", lblsSsq4)

	//
	//
	page = q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("54rem")
	page.WidthMax("48rem")
	page.WidthMax("44rem")

	page.Label = trl.S{
		"de": "Sonderfragen: Kurz- und mittelfristiges Wirtschaftswachstum - extra",
		"en": "Special: Short- and Medium-Term Economic Growth - extra",
	}
	page.Short = trl.S{
		"de": "Wirtschafts-<br>wachstum - extra",
		"en": "Economic<br>Growth - extra",
	}
	page.SuppressInProgressbar = true
	// page.WidthMax("42rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
				Wie groß muss typischerweise die Änderung der Wirtschaftsaussichten für ein Quartal sein, damit Sie Ihre Prognose in der Befragung tatsächlich anpassen?
				<br>
				<br>
				Die Änderung der Wirtschaftsaussichten sollte so groß sein, dass sich meine Prognose mindestens ändert um:

			`,
				"en": `
				todo
			`,
			}.Outline("5.")
		}
	}

	lblsSsq5 := []trl.S{
		{
			"de": `±0,10 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,20 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,30 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,40 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,50 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±`,
			"en": "todo",
		},
		{
			"de": `Ich passe meine Prognose immer an, wenn sich die Wirtschaftsaussichten ändern`,
			"en": "todo",
		},
	}
	special202602QType2(qst.WrapPageT(page), "ssq5", lblsSsq5)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
				Wie groß muss die Änderung des Konsensus (der durchschnittlichen Prognose aller Befragten) sein, damit Sie Ihre Prognose in der Befragung typischerweise anpassen?
				<br>
				<br>
				Der Konsensus sollte sich mindestens ändern um:

			`,
				"en": `
				todo
			`,
			}.Outline("6.")
		}
	}
	lblsSsq6 := []trl.S{
		{
			"de": `±0,10 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,20 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,30 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,40 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `±0,50 Prozentpunkte `,
			"en": "todo",
		},
		{
			"de": `± `,
			"en": "todo",
		},
		{
			"de": `Ich passe meine Prognose immer an, wenn sich der Konsensus ändert `,
			"en": "todo",
		},
		{
			"de": `Der Konsensus ist für meine Prognose irrelevant`,
			"en": "todo",
		},
	}
	special202602QType2(qst.WrapPageT(page), "ssq6", lblsSsq6)

	return nil

}
