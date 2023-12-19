package qst

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func fmt202312(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"de": "Sonderfrage zu CO2-Preiserwartungen",
		"en": "Special questions: CO2 price expectations",
	}
	page.Short = trl.S{
		"de": "Sonderfrage<br>CO2-Preiserwartungen",
		"en": "Special questions:<br>CO2 price",
	}
	page.WidthMax("48rem")

	//
	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.WidthMax("40rem")
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "co2_2030"
			// inp.Placeholder = ph
			inp.Min = 0
			inp.Max = 50000
			inp.Step = 0
			inp.MaxChars = 6

			inp.ColSpan = 6
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 1
			inp.Label = trl.S{
				"de": `
						Was erwarten Sie, was der durchschnittliche 
						<b>CO2-Preis</b> 
						pro Tonne im Rahmen des Emissionshandelssystems 
						der Europäischen Union (EU-EHS) im Jahr 
						<b>2030</b> 
						sein wird?
					`,
				"en": `
						What do you expect the average 
						<b>CO2 price</b> 
						per ton 
						within the European Union Emissions Trading Scheme (EU-ETS) to be in 
						<b>2030</b>?
					`,
			}.Outline("3.")
			inp.Suffix = trl.S{"de": "€", "en": "€"}
			inp.LabelPadRight()
			inp.ControlTop()
			inp.ControlBottom()
		}
	}

	//
	//
	{
		gb := NewGridBuilderRadios(
			[]float32{
				0, 1,
				0, 1,
				0, 1,
				0, 1,
			},
			[]trl.S{
				{
					"de": "sehr sicher",
					"en": "very certain",
				},
				{
					"de": "sicher",
					"en": "certain",
				},
				{
					"de": "unsicher",
					"en": "uncertain",
				},
				{
					"de": "sehr unsicher",
					"en": "very uncertain",
				},
			},
			[]string{"co2_confidence"},
			[]string{"1", "2", "3", "4"},
			nil,
		)
		gb.MainLabel = trl.S{
			"de": "Wie sicher sind Sie sich bei dieser Preiserwartung?",
			"en": "How certain are you about this price expectation? ",
		}.Outline("4.")
		gr := page.AddGrid(gb)
		gr.WidthMax("30rem")
	}

	//
	//
	//
	grIdx := q.UserIDInt() % 2

	lblsGr1 := []trl.S{
		{
			"de": "<75 €",
			"en": "<75 €",
		},
		{
			"de": "75 - 84 €",
			"en": "75 - 84 €",
		},
		{
			"de": "85 - 94 €",
			"en": "85 - 94 €",
		},
		{
			"de": "95 - 104 €",
			"en": "95 - 104 €",
		},
		{
			"de": "105 - 114 €",
			"en": "105 - 114 €",
		},
		{
			"de": "115 - 124 €",
			"en": "115 - 124 €",
		},
		{
			"de": "≥ 125 €",
			"en": "≥ 125 €",
		},
	}
	lblsGr2 := []trl.S{
		{
			"de": "<70 €",
			"en": "<70 €",
		},
		{
			"de": "70 - 89 €",
			"en": "70 - 89 €",
		},
		{
			"de": "90 - 109 €",
			"en": "90 - 109 €",
		},
		{
			"de": "110 - 129 €",
			"en": "110 - 129 €",
		},
		{
			"de": "130 - 149 €",
			"en": "130 - 149 €",
		},
		{
			"de": "150 - 169 €",
			"en": "150 - 169 €",
		},
		{
			"de": "≥ 170 €",
			"en": "≥ 170 €",
		},
	}

	lbls := [][]trl.S{
		lblsGr1,
		lblsGr2,
	}

	inps := []string{
		"co2_brack_1",
		"co2_brack_2",
		"co2_brack_3",
		"co2_brack_4",
		"co2_brack_5",
		"co2_brack_6",
		"co2_brack_7",
	}

	lbl := trl.S{
		"de": fmt.Sprintf(` 
			Wir sind an Ihrer Einschätzung der Wahrscheinlichkeit interessiert, 
			  dass der Preis pro Tonne CO2 im Rahmen des Emissionshandelssystems der Europäischen Union (EU-EHS) 
			  im Jahr 2030 in eine bestimmte Preisspanne fällt.
			<br>
			<br>
			Bitte geben Sie für jede der folgenden Preisspannen die Wahrscheinlichkeit an, 
			dass der durchschnittliche CO2-Preis 
			 im Jahr 2030 in dieser Spanne liegen wird:
			<br>
			<br>
			<span style=font-size:90%%>Bitte beachten Sie, dass alle Wahrscheinlichkeiten in der Summe 100%% ergeben müssen.</span>
		`,
		),
		"en": fmt.Sprintf(`
		  We are interested in your assessment of the likelihood that the CO2 price per ton 
		    within the European Union Emissions Trading Scheme (EU-ETS) will fall into certain price ranges.
			<br>
			<br>
			For each of the following price ranges, please indicate the probability that the average CO2 price in 2030 will be in that range:
			<br>
			<br>
			<span style=font-size:90%%>Please keep in mind that all probabilities have to sum up to 100%%.</span>
		`),
	}.Outline("5.")

	// gr := page.AddGrid(gb)
	gr := page.AddGroup()

	col1 := float32(2)
	col2 := float32(1)
	col3 := float32(3)

	gr.Cols = col1 + col2 + col3
	gr.BottomVSpacers = 2

	gr.Style = css.NewStylesResponsive(gr.Style)
	// gr.Style.Mobile.StyleGridContainer.TemplateColumns = "1fr 1fr 1fr 1fr 1fr 1fr"
	//
	gr.Style.Mobile.StyleGridContainer.TemplateColumns = "5fr 1fr 1fr 0.4fr 0.4fr 0.4fr"

	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = lbl
		inp.ColSpan = gr.Cols
		inp.ColSpanLabel = 1
	}

	for i, inpName := range inps {
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Min = 0
			inp.Max = 100
			inp.Step = 0.1
			inp.Step = 1
			inp.Name = inpName
			inp.MaxChars = 6

			inp.ColSpan = col1 + col2
			inp.ColSpanLabel = col1
			inp.ColSpanControl = 2

			inp.Label = lbls[grIdx][i]
			// inp.Label.AppendStr(" &nbsp; ")

			inp.LabelRight()
			inp.LabelPadRight()

			inp.Suffix = trl.S{"de": "%", "en": "%"}
		}

		{
			// empty column
			inp := gr.AddInput()
			inp.Type = "textblock"
			// inp.Label = trl.S{"de": " &nbsp; ", "en": " &nbsp; "}
			inp.ColSpan = col3
			inp.ColSpanLabel = 1
		}
	}

	{
		// divSum := `&#931; <span id='sum'  style='display:inline-bock; margin-left: 35%; background-color: #ddd; padding: 0.2rem 0.4rem;' > &nbsp; </span>€ `
		{
			inp := gr.AddInput()
			inp.Name = "sum"
			inp.Type = "number"
			inp.Min = -200
			inp.Max = 200
			inp.Step = 1

			inp.MaxChars = 6

			inp.ColSpan = col1 + col2
			inp.ColSpanLabel = col1
			// inp.ColSpanControl = col2
			inp.ColSpanControl = 2

			inp.Label = trl.S{"de": "&#931;", "en": "&#931;"}
			inp.Label = trl.S{
				"de": `Summe der obigen Wahrscheinlichkeiten`,
				"en": `Sum of the above probabilities`,
			}
			// inp.Label.AppendStr(" &nbsp; ")
			inp.LabelRight()
			inp.LabelPadRight()
			// inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
			// inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "start"
			// inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "start"

			inp.Suffix = trl.S{"de": "%", "en": "%"}

			inp.Disabled = true

		}
		{
			// empty column
			inp := gr.AddInput()
			inp.Type = "textblock"
			// inp.Label = trl.S{"de": " &nbsp; ", "en": " &nbsp; "}
			inp.ColSpan = col3
			inp.ColSpanLabel = 1
		}
	}
	//
	{
		inp := gr.AddInput()
		inp.ColSpanControl = 1
		inp.Type = "javascript-block"
		inp.Name = "co2"

		s1 := trl.S{
			"de": "Über hundert Prozent.",
			"en": "Over hundred percent.",
		}
		// s2 := trl.S{
		// 	"de": "Nicht hundert Prozent.",
		// 	"en": "Not hundred percent.",
		// }
		s2 := trl.S{
			"de": "Die Wahrscheinlichkeiten summieren sich nicht zu 100%. Bitte korrigieren Sie Ihre Angaben.",
			"en": "The sum of probabilities is not 100%. Please correct your responses.",
		}
		inp.JSBlockTrls = map[string]trl.S{
			"msg1": s1,
			"msg2": s2,
		}

		inp.JSBlockStrings = map[string]string{}

		ivls := []string{} // intervals
		for _, name := range inps {
			ivl := fmt.Sprintf("\"%v\"", name)
			ivls = append(ivls, ivl)
		}

		inp.JSBlockStrings["inps"] = "[" + strings.Join(ivls, ", ") + "]"

	}

	return nil
}
