package qst

import "github.com/zew/go-questionnaire/pkg/trl"

type rangeConf struct {
	Min, Max, Step float64
	Suffix         trl.S

	RangeType string // labels for ticks
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
	//
	RangeType: `3--0:0;10: ;20:20;30: ;40:40;50: ;60:60;70: ;80:80;90: ;100:100`,
}

// todo: special slider display value "<2"
// 		 0;0.5; nach links rasten;  1;1.5 nach rechts rasten
var sliderPctThreeTen = rangeConf{
	Min:    0,
	Max:    10,
	Step:   0.5,
	Suffix: suffixPercent,
	// RangeType: `3--3:3;5:5;7:7;10:10`,
	// RangeType: `3--3:3;4:4;5: ;6:6;7: ;8:8;9: ;10:10`,
	RangeType: `3--0:<2;2:2;3: ;4:4;5: ;6:6;7: ;8:8;9: ;10:10`,
}

// todo: special slider display value ">2"
// 		 2.25 nach links rasten
var sliderPctZeroTwo = rangeConf{
	Min:    0,
	Max:    2.5,
	Step:   0.25,
	Suffix: suffixPercent,
	// RangeType: `3--0:0;0.5: ;1:1;1.5: ;2:2`,
	// RangeType: `3--0:0;0.5:0.5;1:1;1.5:1.5;2:2`,
	RangeType: `3--0:0;0.5:0.5;1:1;1.5:1.5;2:2;2.5:>2`,
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
// todo:  Einrasten auf 10 über 50
var sliderEBITDAZeroHundred = rangeConf{
	Min:    0,
	Max:    100,
	Step:   5,
	Suffix: suffixMillionEuro,
	// RangeType: `3--0:0;5: ;10:10;15: ;20:20;25: ;30:30;35: ;40:40;45: ;50:50;60:60;70:70;80:80;90:90;100:100`,
	RangeType: `3--0:0;10: ;20:20;30: ;40:40;50: ;60:60;70: ;80:80;90: ;100:100`,
}

// 0-500mn €in 10mn€ brackets
var sliderEVZeroFiveHundred = rangeConf{
	Min:    0,
	Max:    500,
	Step:   10,
	Suffix: suffixMillionEuro,
	// RangeType: `3--0:0;10: ;20: ;30: ;40: ;50: ;60: ;70: ;80: ;90: ;100:100;110: ;120: ;130: ;140: ;150: ;160: ;170: ;180: ;190: ;200:200;210: ;220: ;230: ;240: ;250: ;260: ;270: ;280: ;290: ;300:300;310: ;320: ;330: ;340: ;350: ;360: ;370: ;380: ;390: ;400:400;410: ;420: ;430: ;440: ;450: ;460: ;470: ;480: ;490: ;500:500`,

	// RangeType: `3--0:0;10: ;20: ;30: ;40: ;50: ;60: ;70: ;80: ;90: ;100:100;110: ;120: ;130: ;140: ;150: ;160: ;170: ;180: ;190: ;200:200;210: ;220: ;230: ;240: ;250: ;260: ;270: ;280: ;290: ;300:300;310: ;320: ;330: ;340: ;350: ;360: ;370: ;380: ;390: ;400:400;410: ;420: ;430: ;440: ;450: ;460: ;470: ;480: ;490: ;500:500`,
	RangeType: `3--0:0;50: ;100:100;150: ;200:200;250: ;300:300;350: ;400:400;450: ;500:500`,
}

// todo: smaller than 1
var sliderOneOnePointFive = rangeConf{
	Min:    1,
	Max:    1.5,
	Step:   0.05,
	Suffix: suffixMillionEuro,
	// RangeType: `3--1:1;1.25:1.25;1.5:1.5`,
	RangeType: `3--1:1;1.1:1.1;1.2:1.2;1.3:1.3;1.4:1.4;1.5:1.5`,
}
