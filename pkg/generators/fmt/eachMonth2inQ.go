package fmt

import (
	"fmt"

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
		"de": "Sonderfrage: Inflation und Geldpolitik",
		"en": "Special: Inflation and monetary policy",
	}
	page.Short = trl.S{
		"de": "Inflation,<br>Geldpolitik",
		"en": "Inflation,<br>Mon. Policy",
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
					<b>1.</b> Punktprognose der <b>jährlichen Inflationsrate im Euroraum</b><br>
					Anstieg des HICP von Jan bis Dez; Erwartungswert
				`,
				"en": `
					Forecast of <b>annual inflation rate in the Euro area</b><br>
					Avg. percentage change in HICP from Jan to Dec;
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
				"de": "EZB-Geldpolitik",
				"en": "ECB monetary policy",
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
				"rev_trade_conflicts",
				"rev_brexit",
				"rev_corona",
			},
			radioVals6,
			rowLabelsEconomicAreasShort,
		)
		gb.MainLabel = trl.S{
			"de": "<b>2.</b> Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision Ihrer Inflationsprognosen (ggü. Vormonat) für den Euroraum bewogen und wenn ja in welche Richtung?",
			"en": "<b>2.</b> Which developments have lead you to change your assessment of the inflation outlook for the Euro are compared to the previous month",
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
		gr.BottomVSpacers = 1

	}

	// gr2a
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "rev_free_label"
			inp.MaxChars = 26
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Label = nil
			inp.Placeholder = trl.S{"de": "Sonstige", "en": "Other"}
		}
	}

	// gr2b
	{
		gb := qst.NewGridBuilderRadios(
			columnTemplateLocal,
			nil,
			[]string{"rev_free"},
			radioVals6,
			[]trl.S{
				{
					"de": " &nbsp;  ", // -
					"en": " &nbsp;  ", // -
				},
			},
		)
		gb.MainLabel = nil
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
	}

	// gr3
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
				"de": fmt.Sprintf("<b>3.</b> Den <b>Hauptrefinanzierungssatz</b> der EZB (seit %v) erwarte ich auf Sicht von", latestECBRate),
				"en": fmt.Sprintf("<b>3.</b> I expect the <b>main refinancing operations rate</b> of the ECB (since %v) in", latestECBRate),
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

	return nil

}
