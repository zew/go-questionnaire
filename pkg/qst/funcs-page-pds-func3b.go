package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func radiosLabelsTop(
	page *pageT,
	ac assetClass,
	nm string,
	lbl trl.S,
	cf configMC,
) {

	if len(ac.TrancheTypes) == 0 {
		return
	}

	// numCols := firstColLbl + float32(len(trancheTypeNamesAC1))
	numColsMajor := float32(len(ac.TrancheTypes))
	numColsMinor := numColsMajor * cf.Cols
	idxLastCol := len(ac.TrancheTypes) - 1
	_ = idxLastCol

	grMax := ""
	if len(ac.TrancheTypes) == 3 {
		grMax = "54rem"
	}
	if len(ac.TrancheTypes) == 2 {
		grMax = "38rem"
	}
	if len(ac.TrancheTypes) == 1 {
		grMax = "18rem"
	}

	grSt := css.NewStylesResponsive(nil)
	if cf.GroupLeftIndent != "" {
		grSt.Desktop.StyleBox.Margin = "0 0 0 " + cf.GroupLeftIndent
	}

	// row0 - major label
	if !lbl.Empty() {
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

	// row1 - asset classes
	{
		gr := page.AddGroup()
		gr.Cols = numColsMajor
		gr.BottomVSpacers = 0
		gr.BottomVSpacers = 1
		gr.WidthMax(grMax)

		for idx1 := range ac.TrancheTypes {

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1

			ttLbl := ac.TrancheTypes[idx1].Lbl
			inp.Label = ttLbl

			inp.StyleLbl = styleHeaderCols3
		}

	}

	// radios
	{
		gr := page.AddGroup()
		gr.Cols = numColsMinor
		gr.BottomVSpacers = 3
		if cf.GroupBottomSpacers != 0 {
			gr.BottomVSpacers = cf.GroupBottomSpacers
		}
		gr.WidthMax(grMax)

		// for idx1 := 0; idx1 < len(trancheTypeNamesAC1)+1; idx1++ {
		for idx1, trancheType := range ac.TrancheTypes {

			_ = idx1

			// row1 - inputs
			lastIdx2 := len(PDSLbls[cf.KeyLabels]) - 1

			for idx2 := 0; idx2 < len(PDSLbls[cf.KeyLabels]); idx2++ {
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_%v_%v", ac.Prefix, trancheType.Prefix, nm)
				inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
				inp.Label = PDSLbls[cf.KeyLabels][idx2]

				inp.ColSpan = cf.InpColspan
				inp.ColSpanControl = 1
				inp.Vertical()
				inp.VerticalLabel()

				//
				// label styling
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				if cf.LabelBottom {
					inp.StyleLbl.Desktop.StyleGridItem.Order = 2
				} else {
					// top
					inp.StyleLbl.Desktop.StyleBox.Position = "relative"
					inp.StyleLbl.Desktop.StyleBox.Top = "-0.3rem"
				}
				inp.StyleLbl.Desktop.StyleText.FontSize = 90

				//
				//
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Position = "relative"

				if idx2 == 0 {
					inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "left"
					inp.StyleLbl.Desktop.StyleBox.Left = "0.6rem"
				}
				if idx2 == lastIdx2 {
					inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "right"
					inp.StyleLbl.Desktop.StyleBox.Right = "0.6rem"
				}

				if idx2 < len(cf.XDisplacements) {
					// if idx2 < lastIdx2/2 {
					if idx2 <= lastIdx2/2 {
						inp.Style.Desktop.StyleBox.Left = cf.XDisplacements[idx2]
					} else {
						inp.Style.Desktop.StyleBox.Right = cf.XDisplacements[idx2]
					}
				}

			}

		}
	}

}
