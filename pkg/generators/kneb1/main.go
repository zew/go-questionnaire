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
		"en": "Umfrage zum Thema Finanzentscheidungen",
		"de": "Umfrage zum Thema Finanzentscheidungen",
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
		page.Label = trl.S{
			"de": "",
			"en": "",
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

		// page.Section = trl.S{
		// 	"de": "Soziodemographie",
		// 	"en": "Sociodemographics",
		// }
		page.Label = trl.S{
			"de": "Alter, Herkunft, Erfahrungen",
			"en": "Age, origin, experience",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
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
				}.OutlineHid("D1.")
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
				}.OutlineHid("D2.")
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
				// for year := 1930; year < time.Now().Year()-10; year++ {
				for year := 1968; year <= 2005; year++ {
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
					}.OutlineHid("D3.")
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
		page.Label = trl.S{
			"de": "",
			"en": "",
		}

		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			// single column
			gr.Cols = 3
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
				}.OutlineHid("D4.")
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
			// single column
			gr.Cols = 3
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
				}.OutlineHid("D5.")
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
	}

	// page 2-split
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 6
			// single column
			gr.Cols = 3
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
				}.OutlineHid("D6.")
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

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 6
			// single column
			gr.Cols = 3
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
				}.OutlineHid("D7.")
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

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 6
			// single column
			gr.Cols = 3
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
				}.OutlineHid("D7a.")
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

		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 6
			// single column
			gr.Cols = 3
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
				}.OutlineHid("D7a.")
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

			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebPageD2" // js filename

				s1 := trl.S{
					"de": "no javascript dialog message needed",
					"en": "no javascript dialog message needed",
				}
				inp.JSBlockTrls = map[string]trl.S{
					"msg": s1,
				}

				inp.JSBlockStrings = map[string]string{}
				inp.JSBlockStrings["pageID"] = fmt.Sprintf("pg%02v", len(q.Pages)-1)

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
		page.Label = trl.S{
			"de": "",
			"en": "",
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
			}.OutlineHid("M1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}
	// page 3-split-1
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "",
			"en": "",
		}

		page.SuppressInProgressbar = true

		page.WidthMax("42rem")
		page.WidthMax("48rem")

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
				}.OutlineHid("M2.")
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

	}

	// page 3-split-2
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "",
			"en": "",
		}

		page.SuppressInProgressbar = true

		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
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
			}.OutlineHid("M3.")
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
			}.OutlineHid("M4.")
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
			}.OutlineHid("M5.")
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
			}.OutlineHid("M6.")
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
		page.Label = trl.S{
			"de": "",
			"en": "",
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
					Antworten Sie bitte anhand der folgenden Skala, 
					wobei der Wert 1 bedeutet: nicht bereit, ein Risiko einzugehen 
					und der Wert 5: bereit, ein erhebliches Risiko einzugehen, 
					um potenziell eine höhere Rendite zu erzielen.
					</small>
				`,
				"en": `
					todo
				`,
			}.OutlineHid("F1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 4-split
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr0
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
					"de": `
						Woran haben Sie gedacht, als Sie die Frage nach dem <i>finanziellen</i> Risiko beantwortet haben?
						<br>
						<small>Bitte geben Sie ein oder mehrere Stichwörter an.</small>
					`,
					"en": `todo`,
				}.OutlineHid("F2.")
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

	}

	// page 4-split2
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")

		// gr0
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
						<b>Chancen</b><br>
						<small>(Mehrfachnennung möglich)</small>
					`,
					"en": `todo`,
				}.OutlineHid("F3.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf3_chance1"
				inp.Validator = "must"
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
						<b>Risiken</b><br>
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
				inp.Validator = "must"
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
		page.Label = trl.S{
			"de": "",
			"en": "",
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
			}.OutlineHid("P1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 5-split
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

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

				Persönliche Finanzen finde ich langweilig.

				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				Ich habe große Angst vor finanziellen Verlusten.

				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				Im Allgemeinen kann man den Menschen vertrauen.

				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				Im Allgemeinen kann man Banken und Finanzinstitutionen in Deutschland vertrauen.

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
			gb.MainLabel = lbls[i].Outline(fmt.Sprintf("P%d.", i+2)) // P2., P3., P4., P5.
			// if i == 0 {
			// 	gb.MainLabel.OutlineHid("P2.")
			// }
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 5-new
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("44rem")

		// gr 0
		{
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = 3

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					
						Stellen Sie sich vor, Sie haben 2500&nbsp;Euro in der Lotterie gewonnen
						und können dieses Geld in ein Unternehmen investieren.

						Die Chance, dass das Unternehmen erfolgreich ist, liegt bei 50 Prozent.

						Im Erfolgsfall verdoppelt sich Ihre Investition nach einem Jahr.

						Bei Misserfolg verlieren Sie die Hälfte der investierten Summe.
						Welchen Anteil der 2500&nbsp;Euro würden Sie in dieses Unternehmen investieren?
					`,
					"en": `todo`,
				}.OutlineHid("P6.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "qp6_risky_investment"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 6
				inp.Label = trl.S{
					"de": `					
						Ich würde 
					`,
					"en": `todo`,
				}
				inp.Suffix = trl.S{
					"de": `					
						% in das Unternehmen investieren. 
					`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 4
			}
		}

		// gr 1
		{
			gr := page.AddGroup()
			gr.Cols = 3
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Certainty equivalent experiment<br>
						Siehe Menkhoff und Sakha (2016, 2017) Appendix A.3.<br>
						todo
					`,
					"en": `todo`,
				}.OutlineHid("P7.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

		}
	}

	vorsorgeplanung8x7(&q, 0, 2)
	vorsorgeplanung8x7(&q, 3, 5)
	vorsorgeplanung8x7(&q, 6, 7)

	// page 7-0
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Financial numeracy and literacy",
			"en": "Financial numeracy and literacy",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			"de": "Financial<br>literacy",
			"en": "Financial<br>literacy",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3,
				[]trl.S{

					{
						"de": "<small>höher   als 102&nbsp;€</small>      ",
						"en": "<small>higher than 102&nbsp;€</small>    ",
					},
					{
						"de": "<small>genau   102&nbsp;€</small>        ",
						"en": "<small>exactly 102&nbsp;€</small>        ",
					},
					{
						"de": "<small>niedriger als 102&nbsp;€</small>  ",
						"en": "<small>lower than    102&nbsp;€</small>     ",
					},
				},
				[]string{"qfl1_interest"},
				radioVals3,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Angenommen, Sie haben 100&nbsp;€ Guthaben auf Ihrem Sparkonto.
					Dieses Guthaben wird mit 2% pro Jahr verzinst,
					und Sie lassen es 5&nbsp;Jahre auf diesem Konto.
					Was meinen Sie: Wie hoch wird ihr Guthaben nach 5&nbsp;Jahren sein?
				`,
				"en": `
					todo
				`,
			}.OutlineHid("FL1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 3
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 7-1
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		howSicher(*qst.WrapPageT(page), "qfl1a_free", "FL1a.")

	}

	// page 7-2
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate3,
				[]trl.S{
					{
						"de": "<small>mehr</small>    ",
						"en": "<small>more</small>    ",
					},
					{
						"de": "<small>genauso viel</small>    ",
						"en": "<small>equal</small>           ",
					},
					{
						"de": "<small>weniger als heute</small>  ",
						"en": "<small>less than today</small>    ",
					},
				},
				[]string{"qfl2_inflation"},
				radioVals3,
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
			}.OutlineHid("FL2.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 3
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 7-3
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		howSicher(*qst.WrapPageT(page), "qfl2a_free", "FL2a.")
	}

	// page 7-4
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate2,
				labelsStimmeZuOderNicht(),
				[]string{"qfl3_portfoliorisk"},
				radioVals2,
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
			}.OutlineHid("FL3.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 3
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page 7-5
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		howSicher(*qst.WrapPageT(page), "qfl3a_free", "FL3a.")

	}

	// page 7-6
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 3
			{
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": "Was ist die Hauptfunktion des Aktienmarktes?",
						"en": "todo",
					}.OutlineHid("FL4.")
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

	}

	// page 7-7
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		howSicher(*qst.WrapPageT(page), "qfl4a_free", "FL4a.")

	}

	// page 7-8
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 3
			{
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": "Welche der folgenden Anlageformen zeigt im Laufe der Zeit die höchsten Ertragsschwankungen?",
						"en": "todo",
					}.OutlineHid("FL5.")
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

	}

	// page 7-9
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		howSicher(*qst.WrapPageT(page), "qfl5a_free", "FL5a.")
	}

	// page 7-10
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 3
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
					}.OutlineHid("FL6.")
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
	}

	// page 7-11
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		howSicher(*qst.WrapPageT(page), "qfl6a_free", "FL6a.")

	}

	// page 8
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Selbstvertrauen vor Experiment",
			"en": "Confidence before experiment",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
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

			lbls := labelsSelfKnowledgeXX()
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
					<br>
					Antworten Sie bitte anhand der folgenden Skala,
					wobei der Wert 0 bedeutet: Kein Vertrauen in die eigenen Fähigkeiten
					und der Wert 10: Hohes Vertrauen in die eigenen Fähigkeiten.

				`,
				"en": `
					todo
				`,
			}.OutlineHid("E1.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page chart introduction 1
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "Experiment-Chart-Introduction",
			"de": "Experiment chart-Introduction",
		}
		page.Label = trl.S{
			"en": "",
			"de": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("52rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.Label = trl.S{
					"de": `
					<p>
					In diesem Teil unserer Umfrage geht es um <i>Investitionsentscheidungen im Rahmen der Altersvorsorge</i>. Genügend finanzielle Mittel im Alter sind der Grundstein für einen sicheren und stabilen Ruhestand.

					Vor allem im Ruhestand, wenn das regelmäßige Renteneinkommen im Durchschnitt niedriger ist als das Arbeitseinkommen während der Erwerbstätigkeit, ist es wichtig, dass man finanziell gut abgesichert ist.
					</p>

					<p>
					Es gibt viele verschiedene Möglichkeiten für das Alter vorzusorgen. Im Folgenden stellen wir Ihnen eine dieser Möglichkeiten vor: <i>Einen monatlichen Sparbetrag über einen längeren Zeithorizont am Kapitalmarkt anzulegen</i>.
					</p>

					<p>
					In unserer interaktiven Graphik versuchen wir, die Chancen und Risiken einer Anlage am Kapitalmarkt zu verdeutlichen.
					</p>

	
					`,
					"en": `
							todo
						`,
				}
			}
		}

	}

	// page chart introduction 2
	//   guided tour
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "Experiment-Chart-Introduction",
			"de": "Experiment chart-Introduction",
		}
		page.Label = trl.S{
			"en": "",
			"de": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("52rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				// store current page of the guided tour
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "section"
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
				inp.DynamicFuncParamset = "./slide-show/index.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page chart
	{
		page := q.AddPage()

		page.Label = trl.S{
			"en": "Experiment-Chart",
			"de": "Experiment chart",
		}
		page.Label = trl.S{
			"en": "",
			"de": "",
		}
		page.SuppressInProgressbar = true

		page.WidthMax("42rem")
		page.WidthMax("52rem")
		page.WidthMax("58rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "share_safe_bg"
			}
			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "share_risky_bg"
			}
			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "sparbetrag_bg"
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
				inp.DynamicFuncParamset = "./echart/inner.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

		// advance to next page
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1-1)
				inp.Label = trl.S{
					"de": "Zurück",
					"en": "todo",
				}
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "p"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "start"
			}
			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since next page is appended below
				inp.Label = trl.S{
					"de": "Werte speichern und weiter",
					"en": "todo",
				}
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
			}
		}

	}

	// page x+0
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Feedbackfragen zum Tool",
			"en": "todo",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true

		page.WidthMax("42rem")
		page.WidthMax("48rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						In wieweit stimmen Sie folgenden Aussagen zu:
					`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			inputs := []string{
				"qe2_comprehensive",
				"qe3_helplful",
				"qe4_complex",
			}
			lblsQ := []trl.S{
				{
					"de": `Die Darstellung des Portfolios ist verständlich.`,
					"en": `todo`,
				},
				{
					"de": `Die Darstellung des Portfolios ist hilfreich.`,
					"en": `todo`,
				},
				{
					"de": `Die Darstellung des Portfolios ist kompliziert.`,
					"en": `todo`,
				},
			}

			for i := 0; i < len(inputs); i++ {

				lbls := labelsSelfKnowledge()
				lbls[0] = trl.S{
					"de": "<small>trifft ganz und gar nicht zu</small>     <div>0</div>",
					"en": "<small>todo</small>     <div>0</div>",
				}
				lbls[10] = trl.S{
					"de": "<small>trifft voll und ganz zu</small>       <div>10</div>",
					"en": "<small>todo</small> <div>10</div>",
				}
				gb := qst.NewGridBuilderRadios(
					columnTemplate11,
					lbls,
					[]string{inputs[i]},
					radioVals11,
					[]trl.S{{"de": ``, "en": ``}},
				)
				// gb.MainLabel = lblsQ[i].OutlineHid(fmt.Sprintf("%c)", rune(97+i)))
				gb.MainLabel = lblsQ[i].OutlineHid(fmt.Sprintf("E%v.", i+2)) // .Outline("E2., E3., E4.")
				gr := page.AddGrid(gb)
				_ = gr
				// gr.BottomVSpacers = 2
				gr.Style = css.NewStylesResponsive(gr.Style)
				gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			}

		}

		//
		//
		{

			lbls := labelsSelfKnowledge()
			lbls[0] = trl.S{
				"de": "<small>ganz und gar uninformiert</small>     <div>0</div>",
				"en": "<small>todo</small>     <div>0</div>",
			}
			lbls[10] = trl.S{
				"de": "<small>ganz und gar informiert</small>       <div>10</div>",
				"en": "<small>todo</small> <div>10</div>",
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lbls,
				[]string{"qe5_feelinginformed"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie informiert fühlen Sie sich über das gezeigte Szenario?
					`,
				"en": `
					todo
				`,
			}.OutlineHid("E5.")
			gr := page.AddGrid(gb)
			_ = gr
			// gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

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
					"de": `Haben Sie weitere Anmerkungen zu der Darstellung im Tool?`,
					"en": `todo`,
				}.OutlineHid("E6.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qe6_tool_free"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.12rem"
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `Möchten Sie uns sonst noch etwas mitteilen?`,
					"en": `todo`,
				}.OutlineHid("E7.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qe7_tool_other"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

	}

	// page x+1
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Treatment: Giving advice",
			"en": "todo",
		}
		// page.Label = trl.S{
		// 	"de": "",
		// 	"en": "",
		// }

		// page.SuppressInProgressbar = true
		page.Short = trl.S{
			"de": "Treatment:<br>Giving advice",
			"en": "todo",
		}
		page.WidthMax("42rem")
		page.WidthMax("48rem")

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
					"de": `
						Bitte denken Sie über Ihre Erfahrungen nach, die Sie mit dem Tool gesammelt haben. Was können Sie zukünftigen Nutzerinnen und Nutzern empfehlen? Welche Ratschläge können Sie weitergeben, um anderen zu helfen?
					`,
					"en": `todo`,
				}.OutlineHid("T1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qt1_advice"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.12rem"
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Welche Ihrer Erkenntnisse/ Ergebnisse aus dem Tool könnten für andere Nutzerinnen und Nutzer besonders interessant sein? Was ist besonders hilfreich? 					
					`,
					"en": `todo`,
				}.OutlineHid("T2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qt2_which"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

	}

	// page x+2
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Selbstvertrauen nach dem Tool",
			"en": "Confidence afterwards",
		}
		// page.Label = trl.S{
		// 	"de": "",
		// 	"en": "",
		// }
		page.Short = trl.S{
			"de": "Selbstvertrauen<br>nachher",
			"en": "Confidence<br>after",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("42rem")
		page.WidthMax("48rem")

		// gr0
		{

			lbls := labelsSelfKnowledgeXX()
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lbls,
				[]string{"qe8_confidence_after"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
						Wie viel Vertrauen haben Sie in Ihre Fähigkeit, gute finanzielle Entscheidungen zu treffen?
						<br>
						Antworten Sie bitte anhand der folgenden Skala,
						wobei der Wert 0 bedeutet: Kein Vertrauen in die eigenen Fähigkeiten
						und der Wert 10: Hohes Vertrauen in die eigenen Fähigkeiten.

					`,
				"en": `
						todo
					`,
			}.OutlineHid("E8.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 3
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page x+3
	erfahrungMitFinanzenSplit1(&q, 0, 0)
	erfahrungMitFinanzenSplit2(&q, 0, 0)

	// page x+4
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Wann haben Sie sich das letzte Mal zum Thema Altersvorsorge beraten lassen?
					`,
					"en": `
						todo
					`,
				}.OutlineHid("B1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, lbl := range labelsPensionAdvice() {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qb1_pensionadvice"
				rad.ValueRadio = fmt.Sprintf("%v", idx+1)

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = lbl
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}

		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				[]trl.S{
					{
						"de": "einmal im Monat",
						"en": "todo",
					},
					{
						"de": "einmal im Quartal",
						"en": "todo",
					},
					{
						"de": "einmal im Jahr",
						"en": "todo",
					},
					{
						"de": "seltener als einmal im Jahr",
						"en": "todo",
					},
					{
						"de": "ich habe keine/n Finanzberater/ Finanzberaterin",
						"en": "todo",
					},
				},
				[]string{"qb2_frequency"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie oft sprechen Sie mit Ihrem Finanzberater oder Ihrer Finanzberaterin?
				`,
				"en": `
					todo
				`,
			}.OutlineHid("B2.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate5,
				labelsFrequency(),
				[]string{"qb3_frequency"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie oft sprechen Sie mit Ihrer Familie oder Ihren Freunden über Finanzen?
				`,
				"en": `
					todo
				`,
			}.OutlineHid("B3.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 3
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 0
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebPageCounseling1" // js filename

				s1 := trl.S{
					"de": "no javascript dialog message needed",
					"en": "no javascript dialog message needed",
				}
				inp.JSBlockTrls = map[string]trl.S{
					"msg": s1,
				}

				inp.JSBlockStrings = map[string]string{}
				inp.JSBlockStrings["pageID"] = fmt.Sprintf("pg%02v", len(q.Pages)-1)

			}

		}

	}

	//
	// page x+5
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Haben Sie selbst schon einmal einer anderen Person finanzielle Ratschläge gegeben?
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
				gb.MainLabel.OutlineHid(fmt.Sprintf("B%v.", i+5))
			}
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		//
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
			}.OutlineHid("B7.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 1
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 0
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.ColSpanControl = 1
				inp.Type = "javascript-block"
				inp.Name = "knebPageCounseling2" // js filename

				s1 := trl.S{
					"de": "no javascript dialog message needed",
					"en": "no javascript dialog message needed",
				}
				inp.JSBlockTrls = map[string]trl.S{
					"msg": s1,
				}

				inp.JSBlockStrings = map[string]string{}
				inp.JSBlockStrings["pageID"] = fmt.Sprintf("pg%02v", len(q.Pages)-1)

			}

		}

	}

	// page x+6
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Haushaltseinkommen und Vermögen",
			"en": "Household income and assets",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			"de": "Haushalts-<br>einkommen<br>Vermögen",
			"en": "Household income,<br>assets",
		}
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
						Wer ist in Ihrem Haushalt hauptsächlich für Folgendes zuständig?
					`,
					),
					"en": `todo`,
				}.OutlineHid("H1.")
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
	}

	// page x+7
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Wie hoch <i style='font-size:110%'>schätzen</i> Sie,
						ist das <i style='font-size:110%'>monatlich</i> verfügbare 
						<i style='font-size:110%'>Nettoeinkommen Ihres Haushalts</i>,
						also dasjenige Geld, das dem gesamten Haushalt nach Abzug
						von Steuern und Sozialversicherungsbeiträgen zur Deckung der Ausgaben
						 zur Verfügung steht?
					`,
					"en": `
						todo
					`,
				}.OutlineHid("H2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0

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
				3500,
				4000,
				4500,
				5000,
				//
				7500,
				10000,
				15000,
			}
			lbls := make([]trl.S, 0, len(ranges)+1)
			opts := make([]string, 0, len(ranges)+1)

			for i := 0; i < len(ranges); i++ {

				rLow := ranges[i]
				rHigh := -1
				if i < len(ranges)-1 {
					rHigh = ranges[i+1]
				}
				opt := fmt.Sprintf("upto%d", rHigh)
				lbl := trl.S{
					"de": fmt.Sprintf("%d€ bis unter %d€", rLow, rHigh),
					"en": fmt.Sprintf("%d€ to under  %d€", rLow, rHigh),
				}
				if i == 0 {
					lbl = trl.S{
						"de": fmt.Sprintf("unter %d€", rHigh),
						"en": fmt.Sprintf("under %d€", rHigh),
					}
				}
				if rHigh == -1 {
					opt = "over15000"
					lbl = trl.S{
						"de": fmt.Sprintf("%d€ und mehr", rLow),
						"en": fmt.Sprintf("%d€ and more", rLow),
					}

				}
				lbls = append(lbls, lbl)
				opts = append(opts, opt)
			}
			lbls = append(lbls, trl.S{
				"de": " keine Angabe",
				"en": " no answer",
			})
			opts = append(opts, "noanswer")

			//
			for idx, lbl := range lbls {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qh2_income"
				rad.ValueRadio = fmt.Sprintf("%v", idx+1)

				rad.ColSpan = gr.Cols
				rad.ColSpan = 3
				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = lbl
				rad.ControlFirst()
				rad.LabelTop()
				rad.ControlTop()
			}

			if false {
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
				for idx, lbl := range lbls {
					inp.DD.Add(opts[idx], lbl)
				}
			}

		}
	}

	// page x+8
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
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
				}.OutlineHid("H3.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
		}

		// gr 1
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					Private Lebensversicherungen

					<small>
					z.B. klassische und fondsgebundene Kapitallebensversicherungen,
						<i>nicht</i> reine Risikolebensversicherungen
						oder Direktversicherungen über den Arbeitgeber
					</small>

				`,
				"en": `todo`,
			},
			"qh3a_lifeinsurance",
			"a)",
			true,
		)

		// gr 2
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

		// gr 3
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

		// gr 4
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

		// gr 5
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

	}

	// page x+9
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Abschluss",
			"en": "todo",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

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
				inp.Label = trl.S{
					"de": `
						<strong>
						Danke, dass Sie an unserer Umfrage teilgenommen haben!
						</strong>
						<br>
						<br>

						Gibt es noch irgendetwas, was Sie uns zu diesem Fragebogen 
						oder Thema mitteilen möchten?
					`,
					"en": `
						todo
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
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since next page is appended below
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
