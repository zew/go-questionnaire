package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage4X(q *QuestionnaireT, page *pageT) error {

	// ac := PDSAssetClasses[0]
	ac := PDSAssetClassGlob

	page.Label = trl.S{
		"en": fmt.Sprintf("European Private Debt Markets in %v (continued)", q.Survey.Quarter(0)),
		"de": fmt.Sprintf("European Private Debt Markets in %v (continued)", q.Survey.Quarter(0)),
	}

	page.Label.Outline("2.")

	page.Short = trl.S{
		"en": (" Private <br>   Debt Markets 2"),
		"de": (" Private <br>   Debt Markets 2"),
	}

	page.CounterProgress = "2"

	page.WidthMax("42rem")

	// dynamically recreate the groups
	page.Groups = nil

	//
	// matrix1
	{
		lblMain := trl.S{
			"en": `What do you think are the main risks for your investment strategy over the next 3&nbsp;months? Please rank the top three.`,
			"de": `What do you think are the main risks for your investment strategy over the next 3&nbsp;months? Please rank the top three.`,
		}
		lblMain.Outline("2.1")

		inps := []string{
			"1_business_cycle",
			"2_interest_rates",
			"3_inflation_deflation",
			"4_regulatory_environment",
			"5_supply_chains",
			"6_health_issues",
			"7_demographics",
			"8_other",
		}

		lbls := map[string]string{
			"1_business_cycle":         "Business Cycle",
			"2_interest_rates":         "Interest Rates",
			"3_inflation_deflation":    "Inflation/Deflation",
			"4_regulatory_environment": "Regulatory Environment",
			"5_supply_chains":          "Supply Chain Disruptions",
			"6_health_issues":          "Health Issues (e.g. Covid)",
			"7_demographics":           "Demographic Change",
			"8_other":                  "Other",
		}
		prio3Matrix(page, ac, "q35_risks", lblMain, inps, lbls, true)
	}

	//
	//
	{
		esgImportance1 := trl.S{
			"en": `How important are ESG considerations in your in­vest­ment process?`,
			"de": `How important are ESG considerations in your in­vest­ment process?`,
		}
		desc := trl.S{
			"en": `Please choose the statement that describes most closely the importance of ESG considerations in your in­vest­ment process. `,
			"de": `Please choose the statement that describes most closely the importance of ESG considerations in your in­vest­ment process. `,
		}
		esgImportance1.Append90(desc)
		esgImportance1.Outline("2.2")

		radiosSingleRow(
			page,
			ac,
			"q37_esg_importance",
			esgImportance1,
			mCh3,
		)
	}

	return nil
}
