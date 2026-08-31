package fmtest

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// main
// func
func special202603Climate(q *qst.QuestionnaireT) error {

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
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq1", lblsSsq1, 2, true, true)

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
	randomizedMatrixWithFree(qst.WrapPageT(page), colLabelsSsq3and5, "ssq3", lblsSsq3, 3, nil)

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
	randomizedMatrixWithFree(qst.WrapPageT(page), colLabelsSsq4, "ssq4", lblsSsq4, 0, nil)

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

	lblFree := trl.S{
		"de": "Sonstiges:",
		"en": "Other, namely …",
	}

	randomizedMatrixWithFree(qst.WrapPageT(page), colLabelsSsq3and5, "ssq5", lblsSsq5, 4, lblFree)

	return nil

}
