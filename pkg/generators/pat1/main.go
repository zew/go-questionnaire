package pat1

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/generators/pat"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func CorePatMust(q *qst.QuestionnaireT) error {

	var err error

	err = pat.Part1Entscheidung1bis6(q, pat.VariableElements{AllMandatory: true, ComprehensionCheck1: true})
	if err != nil {
		return fmt.Errorf("Error adding Part1(): %v", err)
	}

	err = pat.Part1Frage1(q, pat.VariableElements{NumberingQuestions: 1, AllMandatory: true})
	if err != nil {
		return fmt.Errorf("Error adding Part1Frage1(): %v", err)
	}

	err = pat.Part2(q, pat.VariableElements{ZumXtenTeil: "zweiten", NumberingSections: 7, NumberingQuestions: 2, AllMandatory: true, ComprehensionCheck2: true})
	if err != nil {
		return fmt.Errorf("Error adding Part2(): %v", err)
	}

	err = pat.Part2Frage4(q, pat.VariableElements{NumberingQuestions: 4, AllMandatory: true})
	if err != nil {
		return fmt.Errorf("Error adding Part2Frage4(): %v", err)
	}

	return nil
}

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

	err = pat.Title(&q, false, true)
	if err != nil {
		return nil, fmt.Errorf("Error adding title page: %v", err)
	}

	err = CorePatMust(&q)
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
