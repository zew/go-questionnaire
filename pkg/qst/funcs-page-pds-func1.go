package qst

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func restrictedTextMultiCols(
	page *pageT,
	ac assetClass,
	cf configRT,
) {

	numCols := firstColLbl + float32(len(ac.TrancheTypes))
	idxLastCol := len(ac.TrancheTypes) - 1

	{
		gr := page.AddGroup()
		gr.Cols = numCols
		// gr.BottomVSpacers =

		// row0 - column headers, 1-3 tranche type names
		for idx1 := 0; idx1 < len(ac.TrancheTypes)+1; idx1++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if idx1 == 0 {
				inp.ColSpan = firstColLbl
			}
			if idx1 > 0 {
				ttLbl := ac.TrancheTypes[idx1-1].Lbl
				inp.Label = ttLbl
			}
			inp.StyleLbl = styleHeaderCols1
		}

		if cf.SuppressSumField {

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols

			combined := trl.S{"en": "", "de": ""}
			combined.Append(cf.LblRow1)
			combined.Append90(cf.LblRow2)
			// combined.AppendStr("<div style='margin-left: 2.2rem;'>")
			// combined.Append(cf.LblRow2)
			// combined.AppendStr("</div>")

			inp.Label = combined

		} else {

			// row1
			// sum field
			for idx1, trancheType := range ac.TrancheTypes {

				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("%v_%v_%v_%v", ac.Prefix, trancheType.Prefix, cf.InputNameP2, "main")

					inp.MaxChars = cf.Chars
					inp.Step = 1
					if cf.Step != 0.0 {
						inp.Step = cf.Step
					}
					inp.Min = 0
					inp.Max = 1000 * 1000 * 1000
					// inp.Validator = "inRange100"

					inp.Placeholder = cf.Placeholder

					inp.ColSpan = 1
					inp.ColSpanLabel = 0
					inp.ColSpanControl = 1

					if idx1 == 0 {
						inp.Label = cf.LblRow1.Bold()
						inp.Label = cf.LblRow1
						inp.ColSpan = firstColLbl + 1
						inp.ColSpanLabel = firstColLbl
						inp.ColSpanControl = 1
					}
					if idx1 == idxLastCol {
						inp.Suffix = cf.Suffix
					}

					if cf.FirstRow100Pct {
						// seems to get lost due to dynamic page?
						inp.Response = "100" // must parse to number
						inp.Placeholder = placeHolder100percent
						inp.Disabled = true
					}
					if cf.FirstRowDisabled {
						inp.Disabled = true
					}
				}
			}

			// row2
			// sub headline
			if !cf.LblRow2.Empty() {
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols

				inp.Label = cf.LblRow2

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0 0 0 2.2rem"
				inp.Style.Mobile.StyleBox.Margin = "0 0 0 0"
				// inp.Style.Desktop.StyleBox.Width = "60%"
				// inp.Style.Mobile.StyleBox.Width = "96%"
			}
		}

		//
		//
		// rows 3,4...
		for _, suffx := range cf.SubNames {

			for idx2, trancheType := range ac.TrancheTypes {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v_%v", ac.Prefix, trancheType.Prefix, cf.InputNameP2, suffx)

				inp.MaxChars = cf.Chars
				inp.Step = 1
				if cf.Step != 0.0 {
					inp.Step = cf.Step
				}
				inp.Min = 0
				inp.Max = 1000 * 1000
				// inp.Validator = "inRange100"

				inp.Placeholder = cf.Placeholder

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				// input vertically bottom
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "start"
				inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "end"

				if idx2 == 0 {
					inp.Label = trl.S{
						"en": fmt.Sprintf("%v", cf.SubLbls[suffx]),
						"de": fmt.Sprintf("%v", cf.SubLbls[suffx]),
					}
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1

					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
					inp.StyleLbl.Desktop.StyleBox.Margin = "0 0 0 2.2rem"
					inp.StyleLbl.Mobile.StyleBox.Margin = "0 0 0 0"
				}

				if idx2 == idxLastCol {
					inp.Suffix = cf.Suffix
				}

			}

		}

		if !cf.SuppressSumField && len(cf.SubNames) > 0 {

			for _, trancheType := range ac.TrancheTypes {

				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "restrictedTextSum-a" // js filename

				inp.JSBlockStrings = map[string]string{}
				inp.JSBlockStrings["InpMain"] = fmt.Sprintf("%v_%v_%v", ac.Prefix, trancheType.Prefix, cf.InputNameP2)

				sns := make([]string, 0, len(cf.SubNames))
				for _, name := range cf.SubNames {
					sn := fmt.Sprintf("\"%v_%v_%v_%v\"", ac.Prefix, trancheType.Prefix, cf.InputNameP2, name)
					sns = append(sns, sn)
					// key := fmt.Sprintf("%v_%v", "inp", idx1+1) // {{.inp_1}}, {{.inp_2}}, ...
				}
				inp.JSBlockStrings["SummandNames"] = "\n\t" + strings.Join(sns, ",\n\t") + "\n\t"

				lblClean := cf.LblRow1.RemoveSomeHTML()
				s1 := trl.S{
					"en": fmt.Sprintf("Please check Question 1.1%v - asset class %v. The sum of transaction volumes for the individual market segments does not add up to the total transaction volume. Really continue?", lblClean["en"], trancheType.Lbl),
					"de": fmt.Sprintf("Please check Question 1.1%v - asset class %v. The sum of transaction volumes for the individual market segments does not add up to the total transaction volume. Really continue?", lblClean["de"], trancheType.Lbl),
				}
				inp.JSBlockStrings["CmpOperator"] = "unequal"
				if cf.AddendsLighterSum {
					s1 = trl.S{
						"en": fmt.Sprintf("Please check Question 1.1 a) asset class %v. The total number of transaction is smaller than the number of transactions in the subgroup below. Really continue?", trancheType.Lbl),
						"de": fmt.Sprintf("Please check Question 1.1 a) asset class %v. The total number of transaction is smaller than the number of transactions in the subgroup below. Really continue?", trancheType.Lbl),
					}
					inp.JSBlockStrings["CmpOperator"] = "noneGreater"
					inp.Name = "restrictedTextSum-b" // js filename
				}
				inp.JSBlockTrls = map[string]trl.S{
					"msg": s1.RemoveSomeHTML(),
				}

			}
		}

	} // /group

}
