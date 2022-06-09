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
	// idxOther int,
	idxOther map[int]bool,
	showVolumeInput int,
) *groupT {

	colSpace := float32(0)
	if showVolumeInput > 0 {
		colSpace = 2
	}

	if len(labels) == 0 || len(names) == 0 {
		labels = []trl.S{
			{"de": "Gewinn"},
			{"de": "Moral"},
			{"de": "Ethik"},
			{"de": "Umsatz"},
			{"de": "Partnerschaft"},
			{"de": "Klima"},
			{"de": "Waldschutz"},
			{"de": "Tierschutz"},
		}
		names = []string{
			"1",
			"2",
			"3",
			"4",
			"5",
			"6",
			"7",
			"8",
		}
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
			inpOth.ColSpanLabel = 1
			inpOth.ColSpanControl = 3

			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = fmt.Sprintf("%v", names[idx])
			// inp.Label = label // no label

			inp.ColSpan = 1
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1

		}

		if showVolumeInput == 1 {
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

		if showVolumeInput == 2 {
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
