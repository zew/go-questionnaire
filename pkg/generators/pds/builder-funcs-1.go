package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// checkBoxColumnCascade for hierarchical checkboxes from page1
func checkBoxColumnCascade(
	page *qst.WrappedPageT,
	lblMain trl.S,
	// numCols float32,
	inps []string,
	lbls []trl.S,
) {

	numCols := float32(5)

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

	}

	// gr2
	{
		gr := page.AddGroup()
		gr.Cols = numCols
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapRow = "0"

		for idx1 := 0; idx1 < len(inps); idx1++ {

			// row1
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "xx_" + inps[idx1]
				inp.Label = lbls[idx1]
				if idx1 > 0 {
					inp.Label["en"] = inp.Label["en"] + " - made available in Q2 23?"
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanControl = 8
				inp.ColSpanLabel = 1
				inp.ControlFirst()
			}

			// row2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 4
				inp.Label = trl.S{
					"en": "Which Strategies do you engage in?",
					"de": "In welchen Strategien engagieren Sie sich?",
				}
				inp.ColSpanLabel = 1
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.FontSize = 85
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "0.2rem"
			}

			// row3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}

			inps := trancheTypeNamesAC1
			lbls2 := allLbls["ac1-tranche-types"]

			if idx1 == 1 {
				inps = trancheTypeNamesAC2
				lbls2 = allLbls["ac2-tranche-types"]
			}
			if idx1 == 2 {
				inps = trancheTypeNamesAC3
				lbls2 = allLbls["ac3-tranche-types"]
			}

			for idx2 := 0; idx2 < len(inps); idx2++ {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("%v_%v_", idx1, idx2) + inps[idx2]
				inp.Label = lbls2[idx2]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.Vertical()
				inp.VerticalLabel()

				labelBottom := false
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				if labelBottom {
					inp.StyleLbl.Desktop.StyleGridItem.Order = 2
				} else {
					// top
					inp.StyleLbl.Desktop.StyleBox.Position = "relative"
					inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
				}
			}

		}
	}
}
