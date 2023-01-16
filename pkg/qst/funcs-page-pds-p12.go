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

	// marker for naviFuncs pds_ac1-3
	page.CounterProgress = "page12"

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
			"en": `Please state the average cash margin over the relevant base rate. Only relevant for <i>floating rate loans</i>.`,
			"de": `Please state the average cash margin over the relevant base rate. Only relevant for <i>floating rate loans</i>.`,
		},
		{
			"en": `Please state the average interest floor. Only relevant for <i>floating rate loans</i>.`,
			"de": `Please state the average interest floor. Only relevant for <i>floating rate loans</i>.`,
		},
		{
			"en": `Please state the average fixed rate copuon. Only relevant for <i>fixed rate loans</i>.`,
			"de": `Please state the average fixed rate copuon. Only relevant for <i>fixed rate loans</i>.`,
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
			&sliderPctTwoTen,
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

	if acIdx == 0 {

		page13Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			// "restricted-text-int",
			// "restricted-text-int",
		}
		page13Inputs := []string{
			"q13a_number_covenants",
			"q13b_contracted_maturity",
			"q13c_opening_leverage",
			"q13d_ebitda_avg",
			"q13e_ev_avg",
			// "q13f_share_sponsored_or_not",
			// "q13g_share_stepdown",
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
				"en": `Opening leverage`,
				"de": `Opening leverage`,
			},
			{
				"en": `Average EBITDA`,
				"de": `Average EBITDA`,
			},
			{
				"en": `Average EV`,
				"de": `Average EV`,
			},
			// {
			// 	"en": `Number of loans with PE sponsor`,
			// 	"de": `Number of loans with PE sponsor`,
			// },
			// {
			// 	"en": `Number of loans with margin step down`,
			// 	"de": `Number of loans with margin step down`,
			// },
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
			// {
			// 	"en": `Please state the number of transactions with a private equity sponsor.`,
			// 	"de": `Please state the number of transactions with a private equity sponsor.`,
			// },
			// {
			// 	"en": `Please state the number of transactions with a margin step down.`,
			// 	"de": `Please state the number of transactions with a margin step down.`,
			// },
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
				// nil,
				// nil,
			},
		)
	}

	if acIdx == 1 {

		page13Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",

			// real estate specific
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			"range-pct",
			// "restricted-text-int",
			// "restricted-text-int",
			// "restricted-text-int",
		}
		page13Inputs := []string{
			"q13a_number_covenants",
			"q13b_contracted_maturity",
			"q13c_opening_leverage",

			// real estate specific
			"q13d_opening_dscr",
			"q13e_opening_icr",
			"q13f_opening_debt_yield",
			"q13g_exit_leverage",
			"q13h_exit_dscr",
			"q13i_exit_icr",
			"q13j_exit_yield",
			// "q13k_num_amortizing",
			// "q13k_num_developmentrisk",
			// "q13m_num_marginstepdown",
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
				"en": `Opening leverage`,
				"de": `Opening leverage`,
			},

			// real estate specific
			{
				"en": `Opening DSCR`,
				"de": `Opening DSCR`,
			},
			{
				"en": `Opening ICR`,
				"de": `Opening ICR`,
			},
			{
				"en": `Opening debt yield`,
				"de": `Opening debt yield`,
			},
			{
				"en": `Expected exit leverage`,
				"de": `Expected exit leverage`,
			},
			{
				"en": `Expected exit DSCR`,
				"de": `Expected exit DSCR`,
			},
			{
				"en": `Expected exit ICR`,
				"de": `Expected exit ICR`,
			},
			{
				"en": `Expected exit yield`,
				"de": `Expected exit yield`,
			},
			// {
			// 	"en": `Number of amortizing loans`,
			// 	"de": `Number of amortizing loans`,
			// },
			// {
			// 	"en": `Number of loans with development risk`,
			// 	"de": `Number of loans with development risk`,
			// },
			// {
			// 	"en": `Number of loans with margin step down`,
			// 	"de": `Number of loans with margin step down`,
			// },
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
				"en": `What is the average opening LTV or LTC?`,
				"de": `What is the average opening LTV or LTC?`,
			},

			// real estate specific
			{
				"en": `What is the average opening DSCR?`,
				"de": `What is the average opening DSCR?`,
			},
			{
				"en": `What is the average opening interest rate coverage ratio?`,
				"de": `What is the average opening interest rate coverage ratio?`,
			},
			{
				"en": `What is the average opening debt yield?`,
				"de": `What is the average opening debt yield?`,
			},
			{
				"en": `What is the average expected exit LTV or LTC?`,
				"de": `What is the average expected exit LTV or LTC?`,
			},
			{
				"en": `What is the average expected exit DSCR?`,
				"de": `What is the average expected exit DSCR?`,
			},
			{
				"en": `What is the average expected exit interest rate coverage ratio?`,
				"de": `What is the average expected exit interest rate coverage ratio?`,
			},
			{
				"en": `What is the average expected exit debt yield?`,
				"de": `What is the average expected exit debt yield?`,
			},
			// {
			// 	"en": `Please state the number of amortizing loans.`,
			// 	"de": `Please state the number of amortizing loans.`,
			// },
			// {
			// 	"en": `Please state the number of loans with development risk.`,
			// 	"de": `Please state the number of loans with development risk.`,
			// },
			// {
			// 	"en": `Please state the number of transactions with a margin step down.`,
			// 	"de": `Please state the number of transactions with a margin step down.`,
			// },
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
				nil,
				&sliderYearsZeroTen,

				&slider30To100,

				// real estate specific
				&slider1To5,
				&slider1To5,
				&sliderPctTwoTen,
				&slider30To100,
				&slider1To5,
				&slider1To5,
				&sliderPctTwoTen,
				// nil,
				// nil,
				// nil,
			},
		)
	}

	if acIdx == 2 {

		page13Types := []string{
			"radios1-4",
			"range-pct",
			"range-pct",

			// infrastruct specific
			"range-pct",
			"range-pct",
			"range-pct",
			// "restricted-text-int",
			// "restricted-text-int",
		}
		page13Inputs := []string{
			"q13a_number_covenants",
			"q13b_contracted_maturity",
			"q13c_opening_leverage",

			// infrastruct specific
			"q13d_maximum_leverage",
			"q13e_average_dscr",
			"q13f_minimum_dscr",
			// "q13g_num_greenfield_risk",
			// "q13h_num_margin_step_down",
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
				"en": `Opening leverage`,
				"de": `Opening leverage`,
			},

			// infrastruct specific
			{
				"en": `Expected maximum Leverage`,
				"de": `Expected maximum Leverage`,
			},
			{
				"en": `Expected average DSCR`,
				"de": `Expected average DSCR`,
			},
			{
				"en": `Expected minimum DSCR`,
				"de": `Expected minimum DSCR`,
			},
			// {
			// 	"en": `Number of loans with greenfield risk`,
			// 	"de": `Number of loans with greenfield risk`,
			// },
			// {
			// 	"en": `Number of loans with margin step down`,
			// 	"de": `Number of loans with margin step down`,
			// },
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
				"en": `What is the average opening leverage?`,
				"de": `What is the average opening leverage?`,
			},

			// infrastruct specific
			{
				"en": `What is the average expected maximum LTV?`,
				"de": `What is the average expected maximum LTV?`,
			},
			{
				"en": `What is the expected average DSCR?`,
				"de": `What is the expected average DSCR?`,
			},
			{
				"en": `What is the expected minimum DSCR?`,
				"de": `What is the expected minimum DSCR?`,
			},
			// {
			// 	"en": `Please state the number of loans with greenfield risk.`,
			// 	"de": `Please state the number of loans with greenfield risk.`,
			// },
			// {
			// 	"en": `Please state the number of transactions with a margin step down.`,
			// 	"de": `Please state the number of transactions with a margin step down.`,
			// },
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
				nil,
				&sliderYearsZeroTen,
				&slider50To100,

				// infrastruct specific
				&slider50To100,
				&slider1To175,
				&slider1To175,
				// nil,
				// nil,
			},
		)
	}

	return nil
}
