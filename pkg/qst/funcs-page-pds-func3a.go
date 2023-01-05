package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func rangesRowLabelsTop(
	page *pageT,
	ac assetClass,
	inputName string,
	lbl trl.S,
	rcf rangeConf,
) {

	numCols := float32(len(ac.TrancheTypes))
	idxLastCol := len(ac.TrancheTypes) - 1

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

		if numCols == 1 {
			gr.WidthMax("18rem")
		}

		// row1
		for idx1, trancheType := range ac.TrancheTypes {

			{
				inp := gr.AddInput()
				inp.Type = "range"
				inp.DynamicFuncParamset = rcf.SerializeExtendedConfig()
				inp.Name = fmt.Sprintf("%v_%v_%v", ac.Prefix, trancheType.Prefix, inputName)
				inp.Label = trancheType.Lbl

				// 0%-100% in 5% brackets
				inp.Min = rcf.Min
				inp.Max = rcf.Max
				inp.Step = rcf.Step

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
					inp.Suffix = rcf.Suffix
				}
				inp.Suffix = rcf.Suffix

			}
		}

	}
}
