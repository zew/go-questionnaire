package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage11AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage11(q, page, 0)
}
func pdsPage11AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage11(q, page, 1)
}
func pdsPage11AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage11(q, page, 2)
}

func pdsPage11(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]
	ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.ValidationFuncName = "pdsPage11-a,pdsPage11-b"

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: &nbsp;&nbsp;  Portfolio changes  (past 3&nbsp;months)", ac.Lbl["en"]),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp;  Portfolio changes  (past 3&nbsp;months)", ac.Lbl["de"]),
	}.Outline(fmt.Sprintf("%c1.", rn))
	page.Short = trl.S{
		"en": "Portfolio<br>changes",
		"de": "Portfolio<br>changes",
	}
	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>changes", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>changes", ac.Short["de"]),
	}

	page.CounterProgress = fmt.Sprintf("%c1", rn)

	page.WidthMax("58rem")
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("42rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("38rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"en": "New transactions",
				"de": "New transactions",
			}.Outline("1.1")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	restrictedTextMultiCols(page, ac, rT11a)

	restrictedTextMultiCols(page, ac, rT11b)

	// c) Average time to close a transaction
	lblDuration := trl.S{
		// "en": "Average time to close a deal in weeks (across all tranche types)",
		"en": "How long does it take on average to close a transaction (across all tranche types)?",
		"de": "Durchschnittl. Zeit bis Abschluss in Wochen",
	}.Outline("c.)")

	/*
		dropdownsLabelsTop(
			page,
			ac,
			"q11c_closing_time",
			lblDuration,
			mCh5,
		)
	*/

	// colspan 1/2/3
	radiosSingleRow(
		page,
		ac,
		"q11c_closing_time",
		lblDuration,
		mCh5,
	)

	if acIdx == 0 {
		restrictedTextMultiCols(page, ac, rT11dCorpLend)
	} else {
		restrictedTextMultiCols(page, ac, rT11dRealEstate)
	}

	restrictedTextMultiCols(page, ac, rT11e)

	if acIdx == 0 {
		restrictedTextMultiCols(page, ac, rT11fCorpLend)
	} else if acIdx == 1 {
		restrictedTextMultiCols(page, ac, rT11fRealEstate)
	} else {
		restrictedTextMultiCols(page, ac, rT11fInfrastruc)
	}

	return nil
}
