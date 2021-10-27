package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202111c(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 11 {
		return nil
	}

	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
		page.Label = trl.S{
			"de": "Sonderfrage: Finanzmarktreport",
			"en": "Special:     Finanzmarktreport",
		}
		page.Short = trl.S{
			"de": "Finanzmarkt<br>report",
			"en": "Finanzmarkt<br>report",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `<p>Als Teilnehmerin bzw. Teilnehmer an unserer Umfrage 
						erhalten Sie von uns monatlich den ZEW-Finanzmarktreport. 
						Mit dem ZEW-Finanzmarktreport möchten wir für Sie einen Mehrwert schaffen. 
						Damit uns dies gelingt, würden wir gerne von Ihnen wissen, 
						wie Sie den ZEW-Finanzmarktreport nutzen 
						und was wir Ihrer Meinung nach verbessern können.</p>
					`,
					"en": `todo`,
				}
				inp.ColSpanLabel = 1
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1

			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style = css.DesktopWidthMaxForGroups(gr.Style, "26rem")

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `<p><b>1.</b> 
						Lesen Sie den ZEW-Finanzmarktreport? </p>`,
					"en": `todo`,
				}
				inp.ColSpanLabel = 1
			}

			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "fmr_frequency"
				rad.ValueRadio = "never"
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "&nbsp; &nbsp; &nbsp;	<b>a.</b> &nbsp; Nein, nie",
					"en": "&nbsp; &nbsp; &nbsp;	<b>a.</b> &nbsp; todo",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "fmr_frequency"
				rad.ValueRadio = "infrequently"
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "&nbsp; &nbsp; &nbsp;	<b>b.</b> &nbsp; Ja, unregelmäßig",
					"en": "&nbsp; &nbsp; &nbsp;	<b>b.</b> &nbsp; todo",
				}
			}
			{
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "fmr_frequency"
				rad.ValueRadio = "regularly"
				rad.ColSpan = 1
				rad.ColSpanLabel = 6
				rad.ColSpanControl = 1
				rad.Label = trl.S{
					"de": "&nbsp; &nbsp; &nbsp;	<b>c.</b> &nbsp; Ja, regelmäßig",
					"en": "&nbsp; &nbsp; &nbsp;	<b>c.</b> &nbsp; todo",
				}
			}

			//
			//
			//
			// gr2
			rowLabelsAreas := []trl.S{
				{
					"de": "Deutschland",
					"en": "Germany",
				},
				{
					"de": "Eurogebiet",
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
				{
					"de": "Sonder&shy;frage",
					"en": "Specials",
				},
			}
			{
				gb := qst.NewGridBuilderRadios(
					columnTemplate6,
					labelsPositiveNeutralNegative(),
					[]string{
						"fmr_germany", "fmr_euroarea", "fmr_us", "fmr_china", "fmr_specials",
					},
					radioVals6,
					rowLabelsAreas,
				)
				gb.MainLabel = trl.S{
					"de": "<b>2.</b> Wie beurteilen Sie die einzelnen Themenblöcke, die im ZEW-Finanzmarktreport behandelt werden?",
					"en": "<b>2.</b> todo",
				}
				gr := page.AddGrid(gb)
				gr.OddRowsColoring = true
			}

		}

		//
		//
		//
		// gr3
		gr := page.AddGroup()
		colWidths := []float32{4, 3, 3, 3}
		for _, cw := range colWidths {
			gr.Cols += cw
		}

		//
		colWidths = []float32{1, 1, 1, 1}
		gr.Cols = 4
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.AutoFlow = "row"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = "4fr   2fr  3fr 3fr"
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = " 1.5fr 1.3fr  3fr 3fr"

		gr.Style.Desktop.StyleGridContainer.GapRow = "1.2rem"

		gr.Style.Mobile.StyleGridContainer.GapColumn = "0"

		// gr.BottomVSpacers = 02

		// header row
		colHeaders := []trl.S{
			{
				"de": " &nbsp;",
				"en": " &nbsp;",
			},
			{
				"de": "Keine zu&shy;sätz&shy;lich&shy;en In&shy;for&shy;ma&shy;tion&shy;en",
				"en": "todo",
			},
			{
				"de": "Eine Grafik",
				"en": "todo",
			},
			{
				"de": "Eine Grafik und mehr Text",
				"en": "todo",
			},
		}
		for col, colHeader := range colHeaders {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = colHeader
				inp.ColSpan = colWidths[col]
				inp.ColSpanLabel = 1

				inp.StyleLbl = css.TextStart(inp.StyleLbl)
				inp.Style = css.NewStylesResponsive(inp.Style)
				if col != 1 {
					inp.Style.Desktop.StyleBox.Padding = "0 0 0 0.4rem"
					// inp.LabelPadRight()
				} else {
					inp.Style.Desktop.StyleBox.Padding = "0 0.4rem 0 0.4rem"
					inp.Style.Mobile.StyleBox.Padding = "0 0 0 0"
				}
			}
		}

		// body rows
		rowInputNames := []string{
			"inflation",
			"rates",
			"stocks",
			"exchanger",
			"sectors",
			"specials",
		}
		rowLabelChapters := []trl.S{
			{
				"de": "Inflation (Frage&nbsp;3)",
				"en": "todo",
			},
			{
				"de": "Kurz- und langfristige Zinsen (Fragen 4 und 5)",
				"en": "todo",
			},
			{
				"de": "Aktien&shy;märkte (Fragen 6a-6c)",
				"en": "todo",
			},
			{
				"de": "Wechselkurse (Frage&nbsp;7)",
				"en": "todo",
			},
			{
				"de": "Ertrags&shy;lage deutscher Unternehmen nach Branche (Frage&nbsp;8)",
				"en": "todo",
			},
			{
				"de": "Sonder&shy;frage",
				"en": "todo",
			},
		}

		frequencies := []string{
			"month",
			"quarter",
			"halfyr",
			"year",
		}
		frequencyLabels := []trl.S{
			{
				"de": "Monatl.",
				"en": "Monthly",
			},
			{
				"de": "Quartal",
				"en": "Quartrly",
			},
			{
				"de": "Halbjrl.",
				"en": "Biyearly",
			},
			{
				"de": "Jährlich",
				"en": "Yearly",
			},
		}

		addRowClosure := func(colName, inpName string) {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "<div> <!-- opening --> ",
					"en": "<div> <!-- opening --> ",
				}
				inp.ColSpan = 0.1
				inp.ColSpan = colWidths[2] // 2 and 3 identical
				inp.ColSpanLabel = 0.1
				inp.ColSpanLabel = 1
			}

			for idx2, fr := range frequencies {
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Type = "radio"
				inp.ValueRadio = fr
				inp.Name = fmt.Sprintf("fmr_%v_%v", colName, inpName)
				inp.Label = frequencyLabels[idx2]
				inp.ColSpan = colWidths[2] // 2 and 3 identical
				inp.ColSpan = 1

				// inp.ColSpanLabel = 4
				// inp.ColSpanControl = 1
				inp.ControlFirst()
				inp.StyleLbl.Desktop.Padding = "0 0 0 0.5rem"
				inp.StyleLbl.Mobile.Padding = "0 0 0 0"

				inp.Style = css.NewStylesResponsive(inp.Style)
				inp.Style.Desktop.Padding = "0 0 0 0.9rem"
				inp.Style.Mobile.Padding = "0 0 0 0.6rem"

			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": " <!-- closing --></div> ",
					"en": " <!-- closing --></div> ",
				}
				inp.ColSpan = 0.1
				inp.ColSpanLabel = 0.1
				inp.ColSpanLabel = 1
			}

		}
		for rowIdx, inpName := range rowInputNames {

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = rowLabelChapters[rowIdx]
				inp.ColSpan = colWidths[0]
				inp.ColSpanLabel = 1

			}
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("fmr_%v_nomore", inpName)
				inp.Label = trl.S{
					"de": `&nbsp;`,
					"en": `&nbsp;`,
				}
				inp.ColSpan = colWidths[1]
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1

				inp.ControlFirst()
				// inp.Style = css.ItemStartCA(inp.Style)
				// inp.Style.Desktop.StyleBox.Padding = "0 0.2rem 0 0"

			}
			addRowClosure("graphics", inpName)
			addRowClosure("grandtxt", inpName)

		}

	} // special page 4

	return nil
}
