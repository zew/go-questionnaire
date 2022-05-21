package biii

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func finish(q *qst.QuestionnaireT) {
	//
	// Finish questionnaire?  - one before last page
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("36rem")

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
					"de": "Wollen Sie uns noch etwas mitteilen?",
					"en": "Any remarks or advice for us?",
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "comment"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = qst.RemainOpen
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Zugang bleibt bestehen.  Daten können in weiteren Sitzungen geändert/ergänzt werden. ",
					"en": "Leave the questionnaire open. Data  can be changed/completed&nbsp;in later sessions.     ",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = qst.Finished
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Fragebogen ist abgeschlossen und kann nicht mehr geöffnet werden. ",
					"en": "The questionnaire is finished. No more edits.                         ",
				}
			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Label = trl.S{
					"de": "Durch Klicken erhalten Sie eine Zusammenfassung Ihrer Antworten",
					"en": "By clicking, you will receive a summary of your answers.",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

		// pge.ExampleSixColumnsLabelRight()

	}

	// Report of results
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Ihre Eingaben", "en": "Summary of results"}
		page.NoNavigation = true
		page.SuppressProgressbar = true
		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "ResponseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.Param = "site-imprint.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}
	}

}
