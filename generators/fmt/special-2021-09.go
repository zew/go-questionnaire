package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var rowLabelsAssetClasses202109 = []trl.S{
	{
		"de": "Aktien",
		"en": "Stocks",
	},
	{
		"de": "Staats&shy;anleihen",
		"en": "Govt. bonds",
	},
	{
		"de": "Unter&shy;nehmens&shy;anleihen",
		"en": "Corporate bonds",
	},
	{
		"de": "Immobilien",
		"en": "Real estate",
	},
}

var inputNamesAssetClasses202109 = []string{
	"ass_euro_stocks",
	"ass_euro_bonds_govt",
	"ass_euro_bonds_corp",
	"ass_euro_re",
}

var inputNamesAssetClassesChange202109 = []string{
	"chg_euro_stocks",
	"chg_euro_bonds_govt",
	"chg_euro_bonds_corp",
	"chg_euro_re",
}

var labelsInputVariables202109 = []trl.S{
	{
		"de": "Gesamtwirtschaftlicher Ausblick",
		"en": "Economic outlook",
	},
	{
		"de": "Geldpolitik der EZB",
		"en": "ECB monetary policy",
	},
	{
		"de": "Ausblick Inflation",
		"en": "Inflation outlook",
	},
	{
		"de": "Politische Rahmenbedingungen",
		"en": "Political framework",
	},
	{
		"de": "Aktuelle Marktbewertung",
		"en": "Current market valuation",
	},
	{
		"de": "Andere",
		"en": "Other",
	},
}

var namesInputVariables202109 = []string{
	"economy",   // overall economic outlook
	"ecb",       // monetary policy ecb
	"inflation", // outlook inflation
	"politics",  // political framework
	"valuation", // market valuation
	"other",     // other
}

func special202109(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 9 {
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
			"de": "Sonderfrage: Anlageklassen im Eurogebiet",
			"en": "Special: Asset Classes in the Euro Area",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Anlageklassen",
			"en": "Special:<br>Asset classes",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "46rem")

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
				inputNamesAssetClasses202109,
				radioVals6,
				rowLabelsAssetClasses202109,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>1.</b> &nbsp;
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
					<b>1.</b> &nbsp;
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
			gr.OddRowsColoring = true
		}

		//
		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.WidthMax = "30rem"
			gr.Style.Mobile.StyleBox.WidthMax = "100%"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
					"de": `
					<b>2.</b>  &nbsp;
					Haben Entwicklungen der folgenden Faktoren 
					Sie zu einer Revision Ihrer Einschätzungen 
					zum Rendite-Risiko-Profil der einzelnen Assetklassen 
					gegenüber Juni 2021 bewogen?
					<br>
					Wenn ja, 
					nach oben (+) oder unten (-)?
					`,

					"en": `
					<b>2.</b>  &nbsp;
					Has your assessment of the risk return characteristics of following 
					asset classes
					been influenced by any of the
					following factors 
					since June of 2021?
					
					<br>
					If so, 
					upwards (+) or downwards (-)?

					`,
				}
				inp.ColSpanLabel = 1
			}

		}

		//
		//
		//
		// gr3 ... gr7
		var columnTemplateLocal = []float32{
			3.6, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}
		for idx, assCl := range inputNamesAssetClassesChange202109 {

			names := []string{}
			for _, nm := range namesInputVariables202109 {
				names = append(names, assCl+"__"+nm)
			}

			lbl := rowLabelsAssetClasses202109[idx]

			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				improvedDeterioratedPlusMinus6(),
				names,
				radioVals6,
				labelsInputVariables202109,
			)

			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
				<p style='position: relative; top: 0.8rem'>
					<span>2.%v.</span> &nbsp;
					%v
					 &nbsp; - &nbsp;  Eurogebiet
				</p>
				`,
					idx+1,
					lbl.Tr("de"),
				),
				"en": fmt.Sprintf(`
				<p style='position: relative; top: 0.8rem'>
					<span>2.%v.</span> &nbsp;

					%v
					 &nbsp; - &nbsp;  euro area
				</p>
				`,
					idx+1,
					lbl.Tr("en"),
				),
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true

			gr.BottomVSpacers = 2
			if idx == 3 {
				gr.BottomVSpacers = 4
			}

		}

	}

	return nil

}
