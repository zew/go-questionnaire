package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var inflationFactorNames = []string{
	"business_cycle",   // Konjunkturentwicklung im Eurogebiet,
	"wages",            // Entwicklung der Löhne im Eurogebiet,
	"energy",           // Entwicklung der Energiepreise,
	"raw_materials",    // Entwicklung der Rohstoffpreise (ohne Energiepreise),
	"exchange_rates",   // Veränderung der Wechselkurse (relativ zum Euro),
	"ecb_money_policy", // Geldpolitik der EZB,
	"trade_war",        // Internationale Handelskonflikte,
	"supply_chain",     // Internationale Lieferengpässe,
	"corona",           // Corona-Pandemie,
	// "other",        // other
}

var inflationFactorLabels = []trl.S{
	{
		"de": "Konjunktur&shy;ent&shy;wicklung im Eurogebiet",
		"en": "todo",
	},
	{
		"de": "Entwicklung der Löhne im Eurogebiet",
		"en": "todo",
	},
	{
		"de": "Entwicklung der Energiepreise",
		"en": "todo",
	},
	{
		"de": "Entwicklung der Rohstoffpreise (ohne Energiepreise)",
		"en": "todo",
	},
	{
		"de": "Veränderung der Wechselkurse (relativ zum Euro)",
		"en": "todo",
	},
	{
		"de": "Geldpolitik der EZB",
		"en": "todo",
	},
	{
		"de": "Internationale Handelskonflikte",
		"en": "todo",
	},
	{
		"de": "Internationale Lieferengpässe",
		"en": "todo",
	},
	{
		"de": "Corona-Pandemie",
		"en": "todo",
	},
	// {
	// 	"de": "Andere",
	// 	"en": "Other",
	// },
}

func special202111a(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 11 {
		return nil
	}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage: Inflation, Prognosetreiber und Geldpolitik",
			"en": "Special: Inflation, forecast drivers and monetary policy",
		}
		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik",
			"en": "Inflation,<br>Mon. Policy",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

		//
		{
			gr := page.AddGroup()
			gr.Cols = 9
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.Width = "70%"
			gr.Style.Mobile.StyleBox.Width = "100%"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 9
				// inp.ColSpanLabel = 12
				inp.Label = trl.S{
					"de": `
					<p>
						Angesichts der Inflationsentwicklung der letzten
						 Monate möchten wir Ihre Einschätzungen 
						 zu den Ursachen und der weiteren Entwicklung 
						 der Inflation im Eurogebiet noch detaillierter 
						 als üblicherweise erfragen
					</p>
				`,
					"en": `
					<p>
						todo english
					</p>
				`,
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 9
				// inp.ColSpanLabel = 12
				inp.Label = trl.S{
					"de": `<b>1.</b> &nbsp; Punktprognose der <b>jährlichen Inflationsrate im Euroraum</b>
				<br>
				Anstieg des HICP von Jan bis Dez; Erwartungswert
				`,
					"en": `<b>1.</b> &nbsp; Forecast <b>yearly inflation rate in the Euro area</b>
				<br>
				HICP  increase from Jan to Dec; expected value
				`,
				}
			}

			for idx := range []int{0, 1, 2} {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("ppjinf_jp%v", idx) //"p1_y1"
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.01
				inp.Label = trl.S{
					"de": q.Survey.YearStr(idx),
					"en": q.Survey.YearStr(idx),
				}
				inp.Suffix = trl.S{
					"de": "%",
					"en": "pct",
				}

				inp.ColSpan = 3
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 2

				inp.StyleLbl = lblStyleRight
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 10
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Mobile.StyleGridContainer.GapRow = "0.02rem"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 10
				inp.Label = trl.S{
					"de": `<b>2.</b> &nbsp; Wir möchten gerne von Ihnen erfahren, 
						für wie wahrscheinlich Sie bestimmte Ausprägungen 
						der durchschnittlichen jährlichen Inflationsrate 
						in den Jahren 2021 bis 2023 halten.`,
					"en": "<b>2.</b> &nbsp; todo",
				}

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0.8rem 0"

			}
			// first row: labels
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Kleiner oder gleich 2&nbsp;Prozent",
					"en": "todo",
				}
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style.Mobile.StyleBox.Padding = "0 0.8rem 0 0"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Zwischen 2 und 3&nbsp;Prozent",
					"en": "todo",
				}
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style.Mobile.StyleBox.Padding = "0 0.8rem 0 0"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Größer als 3&nbsp;Prozent",
					"en": "todo",
				}
				inp.Style = css.ItemStartCA(inp.Style)
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "&#931;",
					"en": "&#931;",
				}
				inp.Style = css.ItemCenteredMCA(inp.Style)
				inp.Style = css.ItemStartCA(inp.Style)
			}
			// second to fourth row: inputs
			for i := 2021; i <= 2023; i++ {

				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("i%v_probgood", i)
					inp.Suffix = trl.S{"de": "%", "en": "%"}
					inp.ColSpan = 3
					inp.ColSpanControl = 3
					inp.Min = 0
					inp.Max = 100
					inp.Step = 0
					inp.MaxChars = 4
				}
				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("i%v_probnormal", i)
					inp.Suffix = trl.S{"de": "%", "en": "%"}
					inp.ColSpan = 3
					inp.ColSpanControl = 3
					inp.Min = 0
					inp.Max = 100
					inp.Step = 0
					inp.MaxChars = 4
				}
				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("i%v_probbad", i)
					inp.Suffix = trl.S{"de": "%", "en": "%"}
					inp.ColSpan = 3
					inp.ColSpanControl = 3
					inp.Min = 0
					inp.Max = 100
					inp.Step = 0
					inp.MaxChars = 4
				}
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.Label = trl.S{
						"de": "100&nbsp;%",
						"en": "100&nbsp;%",
					}
					inp.Style = css.ItemCenteredMCA(inp.Style)
				}
			}
		}

	} // special page 0

	//
	//
	// special page 1
	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage: Inflation, Prognosetreiber und Geldpolitik 2",
			"en": "Special: Inflation, forecast drivers and monetary policy 2",
		}
		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik 2",
			"en": "Inflation,<br>Mon. Policy 2",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

		//
		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.WidthMax = "44rem"
			gr.Style.Mobile.StyleBox.WidthMax = "100%"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
					"de": `
					<b>3.</b>  &nbsp;
						Haben Entwicklungen in den folgenden Bereichen
						 Sie zu einer Revision Ihrer <b>Inflationsprognosen</b>
						 (ggü. August 2021) für den Euroraum bewogen 
						 und wenn ja, nach oben (+) oder unten (-)?
					`,

					"en": `
					<b>3.</b>  &nbsp;
						todo
					
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

		years := []string{"2021", "2022", "2023"}

		for idx, year := range years {

			// lbl := rowLabelsAssetClasses202111[idx]
			names := []string{}

			for _, v := range inflationFactorNames {
				names = append(names, fmt.Sprintf("inff_%v__%v", year, v))
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					improvedDeterioratedPlusMinus6(),
					names,
					radioVals6,
					inflationFactorLabels,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						
						Für das Jahr <b>%v</b>:
					</p>
					`,
						idx+2021,
					),
					"en": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						For the year <b>%v</b>:
					</p>
					`,
						idx+2021,
					),
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
				gr.BottomVSpacers = 2
				if idx == 2 {
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
					inp.Name = "inff_" + year + "__other_label"
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

					rad.Name = "inff_" + year + "__free"
					rad.ValueRadio = fmt.Sprint(idx + 1)

					rad.ColSpan = 1
					rad.ColSpanLabel = colsBelow1[2*(idx+1)]
					rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1

				}

			}

		}

	} // special page 1

	//
	//
	return nil
}
