package pat0

import (
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/qst"
)

// Create for PAT but with zentrally distributed versions.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	// q := &qst.QuestionnaireT{}

	q, err := pat.Create(params)
	if err != nil {
		return q, err

	}

	q.AssignVersion = "round-robin"
	q.VersionEffective = -2 // must be re-set at the end - after validate

	return q, nil

}
