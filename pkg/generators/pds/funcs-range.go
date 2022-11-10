package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func rangeClosingTime(page *qst.WrappedPageT, inputName string, lbl trl.S) {

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		gr.WidthMax("70%")

		{
			inp := gr.AddInput()
			inp.Type = "range"
			inp.Name = fmt.Sprintf("%v_closing_time", inputName)
			inp.Label = lbl
			// below 6 months, 6m-18m in 3m brackets, over 18m

			inp.Min = 3
			inp.Max = 21
			inp.Step = 3

			inp.Suffix = trl.S{
				"en": "weeks",
				"de": "Wochen",
			}

			inp.ColSpan = 1
			inp.ColSpanControl = 1
		}

	}

}

func rangePercentage(page *qst.WrappedPageT, inputName string, lbl trl.S, subName string) {

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		gr.WidthMax("80%")

		{
			inp := gr.AddInput()
			inp.Type = "range"
			inp.Name = fmt.Sprintf("%v_%v", inputName, subName)
			inp.Label = lbl
			// below 6 months, 6m-18m in 3m brackets, over 18m

			inp.Min = 3
			inp.Max = 21
			inp.Step = 3

			inp.Suffix = trl.S{
				"en": "%",
				"de": "%",
			}

			inp.ColSpan = 1
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
		}

	}

}
