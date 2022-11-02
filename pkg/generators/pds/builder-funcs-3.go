package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// sentimentSingleRow - five shades - and no answer
func sentimentSingleRow(
	page *qst.WrappedPageT,
	nm string,
	lblMain trl.S,
) {

	lbls := []trl.S{
		{
			"de": "Improve significantly",
			"en": "Improve significantly",
		},
		{
			"de": "Improve slightly",
			"en": "Improve slightly",
		},
		{
			"de": "Stay constant",
			"en": "Stay constant",
		},
		{
			"de": "Worsen slightly",
			"en": "Worsen slightly",
		},
		{
			"de": "Improve significantly",
			"en": "Improve significantly",
		},
	}

	lblDont := trl.S{
		"de": "Don´t know",
		"en": "Don´t know",
	}

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

	// gr2
	{
		gr := page.AddGroup()
		gr.Cols = 14
		gr.BottomVSpacers = 1

		for idx2 := 0; idx2 < 5; idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = lbls[idx2]
			inp.ColSpan = 2
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()
		}

		for idx2 := 5; idx2 < 6; idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = lblDont
			inp.ColSpan = 4
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

		}

		// {
		// 	inp := gr.AddInput()
		// 	inp.ColSpanControl = 1
		// 	inp.Type = "javascript-block"
		// 	inp.Name = "prio123"

		// 	s1 := trl.S{
		// 		"de": "Keine Prio zweimal",
		// 		"en": "Priorities not twice",
		// 	}
		// 	inp.JSBlockTrls = map[string]trl.S{
		// 		"msg": s1,
		// 	}

		// 	inp.JSBlockStrings = map[string]string{}
		// 	inp.JSBlockStrings["inputBaseName"] = name
		// 	for idx1 := 0; idx1 < 3; idx1++ {
		// 		key := fmt.Sprintf("%v_%v", "inp", idx1+1) // {{.inp_1}}, {{.inp_2}}, ...
		// 		inp.JSBlockStrings[key] = fmt.Sprintf("%v_prio%v", name, idx1)
		// 	}

		// }
	}

}
