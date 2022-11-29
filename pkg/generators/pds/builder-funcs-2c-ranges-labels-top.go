package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

type rangeConf struct {
	Min, Max, Step float64
	Suffix         trl.S

	RangeType string // labels
}

var sliderWeeksClosing = rangeConf{
	Min:    3,
	Max:    21,
	Step:   3,
	Suffix: suffixWeeks,
	// RangeType: `1--3:<6;6:6;9:9;12:12;15:15;18:18;21:>18`,
	RangeType: `1--3:<6;6:6;12:12;18:18;21:>18`,
}

var sliderPctZeroHundredMiddle = rangeConf{
	Min:       0,
	Max:       100,
	Step:      5,
	Suffix:    suffixPercent,
	RangeType: `2--0:0;20:20;40:40;60:60;80:80;100:100`,
}
var sliderPctZeroHundredWide = rangeConf{
	Min:       0,
	Max:       100,
	Step:      5,
	Suffix:    suffixPercent,
	RangeType: `3--0:0;20:20;40:40;60:60;80:80;100:100`,
}

var sliderPctThreeTen = rangeConf{
	Min:       3,
	Max:       10,
	Step:      0.5,
	Suffix:    suffixPercent,
	RangeType: `3--3:3;5:5;7:7;10:10`,
}

var sliderPctZeroTwo = rangeConf{
	Min:       0,
	Max:       2,
	Step:      0.25,
	Suffix:    suffixPercent,
	RangeType: `3--0:0;0.5:0.5;1:1;1.5:1.5;2:2`,
}

var sliderPctZeroFour = rangeConf{
	Min:       0,
	Max:       4,
	Step:      0.25,
	Suffix:    suffixPercent,
	RangeType: `3--0:0;1:1;2:2;3:3;4:4`,
}

var sliderPctThreeTwenty = rangeConf{
	Min:       3,
	Max:       20,
	Step:      0.5,
	Suffix:    suffixPercent,
	RangeType: `3--3:3;5:5;10:10;15:15;20:20`,
}

var sliderPctThreeTwentyfive = rangeConf{
	Min:       3,
	Max:       25,
	Step:      0.5,
	Suffix:    suffixPercent,
	RangeType: `3--3:3;5:5;10:10;15:15;20:20;25:25`,
}

var sliderYearsZeroTen = rangeConf{
	Min:       0,
	Max:       10,
	Step:      0.5,
	Suffix:    suffixYears,
	RangeType: `3--0:0;2:2;4:4;6:6;8:8;10:10`,
}

var sliderEBITDA2x10x = rangeConf{
	Min:       2,
	Max:       10,
	Step:      0.5,
	Suffix:    suffixEBITDA,
	RangeType: `3--2:2;4:4;6:6;8:8;10:10`,
}

func slidersPctRowLabelsTop(
	page *qst.WrappedPageT,
	inputName string,
	lbl trl.S,
	// sfx trl.S,
	cf rangeConf,
) {

	numCols := float32(len(trancheTypeNamesAC1))
	idxLastCol := len(trancheTypeNamesAC1) - 1

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
				inp.Type = "range"
				inp.DynamicFuncParamset = cf.RangeType
				inp.Name = fmt.Sprintf("%v_%v", ttPref, inputName)
				inp.Label = allLbls["ac1-tranche-types"][idx1]

				// 0%-100% in 5% brackets
				inp.Min = cf.Min
				inp.Max = cf.Max
				inp.Step = cf.Step

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				// inp.LabelPadRight()
				inp.LabelCenter()
				inp.Vertical()

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.FontSize = 95

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0 1.4rem 0 1.4rem"

				if idx1 == idxLastCol {
					inp.Suffix = cf.Suffix
				}
				inp.Suffix = cf.Suffix

			}
		}

	}
}
