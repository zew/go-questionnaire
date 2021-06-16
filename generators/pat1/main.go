package pat1

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create paternalismus questionnaire with addtl pers questions
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat1")
	q.Survey = s
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}

	q.VersionMax = 16
	q.AssignVersion = "round-robin"

	q.ShufflingVariations = 8 // for party affiliation and "Entscheidung 7/8"

	var err error

	// err = pat.Title(&q, true)
	err = pat.Title(&q, false)
	if err != nil {
		return nil, fmt.Errorf("Error adding title page: %v", err)
	}

	err = pat.Core(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding core pages: %v", err)
	}

	err = pat.PersonalQuestions1(&q, pat.VariableElements{NumberingQuestions: 5, ZumSchlussOrNunOrNothing: 2, AllMandatory: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 1: %v", err)
	}

	err = pat.PersonalQuestions2(&q, pat.VariableElements{NumberingQuestions: 8, AllMandatory: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 2: %v", err)
	}

	err = pat.End(&q, pat.VariableElements{})
	if err != nil {
		return nil, fmt.Errorf("Error adding core pages: %v", err)
	}

	q.AddFinishButtonNextToLast()

	q.VersionEffective = -2 // must be re-set at the end - after validate

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
