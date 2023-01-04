package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage4AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage4(q, page, 0)
}
func pdsPage4AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage4(q, page, 1)
}
func pdsPage4AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage4(q, page, 2)
}

func pdsPage4(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]

	// Tranche type selection has no effect on this page.
	// Its for the asset class as a whole
	// ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: &nbsp;&nbsp; Qualitative questions", ac.Lbl["en"]),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp; Qualitative questions", ac.Lbl["de"]),
	}.Outline(fmt.Sprintf("%c4.", rn))
	page.Short = trl.S{
		"en": "Quality",
		"de": "Quality",
	}
	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>quality", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>quality", ac.Short["de"]),
	}
	page.CounterProgress = "4"
	page.CounterProgress = fmt.Sprintf("%c4", rn)

	page.WidthMax("42rem")

	// dynamically recreate the groups
	page.Groups = nil

	// matrix1
	{

		lblMain := trl.S{
			"en": `What are the main risks for your investment strategy over the next 3&nbsp;months? <br>Please choose three.`,
			"de": `What are the main risks for your investment strategy over the next 3&nbsp;months? <br>Please choose three.`,
		}
		lblMainDesc := trl.S{
			"en": `Choose from list - For example: Business cycle, Interest rates, ... , Other (free text)`,
			"de": `Choose from list - For example: Business cycle, Interest rates, ... , Other (free text)`,
		}
		lblMain.Append90(lblMainDesc)
		lblMain.Outline("4.1")

		inps := []string{
			"1_business_cycle",
			"2_interest_rates",
			"3_inflation_deflation",
			"4_demographics",
			"5_supply_chains",
			"6_health_issues",
			"7_regulatory_environment",
			"8_other",
		}

		lbls := map[string]string{
			"1_business_cycle":         "Business cycle",
			"2_interest_rates":         "Interest rates",
			"3_inflation_deflation":    "Inflation/deflation",
			"4_demographics":           "Demographics",
			"5_supply_chains":          "Supply chains",
			"6_health_issues":          "Health issues",
			"7_regulatory_environment": "Regulatory environment",
			"8_other":                  "Other",
		}
		prio3Matrix(page, ac, "q41_risks", lblMain, inps, lbls, true)
	}

	//
	// matrix2
	{

		if acIdx == 0 {

			lblMain := trl.S{
				"en": `What GICS sectors do you expect to be most challenging in the next three months? <br>Please rank the top three.`,
				"de": `What GICS sectors do you expect to be most challenging in the next three months? <br>Please rank the top three.`,
			}
			lblMainDesc := trl.S{
				"en": `Choose from sectors based on GICS: Energy, Materials, ... , Real estate`,
				"de": `Choose from sectors based on GICS: Energy, Materials, ... , Real estate`,
			}
			lblMain.Append90(lblMainDesc)
			lblMain.Outline("4.2")

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
			}

			lbls := map[string]string{
				"01_energy":                 "Energy",
				"02_materials":              "Materials",
				"03_industrials":            "Industrials",
				"04_consumer_discretionary": "Consumer discretionary",
				"05_consumer_staples":       "Consumer staples",
				"06_health_care":            "Health care",
				"07_financials":             "Financials",
				"08_information_technology": "Information technology",
				"09_communication_services": "Communication services",
				"10_utilities":              "Utilities",
				"11_real_estate":            "Real estate",
			}

			prio3Matrix(page, ac, "q42_gicsprio", lblMain, inps, lbls, false)

		}

		if acIdx == 1 || acIdx == 2 {

			lblMain := trl.S{
				"en": `What sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
				"de": `What sectors do you expect to be most challenging in the next three months? Please rank the top three.`,
			}
			lblMainDesc := trl.S{
				"en": `Choose from sectors: Office, Retail, Hospitality, Residential, Logistics, Other.`,
				"de": `Choose from sectors: Office, Retail, Hospitality, Residential, Logistics, Other.`,
			}

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

			if acIdx == 2 {

				lblMainDesc = trl.S{
					"en": `Choose from sectors: Transportation, Power, Renewables, Utilities, Telecoms, Social, Other`,
					"de": `Choose from sectors: Transportation, Power, Renewables, Utilities, Telecoms, Social, Other`,
				}

				inps = []string{
					"01_transportation",
					"02_power",
					"03_renewables",
					"04_utilities",
					"05_telecoms",
					"06_social",
					"07_other",
				}

				lbls = map[string]string{
					"01_transportation": "Transportation",
					"02_power":          "Power",
					"03_renewables":     "Renewables",
					"04_utilities":      "Utilities",
					"05_telecoms":       "Telecoms",
					"06_social":         "Social",
					"07_other":          "Other",
				}

			}

			lblMain.Append90(lblMainDesc)
			lblMain.Outline("4.2")

			prio3Matrix(page, ac, "q42_challenge_sectors", lblMain, inps, lbls, false)

		}

	}

	{

		esgImportance1 := trl.S{
			"en": `How important are ESG considerations in your investment process?`,
			"de": `How important are ESG considerations in your investment process?`,
		}
		desc := trl.S{
			"en": `Not relevant (1) to core principal in the investment process and potential dealbreaker (5)`,
			"de": `Not relevant (1) to core principal in the investment process and potential dealbreaker (5)`,
		}
		esgImportance1.Append90(desc)
		esgImportance1.Outline("4.3")

		// rejected at meeting 2022-11-14
		// rangePercentage(qst.WrapPageT(page), "esg", esgImportance1, "importance2")

		radiosSingleRow(
			page,
			ac,
			"q43_esg_importance",
			esgImportance1,
			mCh3,
		)
	}

	/*
		//
		// matrix3
		{
			inps := []string{
				"1_availability",
				"2_quality",
				"3_performance",
				"4_greenwashing",
				"5_regulation",
				"6_opportunities",
				"7_other",
			}

			lbls := map[string]string{
				"1_availability":  "ESG data availability",
				"2_quality":       "ESG data quality",
				"3_performance":   "Concerns about performance/sacrificing returns",
				"4_greenwashing":  "Concerns about greenwashing",
				"5_regulation":    "Complex regulatory landscape",
				"6_opportunities": "Lack of suitable investments",
				"7_other":         "Other",
			}

			lblMain := trl.S{
				"en": `What is the biggest challenge related to the implementation of ESG into your investment strategy?`,
				"de": `What is the biggest challenge related to the implementation of ESG into your investment strategy?`,
			}.Outline("4.4")
			prio3Matrix(page, ac, "q44_esg_challenge", lblMain, inps, lbls, true)
		}

		{

			var inps = []string{
				"q4501_poverty",
				"q4502_hunger",
				"q4503_health",
				"q4504_education",
				"q4505_gender_eq",
				"q4506_water",
				"q4507_energy",
				"q4508_work",
				"q4509_industry",
				"q4510_inequality",
				"q4511_communities",
				"q4512_responsible",
				"q4513_climate",
				"q4514_life_water",
				"q4515_life_land",
				"q4516_peace",
				"q4517_partnership",
			}
			var lbls = []trl.S{
				{
					"de": "No poverty",
					"en": "No poverty",
				},
				{
					"de": "Zero hunger",
					"en": "Zero hunger",
				},
				{
					"de": "Good health and well-being",
					"en": "Good health and well-being",
				},
				{
					"de": "Quality education",
					"en": "Quality education",
				},
				{
					"de": "Gender equality",
					"en": "Gender equality",
				},
				{
					"de": "Clean water and sanitation",
					"en": "Clean water and sanitation",
				},
				{
					"de": "Affordable and clean energy",
					"en": "Affordable and clean energy",
				},
				{
					"de": "Decent work and economic growth",
					"en": "Decent work and economic growth",
				},
				{
					"de": "Industry innovation and infrastructure",
					"en": "Industry innovation and infrastructure",
				},
				{
					"de": "Reduce inequality",
					"en": "Reduce inequality",
				},
				{
					"de": "Sustainable cities and communities",
					"en": "Sustainable cities and communities",
				},
				{
					"de": "Responsible consumption and production",
					"en": "Responsible consumption and production",
				},
				{
					"de": "Climate action",
					"en": "Climate action",
				},
				{
					"de": "Life below water",
					"en": "Life below water",
				},
				{
					"de": "Life on land",
					"en": "Life on land",
				},
				{
					"de": "Peace, justice and strong institutions",
					"en": "Peace, justice and strong institutions",
				},
				{
					"de": "Partnership for the goals",
					"en": "Partnership for the goals",
				},
			}

			unSDG := trl.S{
				"en": `What UN SDGs are supported by your investment strategy?`,
				"de": `What UN SDGs are supported by your investment strategy?`,
			}.Outline("4.5")
			checkBoxColumn(page, ac, unSDG, 2, inps, lbls)
		}
	*/
	return nil
}
