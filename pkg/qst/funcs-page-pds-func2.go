package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

// config multiple choice
type configMC struct {
	KeyLabels          string // key to a map of labels
	Cols               float32
	InpColspan         float32
	LabelBottom        bool
	DontKnow           bool
	GroupBottomSpacers int

	GroupLeftIndent string

	XDisplacements []string
}

var (
	mCh2 = configMC{
		KeyLabels:          "teamsize",
		Cols:               16,
		InpColspan:         16 / 4,
		LabelBottom:        true,
		DontKnow:           false,
		GroupBottomSpacers: 3,
	}

	mCh2a = configMC{
		KeyLabels:          "covenants-per-credit",
		Cols:               4,
		InpColspan:         1,
		LabelBottom:        false,
		DontKnow:           false,
		GroupBottomSpacers: 3,
		GroupLeftIndent:    outline2Indent,

		XDisplacements: []string{
			"1.6rem",
			"0.62rem",
			"0.62rem",
			"1.6rem",
		},
	}

	mCh3 = configMC{
		KeyLabels:   "relevance1-5",
		Cols:        10,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,
	}

	mCh4 = configMC{
		KeyLabels:       "improveDecline1-5",
		Cols:            10,
		InpColspan:      2,
		LabelBottom:     false,
		DontKnow:        false,
		GroupLeftIndent: outline2Indent,
		XDisplacements: []string{
			"1.6rem",
			"0.79rem",
			"",
			"0.79rem",
			"1.6rem",
		},
	}
	mCh5 = configMC{
		KeyLabels:   "closing-time-weeks",
		Cols:        14,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,

		// not yet
		// GroupLeftIndent: outline2Indent,

		XDisplacements: []string{
			"1.46rem",
			"1.27rem",
			"0.64rem",
			"",
			"0.64rem",
			"1.27rem",
			"1.46rem",
		},
	}
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
					inp.DD.AddPleaseSelect(trl.CoreTranslations()["must_one_option"])
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
