package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func restrTextRowLabelsTop(
	page *pageT,
	ac assetClass,
	inputName string,
	lbl trl.S,
	cf configRT,
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

			inp.LabelVertical()
			inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "start"
			inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "left"
		}
	}

	// row1 - tranche types
	{
		gr := page.AddGroup()
		gr.Cols = numCols

		if numCols == 1 {
			gr.WidthMax("18rem")
		}

		for idx1, trancheType := range ac.TrancheTypes {

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v", ac.Prefix, trancheType.Prefix, inputName)
				inp.Label = trancheType.Lbl
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
