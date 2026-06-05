package lix

import (
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create questionnaire
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"de"} // governs default language code
	// q.LangCode = "en"

	q.Survey.Org = trl.S{
		"en": "ZEW",
		"de": "ZEW",
	}
	q.Survey.Name = trl.S{
		"en": "Länderindex Umfrage",
		"de": "Länderindex Umfrage",
	}
	// q.Variations = 1

	// page0
	{
		page := q.AddPage()
		page.ValidationFuncName = ""

		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.NoNavigation = true

		page.Label = trl.S{
			"en": "",
			"de": "",
		}
		// page.Label = trl.S{
		// 	"en": "Dear Madam / Sir,",
		// 	"de": "Sehr geehrter Damen und Herren",
		// }
		// page.Short = trl.S{
		// 	"en": "Greeting",
		// 	"de": "Begrüßung",
		// }

		page.WidthMax("42rem")

		inps := []string{
			"hk_steuern",
			"hk_arbeit",
			"hk_fin",
			"hk_reg",
			"hk_inf",
			"hk_ene",
			"uk_steuern_biz",
			"uk_steuern_erb",
			"uk_steuern_kpl",
			"uk_arbeit_kos",
			"uk_arbeit_pro",
			"uk_fin_krd",
			"uk_fin_vrs",
			"uk_reg_ins",
			"uk_reg_vor",
			"uk_reg_mit",
			"uk_inf_tra",
			"uk_inf_rec",
			"uk_ene_pre",
			"uk_ene_sic",
			"uk_ene_kli",
		}

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3

			for _, inpName := range inps {
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = inpName
			}

			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./main.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	//
	//
	// Report of results
	{
		page := q.AddPage()
		page.NoNavigation = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		page.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": `&nbsp;`,
					"de": `&nbsp;`,
				}
			}
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.DynamicFunc = "RepsonseStatistics"
			// }
		}
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	q.SetColspans()

	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
