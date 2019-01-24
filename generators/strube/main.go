package strube

import (
	"fmt"

	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var xx = trl.S{
	"de": "...",
	"en": "...",
	"es": "...",
	"fr": "...",
	"it": "...",
	"pl": "...",
}

func labelsGoodBad19() []trl.S {

	tm := []trl.S{
		{
			"de": "lehne ab<span class='ordinals'><br>-4</span>",
			"en": "Disagree<span class='ordinals'><br>-4</span>",
			"es": "en desacuerdo<span class='ordinals'><br>-4</span>",
			"fr": "Pas d’accord<span class='ordinals'><br>-4</span>",
			"it": "Non favorevole<span class='ordinals'><br>-4</span>",
			"pl": "jestem przeciwny/a<span class='ordinals'><br>-4</span>",
		},
		{
			"de": "<span class='ordinals'><br>-3</span>",
			"en": "<span class='ordinals'><br>-3</span>",
			"es": "<span class='ordinals'><br>-3</span>",
			"fr": "<span class='ordinals'><br>-3</span>",
			"it": "<span class='ordinals'><br>-3</span>",
			"pl": "<span class='ordinals'><br>-3</span>",
		},
		{
			"de": "<span class='ordinals'><br>-2</span>",
			"en": "<span class='ordinals'><br>-2</span>",
			"es": "<span class='ordinals'><br>-2</span>",
			"fr": "<span class='ordinals'><br>-2</span>",
			"it": "<span class='ordinals'><br>-2</span>",
			"pl": "<span class='ordinals'><br>-2</span>",
		},
		{
			"de": "<span class='ordinals'><br>-1</span>",
			"en": "<span class='ordinals'><br>-1</span>",
			"es": "<span class='ordinals'><br>-1</span>",
			"fr": "<span class='ordinals'><br>-1</span>",
			"it": "<span class='ordinals'><br>-1</span>",
			"pl": "<span class='ordinals'><br>-1</span>",
		},
		{
			"de": "unentschieden<span class='ordinals'><br>0</span>",
			"en": "Undecided<span class='ordinals'><br>0</span>",
			"es": "indeciso<span class='ordinals'><br>0</span>",
			"fr": "Indifférent<span class='ordinals'><br>0</span>",
			"it": "Indeciso<span class='ordinals'><br>0</span>",
			"pl": "jestem niezdecydowany/a<span class='ordinals'><br>0</span>",
		},
		{
			"de": "<span class='ordinals'><br>1</span>",
			"en": "<span class='ordinals'><br>1</span>",
			"es": "<span class='ordinals'><br>1</span>",
			"fr": "<span class='ordinals'><br>1</span>",
			"it": "<span class='ordinals'><br>1</span>",
			"pl": "<span class='ordinals'><br>1</span>",
		},
		{
			"de": "<span class='ordinals'><br>2</span>",
			"en": "<span class='ordinals'><br>2</span>",
			"es": "<span class='ordinals'><br>2</span>",
			"fr": "<span class='ordinals'><br>2</span>",
			"it": "<span class='ordinals'><br>2</span>",
			"pl": "<span class='ordinals'><br>2</span>",
		},
		{
			"de": "<span class='ordinals'><br>3</span>",
			"en": "<span class='ordinals'><br>3</span>",
			"es": "<span class='ordinals'><br>3</span>",
			"fr": "<span class='ordinals'><br>3</span>",
			"it": "<span class='ordinals'><br>3</span>",
			"pl": "<span class='ordinals'><br>3</span>",
		},
		{
			"de": "stimme zu<span class='ordinals'><br>4</span>",
			"en": "Agree<span class='ordinals'><br>4</span>",
			"es": "de acuerdo<span class='ordinals'><br>4</span>",
			"fr": "D’accord<span class='ordinals'><br>4</span>",
			"it": "Favorevole<span class='ordinals'><br>4</span>",
			"pl": "zgadzam się<span class='ordinals'><br>4</span>",
		},
	}

	return tm

}

// Create creates a questionnaire with a few pages and inputs.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {
	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("strube")
	q.Survey.Params = params
	q.Variations = 4

	q.LangCodes = map[string]string{
		"de": "Deutsch",
		"en": "English",
		"es": "Español",
		"fr": "Français",
		"it": "Italiano",
		"pl": "Polski",
	}
	q.LangCodesOrder = []string{
		"en",
		"fr",
		"de",
		"it",
		"es",
		"pl",
	}
	q.LangCode = "en" // No default; forces usage of UserLangCode()

	q.Survey.Org = trl.S{
		"de": "  ",
		"en": "  ",
		"es": "  ",
		"fr": "  ",
		"it": "  ",
		"pl": "  ",
	}
	q.Survey.Name = trl.S{
		"de": "",
		"en": "General attitudes on euro and economic policy",
		"es": "",
		"fr": "",
		"it": "",
		"pl": "",
	}

	i2 := "[groupID]"

	//
	// Page 1
	{

		p := q.AddPage()
		p.NoNavigation = true
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"en": "Do you agree with the following statements?",
		}
		p.Label = trl.S{
			"en": "",
		}
		p.Desc = trl.S{
			"en": "",
		}
		p.Short = trl.S{
			"en": "Survey",
		}

		// 11
		{
			names1stMatrix := []string{"ecb"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Appropriate monetary policy stance – European Central Bank<br>", i2),
			}
			gr.Desc = trl.S{
				"en": "The European Central Bank should take its time to normalize its interest rates. There is no need to increase central bank rates in Europe before the year 2020.",
			}
		}

		{
			gr := p.AddGroup()
			gr.Cols = 1
			gr.Width = 99
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = "1"
				inp.Label = trl.S{
					"de": "Weiter",
					"en": "Submit",
				}
				inp.AccessKey = "n"
				inp.ColSpanControl = 1
				inp.HAlignControl = qst.HRight
			}
		}

	}

	//
	// Page Finish
	{
		p := q.AddPage()
		p.NoNavigation = true
		p.Label = trl.S{
			"de": "Vielen Dank",
			"en": "Thank you",
		}

		{
			// Only one group => shuffling is no problem
			gr := p.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Desc = trl.S{
					"de": "Danke für Ihre Teilnahme an unserer Umfrage.",
					"en": "Thank you for your participation in our survey.",
					"es": "Gracias por haber contestado a nuestro cuestionario.",
					"fr": "Nous vous remercions d'avoir répondu à nos questions.",
					"it": "Grazie per aver risposto al nostro questionario.",
					"pl": "Dziękujemy za uczestnictwo w ankiecie.",
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Desc = trl.S{
					"de": "<span style='font-size: 100%;'>Ihre Eingaben wurden gespeichert.</span>",
					"en": "<span style='font-size: 100%;'>Your entries have been saved.</span>",
					"es": "<span style='font-size: 100%;'>Sus entradas se han guardado.</span>",
					"fr": "<span style='font-size: 100%;'>Vos réponses ont été sauvegardées.</span>",
					"it": "<span style='font-size: 100%;'>Le Sue risposte sono state salvate.</span>",
					"pl": "<span style='font-size: 100%;'>Twoje wpisy zostały zapisane.</span>",
				}
			}

		}

	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	return &q, nil
}
