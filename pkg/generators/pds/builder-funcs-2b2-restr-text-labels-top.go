package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func restrTextRowLabelsTop(
	page *qst.WrappedPageT,
	inputName string,
	lbl trl.S,
	cf configRT,
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

			inp.LabelVertical()
			inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "start"
			inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "left"

		}
	}

	// row1 - tranche types
	{
		gr := page.AddGroup()
		gr.Cols = numCols

		for idx1, trancheType := range trancheTypeNamesAC1 {

			ttPref := trancheType[:3]

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v", ttPref, inputName)
				inp.Label = allLbls["ac1-tranche-types"][idx1]
				inp.Placeholder = cf.Placeholder

				// 0%-100% in 5% brackets
				inp.Min = cf.Min
				inp.Max = cf.Max
				inp.Step = cf.Step

				inp.MaxChars = cf.Chars

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				// inp.LabelPadRight()
				inp.Vertical()
				inp.ControlCenter()

				inp.StyleLbl = styleHeaderCols2

				if idx1 == idxLastCol {
					inp.Suffix = cf.Suffix
				}
				inp.Suffix = cf.Suffix

			}
		}

	}
}
