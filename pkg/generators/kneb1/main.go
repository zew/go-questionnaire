package kneb1

import (
	"fmt"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create questionnaire
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"de", "en"} // governs default language cod,
	q.LangCodes = []string{"de"}
	q.LangCode = "de"

	q.Survey.Org = trl.S{
		"en": "ZEW",
		"de": "ZEW",
	}
	q.Survey.Name = trl.S{
		"en": "Financial Literacy Test",
		"de": "Financial Literacy Test",
	}
	// q.Variations = 1

	// page 0
	{
		page := q.AddPage()
		page.ValidationFuncName = ""

		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true

		page.Label = trl.S{
			"en": "Dear Madam / Sir,",
			"de": "Sehr geehrte Damen und Herren",
		}
		// pge.Short = trl.S{
		// 	"en": "Greeting",
		// 	"de": "Begrüßung",
		// }

		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./welcome-page.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page 1
	{
		page := q.AddPage()

		page.Section = trl.S{
			"de": "Soziodemographie",
			"en": "Sociodemographics",
		}
		page.Label = trl.S{
			"de": "Alter, Herkunft, Erfahrungen",
			"en": "Age, origin, experience",
		}
		page.Short = trl.S{
			"de": "Soziodemo-<br>graphie",
			"en": "Sociodemo-<br>graphics",
		}
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 6
			var radioValues = []string{
				"male",
				"female",
				"diverse",
			}
			var labels = []trl.S{
				{
					"de": "männlich",
					"en": "male",
				},
				{
					"de": "weiblich",
					"en": "female",
				},
				{
					"de": "divers",
					"en": "diverse",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "Sind Sie…",
					"en": "What is your gender?",
				}.Outline("D1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd1_gender"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 2
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 6
			{
				inp := gr.AddInput()
				inp.Type = "dropdown"
				inp.Name = "qd2_birthyear"
				inp.MaxChars = 20
				inp.MaxChars = 10

				inp.Label = trl.S{
					"de": "In welchem Jahr sind Sie geboren?",
					"en": "Your year of birth?",
				}.Outline("D2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 4

				//
				inp.DD = &qst.DropdownT{}
				if false {
					inp.DD.AddPleaseSelect(cfg.Get().Mp["must_one_option"])
				} else {
					inp.DD.Add(
						"",
						trl.S{
							"de": " Bitte ein Jahr wählen ",
							"en": " Please choose a year ",
						},
					)
				}
				for year := 1930; year < time.Now().Year()-10; year++ {
					inp.DD.Add(
						fmt.Sprintf("%d", year),
						trl.S{
							"de": fmt.Sprintf("%d", year),
							"en": fmt.Sprintf("%d", year),
						},
					)
				}

			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 6
			var radioValues = []string{
				"baden_wuerttemberg",
				"bayern",
				"berlin",
				"brandenburg",
				"bremen",
				"hamburg",
				"hessen",
				"mecklenburg_vorpommern",
				"niedersachsen",
				"nordrhein_westfalen",
				"rheinland_pfalz",
				"saarland",
				"sachsen",
				"sachsen_anhalt",
				"schleswig_holstein",
				"thueringen",
			}
			var labels = []trl.S{
				{
					"de": "Baden-Württemberg",
					"en": "Baden-Württemberg",
				},
				{
					"de": "Bayern",
					"en": "Bavaria",
				},
				{
					"de": "Berlin",
					"en": "Berlin",
				},
				{
					"de": "Brandenburg",
					"en": "Brandenburg",
				},
				{
					"de": "Bremen",
					"en": "Bremen",
				},
				{
					"de": "Hamburg",
					"en": "Hamburg",
				},
				{
					"de": "Hessen",
					"en": "Hesse",
				},
				{
					"de": "Mecklenburg-Vorpommern",
					"en": "Mecklenburg-Vorpommern",
				},
				{
					"de": "Niedersachsen",
					"en": "Lower Saxony",
				},
				{
					"de": "Nordrhein-Westfalen",
					"en": "North Rhine-Westphalia",
				},
				{
					"de": "Rheinland-Pfalz",
					"en": "Rhineland-Palatinate",
				},
				{
					"de": "Saarland",
					"en": "Saarland",
				},
				{
					"de": "Sachsen",
					"en": "Saxony",
				},
				{
					"de": "Sachsen-Anhalt",
					"en": "Saxony-Anhalt",
				},
				{
					"de": "Schleswig-Holstein",
					"en": "Schleswig-Holstein",
				},
				{
					"de": "Thüringen",
					"en": "Thuringia",
				},
			}

			{

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": "In welchem Bundesland wohnen Sie?",
						"en": "Which German state you live in?",
					}.Outline("D3.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				for idx, label := range labels {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "qd3_german_state"
					rad.ValueRadio = radioValues[idx]

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}

		if false {
			// gr3
			{
				gr := page.AddGroup()
				gr.Cols = 4
				gr.BottomVSpacers = 1

				{
					inp := gr.AddInput()
					inp.Type = "number"
					inp.Name = "q04_age"

					inp.Label = trl.S{
						"de": "Wie alt sind Sie?",
						"en": "How old are you?",
					}
					inp.MaxChars = 4
					inp.Step = 1
					inp.Min = 15
					inp.Max = 100
					inp.Validator = "inRange100"

					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 3
					inp.ColSpanControl = 2
					inp.Suffix = trl.S{
						"de": "&nbsp; Jahre",
						"en": "&nbsp; years",
					}
				}

				{
					inp := gr.AddInput()
					inp.Type = "text"
					inp.Name = "q05_country_birth"

					inp.Label = trl.S{
						"de": "In welchem Land sind Sie geboren?",
						"en": "What is your country of birth?",
					}

					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 3
					inp.ColSpanControl = 2
					inp.MaxChars = 20
					// inp.Validator = "must"
				}

			}

		}

	}

	// page 2
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Alter, Herkunft, Erfahrungen 2",
			"en": "Age, origin, experience 2",
		}
		page.SuppressInProgressbar = true
		// page.Short = trl.S{
		// 	"de": "Soziodemo-<br>graphie 2",
		// 	"en": "Sociodemo-<br>graphics 2",
		// }
		page.WidthMax("42rem")

		page.ValidationFuncName = "knebPageD2"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 6
			var radioValues = []string{
				"hauptschule",
				"realschule",
				"fachhochschule",
				"abitur",
			}
			var labels = []trl.S{
				{
					"de": "Hauptschul-/ Volksschulabschluss",
					"en": "Primary school",
				},
				{
					"de": "Mittlere Reife/ Realschulabschluss/ Polytechnische Oberschule, 10. Klasse",
					"en": "Secondary school",
				},
				{
					"de": "Fachhochschulreife",
					"en": "College",
				},
				{
					"de": "Allgemeine oder fachgebundene Hochschulreife/ Abitur",
					"en": "High school",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "Welchen höchsten allgemeinbildenden Schulabschluss haben Sie?",
					"en": "Which is your highest degree?",
				}.Outline("D4.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd4_abschluss"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 6
			var radioValues = []string{
				"no",
				"yes",
				"highschool",
			}
			var labels = []trl.S{
				{
					"de": "Nein, keine abgeschlossene Berufsausbildung",
					"en": "todo",
				},
				{
					"de": "Ja, abgeschlossene Berufsausbildung",
					"en": "todo",
				},
				{
					"de": "Ja, Hochschulabschluss (Fachhochschule oder Universität)",
					"en": "todo",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Haben Sie eine abgeschlossene Berufsausbildung?  <br>
						<small> Falls es mehrere Abschlüsse sind, geben Sie bitte nur den höchsten an. </small>
					`,
					"en": `
						todo

					`,
				}.Outline("D5.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd5_vocational_training"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}
		}

		// gr6
		{
			gr := page.AddGroup()
			gr.Cols = 6
			var radioValues = []string{
				"single_livingalone",
				"unmarried_livingtogether",
				"married_livingtogether",
				"divorcedwidowed_livingalone",
				"divorcedwidowed_livingtogether",
			}
			var labels = []trl.S{
				{
					"de": "Ledig ohne Partner/in im Haushalt",
					"en": "todo",
				},
				{
					"de": "Ledig mit Partner/in im Haushalt",
					"en": "todo",
				},
				{
					"de": "Verheiratet und zusammenlebend",
					"en": "todo",
				},
				{
					"de": "Geschieden/getrennt lebend/verwitwet ohne Partner/in im Haushalt",
					"en": "todo",
				},
				{
					"de": "Geschieden/getrennt lebend/verwitwet mit Partner/in im Haushalt",
					"en": "todo",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `Welchen Familienstand haben Sie?`,
					"en": `What is your marital status?`,
				}.Outline("D6.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd6_family_status"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}
		}

		// gr7
		{
			gr := page.AddGroup()
			gr.Cols = 6
			var radioValues = []string{
				"above35hours",
				"between15and35hours",
				"upto15hours",
				"occasionally",
				"none",
			}
			var labels = []trl.S{
				{
					"de": "Vollzeiterwerbstätig mit einer wöchentlichen Arbeitszeit von 35&nbsp;Stunden oder mehr",
					"en": "todo",
				},
				{
					"de": "Teilzeiterwerbstätig mit einer wöchentlichen Arbeitszeit von 15 bis unter 35&nbsp;Stunden",
					"en": "todo",
				},
				{
					"de": "Geringfügig beschäftigt mit einer wöchentlichen Arbeitszeit unter 15&nbsp;Stunden",
					"en": "todo",
				},
				{
					"de": "Gelegentlich erwerbstätig",
					"en": "todo",
				},
				{
					"de": "In keiner Weise erwerbstätig",
					"en": "todo",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Sind Sie zurzeit in irgendeiner Weise erwerbstätig? 
						<br>
						<small>
						Unter Erwerbstätigkeit wird jede bezahlte bzw. mit einem Einkommen verbundene Tätigkeit verstanden, egal welchen zeitlichen Umfang sie hat. Was auf dieser Liste trifft am besten zu?
						</small>
					`,
					"en": `
						todo
					`,
				}.Outline("D7.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd7_employment"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}
		}

		// gr7
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 0
			var radioValues = []string{
				"housewife",
				"unemployed",
				"ineducation",
				"military",
				"parental",
				"other",
			}
			var labels = []trl.S{
				{
					"de": "Hausfrau / Hausmann",
					"en": "todo",
				},
				{
					"de": "Arbeitslos",
					"en": "todo",
				},
				{
					"de": "In Ausbildung, Schule, Lehre, Studium oder Umschulung",
					"en": "todo",
				},
				{
					"de": "Wehr- oder Ersatzdienst",
					"en": "todo",
				},
				{
					"de": "Mutterschafts-/ Erziehungsurlaub bzw. Elternzeit",
					"en": "todo",
				},
				{
					"de": "Sonstiges",
					"en": "todo",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Wenn Sie einmal von den Erwerbstätigkeiten absehen, was von dem Folgenden trifft dann auf Sie zu?
					`,
					"en": `
						todo
					`,
				}.Outline("D7a.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd7a_employment"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}
		}

		// gr8
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 0
			var radioValues = []string{
				"worker",
				"employee",
				"civilservant",
				"selfemployedalone",
				"selfemployedwithemployees",
			}
			var labels = []trl.S{
				{
					"de": "Arbeiter",
					"en": "todo",
				},
				{
					"de": "Angestellter",
					"en": "todo",
				},
				{
					"de": "Beamter",
					"en": "todo",
				},
				{
					"de": "Selbständig ohne Mitarbeiter",
					"en": "todo",
				},
				{
					"de": "Selbständig mit Mitarbeitern",
					"en": "todo",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Sind Sie zurzeit…					`,
					"en": `
						todo
					`,
				}.Outline("D7a.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qd7b_employment"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}
		}

	}

	// page 3
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Generelles Risiko",
			"en": "Risk in general",
		}
		page.Short = trl.S{
			"de": "Generelles<br>Risiko",
			"en": "Risk<br>in general",
		}
		page.Short = trl.S{
			"de": "Risiko",
			"en": "Risk",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				labelsRisk(),
				[]string{"qm1_risk"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie schätzen Sie sich persönlich ein: 
					Sind Sie im Allgemeinen ein risikobereiter Mensch 
					oder versuchen Sie, Risiken zu vermeiden?

					<small>
					Antworten Sie bitte anhand der folgenden Skala, 
					wobei der Wert 0 bedeutet: gar nicht risikobereit 
					und der Wert 10: sehr risikobereit.
					Mit den Werten dazwischen können Sie Ihre Einschätzung abstufen.
					</small>
				`,
				"en": `
					todo
				`,
			}.Outline("M1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.12rem"
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `Woran haben Sie gedacht, als Sie die Frage nach dem allgemeinen Risiko beantwortet haben?
						<small>(Mehrfachnennung möglich)</small>
					`,
					"en": `todo`,
				}.Outline("M2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qm2_risk_free"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsPositiveAspects(),
				[]string{"qm3_averse"},
				radioVals7,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Haben Sie eher an die negativen oder
					 positiven Seiten des Risikos gedacht?

					<small>
					Antworten Sie bitte anhand der folgenden Skala, wobei der Wert 1 bedeutet: 
					nur an die negativen Seiten und der 
					Wert 7: nur an die positiven Seiten.
					Mit den Werten dazwischen können Sie Ihre Einschätzung abstufen.
					</small>
				`,
				"en": `
					todo
				`,
			}.Outline("M3.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsImportantSituations(),
				[]string{"qm4_occasions"},
				radioVals7,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Haben Sie eher an kleine Alltagssituationen 
					oder an große, wichtige Situationen gedacht?

					<small>
					Antworten Sie bitte anhand der folgenden Skala, wobei der Wert 1 bedeutet: 
					kleine Alltagssituationen und der Wert 7: 
					große, wichtige Situationen.
					Mit den Werten dazwischen können Sie Ihre Einschätzung abstufen.

					</small>
				`,
				"en": `
					todo
				`,
			}.Outline("M4.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsReturns(),
				[]string{"qm5_returns"},
				radioVals7,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Haben Sie eher an Situationen gedacht, in denen es kleine oder große Gewinne gibt?
		
					<small>
					Antworten Sie bitte anhand der folgenden Skala, 
					wobei der Wert 1 bedeutet: kleine Gewinne und der Wert 7: 
					große Gewinne.
					Mit den Werten dazwischen können Sie Ihre Einschätzung abstufen.
					</small>
				`,
				"en": `
							todo
						`,
			}.Outline("M5.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr5
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsLosses(),
				[]string{"qm6_losses"},
				radioVals7,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Haben Sie eher an Situationen gedacht, in denen es kleine oder große Verluste gibt?
		
					<small>
					Antworten Sie bitte anhand der folgenden Skala, wobei der Wert 1 bedeutet: 
					kleine Verluste und der Wert 7: große Verluste.
					Mit den Werten dazwischen können Sie Ihre Einschätzung abstufen.
					</small>
				`,
				"en": `
							todo
						`,
			}.Outline("M6.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 4
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Finanzielles Risiko",
			"en": "Financial risk",
		}
		page.SuppressInProgressbar = true
		// page.Short = trl.S{
		// 	"de": "Finanzielles<br>Risiko",
		// 	"en": "Financial<br>risk",
		// }

		page.WidthMax("42rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				labelsFinRisk(),
				[]string{"qf1_risk"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Bitte schätzen Sie Ihre Bereitschaft ein, ein <i>finanzielles</i> Risiko einzugehen.
		
					<small>
					Bewerten Sie bitte anhand der Skala von 1 bis 5. 
					</small>
				`,
				"en": `
					todo
				`,
			}.Outline("F1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.12rem"
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `Woran haben Sie gedacht, als Sie die Frage nach dem allgemeinen Risiko beantwortet haben?
						<small>(Mehrfachnennung möglich)</small>
					`,
					"en": `todo`,
				}.Outline("F2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf2_risk_free"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.12rem"
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `Denken Sie bitte einmal an die Risiken und Chance im Zusammenhang mit dem Finanzmarkt. 
						Was kommt Ihnen dabei in den Sinn?

						<br>
						<br>
						Chancen<br>
						<small>(Mehrfachnennung möglich)</small>
					`,
					"en": `todo`,
				}.Outline("F3.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_chance1"
				inp.MaxChars = 120
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_chance2"
				inp.MaxChars = 120
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_chance3"
				inp.MaxChars = 120
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Risiken<br>
						<small>(Mehrfachnennung möglich)</small>
					`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_risk1"
				inp.MaxChars = 120
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_risk2"
				inp.MaxChars = 120
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_risk3"
				inp.MaxChars = 120
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}

		}

	}

	// page 5
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Präferenzen und Einschätzungen",
			"en": "Preferences and assessments",
		}

		page.Short = trl.S{
			// including next page
			"de": "Präferenzen,<br>Einschätzungen<br>Vorsorge",
			"en": "Preferences,<br>assessments<br>provisions",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				labelsSelfKnowledge(),
				[]string{"qp1_risk"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie beurteilen Sie Ihr persönliches Wissen hinsichtlich finanzieller Angelegenheiten?
				`,
				"en": `
					todo
				`,
			}.Outline("P1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		//
		//
		//
		lblsQP2 := labelsSelfKnowledge()
		lblsQP2[0] = trl.S{
			"de": "<small>stimme gar nicht zu</small>     <div>0</div>",
			"en": "<small>dont agree at all  </small>     <div>0</div>",
		}
		lblsQP2[10] = trl.S{
			"de": "<small>stimme voll und ganz zu</small> <div>10</div>",
			"en": "<small>agree completely       </small> <div>10</div>",
		}
		inputs := []string{
			"qp2a_boring",
			"qp2b_fear",
			"qp2c_trust_people",
			"qp2d_trust_institutions",
		}
		lbls := []trl.S{
			{
				"de": `
				Wie sehr stimmen Sie den folgenden Aussagen zu?

				<br>
				<br>

				<b>a)</b>&nbsp; Persönliche Finanzen finde ich langweilig.
				
				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				<b>b)</b>&nbsp; Ich habe große Angst vor finanziellen Verlusten.
				
				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				<b>c)</b>&nbsp; Im Allgemeinen kann man den Menschen vertrauen.
				
				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				<b>d)</b>&nbsp; Im Allgemeinen kann man Banken und Finanzinstitutionen in Deutschland vertrauen.
				
				`,
				"en": `
					todo
				`,
			},
		}

		for i := 0; i < len(inputs); i++ {
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lblsQP2,
				[]string{inputs[i]},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = lbls[i]
			if i == 0 {
				gb.MainLabel.Outline("P2.")
			}
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 6
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Vorsorgeplanung ",
			"en": "Foresight",
		}
		page.SuppressInProgressbar = true
		// page.Short = trl.S{
		// 	"de": "Vorsorgeplanung",
		// 	"en": "Foresight",
		// }
		page.WidthMax("42rem")
		page.WidthMax("44rem")

		lblsQV1to6 := labelsImportantSituations()
		lblsQV1to6[0] = trl.S{
			"de": "<small>stimme überhaupt nicht zu</small> 0",
			"en": "<small>dont agree at all</small>         0",
		}
		lblsQV1to6[6] = trl.S{
			"de": "<small>stimme voll und ganz zu</small> 7",
			"en": "<small>agree completely</small>        7",
		}

		//
		//
		inputs := []string{
			"saving",
			"knowledge",
			"secure",
			"oversight",
			"avoidance",
			"procrastination",
			"emphasis",
			"fear",
		}
		lbls := []trl.S{
			{
				"de": `
			Inwiefern stimmen Sie den folgenden Aussagen zu?
			
			<small>
			Antworten Sie auf der Skala von 
			1: "Stimme überhaupt nicht zu" 
			bis 
			7: "Stimme voll und ganz zu" .
			</small>

			<br>
			<br>

			<b>%c)</b>&nbsp; Ich spare genug für die Rente.
				
		`,
				"en": `
			<b>%c)</b>&nbsp; todo.
		`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Ich beschäftige mich ausreichend mit dem Thema Rente.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Ich fühle mich gut für das Alter abgesichert.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Ich habe heute einen guten Überblick über meine angesammelten Rentenansprüche.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Ich habe noch genug Zeit bis zum Ruhestand, um mich um meine Altersvorsorge zu kümmern.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Ich habe es noch nicht geschafft, mich um meine Altersvorsorge zu kümmern.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Mir ist es wichtig, dass ich für das Alter ausreichend abgesichert bin.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
			{
				"de": `	
					<b>%c)</b>&nbsp; Ich habe Angst vor Armut im Alter.
				`,
				"en": `
					<b>%c)</b>&nbsp; todo.
				`,
			},
		}

		for i := 0; i < len(inputs); i++ {
			rn := rune(97 + i) // ascii 65 is A; 97 is a
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				lblsQV1to6,
				[]string{fmt.Sprintf("qv1%c_%s", rn, inputs[i])},
				radioVals7,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = lbls[i].Fill(rn)
			gr := page.AddGrid(gb)
			// _ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 7
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Financial numeracy and literacy",
			"en": "Financial numeracy and literacy",
		}
		page.Short = trl.S{
			"de": "Financial<br>literacy",
			"en": "Financial<br>literacy",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// a func to create questions 1a, 2a, 3a, 4a
		howSicher := func(inputName, outlineNumber string) {

			lblsQF1a := labelsSelfKnowledge()
			lblsQF1a[0] = trl.S{
				"de": "<small>nicht sicher</small>     <div>0</div>",
				"en": "<small>not sure</small>         <div>0</div>",
			}
			lblsQF1a[10] = trl.S{
				"de": "<small>sehr sicher</small>      <div>10</div>",
				"en": "<small>very sure</small>        <div>10</div>",
			}

			// append one more
			lblsQF1aCp := make([]trl.S, len(lblsQF1a)+1)
			copy(lblsQF1aCp, lblsQF1a)
			lblsQF1aCp[11] = trl.S{
				"de": "<small style='padding-left: 3.0rem; text-align: left'>ich weiß die Antwort nicht, ich habe geraten</small> ",
				"en": "<small style='padding-left: 3.0rem; text-align: left'>I dont know the answer, I guessed.</small>           ",
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate12,
				lblsQF1aCp,
				[]string{inputName},
				radioVals12,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
						Wie sicher sind Sie sich bei Ihrer Antwort? .
					`,
				"en": `
						todo
					`,
			}.Outline(outlineNumber)
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 4
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"

		}

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				labelsInterestOverTwoYears(),
				[]string{"qfl1_interest"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Angenommen, Sie haben 100&nbsp;€ Guthaben auf Ihrem Sparkonto. 
					Dieses Guthaben wird mit 2% pro Jahr verzinst, 
					und Sie lassen es 5&nbsp;Jahre auf diesem Konto. 
					Was meinen Sie: Wie hoch wird ihr Guthaben nach 5&nbsp;Jahren sein?				
					<div style='color:red'>ist das richtig? fünf Jahre? </div>
				`,
				"en": `
					todo
				`,
			}.Outline("FL1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr1
		howSicher("qfl1a_free", "FL1a.")

		// gr2
		{
			lbls := labelsInterestOverTwoYears()
			lbls[0] = trl.S{
				"de": "<small>mehr</small>    ",
				"en": "<small>more</small>    ",
			}
			lbls[1] = trl.S{
				"de": "<small>genauso viel</small>    ",
				"en": "<small>equal</small>           ",
			}
			lbls[2] = trl.S{
				"de": "<small>weniger als heute</small>  ",
				"en": "<small>less than today</small>    ",
			}

			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				lbls,
				[]string{"qfl2_inflation"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Angenommen, die Verzinsung Ihres Sparkontos beträgt 1% pro Jahr 
					und die Inflationsrate beträgt 2% pro Jahr. 
					Was glauben Sie: 
					Werden Sie nach einem Jahr mit dem Guthaben des Sparkontos genauso viel, 
					mehr oder weniger als heute kaufen können?
				`,
				"en": `
					todo
				`,
			}.Outline("FL2.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr3
		howSicher("qfl2a_free", "FL2a.")

		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsStimmeZuOderNicht(),
				[]string{"qfl3_portfoliorisk"},
				radioVals4,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Stimmen Sie der folgenden Aussage zu: 
					„Die Anlage in Aktien eines einzelnen Unternehmens ist weniger riskant 
					als die Anlage in einem Fonds mit Aktien ähnlicher Unternehmen“?
				`,
				"en": `
					todo
				`,
			}.Outline("FL3.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr5
		howSicher("qfl3a_free", "FL3a.")

		//
		// gr6
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 1
			{
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": "Was ist die Hauptfunktion des Aktienmarktes?",
						"en": "todo",
					}.Outline("FL4.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				for idx, label := range labelsStockMarketPurpose() {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "qfl4_stockmarketpurpose"
					rad.ValueRadio = fmt.Sprintf("%d", idx)

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}

		// // gr6
		// if false {
		// 	gb := qst.NewGridBuilderRadios(
		// 		colsStockMarket,
		// 		labelsStockMarketPurpose(),
		// 		[]string{"qfl4_stockmarketpurpose"},
		// 		radioVals6,
		// 		[]trl.S{{"de": ``, "en": ``}},
		// 	)
		// 	gb.MainLabel = trl.S{
		// 		"de": `
		// 			Was ist die Hauptfunktion des Aktienmarktes?
		// 			<div style='color:red'>D3</div>
		// 		`,
		// 		"en": `
		// 			todo
		// 		`,
		// 	}.Outline("FL4.")
		// 	gr := page.AddGrid(gb)
		// 	_ = gr
		// 	gr.BottomVSpacers = 1
		// 	gr.Style = css.NewStylesResponsive(gr.Style)
		// 	gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		// 	gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		// 	gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		// }

		// gr7
		howSicher("qfl4a_free", "FL4a.")

		//
		// gr8
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 1
			{
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": "Welche der folgenden Anlageformen zeigt im Laufe der Zeit die höchsten Ertragsschwankungen?",
						"en": "todo",
					}.Outline("FL5.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				for idx, label := range labelsAssetClassVola() {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "qfl5_assetclassvola"
					rad.ValueRadio = fmt.Sprintf("%d", idx)

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}

		// // gr8
		// {
		// 	gb := qst.NewGridBuilderRadios(
		// 		colsAssetClasses,
		// 		labelsAssetClassVola(),
		// 		[]string{"qfl5_assetclassvola"},
		// 		radioVals6,
		// 		[]trl.S{{"de": ``, "en": ``}},
		// 	)
		// 	gb.MainLabel = trl.S{
		// 		"de": `
		// 			Welche der folgenden Anlageformen zeigt im Laufe der Zeit die höchsten Ertragsschwankungen?
		// 		`,
		// 		"en": `
		// 			todo
		// 		`,
		// 	}.Outline("FL5.")
		// 	gr := page.AddGrid(gb)
		// 	_ = gr
		// 	gr.BottomVSpacers = 1
		// 	gr.Style = css.NewStylesResponsive(gr.Style)
		// 	gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		// 	gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		// 	gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"
		// }

		// gr9
		howSicher("qfl5a_free", "FL5a.")

		//
		// gr10
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 1
			{

				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": `
						Es besteht eine 50/50 Chance, dass Maliks Auto innerhalb der nächsten 
						sechs Monate eine Motorreparatur benötigt, die 1.000 Euro kosten würde. 
						
						Gleichzeitig besteht eine 10%-ige Chance, 
						dass er die Klimaanlage in seinem Haus ersetzen muss, 
						was 4.000 Euro kosten würde. 
						
						Welches ist das größere finanzielle Risiko für Malik?
							
						`,
						"en": "todo",
					}.Outline("FL6.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				for idx, label := range labelsExpectedValueRisk() {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "qfl6_expectedvalue"
					rad.ValueRadio = fmt.Sprintf("%d", idx)

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}

		// // gr10
		// {
		// 	gb := qst.NewGridBuilderRadios(
		// 		columnTemplate4,
		// 		labelsExpectedValueRisk(),
		// 		[]string{"qfl6_expectedvalue"},
		// 		radioVals4,
		// 		[]trl.S{{"de": ``, "en": ``}},
		// 	)
		// 	gb.MainLabel = trl.S{
		// 		"de": `
		// 			Es besteht eine 50/50 Chance, dass Maliks Auto innerhalb der nächsten
		// 			sechs Monate eine Motorreparatur benötigt, die 1.000 Euro kosten würde.

		// 			Gleichzeitig besteht eine 10%-ige Chance,
		// 			dass er die Klimaanlage in seinem Haus ersetzen muss,
		// 			was 4.000 Euro kosten würde.

		// 			Welches ist das größere finanzielle Risiko für Malik?
		// 		`,
		// 		"en": `
		// 			todo
		// 		`,
		// 	}.Outline("FL6.")
		// 	gr := page.AddGrid(gb)
		// 	_ = gr
		// 	gr.BottomVSpacers = 1
		// 	gr.Style = css.NewStylesResponsive(gr.Style)
		// 	gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		// }

		// gr11
		howSicher("qfl6a_free", "FL6a.")

	}

	// page 8
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Selbstvertrauen vor Experiment",
			"en": "Confidence before experiment",
		}

		// for next three pages
		page.Short = trl.S{
			"de": "Experiment",
			"en": "Epxperiment",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{

			lbls := labelsSelfKnowledge()
			lbls[0] = trl.S{
				"de": "<small>kein Vertrauen in die eigenen Fähigkeiten</small>     <div>0</div>",
				"en": "<small>todo</small>     <div>0</div>",
			}
			lbls[10] = trl.S{
				"de": "<small>hohes Vertrauen in die eigenen Fähigkeiten</small> <div>10</div>",
				"en": "<small>todo</small> <div>10</div>",
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lbls,
				[]string{"qe1_confidence_before"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie viel Vertrauen haben Sie in Ihre Fähigkeit, gute finanzielle Entscheidungen zu treffen?
				`,
				"en": `
					todo
				`,
			}.Outline("E1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page chart
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "Experiment-Chart",
			"de": "Experiment chart",
		}
		page.SuppressInProgressbar = true
		// page.Short = trl.S{
		// 	"de": "Chart",
		// 	"en": "Chart",
		// }
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./echart/inner.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page x+0
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Selbstvertrauen nach Experiment",
			"en": "Confidence after experiment",
		}
		page.SuppressInProgressbar = true
		// page.Short = trl.S{
		// 	"de": "Selbstvertrauen<br>nachher",
		// 	"en": "Confidence<br>after",
		// }
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{

			lbls := labelsSelfKnowledge()
			lbls[0] = trl.S{
				"de": "<small>kein Vertrauen in die eigenen Fähigkeiten</small>     <div>0</div>",
				"en": "<small>todo</small>     <div>0</div>",
			}
			lbls[10] = trl.S{
				"de": "<small>hohes Vertrauen in die eigenen Fähigkeiten</small> <div>10</div>",
				"en": "<small>todo</small> <div>10</div>",
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lbls,
				[]string{"qe1_confidence_after"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
						Wie viel Vertrauen haben Sie in Ihre Fähigkeit, gute finanzielle Entscheidungen zu treffen?
					`,
				"en": `
						todo
					`,
			}.Outline("E1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page x+1
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Erfahrungen mit Finanzen",
			"en": "Experience in finance",
		}

		// includes next page
		page.Short = trl.S{
			"de": "Erfahrungen<br>Finanzen,<br>Beratung",
			"en": "Experience<br>finance,<br>advice",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		//
		//
		//
		lblsQX1 := labelsSelfKnowledge()
		lblsQX1[0] = trl.S{
			"de": "<small>überhaupt nicht</small>     <div>0</div>",
			"en": "<small>not at all     </small>     <div>0</div>",
		}
		lblsQX1[10] = trl.S{
			"de": "<small>sehr gerne</small> <div>10</div>",
			"en": "<small>very like </small> <div>10</div>",
		}
		lblsQX2 := labelsSelfKnowledge()
		lblsQX2[0] = trl.S{
			"de": "<small>überhaupt nicht</small>     <div>0</div>",
			"en": "<small>not at all     </small>     <div>0</div>",
		}
		lblsQX2[10] = trl.S{
			"de": "<small>sehr viel</small> <div>10</div>",
			"en": "<small>very much</small> <div>10</div>",
		}
		inputs := []string{
			"qx1_discuss",
			"qx2_decisionfear",
			"qx3_worry",
			"qx4_instinct",
		}
		lbls := []trl.S{
			{
				"de": `
					Unterhalten Sie sich gerne über Geld und Geldanlagen?				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
					Haben Sie Angst davor, Geld zu investieren oder finanzielle Entscheidungen zu treffen?			
				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
					Machen Sie sich Sorgen über den Erfolg Ihrer finanziellen Entscheidungen?
				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
					Sind Sie der Meinung, dass Investitionsentscheidungen am Ende nur vom Instinkt abhängen?			
				`,
				"en": `
					todo
				`,
			},
		}

		for i := 0; i < len(inputs); i++ {

			hdrs := lblsQX1
			if i > 0 {
				hdrs = lblsQX2
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				hdrs,
				[]string{inputs[i]},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = lbls[i]
			if i == 0 || true {
				gb.MainLabel.Outline(fmt.Sprintf("X%v.", i+1))
			}
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Haben Sie während Ihrer Schul- oder Berufsausbildung an Vorlesungen, 
					Kursen oder Fortbildungen zum Thema Finanzen oder dem Umgang mit Geld teilgenommen?
				`,
				"en": `todo`,
			},
			"qx5_courses",
			"X5.",
			false,
		)

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`
						Wir stellen Ihnen nun eine Frage zum Finanzvermögen: <br>

						Haben Sie (d.h. Ihr Haushalt) im Jahr <i>%d</i> eine der folgenden Vermögensarten besessen? <br>
					
						<small>
						Falls Sie nicht wissen, ob ihr Partner diese Vermögensarten besitzt, 
						beantworten Sie die Fragen bitte für sich selbst.
						</small>
					
					`, time.Now().Year(),
					),
					"en": `todo`,
				}.Outline("X6.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}

		}

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Sparanlagen

					<span style='color:red' >Klammern weglassen?</span>

					<br>

					<small>
					(z.B. Sparbücher, Festgeldkonten, Tagesgeldkonten oder Sparverträge)
					</small>
				`,
				"en": `todo`,
			},
			"qx6a_savingaccount",
			"a)",
			true,
		)

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Bausparverträge

					<br>
					
					<small>
					(die noch nicht in Darlehen umgewandelt wurden)						
					</small>
				`,
				"en": `todo`,
			},
			"qx6b_realestatesa",
			"b)",
			true,
		)

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Festverzinsliche Wertpapiere

					<br>
					
					<small>
					(z. B. Spar- oder Pfandbriefe, Bundesschatzbriefe, Industrieanleihen oder Anteile an Rentenfonds) 						
					</small>
				`,
				"en": `todo`,
			},
			"qx6c_bonds",
			"c)",
			true,
		)

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Aktien oder Aktienfonds und Immobilienfonds

					<br>
					
					<small>
					(auch Aktienanleihen, börsennotierte Fonds, ETFs, Mischfonds oder ähnliche Anlagen)
					</small>
				`,
				"en": `todo`,
			},
			"qx6d_stocks",
			"d)",
			true,
		)

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Sonstige Wertpapiere

					<br>
					
					<small>
					(z.B. Discountzertifikate, Hedgefonds, Filmfonds, Windenergiefonds, Geldmarktfonds und andere Finanzinnovationen)
					</small>
				`,
				"en": `todo`,
			},
			"qx6e_other",
			"e)",
			true,
		)

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Gold
				`,
				"en": `todo`,
			},
			"qx6f_gold",
			"f)",
			true,
		)

	}

	// page x+2
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Erfahrungen mit Beratung",
			"en": "Experience with advice",
		}

		page.SuppressInProgressbar = true
		// page.Short = trl.S{
		// 	"de": "Erfahrungen<br>Beratung",
		// 	"en": "Experience<br>with advice",
		// }
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		page.ValidationFuncName = "knebPageCounseling"
		page.ValidationFuncMsg = trl.S{"de": "no javascript dialog message needed"}

		//
		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				labelsPensionAdvice(),
				[]string{"qb1_pensionadvice"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wann haben Sie sich das letzte Mal zum Thema Altersvorsorge beraten lassen?
				`,
				"en": `
					todo
				`,
			}.Outline("B1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsFrequency(),
				[]string{"qb2_frequency"},
				radioVals4,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie oft sprechen Sie mit Ihrem Berater oder Ihrer Beraterin?
				`,
				"en": `
					todo
				`,
			}.Outline("B2.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsFrequency(),
				[]string{"qb3_frequency"},
				radioVals4,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie oft sprechen Sie mit Ihrer Familie oder Ihren Freunden über Finanzen?
				`,
				"en": `
					todo
				`,
			}.Outline("B3.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Haben Sie selbst schon einmal einer Person finanzielle Ratschläge gegeben?
				`,
				"en": `todo`,
			},
			"qb4_advisingself",
			"B4.",
			false,
		)

		//
		//
		lblsQX1 := labelsSelfKnowledge()
		lblsQX1[0] = trl.S{
			"de": "<small>überhaupt nicht</small>     <div>0</div>",
			"en": "<small>not at all     </small>     <div>0</div>",
		}
		lblsQX1[10] = trl.S{
			"de": "<small>sehr viel</small> <div>10</div>",
			"en": "<small>very much</small> <div>10</div>",
		}
		inputs := []string{
			"qb5_fearofloss",
			"qb6_delegate",
		}
		lbls := []trl.S{
			{
				"de": `
					Haben Sie Angst, bei finanziellen Entscheidungen Verluste zu machen?	
				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
					Würden Sie es vorziehen, wenn eine andere Person finanzielle Entscheidungen für Sie trifft?
				`,
				"en": `
					todo
				`,
			},
		}

		for i := 0; i < len(inputs); i++ {

			hdrs := lblsQX1
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				hdrs,
				[]string{inputs[i]},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = lbls[i]
			if i == 0 || true {
				gb.MainLabel.Outline(fmt.Sprintf("B%v.", i+5))
			}
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr last
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate4,
				labelsWhoIsCompetent(),
				[]string{"qb7_whocompetent"},
				radioVals4,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wer könnte finanzielle Entscheidungen am besten für Sie treffen?
				`,
				"en": `
					todo
				`,
			}.Outline("B7.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page x+3
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Haushaltseinkommen und Vermögen",
			"en": "Household income and assets",
		}
		page.Short = trl.S{
			"de": "Haushalts-<br>einkommen<br>Vermögen",
			"en": "Household income,<br>assets",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr 0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`
						Wer ist in Ihrem Haushalt hauptsächlich für folgendes zuständig?
					`,
					),
					"en": `todo`,
				}.Outline("H1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
		}

		// gr 1
		meOrTogether(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					alltägliche Einkäufe (z.B. Lebensmitteleinkäufe)
				`,
				"en": `todo`,
			},
			"qh1a_shopping",
			"a)",
			true,
		)

		// gr 2
		meOrTogether(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					größere Anschaffungen (z.B. Möbel, Auto)
				`,
				"en": `todo`,
			},
			"qh1b_largerpurchases",
			"b)",
			true,
		)

		// gr 3
		meOrTogether(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Essensplanung und -zubereitung
				`,
				"en": `todo`,
			},
			"qh1c_foodpreparation",
			"c)",
			true,
		)

		// gr 4
		meOrTogether(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Entscheidungen über Spar- und Finanzanlagen
				`,
				"en": `todo`,
			},
			"qh1d_financialdecisions",
			"d)",
			true,
		)

		// gr5
		{
			gr := page.AddGroup()
			gr.Cols = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Wie hoch schätzen Sie, 
						ist das <i>monatlich</i> verfügbare Nettoeinkommen Ihres Haushalts, 
						also dasjenige Geld, das dem gesamten Haushalt nach Abzug 
						von Steuern und Sozialversicherungsbeiträgen zur Deckung der Ausgaben
						 zur Verfügung steht? 
					`,
					"en": `
						todo
					`,
				}.Outline("H2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0

			}

			{
				inp := gr.AddInput()
				inp.Type = "dropdown"
				inp.Name = "qh2_income"
				inp.MaxChars = 20
				inp.MaxChars = 10
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				//
				inp.DD = &qst.DropdownT{}
				if false {
					inp.DD.AddPleaseSelect(cfg.Get().Mp["must_one_option"])
				} else {
					inp.DD.Add(
						"",
						trl.S{
							"de": " Bitte einen Bereich wählen ",
							"en": " Please choose a range",
						},
					)
				}
				ranges := []int{
					0,
					500,
					750,
					1000,
					1000,
					1250,
					1500,
					2000,
					2500,
					3000,
					3000,
					3500,
					4000,
					4500,
					5000,
					//
					7500,
					10000,
					15000,
				}
				for i := 0; i < len(ranges); i++ {

					rLow := ranges[i]
					rHigh := -1
					if i < len(ranges)-1 {
						rHigh = ranges[i+1]
					}
					opt := fmt.Sprintf("upto%d", rHigh)
					chLbl := trl.S{
						"de": fmt.Sprintf("%d€ bis unter %d€", rLow, rHigh),
						"en": fmt.Sprintf("%d€ to under  %d€", rLow, rHigh),
					}
					if i == 0 {
						chLbl = trl.S{
							"de": fmt.Sprintf("unter %d€", rHigh),
							"en": fmt.Sprintf("under %d€", rHigh),
						}

					}
					if rHigh == -1 {
						opt = "over15000"
						chLbl = trl.S{
							"de": fmt.Sprintf("%d€ und mehr", rLow),
							"en": fmt.Sprintf("%d€ and more", rLow),
						}

					}
					inp.DD.Add(opt, chLbl)

				}
				inp.DD.Add(
					"noanswer",
					trl.S{
						"de": " keine Angabe",
						"en": " no answer",
					},
				)

			}
		}

		// gr 6
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": fmt.Sprintf(`
						Nun sind wir an Ihrem Altersvorsorgevermögen interessiert:<br>
						
						Haben Sie (d.h. Ihr Haushalt) im <i>Dezember %d </i>
						einen der folgenden privaten oder betrieblichen Altersvorsorgeverträge besessen?

						<small>
						Falls Sie nicht wissen, ob ihr Partner diese Vermögensarten besitzt, 
						beantworten Sie die Fragen bitte für sich selbst.
						</small>
					
					`,
						// december previous year - for 2023: 2022
						time.Now().Year()-1,
					),
					"en": `todo`,
				}.Outline("H1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
		}

		// gr 7
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Private Lebensversicherungen
					<small> 
					(z.B. klassische und fondsge¬bundene Kapitallebensversicherungen, 
						<i>nicht</i> reine Risikolebensversicherungen 
						oder Direktversicherungen über den Arbeitgeber)
					</small> 
				
				`,
				"en": `todo`,
			},
			"qh3a_lifeinsurance",
			"a)",
			true,
		)

		// gr 8
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Betriebliche Lebensversicherungen
					<small> 
					(z. B. Direktversicherungen)
					</small> 
				
				`,
				"en": `todo`,
			},
			"qh3b_directinsurance",
			"b)",
			true,
		)

		// gr 9
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Sonstige betriebliche Altersvorsorge 
					<small> 
					(z. B. Betriebsrenten aus Pensions- oder Unterstützungskassen und betriebliche Direktzusagen sowie Zusatzversorgung im öffentlichen Dienst; auch aus früheren Beschäftigungsverhältnissen) 
					</small> 
				
				`,
				"en": `todo`,
			},
			"qh3c_otherpensions",
			"c)",
			true,
		)

		// gr 10
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Staatlich geförderte private Altersvorsorge („Riester- Rente“)
					<small> 
					(staatlich geförderte und zertifizierte Sparanlagen, auch „Rürup-“ bzw. Basisrenten) 
					</small> 
				
				`,
				"en": `todo`,
			},
			"qh3d_otherpensions",
			"d)",
			true,
		)

		// gr 11
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Private Rentenversicherungen 
					<small> 
					(z.B. private Rentenversicherungsverträge, die nicht staatlich gefördert werden bzw. abgeschlossen wurden, bevor es solche Fördermöglichkeiten gab) 
					</small> 				
				`,
				"en": `todo`,
			},
			"qh3e_privatepensions",
			"e)",
			true,
		)

		//
		//
		// comment
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
					"de": `
						<hr>
						<br>
						<em>Zum Abschluss des Fragebogens</em><br>
						Wollen Sie uns noch etwas mitteilen?
					`,
					"en": `
						<hr>
						<br>
						<em>Finish</em><br>
						Any remarks or advice for us?
					`,
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "comment"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// advance to last page "data saved"
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				// inp.Label = trl.S{
				// 	"de": "Durch Klicken erhalten Sie eine Zusammenfassung Ihrer Antworten",
				// 	"en": "By clicking, you will receive a summary of your answers.",
				// }
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

	}

	// Report of results
	{
		page := q.AddPage()
		page.NoNavigation = true
		page.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}

		page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "ResponseStatistics"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PersonalLink"
			}
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "site-imprint.md"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
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
