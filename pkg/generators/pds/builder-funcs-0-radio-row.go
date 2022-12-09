package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// radiosSingleRow - five shades - and no answer
// previously "multipleChoice"
func radiosSingleRow(
	page *qst.WrappedPageT,
	nm string,
	lblMain trl.S,
	cf configMC,
) {

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
		gr.Cols = float32(cf.Cols)
		gr.BottomVSpacers = 3
		if cf.GroupBottomSpacers != 0 {
			gr.BottomVSpacers = cf.GroupBottomSpacers
		}
		// gr.Style = css.NewStylesResponsive(gr.Style)
		// gr.Style.Desktop.StyleGridContainer.GapRow = "0"

		for idx2 := 0; idx2 < len(allLbls[cf.KeyLabels]); idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = allLbls[cf.KeyLabels][idx2]
			inp.ColSpan = 2
			inp.ColSpan = cf.InpColspan
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			if cf.LabelBottom {
				inp.StyleLbl.Desktop.StyleGridItem.Order = 2
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "-0.3rem"
			} else {
				// top
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "0.4rem"
			}
		}

		if cf.DontKnow {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", len(allLbls[cf.KeyLabels])+1)
			inp.Label = lblDont
			inp.ColSpan = 4
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()
		}

	}

}
