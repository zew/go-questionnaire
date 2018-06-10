package generators

import (
	"github.com/zew/go-questionaire/generators/fmt"
	"github.com/zew/go-questionaire/generators/min"
	"github.com/zew/go-questionaire/qst"
)

type genT func() *qst.QuestionaireT

var gens = map[string]genT{
	"fmt": fmt.Create,
	"min": min.Create,
}

// Get returns all questionaire generators
func Get() map[string]genT {
	return gens
}
