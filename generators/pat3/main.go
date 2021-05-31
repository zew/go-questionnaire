package pat3

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/generators/pat"
	"github.com/zew/go-questionnaire/generators/pat2"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create paternalismus questionnaire with totally diff questions
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat3")
	q.Survey.Params = params
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}

	q.VersionMax = 16
	q.AssignVersion = "round-robin"

	q.ShufflingsMax = 8 // for party affiliation and "Entscheidung 7/8"
	q.PostponeNavigationButtons = 6

	var err error

	err = pat2.TitlePat23(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding TitlePat23(): %v", err)
	}

	err = pat.PersonalQuestions2(&q, pat.VariableElements{NumberingQuestions: 1, AllMandatory: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 2: %v", err)
	}

	err = pat.PersonalQuestions1(&q, pat.VariableElements{NumberingQuestions: 14, AllMandatory: true, ZumSchlussOrNunOrNothing: 3})
	if err != nil {
		return nil, fmt.Errorf("Error adding personal questions 1: %v", err)
	}

	// core
	err = pat.Part2(&q,
		pat.VariableElements{
			NumberingSections:     1,
			NumberingQuestions:    1,
			AllMandatory:          true,
			ZumXtenTeil:           "1",
			ZumErstenTeilAsNumber: true,
		},
	)
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2(): %v", err)
	}

	err = ComprehensionCheckPop3(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding ComprehensionCheckPop3(): %v", err)
	}

	// core
	err = pat.Part2Frage4(&q, pat.VariableElements{NumberingQuestions: 3, AllMandatory: true})
	if err != nil {
		return nil, fmt.Errorf("Error adding Part2Frage4(): %v", err)
	}

	//
	//
	err = POP3Part1Intro(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part1Intro(): %v", err)
	}

	err = POP3Part1Decision34(&q, 3, "dec3")
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part1Decision34(3): %v", err)
	}

	err = POP3Part1Decision34(&q, 4, "dec4")
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part1Decision34(4): %v", err)
	}

	err = POP3Part2Intro(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part2Intro(): %v", err)
	}

	err = POP3Part2Questions123and456(&q, 1)
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part2Questions123and456(): %v", err)
	}

	err = POP3Part2Questions123and456(&q, 4)
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part2Questions123and456(): %v", err)
	}

	err = POP3Part2Questions78(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding POP3Part2Questions78(): %v", err)
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
