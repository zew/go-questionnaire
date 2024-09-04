package fmt

import fmtt "fmt"
import "strings"

import (
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func forecastDependingOnWinner(
	q *qst.QuestionnaireT,
	page *qst.WrappedPageT,
	lblMain trl.S,
	inpRump string,
) {

	lower := css.NewStylesResponsive(nil)
	lower.Desktop.StyleBox.Position = "relative"
	lower.Desktop.StyleBox.Top = "0.3rem"

	fourYears := []int{0, 1, 2, 4}

	// sq2
	{
		gr := page.AddGroup()
		const colsFirstCol = 8
		gr.Cols = colsFirstCol + 3*float32(len(fourYears))

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = lblMain
		}

		// first row
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "Jahr",
				"en": "Year",
			}
			inp.ColSpan = colsFirstCol
			inp.ColSpanLabel = 1
			inp.StyleLbl = lower
		}
		for idx := range fourYears {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": fmtt.Sprintf(" &nbsp;  %v", q.Survey.YearStr(idx+1)),
				"en": fmtt.Sprintf(" &nbsp;  %v", q.Survey.YearStr(idx+1)),
			}
			inp.ColSpan = 3
			inp.ColSpanLabel = 1
			inp.StyleLbl = lower
		}

		// second and third row
		for rowIdx, name := range []string{"Trump", "Harris"} {

			_ = rowIdx

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmtt.Sprintf("Prognose falls <i>%v</i> zum US-Präsidenten gewählt wird", name),
					"en": fmtt.Sprintf("Forecast if    <i>%v</i> is elected as US president", name),
				}
				inp.ColSpan = colsFirstCol
				inp.ColSpanLabel = 1

				inp.LabelPadRight()
				// inp.StyleLbl = lower
			}

			for colIdx := range fourYears {
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmtt.Sprintf("%v_%v_%v", inpRump, strings.ToLower(name), colIdx+1) //"p1_y1"
				// inp.Min = -0
				// inp.Max = +100
				inp.Min = -10
				inp.Max = +20
				inp.Validator = "inRange20"
				inp.MaxChars = 5
				inp.Step = 0.1

				// inp.Label = trl.S{
				// 	"de": q.Survey.YearStr(colIdx),
				// 	"en": q.Survey.YearStr(colIdx),
				// }
				inp.Suffix = trl.S{
					"de": "%",
					"en": "pct",
				}

				inp.ColSpan = 3
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 2

				// inp.StyleLbl = lower
			}

		}
	}

}

func special202409(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 9
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "Sonderfrage zur US Präsidentschaftswahl",
		"en": "Special questions: US presidential election",
	}
	page.Short = trl.S{
		"de": "Sonderfrage<br>US Präsidentschaftswahl",
		"en": "Special questions:<br>US presidential election",
	}
	page.WidthMax("48rem")

	indent := css.NewStylesResponsive(nil)
	indent.Desktop.StyleBox.Padding = "0  0.8rem  0  1.4rem"
	indent.Mobile.StyleBox.Padding = " 0  0.8rem  0  0.6rem"

	indent.Desktop.StyleBox.Margin = "0  0  0.95rem  0"

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
					Wir möchten Sie gerne zu Ihrer persönlichen Einschätzung 
					der Wahrscheinlichkeit einer Wahl von Donald Trump als nächster 
					Präsident der USA befragen und wie bestimmte Ereignisse 
					Ihre Erwartung beeinflusst haben. 
					Bitte geben Sie Ihre Antworten auf einer Skala von 
					0% (absolut unwahrscheinlich) bis 100% (absolut sicher) an.
					
					<br>
					<br>

					<b>
					Wie hoch war Ihrer Meinung nach die Wahrscheinlichkeit 
					mit der Donald Trump als nächster US-Präsident gewählt wird…
					</b>
				`,
				"en": `
					We would like to ask for your personal assessment of the probability 
					of Donald Trump being elected as the next US president 
					and how certain events affected your expectation. 
					
					Please state your answers on a scale from 
					0% (no chance) to 100% (absolutely certain).

					<br>
					<br>

					<b>
					In your opinion, how likely was it that Donald Trump 
					would be elected as the next US president…
					</b>


				`,
			}.Outline("1.")
			inp.ColSpanLabel = 1
			inp.ColSpan = gr.Cols
		}

		{
			lbls := []trl.S{
				{
					"de": `…nachdem sich Biden am 21. Juli 2024 aus dem Präsidentschaftsrennen zurückgezogen hat aber vor dem Democratic National Convention vom 19.-22. August 2024?`,
					"en": `…after Biden decided to step down from the presidential race on 21 July 2024 but before the Democratic National Convention from 19-22 August 2024?`,
				},
				{
					"de": `…nach dem Democratic National Convention vom 19.-22. August 2024 aber vor dem TV-Duell zwischen Trump und Harris am 10. September 2024?`,
					"en": `…after the Democratic National Convention from 19-22 August 2024 but before the TV debate between Trump and Harris on 10 September 2024?`,
				},
				{
					"de": `…TV-Duell zwischen Trump und Harris am 10. September 2024?`,
					"en": `…after the TV debate between Trump and Harris on 10 September 2024?`,
				},
			}

			for idx, lbl := range lbls {
				inp := gr.AddInput()
				inp.Type = "number"

				if false {
					inp.Label = lbl.Outline(fmtt.Sprintf("%v.", idx+1))
				}
				inp.Label = lbl

				inp.Name = fmtt.Sprintf("sq1_us_elect_%v", idx+1)

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 5
				inp.ColSpanControl = 1

				inp.Min = 0
				inp.Max = 100
				inp.Step = 1
				inp.MaxChars = 5
				inp.Suffix = trl.S{
					"de": "Prozent",
					"en": "percent",
				}
				inp.Placeholder = trl.S{
					"de": "00",
					"en": "00",
				}

				inp.Style = indent

			}

		}
	}

	lbl1 := trl.S{
		"de": `
			Wir möchten Sie nun gerne nach Ihren 
				<i>szenarienbasierten Punktprognosen</i> 
			für die Wachstumsrate des 
				<i>deutschen BIP</i> 
			in Abhängigkeit vom Ausgang der US-Präsidentschaftswahl fragen:
			
			<br>
			<xxsmall>Bitte geben die die jährliche Wachstumsrate in Prozent an.</xxsmall>
		`,
		"en": `
			We would like to ask you for your 
				<i>scenario-based point forecast</i> 
			of the growth rate of the 
				<i>German GDP</i> 
			depending on the outcome of the US presidential election:
			
			<br>
			<xxsmall>Please indicate the annual real GDP growth rate in percent.</xxsmall>
		`,
	}.Outline("2.")

	lbl2 := trl.S{
		"de": `
			Wir möchten Sie nun gerne nach Ihren 
				<i>szenarienbasierten Punktprognosen</i> 
			für die Wachstumsrate des 
				<i>Inflationsrate in Deutschland </i> 
			in Abhängigkeit vom Ausgang der US-Präsidentschaftswahl fragen:
			
			<br>
			<xxsmall>Bitte geben die die durchschnittliche jährliche Veränderung des HVPI in Prozent an.</xxsmall>
		`,
		"en": `
			We would like to ask you for your 
				<i>scenario-based point forecast</i> 
			of the growth rate of the 
				<i>annual inflation rate in Germany</i> 
			depending on the outcome of the US presidential election:
			
			<br>
			<xxsmall>Please indicate the annual average change of the HICP in percent.</xxsmall>
		`,
	}

	forecastDependingOnWinner(
		q,
		qst.WrapPageT(page),
		lbl1,
		"sq2_pp_y",
	)

	forecastDependingOnWinner(
		q,
		qst.WrapPageT(page),
		lbl2,
		"sq2_pp_inf",
	)

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		// gr.Style = css.NewStylesResponsive(gr.Style)
		// gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
					Welche ökonomischen und/oder politischen Maßnahmen 
					sollte die deutsche Regierung idealerweise 
					<i>jetzt schon</i> 
					(d.h. vor dem Ausgang der US-Präsidentschaftswahl) implementieren, 
					um die deutsche Wirtschaft vor potenziellen 
					adversen Effekten der Gesetzgebung durch die neue US-Regierung zu wappnen?
				`,
				"en": `
					Which economic and/or political measures should 
					the German government ideally implement 
					<i>right now</i> 
					(i.e., before the outcome of the US presidential election is known) 
					to safeguard the German economy against potential adverse 
					effects of policies implemented by the new US government? 					
				
				`,
			}.Outline("3.")
			inp.ColSpanLabel = 1
		}
		{
			inp := gr.AddInput()
			inp.Type = "textarea"
			inp.Name = "sq3_free"
			inp.MaxChars = 300
			inp.ColSpanLabel = 0
			inp.ColSpanControl = 1
		}
	}

	// sq4
	{
		lbls := []trl.S{
			{
				"de": `Stabilisierung globaler Finanzmärkte`,
				"en": `Stabilizing global financial markets`,
			},
			{
				"de": `Strengere finanzielle Bedingungen für US-Firmen `,
				"en": `Tighter financial conditions for US firms `,
			},
			{
				"de": `Stärkung des US-Dollar gegenüber dem Euro `,
				"en": `Stronger US Dollar relative to the Euro `,
			},
			{
				"de": `Aufrecherhaltung der Zentralbankunabhängigkeit `,
				"en": `Maintaining independence of the Federal Reserve `,
			},
			{
				"de": `Beschäftigungsanstieg relativ zum aktuellen Niveau`,
				"en": `Increase in employment relative to current levels`,
			},
			{
				"de": `Preisstabilität`,
				"en": `Price stability`,
			},
			{
				"de": `Fiskalische Dominanz`,
				"en": `Fiscal dominance `,
			},
			{
				"de": `Verbesserung der deutsch-amerikanischen wirtschaftlichen Beziehungen `,
				"en": `Improving economic relations between the US and Germany`,
			},
			{
				"de": `Das Finden einer Vereinbarung zum Beenden des Ukraine-Krieges `,
				"en": `Finding an agreement to end the Ukraine War `,
			},
		}

		inps := []string{
			"sq4_1_stabfinmark",
			"sq4_2_finrestrict",
			"sq4_3_dollarappr",
			"sq4_4_independence",
			"sq4_5_employm",
			"sq4_6_infl",
			"sq4_7_dominance",
			"sq4_8_relations",
			"sq4_9_ukraine",
		}

		colTpl := []float32{
			4.0, 1,
			0.0, 1,
			0.0, 1,
			0.0, 1, // no answer slightly apart
		}

		hdrs := []trl.S{
			{
				"de": "Kamala Harris",
				"en": "Kamala Harris",
			},
			{
				"de": "Donald Trump",
				"en": "Donald Trump",
			},
			{
				"de": "keiner der beiden",
				"en": "neither of the two",
			},
			{
				"de": "gleich wahrscheinlich",
				"en": "equally by both",
			},
		}
	

		gb := qst.NewGridBuilderRadios(
			colTpl,
			hdrs,
			inps,
			radioVals4,
			lbls,
		)
		gb.MainLabel = trl.S{
			"de": `
				Geben Sie bitte für die folgenden ökonomischen, 
				politischen oder geldpolitischen Ergebnisse an, 
				ob diese vermutlich eher unter der Präsidentschaft 
				von Donald Trump oder von Kamala Harris erreicht wird, 
				von keinem der beiden oder gleich wahrscheinlich von beiden.
			`,
			"en": `
				Please indicate for the following economic, 
				political, or monetary policy outcomes whether 
				they are more likely to be achieved during a presidency 
				of either Donald Trump, Kamala Harris, neither of the two, 
				or equally by both.
			`,
		}.Outline("4.")
		gr := page.AddGrid(gb)
		_ = gr
	}

	return nil

}
