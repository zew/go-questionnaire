package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// sentimentSingleRow - five shades - and no answer
func sentimentSingleRow(
	page *qst.WrappedPageT,
	nm string,
	lblMain trl.S,
	mode int,
) {

	lbls := []trl.S{
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
			"de": "Improve significantly",
			"en": "Improve significantly",
		},
	}

	if mode == 2 {
		lbls = []trl.S{
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
		}

	}

	lblDont := trl.S{
		"de": "Don´t know",
		"en": "Don´t know",
	}

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
		if mode == 2 {
			gr.Cols = 16
		}
		gr.BottomVSpacers = 1

		for idx2 := 0; idx2 < len(lbls); idx2++ {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
			inp.Label = lbls[idx2]
			inp.ColSpan = 2
			if mode == 2 {
				inp.ColSpan = gr.Cols / 4
			}
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			if mode == 2 {
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleGridItem.Order = 2
			}
		}

		if mode == 1 {
			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = fmt.Sprintf("%v", nm)
			inp.ValueRadio = fmt.Sprintf("%v", len(lbls)+1)
			inp.Label = lblDont
			inp.ColSpan = 4
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()
		}

	}

}
