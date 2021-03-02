package mul

import (
	"fmt"

	"github.com/zew/go-questionnaire/ctr"
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

	ctr.Reset()

	// qst.RadioVali = "mustRadioGroup"
	qst.RadioVali = ""
	qst.HeaderClass = ""
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("aik")
	q.Survey.Params = params
	q.Variations = 0

	q.LangCodes = []string{
		"en",
	} // governs default language code

	q.Survey.Org = trl.S{
		"en": "  ",
	}
	q.Survey.Name = trl.S{
		"en": "Macro Policy Survey",
	}

	groupOrdinal := "[groupID]"

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
				"en": fmt.Sprintf("%v. Appropriate monetary policy stance – European Central Bank<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"en": "The European Central Bank should take its time to normalize its interest rates. There is no need to increase central bank rates in Europe before the year 2020.",
			}
		}

		// 12
		{
			names1stMatrix := []string{"fed"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Appropriate monetary policy stance – US Federal Reserve<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				// "en": "The EU should get a stronger role in immigration policy (e.g. decisions over admission standards or allocation of refugees).",
				"en": "The US Federal Reserve should continue with its course of interest rate increases well throughout 2019.",
			}
		}

		// 13
		{
			names1stMatrix := []string{"fiscal"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Appropriate fiscal policy stance<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"en": "Overall, the current economic situation in industrial countries allows more fiscal consolidation. OECD countries should consolidate more and try to reduce government debt. ",
			}
		}

		// 14
		{
			names1stMatrix := []string{"privatization"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Need for privatization<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"en": "In general, the large involvement of governments in market activities still impairs the growth potential of industrial countries. Privatization should be one of the priorities in strategies to boost the growth potential in OECD countries.",
			}
		}

		// 15
		{
			names1stMatrix := []string{"labor"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Extent of labor market regulation<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"en": "High long-run unemployment in some OECD countries largely reflects an excessive level of labor market regulation. In order to reduce this unemployment, countries would have to deregulate labor markets.",
			}
		}

		// 16
		{
			names1stMatrix := []string{"redistribution"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Redistribution<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"en": "Current inequalities in OECD countries are not just a fairness issue but also detrimental for the growth potential. Governments should address these inequalities by more intense redistribution.",
			}
		}

		// 17
		{
			names1stMatrix := []string{"rules"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"en": fmt.Sprintf("%v. Fiscal Rules<br>", groupOrdinal),
			}
			gr.Desc = trl.S{
				"en": "Fiscal rules like the European Stability and Growth Pact constrain government on the size of the government deficit. Rules like this may not be perfectly effective but, in principle, they are helpful and support long-run economic stability.",
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
				inp.CSSLabel = "special-line-height-higher"
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
				inp.CSSLabel = "special-line-height-higher"
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
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
