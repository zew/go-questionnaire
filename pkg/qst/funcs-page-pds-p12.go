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

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": "<span style='font-size:85%; font-weight: normal'>Portfolio changes continued: <br></span> New transactions",
		"de": "<span style='font-size:85%; font-weight: normal'>Portfolio changes continued: <br></span> New transactions",
	}
	page.Short = trl.S{
		"en": "Portfolio changes - 2",
		"de": "Portfolio changes - 2",
	}
	page.CounterProgress = "1b"
	page.SuppressInProgressbar = true

	page.WidthMax("58rem")

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
				"en": "Return (unlevered)",
				"de": "Return (unlevered)",
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
		"range-pct",
	}
	page12Inputs := []string{
		"q12a_cash_margin",
		"q12b_interest_floor",
		"q12c_upfront_fee",
		"q12d_fixed_rate_coupon",
		"q12e_irr_expected",
		"q12f_share_floating_rate",
	}
	page12Lbls := []trl.S{
		{
			"en": `
					Margin (over 3m Euribor) <br>
					<span class=font-size-90 >Average cash margin (only relevant for floating rate loans)</span>`,
			"de": `
					Margin (over 3m Euribor) <br>
					<span class=font-size-90 >Average cash margin (only relevant for floating rate loans)</span>`,
		},
		{
			"en": `
					Interest floor <br>
					<span class=font-size-90 >Average interest floor (only relevant for floating rate loans)</span>`,
			"de": `
					Interest floor <br>
					<span class=font-size-90 >Average interest floor (only relevant for floating rate loans)</span>`,
		},
		{
			"en": `
					Upfront fee <br>
					<span class=font-size-90 >Average upfront fee (percent of loan value)</span>`,
			"de": `
					Upfront fee <br>
					<span class=font-size-90 >Average upfront fee (percent of loan value)</span>`,
		},

		{
			"en": `
					Fixed rate coupon <br>
					<span class=font-size-90 > Average fixed rate coupon (only relevant for fixed rate loans) </span>`,
			"de": `
					Fixed rate coupon <br>
					<span class=font-size-90 > Average fixed rate coupon (only relevant for fixed rate loans) </span>`,
		},
		{
			"en": `
					 Expected IRR <br>
					<span class=font-size-90 > Average expected IRR  </span>`,
			"de": `
					 Expected IRR <br>
					<span class=font-size-90 > Average expected IRR  </span>`,
		},
		{
			"en": `
					Share of floating rate debt <br>
					<span class=font-size-90 > Share of floating rate debt </span>`,
			"de": `
					Share of floating rate debt <br>
					<span class=font-size-90 > Share of floating rate debt </span>`,
		},
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
			&sliderPctZeroFour,
			&sliderPctThreeTwenty,
			&sliderPctThreeTwentyfive,
			&sliderPctZeroHundredWide,
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
		"range-pct",
	}
	page13Inputs := []string{
		"q13a_number_covenants",
		"q13b_contracted_maturity",
		"q13c_opening_leverage",
		"q13d_ebitda_avg",
		"q13e_ev_avg",
		"q13f_share_sponsored_or_not",
	}
	page13Lbls := []trl.S{
		{
			"en": `
					Average # of covenants <br>
					<span class=font-size-90 > Average number of financial covenants per credit </span>`,
			"de": `
					Average # of covenants <br>
					<span class=font-size-90 > Average number of financial covenants per credit </span>`,
		},
		{
			"en": `
					Contracted maturity <br>
					<span class=font-size-90 > Average contracted maturity </span>`,
			"de": `
					Contracted maturity <br>
					<span class=font-size-90 > Average contracted maturity </span>`,
		},
		{
			"en": `
					 Opening leverage <br>
					<span class=font-size-90 > Measured as a multiple of EBITDA  </span>`,
			"de": `
					 Opening leverage <br>
					<span class=font-size-90 > Measured as a multiple of EBITDA  </span>`,
		},
		{
			"en": `
					 Average EBITDA <br>
					<span class=font-size-90 > Average EBITDA of companies ; todo: Einrasten auf 10 über 50  </span>`,
			"de": `
					 Average EBITDA <br>
					<span class=font-size-90 > Average EBITDA of companies ; todo: Einrasten auf 10 über 50  </span>`,
		},
		{
			"en": `
					 Average EV <br>
					<span class=font-size-90 > Average EV of companies </span>`,
			"de": `
					 Average EV <br>
					<span class=font-size-90 > Average EV of companies </span>`,
		},
		{
			"en": `
					 Share of sponsored vs. sponsor-less <br>
					<span class=font-size-90 > Percentage of deals with private equity sponsor </span>`,
			"de": `
					 Share of sponsored vs. sponsor-less <br>
					<span class=font-size-90 > Percentage of deals with private equity sponsor </span>`,
		},
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
			&sliderEBITDAZeroHundred,
			&sliderEVZeroFiveHundred,
			&sliderPctZeroHundredWide,
		},
	)

	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"en": "Realizations",
				"de": "Realizations",
			}.Outline("1.4")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	page14Types := []string{
		"restricted-text-million",
		"range-pct",
		"range-pct",
		"range-pct",
	}
	page14Inputs := []string{
		"q14a_vol_realized_loans",
		"q14b_time_to_maturity",
		"q14c_gross_irr",
		"q14d_gross_moic",
	}
	page14Lbls := []trl.S{
		{
			"en": `
					Realisations <br>
					<span class=font-size-90 > Volume of realized loans in € </span>`,
			"de": `
					Realisations <br>
					<span class=font-size-90 > Volume of realized loans in € </span>`,
		},
		{
			"en": `
					Time to maturity <br>
					<span class=font-size-90 > Average time to maturity of realized deals </span>`,
			"de": `
					Time to maturity <br>
					<span class=font-size-90 > Average time to maturity of realized deals </span>`,
		},
		{
			"en": `
					Realized gross IRR <br>
					<span class=font-size-90 > Average realized gross IRR </span>`,
			"de": `
					Realized gross IRR <br>
					<span class=font-size-90 > Average realized gross IRR </span>`,
		},
		{
			"en": `
					Realized gross MOIC <br>
					<span class=font-size-90 > Average realized gross MOIC </span>`,
			"de": `
					Realized gross MOIC <br>
					<span class=font-size-90 > Average realized gross MOIC </span>`,
		},
	}

	for i := 0; i < len(page14Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page14Lbls[i] = page14Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	createRows(
		page,
		ac,
		page14Inputs,
		page14Types,
		page14Lbls,
		[]*rangeConf{
			nil,
			&sliderYearsZeroTen,
			&sliderPctThreeTwentyfive,
			&sliderOneOnePointFive,
		},
	)

	return nil
}
