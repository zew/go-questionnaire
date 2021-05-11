package pat

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create paternalismus questionnaire
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat")
	q.Survey.Params = params
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}

	q.VersionMax = 16
	// q.AssignVersion = "round-robin"

	var err error

	err = Title(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding title page: %v", err)
	}

	err = Core(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding core pages: %v", err)
	}

	err = PersonalQuestions1(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding core pages: %v", err)
	}

	err = End(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding core pages: %v", err)
	}

	q.AddFinishButtonNextToLast()

	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}

	q.VersionEffective = -2 // re-set after validate

	return &q, nil

}
