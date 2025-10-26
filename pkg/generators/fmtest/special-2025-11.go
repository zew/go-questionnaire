package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202511(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2025 && q.Survey.Month == 11
	if !cond {
		return nil
	}

	// this is page I
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "&nbsp;",
			"en": "&nbsp;",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			// "de": "Wachstums&shy;chancen I",
			"de": "Wachstums-<br>chancen",
			"en": "todo",
		}
		page.WidthMax("56rem")

		//
		{
			gr := page.AddGroup()
			gr.Cols = 12

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Label = trl.S{
					"de": `
						<p style='font-size:120%; font-weight:bold; margin-top: 0;'>
							Ihre Sicht auf Wachstumschancen für Deutschland 
						</p>
											
						<p>
						Die deutsche Wirtschaft leidet seit Jahren unter niedrigem Wachstum.
						Um die Aussichten für Deutschland besser zu verstehen, sind wir auf Ihre Expertise angewiesen.
						</p>

						<p>
						Da Expertenprognosen über das Wirtschaftswachstum häufig variieren,
						möchten wir <!-- zunächst --> Ihre Einschätzung hinsichtlich der anderen Umfrageteilnehmer erfahren.					
						</p>

						<br>
						<br>

				`,
					"en": `
					todo
				`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.Label = trl.S{
					"de": `
						Was glauben Sie: Wie hoch ist die durchschnittliche Wirtschaftswachstumsprognose 
						<br>
						<i>unter allen Befragten in der aktuellen Befragung</i>?
							
						<small>
						Bitte nicht-annualisiertes Quartalswachstum des realen & saisonbereinigten BIP angeben.
						</small>

				`,
					"en": `
					todo
				`,
				}.Outline("3a.")
			}

			sLbl1 := css.NewStylesResponsive(nil)
			sLbl1.Desktop.StyleGridItem.JustifySelf = "end"
			sLbl1.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
			sLbl1.Mobile.StyleBox.Padding = " 0 2.7rem 0 0.2rem"

			sLbl2 := *sLbl1
			sLbl2.Mobile.StyleGridItem.JustifySelf = "start"
			sLbl2.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
			sLbl2.Mobile.StyleBox.Padding = " 0 1.5rem 0 0.2rem"

			// row 1 - four quarters - label
			//   removed
			// row 2 - four quarters - inputs
			for i := 0; i < 4; i++ {
				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("ssq3a_%v", i+1)
					inp.ColSpan = 3
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 1
					inp.Min = -20
					inp.Max = 20
					inp.Step = 0.01
					inp.MaxChars = 4
					inp.Label = trl.S{
						"de": q.Survey.Quarter(i),
						"en": q.Survey.Quarter(i),
					}

					inp.Suffix = trl.S{
						"de": "%",
						"en": "pct",
					}
					inp.StyleLbl = sLbl1

					inp.Style = css.MobileVertical(inp.Style)
					inp.StyleLbl.Mobile.StyleGridItem.JustifySelf = "start"
					// inp.StyleLbl.Mobile.StyleGridItem.AlignSelf = "end"
				}
			}

		}

		qst.ChangeHistoryJS(q, page)

	}

	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511Pg2"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511Pg3"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511Pg4"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511Pg5"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511Pg6"
	}

	return nil

}
