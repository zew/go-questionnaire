package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage23AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage23(q, page, 0)
}
func pdsPage23AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage23(q, page, 1)
}
func pdsPage23AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage23(q, page, 2)
}

func pdsPage23(q *QuestionnaireT, page *pageT, acIdx int) error {

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": "2.3 Portfolio risk",
		"de": "2.3 Portfolio risk",
	}
	page.Short = trl.S{
		"en": "Portfolio risk",
		"de": "Portfolio risk",
	}
	page.CounterProgress = "2.3"
	page.CounterProgress = "2b"
	page.SuppressInProgressbar = true

	page.WidthMax("58rem")

	// dynamically recreate the groups
	page.Groups = nil

	page23Types := []string{
		"radios1-4",
		"range-pct",
		"range-pct",

		"range-pct",
		"restricted-text-pct",
		"range-pct",

		"range-pct",
		"range-pct",
		"range-pct",
	}
	page23Inputs := []string{
		"q23a_covenants_per_credit",
		"q23b_share_covenant_holiday",
		"q23c_share_covenant_reset",

		"q23d_share_covenant_breach",
		"q23e_share_loan_defaults",
		"q23f_share_default_recovered",

		"q23g_share_esg_kpis",
		"q23h_share_esg_ratchets",
		"q23i_share_esg_15degrees",
	}
	page23Lbls := []trl.S{
		{
			"en": "Average number of financial covenants per credit",
			"de": "Average number of financial covenants per credit",
		},
		{
			"en": "What is the share of portfolio (at fair market value) with a covenant holiday?",
			"de": "What is the share of portfolio (at fair market value) with a covenant holiday?",
		},
		{
			"en": "What is the share of portfolio (at fair market value) with a covenant reset?",
			"de": "What is the share of portfolio (at fair market value) with a covenant reset?",
		},
		{
			"en": "What is the share of portfolio (at fair market value) with a covenant breach?",
			"de": "What is the share of portfolio (at fair market value) with a covenant breach?",
		},
		{
			"en": "Share of defaulted loans (measured at cost/principal amount of loan)",
			"de": "Share of defaulted loans (measured at cost/principal amount of loan)",
		},
		{
			"en": "If you had a default in the past. What was the recovery rate (share of principal)?",
			"de": "If you had a default in the past. What was the recovery rate (share of principal)?",
		},
		// esg
		{
			"en": "What is the share of portfolio (at fair market value) with explicit ESG targets in the credit documentation?",
			"de": "What is the share of portfolio (at fair market value) with explicit ESG targets in the credit documentation?",
		},
		{
			"en": "What is the share of portfolio (at fair market value) with ESG ratchets?",
			"de": "What is the share of portfolio (at fair market value) with ESG ratchets?",
		},
		{
			"en": "What is the share of portfolio (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?",
			"de": "What is the share of portfolio (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?",
		},
	}

	for i := 0; i < len(page23Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page23Lbls[i] = page23Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	createRows(
		page,
		PDSAssetClasses[acIdx],
		page23Inputs,
		page23Types,
		page23Lbls,
		[]*rangeConf{
			nil,
			&sliderPctZeroHundredWide,
			&sliderPctZeroHundredWide,
			&sliderPctZeroHundredWide,
			nil,
			&sliderPctZeroHundredWide,
			&sliderPctZeroHundredWide,
			&sliderPctZeroHundredWide,
			&sliderPctZeroHundredWide,
		},
	)

	return nil
}
