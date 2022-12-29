package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// checkBoxCascade for hierarchical checkboxes from page1
func checkBoxCascade(
	page *qst.WrappedPageT,
	lblMain trl.S,
	// numCols float32,
	inpsL1 []string,
	lblsL1 []trl.S,
) {

	numCols := float32(5)
	// numCols := float32(4)

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblMain
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

	}

	// gr2 - 3-4, 5-6
	for idx1 := 0; idx1 < len(inpsL1); idx1++ {

		gr := page.AddGroup()
		gr.Cols = numCols
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.05rem"
		gr.BottomVSpacers = 2
		{
			ttPrefL1 := inpsL1[idx1][:3] // ac1, ac2, ac3

			// row1
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "q03_" + inpsL1[idx1]
				inp.Label = lblsL1[idx1]
				if idx1 > 0 {
					inp.Label["en"] += "<br> <span style=font-size:80%>(not yet included in following pages)</span>"
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanControl = 10
				inp.ColSpanLabel = 1
				inp.ControlFirst()

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0.4rem 0 0 0"
			}

			//
			//
			// row2
			// 		indented label
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
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
				inp.DisplayNone()
			}

			// row3
			// 		indented level2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
			}

			inpsL2 := trancheTypeNamesAC1
			lblsL2 := allLbls["ac1-tranche-types"]

			if idx1 == 1 {
				inpsL2 = trancheTypeNamesAC2
				lblsL2 = allLbls["ac2-tranche-types"]
			}
			if idx1 == 2 {
				inpsL2 = trancheTypeNamesAC3
				lblsL2 = allLbls["ac3-tranche-types"]
			}

			for idx2 := 0; idx2 < len(inpsL2); idx2++ {

				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("q031_%v_%v", ttPrefL1, inpsL2[idx2])
				inp.Label = lblsL2[idx2]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.Vertical()
				inp.VerticalLabel()

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Position = "relative"
				inp.Style.Desktop.StyleBox.Top = "-0.4rem"

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.FontSize = 90
				inp.DisplayNone()

			}
			if len(inpsL2) == 3 {
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.ColSpanLabel = 1
					inp.DisplayNone()
				}
			}
			if len(inpsL2) == 2 {
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.ColSpanLabel = 1
					inp.DisplayNone()
				}
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.ColSpanLabel = 1
					inp.DisplayNone()
				}
			}

			//
			//
			// row4
			// 		indented label
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 4
				inp.Label = trl.S{
					"de": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
					"en": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
				}
				inp.ColSpanLabel = 1
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.FontSize = 85

				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "0.3rem"
				inp.DisplayNone()
			}
			// row5
			// 		indented level2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
			}
			for idx2 := 0; idx2 < len(allLbls["teamsize"]); idx2++ {
				inp := gr.AddInput()
				inp.Type = "radio"
				// inp.Name = fmt.Sprintf("%v", nm)
				inp.Name = fmt.Sprintf("q032_%v", ttPrefL1)
				// inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
				inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
				inp.Label = allLbls["teamsize"][idx2]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.Vertical()
				inp.VerticalLabel()

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				if false {
					inp.StyleLbl.Desktop.StyleGridItem.Order = 2
					inp.StyleLbl.Desktop.StyleBox.Position = "relative"
					inp.StyleLbl.Desktop.StyleBox.Top = "-0.3rem"
				} else {
					// top
					// inp.StyleLbl.Desktop.StyleBox.Position = "relative"
					// inp.StyleLbl.Desktop.StyleBox.Top = "0.4rem"
				}
				inp.DisplayNone()

			}

		}

	}
}
