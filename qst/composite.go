package qst

import "github.com/zew/go-questionaire/trl"

// AddRadioMatrixGroup adds several inputs of type radiogroup
// and prepends a row with labels.
func (p *pageT) AddRadioMatrixGroup(headerLabels []trl.S, inpNames []string, rowLabels []trl.S, opt ...int) *groupT {

	gr := p.AddGroup()

	colSpanLabel := 1
	if len(opt) > 0 {
		colSpanLabel = opt[0]
	}

	// Header row - first column - empty cell
	if len(rowLabels) > 0 {
		if len(rowLabels) != len(inpNames) { // consistence check
			panic("radioMatrix(): if row labels exist, they should exist for *all* rows")
		}
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = trl.S{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		}
		inp.ColSpanLabel = colSpanLabel
	}

	// Header row - next columns
	for _, lbl := range headerLabels {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.HAlignLabel = HCenter
		inp.Desc = lbl // for instance trl.S{"de": "gut", "en": "good"}
	}

	//
	for i, name := range inpNames {
		inp := gr.AddInput()
		inp.Type = "radiogroup"
		inp.Name = name // "y0_euro"
		if len(rowLabels) > 0 {
			// for instance trl.S{"de": "Euroraum", "en": "euro area"}
			// inp.Label = rowLabels[i]
			inp.Desc = rowLabels[i]
		}
		for i2 := 0; i2 < len(headerLabels); i2++ {
			rad := inp.AddRadio()
			rad.HAlign = HCenter
		}
		inp.ColSpanLabel = colSpanLabel

	}

	return gr

}

func (p *pageT) ExampleSixColumnsLabelRight() *groupT {

	grp := p.AddGroup()
	inp := grp.AddInput()
	inp.Type = "radiogroup"
	inp.Name = "layoutlang"
	inp.Label = map[string]string{"de": "Layouting Archtitektur", "en": "Layout architecture"}
	inp.Desc = map[string]string{"de": "<br>Label rechts, Rest zentriert", "en": "<br>label right, rest centered"}
	inp.HAlignLabel = HRight
	inp.HAlignControl = HLeft
	inp.Radios = []*radioT{
		{Label: trl.S{"de": "HTML", "en": "HTML"}, HAlign: HCenter},
		{Label: trl.S{"de": "Winword", "en": "Microsoft Word"}, HAlign: HCenter},
		{Label: trl.S{"de": "Tec", "en": "Tec - Latec"}, HAlign: HCenter},
		{Label: trl.S{"de": "Markdown", "en": "Markdown"}, HAlign: HCenter},
		{Label: trl.S{"de": "Breakdown", "en": "Breakdown"}, HAlign: HCenter},
	}

	grp.Cols = 6
	grp.Label = map[string]string{"de": "Fünf mit Label", "en": "Five with label"}
	grp.Desc = map[string]string{"de": "", "en": ""}
	return grp
}
