package fmt

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Seasonal questions revolve around the month of the Quarter.
//
// 	1 of quarter: Business cycle & drivers: 		         Januar, April, Juli, October
// 	2 of quarter: Inflation, drivers, central bank rates:    Februar, May, August, November
// 	3 of quarter: Free special questoins:                    March, June, September, December
func addSeasonal1(q *qst.QuestionnaireT) error {

	if monthOfQuarter() != 1 && false {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{
		"de": "Sonderfrage: Kurz- und mittelfristiges Wirtschaftswachstum",
		"en": "Special: Short and Medium Term Economic Growth",
	}
	page.Short = trl.S{"de": "Wachstum", "en": "Growth"}
	page.Style = css.DesktopWidthMax(page.Style, "42rem")

	{
		gr := page.AddGroup()
		gr.Cols = 12

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			// inp.ColSpanLabel = 6
			inp.Label = trl.S{
				"de": `<b>1.</b> 
				Punktprognose der Wachstumsrate des deutschen BIP: <br>
				<div class='vspacer-08' ></div>
				<p style='font-size: 90%'>
				Bei den Quartalwerten bitte nicht-annualisiertes Quartalswachstum 
						des realen & saisonbereinigten BIP angeben. 
				Bei den Jahreswerten die Jahreswachstumsrate des realen BIP.
				</p>
				`,
				"en": `<b>1.</b> 
				German GDP growth rate - point forecate: <br>
				<div class='vspacer-08' ></div>
				<p style='font-size: 90%'>
				For the quarterly values, please quote the non-annualized quarterly growth
						of the real & seasonal adjusted GDP.
				For the yearly values, please quote the yearly growth rate of the real GDP.
				</p>
				`,
			}
		}

		sLbl := css.NewStylesResponsive(nil)
		// inp.StyleLbl.Desktop.StyleGridItem.AlignSelf = "start"
		sLbl.Desktop.StyleGridItem.JustifySelf = "start"
		sLbl.Desktop.StyleGridItem.JustifySelf = "end"
		sLbl.Desktop.StyleText.AlignHorizontal = "right"

		sLbl.Mobile.StyleText.FontSize = 85

		sLbl.Desktop.StyleBox.Padding = "0 0.4rem 0 0"
		sLbl.Mobile.StyleBox.Padding = "0 0.1rem 0 0"

		/*
			Quarterly estimates.
			Quarterly results are published by Destatis six weeks after quarter ends. I.e. 15.May for Q1.
			We dont want estimates, if final results are already published.

			We are in first monthOfQuarter() == 1, i.e. April.

			Thus: Previous quarter, current, next ...

		*/
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "Prognose Wachstum des BIP je <b>Quartal</b> <br>\n (real, saisonbereinigt, nicht annualisiert)",
				"en": "Forecast <b>quarterly</b> GDP growth <br>\n(real, seasonally adjusted, non annualized)",
			}
			inp.Label = trl.S{
				"de": "Prognose <b>Quartal</b>",
				"en": "<b>Quarterly</b> forecast ",
			}
			inp.ColSpan = 12
			// inp.ColSpanLabel = 6
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xquart1"
			inp.ColSpan = 3
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextQ(0),
				"en": nextQ(0),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.StyleLbl = sLbl
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xquart2"
			inp.ColSpan = 3
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextQ(1),
				"en": nextQ(1),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.StyleLbl = sLbl
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xquart3"
			inp.ColSpan = 3
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextQ(2),
				"en": nextQ(2),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.StyleLbl = sLbl
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xquart4"
			inp.ColSpan = 3
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
			inp.MaxChars = 3
			// inp.Validator = "inRange20"
			inp.Label = trl.S{
				"de": nextQ(3),
				"en": nextQ(3),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.StyleLbl = sLbl
		}

		//
		// row 1
		//

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "Prognose Wachstum des BIP aufs&nbsp;<b>Jahr</b> <br>\n(real, saisonbereinigt)",
				"en": "Forecast GDP growth per&nbsp;<b>year</b> <br>\n(real, seasonally adjusted)",
			}
			inp.Label = trl.S{
				"de": "Prognose  <b>Jahr</b>",
				"en": "Forecast  <b>Year</b>",
			}
			inp.ColSpan = 12
			// inp.ColSpanLabel = 6
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xyear1"
			inp.ColSpan = 4 - 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
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
			inp.StyleLbl = sLbl
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xyear2"
			inp.ColSpan = 4 - 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
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
			inp.StyleLbl = sLbl
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xyear3"
			inp.ColSpan = 4 - 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = 0
			inp.Max = 20
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
			inp.StyleLbl = sLbl
		}
	}

	//
	//
	rowLabelsEconomicAreas := []trl.S{
		{
			"de": "Konjunkturdaten Deutschland",
			"en": "Business cycle data Germany",
		},
		{
			"de": "Exportmärkte",
			"en": "Export markets",
		},
		{
			"de": "Wechselkurse",
			"en": "Exchange rates",
		},
		{
			"de": "Int. Handelskonflikte",
			"en": "Intl. trade conflicts",
		},
		{
			"de": "EZB-Geldpolitik",
			"en": "ECB monetary policy",
		},
		{
			"de": "US-Geldpolitik",
			"en": "FED monetary policy",
		},
		{
			"de": "Brexit",
			"en": "Brexit",
		},
		{
			"de": "Corona Pandemie",
			"en": "Corona pandemic",
		},
		// {
		// 	// dummy for 'iobc_free' be replaced
		// 	"de": " &nbsp; ",
		// 	"en": " &nbsp; ",
		// },
	}

	// gr2
	// iobc => impact on business cycle
	{
		gb := qst.NewGridBuilderRadios(
			columnTemplate6,
			labelsStronglyPositiveStronglyNegativeInfluence(),
			[]string{"iobc_cycle_data_deu", "iobc_exp_markets", "iobc_exch_rates",
				"iobc_trade_conflicts",
				"iobc_mp_ecb", "iobc_mp_fed", "iobc_brexit", "iobc_corona",
				// "iobc_free",
			},
			radioVals6,
			rowLabelsEconomicAreas,
		)
		gb.MainLabel = trl.S{
			"de": "<b>2.</b> Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision Ihrer Konjunkturprognosen für die deutsche Wirtschaft bewogen?",
			"en": "<b>2.</b> Which developments have lead you to change your assessment of the business cycle outlook for the German economy?",
		}
		gr := page.AddGrid(gb)
		gr.OddRowsColoring = true
		gr.BottomVSpacers = 1
	}
	// gr3
	{
		gr := page.AddGroup()
		gr.Cols = 8
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.TemplateColumns =
			"1.92fr 1fr  1fr  1fr  1fr  1fr  0.35fr 1fr"
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "iobc_free_label"
			inp.MaxChars = 13 // more would break the col widthcase
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Label = nil
		}
		for i := 0; i < len(radioVals6); i++ {

			if i == 5 {
				txt := gr.AddInput()
				txt.Type = "textblock"
			}

			inp := gr.AddInput()
			inp.Type = "radio"
			inp.Name = "iobc_free"
			inp.ValueRadio = radioVals6[i]
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.Label = nil
		}
	}

	return nil

}
