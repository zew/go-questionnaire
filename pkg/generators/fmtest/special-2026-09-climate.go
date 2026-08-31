package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// main
// func
func special202609Climate(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2026 && q.Survey.Month == 9
	if !cond {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.Short = trl.S{
		"de": "Klima-<br>erwartungen",
		"en": "Climate<br>Expectations",
	}
	page.Label = trl.S{
		"de": "",
		"en": "Beliefs",
	}
	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.WidthMax("48rem")

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
					Das Pariser Klimaabkommen zielt darauf ab, den „Anstieg der durchschnittlichen Erdtemperatur deutlich unter 2 °C über dem vorindustriellen Niveau“ zu halten.  

					<br>
					<br>

					Angenommen, die globale Klimapolitik bleibt so, wie sie heute ist, und es werden keine zusätzlichen Maßnahmen zur Bekämpfung des Klimawandels ergriffen: 
				`,
				"en": `
					The Paris Agreement aims to hold “the increase in the global average temperature to well below 2&nbsp;°C above pre-industrial levels”.  

					<br>
					<br>

					Assuming that global climate policies stay the way they are today, and no additional measures are taken to tackle climate change: 
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
					Wie hoch schätzen Sie unter diesen Bedingungen den Anstieg der durchschnittlichen Erdtemperatur bis zum Ende dieses Jahrhunderts ein? 
					<br>
					<br>
				`,
				"en": `
					What is your expectation for the global average temperature rise by the end of this century under these conditions? 
					<br>
					<br>
				`,
			}.Outline("1.")
		}
	}

	lblsSsq1 := []trl.S{
		{
			"de": `kein signifikanter Anstieg`,
			"en": "no significant rise",
		},
		{
			"de": `etwa  1.5&nbsp;°C`,
			"en": "about 1.5&nbsp;°C",
		},
		{
			"de": `etwa  2&nbsp;°C`,
			"en": "about 2&nbsp;°C",
		},

		{
			"de": `etwa  3&nbsp;°C`,
			"en": "about 3&nbsp;°C",
		},
		{
			"de": `etwa  4&nbsp;°C`,
			"en": "about 4&nbsp;°C",
		},
		{
			"de": `mehr als  4&nbsp;°C`,
			"en": "more than 4&nbsp;°C",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq1", lblsSsq1, 0, false, false)

	//
	//
	//
	//
	//
	//
	page = q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("44rem")

	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.SuppressInProgressbar = true

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Auf einer Skala von 1 (überhaupt nicht sicher) bis 5 (sehr sicher), wie sicher sind Sie sich hinsichtlich Ihrer Einschätzung in der vorherigen Frage? 
				`,
				"en": `
					On a scale from 1 (not at all confident) to 5 (very confident), how confident are you about your assessment in the previous question?`,
			}.Outline("2.")
		}
	}

	{

		gb := qst.NewGridBuilderRadios(
			[]float32{
				0, 1,
				0, 1,
				0, 1,
				0, 1,
				0, 1,
			},
			[]trl.S{
				{
					"de": "überhaupt nicht sicher",
					"en": "not at all confident",
				},
				{
					"de": "&nbsp;",
					"en": "&nbsp;",
				},
				{
					"de": "&nbsp;",
					"en": "&nbsp;",
				},
				{
					"de": "&nbsp;",
					"en": "&nbsp;",
				},
				{
					"de": "sehr sicher",
					"en": "very confident",
				},
			},
			[]string{"ssq2"},
			radioVals5,
			nil,
		)
		gr := page.AddGrid(gb)
		gr.BottomVSpacers = 4
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
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.SuppressInProgressbar = true

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Bitte geben Sie an, inwieweit Sie den folgenden Aussagen zustimmen. 
				`,
				"en": `
					Please indicate to what extent you agree with the following statements.`,
			}.Outline("3.")
		}
	}

	colLabelsSsq3and5 := []trl.S{
		{
			"de": "stimme überhaupt nicht zu <br>  <span class='ordinal-numbers'> 1 </span> ",
			"en": "strongly disagree         <br>  <span class='ordinal-numbers'> 1 </span> ",
		},
		{
			"de": "stimme nicht zu           <br>  <span class='ordinal-numbers'> 2 </span> ",
			"en": "disagree                  <br>  <span class='ordinal-numbers'> 2 </span> ",
		},
		{
			"de": "weder noch                <br>  <span class='ordinal-numbers'> 3 </span> ",
			"en": "undecided                 <br>  <span class='ordinal-numbers'> 3 </span> ",
		},
		{
			"de": "stimme zu                 <br>  <span class='ordinal-numbers'> 4 </span> ",
			"en": "agree                     <br>  <span class='ordinal-numbers'> 4 </span> ",
		},
		{
			"de": "stimme voll und ganz zu   <br>  <span class='ordinal-numbers'> 5  </span> ",
			"en": "strongly agree            <br>  <span class='ordinal-numbers'> 5  </span> ",
		},
		{
			"de": "keine<br>Angabe           <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
			"en": "no answer                 <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
		},
	}
	lblsSsq3 := []trl.S{
		{
			"de": `Der Klimawandel stellt ein bedeutendes Problem für Volkswirtschaften und Finanzmärkte dar. `,
			"en": `Climate change represents a significant issue for economies and financial markets.`,
		},
		{
			"de": `Mit den richtigen Maßnahmen ist es möglich, bis 2050 eine klimaneutrale Wirtschaft zu erreichen. `,
			"en": `With the right measures, it is possible to achieve a climate-neutral economy by 2050.`,
		},
		{
			"de": `Die Wirtschaft kann klimaneutral werden und dabei weiterwachsen. `,
			"en": `The economy can become climate-neutral while growing at the same time.`,
		},
		{
			"de": `Solange keine geeignete Ersatztechnologie verfügbar ist, sollten weiterhin Investitionen in emissionsintensive Sektoren fließen. `,
			"en": `As long as there is no suitable replacement technology, investment should still flow into emissions-intensive industries.`,
		},
		{
			"de": `Die Bewältigung des Klimawandels erfordert, dass emissionsintensive Unternehmen über die nötigen Finanzmittel verfügen, um auf emissionsarme Technologien umzustellen. `,
			"en": `Responding to climate change requires that emissions-intensive companies have the funding to transition to low-emission technologies.`,
		},
		{
			"de": `Die Bewältigung des Klimawandels erfordert den Rückbau emissionsintensiver Sektoren und den Ausbau emissionsarmer Sektoren. `,
			"en": `Responding to climate change requires shrinking emissions-intensive industries and growing low-emissions industries.`,
		},
		{
			"de": `Technologische Innovation wird der entscheidende Faktor für das Erreichen einer klimaneutralen Wirtschaft sein. `,
			"en": `Technological innovation will be the decisive determinant of achieving a climate-neutral economy.`,
		},
		{
			"de": `Klimaanpassung (d. h. Anpassung an die Auswirkungen des Klimawandels) ist wichtiger als Klimaschutz (d. h. Vermeidung oder Minderung von Treibhausgasemissionen).`,
			"en": `Climate adaptation (i.e., adapting to the effects of climate change) is more important than climate mitigation (i.e., preventing or reducing greenhouse gas emissions).`,
		},
		{
			"de": `Dem Umweltschutz sollte Vorrang eingeräumt werden, selbst dann, wenn dies zu einem langsameren Wirtschaftswachstum und zu Arbeitsplatzverlusten führen würde.`,
			"en": `Protecting the environment should be given priority, even if this were to result in slower economic growth and job losses.`,
		},
	}
	randomizedMatrixWithFree(qst.WrapPageT(page), colLabelsSsq3and5, "ssq3", lblsSsq3, 3, nil)

	//
	if false {
		lblsSsq5 := []trl.S{
			{
				"de": `sub-q1`,
				"en": `sub-q1`,
			},
			{
				"de": `sub-q2`,
				"en": `sub-q2`,
			},
		}
		lblFree := trl.S{
			"de": "Sonstiges:",
			"en": "Other, namely …",
		}
		randomizedMatrixWithFree(qst.WrapPageT(page), colLabelsSsq3and5, "ssq5", lblsSsq5, 4, lblFree)
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
					"de": `label for freetext german`,
					"en": `label for freetext english`,
				}.Outline("2.")
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "ssqxx2"
				inp.MaxChars = 1000
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}
	}

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.WidthMax("64rem")
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.Label = trl.S{
		"de": "",
		"en": "Professional role",
	}
	page.Label = trl.S{
		"en": "",
		"de": "",
	}

	page.SuppressInProgressbar = true
	page.WidthMax("64rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Mit welchen der folgenden Tätigkeiten sind Sie in Ihrer derzeitigen Funktion befasst?
					<br>
					<small>Bitte wählen Sie alle zutreffenden Optionen aus.</small>

				`,
				"en": `
					In your current role, which of the following activities do you contribute to?					
					<br>
					<small>Please select all that apply.</small>
					`,
			}.Outline("4.")
		}
	}

	lblsSsq4 := []trl.S{
		{
			"de": `Investitions- oder Portfolioentscheidungen`,
			"en": `investment or portfolio decision-making`,
		},
		{
			"de": `Kreditvergabe oder Kreditentscheidungen`,
			"en": `lending or credit decision-making`,
		},
		{
			"de": `Versicherungsunderwriting`,
			"en": `insurance underwriting`,
		},
		{
			"de": `Risikobewertung`,
			"en": `risk assessment`,
		},
		{
			"de": `Unternehmensstrategie oder Entscheidungen über Investitionsausgaben`,
			"en": `corporate strategy or capital expenditure decisions`,
		},
		{
			"de": `Aufsichtsrechtliche Berichterstattung`,
			"en": `regulatory reporting`,
		},
		{
			"de": `Nachhaltigkeitsberichterstattung`,
			"en": `sustainability reporting`,
		},
		{
			"de": `Produktgestaltung oder -klassifizierung`,
			"en": `product design or classification`,
		},
		{
			"de": `Kundenberatung oder -kommunikation`,
			"en": `client advice or communication`,
		},
		{
			"de": `Wirtschafts-, Markt- oder Branchenanalysen`,
			"en": `economic, market, or sector analysis`,
		},
		{
			"de": `Politikanalyse oder Politikberatung`,
			"en": `policy analysis or advice`,
		},
		{
			"de": `Stewardship oder Engagement`,
			"en": `stewardship or engagement`,
		},
		{
			"de": `Handel`,
			"en": `trading`,
		},
		{
			"de": `sonstiges`,
			"en": `other`,
		},
		// {
		// 	"de": `todo`,
		// 	"en": `none of the above`,
		// },
	}

	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 2
		for i := 0; i < len(lblsSsq4); i++ {

			secondToLast := i == len(lblsSsq4)-2
			secondToLast = i == len(lblsSsq4)-1

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("ssq4_%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = lblsSsq4[i]
			inp1.ControlFirst()

			if secondToLast {
				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7
				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = "ssq4_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1

			}
		}
	}

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.WidthMax("64rem")
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.Label = trl.S{
		"de": "",
		"en": "Taxonomy exposure and knowledge",
	}
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}

	page.SuppressInProgressbar = true
	page.WidthMax("64rem")

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
					Wie vertraut sind Sie mit der EU-Taxonomie für nachhaltige Wirtschaftstätigkeiten?
					<br>
					<br>
				`,
				"en": `
					How familiar are you with the EU Taxonomy for sustainable activities?
					<br>
					<br>
				`,
			}.Outline("5.")
		}
	}

	lblsSsq5 := []trl.S{
		{
			"de": `Ich habe noch nie davon gehört.`,
			"en": "I have never heard of it.",
		},
		{
			"de": `Ich habe davon gehört, weiß aber nur sehr wenig darüber.`,
			"en": "I have heard of it but know very little about it.",
		},
		{
			"de": `Ich habe ein allgemeines Verständnis davon.`,
			"en": "I have a general understanding of it.",
		},

		{
			"de": `Ich verstehe die für meine Arbeit relevanten Aspekte.`,
			"en": "I understand the aspects relevant to my work.",
		},
		{
			"de": `Ich verfüge über detaillierte praktische Kenntnisse.`,
			"en": "I have detailed working knowledge of it.",
		},
		{
			"de": `Ich verfüge über Expertenwissen.`,
			"en": "I have expert-level knowledge of it.",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq5", lblsSsq5, 0, false, false)

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.WidthMax("64rem")
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.SuppressInProgressbar = true
	page.NavigationCondition = "fmt202609Droput"
	page.WidthMax("64rem")

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
					Bestehen für Ihre Organisation oder deren Produkte aufgrund der EU-Taxonomie Berichtspflichten oder Offenlegungspflichten?
					<br>
					<br>
				`,
				"en": `
					Does the EU Taxonomy create any reporting or disclosure obligation for your organisation or its products?
					<br>
					<br>
				`,
			}.Outline("6.")
		}
	}

	lblsSsq6 := []trl.S{
		{
			"de": `ja – eine verpflichtende Vorgabe (z. B. CSRD-Berichterstattung oder SFDR-Produktoffenlegung)`,
			"en": "yes - a mandatory obligation (e.g., CSRD reporting or SFDR product disclosure)",
		},
		{
			"de": `keine verpflichtende Vorgabe – wir berichten bzw. legen Informationen jedoch freiwillig oder auf Nachfrage von Kunden oder Investoren offen`,
			"en": "no mandatory requirement - but we report or disclose voluntarily or in response to client or investor demand",
		},
		{
			"de": `nein`,
			"en": "no",
		},

		{
			"de": `weiß nicht`,
			"en": "don't know",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq6", lblsSsq6, 0, false, false)

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
					In meiner beruflichen Funktion ist die Anwendung der EU-Taxonomie in erster Linie
					<br>
					<br>
				`,
				"en": `
					In my professional role, using the EU Taxonomy is primarily
					<br>
					<br>
				`,
			}.Outline("7.")
		}
	}

	lblsSsq7 := []trl.S{
		{
			"de": `vorgeschrieben`,
			"en": "required ",
		},
		{
			"de": `erwartet`,
			"en": "expected ",
		},
		{
			"de": `erwünscht`,
			"en": "encouraged",
		},
		{
			"de": `optional`,
			"en": "optional ",
		},

		{
			"de": `für meine Funktion nicht relevant`,
			"en": "not relevant to my role",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq7", lblsSsq7, 0, false, false)

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.Label = trl.S{
		"en": "Usage of the Taxonomy <!-- taxonomy exposure follow up -->",
		"de": "todo",
	}
	page.Label = trl.S{
		"en": " <!-- taxonomy exposure follow up --> ",
		"de": "&nbsp;",
	}

	page.SuppressInProgressbar = true
	page.NavigationCondition = "fmt202609Droput"
	page.WidthMax("64rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Welche der folgenden Aussagen beschreibt am besten, wie Sie in Ihrer derzeitigen Funktion mit der EU-Taxonomie befasst sind?
					<br>
					<small>Bitte wählen Sie alle zutreffenden Antworten aus.</small>
				`,
				"en": `
					Which of the following best describes your involvement with the EU Taxonomy in your current role?
					<br>
					<small>Please select all that apply.</small>
				`,
			}.Outline("8.")
		}
	}
	lblsSsq8 := []trl.S{
		{
			"de": `Ich analysiere Informationen mit Bezug zur EU-Taxonomie.`,
			"en": `I analyse Taxonomy-related information. `,
		},
		{
			"de": `Ich erstelle Analysen, Berichte oder Offenlegungen mit Bezug zur EU-Taxonomie.`,
			"en": `I prepare Taxonomy-related analysis, reports, or disclosures. `,
		},
		{
			"de": `Ich prüfe oder beaufsichtige Arbeiten mit Taxonomiebezug.`,
			"en": `I review or supervise Taxonomy-related work. `,
		},
		{
			"de": `Ich vermittle oder erläutere anderen Informationen zur EU-Taxonomie.`,
			"en": `I communicate or explain the Taxonomy to others.`,
		},
		{
			"de": `Ich treffe Entscheidungen, in die Taxonomie-Informationen einfließen.`,
			"en": `I make decisions that use Taxonomy information. `,
		},
		{
			"de": `Ich habe in meiner Arbeit selten mit der EU-Taxonomie zu tun.`,
			"en": `I rarely encounter the Taxonomy in my work. `,
		},
		{
			"de": `sonstiges`,
			"en": `other`,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq8); i++ {

			reallyLast := i == len(lblsSsq8)-1

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("ssq8_%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = lblsSsq8[i]
			inp1.ControlFirst()

			if reallyLast {
				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7
				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = "ssq8_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1

			}
		}
	}

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
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Für welche Zwecke haben Sie die EU-Taxonomie genutzt?
					<br>
					<small>Bitte wählen Sie alle zutreffenden Antworten aus.</small>				
				`,
				"en": `
					For what purpose have you used the EU Taxonomy?
					<br>
					<small>Please select all that apply.</small>
					`,
			}.Outline("9.")
		}
	}
	lblsSsq9 := []trl.S{
		{
			"de": `Berichterstattung oder Offenlegung`,
			"en": `reporting or disclosure`,
		},
		{
			"de": `Investitionen, Kreditvergabe oder damit zusammenhängende Analysen`,
			"en": `investment, lending or related analysis `,
		},
		{
			"de": `Risikobewertung`,
			"en": `risk assessment `,
		},
		{
			"de": `Unternehmensstrategie oder -planung`,
			"en": `corporate strategy or planning`,
		},
		{
			"de": `Produktentwicklung oder -klassifizierung`,
			"en": `product development or classification`,
		},
		{
			"de": `Politik- oder Wirtschaftsanalysen`,
			"en": `policy or economic analysis `,
		},
		{
			"de": `Kundenberatung`,
			"en": `client advice `,
		},

		{
			"de": `sonstiges`,
			"en": `other`,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq9); i++ {

			reallyLast := i == len(lblsSsq9)-1

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("ssq9_%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = lblsSsq9[i]
			inp1.ControlFirst()

			if reallyLast {
				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7
				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = "ssq9_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1

			}
		}
	}

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.Label = trl.S{
		"en": "Where the information comes from <!-- taxonomy exposure follow up -->",
		"de": "",
	}
	page.Label = trl.S{
		"en": "<!-- taxonomy exposure follow up -->",
		"de": "&nbsp;",
	}
	page.SuppressInProgressbar = true
	page.NavigationCondition = "fmt202609Droput"
	page.WidthMax("64rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Wie kommen Sie normalerweise an Informationen zur EU-Taxonomie?
					<br>
					<small>Bitte wählen Sie alle zutreffenden Antworten aus.</small>
				
				`,
				"en": `
					How do you normally obtain EU Taxonomy-related information?
					<br>
					<small>Please select all that apply.</small>
					`,
			}.Outline("10.")
		}
	}
	lblsSsq10 := []trl.S{
		{
			"de": `Ich bewerte wirtschaftliche Aktivitäten direkt anhand der Kriterien der EU-Taxonomie.`,
			"en": `I assess economic activities directly against the EU Taxonomy criteria. `,
		},
		{
			"de": `Ich berechne Taxonomieindikatoren selbst.`,
			"en": `I calculate Taxonomy indicators myself. `,
		},
		{
			"de": `Ich nutze Taxonomieinformationen, die von Unternehmen oder Gegenparteien gemeldet werden.`,
			"en": `I use Taxonomy information reported by companies or counterparties. `,
		},
		{
			"de": `Ich nutze Taxonomieschätzungen oder -klassifizierungen von externen Datenanbietern.`,
			"en": `I use Taxonomy estimates or classifications from external data providers. `,
		},
		{
			"de": `Ich nutze Taxonomieanalysen von Kollegen.`,
			"en": `I use Taxonomy analyses prepared by colleagues. `,
		},
		{
			"de": `Ich nutze Anlageprodukte, Scores oder Ratings, in die Taxonomie-Informationen einfließen.`,
			"en": `I use investment products, scores, or ratings that incorporate Taxonomy information. `,
		},
		{
			"de": `Ich nutze die EU-Taxonomie hauptsächlich als allgemeine Hintergrundinformation oder Referenz.`,
			"en": `I use the EU Taxonomy mainly as general background or reference information. `,
		},

		{
			"de": `sonstiges`,
			"en": `other`,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq10); i++ {

			reallyLast := i == len(lblsSsq10)-1

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("ssq10_%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = lblsSsq10[i]
			inp1.ControlFirst()

			if reallyLast {
				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7
				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = "ssq10_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1

			}
		}
	}

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.Label = trl.S{
		"en": "Intermediation / transmission   <!-- question 11 relevant -->",
		"de": "",
	}
	page.Label = trl.S{
		"en": " <!-- question 11 relevant --> ",
		"de": "&nbsp;",
	}

	page.SuppressInProgressbar = true
	page.NavigationCondition = "fmt202609Droput"
	page.WidthMax("64rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Auf welche der folgenden Arten haben Sie in den letzten 12 Monaten Informationen zur EU-Taxonomie (z. B. Taxonomiekriterien, Angaben zur Taxonomiekonformität oder darauf basierende Entscheidungen) verwendet, kommuniziert oder, weitergegeben?
					<br>
					<small>Bitte wählen Sie alle zutreffenden Antworten aus.</small>				
				`,
				"en": `
					In which of the following ways have you communicated, shared, 
					or incorporated EU Taxonomy information 
					(e.g., its criteria, alignment figures, or decisions based on it) 
					during the past 12 months?
					<br>
					<small>Please select all that apply.</small>
					`,
			}.Outline("11.")
		}
	}
	lblsSsq11a := []trl.S{
		{
			"de": `Ich habe keine Informationen zur EU-Taxonomie verwendet, kommuniziert oder weitergegeben.`,
			"en": `I have not communicated or shared information about the EU Taxonomy.`,
		},
		{
			"de": `Ich habe die EU-Taxonomie Kollegen erläutert oder mit ihnen besprochen.`,
			"en": `I have explained or discussed the EU Taxonomy with colleagues. `,
		},
		{
			"de": `Ich habe die EU-Taxonomie Kunden oder anderen externen Interessengruppen erläutert oder mit ihnen besprochen.`,
			"en": `I have explained or discussed the EU Taxonomy with clients or other external stakeholders.  `,
		},
		{
			"de": `Ich habe taxonomiebezogene Offenlegungen, Berichte oder andere veröffentlichte Informationen erstellt oder daran mitgewirkt.`,
			"en": `I have prepared or contributed to Taxonomy-related disclosures, reports, or other published information.  `,
		},
		{
			"de": `Ich habe die EU-Taxonomie bei der Beratung von Kunden oder bei der Bewertung ihrer Nachhaltigkeitspräferenzen herangezogen.`,
			"en": `I have used the EU Taxonomy when advising clients or assessing their sustainability preferences.  `,
		},
		{
			"de": `Ich habe die EU-Taxonomie bei Empfehlungen an Kollegen berücksichtigt.`,
			"en": `I have incorporated the EU Taxonomy into advice or recommendations provided to colleagues. `,
		},
		{
			"de": `Ich habe Finanzprodukte konzipiert, gekennzeichnet oder vermarktet, die die EU-Taxonomie berücksichtigen.`,
			"en": `I have designed, labeled, or marketed financial products that incorporate the EU Taxonomy.  `,
		},
		{
			"de": `Ich habe im Namen von Kunden oder meiner Organisation Investitions-, Kredit- oder Finanzierungsentscheidungen unter Verwendung der EU-Taxonomie getroffen oder daran mitgewirkt.`,
			"en": `I have made or contributed to investment, lending or financing decisions on behalf of clients or my organisation using the EU Taxonomy. `,
		},

		{
			"de": `sonstiges`,
			"en": `other`,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq11a); i++ {

			secondLast := i == len(lblsSsq11a)-2
			secondLast = i == len(lblsSsq11a)-1

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("ssq11a_%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = lblsSsq11a[i]
			inp1.ControlFirst()

			if secondLast {
				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7
				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = "ssq11a_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1

			}
		}

		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "grey-out-other-checkboxes"

			s1 := trl.S{
				"de": "",
				"en": "",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}
			inp.JSBlockStrings = map[string]string{
				"core_id": "ssq11a_",
			}
		}

	}

	//
	//
	//
	//
	//
	page = q.AddPage()
	page.Label = trl.S{
		"en": "&nbsp;",
		"de": "&nbsp;",
	}
	page.Label = trl.S{
		"en": "Intermediation / transmission   <!-- question 11 irrelevant -->",
		"de": "",
	}
	page.Label = trl.S{
		"en": " <!-- question 11 irrelevant --> ",
		"de": "&nbsp;",
	}

	page.SuppressInProgressbar = true
	page.NavigationCondition = "fmt202609Droput"
	page.WidthMax("64rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": `
					Auch wenn die EU-Taxonomie für Ihre derzeitige Funktion nicht relevant ist, kann es dennoch vorkommen, dass Sie beruflich damit in Berührung kommen oder darüber diskutieren. Haben Sie in den letzten 12 Monaten auf eine der folgenden Arten Informationen zur EU-Taxonomie verwendet, kommuniziert oder weitergegeben?
					<br>
					<small>Bitte wählen Sie alle zutreffenden Antworten aus.</small>
				`,
				"en": `
					Even though the EU Taxonomy is not relevant to your current role, 
					you may still encounter or discuss it professionally. 
					Have you communicated, shared, or incorporated EU Taxonomy information 
					in any of the following ways during the past 12 months?
					<br>
					<small>Please select all that apply.</small>
					`,
			}.Outline("8.")
			// .Outline("11.")
		}
	}
	lblsSsq11b := []trl.S{
		{
			"de": `Ich habe keine Informationen zur EU-Taxonomie verwendet, kommuniziert oder weitergegeben.`,
			"en": `I have not communicated or shared information about the EU Taxonomy. `,
		},
		{
			"de": `Ich habe die EU-Taxonomie Kollegen erläutert oder mit ihnen besprochen.`,
			"en": `I have explained or discussed the EU Taxonomy with colleagues. `,
		},
		{
			"de": `Ich habe die EU-Taxonomie Kunden oder anderen externen Interessengruppen erläutert oder mit ihnen besprochen.`,
			"en": `I have explained or discussed the EU Taxonomy with clients or other external stakeholders. `,
		},
		{
			"de": `Ich habe in Präsentationen, Veröffentlichungen, im Unterricht oder in anderer beruflicher Kommunikation auf die EU-Taxonomie Bezug genommen.`,
			"en": `I have referred to the EU Taxonomy in presentations, publications, teaching, or other professional communication. `,
		},
		{
			"de": `Ich habe die EU-Taxonomie in beruflichen Diskussionen als Beispiel oder Bezugspunkt verwendet.`,
			"en": `I have used the EU Taxonomy as an example or point of reference in professional discussions. `,
		},

		{
			"de": `sonstiges`,
			"en": `other`,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq11b); i++ {

			secondLast := i == len(lblsSsq11b)-2
			secondLast = i == len(lblsSsq11b)-1

			inp1 := gr.AddInput()
			inp1.Type = "checkbox"
			inp1.Name = fmt.Sprintf("ssq11b_%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = lblsSsq11b[i]
			inp1.ControlFirst()

			if secondLast {
				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7
				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = "ssq11b_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1

			}
		}

		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "grey-out-other-checkboxes"

			s1 := trl.S{
				"de": "",
				"en": "",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}
			inp.JSBlockStrings = map[string]string{
				"core_id": "ssq11b_",
			}
		}

	}

	return nil

}
