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

	// page0
	{
		page := q.AddPage()
		page.ValidationFuncName = ""

		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true

		page.Label = trl.S{
			"en": "Dear Madam / Sir,",
			"de": "Sehr geehrte Damen und Herren",
		}
		// page.Short = trl.S{
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
					"de": "Männlich",
					"en": "Male",
				},
				{
					"de": "Weiblich",
					"en": "Female",
				},
				{
					"de": "Divers",
					"en": "Diverse",
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
		page.Short = trl.S{
			"de": "Soziodemo-<br>graphie 2",
			"en": "Sociodemo-<br>graphics 2",
		}
		page.SuppressInProgressbar = true
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
		page.Short = trl.S{
			"de": "Finanzielles<br>Risiko",
			"en": "Financial<br>risk",
		}
		page.WidthMax("42rem")
		// page.WidthMax("48rem")

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
			"de": "Präferenzen,<br>Einschätzungen",
			"en": "Preferences,<br>assessments",
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

		lblsQP2aTo2d := labelsSelfKnowledge()
		lblsQP2aTo2d[0] = trl.S{
			"de": "<small>stimme gar nicht zu</small> 0",
			"en": "<small>dont agree at all</small> 0",
		}
		lblsQP2aTo2d[10] = trl.S{
			"de": "<small>stimme voll und ganz zu</small> 10",
			"en": "<small>agree completely</small>        10",
		}

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lblsQP2aTo2d,
				[]string{"qp2a_boring"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
				Wie sehr stimmen Sie den folgenden Aussagen zu?

				<br>
				<br>

				<b>a)</b>&nbsp; Persönliche Finanzen finde ich langweilig.
				
				`,
				"en": `
					todo
				`,
			}.Outline("P2.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lblsQP2aTo2d,
				[]string{"qp2b_fear"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
				<b>b)</b>&nbsp; Ich habe große Angst vor finanziellen Verlusten.
				
				`,
				"en": `
					todo
				`,
			}
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lblsQP2aTo2d,
				[]string{"qp2c_trust_people"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
				<b>c)</b>&nbsp; Im Allgemeinen kann man den Menschen vertrauen.
				
				`,
				"en": `
					todo
				`,
			}
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		// gr4
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate11,
				lblsQP2aTo2d,
				[]string{"qp2d_trust_institutions"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
			)
			gb.MainLabel = trl.S{
				"de": `
				<b>d)</b>&nbsp; Im Allgemeinen kann man Banken und Finanzinstitutionen in Deutschland vertrauen.
				
				`,
				"en": `
					todo
				`,
			}
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
			"en": "Chart",
			"de": "Chart",
		}
		page.Short = trl.S{
			"de": "Chart",
			"en": "Chart",
		}
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

	// Report of results
	{
		p := q.AddPage()
		p.NoNavigation = true
		p.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
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
