package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202603IranSsq1(
	page *qst.WrappedPageT,
	inputStem string,
	rowLbls []trl.S,
) {

	firstCol := float32(5)

	//
	gr := page.AddGroup()
	gr.Cols = firstCol + 1

	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
	gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

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

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Padding = "0 0 0 1.2ch"
			inp.Style.Mobile.StyleBox.Left = "0"
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v", inputStem, i1+1)
			inp.Min = 0
			inp.Max = 100
			inp.Step = 1
			inp.MaxChars = 5
			inp.Suffix = trl.S{
				"de": "%",
				"en": "%",
			}
			// inp.Placeholder = trl.S{
			// 	"de": "0.0",
			// 	"en": "0.0",
			// }

			inp.ColSpan = 1
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1

			inp.ControlCenter()

			// summe
			if i1 == 3 {
				inp.Disabled = true

				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleBox.Disabled = true

			}

		}
	}

	{
		inp := gr.AddInput()
		inp.ColSpanControl = 1
		inp.Type = "javascript-block"
		inp.Name = "inputs-adding-up-100-iran"
		s1 := trl.S{
			"de": "Ihre Antworten  addieren sich nicht zu 100%. Wirklich weiter?",
			"en": "Your answers    dont add up to 100%.         Continue anyway?",
		}
		inp.JSBlockTrls = map[string]trl.S{
			"msg": s1,
		}
		inp.JSBlockStrings = map[string]string{
			"inp1": inputStem + "_1",
			"inp2": inputStem + "_2",
			"inp3": inputStem + "_3",
			"inp4": inputStem + "_4",
		}
	}

}

func special202603IranSsq2(
	page *qst.WrappedPageT,
	colLabels []trl.S,
	inputStem string,
	rowLbls []trl.S,
) {

	firstCol := float32(4)

	//
	gr := page.AddGroup()
	gr.Cols = firstCol + 1 + 1 + 1

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

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Padding = "0 0 0 1.2ch"
			inp.Style.Mobile.StyleBox.Left = "0"
		}

		for i2 := 0; i2 < 3; i2++ {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v_%v", inputStem, i1+1, i2+1)
			inp.Min = 0
			inp.Max = 1000
			inp.Step = 1
			inp.MaxChars = 5
			inp.Suffix = trl.S{
				"de": "USD",
				"en": "USD",
			}
			// inp.Placeholder = trl.S{
			// 	"de": "00",
			// 	"en": "00",
			// }

			inp.ColSpan = 1
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1

			inp.ControlCenter()
			// inp.ControlBottom()
		}
	}

}

func special202603IranSsq3(
	page *qst.WrappedPageT,
	colLabels []trl.S,
	inputStem string,
	rowLbls []trl.S,
) {

	firstCol := float32(4)

	//
	gr := page.AddGroup()
	gr.Cols = firstCol + 1 + 1 + 1

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

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Padding = "0 0 0 1.2ch"
			inp.Style.Mobile.StyleBox.Left = "0"
		}

		// empty second col
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.Label = trl.S{
				"de": "&nbsp;",
				"en": "&nbsp;",
			}
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
			// inp.ControlBottom()
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

	{

		page := q.AddPage()
		// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

		page.WidthMax("72rem")
		page.WidthMax("64rem")
		page.WidthMax("48rem")
		page.WidthMax("52rem")
		page.WidthMax("58rem")

		page.Label = trl.S{
			"de": "Sonderfragen zur Eskalation im Nahost-Konflikt",
			"en": "Special questions: Near East conflict",
		}
		page.Short = trl.S{
			"de": "Nahost",
			"en": "Near<br>East",
		}
		// page.WidthMax("42rem")

		lblsSsq1 := []trl.S{
			{
				"de": `<b> a</b>.)  &nbsp;    			
				 <span style='font-size: 110%'> Kurzfristiger und begrenzter militärischer Konflikt ohne wesentliche Beeinträchtigung der Energieversorgung
				(<&nbsp;1&nbsp;Monat): </span>
				<xxbr>
				Konflikt bleibt zeitlich begrenzt; die globale Energieversorgung wird nicht wesentlich beeinträchtigt. 
			`,
				"en": `todo`,
			},
			{
				"de": `<b> b</b>.)  &nbsp;  
 				 <span style='font-size: 110%'> Anhaltender militärischer Konflikt mit spürbarer Belastung der Energieversorgung
				(1&#8209;3&nbsp;Monate):  </span>
				<xxbr>
				Konflikt hält an; es kommt zu wiederholten oder anhaltenden Beeinträchtigungen der Öl- und Gasinfrastruktur; die globale Energieversorgung ist eingeschränkt, bleibt jedoch funktionsfähig.
			`,
				"en": `todo`,
			},
			{
				"de": `<b> c</b>.)  &nbsp;  
				 <span style='font-size: 110%'> Länger anhaltender militärischer Konflikt mit erheblichen Versorgungsstörungen
				 (>&nbsp;3&nbsp;Monate):  </span>
				 <xxbr>
				 Stark ausgeprägter und länger andauernder Konflikt; zentrale Öl- und Gasinfrastruktur wird massiv gestört; die globale Energieversorgung wird substanziell beeinträchtigt.
			`,
				"en": `todo`,
			},
			{
				"de": `Summe`,
				"en": `Sum`,
			},
		}

		lblsSsq2and3 := []trl.S{
			{
				"de": `<b> a</b>.)  &nbsp;    			
				Kurzfristiger und begrenzter militärischer Konflikt ohne wesentliche Beeinträchtigung der Energieversorgung
			`,
				"en": `todo`,
			},
			{
				"de": `<b> b</b>.)  &nbsp;  
				Anhaltender militärischer Konflikt mit spürbarer Belastung der Energieversorgung
			`,
				"en": `todo`,
			},
			{
				"de": `<b> c</b>.)  &nbsp;  
				Länger anhaltender militärischer Konflikt mit erheblichen Versorgungsstörungen
			`,
				"en": `todo`,
			},
		}

		//
		//
		//
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `
					Für wie wahrscheinlich halten Sie die folgenden Szenarien im Zusammenhang mit der aktuellen Eskalation im Nahen Osten?

					<br>

					<i>Bitte stellen Sie sicher, dass die Summe der Wahrscheinlichkeiten in den Zeilen jeweils 100% ergeben</i>

				`,
					"en": `
					todo
				`,
				}.Outline("1.")
			}
		}
		special202603IranSsq1(qst.WrapPageT(page), "ssq1", lblsSsq1)

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
					Was ist Ihre Prognose für den durchschnittlichen Brent Rohölpreis (USD pro Barrel) für die folgenden Perioden in jedem Szenario?
				`,
					"en": `
					todo
				`,
				}.Outline("2.")
			}
		}

		colLabelsSsq2 := []trl.S{
			{
				"de": "Szenario",
				"en": "Scenario",
			},
			{
				"de": "Q2 2026 &nbsp; &nbsp; &nbsp; ",
				"en": "Q2 2026 &nbsp; &nbsp; &nbsp; ",
			},
			{
				"de": "Q3 2026 &nbsp; &nbsp; &nbsp; ",
				"en": "Q3 2026 &nbsp; &nbsp; &nbsp; ",
			},
			{
				"de": "Q4 2026 &nbsp; &nbsp; &nbsp; ",
				"en": "Q4 2026 &nbsp; &nbsp; &nbsp; ",
			},
		}
		special202603IranSsq2(qst.WrapPageT(page), colLabelsSsq2, "ssq2", lblsSsq2and3)

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
				}.Outline("3.")
			}
		}

		colLabelsSsq3 := []trl.S{
			{
				"de": "Szenario",
				"en": "Scenario",
			},
			// empty second  col
			{
				"de": "&nbsp;",
				"en": "&nbsp;",
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
		special202603IranSsq3(qst.WrapPageT(page), colLabelsSsq3, "ssq3", lblsSsq2and3)

	} // page 1

	//
	//
	//
	{
		page := q.AddPage()
		// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

		page.WidthMax("72rem")
		page.WidthMax("48rem")

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			"de": "Nahost",
			"en": "Near<br>East",
		}
		page.SuppressInProgressbar = true
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

	} // page 2

	return nil

}
