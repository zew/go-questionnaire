package pds

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var assetClassesInputs = []string{
	"real_estate",
	"xx",
	"corp_lending",
}
var assetClassesLabels = []trl.S{
	{
		"en": "Real estate",
		"de": "Immobilien",
	},
	{
		"en": "xx",
		"de": "xx",
	},
	{
		"en": "Corporate lending",
		"de": "Corporate lending",
	},
}

func assetClass(page *qst.WrappedPageT) {

	lblMain := trl.S{
		"en": `Which asset classes do you invest in?<br> 
				We ask a similar range of questions for each selected asset class.
			`,
		"de": `Wählen Sie Ihre Assetklassen.<br>
				Wir fragen für jede Anlageklasse die gleiche Serie von Fragen.
			`,
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
		gr.Cols = 3
		// gr.BottomVSpacers = 1

		for idx1 := 0; idx1 < len(assetClassesInputs); idx1++ {
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = assetClassesInputs[idx1]
			inp.Label = assetClassesLabels[idx1]
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Vertical()
			inp.VerticalLabel()

			labelBottom := false
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			if labelBottom {
				inp.StyleLbl.Desktop.StyleGridItem.Order = 2
			} else {
				// top
				inp.StyleLbl.Desktop.StyleBox.Position = "relative"
				inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			}
		}

	}
}
