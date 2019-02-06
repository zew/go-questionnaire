package strube

import (
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
	q.Variations = 0 // attention => shuffles submit buttons if > 0

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
	i2 := "[attr-has-euro]"
	_, _ = i1, i2

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
			// Dynamic header for following headless radio matrix
			gr := p.AddGroup()
			gr.BottomVSpacers = 0
			inp := gr.AddInput()
			inp.Type = "dynamic"
			inp.DynamicFunc = "HasEuroQuestion"
		}
		{
			// Headless radio matrix
			names1stMatrix := []string{"benefit"}
			emptyRowLabels := []trl.S{}
			gr2 := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr2.Cols = 9 // necessary, otherwise no vspacers
			gr2.OddRowsColoring = true

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

	}

	//
	// Page 2
	{

		p := q.AddPage()
		// p.NoNavigation = true
		p.Width = 70
		p.Section = trl.S{
			"de": "Zur Aufgaben- und Kompetenzverteilung in Europa",
			"en": "EU competencies",
			"fr": "Répartition des missions et des compétences en Europe",
			"it": "Competenze dell'Unione europea (UE). È d'accordo con le seguenti affermazioni?",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Desc = trl.S{
			"de": "",
			"en": "",
			"fr": "",
			"it": "",
		}
		p.Short = trl.S{
			"de": "Aufgaben- und Kompetenzverteilung",
			"en": "EU competencies",
			"fr": "Répartition des compétences",
			"it": "Competenze UE",
		}

		// 21
		{
			names1stMatrix := []string{"tax"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Steuerpolitik<br>",
				"en": "Tax policy<br>",
				"fr": "Politique fiscale<br>",
				"it": "Tassazione<br>",
			}
			gr.Desc = trl.S{
				"de": "Der Rat der EU sollte mit qualifizierter Mehrheit anstelle von Einstimmigkeit über Steuern beschließen können (z.B. über verbindliche Ober- oder Untergrenzen für Unternehmenssteuern).",
				"en": "The European Council should be able to vote on tax issues with a qualified majority instead of una-nimity (e.g. common caps or floors for corporate taxes binding for member states).",
				"fr": "Le Conseil européen devrait pouvoir statuer avec une majorité qualifiée sur les impôts perçus dans les États membres (par exemple sur des taux planchers et plafonds de l’impôt sur les Sociétés).",
				"it": "Il Consiglio europeo dovrebbe poter votare su questioni tributarie a maggioranza qualificata invece che all’unanimità (ad esempio su limiti massimi e minimi per le imposte sulle aziende comuni a tutti gli Stati membri).",
			}
		}

		// 22
		{
			names1stMatrix := []string{"redist"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad19(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": "Umverteilung<br>",
				"en": "Redistribution<br>",
				"fr": "Redistribution des revenus<br>",
				"it": "Ridistribuzione<br>",
			}
			gr.Desc = trl.S{
				"de": "Es sollte mehr Umverteilung von reichen zu armen EU-Mitgliedstaaten geben.",
				"en": "There should be more redistribution from richer to poorer EU member states.",
				"fr": "Il devrait y avoir davantage de redistribution des États membres de l'UE les plus riches vers les plus pauvres.",
				"it": "Ci dovrebbe essere maggiore ridistribuzione di risorse dagli Stati membri più ricchi a quelli più po-veri dell'UE.",
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
				inp.Response = "2"
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
