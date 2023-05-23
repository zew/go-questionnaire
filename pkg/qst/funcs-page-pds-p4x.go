package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage4X(q *QuestionnaireT, page *pageT) error {

	// ac := PDSAssetClasses[0]
	ac := PDSAssetClassGlob
	ac = onlySelectedTranchTypes2(q, ac)

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
			"en": `<i>All asset classes</i>: What do you think are the main risks for your investment strategy over the next 3&nbsp;months? Please rank the top three.`,
			"de": `<i>All asset classes</i>: What do you think are the main risks for your investment strategy over the next 3&nbsp;months? Please rank the top three.`,
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

	counter := 2

	//
	// matrix2
	{

		if ac.Has("ac1") {

			lblMain := trl.S{
				"en": fmt.Sprintf(`
					<i>%v</i>: 
					What GICS sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
					ac.Get("ac1").Lbl,
				),
				"de": fmt.Sprintf(`
					<i>%v</i>: 
					What GICS sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
					ac.Get("ac1").Lbl,
				),
			}
			lblMain.Outline(fmt.Sprintf("2.%v", counter))
			counter++

			inps := []string{
				"01_energy",
				"02_materials",
				"03_industrials",
				"04_consumer_discretionary",
				"05_consumer_staples",
				"06_health_care",
				"07_financials",
				"08_information_technology",
				"09_communication_services",
				"10_utilities",
				"11_real_estate",
				"12_other",
			}

			lbls := map[string]string{
				"01_energy":                 "Energy",
				"02_materials":              "Materials",
				"03_industrials":            "Industrials",
				"04_consumer_discretionary": "Consumer Discretionary",
				"05_consumer_staples":       "Consumer Staples",
				"06_health_care":            "Health Care",
				"07_financials":             "Financials",
				"08_information_technology": "Information Technology",
				"09_communication_services": "Communication Services",
				"10_utilities":              "Utilities",
				"11_real_estate":            "Real Estate",
				"12_other":                  "Other",
			}

			prio3Matrix(page, ac, "ac1_q36_gicsprio", lblMain, inps, lbls, true)

		}

		if ac.Has("ac2") {

			lblMain := trl.S{
				"en": fmt.Sprintf(`
					<i>%v</i>: 
					What sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
					ac.Get("ac2").Lbl,
				),
				"de": fmt.Sprintf(`
					<i>%v</i>: 
					What sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
					ac.Get("ac2").Lbl,
				),
			}

			lblMain.Outline(fmt.Sprintf("2.%v", counter))
			counter++

			inps := []string{
				"01_office",
				"02_retail",
				"03_hospitality",
				"04_residential",
				"05_logistics",
				"06_other",
			}

			lbls := map[string]string{
				"01_office":      "Office",
				"02_retail":      "Retail",
				"03_hospitality": "Hospitality",
				"04_residential": "Residential",
				"05_logistics":   "Logistics",
				"06_other":       "Other",
			}

			prio3Matrix(page, ac, "ac2_q36_challenge_sectors", lblMain, inps, lbls, true)

		}
		if ac.Has("ac3") {

			lblMain := trl.S{
				"en": fmt.Sprintf(`
					<i>%v</i>: 
					What sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
					ac.Get("ac3").Lbl,
				),
				"de": fmt.Sprintf(`
					<i>%v</i>: 
					What sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
					ac.Get("ac3").Lbl,
				),
			}

			lblMain.Outline(fmt.Sprintf("2.%v", counter))
			counter++

			inps := []string{
				"01_transportation",
				"02_power",
				"03_renewables",
				"04_utilities",
				"05_telecoms",
				"06_social",
				"07_other",
			}

			lbls := map[string]string{
				"01_transportation": "Transportation",
				"02_power":          "Power",
				"03_renewables":     "Renewables",
				"04_utilities":      "Utilities",
				"05_telecoms":       "Telecoms",
				"06_social":         "Social",
				"07_other":          "Other",
			}

			prio3Matrix(page, ac, "ac3_q36_challenge_sectors", lblMain, inps, lbls, true)

		}

	}

	//
	//
	{
		esgImportance1 := trl.S{
			"en": `<i>All asset classes</i>: How important are ESG considerations in your investment process?`,
			"de": `<i>All asset classes</i>: How important are ESG considerations in your investment process?`,
		}
		desc := trl.S{
			"en": `Please choose the statement that describes most closely the importance of ESG considerations in your investment process. `,
			"de": `Please choose the statement that describes most closely the importance of ESG considerations in your investment process. `,
		}
		esgImportance1.Append90(desc)
		esgImportance1.Outline(fmt.Sprintf("2.%v", counter))
		counter++

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
