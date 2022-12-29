package qst

import "log"

// WrappedPageT for creating helper funcs outside
// the package - *temporarily*.
// compare EditPage
type WrappedPageT struct {
	*pageT
}

// WrapPageT is a wrapper for pageT
func WrapPageT(pt *pageT) *WrappedPageT {
	return &WrappedPageT{pt}
}

func init() {
	var qif Q

	qif = Q(nil) // first

	var qTmp *QuestionnaireT
	qif = Q(qTmp) // second
	qif.Version()

	qif = Q(&QuestionnaireT{}) // third

	// extracting a questionnaire from first and second fails
	qExtr, ok := qif.(*QuestionnaireT)
	if !ok {
		log.Fatal("(1) cannot convert qst.Q to Questionnaire")
	}
	qExtr.Version()
}

// Q decouples
type Q interface {

	// Questionnaire level
	UserIDInt() int
	Version() int
	GetLangCode() string

	// does not work, because G interface
	// AddGroupWithInputs([]string)

	ResponseByName(n string) (string, error)
	ErrByName(n string) (string, error)
}
