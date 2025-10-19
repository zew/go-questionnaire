package fmtest

import (
	"github.com/zew/go-questionnaire/pkg/qst"
)

func special202511(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2025 && q.Survey.Month == 11
	if !cond {
		return nil
	}

	// page experiment chart 0
	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511_0"
	}

	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202511_1"
	}

	return nil

}
