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
		"en": fmt.Sprintf("%v: &nbsp;&nbsp; Realizations in %v", ac.Lbl["en"], q.Survey.Quarter(-1)),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp; Realizations in %v", ac.Lbl["de"], q.Survey.Quarter(-1)),
	}.Outline(fmt.Sprintf("%c2.", rn))

	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>realizations", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>realizations", ac.Short["de"]),
	}
	page.Short = trl.S{
		"en": fmt.Sprintf("<b> &nbsp;<br> &nbsp; </b> <br>Repaid Loans"),
		"de": fmt.Sprintf("<b> &nbsp;<br> &nbsp; </b> <br>Repaid Loans"),
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
			"en": `Volume of realized loans`,
			"de": `Volume of realized loans`,
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
			"en": `Please state the volume (in mn EUR) of realized loans closed in [quarter-1].`,
			"de": `Please state the volume (in mn EUR) of realized loans closed in [quarter-1].`,
		},
		{
			"en": `Please state the average time to repayment of repayed loans.`,
			"de": `Please state the average time to repayment of repayed loans.`,
		},
		{
			"en": `Please state the average realized Gross IRR.`,
			"de": `Please state the average realized Gross IRR.`,
		},
		{
			"en": `Please state the average realized Gross MOIC.`,
			"de": `Please state the average realized Gross MOIC.`,
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
				nil,
				&range0To10,
				&range0To25,
				&range0To2StepDot1,
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
				nil,
				&range0To10,
				&range0To25,
				// &sliderOneOnePointFive,
			},
		)
	}

	// here not possible, because necessary values in other pages are not yet fully populated at this point
	// pdsSpecialDisableColumns(...)

	return nil
}
