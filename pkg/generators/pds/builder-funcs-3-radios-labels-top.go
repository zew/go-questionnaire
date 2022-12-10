package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func radiosLabelsTop(
	page *qst.WrappedPageT,
	nm string,
	lbl trl.S,
	cf configMC,
) {

	// numCols := firstColLbl + float32(len(trancheTypeNamesAC1))
	numColsMajor := float32(len(trancheTypeNamesAC1))
	numColsMinor := numColsMajor * cf.Cols
	idxLastCol := len(trancheTypeNamesAC1) - 1
	_ = idxLastCol

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

		for idx1 := range trancheTypeNamesAC1 {

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1

			ttLbl := allLbls["ac1-tranche-types"][idx1]
			// inp.Label = ttLbl.Bold()
			inp.Label = ttLbl

			// inp.LabelVertical()

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

		// for idx1 := 0; idx1 < len(trancheTypeNamesAC1)+1; idx1++ {
		for idx1, trancheType := range trancheTypeNamesAC1 {

			_ = idx1

			// row1 - inputs
			ttPref := trancheType[:3]

			lastIdx2 := len(allLbls[cf.KeyLabels]) - 1

			for idx2 := 0; idx2 < len(allLbls[cf.KeyLabels]); idx2++ {
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_%v", ttPref, nm)
				inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
				inp.Label = allLbls[cf.KeyLabels][idx2]

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

			// if cf.DontKnow {
			// 	inp := gr.AddInput()
			// 	inp.Type = "radio"
			// 	inp.Name = fmt.Sprintf("%v", nm)
			// 	inp.ValueRadio = fmt.Sprintf("%v", len(allLbls[cf.KeyLabels])+1)
			// 	inp.Label = lblDont
			// 	inp.ColSpan = 4
			// 	inp.ColSpanControl = 1
			// 	inp.Vertical()
			// 	inp.VerticalLabel()
			// }

		}
	}

}
