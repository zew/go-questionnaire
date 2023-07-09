package kneb1

import (
	"fmt"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
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
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 4

				//
				inp.DD = &qst.DropdownT{}
				if true {
					inp.DD.AddPleaseSelect(cfg.Get().Mp["must_one_option"])
				} else {
					inp.DD.Add(
						"",
						trl.S{
							"de": " Bitte wählen",
							"en": " Please choose",
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
			"en": "Sociodemo-<br>graphics 1",
		}
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
			var inpNames = []string{
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
				inp := gr.AddInput()
				inp.Type = "checkbox"
				inp.Name = "qd5_vocational_" + inpNames[idx]

				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 6

				inp.Label = label
				inp.ControlFirst()
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
					"de": "Welchen Familienstand haben Sie?",
					"en": "What is your marital status?",
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
