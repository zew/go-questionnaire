package fmt

// todo: inflation fields for JS dynamic - see JSBlockStrings
// 			see inflationRange.js

import (
	"fmt"
	"time"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func eachMonth2inQ(q *qst.QuestionnaireT) error {

	if q.Survey.MonthOfQuarter() != 2 {
		return nil
	}

	if q.Survey.Year == 2021 && q.Survey.Month == 8 {
		return nil
	}

	if q.Survey.Year == 2021 && q.Survey.Month == 11 {
		return nil
	}

	if q.Survey.Year == 2022 && q.Survey.Month == 2 {
		return nil
	}

	if q.Survey.Year == 2022 && q.Survey.Month == 5 {
		return nil
	}

	// not 6 as in m3 of q
	monthsBack := 3

	idxThreeMonthsBefore := trl.MonthsShift(int(q.Survey.Month), -monthsBack)
	monthMinus3 := trl.MonthByInt(idxThreeMonthsBefore)

	loc := time.Now().Location()
	yearMinus1Q := time.Date(q.Survey.Year, time.Month(q.Survey.Month), 2, 0, 0, 0, 0, loc)
	yearMinus1Q = yearMinus1Q.Local().AddDate(0, -monthsBack, 0)

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	/*
		SELECT
			frage_kurz,
			GROUP_CONCAT( DISTINCT antwort ORDER BY antwort) aw,
			count(*) anz
		FROM sonderfragen_ger
		WHERE survey_id = 202005
		GROUP BY frage_kurz
	*/

	page := q.AddPage()
	// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{
		"de": "Sonderfrage: Inflation, Inflationstreiber und Geldpolitik",
		"en": "Special Questions: Inflation, its causes, and monetary policy ",
	}
	page.Short = trl.S{
		"de": "Inflation,<br>Geldpolitik",
		"en": "Inflation,<br>Monetary Policy",
	}
	page.WidthMax("48rem")

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
					<b>1a.</b> &nbsp; Punktprognose der <b>jährlichen Inflationsrate im Euroraum</b>
					<br>
					(durchschnittliche jährliche Veränderung des HICP in Prozent)
					<!-- Anstieg des HICP von Jan bis Dez; Erwartungswert -->
				`,
				"en": `
					<b>1a.</b> &nbsp; Point forecast of the <b>annual inflation rate in the euro area</b>
					<br>
					(annual average change of the HICP, in percent)
					<!-- Avg. percentage change in HICP from Jan to Dec; -->
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
	//
	//
	//
	//
	// gr2
	{

		// colspan := float32(2 + 4*3 + 2 + 2)

		gr := page.AddGroup()
		gr.Cols = 6
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Mobile.StyleGridContainer.GapRow = "0.02rem"
		gr.ColWidths("1.6fr    2.7fr 3.1fr 3.1fr 3.4fr 2.4fr    2.4fr  1.4fr")

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 8
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
						How likely are specific future realizations of inflation? 
						
						Please give us your assessments for the annual average inflation rate 
						in the euro area.
						
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
				"de": "unter <br>4&nbsp;Prozent",
				"en": "below <br>4&nbsp;percent",
			},
			{
				"de": "zwischen 4 u.  <br>6&nbsp;Prozent",
				"en": "between  4 and <br>6&nbsp;percent",
			},
			{
				"de": "zwischen 6 u.  <br>8&nbsp;Prozent",
				"en": "between  6 and <br>8&nbsp;percent",
			},
			{
				"de": "zwischen 8 u.  <br>10&nbsp;Prozent",
				"en": "between  8 and <br>10&nbsp;percent",
			},
			{
				"de": "größer als 10&nbsp;Prozent",
				"en": "above  <br>10&nbsp;percent",
			},
		}

		// first row - cols 2-5
		for _, lbl := range labels {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = lbl
			inp.Style = css.ItemStartCA(inp.Style)
			inp.Style.Desktop.StyleGridItem.AlignSelf = "end"
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
			inp.Style.Desktop.StyleGridItem.AlignSelf = "end"
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
			inp.Style.Desktop.StyleGridItem.AlignSelf = "end"

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Desktop.StyleText.LineHeight = 118
		}

		//
		//
		// second to fourth row: inputs
		for i := q.Survey.Year; i <= q.Survey.Year+2; i++ {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": fmt.Sprint(i),
					"en": fmt.Sprint(i),
				}
			}

			// inpNames := []string{
			// 	"under2", "between2and3", "between3and4", "above4",
			// }
			inpNames := []string{
				"under4", "between4and6", "between6and8", "between8and10", "above10",
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
					"de": "100&nbsp;%",
					"en": "100&nbsp;%",
				}
				inp.Style = css.ItemStartMA(inp.Style)
				inp.Style = css.TextStart(inp.Style)
			}
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.ColSpan = 1
				inp.Name = fmt.Sprintf("inf%v__noanswer", i)
				inp.ColSpanControl = 1
				inp.Style = css.ItemStartMA(inp.Style)
			}

		}

		//
		//
		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "inflationRange"

			s1 := trl.S{
				"de": "Ihre Antworten auf Frage 1b addieren sich nicht zu 100%. Wirklich weiter?",
				"en": "Your answers to question 1b dont add up to 100%. Continue anyway?",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}

			// inp.JSBlockStrings = map[string]string{}
			// inp.JSBlockStrings["inputBaseName"] = name
			// for idx1 := 0; idx1 < 3; idx1++ {
			// 	key := fmt.Sprintf("%v_%v", "inp", idx1+1) // {{.inp_1}}, {{.inp_2}}, ...
			// 	inp.JSBlockStrings[key] = fmt.Sprintf("%v_prio%v", name, idx1)
			// }

		}

	}

	// gr2
	colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	{
		rowLabelsEconomicAreasShort := []trl.S{
			{
				"de": "Konjunkturentwicklung im Eurogebiet",
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
				"en": "Development of prices for raw materials (except energy) ",
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
				"en": "Covid pandemic",
			},
			{
				"de": "Grüne Transformation",
				"en": "Green transformation",
			},
			{
				"de": "Krieg in der Ukraine",
				"en": "War in Ukraine",
			},
		}

		gb := qst.NewGridBuilderRadios(
			colTemplate,
			labelsPlusPlusMinusMinus(),
			// prefix ioi_ => impact on inflation
			//   but we stick to rev_ => revision
			[]string{
				"rev_bus_cycle",
				"rev_wages",
				"rev_energy_prices",
				"rev_commodity_prices",
				"rev_exch_rates",
				"rev_mp_ecb",
				"rev_trade_conflicts",
				"rev_supply_shortages",
				"rev_corona",
				"rev_green_trafo",
				"rev_war_ukraine",
			},
			radioVals6,
			rowLabelsEconomicAreasShort,
		)

		gb.MainLabel = trl.S{
			"de": fmt.Sprintf(`<b>2.</b> &nbsp; 
				Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision 
				Ihrer <b>Inflationsprognosen</b> für den Euroraum (ggü. %v %v) bewogen 
				und wenn ja, nach oben (+) oder unten (-)?
				<br>
				<br>
				<b>Für die Jahre %d, %d und %d</b>
			`,
				monthMinus3.Tr("de"), yearMinus1Q.Year(),
				q.Survey.Year+0, q.Survey.Year+1, q.Survey.Year+2,
			),
			"en": fmt.Sprintf(`<b>2.</b> &nbsp;
				What are the main factors leading you to change your inflation forecasts
				for the euro area (in comparison to your forecasts as of %v %v).
				(+) means increase in inflation forecast,
				(-) means decrease in inflation forecast.
				<br>
				<br>
				<b>For the years %d, %d, and %d</b>
			`,
				monthMinus3.Tr("en"), yearMinus1Q.Year(),
				q.Survey.Year+0, q.Survey.Year+1, q.Survey.Year+2,
			),
		}
		gr := page.AddGrid(gb)
		gr.BottomVSpacers = 1

	}
	{

		//
		// row free input
		gr := page.AddGroup()
		gr.Cols = float32(len(improvedDeterioratedPlusMinus6()) + 1)
		gr.Cols = 7

		gr.Style = css.NewStylesResponsive(gr.Style)
		if gr.Style.Desktop.StyleGridContainer.TemplateColumns == "" {
			gr.Style.Desktop.StyleBox.Display = "grid"
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = styleRowFree
			// log.Printf("fmt special 2021-09: grid template - %v", stl)
		} else {
			return fmt.Errorf("GridBuilder.AddGrid() - another TemplateColumns already present.\nwnt%v\ngot%v", styleRowFree, gr.Style.Desktop.StyleGridContainer.TemplateColumns)
		}

		gr.BottomVSpacers = 4

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "rev_free_label"
			// inp.MaxChars = 17
			inp.MaxChars = 15
			inp.ColSpan = 1
			inp.ColSpanLabel = 2.4
			inp.ColSpanControl = 4
			inp.Label = trl.S{
				"de": "Andere",
				"en": "Other",
			}
		}

		//
		for idx := 0; idx < len(improvedDeterioratedPlusMinus6()); idx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "rev_free"
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1
			rad.ColSpanLabel = colsRowFree[2*(idx+1)]
			rad.ColSpanControl = colsRowFree[2*(idx+1)] + 1
		}

	}

	// gr3
	{
		latestECBRate, err := q.Survey.Param("main_refinance_rate_ecb")
		// www.euribor-rates.eu/en/ecb-refinancing-rate/
		// www.euribor-rates.eu/en/ecb-refinancing-rate/
		if err != nil {
			return fmt.Errorf("set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%%` as in `main refinancing operations rate of the ECB (01.02.2018: 3.2%%)`; error was %v", err)
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
				"de": fmt.Sprintf(
					`<b>3.</b> &nbsp; 
					Den <b>Hauptrefinanzierungssatz</b> der EZB (derzeit %v) erwarte ich 
					`, latestECBRate,
				),
				"en": fmt.Sprintf(
					`<b>3.</b> &nbsp; 
					I expect the <b>main refinancing facility rate</b> of the ECB (currently at %v) to be 
					`, latestECBRate,
				),
			}
		}

		lbls := []trl.S{
			{
				"de": "in 6&nbsp;Monaten",
				"en": "in 6&nbsp;months",
			},
			{
				"de": fmt.Sprintf("Ende   %v", q.Survey.Year+0),
				"en": fmt.Sprintf("End of %v", q.Survey.Year+0),
			},
			{
				"de": fmt.Sprintf("Ende   %v", q.Survey.Year+1),
				"en": fmt.Sprintf("End of %v", q.Survey.Year+1),
			},
			{
				"de": fmt.Sprintf("Ende   %v", q.Survey.Year+2),
				"en": fmt.Sprintf("End of %v", q.Survey.Year+2),
			},
		}

		inputs := []string{
			"ezb6",
			fmt.Sprintf("ezb%d", q.Survey.Year+0),
			fmt.Sprintf("ezb%d", q.Survey.Year+1),
			fmt.Sprintf("ezb%d", q.Survey.Year+2),
		}

		// rows 2...5
		for i := 0; i < 4; i++ {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = lbls[i]
				inp.StyleLbl = lblStyleRight
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%vmin", inputs[i]) // "ezb6min"
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
				inp.Name = fmt.Sprintf("%vmax", inputs[i]) // "ezb6max"
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
		// row-6
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
				"en": "[central 90% confidence interval]",
			}
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)

			inp.StyleLbl.Desktop.StyleBox.Position = "relative"

			inp.StyleLbl.Desktop.StyleBox.Left = "2rem"
			inp.StyleLbl.Desktop.StyleBox.Top = "-0.2rem"
			inp.StyleLbl.Mobile.StyleBox.Left = "0"
			inp.StyleLbl.Mobile.StyleBox.Top = "0"

			inp.StyleLbl.Desktop.StyleText.FontSize = 90

		}

	}

	if q.Survey.Year == 2022 && q.Survey.Month == 11 {

		{
			gr := page.AddGroup()
			gr.Cols = 14

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 14
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": `<b>4.</b> &nbsp; 
					Mit Blick auf das Jahr 2023, wie beeinflusst die aktuelle Entwicklung der Inflation Ihre Beurteilung des Rendite‐Risiko‐Profils des DAX?
				`,
					"en": `<b>4.</b> &nbsp; 
					How do current developments of inflation affect your assessment of the return-risk-profile of the DAX for the year 2023?
				`,
				}
			}

			lbls := labelsPositiveNeutralNegative()

			{
				for idx2 := 0; idx2 < len(lbls); idx2++ {
					inp := gr.AddInput()
					inp.Type = "radio"
					inp.Name = fmt.Sprintf("%v", "spec_4")
					inp.ValueRadio = fmt.Sprintf("%v", idx2+1) // row idx1
					inp.Label = lbls[idx2]
					inp.ColSpan = 2
					inp.ColSpanControl = 1
					inp.Vertical()
					inp.VerticalLabel()

					if idx2 == len(lbls)-1 {
						inp.ColSpan = 4
					}

				}

			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": `<b>5.</b> &nbsp; 
					Beschreiben Sie kurz in ganzen Sätzen über welche Mechanismen die Inflation Ihre Rendite- und Risiko-Erwartungen für den DAX in 2023 beeinflusst bzw. warum Sie keinen Zusammenhang sehen.
				`,
					"en": `<b>5.</b> &nbsp; 
					Please describe briefly in whole sentences via which mechanisms inflation affects your return-risk-expectations of the DAX for the year 2023 or why you see no relationship.
				`,
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "spec_5"
				inp.MaxChars = 300
				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

	}

	return nil

}
