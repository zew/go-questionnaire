package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage21AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage21(q, page, 0)
}
func pdsPage21AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage21(q, page, 1)
}
func pdsPage21AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage21(q, page, 2)
}

func pdsPage21(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]
	ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.ValidationFuncName = "pdsRange,pdsPage21"

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: &nbsp;&nbsp; Repaid loans in %v", ac.Lbl["en"], q.Survey.Quarter(-1)),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp; Repaid loans in %v", ac.Lbl["de"], q.Survey.Quarter(-1)),
	}.Outline(fmt.Sprintf("%c2.", rn))

	page.Short = trl.S{
		"en": ("<b> &nbsp;<br> &nbsp; </b> Repaid<br>&nbsp;&nbsp;Loans"),
		"de": ("<b> &nbsp;<br> &nbsp; </b> Repaid<br>&nbsp;&nbsp;Loans"),
	}

	page.CounterProgress = fmt.Sprintf("%c2", rn)

	page.WidthMax("58rem")
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("42rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("38rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	page21Types := []string{
		"restricted-text-million",
		"range-pct",
		"range-pct",
		"range-pct",
	}
	page21Inputs := []string{
		"q2a_vol_realized_loans",
		"q2b_time_to_maturity",
		"q2c_gross_irr",
		"q2d_gross_moic",
	}

	page21Lbls := []trl.S{
		{
			"en": `Volume of repaid loans`,
			"de": `Volume of repaid loans`,
		},
		{
			"en": `Time to maturity`,
			"de": `Time to maturity`,
		},
		{
			"en": `Realized Gross IRR`,
			"de": `Realized Gross IRR`,
		},
		{
			"en": `Realized Gross MOIC`,
			"de": `Realized Gross MOIC`,
		},
	}

	page21LblsDescr := []trl.S{
		{
			"en": `Please state the volume (in mn EUR) of loans repaid in [quarter-1].`,
			"de": `Please state the volume (in mn EUR) of loans repaid in [quarter-1].`,
		},
		{
			"en": `Please state the average time to repayment of loans repaid in [quarter-1].`,
			"de": `Please state the average time to repayment of loans repaid in [quarter-1].`,
		},
		{
			"en": `Please state the average realized Gross Internal Rate of Return (IRR) of loans repaid in [quarter-1].`,
			"de": `Please state the average realized Gross Internal Rate of Return (IRR) of loans repaid in [quarter-1].`,
		},
		{
			"en": `Please state the average realized Gross Multiple on Invested Capital (MOIC) of loans repaid in [quarter-1].`,
			"de": `Please state the average realized Gross Multiple on Invested Capital (MOIC) of loans repaid in [quarter-1].`,
		},
	}

	for i := 0; i < len(page21Lbls); i++ {
		page21Lbls[i].Append90(page21LblsDescr[i])
	}

	for i := 0; i < len(page21Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page21Lbls[i] = page21Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	if acIdx == 0 {
		createRows(
			page,
			ac,
			page21Inputs,
			page21Types,
			page21Lbls,
			[]*rangeConf{
				nil,             // q2a
				&range0to10,     // q2b
				&rangeMin05to25, // q2c
				&range0to2,      // q2d
			},
		)
	}

	if acIdx == 1 || acIdx == 2 {

		// option four omitted
		red := len(page21Inputs) - 1
		createRows(
			page,
			ac,
			page21Inputs[:red],
			page21Types[:red],
			page21Lbls[:red],
			[]*rangeConf{
				nil,             // q2a
				&range0to10,     // q2b
				&rangeMin05to25, // q2c
				// &range0To2,      // q2d
			},
		)
	}

	// here not possible, because necessary values in other pages are not yet fully populated at this point
	// pdsSpecialDisableColumns(...)

	return nil
}
