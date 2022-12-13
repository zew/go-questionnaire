package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create questionnaire
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"en"} // governs default language code
	// q.LangCode = "en"

	q.Survey.Org = trl.S{
		"en": "ZEW",
		"de": "ZEW",
	}
	q.Survey.Name = trl.S{
		"en": "Private Debt Survey",
		"de": "Private Debt Survey",
	}
	// q.Variations = 1

	// page0
	{
		page := q.AddPage()
		page.ValidationFuncName = ""

		page.SuppressInProgressbar = true

		page.Label = trl.S{
			"en": "Greeting",
			"de": "Begrüßung",
		}
		page.Short = trl.S{
			"en": "Greeting",
			"de": "Begrüßung",
		}

		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./welcome-page.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page1 - asset classes
	{
		page := q.AddPage()
		// page.SuppressInProgressbar = true
		page.ValidationFuncName = "pdsPage1"

		page.Label = trl.S{
			"en": "Identification and asset classes",
			"de": "Identification and asset classes",
		}
		page.Short = trl.S{
			"en": "Asset classes,<br>tranches",
			"de": "Asset classes,<br>tranches",
		}
		page.CounterProgress = "-"
		// https://www.fileformat.info/info/charset/UTF-8/list.htm?start=2048
		page.CounterProgress = "௵"
		page.CounterProgress = "᎒" // e18e92

		page.WidthMax("42rem")

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q01_identification"
				inp.MaxChars = 24
				inp.Placeholder = trl.S{
					"en": "name of manager",
					"de": "Name Manager",
				}
				inp.Label = trl.S{
					"en": "Identification",
					"de": "Identifikation",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
			}
		}

		// gr2
		if false {
			lblMain := trl.S{
				"de": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
				"en": `How big is your investment team? Please choose the team size in terms of full time equivalents.`,
			}
			radiosSingleRow(
				qst.WrapPageT(page),
				"q02_teamsize",
				lblMain,
				mCh2,
			)
		}

		// if false {
		// 	// gr3
		// 	{
		// 		lblMain := trl.S{
		// 			"en": `Which asset classes do you invest in?`,
		// 			"de": `Wählen Sie Ihre Assetklassen.`,
		// 		}
		// 		checkBoxRow(
		// 			qst.WrapPageT(page),
		// 			lblMain,
		// 			assetClassesInputs,
		// 			assetClassesLabels,
		// 		)
		// 	}

		// 	// gr4
		// 	{
		// 		lblMain := trl.S{
		// 			"en": `Your  strategies`,
		// 			"de": `Ihre Strategie`,
		// 		}
		// 		checkBoxRow(
		// 			qst.WrapPageT(page),
		// 			lblMain,
		// 			trancheTypeNamesAC1,
		// 			allLbls["ac1-tranche-types"],
		// 		)
		// 	}
		// }

		//
		// gr5
		{
			lblMain := trl.S{
				"en": `Which asset classes do you invest in?`,
				"de": `Wählen Sie Ihre Assetklassen.`,
			}
			checkBoxCascade(
				qst.WrapPageT(page),
				lblMain,
				assetClassesInputs,
				assetClassesLabels,
			)

		}

	}

	// page11
	{
		page := q.AddPage()
		page.ValidationFuncName = "pdsRange"

		// page.Section = trl.S{
		// 	"en": "Section 1",
		// 	"de": "Section 1",
		// }
		page.Label = trl.S{
			"en": "1.1 Portfolio changes (past 3 months)",
			"de": "1.1 Portfolio changes (past 3 months)",
		}
		page.Short = trl.S{
			"en": "Portfolio changes",
			"de": "Portfolio changes",
		}
		page.CounterProgress = "1a"

		page.WidthMax("42rem")
		page.WidthMax("64rem")

		restrictedTextMultiCols(qst.WrapPageT(page), rT1)

		lblDuration := trl.S{
			"en": "Average time to close a deal in weeks",
			"de": "Durchschnittl. Zeit bis Abschluss in Wochen",
		}.Outline("b.)")

		/*
			closing weeks - three display variations:
				range
				radios
				dropdown
		*/

		if false {
			UNUSEDrangesRowLabelsLeft(
				qst.WrapPageT(page),
				"closing_time",
				lblDuration,
				sliderWeeksClosing,
			)

			radiosLabelsTop(
				qst.WrapPageT(page),
				"closing_time",
				lblDuration,
				mCh5,
			)
		}

		dropdownsLabelsTop(
			qst.WrapPageT(page),
			"q11b_closing_time",
			lblDuration,
			mCh5,
		)

		restrictedTextMultiCols(qst.WrapPageT(page), rT2)

		restrictedTextMultiCols(qst.WrapPageT(page), rT3)

		restrictedTextMultiCols(qst.WrapPageT(page), rT4)

		page11fghInputs := []string{
			"q11f_esg",
			"q11g_ratch",
			"q11h_degrees",
		}

		page11fghTypes := []string{
			"range-pct",
			"range-pct",
			"range-pct",
		}

		page11fghLbls := []trl.S{
			{
				"en": `<bb>Share ESG KPIs</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
				"de": `<bb>Share ESG KPIs</bb> <br>
					<span class=font-size-90 >What is the share of new deals (at fair market value) with explicit ESG targets in the credit documentation? </span>`,
			},
			{
				"en": `<bb>Share ESG ratchets</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
				"de": `<bb>Share ESG ratchets</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) with ESG ratchets? </span>`,
			},
			{
				"en": `<bb>Share 1.5°C target</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?</span>`,
				"de": `<bb>Share 1.5°C target</bb> <br> 
					<span class=font-size-90 >What is the share of new deals (at fair market value) where the creditor explicitly states a strategy to add to the 1.5°C target?</span>`,
			},
		}

		if false {

			// 5cols layout

			for i := 0; i < len(page11fghInputs); i++ {
				outline := page11fghInputs[i][4:5]
				UNUSEDrangesRowLabelsLeft(
					qst.WrapPageT(page),
					page11fghInputs[i],
					page11fghLbls[i].Outline(outline+".)"),
					sliderPctZeroHundredMiddle,
				)
			}

		} else {

			// 4cols layout

			for i := 0; i < len(page11fghLbls); i++ {
				rn := rune(102 + i) // 102 is f
				page11fghLbls[i] = page11fghLbls[i].Outline(fmt.Sprintf("%c.)", rn))
			}

			createRows(
				qst.WrapPageT(page),
				page11fghInputs,
				page11fghTypes,
				page11fghLbls,
				[]*rangeConf{
					&sliderPctZeroHundredWide, // &sliderPctZeroHundredMiddle,
					&sliderPctZeroHundredWide,
					&sliderPctZeroHundredWide,
				},
			)

		}

	}

	// page12
	{
		page := q.AddPage()
		page.ValidationFuncName = "pdsRange"

		page.Label = trl.S{
			"en": "Portfolio changes continued: <br>New transactions",
			"de": "Portfolio changes continued: <br>New transactions",
		}
		page.Short = trl.S{
			"en": "Portfolio changes - 2",
			"de": "Portfolio changes - 2",
		}
		page.CounterProgress = "1..."
		page.CounterProgress = "-"
		page.CounterProgress = "1b"

		page.WidthMax("64rem")

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
			qst.WrapPageT(page),
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
			qst.WrapPageT(page),
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
			qst.WrapPageT(page),
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

	}

	// page20
	{
		page := q.AddPage()
		page.ValidationFuncName = "pdsRange"

		page.Label = trl.S{
			"en": "2. Overall (existing) Portfolio",
			"de": "2. Overall (existing) Portfolio",
		}
		page.Short = trl.S{
			"en": "Portfolio base",
			"de": "Portfolio base",
		}
		page.CounterProgress = "2a"

		page.WidthMax("64rem")

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
			qst.WrapPageT(page),
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

		restrictedTextMultiCols(qst.WrapPageT(page), r221)
		restrictedTextMultiCols(qst.WrapPageT(page), r222)

	}

	// page23
	{
		page := q.AddPage()
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

		page.WidthMax("64rem")

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
			qst.WrapPageT(page),
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

	}

	// page3
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "3. Index Questions",
			"de": "3. Index Questions",
		}
		page.Short = trl.S{
			"en": "Indizes",
			"de": "Indizes",
		}
		page.CounterProgress = "3"

		page.WidthMax("64rem")

		page3Inputs := []string{
			"q31_financing_situation_pricing",
			"q32_deal_quality",
			"q33_deal_documentation",
			"q34_deal_amount",
		}
		page3Lbls := []trl.S{
			{
				"en": `Financing situation/pricing <br>
						<span class=font-size-90> Assess the development of expected returns  </span>`,
				"de": `Financing situation/pricing <br>
						<span class=font-size-90> Assess the development of expected returns  </span>`,
			},
			{
				"en": `Assess the change in deal quality with respect to the risk return profile <br>
						<span class=font-size-90> Assess the change in deal quality with respect to the risk return profile </span>`,
				"de": `Assess the change in deal quality with respect to the risk return profile <br>
						<span class=font-size-90> Assess the change in deal quality with respect to the risk return profile </span>`,
			},
			{
				"en": `Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) <br>
						<span class=font-size-90>  Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) </span>`,
				"de": `Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) <br>
						<span class=font-size-90>  Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) </span>`,
			},
			{
				"en": `Do you observe more deals, same amount of deals or less deals <br>
						<span class=font-size-90> Do you observe more deals, same amount of deals or less deals  </span>`,
				"de": `Do you observe more deals, same amount of deals or less deals <br>
						<span class=font-size-90> Do you observe more deals, same amount of deals or less deals  </span>`,
			},
		}

		for i := 0; i < len(page3Lbls); i++ {
			page3Lbls[i] = page3Lbls[i].Outline(fmt.Sprintf("3.%v", i+1))
		}

		for idx1, inpName := range page3Inputs {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 0
				gr.BottomVSpacers = 1
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = page3Lbls[idx1].Bold()
					inp.Label = page3Lbls[idx1]

					inp.ColSpan = 1
					// inp.ColSpanLabel = 1
				}
			}

			radiosLabelsTop(
				qst.WrapPageT(page),
				inpName+"_past3m",
				trl.S{
					"en": "<i>Last</i> 3&nbsp;months",
					"de": "<i>Last</i> 3&nbsp;months",
				}.Outline("a.)"),
				mCh4,
			)
			radiosLabelsTop(
				qst.WrapPageT(page),
				inpName+"_next3m",
				trl.S{
					"en": "<i>Next</i> 3&nbsp;months",
					"de": "<i>Next</i> 3&nbsp;months",
				}.Outline("b.)"),
				mCh4,
			)

		}

	}

	// page4
	{
		page := q.AddPage()

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
			prio3Matrix(qst.WrapPageT(page), "risks", lblMain, inps, lbls, true)
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
			prio3Matrix(qst.WrapPageT(page), "gicsprio", lblMain, inps, lbls, false)
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
				qst.WrapPageT(page),
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
			prio3Matrix(qst.WrapPageT(page), "esg_challenge", lblMain, inps, lbls, true)
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
			checkBoxColumn(qst.WrapPageT(page), unSDG, 2, inps, lbls)
		}

	}

	//
	//
	if false {
		gn := "A"
		page := q.AddPage()
		page.GeneratorFuncName = "pds01"
		page.Label = trl.S{
			"en": fmt.Sprintf("dyn page label %v", gn),
			"de": fmt.Sprintf("dyn page label %v", gn),
		}
		page.Short = trl.S{
			"en": fmt.Sprintf("dyn %v", gn),
			"de": fmt.Sprintf("dyn %v", gn),
		}
		page.WidthMax("42rem")
		page.NoNavigation = false
	}
	if false {
		gn := "B"
		page := q.AddPage()
		page.GeneratorFuncName = "pds01"
		page.Label = trl.S{
			"en": fmt.Sprintf("dyn page label %v", gn),
			"de": fmt.Sprintf("dyn page label %v", gn),
		}
		page.Short = trl.S{
			"en": fmt.Sprintf("dyn %v", gn),
			"de": fmt.Sprintf("dyn %v", gn),
		}
		page.WidthMax("42rem")
		page.NoNavigation = false
	}

	// page6 - finish
	{
		page := q.AddPage()
		page.Label = trl.S{
			"en": "Finish",
			"de": "Abschluss<br><br>",
		}
		page.Short = trl.S{
			"en": "Finish",
			"de": "DSGVO",
		}
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "q61_dsgvo"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 6
				inp.Validator = "must"
				inp.Label = trl.S{
					"en": `
						Todo: Abstimmung des exakten Textes zwischen ZEW und Partner
						<br>

						<b>Einwilligungserklärung gemäß DSGVO</b>

						<br>

						Die Antworten dieser Online-Umfrage werden von uns streng vertraulich, 
						DSGVO-konform behandelt und nur in anonymer bzw. aggregierter Form benutzt.

						<br>

						Im <a href="/doc/site-imprint.md" >Impressum</a> finden Sie umfangreiche Angaben zum Datenschutz.

						<br>

						Hiermit willige ich ein, dass meine gesammelten Daten 
						für [Private Debt Survey] des [ZEW] verwendet werden.

						<br>

					`,
				}

				inp.ControlFirst()
				inp.ControlTop()
			}

		}

		// gr0
		{
			labels := []trl.S{
				{
					"en": `Ich erkläre mich einverstanden, 
					dass meine angegebenen Daten zu Auswertungszwecken an [partner_1] 
					weitergeleitet werden.
					`,
				},

				{
					"en": `Meine Daten sollen <i>nicht</i> an [partner_1] 
					weitergeleitet werden.
					`,
				},
			}
			radioValues := []string{
				"datasharing_yes",
				// "datasharing_anonymous",
				"datasharing_not",
			}

			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": `
				Todo: <br>
				Text Weitergabe meiner Daten an [partner_2]<br>

				Zusammen mit Identifikation am Anfang?<br>
				Identifikation hierher ans Ende?<br>


				`,
				}
				inp.ColSpan = gr.Cols
			}

			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q62_sharing"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
				rad.ControlTop()

				rad.Validator = "mustRadioGroup"

			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": `Fragebogen abschließen um die Daten final zu speichern.`,
					"de": `Fragebogen abschließen um die Daten final zu speichern.`,
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

		// pge.ExampleSixColumnsLabelRight()

	}

	//
	//
	// Report of results
	{
		p := q.AddPage()
		p.NoNavigation = true
		p.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}
		{
			// gr := p.AddGroup()
			// gr.Cols = 1
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.DynamicFunc = "RepsonseStatistics"
			// }
		}
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	q.SetColspans()

	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
