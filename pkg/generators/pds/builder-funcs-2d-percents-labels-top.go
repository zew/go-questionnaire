package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func percentsRowLabelsTop(
	page *qst.WrappedPageT,
	inputName string,
	lbl trl.S,
	sfx trl.S,
) {

	numCols := float32(len(trancheTypeNamesAC1))
	idxLastCol := len(trancheTypeNamesAC1) - 1

	// row0 - major label
	if !lbl.Empty() {
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		gr.BottomVSpacers = 0
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
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v", ttPref, inputName)
				inp.Label = allLbls["ac1-tranche-types"][idx1]
				inp.Placeholder = trl.S{
					"en": "#",
					"de": "#",
				}

				// 0%-100% in 5% brackets
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.1

				inp.MaxChars = 4

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				// inp.LabelPadRight()
				inp.Vertical()

				inp.LabelCenter()
				inp.ControlCenter()

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.FontSize = 95

				// inp.Style = css.NewStylesResponsive(inp.Style)
				// inp.Style.Desktop.StyleBox.Margin = "0 1.4rem 0 1.4rem"

				if idx1 == idxLastCol {
					inp.Suffix = sfx
				}
				inp.Suffix = sfx

			}
		}

	}
}
