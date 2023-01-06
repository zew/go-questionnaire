package qst

import (
	"fmt"
	"strings"

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

	// master-row - or columns - asset classes
	{
		grMaster := page.AddGroup()
		if cf.GroupBottomSpacers != 0 {
			grMaster.BottomVSpacers = cf.GroupBottomSpacers
		}
		grMaster.BottomVSpacers = 0

		grMaster.Cols = float32(len(ac.TrancheTypes))
		grMaster.ChildGroups = len(ac.TrancheTypes)

		grMaster.WidthMax(grMax)

		//
		//
		grMaster.Style = css.NewStylesResponsive(grMaster.Style)
		grMaster.Style.Desktop.StyleGridContainer.TemplateColumns = strings.Repeat("1fr ", len(ac.TrancheTypes))
		grMaster.Style.Mobile.StyleGridContainer.TemplateColumns = "1fr "

		grMaster.Style.Desktop.StyleBox.Margin = "0 0 2.2rem 0"          // instead of BottomVSpacers
		grMaster.Style.Mobile.StyleBox.Margin = "0 1.5rem 2.2rem 1.5rem" // left and right

		for idx1, trancheType := range ac.TrancheTypes {

			_ = idx1

			grChild := page.AddGroup()
			grChild.Cols = cf.Cols
			grChild.BottomVSpacers = 0

			// subrow 0 - tranche type name
			{
				inp := grChild.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = grChild.Cols
				inp.Label = trancheType.Lbl
				inp.StyleLbl = styleHeaderCols3
			}

			// subrow 1 - radio inputs
			lastIdx2 := int(cf.Cols) - 1
			for idx2 := 0; idx2 < int(cf.Cols); idx2++ {
				inp := grChild.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_%v_%v", ac.Prefix, trancheType.Prefix, nm)
				inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
				inp.Label = PDSLbls[cf.KeyLabels][idx2]

				inp.ColSpan = cf.InpColspan
				inp.ColSpanControl = 1
				inp.Vertical()
				inp.VerticalLabel()

				//
				// label top or bottom
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				if cf.LabelBottom {
					inp.StyleLbl.Desktop.StyleGridItem.Order = 2
				} else {
					inp.StyleLbl.Desktop.StyleBox.Position = "relative"
					inp.StyleLbl.Desktop.StyleBox.Top = "-0.3rem"
				}

				//
				// label styling and displacement for "dense group pattern"
				inp.StyleLbl.Desktop.StyleText.FontSize = 85
				inp.StyleLbl.Mobile.StyleText.FontSize = 85
				if idx2 == 0 {
					inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "left"
					inp.StyleLbl.Desktop.StyleBox.Left = "0.6rem"
				}
				if idx2 == lastIdx2 {
					inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "right"
					inp.StyleLbl.Desktop.StyleBox.Right = "0.6rem"
				}
				inp.StyleLbl.Mobile.StyleBox.Left = "0"
				inp.StyleLbl.Mobile.StyleBox.Right = "0"

				//
				// entire input displacement for "dense group pattern"
				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Position = "relative"
				if idx2 < len(cf.XDisplacements) {
					if idx2 <= lastIdx2/2 {
						inp.Style.Desktop.StyleBox.Left = cf.XDisplacements[idx2]
					} else {
						inp.Style.Desktop.StyleBox.Right = cf.XDisplacements[idx2]
					}
				}
				inp.Style.Mobile.StyleBox.Left = "0"
				inp.Style.Mobile.StyleBox.Right = "0"

				//
				//

				// 		/*
				// 			inp.Style.Mobile.StyleBox.Left = "0"
				// 			inp.Style.Mobile.StyleBox.Right = "0"

				// 			inp.StyleLbl.Mobile.StyleBox.Left = "0"
				// 			inp.StyleLbl.Mobile.StyleBox.Right = "0"

				// 			if idx2 == 0 {
				// 				// inp.StyleLbl.Mobile.StyleBox.Left = "0.2rem"
				// 				inp.StyleLbl.Mobile.StyleText.AlignHorizontal = "right"
				// 				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				// 				inp.StyleCtl.Mobile.StyleGridItem.JustifySelf = "end"
				// 			}
				// 			if idx2 == lastIdx2 {
				// 				// inp.StyleLbl.Mobile.StyleBox.Right = "0.2rem"
				// 				inp.StyleLbl.Mobile.StyleText.AlignHorizontal = "left"
				// 				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				// 				inp.StyleCtl.Mobile.StyleGridItem.JustifySelf = "start"
				// 			}

				// 			// inp.StyleLbl.Desktop.StyleBox.Padding = "0 0.3rem 0 0"
				// 			// inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0 0"

				// 		*/

			}
		}
	}
}
