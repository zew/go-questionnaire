package kneb1

import (
	"math"

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

	q.Survey.Org = trl.S{"en": "ZEW"}
	q.Survey.Name = trl.S{"en": "Financial Literacy Test"}
	// q.Variations = 1

	// page 1
	{
		p := q.AddPage()

		p.Section = trl.S{"en": "Sociodemographics"}
		p.Label = trl.S{"en": "Age, origin, experience"}
		p.Short = trl.S{"en": "Sociodemo-<br>graphics"}
		p.WidthMax("42rem")

		// gr0
		{
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q01_age"

				inp.Label = trl.S{"en": "How old are you?"}
				inp.MaxChars = 4
				inp.Step = 1
				inp.Min = 15
				inp.Max = 150
				inp.Validator = "inRange100"

				inp.ColSpan = 4
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
				// inp.Suffix = trl.S{"en": "years"}
				inp.Suffix = trl.S{"en": "&nbsp; years"}
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q02_country_birth"

				inp.Label = trl.S{"en": "What is your country of birth?"}

				inp.ColSpan = 4
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
				inp.MaxChars = 20
				// inp.Validator = "must"
			}

		}

		// gr1
		{
			var radioValues = []string{
				"male",
				"female",
				"diverse",
			}
			var labels = []trl.S{
				{"de": "Male"},
				{"de": "Female"},
				{"de": "Diverse"},
			}

			gr := p.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "What is your gender?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q03_gender"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.ColSpan = gr.Cols
				rad.Label = label
				rad.ControlFirst()
			}
		}

		// gr2
		{
			var radioValues = []string{
				"yes",
				"no",
			}
			var labels = []trl.S{
				{"de": "Yes"},
				{"de": "No"},
			}

			gr := p.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "Are you currently an <i>exchange</i> student in Mannheim?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q04_visiting_student"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.ColSpan = gr.Cols
				rad.Label = label
				rad.ControlFirst()
			}

			//
			{
				inp := gr.AddInput()
				inp.Name = "q04a_country_study"

				inp.Label = trl.S{"en": "If yes: from which country?"}
				inp.Type = "text"
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 5
				inp.MaxChars = 20

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0 0 0 2.1rem"

			}
		}

		// gr3
		{
			var radioValues = []string{
				"yes",
				"no",
			}
			var labels = []trl.S{
				{"de": "Yes"},
				{"de": "No"},
			}

			gr := p.AddGroup()
			gr.Cols = 7
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "Do you already have professional experience from working in a job (other than a student job)?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q05_professional_experience"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.ColSpan = gr.Cols
				rad.Label = label
				rad.ControlFirst()
			}
		}

		// gr4
		{
			var radioValues = []string{
				"yes",
				"no",
			}
			var labels = []trl.S{
				{"de": "Yes"},
				{"de": "No"},
			}

			gr := p.AddGroup()
			gr.Cols = 7
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.Margin = "0 0 0 4rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "If yes: was this experience related to finance?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q06_professional_finance"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.ColSpan = gr.Cols
				rad.Label = label
				rad.ControlFirst()
			}
		}

		// gr5
		{
			var radioValues = []string{
				"yes",
				"no",
			}
			var labels = []trl.S{
				{"de": "Yes"},
				{"de": "No"},
			}

			gr := p.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "Have you ever worked with a statistics program like Stata or&nbspR?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q07_stata_or_r"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.ColSpan = gr.Cols
				rad.Label = label
				rad.ControlFirst()
			}
		}

		// gr6
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6a,
				labelsGoodBad6(),
				[]string{"q07_health"},
				[]string{"very_good", "good", "average", "bad", "very_bad", "no_answer"},
				[]trl.S{
					{"de": ``},
					{"en": ``},
				},
			)

			gb.MainLabel = trl.S{
				"de": "Sie schätzen Ihren Gesundheitszustand als",
				"en": "Would you say your health&nbsp;is",
			}
			gr := p.AddGrid(gb)
			gr.BottomVSpacers = 4
		}

	} // page1

	//
	// page 2
	{

		p := q.AddPage()
		p.WidthMax("40rem")
		p.Section = trl.S{"en": "Financial decisions"}
		p.Label = trl.S{"en": "Wealth, assets"}
		p.Short = trl.S{"en": "Financial<br>decisions"}

		{
			//
			// Income
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q08_income"

				inp.Label = trl.S{"en": `
					<b>What is your current income?</b> 
					<br>
					Please consider all sources of current income from work, transfers from your parents, stipends etc.
				`}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}

				inp.ControlBottom()
				inp.LabelPadRight()
			}

			//
			// Wealth
			gr = p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "<b>What is approximately the amount of wealth you own</b>"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q09a_wealth_liquid"

				inp.Label = trl.S{"en": "In cash (also on a debit account, savings account etc.)"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q09b_wealth_fungible"

				inp.Label = trl.S{"en": "In assets, e.g. in stocks, bonds, mutual funds etc."}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q09c_wealth_re"

				inp.Label = trl.S{"en": "In real estate"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q09d_wealth_other"

				inp.Label = trl.S{"en": "In other valuables (e.g. cars, durables, art etc.)"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}
			}

			//
			// Debt
			gr = p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "<b>What is the amount of debt you owe</b>"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 4
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q10a_debt_bank"

				inp.Label = trl.S{"en": "To a bank"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q10b_debt_other"

				inp.Label = trl.S{"en": "To another institution"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}

			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q10c_debt_parents"

				inp.Label = trl.S{"en": "To your parents"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q10d_debt_friends"

				inp.Label = trl.S{"en": "To other friends or relatives"}
				inp.MaxChars = 9
				inp.Step = 1
				inp.Min = 0
				inp.Max = math.MaxInt
				inp.Validator = "inRange1Mio"

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"en": "&nbsp€"}
			}

		} // income, wealth, debt

		//
		// type of fin assets
		{
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "<b>Do you currently own/use any of the following financial products?</b>"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			fieldnames := []string{
				"q11_asset_investment_account",
				"q11_asset_savings_account",
				"q11_asset_stocks",
				"q11_asset_mutual",
				"q11_asset_bonds",
				"q11_asset_credit_card",
				"q11_asset_insurance",
				"q11_asset_mobile_account",
				"q11_asset_crypto",
				"q11_asset_no_answer",
			}

			labels := []string{
				"An investment account",
				"A savings account",
				"Stocks and shares",
				"Mutual funds",
				"Bonds",
				"A credit card",
				"An insurance contract (except mandatory health insurance)",
				"A mobile phone payment app (e.g. Google pay, Apple pay etc.)",
				"Crypto assets (e.g. Bitcoin) or ICOs (Initial Coin Offerings)",
				"No answer",
			}

			for idx, fn := range fieldnames {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fn

				inp.Label = trl.S{"en": labels[idx]}
				inp.ColSpan = 2
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
			}

		}

		//
		// fin activity
		{
			gr := p.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"en": "<b>Do you do any of the following for yourself or your household?</b>"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			fieldnames := []string{
				"q12_beh_planning",
				"q12_beh_expenses_notes",
				"q12_beh_money_stashes",
				"q12_beh_planning_due_payments",
				"q12_beh_app",
				"q12_beh_automatic_payments",
				"q12_beh_no_reply",
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

			for idx, fn := range fieldnames {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fn

				inp.Label = trl.S{"en": labels[idx]}
				inp.ColSpan = 2
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 2
			}
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5a,
				labelsYesNo5(),
				[]string{"q13_broke_before_paycheck"},
				[]string{"yes", "no", "dont_know", "not_applicable", "no_answer"},
				[]trl.S{
					{"en": ``},
				},
			)

			gb.MainLabel = trl.S{
				"en": `Sometimes people find that their income does not quite cover their living expenses. 
					In the last 12&nbsp;months, has this happened to you personally?`,
			}
			gr := p.AddGrid(gb)
			gr.BottomVSpacers = 4
		}

	}

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
