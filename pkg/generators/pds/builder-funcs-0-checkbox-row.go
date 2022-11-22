package pds

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func checkBoxRow(
	page *qst.WrappedPageT,
	lblMain trl.S,
	names []string,
	lbls []trl.S,
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
		gr.Cols = float32(len(names))
		// gr.BottomVSpacers = 1

		for idx1 := 0; idx1 < len(names); idx1++ {
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = names[idx1]
			inp.Label = lbls[idx1]
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

func multipeChoiceColumn(
	page *qst.WrappedPageT,
	lblMain trl.S,
	numCols float32,
	inps []string,
	lbls []trl.S,
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
		gr.Cols = numCols
		// gr.BottomVSpacers = 1

		for idx1 := 0; idx1 < len(inps); idx1++ {
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = inps[idx1]
			inp.Label = lbls[idx1]
			inp.ColSpan = 1
			inp.ColSpanControl = 4
			inp.ColSpanLabel = 1
			inp.ControlFirst()
			// inp.Vertical()
			// inp.VerticalLabel()

			// labelBottom := false
			// inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			// if labelBottom {
			// 	inp.StyleLbl.Desktop.StyleGridItem.Order = 2
			// } else {
			// 	// top
			// 	inp.StyleLbl.Desktop.StyleBox.Position = "relative"
			// 	inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			// }
		}

	}
}
