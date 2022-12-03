package fmt

import (
	"fmt"
	"strconv"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create creates a JSON file for a financial markets survey
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"de", "en"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW", "en": "ZEW"}

	// Number of approx google hits
	// 55.000 - ZEW Index - and wikipedia entry
	// 32.000 - ZEW Konjunkturerwartungen
	//  7.000 - ZEW Indicator of Economic Sentiment
	//  4.000 - ZEW Financial market survey
	//  2.500 - ZEW Finanzmarkttest
	q.Survey.Name = trl.S{
		"de": "Index / Finanzmarkttest",
		"en": "Index / Indicator of Econ. Sentiment",
	}

	// page 0
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Begrüßung", "en": "Greeting"}
		page.NoNavigation = true
		page.SuppressProgressbar = true
		page.WidthMax("36rem")

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
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "page-0-data-protection.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

		// gr1
		{

			gr := page.AddGroup()
			gr.Cols = 6
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.WidthMax("26rem")
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
					"en": "No, I am filling in for the addressee.",
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
				inp.Label = cfg.Get().Mp["next"].Pad(3)
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
		page.WidthMax("34rem") // 55

		page.ValidationFuncName = "fmtPage1"
		page.ValidationFuncMsg = trl.S{
			// "de": "Summiert sich nicht zu 100. Wirklich weiter?",
			// Möchten Sie dies ändern?
			"de": "Ihre Antworten auf Frage 2b addieren sich nicht zu 100%. Wirklich weiter?",
			"en": "Your answers to question 2b dont add up to 100%. Continue anyway?",
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsGoodBad(),
				[]string{"y0_ez", "y0_deu", "y0_usa", "y0_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob1,
			)
			gb.MainLabel = trl.S{
				"de": "<b>1.</b> &nbsp; Die gesamtwirtschaftliche Situation beurteilen wir zurzeit als",
				"en": "<b>1.</b> &nbsp; We estimate the current overall macroeconomic situation as being",
			}
			gr := page.AddGrid(gb)
			_ = gr

		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsImproveDeteriorate(),
				[]string{"y_ez", "y_deu", "y_usa", "y_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob1,
			)
			gb.MainLabel = trl.S{
				"de": "<b>2a.</b> &nbsp; Die gesamtwirtschaftliche Situation wird sich mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>2a.</b> &nbsp; In the medium-term (<bx>6</bx>&nbsp;months), the overall macroeconomic situation will",
			}
			gr := page.AddGrid(gb)
			_ = gr
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
					"de": "<b>2b.</b> &nbsp; Für wie wahrscheinlich halten Sie die folgenden mittelfristigen (<bx>6</bx>&nbsp;Mo.) Entwicklungen der gesamtwirtschaftlichen Situation in Deutschland?",
					"en": "<b>2b.</b> &nbsp; Please assess the probability of the following medium-term (<bx>6</bx>&nbsp;months) developments of the overall macroeconomic situation in Germany (in percent).",
				}

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Mobile.StyleBox.Padding = "0 0 0.8rem 0"

			}
			// first row: labels
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Verbesserung",
					"en": "improvement",
				}
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style.Mobile.StyleBox.Padding = "0 0.8rem 0 0"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Gleich bleiben",
					"en": "no change",
				}
				inp.Style = css.ItemStartCA(inp.Style)
				inp.Style.Mobile.StyleBox.Padding = "0 0.8rem 0 0"
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "Verschlechterung",
					"en": "worsening",
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
			// second row: inputs
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "y_probgood"
				inp.Suffix = trl.S{"de": "%", "en": "%"}
				inp.ColSpan = 3
				inp.ColSpanControl = 3
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0
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
				inp.Step = 0
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
				inp.Step = 0
				inp.MaxChars = 4
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Label = trl.S{
					"de": "100%", //  "100&nbsp;%",
					"en": "100%", //  "100&nbsp;%",
				}
				inp.Style = css.ItemCenteredMCA(inp.Style)
			}
		}

		// gr3a
		offsetDestatis := 0 // next quarter
		if osd, err := s.Param("destatis"); err == nil {
			offsetDestatis, _ = strconv.Atoi(osd)
		}

		{
			gr := page.AddGroup()
			gr.Cols = 5
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 5
				inp.ColSpanLabel = 5
				inp.Label = trl.S{
					"de": fmt.Sprint("<b>2c.</b> &nbsp; Die Wahrscheinlichkeit eines negativen BIP-Wachstums in Deutschland (Wachstum des realen & saisonbereinigten BIP zum Vorquartal) liegt bei"), //nextQ()
					"en": fmt.Sprint("<b>2c.</b> &nbsp; The probability of a negative GDP growth in Germany (quarterly growth of the seasonally adjusted real GDP) will be"),                         // nextQ()
				}
			}
		}
		// gr3b
		{
			gr := page.AddGroup()
			gr.Cols = 4
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.WidthMax("15rem")

			{
				inp := gr.AddInput()
				inp.Label = trl.S{
					"de": fmt.Sprintf("Aktuelles Quartal (%v)", q.Survey.Quarter(0)),
					"en": fmt.Sprintf("Current quarter  (%v)", q.Survey.Quarter(0)),
				}
				if offsetDestatis != 0 {
					for key, val := range inp.Label {
						inp.Label[key] = val + "*"
					}
				}

				// inp.Tooltip = trl.S{
				// 	"de": fmt.Sprintf("Unmittelbar zurückliegendes Quartal"),
				// 	"en": fmt.Sprintf("Most recent quarter"),
				// }
				inp.Type = "number"
				inp.Name = "y_recession_q0"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.01
				inp.MaxChars = 5

				inp.ColSpan = 4
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
			}
			{
				inp := gr.AddInput()
				inp.Label = trl.S{
					"de": fmt.Sprintf("Folgendes Quartal (%v)", q.Survey.Quarter(1)),
					"en": fmt.Sprintf("Next quarter  (%v)", q.Survey.Quarter(1)),
				}
				inp.Type = "number"
				inp.Name = "y_recession_q1"
				inp.Min = 0
				inp.Max = 100
				inp.Step = 0.01
				inp.MaxChars = 5

				inp.ColSpan = 4
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				inp.Suffix = trl.S{"de": "%", "en": "pct"}
			}

			// row 2c quarter explanation / footnote

		}

		if offsetDestatis != 0 {
			var lblFootnote = trl.S{
				"de": fmt.Sprintf("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<superscript>*</superscript><span style='font-size:80%%'> Die realisierten Zahlen für %v werden erst <a  target='_blank'  href='https://www.destatis.de/SiteGlobals/Forms/Suche/Termine/DE/Terminsuche_Formular.html' >später</a> veröffentlicht.<span>", q.Survey.Quarter(0)),
				"en": fmt.Sprintf("&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;<superscript>*</superscript><span style='font-size:80%%'> Realized numbers for %v are only published <a  target='_blank'  href='https://www.destatis.de/SiteGlobals/Forms/Suche/Termine/DE/Terminsuche_Formular.html' >later</a>.<span>", q.Survey.Quarter(0)),
			}

			// gr3c
			// 	full width
			{
				gr := page.AddGroup()
				gr.Cols = 1
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = lblFootnote
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1

					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
					inp.StyleLbl.Mobile.StyleBox.Position = "relative"
					inp.StyleLbl.Mobile.StyleBox.Top = "0.6rem"
				}
			}

		}

	}

	//
	// page 2 - inflation + zinsen
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Inflation und Zinsen", "en": "Inflation and Interest Rates"}
		page.Short = trl.S{"de": "Inflation,<br>Zinsen", "en": "Inflation,<br>Interest Rates"}
		page.WidthMax("36rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				[]string{"pi_ez", "pi_deu", "pi_usa", "pi_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob1,
			)
			gb.MainLabel = trl.S{
				"de": "<b>3.</b> &nbsp; Die jährliche gesamtwirtschaftliche Inflationsrate wird mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>3.</b> &nbsp; In the medium-term (<bx>6</bx>&nbsp;months), the annual inflation rate will",
			}
			gr := page.AddGrid(gb)
			_ = gr

		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				// []string{"i_ez", "i_deu", "i_usa", "i_chn"},
				// i_ez goes into i_deu
				[]string{"i_deu", "i_usa", "i_chn"},
				radioVals4,
				rowLabelsEuroGerUSGlob2,
			)
			gb.MainLabel = trl.S{
				"de": "<b>4.</b> &nbsp; Die <i>kurzfristigen</i> Zinsen (3-Mo.-Interbanksätze) werden mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>4.</b> &nbsp; In the medium-term (<bx>6</bx>&nbsp;months), <i>short-term</i> interest rates (3-month interbank rates) will",
			}
			gr := page.AddGrid(gb)
			_ = gr

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
				"de": "<b>5.</b> &nbsp; Die <i>langfristigen</i> Zinsen (zehnjähriger Staatsanleihen) werden mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>5.</b> &nbsp; In the medium-term, <i>long-term</i> interest rates (yields on 10-year sovereign bonds) will",
			}
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 4
		}

	}

	//
	// page 3 - financial markets
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Aktienmärkte", "en": "Stock Markets"}
		page.Short = trl.S{"de": "Aktien-<br>märkte", "en": "Stock<br>Markets"}
		page.WidthMax("36rem")

		page.ValidationFuncName = "fmtPage3"
		page.ValidationFuncMsg = trl.S{
			// Möchten Sie dies ändern?
			"de": "Ihre Punktprognose für den DAX in sechs Monaten liegt nicht in dem von Ihnen angegebenen Intervall. Wirklich weiter?",
			"en": "Your forcecast for the DAX in six months is not inside your MIN/MAX interval. Continue anyway?",
		}

		rowLabelsUncorrelatedAssets := []trl.S{
			{
				"de": "EURO STOXX 50",
				"en": "EURO STOXX 50",
			},
			{
				"de": "DAX",
				"en": "DAX (Germany)",
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

		ph := trl.S{
			"de": "0000",
			"en": "0000",
		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsIncreaseDecrease(),
				// []string{"euro_stoxx50", "sto_dax", "dow_jones", "sto_sse_comp_chn"},
				[]string{"sto_ez", "sto_dax", "sto_usa", "sto_sse_comp_chn"},
				radioVals4,
				rowLabelsUncorrelatedAssets,
			)
			gb.MainLabel = trl.S{
				"de": "<b>6a.</b> &nbsp; Die folgenden Aktienindizes werden mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>6a.</b> &nbsp; In the medium-term (<bx>6</bx>&nbsp;months), the following stock market indices will",
			}
			gr := page.AddGrid(gb)
			_ = gr

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
				inp.Placeholder = ph
				inp.Min = 2000
				inp.Max = 50000
				inp.Step = 0
				inp.MaxChars = 6

				inp.ColSpan = 6
				inp.ColSpanLabel = 4
				inp.ColSpanControl = 2
				inp.Label = trl.S{
					"de": `
					   <b>6b.</b> &nbsp;  
					   Den DAX erwarte ich in 6&nbsp;Monaten&nbsp;bei 
					`,
					"en": `
					   <b>6b.</b> &nbsp; 
					   Six&nbsp;months ahead, I expect the DAX to stand&nbsp;at
					`,
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
					"en": " <br>With a probability of 90&nbsp;per&nbsp;cent, the DAX will then range between",
				}
			}
			// third row
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "dax_min"
				inp.Placeholder = ph
				inp.Min = 2000
				inp.Max = 50000
				inp.Step = 0
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
				inp.Placeholder = ph
				inp.Min = 2000
				inp.Max = 50000
				inp.Step = 0
				inp.MaxChars = 6
				inp.ColSpan = 2
				inp.ColSpanControl = 2
				inp.Suffix = trl.S{}
				inp.Suffix = trl.S{
					"de": "Punkten",
					"en": "points",
				}
			}
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
					"en": "<b>6c.</b> &nbsp; In view of the fundamentals of the DAX companies, the DAX is currently",
					"de": "<b>6c.</b> &nbsp; Aus Sicht der Fundamentaldaten der DAX-Unternehmen ist der DAX derzeit",
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
					0, 1,
				},
				labelsOvervaluedFairUndervalued(),
				[]string{"dax_fund"},
				radioVals4,
				nil,
			)
			// gb.MainLabel = trl.S{
			// }
			gr := page.AddGrid(gb)
			gr.WidthMax("30rem")
			gr.Style = css.NewStylesResponsive(gr.Style)

			gr.Style.Desktop.StyleGridContainer.GapColumn = "0.2rem"
			gr.Style.Mobile.StyleGridContainer.GapColumn = "0.9rem" // force some wrapping in mobile

			gr.Style.Desktop.StyleGridContainer.GapRow = "0.5rem"

			gr.Style.Desktop.StyleBox.Position = "relative"
			gr.Style.Desktop.StyleBox.Left = "-1.1rem"
			gr.Style.Desktop.StyleBox.Left = "2rem"
			gr.Style.Mobile.StyleBox.Left = "0"
		}
	}

	//
	// page 4
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Währungen", "en": "Currencies"}
		page.Short = trl.S{"de": "Währungen", "en": "Currencies"}
		page.WidthMax("36rem")

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
				"de": "<b>7.</b> &nbsp; Folgende Währungen werden gegenüber dem Euro mittelfristig (<bx>6</bx>&nbsp;Mo.)",
				"en": "<b>7.</b> &nbsp; In the medium-term (<bx>6</bx>&nbsp;months), the following currencies compared to the Euro will",
			}
			gr := page.AddGrid(gb)
			_ = gr
		}
	}

	//
	// page 5
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Sektoren", "en": "Sectors"}
		page.Short = trl.S{"de": "Sektoren", "en": "Sectors"}
		page.WidthMax("36rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsImproveDeteriorateSectoral(),
				[]string{
					"sec_banks", "sec_insur", "sec_cars", "sec_chemi", "sec_steel", "sec_elect", "sec_mecha",
					// "sec_c",
					"sec_consu", "sec_const", "sec_utili", "sec_servi", "sec_telec", "sec_infor",
				},
				radioVals4,
				rowLabelsSectors,
			)
			gb.MainLabel = trl.S{
				"de": "<b>8.</b> &nbsp; Die Ertragslage der Unternehmen in Deutschland wird mittelfristig (<bx>6</bx>&nbsp;Mo.) in den folgenden Branchen ",
				"en": "<b>8.</b> &nbsp; In the medium-term (<bx>6</bx>&nbsp;months), the profit situation of German companies in the following sectors will",
			}
			gr := page.AddGrid(gb)
			_ = gr
		}
	}

	// log.Printf("q.Survey.MonthOfQuarter() is %v  (from %v - %v)", q.Survey.MonthOfQuarter(), q.Survey.Year, q.Survey.Month)

	var err error
	err = eachMonth1inQ(&q)
	if err != nil {
		return nil, fmt.Errorf("error adding month 1 per quarter: %v", err)
	}
	err = eachMonth2inQ(&q)
	if err != nil {
		return nil, fmt.Errorf("error adding month 2 per quarter: %v", err)
	}
	err = eachMonth3inQ(&q)
	if err != nil {
		return nil, fmt.Errorf("error adding specialQ3(): %v", err)
	}

	// err = special202106(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202106(): %v", err)
	// }
	// err = special202108(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202108(): %v", err)
	// }
	// err = special202110(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202110(): %v", err)
	// }
	// err = special202111a(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202111a(): %v", err)
	// }
	// err = special202111b(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202111b(): %v", err)
	// }
	// err = special202111c(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202111c(): %v", err)
	// }
	// err = special_2022_01(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special_2022_01(): %v", err)
	// }
	// err = special202202a(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202202a(): %v", err)
	// }
	// err = special202202b(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202202b(): %v", err)
	// }
	// err = special202203(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202203(): %v", err)
	// }
	// err = special202204(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202204(): %v", err)
	// }
	// err = special202205(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202205(): %v", err)
	// }
	// err = special202206(&q)
	// if err != nil {
	// 	return nil, fmt.Errorf("error adding special202206(): %v", err)
	// }

	err = special202212(&q)
	if err != nil {
		return nil, fmt.Errorf("error adding special202212(): %v", err)
	}

	//
	// page 7 - after seasonal
	// Finish questionnaire?  - one before last page
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "Abschluss", "en": "Finalization"}
		page.Short = trl.S{"de": "Abschluss", "en": "Finalization"}
		page.WidthMax("36rem")

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
				inp.Name = "comment"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = qst.RemainOpen
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Zugang bleibt bestehen.  Daten können in weiteren Sitzungen geändert/ergänzt werden. ",
					"en": "Leave the questionnaire open. Data  can be changed/completed&nbsp;in later sessions.     ",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "finished"
				rad.ValueRadio = qst.Finished
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "Fragebogen ist abgeschlossen und kann nicht mehr geöffnet werden. ",
					"en": "The questionnaire is finished. No more edits.                         ",
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
					"en": "By clicking, you will receive a summary of your answers.",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
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
		page.SuppressProgressbar = true

		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "ResponseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "site-imprint.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}
	return &q, nil

}
