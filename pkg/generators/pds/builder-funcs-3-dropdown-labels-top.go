package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func dropdownsLabelsTop(
	page *qst.WrappedPageT,
	nm string,
	lbl trl.S,
	cf configMC,
) {

	numCols := firstColLbl + float32(len(trancheTypeNamesAC1))
	// numCols := float32(len(trancheTypeNamesAC1))

	{
		gr := page.AddGroup()
		gr.Cols = numCols

		// row0 - headers
		for idx1 := 0; idx1 < len(trancheTypeNamesAC1)+1; idx1++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if idx1 == 0 {
				inp.ColSpan = firstColLbl
			}
			if idx1 > 0 {
				ttLbl := allLbls["ac1-tranche-types"][idx1-1]
				// inp.Label = ttLbl.Bold()
				inp.Label = ttLbl
			}
			inp.LabelVertical()

			inp.StyleLbl = trancheNameStyle
		}

		// row1
		for idx1, trancheType := range trancheTypeNamesAC1 {

			ttPref := trancheType[:3]

			{
				inp := gr.AddInput()
				inp.Type = "dropdown"
				inp.Name = fmt.Sprintf("%v_%v", ttPref, nm)
				inp.MaxChars = 20
				inp.MaxChars = 10

				inp.ColSpan = 1
				inp.ColSpanControl = 1

				inp.DD = &qst.DropdownT{}

				if false {
					inp.DD.AddPleaseSelect(trl.CoreTranslations()["must_one_option"])
					for idx2 := 0; idx2 < len(allLbls[cf.KeyLabels]); idx2++ {
						inp.DD.Add(
							fmt.Sprintf("opt_%02v", idx2),
							allLbls[cf.KeyLabels][idx2],
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
