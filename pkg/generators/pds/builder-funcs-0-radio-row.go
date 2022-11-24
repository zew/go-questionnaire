package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
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
}

var (
	chb01 = configMC{
		Cols: 3,
	}
	chb02 = configMC{
		Cols: 4,
	}

	mCh1 = configMC{
		KeyLabels:   "ac1-tranche-types",
		Cols:        8,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,
	}
	mCh2 = configMC{
		KeyLabels:   "teamsize",
		Cols:        16,
		InpColspan:  16 / 4,
		LabelBottom: true,
		DontKnow:    false,
	}
	mCh3 = configMC{
		KeyLabels:   "relevance1-5",
		Cols:        10,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,
	}
	mCh4 = configMC{
		KeyLabels:   "improveDecline1-5",
		Cols:        10,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    false,
	}
	mChExample1 = configMC{
		KeyLabels:   "improveWorsen1-5",
		Cols:        14,
		InpColspan:  2,
		LabelBottom: false,
		DontKnow:    true,
	}
	mChExample2 = configMC{
		KeyLabels:   "smaller5larger20",
		Cols:        16,
		InpColspan:  16 / 4,
		LabelBottom: true,
		DontKnow:    false,
	}
)

var assetClassesInputs = []string{
	"ac1_corplending",
	"ac2_realestate",
	"ac3_infrastruct",
}

var assetClassesLabels = []trl.S{
	{
		"en": "Corporate / direct lending",
		"de": "Corporate / direct lending",
	},
	{
		"en": "Real estate debt",
		"de": "Real estate debt",
	},
	{
		"en": "Infrastructure debt",
		"de": "Infrastructure debt",
	},
}

// strategy, strategies
var trancheTypeNamesAC1 = []string{
	"st1_senior",
	"st2_unittranche",
	"st3_subordinated",
	"st4_mezzanine", // "mezzanine_pik_other",
}
var trancheTypeNamesAC2 = []string{
	"st1_wholeloan",
	"st2_subordinated",
}
var trancheTypeNamesAC3 = []string{
	"st1_senior",
	"st2_subordinated",
}

var allLbls = map[string][]trl.S{
	"ac1-tranche-types": {
		{
			"en": "Senior",
			"de": "Senior",
		},
		{
			"en": "Unitranche",
			"de": "Unitranche",
		},
		{
			"en": "Subordinated",
			"de": "Subordinated",
		},
		{
			"en": "Mezzanine / PIK / Other",
			"de": "Mezzanine / PIK / Other",
		},
	},
	"ac2-tranche-types": {
		{
			"en": "Whole Loan",
			"de": "Whole Loan",
		},
		{
			"en": "Subordinated",
			"de": "Subordinated",
		},
	},
	"ac3-tranche-types": {
		{
			"en": "Senior",
			"de": "Senior",
		},
		{
			"en": "Subordinated",
			"de": "Subordinated",
		},
	},
	"teamsize": {
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
	},
	"relevance1-5": {
		{
			"en": "(1)<br>Not relevant",
			"de": "(1)<br>Not relevant",
		},
		{
			"en": "(2)<br>",
			"de": "(2)<br>",
		},
		{
			"en": "(3)<br>",
			"de": "(3)<br>",
		},
		{
			"en": "(4)<br>",
			"de": "(4)<br>",
		},
		{
			"en": "(5)<br> Potential dealbreaker",
			"de": "(5)<br> Potential dealbreaker",
		},
	},
	"improveDecline1-5": {
		{
			"en": "Improved",
			"de": "Improved",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "Same",
			"de": "Same",
		},
		{
			"en": "&nbsp;",
			"de": "&nbsp;",
		},
		{
			"en": "Declined",
			"de": "Declined",
		},
	},

	"improveWorsen1-5": {
		{
			"en": "Improve significantly",
			"de": "Improve significantly",
		},
		{
			"en": "Improve slightly",
			"de": "Improve slightly",
		},
		{
			"en": "Stay constant",
			"de": "Stay constant",
		},
		{
			"en": "Worsen slightly",
			"de": "Worsen slightly",
		},
		{
			"en": "Worsen significantly",
			"de": "Worsen significantly",
		},
	},
	"smaller5larger20": {
		{
			"en": "<5&nbsp;FTE",
			"de": "<5&nbsp;FTE",
		},
		{
			"en": "5-10&nbsp;FTE",
			"de": "5-10&nbsp;FTE",
		},
		{
			"en": "11-20&nbsp;FTE",
			"de": "11-20&nbsp;FTE",
		},
		{
			"en": ">20&nbsp;FTE",
			"de": ">20&nbsp;FTE",
		},
	},
}

var lblDont = trl.S{
	"de": "Don´t know",
	"en": "Don´t know",
}

// radiosSingleRow - five shades - and no answer
// previously "multipleChoice"
func radiosSingleRow(
	page *qst.WrappedPageT,
	nm string,
	lblMain trl.S,
	cf configMC,
) {

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblMain
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	// gr2
	{
		gr := page.AddGroup()
		gr.Cols = 14
		gr.Cols = float32(cf.Cols)
		// if mode == 2 {
		// 	gr.Cols = 16
		// }
		gr.BottomVSpacers = 3
		if cf.GroupBottomSpacers != 0 {
			gr.BottomVSpacers = cf.GroupBottomSpacers
		}

		for idx2 := 0; idx2 < len(allLbls[cf.KeyLabels]); idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = allLbls[cf.KeyLabels][idx2]
			inp.ColSpan = 2
			inp.ColSpan = cf.InpColspan
			// if mode == 2 {
			// 	inp.ColSpan = gr.Cols / 4
			// }
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			if cf.LabelBottom {
				inp.StyleLbl.Desktop.StyleGridItem.Order = 2
			} else {
				// top
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			}
		}

		if cf.DontKnow {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", len(allLbls[cf.KeyLabels])+1)
			inp.Label = lblDont
			inp.ColSpan = 4
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()
		}

	}

}
