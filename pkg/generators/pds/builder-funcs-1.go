package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var lblsTeamSize = []trl.S{
	{
		"en": "<5",
		"de": "<5",
	},
	{
		"en": "5-10",
		"de": "5-10",
	},
	{
		"en": "11-20",
		"de": "11-20",
	},
	{
		"en": ">20",
		"de": ">20",
	},
}

// [<1 bn EUR, 1-5 bn EUR, 5-10 bn EUR, >10 bn EUR]

var lblsAssetsUnderMgt = []trl.S{
	{
		"en": "<1bn&nbsp;€",
		"de": "<1bn&nbsp;€",
	},
	{
		"en": "1-5bn&nbsp;€",
		"de": "1-5bn&nbsp;€",
	},
	{
		"en": "5-10bn&nbsp;€",
		"de": "5-10bn&nbsp;€",
	},
	{
		"en": ">10bn&nbsp;€",
		"de": ">10bn&nbsp;€",
	},
}

// checkBoxCascade for hierarchical checkboxes from page1
func checkBoxCascade(
	page *qst.WrappedPageT,
	lblMain trl.S,
) {

	numCols := float32(5)
	// numCols := float32(4)

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblMain
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

	}

	// gr2 - 3-4, 5-6
	// for idx1 := 0; idx1 < len(inpsL1); idx1++ {
	for idx1, ac := range qst.PDSAssetClasses {

		_ = idx1

		gr := page.AddGroup()
		gr.Cols = numCols
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.05rem"
		gr.BottomVSpacers = 2

		// row1
		{
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = fmt.Sprintf("%v_q03", ac.Prefix)
			// inp.Name = fmt.Sprintf("%v_q03", ac.Name)
			inp.Label = ac.Lbl
			inp.ColSpan = gr.Cols
			inp.ColSpanControl = 10
			inp.ColSpanLabel = 1
			inp.ControlFirst()

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Margin = "0.4rem 0 0 0"
		}

		//
		//
		// row2
		// 		indented label
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.DisplayNone()
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 4
			inp.Label = trl.S{
				"en": "Which strategies do you engage in?",
				"de": "In welchen Strategien engagieren Sie sich?",
			}
			inp.ColSpanLabel = 1
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.FontSize = 90
			inp.DisplayNone()
		}

		// row3
		// 		indented level2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.DisplayNone()
		}

		for idx2, trancheType := range ac.TrancheTypes {

			_ = idx2

			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = fmt.Sprintf("%v_%v_q031", ac.Prefix, trancheType.Prefix)
			inp.Label = trancheType.Lbl
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Position = "relative"
			inp.Style.Desktop.StyleBox.Top = "-0.4rem"

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.FontSize = 90

			//
			inp.StyleLbl.Desktop.StyleBox.Position = "relative"
			inp.StyleLbl.Desktop.StyleBox.Top = "0"
			inp.StyleLbl.Mobile.StyleBox.Top = "0.3rem"

			inp.DisplayNone()

		}
		if len(ac.TrancheTypes) == 3 {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
			}
		}
		if len(ac.TrancheTypes) == 2 {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.DisplayNone()
			}
		}

		//
		//
		// row4
		// 		indented label
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.DisplayNone()
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 4
			inp.Label = trl.S{
				"de": `<br>How big is your investment team? Please choose team size in full time equivalents.`,
				"en": `<br>How big is your investment team? Please choose team size in full time equivalents.`,
			}
			inp.ColSpanLabel = 1
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.FontSize = 90

			inp.StyleLbl.Desktop.StyleBox.Position = "relative"
			inp.StyleLbl.Desktop.StyleBox.Top = "0.3rem"
			inp.DisplayNone()
		}
		// row5
		// 		indented level2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.DisplayNone()
		}
		for idx2 := 0; idx2 < len(lblsTeamSize); idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v_q032", ac.Prefix)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = lblsTeamSize[idx2]
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			if false {
				// inp.StyleLbl.Desktop.StyleGridItem.Order = 2
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			}
			inp.DisplayNone()

		}

		//
		//
		// row6
		// 		indented label
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.DisplayNone()
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 4
			inp.Label = trl.S{
				"de": `<br>What are your Assets under Management in this asset class?`,
				"en": `<br>What are your Assets under Management in this asset class?`,
			}
			inp.ColSpanLabel = 1
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.FontSize = 90

			inp.StyleLbl.Desktop.StyleBox.Position = "relative"
			inp.StyleLbl.Desktop.StyleBox.Top = "0.3rem"
			inp.DisplayNone()
		}
		// row7
		// 		indented level2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.DisplayNone()
		}
		for idx2 := 0; idx2 < len(lblsAssetsUnderMgt); idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v_q033", ac.Prefix)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = lblsAssetsUnderMgt[idx2]
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			if false {
				// inp.StyleLbl.Desktop.StyleGridItem.Order = 2
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			}
			inp.DisplayNone()

		}

	}
}

// consent for repeating DSGVO
func consent(
	q *qst.QuestionnaireT,
	page *qst.WrappedPageT,
	instance int,
) {

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"en": `
						<span style='font-size:110%'>

						<b>Declaration of consent according to GDPR</b> 
						
						<br>

						We will treat the answers to this online survey strictly confidential, 
						GDPR-compliant and only use them in anonymous or aggregated form. 
						
						We will pass on your answers to the questions 
						within the ZEW Private Debt Survey to our cooperation 
						partner Prime Capital AG in an aggregated and anonymous form. 
					
						 
						In the <a href="/doc/site-imprint.md" >imprint</a> you find extensive information on data protection.						
						</span>

						<br> <!-- vertical space for the must error message -->
						<br> <!-- vertical space for the must error message -->
						
						`,

				"de": `---`,
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = fmt.Sprintf("q61_consent_%v", instance)
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 6
			// inp.Validator = "must"
			inp.Label = trl.S{
				"en": `I hereby consent to my collected data being used for the ZEW Private Debt Survey.`,
				"de": `I hereby consent to my collected data being used for the ZEW Private Debt Survey.`,
			}
			// inp.Response = "1"
			inp.ControlFirst()
			inp.ControlTop()
		}
	}

}
