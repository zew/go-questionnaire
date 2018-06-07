package generators

import (
	"log"
	"path/filepath"

	"github.com/zew/go-questionaire/generators/fmt"
	"github.com/zew/go-questionaire/generators/min"
	"github.com/zew/go-questionaire/qst"
)

type genT func() *qst.QuestionaireT

var gens = map[string]genT{}

// Run creates all registered questionaire templates
func Run() {

	gens["fmt"] = fmt.Create
	gens["min"] = min.Create

	for key, fnc := range gens {

		q := fnc()

		fn := filepath.Join(".", "generators", key+".json")

		err := q.Save(fn)
		if err != nil {
			log.Fatalf("Error saving %v: %v", fn, err)
		}

	}
}
