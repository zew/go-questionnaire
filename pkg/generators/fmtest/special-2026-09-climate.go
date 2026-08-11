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
		"de": "todo",
		"en": "Beliefs",
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
					todo
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
					todo
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
			"de": `todo`,
			"en": "no significant rise",
		},
		{
			"de": `todo`,
			"en": "about 1.5&nbsp;°C",
		},
		{
			"de": `todo`,
			"en": "about 2&nbsp;°C",
		},

		{
			"de": `todo`,
			"en": "about 2&nbsp;°C",
		},
		{
			"de": `todo`,
			"en": "about 4&nbsp;°C",
		},
		{
			"de": `todo`,
			"en": "more than 4&nbsp;°C",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq1", lblsSsq1, 0, false)

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
				"de": `todo`,
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
					"de": "todo",
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
					"de": "todo",
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
				"de": `todo`,
				"en": `
					Please indicate to what extent you agree with the following statements.`,
			}.Outline("3.")
		}
	}

	colLabelsSsq3and5 := []trl.S{
		{
			"de": "todo   <br>  <span class='ordinal-numbers'> 1 </span> ",
			"en": "strongly disagree      <br>  <span class='ordinal-numbers'> 1 </span> ",
		},
		{
			"de": "todo<br>      <span class='ordinal-numbers'> 2 </span> ",
			"en": "disagree<br>  <span class='ordinal-numbers'> 2 </span> ",
		},
		{
			"de": "todo<br>       <span class='ordinal-numbers'> 3 </span> ",
			"en": "undecided<br>  <span class='ordinal-numbers'> 3 </span> ",
		},
		{
			"de": "todo<br>   <span class='ordinal-numbers'> 4 </span> ",
			"en": "agree<br>  <span class='ordinal-numbers'> 4 </span> ",
		},
		{
			"de": "todo                 <br>  <span class='ordinal-numbers'> 5  </span> ",
			"en": "strongly agree       <br>  <span class='ordinal-numbers'> 5  </span> ",
		},
		{
			"de": "keine<br>Angabe    <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
			"en": "no answer          <br>  <span class='ordinal-numbers'> &nbsp;  </span>",
		},
	}
	lblsSsq3 := []trl.S{
		{
			"de": `todo`,
			"en": `Climate change represents a significant issue for economies and financial markets.`,
		},
		{
			"de": `todo`,
			"en": `With the right measures, it is possible to achieve a climate-neutral economy by 2050.`,
		},
		{
			"de": `todo`,
			"en": `The economy can become climate-neutral while growing at the same time.`,
		},
		{
			"de": `todo`,
			"en": `As long as there is no suitable replacement technology, investment should still flow into emissions-intensive industries.`,
		},
		{
			"de": `todo`,
			"en": `Responding to climate change requires that emissions-intensive companies have the funding to transition to low-emission technologies.`,
		},
		{
			"de": `todo`,
			"en": `Responding to climate change requires shrinking emissions-intensive industries and growing low-emissions industries.`,
		},
		{
			"de": `todo`,
			"en": `Technological innovation will be the decisive determinant of achieving a climate-neutral economy.`,
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
		"de": "todo",
		"en": "Professional role",
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
				"de": `todo`,
				"en": `
					In your current role, which of the following activities do you personally contribute to?					
					<br>
					<small>Select all that apply.</small>
					`,
			}.Outline("4.")
		}
	}

	lblsSsq4 := []trl.S{
		{
			"de": `todo`,
			"en": `investment or portfolio decision-making`,
		},
		{
			"de": `todo`,
			"en": `lending or credit decision-making`,
		},
		{
			"de": `todo`,
			"en": `insurance underwriting`,
		},
		{
			"de": `todo`,
			"en": `risk assessment`,
		},
		{
			"de": `todo`,
			"en": `corporate strategy or capital expenditure decisions`,
		},
		{
			"de": `todo`,
			"en": `regulatory reporting`,
		},
		{
			"de": `todo`,
			"en": `sustainability reporting`,
		},
		{
			"de": `todo`,
			"en": `product design or classification`,
		},
		{
			"de": `todo`,
			"en": `client advice or communication`,
		},
		{
			"de": `todo`,
			"en": `economic, market, or sector analysis`,
		},
		{
			"de": `todo`,
			"en": `policy analysis or advice`,
		},
		{
			"de": `todo`,
			"en": `stewardship or engagement`,
		},
		{
			"de": `todo`,
			"en": `trading`,
		},
		{
			"de": `todo`,
			"en": `other`,
		},
		{
			"de": `todo`,
			"en": `none of the above`,
		},
	}

	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 2
		for i := 0; i < len(lblsSsq4); i++ {

			secondToLast := i == len(lblsSsq4)-2

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
		"de": "todo",
		"en": "Taxonomy exposure and knowledge",
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
					todo
				`,
				"en": `
					How familiar are you with the EU Taxonomy?
					<br>
					<br>
				`,
			}.Outline("5.")
		}
	}

	lblsSsq5 := []trl.S{
		{
			"de": `todo`,
			"en": "I have never heard of it.",
		},
		{
			"de": `todo`,
			"en": "I have heard of it but know very little about it.",
		},
		{
			"de": `todo`,
			"en": "I have a general understanding of it.",
		},

		{
			"de": `todo`,
			"en": "I understand the aspects relevant to my work.",
		},
		{
			"de": `todo`,
			"en": "I have detailed working knowledge of it.",
		},
		{
			"de": `todo`,
			"en": "I have expert-level knowledge of it.",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq5", lblsSsq5, 0, false)

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
					todo
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
			"de": `todo`,
			"en": "yes - a mandatory obligation (e.g., CSRD reporting or SFDR product disclosure)",
		},
		{
			"de": `todo`,
			"en": "no mandatory requirement - but we report or disclose voluntarily or in response to client or investor demand",
		},
		{
			"de": `todo`,
			"en": "no",
		},

		{
			"de": `todo`,
			"en": "don't know",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq6", lblsSsq6, 0, false)

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
					todo
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
			"de": `todo`,
			"en": "required ",
		},
		{
			"de": `todo`,
			"en": "expected ",
		},
		{
			"de": `todo`,
			"en": "encouraged",
		},
		{
			"de": `todo`,
			"en": "optional ",
		},

		{
			"de": `todo`,
			"en": "not relevant to my role",
		},
	}
	randomizedVerticalRadiosWithFree(qst.WrapPageT(page), "ssq7", lblsSsq7, 0, false)

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
		"de": "todo",
		"en": "Usage of the Taxonomy <!-- taxonomy exposure follow up -->",
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
				"de": `todo`,
				"en": `
					Which of the following best describes your involvement with the EU Taxonomy in your current role?
					<br>
					<small>Select all that apply.</small>
					`,
			}.Outline("8.")
		}
	}
	lblsSsq8 := []trl.S{
		{
			"de": `todo`,
			"en": `I analyse Taxonomy-related information. `,
		},
		{
			"de": `todo`,
			"en": `I prepare Taxonomy-related analysis, reports, or disclosures. `,
		},
		{
			"de": `todo`,
			"en": `I review or supervise Taxonomy-related work. `,
		},
		{
			"de": `todo`,
			"en": `I communicate or explain the Taxonomy to others.`,
		},
		{
			"de": `todo`,
			"en": `I make decisions that use Taxonomy information. `,
		},
		{
			"de": `todo`,
			"en": `I rarely encounter the Taxonomy in my work. `,
		},
		{
			"de": `todo`,
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
				"de": `todo`,
				"en": `
					For what purpose have you used the EU Taxonomy?
					<br>
					<small>Select all that apply.</small>
					`,
			}.Outline("9.")
		}
	}
	lblsSsq9 := []trl.S{
		{
			"de": `todo`,
			"en": `reporting or disclosure`,
		},
		{
			"de": `todo`,
			"en": `investment, lending or related analysis `,
		},
		{
			"de": `todo`,
			"en": `risk assessment `,
		},
		{
			"de": `todo`,
			"en": `corporate strategy or planning`,
		},
		{
			"de": `todo`,
			"en": `product development or classification`,
		},
		{
			"de": `todo`,
			"en": `policy or economic analysis `,
		},
		{
			"de": `todo`,
			"en": `client advice `,
		},

		{
			"de": `todo`,
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
		"de": "todo",
		"en": "Where the information comes from <!-- taxonomy exposure follow up -->",
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
				"de": `todo`,
				"en": `
					How do you normally obtain EU Taxonomy-related information?
					<br>
					<small>Select all that apply.</small>
					`,
			}.Outline("10.")
		}
	}
	lblsSsq10 := []trl.S{
		{
			"de": `todo`,
			"en": `I assess economic activities directly against the EU Taxonomy criteria. `,
		},
		{
			"de": `todo`,
			"en": `I calculate Taxonomy indicators myself. `,
		},
		{
			"de": `todo`,
			"en": `I use Taxonomy information reported by companies or counterparties. `,
		},
		{
			"de": `todo`,
			"en": `I use Taxonomy estimates or classifications from external data providers. `,
		},
		{
			"de": `todo`,
			"en": `I use Taxonomy analyses prepared by colleagues. `,
		},
		{
			"de": `todo`,
			"en": `I use investment products, scores, or ratings that incorporate Taxonomy information. `,
		},
		{
			"de": `todo`,
			"en": `I use the EU Taxonomy mainly as general background or reference information. `,
		},

		{
			"de": `todo`,
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
		"de": "todo",
		"en": "Intermediation / transmission   <!-- question 11 relevant -->",
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
				"de": `todo`,
				"en": `
					In which of the following ways have you communicated, shared, 
					or incorporated EU Taxonomy information 
					(e.g., its criteria, alignment figures, or decisions based on it) 
					during the past 12 months?
					<br>
					<small>Select all that apply.</small>
					`,
			}.Outline("11.")
		}
	}
	lblsSsq11a := []trl.S{
		{
			"de": `todo`,
			"en": `I have explained or discussed the EU Taxonomy with colleagues. `,
		},
		{
			"de": `todo`,
			"en": `I have explained or discussed the EU Taxonomy with clients or other external stakeholders.  `,
		},
		{
			"de": `todo`,
			"en": `I have prepared or contributed to Taxonomy-related disclosures, reports, or other published information.  `,
		},
		{
			"de": `todo`,
			"en": `I have used the EU Taxonomy when advising clients or assessing their sustainability preferences.  `,
		},
		{
			"de": `todo`,
			"en": `I have incorporated the EU Taxonomy into advice or recommendations provided to colleagues. `,
		},
		{
			"de": `todo`,
			"en": `I have designed, labeled, or marketed financial products that incorporate the EU Taxonomy.  `,
		},
		{
			"de": `todo`,
			"en": `I have made or contributed to investment, lending or financing decisions on behalf of clients or my organisation using the EU Taxonomy. `,
		},

		{
			"de": `todo`,
			"en": `other`,
		},

		{
			"de": `todo`,
			"en": `I have not communicated or shared information about the EU Taxonomy. `,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq11a); i++ {

			secondLast := i == len(lblsSsq11a)-2

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
		"de": "todo",
		"en": "Intermediation / transmission   <!-- question 11 irrelevant -->",
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
				"de": `todo`,
				"en": `
					Even though the EU Taxonomy is not relevant to your current role, 
					you may still encounter or discuss it professionally. 
					Have you communicated, shared, or incorporated EU Taxonomy information 
					in any of the following ways during the past 12 months?
					<br>
					<small>Select all that apply.</small>
					`,
			}.Outline("11.")
		}
	}
	lblsSsq11b := []trl.S{
		{
			"de": `todo`,
			"en": `I have explained or discussed the EU Taxonomy with colleagues. `,
		},
		{
			"de": `todo`,
			"en": `I have explained or discussed the EU Taxonomy with clients or other external stakeholders. `,
		},
		{
			"de": `todo`,
			"en": `I have referred to the EU Taxonomy in presentations, publications, teaching, or other professional communication. `,
		},
		{
			"de": `todo`,
			"en": `I have used the EU Taxonomy as an example or point of reference in professional discussions. `,
		},

		{
			"de": `todo`,
			"en": `other`,
		},

		{
			"de": `todo`,
			"en": `I have not communicated or shared information about the EU Taxonomy. `,
		},
	}
	{
		gr := page.AddGroup()
		gr.Cols = 6
		gr.BottomVSpacers = 3
		for i := 0; i < len(lblsSsq11b); i++ {

			secondLast := i == len(lblsSsq11b)-2

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
	}

	return nil

}
