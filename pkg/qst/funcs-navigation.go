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
	"GermanOnly":  GermanOnly,
	"BIIINow":     BIIINow,
	"BIIILater":   BIIILater,
	"BIIIMeasure": BIIIMeasure,
}

func GermanOnly(q *QuestionnaireT, pageIdx int) bool {
	if q.LangCode != "de" {
		return false
	}
	return true
}
func BIIINow(q *QuestionnaireT, pageIdx int) bool {
	// input[0] is a text element
	// input[0] is an error dyn element
	inp := q.Pages[1].Groups[0].Inputs[2]
	if inp.Response == "now" {
		return true
	}
	return false
}

func BIIILater(q *QuestionnaireT, pageIdx int) bool {
	inp := q.Pages[1].Groups[0].Inputs[2]
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
	inp := q.Pages[10].Groups[0].Inputs[1]
	if inp.Response != "" && inp.Response != "1" {
		return true
	}
	return false
}
