package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func dropdownsLabelsTop(
	page *pageT,
	ac assetClass,
	nm string,
	lbl trl.S,
	cf configMC,
) {

	numCols := firstColLbl + float32(len(ac.TrancheTypes))

	{
		gr := page.AddGroup()
		gr.Cols = numCols

		// row0 - headers
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
			inp.LabelVertical()

			inp.StyleLbl = styleHeaderCols1
		}

		// row1
		for idx1, trancheType := range ac.TrancheTypes {

			{
				inp := gr.AddInput()
				inp.Type = "dropdown"
				inp.Name = fmt.Sprintf("%v_%v_%v", ac.Prefix, trancheType.Prefix, nm)
				inp.MaxChars = 20
				inp.MaxChars = 10

				inp.ColSpan = 1
				inp.ColSpanControl = 1

				inp.DD = &DropdownT{}

				if false {
					inp.DD.AddPleaseSelect(cfg.Get().Mp["must_one_option"])
					for idx2 := 0; idx2 < len(PDSLbls[cf.KeyLabels]); idx2++ {
						inp.DD.Add(
							fmt.Sprintf("opt_%02v", idx2),
							PDSLbls[cf.KeyLabels][idx2],
						)
					}
				}
				inp.DD.Add(
					"",
					trl.S{"en": " Please choose"},
				)
				inp.DD.Add(
					"lt6weeks",
					trl.S{"en": "&nbsp;<6 weeks"},
				)
				inp.DD.Add(
					"6weeks",
					trl.S{"en": "&nbsp;&nbsp;&nbsp;6 weeks"},
				)
				inp.DD.Add(
					"9weeks",
					trl.S{"en": "&nbsp;&nbsp;&nbsp;9 weeks"},
				)
				inp.DD.Add(
					"12weeks",
					trl.S{"en": "&nbsp;12 weeks"},
				)
				inp.DD.Add(
					"18weeks",
					trl.S{"en": "&nbsp;18 weeks"},
				)
				inp.DD.Add(
					"gt18weeks",
					trl.S{"en": ">18 weeks"},
				)

				if idx1 == 0 {
					inp.Label = lbl
					inp.LabelPadRight()
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1
				}

			}
		}

	}
}
