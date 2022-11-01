package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func prio3Matrix(
	page *qst.WrappedPageT,
	name string,
) {

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"

			inp.Label = trl.S{
				"de": `What GICS sectors provides the most attractive 
					investment opportunities in the next three months? 
					Please rank the top three.`,
			}

			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

	}

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 9
		gr.BottomVSpacers = 3

		inps := []string{
			"energy",
			"materials",
			"industrials",
			"consumer_discretionary",
			"consumer_staples",
			"health_care",
			"financials",
			"information_technology",
			"communication_services",
			"utilities",
			"real_estate",
		}

		lbls := map[string]string{
			"energy":                 "Energy",
			"materials":              "Materials",
			"industrials":            "Industrials",
			"consumer_discretionary": "Consumer Discretionary",
			"consumer_staples":       "Consumer Staples",
			"health_care":            "Health Care",
			"financials":             "Financials",
			"information_technology": "Information Technology",
			"communication_services": "Communication Services",
			"utilities":              "Utilities",
			"real_estate":            "Real Estate",
		}

		for idx1, nm := range inps {

			_ = idx1

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": lbls[nm],
			}

			inp.ColSpan = 3
			inp.ColSpanLabel = 1

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
