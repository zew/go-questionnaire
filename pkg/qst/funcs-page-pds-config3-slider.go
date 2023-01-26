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
	CXLabelTick    string // cx value: label: tick
	Lower, Upper   string

	Suffix trl.S

	CSSType string // CSS file containing variable

	xs []float64 // range step, where label or tick should appear
	ls []string  // label under tick; can be empty
	ts []string  // draw tick or not

	cxMin, cxMax float64
}

func (rc rangeConf) SerializeExtendedConfig() string {

	lowerUpper := fmt.Sprintf(
		"%v;%v",
		rc.Lower,
		rc.Upper,
	)

	return fmt.Sprintf(
		"%v--%v--%v",
		rc.CSSType,
		rc.CXLabelTick,
		lowerUpper,
	)
}

func (rc *rangeConf) New(inp *inputT) {

	if rc == nil {
		*rc = rangeConf{}
		rc.xs = []float64{}
		rc.ls = []string{}
	}

	rc.Min = inp.Min
	rc.Max = inp.Max
	rc.Step = inp.Step

	rc.Suffix = inp.Suffix
	if rc.Suffix == nil {
		log.Printf("invalid range config string - no suffix %s\n\t%s", inp.Name, inp.DynamicFuncParamset)
	}

	tokensL1 := strings.Split(inp.DynamicFuncParamset, "--") // tokens level 1
	if len(tokensL1) != 3 {
		log.Printf("invalid range config string %s\n\t%s", inp.Name, inp.DynamicFuncParamset)
		return
	}

	//
	//
	//
	rc.CSSType = tokensL1[0]
	rc.CSSType = "3"

	//
	//
	//
	rc.CXLabelTick = tokensL1[1]
	scalePartsStr := strings.Split(rc.CXLabelTick, ";")
	for _, scaleEl := range scalePartsStr {
		triple := strings.Split(scaleEl, ":")

		triple[0] = strings.ReplaceAll(triple[0], ",", ".")
		triple[0] = strings.TrimSpace(triple[0]) // leading space causes ParseFloat to return an
		cxFl, err := strconv.ParseFloat(triple[0], 64)
		if err != nil {
			log.Printf("cannot convert range tick %s - %v -\n\t%+v", triple[0], err, inp.DynamicFuncParamset)
		}

		rc.xs = append(rc.xs, cxFl)

		triple[1] = strings.ReplaceAll(triple[1], ">=", "≥")
		rc.ls = append(rc.ls, triple[1])
		rc.ts = append(rc.ts, strings.TrimSpace(triple[2]))
	}

	//
	//
	//
	lu := strings.Split(tokensL1[2], ";") // lower upper - four values
	if len(lu) != 2 {
		log.Printf("lower upper invalid  %s \n\t%+v", inp.Name, inp.DynamicFuncParamset)
		return
	}
	rc.Lower = lu[0]
	rc.Upper = lu[1]

	last := len(rc.xs) - 1

	if rc.Lower == "" {
		rc.Lower = rc.ls[0]
	}
	rc.cxMin = rc.xs[0]

	if rc.Upper == "" {
		rc.Upper = rc.ls[last]
	}
	rc.cxMax = rc.xs[last]
	//

}

func (rc *rangeConf) lowerUpperAttrs() string {
	return fmt.Sprintf(
		"  data-ld='%v' data-lfr='%.3f'       data-ud='%v' data-ulr='%.3f'  ",
		rc.Lower,
		rc.cxMin,
		//
		rc.Upper,
		rc.cxMax,
	)
}

//
//
//

var (
	range0Too10 = rangeConf{
		Min:         0,
		Max:         10,
		Step:        0.5,
		CXLabelTick: `0:0:t; 1: :t; 2:2:t; 3: :t; 4:4:t; 5: :t; 6:6:t; 7: :t; 8:8:t; 9: :t; 10:>=10:t`,
		Suffix:      suffixYears,
	}

	range0To150 = rangeConf{
		Min:         0,
		Max:         150,
		Step:        5,
		CXLabelTick: `0:0:t; 25: :t; 50:50:t; 75: :t; 100:100:t; 125: :t; 150:>=150:t`,
		Suffix:      suffixMillionEuro,
	}

	range0To2 = rangeConf{
		Min:         0,
		Max:         2.0,
		Step:        0.1,
		CXLabelTick: `0:0:t; 0.5:0.5:t; 1:1:t; 1.5:1.5:t; 2:>=2:t`,
		Suffix:      suffixInvestedCapital,
	}

	range0Too375 = rangeConf{
		Min:         -0.25,
		Max:         3.75,
		Step:        0.5,
		CXLabelTick: `-0.25:0:t; 0.25: :t; 0.75:0.75:t; 1.25: :t; 1.75:1.75:t; 2.25: :t; 2.75:2.75:t; 3.25: :t; 3.75:>=3.75:t`,
		Lower:       "0 - <0.25",
		Suffix:      suffixPercent,
	}

	range0To500 = rangeConf{
		Min:         0,
		Max:         500,
		Step:        10,
		CXLabelTick: `0:0:t; 50: :t; 100:100:t; 150: :t; 200:200:t; 250: :t; 300:300:t; 350: :t; 400:400:t; 450: :t; 500:>=500:t`,
		Suffix:      suffixMillionEuro,
	}

	rangeMin05To25 = rangeConf{
		Min:         -0.5,
		Max:         25.0,
		Step:        0.5,
		CXLabelTick: `-0.5:<0:nt; 5:5:t; 10:10:t; 15:15:t; 20:20:t; 25:>=25:t`,
		Lower:       "<0",
		Suffix:      suffixPercent,
	}

	range075to5a = rangeConf{
		Min:  0.5,
		Max:  5,
		Step: 0.5,
		// CXLabelTick: `0.75:<1:nt; 1: :t; 1.5: :t; 2:2:t; 2.5: :t; 3:3:t; 3.5: :t; 4:4:t; 4.5: :t; 5:>=5:t`,
		CXLabelTick: `0.5:<1:nt;   1.0: :t;   1.5: :t;   2.0:2:t;   2.5: :t;   3.0:3:t;   3.5: :t;   4.0:4:t;   4.5: :t;   5.0:>=5:t`,
		Lower:       "<1",
		Suffix:      suffixDebtService, // only difference to range075to5b
	}

	range075to5b = rangeConf{
		Min:  0.5,
		Max:  5,
		Step: 0.5,
		// CXLabelTick: `0.75:<1:nt; 1: :t; 1.5: :t; 2:2:t; 2.5: :t; 3:3:t; 3.5: :t; 4:4:t; 4.5: :t; 5:>=5:t`,
		CXLabelTick: `0.5:<1:nt;   1.0: :t;   1.5: :t;   2.0:2:t;   2.5: :t;   3.0:3:t;   3.5: :t;   4.0:4:t;   4.5: :t;   5.0:>=5:t`,
		Lower:       "<1",
		Suffix:      suffixInterestPayment, // only difference to range075to5a
	}

	range1to175 = rangeConf{
		Min:         1,
		Max:         1.75 + 0.01, // rounding trouble
		Step:        0.05,
		CXLabelTick: `1:1:t; 1.25:1.25:t; 1.5:1.5:t; 1.75:>=1.75:t`,
		Suffix:      suffixDebtService, // only difference to range075to5a
		Upper:       "≥1.75",
		// Upper:       ">=1.75",
	}

	range2to10a = rangeConf{
		Min:         1.5,
		Max:         10,
		Step:        0.5,
		CXLabelTick: `1.5:<2:nt; 2: :t; 3:3:t; 4: :t; 5: :t; 5.5:5.5:nt; 6: :t; 7: :t; 8:8:t; 9: :t; 10:>=10:t`,
		Lower:       "<2",
		Suffix:      suffixPercent,
	}
	range2to10b = rangeConf{
		Min:         1.5,
		Max:         10,
		Step:        0.5,
		CXLabelTick: `1.5:<2:nt; 2: :t; 3:3:t; 4: :t; 5: :t; 5.5:5.5:nt; 6: :t; 7: :t; 8:8:t; 9: :t; 10:>=10:t`,
		Lower:       "<2",
		Suffix:      suffixEBITDA,
	}

	// badly named
	range25to200 = rangeConf{
		Min:         2,
		Max:         20,
		Step:        0.5,
		CXLabelTick: `2:<2.5:nt; 5:5:t; 10:10:t; 15:15:t; 20:>=20:t`,
		Suffix:      suffixPercent,
	}

	// badly named
	range25to250 = rangeConf{
		Min:         2,
		Max:         25,
		Step:        0.5,
		CXLabelTick: `2:<2.5:nt; 5:5:t; 10:10:t; 15:15:t; 20:20:t; 25:>=25:t`,
		Suffix:      suffixPercent,
	}

	range20to100 = rangeConf{
		Min:         20,
		Max:         95,
		Step:        5,
		CXLabelTick: `20:<25:nt; 25: :t; 50:50:t; 75:75:t; 95:100:t`,
		Suffix:      suffixPercent,
	}

	range50to100 = rangeConf{
		Min:         45,
		Max:         95,
		Step:        5,
		CXLabelTick: `45:<50:nt; 75:75:t; 95:100:t`,
		Suffix:      suffixPercent,
	}
)
