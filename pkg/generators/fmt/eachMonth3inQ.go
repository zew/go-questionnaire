package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// q1
var rowLabelsAssetClassesEuroZoneQ3 = []trl.S{
	{
		"de": "Aktien",
		"en": "Stocks",
	},
	{
		"de": "Staats&shy;anleihen",
		"en": "Sovereign bonds",
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

var inputNamesAssetClassesEuroZoneQ3 = []string{
	"ass_euro_stocks",
	"ass_euro_bonds_govt",
	"ass_euro_bonds_corp",
	"ass_euro_re",
}

// q2
var rowLabelsAssetClassesGlobalQ3 = []trl.S{
	{
		"de": "Aktien",
		"en": "Stocks",
	},
	{
		"de": "Staats&shy;anleihen",
		"en": "Sovereign bonds",
	},
	{
		"de": "Unter&shy;nehmens&shy;anleihen",
		"en": "Corporate bonds",
	},
	{
		"de": "Immobilien",
		"en": "Real estate",
	},
	{
		"de": "Gold",
		"en": "Gold",
	},
	{
		"de": "Rohstoffe",
		"en": "Raw materials",
	},
	{
		"de": "Krypto&shy;währungen",
		"en": "Crypto currencies",
	},
}

var inputNamesAssetClassesGlobalQ3 = []string{
	"ass_global_stocks",
	"ass_global_bonds_govt",
	"ass_global_bonds_corp",
	"ass_global_re",
	"ass_global_gold",
	"ass_global_raw_materials",
	"ass_global_crypto",
}

// q3
var inputNamesAssetClassesChangeQ3 = []string{
	"chg_euro_stocks",
	"chg_euro_bonds_govt",
	"chg_euro_bonds_corp",
	"chg_euro_re",
}

var influenceFactorLabelsQ3 = []trl.S{
	{
		"de": "Gesamtwirtschaftlicher Ausblick im Eurogebiet",
		"en": "Economic outlook",
	},
	{
		"de": "Geldpolitik der EZB",
		"en": "Monetary policy of the ECB",
	},
	{
		"de": "Geldpolitik der US-Notenbank",
		"en": "Monetary policy of the US Fed",
	},
	{
		"de": "Ausblick Inflation im Eurogebiet",
		"en": "Outlook Inflation",
	},
	{
		"de": "Politische Rahmen&shy;bedingungen",
		"en": "Political situation",
	},
	{
		"de": "Markt&shy;bewertung",
		"en": "Market valuation",
	},
	{
		"de": "Krieg Russ&shy;land - Ukraine",
		"en": "Russia's war with Ukraine",
	},
}

var rowLabelsAssetClassesEuroZoneQ3B = []trl.S{
	{
		"de": "Aktien (Eurogebiet)",
		"en": "Stocks (euro area)",
	},
	{
		"de": "Staats&shy;anleihen (Eurogebiet)",
		"en": "Sovereign bonds (euro area)",
	},
	{
		"de": "Unter&shy;nehmens&shy;anleihen (Eurogebiet)",
		"en": "Corporate bonds (euro area)",
	},
	{
		"de": "Immobilien (Eurogebiet)",
		"en": "Real estate (euro area)",
	},
}

var influenceFactorNamesQ3 = []string{
	"economy",    // overall economic outlook
	"ecb",        // monetary policy ecb
	"fed",        // monetary policy fed
	"inflation",  // outlook inflation
	"politics",   // political framework
	"valuation",  // market valuation
	"warukraine", //
	// "other",     // other
}

func eachMonth3inQ(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2021 && q.Survey.Month == 9
	cond = cond || q.Survey.Year == 2021 && q.Survey.Month == 12
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
				inputNamesAssetClassesEuroZoneQ3,
				radioVals6,
				rowLabelsAssetClassesEuroZoneQ3,
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
					How do you assess the return-risk profile of the following asset classes 
					in the <b><i>euro area</i></b> for the next 6 months? 

					Please consider well-diversified indices.
					
					
				</p>
				<p style=''>
					My assessment of the return-risk profile is …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr2
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
				inputNamesAssetClassesGlobalQ3,
				radioVals6,
				rowLabelsAssetClassesGlobalQ3,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					<b>2.</b> &nbsp;
					Mit Blick auf die nächsten sechs Monate, 
					wie beurteilen Sie das Rendite-Risko-Profil der folgenden Anlageklassen? 
					
					Orientieren Sie sich an <b><i>globalen</i></b>, breit gestreuten Indizes.
				</p>

				<p style=''>
					Das Rendite-Risiko-Profil beurteile ich …
				</p>
				`,
				"en": `
				<p style=''>
					<b>2.</b> &nbsp;
					How do you assess the return-risk profile  
					of the following <b><i>global</i></b> asset classes for the next 6 months? 

					Please consider well-diversified indices.
					
					
				</p>
				<p style=''>
					My assessment of the return-risk profile is …
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr3
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
				inp.Label = trl.S{
					"de": `
				<p style=''>
					<b>3.</b>  &nbsp;
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
					<b>3.</b>  &nbsp;
					Do the following factors motivate you
					to change your expectations (compared to December 2021) 
					regarding the return-risk-profile of 
					asset classes in the euro area?					
				</p>

				<p style=''>
					(+) = upward change, (-) = downward change
				</p>
					

					`,
				}
				inp.ColSpanLabel = 1
			}

		}

		//
		//
		//
		// gr4 ... gr11
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

		for idx, assCl := range inputNamesAssetClassesChangeQ3 {

			names := []string{}
			for _, nm := range influenceFactorNamesQ3 {
				names = append(names, assCl+"__"+nm)
			}

			lbl := rowLabelsAssetClassesEuroZoneQ3B[idx]

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					improvedDeterioratedPlusMinus6(),
					names,
					radioVals6,
					influenceFactorLabelsQ3,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						<span>3.%v.</span> &nbsp;
						%v
						&nbsp; - &nbsp;  Eurogebiet
					</p>
					`,
						idx+1,
						lbl.Tr("de"),
					),
					"en": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						<span>3.%v.</span> &nbsp;

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
