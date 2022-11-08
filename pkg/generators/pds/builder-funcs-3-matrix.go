package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func prio3Matrix(
	page *qst.WrappedPageT,
	name string,
	lblMain trl.S,
	inps []string,
	lbls map[string]string,
) {

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblMain
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

	}

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 9
		gr.BottomVSpacers = 3

		// first row
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label.Empty()
			inp.ColSpan = 3
			inp.ColSpanLabel = 1
			inp.LabelRight()
		}
		for idx2 := 0; idx2 < 3; idx2++ {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf("Prio %v", idx2+1),
					"en": fmt.Sprintf("Prio %v", idx2+1),
				}
				inp.ColSpan = 2
				inp.ColSpanLabel = 1
				inp.LabelCenter()
			}
		}

		// second ... n-th row
		for _, nm := range inps {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": lbls[nm],
					"en": lbls[nm],
				}
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
			}

			for idx2 := 0; idx2 < 3; idx2++ {
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_prio%v", name, idx2+1)
				inp.ValueRadio = fmt.Sprintf("%v", nm) // row idx1
				// inp.Label = trl.S{
				// 	"de": lbls[nm],
				// }

				inp.ColSpan = 2
				// inp.ColSpanLabel = 1
				inp.ColSpanControl = 1

			}

		}

		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "prio123"

			s1 := trl.S{
				"de": "Keine Prio zweimal",
				"en": "Priorities not twice",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}

			inp.JSBlockStrings = map[string]string{}
			inp.JSBlockStrings["inputBaseName"] = name
			for idx1 := 0; idx1 < 3; idx1++ {
				key := fmt.Sprintf("%v_%v", "inp", idx1+1) // {{.inp_1}}, {{.inp_2}}, ...
				inp.JSBlockStrings[key] = fmt.Sprintf("%v_prio%v", name, idx1)
			}

		}
	}

}
