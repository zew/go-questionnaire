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

func pdsAssetClass(q *QuestionnaireT, pageIdx int, acIdx int) bool {
	ac := PDSAssetClasses[acIdx]

	// inp := q.Pages[11].Groups[0].Inputs[1]
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
