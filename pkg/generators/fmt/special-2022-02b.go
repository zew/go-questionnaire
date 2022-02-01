package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202202b(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2022 || q.Survey.Month != 2 {
		return nil
	}

	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage: Inflation, Prognosetreiber und Geldpolitik 3",
			"en": "Special: Inflation, forecast drivers and monetary policy 3",
		}
		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik 3",
			"en": "Inflation,<br>Mon. Policy 3",
		}
		page.WidthMax("48rem")

		//
		//
		//
		// gr3 ... gr10
		var columnTemplateLocal1 = []float32{
			0.2, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal1,
				labelsStrongIncreaseStrongDecrease2(),
				[]string{"infl_euro_area_2025_2030"},
				radioVals6,
				[]trl.S{
					{
						"de": " &nbsp; ",
						"en": " &nbsp; ",
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
					<b>4.</b>  &nbsp;
						Gegenüber dem Zeitraum 2022-2024 wird die durchschnittliche Inflationsrate im Euroraum im 
						<b>Zeitraum 2025-2030</b>
						
					`,

				"en": `
					<b>4.</b>  &nbsp;
					
					`,
			}

			gr := page.AddGrid(gb)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.6rem"
			// gr.BottomVSpacers = 1
		}

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal1,
				labelsStrongIncreaseStrongDecrease2(),
				[]string{"ezb_rate_2025_2030"},
				radioVals6,
				[]trl.S{
					{
						"de": " &nbsp; ",
						"en": " &nbsp; ",
					},
				},
			)

			gb.MainLabel = trl.S{
				"de": `
					<b>5.</b>  &nbsp;
						Gegenüber dem Zeitraum 2022-2024 wird der durchschnittliche Hauptrefinanzierungssatz der EZB im 
						<b>Zeitraum 2025-2030</b>
						
					`,

				"en": `
					<b>5.</b>  &nbsp;
					
					`,
			}

			gr := page.AddGrid(gb)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.6rem"
			// gr.BottomVSpacers = 1
		}

		//
		//
		//
		//
		//
		//
		// gr3 ... gr10
		var columnTemplateLocal2 = []float32{
			3.6, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1,
			0.4, 1,
		}
		// additional row below each block
		colsBelow1 := append([]float32{1.0}, columnTemplateLocal2...)
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

		{

			var ecbFactorName2 = []string{
				"growth_euro_area",
				"demographics",
				"green_transform",
				"globalization",
				"govt_debt",
			}

			var ecbFactorLabels = []trl.S{
				{
					"de": "Wirtschaftswachstum im Euroraum",
					"en": "Euro area economic growth",
				},
				{
					"de": "Demographische Entwicklungen",
					"en": "Demographic trends",
				},
				{
					"de": "Grüne Transformation",
					"en": "Green transformation",
				},
				{
					"de": "Entwicklung der Globalisierung",
					"en": "Globalisation trends",
				},
				{
					"de": "Staatsverschuldung im Eurogebiet",
					"en": "Government debt of Euro area",
				},
			}

			// lbl := rowLabelsAssetClasses202111[idx]
			names := []string{}

			for _, v := range ecbFactorName2 {
				names = append(names, fmt.Sprintf("ecb_rate__%v", v))
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal2,
					improvedDeterioratedPlusMinus6(),
					names,
					radioVals6,
					ecbFactorLabels,
				)

				gb.MainLabel = trl.S{
					"de": `
					<b>6.</b>  &nbsp;
						Wie beurteilen Sie den Einfluss der folgenden Faktoren auf den <b>Hauptrefinanzierungssatz</b> der EZB im 
						<b>Zeitraum 2025-2030</b>
						
					`,

					"en": `
					<b>6.</b>  &nbsp;
					
					`,
				}

				gr := page.AddGrid(gb)
				gr.BottomVSpacers = 1
			}

			//
			//
			// row free input
			{
				gr := page.AddGroup()
				gr.Cols = 7
				gr.BottomVSpacers = 4
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
					inp.Name = fmt.Sprintf("ecb_rate__other_label")
					inp.MaxChars = 17
					inp.ColSpan = 1
					inp.ColSpanLabel = 2.4
					inp.ColSpanControl = 4
					// inp.Placeholder = trl.S{"de": "Andere: Welche?", "en": "Other: Which?"}
					inp.Label = trl.S{
						"de": "Andere",
						"en": "Other",
					}
				}

				//
				for idx := 0; idx < len(improvedDeterioratedPlusMinus6()); idx++ {
					rad := gr.AddInput()
					rad.Type = "radio"

					rad.Name = fmt.Sprintf("ecb_rate__free")
					rad.ValueRadio = fmt.Sprint(idx + 1)

					rad.ColSpan = 1
					rad.ColSpanLabel = colsBelow1[2*(idx+1)]
					rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1

				}

			}

		}

	} // page

	//
	//
	return nil
}
