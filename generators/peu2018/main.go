package peu2018

import (
	"fmt"
	"log"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/tpl"
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
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
		},
		{
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
		},
		{
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
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
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
		},
		{
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
		},
		{
			"de": "",
			"en": "",
			"es": "",
			"fr": "",
			"it": "",
			"pl": "",
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
	q.Survey = qst.NewSurvey("peu2018")
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
	q.LangCode = "" // No default; forces usage of UserLangCode()

	q.Survey.Org = trl.S{
		"de": "  ",
		"en": "  ",
		"es": "  ",
		"fr": "  ",
		"it": "  ",
		"pl": "  ",
	}
	q.Survey.Name = trl.S{
		"de": "Umfrage: Zur Zukunft der Europäischen Union (EU)",
		"en": "Survey: On the prospects of the European Union (EU)",
		"es": "Cuestionario: El futuro de la Unión Europea (UE)",
		"fr": "Questionnaire: De l’avenir de l’Union Européenne",
		"it": "Questionario: le prospettive dell’Unione Europea (UE)",
		"pl": "Badanie na temat: Przyszłość Unii Europejskiej (UE)",
	}

	i2 := "[groupID]"

	//
	// Page Welcome
	{
		p := q.AddPage()
		p.Label = trl.S{
			"de": "Willkommen",
			"en": "Welcome",
			"es": "Bienvenido",
			"fr": "Bienvenue",
			"it": "Benvenuto",
			"pl": "Zapraszamy",
		}
		p.NoNavigation = true

		{
			// Only one group => shuffling is no problem
			gr := p.AddGroup()
			gr.Cols = 1
			{

				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				inp.ColSpanLabel = 3
				impr := trl.S{}
				for lc := range q.LangCodes {
					cnt, err := tpl.MarkDownFromFile("./static/doc/welcome.md", q.Survey.Type, lc)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = cnt
				}
				inp.Desc = impr

				{
					inp := gr.AddInput()
					inp.Type = "button"
					inp.Name = "submitBtn"
					inp.Response = "1"
					inp.Label = cfg.Get().Mp["start"]
					inp.AccessKey = "n"
					inp.ColSpanControl = 1
					inp.HAlignControl = qst.HRight
				}

			}

		}
	}

	//
	// Page 1
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
			gr.Cols = 9 // necessary, otherwise no vspacers
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
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Einwanderungspolitik<br>", i2),
				"en": fmt.Sprintf("%v. Immigration policy<br>", i2),
				"es": fmt.Sprintf("%v. Política migratoria<br>", i2),
				"fr": fmt.Sprintf("%v. Politique d’immigration<br>", i2),
				"it": fmt.Sprintf("%v. Immigrazione<br>", i2),
				"pl": fmt.Sprintf("%v. Polityka imigracyjna<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte eine stärkere Rolle in der Einwanderungspolitik erhalten (z.B. Aufnahmestandards festlegen oder über die Verteilung von Flüchtlingen entscheiden).",
				"en": "The EU should get a stronger role in immigration policy (e.g. decisions over admission standards or allocation of refugees).",
				"es": "La UE debería jugar un papel importante en las políticas migratorias (por ejemplo, definir las condiciones de acogida y decidir acerca del reparto de refugiados).",
				"fr": "L’UE devrait jouer un rôle renforcé dans la politique d’immigration des États membres (par exemple en fixant les normes d’accueil ou en décidant de la répartition des réfugiés).",
				"it": "L’UE dovrebbe avere un ruolo più forte sull’immigrazione (ad esempio sulle decisioni rispetto ai criteri di ammissione o sulla distribuzione dei rifugiati tra i paesi).",
				"pl": "UE powinna odgrywać większą rolę w polityce imigracyjnej (np. ustalając standardy przyjmowania lub decydując o rozmieszczeniu uchodźców).",
			}
		}

		// 13
		{
			names1stMatrix := []string{"defence"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Verteidigungspolitik<br>", i2),
				"en": fmt.Sprintf("%v. Defence policy<br>", i2),
				"es": fmt.Sprintf("%v. Política de defensa<br>", i2),
				"fr": fmt.Sprintf("%v. Politique de défense<br>", i2),
				"it": fmt.Sprintf("%v. Difesa<br>", i2),
				"pl": fmt.Sprintf("%v. Polityka obronna<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Eine unter dem Befehl der EU stehende und aus ihrem Haushalt finanzierte europäische Armee sollte Aufgaben der nationalen Streitkräfte für internationale Kriseneinsätze übernehmen.",
				"en": "A European army under the command of the EU and financed from its budget should take over duties from national armies regarding international conflict deployments.",
				"es": "Un ejército europeo financiado con el presupuesto de la UE y bajo su mando debería asumir las tareas de las fuerzas armadas nacionales en los conflictos internacionales.",
				"fr": "L’UE pourrait constituer une armée européenne placée sous son commandement et financée par son budget avec pour mission d’assurer les opérations extérieures à la place des armées nationales.",
				"it": "Un esercito europeo sotto il comando dell’UE, finanziato dal budget europeo, dovrebbe subentrare alle forze armate nazionali nei conflitti internazionali.",
				"pl": "Armia europejska pod dowództwem UE i finansowana z jej budżetu powinna przejąć zadania krajowych sił zbrojnych w międzynarodowych operacjach kryzysowych.",
			}
		}

		// 14
		{
			names1stMatrix := []string{"wages"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Lohnpolitik<br>", i2),
				"en": fmt.Sprintf("%v. Wage policy<br>", i2),
				"es": fmt.Sprintf("%v. Política salarial<br>", i2),
				"fr": fmt.Sprintf("%v. Politique salariale<br>", i2),
				"it": fmt.Sprintf("%v. Politica salariale<br>", i2),
				"pl": fmt.Sprintf("%v. Polityka płacowa<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte in der Lohnpolitik stärkere Eingriffsrechte erhalten (z.B. bezüglich der Höhe für allgemeine gesetzliche Mindestlöhne).",
				"en": "The EU should have stronger rights to intervene in the wage policies (e.g. regarding the level of general statutory minimum wages).",
				"es": "La UE debería poder intervenir más en la política salarial (por ejemplo, en lo relativo al nivel del salario mínimo interprofesional).",
				"fr": "L’UE pourrait avoir des droits d’intervention plus importants dans les politiques salariales des États membres (par exemple sur le niveau du salaire minimum légal).",
				"it": "L’UE dovrebbe avere maggiore diritto di intervenire sulle politiche salariali (ad esempio riguardo alla definizione del salario minimo obbligatorio).",
				"pl": "UE powinna uzyskać większe prawa interwencji w politykę płacową (np. w odniesieniu do poziomu ogólnego ustawowego wynagrodzenia minimalnego).",
			}
		}

		// 15
		{
			names1stMatrix := []string{"flexibility"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Regeln des Arbeitsmarktes<br>", i2),
				"en": fmt.Sprintf("%v. Labour market regulation<br>", i2),
				"es": fmt.Sprintf("%v. Regulación del mercado laboral<br>", i2),
				"fr": fmt.Sprintf("%v. Marché du travail<br>", i2),
				"it": fmt.Sprintf("%v. Regolamentazione del mercato del lavoro<br>", i2),
				"pl": fmt.Sprintf("%v. Zasady rynku pracy<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die EU sollte den Mitgliedstaaten verbindliche Vorgaben zum Arbeitsmarkt machen dürfen (z.B. zur Ausgestaltung des Kündigungsschutzes oder von Zeitverträgen).",
				"en": "The EU should be able to make binding guidelines to member states regarding the labour market (e.g. regarding the design of dismissal protection or temporary contracts).",
				"es": "La UE debería poder establecer normas vinculantes sobre el mercado laboral (por ejemplo, sobre la definición de la indemnización por despido o los contratos temporales).",
				"fr": "L’UE pourrait donner des directives contraignantes aux États membres relatives au marché du travail (par exemple sur les modalités de la protection des salariés contre les licenciements ou sur les clauses du contrat de travail à durée déterminée).",
				"it": "L’UE dovrebbe essere in grado di fornire linee guida vincolanti agli stati membri sul mercato del lavoro (ad esempio sulle protezioni ai lavoratori per i licenziamenti o per i contratti a tempo determinato).",
				"pl": "UE powinna mieć możliwość ustanowienia wiążących zasad dotyczących rynku pracy dla państw członkowskich (na przykład w odniesieniu do ochrony przed zwolnieniem lub umów na czas określony).",
			}
		}

		// 16
		{
			names1stMatrix := []string{"eutax"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. EU-Steuer als neue Eigenmittelart<br>", i2),
				"en": fmt.Sprintf("%v. EU tax as a new own resource<br>", i2),
				"es": fmt.Sprintf("%v. Impuestos de la UE como nuevo recurso propio<br>", i2),
				"fr": fmt.Sprintf("%v. Taxe commune pour participer au financement du budget européen.<br>", i2),
				"it": fmt.Sprintf("%v. Imposte europee <br>", i2),
				"pl": fmt.Sprintf("%v. Podatek unijny jako nowy rodzaj funduszy własnych<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Es sollte eine neue steuerbasierte Eigenmittelart für den EU-Haushalt unter direkter Kontrolle der EU geben (z.B. eine EU-Steuer auf eine gemeinsame Körperschaftsteuer-Bemessungsgrundlage).",
				"en": "There should be a new tax-based own resource for the EU budget under direct control of the EU (e.g. an EU tax on a common corporate tax base).",
				"es": "Se debe crear un nuevo recurso fiscal propio para financiar el presupuesto europeo bajo el control directo de la UE (por ejemplo, un impuesto de la UE con una base imponible conjunta del impuesto sobre sociedades).",
				"fr": "Le budget européen devrait pouvoir être financé par une taxe commune sous contrôle direct de l’Union Européenne (par exemple un impôt sur les sociétés sur une base fiscale commune dans l’Union Européenne).",
				"it": "Il budget dell’Unione Europea dovrebbe essere finanziato con una nuova entrata tributaria posta sotto il diretto controllo dell’Unione (ad esempio un’imposta sulle società di capitali definita su una base imponibile comune per le imprese europee).",
				"pl": "Do budżetu UE należy wprowadzić nowe, oparte na opodatkowaniu zasoby własne, które podlegałyby bezpośredniej kontroli UE (np. podatek UE od wspólnej podstawy opodatkowania osób prawnych).",
			}
		}

		// 17
		{
			names1stMatrix := []string{"taxpolicy"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Steuerpolitik<br>", i2),
				"en": fmt.Sprintf("%v. Tax policy<br>", i2),
				"es": fmt.Sprintf("%v. Política fiscal<br>", i2),
				"fr": fmt.Sprintf("%v. Politique fiscale<br>", i2),
				"it": fmt.Sprintf("%v. Tassazione<br>", i2),
				"pl": fmt.Sprintf("%v. Polityka podatkowa<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Der Rat der EU sollte mit qualifizierter Mehrheit anstelle von Einstimmigkeit über Steuern beschließen können (z.B. über verbindliche Ober- oder Untergrenzen für Unternehmenssteuern).",
				"en": "The European Council should be able to vote on tax issues with a qualified majority instead of unanimity (e.g. common caps or floors for corporate taxes binding for member states).",
				"es": "El Consejo de la UE debería poder tomar decisiones sobre la fiscalidad por mayoría cualificada (por ejemplo, sobre el límite superior o inferior de los impuestos de sociedades).",
				"fr": "Le Conseil européen devrait pouvoir statuer avec une majorité qualifiée sur les impôts perçus dans les États membres (par exemple sur des taux planchers et plafonds de l’impôt sur les Sociétés).",
				"it": "Il Consiglio Europeo dovrebbe poter votare sulle questioni tributarie a maggioranza qualificata invece che all’unanimità (ad esempio su limiti massimi e minimi comuni a tutti gli stati membri per le imposte sulle imprese).",
				"pl": "Rada UE powinna mieć możliwość podejmowania decyzji w sprawie podatków kwalifikowaną większością głosów, a nie jednomyślnie (na przykład w odniesieniu do wiążących górnych lub dolnych limityów podatków od przedsiębiorstw).",
			}
		}

		// 18
		{
			names1stMatrix := []string{"initiative"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Initiativrechte für das Europäische Parlament<br>", i2),
				"en": fmt.Sprintf("%v. European Parliament and legislative initiative<br>", i2),
				"es": fmt.Sprintf("%v. Derecho de iniciativa del Parlamento Europeo<br>", i2),
				"fr": fmt.Sprintf("%v. Parlement européen et initiative législative<br>", i2),
				"it": fmt.Sprintf("%v. Iniziativa legislative e Parlamento europeo <br>", i2),
				"pl": fmt.Sprintf("%v. Prawo inicjatywy Parlamentu Europejskiego <br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Das Europäische Parlament sollte das Recht erhalten, neue EU-Gesetze vorzuschlagen. Diese Gesetzesinitiative ist bisher der Europäischen Kommission vorbehalten.",
				"en": "The European Parliament should get the right to propose new EU laws (i.e. the legislative initiative) which is currently confined to the European Commission.",
				"es": "El Parlamento Europeo debería tener derecho a proponer nuevas leyes de la UE. Las iniciativas legislativas están por por ahora exclusivamente reservadas a la Comisión Europea.",
				"fr": "Le Parlement Européen devrait pouvoir proposer des lois pour l’Union Européenne (initiative législative), pouvoir qui est aujourd’hui réservé à la Commission Européenne.",
				"it": "Il Parlamento europeo dovrebbe essere dotato del potere di iniziativa legislativa, cioè il diritto di proporre nuove leggi e direttive, potere che al momento è attribuito solo alla Commissione europea.",
				"pl": "Parlament Europejski powinien mieć prawo do proponowania nowych przepisów UE. Ta inicjatywa ustawodawcza przysługuje obecnie Komisji Europejskiej.",
			}
		}

	}

	//
	// Page 2
	{
		p := q.AddPage()
		p.Width = 70
		p.Section = trl.S{
			"de": "2. Reforminitiativen in der Europäischen Währungsunion (EWU)",
			"en": "2. Reform initiatives in the European Monetary Union (EMU)",
			"es": "2. Iniciativas de reformas de la Unión Monetaria Europea (UEM)",
			"fr": "2. Politique monétaire et des finances dans la zone euro",
			"it": "2. Iniziative di riforma dell’Unione Monetaria Europea (UME)",
			"pl": "2. Inicjatywy reformatorskie w ramach Europejskiej Unii Walutowej (EUW)",
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
			"es": "¿En qué medida está de acuerdo con las siguientes afirmaciones?",
			"fr": "La politique monétaire et des finances dans la zone euro fait l’objet de points de vue divergents.  Approuvez-vous les propositions suivantes ?",
			"it": "cosa pensa delle seguenti proposte?",
			"pl": "Wskaż, w jakim stopniu zgadzasz się z poniższymi zadaniami?",
		}
		p.Short = trl.S{
			"de": "Reforminitiativen",
			"en": "Reform initiatives",
			"es": "Iniciativas de reformas",
			"fr": "Politique monétaire et des finances",
			"it": "Iniziative di riforma",
			"pl": "Inicjatywy reformatorskie",
		}

		// 21
		{
			names1stMatrix := []string{"investment"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Höhere Investitionen <br>", i2),
				"en": fmt.Sprintf("%v. Higher investment<br>", i2),
				"es": fmt.Sprintf("%v. Mayores inversiones<br>", i2),
				"fr": fmt.Sprintf("%v. Investissements plus élevés<br>", i2),
				"it": fmt.Sprintf("%v. Investimenti pubblici<br>", i2),
				"pl": fmt.Sprintf("%v. Wyższe inwestycje<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein höheres Wachstum der Eurozone ist es unverzichtbar, dass die Staaten der Eurozone ihre Ausgaben für Investitionen erhöhen.",
				"en": "For higher economic growth of the EMU it is essential that its member states increase their investment expenditures.",
				"es": "Para aumentar el crecimiento en la eurozona es imprescindible que los estados miembros aumenten sus esfuerzos de inversión.",
				"fr": "Une augmentation des dépenses d’investissements des États membres est indispensable pour une croissance plus forte dans la zone euro.",
				"it": "Per una maggiore crescita economica dell’UME è essenziale che gli stati membri aumentino la loro spesa per investimenti.",
				"pl": "Dla zwiększenia wzrostu gospodarczego w strefie euro konieczne jest, aby kraje strefy euro zwiększyły swoje wydatki inwestycyjne.",
			}
		}

		// 22
		{
			names1stMatrix := []string{"more_labor_flex"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Flexiblere Arbeitsmärkte<br>", i2),
				"en": fmt.Sprintf("%v. Flexible labour markets<br>", i2),
				"es": fmt.Sprintf("%v. Mercados de trabajo más flexibles<br>", i2),
				"fr": fmt.Sprintf("%v. Marché du travail plus flexible<br>", i2),
				"it": fmt.Sprintf("%v. Mercato del lavoro flessibile<br>", i2),
				"pl": fmt.Sprintf("%v. Elastyczniejsze rynki pracy<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein höheres Wachstum der Eurozone ist es unverzichtbar, dass insbesondere die Staaten mit anhaltend hoher Arbeitslosigkeit ihre Arbeitsmärkte flexibler machen (z.B. durch eine Lockerung des Kündigungsschutzes oder eine Absenkung von gesetzlichen Mindestlöhnen).",
				"en": "For higher economic growth of the EMU it is essential that especially countries with permanently high levels of unemployment make their labour markets more flexible (e.g. via an easing of dismissal protection regulations or a decrease of the statutory minimum wage).",
				"es": "Para aumentar el crecimiento en la eurozona es imprescindible que los estados miembros, particularmente aquellos con un índice de paro elevado persistente, flexibilicen sus mercados laborales (p. ej. mediante la reducción de la indemnización por despido o la reducción del salario mínimo interprofesional).",
				"fr": "Une croissance plus forte dans la zone euro requiert que les États comptabilisant un nombre important de chômeurs de longue durée rendent plus flexible leur marché du travail (par exemple en assouplissant les conditions de licenciement ou en baissant le niveau du salaire minimum). ",
				"it": "Per una maggiore crescita economica dell’UME è essenziale che i paesi con livelli elevati e permanenti di disoccupazione rendano più flessibile il proprio mercato del lavoro (ad esempio intervenendo sulla regolamentazione relativa ai licenziamenti o, laddove esiste, riducendo il salario minimo obbligatorio).",
				"pl": "Dla zwiększenia wzrostu gospodarczego w strefie euro konieczne jest, aby kraje, w których utrzymuje się wysokie bezrobocie, uelastyczniły swoje rynki pracy (np. poprzez rozluźnienie ochrony przed zwolnieniami lub obniżenie ustawowej płacy minimalnej).",
			}
		}

		// 23
		{
			names1stMatrix := []string{"unemployment_insurance"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Europäische Arbeitslosenversicherung<br>", i2),
				"en": fmt.Sprintf("%v. European unemployment insurance<br>", i2),
				"es": fmt.Sprintf("%v. Seguro de desempleo europeo<br>", i2),
				"fr": fmt.Sprintf("%v. Assurance chômage européenne <br>", i2),
				"it": fmt.Sprintf("%v. Assicurazione europea contro la disoccupazione<br>", i2),
				"pl": fmt.Sprintf("%v. Europejskie ubezpieczenie na wypadek bezrobocia <br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Eine gemeinsame europäische Arbeitslosenversicherung sollte eingeführt werden, um Rezessionen in einzelnen Mitgliedsstaaten der Eurozone abzufedern.",
				"en": "A common European unemployment insurance should be introduced to absorb recessions in individual member states of the EMU.",
				"es": "Se debería introducir un seguro de desempleo europeo para hacer frente a las recesiones en los estados miembros.",
				"fr": "Pour pallier une éventuelle récession de certains États membres dans la zone euro il faudrait créer une assurance chômage européenne commune.",
				"it": "Un sistema comune di assicurazione europea contro la disoccupazione dovrebbe essere introdotto per moderare gli effetti delle crisi economiche nei paesi appartenenti all’UME.",
				"pl": "Należy wprowadzić wspólny europejski system ubezpieczeń na wypadek bezrobocia, aby złagodzić recesje w poszczególnych państwach członkowskich strefy euro.",
			}
		}

		// 24
		{
			names1stMatrix := []string{"eurobonds"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Eurobonds<br>", i2),
				"en": fmt.Sprintf("%v. Eurobonds<br>", i2),
				"es": fmt.Sprintf("%v. Eurobonos<br>", i2),
				"fr": fmt.Sprintf("%v. Euro-obligations<br>", i2),
				"it": fmt.Sprintf("%v. Eurobond<br>", i2),
				"pl": fmt.Sprintf("%v. Euroobligacje<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für Eurobonds haften alle Euro-Staaten gemeinsam und alle Euro-Staaten zahlen den gleichen Zins. Die EWU sollte Eurobonds ausgeben.",
				"en": "All euro countries are jointly liable for Eurobonds and all euro countries pay the same interest. The EMU should issue Eurobonds.",
				"es": "Todos los estados de la zona euro tienen la misma responsabilidad ante los eurobonos y pagan los mismos intereses. La UEM debería emitir bonos europeos.",
				"fr": "La zone euro devrait émettre des euro-obligations et les États membres s’en porter tous garants solidairement et bénéficier du même taux d’intérêt.",
				"it": "Gli Eurobond sono titoli pubblici di debito  di cui tutti i paesi euro sono collettivamente responsabili e su cui tutti i paesi euro pagano gli stessi interessi. L’UME dovrebbe iniziare a emettere Eurobond.",
				"pl": "Wszystkie kraje strefy euro ponoszą wspólną odpowiedzialność za euroobligacje i wszystkie kraje strefy euro płacą takie same odsetki. EUW powinna emitować euroobligacje.",
			}
		}

		// 25
		{
			names1stMatrix := []string{"asset_purchase"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Wertpapierkaufprogramm der EZB<br>", i2),
				"en": fmt.Sprintf("%v. Asset purchase program of ECB<br>", i2),
				"es": fmt.Sprintf("%v. Programa de compra de bonos del BCE<br>", i2),
				"fr": fmt.Sprintf("%v. Achats d’emprunts par la BCE<br>", i2),
				"it": fmt.Sprintf("%v. Programma di acquisti di attività finanziare da parte della BCE<br>", i2),
				"pl": fmt.Sprintf("%v. Program zakupu papierów wartościowych EBC<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Die Europäische Zentralbank (EZB) hat in den zurückliegenden Jahren durch den Kauf von Staatsanleihen von Euro-Staaten eine sehr aktive Rolle gespielt. Diese starke Rolle der EZB sollte fortgesetzt werden.",
				"en": "The European Central Bank (ECB) did take a strongly active position in recent years by purchasing sovereign bonds of euro countries. This strongly active position of the ECB should continue.",
				"es": "El Banco Central Europeo (BCE) ha jugado un papel muy activo en los últimos años con la compra de deuda soberana de los estados miembros. Se debería continuar con este importante papel del BCE.",
				"fr": "Pour combattre la crise, la Banque centrale européenne s’est engagée fortement dans les années passées en achetant des emprunts d’États de la zone euro. Cet engagement devrait se poursuivre.",
				"it": "Negli ultimi anni la Banca Centrale Europea (BCE) ha attuato una politica monetaria molto espansiva comprando titoli di stato dei paesi euro. Questa politica della BCE dovrebbe continuare in futuro.",
				"pl": "Europejski Bank Centralny (EBC) odegrał w ostatnich latach bardzo aktywną rolę, kupując obligacje skarbowe od krajów strefy euro. Ta silna rola EBC powinna być nadal utrzymana.",
			}
		}

		// 26
		{
			names1stMatrix := []string{"growth_pact"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Stabilitäts- und Wachstumspakt (SWP)<br>", i2),
				"en": fmt.Sprintf("%v. Stability and Growth Pact (SGP)<br>", i2),
				"es": fmt.Sprintf("%v. Pacto de estabilidad y crecimiento<br>", i2),
				"fr": fmt.Sprintf("%v. Pacte de Stabilité et de Croissance (PSC)<br>", i2),
				"it": fmt.Sprintf("%v. Patto di Stabilità e Crescita (PSC)<br>", i2),
				"pl": fmt.Sprintf("%v. Pakt na rzecz stabilności i wzrostu (PSW)<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Der SWP definiert Defizit- und Schuldengrenzen für EU-Mitgliedsstaaten. Der SWP schränkt die Fiskalpolitik der Mitgliedsstaaten unangemessen stark ein und sollte gelockert werden.",
				"en": "The SGP defines deficit and debt limits for EU member states. The SGP inappropriately constrains fiscal policy in member states, and should be relaxed.",
				"es": "El Pacto de estabilidad y crecimiento define el límite de déficit y endeudamiento de los estados miembros de la UE. El Pacto de estabilidad y crecimiento limita la política fiscal de los estados miembros de forma excesiva y debería suavizarse.",
				"fr": "Le PCS définit des limites aux déficits et aux dettes des États membres. Le PCS représente une contrainte excessive sur les politiques fiscales des États membres et devrait être assoupli.",
				"it": "Il PSC definisce i limiti per il deficit e il debito pubblico dei paesi membri dell’UE. Il PSC limita eccessivamente la politica fiscale degli stati membri e dovrebbe essere allentato.",
				"pl": "PSW określa limity deficytu i długu państw członkowskich UE. PSW w zbyt wysokim stopniu ogranicza politykę fiskalną państw członkowskich i należałoby to złagodzić.",
			}
		}

		// 27
		{
			names1stMatrix := []string{"emu_institutions"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Neue EWU-Institutionen<br>", i2),
				"en": fmt.Sprintf("%v. New EMU institutions<br>", i2),
				"es": fmt.Sprintf("%v. Nuevas instituciones de la UEM<br>", i2),
				"fr": fmt.Sprintf("%v. Nouvelles institutions pour la zone euro<br>", i2),
				"it": fmt.Sprintf("%v. Nuove istituzioni dell’UME<br>", i2),
				"pl": fmt.Sprintf("%v. Nowe instytucje EUW<br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein angemessenes Funktionieren benötigt die EWU neue Fiskalinstitutionen (z.B. ein Eurozonenbudget oder einen europäischen Finanzminister).",
				"en": "For a proper functioning, the EMU needs new fiscal institutions (e.g. a euro area budget or a European Minister of Finance).",
				"es": "Para garantizar un funcionamiento adecuado, la UEM necesita nuevas instituciones fiscales (por ejemplo, un presupuesto de la eurozona o un ministro de economía europeo).",
				"fr": "Pour son bon fonctionnement, la zone euro devrait développer de nouvelles institutions en matière de fiscalité (par exemple en mettant en place un budget de la zone euro ou un ministre des Finances européen).",
				"it": "Per funzionare correttamente l’UME ha bisogno di nuove istituzioni fiscali (ad esempio un bilancio specifico per l’Eurozona e/o un Ministro delle Finanze europeo).",
				"pl": "Aby EUW mogła właściwie funkcjonować, potrzebuje nowych instytucji fiskalnych (np. budżetu strefy euro lub europejskiego ministra finansów). ",
			}
		}

		// 28
		{
			names1stMatrix := []string{"banking_union"}
			emptyRowLabels := []trl.S{}
			gr := p.AddRadioMatrixGroup(labelsGoodBad17(), names1stMatrix, emptyRowLabels, 1)
			gr.Cols = 9 // necessary, otherwise no vspacers
			gr.OddRowsColoring = true
			gr.Label = trl.S{
				"de": fmt.Sprintf("%v. Vollendung der Bankenunion<br>", i2),
				"en": fmt.Sprintf("%v. Completion of Banking Union<br>", i2),
				"es": fmt.Sprintf("%v. Culminación de la unión bancaria<br>", i2),
				"fr": fmt.Sprintf("%v. Union bancaire<br>", i2),
				"it": fmt.Sprintf("%v. Completamento dell’Unione Bancaria<br>", i2),
				"pl": fmt.Sprintf("%v. Dokończenie budowy unii bankowej <br>", i2),
			}
			gr.Desc = trl.S{
				"de": "Für ein angemessenes Funktionieren sollte die europäische Bankenunion durch die Europäische Einlagensicherung (European Deposit Insurance System: EDIS) vollendet werden.",
				"en": "For its proper functioning, the European Banking Union should be completed through the European Deposit Insurance Scheme (EDIS).",
				"es": "Para garantizar un funcionamiento correcto, debería culminarse la unión bancaria europea mediante la garantía de los depósitos europeos (European Deposit Insurance System: EDIS).",
				"fr": "Pour son bon fonctionnement, l’Union bancaire européenne devrait être complétée par le Système Européen de Garanties des Dépôts (SEGD).",
				"it": "Per funzionare correttamente, l’Unione Bancaria europea dovrebbe essere completata tramite l’introduzione di un sistema europeo di assicurazione dei depositi (EDIS). ",
				"pl": "W celu prawidłowego funkcjonowania należy ukończyć tworzenie Europejskiej Unii Bankowej za pośrednictwem europejskiego systemu gwarantowania depozytów; (European Deposit Insurance System: EDIS).",
			}
		}

	}

	//
	// Page Finish
	{
		p := q.AddPage()
		p.Label = trl.S{
			"de": "Zusammenfassung",
			"en": "Summary of results",
			"es": "Resumen",
			"fr": "Résumé",
			"it": "Riepilogo",
			"pl": "Krótki opis",
		}
		// p.NoNavigation = true
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
					"en": "<span style='font-size: 100%;'>Your entries have been saved</span>",
					"es": "<span style='font-size: 100%;'>Sus entradas se han guardado</span>",
					"fr": "<span style='font-size: 100%;'>Vos entrées ont été sauvegardées</span>",
					"it": "<span style='font-size: 100%;'>Le tue iscrizioni sono state salvate</span>",
					"pl": "<span style='font-size: 100%;'>Twoje wpisy zostały zapisane</span>",
				}
			}

			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dynamic"
			// 	inp.CSSLabel = "special-input-vert-wider"
			// 	inp.DynamicFunc = "RepsonseStatistics"
			// }

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-input-vert-wider"
				impr := trl.S{}
				for lc := range q.LangCodes {
					cnt, err := tpl.MarkDownFromFile("./static/doc/site-imprint.md", q.Survey.Type, lc)
					if err != nil {
						log.Print(err)
					}
					impr[lc] = cnt
				}
				inp.Desc = impr
			}

		}
	}

	(&q).Hyphenize()
	(&q).ComputeMaxGroups()
	return &q, nil
}
