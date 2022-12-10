package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func rangesRowLabelsTop(
	page *qst.WrappedPageT,
	inputName string,
	lbl trl.S,
	// sfx trl.S,
	cf rangeConf,
) {

	numCols := float32(len(trancheTypeNamesAC1))
	idxLastCol := len(trancheTypeNamesAC1) - 1

	// row0 - major label
	if !lbl.Empty() {

		grSt := css.NewStylesResponsive(nil)
		grSt.Desktop.StyleBox.Margin = "0 0 0 " + outline2Indent

		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		gr.BottomVSpacers = 0
		gr.Style = grSt
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lbl
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = numCols

		// row1
		for idx1, trancheType := range trancheTypeNamesAC1 {

			ttPref := trancheType[:3]

			{
				inp := gr.AddInput()
				inp.Type = "range"
				inp.DynamicFuncParamset = cf.RangeType
				inp.Name = fmt.Sprintf("%v_%v", ttPref, inputName)
				inp.Label = allLbls["ac1-tranche-types"][idx1]

				// 0%-100% in 5% brackets
				inp.Min = cf.Min
				inp.Max = cf.Max
				inp.Step = cf.Step

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				inp.Vertical()

				// inp.LabelPadRight()
				// inp.LabelCenter()
				inp.StyleLbl = styleHeaderCols2

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0 1.4rem 0 1.4rem"

				if idx1 == idxLastCol {
					inp.Suffix = cf.Suffix
				}
				inp.Suffix = cf.Suffix

			}
		}

	}
}
