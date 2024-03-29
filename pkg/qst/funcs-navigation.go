package qst

import (
	"fmt"
	"strconv"
	"time"
)

type naviFuncT func(*QuestionnaireT, int) bool

/*
The navi funcs decide whether or not
to show a particular page
in progress bar and buttons previous next.

Required login characteristics should be transferred to
the questionnaire during login.
*/
var naviFuncs = map[string]func(*QuestionnaireT, int) bool{
	"GermanOnly":  GermanOnly,
	"BIIINow":     BIIINow,
	"BIIILater":   BIIILater,
	"BIIIMeasure": BIIIMeasure,

	"pds_ac1": pdsAssetClass1,
	"pds_ac2": pdsAssetClass2,
	"pds_ac3": pdsAssetClass3,

	"kneb_t1a":               knebTreatment1NeutraVsFinance_A,
	"kneb_t1b":               knebTreatment1NeurtraVsFinance_B,
	"kneb_t2a":               knebTreatment2AdviceNoOrYes_A,
	"kneb_t2b":               knebTreatment2AdviceNoOrYes_B,
	"kneb_d7_unemployed":     knebD7unemployed,
	"kneb_d7_employed":       knebD7employed,
	"kneb_too_old":           knebTooOld,
	"kneb_b6_who_competent":  knebB6WhoIsCompetent,
	"kneb_h1_who_responsibe": knebH1WhoIsResponsible,
}

func GermanOnly(q *QuestionnaireT, pageIdx int) bool {
	if q.LangCode != "de" {
		return false
	}
	return true
}
func BIIINow(q *QuestionnaireT, pageIdx int) bool {
	inp := q.Pages[2].Groups[0].Inputs[2]
	if inp.Response == "now" {
		return true
	}
	return false
}

func BIIILater(q *QuestionnaireT, pageIdx int) bool {
	inp := q.Pages[2].Groups[0].Inputs[2]
	if inp.Response != "" && inp.Response != "now" {
		return true
	}
	return false
}

func BIIIMeasure(q *QuestionnaireT, pageIdx int) bool {
	if BIIINow(q, pageIdx) == false {
		return false
	}
	// q20 - we measure impact of our investments
	inp := q.Pages[11].Groups[0].Inputs[1]
	if inp.Response != "" && inp.Response != "1" {
		return true
	}
	return false
}

func pdsAssetClass1(q *QuestionnaireT, pageIdx int) bool {
	return pdsAssetClass(q, pageIdx, 0)
}
func pdsAssetClass2(q *QuestionnaireT, pageIdx int) bool {
	return pdsAssetClass(q, pageIdx, 1)
}
func pdsAssetClass3(q *QuestionnaireT, pageIdx int) bool {
	return pdsAssetClass(q, pageIdx, 2)
}

// pdsAssetClass governs
//   - visibility for page type 12    - based on specific page11 values
//   - visibility for all pages       - based on page1 values
func pdsAssetClass(q *QuestionnaireT, pageIdx int, acIdx int) bool {

	// special rule for page12
	//  if (all number of transactions for tranche types are "0")
	//  then skip page12
	// depends on  setting
	//    `page.CounterProgress = "page12"`
	page := q.Pages[pageIdx]
	if page.CounterProgress == "page12" {
		// ac1_tt2_q11a_numtransact_main
		tts := PDSAssetClasses[acIdx].TrancheTypes
		allNull := true
		for _, tt := range tts {
			name := fmt.Sprintf("ac%v_%v_q11a_numtransact_main", acIdx+1, tt.Prefix)
			inp := q.ByName(name)
			if inp == nil { // page not initialized
				break
			}
			if inp.Response == "" || inp.Response != "0" {
				allNull = false
				break
			}
		}
		if allNull {
			return false
		}

	}

	//
	// visibility; depending on selection on page1
	ac := PDSAssetClasses[acIdx]
	// inp := q.Pages[1].Groups[xxx].Inputs[yyy]
	name := fmt.Sprintf("%v_q03", ac.Prefix)
	inp := q.ByName(name)
	if inp.Response != "" && inp.Response != "0" {
		// at least one tranchetype must be selected
		for _, tt := range ac.TrancheTypes {
			subName := fmt.Sprintf("%v_%v_q031", ac.Prefix, tt.Prefix)
			subInp := q.ByName(subName)
			if subInp.Response != "" && subInp.Response != "0" {
				return true
			}
		}
		// asset class selected, but not a single tranche type
		return false
	}
	return false
}

func knebTreatment1NeutraVsFinance_A(q *QuestionnaireT, pageIdx int) bool {
	if q.Version()%2 == 0 {
		return true
	}
	return false
}
func knebTreatment1NeurtraVsFinance_B(q *QuestionnaireT, pageIdx int) bool {
	return !knebTreatment1NeutraVsFinance_A(q, pageIdx)
}

// 1000,1001  => 500 =>  0 => false
// 1002,1003, => 501 =>  1 => true
func knebTreatment2AdviceNoOrYes_A(q *QuestionnaireT, pageIdx int) bool {
	id := q.Version()
	id = id / 2
	if id%2 == 0 {
		return false
	}
	return true
}
func knebTreatment2AdviceNoOrYes_B(q *QuestionnaireT, pageIdx int) bool {
	return !knebTreatment2AdviceNoOrYes_A(q, pageIdx)
}

func knebTooOld(q *QuestionnaireT, pageIdx int) bool {

	inp := q.ByName("qd2_birthyear")
	if inp.Response != "" {
		yrBirth, _ := strconv.Atoi(inp.Response)
		tolerance := 1
		tooYng := time.Now().Year()-yrBirth < 18-tolerance
		tooOld := time.Now().Year()-yrBirth > 55+tolerance
		if tooYng || tooOld {
			return true
		}
	}
	return false
}

func knebD7employed(q *QuestionnaireT, pageIdx int) bool {
	inp := q.ByName("qd7_employment")
	employed := inp.Response == "above35hours" || inp.Response == "between15and35hours"
	if employed {
		return true
	}
	return false
}
func knebD7unemployed(q *QuestionnaireT, pageIdx int) bool {
	inp := q.ByName("qd7_employment")
	employed := inp.Response == "above35hours" || inp.Response == "between15and35hours"
	if !employed && inp.Response != "" {
		return true
	}
	return false
}

func knebB6WhoIsCompetent(q *QuestionnaireT, pageIdx int) bool {
	inp := q.ByName("qb5_delegate")
	if inp.Response == "yes" {
		return true
	}
	return false
}
func knebH1WhoIsResponsible(q *QuestionnaireT, pageIdx int) bool {
	inp := q.ByName("qd5_family_status")
	if inp.Response == "unmarried_livingtogether" || inp.Response == "married_livingtogether" || inp.Response == "divorcedwidowed_livingtogether" {
		return true
	}
	return false
}
