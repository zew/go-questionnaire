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
	idxOther int,
) *groupT {

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
	gr.Cols = 7
	// gr.WidthMax("17rem")

	if !mainLbl.Empty() {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = mainLbl
	}

	for idx, label := range labels {

		if idx != idxOther {

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v", names[idx])
			inp.Label = label
			inp.Min = 1
			inp.Max = 5
			inp.MaxChars = 2

			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 6
			inp.ColSpanControl = 1

		} else {

			inp2 := gr.AddInput()
			inp2.Type = "text"
			inp2.Name = fmt.Sprintf("%v_label", names[idx])
			inp2.Label = label
			inp2.MaxChars = 17

			inp2.ColSpan = 6
			inp2.ColSpanLabel = 1
			inp2.ColSpanControl = 3

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v", names[idx])
			// inp.Label = label // no label
			inp.Min = 1
			inp.Max = 5
			inp.MaxChars = 2

			inp.ColSpan = 1
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1

		}

		// inp.ControlFirst()
	}

	return gr

}
