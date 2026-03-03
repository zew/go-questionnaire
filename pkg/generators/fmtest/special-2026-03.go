package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202603tpA(page *qst.WrappedPageT, inputStem string, rowLbls []trl.S) {

	for i := 0; i < len(rowLbls); i++ {

		gr := page.AddGroup()
		gr.Cols = 6
		gr.RandomizationGroup = 2
		gr.BottomVSpacers = 1
		if i == 5 {
			gr.RandomizationGroup = 0
			gr.BottomVSpacers = 3
		}

		{

			inp1 := gr.AddInput()
			inp1.Type = "radio"
			inp1.Name = inputStem
			inp1.ValueRadio = fmt.Sprintf("%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = rowLbls[i]

			inp1.ControlFirst()

			if i == 5 {

				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7

				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = inputStem + "_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1
			}

		}

		// {
		// 	inp := gr.AddInput()
		// 	inp.ColSpanControl = 1
		// 	inp.Type = "javascript-block"
		// 	inp.Name = "radio-xor-number"
		// 	s1 := trl.S{
		// 		"de": "unused",
		// 		"en": "unused",
		// 	}
		// 	inp.JSBlockTrls = map[string]trl.S{
		// 		"msg": s1,
		// 	}
		// 	inp.JSBlockStrings = map[string]string{
		// 		"inp1":    inputStem,
		// 		"inp2":    inputStem + "_pfc",
		// 		"radioOn": inputStem + "6",
		// 	}
		// }

	}
}

func special202603B(page *qst.WrappedPageT, colLabels []trl.S, inputStem string, rowLblsRandomized []trl.S, randGroup int, showFree bool) {

	// colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	colTemplateStr := "7fr       1fr 1fr 1fr 1fr 1fr   1.4fr"
	styleRowFree := "  7fr       1fr 1fr 1fr 1fr 1fr   1.4fr"

	//
	{
		gr := page.AddGroup()
		gr.Cols = 7
		gr.BottomVSpacers = 0

		// equal to below
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		gr.Style.Desktop.StyleText.FontSize = 90

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": "&nbsp;",
				"en": "&nbsp;",
			}
		}
		for i := 0; i < len(colLabels); i++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = colLabels[i]
			inp.LabelCenter()
			inp.LabelBottom()
		}
	}

	//
	//
	for i1 := 0; i1 < len(rowLblsRandomized); i1++ {
		gr := page.AddGroup()
		firstCol := float32(1)
		gr.Cols = firstCol + 6
		gr.RandomizationGroup = randGroup
		gr.BottomVSpacers = 0
		if i1 == (len(rowLblsRandomized) - 1) {
			gr.BottomVSpacers = 2 // bad, because of shuffling
			gr.BottomVSpacers = 0
		}

		// equal to above
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.Display = "grid"
		// gr.Style.Desktop.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		// gr.Style.Mobile.StyleGridContainer.TemplateColumns =  "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		// distinct
		gr.Style.Desktop.StyleBox.Margin = "0 0 0.6rem" // bottom margin

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = firstCol
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = rowLblsRandomized[i1]
		}
		for i2 := 0; i2 < 6; i2++ {
			{
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_%v", inputStem, i1+1)
				inp.ValueRadio = fmt.Sprintf("%v", i2+1)
				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}
	}

	//
	//
	//
	//
	// row free input
	if !showFree {

		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3

	} else {

		gr := page.AddGroup()
		gr.Cols = 7

		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = styleRowFree
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = styleRowFree
		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		gr.BottomVSpacers = 3

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = fmt.Sprintf("%v_free", inputStem)
			// inp.MaxChars = 17
			inp.MaxChars = 25
			inp.ColSpan = 1
			inp.ColSpanLabel = 2.4
			inp.ColSpanLabel = 0.9
			inp.ColSpanControl = 4
			inp.Label = trl.S{
				"de": "Sonstiges:",
				"en": "Other, namely …",
			}
		}

		//
		for idx := 0; idx < len(colLabels); idx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = fmt.Sprintf("%v_free_val", inputStem)
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1

			rad.ColSpanLabel = 0
			rad.ColSpanControl = 1
		}

	}

}

// main
// func
func special202603(q *qst.QuestionnaireT) error {

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
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Klima-<br>erwartungen",
		"en": "Climate<br>Expectations",
	}
	// page.WidthMax("42rem")

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
					Im Dezember 2025 wurden Sie gebeten, mehrere Fragen zu Ihren Überzeugungen zum Klimawandel zu beantworten. Diese Fragen sowie die nachfolgenden sind Teil eines laufenden Forschungsprojekts. Eine weitere Erhebungswelle zu diesem Thema wird im Juni durchgeführt.
				`,
				"en": `
					In December 2025, you were asked to respond to several questions regarding your beliefs about climate change. These questions, together with those that follow, are part of an ongoing research project. A subsequent survey wave on this topic will be conducted in June.
				`,
			}
		}
	}

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
					Wer sollte primär vorangehen, um Unternehmen dazu zu bewegen, ihre Treibhausgasemissionen zu reduzieren?
				`,
				"en": `
					Who should take the lead in encouraging companies to reduce their greenhouse gas emissions?
				`,
			}.Outline("1.")
		}
	}

	lblsSsq1 := []trl.S{
		{
			"de": `Regierungen / Aufsichtsbehörden`,
			"en": "governments / regulators",
		},
		{
			"de": `Konsumenten`,
			"en": "consumers",
		},
		{
			"de": `Institutionelle Investoren (z. B. Pensionsfonds, Investmentfonds, Versicherungsgesellschaften)`,
			"en": "institutional investors (e.g., pension funds, investment funds, insurance companies)",
		},
		{
			"de": `Banken`,
			"en": "banks",
		},
		{
			"de": `Privatanleger`,
			"en": "private investors",
		},
		{
			"de": `Sonstiges`,
			"en": "other, namely...",
		},
	}
	special202603tpA(qst.WrapPageT(page), "ssq1", lblsSsq1)

	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `Skizzieren Sie, wenn Sie möchten, stichpunktartig, welche Ansätze Sie für besonders vielversprechend halten, um eine klimaneutrale Wirtschaft zu erreichen.`,
				"en": `Please feel free to outline, in bullet points, which approaches you consider particularly promising for achieving a climate-neutral economy.`,
			}.Outline("2.")
		}
		{
			inp := gr.AddInput()
			inp.Type = "textarea"
			inp.Name = "ssq2"
			inp.MaxChars = 1000
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1
		}
	}

	//
	//
	//
	//
	//
	page = q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("64rem")

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.SuppressInProgressbar = true

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
					Nun möchten wir den Fokus auf die Rolle des Finanzsektors legen. Bitte bewerten Sie die allgemeine Bedeutung der folgenden Mechanismen, um Unternehmen zur Reduzierung ihrer Treibhausgasemissionen zu bewegen.
				`,
				"en": `
					We now want to focus on the role of the financial sector. Please rate the general importance of the following financial-sector mechanisms in moving corporations to reduce their greenhouse gas emissions.
				`,
			}.Outline("3.")
		}
	}

	colLabelsSsq3and5 := []trl.S{
		{
			"de": "überhaupt nicht wichtig   <br>  <span class='ordinal-numbers'> 1 </span> ",
			"en": "not at all important      <br>  <span class='ordinal-numbers'> 1 </span> ",
		},
		{
			"de": "<br>  <span class='ordinal-numbers'> 2 </span> ",
			"en": "<br>  <span class='ordinal-numbers'> 2 </span> ",
		},
		{
			"de": "<br>  <span class='ordinal-numbers'> 3 </span> ",
			"en": "<br>  <span class='ordinal-numbers'> 3 </span> ",
		},
		{
			"de": "<br>  <span class='ordinal-numbers'> 4 </span> ",
			"en": "<br>  <span class='ordinal-numbers'> 4 </span> ",
		},
		{
			"de": "äußerst wichtig           <br>  <span class='ordinal-numbers'> 5  </span> ",
			"en": "extremely important       <br>  <span class='ordinal-numbers'> 5  </span> ",
		},
		{
			"de": "keine<br>Angabe    <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
			"en": "no answer          <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
		},
	}
	lblsSsq3 := []trl.S{
		{
			"de": `Anpassung des Kreditvolumens durch Banken`,
			"en": `Adjustment of bank lending volumes `,
		},
		{
			"de": `Anpassung der Kreditzinsen durch Banken`,
			"en": `Adjustment of bank lending rates `,
		},
		{
			"de": `Marktbasierte Bepreisung von Klimarisiken`,
			"en": `Market pricing of climate risk `,
		},
		{
			"de": `Nachhaltigkeitsgebundene (“sustainability-linked”) Finanzierungsinstrumente`,
			"en": `Sustainability-linked financial instruments`,
		},
		{
			"de": `Managementfokus auf Nachhaltigkeit durch grüne Finanzierung`,
			"en": `Green financing elevating executive attention`,
		},
		{
			"de": `Dialog (Engagement) institutioneller Investoren mit Unternehmen`,
			"en": `Bilateral engagement by institutional investors`,
		},
		{
			"de": `Aktionärsanträge und Ausübung von Stimmrechten`,
			"en": `Shareholder proposals and voting`,
		},
		{
			"de": `Divestment durch institutionelle Investoren`,
			"en": `Divestment by institutional investors`,
		},
		{
			"de": `Verweigerung von Versicherungsschutz aufgrund von Klimarisiken`,
			"en": `Denying insurance due to climate risk`,
		},
		{
			"de": `Investorenpräferenzen für nachhaltige Finanzanlagen`,
			"en": `Investor preferences for sustainable investments`,
		},
	}
	special202603B(qst.WrapPageT(page), colLabelsSsq3and5, "ssq3", lblsSsq3, 3, false)

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
					Inwieweit sollte der Finanzsektor über seine Rolle als Intermediär hinaus eine proaktive Rolle bei der Finanzierung der grünen Transformation übernehmen? 
					
					<!-  (Überhaupt nicht 1 - In sehr hohem Maße 5)  -->
				`,
				"en": `
					Beyond being an intermediary, to what extent should the financial sector play a proactive role in financing the green transition? 
					
					<!-- (Not at all 1 - To a great extent 5). -->
				`,
			}.Outline("4.")
		}
	}
	//
	//
	colLabelsSsq4 := []trl.S{
		{
			"de": "überhaupt nicht     <br>  <span class='ordinal-numbers'> 1 </span> ",
			"en": "not at all          <br>  <span class='ordinal-numbers'> 1 </span> ",
		},
		{
			"de": "<br>  <span class='ordinal-numbers'> 2 </span> ",
			"en": "<br>  <span class='ordinal-numbers'> 2 </span> ",
		},
		{
			"de": "<br>  <span class='ordinal-numbers'> 3 </span> ",
			"en": "<br>  <span class='ordinal-numbers'> 3 </span> ",
		},
		{
			"de": "<br>  <span class='ordinal-numbers'> 4 </span> ",
			"en": "<br>  <span class='ordinal-numbers'> 4 </span> ",
		},
		{
			"de": "in sehr hohem Maße       <br>  <span class='ordinal-numbers'> 5  </span> ",
			"en": "to a great extent        <br>  <span class='ordinal-numbers'> 5  </span> ",
		},
		{
			"de": "keine<br>Angabe    <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
			"en": "no answer          <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
		},
	}
	lblsSsq4 := []trl.S{
		{
			"de": ` &nbsp; `,
			"en": ` &nbsp; `,
		},
	}
	special202603B(qst.WrapPageT(page), colLabelsSsq4, "ssq4", lblsSsq4, 0, false)

	//
	//
	//
	//
	//
	page = q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("64rem")

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.SuppressInProgressbar = true

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
					Bitte bewerten Sie die Bedeutung der folgenden Hindernisse für grüne Investitionen.
				`,
				"en": `
					Please rate the importance of the following barriers for green investments.
				`,
			}.Outline("5.")
		}
	}

	lblsSsq5 := []trl.S{
		{
			"de": `Informationsfriktionen bei der Klassifizierung (z. B. welche wirtschaftlichen Aktivitäten als „grün“ gelten)`,
			"en": `Information frictions in classification (e.g. which economic activities qualify as green)`,
		},
		{
			"de": `Informationsfriktionen bei der Überprüfung (z. B. ob Unternehmen grüne Standards einhalten oder Greenwashing betreiben)`,
			"en": `Information frictions in verification (e.g. whether firms maintain green standards; greenwashing)`,
		},
		{
			"de": `Fehlende materielle Sicherheiten`,
			"en": `Lack of tangible collateral`,
		},
		{
			"de": `Lange Amortisationszeiten`,
			"en": `Long payoff horizons`,
		},
		{
			"de": `Renditeunsicherheit aufgrund künftiger klimapolitischer Maßnahmen`,
			"en": `Payoffs depending on future climate policies`,
		},
		{
			"de": `Renditeunsicherheit aufgrund künftiger technologischer Entwicklungen`,
			"en": `Payoffs depending on future technological developments`,
		},
		{
			"de": `Renditeunsicherheit aufgrund künftiger technologischer Entwicklungen`,
			"en": `Payoffs depending on physical climate risks`,
		},
		{
			"de": `Unattraktive Risiko-Rendite-Profile`,
			"en": `Unattractive risk-return profiles`,
		},
		{
			"de": `Illiquide oder fragmentierte Kapitalmärkte`,
			"en": `Shallow or fragmented capital markets`,
		},
		{
			"de": `Begrenzte Verfügbarkeit investierbarer grüner Projekte`,
			"en": `Limited availability of investable green projects`,
		},
		{
			"de": `Koordinationsprobleme zwischen verschiedenen Interessengruppen`,
			"en": `Coordination challenges across multiple stakeholders`,
		},
		{
			"de": `Kleine und fragmentierte Projekte`,
			"en": `Small-scale and fragmented projects`,
		},
		{
			"de": `Regulatorische Investitionsbeschränkungen`,
			"en": `Regulatory investment constraints`,
		},
		{
			"de": `Investitionsbeschränkungen durch Kunden- und Treuhandmandate`,
			"en": `Client and fiduciary mandate constraints`,
		},
		{
			"de": `Höhere Compliance-Anforderungen bei grünen Finanzierungen`,
			"en": `Higher compliance requirements for green financing`,
		},
	}
	special202603B(qst.WrapPageT(page), colLabelsSsq3and5, "ssq5", lblsSsq5, 4, true)

	return nil

}
