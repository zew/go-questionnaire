package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
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

var influenceFactorLabels202109 = []trl.S{
	{
		"de": "Gesamtwirtschaftlicher Ausblick im Eurogebiet",
		"en": "Economic outlook for the euro area",
	},
	{
		"de": "Geldpolitik der EZB",
		"en": "ECB monetary policy",
	},
	{
		"de": "Geldpolitik der US-Notenbank",
		"en": "US Federal Reserve monetary policy",
	},
	{
		"de": "Ausblick Inflation im Eurogebiet",
		"en": "Inflation outlook for the euro area",
	},
	{
		"de": "Politische Rahmen&shy;bedingungen in der Eurozone",
		"en": "Political framework in the Eurozone",
	},
	// {
	// 	"de": "Politische Rahmen&shy;bedingungen Eurogebiet",
	// 	"en": "Political framework euro area",
	// },
	// {
	// 	"de": "Geopolitische Rahmen&shy;bedingungen",
	// 	"en": "Geopolitical framework",
	// },
	{
		"de": "Aktuelle Markt&shy;bewertung",
		"en": "Current valuation multiples",
	},
	// {
	// 	"de": "Andere",
	// 	"en": "Other",
	// },
}

var influenceFactorNames202109 = []string{
	"economy",   // overall economic outlook
	"ecb",       // monetary policy ecb
	"fed",       // monetary policy fed
	"inflation", // outlook inflation
	"politics",  // political framework
	// "politics_euro",   // political framework euro area
	// "politics_global", // political framework global
	"valuation", // market valuation
	// "other",     // other
}

func special202109(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || (q.Survey.Month != 9 && q.Survey.Month != 12) {
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
			// "en": "Special: Asset Classes in the Euro Area",
			"en": "Additional questions on the attractiveness of different asset classes",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Anlageklassen",
			"en": "Special:<br>Asset classes",
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
					What is your assessment of the risk-return profile 
					of the following asset classes
 					over the coming six months?
					 
					Think about diversified investments in assets from the <b>Eurozone</b>
				</p>
				<p style=''>
					My assessment of the risk-return profile is …
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
				<p style=''>
					<b>2.</b>  &nbsp;
					Haben Entwicklungen der folgenden Faktoren 
					Sie zu einer Revision Ihrer Einschätzungen 
					zum Rendite-Risiko-Profil der einzelnen Assetklassen 
					gegenüber September 2021 bewogen?

				</p>

				<p style=''>
					Wenn ja, 
					nach oben (+) oder unten (-) ?
				</p>

					`,

					"en": `
				<p style=''>
					<b>2.</b>  &nbsp;
					Did developments in the following areas lead you to 
					change your assessment of the risk-return profiles 
					of the following four asset classes
					(relative to September 2021)?
				</p>

				<p style=''>
					If yes, did you revise them up (+) or down (-) ?
				</p>
					

					`,
				}
				inp.ColSpanLabel = 1
			}

		}

		//
		//
		//
		// gr3 ... gr10
		var columnTemplateLocal = []float32{
			3.6, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}
		// additional row below each block
		colsBelow1 := append([]float32{1.0}, columnTemplateLocal...)
		colsBelow1 = []float32{
			// 1.4, 2.2, //   3.0, 1,  |  4.6 separated to two cols
			1.38, 2.1, //   3.0, 1,  |  4.6 separated to two cols
			0.0, 1, //     3.0, 1,  |  4.6 separated to two cols
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}
		colsBelow2 := []float32{}
		for i := 0; i < len(colsBelow1); i += 2 {
			colsBelow2 = append(colsBelow2, colsBelow1[i]+colsBelow1[i+1])
		}
		// log.Printf("colsBelow1 %+v", colsBelow1)
		// log.Printf("colsBelow2 %+v", colsBelow2)

		for idx, assCl := range inputNamesAssetClassesChange202109 {

			names := []string{}
			for _, nm := range influenceFactorNames202109 {
				names = append(names, assCl+"__"+nm)
			}

			lbl := rowLabelsAssetClasses202109[idx]

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					improvedDeterioratedPlusMinus6(),
					names,
					radioVals6,
					influenceFactorLabels202109,
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
						<!-- &nbsp; - &nbsp;  euro area -->
					</p>
					`,
						idx+1,
						lbl.Tr("en"),
					),
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 1
			}

			{

				gr := page.AddGroup()
				gr.Cols = 7
				gr.BottomVSpacers = 2
				if idx == 3 {
					gr.BottomVSpacers = 4
				}
				stl := ""
				for colIdx := 0; colIdx < len(colsBelow2); colIdx++ {
					stl = fmt.Sprintf(
						"%v   %vfr ",
						stl,
						colsBelow2[colIdx],
					)
				}
				gr.Style = css.NewStylesResponsive(gr.Style)
				if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
					gr.Style.Desktop.StyleBox.Display = "grid"
					gr.Style.Desktop.StyleGridContainer.TemplateColumns = stl
					// log.Printf("fmt special 2021-09: grid template - %v", stl)
				} else {
					log.Printf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", stl, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
				}

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = assCl + "__other_label"
					inp.MaxChars = 17
					inp.ColSpan = 1
					inp.ColSpanLabel = 2.4
					inp.ColSpanControl = 4
					// inp.Placeholder = trl.S{"de": "Andere: Welche?", "en": "Other: Which?"}
					inp.Label = trl.S{
						"de": "Andere",
						"en": "Other",
					}

					// inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
					// inp.StyleCtl.Desktop.StyleBox.WidthMax = "14.0rem"
					// inp.StyleCtl.Mobile.StyleBox.WidthMax = "4.0rem"

				}

				//
				for idx := 0; idx < len(improvedDeterioratedPlusMinus6()); idx++ {
					rad := gr.AddInput()
					rad.Type = "radio"

					rad.Name = assCl + "__other"
					rad.ValueRadio = fmt.Sprint(idx + 1)

					rad.ColSpan = 1
					rad.ColSpanLabel = colsBelow1[2*(idx+1)]
					rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1

					// rad.Label = lbl
					// rad.ControlFirst()
					// rad.LabelRight()

					// 	rad.Validator = "must;comprehensionPOP2"
				}

			}

		}

	}

	return nil

}
