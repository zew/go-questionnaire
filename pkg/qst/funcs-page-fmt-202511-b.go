package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func addingThreeCharts(q *QuestionnaireT, page *pageT) error {

	// grIdx := q.Version() % 2

	//
	// gr1
	// 	hidden inputs - making previous data from previous page available as hidden field
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0

		// prevRsp = strings.TrimSpace(prevRsp)
		// log.Printf("%v value is     '%v'", prevInp.Name, prevRsp)

		/*
			We want to make an input from previous page available.
				q.ByName("param1_pg2_bg")

			This yields the correct number.
			But the following fails:

				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "user_share_prev"
				inp.Response = rsp

			inp.Response is overwritten later - when  dynamic page values are set.
			Thus we
		*/
		prevInp := q.ByName("param1_pg2_bg")
		prevRsp := prevInp.Response
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": fmt.Sprintf(
					`<input type='hidden' name='user_share_prev' id='user_share_prev'  value='%v' `,
					prevRsp,
				),
			}
			inp.Label["en"] = inp.Label["de"]
			inp.ColSpan = gr.Cols
		}

	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
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
