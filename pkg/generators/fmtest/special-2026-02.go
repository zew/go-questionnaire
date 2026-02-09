package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202602QType1(page *qst.WrappedPageT, colLabels []trl.S, inputStem string, rowLblsRandomized []trl.S, randGroup int) {

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
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

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
		gr.RandomizationGroup = randGroup
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
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

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
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

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
					"en": "Percentage points",
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
		"de": "Sonderfragen: Kurz- und mittelfristiges Wirtschaftswachstum - Prognoserevisionen",
		"en": "Special: Short- and Medium-Term Economic Growth - Revisions",
	}
	page.Short = trl.S{
		"de": "Wirtschafts-<br>wachstum - Revisionen",
		"en": "Economic<br>Growth - Revisions",
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
					For some quarter (e.g., Q3 2026), participants submit their assessment of economic growth across several FMS survey waves. In this context, not only the forecasts but also their revisions are of great importance. We refer to a revision when the economic growth forecast for a given quarter deviates from the forecast in the preceding survey wave.
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
				When you revise your economic growth forecast for a given quarter: How important are the following factors <i>typically</i> for your revision?
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
			"en": "somewhat important",
		},
		{
			"de": "weder noch",
			"en": "undecided",
		},
		{
			"de": "eher unwichtig",
			"en": "somewhat unimportant",
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
			"en": "New public economic data (e.g., the inflation rate or industrial production",
		},
		{
			"de": `Neue nichtöffentliche Daten (z.B. Verkaufszahlen Ihres Unternehmens oder andere Interna)`,
			"en": "New non-public data (e.g., sales of your company or other internal information)",
		},
		{
			"de": `Neue wirtschaftspolitische Maßnahmen (z.B. eine Leitzinsänderung)`,
			"en": "New economic policy measures (e.g., a change of the monetary policy interest rate)",
		},
		{
			"de": `Änderungen der durchschnittlichen Prognose aller Befragten (der Konsensusprognose)`,
			"en": "Changes of the average forecast among all participants (the consensus forecast)",
		},
	}
	special202602QType1(qst.WrapPageT(page), colLabelsSsq3, "ssq3", lblsSsq3, 2)

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
				When you do not revise your economic growth forecast for a given quarter: Which of the following reasons for your decision not to make a revision <i>typically</i> apply?
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
			"en": "tends to apply",
		},
		{
			"de": "weder noch",
			"en": "undecided",
		},
		{
			"de": "trifft eher nicht zu",
			"en": "tends not to apply",
		},
		{
			"de": "trifft nicht zu",
			"en": "does not apply",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}
	lblsSsq4 := []trl.S{
		{
			"de": `Der Zeitaufwand bzw. die Kosten der Prognoseanpassung sind zu hoch, wenn es nur geringfügige Änderungen in den Wirtschaftsaussichten gibt`,
			"en": "The workload or the costs of the forecast adjustment are too high if there are only minor changes in the economic outlook",
		},
		{
			"de": `Es gibt keine neuen Informationen, die eine Änderung der Wirtschaftsaussichten nahelegen`,
			"en": "There is no new information that suggests a change in the economic outlook",
		},
		{
			"de": `Meine Prognose passe ich nur dann an, wenn sie entscheidungsrelevant ist`,
			"en": "I change my forecast only if it is decision-relevant",
		},
		{
			"de": `Meine Prognose ist glaubwürdiger, wenn sie seltener angepasst wird`,
			"en": "My forecast is more credible if it is adjusted less frequently",
		},
		{
			"de": `Es gibt keine Änderung der durchschnittlichen Prognose aller Befragten (der Konsensusprognose)`,
			"en": "There is no change in the average forecast among all participants (the consensus forecast)",
		},
	}
	special202602QType1(qst.WrapPageT(page), colLabelsSsq4, "ssq4", lblsSsq4, 3)

	//
	//
	page = q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("54rem")
	page.WidthMax("48rem")
	page.WidthMax("44rem")

	page.Label = trl.S{
		"de": "Sonderfragen: Kurz- und mittelfristiges Wirtschaftswachstum - Prognoserevisionen",
		"en": "Special: Short- and Medium-Term Economic Growth - Revisions",
	}
	page.Short = trl.S{
		"de": "Wirtschafts-<br>wachstum - Prognose",
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
				How large must the change in the economic outlook be for you to typically adjust your forecast in the survey?
				<br>
				The change in the economic outlook should be so large that my forecast changes at least by:
			`,
			}.Outline("5.")
		}
	}

	lblsSsq5 := []trl.S{
		{
			"de": `±0,10 Prozentpunkte `,
			"en": "±0,10 Percentage points",
		},
		{
			"de": `±0,20 Prozentpunkte `,
			"en": "±0,20 Percentage points",
		},
		{
			"de": `±0,30 Prozentpunkte `,
			"en": "±0,30 Percentage points",
		},
		{
			"de": `±0,40 Prozentpunkte `,
			"en": "±0,40 Percentage points",
		},
		{
			"de": `±0,50 Prozentpunkte `,
			"en": "±0,50 Percentage points",
		},
		{
			"de": `±`,
			"en": "±",
		},
		{
			"de": `Ich passe meine Prognose immer an, wenn sich die Wirtschaftsaussichten ändern`,
			"en": "I always adjust my forecast when the economic outlook changes",
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
				How large must the change in the consensus forecast be for you to typically adjust your forecast in the survey?
				<br>
				The consensus should change by at least how many percentage points:
			`,
			}.Outline("6.")
		}
	}
	lblsSsq6 := []trl.S{
		{
			"de": `±0,10 Prozentpunkte `,
			"en": "±0,10 Percentage points",
		},
		{
			"de": `±0,20 Prozentpunkte `,
			"en": "±0,20 Percentage points",
		},
		{
			"de": `±0,30 Prozentpunkte `,
			"en": "±0,30 Percentage points",
		},
		{
			"de": `±0,40 Prozentpunkte `,
			"en": "±0,40 Percentage points",
		},
		{
			"de": `±0,50 Prozentpunkte `,
			"en": "±0,50 Percentage points",
		},
		{
			"de": `±`,
			"en": "±",
		},
		{
			"de": `Ich passe meine Prognose immer an, wenn sich der Konsensus ändert `,
			"en": "I always adjust my forecast when the consensus changes",
		},
		{
			"de": `Der Konsensus ist für meine Prognose irrelevant`,
			"en": "The consensus is not relevant for my own forecast",
		},
	}
	special202602QType2(qst.WrapPageT(page), "ssq6", lblsSsq6)

	return nil

}
