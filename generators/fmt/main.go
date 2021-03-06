package fmt

import (
	"fmt"
	"log"
	"strings"

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
	1, 1,
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
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
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
				inp.Label = impr
			}
		}

		// gr1
		{

			gr := page.AddGroup()
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleBox.WidthMax = "26rem"
			gr.Style.Mobile.StyleBox.WidthMax = "none"

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
				rad.Name = "proxy"
				rad.ValueRadio = "no"
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
				rad.Name = "proxy"
				rad.ValueRadio = "yes"
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
				inp.Name = "address_change"
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
		page.Width = 60

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
			// gr.Style.Desktop.GridContainerStyle.TemplateColumns = "1fr 1fr 1fr 1fr 1fr 0.4fr 1fr"
			// gr.Style.Desktop.GridContainerStyle.ColumnGap = "0.1rem"
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
				"de": "<b>2a.</b> Die gesamtwirtschaftliche Situation wird sich mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>2a.</b> The overall economic situation medium term (<b>6</b>&nbsp;months) will",
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		/*
			{
				gb := qst.NewGridBuilderRadios(
					columnTemplate4,
					labelsImproveDeteriorate(),
					[]string{"y24_ez", "y24_deu", "y24_usa", "y24_chn"},
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

		*/

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
				// inp.ColSpanLabel = 10
				inp.Label = trl.S{
					"de": "<b>2b.</b> Für wie wahrscheinlich halten Sie die folgenden mittelfristigen (<b>6</b>&nbsp;Mo.) Entwicklungen der gesamtwirtschaftlichen Situation in Deutschland?",
					"en": "<b>2b.</b> How likely are the following medium term (<b>6</b>&nbsp;months) developments of the general economic situation in Germany?",
				}

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0.8rem 0"

			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				// inp.ColSpanLabel = 3
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
				// inp.ColSpanLabel = 3
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
				// inp.ColSpanLabel = 3
				inp.Label = trl.S{
					"de": "Verschlechterung",
					"en": "Deterioration",
				}
				inp.Style = css.ItemStartCA(inp.Style)
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.ColSpanLabel = 1
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
				// inp.ColSpanLabel = 1
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
		page.Width = 60

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
				"de": "<b>3.</b> Die jährliche gesamtwirtschaftliche Inflationsrate wird mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>3.</b> Yearly overall inflation rate in the medium term (<b>6</b>&nbsp;months)  will",
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
				"de": "<b>4.</b> Die <b>kurzfristigen</b> Zinsen (3-Mo.-Interbanksätze) werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>4.</b> <b>Short term</b> interest rates (3&nbsp;months interbank) in the medium term (<b>6</b>&nbsp;months) will",
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
				"de": "<b>5.</b> Die <b>langfristigen</b> Zinsen (Renditen 10jg. Staatsanleihen) werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>5.</b> <b>Long term</b> interest rates (10-year govt. bonds) in the medium term (6&nbsp;months) will",
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
		// page.Section = trl.S{"de": "Finanzmärkte", "en": "Financial markets"}
		page.Label = trl.S{"de": "Aktienmärkte", "en": "Stock Markets"}
		page.Short = trl.S{"de": "Aktien-<br>märkte", "en": "Stock<br>Markets"}
		page.Width = 60

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
				[]string{"euro_stoxx50", "sto_dax", "dow_jones", "sse_comp_chn"},
				radioVals4,
				rowLabelsUncorrelatedAssets,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6a.</b> Die folgenden Aktienindizes werden mittelfristig (<b>6</b>&nbsp;Mo.)",
				"en": "<b>6a.</b> Following stock indices in the medium term (<b>6</b>&nbsp;months) will",
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
			gr.Style.Desktop.StyleGridContainer.GapRow = "0rem"
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_6"
				inp.Min = 1000
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
				inp.Name = "dax_6_low"
				inp.Min = 1000
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
				inp.Name = "dax_6_high"
				inp.Min = 1000
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

			/* 	{
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
			} */

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
			gr.Style.Desktop.StyleGridContainer.GapColumn = "0"
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.5rem"
			gr.Style.Desktop.StyleBox.WidthMax = "30rem"
			gr.Style.Mobile.StyleBox.WidthMax = "none"

			gr.Style.Desktop.StyleBox.Position = "relative"
			gr.Style.Desktop.StyleBox.Left = "-1.1rem"
			gr.Style.Mobile.StyleBox.Left = "0"
			gr.OddRowsColoring = true
		}

		/*
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 0
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": "<b>8.</b> Die Wahrscheinlichkeit für ein Extremereignis im deutschen Finanzmarkt liegt",
						"en": "<b>8.</b> The probability for an extreme event in the German financial markets is",
					}
				}
			}

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
					inp.Label = trl.S{
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
					inp.Label = trl.S{
						"de": "und langfristig (<b>24</b>&nbsp;Mo.) bei ",
						"en": "and long term (<b>24</b>&nbsp;months) at ",
					}

					inp.Suffix = trl.S{"de": "%", "en": "pct"}
					inp.HAlignLabel = qst.HRight
					inp.HAlignControl = qst.HLeft
				}
			}
		*/

	}

	//
	// page 4
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Währungen", "en": "Currencies"}
		page.Short = trl.S{"de": "Währungen", "en": "Currencies"}
		page.Width = 60
	}

	//
	// page 4
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Wachstum", "en": "Growth"}
		page.Short = trl.S{"de": "Wachstum", "en": "Growth"}
		page.Width = 60

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 5
				inp.Label = trl.S{
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
				inp.Label = trl.S{
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
				inp.Label = trl.S{
					"de": fmt.Sprintf("Unsere Prognose für das BIP Wachstum für das Jahr %v <xxbr/>\n(real, saisonbereinigt)", nextY()),
					"en": fmt.Sprintf("Our estimate for the GDP growth in %v                <xxbr/>\n(real, seasonally adjusted)", nextY()),
				}
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
				inp.HAlignLabel = qst.HLeft
				inp.HAlignControl = qst.HLeft
				inp.ColSpanControl = 1
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 5 // necessary, otherwise no vspacers

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpanLabel = 5
				inp.Label = trl.S{
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
				inp.Label = trl.S{
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
				inp.Label = trl.S{
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
	// page 5 - Credit situation
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
	// page 6 - Credit influence factors
	{
		page := q.AddPage()
		// pge.Section = trl.S{"de": "Kreditsituation", "en": "Credit situation"}
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

		{
			gr := page.AddGroup()
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
				inp.Label = trl.S{
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
				inp.Label = trl.S{
					"de": "OK",
					"en": "OK",
				}
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.HAlignControl = qst.HCenter
				inp.HAlignControl = qst.HLeft
			}
		}

		// pge.ExampleSixColumnsLabelRight()

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
