package fmt

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202401(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 1
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "Sonderfrage zu Schuldenbremse und Investitionen",
		"en": "Special questions: Debt brake and investments",
	}
	page.Short = trl.S{
		"de": "Sonderfrage<br>Schuldenbremse",
		"en": "Special questions:<br>Debt brake",
	}
	page.WidthMax("48rem")

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
				Seit 2009 hat Deutschland die so genannte <i>„Schuldenbremse“</i> im Grundgesetz. 
				Sie erlaubt dem Bund eine maximale jährliche Neuverschuldung 
				im Umfang von 3,5&nbsp;Prozent des BIP. 
				
				Die Länder müssen seit 2020 stets einen ausgeglichenen Haushalt vorlegen. 
				Allerdings bleiben Defizite im Falle einer Rezession oder außergewöhnlicher Ereignisse erlaubt. 
				Aktuell wird eine Reform der Schuldenbremse diskutiert. 
				Ein Argument der Reform-Befürworter ist, 
				dass die Schuldenbremse in ihrer heutigen Gestalt ein Hindernis für Investitionen darstelle. 
				
				Wie ist Ihre Einschätzung?  
				`,
				"en": `  
				Since 2009, Germany has added the so-called <i>„Schuldenbremse“</i> (debt brake) 
				to its constitution. 
				It allows the government to accumulate additional annual debt of at most 3.5&nbsp;percent of GDP. 
				Since 2020, the federal states are required to always maintain a balanced budget. 
				However, deficits are allowed in case of a recession or other extraordinary events. 
				A reformed version of the debt brake is currently being discussed. 
				An argument by the proponents of the reform is that the current version 
				of the debt brake might be an obstacle to investment. What is your assessment?
				
				`,
			}
			inp.ColSpanLabel = 1
			inp.ColSpan = gr.Cols
		}
	}

	//
	// gr1
	{
		colTpl := []float32{
			4, 1,
			0, 1,
		}
		gb := qst.NewGridBuilderRadios(
			colTpl,
			[]trl.S{
				{
					"de": "stimme zu",
					"en": "I agree",
				},
				{
					"de": "stimme nicht zu",
					"en": "I do not agree",
				},
			},
			[]string{"sq3_unchanged", "sq3_amended", "sq3_dropped"},
			[]string{"1", "2"},
			[]trl.S{
				{
					"de": `…unverändert bleiben`,
					"en": `…left unchanged`,
				},
				{
					"de": `…ergänzt werden, so dass höhere Defizite für öffentliche Investitionen erlaubt sind`,
					"en": `…augmented so that higher deficits are allowed for public investments`,
				},
				{
					"de": `…ganz entfallen`,
					"en": `…dropped entirely`,
				},
			},
		)
		gb.MainLabel = trl.S{
			"de": `
					Die Schuldenbremse sollte…
				`,
			"en": `  
					The debt brake should be…
				`,
		}.Outline("3.")
		gr := page.AddGrid(gb)
		_ = gr
		// gr.WidthMax("40rem")
	}

	//
	// gr2
	gr := page.AddGroup()
	{
		lblQ4 := trl.S{
			"de": ` 
				Was sind Ihrer Ansicht nach aktuell die wichtigsten Hindernisse 
				für höhere öffentliche Investitionen (Mehrfachnennungen möglich)?
			`,
			"en": `
				In your opinion, what are the biggest obstacles 
				to higher public investments (multiple choice)?
			`,
		}.Outline("4.")

		lblsQ4 := []trl.S{
			{
				"de": "Die Schuldenbremse",
				"en": "The debt brake",
			},
			{
				"de": "Geringe Steuereinnahmen",
				"en": "Insufficient tax revenues",
			},
			{
				"de": "Hohe nicht-investive Staatsausgaben (z.B. Sozialleistungen, Personalausgaben)",
				"en": "High public spending that is not related to investments (e.g., social security, personnel expenses)",
			},
			{
				"de": "Der Widerstand von betroffenen Bürgerinnen und Bürgern gegen Investitionsvorhaben",
				"en": "The resistance of affected citizens against investment projects",
			},
			{
				"de": "Eine schwerfällige Bürokratie und zu lange Genehmigungsverfahren",
				"en": "A cumbersome bureaucracy and long approval processes",
			},
			{
				"de": "Die Kapazitätsengpässe in der Bauwirtschaft",
				"en": "Capacity bottlenecks in the construction sector",
			},
			{
				"de": "Andere Gründe:",
				"en": "Other reasons:",
			},
		}

		inps := []string{
			"sq4_debt_brake",
			"sq4_tax_revenue",
			"sq4_govt_consumpt",
			"sq4_resistance",
			"sq4_bureaucracy",
			"sq4_bottlenecks",
			"sq4_other",
		}

		col1 := float32(4)
		col2 := float32(1)

		gr.Cols = col1 + col2
		gr.BottomVSpacers = 3
		// gr.Style = css.NewStylesResponsive(gr.Style)
		// gr.Style.Mobile.StyleGridContainer.TemplateColumns = "5fr 1fr 1fr 0.4fr 0.4fr 0.4fr"

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblQ4
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
		}

		for i, inpName := range inps {
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = inpName

				inp.ColSpan = col1
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1

				inp.Label = lblsQ4[i]

				// inp.LabelRight()
				// inp.LabelPadRight()
				// inp.Suffix = trl.S{"de": "%", "en": "%"}
			}

			if i < len(inps)-1 {
				// empty column
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": " &nbsp; ", "en": " &nbsp; "}
				inp.ColSpan = col2
				inp.ColSpanLabel = 1
			} else {
				// free input
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "sq4_other_free"
				inp.MaxChars = 24
				inp.ColSpan = col2
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}
	}

	return nil

}
