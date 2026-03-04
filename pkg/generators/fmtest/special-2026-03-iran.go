package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202603IranSsq3(
	page *qst.WrappedPageT,
	colLabels []trl.S,
	inputStem string,
	rowLbls []trl.S,
) {

	firstCol := float32(4)

	//
	gr := page.AddGroup()
	gr.Cols = firstCol + 1 + 1

	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
	gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

	{

		// {
		// 	inp := gr.AddInput()
		// 	inp.Type = "textblock"
		// 	inp.ColSpan = firstCol
		// 	inp.ColSpanLabel = 1
		// 	inp.ColSpanControl = 0
		// 	inp.Label = trl.S{
		// 		"de": "&nbsp;",
		// 		"en": "&nbsp;",
		// 	}
		// }
		for i := 0; i < len(colLabels); i++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if i == 0 {
				inp.ColSpan = firstCol
			}
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = colLabels[i]
			inp.LabelCenter()
			inp.LabelBottom()
		}
	}

	//
	//
	for i1 := 0; i1 < len(rowLbls); i1++ {

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = firstCol
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = rowLbls[i1]

		}

		for i2 := 0; i2 < 2; i2++ {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v_%v", inputStem, i1+1, i2+1)
			inp.Min = -10
			inp.Max = 40
			inp.Step = 0.1
			inp.MaxChars = 4
			inp.Suffix = trl.S{
				"de": "%",
				"en": "%",
			}
			inp.Placeholder = trl.S{
				"de": "0.0",
				"en": "0.0",
			}

			inp.ColSpan = 1
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1

			inp.ControlCenter()
		}
	}

}

func special202603IranSsq5(page *qst.WrappedPageT, inputStem string, rowLbls []trl.S) {

	for i := 0; i < len(rowLbls); i++ {

		gr := page.AddGroup()
		gr.Cols = 12
		gr.RandomizationGroup = 0 //  change
		gr.BottomVSpacers = 1
		if i == len(rowLbls)-1 {
			gr.RandomizationGroup = 0
			gr.BottomVSpacers = 3
		}

		{

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("%v_%v", inputStem, i+1)

			if i < 6 {

				inp1.ColSpan = gr.Cols
				inp1.ColSpanLabel = 1
				inp1.ColSpanControl = 11
				inp1.Label = rowLbls[i]

				inp1.ControlFirst()

			} else {

				inp1.ColSpan = 1
				inp1.ColSpanLabel = 0
				inp1.ColSpanControl = 1
				// inp1.Label = nil

				//
				inp2 := gr.AddInput()
				inp2.Label = rowLbls[i]
				inp2.Type = "text"
				inp2.Name = inputStem + "_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 1
				inp2.ColSpanControl = 6

			}

		}

	}
}

// main
// func
func special202603Iran(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2026 && q.Survey.Month == 3
	if !cond {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("72rem")
	page.WidthMax("64rem")
	page.WidthMax("48rem")

	page.Label = trl.S{
		"de": "Sonderfragen zur Eskalation im Nahost-Konflikt",
		"en": "Special questions: Near East conflict",
	}
	page.Short = trl.S{
		"de": "Nahost",
		"en": "Near<br>East",
	}
	// page.WidthMax("42rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Wie wichtig sind folgende Wirkungskanäle für die Veränderung Ihrer Wachstumsprognose für das Jahr 2026 infolge der Eskalation im Nahen Osten?
				`,
				"en": `
					todo
				`,
			}.Outline("3.")
		}
	}

	colLabelsSsq3 := []trl.S{
		{
			"de": "Szenario",
			"en": "Scenario",
		},
		{
			"de": "reale BIP Wachstumsrate 2026",
			"en": "real GDP growth rate 2026",
		},
		{
			"de": "Inflationsrate 2026",
			"en": "inflation rate 2026",
		},
	}

	lblsSsq3 := []trl.S{
		{
			"de": `<b> a) </b>   Kurzfristiger und begrenzter militärischer 
				Konflikt ohne wesentliche Beeinträchtigung 
				der Energieversorgung
			`,
			"en": `todo`,
		},
		{
			"de": `<b> b) </b> Anhaltender militärischer Konflikt mit
			spürbarer Belastung der Energieversorgung
			`,
			"en": `todo`,
		},
		{
			"de": `<b> c) </b> Länger anhaltender militärischer Konflikt
				mit erheblichen Versorgungsstörungen
			`,
			"en": `todo`,
		},
	}
	special202603IranSsq3(qst.WrapPageT(page), colLabelsSsq3, "ssq3", lblsSsq3)

	//
	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Wie wichtig sind folgende Wirkungskanäle für die Veränderung Ihrer Wachstumsprognose für das Jahr 2026 infolge der Eskalation im Nahen Osten?
				`,
				"en": `
					todo
				`,
			}.Outline("4.")
		}
	}

	colLabelsSsq4 := []trl.S{
		{
			"de": "sehr wichtig",
			"en": "very important",
		},
		{
			"de": "eher   wichtig",
			"en": "rather important",
		},
		{
			"de": "weder/ noch",
			"en": "neither nor",
		},
		{
			"de": "eher   unwichtig",
			"en": "rather unimportant",
		},
		{
			"de": "sehr unwichtig",
			"en": "very unimportant",
		},
		{
			"de": "keine<br>Angabe  ",
			"en": "no answer        ",
		},
	}
	lblsSsq4 := []trl.S{
		{
			"de": `Höhere Energiepreise für Unternehmen`,
			"en": `todo`,
		},
		{
			"de": `Höhere Energiepreise für Haushalte (z. B. Benzin, Heizkosten)`,
			"en": `todo`,
		},
		{
			"de": `Niedrigere globale Nachfrage nach deutschen Exporten`,
			"en": `todo`,
		},
		{
			"de": `Investitionszurückhaltung aufgrund höherer Unsicherheit`,
			"en": `todo`,
		},
		{
			"de": `Störungen von Lieferketten oder Transportwegen`,
			"en": `todo`,
		},
	}
	lblFree := trl.S{
		"de": "Andere",
		"en": "Other",
	}

	special202603ClimateTpB(qst.WrapPageT(page), colLabelsSsq4, "ssq4", lblsSsq4, 0, lblFree)

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Welche wirtschaftspolitischen Reaktionen erwarten Sie bei einem anhaltenden militärischen Konflikt mit spürbarer Belastung der Energieversorgung (1-3 Monate)? Mehrfachauswahl möglich. 
				`,
				"en": `
					todo?
				`,
			}.Outline("5.")
		}
	}

	lblsSsq5 := []trl.S{
		{
			"de": `Keine wesentliche wirtschaftspolitische Reaktion`,
			"en": `todo`,
		},
		{
			"de": `Lockerung der Geldpolitik im Euroraum`,
			"en": `todo`,
		},
		{
			"de": `Restriktivere Geldpolitik im Euroraum`,
			"en": `todo`,
		},
		{
			"de": `Fiskalische Stützungsmaßnahmen (z. B. Entlastungen, Transfers)`,
			"en": `todo`,
		},
		{
			"de": `Maßnahmen zur Stabilisierung der Energieversorgung`,
			"en": `todo`,
		},
		{
			"de": `Ausweitung verteidigungsbezogener Staatsausgaben`,
			"en": `todo`,
		},
		{
			"de": `Sonstige`,
			"en": `todo`,
		},
	}
	special202603IranSsq5(qst.WrapPageT(page), "ssq5", lblsSsq5)

	return nil

}
