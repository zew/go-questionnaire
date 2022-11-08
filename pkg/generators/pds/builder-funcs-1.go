package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var subLbl1 = trl.S{
	"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by market segment: `,
	"de": `Please state the volume (in million Euro) of deals closed in Q4 2022 by market segment: `,
}

func restrictedText(
	page *qst.WrappedPageT,
	name string,
	lbl trl.S,
	subLbl trl.S,
) {

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_main", name)

			inp.Label = lbl
			inp.Suffix = trl.S{"de": "Million Euro"}

			inp.MaxChars = 12
			inp.Step = 1
			inp.Min = 0
			inp.Max = 1000 * 1000
			// inp.Validator = "inRange100"

			inp.ColSpan = 1
			inp.ColSpanLabel = 3
			inp.ColSpanControl = 2
		}

	}

	// gr2
	{
		gr := page.AddGroup()
		gr.Cols = 2
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Margin = "0 0 0 2rem"
		gr.BottomVSpacers = 3

		{

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 2
			inp.Label = subLbl

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Width = "60%"
			inp.Style.Mobile.StyleBox.Width = "96%"

		}

		lbls := map[string]string{
			"low":   "Lower Mid-Market (0-15m € EBITDA)",
			"mid":   "Core Mid-Market (15-50m € EBITDA)",
			"upper": "Upper Mid-Market (>50m € EBITDA)",
		}

		for _, suffx := range []string{"low", "mid", "upper"} {

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v", name, suffx)

			inp.Label = trl.S{
				"de": fmt.Sprintf("- %v", lbls[suffx]),
			}
			inp.Suffix = trl.S{"de": "Million Euro"}

			inp.MaxChars = 10
			inp.Step = 1
			inp.Min = 0
			inp.Max = 1000 * 1000
			// inp.Validator = "inRange100"

			inp.ColSpan = 2
			inp.ColSpanLabel = 3
			inp.ColSpanControl = 2
		}

		inp := gr.AddInput()
		inp.ColSpanControl = 1
		inp.Type = "javascript-block"
		inp.Name = "lowMidUpperSum"

		s1 := trl.S{
			"de": "Addiert sicht nicht auf. Trotzdem weiter?",
			"en": "Does not add up. Continue?",
		}
		inp.JSBlockTrls = map[string]trl.S{
			"msg": s1,
		}

		inp.JSBlockStrings = map[string]string{}
		for idx, suffx := range []string{"main", "low", "mid", "upper"} {
			inpNamePlaceholder := fmt.Sprintf("%v_%v", "inp", idx+1) // {{.inp_1}}, {{.inp_2}}, ...
			nm := fmt.Sprintf("%v_%v", name, suffx)
			inp.JSBlockStrings[inpNamePlaceholder] = nm
		}

	}

}
