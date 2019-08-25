package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/trl"
)

var RadioVali = ""
var CSSLabelHeader = ""
var CSSLabelRow = ""

/* AddRadioGroupVertical prints ooptions vertically

   Green   x
   Red     x
   Black   x

*/
func (p *pageT) AddRadioGroupVertical(name string, rowLabels []trl.S) *groupT {

	//
	var gr *groupT
	gr = p.AddGroup()
	gr.Cols = 1
	radGroup := gr.AddInput()
	radGroup.Type = "radiogroup"
	radGroup.Name = name // "y0_euro"
	radGroup.CSSLabel = CSSLabelRow
	radGroup.Validator = RadioVali

	for idx, lbl := range rowLabels {

		rad := radGroup.AddRadio()
		rad.HAlign = HRight
		rad.Label = lbl
		rad.Col = idx % 2 // 0 1 --  0 1
		rad.Cols = 2

	}

	return gr
}

// AddRadioMatrixGroup adds several inputs of type radiogroup
// and prepends a row with labels.
func (p *pageT) AddRadioMatrixGroup(headerLabels []trl.S, inpNames []string, rowLabels []trl.S, opt ...int) *groupT {

	gr := p.AddGroup()

	colSpanLabel := 1
	if len(opt) > 0 {
		colSpanLabel = opt[0]
	}

	// top-left cell => empty
	if len(rowLabels) > 0 {
		if len(rowLabels) != len(inpNames) { // consistence check, deliberately inside if condition
			panic("AddRadioMatrixGroup(): if row labels exist, they should exist for *all* rows")
		}
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = trl.S{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		}
		inp.ColSpanLabel = colSpanLabel
		inp.CSSLabel = CSSLabelHeader // apply even if its empty
	}

	// Header row - next columns
	for _, lbl := range headerLabels {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.HAlignLabel = HCenter
		inp.Desc = lbl // for instance trl.S{"de": "gut", "en": "good"}
		inp.CSSLabel = CSSLabelHeader
	}

	//
	for i1, name := range inpNames {
		radGroup := gr.AddInput()
		radGroup.Type = "radiogroup"
		radGroup.Name = name // "y0_euro"
		radGroup.ColSpanLabel = colSpanLabel
		radGroup.CSSLabel = CSSLabelRow
		radGroup.Validator = RadioVali
		if len(rowLabels) > 0 {
			// for instance trl.S{"de": "Euroraum", "en": "euro area"}
			// inp.Label = rowLabels[i]
			radGroup.Desc = rowLabels[i1]
		}
		for i2 := 0; i2 < len(headerLabels); i2++ {
			rad := radGroup.AddRadio()
			rad.HAlign = HCenter
		}

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

// EmptyCells to fill up rows
func (gr *groupT) EmptyCells(rowSpan int) {
	inp := gr.AddInput()
	inp.Type = "textblock"
	inp.Name = fmt.Sprintf("vspacer%02d", ctr.Increment())
	inp.Label = trl.S{"de": " ", "en": " "}
	inp.Desc = trl.S{"de": " ", "en": " "}
	inp.ColSpanLabel = rowSpan
}
