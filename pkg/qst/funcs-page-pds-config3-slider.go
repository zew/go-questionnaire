package qst

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/pkg/trl"
)

type rangeConf struct {
	Min, Max, Step float64
	Suffix         trl.S

	CSSType string // CSS file containing variable

	TicksLabels string    // ticks and labels; list elmeents separated by semicolon; pairs separated by colon
	xs          []float64 // range steps, where ticks should appear
	lbls        []string  // the tick-label; can be empty

	LowerThreshold float64
	LowerDisplay   string

	UpperThreshold float64
	UpperDisplay   string
}

func (rc *rangeConf) New(inp *inputT) {

	if rc == nil {
		*rc = rangeConf{}
		rc.xs = []float64{}
		rc.lbls = []string{}
	}

	rc.Min = inp.Min
	rc.Max = inp.Max
	rc.Step = inp.Step
	rc.Suffix = inp.Suffix

	tokensL1 := strings.Split(inp.DynamicFuncParamset, "--") // tokens level 1
	if len(tokensL1) != 3 {
		log.Panicf("invalid range config string %s - %s", inp.Name, inp.DynamicFuncParamset)
		return
	}

	//
	//
	//
	rc.CSSType = tokensL1[0]

	//
	//
	//
	rc.TicksLabels = tokensL1[1]
	scalePartsStr := strings.Split(rc.TicksLabels, ";")
	for _, scaleEl := range scalePartsStr {
		valLbl := strings.Split(scaleEl, ":")
		valLbl[0] = strings.ReplaceAll(valLbl[0], ",", ".")
		vl, err := strconv.ParseFloat(valLbl[0], 64)
		if err != nil {
			log.Printf("cannot convert range tick %s - %v -\n\t%+v", valLbl[0], err, inp.DynamicFuncParamset)
		}
		rc.xs = append(rc.xs, vl)
		rc.lbls = append(rc.lbls, valLbl[1])
	}

	//
	//
	//
	lu := strings.Split(tokensL1[2], ";") // lower upper
	//
	lu[0] = strings.ReplaceAll(lu[0], ",", ".")
	lowerFlt, err := strconv.ParseFloat(lu[0], 64)
	if err != nil {
		log.Printf("cannot convert lower thresh %s - %v -\n\t%+v", lu[0], err, inp.DynamicFuncParamset)
	}
	rc.LowerThreshold = lowerFlt
	rc.LowerDisplay = lu[1]
	//
	lu[2] = strings.ReplaceAll(lu[2], ",", ".")
	upperFlt, err := strconv.ParseFloat(lu[2], 64)
	if err != nil {
		log.Printf("cannot convert upper thresh %s - %v -\n\t%+v", lu[0], err, inp.DynamicFuncParamset)
	}
	rc.UpperThreshold = upperFlt
	rc.UpperDisplay = lu[3]

}

func (rc rangeConf) SerializeExtendedConfig() string {
	lowerUpper := fmt.Sprintf("%5.3f;%v;%5.3f;%v", rc.LowerThreshold, rc.LowerDisplay, rc.UpperThreshold, rc.UpperDisplay)
	return fmt.Sprintf("%v--%v--%v", rc.CSSType, rc.TicksLabels, lowerUpper)
}

func (rc *rangeConf) lowerUpperAttrs() string {

	lt := fmt.Sprint(rc.LowerThreshold)
	if rc.Min == 0 && rc.LowerThreshold == 0 {
		lt = ""
	}

	ut := fmt.Sprint(rc.UpperThreshold)
	if rc.Max == 0 && rc.UpperThreshold == 0 {
		ut = ""
	}

	return fmt.Sprintf(
		" data-lt='%v' data-ld='%v' data-ut='%v' data-ud='%v' ",
		lt,
		rc.LowerDisplay,
		ut,
		rc.UpperDisplay,
	)
}

var sliderWeeksClosing = rangeConf{
	Min:    3,
	Max:    21,
	Step:   3,
	Suffix: suffixWeeks,

	// Following fields are extended parameters,
	// which must be crammed into
	// DynamicFuncParamset as one string.
	CSSType: "1",

	TicksLabels: `3:<6;6:6;9: ;12:12;15: ;18:18;21:>18`,

	LowerThreshold: 5.9,
	LowerDisplay:   "<6",
	UpperThreshold: 18.1,
	UpperDisplay:   ">18",
}

var sliderPctZeroHundredMiddle = rangeConf{
	Min:    0,
	Max:    100,
	Step:   5,
	Suffix: suffixPercent,
	//
	CSSType:     "2",
	TicksLabels: `0:0;25:25;50:50;75:75;100:100`,
}
var sliderPctZeroHundredWide = rangeConf{
	Min:    0,
	Max:    100,
	Step:   5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:0;10: ;20:20;30: ;40:40;50: ;60:60;70: ;80:80;90: ;100:100`,
}

var sliderPctTwoTen = rangeConf{
	Min:    0,
	Max:    12,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:<2;2:2;3: ;4:4;5: ;6:6;7: ;8:8;9: ;10:10;12:>10`,

	LowerThreshold: 1.9,
	LowerDisplay:   "<2",
	UpperThreshold: 10.1,
	UpperDisplay:   ">10",
}

var sliderPctZeroTwo = rangeConf{
	Min:    0,
	Max:    2.5,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:0;0.5:0.5;1:1;1.5:1.5;2:2;2.5:>2`,

	UpperThreshold: 2.1,
	UpperDisplay:   ">2",
}

var sliderPctThreeTwenty = rangeConf{
	Min:    -2,
	Max:    25,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `-2:<3;3:3;5:5;10:10;15:15;20:20;25:>20`,

	LowerThreshold: 2.9,
	LowerDisplay:   "<3",
	UpperThreshold: 20.1,
	UpperDisplay:   ">20",
}

var sliderPctZeroFour = rangeConf{
	Min:    0,
	Max:    5,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:0;1:1;2:2;3:3;4:4;5:>4`,

	UpperThreshold: 4.1,
	UpperDisplay:   ">4",
}

var sliderPctThreeTwentyfive = rangeConf{
	Min:    -2,
	Max:    30,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `-2:<3;3:3;5:5;10:10;15:15;20:20;25:25;30:>25`,

	LowerThreshold: 2.9,
	LowerDisplay:   "<3",
	UpperThreshold: 25.1,
	UpperDisplay:   ">25",
}

var sliderPctZeroTwentyfive = rangeConf{
	Min:    -5,
	Max:    30,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `-5:<0;0:0;5:5;10:10;15:15;20:20;25:25;30:>25`,

	LowerThreshold: -0.1,
	LowerDisplay:   "<0",
	UpperThreshold: 25.1,
	UpperDisplay:   ">25",
}

var sliderYearsZeroTen = rangeConf{
	Min:    0,
	Max:    12,
	Step:   0.5,
	Suffix: suffixYears,
	//
	CSSType:     "3",
	TicksLabels: `0:0;2:2;4:4;6:6;8:8;10:10;12:>10`,

	UpperThreshold: 10.1,
	UpperDisplay:   ">10",
}

var sliderEBITDA2x10x = rangeConf{
	Min:    0,
	Max:    12,
	Step:   0.5,
	Suffix: suffixEBITDA,
	//
	CSSType:     "3",
	TicksLabels: `0:<2;2:2;4:4;6:6;8:8;10:10;12:>10`,

	LowerThreshold: 1.9,
	LowerDisplay:   "<2",
	UpperThreshold: 10.1,
	UpperDisplay:   ">10",
}

// _0- 50 mn € in  5 mn€ brackets
// 50-100 mn € in 10 mn€ brackets
var sliderEBITDAZero150 = rangeConf{
	Min:    0,
	Max:    200,
	Step:   5,
	Suffix: suffixMillionEuro,
	//
	CSSType: "3",
	// TicksLabels: `0:0;25:25;50:50;75:75;100:100;125:125;150:150;200:>150`,
	TicksLabels: `0:0;25: ;50:50;75: ;100:100;125: ;150:150;200:>150`,

	UpperThreshold: 151,
	UpperDisplay:   ">150",
}

// 0-500mn €in 10mn€ brackets
var sliderEVZeroFiveHundred = rangeConf{
	Min:    0,
	Max:    650,
	Step:   10,
	Suffix: suffixMillionEuro,
	//
	CSSType: "3",
	// TicksLabels: `0:0;50: ;100:100;150: ;200:200;250: ;300:300;350: ;400:400;450: ;500:500;650:>500`,
	TicksLabels: `0:0;50: ;100:100;150: ;200: ;250: ;300:300;350: ;400: ;450: ;500:500;650:>500`,

	UpperThreshold: 501,
	UpperDisplay:   ">500",
}

var sliderOneOnePointFive = rangeConf{
	Min:    0,
	Max:    2.55, // should be 2.5 but rounding stuff
	Step:   0.1,
	Suffix: suffixInvestedCapital,
	//
	CSSType:     "3",
	TicksLabels: `0:0;0.5:0.5;1:1;1.5:1.5;2:2;2.5:>2`,

	UpperThreshold: 2.1,
	UpperDisplay:   ">2",
}

var slider50To100 = rangeConf{
	Min:    30,
	Max:    120,
	Step:   10,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `30:<50;50:50;60:60;70:70;80:80;90:90;100:100;120:>100`,

	LowerThreshold: 49,
	LowerDisplay:   "<50",
	UpperThreshold: 101,
	UpperDisplay:   ">100",
}

var slider1To175 = rangeConf{
	Min:  1,
	Max:  2.0,
	Step: 0.05,
	// Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `1:1;1.25:1.25;1.5:1.5;1.75:1.75;2:>1.75`,

	UpperThreshold: 1.755,
	UpperDisplay:   ">1.75",
}

var slider30To100 = rangeConf{
	Min:    30 - 20,
	Max:    100 + 20,
	Step:   5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `10:<30;30:30;50:50;75:75;100:100;120:>100`,

	LowerThreshold: 29,
	LowerDisplay:   "<30",
	UpperThreshold: 101,
	UpperDisplay:   ">100",
}

var slider1To5 = rangeConf{
	Min:    1 - 1.5,
	Max:    5 + 1.5,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `-0.5:<1;1:1;2:2;3:3;4:4;5:5;6.5:>5`,

	LowerThreshold: 0.9,
	LowerDisplay:   "<1",
	UpperThreshold: 5.1,
	UpperDisplay:   ">5",
}
