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

	page.ValidationFuncName = "pdsRange"

	page.Label = trl.S{
		"en": "2. Overall (existing) Portfolio",
		"de": "2. Overall (existing) Portfolio",
	}
	page.Short = trl.S{
		"en": "Portfolio<br>base + risk",
		"de": "Portfolio<br>base + risk",
	}
	page.CounterProgress = "2"

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
				"en": "Assets under management",
				"de": "Assets under management",
			}.Outline("2.1")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	page21Types := []string{
		"restricted-text-million",
		"restricted-text-million",
		"restricted-text-million",
		"restricted-text-million",
	}
	page21Inputs := []string{
		"q21a_portfolio_value",
		"q21b_capital_called",
		"q21c_capital_repaid",
		"q21d_capital_reserve",
	}
	page21Lbls := []trl.S{
		{
			"en": "Fair market value of current portfolio in mn €",
			"de": "Fair market value of current portfolio in mn €",
		},
		{
			"en": "Capital called from investor in mn €",
			"de": "Capital called from investor in mn €",
		},
		{
			"en": "Repaid capital either reinvested or distributed to investor in mn €",
			"de": "Repaid capital either reinvested or distributed to investor in mn €",
		},
		{
			"en": "Dry powder in mn €",
			"de": "Dry powder in mn €",
		},
	}

	for i := 0; i < len(page21Lbls); i++ {
		rn := rune(97 + i) // 97 is a
		page21Lbls[i] = page21Lbls[i].Outline(fmt.Sprintf("%c.)", rn))
	}

	createRows(
		page,
		ac,
		page21Inputs,
		page21Types,
		page21Lbls,
		[]*rangeConf{
			nil,
			nil,
			nil,
			nil,
			nil,
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
				"en": "Portfolio composition",
				"de": "Portfolio composition",
			}.Outline("2.2")
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	restrictedTextMultiCols(page, ac, r221)
	restrictedTextMultiCols(page, ac, r222)

	return nil
}
