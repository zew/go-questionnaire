package qst

import (
	"fmt"
	"log"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/trl"
)

// RadioVali - name of validation func for radio in composite funcs
var RadioVali = ""

// HeaderClass -  CSS class for first row with labels
var HeaderClass = ""

// CSSLabelRow -  CSS class radio input in rows
var CSSLabelRow = ""

/* AddRadioGroupVertical prints options vertically

   Green   x
   Red     x
   Black   x

*/

// gbCol - a grid builder column;
// element of grid-column-template;
// optional header content
type gbCol struct {
	header      trl.S // first row / header - horizontally over input
	spanLabel   float32
	spanControl float32
	cells       []inputT // filled programmatically by addRadioRow...
}

// GridBuilder to generate a matrix or grid or table of labels and inputs;
// cells are stored column-wise - enforcing equal spans for label and control;
// columns together constitute a CSS-grid-column-template
// rendering is done row-wise;
type GridBuilder struct {
	MainLabel trl.S // first row - before column headers - as wide as the group
	cols      []gbCol
}

// AddCol adds a column;
// to the grid builder template
func (gb *GridBuilder) AddCol(headerCell trl.S, spanLabel, spanControl float32) {
	col := gbCol{}
	col.header = headerCell
	col.spanLabel = spanLabel
	col.spanControl = spanControl
	gb.cols = append(gb.cols, col)
}

// AddRadioRow adds radio inputs - empty columns are filled with empty text
func (gb *GridBuilder) AddRadioRow(name string, vals []string, sparseLabels map[int]trl.S) {

	if len(gb.cols) < 1 {
		log.Panicf("RadioMatrix2.addRadioRow(name) - no cols defined")
	}
	if name == "" {
		log.Panicf("RadioMatrix2.addRadioRow(name) - name is empty")
	}
	if len(vals) < 2 {
		log.Panicf("RadioMatrix2.addRadioRow(name) - at least 2 radio values")
	}

	for colIdx := 0; colIdx < len(gb.cols); colIdx++ {

		rad := InputEmpty()
		if colIdx < len(vals) {
			rad = InputEmpty()
			rad.Label = nil
			rad.Type = "radio"
			rad.Name = name // "y_euro"
			rad.ValueRadio = vals[colIdx]

		}

		if _, ok := sparseLabels[colIdx]; ok {
			rad.Label = sparseLabels[colIdx]
			rad.StyleLbl = css.ItemStart(rad.StyleLbl)
		}

		rad.ColSpanLabel = gb.cols[colIdx].spanLabel
		rad.ColSpanControl = gb.cols[colIdx].spanControl

		gb.cols[colIdx].cells = append(gb.cols[colIdx].cells, *rad)

	}

}

// AddRadioRow adds radio inputs - empty columns are filled with empty text
func (gb *GridBuilder) dumpCols() {
	w := &strings.Builder{}
	cntr := float32(0.0)
	for colIdx, col := range gb.cols {
		fmt.Fprintf(w, "%v - %v %v \n", colIdx, col.spanLabel, col.spanControl)
		cntr += col.spanLabel
		cntr += col.spanControl
	}
	log.Printf("\n%v  total %v", w.String(), cntr)
}

// AddGrid creates static entries to the page;
// being prepared using AddCol() and AddRadioRow()
func (p *pageT) AddGrid(gb *GridBuilder) *groupT {

	// gb.dumpCols()

	gr := p.AddGroup()

	gr.Cols = 0
	for colIdx := 0; colIdx < len(gb.cols); colIdx++ {
		gr.Cols += gb.cols[colIdx].spanLabel
		gr.Cols += gb.cols[colIdx].spanControl
	}

	gr.Style = css.NewStylesResponsive()
	stl := ""
	for colIdx := 0; colIdx < len(gb.cols); colIdx++ {
		if gb.cols[colIdx].spanLabel != 0 {
			stl = fmt.Sprintf("%v   %vfr ", stl, gb.cols[colIdx].spanLabel)
		}
		if gb.cols[colIdx].spanControl != 0 {
			stl = fmt.Sprintf("%v   %vfr ", stl, gb.cols[colIdx].spanControl)
		}
	}
	// gr.Style.Desktop.GridContainerStyle.TemplateColumns = stl

	// first row - main label
	if gb.MainLabel != nil {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = gb.MainLabel
		inp.ColSpanLabel = gr.Cols
	}

	// second row - headers - preflight
	headersExist := false
	for colIdx := 0; colIdx < len(gb.cols); colIdx++ {
		if gb.cols[colIdx].header != nil {
			headersExist = true
			break
		}
	}
	// second row - headers - execution
	if headersExist {
		for colIdx := 0; colIdx < len(gb.cols); colIdx++ {
			if gb.cols[colIdx].header != nil {

				// space cell - horizontally over labels
				if gb.cols[colIdx].spanLabel > 0 {
					inp2 := gr.addInputEmpty()
					// s := fmt.Sprintf("col%v", colIdx)
					// inp2.Label = trl.S{"de": s, "en": s}
					inp2.ColSpanLabel = gb.cols[colIdx].spanLabel
				}

				// label cell - horizontally over controls
				inp2 := gr.addInputEmpty()
				inp2.Label = gb.cols[colIdx].header
				inp2.ColSpanLabel = gb.cols[colIdx].spanControl
				inp2.Style = css.ItemCentered(inp2.Style)
				inp2.Style.Desktop.GridItemStyle.AlignSelf = "end"
				inp2.Style.Desktop.GridItemStyle.AlignSelf = "start"
				// inp2.Style.Desktop.BoxStyle.Padding = "0 0.2rem"

			} else {
				inp := gr.addInputEmpty()
				inp.ColSpanLabel = gb.cols[colIdx].spanLabel + gb.cols[colIdx].spanControl
			}
		}
	}

	// subsequent rows - consisting of inputs
	// notice nesting inside out - to resolve column-wise structuring
	for rowIdx := 0; rowIdx < len(gb.cols[0].cells); rowIdx++ {
		for colIdx := 0; colIdx < len(gb.cols); colIdx++ {
			gr.addInputArg(&gb.cols[colIdx].cells[rowIdx])
		}
	}

	return gr

}

// NewGridBuilderRadios for FMT questionnaire
func NewGridBuilderRadios(
	columnTemplate []float32,
	hdrLabels []trl.S,
	inputNames []string,
	radioVals []string,
	firstColLabels []trl.S,
) *GridBuilder {

	gb := &GridBuilder{}

	if len(columnTemplate) != len(hdrLabels)*2 {
		log.Panicf("NewGridBuilderRadios(): len(columnTemplate) != len(hdrLabels)*2 - %v != %v", len(columnTemplate), len(hdrLabels)*2)
	}

	// Setup of columns
	for i := 0; i < len(columnTemplate); i += 2 {
		gb.AddCol(hdrLabels[i/2], columnTemplate[i], columnTemplate[i+1])
	}

	// adding rows
	// for rowIdx := 0; rowIdx < len(firstColLabels); rowIdx++ {
	for rowIdx := 0; rowIdx < len(inputNames); rowIdx++ {
		name := inputNames[rowIdx]

		lbl := trl.S{}
		sparseLbls := map[int]trl.S{}
		if rowIdx < len(firstColLabels) {
			lbl = firstColLabels[rowIdx]
			sparseLbls[0] = lbl
		}

		// sparseLbls[3] = trl.S{"de": "--", "en": "--"}
		gb.AddRadioRow(name, radioVals, sparseLbls)
	}

	return gb

}
