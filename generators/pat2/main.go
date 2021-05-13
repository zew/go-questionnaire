package pat2

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create paternalismus questionnaire with totally diff questions
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat2")
	q.Survey.Params = params
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}

	q.VersionMax = 16
	// q.AssignVersion = "round-robin"

	var err error

	err = TitlePat2(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding title pat2 page: %v", err)
	}

	err = pat.PersonalQuestions2(&q, pat.VariableElements{NumberingStart: 1})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 2: %v", err)
	}

	err = pat.PersonalQuestions1(&q, pat.VariableElements{NumberingStart: 9})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 1: %v", err)
	}

	err = Part1Intro(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1Intro(): %v", err)
	}

	// core
	err = pat.Part1(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1(): %v", err)
	}

	// core
	err = pat.Part1Frage1(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1Frage1(): %v", err)
	}

	err = Part1Entscheidung78(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1Entscheidung78(): %v", err)
	}

	err = Part2Intro(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Intro(): %v", err)
	}

	err = Part2Block12(&q, 0)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Block1(0): %v", err)
	}

	err = Part2Block12(&q, 3)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Block1(3): %v", err)
	}

	err = pat.End(&q)
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
