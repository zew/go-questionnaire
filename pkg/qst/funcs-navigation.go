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

	"fmt202511Include": fmt202511Include,
	"fmt202512Include": fmt202512Include,

	"fmtAgreeTerms": fmtAskConsent_GDPR,
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

func fmt202511Include(q *QuestionnaireT, pageIdx int) bool {
	_, found := ForecastData(q.UserIDInt())
	return found
}

func fmt202512Include(q *QuestionnaireT, pageIdx int) bool {
	inp := q.ByName("ssq5")
	if inp.Response == "" {
		return false
	}
	if inp.Response != "neutral" {
		return true
	}
	return false
}

var askForConsent = map[int]bool{

	9990: true,
	9991: false,
	9992: true,  // already agreed, do not include
	9993: false, // already agreed, do not include

	10003: false,
	10009: false,
	10014: false,
	10015: false,
	10016: false,
	10017: false,
	10021: false,
	10023: false,
	10025: false,
	10033: false,
	10034: false,
	10035: false,
	10040: false,
	10056: false,
	10058: false,
	10062: false,
	10063: false,
	10070: false,
	10073: false,
	10079: false,
	10080: false,
	10084: false,
	10086: false,
	10089: false,
	10090: false,
	10095: false,
	10105: false,
	10115: false,
	10129: false,
	10133: false,
	10134: false,
	10140: false,
	10143: false,
	10146: false,
	10147: false,
	10150: false,
	10154: false,
	10161: false,
	10162: false,
	10163: false,
	10172: false,
	10178: false,
	10179: false,
	10180: false,
	10185: false,
	10205: false,
	10209: false,
	10210: false,
	10231: false,
	10235: false,
	10263: false,
	10267: false,
	10268: false,
	10274: false,
	10278: false,
	10307: false,
	10315: false,
	10330: false,
	10343: false,
	10344: false,
	10345: false,
	10364: false,
	10366: false,
	10367: false,
	10369: false,
	10372: false,
	10374: false,
	10377: false,
	10381: false,
	10385: false,
	10391: false,
	10418: false,
	10420: false,
	10421: false,
	10802: false,
	10806: false,
	10812: false,
	10813: false,
	10828: false,
	10830: false,
	10844: false,
	10947: false,
	11241: false,
	11246: false,
	11272: false,
	11275: false,
	11400: false,
	11425: false,
	11435: false,
	11445: false,
	11452: false,
	11465: false,
	11478: false,
	11482: false,
	11499: false,
	11500: false,
	11514: false,
	11558: false,
	11565: false,
	11568: false,
	11569: false,
	11575: false,
	11588: false,
	11589: false,
	11600: false,
	11602: false,
	11603: false,
	11605: false,
	11606: false,
	11607: false,
	11614: false,
	11616: false,
	11619: false,
	11622: false,
	11628: false,
	11632: false,
	11633: false,
	11635: false,
	11639: false,
	11640: false,
	11642: false,
	11643: false,
	11650: false,
	11653: false,
	11659: false,
	11661: false,
	11673: false,
	11677: false,
	11680: false,
	11682: false,
	11686: false,
	11687: false,
	11689: false,
	11692: false,
	11693: false,
	11698: false,
	11699: false,
	11702: false,
	11703: false,
	11705: false,
	11708: false,
	11709: false,
	11710: false,
	11712: false,
	11714: false,
	11715: false,
	11716: false,
	11717: false,
	11718: false,
	11719: false,
	11720: false,
	11721: false,
	11726: false,
	11728: false,
	11729: false,
	11730: false,
	11732: false,
	11739: false,
	11740: false,
	11743: false,
	11750: false,
	11753: false,
	11755: false,
	11756: false,
	11758: false,
	11760: false,
	11765: false,
	11767: false,
	11768: false,
	11769: false,
	11770: false,
	11771: false,
	11773: false,
	11774: false,
	11777: false,
	11779: false,
	11781: false,
	11783: false,
	11786: false,
	11787: false,
}

// fmtAskConsent_GDPR - which users get to see the "agree to terms and conditions" page
func fmtAskConsent_GDPR(q *QuestionnaireT, pageIdx int) bool {

	for i := 9970; i < 10000; i++ {
		if i%2 == 0 {
			askForConsent[i] = true
		} else {
			askForConsent[i] = false
		}
	}

	askAgain := map[int]bool{
		11653: true,
		11709: true,
	}

	//
	//
	userId := q.UserIDInt()

	_, ok1 := askAgain[userId]
	if ok1 {
		// ask again
		return true
	}

	val, ok2 := askForConsent[userId]
	if !ok2 {
		// not in list - not agreed yet - include
		return true
	}
	return val

}
