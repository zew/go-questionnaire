package flit

import (
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	qst.RadioVali = "mustRadioGroup"
	qst.RadioCSSControl = "special-input-margin-vertical"

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("flit")
	q.Survey.Params = params
	q.LangCodes = map[string]string{"en": "English"}
	q.LangCodesOrder = []string{"en"} // governs default language code

	q.Survey.Org = trl.S{"en": "ZEW"}
	q.Survey.Name = trl.S{"en": "Financial Literacy Test"}
	// q.Variations = 1

	// page 1
	{
		p := q.AddPage()

		p.Section = trl.S{"en": "Sociodemographics"}
		p.Label = trl.S{"en": "Experience"}
		p.Short = trl.S{"en": "Experience"}
		p.Width = 75

		gr := p.AddGroup()
		gr.Cols = 4

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "age"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "How old are you?"}
			inp.MaxChars = 3
			inp.Step = 1
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.Suffix = trl.S{"en": "&nbsp; years"}
			inp.Validator = "inRange100"
		}

		gr.EmptyCells(1)

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "country_birth"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "What is your country of birth?"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.MaxChars = 20
			inp.Validator = "must"
		}

		gr.EmptyCells(1)

		{
			inp := gr.AddInput()
			inp.Type = "radiogroup"
			inp.Name = "gender"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "What is your gender?"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			inp.Validator = "mustRadioGroup"
			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "Male",
				}
			}

			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "Female",
				}
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "radiogroup"
			inp.Name = "visiting_student"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "Are you currently an exchange student in Mannheim?"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			inp.Validator = "mustRadioGroup"

			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "Yes",
				}
			}
			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "No",
				}
			}
		}

		{
			inp := gr.AddInput()
			inp.Name = "country_study"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider special-input-left-padding"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "If yes: from which country?"}
			// inp.HAlignLabel = qst.HRight
			inp.Type = "text"
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.MaxChars = 20

		}

		gr.EmptyCells(1)

		// row
		{
			inp := gr.AddInput()
			inp.Type = "radiogroup"
			inp.Name = "professional_experience"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "Do you already have professional experience from working in a job (other than a student job)?"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			inp.Validator = "mustRadioGroup"

			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "Yes",
				}
			}
			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "No",
				}
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "radiogroup"
			inp.Name = "professional_finance"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider special-input-left-padding"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "If yes: was this experience related to finance?"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "Yes",
				}
			}
			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "No",
				}
			}
		}

		//
		// stata
		{
			inp := gr.AddInput()
			inp.Type = "radiogroup"
			inp.Name = "experience_stata_or_r"
			inp.CSSLabel = "special-input-margin-vertical special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": "Have you ever worked with a statistics program like Stata or&nbspR?"}
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2
			inp.Validator = "mustRadioGroup"

			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "Yes",
				}
			}
			{
				rad := inp.AddRadio()
				// rad.HAlign = qst.HLeft
				rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"en": "No",
				}
			}
		}

		gr.EmptyCells(1)

	} // page1

	// page 2
	{

		p := q.AddPage()
		p.Width = 80
		p.Section = trl.S{"en": "Sociodemographics"}
		// p.Label = trl.S{"en": "Financial literacy, wealth, assets, preferences"}
		p.Label = trl.S{"en": "Health, wealth, assets"}
		p.Short = trl.S{"en": "Sociodemo&shy;graphics&nbsp;"}

		{
			lbls := []trl.S{
				{
					"de": "Sie schätzen Ihren Gesundheitszustand als",
					"en": "Would you say your health&nbsp;is",
				},
			}
			flds := []string{
				"health",
			}
			gr := p.AddRadioMatrixGroup(labelsGoodBad6(), flds, lbls, 2)
			gr.Cols = 8 //
			gr.BottomVSpacers = 2
		}

		{
			//
			// Income
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "income"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": "What is your current income"}
				inp.Desc = trl.S{"en": "<br>Please consider all sources of current income from work, transfers from your parents, stipends etc."}
				inp.MaxChars = 6
				inp.Step = 10
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange50000"
			}
			gr.EmptyCells(1)

			//
			// Wealth
			gr = p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "textblock1"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": "What is approximately the amount of wealth you own"}
				inp.Desc = trl.S{"en": " "}
				inp.ColSpanLabel = 2
			}
			gr.EmptyCells(2)

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "wealth_liquid"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "In cash (also on a debit account, savings account etc.)"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"

			}
			gr.EmptyCells(1)
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "wealth_fungible"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "In assets, e.g. in stocks, bonds, mutual funds etc."}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"

			}
			gr.EmptyCells(1)
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "wealth_re"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "In real estate"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"

			}
			gr.EmptyCells(1)
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "wealth_other"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "In other valuables (e.g. cars, durables, art etc.)"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"

			}
			gr.EmptyCells(1)

			//
			// Debt
			gr = p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "textblock2"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": "What is the amount of debt you owe"}
				inp.Desc = trl.S{"en": " "}
				inp.ColSpanLabel = 2

			}
			gr.EmptyCells(2)

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "debt_bank"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "To a bank"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"

			}
			gr.EmptyCells(1)
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "debt_other"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "To another institution"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"

			}
			gr.EmptyCells(1)
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "debt_parents"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "To your parents"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"
			}
			gr.EmptyCells(1)
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "debt_friends"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": "To other friends or relatives"}
				inp.MaxChars = 9
				inp.Step = 100
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbspEuros"}
				inp.Validator = "inRange1Mio"
			}
			gr.EmptyCells(1)

		} // income, wealth, debt

		{
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "textblock4"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": "Do you currenlty own/use any of the following financial products?"}
				inp.Desc = trl.S{"en": " "}
				inp.ColSpanLabel = 2

			}
			gr.EmptyCells(2)

			assets := []string{
				"asset_investment_account",
				"asset_savings_account",
				"asset_stocks",
				"asset_mutual",
				"asset_bonds",
				"asset_credit_card",
				"asset_insurance",
				"asset_mobile_account",
				"asset_crypto",
				"asset_no_answer",
			}

			labels := []string{
				"An investment account",
				"A savings account",
				"Stocks and shares",
				"Mutual funds",
				"Bonds",
				"A credit card",
				"An insurance contract (except health insurance, which is mandatory)",
				"A mobile phone payment app (e.g. google pay, Apple pay etc.)",
				"Crypto assets (e.g. Bitcoin) or ICOs (Initial Coin Offerings)",
				"No answer",
			}

			for idx, ass := range assets {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = ass
				// inp.CSSLabel = "special-input-vert-wider"
				inp.CSSLabel = "special-input-margin-vertical"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": labels[idx]}
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
			}

		}

		{
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "textblock3"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Label = trl.S{"en": "Do you do any of the following for yourself or your household?"}
				inp.Desc = trl.S{"en": " "}
				inp.ColSpanLabel = 2

			}
			gr.EmptyCells(2)

			assets := []string{
				"beh_planning",
				"beh_expenses_notes",
				"beh_money_stashes",
				"beh_planning_due_payments",
				"beh_app",
				"beh_automatic_payments",
				"beh_no_reply",
			}

			labels := []string{
				"Make a plan to manage your income and expenses",
				"Keep a note of your spending",
				"Keep money for bills separate from day-to-day spending money",
				"Make a note of upcoming bills to make sure you  don't miss them",
				"Use a banking app or money management tool to keep track of your outgoings",
				"Arrange automatic payments for regular outgoings",
				"No answer",
			}

			for idx, ass := range assets {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = ass
				// inp.CSSLabel = "special-input-vert-wider"
				inp.CSSLabel = "special-input-margin-vertical"
				inp.Label = trl.S{"en": " "}
				inp.Desc = trl.S{"en": labels[idx]}
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
			}
			gr.EmptyCells(2)

		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Sometimes people find that their income does not quite cover their living expenses. 
					In the last 12&nbsp;months, has this happened to you personally?`,
				},
			}
			flds := []string{
				"broke_before_paycheck",
			}
			gr := p.AddRadioMatrixGroup(labelsYesNo5(), flds, lbls, 3)
			gr.Cols = 8 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": "How you would rate your overall knowledge about financial matters compared with other adults?",
				},
			}
			flds := []string{
				"self_assessment",
			}
			gr := p.AddRadioMatrixGroup(labelsHighLow6(), flds, lbls, 2)
			gr.Cols = 8 //
			gr.BottomVSpacers = 2
		}

	} // p2

	// p3
	{
		p := q.AddPage()
		p.Width = 80
		p.Section = trl.S{"en": "Financial literacy 1"}
		// p.Label = trl.S{"en": "Financial literacy, wealth, assets, preferences"}
		p.Label = trl.S{"en": ""}
		p.Short = trl.S{"en": "Fin. literacy 1"}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": "Which of the following statements describes the main function of the stock market?",
				},
			}
			flds := []string{
				"stock_market_function",
			}
			gr := p.AddRadioMatrixGroup(stockMarketFuncs6(), flds, lbls, 2)
			gr.Cols = 8 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": "Which of the following statements is correct?",
				},
			}
			flds := []string{
				"mutual_fund_withdrawal",
			}
			gr := p.AddRadioMatrixGroup(mutualFunds6(), flds, lbls, 2)
			gr.Cols = 8 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Do you think that the following statement is true or false? 
					Buying a single company stock usually provides a safer return than a stock mutual fund.`,
				},
			}
			flds := []string{
				"mutual_fund_risk",
			}
			gr := p.AddRadioMatrixGroup(trueFalse4(), flds, lbls, 3)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": "Normally, which asset displays the highest fluctuations over time?",
				},
			}
			flds := []string{
				"asset_variance",
			}
			gr := p.AddRadioMatrixGroup(assetVariance5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Considering a long time period (for example 20 years), which asset normal gives the highest average return`,
				},
			}
			flds := []string{
				"max_long_term_return",
			}
			gr := p.AddRadioMatrixGroup(highestReturn5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `When an investor spreads his money among different assets, does the risk of losing money `,
				},
			}
			flds := []string{
				"diversification_risk",
			}
			gr := p.AddRadioMatrixGroup(diversification5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

	} // p3

	// p4
	{
		p := q.AddPage()
		p.Width = 80
		p.Section = trl.S{"en": "Financial literacy 2"}
		// p.Label = trl.S{"en": "Financial literacy, wealth, assets, preferences"}
		p.Label = trl.S{"en": ""}
		p.Short = trl.S{"en": "Fin. literacy 2"}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Suppose you had 100 Euros in a savings account and the interest rate was 2% per year. 
					After 5 years, how much do you think you would have in the account if you left the money to grow?`,
				},
			}
			flds := []string{
				"compound_interest",
			}
			gr := p.AddRadioMatrixGroup(compounding5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Imagine that the interest rate on your savings account was 1% per year and inflation was 2% per year. 
					After 1 year, would you be able to buy more than, exactly the same as, or less than today with the money in this account?`,
				},
			}
			flds := []string{
				"real_interest",
			}
			gr := p.AddRadioMatrixGroup(realRate5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Suppose you had 100 Euros in a savings account and the interest rate is 20% per year and you never withdraw money or interest payments.  
					After 5 years, how much would you have on this account in total?`,
				},
			}
			flds := []string{
				"compound_multi",
			}
			gr := p.AddRadioMatrixGroup(compoundingMulti5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `Suppose that in the year 2020, your income has doubled and prices of all goods have doubled too. 
					In 2020, how much will you be able to buy with your income?`,
				},
			}
			flds := []string{
				"inflation_indexing",
			}
			gr := p.AddRadioMatrixGroup(inflationIndexIncome5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `If the interest rates fall, what happens to bond prices?`,
				},
			}
			flds := []string{
				"interest_bond_price",
			}
			gr := p.AddRadioMatrixGroup(interestBondPrice5(), flds, lbls, 2)
			gr.Cols = 7 //
			gr.BottomVSpacers = 2
		}

	} // p4

	// p5
	{
		p := q.AddPage()
		p.Width = 80
		p.Section = trl.S{"en": "Financial literacy 3"}
		// p.Label = trl.S{"en": "Financial literacy, wealth, assets, preferences"}
		p.Label = trl.S{"en": ""}
		p.Short = trl.S{"en": "Fin. literacy 2"}

		gr := p.AddGroup()
		gr.Cols = 3
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "cost_of_ball"
			inp.CSSLabel = "special-input-vert-wider special-input-margin-vertical"
			inp.Label = trl.S{"en": ""}
			inp.Desc = trl.S{"en": `A bat and a ball cost $1.10 in total. The bat costs $1.00 more than the ball.
			How much does the ball cost?`}
			inp.MaxChars = 4
			inp.Step = 0.01
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.Suffix = trl.S{"en": "&nbspCents"}
			inp.Validator = "inRange100"
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "minutes_100_widgets"
			inp.CSSLabel = "special-input-vert-wider special-input-margin-vertical"
			inp.Label = trl.S{"en": ""}
			inp.Desc = trl.S{"en": `If it takes 5&nbsp;machines 5&nbsp;minutes to make 5&nbsp;widgets, 
			how long would it take 100 machines to make 100 widgets? `}
			inp.MaxChars = 4
			inp.Step = 0.1
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.Suffix = trl.S{"en": "&nbspMinutes"}
			inp.Validator = "inRange100"
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "days_covering_half_lake"
			inp.CSSLabel = "special-input-vert-wider special-input-margin-vertical"
			inp.Label = trl.S{"en": ""}
			inp.Desc = trl.S{"en": `In a lake there is a patch of lily pads. Every day, the patch doubles in size. 
			If it takes 48 days for the patch to cover the entire lake, how long would it take
			for the patch to cover half of the lake?`}
			inp.MaxChars = 4
			inp.Step = 1
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.Suffix = trl.S{"en": "&nbspDays"}
			inp.Validator = "inRange100"
		}

		{
			lbls := []trl.S{
				{
					"de": ``,
					"en": `How do you see yourself – how willing are you in general to take risks?`,
				},
			}
			flds := []string{
				"risk_preference",
			}
			gr := p.AddRadioMatrixGroup(riskPreference7(), flds, lbls, 3)
			gr.Cols = 10 //
			gr.BottomVSpacers = 3
		}

		{
			gr := p.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = 0

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Name = "textblock6"
			inp.CSSLabel = "special-input-vert-wider"
			inp.Label = trl.S{"en": " "}
			inp.Desc = trl.S{"en": `How do you see yourself on a scale from<br>
			1 – I completely disagree<br>
			&nbsp; to <br> 
			7 - I completely agree`}
			inp.Desc = trl.S{"en": `How do you see yourself on a scale from ...`}
			inp.ColSpanLabel = 3
		}

		{
			lbls := []trl.S{
				{
					"de": "xxx",
					"en": `I am good at resisting temptation`,
				},
				{
					"de": "xxx",
					"en": `I refuse things that are bad for me`,
				},
				{
					"de": "xxx",
					"en": `I wish I had more self discipline (R)`,
				},
				{
					"de": "xxx",
					"en": `Pleasure and fun sometimes keep me from getting work done (R)`,
				},
				{
					"de": "xxx",
					"en": `I have trouble concentrating (R)`,
				},
				{
					"de": "xxx",
					"en": `I am able to work effectively towards long term goals`,
				},
				{
					"de": "xxx",
					"en": `Sometimes I cant stop myself from doing something, even if I know it is wrong (R)`,
				},
				{
					"de": "xxx",
					"en": `I often act without thinking through the alternatives (R)`,
				},
				{
					"de": "xxx",
					"en": `I am impulsive and tend to buy things even when I can’t really afford them.`,
				},
				{
					"de": "xxx",
					"en": `I set financial goals for the next 1–2 months for what I want to achieve with my money.`,
				},
			}
			flds := []string{
				"temptation_resistance",
				"bad_things_refusal",
				"more_self_discipline",
				"pleasure_prevents_me",
				"trouble_concentrating",
				"long_term_goal_oriented",
				"cannot_stop_myself",
				"acting_without_thinking",
				"impulsive_purchases",
				"financial_goals",
			}
			gr := p.AddRadioMatrixGroup(agreeToDisagree7(), flds, lbls, 3)
			gr.Cols = 10 //
			gr.BottomVSpacers = 2
		}

	} // p5

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
