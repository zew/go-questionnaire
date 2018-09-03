package peu2018

import (
	"fmt"

	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/trl"
)

var xx = trl.S{
	"de": "...",
	"en": "...",
	"es": "...",
	"fr": "...",
	"it": "...",
	"pl": "...",
}

func labelsGoodBad17() []trl.S {

	tm := []trl.S{
		{
			"de": "lehne ab",
			"en": "Disagree",
			"es": "en desacuerdo",
			"fr": "Pas d’accord",
			"it": "Non favorevole",
			"pl": "jestem przeciwny/a",
		},
		{
			"de": "...",
			"en": "...",
			"es": "...",
			"fr": "...",
			"it": "...",
			"pl": "...",
		},
		{
			"de": "...",
			"en": "...",
			"es": "...",
			"fr": "...",
			"it": "...",
			"pl": "...",
		},
		{
			"de": "unentschieden",
			"en": "Undecided",
			"es": "indeciso",
			"fr": "Indifférent",
			"it": "Indeciso",
			"pl": "jestem niezdecydowany/a",
		},
		{
			"de": "...",
			"en": "...",
			"es": "...",
			"fr": "...",
			"it": "...",
			"pl": "...",
		},
		{
			"de": "...",
			"en": "...",
			"es": "...",
			"fr": "...",
			"it": "...",
			"pl": "...",
		},
		{
			"de": "stimme zu",
			"en": "Agree",
			"es": "de acuerdo",
			"fr": "D’accord",
			"it": "Favorevole",
			"pl": "zgadzam się",
		},
	}

	return tm

}

// Create creates an minimal example questionaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionaireT, error) {
	q := qst.QuestionaireT{}
	q.Survey = qst.NewSurvey("eup")
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
	q.LangCode = "de"

	q.Survey.Org = trl.S{
		"de": "ZEW",
		"en": "ZEW",
		"es": "ZEW",
		"fr": "ZEW",
		"it": "ZEW",
		"pl": "ZEW",
	}
	q.Survey.Name = trl.S{
		"de": "Umfrage: Zur Zukunft der Europäischen Union (EU)",
		"en": "Survey: On the prospects of the European Union (EU)",
		"es": "Cuestionario: El futuro de la Unión Europea (UE)",
		"fr": "De l’avenir de l’Union Européenne",
		"it": "Questionario: le prospettive dell’Unione Europea (UE)",
		"pl": "Badanie na temat: Przyszłość Unii Europejskiej (UE)",
	}

	i2 := "[groupID]"
	{

		p := q.AddPage()
		p.Width = 70
		// p.Label = trl.S{
		p.Section = trl.S{
			"de": "1. Zur Aufgaben- und Kompetenzverteilung in Europa",
			"en": "1. Competency allocation in Europe",
			"es": "1. Sobre la distribución de las tareas y competencias en Europa",
			"fr": "1. Répartition des missions et des compétences en Europe",
			"it": "1. La distribuzione delle competenze in Europa tra i Paesi membri e la UE",
			"pl": "1. Podział zadań i kompetencji w Europie",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
		}
		p.Desc = trl.S{
			"de": "Inwieweit stimmen Sie den folgenden Aufgaben zu?",
			"en": "Do you approve the following proposals?",
			"es": "¿En qué medida está de acuerdo con las siguientes competencias?",
			"fr": "Actuellement l’élargissement des compétences de l’UE dans certains domaines politiques est en débat. Approuvez-vous les propositions suivantes ?",
			"it": "cosa pensa delle seguenti proposte?",
			"pl": "Wskaż, w jakim stopniu zgadzasz się z poniższymi stwierdzeniami?",
		}
		p.Short = trl.S{
			"de": "Kompetenzverteilung",
			"en": "Competency allocation",
			"es": "distribución de las tareas y competencias",
			"fr": "Répartition des missions et des compétences",
			"it": "La distribuzione delle competenze",
			"pl": "Podział zadań i kompetencji",
		}

		// 11
		{
			names1stMatrix := []string{"energy"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Energiepolitik<br>", i2),
				"en": fmt.Sprintf("%v. Energy policy<br>", i2),
				"es": fmt.Sprintf("%v. Política energética<br>", i2),
				"fr": fmt.Sprintf("%v. Politique énergétique<br>", i2),
				"it": fmt.Sprintf("%v. Politica energetica<br>", i2),
				"pl": fmt.Sprintf("%v. Polityka energetyczna<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte den Mitgliedstaaten verbindliche Vorgaben zum Energiemix machen dürfen (z.B. über den Anteil von erneuerbaren Energien, Kohle oder Kernenergie).",
				"en": "The EU should be able to make binding guidelines to the member states regarding the energy mix (e.g. regarding the share of renewable energies, coal or nuclear power).",
				"es": "La UE debería poder establecer normas vinculantes sobre el mix energético (por ejemplo, sobre la cuota de energías renovables, energía del carbón o nuclear).",
				"fr": "L’UE pourrait donner des directives contraignantes aux États membres quant à leurs choix de mix énergétique dans leur production d’électricité (par exemple en fixant la proportion des énergies renouvelables, du charbon, et de l’énergie nucléaire).",
				"it": "L’UE dovrebbe essere in grado di fornire linee guida vincolanti agli Stati membri sul mix energetico (ad esempio riguardo la quota di energie rinnovabili, carbone o energia nucleare).",
				"pl": "UE powinna mieć możliwość wyznaczania wiążących zasad wobec państw członkowskich w zakresie koszyka energetycznego (na przykład udziału energii ze źródeł odnawialnych, węgla lub energii jądrowej).",
			}
		}

		// 12
		{
			names1stMatrix := []string{"immigration"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Einwanderungspolitik<br>", i2),
				"en": fmt.Sprintf("%v. Immigration policy<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte eine stärkere Rolle in der Einwanderungspolitik erhalten (z.B. Aufnahmestandards festlegen oder über die Verteilung von Flüchtlingen entscheiden).",
				"en": "The EU should get a stronger role in immigration policy (e.g. decisions over admission standards or allocation of refugees).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 13
		{
			names1stMatrix := []string{"defence"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Verteidigungspolitik<br>", i2),
				"en": fmt.Sprintf("%v. Defence policy<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Eine unter dem Befehl der EU stehende und aus ihrem Haushalt finanzierte europäische Armee sollte Aufgaben der nationalen Streitkräfte für internationale Kriseneinsätze übernehmen.",
				"en": "A European army under the command of the EU and financed from its budget should take over duties from national armies regarding international conflict deployments.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 14
		{
			names1stMatrix := []string{"wages"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Lohnpolitik<br>", i2),
				"en": fmt.Sprintf("%v. Wage policy<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte in der Lohnpolitik stärkere Eingriffsrechte erhalten (z.B. bezüglich der Höhe für allgemeine gesetzliche Mindestlöhne).",
				"en": "The EU should have stronger rights to intervene in the wage policies (e.g. regarding the level of general statutory minimum wages).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 15
		{
			names1stMatrix := []string{"flexibility"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Regeln des Arbeitsmarktes<br>", i2),
				"en": fmt.Sprintf("%v. Labour market regulation<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte den Mitgliedstaaten verbindliche Vorgaben zum Arbeitsmarkt machen dürfen (z.B. zur Ausgestaltung des Kündigungsschutzes oder von Zeitverträgen).",
				"en": "The EU should be able to make binding guidelines to member states regarding the labour market (e.g. regarding the design of dismissal protection or temporary contracts).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 16
		{
			names1stMatrix := []string{"eutax"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. EU-Steuer als neue Eigenmittelart<br>", i2),
				"en": fmt.Sprintf("%v. EU tax as a new own resource<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Es sollte eine neue steuerbasierte Eigenmittelart für den EU-Haushalt unter direkter Kontrolle der EU geben (z.B. eine EU-Steuer auf eine gemeinsame Körperschaftsteuer-Bemessungsgrundlage).",
				"en": "There should be a new tax-based own resource for the EU budget under direct control of the EU (e.g. an EU tax on a common corporate tax base).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 17
		{
			names1stMatrix := []string{"taxpolicy"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Steuerpolitik<br>", i2),
				"en": fmt.Sprintf("%v. Tax policy<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Der Rat der EU sollte mit qualifizierter Mehrheit anstelle von Einstimmigkeit über Steuern beschließen können (z.B. über verbindliche Ober- oder Untergrenzen für Unternehmenssteuern).",
				"en": "The European Council should be able to vote on tax issues with a qualified majority instead of unanimity (e.g. common caps or floors for corporate taxes binding for member states).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 18
		{
			names1stMatrix := []string{"initiative"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Initiativrechte für das Europäische Parlament<br>", i2),
				"en": fmt.Sprintf("%v. European Parliament and legislative initiative<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Das Europäische Parlament sollte das Recht erhalten, neue EU-Gesetze vorzuschlagen. Diese Gesetzesinitiative ist bisher der Europäischen Kommission vorbehalten.",
				"en": "The European Parliament should get the right to propose new EU laws (i.e. the legislative initiative) which is currently confined to the European Commission.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

	}
	{
		p := q.AddPage()
		p.Width = 70
		p.Section = trl.S{
			"de": "2. Reforminitiativen in der Europäischen Währungsunion (EWU)",
			"en": "2. Reform initiatives in the European Monetary Union (EMU)",
			"es": "2. ",
			"fr": "2. ",
			"it": "2. ",
			"pl": "2. ",
		}
		p.Label = trl.S{
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
		}
		p.Desc = trl.S{
			"de": "Inwieweit stimmen Sie den folgenden Aussagen zu?",
			"en": "Do you approve the following proposals?",
			"es": "xxx",
			"fr": "xxx",
			"it": "xxx",
			"pl": "xxx",
		}
		p.Short = trl.S{
			"de": "Reforminitiativen",
			"en": "Reform initiatives",
			"es": "xxx",
			"fr": "xxx",
			"it": "xxx",
			"pl": "xxx",
		}

		// 21
		{
			names1stMatrix := []string{"investment"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Höhere Investitionen <br>", i2),
				"en": fmt.Sprintf("%v. Higher investment<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein höheres Wachstum der Eurozone ist es unverzichtbar, dass die Staaten der Eurozone ihre Ausgaben für Investitionen erhöhen.",
				"en": "For higher economic growth of the EMU it is essential that its member states increase their investment expenditures.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 22
		{
			names1stMatrix := []string{"more_labor_flex"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Flexiblere Arbeitsmärkte<br>", i2),
				"en": fmt.Sprintf("%v. Flexible labour markets<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein höheres Wachstum der Eurozone ist es unverzichtbar, dass insbesondere die Staaten mit anhaltend hoher Arbeitslosigkeit ihre Arbeitsmärkte flexibler machen (z.B. durch eine Lockerung des Kündigungsschutzes oder eine Absenkung von gesetzlichen Mindestlöhnen).",
				"en": "For higher economic growth of the EMU it is essential that especially countries with permanently high levels of unemployment make their labour markets more flexible (e.g. via an easing of dismissal protection regulations or a decrease of the statutory minimum wage).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 23
		{
			names1stMatrix := []string{"unemployment_insurance"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Europäische Arbeitslosenversicherung<br>", i2),
				"en": fmt.Sprintf("%v. European unemployment insurance<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Eine gemeinsame europäische Arbeitslosenversicherung sollte eingeführt werden, um Rezessionen in einzelnen Mitgliedsstaaten der Eurozone abzufedern.",
				"en": "A common European unemployment insurance should be introduced to absorb recessions in individual member states of the EMU.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 24
		{
			names1stMatrix := []string{"eurobonds"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Eurobonds<br>", i2),
				"en": fmt.Sprintf("%v. Eurobonds<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für Eurobonds haften alle Euro-Staaten gemeinsam und alle Euro-Staaten zahlen den gleichen Zins. Die EWU sollte Eurobonds ausgeben.",
				"en": "All euro countries are jointly liable for Eurobonds and all euro countries pay the same interest. The EMU should issue Eurobonds.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 25
		{
			names1stMatrix := []string{"asset_purchase"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Wertpapierkaufprogramm der EZB<br>", i2),
				"en": fmt.Sprintf("%v. Asset purchase program of ECB<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die Europäische Zentralbank (EZB) hat in den zurückliegenden Jahren durch den Kauf von Staatsanleihen von Euro-Staaten eine sehr aktive Rolle gespielt. Diese starke Rolle der EZB sollte fortgesetzt werden.",
				"en": "The European Central Bank (ECB) did take a strongly active position in recent years by purchasing sovereign bonds of euro countries. This strongly active position of the ECB should continue.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 26
		{
			names1stMatrix := []string{"growth_pact"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Stabilitäts- und Wachstumspakt (SWP)<br>", i2),
				"en": fmt.Sprintf("%v. Stability and Growth Pact (SGP)<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Der SWP definiert Defizit- und Schuldengrenzen für EU-Mitgliedsstaaten. Der SWP schränkt die Fiskalpolitik der Mitgliedsstaaten unangemessen stark ein und sollte gelockert werden.",
				"en": "The SGP defines deficit and debt limits for EU member states. The SGP inappropriately constrains fiscal policy in member states, and should be relaxed.",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 27
		{
			names1stMatrix := []string{"emu_institutions"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Neue EWU-Institutionen<br>", i2),
				"en": fmt.Sprintf("%v. New EMU institutions<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein angemessenes Funktionieren benötigt die EWU neue Fiskalinstitutionen (z.B. ein Eurozonenbudget oder einen europäischen Finanzminister).",
				"en": "For a proper functioning, the EMU needs new fiscal institutions (e.g. a euro area budget or a European Minister of Finance).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

		// 28
		{
			names1stMatrix := []string{"banking_union"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 7 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Vollendung der Bankenunion<br>", i2),
				"en": fmt.Sprintf("%v. Completion of Banking Union<br>", i2),
				"es": fmt.Sprintf("%v. xx<br>", i2),
				"fr": fmt.Sprintf("%v. xx<br>", i2),
				"it": fmt.Sprintf("%v. xx<br>", i2),
				"pl": fmt.Sprintf("%v. xx<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein angemessenes Funktionieren sollte die europäische Bankenunion durch die Europäische Einlagensicherung (European Deposit Insurance System: EDIS) vollendet werden.",
				"en": "For its proper functioning, the European Banking Union should be completed through the European Deposit Insurance Scheme (EDIS).",
				"es": "es Desc",
				"fr": "fr Desc",
				"it": "it Desc",
				"pl": "pl Desc",
			}
		}

	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	return &q, nil
}
