package biii

import (
	"fmt"
	"math"

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
		page.Short = trl.S{"de": "II not or later"}
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
			}
			radioValues := []string{
				"strategic_prep",
				"due_dilligence",
				"collect_info",
				"no_necessity",
			}
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "<b>42.</b> &nbsp;	Falls Ihre Organisation bisher noch nicht im Impact Investing-Markt tätig ist, wie bewerten Sie dieses Feld für die Zukunft: "}
				// (bitte nur eine Auswahl)
				inp.ColSpan = gr.Cols
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "q42"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = 1
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label

				rad.ControlFirst()
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "q42_other"
				inp.MaxChars = 150
				inp.Label = trl.S{"de": "Wir sehen diesen Markt eher skeptisch aus den folgenden Gründen:"}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
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
						inp.Placeholder = trl.S{"de": "20xx"}
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
					inp.Step = 1000
					inp.Step = 1
					inp.MaxChars = 15
					inp.Suffix = trl.S{"de": "€"}
					inp.Placeholder = trl.S{"de": "0.000.000"}

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
			labels := []trl.S{
				{"de": "Bildung"},
				{"de": "Ernährung"},
				{"de": "Gesundheit"},
				{"de": "Armutsbekämpfung"},

				{"de": "Soziale / Geschlechter-Gerechtigkeit"},
				{"de": "Demografischer Wandel / Altenpflege"},
				{"de": "Kinder- und Jugendhilfe"},
				{"de": "Migration und Integration"},

				{"de": "(Sozialer) Wohnungsbau"},
				{"de": "Saubere Energie"},
				{"de": "Kunst und Kultur"},
				{"de": "Wasser, Sanitärversorgung und Hygiene"},

				{"de": "Finanzielle Inklusion"},
				{"de": "Nachhaltiges Wirtschaften"},
				{"de": "Arbeitsmarktintegration"},
				{"de": "Digitalisierung"},

				{"de": "Mobilität"},
				{"de": "Klimaschutz"},
				{"de": "Schutz von Ökosystemen und Biodiversität"},
			}
			inpNames := []string{
				"q43a_education",
				"q43a_nutrition",
				"q43a_health",
				"q43a_poverty",

				"q43a_justice",
				"q43a_elderly_care",
				"q43a_youth",
				"q43a_migration",

				"q43a_residential",
				"q43a_clean_energy",
				"q43a_culture",
				"q43a_sanitation",

				"q43a_inclusion",
				"q43a_sustainability",
				"q43a_labor_market",
				"q43a_digital",

				"q43a_mobility",
				"q43a_climate",
				"q43a_biodiversity",
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
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "checkbox"
				rad.Name = fmt.Sprintf("q43a_%v", inpNames[idx])

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
