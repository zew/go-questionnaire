package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202504(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2025 || q.Survey.Month != 4 {
		return nil
	}

	page := q.AddPage()
	yearsEffective := []int{0, 1, 2}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 0.8rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	page.Label = trl.S{
		"de": "Sonderfragen: Konsequenzen der protektionistischen Handelspolitik der USA",
		"en": "Special questions: Consequences of the US trade protectionism and tariffs",
	}
	page.Short = trl.S{
		"de": "USA Zölle",
		"en": "Tariffs",
	}

	page.WidthMax("54rem")

	{
		gr := page.AddGroup()
		gr.Cols = 3*float32(len(yearsEffective)) + 2
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Width = "80%"
		gr.Style.Mobile.StyleBox.Width = "100%"

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Wir möchten Sie nach Ihren Punktprognosen für die Wachstumsrate des realen BIP in Deutschland, dem Euroraum und den USA befragen. 
					Wie wird sich Ihrer Meinung nach das <i>reale Wachstum</i> entwickeln in den kommenden Jahren entwickeln?
				`,
				"en": `
					We would like to ask you about your point forecasts for the real GDP growth rate for Germany, the euro zone and the USA. 
					How do you think <i>growth</i> will develop in the coming years?
				`,
			}.Outline("1.")
		}

		//
		pointForecastRegion := []string{
			"ger",
			"eu",
			"us",
		}
		pointForecastLabels := []trl.S{
			{
				"de": "Deutschland",
				"en": "Germany",
			},
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "USA",
				"en": "USA",
			},
		}

		for idx1, region := range pointForecastRegion {

			rowLbl := pointForecastLabels[idx1]
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = rowLbl
				inp.ColSpan = 2
				// inp.StyleLbl = lblStyleRight
			}

			for idx2 := range yearsEffective {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("pp_y1_%v_jp%v", idx2, region)
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.1
				inp.Label = trl.S{
					"de": q.Survey.YearStr(idx2),
					"en": q.Survey.YearStr(idx2),
				}
				inp.Suffix = trl.S{
					"de": "%",
					"en": "pct",
				}

				inp.ColSpan = 3

				inp.ColSpanLabel = 2
				inp.ColSpanControl = 2

				inp.StyleLbl = lblStyleRight
			}

		}

	}

	{
		gr := page.AddGroup()
		gr.Cols = 3 + 5 + 1 + 2
		gr.BottomVSpacers = 4

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Verglichen mit Ihren letzten prognostizierten Wachstumsraten des realen BIP (vor der Ankündigung der neuen US-Zölle am 2. April), 
					sind diese Punktprognosen gleich geblieben, höher oder niedriger ausgefallen? 
					Bitte geben Sie an, ob 
					<i>Ihre Revision</i> 
					positiv oder negativ ist und wie groß diese ist (in Prozentpunkten).
					<br>
				`,
				"en": `
					Compared with your most recent forecasted real GDP growth rates (before the new US tariffs were announced on April 2), 
					are these point forecasts the same, higher or lower? Please indicate whether 
					<i>your revision</i> 
					is positive or negative, and how large it is in percentage points.
					<br>
				`,
			}.Outline("2.")
		}

		// first row with headers
		//    top-left
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "&nbsp;",
				"en": "&nbsp;",
			}
			inp.ColSpan = 3
		}
		lblHeaders := []trl.S{
			{
				"de": "+",
				"en": "+",
			},
			{
				"de": "-",
				"en": "-",
			},
			{
				"de": "0-0.5<br>PP",
				"en": "0-0.5<br>PP",
			},
			{
				"de": ">0.5-1<br>PP",
				"en": ">0.5-1<br>PP",
			},
			{
				"de": ">1<br>PP",
				"en": ">1<br>PP",
			},
			{
				"de": "gleich<br> geblieben",
				"en": "stayed<br>the same ",
			},
			{
				"de": "keine Angabe",
				"en": "don't know",
			},
		}

		stlShiftDown := css.NewStylesResponsive(nil)
		stlShiftDown.Desktop.StyleBox.Position = "relative"
		stlShiftDown.Desktop.StyleBox.Top = "0.95rem"

		for idx3 := 0; idx3 < 7; idx3++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblHeaders[idx3].Italic()
			inp.ColSpan = 1
			if idx3 == 5 {
				inp.ColSpan = 2
			}
			inp.StyleLbl = stlShiftDown
			inp.LabelCenter()
			inp.LabelBottom()
		}

		//
		//
		pointForecastRegion := []string{
			"ger",
			"eu",
			"us",
		}
		pointForecastLabels := []trl.S{
			{
				"de": "Deutschland",
				"en": "Germany",
			},
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "USA",
				"en": "USA",
			},
		}

		for idx1, region := range pointForecastRegion {

			rowLbl := pointForecastLabels[idx1]
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = rowLbl.Bold()
				inp.ColSpan = gr.Cols
			}

			for idx2, year := range yearsEffective {
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": q.Survey.YearStr(year),
						"en": q.Survey.YearStr(year),
					}
					inp.ColSpan = 3
				}
				for idx3 := 0; idx3 < 7; idx3++ {

					inp := gr.AddInput()
					inp.Type = "checkbox"
					inp.Name = fmt.Sprintf("pprev_y%v_%v_%v", idx2, region, idx3)
					inp.ColSpan = 1
					if idx3 == 5 {
						inp.ColSpan = 2
					}
					inp.ControlCenter()
				}

			}

		}

	}

	return nil

}
