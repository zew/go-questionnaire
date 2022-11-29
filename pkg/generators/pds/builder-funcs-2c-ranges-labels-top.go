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
	RangeType: `1--3:<6;6:6;9: ;12:12;15: ;18:18;21:>18`,
}

var sliderPctZeroHundredMiddle = rangeConf{
	Min:    0,
	Max:    100,
	Step:   5,
	Suffix: suffixPercent,
	// RangeType: `2--0:0;20:20;40:40;60:60;80:80;100:100`,
	RangeType: `2--0:0;25:25;50:50;75:75;100:100`,
}
var sliderPctZeroHundredWide = rangeConf{
	Min:    0,
	Max:    100,
	Step:   5,
	Suffix: suffixPercent,
	// RangeType: `3--0:0;20:20;40:40;60:60;80:80;100:100`,
	RangeType: `3--0:0;10: ;20:20;30: ;40:40;50: ;60:60;70: ;80:80;90: ;100:100`,
}

var sliderPctThreeTen = rangeConf{
	Min:       3,
	Max:       10,
	Step:      0.5,
	Suffix:    suffixPercent,
	RangeType: `3--3:3;5:5;7:7;10:10`,
}

// todo: smaller than 1
var sliderPctZeroTwo = rangeConf{
	Min:    0,
	Max:    2,
	Step:   0.25,
	Suffix: suffixPercent,
	// comma as decimal separator?
	// RangeType: `3--0:0;0,5: ;1:1;1,5: ;2:2`,
	RangeType: `3--0:0;1:1;2:2`,
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

// _0- 50 mn € in  5 mn€ brackets
// 50-100 mn € in 10 mn€ brackets
var sliderEBITDAZeroHundred = rangeConf{
	Min:       0,
	Max:       100,
	Step:      5,
	Suffix:    suffixMillionEuro,
	RangeType: `3--0:0;5: ;10:10;15: ;20:20;25: ;30:30;35: ;40:40;45: ;50:50;60:60;70:70;80:80;90:90;100:100`,
}

// 0-500mn €in 10mn€ brackets
var sliderEVZeroFiveHundred = rangeConf{
	Min:    0,
	Max:    500,
	Step:   10,
	Suffix: suffixMillionEuro,
	// RangeType: `3--0:0;10:10;20:20;30:30;40:40;50:50;60:60;70:70;80:80;90:90;100:100;110:110;120:120;130:130;140:140;150:150;160:160;170:170;180:180;190:190;200:200;210:210;220:220;230:230;240:240;250:250;260:260;270:270;280:280;290:290;300:300;310:310;320:320;330:330;340:340;350:350;360:360;370:370;380:380;390:390;400:400;410:410;420:420;430:430;440:440;450:450;460:460;470:470;480:480;490:490;500:500`,
	// RangeType: `3--0:0;10: ;20:20;30: ;40:40;50: ;60:60;70: ;80:80;90: ;100:100;110: ;120:120;130: ;140:140;150: ;160:160;170: ;180:180;190: ;200:200;210: ;220:220;230: ;240:240;250: ;260:260;270: ;280:280;290: ;300:300;310: ;320:320;330: ;340:340;350: ;360:360;370: ;380:380;390: ;400:400;410: ;420:420;430: ;440:440;450: ;460:460;470: ;480:480;490: ;500:500`,
	// RangeType: `3--0:0;10: ;20: ;30: ;40: ;50:50;60: ;70: ;80: ;90: ;100:100;110: ;120: ;130: ;140: ;150:150;160: ;170: ;180: ;190: ;200:200;210: ;220: ;230: ;240: ;250:250;260: ;270: ;280: ;290: ;300:300;310: ;320: ;330: ;340: ;350:350;360: ;370: ;380: ;390: ;400:400;410: ;420: ;430: ;440: ;450:450;460: ;470: ;480: ;490: ;500:500`,
	RangeType: `3--0:0;10: ;20: ;30: ;40: ;50: ;60: ;70: ;80: ;90: ;100:100;110: ;120: ;130: ;140: ;150: ;160: ;170: ;180: ;190: ;200:200;210: ;220: ;230: ;240: ;250: ;260: ;270: ;280: ;290: ;300:300;310: ;320: ;330: ;340: ;350: ;360: ;370: ;380: ;390: ;400:400;410: ;420: ;430: ;440: ;450: ;460: ;470: ;480: ;490: ;500:500`,
}

// todo: smaller than 1
var sliderOneOnePointFive = rangeConf{
	Min:       1,
	Max:       1.5,
	Step:      0.05,
	Suffix:    suffixMillionEuro,
	RangeType: `3--1:1;1.5:1.5`,
}

func rangesRowLabelsTop(
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

				inp.StyleLbl = trancheNameStyle

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
