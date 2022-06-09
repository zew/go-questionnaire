package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

// AddBiiiPrio question block
func (p *pageT) AddBiiiPrio(
	mainLbl trl.S,
	labels []trl.S,
	names []string,
	idxOther map[int]bool,
	showFreeInput int,
) *groupT {

	colSpace := float32(0)
	if showFreeInput > 0 {
		colSpace = 2
	}

	gr := p.AddGroup()
	gr.Cols = 7 + colSpace
	// gr.WidthMax("17rem")

	if !mainLbl.Empty() {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = mainLbl
	}

	for idx, label := range labels {

		if _, ok := idxOther[idx]; !ok {

			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = fmt.Sprintf("%v", names[idx])
			inp.Label = label

			inp.ColSpan = gr.Cols - colSpace
			inp.ColSpanLabel = 6
			inp.ColSpanControl = 1

		} else {

			inpOth := gr.AddInput()
			inpOth.Type = "text"
			inpOth.Name = fmt.Sprintf("%v_label", names[idx])
			inpOth.Label = label
			inpOth.MaxChars = 17

			inpOth.ColSpan = gr.Cols - colSpace - 1
			inpOth.ColSpanLabel = 3
			inpOth.ColSpanControl = 3

			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = fmt.Sprintf("%v", names[idx])
			// inp.Label = label // no label

			inp.ColSpan = 1
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1

		}

		if showFreeInput == 1 {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_addl", names[idx])
			inp.MaxChars = 6
			inp.Placeholder = trl.S{"de": "0,0"}
			inp.Suffix = trl.S{"de": "Mio. €"}

			inp.Min = 0
			inp.Max = 1000 * 1000
			inp.Step = 0.1

			inp.ColSpan = colSpace
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1
		}

		if showFreeInput == 2 {
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = fmt.Sprintf("%v_addl", names[idx])
			inp.MaxChars = 16
			inp.Placeholder = trl.S{"de": "opt. Kommentar"}

			inp.ColSpan = colSpace
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1
		}

		// inp.ControlFirst()
	}

	return gr

}

//
//
// AddBiiiPrio question block
func (p *pageT) AddBiiiPrio2Cols(
	mainLbl trl.S,
	labels []trl.S,
	names []string,
	idxOther map[int]bool,
) *groupT {

	colsWithLabel := float32(7)
	colsCheckbox := float32(1)
	colsComment := float32(3)

	gr := p.AddGroup()
	gr.Cols = colsWithLabel + colsCheckbox + 2*colsComment
	// gr.WidthMax("17rem")

	if !mainLbl.Empty() {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = mainLbl
	}

	for idx, label := range labels {

		for colIdx := 0; colIdx < 2; colIdx++ {

			nameCol := ""
			if colIdx == 0 {
				nameCol = names[idx] + "_need"
			}
			if colIdx == 1 {
				nameCol = names[idx] + "_pot"
			}

			if _, ok := idxOther[idx]; !ok {

				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("%v", nameCol)

				if colIdx == 0 {
					inp.Label = label

					inp.ColSpan = colsWithLabel
					inp.ColSpanLabel = 6
					inp.ColSpanControl = 1
				} else {
					inp.ColSpan = colsCheckbox
					inp.ColSpanLabel = 0
					inp.ColSpanControl = 1
				}

			} else {

				if colIdx == 0 {
					inpOth := gr.AddInput()
					inpOth.Type = "text"
					inpOth.Name = fmt.Sprintf("%v_label", nameCol)
					inpOth.MaxChars = 17

					inpOth.Label = label
					inpOth.ColSpan = colsWithLabel - colsCheckbox
					inpOth.ColSpanLabel = 3
					inpOth.ColSpanControl = 3
				}

				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("%v", nameCol)
				// inp.Label = label // no label

				inp.ColSpan = colsCheckbox
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = fmt.Sprintf("%v_addl", nameCol)
				inp.MaxChars = 14
				inp.Placeholder = trl.S{"de": "opt. Kommentar"}

				inp.ColSpan = colsComment
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}

			// inp.ControlFirst()
		}
	}

	return gr

}
