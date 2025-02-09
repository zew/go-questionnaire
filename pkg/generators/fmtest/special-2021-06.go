package fmtest

import (
	"strings"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var rowLabelsAssetClasses1 = []trl.S{
	{
		"de": "Aktien",
		"en": "Stocks",
	},
	{
		"de": "Staats&shy;anleihen",
		"en": "Govt. bonds",
	},
	{
		"de": "Inflations&shy;indexierte Staats&shy;anleihen",
		"en": "Inflation adjusted govt. bonds",
	},
	{
		"de": "Unter&shy;nehmens&shy;anleihen",
		"en": "Corporate bonds",
	},
	{
		"de": "Grüne Anleihen (Staats- und Unter&shy;nehmens&shy;anleihen)",
		"en": "Green bonds (govt. or corporate)",
	},
	{
		"de": "Rohstoffe",
		"en": "Commodities",
	},
	{
		"de": "Immobilien",
		"en": "Real estate",
	},
	{
		"de": "Krypto&shy;währungen",
		"en": "Crypto currencies",
	},
}

var rowLabelsAssetClasses2 = []trl.S{
	{
		"de": "Aktien",
		"en": "Stocks",
	},
	{
		"de": "Staats&shy;anleihen",
		"en": "Govt. bonds",
	},
	{
		"de": "Inflations&shy;indexierte Staats&shy;anleihen",
		"en": "Inflation adjusted govt. bonds",
	},
	{
		"de": "Unter&shy;nehmens&shy;anleihen",
		"en": "Corporate bonds",
	},
	{
		"de": "Grüne Anleihen (Staats- und Unter&shy;nehmens&shy;anleihen)",
		"en": "Green bonds (govt. or corporate",
	},
	{
		"de": "Immobilien",
		"en": "Real estate",
	},
}

func special202106(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 6 {
		return nil
	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfrage: Anlageklassen 1",
			"en": "Special: Asset classes 1",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Anlageklassen 1",
			"en": "Special:<br>Asset classes 1",
		}
		page.WidthMax("46rem")

		//
		// gr1
		{
			var columnTemplateLocal = []float32{
				3.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				positiveNegative5(),
				[]string{
					"ass_glob_stocks",
					"ass_glob_bonds_govt",
					"ass_glob_bonds_tips",
					"ass_glob_bonds_corp",
					"ass_glob_bonds_green",
					"ass_glob_comm",
					"ass_glob_re",
					"ass_glob_crypto",
				},
				radioVals6,
				rowLabelsAssetClasses1,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>1.</b> &nbsp;
					Mit Blick auf die nächsten sechs Monate, 
					wie beurteilen Sie das Rendite-Risko-Profil der folgenden Anlageklassen? 
					
					Orientieren Sie sich an breit gestreuten, <b><i>globalen</i></b> (alle Länder) Indizes.
				</p>

				<p style=''>
					Das Rendite-Risiko-Profil beurteile ich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>1.</b> &nbsp;
					How do you assess the risk return characteristics of following asset classes,
					in the next six months?
					
					Base your consideration on broadly diversified <b><i>global</i></b> indices (across all countries).
				</p>
				<p style=''>
					I assess the risk return characteristics as …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			_ = gr
		}

		//
		// gr2
		{
			var columnTemplateLocal = []float32{
				3.6, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.4, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				improvedDeteriorated6(),
				[]string{
					"chg_glob_stocks",
					"chg_glob_bonds_govt",
					"chg_glob_bonds_tips",
					"chg_glob_bonds_corp",
					"chg_glob_bonds_green",
					"chg_glob_comm",
					"chg_glob_re",
					"chg_glob_crypto",
				},
				radioVals6,
				rowLabelsAssetClasses1,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>2.</b> &nbsp;
					Wie hat sich Ihre Einschätzung zum Rendite-Risko-Profil 
					der folgenden Anlageklassen im Vergleich zu Anfang 2021 verändert?
					
					Orientieren Sie sich an breit gestreuten, <b><i>globalen</i></b> (alle Länder) Indizes.
				</p>
				<p style=''>
					Meine Einschätzung zum Rendite-Risiko-Profil hat sich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>2.</b> &nbsp;
					How has your assessment the risk return characteristics of following 
					asset classes
					changed since beginning of 2021?
					
					Base your consideration on broadly diversified <b><i>global</i></b> indices (across all countries).
				</p>
				<p style=''>
					My assessment of the risk return characteristics has …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			_ = gr
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
			"de": "Sonderfrage: Anlageklassen 2",
			"en": "Special: Asset classes 2",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Anlageklassen 2",
			"en": "Special:<br>Asset classes 2",
		}
		page.WidthMax("46rem")

		//
		// gr1
		{
			var columnTemplateLocal = []float32{
				3.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				positiveNegative5(),
				[]string{
					"ass_euro_stocks",
					"ass_euro_bonds_govt",
					"ass_euro_bonds_tips",
					"ass_euro_bonds_corp",
					"ass_euro_bonds_green",
					"ass_euro_re",
				},
				radioVals6,
				rowLabelsAssetClasses2,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>3.</b> &nbsp;
					Mit Blick auf die nächsten sechs Monate, 
					wie beurteilen Sie das Rendite-Risko-Profil der folgenden Anlageklassen? 
					
					Orientieren Sie sich an breit gestreuten Indizes
					für das <b><i>Eurogebiet</i></b>.
				</p>

				<p style=''>
					Das Rendite-Risiko-Profil beurteile ich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>3.</b> &nbsp;
					How do you assess the risk return characteristics of following asset classes,
					in the next six months?
					
					Base your consideration on broadly diversified indices 
					in the <b><i>euro area</i></b>.
				</p>
				<p style=''>
					I assess the risk return characteristics as …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			_ = gr
		}

		//
		// gr2
		{
			var columnTemplateLocal = []float32{
				3.6, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.4, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				improvedDeteriorated6(),
				[]string{
					"chg_euro_stocks",
					"chg_euro_bonds_govt",
					"chg_euro_bonds_tips",
					"chg_euro_bonds_corp",
					"chg_euro_bonds_green",
					"chg_euro_re",
				},
				radioVals6,
				rowLabelsAssetClasses2,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>4.</b> &nbsp;
					Wie hat sich Ihre Einschätzung zum Rendite-Risko-Profil 
					der folgenden Anlageklassen im Vergleich zu Anfang 2021 verändert?
					
					Orientieren Sie sich an breit gestreuten Indizes
					für das <b><i>Eurogebiet</i></b>.
				</p>
				<p style=''>
					Meine Einschätzung zum Rendite-Risiko-Profil hat sich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>4.</b> &nbsp;
					How has your assessment the risk return characteristics of following 
					asset classes
					changed since beginning of 2021?
					
					Base your consideration on broadly diversified indices 
					in the <b><i>euro area</i></b>.
				</p>
				<p style=''>
					My assessment of the risk return characteristics has …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			_ = gr
		}

		{

			{
				gr := page.AddGroup()
				gr.Cols = 17
				gr.BottomVSpacers = 3

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 17
					inp.Desc = trl.S{
						"de": `
					<p style='position: relative; top: 0.2rem;'>
					<b>5.</b> &nbsp;
						Angenommen, der DAX erreicht in sechs Monaten exakt 
						Ihre Punktprognose aus Frage 6b. 
						Wäre der DAX dann aus heutiger Sicht gemäß der Fundamentaldaten 					
					</p>
					`,
						"en": `
					<p style='position: relative; top: 0.2rem;'>
					<b>5.</b> &nbsp;
						Assuming the DAX reaches exactly your point forecast from question 6b.
						within 6&nbsp;months.

						From today's perspective
						and judging by the fundamentals
						and within the timeframe of 6&nbsp;months
						would the DAX be 
					</p>
					`,
					}
				}

				partIGroupsShort := []string{
					"1:überbewertet&nbsp;:overvalued",
					"2:fair bewertet&nbsp;:fairly valued",
					"3:unterbewertet&nbsp;:undervalued",
					"4:keine Angabe&nbsp;:no answer",
				}

				for _, kv := range partIGroupsShort {
					sp := strings.Split(kv, ":")
					radVal := sp[0]
					trlDe := sp[1]
					trlEn := sp[2]
					lbl := trl.S{"de": trlDe, "en": trlEn}

					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "dax_reversion"
					rad.ValueRadio = radVal
					rad.ColSpan = 4
					rad.ColSpanLabel = 3
					rad.ColSpanControl = 1
					rad.Label = lbl
					// rad.ControlFirst()
					rad.LabelRight()
					// rad.Validator = "must"
				}

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.Desc = trl.S{
						"de": `
					<p style=''>&nbsp;&nbsp;?</p>
					`,
						"en": `
					<p style=''>&nbsp;&nbsp;?</p>
					`,
					}
				}

			}

		}

	}

	return nil

}
