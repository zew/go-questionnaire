package strube

import (
	"fmt"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var xx = trl.S{
	"de": "...",
	"en": "...",
	"fr": "...",
	"it": "...",
}

func labelsGoodBad19() []trl.S {

	tm := []trl.S{
		{
			"de": "lehne ab<span class='ordinals'><br>-4</span>",
			"en": "Disagree<span class='ordinals'><br>-4</span>",
			"fr": "Pas d’accord<span class='ordinals'><br>-4</span>",
			"it": "Non favorevole<span class='ordinals'><br>-4</span>",
		},
		{
			"de": "<span class='ordinals'><br>-3</span>",
			"en": "<span class='ordinals'><br>-3</span>",
			"fr": "<span class='ordinals'><br>-3</span>",
			"it": "<span class='ordinals'><br>-3</span>",
		},
		{
			"de": "<span class='ordinals'><br>-2</span>",
			"en": "<span class='ordinals'><br>-2</span>",
			"fr": "<span class='ordinals'><br>-2</span>",
			"it": "<span class='ordinals'><br>-2</span>",
		},
		{
			"de": "<span class='ordinals'><br>-1</span>",
			"en": "<span class='ordinals'><br>-1</span>",
			"fr": "<span class='ordinals'><br>-1</span>",
			"it": "<span class='ordinals'><br>-1</span>",
		},
		{
			"de": "unentschieden<span class='ordinals'><br>0</span>",
			"en": "Undecided<span class='ordinals'><br>0</span>",
			"fr": "Indifférent<span class='ordinals'><br>0</span>",
			"it": "Indeciso<span class='ordinals'><br>0</span>",
		},
		{
			"de": "<span class='ordinals'><br>1</span>",
			"en": "<span class='ordinals'><br>1</span>",
			"fr": "<span class='ordinals'><br>1</span>",
			"it": "<span class='ordinals'><br>1</span>",
		},
		{
			"de": "<span class='ordinals'><br>2</span>",
			"en": "<span class='ordinals'><br>2</span>",
			"fr": "<span class='ordinals'><br>2</span>",
			"it": "<span class='ordinals'><br>2</span>",
		},
		{
			"de": "<span class='ordinals'><br>3</span>",
			"en": "<span class='ordinals'><br>3</span>",
			"fr": "<span class='ordinals'><br>3</span>",
			"it": "<span class='ordinals'><br>3</span>",
		},
		{
			"de": "stimme zu<span class='ordinals'><br>4</span>",
			"en": "Agree<span class='ordinals'><br>4</span>",
			"fr": "D’accord<span class='ordinals'><br>4</span>",
			"it": "Favorevole<span class='ordinals'><br>4</span>",
		},
	}

	return tm

}

// Create creates a questionnaire with a few pages and inputs.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {
	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("strube")
	q.Survey.Params = params
	// q.Variations = 4

	q.LangCodes = map[string]string{
		"de": "Deutsch",
		"en": "English",
		"fr": "Français",
		"it": "Italiano",
	}
	q.LangCodesOrder = []string{
		"en",
		"fr",
		"de",
		"it",
	}
	q.LangCode = "en" // No default; forces usage of UserLangCode()

	q.Survey.Org = trl.S{
		"de": "ZEW",
		"en": "ZEW",
		"fr": "ZEW",
		"it": "ZEW",
	}
	q.Survey.Name = trl.S{
		"de": "Umfrage: Gestaltung der Europäischen Union",
		"en": "Survey: Design of the European Union",
		"fr": "Questionnaire: Design de l’Union Européenne",
		"it": "Questionario: Design dell'Unione Europea",
	}

	i1 := "[attr-country]"
	i2 := "[attr2]"
	_ = i2

	//
	// Page 1
	{

		p := q.AddPage()
		// p.NoNavigation = true
		p.Width = 70
		p.Section = trl.S{
			"de": "Allgemeine Einstellungen zum Euro und zur Wirtschaftspolitik",
			"en": "General attitudes on euro and economic policy",
			"fr": "Attitudes générales vis-à-vis de l'euro et de la politique économique",
			"it": "Posizioni generali sull'Euro e sulla politica economica europea.",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Desc = trl.S{
			"de": "Inwieweit stimmen Sie den folgenden Aufgaben zu?",
			"en": "Do you agree with the following statements?",
			"fr": "Approuvez-vous les propositions suivantes ?",
			"it": "In che misura si trova d’accordo con le seguenti affermazioni?",
		}
		p.Short = trl.S{
			"de": "Allgemeine Einstellungen",
			"en": "General attitudes",
			"fr": "Attitudes générales",
			"it": "Posizioni generali",
		}

		// 11
		{
			names1stMatrix := []string{"benefit"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Wirtschaftlicher Nutzen des Euro<br>",
				"en": "Economic benefits of euro<br>",
				"fr": "Avantages économiques de l'euro<br>",
				"it": "Benefici economici dell'Euro<br>",
			}
			gr.Desc = trl.S{
				"de": fmt.Sprintf("Den Euro in %v als die offizielle Währung zu haben, ist wirtschaftlich vorteilhaft.", i1),
				"en": fmt.Sprintf("Having the euro in %v as the official currency is economically beneficial.", i1),
				"fr": fmt.Sprintf("Avoir l'euro dans %v comme monnaie officielle est économiquement avantageux.", i1),
				"it": fmt.Sprintf("Avere l'Euro come valuta ufficiale nel %v è economicamente vantaggioso.", i1),
			}
		}

		// 12
		{
			names1stMatrix := []string{"supply"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Nachfrage- vs. angebotsbezogene Politikmaßnahmen<br>",
				"en": "Demand-side versus supply-side policies<br>",
				"fr": "Politiques axées sur la demande et politiques axées sur l’offre<br>",
				"it": "Politiche della domanda e dell'offerta<br>",
			}
			gr.Desc = trl.S{
				"de": "Regierungen können versuchen, das Wirtschaftswachstum durch verschiedene Instrumente zu stimulieren. Einige argumentieren, dass nachfragebezogene Politikmaßnahmen (z. B. eine Erhöhung der durch Schulden finanzierten öffentlichen Ausgaben) wirksamer sind als angebotsbezogene Maßnahmen (z. B. eine Verringerung der Regulierung der Arbeits- und Gütermärkte).",
				"en": "Governments can try to stimulate economic growth through different instruments. Some argue that demand-side policies (e.g. an increase of debt-financed public spending) are more effective than supply-side policies (e.g. a reduction of regulation in labour and good markets).",
				"fr": "Les gouvernements peuvent essayer de stimuler la croissance économique grâce à différents instruments. Certains font valoir que les politiques axées sur la demande (par exemple, une augmentation des dépenses publiques financées par l’endettement) sont plus efficaces que les politiques axées sur l'offre (par exemple, une réduction de la réglementation sur le marché du travail et des biens).",
				"it": "I governi possono cercare di stimolare la crescita economica attraverso diversi strumenti. Alcuni sostengono che le politiche di sostegno della domanda (ad esempio un aumento della spesa pubblica finanziato in deficit) siano più efficaci delle politiche di sostegno dell’offerta (ad esempio un alleggerimento della regolamentazione del mercato del lavoro e dei beni).",
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
				inp.Label = cfg.Get().Mp["next"]
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
			"fr": "Nous vous remercions d'avoir répondu à nos questions.",
			"it": "Grazie per aver risposto al nostro questionario.",
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
					"fr": "Nous vous remercions d'avoir répondu à nos questions.",
					"it": "Grazie per aver risposto al nostro questionario.",
				}
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				inp.Desc = trl.S{
					"de": "<span style='font-size: 100%;'>Ihre Eingaben wurden gespeichert.</span>",
					"en": "<span style='font-size: 100%;'>Your entries have been saved.</span>",
					"fr": "<span style='font-size: 100%;'>Vos réponses ont été sauvegardées.</span>",
					"it": "<span style='font-size: 100%;'>Le Sue risposte sono state salvate.</span>",
				}
			}

		}

	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	return &q, nil
}
