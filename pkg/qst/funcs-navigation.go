package qst

type naviFuncT func(*QuestionnaireT, int) bool

/*
	The navi funcs decide whether or not
	to show a particular page
	in progress bar and buttons previous next.

	Required login characteristics should be transferred to
	the questionnaire during login.
*/
var naviFuncs = map[string]func(*QuestionnaireT, int) bool{
	"GermanOnly": GermanOnly,
	"BIIINow":    BIIINow,
	"BIIILater":  BIIILater,
}

func GermanOnly(q *QuestionnaireT, pageIdx int) bool {
	if q.LangCode != "de" {
		return false
	}
	return true
}
func BIIINow(q *QuestionnaireT, pageIdx int) bool {
	// input[0] is a text element
	inp := q.Pages[1].Groups[0].Inputs[1]
	if inp.Response == "now" {
		// log.Printf(" => branch now; Response is %q", inp.Response)
		return true
	}
	// log.Printf(" =>  branch later;  Response is %q", inp.Response)
	return false
}

// inverse of BIIINow
func BIIILater(q *QuestionnaireT, pageIdx int) bool {
	return !BIIINow(q, pageIdx)
}
