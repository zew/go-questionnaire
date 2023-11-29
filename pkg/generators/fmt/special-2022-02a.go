package fmt

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var inflationFactorNames02 = []string{
	"business_cycle",   // Konjunkturentwicklung im Eurogebiet,
	"wages",            // Entwicklung der Löhne im Eurogebiet,
	"energy",           // Entwicklung der Energiepreise,
	"raw_materials",    // Entwicklung der Rohstoffpreise (ohne Energiepreise),
	"exchange_rates",   // Veränderung der Wechselkurse (relativ zum Euro),
	"ecb_money_policy", // Geldpolitik der EZB,
	"trade_war",        // Internationale Handelskonflikte,
	"supply_chain",     // Internationale Lieferengpässe,
	"corona",           // Corona-Pandemie,
	"green_transform",  // Grüne Transformation,
	// "other",        // other
}

var inflationFactorLabel02s = []trl.S{
	{
		"de": "Konjunktur&shy;ent&shy;wicklung im Eurogebiet",
		"en": "Development of GDP in the euro area",
	},
	{
		"de": "Entwicklung der Löhne im Eurogebiet",
		"en": "Development of wages in the euro area",
	},
	{
		"de": "Entwicklung der Energiepreise",
		"en": "Development of energy prices",
	},
	{
		"de": "Entwicklung der Rohstoffpreise (ohne Energiepreise)",
		"en": "Development of prices for raw materials (except energy)",
	},
	{
		"de": "Veränderung der Wechselkurse (relativ zum Euro)",
		"en": "Changes in exchange rates (relative to the euro)",
	},
	{
		"de": "Geldpolitik der EZB",
		"en": "Monetary policy of the ECB",
	},
	{
		"de": "Internationale Handelskonflikte",
		"en": "International trade conflicts",
	},
	{
		"de": "Internationale Lieferengpässe",
		"en": "International supply bottlenecks",
	},
	{
		"de": "Corona-Pandemie",
		"en": "Covid-19 pandemic",
	},
	{
		"de": "Grüne Transformation",
		"en": "Green transformation",
	},
	// {
	// 	"de": "Andere",
	// 	"en": "Other",
	// },
}

func special202202a(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2022 || q.Survey.Month != 2 {
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
			"en": "Special Questions: Inflation, its causes, and monetary policy",
		}

		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik 1",
			"en": "Inflation,<br>Mon. Policy 1",
		}
		page.WidthMax("42rem")

		page.ValidationFuncName = "fmt-m2-p6"
		page.ValidationFuncMsg = trl.S{
			"de": "Ihre Antworten auf Frage 1b addieren sich nicht zu 100%. Wirklich weiter?",
			"en": "Your answers to question 1b dont add up to 100%. Continue anyway?",
		}

		//
		{
			gr := page.AddGroup()
			gr.Cols = 9
			gr.WidthMax("30rem")

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
						 als üblicherweise erfragen.
					</p>
				`,
					"en": `
					<p>
						Against the background of the current inflation environment, 
						we would like to ask for your assessments 
						of the causes and the further development 
						of inflation in the euro area in more detail.
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
					"de": `<b>1a.</b> &nbsp; Punktprognose der <b>jährlichen Inflationsrate im Euroraum</b><br>
					(durchschnittliche jährliche Veränderung des HICP in Prozent):  
				`,
					"en": `<b>1a.</b> &nbsp;	
					Point forecast of the <b>annual inflation rate in the euro area</b><br>
					(average annual change of the HICP, in percent):
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

		//
		// gr2
		{

			// colspan := float32(2 + 4*3 + 2 + 2)

			gr := page.AddGroup()
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Mobile.StyleGridContainer.GapRow = "0.02rem"
			// gr.WidthMax("30rem")
			gr.ColWidths("1.6fr    2.7fr 3.1fr 3.1fr 2.4fr    2.4fr  1.4fr")

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 7
				inp.Label = trl.S{
					"de": `<b>1b.</b> &nbsp; Wir möchten gerne von Ihnen erfahren, 
						für wie wahrscheinlich Sie bestimmte Ausprägungen 
						der durchschnittlichen jährlichen Inflationsrate 
						in den folgenden Jahren halten.
						
						<br>
						<i>
						Bitte stellen Sie sicher, 
						dass die Summen der Wahrscheinlichkeiten 
						in den Zeilen jeweils 100% ergeben.
						</i>

						
						`,
					"en": `<b>1b.</b> &nbsp; 
						Please assess the probabilities of the following 
						realizations of the average annual inflation 
						from 2022 to 2024 in the euro area 
						
						<br>
						<i>
						Please ensure that the probabilities 
						in every line add up to 100%.
						</i>
						`,
				}

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0.8rem 0"

			}
			// first row: labels
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": "&nbsp;",
					"en": "&nbsp;",
				}
			}

			labels := []trl.S{
				{
					"de": "unter 1&nbsp;Prozent",
					"en": "below 1&nbsp;percent",
				},
				{
					// "de": "zwischen 1&nbsp;u.&nbsp;2&nbsp;Prozent",
					"de": "zwischen 1 u. 2&nbsp;Prozent",
					"en": "between  1 and 2&nbsp;percent",
				},
				{
					// "de": "zwischen 2&nbsp;u.&nbsp;3&nbsp;Prozent",
					"de": "zwischen 2 u. 3&nbsp;Prozent",
					"en": "between  2 and 3&nbsp;percent",
				},
				{
					"de": "größer als 3&nbsp;Prozent",
					"en": "above 3&nbsp;percent",
				},
			}

			// first row - cols 2-5
			for _, lbl := range labels {
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = lbl
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style.Mobile.StyleBox.Padding = " 0 0.3rem 0 0" // prevent overlapping of columns in narrow mobile view

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.LineHeight = 118

			}

			//
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": "&nbsp;&nbsp;&nbsp;&#931;", // greek SUM sign
					"en": "&nbsp;&nbsp;&nbsp;&#931;",
				}
				inp.Style = css.ItemCenteredMCA(inp.Style)
				inp.Style = css.ItemStartMA(inp.Style)
				inp.Style = css.ItemCenteredCA(inp.Style)

				inp.Style = css.TextStart(inp.Style)
				inp.Style = css.TextCACenter(inp.Style)

				inp.Style.Desktop.StyleText.FontSize = 120

			}
			{
				inp := gr.AddInput()
				inp.ColSpan = 1
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "keine Ang.",
					"en": "no answer",
				}
				inp.Style = css.ItemCenteredMCA(inp.Style)
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style = css.TextStart(inp.Style)

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleText.LineHeight = 118
			}

			//
			// second to fourth row: inputs
			for i := 2022; i <= 2024; i++ {

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.Label = trl.S{
						"de": fmt.Sprint(i),
						"en": fmt.Sprint(i),
					}
				}

				inpNames := []string{
					"under1", "between1and2", "between2and3", "above3",
				}

				for _, inpname := range inpNames {
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("inf%v_%v", i, inpname)
					inp.Suffix = trl.S{"de": "%", "en": "%"}
					// inp.Suffix = trl.S{"de": "%", "en": "pct"}
					inp.ColSpan = 1
					inp.ColSpanControl = 3
					inp.Min = 0
					inp.Max = 100
					inp.Step = 0
					inp.MaxChars = 3
				}

				// last two cols
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 1
					inp.Label = trl.S{
						"de": "100%", //  "100&nbsp;%",
						"en": "100%", //  "100&nbsp;%",
					}
					// inp.Style = css.ItemCenteredMCA(inp.Style)
					inp.Style = css.ItemStartMA(inp.Style)
					inp.Style = css.ItemStartCA(inp.Style)
					inp.Style = css.TextStart(inp.Style)

				}
				{
					inp := gr.AddInput()
					inp.Type = "checkbox"
					inp.ColSpan = 1
					inp.Name = fmt.Sprintf("inf%v__noanswer", i)
					inp.ColSpanControl = 1
					// inp.Style = css.ItemCenteredMCA(inp.Style)
					inp.Style = css.ItemStartMA(inp.Style)
					inp.Style = css.ItemStartCA(inp.Style)

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
			"en": "Special Questions: Inflation, its causes, and monetary policy 2",
		}
		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik 2",
			"en": "Inflation,<br>Mon. Policy 2",
		}
		page.WidthMax("48rem")

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
					<b>2.</b>  &nbsp;
						Haben Entwicklungen in den folgenden Bereichen
						 Sie zu einer Revision Ihrer <b>Inflationsprognosen</b> für den Euroraum
						 (ggü. November 2021) bewogen 
						 und wenn ja, nach oben (+) oder unten (-)?
					`,

					"en": `
					<b>2.</b>  &nbsp;
						Did developments in the following areas make you change 
						your inflation forecasts for the euro area  
						(relative to November 2021)? 
						
						If yes, did you revise them up or down?

						<br>

						(response categories: 
							strongly positive (++), positive (+), 
							no influence (0), 
							negative (-), strongly negative (--)) 
					
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

		{
			year := "22_24"
			idx := 2

			// lbl := rowLabelsAssetClasses202111[idx]
			names := []string{}

			for _, v := range inflationFactorNames02 {
				names = append(names, fmt.Sprintf("inff_%v__%v", year, v))
			}

			{
				gb := qst.NewGridBuilderRadios(
					columnTemplateLocal,
					improvedDeterioratedPlusMinus6(),
					names,
					radioVals6,
					inflationFactorLabel02s,
				)

				gb.MainLabel = trl.S{
					"de": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						<b>Für die Jahre 2022, 2023 und 2024:</b>
					</p>
					`,
					),
					"en": fmt.Sprintf(`
					<p style='position: relative; top: 0.8rem'>
						For the years 2022, 2023 and 2024:
					</p>
					`,
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
					inp.Name = fmt.Sprintf("inff_%v__other_label", year)
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

					rad.Name = fmt.Sprintf("inff_%v__other", year)
					rad.ValueRadio = fmt.Sprint(idx + 1)

					rad.ColSpan = 1
					rad.ColSpanLabel = colsBelow1[2*(idx+1)]
					rad.ColSpanControl = colsBelow1[2*(idx+1)] + 1

				}

			}

		}

		// gr4
		/*
			the “central 90 percent confidence interval”
			directly under the Forecast of the ECB’s main refinancing rate

		*/
		{
			// 2019	18 Sep. 0.00
			latestECBRate, err := q.Survey.Param("main_refinance_rate_ecb")
			if err != nil {
				return fmt.Errorf("Set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%%` as in `main refinancing operations rate of the ECB (01.02.2018: 3.2%%)`; error was %v", err)
			}

			//
			//
			gr := page.AddGroup()
			gr.Cols = 12

			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.Width = "70%"
			gr.Style.Mobile.StyleBox.Width = "100%"

			// row-1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.ColSpanLabel = 12
				inp.Label = trl.S{
					"de": fmt.Sprintf(`
						<b>3.</b> &nbsp; 
							Den Hauptrefinanzierungssatz der EZB (derzeit %v) erwarte ich
							<br>
							 [zentrales 90%% Konfidenzintervall]:	
						`,
						latestECBRate),
					"en": fmt.Sprintf(`
						<b>3.</b> &nbsp;
							 I expect the main refinancing facility rate of the ECB (currently at %v) 
							 to lie 
							 <br>
							 [central 90%%  confidence interval]
							`,
						latestECBRate),
				}
			}

			// row-2

			type row struct {
				lbl trl.S
				inp string
			}
			rows := []row{
				{
					lbl: trl.S{
						"de": "in 6&nbsp;Monaten",
						"en": "in&nbsp;6&nbsp;months",
					},
					inp: "ezb6",
				},
				{
					lbl: trl.S{
						"de": "Ende&nbsp;2022",
						"en": "end of 2022",
					},
					inp: "ezb2022",
				},
				{
					lbl: trl.S{
						"de": "Ende&nbsp;2023",
						"en": "end of 2023",
					},
					inp: "ezb2023",
				},
				{
					lbl: trl.S{
						"de": "Ende&nbsp;2024",
						"en": "end of 2024",
					},
					inp: "ezb2024",
				},
			}

			for _, row := range rows {

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 3
					inp.Label = row.lbl
					inp.StyleLbl = lblStyleRight
				}

				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("%vmin", row.inp) //"i_ez_06_low"
					inp.Min = -10
					inp.Max = +20
					inp.Validator = "inRange20"
					inp.MaxChars = 5
					inp.Step = 0.01

					inp.ColSpan = 5
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 2
					inp.Label = trl.S{
						"de": "zwischen&nbsp;",
						"en": "between&nbsp;",
					}
					inp.Suffix = trl.S{"de": "%", "en": "pct"}
					inp.StyleLbl = lblStyleRight
				}

				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("%vmax", row.inp) //"i_ez_06_high"
					inp.Min = -10
					inp.Max = +20
					inp.Validator = "inRange20"
					inp.MaxChars = 5
					inp.Step = 0.01

					inp.ColSpan = 4
					inp.ColSpanLabel = 2
					inp.ColSpanControl = 2
					inp.Label = trl.S{
						"de": "und",
						"en": "and",
					}
					inp.Suffix = trl.S{"de": "%", "en": "pct"}
				}
			}

			//
			// row last
			/*
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 5
					inp.Label = trl.S{
						"de": " &nbsp;",
						"en": " &nbsp;",
					}
				}
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpan = 7
					inp.Label = trl.S{
						"de": "[zentrales 90% Konfidenzintervall]",
						"en": "[central 90%  confidence interval]",
					}
					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)

					inp.StyleLbl.Desktop.StyleBox.Position = "relative"

					inp.StyleLbl.Desktop.StyleBox.Left = "2rem"
					inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
					inp.StyleLbl.Mobile.StyleBox.Left = "0"
					inp.StyleLbl.Mobile.StyleBox.Top = "0"

					inp.StyleLbl.Desktop.StyleText.FontSize = 90

				}

			*/
		}

	} // special page 1

	//
	//
	return nil
}
