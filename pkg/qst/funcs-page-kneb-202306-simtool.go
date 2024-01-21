package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func kneb202306simtool0(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 0)
}
func kneb202306simtool1(q *QuestionnaireT, page *pageT) error {
	return kneb202306simtool(q, page, 1)
}

// Dimension-1
// The same page is shown in two variations depending on version/group index.
// The variations are neutral frame and financal frame (nf, ff).
// Distinction is only in the name JS config.
// Dimension-2
// The same page is shown several times (instance).
// Validations differ.
func kneb202306simtool(q *QuestionnaireT, page *pageT, instance int) error {

	page.Groups = nil // dynamically recreate the groups

	page.Label = trl.S{
		"en": "Experiment-Chart",
		"de": "Experiment chart",
	}
	page.Label = trl.S{
		"en": "",
		"de": "",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("58rem")

	// group index determines nf or ff
	grIdx := q.Version() % 2

	validationType := fmt.Sprint(instance)
	if instance > 0 {
		validationType = "1"
	}

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			// inp.Name = "simtool_instance"
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
			inp.Validator = fmt.Sprintf("kneb_simtool_inst_%v", validationType)
			inp.Name = fmt.Sprintf("sparbetrag_bg_%v", instance)
		}

		// inversely related
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Validator = fmt.Sprintf("kneb_simtool_inst_%v", validationType)
			inp.Name = fmt.Sprintf("share_safe_bg_%v", instance)
		}
		{
			inp := gr.AddInput()
			inp.Type = "hidden"
			inp.Validator = fmt.Sprintf("kneb_simtool_inst_%v", validationType)
			inp.Name = fmt.Sprintf("share_risky_bg_%v", instance)
		}
	}

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.DynamicFunc = "RenderStaticContent"
			inp.DynamicFuncParamset = fmt.Sprintf("./echart/inner-%d.html", grIdx)
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}

		//
		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "knebRenameNext" // js filename

			s1 := trl.S{
				"de": "Werte speichern und weiter",
				"en": "todo",
			}
			if instance == 0 {
				s1 = trl.S{
					"de": "Weiter",
					"en": "todo",
				}
			}

			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}

			inp.JSBlockStrings = map[string]string{}
			// inp.JSBlockStrings["pageID"] = fmt.Sprintf("pg%02v", len(q.Pages)-1)

		}
	}

	return nil
}
