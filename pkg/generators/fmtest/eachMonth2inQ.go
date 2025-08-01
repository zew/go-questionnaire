package fmtest

// exchanged from 1 - 2025-01

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Seasonal questions revolve around the month of the Quarter.
//
//	1 of quarter: Business cycle & drivers: 		         Januar, April, Juli, October
//	2 of quarter: Inflation, drivers, central bank rates:    Februar, May, August, November
//	3 of quarter: Free special questoins:                    March, June, September, December
func eachMonth2inQ(q *qst.QuestionnaireT) error {

	if q.Survey.MonthOfQuarter() != 2 {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{
		"de": "Sonderfragen: Kurz- und mittelfristiges Wirtschaftswachstum",
		"en": "Special: Short- and Medium-Term Economic Growth",
	}
	page.Short = trl.S{
		"de": "Wirtschafts-<br>wachstum",
		"en": "Economic<br>Growth",
	}
	page.WidthMax("42rem")

	if q.Survey.Year == 2025 && q.Survey.Month == 2 {
		gr := page.AddGroup()
		gr.Cols = 1
		gr.Style = css.NewStylesResponsive(gr.Style)
		// gr.Style.Desktop.StyleBox.Width = "70%"
		// gr.Style.Mobile.StyleBox.Width = "100%"

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					<i>Hinweis</i>:
					Beginnend mit der Januar-Umfragewelle wurde der Turnus der Sonderfragen angepasst. Im ersten Monat jedes Quartals werden nun Fragen zur Inflation im Euroraum gestellt und im zweiten Monat jedes Quartals Fragen zum Wirtschaftswachstum in Deutschland.
				`,
				"en": `
					<i>Note</i>:
					Starting with the January survey wave, the schedule for the special questions has been adjusted. In the first month of each quarter, questions are asked about inflation in the euro area and in the second month of each quarter questions about economic growth in Germany.
				`,
			}
		}
	}

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
				Bei den Quartalswerten bitte nicht-annualisiertes Quartalswachstum
						des realen & saisonbereinigten BIP angeben.
				Bei den Jahreswerten die Jahreswachstumsrate des realen BIP.
				</p>
				`,
				"en": `<b>1.</b>
				Point forecast of the growth rate of the <b>German GDP</b> <br>
				<div class='vspacer-08' ></div>
				<p style='font-size: 90%'>
				For the quarterly values,
				please indicate non-annualized quarterly
				real & seasonally adjusted GDP growth.
				For the yearly values,
				please indicate the annual real GDP growth rate.
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
			We      want *estimates*,  if final results are not     published yet.
			Put another way
			We dont want *estimates*,  if final results are already published.

			We are in first MonthOfQuarter() == 1, i.e. April.

			Thus: Previous quarter, current, next ...

		*/

		offsetDestatis := 0 // next quarter
		if osd, err := q.Survey.Param("destatis"); err == nil {
			offsetDestatis, _ = strconv.Atoi(osd)
		}

		// row 1 - four quarters - label
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
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

				// append asterisk
				if offsetDestatis != 0 {
					if i == 0 {
						inp.Label = trl.S{
							"de": q.Survey.Quarter(i) + "*",
							"en": q.Survey.Quarter(i) + "*",
						}
					}
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

		// destatis correction for the year
		yrQuarterAndYearCorrected := strings.Split(q.Survey.Quarter(0), "&nbsp;")
		yearCorrectedS := yrQuarterAndYearCorrected[1]
		yearCorrected, err := strconv.Atoi(yearCorrectedS)
		if err != nil {
			return err
		}

		var lblFootnote = trl.S{
			"de": fmt.Sprintf("&nbsp;"),
			"en": fmt.Sprintf("&nbsp;"),
		}

		if offsetDestatis != 0 {
			lblFootnote = trl.S{
				"de": fmt.Sprintf("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<superscript>*</superscript><span style='font-size:80%%'> Die realisierten Zahlen für %v werden erst <a  target='_blank'  href='https://www.destatis.de/SiteGlobals/Forms/Suche/Termine/DE/Terminsuche_Formular.html' >später</a> veröffentlicht.<span>", q.Survey.Quarter(0)),
				"en": fmt.Sprintf("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<superscript>*</superscript><span style='font-size:80%%'> Realized numbers for %v are only published <a  target='_blank'  href='https://www.destatis.de/SiteGlobals/Forms/Suche/Termine/DE/Terminsuche_Formular.html' >later</a>.<span>", q.Survey.Quarter(0)),
			}
		}

		// row 2a quarter explanation / footnote
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblFootnote
			inp.ColSpan = 12
			inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
			inp.StyleLbl.Mobile.StyleBox.Position = "relative"
			inp.StyleLbl.Mobile.StyleBox.Top = "0.6rem"
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
				// "de": q.Survey.YearStr(i),
				// "en": q.Survey.YearStr(i),
				"de": fmt.Sprint(yearCorrected + i),
				"en": fmt.Sprint(yearCorrected + i),
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
	// gr2
	rowLabelsEconomicAreas := []trl.S{
		{
			// "de": "Konjunkturdaten Deutschland",
			"de": "Konjunkturentwicklung Deutschland",
			"en": "Economic development<br>Germany",
		},
		{
			"de": "Exportmärkte",
			"en": "Export markets",
		},
		{
			"de": "Wechselkurse",
			// "en": "Exchange rates (relative to the Euro)",
			"en": "Exchange rates",
		},
		{
			"de": "Int. Handelskonflikte",
			"en": "Intl. trade conflicts",
			// "en": "Intl. trade disputes",
		},
		{
			"de": "EZB-Geldpolitik",
			"en": "ECB monetary policy",
		},
		{
			"de": "US-Geldpolitik",
			"en": "FED monetary policy",
			// "en": "US monetary policy",
		},
		// {
		// 	// https://www.duden.de/rechtschreibung/Coronapandemie
		// 	"de": "Corona Pandemie",
		// 	"en": "Corona pandemic",
		// },
		{
			"de": "Internationale Lieferengpässe",
			// "en": "International supply chain bottlenecks",
			"en": "International supply bottlenecks",
			// "en": "Supply chain disruptions",
			// -	Supply chain bottlenecks
		},
		{
			"de": "Energiepreise",
			"en": "Energy prices",
		},
		{
			"de": "Engpässe bei Rohstoffen",
			"en": "Raw material shortages",
			// "en": "Raw material bottlenecks",
		},
		{
			"de": "Inflation (ohne Energiepreise)",
			"en": "Inflation (excl. energy prices)",
		},
		{
			"de": "Krieg in der Ukraine",
			"en": "War in Ukraine",
		},
		// 2025-05
		{
			"de": "Handels&shy;protektionismus und Zölle seitens USA",
			"en": "US trade protectionism/ tariffs",
		},
		{
			"de": "Ankündigungen zu Verteidigungs- und Staats&shy;ausgaben",
			"en": "Announcements from the German government on military and fiscal spending",
		},
		// {
		// 	"de": "Spannungen im Bankensystem",
		// 	"en": "Tensions in the banking sector",
		// },
		// {
		// 	"de": "Die Schuldenbremse",
		// 	"en": "The debt brake",
		// },
	}

	colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	{
		gb := qst.NewGridBuilderRadios(
			colTemplate,
			improvedDeterioratedPlusMinus6(),
			// prefix iogf_ => impact on growth forecast
			//   but we stick to rev_ => revision
			[]string{
				"rev_bus_cycle_ger",
				"rev_exp_markets",
				"rev_exch_rates",
				"rev_trade_conflicts",
				"rev_mp_ecb",
				"rev_mp_fed",
				// "rev_corona",  // sunsetted 2023-07
				"rev_supply_disrupt",
				"rev_energy_prices",
				"rev_shortages_raw_mat",
				"rev_inflation",
				"rev_ukraine",
				"rev_us_tariffs",
				"rev_defence_spending",
				// "rev_banking_strain",
				// "rev_debt_brake",
				// "rev_free",
			},
			radioVals6,
			rowLabelsEconomicAreas,
		)

		// not 6 as in m3 of q
		monthsBack := 3

		if q.Survey.Year == 2025 && q.Survey.Month == 2 {
			monthsBack = 4
		}

		oneMonthPrev := time.Date(
			q.Survey.Year, time.Month(q.Survey.Month), 2,
			0, 0, 0, 0, cfg.Get().Loc,
		)
		oneMonthPrev = oneMonthPrev.Local().AddDate(0, -monthsBack, 0)
		month := int(oneMonthPrev.Month())

		gb.MainLabel = trl.S{
			"de": fmt.Sprintf(`<b>2.</b>
					Haben Entwicklungen in den folgenden Bereichen Sie
					zu einer Revision
					(ggü. <i>%v %v</i>)
					Ihrer Konjunkturprognosen
					für die deutsche Wirtschaft bewogen
					und wenn ja in welche Richtung?
					<br>
					(Erhöhung (+), Senkung (-))
			`, trl.MonthByInt(month)["de"], oneMonthPrev.Year(),
			),
			"en": fmt.Sprintf(`<b>2.</b>
					Which developments have led you to change
					(relative to <i>%v %v</i>)
					your assessment
					of the business cycle outlook for the German economy?
					<br>
					If they made you change your assessment,
					did they make you revise your assessment up (+) or down (-)?
					`, trl.MonthByInt(month)["en"], oneMonthPrev.Year(),
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
			rad.Name = "rev_growth_free"
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1
			rad.ColSpanLabel = colsRowFree[2*(idx+1)]
			rad.ColSpanControl = colsRowFree[2*(idx+1)] + 1
		}

	}

	return nil

}
