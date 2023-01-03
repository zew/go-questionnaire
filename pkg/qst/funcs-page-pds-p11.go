package qst

import (
	"fmt"
	"strings"

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

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: &nbsp;&nbsp;  Portfolio changes  (past 3 months)", ac.Lbl["en"]),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp;  Portfolio changes  (past 3 months)", ac.Lbl["de"]),
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
		page.WidthMax("34rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	//
	// todo: current or previous quarter
	if strings.Contains(rT1.LblRow2["en"], "[quarter]") {
		rT1.LblRow2["en"] = strings.ReplaceAll(rT1.LblRow2["en"], "[quarter]", q.Survey.Quarter())
		rT1.LblRow2["en"] = strings.ReplaceAll(rT1.LblRow2["en"], "[year]", "")
	}

	restrictedTextMultiCols(page, ac, rT1)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = trl.S{
			"en": "Total transaction volume (in mn €)",
			"de": "Total transaction volume (in mn €)",
		}.Outline("b.)")
	}
	restrictedTextMultiCols(page, ac, rT1b)

	lblDuration := trl.S{
		"en": "Average time to close a deal in weeks (across all tranche types)",
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
		restrictedTextMultiCols(page, ac, rT2)
	} else {
		restrictedTextMultiCols(page, ac, rT2RealEstate)
	}

	restrictedTextMultiCols(page, ac, rT3)

	if acIdx == 0 {
		restrictedTextMultiCols(page, ac, rT4)
	} else if acIdx == 1 {
		restrictedTextMultiCols(page, ac, rT4RealEstate)
	} else {
		restrictedTextMultiCols(page, ac, rT4Infrastruc)
	}

	/* 	page11fghInputs := []string{
	   		"q11f_esg",
	   		"q11g_ratch",
	   		"q11h_degrees",
	   	}

	   	page11fghTypes := []string{
	   		"range-pct",
	   		"range-pct",
	   		"range-pct",
	   	}

	   	page11fghLbls := []trl.S{
	   		{
	   			"en": `<bb>Share ESG KPIs</bb> <br>
	   					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
	   			"de": `<bb>Share ESG KPIs</bb> <br>
	   					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
	   		},
	   		{
	   			"en": `<bb>Share ESG ratchets</bb> <br>
	   					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
	   			"de": `<bb>Share ESG ratchets</bb> <br>
	   					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
	   		},
	   		{
	   			"en": `<bb>Share 1.5°C target</bb> <br>
	   					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?</span>`,
	   			"de": `<bb>Share 1.5°C target</bb> <br>
	   					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?</span>`,
	   		},
	   	}

	   	{

	   		// 4cols layout

	   		for i := 0; i < len(page11fghLbls); i++ {
	   			rn := rune(102 + i) // 102 is f
	   			page11fghLbls[i] = page11fghLbls[i].Outline(fmt.Sprintf("%c.)", rn))
	   		}

	   		createRows(
	   			page,
	   			ac,
	   			page11fghInputs,
	   			page11fghTypes,
	   			page11fghLbls,
	   			[]*rangeConf{
	   				&sliderPctZeroHundredWide, // &sliderPctZeroHundredMiddle,
	   				&sliderPctZeroHundredWide,
	   				&sliderPctZeroHundredWide,
	   			},
	   		)

	   	}
	*/
	return nil
}
