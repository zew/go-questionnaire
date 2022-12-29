// package compositeif serves to de-cycle the import chain
// between qst.CompositeFuncT
// and the packages references therein.
package qstif

// Q decouples qst.QuestionnaireT from  CompositeFuncT(q...)
type Q interface {

	// Questionnaire level
	UserIDInt() int
	Version() int
	GetLangCode() string

	ResponseByName(n string) (string, error)
	ErrByName(n string) (string, error)
}
