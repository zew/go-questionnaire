package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// q1
var rowLabels01 = []trl.S{
	{
		"de": "todo",
		"en": "SWIFT ban",
	},
	{
		"de": "todo",
		"en": "Measures against the Russian Central Bank",
	},
	{
		"de": "todo",
		"en": "Freezing foreign assets of Russian oligarchs",
	},
	{
		"de": "todo",
		"en": "Political pressure on Western companies to end business relationships with Russian firms",
	},
}

var inputNames01 = []string{
	"sanction_swift",
	"sanction_rcb",
	"sanction_freezing",
	"sanction_trade",
}

// q4
var rowLabels02 = []trl.S{
	{
		"de": "todo",
		"en": "Financial stability Russia",
	},
	{
		"de": "todo",
		"en": "Financial stability Germany",
	},
	{
		"de": "todo",
		"en": "Financial stability EU",
	},
	{
		"de": "todo",
		"en": "GDP Russia",
	},
	{
		"de": "todo",
		"en": "GDP Germany",
	},
	{
		"de": "todo",
		"en": "GDP EU ",
	},
}

var inputNames02 = []string{
	"effects_financial_russia",
	"effects_financial_germany",
	"effects_financial_eu",
	"effects_gdp_russia",
	"effects_gdp_germany",
	"effects_gdp_eu",
}

func special202203(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 3
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
			"de": "Sonderfrage: Ukrainekrieg",
			"en": "Special: War in Ukraine",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Ukrainekrieg",
			"en": "Special:<br>War in Ukraine",
		}
		page.WidthMax("46rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						<p style=''>
							todo
						</p>
					`,

					"en": `
						<p style=''>
							<b>
							In response to the Russian attack on Ukraine,
							a majority of Western countries decided
							to implement economic sanctions against Russia.
							</b>
						</p>

					`,
				}
				inp.ColSpanLabel = 1
			}

		}

		var columnTemplate = []float32{
			5.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.5, 1,
		}

		//
		// gr1-3
		for i1 := 0; i1 < 3; i1++ {

			inputNamesLp := make([]string, len(inputNames01))
			copy(inputNamesLp, inputNames01)
			for i2, inpn := range inputNamesLp {
				inputNamesLp[i2] = fmt.Sprintf("%v%v", inpn, i1+1)
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate,
				zeroToFive(),
				inputNamesLp,
				[]string{"1", "2", "3", "4", "5", "6", "7"},
				rowLabels01,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>

					todo
				</p>
				`,
				"en": fmt.Sprintf(`
				<p style=''>
					<b>%v.</b> &nbsp;
					How do you think the following sanctions are effective
					to end the military conflict in Ukraine
				</p>
				<p style=''>
					(from 0: not effective to 5: very effective):
				</p>
				`, i1+4),
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate,
				zeroToFive(),
				inputNames02,
				[]string{"1", "2", "3", "4", "5", "6", "7"},
				rowLabels02,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					todo
				</p>

				`,
				"en": `
				<p style=''>
					<b>7.</b> &nbsp;
					How strong will be the
					<b>total effects of all sanctions against Russia</b>
					for financial stability and GDP in Russia, Germany, and the EU:
				</p>

				<p style=''>
					(from 0: no damage to 5: high damages):
				</p>

				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	}

	return nil

}
