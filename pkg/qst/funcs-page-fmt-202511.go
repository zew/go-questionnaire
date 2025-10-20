package qst

import (
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

	page.WidthMax("58rem")

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
			inp.Name = "change_history_pg2"
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
	// {
	// 	inp := gr.AddInput()
	// 	inp.Type = "number"
	// 	inp.Name = fmt.Sprintf("param2_%v", instance)
	// 	inp.Min = 0
	// 	inp.Max = 100
	// 	inp.Step = 5
	// 	inp.MaxChars = 5
	// 	inp.Response = "15"
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
	// page.SuppressInProgressbar = true

	page.WidthMax("58rem")

	// grIdx := q.Version() % 2

	{
		gr := page.AddGroup()
		gr.Cols = 1
		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.DynamicFunc = "RenderStaticContent"
			inp.DynamicFuncParamset = "./experiment-1/pg3.html"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
	}

	return nil
}
