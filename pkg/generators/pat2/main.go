package pat2

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/generators/pat"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create paternalismus questionnaire with totally diff questions
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat2")
	q.Survey = s
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}

	q.VersionMax = 16
	q.AssignVersion = "round-robin"

	q.ShufflingVariations = 8 // for party affiliation and "Entscheidung 7/8"
	q.PostponeNavigationButtons = 6

	var err error

	err = TitlePat23(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding title pat2 page: %v", err)
	}

	err = pat.PersonalQuestions2(&q, pat.VariableElements{NumberingQuestions: 1, AllMandatory: true, NonGermansOut: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 2: %v", err)
	}

	err = pat.PersonalQuestions1(&q, pat.VariableElements{NumberingQuestions: 14, AllMandatory: true, ZumSchlussOrNunOrNothing: 3})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 1: %v", err)
	}

	err = Part1Intro(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1Intro(): %v", err)
	}

	// core
	err = pat.Part1Entscheidung1bis6(&q, pat.VariableElements{AllMandatory: true, ComprehensionCheck1: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1(): %v", err)
	}

	// core
	err = pat.Part1Frage1(&q, pat.VariableElements{NumberingQuestions: 0, AllMandatory: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1Frage1(): %v", err)
	}

	err = Part2IntroA(&q)
	err = Part2IntroBUndEntscheidung78(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part1Entscheidung78(): %v", err)
	}

	err = Part3Intro(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Intro(): %v", err)
	}

	err = Part3Block12(&q, 0)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Block1(0): %v", err)
	}

	err = Part3Block12(&q, 3)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Block1(3): %v", err)
	}

	err = pat.End(&q, pat.VariableElements{Pop2FinishParagraph: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding core pages: %v", err)
	}

	q.AddFinishButtonNextToLast()

	q.VersionEffective = -2 // must be re-set at the end - after validate

	q.Hyphenize()
	q.ComputeMaxGroups()
	q.SetColspans()

	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}

	q.VersionEffective = -2 // re-set after validate

	return &q, nil

}
