package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func prio3Matrix(
	page *pageT,
	ac assetClass,
	name string,
	lblMain trl.S,
	inps []string,
	lbls map[string]string,
	freeText bool,
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
		for idx1, nm := range inps {

			if freeText && idx1 == len(inps)-1 {
				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = fmt.Sprintf("%v_%v_label", ac.Prefix, name)
					inp.Label = trl.S{
						"de": lbls[nm],
						"en": lbls[nm],
					}
					inp.MaxChars = 18
					inp.ColSpan = 3
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 4

					inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
					inp.StyleCtl.Desktop.WidthMax = "6rem"

				}

			} else {
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
				inp.Name = fmt.Sprintf("%v_%v_prio%v", ac.Prefix, name, idx2+1)
				inp.ValueRadio = fmt.Sprintf("%v", nm) // row idx1

				inp.ColSpan = 2
				inp.ColSpanControl = 1
			}

		}

		//
		//
		//
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
				inp.JSBlockStrings[key] = fmt.Sprintf("%v_%v_prio%v", ac.Prefix, name, idx1)
			}

		}
	}

}
