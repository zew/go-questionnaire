package qst

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// func fmt202511Pages(q *QuestionnaireT, page *pageT) error {
// 	return nil
// }

func fmt202511Pg2(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	_, found := ForecastData(q.UserIDInt())
	page.SuppressInProgressbar = found
	// page.NoNavigation = !found
	page.NavigationCondition = "fmt202511Include"

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen II",
		"en": "Growth Prospects II",
	}

	page.WidthMax("62rem")

	grIdx := q.Version() % 2

	dta := addForecastData(q, page)
	mnth, _ := dta["month_de"].(string)
	if q.LangCode == "en" {
		mnth = strings.ReplaceAll(mnth, "Februar", "February")
		mnth = strings.ReplaceAll(mnth, "Mai", "May")
	}

	//
	// gr1
	// 	hidden inputs saving the values from the echart
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = "param1_pg2_bg"
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = "param2_pg2_bg"
		}
		{
			inp := gr.AddInput()
			inp.Type = "javascript-block"
			inp.Name = "warnEmpty"
		}
	}

	//
	// gr2
	// 	visible input fields - "foreground"
	if grIdx == 0 || true {

		gr := page.AddGroup()
		gr.Cols = 3
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": fmt.Sprintf(`
					
					Zuletzt haben Sie im 
						<!-- <u> [August|May|Februar] 2025</u> -->
						<i> %v 2025</i> 
					eine Prognose 
					für das Quartalswachstum in Q4 2025 angegeben.
					<br><br>

					Was denken Sie über die Prognosen der anderen Teilnehmenden in der damaligen Befragung?
					
					<br><br>
										
					`, mnth),
				"en": fmt.Sprintf(`
					
					Previously, 
						in <i> %v 2025</i>, 
					you provided a forecast for quarterly growth for Q4 2025.
					<br><br>

					What do you think about the forecasts of the other participants in that previous survey?
					
					<br><br>

					
					`, mnth),
			}.Outline("4.")
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": fmt.Sprintf(`
					Der 
					<i>Anteil unter allen Befragten</i>
					 (in %%),
					die im %v 2025 ein 
					<i>niedrigeres</i> 
					Wachstum für Q4 2025 als Sie angegeben haben, 
					lag bei...										
					`,
					mnth,
				),

				"en": fmt.Sprintf(`
					The 
					<i>share of all respondents</i> 
					in %%
					who, in %v 2025, stated a 
					<i>lower</i> 
					growth rate for Q4 2025 than you, amounted to…
				`, mnth),
			}.OutlineHid("3b.")
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
		}
	} else {
	}

	//
	// {
	// 	inp := gr.AddInput()
	// 	inp.Type = "number"
	// 	inp.Name = fmt.Sprintf("param1_%v", instance)
	// 	inp.Min = 0
	// 	inp.Max = 100
	// 	inp.Step = 5
	// 	inp.MaxChars = 5
	// 	inp.Response = "55"
	// }

	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.DynamicFunc = "RenderStaticContent"
			inp.DynamicFuncParamset = "./experiment-1/pg2.html"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	ChangeHistoryJS(q, page)

	return nil
}

func fmt202511Pg3(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	_, found := ForecastData(q.UserIDInt())
	page.SuppressInProgressbar = found
	// page.NoNavigation = !found
	page.NavigationCondition = "fmt202511Include"

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen III",
		"en": "Growth Prospects III",
	}

	page.WidthMax("62rem")

	addForecastData(q, page)
	addingThreeCharts(q, page, 3)

	ChangeHistoryJS(q, page)

	return nil
}

func fmt202511Pg4(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	_, found := ForecastData(q.UserIDInt())
	page.SuppressInProgressbar = found
	// page.NoNavigation = !found
	page.NavigationCondition = "fmt202511Include"

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen IV",
		"en": "Growth Prospects IV",
	}

	page.WidthMax("62rem")

	addForecastData(q, page)
	// addingThreeCharts(q, page, 4)

	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
					<p style='font-size:120%; font-weight:bold; margin-top: 0;'>
						Ihre Meinung zu aktuellen Aussichten
					</p>
											
					<p>
						Denken Sie nun bitte erneut über das deutsche Wirtschaftswachstum nach und
						geben Sie Ihre Einschätzung ab.
					</p>
					
					<small>
						Bitte nicht-annualisiertes Quartalswachstum des realen &amp; saisonbereinigten BIP angeben.
					</small>
					<br>
					<br>

				`,
				"en": `
					<p style='font-size:120%; font-weight:bold; margin-top: 0;'>
						Your view on the current outlook
					</p>
											
					<p>
						Please think again about Germany's economic growth prospects and provide your assessment.
					</p>
					
					<small>
						Please indicate non-annualized quarterly real & seasonally adjusted GDP growth.
					</small>
					<br>
					<br>


				`,
			}
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
					Was glauben Sie: Wie hoch wird das Wirtschaftswachstum ausfallen

				`,
				"en": `
					What do you think: What will the economic growth rate be?
				`,
			}.Outline("5a.")
		}

		sLbl1 := css.NewStylesResponsive(nil)
		sLbl1.Desktop.StyleGridItem.JustifySelf = "end"
		sLbl1.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
		sLbl1.Mobile.StyleBox.Padding = " 0 2.7rem 0 0.2rem"

		sLbl2 := *sLbl1
		sLbl2.Mobile.StyleGridItem.JustifySelf = "start"
		sLbl2.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
		sLbl2.Mobile.StyleBox.Padding = " 0 1.5rem 0 0.2rem"

		// row 1 - four quarters - label
		//   removed
		// row 2 - four quarters - inputs
		for i := 0; i < 4; i++ {
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("ssq5a_%v", i+1)
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

	}

	// 5b
	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
					Was glauben Sie: Wie hoch ist die durchschnittliche Wirtschaftswachstumsprognose 
					<i>unter allen Befragten in der aktuellen Befragung?</i>`,
				"en": `
					What do you think: What is the average economic growth forecast 
					<i>among all respondents</i> in the <i>current survey</i>?
				`,
			}.Outline("5b.")
		}

		sLbl1 := css.NewStylesResponsive(nil)
		sLbl1.Desktop.StyleGridItem.JustifySelf = "end"
		sLbl1.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
		sLbl1.Mobile.StyleBox.Padding = " 0 2.7rem 0 0.2rem"

		sLbl2 := *sLbl1
		sLbl2.Mobile.StyleGridItem.JustifySelf = "start"
		sLbl2.Desktop.StyleBox.Padding = "0 0.2rem 0 0"
		sLbl2.Mobile.StyleBox.Padding = " 0 1.5rem 0 0.2rem"

		// row 1 - four quarters - label
		//   removed
		// row 2 - four quarters - inputs
		for i := 0; i < 4; i++ {
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("ssq5b_%v", i+1)
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

	}

	ChangeHistoryJS(q, page)

	return nil

}

func fmt202511Pg5(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	_, found := ForecastData(q.UserIDInt())
	page.SuppressInProgressbar = found
	// page.NoNavigation = !found
	page.NavigationCondition = "fmt202511Include"

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen V",
		"en": "Growth Prospects V",
	}

	page.WidthMax("62rem")

	addForecastData(q, page)
	// addingThreeCharts(q, page, 6)

	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.WidthMax("40rem")
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
						Wie treffen Sie Ihre Prognosen am ehesten?

						<script>
							document.addEventListener("DOMContentLoaded", () => {
							const element = document.getElementById("ssq6_1");

							if (element) {
								element.focus();
								element.scrollIntoView({ behavior: "smooth", block: "center" });
							}
							});
						</script>
				`,
				"en": `
					How do you usually make your forecast?
				`,
			}.Outline("6.")
		}

		lbls := []trl.S{
			{
				// "de": "Bauchgefühl",
				"de": "erfahrungsbasierte Beurteilung",
				"en": "judgement-based",
			},
			{
				// "de": "modellbasiert",
				"de": "modellbasierte Berechnung",
				"en": "model-based",
			},
			{
				"de": "anders",
				"en": "other",
			},
		}

		for i := 0; i < 3; i++ {
			{
				inp := gr.AddInput()
				// inp.Type = "checkbox"
				inp.Type = "radio"
				inp.ValueRadio = fmt.Sprintf("%v", i+1)
				inp.Name = "ssq6"
				inp.ColSpan = 12
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 11
				inp.Label = lbls[i]
				inp.ControlFirst()
				inp.ControlRight()
				if i == 2 {
					inp.ColSpan = 3
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 2
				}
			}

			if i == 2 {
				inp := gr.AddInput()
				inp.Type = "text"
				inp.MaxChars = 36
				inp.Name = "ssq6_other"
				inp.ColSpan = 9
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}

		}
	}

	//
	{
		gr := page.AddGroup()
		gr.Cols = 12
		// gr.WidthMax("44rem")
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Fanden Sie die graphisch bereitgestellten Informationen über die Wirtschaftswachstumsprognosen nützlich?
				`,
				"en": `
					Did you find the graphical information on the economic growth forecasts useful?
				`,
			}.Outline("7.")
		}
		{
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = "ssq7"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 11
			inp.Label = trl.S{
				"de": "ja",
				"en": "yes",
			}
			inp.ControlFirst()
			inp.ControlRight()
		}

	}

	ChangeHistoryJS(q, page)

	return nil

}
