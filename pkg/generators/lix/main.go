package lix

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create questionnaire
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"de"} // governs default language code
	// q.LangCode = "en"

	q.Survey.Org = trl.S{
		"en": "ZEW",
		"de": "ZEW",
	}
	q.Survey.Name = trl.S{
		"en": "Länderindex Umfrage",
		"de": "Länderindex Umfrage",
	}
	// q.Variations = 1

	// page0
	{
		page := q.AddPage()
		page.ValidationFuncName = ""

		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true

		page.Label = trl.S{
			"en": "Dear Madam / Sir,",
			"de": "Sehr geehrter Damen und Herren",
		}
		// page.Short = trl.S{
		// 	"en": "Greeting",
		// 	"de": "Begrüßung",
		// }

		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./main.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page6 - finish
	{
		page := q.AddPage()
		page.Label = trl.S{
			"en": "Consent",
			"de": "Abschluss<br><br>",
		}
		page.Short = trl.S{
			"en": "Finish",
			"de": "DSGVO",
		}
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		// gr2
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
				inp.Name = "q63_comment"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
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
				inp.Label = cfg.Get().Mp["finish_save_questionnaire"]
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

	//
	//
	// Report of results
	{
		page := q.AddPage()
		page.NoNavigation = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		page.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": `&nbsp;`,
					"de": `&nbsp;`,
				}
			}
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.DynamicFunc = "RepsonseStatistics"
			// }
		}
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	q.SetColspans()

	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
