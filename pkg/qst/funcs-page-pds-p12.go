package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage12AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage12(q, page, 0)
}
func pdsPage12AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage12(q, page, 1)
}
func pdsPage12AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage12(q, page, 2)
}

func pdsPage12(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]
	ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": fmt.Sprintf(`
					New transactions
				<span style='font-size:85%%; font-weight: normal'> &nbsp;&nbsp;&nbsp; (portfolio changes continued: %v)</span>
				`, ac.Lbl["en"]),
		"de": fmt.Sprintf(`
					New transactions
				<span style='font-size:85%%; font-weight: normal'> &nbsp;&nbsp;&nbsp; (portfolio changes continued: %v)</span>
				`, ac.Lbl["de"]),
	}.Outline(fmt.Sprintf("%c1.", rn))
	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>Changes 2", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>Changes 2", ac.Short["de"]),
	}
	page.CounterProgress = fmt.Sprintf("%c1b", rn)
	page.SuppressInProgressbar = true

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
				"en": "Unlevered returns",
				"de": "Unlevered returns",
			}.Outline("1.2")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	page12Types := []string{
		"range-pct",
		"range-pct",
		"range-pct",
		"range-pct",
		"range-pct",
	}
	page12Inputs := []string{
		"q12a_cash_margin",
		"q12b_interest_floor",
		"q12c_upfront_fee",
		"q12d_fixed_rate_coupon",
		"q12e_irr_expected",
		// "q12f_share_floating_rate",
	}
	page12Lbls := []trl.S{
		{
			"en": `Average cash margin over the relevant base rate`,
			"de": `Average cash margin over the relevant base rate`,
		},
		{
			"en": `Average interest rate floor`,
			"de": `Average interest rate floor`,
		},
		{
			"en": `Average fixed rate coupon`,
			"de": `Average fixed rate coupon`,
		},
		{
			"en": `Average upfront fee`,
			"de": `Average upfront fee`,
		},
		{
			"en": `Average expected Gross IRR`,
			"de": `Average expected Gross IRR`,
		},
	}

	page12LblsDescr := []trl.S{
		{
			"en": `Please state the average cash margin over the relevant base rate. Only relevant for floating rate loans.`,
			"de": `Please state the average cash margin over the relevant base rate. Only relevant for floating rate loans.`,
		},
		{
			"en": `Please state the average interest floor. Only relevant for floating rate loans.`,
			"de": `Please state the average interest floor. Only relevant for floating rate loans.`,
		},
		{
			"en": `Please state the average fixed rate copuon. Only relevant for fixed rate loans.`,
			"de": `Please state the average fixed rate copuon. Only relevant for fixed rate loans.`,
		},
		{
			"en": `Please state the average upfront fees charged to the borrower.`,
			"de": `Please state the average upfront fees charged to the borrower.`,
		},
		{
			"en": `Please state the average expected Gross IRR.`,
			"de": `Please state the average expected Gross IRR.`,
		},
	}

	for i := 0; i < len(page12Lbls); i++ {
		page12Lbls[i].Append90(page12LblsDescr[i])
	}

	for i := 0; i < len(page12Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page12Lbls[i] = page12Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	createRows(
		page,
		ac,
		page12Inputs,
		page12Types,
		page12Lbls,
		[]*rangeConf{
			&sliderPctThreeTen,
			&sliderPctZeroTwo,
			&sliderPctThreeTwenty,
			&sliderPctZeroFour,
			&sliderPctThreeTwentyfive,
			// &sliderPctZeroHundredWide,
		},
	)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"en": "Risk",
				"de": "Risk",
			}.Outline("1.3")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	page13Types := []string{
		"radios1-4",
		"range-pct",
		"range-pct",
		"range-pct",
		"range-pct",
		"restricted-text-int",
		"restricted-text-int",
	}
	page13Inputs := []string{
		"q13a_number_covenants",
		"q13b_contracted_maturity",
		"q13c_opening_leverage",
		"q13d_ebitda_avg",
		"q13e_ev_avg",
		"q13f_share_sponsored_or_not",
		"q13g_share_stepdown",
	}
	page13Lbls := []trl.S{
		{
			"en": `Average number of covenants`,
			"de": `Average number of covenants`,
		},
		{
			"en": `Contracted maturity`,
			"de": `Contracted maturity`,
		},
		{
			"en": `Opening Leverage`,
			"de": `Opening Leverage`,
		},
		{
			"en": `Average EBITDA`,
			"de": `Average EBITDA`,
		},
		{
			"en": `Average EV`,
			"de": `Average EV`,
		},
		{
			"en": `Number of loans with PE sponsor`,
			"de": `Number of loans with PE sponsor`,
		},
		{
			"en": `Number of loans with margin step down`,
			"de": `Number of loans with margin step down`,
		},
	}

	page13LblsDescr := []trl.S{
		{
			"en": `What is the average number of financial covenants per loan?`,
			"de": `What is the average number of financial covenants per loan?`,
		},
		{
			"en": `What is the average contracted maturity?`,
			"de": `What is the average contracted maturity?`,
		},
		{
			"en": `What is the average opening leverage, measured as a multiple of EBITDA?`,
			"de": `What is the average opening leverage, measured as a multiple of EBITDA?`,
		},
		{
			"en": `What is the average EBITDA of borrower companies?`,
			"de": `What is the average EBITDA of borrower companies?`,
		},
		{
			"en": `What is the average enterprise value of borrower companies?`,
			"de": `What is the average enterprise value of borrower companies?`,
		},
		{
			"en": `Please state the number of transactions with a private equity sponsor.`,
			"de": `Please state the number of transactions with a private equity sponsor.`,
		},
		{
			"en": `Please state the number of transactions with a margin step down.`,
			"de": `Please state the number of transactions with a margin step down.`,
		},
	}

	for i := 0; i < len(page13Lbls); i++ {
		page13Lbls[i].Append90(page13LblsDescr[i])
	}

	for i := 0; i < len(page13Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page13Lbls[i] = page13Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	createRows(
		page,
		ac,
		page13Inputs,
		page13Types,
		page13Lbls,
		[]*rangeConf{
			nil, // unused
			&sliderYearsZeroTen,
			&sliderEBITDA2x10x,
			&sliderEBITDAZero150,
			&sliderEVZeroFiveHundred,
			nil,
			nil,
		},
	)

	return nil
}
