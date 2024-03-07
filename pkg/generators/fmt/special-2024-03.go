package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/qstcp/cpfmt"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202403(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 3
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "todo",
		"en": "Questions about climate transition",
	}
	page.Short = trl.S{
		"de": "todo",
		"en": "Questions about<br>climate transition",
	}
	page.WidthMax("75rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3

		{
			inp := gr.AddInput()
			inp.Type = "dyn-composite"
			inp.DynamicFunc = fmt.Sprintf("Special202403QS1__%v__%v", 0, 0)
			inp.DynamicFuncParamset = ""
			inp.ColSpanControl = 1
		}

		_, inputNames, _ := cpfmt.Special202403QS1(q, 0, 0, true)
		for _, inpName := range inputNames {
			inp := gr.AddInput()
			inp.Type = "dyn-composite-scalar"
			inp.Name = inpName
		}

	}

	return nil
}
