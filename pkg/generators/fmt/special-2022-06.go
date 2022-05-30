package fmt

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202206(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 6
	if !cond {
		return nil
	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.NoNavigation = true
		// page.SuppressProgressbar = true

		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "/echart/inner.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}
	}

	return nil

}
