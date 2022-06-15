package biii

import (
	"fmt"
	"math"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func later(q *qst.QuestionnaireT) {

	//
	// branch "not now"
	//
	//
	// page 2b-01 II no or later
	{
		page := q.AddPage()
		page.Short = trl.S{"de": "Zukunftsposition"}
		page.Label = trl.S{"de": ""}
		page.NavigationCondition = "BIIILater"
		page.WidthMax("42rem")

		// gr0
		{
			labels := []trl.S{
				{"de": "Wir bereiten einen Markteintritt strategisch vor"},
				{"de": "Wir prüfen künftige Marktchancen sorgfältig "},
				{"de": "Wir beginnen, Informationen zu diesem Markt zu sammeln"},
				{"de": "Wir sehen keine Notwendigkeit, uns damit zu befassen "},
				{"de": "Wir sehen diesen Markt eher skeptisch aus den folgenden Gründen: "},
			}
			radioValues := []string{
				"strategic_prep",
				"due_dilligence",
				"collect_info",
				"no_necessity",
				"sceptical",
			}
			gr := page.AddGroup()
			gr.Cols = 7
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>42.</b> &nbsp;	Falls Ihre Organisation bisher noch nicht im Impact Investing-Markt tätig ist, was planen Sie für die Zukunft? "}
				// (bitte nur eine Auswahl)
				inp.ColSpan = gr.Cols
			}

			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q42"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				// all rows except last
				if idx < len(labels)-1 {
					rad.ColSpan = gr.Cols
					rad.Label = label
					rad.ControlFirst()
				} else {
					// last row: now label
					rad.ColSpan = 1
					rad.ColSpanLabel = 0 // value 0 prevents the label from taking any place
					rad.ColSpanControl = 1

					inp := gr.AddInput()
					inp.Type = "textarea"
					inp.Name = "q42_other"
					inp.MaxChars = 150
					inp.Label = label

					inp.ColSpan = gr.Cols - 1
					inp.ColSpanLabel = 4
					inp.ColSpanControl = 5
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				}

			}

		}

		// gr1
		{
			labels := []trl.S{
				{"de": "N/A (Nicht anwendbar)"},
				{"de": "Für welches Jahr/Monat?"},
				{"de": "Für welches Marktsegment (Eigenkapital, Fremdkapital, strukturierte Finanzierung)? "},
				{"de": "Mit welchem geplanten Volumen? "},
			}
			subNames := []string{
				"not_applicable",
				"year_month",
				"segment",
				"volume",
			}
			gr := page.AddGroup()
			gr.Cols = 5
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": `
					<b>43.</b> &nbsp;	
					Falls Sie einen Markteintritt strategisch vorbereiten… 
				`}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q43_%v", subNames[idx])
				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6
				rad.Label = label
				rad.Style = css.NewStylesResponsive(rad.Style)
				rad.ControlFirst()

				rad.ColSpan = 3
				if idx > 0 {
					rad.ColSpan = 3
				}

				if idx == 1 || idx == 2 {
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = fmt.Sprintf("q43_%v_detail", subNames[idx])
					// inp.Label = label
					inp.ColSpan = 2
					inp.ColSpanControl = 7
					inp.MaxChars = 20

					if idx == 1 {
						inp.MaxChars = 8
						inp.Placeholder = trl.S{"de": "00/20xx"}
					}
					// inp.Suffix = trl.S{"de": "%"}

					inp.LabelPadRight()

					inp.Style = css.NewStylesResponsive(inp.Style)
					inp.Style.Desktop.StyleBox.Position = "relative"
					inp.Style.Desktop.StyleBox.Left = "-1rem"

				}

				if idx == 3 {
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("q43_%v_detail", subNames[idx])

					inp.ColSpan = 2
					inp.ColSpanControl = 7
					inp.Min = 0
					inp.Max = math.MaxFloat64
					inp.Step = 0.01
					inp.MaxChars = 9
					inp.Suffix = trl.S{"de": "Mio €"}
					inp.Placeholder = trl.S{"de": "0,00"}

					inp.LabelPadRight()
					// inp.Style = css.NewStylesResponsive(inp.Style)
					// inp.Style.Desktop.StyleBox.Margin = "0 0 0 2.5rem"
					inp.Style = css.NewStylesResponsive(inp.Style)
					inp.Style.Desktop.StyleBox.Position = "relative"
					inp.Style.Desktop.StyleBox.Left = "-1rem"
				}
			}
		}

		// gr2
		{
			tmp1 := q12inputNames[:len(q12inputNames)-2]
			tmp2 := q12Labels[:len(q12inputNames)-2]
			inpNames := make([]string, len(tmp1))
			copy(inpNames, tmp1)
			for i := 0; i < len(inpNames); i++ {
				inpNames[i] = strings.ReplaceAll(inpNames[i], "q12", "q43a")
			}

			gr := page.AddGroup()
			gr.Cols = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": "<b>43a.</b> &nbsp;	Für welche Themen?"}
				inp.Label = trl.S{"de": "Für welche Themen?"}
				inp.ColSpan = gr.Cols
			}
			for idx, label := range tmp2 {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("%v", inpNames[idx])

				rad.ColSpan = 1
				rad.ColSpanLabel = 2
				rad.ColSpanControl = 10

				rad.Label = label

				rad.ControlFirst()
			}
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q43_other"
				inp.MaxChars = 20
				inp.Label = trl.S{"de": "Andere, bitte nennen"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleBox.Padding = "0 0 0 3.4rem"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.StyleBox.Margin = "0.35rem 0 0 0"
			}
		}

	}

}
