package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/qstcp/cpfmt"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202303(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2023 && q.Survey.Month == 3
	if !cond {
		return nil
	}

	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Korrelation Anlageklassen",
			"en": "Correlation of asset classes",
		}
		page.Short = trl.S{
			"de": "<br>Korrelation",
			"en": "<br>Correlation",
		}
		page.WidthMax("80rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			// gr.WidthMax("45rem")
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Mit Blick auf die nächsten sechs Monate 
						erwarte ich die folgenden Korrelationen zwischen den Gesamtrenditen 
						der folgenden Anlageklassen (breit gestreute Indizes) 
						aus dem&nbsp;<b><i>Eurogebiet</i></b>. 
					`,

					"en": `
						Over the coming 6 months, 
						I expect the following correlations 
						between the total returns of the following asset classes 
						in the&nbsp;<b><i>euro&nbsp;area</i></b>.
					`,
				}.Outline("4.")
				inp.ColSpanLabel = 1
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.DynamicFunc = fmt.Sprintf("Special202303__%v__%v", 0, 0)
				inp.DynamicFuncParamset = ""
				inp.ColSpanControl = 1
			}
			// would be good, if we could write this as
			// inputNames, _,_ = q.HasComposit(pageIdx, groupIdx)
			_, inputNames, _ := cpfmt.Special202303(q, 0, 0, true)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "dyn-composite-scalar"
				inp.Name = inpName
			}

		}

	}

	return nil
}
