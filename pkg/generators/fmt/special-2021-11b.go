package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202111b(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 11 {
		return nil
	}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage: Inflation, Prognosetreiber und Geldpolitik 3",
			"en": "Special: Inflation, forecast drivers and monetary policy 3",
		}
		page.Short = trl.S{
			"de": "Inflation,<br>Geldpolitik 3",
			"en": "Inflation,<br>Mon. Policy 3",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

		// gr1

		/*
			todo
			the “central 90 percent confidence interval”
			directly under the Forecast of the ECB’s main refinancing rate

		*/
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
					"de": fmt.Sprintf(`
						<b>4.</b> &nbsp; 
							Den <b>Hauptrefinanzierungssatz</b> der EZB (seit %v) erwarte ich auf Sicht von`,
						latestECBRate),
					"en": fmt.Sprintf(`
						<b>4.</b> &nbsp; 
							Forecast of the ECB's main refinancing rate (since %v) `,
						latestECBRate),
				}
			}

			// row-2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 3
				inp.Label = trl.S{
					"de": "6&nbsp;Monaten",
					"en": "in&nbsp;6&nbsp;months",
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
					"en": " in&nbsp;24&nbsp;months",
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
	} // special page 3

	return nil
}
