package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Seasonal questions revolve around the month of the Quarter.
//
// 	1 of quarter: Business cycle & drivers: 		         Januar, April, Juli, October
// 	2 of quarter: Inflation, drivers, central bank rates:    Februar, May, August, November
// 	3 of quarter: Free special questoins:                    March, June, September, December
func eachMonth1inQ(q *qst.QuestionnaireT) error {

	if q.Survey.MonthOfQuarter() != 1 {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{
		"de": "Sonderfrage: Kurz- und mittelfristiges Wirtschaftswachstum",
		"en": "Special: Short and Medium Term Economic Growth",
	}
	page.Short = trl.S{
		"de": "Wachstum",
		"en": "Growth",
	}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "42rem")

	{
		gr := page.AddGroup()
		gr.Cols = 12

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `<b>1.</b> 
				Punktprognose der Wachstumsrate des deutschen BIP <br>
				<div class='vspacer-08' ></div>
				<p style='font-size: 90%'>
				Bei den Quartalwerten bitte nicht-annualisiertes Quartalswachstum 
						des realen & saisonbereinigten BIP angeben. 
				Bei den Jahreswerten die Jahreswachstumsrate des realen BIP.
				</p>
				`,
				"en": `<b>1.</b> 
				German GDP growth rate - point forecast <br>
				<div class='vspacer-08' ></div>
				<p style='font-size: 90%'>
				For the quarterly values, please quote the non-annualized quarterly growth
						of the real & seasonal adjusted GDP.
				For the yearly values, please quote the yearly growth rate of the real GDP.
				</p>
				`,
			}
		}

		sLbl1 := css.NewStylesResponsive(nil)
		sLbl1.Desktop.StyleGridItem.JustifySelf = "end"
		sLbl1.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
		sLbl1.Mobile.StyleBox.Padding = " 0 2.7rem 0 0.2rem"

		sLbl2 := *sLbl1
		sLbl2.Mobile.StyleGridItem.JustifySelf = "start"
		sLbl2.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
		sLbl2.Mobile.StyleBox.Padding = " 0 1.5rem 0 0.2rem"

		/*
			Quarterly estimates.
			Quarterly results are published by Destatis six weeks after quarter ends. I.e. 15.May for Q1.
			We dont want estimates, if final results are already published.

			We are in first MonthOfQuarter() == 1, i.e. April.

			Thus: Previous quarter, current, next ...

		*/
		// row 1 - four quarters - label
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "Prognose Wachstum des BIP je <b>Quartal</b> <br>\n (real, saisonbereinigt, nicht annualisiert)",
				"en": "Forecast <b>quarterly</b> GDP growth <br>\n(real, seasonally adjusted, non annualized)",
			}
			inp.Label = trl.S{
				"de": "Prognose <bx>Quartal</bx>",
				"en": "Forecast <bx>Quarter</bx>",
			}
			inp.ColSpan = 12

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Mobile.StyleBox.Position = "relative"
			inp.StyleLbl.Mobile.StyleBox.Top = "0.6rem"

		}
		// row 2 - four quarters - inputs
		for i := 0; i < 4; i++ {
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("pprwbipq%v", i+1)
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
				inp.Min = -20
				inp.Max = 20
				inp.Step = 0.01
				inp.MaxChars = 4
				inp.Label = trl.S{
					"de": q.Survey.Quarter(i),
					"en": q.Survey.Quarter(i),
				}
				inp.Suffix = trl.S{
					"de": "%",
					"en": "pct",
				}
				inp.StyleLbl = sLbl1

				inp.Style = css.MobileVertical(inp.Style)
				inp.StyleLbl.Mobile.StyleGridItem.JustifySelf = "start"
				// inp.StyleLbl.Mobile.StyleGridItem.AlignSelf = "end"
			}
		}

		// row 3 - three years - label
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "Prognose Wachstum des BIP aufs&nbsp;<bx>Jahr</bx> <br>\n(real, saisonbereinigt)",
				"en": "Forecast GDP growth per&nbsp;<bx>year</bx> <br>\n(real, seasonally adjusted)",
			}
			inp.Label = trl.S{
				"de": "Prognose  <bx>Jahr</bx>",
				"en": "Forecast  <bx>Year</bx>",
			}
			inp.ColSpan = 12

			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Mobile.StyleBox.Position = "relative"
			inp.StyleLbl.Mobile.StyleBox.Top = "0.6rem"
		}
		// row 4 - three years - inputs
		for i := 0; i < 3; i++ {
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("pprwbipj%v", i+1)
			inp.ColSpan = 4 - 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.Min = -20
			inp.Max = 20
			inp.Step = 0.01
			inp.MaxChars = 5
			inp.Label = trl.S{
				"de": q.Survey.YearStr(i),
				"en": q.Survey.YearStr(i),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.Style = css.MobileVertical(inp.Style)
			inp.StyleLbl = &sLbl2
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
	}

	// gr2
	{
		gb := qst.NewGridBuilderRadios(
			columnTemplate6,
			labelsStronglyPositiveStronglyNegativeInfluence(),
			// prefix iogf_ => impact on growth forecast
			//   but we stick to rev_ => revision
			[]string{
				"rev_bus_cycle_ger",
				"rev_exp_markets",
				"rev_exch_rates",
				"rev_trade_conflicts",
				"rev_mp_ecb",
				"rev_mp_fed",
				"rev_brexit",
				"rev_corona",
				// "rev_free",
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

	// gr4
	{
		gb := qst.NewGridBuilderRadios(
			columnTemplate6,
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

	return nil

}
