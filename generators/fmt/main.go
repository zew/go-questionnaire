package fmt

import (
	"fmt"
	"log"
	"strings"

	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/go-questionnaire/trl"
)

var radioVals4 = []string{"1", "2", "3", "4"}
var radioVals6 = []string{"1", "2", "3", "4", "5", "6"}
var columnTemplate4 = []int{
	2, 1,
	0, 1,
	0, 1,
	1, 1,
}
var columnTemplate6 = []int{
	2, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	1, 1,
}

var rowLabelsSmallLargeEnterprises = []trl.S{
	{
		"de": "Großunternehmen",
		"en": "Large enterprises",
	},
	{
		"de": "KMU",
		"en": "Small+medium enterprises",
	},
	{
		"de": "Immobilienkredite",
		"en": "Real estate credit",
	},
	{
		"de": "Konsumentenkredite",
		"en": "Consumer credit",
	},
}

// Create creates a JSON file for a financial markets survey
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	// qst.RadioVali = "mustRadioGroup"
	qst.RadioVali = ""
	qst.HeaderClass = ""
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("fmt")
	q.Survey.Params = params
	q.LangCodes = []string{"de", "en"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{"de": "Finanzmarkttest", "en": "Financial Markets Survey"}

	q.Version = 2

	// page 0
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Begrüßung", "en": "Greeting"}
		page.NoNavigation = true
		page.Width = 60

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
				inp.ColSpanLabel = 3
				impr := trl.S{}
				for _, lc := range q.LangCodes {
					w1 := &strings.Builder{}
					err := tpl.RenderStaticContent(
						w1, "data-protection.md", q.Survey.Type, lc,
					)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = w1.String()

				}
				inp.Desc = impr
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 4
			// gr.Width = 75  // squeeze on mobile
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Sind Sie die angeschriebene Person?", "en": "Are you the addressee?"}

				inp.ColSpanLabel = 3
			}

			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "proxy"
				rad.ValueRadio = "no"
				rad.ColSpanLabel = 3
				rad.ColSpanControl = 1
				// rad.HAlign = qst.HLeft
				// rad.HAlign = qst.HCenter
				rad.Label = trl.S{
					"de": " Ja, ich bin die angeschriebene Person.",
					"en": "Yes, I am the addressee.",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "proxy"
				rad.ValueRadio = "yes"
				rad.ColSpanLabel = 3
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Nein, ich fülle den Fragebogen in Vertretung der angeschriebenen Person aus.",
					"en": "No. I am filling in for the addressee.",
				}
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 3
			// gr.Width = 75
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "Meine Adresse hat sich geändert",
					"en": "My address has changed",
				}
				inp.ColSpanLabel = 3
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "address_change"
				inp.MaxChars = 150
				inp.ColSpanControl = 3
			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.Width = 75
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = "1"
				inp.Label = trl.S{
					"de": "Weiter",
					"en": "Next",
				}
				inp.AccessKey = "n"
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HRight
			}
		}

	}

	// page 1
	{
		page := q.AddPage()
		page.Section = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		page.Label = trl.S{"de": "Status und Ausblick", "en": "Status and outlook"}
		page.Short = trl.S{"de": "Konjunktur:<br/>Status,<br/>Ausblick", "en": "Business cycle:<br/>Status,<br/>Outlook"}
		page.Short = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		page.Width = 60

		//
		//
		rowLabelsEuroGerUSGlob := []trl.S{
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "Deutschland",
				"en": "Germany",
			},
			{
				"de": "USA",
				"en": "US",
			},
			{
				"de": "Weltwirtschaft",
				"en": "Global economy",
			},
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsGoodBad(),
				[]string{"y0_ez", "y0_deu", "y0_usa", "y0_glob"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>1.</b> Die gesamtwirtschaftliche Situation beurteilen wir als",
				"en": "<b>1.</b> We assess the overall economic situation as",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr2a
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsImproveDeteriorate(),
				[]string{"y_ez", "y_deu", "y_usa", "y_glob"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>2a.</b> Die gesamtwirtschaftliche Situation wird sich mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>2a.</b> The overall economic situation medium term (<b>6</b>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr2b
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsImproveDeteriorate(),
				[]string{"y24_ez", "y24_deu", "y24_usa", "y24_glob"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>2b.</b> Die gesamtwirtschaftliche Situation wird sich langfristig (<b>24</b>&nbsp;Mo.)",
				"en": "<b>2b.</b> The overall economic situation long term (<b>24</b>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	}

	//
	// page 2
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Wachstum", "en": "Growth"}
		page.Short = trl.S{"de": "Wachstum", "en": "Growth"}
		page.Width = 60

		{
			gr := page.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 5
				inp.Desc = trl.S{
					"de": "<b>3a.</b> ",
					"en": "<b>3a.</b> ",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_q_deu"
				inp.Min = 0
				inp.Max = 20
				inp.MaxChars = 2
				inp.MaxChars = 3
				// inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das <b>deutsche</b> BIP Wachstum in %v <xxbr/>\n(real, saisonbereinigt, nicht annualisiert)", nextQ()),
					"en": fmt.Sprintf("Our estimate for the <b>German</b> GDP growth in %v        <xxbr/>\n(real, seasonally adjusted, non annualized)", nextQ()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
				inp.ColSpanControl = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_y_deu"
				inp.Min = 0
				inp.Max = 50
				inp.MaxChars = 3
				// inp.Validator = "inRange20"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das BIP Wachstum für das Jahr %v <xxbr/>\n(real, saisonbereinigt)", nextY()),
					"en": fmt.Sprintf("Our estimate for the GDP growth in %v                <xxbr/>\n(real, seasonally adjusted)", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
				inp.ColSpanControl = 1
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 5
				inp.Desc = trl.S{
					"de": "<b>3b.</b> ",
					"en": "<b>3b.</b> ",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "yshr_q_deu"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
				// inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Die Wahrscheinlichkeit eines negativen Wachstums des <b>deutschen</b> BIP in %v liegt bei", nextQ()),
					"en": fmt.Sprintf("The probability of negative growth for the <b>German</b> GDP in %v is", nextQ()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
				inp.ColSpanControl = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "yshr_y_deu"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
				// inp.Validator = "inRange100"

				inp.ColSpanLabel = 4
				inp.Desc = trl.S{
					"de": fmt.Sprintf("Die Wahrscheinlichkeit einer Rezession in Deutschland <xxbr/>\n(mind. 2&nbsp;Quartale neg. Wachstum) bis Q4 %v liegt bei", nextY()),
					"en": fmt.Sprintf("The probability of a recession in Germany             <xxbr/>\n(at least 2&nbsp;quarters neg. growth) until Q4 %v is", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
				inp.ColSpanControl = 1
			}

		}

	}

	//
	// page 3 - inflation
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Inflation und Zinsen", "en": "Inflation and rates"}
		page.Short = trl.S{"de": "Inflation,<br/>Zinsen", "en": "Inflation,<br/>Rates"}
		page.Width = 60

		rowLabelsEuroGer := []trl.S{
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "Deutschland",
				"en": "Germany",
			},
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				[]string{"pi_ez", "pi_deu"},
				radioVals4,
				rowLabelsEuroGer,
			)
			gb.MainLabel = trl.S{
				"de": "<b>4.</b> Die jährl. gesamtwirtschaftl. Inflationsrate wird mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>4.</b> Medium term (<b>6</b>&nbsp;months) yearly overall inflation rate will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.HeaderBottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 9
				inp.Desc = trl.S{
					"de": "<b>5a.</b> Die <b>kurzfristigen</b> Zinsen (3-Mo.-Interbanksätze) im <b>Euroraum</b> erwarten wir auf Sicht von 6&nbsp;Monaten",
					"en": "<b>5a.</b> We expect <b>short term</b> interest rates (3 months interbank) in the <b>euro area</b>",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "i_ez_low"
				inp.Min = 0
				inp.Max = 20
				inp.MaxChars = 4
				// inp.Validator = "inRange20"

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 2
				inp.Desc = trl.S{
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
				inp.Name = "i_ez_high"
				inp.Min = 0
				inp.Max = 20
				inp.MaxChars = 4
				// inp.Validator = "inRange20"

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.Desc = trl.S{
					"de": "und&nbsp;&nbsp;&nbsp;&nbsp;",
					"en": "and&nbsp;&nbsp;&nbsp;&nbsp;",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			// row below
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 2
				inp.Desc = trl.S{
					"de": " &nbsp;",
					"en": " &nbsp;",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 7
				inp.Desc = trl.S{
					"de": " [zentrales 90% Konfidenzintervall]",
					"en": " [central 90&nbsp;pct confidence interval]",
				}
				inp.CSSLabel = "special-input-textblock-smaller"
			}

		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.HeaderBottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 9
				inp.Desc = trl.S{
					"de": "<b>5b.</b> Die <b>langfristigen</b> Zinsen (Renditen 10jg. Staatsanleihen) in <b>Deutschland</b> erwarten wir auf Sicht von 6&nbsp;Monaten",
					"en": "<b>5b.</b> We expect <b>long term</b> interest rates in <b>Germany</b> in 6&nbsp;months",
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "r_deu_low"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
				// inp.Validator = "inRange100"

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 2
				inp.Desc = trl.S{
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
				inp.Name = "r_deu_high"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
				// inp.Validator = "inRange100"

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 3
				inp.Desc = trl.S{
					"de": "und&nbsp;&nbsp;&nbsp;&nbsp;",
					"en": "and&nbsp;&nbsp;&nbsp;&nbsp;",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			// row below
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 2
				inp.Desc = trl.S{
					"de": " &nbsp;",
					"en": " &nbsp;",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 7
				inp.Desc = trl.S{
					"de": " [zentrales 90% Konfidenzintervall]",
					"en": " [central 90&nbsp;pct confidence interval]",
				}
				inp.CSSLabel = "special-input-textblock-smaller"
			}

		}

	}

	//
	// page 4 - Credit situation
	{
		page := q.AddPage()
		page.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
		page.Label = trl.S{"de": "Markt", "en": "Market"}
		page.Short = trl.S{"de": "Kredit-<br/>markt", "en": "Credit<br/>Markets"}
		page.Width = 70

		rowLabelsCreditDemandSupply := []trl.S{
			{
				"de": "Kreditnachfrage",
				"en": "Credit demand",
			},
			{
				"de": "Kreditangebot",
				"en": "Credit supply",
			},
		}

		rowLabelsMediumLongTerm := []trl.S{
			{
				"de": "mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "medium term (<b>6</b>&nbsp;months)",
			},
			{
				"de": "langfristig (<b>24</b>&nbsp;Mo.)",
				"en": "long term (<b>24</b>&nbsp;months)",
			},
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				labelsVeryHighVeryLow(),
				[]string{"cd_deu", "cs_deu"},
				radioVals6,
				rowLabelsCreditDemandSupply,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6a.</b> Wie schätzen Sie die Kreditsituation in Deutschland ein?",
				"en": "<b>6a.</b> How do you assess credit conditions in Germany?",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				labelsStrongIncreaseStrongDecrease(),
				[]string{"c0_6", "c0_24"},
				radioVals6,
				rowLabelsMediumLongTerm,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6b.</b> Das (saisonbereinigte) Gesamtvolumen der Neukreditvergabe in Deutschland wird",
				"en": "<b>6b.</b> The seasonally adjusted volume of new credit in Germany will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				labelsStrongIncreaseStrongDecrease(),
				[]string{"cd_24_le", "cd_24_sme", "cd_24_re", "cd_24_co"},
				radioVals6,
				rowLabelsSmallLargeEnterprises,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6c.</b> Die (saisonbereinigte) Kreditnachfrage wird mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>6c.</b> The seasonally adjusted credit demand medium term (<b>6</b>&nbsp;months) will be",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	}

	//
	// page 5 - Credit influence factors
	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
		page.Label = trl.S{"de": "Einflussfaktoren", "en": "Influence Factors"}
		page.Short = trl.S{"de": "Kredit-<br/>faktoren", "en": "Credit<br/>Influencers"}
		page.Width = 70

		rowLabelsFinancingFactors := []trl.S{
			{
				"de": "Ausfallrisiken",
				"en": "Default risk",
			},
			{
				"de": "Risikotragfähigkeit",
				"en": "Risk profile",
			},
			{
				"de": "Refinanzierung",
				"en": "Refinancing",
			},
			{
				"de": "Wettbewerbssituation",
				"en": "Competitive environment",
			},
			{
				"de": "Regulierung",
				"en": "Regulation",
			},
			{
				"de": "EZB Politik",
				"en": "ECB policy",
			},
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				labelsVeryPositiveVeryNegative(),
				[]string{"c_inf_6_dr", "c_inf_6_ri", "c_inf_6_re", "c_inf_6_ce", "c_inf_6_rg", "c_inf_6_ep"},
				radioVals6,
				rowLabelsFinancingFactors,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6d.</b> Wie schätzen Sie den Einfluss folgender Faktoren auf die mittelfristige (<b>6</b>&nbsp;Mo.) Veränderung des Kreditangebots ein?",
				"en": "<b>6d.</b> How do you assess the influence of following factors on the medium term (<b>6</b>&nbsp;months) change of credit supply?",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				labelsStrongIncreaseStrongDecrease(),
				[]string{"c_std_6_le", "c_std_6_sme", "c_std_6_re", "c_std_6_co"},
				radioVals6,
				rowLabelsSmallLargeEnterprises,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6e.</b> Die (saisonbereinigte) Kreditstandards für Neukredite werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>6e.</b> The seasonally adjusted credit standards medium term (<b>6</b>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	}

	//
	// page 6 - Financial markets
	{
		page := q.AddPage()
		page.Section = trl.S{"de": "Finanzmärkte", "en": "Financial markets"}
		page.Label = trl.S{"de": "Preise", "en": "Prices"}
		page.Short = trl.S{"de": "Finanz-<br/>märkte", "en": "Financial<br/>Markets"}
		page.Width = 70

		rowLabelsUncorrelatedAssets := []trl.S{
			{
				"de": "DAX",
				"en": "German DAX",
			},
			{
				"de": "Rohöl (Nordsee Brent)",
				"en": "Brent Crude",
			},
			{
				"de": "Gold",
				"en": "Gold",
			},
			{
				"de": "US-Dollar (ggü. €)",
				"en": "Dollar / Euro",
			},
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				[]string{"sto_dax", "oil", "gold", "fx_usa"},
				radioVals4,
				rowLabelsUncorrelatedAssets,
			)
			gb.MainLabel = trl.S{
				"de": "<b>7a.</b> Die folgenden Aktienindizes / Rohstoffpreise / Wechselkurse werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>7a.</b> Following stock indices / raw materials / exchange rates will medium term (<b>6</b>&nbsp;months)",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": "<b>7b.</b>",
					"en": "<b>7b.</b>",
				}
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 100
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_6"
				inp.Min = 1000
				inp.Max = 50000
				inp.MaxChars = 6
				// inp.Validator = "inRange50000"

				inp.ColSpanLabel = 55
				// inp.CSSLabel = "special-line-height-higher"
				inp.ColSpanControl = 45
				inp.Desc = trl.S{
					"de": `Den DAX erwarten wir in 6&nbsp;Monaten bei `,
					"en": "We expect the German DAX in 6&nbsp;month at",
				}
				inp.Suffix = trl.S{"de": "Punkten", "en": "points"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_6_low"
				inp.Min = 1000
				inp.Max = 50000
				inp.MaxChars = 6
				// inp.Validator = "inRange50000"

				inp.ColSpanLabel = 55
				inp.ColSpanControl = 21
				inp.Desc = trl.S{
					"de": `Mit einer Wahrscheinlichkeit von 90&nbsp;Prozent wird der DAX dann zwischen `,
					"en": "With 90&nbsp;percent probability, the DAX will then be between",
				}
				inp.Suffix = trl.S{"de": "Punkten  &nbsp; und ", "en": "points &nbsp; and "}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_6_high"
				inp.Min = 1000
				inp.Max = 50000
				inp.MaxChars = 6
				// inp.Validator = "inRange50000"

				inp.ColSpanControl = 24
				inp.Suffix = trl.S{"de": "Punkten liegen.", "en": "points"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
			}
		}

		//
		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": "<b>7c.</b> Aus Sicht der Fundamentaldaten der DAX-Unternehmen ist der DAX derzeit",
					"en": "<b>7c.</b> The fundamentals of the companies comprising the DAX make the DAX currently",
				}
			}
		}

		// gr4
		{
			// gr := p.AddRadioMatrixGroupCSSGrid(labelsOvervaluedFairUndervalued(), names3rdMatrix, labels123Matrix)
			gr := page.AddGroup()
			gr.Cols = 12 // necessary, otherwise no vspacers
			gr.HeaderBottomVSpacers = 1

			for i2, lbl := range labelsOvervaluedFairUndervalued() {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "dax_fund"
				rad.ValueRadio = fmt.Sprintf("%v", i2+1)
				rad.Label = lbl
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 1
			}

		}

		//
		// gr5
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": "<b>8.</b> Die Wahrscheinlichkeit für ein Extremereignis im deutschen Finanzmarkt liegt",
					"en": "<b>8.</b> The probability for an extreme event in the German financial markets is",
				}
			}
		}

		// gr6
		{
			gr := page.AddGroup()
			gr.Cols = 100
			gr.HeaderBottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "crash_low"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
				// inp.Validator = "inRange100"

				inp.ColSpanLabel = 22
				inp.ColSpanControl = 12
				inp.Desc = trl.S{
					"de": " mittelfristig (<b>6</b>&nbsp;Mo.) bei  ",
					"en": " medium term (<b>6</b>&nbsp;months) at ",
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HRight
				inp.HAlignControl = qst.HLeft
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "crash_high"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
				// inp.Validator = "inRange100"

				inp.ColSpanLabel = 26
				inp.ColSpanControl = 40
				inp.Desc = trl.S{
					"de": "und langfristig (<b>24</b>&nbsp;Mo.) bei ",
					"en": "and long term (<b>24</b>&nbsp;months) at ",
				}

				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HRight
				inp.HAlignControl = qst.HLeft
			}
		}

	}

	err := addSeasonal1(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding seasonal1: %v", err)
	}

	err = addSeasonal2(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding seasonal2: %v", err)
	}

	//
	// page 7 - after seasonal
	// Finish questionnaire?  - one before last page
	{
		page := q.AddPage()
		page.Section = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Label = trl.S{"de": "", "en": ""}
		page.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Width = 65

		{
			gr := page.AddGroup()
			gr.Cols = 1 // necessary, otherwise no vspacers

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Desc = trl.S{
					"de": "Wollen Sie uns noch etwas mitteilen?",
					"en": "Any remarks or advice for us?",
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "remark"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		{
			gr := page.AddGroup()
			gr.Width = 100
			gr.Cols = 2 // necessary, otherwise no vspacers

			// todo
			// inp.Validator = "mustRadioGroup"
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = "no"
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Zugang bleibt bestehen.  Daten können in weiteren Sitzungen geändert/ergänzt werden. <br/>\n &nbsp;",
					"en": "Leave questionnaire open. Data  can be changed/completed&nbsp;in later sessions.     <br/>\n &nbsp;",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = "yes"
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Fragebogen ist abgeschlossen und kann nicht mehr geöffnet werden. <br/>\n &nbsp;",
					"en": "Questionnaire is finished. No more edits.                         <br/>\n &nbsp;",
				}
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 2 // necessary, otherwise no vspacers
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "Durch Klicken auf 'OK' erhalten Sie eine Zusammenfassung Ihrer Antworten",
					"en": "By Clicking 'OK' you receive a summary of your answers",
				}
				inp.ColSpanLabel = 1
				inp.CSSLabel = "special-line-height-higher"
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = trl.S{
					"de": "OK",
					"en": "OK",
				}
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.HAlignControl = qst.HCenter
				inp.HAlignControl = qst.HLeft
			}
		}

		// page.ExampleSixColumnsLabelRight()

	}

	// page 8 - after seasonal
	// Report of results
	{
		p := q.AddPage()
		p.Label = trl.S{"de": "Ihre Eingaben", "en": "Summary of results"}
		p.NoNavigation = true
		{
			gr := p.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.CSSLabel = "special-line-height-higher"
				inp.DynamicFunc = "RepsonseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.CSSLabel = "special-line-height-higher"
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 1
				inp.CSSLabel = "special-line-height-higher"
				impr := trl.S{}
				for _, lc := range q.LangCodes {
					w1 := &strings.Builder{}
					err := tpl.RenderStaticContent(
						w1, "site-imprint.md", q.Survey.Type, lc,
					)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = w1.String()

				}
				inp.Desc = impr
			}
		}
	}

	// quest.ClosingTime = time.Now()

	err = q.Validate()
	if err != nil {
		return nil, fmt.Errorf("Error validating questionnaire: %v", err)
	}
	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}

	return &q, nil
}
