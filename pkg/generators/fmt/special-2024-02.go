package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202402(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 2
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "Klima- und Umweltrisiken bei der Kreditvergabe",
		"en": "Climate and environmental risks in lending",
	}
	page.Short = trl.S{
		"de": "Sonderfrage<br>Umweltrisiken",
		"en": "Special questions:<br>Climate risks",
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
					Klima- und Umweltrisiken gewinnen zunehmend an Bedeutung für Finanzmärkte. 
					Im Folgenden möchten wir Sie zum Einfluss Klima- und Umweltrisiken auf die Kreditvergabe befragen. 
				`,
				"en": `  
					Climate and environmental risks are becoming increasingly important for financial markets. 
					In the following, we would like to ask you about the influence of climate and environmental risks on lending.			
				`,
			}
			inp.ColSpanLabel = 1
			inp.ColSpan = gr.Cols
		}
	}

	// gr1
	{
		gb := qst.NewGridBuilderRadios(
			[]float32{
				0.2, 1,
				0.0, 1,
				0.0, 1,
				0.4, 1, // no answer slightly apart
			},
			[]trl.S{
				{
					"de": "Kreditgeber",
					"en": "Lender",
				},
				{
					"de": "Kreditnehmer",
					"en": "Borrower",
				},
				{
					"de": "weder noch",
					"en": "Neither nor",
				},
				{
					"de": "keine<br>Angabe",
					"en": "No answer",
				},
			},
			[]string{"sq4"},
			radioVals4,
			[]trl.S{
				{
					"de": "&nbsp;",
					"en": "&nbsp;",
				},
			},
		)
		gb.MainLabel = trl.S{
			"de": `
				Aus welcher Perspektive blicken sie primär auf den deutschen Kreditmarkt?
			`,
			"en": `  
				What is your primary perspective on the German credit market?
			`,
		}.Outline("4.")

		gr := page.AddGrid(gb)
		_ = gr
	}

	//
	// gr2
	colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	ln1 := len(colTemplate)
	colTemplate[ln1-2] = 0.0
	colTemplate[ln1-4] = 0.2

	ln2 := len(colsRowFree)
	colsRowFree[ln2-1] = 0.0
	colsRowFree[ln2-2] = 0.2

	styleRowFree = "3.6fr    1fr    1fr    1fr    1fr    1fr    1.4fr"
	styleRowFree = "3.6fr    1fr    1fr    1fr    1fr    1.2fr    1fr"

	colLbls := []trl.S{
		{
			"de": "0",
			"en": "0",
		},
		{
			"de": "+",
			"en": "+",
		},
		{
			"de": "++",
			"en": "++",
		},
		{
			"de": "+++",
			"en": "+++",
		},
		{
			"de": "nicht zutreffend",
			"en": "not applicable",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}
	lenColLbls := len(colLbls)

	{
		lbls := []trl.S{
			{
				"de": "Aktuelle regulatorische Anforderungen",
				"en": "Current regulatory requirements",
			},
			{
				"de": "Zukünftige regulatorische Anforderungen",
				"en": "Future regulatory requirements",
			},
			{
				"de": "Stakeholder",
				"en": "Stakeholders",
			},
			{
				"de": "Shareholder",
				"en": "Shareholders",
			},
			{
				"de": "Risikopräferenz der Banken",
				"en": "Risk preference of the banks",
			},
			{
				"de": "Intrinsische grüne Präferenz der Banken",
				"en": "Intrinsic green preference of banks",
			},
		}

		gb := qst.NewGridBuilderRadios(
			colTemplate,
			colLbls,
			[]string{
				"qs5_reg_present",
				"qs5_reg_future",
				"qs5_stakeholder",
				"qs5_shareholder",
				"qs5_risk_pref",
				"qs5_green_ref",
			},
			[]string{"1", "2", "3", "4", "5", "6"},
			lbls,
		)

		gb.MainLabel = trl.S{
			"de": fmt.Sprint(` 
				Wie stark beeinflussen, Ihrer Meinung, folgende Faktoren aktuell, 
				ob Klima- und Umweltrisiken bei der Kreditvergabe durch Banken berücksichtigt werden? 
				<small>
				Gar nicht (0), gering (+), mäßig (++), stark (+++)
				</small>
			`,
			),
			"en": fmt.Sprint(`
				In your opinion, to what extent do the following factors currently influence 
				whether climate and environmental risks are considered by banks when granting loans? 
				<small>
				Not at all (0), slightly (+), moderately (++), strongly (+++)
				</small>
			`,
			),
		}.Outline("5.")
		gr := page.AddGrid(gb)
		gr.BottomVSpacers = 1

	}
	{
		//
		// row free input
		gr := page.AddGroup()
		gr.Cols = float32(lenColLbls) + 1
		gr.BottomVSpacers = 4

		gr.Style = css.NewStylesResponsive(gr.Style)
		if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
			gr.Style.Desktop.StyleBox.Display = "grid"
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = styleRowFree
			// log.Printf("fmt special 2021-09: grid template - %v", stl)
		} else {
			return fmt.Errorf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", styleRowFree, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
		}

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "qs5_free_label"
			inp.MaxChars = 15
			inp.ColSpan = 1
			inp.ColSpanLabel = 2.4
			inp.ColSpanControl = 4
			inp.Label = trl.S{
				"de": "Andere",
				"en": "Other",
			}
		}
		for idx := 0; idx < lenColLbls; idx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "qs5_free"
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1
			rad.ColSpanLabel = colsRowFree[2*(idx+1)]
			rad.ColSpanControl = colsRowFree[2*(idx+1)] + 1
		}

	}

	// gr3
	{
		gb := qst.NewGridBuilderRadios(
			[]float32{
				0.2, 1,
				0.0, 1,
				0.0, 1,
				0.4, 1, // no answer slightly apart
			},
			[]trl.S{
				{
					"de": "Ja, über eine separate Risikokomponente",
					"en": "Yes, through a separate risk component",
				},
				{
					"de": "Ja, indirekt über bestehende Risikokomponenten",
					"en": "Yes, indirectly through existing risk components",
				},
				{
					"de": "Nein, derzeit nicht berücksichtigt",
					"en": "No, currently not considered",
				},
				{
					"de": "keine<br>Angabe",
					"en": "no answer",
				},
			},
			[]string{"sq6"},
			radioVals4,
			[]trl.S{
				{
					"de": "&nbsp;",
					"en": "&nbsp;",
				},
			},
		)
		gb.MainLabel = trl.S{
			"de": `
				Haben Umwelt- und Klimarisiken, Ihrer Einschätzung nach, aktuell einen Einfluss auf die Kreditzinsen?
			`,
			"en": `  
				In your opinion, do environmental and climate risks currently have an impact on lending rates?
			`,
		}.Outline("6.")

		gr := page.AddGrid(gb)
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.6rem"

	}

	// gr2
	gr := page.AddGroup()
	{
		lbl := trl.S{
			"de": `
				Was denken Sie, auf welcher Basis erfolgt die Einbeziehung von Klima- und Umweltrisiken bei der Bepreisung in der Kreditvergabe? 
				<br>
				<small>
				(mehrere Antworten möglich)
				</small>
			`,
			"en": `  
				On what basis do you think climate and environmental risks are included in the pricing of loans? 
				<br>
				<small>
				(multiple answers possible) 
				</small>
			`,
		}.Outline("7.")

		lbls := []trl.S{
			{
				"de": `CO2-Emissionen`,
				"en": `CO2 emissions`,
			},
			{
				"de": `ESG-Ratings`,
				"en": `ESG ratings`,
			},
			{
				"de": `Physische Risiken`,
				"en": `Physical risks`,
			},
			{
				"de": `Transitorische Risiken`,
				"en": `Transitory risks`,
			},
			{
				"de": `Projektbezogene Merkmale (z.B. Transformationsprojekt)`,
				"en": `Project related characteristics (e.g. transformation project)`,
			},
			{
				"de": `Andere Kennwerte`,
				"en": `Other parameters`,
			},
		}

		inps := []string{
			"sq7_co2",
			"sq7_esg",
			"sq7_physical",
			"sq7_transitory",
			"sq7_project",
			"sq7_other",
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
			inp.Label = lbl
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

				inp.Label = lbls[i]
			}

		}
	}

	return nil

}
