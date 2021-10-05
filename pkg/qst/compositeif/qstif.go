// package compositeif serves to de-cycle the import chain
// between qst.CompositeFuncT
// and the packages references therein.
package compositeif

// Q decouples qst.QuestionnaireT from  CompositeFuncT(q...)
type Q interface {
	UserIDInt() int
	Version() int
	ResponseByName(n string) (string, error)
	ErrByName(n string) (string, error)
}
