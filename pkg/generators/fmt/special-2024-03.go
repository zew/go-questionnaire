package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/qstcp/cpfmt"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202403(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 3
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "Fragen zur Transition der Wirtschaft zur Klimaneutralität",
		"en": "Questions about climate transition",
	}
	page.Short = trl.S{
		"de": "Transition zur<br>Klimaneutralität",
		"en": "Questions about<br>climate transition",
	}
	page.WidthMax("75rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1

			inp.Label = trl.S{
				"de": `
					Für wie wahrscheinlich halten Sie es aus technologischer Sicht, 
					dass eine Mehrheit der Unternehmen aus den folgenden Branchen zu den folgenden Zeitpunkten klimaneutral wird?
					<br>
					<small>Kategorien: --: sehr unwahrscheinlich, -: unwahrscheinlich, +: wahrscheinlich, ++: sehr wahrscheinlich </small>
				`,
				"en": `
					What do you think how likely it is, from a technological standpoint, 
					that a majority of firms from the following sectors will become climate-neutral by the following years?
					<br>
					<small>Categories: --: very unlikely, -: unlikely, +: likely, ++: very likely </small>
				
				`,
			}.Outline("1.")
		}
	}
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "dyn-composite"
			inp.DynamicFunc = fmt.Sprintf("Special202403QS1__%v__%v", 0, 0)
			inp.DynamicFuncParamset = ""
			inp.ColSpanControl = 1
		}

		_, inputNames, _ := cpfmt.Special202403QS1(q, 0, 0, true)
		for _, inpName := range inputNames {
			inp := gr.AddInput()
			inp.Type = "dyn-composite-scalar"
			inp.Name = inpName
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
			inp.ColSpanLabel = 1

			inp.Label = trl.S{
				"de": `
					Wie hoch schätzen Sie den wirtschaftlichen Nutzen und die Kosten des Übergangs zur Klimaneutralität für Unternehmen aus den folgenden Sektoren insgesamt ein? 
					<br>
					<small>Kategorien:  0: irrelevant, +: niedrig, ++: mittel, +++: groß, ++++: sehr groß</small>
				`,
				"en": `
					How significant in economic terms do you think will the benefits and costs of the transition to climate-neutrality be for firms from the following sectors? 
					<br>
					<small>Categories:  0: insignificant, +: low significance, ++: medium significance, +++: large significance, ++++: very large significance </small>
				
				`,
			}.Outline("2.")
		}
	}
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "dyn-composite"
			inp.DynamicFunc = fmt.Sprintf("Special202403QS2__%v__%v", 0, 0)
			inp.DynamicFuncParamset = ""
			inp.ColSpanControl = 1
		}

		_, inputNames, _ := cpfmt.Special202403QS2(q, 0, 0, true)
		for _, inpName := range inputNames {
			inp := gr.AddInput()
			inp.Type = "dyn-composite-scalar"
			inp.Name = inpName
		}

	}

	//
	// gr1
	rowLbls202403QSS := []trl.S{
		{
			"de": `Fahrzeugbau`,
			"en": `Automotive`,
		},
		{
			"de": `Industrieunternehmen <ssmall>(Chemie, Pharma, Stahl, NE-Metalle, Elektro, Maschinenbau)</ssmall>`,
			"en": `Industrials <ssmall>(Chemicals, Pharma, Steel, Metal Products, Electronics, Machinery)</ssmall>`,
		},
		{
			"de": `Baugewerbe`,
			"en": `Construction`,
		},
		{
			"de": `Versorger  <ssmall>(e.g. Elektrizität, Gas, Wasser)</ssmall>`,
			"en": `Utilities  <ssmall>(e.g. electricity, gas, water)</ssmall>`,
		},
	}
	hrdsQSS3and4 := []trl.S{
		{
			"de": "2024-2030",
			"en": "2024-2030",
		},
		{
			"de": "2030-2040",
			"en": "2030-2040",
		},
		{
			"de": "2040-2050",
			"en": "2040-2050",
		},
		{
			"de": "nie",
			"en": "never",
		},
		{
			"de": "keine <br> Angabe",
			"en": "no <br> answer",
		},
	}
	columnTplQSS3and4 := []float32{
		3.0, 1,
		0.0, 1,
		0.0, 1,
		0.0, 1,
		0.5, 1,
	}

	inpNamesQSS3 := []string{
		"qss3_automotive",
		"qss3_industr", // industrials gets hyphenated
		"qss3_construction",
		"qss3_utilities",
	}
	inpNamesQSS4 := []string{
		"qss4_automotive",
		"qss4_industr", // industrials gets hyphenated
		"qss4_construction",
		"qss4_utilities",
	}

	{
		gb := qst.NewGridBuilderRadios(
			columnTplQSS3and4,
			hrdsQSS3and4,
			inpNamesQSS3,
			radioVals6,
			rowLbls202403QSS,
		)

		gb.MainLabel = trl.S{
			"de": `

						<style> i {font-size:110%} </style>


						Stellen Sie sich bitte vor, 
						dass es in jedem der folgenden Sektoren die folgenden zwei Typen von Unternehmen gibt: 						

					<div style='margin-left: 2rem;'>
						<p>
							<b>A:</b>&nbsp; Unternehmen, die bis zum Jahr 2050 klimaneutral sein werden,
						</p>
						<p>
							<b>B:</b>&nbsp; Unternehmen, die sich nicht anpassen und bis zum Jahr 2050 nicht klimaneutral sein werden.
						</p>
					</div>
						
					<p >
						Wann, glauben Sie, werden Unternehmen des Typs A im Durchschnitt <i>profitabler</i> sein 
						als Unternehmen des Typs B?
					</p>
						
				`,
			"en": `

						<style> i {font-size:110%} </style>

						Imagine that in every of the following sectors, there are firms that

					<div style='margin-left: 2rem;'>
						<p>
							<b>A:</b>&nbsp; will transition to climate-neutrality by 2050
						</p>
						<p>
							<b>B:</b>&nbsp; do not want to change and will not be climate neutral by 2050
						</p>
					</div>
						
					<p >
						When do you think will firms of type A begin to be on average more <i>profitable</i> than firms of type B?	
					</p>


				`,
		}.Outline("3.")

		gr := page.AddGrid(gb)
		_ = gr
	}

	//
	//
	{
		gb := qst.NewGridBuilderRadios(
			columnTplQSS3and4,
			hrdsQSS3and4,
			inpNamesQSS4,
			radioVals6,
			rowLbls202403QSS,
		)

		gb.MainLabel = trl.S{
			"de": `
					Stellen Sie sich bitte erneut vor, 
					dass es in jedem der folgenden Sektoren die folgenden zwei Typen von Unternehmen gibt: 						

					<div style='margin-left: 2rem;'>
						<p>
							<b>A:</b>&nbsp; Unternehmen, die bis zum Jahr 2050 klimaneutral sein werden,
						</p>
						<p>
							<b>B:</b>&nbsp; Unternehmen, die sich nicht anpassen und bis zum Jahr 2050 nicht klimaneutral sein werden.
						</p>
					</div>
						
					<p >
						Wann, glauben Sie, werden Unternehmen des Typs A im Durchschnitt 
						<i>weniger riskant</i> (mit Blick auf das Ausfallrisiko) sein als Unternehmen des Typs B?
					</p>
						
				`,
			"en": `
						Imagine again that in every of the following sectors, there are firms that

					<div style='margin-left: 2rem;'>
						<p>
							<b>A:</b>&nbsp; will transition to climate-neutrality by 2050
						</p>
						<p>
							<b>B:</b>&nbsp; do not want to change and will not be climate neutral by 2050
						</p>
					</div>
						
					<p >
						When do you think will firms of type A begin to be on average less <i>risky</i> (in terms of default risk) than firms of type B?
					</p>


				`,
		}.Outline("4.")

		gr := page.AddGrid(gb)
		_ = gr
	}

	return nil
}
