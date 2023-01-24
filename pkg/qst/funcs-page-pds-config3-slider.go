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

	LowerThreshold float64 // below this value, we lock/jump the value to Min,
	LowerDisplay   string  // display value for min ("<0")

	UpperThreshold float64
	UpperDisplay   string

	UpperLastRegular float64 // last regular upper value - for suppressing the interval
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
	lu := strings.Split(tokensL1[2], ";") // lower upper - four values
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

	lu[4] = strings.ReplaceAll(lu[4], ",", ".")
	upperLastRegular, err := strconv.ParseFloat(lu[4], 64)
	if err != nil {
		log.Printf("cannot convert last regular upper  %s - %v -\n\t%+v", lu[4], err, inp.DynamicFuncParamset)
	}
	rc.UpperLastRegular = upperLastRegular

}

func (rc rangeConf) SerializeExtendedConfig() string {
	lowerUpper := fmt.Sprintf(
		"%5.3f;%v;%5.3f;%v;%5.3f",
		rc.LowerThreshold,
		rc.LowerDisplay,
		rc.UpperThreshold,
		rc.UpperDisplay,
		rc.UpperLastRegular,
	)
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
		" data-lt='%v' data-ld='%v' data-ut='%v' data-ud='%v' data-ulr='%v'  ",
		lt,
		rc.LowerDisplay,
		ut,
		rc.UpperDisplay,
		rc.UpperLastRegular,
	)
}

var range0To100 = rangeConf{
	Min:    0,
	Max:    100,
	Step:   5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:0;10: ;20:20;30: ;40:40;50: ;60:60;70: ;80:80;90: ;100:100`,
}

var range0To10 = rangeConf{
	Min:    0,
	Max:    12,
	Step:   0.5,
	Suffix: suffixYears,
	//
	CSSType:     "3",
	TicksLabels: `0:0;2:2;4:4;6:6;8:8;10:10;12:>10`,

	UpperThreshold:   10.1,
	UpperDisplay:     ">10",
	UpperLastRegular: 10.0,
}

var range2To10 = rangeConf{
	Min:    0,
	Max:    12,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:<2;2:2;3: ;4:4;5: ;6:6;7: ;8:8;9: ;10:10;12:>10`,

	LowerThreshold:   1.9,
	LowerDisplay:     "<2",
	UpperThreshold:   10.1,
	UpperDisplay:     ">10",
	UpperLastRegular: 10.0,
}

var range0To2a = rangeConf{
	Min:    0,
	Max:    2.5,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:0;0.5:0.5;1:1;1.5:1.5;2:2;2.5:>2`,

	UpperThreshold:   2.1,
	UpperDisplay:     ">2",
	UpperLastRegular: 2.0,
}

// different stepping
var range0To2b = rangeConf{
	Min:    0,
	Max:    2.55, // should be 2.5 but rounding stuff
	Step:   0.1,
	Suffix: suffixInvestedCapital,
	//
	CSSType:     "3",
	TicksLabels: `0:0;0.5:0.5;1:1;1.5:1.5;2:2;2.5:>2`,

	UpperThreshold:   2.1,
	UpperDisplay:     ">2",
	UpperLastRegular: 2.0,
}

var range3To20 = rangeConf{
	Min:    -2,
	Max:    25,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType: "3",
	// TicksLabels: `-2:<3;3:3;5:5;10:10;15:15;20:20;25:>20`,
	TicksLabels: `-2:<3;3: ;5:5;10:10;15:15;20:20;25:>20`,

	LowerThreshold:   2.9,
	LowerDisplay:     "<3",
	UpperThreshold:   20.1,
	UpperDisplay:     ">20",
	UpperLastRegular: 20.0,
}

var range3To25 = rangeConf{
	Min:    -2,
	Max:    30,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType: "3",
	// TicksLabels: `-2:<3;3:3;5:5;10:10;15:15;20:20;25:25;30:>25`,
	TicksLabels: `-2:<3;3: ;5:5;10:10;15:15;20:20;25:25;30:>25`,

	LowerThreshold:   2.9,
	LowerDisplay:     "<3",
	UpperThreshold:   25.1,
	UpperDisplay:     ">25",
	UpperLastRegular: 25.0,
}

var range0To25 = rangeConf{
	Min:    -5,
	Max:    30,
	Step:   0.5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `-5:<0;0:0;5:5;10:10;15:15;20:20;25:25;30:>25`,

	LowerThreshold:   -0.1,
	LowerDisplay:     "<0",
	UpperThreshold:   25.1,
	UpperDisplay:     ">25",
	UpperLastRegular: 25.0,
}

var range0To4 = rangeConf{
	Min:    0,
	Max:    5,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `0:0;1:1;2:2;3:3;4:4;5:>4`,

	UpperThreshold:   4.1,
	UpperDisplay:     ">4",
	UpperLastRegular: 4.0,
}

var range1To5 = rangeConf{
	Min:    1 - 1.5,
	Max:    5 + 1.5,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `-0.5:<1;1:1;2:2;3:3;4:4;5:5;6.5:>5`,

	LowerThreshold:   0.9,
	LowerDisplay:     "<1",
	UpperThreshold:   5.1,
	UpperDisplay:     ">5",
	UpperLastRegular: 5.0,
}

var rangeEBITDA2x10x = rangeConf{
	Min:    0,
	Max:    12,
	Step:   0.5,
	Suffix: suffixEBITDA,
	//
	CSSType:     "3",
	TicksLabels: `0:<2;2:2;4:4;6:6;8:8;10:10;12:>10`,

	LowerThreshold:   1.9,
	LowerDisplay:     "<2",
	UpperThreshold:   10.1,
	UpperDisplay:     ">10",
	UpperLastRegular: 10.0,
}

// _0- 50 mn € in  5 mn€ brackets
// 50-100 mn € in 10 mn€ brackets
var rangeEBITDAZero150 = rangeConf{
	Min:    0,
	Max:    200,
	Step:   5,
	Suffix: suffixMillionEuro,
	//
	CSSType:     "3",
	TicksLabels: `0:0;25: ;50:50;75: ;100:100;125: ;150:150;200:>150`,

	UpperThreshold:   151,
	UpperDisplay:     ">150",
	UpperLastRegular: 150.0,
}

// 0-500mn €in 10mn€ brackets
var rangeEV0To500 = rangeConf{
	Min:    0,
	Max:    650,
	Step:   10,
	Suffix: suffixMillionEuro,
	//
	CSSType: "3",
	// TicksLabels: `0:0;50: ;100:100;150: ;200:200;250: ;300:300;350: ;400:400;450: ;500:500;650:>500`,
	TicksLabels: `0:0;50: ;100:100;150: ;200: ;250: ;300:300;350: ;400: ;450: ;500:500;650:>500`,

	UpperThreshold:   501,
	UpperDisplay:     ">500",
	UpperLastRegular: 500.0,
}

var range50To100 = rangeConf{
	Min:    30,
	Max:    120,
	Step:   10,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `30:<50;50:50;60:60;70:70;80:80;90:90;100:100;120:>100`,

	LowerThreshold:   49,
	LowerDisplay:     "<50",
	UpperThreshold:   101,
	UpperDisplay:     ">100",
	UpperLastRegular: 100.0,
}

var range1To175 = rangeConf{
	Min:  1,
	Max:  2.0,
	Step: 0.05,
	// Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `1:1;1.25:1.25;1.5:1.5;1.75:1.75;2:>1.75`,

	UpperThreshold:   1.755,
	UpperDisplay:     ">1.75",
	UpperLastRegular: 1.75,
}

var range30To100 = rangeConf{
	Min:    30 - 20,
	Max:    100 + 20,
	Step:   5,
	Suffix: suffixPercent,
	//
	CSSType:     "3",
	TicksLabels: `10:<30;30:30;50:50;75:75;100:100;120:>100`,

	LowerThreshold:   29,
	LowerDisplay:     "<30",
	UpperThreshold:   101,
	UpperDisplay:     ">100",
	UpperLastRegular: 100.0,
}

var range0To175 = rangeConf{
	Min:    0,
	Max:    2.0,
	Step:   0.25,
	Suffix: suffixPercent,
	//
	CSSType: "3",
	// TicksLabels: `0:0;0.5:0.5;1:1;1.5:1.5;1.75:1.75;2.0:>1.75`,
	TicksLabels: `0:0-0.25;  0.25:0.25;  1.25:1.25;  2:>1.75`,

	UpperThreshold:   1.76,
	UpperDisplay:     ">1.75",
	UpperLastRegular: 1.75,
}
