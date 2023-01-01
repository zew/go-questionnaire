package qst

import (
	"log"

	"github.com/zew/go-questionnaire/pkg/qst/internal/x"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// ============ internal package ==========
/*
	we could move the inner implementation into an internal package;
	thus only qst could make use of them
*/
func init() {
	_ = &x.QuestionnaireT{}
	_ = &x.PageT{}
}

//
// ============ exported type ==========
//

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

//
// ============ interface ==========
//
/* 	This is a study of creating an interface for QuestionnaireT, pageT, groupT, inputT.
Generators would only use the interface.
Changes in package qst would not require recompilation of all generator packages.
*/

//
// special methods for interface

// AddPageIf like AddPage
// compare WrapPage
func (q *QuestionnaireT) AddPageIf() P {
	return q.AddPage()
}

// SetSection corresponds to the struct field Section.
// Setters, and some getters for every struct field required
func (p *pageT) SetSection(s trl.S) {
	p.Section = s
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

type P interface {
	// Setters, and some getters for every struct field are required.
	// We could use reflection:
	//
	//     SetField(name string, value interface{})
	//
	// sacrificing type safety.
	//
	// We could use AST code generation to
	// create setters and getters like in Java.
	// => 15 pairs for pageT
	// => 25 pairs for inputT
	SetSection(trl.S)

	WidthMax(string) // setting max width
}

// Q decouples
type Q interface {

	// Questionnaire level
	UserIDInt() int
	Version() int
	GetLangCode() string

	// AddPage() *pageT
	AddPageIf() P

	// does not work, because G interface
	// AddGroupWithInputs([]string)

	ResponseByName(n string) (string, error)
	ErrByName(n string) (string, error)
}
