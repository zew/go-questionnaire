package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// const firstColLbl = float32(4)
const firstColLbl = float32(3)

var suffixWeeks = trl.S{
	"en": "weeks",
	"de": "Wochen",
}
var suffixPercent = trl.S{
	"en": "%",
	"de": "%",
}

func slidersPctRowLabelsLeft(
	page *qst.WrappedPageT,
	inputName string,
	lbl trl.S,
	sfx trl.S,
	rangeType string,
) {

	numCols := firstColLbl + float32(len(trancheTypeNamesAC1))
	idxLastCol := len(trancheTypeNamesAC1) - 1

	{
		gr := page.AddGroup()
		gr.Cols = numCols

		// row0 - headers
		for idx1 := 0; idx1 < len(trancheTypeNamesAC1)+1; idx1++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if idx1 == 0 {
				inp.ColSpan = firstColLbl
			}
			if idx1 > 0 {
				ttLbl := allLbls["ac1-tranche-types"][idx1-1]
				// inp.Label = ttLbl.Bold()
				inp.Label = ttLbl
			}
			inp.LabelVertical()
			// label left
			inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "start"
			inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "left"
			inp.StyleLbl.Desktop.StyleText.FontSize = 90
		}

		// row1
		for idx1, trancheType := range trancheTypeNamesAC1 {

			ttPref := trancheType[:3]

			{
				inp := gr.AddInput()
				inp.Type = "range"
				inp.DynamicFuncParamset = rangeType

				inp.Name = fmt.Sprintf("%v_%v", ttPref, inputName)

				// 0%-100% in 5% brackets
				inp.Min = 0
				inp.Max = 100
				inp.Step = 5

				if rangeType == "3" {
					// below 6 months, 6m-18m in 3m brackets, over 18m
					inp.Min = 3
					inp.Max = 21
					inp.Step = 3
				}

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.WidthMax = "90%"

				if idx1 == 0 {
					inp.Label = lbl
					inp.LabelPadRight()
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1
				}

				if idx1 == idxLastCol {
					inp.Suffix = sfx
				}
				inp.Suffix = sfx

			}
		}

	}
}
