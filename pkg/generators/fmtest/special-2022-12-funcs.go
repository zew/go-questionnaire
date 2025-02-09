package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// matrixOfPercentageInputs for hierarchical checkboxes from page1
func matrixOfPercentageInputs(
	page *qst.WrappedPageT,
	lblMain trl.S,
	lblsCols []trl.S,
	inpNames []string,
	lblsRows []trl.S,
) {

	const col1Width = 4
	const col23Width = 2
	const col4Width = 1

	gr := page.AddGroup()
	gr.Cols = col1Width + 2*col23Width + col4Width
	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleGridContainer.GapColumn = "0.7rem"

	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = lblMain
	}

	// header row with column labels
	for i1 := 0; i1 < len(lblsCols); i1++ {
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = col23Width
			if i1 == 0 {
				inp.ColSpan = col1Width
			}
			if i1 == 3 {
				inp.ColSpan = col4Width
			}
			inp.Label = lblsCols[i1]
			inp.LabelCenter()
			inp.LabelBottom()
		}
	}

	for i1 := 0; i1 < len(inpNames); i1++ {

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = col1Width
			inp.Label = lblsRows[i1]
		}

		for _, suffix := range []string{"lb", "ub"} {

			inp := gr.AddInput()
			inp.Type = "number"
			// inp.Name = fmt.Sprintf("inf%v_%v", inpNames[i1], suffix)
			inp.Name = fmt.Sprintf("%v_%v", inpNames[i1], suffix)
			inp.Suffix = trl.S{"de": "%", "en": "%"}
			inp.ColSpan = col23Width
			inp.Min = -40
			inp.Max = 40
			inp.Step = 0.01
			// different steps for growth...
			inp.MaxChars = 5
			inp.ControlCenter()
			// inp.ControlTop()
		}
		{

			inp := gr.AddInput()
			inp.Type = "checkbox"
			// inp.Name = fmt.Sprintf("inf%v_%v", inpNames[i1], "no_answer")
			inp.Name = fmt.Sprintf("%v_%v", inpNames[i1], "no_answer")
			inp.ColSpan = col4Width
			inp.ControlTopNudge()
		}

	}

}
