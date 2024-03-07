package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/qstcp/cpfmt"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202403(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2024 && q.Survey.Month == 3
	if !cond {
		return nil
	}

	page := q.AddPage()

	page.Label = trl.S{
		"de": "Fragen zur Transition der Wirtschaft zur Klimaneutralität",
		"en": "Questions about climate transition",
	}
	page.Short = trl.S{
		"de": "Transition zur<br>Klimaneutralität",
		"en": "Questions about<br>climate transition",
	}
	page.WidthMax("75rem")

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1

			inp.Label = trl.S{
				"de": `
					Für wie wahrscheinlich halten Sie es aus technologischer Sicht, 
					dass eine Mehrheit der Unternehmen aus den folgenden Branchen zu den folgenden Zeitpunkten klimaneutral wird?
					<br>
					<small>Kategorien: --: sehr unwahrscheinlich, -: unwahrscheinlich, +: wahrscheinlich, ++: sehr wahrscheinlich </small>
				`,
				"en": `
					What do you think how likely it is, from a technological standpoint, 
					that a majority of firms from the following sectors will become climate-neutral by the following years?
					<br>
					<small>Categories: --: very unlikely, -: unlikely, +: likely, ++: very likely </small>
				
				`,
			}.Outline("1.")
		}
	}
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "dyn-composite"
			inp.DynamicFunc = fmt.Sprintf("Special202403QS1__%v__%v", 0, 0)
			inp.DynamicFuncParamset = ""
			inp.ColSpanControl = 1
		}

		_, inputNames, _ := cpfmt.Special202403QS1(q, 0, 0, true)
		for _, inpName := range inputNames {
			inp := gr.AddInput()
			inp.Type = "dyn-composite-scalar"
			inp.Name = inpName
		}

	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1

			inp.Label = trl.S{
				"de": `
					Wie hoch schätzen Sie den wirtschaftlichen Nutzen und die Kosten des Übergangs zur Klimaneutralität für Unternehmen aus den folgenden Sektoren insgesamt ein? 
					<br>
					<small>Kategorien:  0: irrelevant, +: niedrig, ++: mittel, +++: groß, ++++: sehr groß</small>
				`,
				"en": `
					How significant in economic terms do you think will the benefits and costs of the transition to climate-neutrality be for firms from the following sectors? 
					<br>
					<small>Categories:  0: insignificant, +: low significance, ++: medium significance, +++: large significance, ++++: very large significance </small>
				
				`,
			}.Outline("2.")
		}
	}
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3
		{
			inp := gr.AddInput()
			inp.Type = "dyn-composite"
			inp.DynamicFunc = fmt.Sprintf("Special202403QS2__%v__%v", 0, 0)
			inp.DynamicFuncParamset = ""
			inp.ColSpanControl = 1
		}

		_, inputNames, _ := cpfmt.Special202403QS2(q, 0, 0, true)
		for _, inpName := range inputNames {
			inp := gr.AddInput()
			inp.Type = "dyn-composite-scalar"
			inp.Name = inpName
		}

	}

	return nil
}
