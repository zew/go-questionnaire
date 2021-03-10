package fmt

import (
	"fmt"
	"log"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/go-questionnaire/trl"
)

var radioVals4 = []string{"1", "2", "3", "4"}
var radioVals6 = []string{"1", "2", "3", "4", "5", "6"}
var columnTemplate4 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0.4, 1, // no answer slightly apart
}
var columnTemplate6 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

var rowLabelsEuroGerUSGlob = []trl.S{
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
		"de": "China",
		"en": "China",
	},
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

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("fmt")
	q.Survey.Params = params
	q.LangCodes = []string{"de", "en"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}
	q.Survey.Name = trl.S{
		"de": "Index / Finanzmarkttest",
		"en": "Indicator of Economic Sentiment",
	}

	q.Version = 2

	// page 0
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Begrüßung", "en": "Greeting"}
		page.NoNavigation = true
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": "Sehr geehrte Finanzmarktexpertin, sehr geehrter Finanzmarktexperte, herzlich willkommen beim ZEW-Finanzmarkttest.",
					"en": "Dear expert, welcome to the ZEW financial markets survey.",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				impr := trl.S{}
				for _, lc := range q.LangCodes {
					w1 := &strings.Builder{}
					err := tpl.RenderStaticContent(
						w1, "page-0-data-protection.md", q.Survey.Type, lc,
					)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = w1.String()

				}
				inp.Label = impr
			}
		}

		// gr1
		{

			gr := page.AddGroup()
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style = css.DesktopWidthMax(gr.Style, "26rem")
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Sind Sie die angeschriebene Person?", "en": "Are you the addressee?"}
				inp.ColSpan = 6
				inp.ColSpanLabel = 6
			}

			lblStyle := css.NewStylesResponsive(nil)
			lblStyle.Desktop.StyleText.AlignHorizontal = "left"
			lblStyle.Desktop.StyleBox.Padding = "0 0 0 1rem"
			lblStyle.Mobile.StyleBox.Padding = "0 0 0 2rem"
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "selbst"
				rad.ValueRadio = "1"
				rad.ColSpan = 6
				rad.ColSpanLabel = 5
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Ja, ich bin die angeschriebene Person.",
					"en": "Yes, I am the addressee.",
				}
				rad.StyleLbl = lblStyle
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "selbst"
				rad.ValueRadio = "2"
				rad.ColSpan = 6
				rad.ColSpanLabel = 5
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Nein, ich fülle den Fragebogen in Vertretung der angeschriebenen Person aus.",
					"en": "No. I am filling in for the addressee.",
				}
				rad.StyleLbl = lblStyle
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "Meine Adresse hat sich geändert",
					"en": "My address has changed",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "contact"
				inp.MaxChars = 150
				inp.ColSpanControl = 1
			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
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
				inp.StyleCtl = css.ItemEndMA(inp.StyleCtl)
			}
		}

	}

	// page 1
	{
		page := q.AddPage()
		// pge.Section = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		page.Label = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		page.Short = trl.S{"de": "Konjunktur", "en": "Business cycle"}
		page.Style = css.DesktopWidthMax(page.Style, "34rem") // 55

		page.ValidationFuncName = "fmtPage1"
		page.ValidationFuncMsg = trl.S{
			"de": "Summiert sich nicht zu 100. Wirklich weiter?",
			"en": "Does not add up. Really continue?",
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsGoodBad(),
				[]string{"y0_ez", "y0_deu", "y0_usa", "y0_chn"},
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

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsImproveDeteriorate(),
				[]string{"y_ez", "y_deu", "y_usa", "y_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>2a.</b> Die gesamtwirtschaftliche Situation wird sich mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>2a.</b> The overall economic situation medium term (<bx>6</bx>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 10
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Mobile.StyleGridContainer.GapRow = "0.02rem"

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 10
				inp.Label = trl.S{
					"de": "<b>2b.</b> Für wie wahrscheinlich halten Sie die folgenden mittelfristigen (<bx>6</bx>&nbsp;Mo.) Entwicklungen der gesamtwirtschaftlichen Situation in Deutschland?",
					"en": "<b>2b.</b> How likely are the following medium term (<bx>6</bx>&nbsp;months) developments of the general economic situation in Germany?",
				}

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0.8rem 0"

			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Verbesserung",
					"en": "Improvement",
				}
				inp.Style = css.ItemStartCA(inp.Style)
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Gleich bleiben",
					"en": "Remain the same",
				}
				inp.Style = css.ItemStartCA(inp.Style)
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Verschlechterung",
					"en": "Deterioration",
				}
				inp.Style = css.ItemStartCA(inp.Style)
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "&#931;",
					"en": "&#931;",
				}
				inp.Style = css.ItemCenteredMCA(inp.Style)
				inp.Style = css.ItemStartCA(inp.Style)
			}
			// second row
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_probgood"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = 3
				inp.ColSpanControl = 3
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_probnormal"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = 3
				inp.ColSpanControl = 3
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_probbad"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = 3
				inp.ColSpanControl = 3
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": "100&nbsp;%",
					"en": "100&nbsp;%",
				}
				inp.Style = css.ItemCenteredMCA(inp.Style)
			}

		}
	}

	//
	// page 2 - inflation + zinsen
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Inflation und Zinsen", "en": "Inflation and Interest Rates"}
		page.Short = trl.S{"de": "Inflation,<br/>Zinsen", "en": "Inflation,<br/>Inter. Rates"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				[]string{"pi_ez", "pi_deu", "pi_usa", "pi_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>3.</b> Die jährliche gesamtwirtschaftliche Inflationsrate wird mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>3.</b> Yearly overall inflation rate in the medium term (<bx>6</bx>&nbsp;months)  will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr1
		{
			var rowLabelsEuroGerUSGlob = []trl.S{
				{
					"de": "Euroraum",
					"en": "Euro area",
				},
				{
					"de": "USA",
					"en": "US",
				},
				{
					"de": "China",
					"en": "China",
				},
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				// []string{"i_ez", "i_deu", "i_usa", "i_chn"},
				[]string{"i_ez", "i_usa", "i_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>4.</b> Die <i>kurzfristigen</i> Zinsen (3-Mo.-Interbanksätze) werden mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>4.</b> <i>Short term</i> interest rates (3&nbsp;months interbank) in the medium term (<bx>6</bx>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr2
		{
			var rowLabelsEuroGerUSGlob = []trl.S{
				{
					"de": "Deutschland",
					"en": "Germany",
				},
				{
					"de": "USA",
					"en": "US",
				},
				{
					"de": "China",
					"en": "China",
				},
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				// []string{"r_ez", "r_deu", "r_usa", "r_chn"},
				[]string{"r_deu", "r_usa", "r_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob,
			)
			gb.MainLabel = trl.S{
				"de": "<b>5.</b> Die <i>langfristigen</i> Zinsen (zehnjähriger Staatsanleihen) werden mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>5.</b> <i>Long term</i> interest rates (10-year govt. bonds) in the medium term (6&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.BottomVSpacers = 4
		}

	}

	//
	// page 3 - financial markets
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Aktienmärkte", "en": "Stock Markets"}
		page.Short = trl.S{"de": "Aktien-<br>märkte", "en": "Stock<br>Markets"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		rowLabelsUncorrelatedAssets := []trl.S{
			{
				"de": "EURO STOXX 50",
				"en": "EURO STOXX 50",
			},
			{
				"de": "DAX",
				"en": "German DAX",
			},
			{
				"de": "Dow Jones (USA)",
				"en": "Dow Jones (USA)",
			},
			{
				"de": "SSE Composite (China)",
				"en": "SSE Composite (China)",
			},
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				// []string{"euro_stoxx50", "sto_dax", "dow_jones", "sto_sse_comp_chn"},
				[]string{"sto_ez", "sto_dax", "dow_usa", "sto_sse_comp_chn"},
				radioVals4,
				rowLabelsUncorrelatedAssets,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6a.</b> Die folgenden Aktienindizes werden mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>6a.</b> Following stock indices in the medium term (<bx>6</bx>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0rem"
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.4rem"
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_erw"
				inp.Min = 2000
				inp.Max = 50000
				inp.MaxChars = 6

				inp.ColSpan = 6
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 2
				inp.Label = trl.S{
					"de": "<b>6b.</b> Den DAX erwarte ich in 6&nbsp;Monaten bei ",
					"en": "<b>6b.</b> We expect the German DAX in 6&nbsp;month at",
				}
				inp.Suffix = trl.S{"de": "Punkten", "en": "points"}
			}
			// second row
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 6
				inp.ColSpanLabel = 6
				inp.Label = trl.S{
					"de": " <br>Mit einer Wahrscheinlichkeit von 90&nbsp;Prozent liegt der DAX dann zwischen ",
					"en": " <br>With 90&nbsp;percent probability, the DAX will then be between",
				}
			}
			// third row
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_min"
				inp.Min = 2000
				inp.Max = 50000
				inp.MaxChars = 6
				inp.ColSpan = 3
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
				inp.Label = trl.S{
					"de": " &nbsp; ",
					"en": " &nbsp; ",
				}
				inp.Suffix = trl.S{
					"de": "Punkten ",
					"en": "points",
				}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": " und &nbsp; ",
					"en": " and &nbsp; ",
				}
				inp.StyleLbl = css.TextEnd(inp.StyleLbl)
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_max"
				inp.Min = 2000
				inp.Max = 50000
				inp.MaxChars = 6
				inp.ColSpan = 2
				inp.ColSpanControl = 2
				inp.Suffix = trl.S{}
				inp.Suffix = trl.S{
					"de": "Punkten",
					"en": "points",
				}
			}
			/*
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.ColSpanLabel = 6
					inp.Label = trl.S{
						"de": " liegen.",
						"en": " &nbsp; ",
					}
					inp.Style = css.NewStylesResponsive(inp.Style)
					inp.Style.Desktop.BoxStyle.Position = "relative"
					inp.Style.Desktop.BoxStyle.Top = "-0.4rem"
				}
			*/
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "<b>6c.</b> Aus Sicht der Fundamentaldaten der DAX-Unternehmen ist der DAX derzeit",
					"en": "<b>6c.</b> The fundamentals of the companies comprising the DAX make the DAX currently",
				}
			}
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.Padding = "0 0 0.5rem  0"
		}

		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				[]float32{
					0, 1,
					0, 1,
					0, 1,
				},
				labelsOvervaluedFairUndervalued(),
				[]string{"dax_fund"},
				[]string{"1", "2", "3"},
				nil,
			)
			// gb.MainLabel = trl.S{
			// }
			gr := page.AddGrid(gb)
			gr.Style = css.DesktopWidthMax(gr.Style, "30rem")

			gr.Style.Desktop.StyleGridContainer.GapColumn = "0"
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.5rem"

			gr.Style.Desktop.StyleBox.Position = "relative"
			gr.Style.Desktop.StyleBox.Left = "-1.1rem"
			gr.Style.Mobile.StyleBox.Left = "0"
			gr.OddRowsColoring = true
		}
	}

	//
	// page 4
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Währungen", "en": "Currencies"}
		page.Short = trl.S{"de": "Währungen", "en": "Currencies"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		rowLabelsCurrencies := []trl.S{
			{
				"de": "US Dollar",
				"en": "US Dollar",
			},
			{
				"de": "Yuan",
				"en": "Yuan",
			},
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecreaseCurrency(),
				[]string{"fx_usa", "fx_chn"},
				radioVals4,
				rowLabelsCurrencies,
			)
			gb.MainLabel = trl.S{
				"de": "<b>7.</b> Folgende Währungen werden gegenüber dem Euro mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>7.</b> In the medium term (<bx>6</bx>&nbsp;months), following currencies will against the Euro",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}
	}

	//
	// page 5
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Sektoren", "en": "Sectors"}
		page.Short = trl.S{"de": "Sektoren", "en": "Sectors"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		rowLabelsCurrencies := []trl.S{
			{
				"de": "Banken",
				"en": "Banking",
			},
			{
				"de": "Versicherungen",
				"en": "Insurance",
			},
			{
				"de": "Fahrzeugbau",
				"en": "Automotive",
			},
			{
				"de": "Chemie, Pharma",
				"en": "Chem. Pharmac.",
			},
			{
				"de": "Stahl/NE-Metalle",
				"en": "Metallurgy",
			},
			{
				"de": "Elektro",
				"en": "Electrical Engineering",
			},
			{
				"de": "Maschinen&shy;bau",
				"en": "Mechanical Engineering",
			},
			// row 2
			{
				"de": "Konsum, Handel",
				"en": "Retail",
			},
			{
				"de": "Baugewerbe",
				"en": "Construction",
			},
			{
				"de": "Versorger",
				"en": "Utilities",
			},
			{
				"de": "Dienstleister",
				"en": "Service Sect.",
			},
			{
				"de": "Telekommunikation",
				"en": "Telco",
			},
			{
				"de": "Inform.-Technologien",
				"en": "IT",
			},
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsImproveDeteriorateSectoral(),
				[]string{
					"sec_banks", "sec_insur", "sec_cars", "sec_chemi", "sec_steel", "sec_elect", "sec_mecha",
					// "sec_c",
					"sec_consu", "sec_const", "sec_utili", "sec_servi", "sec_telec", "sec_infor"},
				radioVals4,
				rowLabelsCurrencies,
			)
			gb.MainLabel = trl.S{
				"de": "<b>8.</b> Die Ertragslage der Unternehmen in Deutschland wird mittelfristig (<bx>6</bx>&nbsp;Mo.) in den folgenden Branchen ",
				"en": "<b>8.</b> Revenues of German enterprise will medium term (<bx>6</bx>&nbsp;months)",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}
	}

	// page 6
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Rezession", "en": "Recession"}
		page.Short = trl.S{"de": "Rezession", "en": "Recession"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 5

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 5
				inp.ColSpanLabel = 5
				inp.Label = trl.S{
					"de": fmt.Sprintf("<b>9.</b> Die Wahrscheinlichkeit eines negativen BIP-Wachstums in Deutschland (Wachstum des realen & saisonbereinigten BIP zum Vorquartal) liegt bei"), //nextQ()
					"en": fmt.Sprintf("<b>9.</b> The probability of negative GDP growth in Germany (growth of real & seasonal adjusted GDP against previous quarter) is"),                     // nextQ()
				}
			}
			{
				inp := gr.AddInput()
				inp.Label = trl.S{
					"de": fmt.Sprintf("Aktuelles Quartal"),
					"en": fmt.Sprintf("Current quarter"),
				}
				inp.Type = "number"
				inp.Name = "y_recession_q0"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4

				inp.ColSpan = 5
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
			}
			{
				inp := gr.AddInput()
				inp.Label = trl.S{
					"de": fmt.Sprintf("Folgendes Quartal"),
					"en": fmt.Sprintf("Next quarter"),
				}
				inp.Type = "number"
				inp.Name = "y_recession_q1"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 4

				inp.ColSpan = 5
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
			}
		}

	}

	err := addSeasonal1(&q)
	if err != nil {
		return nil, fmt.Errorf("Error adding seasonal1: %v", err)
	}
	/*
		err = addSeasonal2(&q)
		if err != nil {
			return nil, fmt.Errorf("Error adding seasonal2: %v", err)
		}
	*/

	//
	// page 7 - after seasonal
	// Finish questionnaire?  - one before last page
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Short = trl.S{"de": "Abschluss", "en": "Finish"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem")

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
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

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			// todo
			// inp.Validator = "mustRadioGroup"
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = "no"
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Zugang bleibt bestehen.  Daten können in weiteren Sitzungen geändert/ergänzt werden. ",
					"en": "Leave questionnaire open. Data  can be changed/completed&nbsp;in later sessions.     ",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = "yes"
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Fragebogen ist abgeschlossen und kann nicht mehr geöffnet werden. ",
					"en": "Questionnaire is finished. No more edits.                         ",
				}
			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Label = trl.S{
					"de": "Durch Klicken erhalten Sie eine Zusammenfassung Ihrer Antworten",
					"en": "By clicking you receive a summary of your answers",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				// inp.Label = trl.S{"de": "", "en": ""}
				// inp.Label = trl.S{
				// 	"de": "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;OK&nbsp;&nbsp;&nbsp;&nbsp;",
				// 	"en": "&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;OK&nbsp;&nbsp;&nbsp;&nbsp;",
				// }
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

		// pge.ExampleSixColumnsLabelRight()

	}

	// page 8 - after seasonal
	// Report of results
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Ihre Eingaben", "en": "Summary of results"}
		page.NoNavigation = true
		page.Style = css.DesktopWidthMax(page.Style, "calc(100% - 1.2rem)")
		page.Style = css.DesktopWidthMax(page.Style, "40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "RepsonseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 1
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
				inp.Label = impr
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
