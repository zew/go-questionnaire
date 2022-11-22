package pds

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func restrictedText2b(
	page *qst.WrappedPageT,
	// ttPref string,
	lbl trl.S,
	cf configRT,
) {

	for idx0, trancheType := range trancheTypeNames {

		ttPref := trancheType[:3]
		ttLbl := allLbls["tranche-types"][idx0]

		_, _ = ttPref, ttLbl

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_main", ttPref, cf.InputToken2)

				inp.Label = lbl
				inp.Suffix = cf.Suffix

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
				inp.Label = cf.SubLbl

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Width = "60%"
				inp.Style.Mobile.StyleBox.Width = "96%"

			}

			for _, suffx := range cf.SubNames {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v", ttPref, cf.InputToken2, suffx)

				inp.Label = trl.S{
					"en": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
					"de": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
					//
				}
				inp.Suffix = cf.Suffix

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
			inp.Name = "restrictedTextSum"

			s1 := trl.S{
				"de": "Addiert sicht nicht auf. Trotzdem weiter?",
				"en": "Does not add up. Continue?",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}

			inp.JSBlockStrings = map[string]string{}
			//

			inpMain := fmt.Sprintf("%v_%v", ttPref, cf.InputToken2)
			inp.JSBlockStrings["InpMain"] = inpMain

			summands := []string{}
			for _, suffx := range cf.SubNames {
				nm := fmt.Sprintf("%v_%v", inpMain, suffx)
				summands = append(summands, nm)
			}

			inp.JSBlockStrings["SummandNames"] = `"` + strings.Join(summands, `", "`) + `"`

		}
	}

}
