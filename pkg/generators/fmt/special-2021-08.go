package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202108(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 8 {
		return nil
	}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	var columnTemplateLocal = []float32{
		3, 1,
		0, 1,
		0, 1,
		0, 1,
		0, 1,
		0.4, 1,
	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfragen zu Inflation, Geldpolitik, Prognosetreiber und zur neuen geldpolitischen Strategie der Europäischen Zentralbank (EZB) - Teil&nbsp;1",
			"en": "FMT Special Questions on Inflation, Monetary Policy, and the new Strategy of the European Central Bank (ECB) - Part 1",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>EZB 1",
			"en": "Special Questions&nbsp;1 - ECB",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

		//
		// gr 0 - einleitung
		{
			gr := page.AddGroup()
			gr.Cols = 2
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleBox.Width = "70%"
			// gr.Style.Mobile.StyleBox.Width = "100%"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 2
				inp.Desc = trl.S{
					"de": `

<div style='max-width: 36rem; margin-left: 1rem; margin-bottom: 0.7rem'>

<p style='font-weight: bold'>
    Am 8. Juli 2021 hat die EZB ihre neue geldpolitische Strategie veröffentlicht.
    Am 22. Juli 2021 hat die EZB zum ersten Mal geldpolitische Beschlüsse
    auf Basis der neuen Strategie gefasst.
</p>

<p>
    Die neue Strategie der EZB enthält folgende zentrale Elemente:
</p>

<ol>
	<li>
	das Inflationsziel von mittelfristig 2&nbsp;Prozent wird symmetrisch definiert. Negative Abweichungen vom Zielwert
	sind nun genauso unerwünscht wie positive,
	</li>

	<li>
	der HVPI (Harmonisierter Verbraucherpreisindex)
	soll schrittweise um selbst genutztes Wohneigentum erweitert werden,
	</li>

	<li>
	Klimaschutzaspekte sollen bei der Geldpolitik stärker berücksichtigt werden.
	</li>

</ol>

<p>
    Im Folgenden würden wir gerne von Ihnen erfahren, 
    welche Auswirkungen Sie von der neuen Strategie der EZB für Zinsen und Inflationsentwicklung erwarten.
</p>

</div>
					`,

					"en": `

<div style='max-width: 36rem; margin-left: 1rem; margin-bottom: 0.7rem'>


<p style='font-weight: bold'>
	On July 8, 2021, the ECB informed the public of its new monetary policy strategy. 
	On July 22, 2021, the ECB released 
	its first monetary policy decision based on this new strategy.
</p>

<p>
	The new strategy of the ECB, consists of three main elements. 
</p>

<ol>
	<li>
	The inflation target of 2 percent is now defined symmetrically. This means that negative deviations from the target value are equally undesirable as positive ones.
	</li>

	<li>
	The Harmonized Consumer Price Index (HCPI) will be augmented by owner-occupied housing. 
	</li>

	<li>
	Climate-policy aspects will have an increased weight in future in monetary policy decisions. 
	</li>

</ol>

<p>
	In the following, we would like to know how the new strategy of the ECB has affected your expectations of future interest rates and inflation. 
</p>

</div>

					`,
				}
			}

		}

		// gr 1
		{
			gr := page.AddGroup()
			gr.Cols = 9
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleBox.Width = "70%"
			// gr.Style.Mobile.StyleBox.Width = "100%"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 9
				// inp.ColSpanLabel = 12
				inp.Label = trl.S{
					"de": `<b>1.</b> 
					
					Welche jährlichen Inflationsraten erwarten Sie für den Euroraum 
					in den Jahren 2021, 2022 und 2023 
					(Punktprognose, prozentualer Anstieg des HVPI von Jan. bis Dez.; Erwartungswert)

					
				`,
					"en": `<b>1.</b> Forecast <b>yearly inflation rate in the Euro area</b>
				<br>
				HICP  increase from Jan to Dec; expected value
				`,
				}
			}

			for idx := range []int{0, 1, 2} {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("ppjinf_jp%v", idx)
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

		// gr 2
		{
			rowLabelsEconomicAreasShort := []trl.S{
				{
					"de": "Konjunkturdaten Euroraum",
					"en": "Business cycle data Euro area",
				},
				{
					"de": "Löhne Euroraum",
					"en": "Wages Euro area",
				},
				{
					"de": "Rohstoffpreise",
					"en": "Raw material prices",
				},
				{
					"de": "Wechselkurse",
					"en": "Exchange rates",
				},
				{
					"de": "EZB-Geldpolitik insgesamt",
					"en": "ECB´s monetary policy as a whole",
				},
				{
					"de": "Wechsel der EZB zu einer neuen geldpol. Strategie seit 08.07.2021",
					"en": "ECB´s change to a new monetary strategy since July 8, 2021",
				},
				{
					"de": "Internat. Handelskonflikte",
					"en": "Internat. trade conflicts",
				},
				{
					"de": "Brexit",
					"en": "Brexit",
				},
				{
					"de": "Corona Pandemie",
					"en": "Corona pandemic",
				},
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				labelsPlusPlusMinusMinus(),
				// prefix ioi_ => impact on inflation
				//   but we stick to rev_ => revision
				[]string{
					"rev_bus_cycle_ea",
					"rev_wages_ea",
					"rev_commodity_prices",
					"rev_exch_rates",
					"rev_mp_ecb",
					"rev_mp_ecb_change21",
					"rev_trade_conflicts",
					"rev_brexit",
					"rev_corona",
				},
				radioVals6,
				rowLabelsEconomicAreasShort,
			)
			gb.MainLabel = trl.S{
				"de": "<b>2.</b> Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision Ihrer Inflationsprognosen (ggü. Mai 2021) für den Euroraum bewogen und wenn ja in welche Richtung?",
				"en": "<b>2.</b> Which developments have lead you to change your assessment of the inflation outlook for the Euro are compared to the previous month",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.BottomVSpacers = 3

		}

		// gr 3
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 3
			gr.BottomVSpacers = 4
			gr.Cols = 9
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleBox.Width = "70%"
			// gr.Style.Mobile.StyleBox.Width = "100%"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 9
				// inp.ColSpanLabel = 12
				inp.Desc = trl.S{
					"de": `

<p>
	<b>3.</b>

    Der Wechsel zu einem symmetrischen Inflationsziel
    von 2&nbsp;Prozent hat mich zu folgenden Revisionen meiner Inflationsprognosen bewogen:
</p>

<p>
    Veränderung Ihrer Inflationsprognose (von Sonderfrage&nbsp;1)
    gegenüber der Prognose vor
    der Entscheidung der EZB am 08.&nbsp;Juli&nbsp;2021
</p>
					`,
					"en": `
<p>
	<b>3.</b>
	
	The change in the ECB´s strategy to a symmetric 2&nbsp;percent inflation target made me change my inflation forecasts
</p>

<p>
	How did your inflation forecast (special question 1) change after the ECB´s publication of a new strategy (i.e. after 8 July 2021)?
</p>
					
					`,
				}

			}

			for idx := range []int{0, 1, 2} {
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					// inp.Name = fmt.Sprintf("ezb_change_na_%v", q.Survey.YearStr(idx))
					inp.Label = trl.S{
						"de": q.Survey.YearStr(idx),
						"en": q.Survey.YearStr(idx),
					}
					inp.ColSpan = 3
					// inp.StyleLbl = lblStyleRight
				}
			}

			for idx := range []int{0, 1, 2} {

				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = fmt.Sprintf("ezb_inflation_chg_%v", q.Survey.YearStr(idx))
					inp.Min = -10
					inp.Max = +20
					inp.Validator = "inRange20"
					inp.MaxChars = 5
					inp.Step = 0.01
					// inp.Label = trl.S{
					// 	"de": q.Survey.YearStr(idx),
					// 	"en": q.Survey.YearStr(idx),
					// }
					inp.Suffix = trl.S{
						"de": "Prozent&shy;punkte",
						"en": "percentage pts",
					}

					inp.ColSpan = 3
					// inp.ColSpanLabel = 2
					inp.ColSpanControl = 2

					// inp.StyleLbl = lblStyleRight
				}
			}

			for idx := range []int{0, 1, 2} {
				{
					inp := gr.AddInput()
					inp.Type = "checkbox"
					inp.Name = fmt.Sprintf("ezb_inflation_chg_dk_%v", q.Survey.YearStr(idx))
					inp.Label = trl.S{
						"de": "weiß<br>nicht",
						"en": "dont<br>know",
					}
					inp.ColSpan = 3
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 4
					// inp.StyleLbl = lblStyleRight
					inp.ControlFirst()
					inp.LabelPadRight()
				}
			}

			for idx := range []int{0, 1, 2} {
				{
					inp := gr.AddInput()
					inp.Type = "checkbox"
					inp.Name = fmt.Sprintf("ezb_inflation_chg_na_%v", q.Survey.YearStr(idx))
					inp.Label = trl.S{
						"de": "keine<br>Antwort",
						"en": "no<br>answer",
					}
					inp.ColSpan = 3
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 4
					// inp.StyleLbl = lblStyleRight
					inp.ControlFirst()
					inp.LabelPadRight()
				}
			}

		}

	} // page 1 end

	// page 2
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfragen zu Inflation, Geldpolitik, Prognosetreiber und zur neuen geldpolitischen Strategie der Europäischen Zentralbank (EZB) - Teil&nbsp;2",
			"en": "FMT Special Questions on Inflation, Monetary Policy, and the new Strategy of the European Central Bank (ECB) - Part 2",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>EZB 2",
			"en": "Special Questions&nbsp;2 - ECB",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

		// gr1
		{
			// 2019	18 Sep. 0.00
			latestECBRate, err := q.Survey.Param("main_refinance_rate_ecb")
			if err != nil {
				return fmt.Errorf("Set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%%` as in `main refinance rate of the ECB (01.02.2018: 3.2%%)`; error was %v", err)
			}

			//
			//
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			gr.Cols = 12

			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleBox.Width = "70%"
			// gr.Style.Mobile.StyleBox.Width = "100%"

			// row-1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 12
				inp.ColSpanLabel = 12
				inp.Label = trl.S{
					"de": fmt.Sprintf("<b>4.</b> Den <b>Hauptrefinanzierungssatz</b> der EZB (seit %v) erwarte ich auf Sicht von", latestECBRate),
					"en": fmt.Sprintf("<b>4.</b> I expect the <b>main refinance rate</b> of the ECB (since %v) in", latestECBRate),
				}
			}

			// row-2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "6&nbsp;Monaten",
					"en": "6&nbsp;months",
				}
				inp.StyleLbl = lblStyleRight
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "ezb6min" //"i_ez_06_low"
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
				inp.Name = "ezb6max" //"i_ez_06_high"
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

			//
			// row-3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": " 24&nbsp;Monaten",
					"en": " 24&nbsp;months",
				}
				inp.StyleLbl = lblStyleRight
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "ezb24min" //"i_ez_24_low"
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
				inp.Name = "ezb24max" //"i_ez_24_high"
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

			//
			// row-4
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
					"en": "[central 90&nbsp;pct confidence interval]",
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

		lblsGr2and3 := []trl.S{
			{"de": "Prognosehorizont<br>in  6&nbsp;Monaten", "en": "Forecast horizon:<br>  6&nbsp;months"},
			{"de": "Prognosehorizont<br>in 24&nbsp;Monaten", "en": "Forecast horizon:<br> 24&nbsp;months"},
		}

		inpsGr2and3 := []string{"6m", "24mn"}

		// gr 2a
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleBox.Width = "70%"
			// gr.Style.Mobile.StyleBox.Width = "100%"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 6
				// inp.ColSpanLabel = 12
				inp.Desc = trl.S{
					"de": `

<p>
	<b>5.</b>

	Die Einführung des symmetrischen Inflationsziels von 2&nbsp;Prozent hat mich zu folgenden Revisionen meiner Prognosen des EZB-Hauptrefinanzierungssatzes bewogen:
</p>

<p>
	Veränderung der <b>Untergrenze</b> des zentralen 90%&nbsp;Konfidenzintervalls (Sonderfrage 4) gegenüber der Prognose vor der Entscheidung der EZB am 08.Juli 2021
</p>
					`,
					"en": `
<p>
	<b>5.</b>

	The change in the ECB´s strategy to a symmetric 2&nbsp;percent inflation target made me change my forecasts of the ECB´s main refinancing operations rate:
</p>

<p>
	How much did you change the <b>lower bound</b> of the central 90&nbsp;percent confidence interval (special question 4) after the ECB´s publication of a new strategy (i.e. after 8 July 2021)?
</p>
					
					`,
				}

			}

			for _, lbl := range lblsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": lbl.Tr("de"),
					"en": lbl.Tr("en"),
				}
				inp.ColSpan = 3
			}

			for _, inpName := range inpsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("ezb_rate_chg_lb_%v", inpName)
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.01
				inp.Suffix = trl.S{
					"de": "Prozent&shy;punkte",
					"en": "percentage pts",
				}
				inp.ColSpan = 3
				inp.ColSpanControl = 2
			}

			for _, inpName := range inpsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("ezb_rate_chg_lb_dk_%v", inpName)
				inp.Label = trl.S{
					"de": "weiß<br>nicht",
					"en": "dont<br>know",
				}
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.ControlFirst()
				inp.LabelPadRight()
			}

			for _, inpName := range inpsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("ezb_rate_chg_lb_na_%v", inpName)
				inp.Label = trl.S{
					"de": "keine<br>Antwort",
					"en": "no<br>answer",
				}
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.ControlFirst()
				inp.LabelPadRight()
			}

		}

		// gr 2b
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 3
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleBox.Width = "70%"
			// gr.Style.Mobile.StyleBox.Width = "100%"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 6
				// inp.ColSpanLabel = 12
				inp.Desc = trl.S{
					"de": `


<p>
	Veränderung der <b>Obergrenze</b> des zentralen 90%&nbsp;Konfidenzintervalls (Sonderfrage 4) gegenüber der Prognose vor der Entscheidung der EZB am 08.Juli 2021
</p>
					`,
					"en": `
<p>
	How much did you change the <b>upper bound</b> of the central 90&nbsp;percent confidence interval (special question 4) after the ECB´s publication of a new strategy (i.e. after 8 July 2021)?
</p>

					`,
				}

			}

			for _, lbl := range lblsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": lbl.Tr("de"),
					"en": lbl.Tr("en"),
				}
				inp.ColSpan = 3
			}

			for _, inpName := range inpsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("ezb_rate_chg_ub_%v", inpName)
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.01
				inp.Suffix = trl.S{
					"de": "Prozent&shy;punkte",
					"en": "percentage pts",
				}
				inp.ColSpan = 3
				inp.ColSpanControl = 2
			}

			for _, inpName := range inpsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("ezb_rate_chg_ub_dk_%v", inpName)
				inp.Label = trl.S{
					"de": "weiß<br>nicht",
					"en": "dont<br>know",
				}
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.ControlFirst()
				inp.LabelPadRight()
			}

			for _, inpName := range inpsGr2and3 {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("ezb_rate_chg_ub_na_%v", inpName)
				inp.Label = trl.S{
					"de": "keine<br>Antwort",
					"en": "no<br>answer",
				}
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 4
				inp.ControlFirst()
				inp.LabelPadRight()
			}

		}

		//
		//
		// gr 3
		{
			var columnTemplateLocal = []float32{
				0, 1,
				0, 1,
				0, 1,
				0, 1,
				0, 1,
				0.4, 1,
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				raiseDecrease6(),
				[]string{"inflation_chg_by_self_residential"},
				radioVals6,
				nil,
			)
			gb.MainLabel = trl.S{
				"de": `<b>6.</b> 
					Im Vergleich zur bisherigen Definition 
					wird die Berücksichtigung von 
					<b>selbst genutztem Wohneigentum im HVPI </b>
					die Inflationsrate im Zeitraum 2021-2023 
					folgendermaßen verändern:

					`,
				"en": `<b>6.</b> 
						The HCPI will be augmented by <b>owner-owned housing</b>. Which effect do you expect for inflation in the coming years 2021-2023?
					`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleGridContainer.GapColumn = "2.4rem"
			// gr.Style.Mobile.StyleGridContainer.GapColumn = "0.4rem"
		}

		//
		//
		// gr 4a
		var columnTemplate3plusDK = []float32{
			0, 1,
			0, 1,
			0, 1,
			0.4, 1,
		}
		{

			gb := qst.NewGridBuilderRadios(
				columnTemplate3plusDK,
				special202108A(),
				[]string{"climate_on_inflation_target"},
				radioVals4,
				nil,
			)
			gb.MainLabel = trl.S{
				"de": `
				<p>
					<b>7.</b> 
					Die Berücksichtigung von <b>Klimaschutzaspekten</b> in der geldpolitischen Strategie der EZB … 
				</p>

				<p>
					&nbsp;&nbsp; <b>a)</b>	wird die Umsetzung der Geldpolitik in Bezug auf das symmetrische 2-Prozent-Ziel 
				</p>


					`,
				"en": `
				<p>
					<b>7.</b> 
					The increased weight of <b>climate policy aspects</b> in the ECB´s monetary policy …  
				</p>

				<p>
					&nbsp;&nbsp; <b>a)</b>	will make the implementation of monetary policy with regard to the symmetric 2 percent target … 
				</p>
				
					`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleGridContainer.GapColumn = "2.4rem"
			// gr.Style.Mobile.StyleGridContainer.GapColumn = "0.4rem"
		}

		//
		// gr 4b
		{

			gb := qst.NewGridBuilderRadios(
				columnTemplate3plusDK,
				special202108B(),
				[]string{"climate_on_communication"},
				radioVals4,
				nil,
			)
			gb.MainLabel = trl.S{
				"de": `

				<p>
					&nbsp;&nbsp; <b>b)</b>	wird die Transparenz der Kommunikation der geldpolitischen Entscheidungen …
				</p>


					`,
				"en": ` 
				<p>
					&nbsp;&nbsp; <b>b)</b>	will make the transparency of ECB´s monetary policy communication …
				</p>
					`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleGridContainer.GapColumn = "2.4rem"
			// gr.Style.Mobile.StyleGridContainer.GapColumn = "0.4rem"
		}

		//
		// gr 4c
		{

			gb := qst.NewGridBuilderRadios(
				columnTemplate3plusDK,
				special202108C(),
				[]string{"climate_on_eu_targets"},
				radioVals4,
				nil,
			)
			gb.MainLabel = trl.S{
				"de": `

				<p>
					&nbsp;&nbsp; <b>c)</b>	wird die Erreichung der klimapolitischen Ziele der EU …
				</p>


					`,
				"en": ` 
				<p>
					&nbsp;&nbsp; <b>c)</b>	will make reaching the climate goals of the EU …  
				</p>
					`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = css.NewStylesResponsive(gr.Style)
			// gr.Style.Desktop.StyleGridContainer.GapColumn = "2.4rem"
			// gr.Style.Mobile.StyleGridContainer.GapColumn = "0.4rem"
		}

	} // page 2 end

	return nil

}
