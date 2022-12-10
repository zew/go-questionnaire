package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func restrictedTextMultiCols(
	page *qst.WrappedPageT,
	cf configRT,
) {

	numCols := firstColLbl + float32(len(trancheTypeNamesAC1))
	idxLastCol := len(trancheTypeNamesAC1) - 1

	{
		gr := page.AddGroup()
		gr.Cols = numCols
		// gr.BottomVSpacers =

		// row0 - column headers, 1-4 tranche type names
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
			// inp.LabelVertical()
			// inp.LabelBottom()
			inp.StyleLbl = styleHeaderCols1
		}

		// row1
		for idx1, trancheType := range trancheTypeNamesAC1 {

			ttPref := trancheType[:3]

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v", ttPref, cf.InputNameP2, "main")

				inp.MaxChars = cf.Chars
				inp.Step = 1
				inp.Min = 0
				inp.Max = 1000 * 1000
				// inp.Validator = "inRange100"

				inp.Placeholder = cf.Placeholder

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				if idx1 == 0 {
					inp.Label = cf.LblRow1.Bold()
					inp.Label = cf.LblRow1
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1
				}
				if idx1 == idxLastCol {
					inp.Suffix = cf.Suffix
				}

				if cf.FirstRow100Pct {
					inp.Response = "100" // must parse to number
					inp.Disabled = true

				}

			}
		}

		// row2
		if !cf.LblRow2.Empty() {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols

			inp.Label = cf.LblRow2

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Width = "60%"
			inp.Style.Mobile.StyleBox.Width = "96%"
		}

		// rows 3,4...
		for _, suffx := range cf.SubNames {

			for idx2, trancheType := range trancheTypeNamesAC1 {

				ttPref := trancheType[:3]

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v", ttPref, cf.InputNameP2, suffx)

				inp.MaxChars = cf.Chars
				inp.Step = 1
				inp.Min = 0
				inp.Max = 1000 * 1000
				// inp.Validator = "inRange100"

				inp.Placeholder = cf.Placeholder

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				if idx2 == 0 {
					inp.Label = trl.S{
						"en": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
						"de": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
					}
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1

					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
					inp.StyleLbl.Desktop.StyleBox.Margin = "0 0 0 1.2rem"
				}

				if idx2 == idxLastCol {
					inp.Suffix = cf.Suffix
				}

			}

		}

	} // /group

}
