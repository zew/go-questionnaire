package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

func addSeasonal2(q *qst.QuestionnaireT) error {

	if monthOfQuarter() != 2 && false {
		return nil
	}

	page := q.AddPage()
	page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{"de": "Prognosetreiber Inflation und Geldpolitik", "en": "Inflation and monetary policy drivers"}
	page.Short = trl.S{"de": "Sonderfrage:<br>Inflation,<br>Geldpolitik", "en": "Special:<br>Inflation,<br>Mon. Policy"}
	page.Style = css.DesktopWidthMax(page.Style, "60rem") // 90

	{
		gr := page.AddGroup()
		gr.Cols = 9

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 9
			inp.Label = trl.S{
				"de": "<b>1.</b> Punktprognose der jährlichen Inflationsrate im Euroraum",
				"en": "<b>1.</b> Forecast yearly inflation rate in the Euro area",
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "Anstieg des HICP von Jan bis Dez; Erwartungswert",
				"en": "HICP  increase from Jan to Dec; expected value",
			}
			inp.ColSpanLabel = 3
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "p1_y1"
			inp.ColSpanControl = 1
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextY(0),
				"en": nextY(0),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight

		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "p1_y2"
			inp.ColSpanControl = 1
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextY(1),
				"en": nextY(1),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "p1_y3"
			inp.ColSpanControl = 1
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextY(2),
				"en": nextY(2),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight

		}

	}

	rowLabelsEconomicAreasShort := []trl.S{
		{
			"de": "Konjunkturdaten Euroraum",
			"en": "Business cycle data Euro area",
		},
		{
			"de": "Konjunkturdaten global",
			"en": "Business cycle data globally",
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
	}

	// gr2
	// ioi => impact on inflation
	{
		gb := qst.NewGridBuilderRadios(
			columnTemplate6,
			labelsStronglyPositiveStronglyNegativeInfluence(),
			[]string{"ioi_cycle_data_ea", "ioi_cycle_data_gl", "ioi_wages_ea", "ioi_rmp", "ioi_exch_rates", "ioi_mp_ecb"},
			radioVals6,
			rowLabelsEconomicAreasShort,
		)
		gb.MainLabel = trl.S{
			"de": "<b>2.</b> Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision Ihrer Inflationsprognosen (ggü. Vormonat) für den Euroraum bewogen und wenn ja in welche Richtung?",
			"en": "<b>2.</b> Which developments have lead you to change your assessment of the inflation outlook for the Euro are compared to the previous month",
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true

	}

	// gr3
	{
		gr := page.AddGroup()
		gr.Cols = 100
		val, err := q.Survey.Param("main_refinance_rate_ecb") // 01.02.2018: 0,0
		if err != nil {
			return fmt.Errorf("Set field 'main_refinance_rate_ecb' to `01.02.2018: 3.2%%` as in `main refinance rate of the ECB (01.02.2018: 3.2%%)`; error was %v", err)
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 100
			inp.Label = trl.S{
				"de": fmt.Sprintf("<b>3.</b> Den Hauptrefinanzierungssatz der EZB (am %v) erwarte ich auf Sicht von", val),
				"en": fmt.Sprintf("<b>3.</b> I expect the main refinance rate of the ECB (%v) in", val),
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 8
			inp.Label = trl.S{
				"de": "6&nbsp;Monaten",
				"en": "6&nbsp;months",
			}
			inp.HAlignLabel = qst.HRight
			inp.CSSLabel = "special-min-width-85"
		}

		{

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "i_ez_06_low"
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"

			inp.ColSpanLabel = 7
			inp.CSSLabel = "special-line-height-higher"
			inp.ColSpanControl = 9
			inp.Label = trl.S{
				"de": "zwischen&nbsp;",
				"en": "between&nbsp;",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HRight
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "i_ez_06_high"
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"

			inp.ColSpanLabel = 3
			inp.ColSpanControl = 73
			inp.Label = trl.S{
				"de": "und",
				"en": "and",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 8
			inp.Label = trl.S{
				"de": " 24&nbsp;Monaten",
				"en": " 24&nbsp;months",
			}
			inp.HAlignLabel = qst.HRight
			inp.CSSLabel = "mobile-wider"

		}

		// Second row
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "i_ez_24_low"
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"

			inp.ColSpanLabel = 7
			inp.CSSLabel = "special-line-height-higher"
			inp.ColSpanControl = 9
			inp.Label = trl.S{
				"de": "zwischen&nbsp;",
				"en": "between&nbsp;",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HRight
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "i_ez_24_high"
			inp.Min = -10
			inp.Max = +20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"

			inp.ColSpanLabel = 3
			inp.ColSpanControl = 73
			inp.Label = trl.S{
				"de": "und",
				"en": "and",
			}
			inp.Suffix = trl.S{"de": "%", "en": "pct"}
			inp.HAlignLabel = qst.HLeft
			inp.HAlignControl = qst.HLeft
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 15
			inp.Label = trl.S{
				"de": " &nbsp;",
				"en": " &nbsp;",
			}
			inp.CSSLabel = "mobile-wider"
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpanLabel = 85
			inp.Label = trl.S{
				"de": "&nbsp; [zentrales 90% Konfidenzintervall]",
				"en": "&nbsp; [central 90&nbsp;pct confidence interval]",
			}
			inp.CSSLabel = "special-input-textblock-smaller"

		}

	}

	return nil

}
