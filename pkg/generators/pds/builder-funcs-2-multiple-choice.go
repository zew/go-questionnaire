package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// config multiple choice
type configMC struct {
	KeyLabels          string
	Cols               float32
	InpColspan         float32
	LabelBottom        bool
	DontKnow           bool
	GroupBottomSpacers int
}

var (
	mCh1 = configMC{
		KeyLabels:   "tranche-types",
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

var trancheTypeNames = []string{
	"senior",
	"unit",       // unitranche
	"subord",     // subordinated
	"mezzpiketc", // "mezzanine_pik_other",
}

var allLbls = map[string][]trl.S{
	"tranche-types": {
		{
			"de": "Senior",
			"en": "Senior",
		},
		{
			"de": "Unitranche",
			"en": "Unitranche",
		},
		{
			"de": "Subordinated",
			"en": "Subordinated",
		},
		{
			"de": "Mezzanine / PIK / Other",
			"en": "Mezzanine / PIK / Other",
		},
	},
	"teamsize": {
		{
			"de": "<5",
			"en": "<5",
		},
		{
			"de": "5-10",
			"en": "5-10",
		},
		{
			"de": "11-20",
			"en": "11-20",
		},
		{
			"de": ">20",
			"en": ">20",
		},
	},
	"improveWorsen1-5": {
		{
			"de": "Improve significantly",
			"en": "Improve significantly",
		},
		{
			"de": "Improve slightly",
			"en": "Improve slightly",
		},
		{
			"de": "Stay constant",
			"en": "Stay constant",
		},
		{
			"de": "Worsen slightly",
			"en": "Worsen slightly",
		},
		{
			"de": "Worsen significantly",
			"en": "Worsen significantly",
		},
	},
	"smaller5larger20": {
		{
			"de": "<5&nbsp;FTE",
			"en": "<5&nbsp;FTE",
		},
		{
			"de": "5-10&nbsp;FTE",
			"en": "5-10&nbsp;FTE",
		},
		{
			"de": "11-20&nbsp;FTE",
			"en": "11-20&nbsp;FTE",
		},
		{
			"de": ">20&nbsp;FTE",
			"en": ">20&nbsp;FTE",
		},
	},
}

var lblDont = trl.S{
	"de": "Don´t know",
	"en": "Don´t know",
}

// multipleChoiceSingleRow - five shades - and no answer
func multipleChoiceSingleRow(
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
		gr.BottomVSpacers = 1
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
