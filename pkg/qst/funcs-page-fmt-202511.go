package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func fmt202511Pg2(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen II",
		"en": "todo %v",
	}
	// page.SuppressInProgressbar = true

	page.WidthMax("62rem")

	grIdx := q.Version() % 2

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
			inp.Name = "history_stack_pg2"
		}
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
	}

	//
	// gr2
	// 	visible input fields - "foreground"
	if grIdx == 0 || true {

		gr := page.AddGroup()
		gr.Cols = 3
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
					
					Zuletzt haben Sie im 
						<u>August [May/Februar] 2025</u> 
					eine Prognose 
					für das Quartalswachstum in Q4 2025 angegeben.
					<br><br>
					Was denken Sie über Prognosen der anderen Teilnehmer in der damaligen Befragung?<br><br>
										
					`,
				"en": `todo`,
			}
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
					
					Der Anteil unter allen Befragten, 
					die im August 2025 ein niedrigeres Wachstum für Q4 2025 als Sie angegeben haben, 
					lag bei...										
					`,

				"en": `todo`,
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

	return nil
}

func fmt202511Pg3(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen III",
		"en": "todo %v",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("62rem")

	addingThreeCharts(q, page, 3)

	return nil
}

func fmt202511Pg4(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen IV",
		"en": "todo %v",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("62rem")

	addingThreeCharts(q, page, 4)

	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.BottomVSpacers = 2
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




				`,
				"en": `
					todo
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


						<script>
							document.addEventListener("DOMContentLoaded", () => {
							const element = document.getElementById("ssq5a_1");

							if (element) {
								element.focus();
								element.scrollIntoView({ behavior: "smooth", block: "center" });
							}
							});
						</script>

				`,
				"en": `
					todo
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
		for i := 0; i < 3; i++ {
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

	return nil
}

func fmt202511Pg5(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen V",
		"en": "todo %v",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("62rem")

	addingThreeCharts(q, page, 5)

	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
						Was glauben Sie: Wie hoch ist die durchschnittliche Wirtschaftswachstumsprognose 
						<br>
						<i>unter allen Befragten in der aktuellen Befragung?</i>



						<script>
							document.addEventListener("DOMContentLoaded", () => {
							const element = document.getElementById("ssq5b_1");

							if (element) {
								element.focus();
								element.scrollIntoView({ behavior: "smooth", block: "center" });
							}
							});
						</script>
				`,
				"en": `
					todo
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
		for i := 0; i < 3; i++ {
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

	return nil

}

func fmt202511Pg6(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Wachtsumschancen VI",
		"en": "todo %v",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("62rem")

	addingThreeCharts(q, page, 6)

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
						Wie treffen Sie Ihre Vorhersagen am ehesten?

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
					todo
				`,
			}.Outline("6.")
		}

		lbls := []trl.S{
			{
				"de": "Bauchgefühl",
				"en": "todo",
			},
			{
				"de": "modellbasiert",
				"en": "todo",
			},
			{
				"de": "anders",
				"en": "todo",
			},
		}

		for i := 0; i < 3; i++ {
			{
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = fmt.Sprintf("ssq6_%v", i+1)
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
		gr.WidthMax("40rem")
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `
					Möchten Sie über die Studienergebnisse mit 
					Ihren Angaben zu den deutschen Wachstumschancen per Email informiert werden?
				`,
				"en": `
					todo
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
				"de": "Ja",
				"en": "todo",
			}
			inp.ControlFirst()
			inp.ControlRight()
		}

	}

	return nil

}
