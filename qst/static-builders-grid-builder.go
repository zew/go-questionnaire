package qst

import (
	"fmt"
	"log"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/trl"
)

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
	validator string
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

// AddRadioRow adds a row radio inputs - empty columns are filled with empty text;
// sparse labels allows to set labels for selected columns only; often only the first columns has a label
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

		rad := emptyTextblock()
		if colIdx < len(vals) {
			rad = emptyTextblock()
			rad.Label = nil
			rad.Type = "radio"
			rad.Name = name // "y_euro"
			rad.Validator = gb.validator
			rad.ValueRadio = vals[colIdx]
		}

		if _, ok := sparseLabels[colIdx]; ok {
			rad.Label = sparseLabels[colIdx]
			rad.StyleLbl = css.TextStart(rad.StyleLbl)
		}

		rad.ColSpanLabel = gb.cols[colIdx].spanLabel
		rad.ColSpanControl = gb.cols[colIdx].spanControl

		gb.cols[colIdx].cells = append(gb.cols[colIdx].cells, *rad)

	}

}

func (gb *GridBuilder) dumpCols() {
	w := &strings.Builder{}
	cntr := float32(0.0)
	for colIdx, col := range gb.cols {
		fmt.Fprintf(w, "%v - %4.1f %4.1f \n", colIdx, col.spanLabel, col.spanControl)
		cntr += col.spanLabel
		cntr += col.spanControl
	}
	log.Printf("\n%v  total %v", w.String(), cntr)
}

// -----------------------------------------------------------------

// NewGridBuilderRadios - see NewGridBuilderRadiosWithValidator
func NewGridBuilderRadios(
	columnTemplate []float32,
	hdrLabels []trl.S,
	inputNames []string,
	radioVals []string,
	firstColLabels []trl.S,
) *GridBuilder {

	return NewGridBuilderRadiosWithValidator(
		columnTemplate,
		hdrLabels,
		inputNames,
		radioVals,
		firstColLabels,
		"",
	)

}

// NewGridBuilderRadiosWithValidator
// uses GridBuilder.AddCol() and GridBuilder.AddRadioRow()
// to compile a matrix of cells;
// basically, each cell can contain input names, input values and input labels
//
// firstColLabels can be nil
func NewGridBuilderRadiosWithValidator(
	columnTemplate []float32,
	hdrLabels []trl.S,
	inputNames []string,
	radioVals []string,
	firstColLabels []trl.S,
	Validator string,
) *GridBuilder {

	gb := &GridBuilder{}

	gb.validator = Validator

	if len(hdrLabels) > 0 {
		if len(columnTemplate) != len(hdrLabels)*2 {
			log.Panicf("NewGridBuilderRadios(): len(columnTemplate) != len(hdrLabels)*2 - %v != %v", len(columnTemplate), len(hdrLabels)*2)
		}
		// Setup of columns with header labels
		for i := 0; i < len(columnTemplate); i += 2 {
			gb.AddCol(hdrLabels[i/2], columnTemplate[i], columnTemplate[i+1])
		}
	} else {
		// Setup of columns without
		for i := 0; i < len(columnTemplate); i += 2 {
			gb.AddCol(nil, columnTemplate[i], columnTemplate[i+1])
		}
	}

	// gb.dumpCols()

	// adding rows
	for rowIdx := 0; rowIdx < len(inputNames); rowIdx++ {
		name := inputNames[rowIdx]

		sparseLbls := map[int]trl.S{}
		if rowIdx < len(firstColLabels) {
			lbl := firstColLabels[rowIdx]
			sparseLbls[0] = lbl
		}

		// sparseLbls[3] = trl.S{"de": "--", "en": "--"}
		gb.AddRadioRow(name, radioVals, sparseLbls)
	}

	return gb

}

// -----------------------------------------------------------------

// AddGrid statically adds inputs to the page;
// these inputs are based on the preconfigured
// GridBuilder's cells
func (p *pageT) AddGrid(gb *GridBuilder) *groupT {

	gr := p.AddGroup()

	gr.Cols = float32(len(gb.cols))

	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleBox.Display = "grid"           // prevent defaults being set while rendering
	gr.Style.Desktop.StyleGridContainer.AutoFlow = "row" // prevent defaults being set while rendering
	stl := ""
	for colIdx := 0; colIdx < len(gb.cols); colIdx++ {
		stl = fmt.Sprintf(
			"%v   %vfr ",
			stl,
			gb.cols[colIdx].spanLabel+gb.cols[colIdx].spanControl,
		)
	}
	if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = stl
		// log.Printf("GridBuilder.AddGrid() - %v", stl)
	} else {
		log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
	}

	// gb.dumpCols()
	// log.Printf("gr.Cols %v  ", gr.Cols)
	// log.Printf("gr.Style.GridContainerStyle %v  ", util.IndentedDump(gr.Style.Desktop.GridContainerStyle))

	// first row - main label
	if gb.MainLabel != nil {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = gb.MainLabel
		inp.ColSpan = gr.Cols
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

			inp := gr.addEmptyTextblock()

			if gb.cols[colIdx].header == nil {
				// inp.Label = trl.S{"de": "---", "en": "---"}
				inp.ColSpanLabel = gb.cols[colIdx].spanLabel + gb.cols[colIdx].spanControl
			} else {
				inp.Type = "label-as-input"
				inp.Label = gb.cols[colIdx].header
				inp.ColSpanLabel = gb.cols[colIdx].spanLabel
				inp.ColSpanControl = gb.cols[colIdx].spanControl
				inp.StyleCtl = css.ItemCenteredCA(inp.StyleCtl)
				inp.StyleCtl = css.ItemStartCA(inp.StyleCtl)

				inp.Style = css.NewStylesResponsive(inp.Style)
				// inp.Style.Desktop.GridContainerStyle.GapColumn is distorting cell widths..
				inp.Style.Mobile.StyleBox.Padding = "0 0.1rem"
				inp.Style.Mobile.StyleText.FontSize = 90
			}

			// log.Printf("colIdx %v  -  lbl/ctl  %v  %v ", colIdx, inp.ColSpanLabel, inp.ColSpanControl)

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
