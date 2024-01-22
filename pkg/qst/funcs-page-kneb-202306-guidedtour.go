package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// obsolete
func kneb202306guidedtour_REVEAL(q *QuestionnaireT, page *pageT) error {

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
	grIdx := q.Version() % 2
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

func kneb202306guidedtourN0(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 0)
}
func kneb202306guidedtourN1(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 1)
}
func kneb202306guidedtourN2(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 2)
}
func kneb202306guidedtourN3(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 3)
}
func kneb202306guidedtourN4(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 4)
}
func kneb202306guidedtourN5(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 5)
}
func kneb202306guidedtourN6(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 6)
}

func kneb202306guidedtourN7(q *QuestionnaireT, page *pageT) error {
	return kneb202306guidedtourSepSingle(q, page, 7)
}

func kneb202306guidedtourSepSingle(q *QuestionnaireT, page *pageT, pageIdx int) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"en": "",
		"de": "",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("60rem")

	// gr0
	grIdx := q.Version() % 2

	{
		gr := page.AddGroup()
		gr.BottomVSpacers = 0
		gr.Cols = 1
		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "knebGuidedTourImg" // js filename

			s1 := trl.S{
				"de": "no javascript dialog message needed",
				"en": "no javascript dialog message needed",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}

			inp.JSBlockStrings = map[string]string{}
			inp.JSBlockStrings["appPrefix"] = cfg.Pref()
			inp.JSBlockStrings["treatIdx"] = fmt.Sprintf("%v", grIdx)
			inp.JSBlockStrings["pageIdx"] = fmt.Sprintf("%v", pageIdx)

		}

	}

	return nil
}
