package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// q4, q5a, q5b, q5c
var rowLabels01 = []trl.S{
	{
		"de": "Ausschluss von SWIFT",
		"en": "SWIFT ban",
	},
	{
		"de": "Maßnahmen gegen die russische Zentralbank",
		"en": "Measures against the Russian Central Bank",
	},
	{
		"de": "Einfrieren der Konten russischer Oligarchen",
		"en": "Freezing foreign assets of Russian oligarchs",
	},
	{
		"de": "Politischer Druck auf westliche Unternehmen, die Geschäftsbeziehungen mit russischen Unternehmen einzustellen",
		"en": "Political pressure on Western companies to end business relationships with Russian firms",
	},
}

var inputNames01 = []string{
	"sanction_swift",
	"sanction_rcb",
	"sanction_freezing",
	"sanction_trade",
}

// q6
var rowLabels02 = []trl.S{
	{
		"de": "BIP",
		"en": "GDP",
	},
	{
		"de": "Inflation",
		"en": "Inflation",
	},
	{
		"de": "Haupt&shy;refinanzierungs&shy;fazilität der EZB",
		"en": "Main refinancing rate of the ECB",
	},
}

var inputNames02 = []string{
	"sanction_effect_gdp",
	"sanction_effect_inflation",
	"sanction_effect_ecb_rate",
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
		//
		//

		// gr1 - q4 intro
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
							Als Reaktion auf den russischen Angriff auf die Ukraine 
							hat eine große Zahl an Ländern entschieden, 
							Sanktionen gegen Russland einzuführen.
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

		// gr2 - q4
		{
			inputNamesLp := make([]string, len(inputNames01))
			copy(inputNamesLp, inputNames01)
			for i2, inpn := range inputNamesLp {
				inputNamesLp[i2] = fmt.Sprintf("%v%v", inpn, "_ukraine")
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate,
				zeroToFive(),
				inputNamesLp,
				[]string{"1", "2", "3", "4", "5", "6", "7"},
				rowLabels01,
			)

			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
				<p style=''>
					<b>%v.</b> &nbsp;
					In welchem Maße werden die folgenden Sanktionen Ihrer Einschätzung 
					nach dazu beitragen, den Krieg in der Ukraine zu beenden?
				</p>
				<p style=''>
					(von 0: nicht wirksam bis 5: hochwirksam):
				</p>
				`, 4),

				"en": fmt.Sprintf(`
				<p style=''>
					<b>%v.</b> &nbsp;
					How do you think the following sanctions are effective
					to end the military conflict in Ukraine?
				</p>
				<p style=''>
					(from 0: not effective to 5: very effective):
				</p>
				`, 4),
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true

		}

		// gr3 - intro to q5 a, b, c
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
							<b>5.</b> &nbsp;
							Wie hoch schätzen Sie den Schaden für die jeweilige Volkswirtschaft ein, 
							die durch die folgenden Sanktionsmaßnahmen verursacht werden?
						</p>
						<p style=''>
							(von 0: kein Schaden bis 5: hoher Schaden):
						</p>
					`,

					"en": `
						<p style=''>
							<b>5.</b> &nbsp;
							How damaging do you think the following sanctions will be 
							for the following economies?
						</p>
						<p style=''>
							(from 0: no damage to 5: high damages):
						</p>

					`,
				}
				inp.ColSpanLabel = 1
			}

		}

		//
		// gr4-6 - q 5a, 5b, 5c
		inpNamesSuffixes := []string{"russian", "german", "euro_area"}
		lbls := []trl.S{
			{
				"de": "russische",
				"en": "Russian",
			},
			{
				"de": "deutsche",
				"en": "German",
			},
			{
				"de": "Eurogebiet",
				"en": "euro area",
			},
		}

		for i1, suffix := range inpNamesSuffixes {

			inputNamesLp := make([]string, len(inputNames01))
			copy(inputNamesLp, inputNames01)
			for i2, inpn := range inputNamesLp {
				inputNamesLp[i2] = fmt.Sprintf("%v_%v", inpn, suffix)
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate,
				zeroToFive(),
				inputNamesLp,
				[]string{"1", "2", "3", "4", "5", "6", "7"},
				rowLabels01,
			)

			gb.MainLabel = trl.S{

				"de": fmt.Sprintf(`
				<p style=''>
					<!-- <b>%v.</b> &nbsp; -->
					Für die <b>%v</b> Wirtschaft
				</p>
				<p style='position: relative; top: 1.1rem; height: 0.1rem;'>
					<bx>Sanktionsmaßnahme</bx>
				</p>


				`, i1+5, lbls[i1]["de"]),

				"en": fmt.Sprintf(`
				<p style=''>
					<!-- <b>%v.</b> &nbsp; -->
					For the <b>%v</b> economy
				</p>
				<p style='position: relative; top: 1.1rem; height: 0.1rem;'>
					<bx>Sanction</bx>
				</p>


				`, i1+5, lbls[i1]["en"]),
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr7 - q6
		{

			var columnTemplate = []float32{
				5.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate,
				improvedDeterioratedPlusMinus6(),
				inputNames02,
				// []string{"1", "2", "3", "4", "5", "6", "7"},
				radioVals6,
				rowLabels02,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>6.</b> &nbsp;
					Wie hoch sind Ihrer Einschätzung nach die Wirkungen der gesamten 
					Sanktionsmaßnahmen gegen Russland für Bruttoinlandsprodukt, 
					Inflation und Geldpolitik im Eurogebiet?
				</p>

				<p style=''>
					(-) = Rückgang, (+) = Steigerung:
				</p>

				`,
				"en": `
				<p style=''>
					<b>6.</b> &nbsp;
					How strong will be the
					<b>total effects of all sanctions against Russia</b>
					for GDP, inflation, and monetary policy in the euro area?
				</p>

				<p style=''>
					(-) means downward change, (+) means upward change:
				</p>

				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.BottomVSpacers = 3
		}

	}

	return nil

}
