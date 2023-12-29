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

// param iter is either 0 or 1
// because we want two distinct instances of this page
// with two distinct sets of stored params
func kneb202306simtool(q *QuestionnaireT, page *pageT, iter int) error {

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

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = fmt.Sprintf("share_safe_bg_%v", iter)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = fmt.Sprintf("share_risky_bg_%v", iter)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = fmt.Sprintf("sparbetrag_bg_%v", iter)
		}
	}

	// gr1
	grIdx := q.Version() % 2
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
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
				"de": "Werte speichern und weiter",
				"en": "todo",
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
