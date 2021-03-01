package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/trl"
)

// RadioVali - name of validation func for radio in composite funcs
var RadioVali = ""

// CSSLabelHeader -  CSS class for first row with labels
var CSSLabelHeader = ""

// CSSLabelRow -  CSS class radio input in rows
var CSSLabelRow = ""

/* AddRadioGroupVertical prints options vertically

   Green   x
   Red     x
   Black   x

*/
func (p *pageT) AddRadioGroupVertical(name string, rowLabels []trl.S, cols int) *groupT {

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

		if cols == 2 {
			rad.Col = idx % 2 // 0 1 --  0 1
			rad.Cols = 2
		}

		// does not work yet
		if cols == 1 {
			rad.Col = 0
			rad.Cols = 2
		}

	}

	return gr
}

// AddRadioMatrixGroup adds several inputs of type radiogroup
// and prepends a row with labels.
// len(rowLabels) == 0  => no first column
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

// AddRadioMatrixGroupNoLabels - like AddRadioMatrixGroup
// but no first column with labels
func (p *pageT) AddRadioMatrixGroupNoLabels(headerLabels []trl.S, inpNames []string, opt ...int) *groupT {

	gr := p.AddGroup()

	colSpanLabel := 1
	if len(opt) > 0 {
		colSpanLabel = opt[0]
	}

	// Header row - next columns
	for _, lbl := range headerLabels {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.HAlignLabel = HCenter
		inp.Desc = lbl // for instance trl.S{"de": "gut", "en": "good"}
		inp.CSSLabel = CSSLabelHeader
		inp.CSSLabel = "special-vertical-align-top"
	}

	//
	for _, name := range inpNames {
		radGroup := gr.AddInput()
		radGroup.Type = "radiogroup"
		radGroup.Name = name // "y0_euro"
		radGroup.ColSpanLabel = colSpanLabel
		radGroup.CSSLabel = CSSLabelRow
		radGroup.Validator = RadioVali
		for i2 := 0; i2 < len(headerLabels); i2++ {
			rad := radGroup.AddRadio()
			rad.HAlign = HCenter
		}

	}

	return gr

}

// ExampleSixColumnsLabelRight - five radio in horizontal row
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

//
//
// AddRadioMatrixGroupCSSGrid adds several inputs of type radiogroup;
// labels1stRow => header row
// labels1stCol => header col
// spanCol1     => wider first col
func (p *pageT) AddRadioMatrixGroupCSSGrid(
	inpNames []string, inpCols int,
	labels1stRow []trl.S,
	labels1stCol []trl.S,
	span1stCol int,
) *groupT {

	has1stRowLabel := false
	if len(labels1stRow) > 0 {
		has1stRowLabel = true
	}

	// consistence check
	if has1stRowLabel {
		if inpCols != len(labels1stRow) {
			panic("AddRadioMatrixGroup(): if labels for 1st row exist, they should exist for *all* cols")
		}
	}

	has1stColLabel := false
	if len(labels1stCol) > 0 {
		has1stColLabel = true
	}

	// consistence check
	if has1stColLabel {
		// deliberately inside if condition
		if len(labels1stCol) != len(inpNames) {
			panic("AddRadioMatrixGroup(): if labels for 1st col exist, they should exist for *all* rows")
		}
		if span1stCol == 0 {
			span1stCol = 1
		}
	} else {
		span1stCol = 0
	}

	//
	gr := p.AddGroup()
	gr.Cols = span1stCol + inpCols
	// log.Printf("group cols = %v  -  span1stCol + inpCols  <=> %v  %v", inpCols+span1stCol, span1stCol, inpCols)

	//
	// Header row - first column
	// (empty cell in top-left)
	if has1stRowLabel && has1stColLabel {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = trl.S{
			"de": " &nbsp; ",
			"en": " &nbsp; ",
		}
		inp.ColSpanLabel = span1stCol
		inp.CSSLabel = CSSLabelHeader // apply even if its empty
	}

	//
	// Header row - next columns
	for _, lbl := range labels1stRow {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Style = css.NewStylesResponsive()
		inp.Style.Desktop.GridItemStyle.JustifySelf = "center"
		inp.Style.Desktop.GridItemStyle.AlignSelf = "center"
		inp.Style.Desktop.TextStyle.AlignHorizontal = "center"
		inp.StyleLbl = inp.Style
		inp.Desc = lbl // for instance trl.S{"de": "gut", "en": "good"}
		inp.CSSLabel = CSSLabelHeader
	}

	//
	// Input rows
	for rowIdx, name := range inpNames {

		if has1stColLabel {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = labels1stCol[rowIdx]
			inp.ColSpanLabel = span1stCol
		}

		radGroup := gr.AddInput()
		radGroup.Type = "radiogroup"
		radGroup.Name = name // "y0_euro"
		radGroup.ColSpanLabel = 0
		radGroup.ColSpanControl = inpCols
		radGroup.Validator = RadioVali

		for inpCol := 0; inpCol < inpCols; inpCol++ {
			rad := radGroup.AddRadio()
			rad.HAlign = HCenter
		}

	}

	return gr

}
