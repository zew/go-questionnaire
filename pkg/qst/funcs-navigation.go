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
}

func GermanOnly(q *QuestionnaireT, pageIdx int) bool {
	if q.LangCode != "de" {
		return false
	}
	return true
}
