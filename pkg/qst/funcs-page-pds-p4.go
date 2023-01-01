package qst

import "github.com/zew/go-questionnaire/pkg/trl"

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
	ac = onlySelectedTranchTypes(q, ac)

	page.Label = trl.S{
		"en": "4. Qualitative questions",
		"de": "4. Qualitative questions",
	}
	page.Short = trl.S{
		"en": "Quality",
		"de": "Quality",
	}
	page.CounterProgress = "4"

	page.WidthMax("42rem")

	// dynamically recreate the groups
	page.Groups = nil

	// matrix1
	{
		inps := []string{
			"411_business_cycle",
			"412_interest_rates",
			"413_inflation_deflation",
			"414_demographics",
			"415_supply_chains",
			"416_health_issues",
			"417_regulatory_environment",
			"418_other",
		}

		lbls := map[string]string{
			"411_business_cycle":         "Business cycle",
			"412_interest_rates":         "Interest rates",
			"413_inflation_deflation":    "Inflation/deflation",
			"414_demographics":           "Demographics",
			"415_supply_chains":          "Supply chains",
			"416_health_issues":          "Health issues",
			"417_regulatory_environment": "Regulatory environment",
			"418_other":                  "Other",
		}

		lblMain := trl.S{
			"en": `What do you think are the main risks for your investment strategy over the next 3 months? Please choose three.`,
			"de": `What do you think are the main risks for your investment strategy over the next 3 months? Please choose three.`,
		}.Outline("4.1")
		prio3Matrix(page, ac, "risks", lblMain, inps, lbls, true)
	}

	//
	// matrix2
	{
		inps := []string{
			"q4201_energy",
			"q4202_materials",
			"q4203_industrials",
			"q4204_consumer_discretionary",
			"q4205_consumer_staples",
			"q4206_health_care",
			"q4207_financials",
			"q4208_information_technology",
			"q4209_communication_services",
			"q4210_utilities",
			"q4211_real_estate",
		}

		lbls := map[string]string{
			"q4201_energy":                 "Energy",
			"q4202_materials":              "Materials",
			"q4203_industrials":            "Industrials",
			"q4204_consumer_discretionary": "Consumer discretionary",
			"q4205_consumer_staples":       "Consumer staples",
			"q4206_health_care":            "Health care",
			"q4207_financials":             "Financials",
			"q4208_information_technology": "Information technology",
			"q4209_communication_services": "Communication services",
			"q4210_utilities":              "Utilities",
			"q4211_real_estate":            "Real estate",
		}

		lblMain := trl.S{
			"en": `What GICS sectors provides the most attractive 
					investment opportunities in the next three months? 
					Please rank the top three.`,
			"de": `What GICS sectors provides the most attractive 
					investment opportunities in the next three months? 
					Please rank the top three.`,
		}.Outline("4.2")
		prio3Matrix(page, ac, "gicsprio", lblMain, inps, lbls, false)
	}

	{
		esgImportance1 := trl.S{
			"en": `How important are ESG considerations 
					to core principal in your investment process?
			`,
			"de": `How important are ESG considerations 
					to core principal in your investment process?
			`,
		}.Outline("4.3")

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

	//
	// matrix3
	{
		inps := []string{
			"q441_availability",
			"q442_quality",
			"q443_performance",
			"q444_greenwashing",
			"q445_regulation",
			"q446_opportunities",
			"q447_other",
		}

		lbls := map[string]string{
			"q441_availability":  "ESG data availability",
			"q442_quality":       "ESG data quality",
			"q443_performance":   "Concerns about performance/sacrificing returns",
			"q444_greenwashing":  "Concerns about greenwashing",
			"q445_regulation":    "Complex regulatory landscape",
			"q446_opportunities": "Lack of suitable investments",
			"q447_other":         "Other",
		}

		lblMain := trl.S{
			"en": `What is the biggest challenge related to the implementation of ESG into your investment strategy?`,
			"de": `What is the biggest challenge related to the implementation of ESG into your investment strategy?`,
		}.Outline("4.4")
		prio3Matrix(page, ac, "esg_challenge", lblMain, inps, lbls, true)
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

	return nil
}
