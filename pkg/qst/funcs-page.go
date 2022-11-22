package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/trl"
)

type funcPageGeneratorT func(*QuestionnaireT, *pageT) error

var pageGens = ctr.New() // page generations

var funcPGs = map[string]funcPageGeneratorT{
	"pds01": pds01,
}

func KeysValuesPage(q *QuestionnaireT, page *pageT) (keys, vals []string) {

	// finishes, ks, vs := q.KeysValues(true)
	// _, _, _ = finishes, ks, vs

	cleanse := false
	for i2 := 0; i2 < len(page.Groups); i2++ {
		for i3 := 0; i3 < len(page.Groups[i2].Inputs); i3++ {
			if page.Groups[i2].Inputs[i3].IsLayout() {
				continue
			}
			keys = append(keys, page.Groups[i2].Inputs[i3].Name)
			val := page.Groups[i2].Inputs[i3].Response
			if cleanse {
				if page.Groups[i2].Inputs[i3].Type == "number" {
					val = DelocalizeNumber(val)
				}
				val = EnglishTextAndNumbersOnly(val)
			}
			vals = append(vals, val)
		}
	}

	return
}

func pds01(q *QuestionnaireT, page *pageT) error {

	gn := pageGens.Increment()

	// depending on q
	page.NoNavigation = false

	/* 	if false {

	   		// preserved stuff
	   		gf := page.GeneratorFuncName

	   		*page = pageT{
	   			Label: trl.S{
	   				"en": fmt.Sprintf("dyn page label %v", gn),
	   				"de": fmt.Sprintf("dyn page label %v", gn),
	   			},
	   			Short: trl.S{
	   				"en": fmt.Sprintf("dyn %v", gn),
	   				"de": fmt.Sprintf("dyn %v", gn),
	   			},
	   			Desc:              trl.S{"en": "", "de": ""},
	   			GeneratorFuncName: gf,
	   		}
	   		page.WidthMax("42rem")

	   	}

	*/

	lblMain := trl.S{
		"en": fmt.Sprintf(`lbl main %v - lbl main lbl main lbl main`, gn),
		"de": fmt.Sprintf(`lbl main %v - lbl main lbl main lbl main`, gn),
	}

	if page.Finished.IsZero() {

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
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = fmt.Sprintf("text%v", gn)
				inp.Label = trl.S{
					"en": "label input",
					"de": "label input",
				}
				inp.ColSpan = 1
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.MaxChars = 40
			}

		}

	} else {
		// previous input values
		ks, vs := KeysValuesPage(q, page)
		_, _ = ks, vs
		page.Groups[0].Inputs[1].Response = vs[0]
	}

	// dynpg.MyTest()

	return nil

}
