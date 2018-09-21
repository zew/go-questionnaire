package fmt

import (
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

// Seasonal questions revolve around the month of the Quarter.
//
// 	1 of quarter: Business cycle & drivers: 		         Januar, April, Juli, October
// 	2 of quarter: Inflation, drivers, central bank rates:    Februar, May, August, November
// 	3 of quarter: Free special questoins:                    March, June, September, December
func addSeasonal1(q *qst.QuestionaireT) error {

	if monthOfQuarter() != 1 && false {
		return nil
	}

	p := q.AddPage()
	p.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	p.Label = trl.S{"de": "Prognosetreiber Wachstum", "en": "Growth drivers"}
	p.Short = trl.S{"de": "Wachstum", "en": "Growth"}
	p.Width = 80

	{
		gr := p.AddGroup()
		gr.Cols = 9
		gr.Label = trl.S{
			"de": "1.",
			"en": "1.",
		}
		gr.Desc = trl.S{
			"de": "Punktprognose der Wachstumsrate des deutschen BIP",
			"en": "Forecast growth rate German GDP",
		}

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
			inp.Desc = trl.S{
				"de": "Prognose Wachstum des BIP je <b>Quartal</b> <br>\n (real, saisonbereinigt, nicht annualisiert)",
				"en": "Forecast <b>quarterly</b> GDP growth <br>\n(real, seasonally adjusted, non annualized)",
			}
			inp.ColSpanLabel = 3
			inp.CSSLabel = "special-input-margin-vertical"
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xquart1"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
				"de": nextQ(-1),
				"en": nextQ(-1),
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
			inp.Name = "xquart2"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
				"de": nextQ(0),
				"en": nextQ(0),
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
			inp.Name = "xquart3"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
				"de": nextQ(),
				"en": nextQ(),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}
			inp.HAlignLabel = qst.HRight

		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": "Prognose Wachstum des BIP aufs&nbsp;<b>Jahr</b> <br>\n(real, saisonbereinigt)",
				"en": "Forecast GDP growth per&nbsp;<b>year</b> <br>\n(real, seasonally adjusted)",
			}
			inp.ColSpanLabel = 3

		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = "xyear1"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
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
			inp.Name = "xyear2"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
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
			inp.Name = "xyear3"
			inp.MaxChars = 4
			inp.Validator = "inRange20"
			inp.Desc = trl.S{
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

	// gr1
	{
		labels123Matrix := []trl.S{
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
				"de": "EZB-Geldpolitik",
				"en": "ECB monetary policy",
			},
			{
				"de": "US-Geldpolitik",
				"en": "FED monetary policy",
			},
			{
				"de": "Geopol. Ereignisse",
				"en": "Geopolitical events",
			},
			{
				"de": "Regierungsbildung Deutschland",
				"en": "Government formation Germany",
			},
			{
				"de": "Sonstige",
				"en": "Other",
			},
		}
		// iobc => impact on business cycle
		names1stMatrix := []string{
			"iobc_cycle_data_deu",
			"iobc_exp_markets",
			"iobc_exch_rates",
			"iobc_mp_ecb",
			"iobc_mp_fed",
			"iobc_geopol",
			"iobc_gvt_form_deu",
			"iobc_other",
		}
		gr := p.AddRadioMatrixGroup(labelsStronglyPositiveStronglyNegativeInfluence(),
			names1stMatrix, labels123Matrix, 2)
		gr.Cols = 8 // necessary, otherwise no vspacers
		gr.OddRowsColoring = true
		gr.Label = trl.S{
			"de": "2.",
			"en": "2.",
		}
		gr.Desc = trl.S{
			"de": "Haben Entwicklungen in den folgenden Bereichen Sie zu einer Revision Ihrer Konjunkturprognosen (ggü. Vormonat) für die deutsche Wirtschaft bewogen und wenn ja in welche Richtung?",
			"en": "Which developments have lead you to change your assessment of the business cycle outlook for the German economy compared to the previous month",
		}

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = "other_cycle_infl"
			inp.Desc = trl.S{
				"de": "Wenn sonstige - welche?",
				"en": "If other - which?",
			}
			inp.MaxChars = 30
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 6
			inp.CSSControl = "mobile-input-smaller"
		}

	}

	return nil

}
