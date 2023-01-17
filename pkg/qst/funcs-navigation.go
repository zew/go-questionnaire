package qst

import "fmt"

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
//   * visibility for page type 12    - based on specific page11 values
//   * visibility for all pages       - based on page1 values
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
