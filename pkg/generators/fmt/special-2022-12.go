package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var agree6a = []trl.S{
	{
		"de": "Stimme voll zu",
		"en": "Strongly agree",
	},
	{
		"de": "Stimme zu",
		"en": "Agree",
	},
	{
		"de": "Stimme weder zu noch lehne ab",
		"en": "Undecided",
	},
	{
		"de": "Stimme nicht zu",
		"en": "Disagree",
	},
	{
		"de": "Stimme überhaupt nicht zu",
		"en": "Strongly disagree",
	},

	{
		"de": "Keine<br>Angabe",
		"en": "No answer",
	},
}

func special202212(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 12
	if !cond {
		return nil
	}

	//
	//
	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderbefragung zum Jahresende 2022",
			"en": "Special end of year 2022",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022",
			"en": "Special<br>end of 2022",
		}
		page.WidthMax("46rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1

			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.WidthMax = "30rem"
			gr.Style.Mobile.StyleBox.WidthMax = "100%"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
				<p style=''>
					Todo Einleitungstext
				</p>
					`,

					"en": `
				<p style=''>
					Todo Einleitungstext
				</p>
					`,
				}
				inp.ColSpanLabel = 1
			}

		}

	}
	//
	//
	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderbefragung zum Jahresende 2022 - 2",
			"en": "Special end of year 2022 - 2",
		}
		page.Short = trl.S{
			"de": "Sonderfragen<br>Ende 2022 - 2",
			"en": "Special<br>end of 2022 - 2",
		}
		page.WidthMax("46rem")

		//
		// gr1
		{

			const col1Width = 2
			gr := page.AddGroup()
			gr.Cols = col1Width + 3

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = gr.Cols
				inp.Label = trl.S{
					"de": `
				<p style=''>
					Mit einer Wahrscheinlichkeit von 90 Prozent werden die durchschnittliche 
					<b>jährliche Inflationsrate in Deutschland</b> (durchschnittliche jährliche Veränderung des HICP in Prozent) 
					bzw. die durchschnittliche jährliche 
					<b>Wachstumsrate des realen Bruttoinlandprodukts</b> in Deutschland  

					Sicht von <b>zwölf Monaten</b> bzw. <b>drei Jahren</b>   

					zwischen den folgenden Werten liegen:
				</p>
				`,
					"en": `
				<p style=''>
					todo
				</p>
				`,
				}
			}

			colLbls := []trl.S{
				{
					"de": "&nbsp;",
					"en": "&nbsp;",
				},
				{
					"de": "Untergrenze des 90-Prozent-Konfidenzintervalls",
					"en": "Untergrenze des 90-Prozent-Konfidenzintervalls",
				},
				{
					"de": "Obergrenze des 90-Prozent-Konfidenzintervalls",
					"en": "Obergrenze des 90-Prozent-Konfidenzintervalls",
				},
				{
					"de": "Keine Angabe",
					"en": "Keine Angabe",
				},
			}
			for i1 := 0; i1 < len(colLbls); i1++ {
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					if i1 == 0 {
						inp.ColSpan = col1Width
					}
					inp.Label = colLbls[i1]
					inp.LabelCenter()
				}
			}

			inpNames := []string{
				"inflation_12m",
				"inflation_36m",
				"growth_12m",
				"growth_36m",
			}

			rowlblsQ1 := []trl.S{
				{
					"de": "Inflationsrate in Deutschland, 2023",
					"en": "Inflationsrate in Deutschland, 2023",
				},
				{
					"de": "Durchschn. Inflationsrate in Deutschland, 2023-2025",
					"en": "Durchschn. Inflationsrate in Deutschland, 2023-2025",
				},
				{
					"de": "BIP-Wachstumsrate in Deutschland, 2023",
					"en": "BIP-Wachstumsrate in Deutschland, 2023",
				},
				{
					"de": "Durchschn. BIP-Wachstumsrate in Deutschland, 2023-2025",
					"en": "Durchschn. BIP-Wachstumsrate in Deutschland, 2023-2025",
				},
			}

			for i1 := 0; i1 < len(inpNames); i1++ {

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = col1Width
					inp.Label = rowlblsQ1[i1]
				}

				for _, i2 := range []string{"ub", "lb"} {

					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("inf%v_%v", inpNames[i1], i2)
					inp.Suffix = trl.S{"de": "%", "en": "%"}
					inp.ColSpan = 1
					inp.Min = 0
					inp.Max = 40
					inp.Step = 0.05
					inp.MaxChars = 4
				}
				{

					inp := gr.AddInput()
					inp.Type = "checkbox"
					inp.Name = fmt.Sprintf("inf%v_%v", inpNames[i1], "no_answer")
					inp.ColSpan = 1
				}

			}

		}

		//

	}

	return nil
}
