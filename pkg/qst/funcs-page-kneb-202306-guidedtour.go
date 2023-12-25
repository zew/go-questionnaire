package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func kneb202306guidedtour(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"en": "Experiment-Chart-Introduction",
		"de": "Experiment chart-Introduction",
	}
	page.Label = trl.S{
		"en": "",
		"de": "",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("52rem")

	// gr0
	grIdx := q.UserIDInt() % 2
	{
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				// store current page of the guided tour
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "section"
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = fmt.Sprintf("./slide-show/index-%d.html", grIdx)
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}
	}

	return nil
}
