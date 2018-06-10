package generators

import (
	"log"
	"path/filepath"

	"github.com/zew/go-questionaire/generators/fmt"
	"github.com/zew/go-questionaire/generators/min"
	"github.com/zew/go-questionaire/qst"
)

type genT func() *qst.QuestionaireT

var gens = map[string]genT{
	"fmt": fmt.Create,
	"min": min.Create,
}

// Run creates all registered questionaire templates
func Run() {

	for key, fnc := range gens {

		q := fnc()

		fn := filepath.Join(qst.BasePath(), key+".json")
		err := q.Save1(fn)
		if err != nil {
			log.Fatalf("Error saving %v: %v", fn, err)
		}

	}
}
