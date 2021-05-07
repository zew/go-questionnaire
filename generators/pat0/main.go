package pat0

import (
	"fmt"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/qst"
)

// Create for PAT but with zentrally distributed versions.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.3rem 0 0"

	q, err := pat.Create(params)
	if err != nil {
		return q, err
	}

	q.AssignVersion = "round-robin"
	q.VersionEffective = -2 // must be re-set at the end - after validate

	q.RemoveGroup(8, 3)

	pat.AddPersonalQuestions(q, 8)

	{

		l := len(q.Pages)
		page := q.EditPage(l - 1 - 1)

		//
		// finish button
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			gr.Cols = 2
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "finished"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 2
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"

				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
			}
		}

	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := q.TranslationCompleteness(); err != nil {
		return q, err
	}
	if err := q.Validate(); err != nil {
		return q, err
	}

	q.VersionEffective = -2 // re-set after validate

	return q, nil

}
