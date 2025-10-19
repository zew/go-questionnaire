package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func fmt202511_0(q *QuestionnaireT, page *pageT) error {
	return fmt202511(q, page, 0)
}
func fmt202511_1(q *QuestionnaireT, page *pageT) error {
	return fmt202511(q, page, 1)
}

func fmt202511(q *QuestionnaireT, page *pageT, instance int) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"en": fmt.Sprintf("Experiment-Chart %v", instance),
		"de": fmt.Sprintf("Experiment chart %v", instance),
	}
	page.Label = trl.S{
		"en": "",
		"de": "",
	}
	// page.SuppressInProgressbar = true

	page.WidthMax("58rem")

	btmSpacers := 0

	grIdx := q.Version() % 2

	if instance < 10 {
		gr := page.AddGroup()
		gr.Cols = 3
		gr.BottomVSpacers = btmSpacers

		if grIdx == 0 {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `intro text treatment A`,
					"en": `todo`,
				}.OutlineHid("C24.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
		} else {
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `intro text treatment B`,
					"en": `todo`,
				}.OutlineHid("C24.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("param1_%v", instance)
			inp.Min = 0
			inp.Max = 100
			inp.Step = 5
			inp.MaxChars = 5
			inp.Response = "55"
		}
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("param2_%v", instance)
			inp.Min = 0
			inp.Max = 100
			inp.Step = 5
			inp.MaxChars = 5
			inp.Response = "15"
		}

	}

	// gr1 - hidden inputs saving the values from the echart
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = btmSpacers
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			// inp.Name = "simtool_instance"
			// we render this as textblock, to circumvent the uniqueness of fieldnames.
			// it only serves as read-only indicator for the embedded echart javascript
			inp.Label = trl.S{
				"de": fmt.Sprintf(`
					<div style='visible: none; height: 0.1rem; font-size: 4px;'> 
						<input type='hidden' value='%v'  name='simtool_instance' id='simtool_instance' />
					</div> 
				`, instance),
				"en": `todo`,
			}
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = fmt.Sprintf("sim_history_%v", instance)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = fmt.Sprintf("param1_bg_%v", instance)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Name = fmt.Sprintf("param2_bg_%v", instance)
		}

	}

	//
	// gr2 - embedding html with echart
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = btmSpacers
		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.DynamicFunc = "RenderStaticContent"
			// inp.DynamicFuncParamset = fmt.Sprintf("./experiment-1/inner-%d.html", grIdx)
			inp.DynamicFuncParamset = fmt.Sprintf("./experiment-1/inner-%d.html", 0)
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

	}

	return nil
}
