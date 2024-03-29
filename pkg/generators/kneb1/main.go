package kneb1

import (
	"fmt"
	"strings"
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

	// version stuff
	q.VersionMax = 4
	q.AssignVersion = "version-from-login-url"
	q.VersionEffective = -2 // prevent 0 - 0 would be a valid version
	if false {
		// version is determined dynamically upon loading base fresh from file in loadQuestionnaire()
		_ = q.Version()
	}

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

			// keep in this position: page-0, gr-0 input-2
			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "panel_type"
				inp.ColSpan = gr.Cols
				inp.ColSpanControl = 0
			}
			{
				inp := gr.AddInput()
				inp.Type = "hidden"
				inp.Name = "panel_id"
				inp.ColSpan = gr.Cols
				inp.ColSpanControl = 0
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
					"de": "Sind Sie… ",
					"en": "What is your gender?",
				}.OutlineHid("D1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Validator = "mustRadioGroup"
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
				inp.Validator = "must"
				inp.Validator = "must; kneb-age-bracket"
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
				yrStart := 1968
				yrStop := 2005

				yrStart = 1950
				yrStop = time.Now().Year() - 10

				for year := yrStart; year <= yrStop; year++ {
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
						"de": "In welchem Bundesland wohnen Sie? ",
						"en": "Which German state you live in?",
					}.OutlineHid("D3.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				for idx, label := range labels {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Validator = "mustRadioGroup"

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

		/*
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
		*/

	}

	// page screenout age
	// disabled for Frau Selz
	if false {
		page := q.AddPage()
		// "Altersgrenzen nicht erfüllt"
		//   => headline inside markdown file
		page.Label = trl.S{
			"de": "",
			"en": "",
		}

		page.SuppressInProgressbar = true

		//  would invalidate NavigationCondition
		// page.NoNavigation = true
		page.NavigationCondition = "kneb_too_old"

		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./must-between-18-and-55.md"
				inp.ColSpan = gr.Cols
				inp.ColSpanControl = 0
				inp.ColSpanLabel = 1
			}

			/*
				{
					inp := gr.AddInput()
					inp.Type = "dyn-textblock"
					inp.DynamicFunc = "knebLinkBackToPanel"
					inp.DynamicFuncParamset = "screenout"
					inp.ColSpan = gr.Cols
					inp.ColSpanControl = 0
					inp.ColSpanLabel = 1
				}
			*/

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since next page is appended below
				inp.Label = trl.S{
					// "de": ` &nbsp;  &nbsp;  &nbsp; Zurück zu Ihrem Panel  &nbsp;  &nbsp;  &nbsp; `,
					"de": ` &nbsp;  &nbsp;  &nbsp; Weiter  &nbsp;  &nbsp;  &nbsp; `,
					"en": `todo`,
				}
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "start"
			}

		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Forward to Panel - screenout",
			"en": "Forward to Panel - screenout",
		}
		page.RedirectFunc = "pageForwardKnebScreenout"
		page.NavigationCondition = "kneb_too_old"
	}

	// page 2a
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Alter, Herkunft, Erfahrungen 2 - Bildung",
			"en": "Age, origin, experience 2 - education",
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
				"none",
				// "inprogress",

				"haupt_noapp",
				"haupt_withapp",
				"real",

				"abitur",
				"uni",
			}
			var labels = []trl.S{
				{
					"de": "Keinen allgemeinen Schulabschluss",
					"en": "todo",
				},
				// {
				// 	"de": "Noch in schulischer Ausbildung",
				// 	"en": "todo",
				// },

				{
					"de": "Hauptschule/ Volksschule <i>ohne</i> abgeschlossene Lehre",
					"en": "todo",
				},
				{
					"de": "Hauptschule/ Volksschule <i>mit</i> abgeschlossener Lehre",
					"en": "todo",
				},
				{
					"de": "Mittel-, Real-, Höhere-, Fach-, Handelsschule ohne Abitur",
					"en": "todo",
				},

				{
					"de": "Abitur / Hochschulreife",
					"en": "todo",
				},
				{
					"de": "Abgeschlossenes Studium (Hochschule oder Universität)",
					"en": "todo",
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": "Was ist Ihr höchster Schulabschluss?",
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
				rad.Validator = "mustRadioGroup"

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

	// page 2b - household type and size
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
				}.OutlineHid("D5.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Validator = "mustRadioGroup"

				rad.Name = "qd5_family_status"
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
			gr.Cols = 3
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
					Wie viele Personen leben insgesamt, d.h. mit Ihnen eingerechnet, in Ihrem Haushalt?
					<br>
					<small>
					Sollten Sie in einer (Studierenden-)WG wohnen, 
					so sollte die Anzahl aller Haushaltsmitglieder auf 1 gesetzt werden
					</small>
					`,
					"en": `
						todo
					`,
				}.OutlineHid("D6.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must"

				inp.Name = "qp6_householdsize"
				inp.Min = 0
				inp.Max = 100
				inp.MaxChars = 6
				inp.Label = trl.S{
					"de": `
						Zahl der Haushaltsmitglieder einschließlich Ihnen selbst, 
						(Ehe-) Partner/in, Kindern und sonstigen Personen
					
					`,
					"en": `todo`,
				}
				inp.Suffix = trl.S{
					"de": `Personen`,
					"en": `people`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 2
				inp.ColSpanControl = 1
			}
		}

	}

	// page 2c - occupational status
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

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
				rad.Validator = "mustRadioGroup"

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
	}

	/*

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

	*/

	// page 2c-condtional-b - occupational status if not employed
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_d7_unemployed"

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
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
							Wenn Sie einmal von den Erwerbstätigkeiten absehen, 
							was von dem Folgenden trifft dann auf Sie zu?
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
				rad.Name = "qd7a_notemployed"
				rad.Validator = "mustRadioGroup"

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

	// page 2c-condtional-b - occupational status if employed
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_d7_employed"

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
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
					"de": "Arbeiter/in",
					"en": "todo",
				},
				{
					"de": "Angestellte/r",
					"en": "todo",
				},
				{
					"de": "Beamtin/Beamter",
					"en": "todo",
				},
				{
					"de": "Selbständig ohne Mitarbeiter/innen",
					"en": "todo",
				},
				{
					"de": "Selbständig mit Mitarbeiter/innen",
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
				rad.Validator = "mustRadioGroup"

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

		/*

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

		*/
	}

	/*
		// page 2e - vocational
		{
			page := q.AddPage()

			page.Label = trl.S{
				"de": "Alter, Herkunft, Erfahrungen 3 - Berufsausbildung",
				"en": "Age, origin, experience 3 - vocational training",
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
					}.OutlineHid("D8.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				//
				for idx, label := range labels {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "qd8_vocational_training"
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
	*/

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
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				labelsRisk(),
				[]string{"qr1_averse_common"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
					<span style="font-size:120%">
					Wie schätzen Sie sich persönlich ein:
					Sind Sie im Allgemeinen ein risikobereiter Mensch
					oder versuchen Sie, Risiken zu vermeiden?
					</span>

					<small>
					Antworten Sie bitte anhand der folgenden Skala,
					wobei der Wert&nbsp;0 bedeutet "gar nicht risikobereit"
					und der Wert&nbsp;10 "sehr risikobereit".
					Mit den Werten dazwischen können Sie Ihre Einschätzung abstufen.
					</small>
				`,
				"en": `
					todo
				`,
			}.OutlineHid("R1.")
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
					"de": `
					An welchen Lebensbereich oder an welche Entscheidungen haben Sie gedacht, 
					als Sie diese Frage beantwortet haben?
						<small>Bitte geben Sie ein oder mehrere Stichwörter an.</small>
					`,
					"en": `todo`,
				}.OutlineHid("R2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Validator = "must"

				inp.Name = "qr2_averse_common_free"
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
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				// labelsRiskFin(),
				labelsRisk(),
				[]string{"qr3_averse_fin"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
					<span style="font-size:120%">
					Bitte schätzen Sie Ihre Bereitschaft ein, ein <i>finanzielles</i> Risiko einzugehen. 
					</span>

					<small>
							Antworten Sie bitte anhand der folgenden Skala, 
							wobei der Wert&nbsp;0 bedeutet "gar nicht risikobereit" 
							und der Wert&nbsp;10 "sehr risikobereit", 
							um potenziell eine höhere Rendite zu erzielen.
					</small>
				`,
				"en": `
					todo
				`,
			}.OutlineHid("R3.")
			gr := page.AddGrid(gb)
			_ = gr
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

		/*
			//
			// gr1
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

			// gr2
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

			// gr3
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
		*/

	}

	/*
		// page 3-split-1
		{
			page := q.AddPage()

			page.Label = trl.S{
				"de": "",
				"en": "",
			}

			page.SuppressInProgressbar = true
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
						"de": `Woran haben Sie gedacht, als Sie die Frage nach dem <i>finanziellen</i> Risiko beantwortet haben?
							<small>Bitte geben Sie ein oder mehrere Stichwörter an.</small>
						`,
						"en": `todo`,
					}.OutlineHid("R4.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
				}
				{
					inp := gr.AddInput()
					inp.Type = "textarea"
					inp.Name = "qr4_averse_fin_free"
					inp.MaxChars = 200
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 0
					inp.ColSpanControl = 1
				}
			}

		}
	*/

	/*
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
						<span style="font-size:120%">
						Bitte schätzen Sie Ihre Bereitschaft ein, ein <i>finanzielles</i> Risiko einzugehen.
						</span>

						<small>
							Antworten Sie bitte anhand der folgenden Skala,
							wobei der Wert&nbsp;0 bedeutet: nicht bereit, ein Risiko einzugehen
							und der Wert&nbsp;10: bereit, ein erhebliches Risiko einzugehen,
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
	*/

	/*
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
	*/

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
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				labelsSelfKnowledge(),
				[]string{"qp1_risk"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
					Wie beurteilen Sie Ihr eigenes Wissen hinsichtlich finanzieller Angelegenheiten?
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

				Ich finde es langweilig, mich mit meinen Finanzen auseinanderzusetzen.

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
				Im Allgemeinen vertraue ich anderen Menschen.

				`,
				"en": `
					todo
				`,
			},
			{
				"de": `
				Im Allgemeinen vertraue ich Banken und Finanzinstitutionen in Deutschland.

				`,
				"en": `
					todo
				`,
			},
		}

		for i := 0; i < len(inputs); i++ {
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				lblsQP2,
				[]string{inputs[i]},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = lbls[i].OutlineHid(fmt.Sprintf("P%d.", i+2)) // P2., P3., P4., P5.
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

	// page
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

		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
		{

			lbls := labelsSelfKnowledgeXX()
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				lbls,
				[]string{"qe1_confidence_before"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
						<span style="font-size:120%">
						Wie viel Vertrauen haben Sie in Ihre Fähigkeit, 
						gute finanzielle Entscheidungen zu treffen?
						</span>


						<small>
						Antworten Sie bitte anhand der folgenden Skala, 
						wobei der Wert&nbsp;0 "kein Vertrauen in die eigene Fähigkeit" bedeutet 
						und der Wert&nbsp;10 "hohes Vertrauen in die eigene Fähigkeit".


						</small>


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

	// page p6
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
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
					
						Stellen Sie sich vor, Sie haben 6000&nbsp;Euro in der Lotterie gewonnen
						und können dieses Geld in ein Unternehmen investieren.

						Die Chance, dass das Unternehmen erfolgreich ist, liegt bei 50 Prozent.

						Im Erfolgsfall verdoppelt sich Ihre Investition nach einem Jahr.

						Bei Misserfolg verlieren Sie die Hälfte der investierten Summe.

						Welchen Anteil der 6000&nbsp;Euro würden Sie in dieses Unternehmen investieren?
					`,
					"en": `todo`,
				}.OutlineHid("P6.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Validator = "must"

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

	}

	if false {
		vorsorgeplanung8x7(&q, 0, 1)
		vorsorgeplanung8x7(&q, 2, 3)
	}
	// vorsorgeplanung8x7(&q, 6, 7)

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
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate5,
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
					{
						"de": "<small>weiß nicht</small>  ",
						"en": "<small>dont know</small>     ",
					},
					{
						"de": "<small>keine Angabe</small>  ",
						"en": "<small>no answer</small>     ",
					},
				},
				[]string{"qfl1_interest"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
					<span style="font-size:110%">
					Im Folgenden haben wir ein Quiz zu finanziellem Wissen. 
					Bitte antworten Sie spontan. 
					</span>

					<br>
					<br>


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

	/*
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

			howSicherPlusGuessed(*qst.WrapPageT(page), "qfl1a_free", "FL1a.")

		}
	*/

	// page 7-2
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate5,
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
					{
						"de": "<small>weiß nicht</small>  ",
						"en": "<small>dont know</small>     ",
					},
					{
						"de": "<small>keine Angabe</small>  ",
						"en": "<small>no answer</small>     ",
					},
				},
				[]string{"qfl2_inflation"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
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

	/*
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

			howSicherPlusGuessed(*qst.WrapPageT(page), "qfl2a_free", "FL2a.")
		}
	*/

	// page 7-4
	{

		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate4,
				labelsStimmeZuOderNicht(),
				[]string{"qfl3_portfoliorisk"},
				radioVals4,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
					Stimmen Sie der folgenden Aussage zu:
					"Die Anlage in Aktien eines einzelnen Unternehmens ist weniger riskant
					als die Anlage in einem Fonds mit Aktien ähnlicher Unternehmen"?
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

	/*
		// page 7-7
		{

			page := q.AddPage()
			page.Label = trl.S{
				"de": "",
				"en": "",
			}
			page.SuppressInProgressbar = true
			page.WidthMax("48rem")

			howSicherPlusGuessed(*qst.WrapPageT(page), "qfl4a_free", "FL4a.")

		}
	*/

	//
	// experiment sequence
	//
	// page experiment introduction 1 neutral
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_t1a"

		page.Label = trl.S{
			"en": "Experiment-Chart-Introduction",
			"de": "Experiment chart-Introduction",
		}
		page.Label = trl.S{
			"en": "",
			"de": "",
		}

		// for next x pages
		page.Short = trl.S{
			"de": "Experiment",
			"en": "Epxperiment",
		}

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
						Im Folgenden geht es um  
						<i>Anbauentscheidungen in der Waldwirtschaft</i>.
						<br>
						<br>

						Es gibt verschiedene Möglichkeiten, ein Waldstück zu bepflanzen. 						
						Ein Beispiel: 
						Eine Person besitzt einen Waldabschnitt und muss entscheiden, 
						welche <i>Baumarten</i> sie pflanzen möchte. Sie kann sich 
						<i>zwischen zwei Arten</i>
						entscheiden. 
						
						
					</p>

					<ul>
						<li>
						<b>Baumart 1:</b> Diese Art wächst <i>langsamer</i> 
							und erzielt im Durchschnitt <i>geringere Erträge</i>. 
							Gleichzeitig ist sie <i>widerstandsfähiger</i> gegen Schädlinge. 
							Das bedeutet, dass die Erträge <i>kaum schwanken</i>.
							<br>
							<br>
						</li>
						<li>
							<b>Baumart 2:</b> Diese Art wächst <i>schneller</i> 
							und erzielt im Durchschnitt höhere Erträge. 
							Gleichzeitig ist sie <i>anfälliger</i> für Schädlinge. 
							Das bedeutet, dass die Erträge <i>stärker schwanken</i>.
							<br>
							<br>
						</li>			
					</ul>

					<p>
						In einer interaktiven Graphik versuchen wir, 
						die <i>Abwägung zwischen Ertrag und Widerstandsfähigkeit</i> 
						beider Baumarten zu verdeutlichen. 
					
						Der Preis für beide Baumarten ist in unserem Beispiel gleich.
					
					</p>
					`,
					"en": `
						todo
					`,
				}
			}
		}

	}

	// page experiment introduction 1 finance
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_t1b"

		page.Label = trl.S{
			"en": "Experiment-Chart-Introduction",
			"de": "Experiment chart-Introduction",
		}
		page.Label = trl.S{
			"en": "",
			"de": "",
		}
		// for next x pages
		page.Short = trl.S{
			"de": "Experiment",
			"en": "Epxperiment",
		}

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
						Im Folgenden geht es um
						<i>Investitionsentscheidungen im Rahmen der Altersvorsorge</i>.
					</p>
						
					<p>
						Es gibt verschiedene Möglichkeiten für das Alter vorzusorgen.
						Eine Möglichkeit ist
						<i>einen monatlichen Geldbetrag</i> über einen <i>langen Zeithorizont</i> zu sparen.
					</p>

					<p>In diesem Beispiel kann eine Person zwischen zwei Anlageformen wählen.</p>

					<ul>							
						<li>
							<i>Möglichkeit 1:</i> Sparen mit <b>Sparbuch</b>.
							<br>
							Ein Sparbuch ist eine Anlage mit <i>geringem Risiko</i> 
							und erzielt <i>konstante Erträge</i>. 
							Das <i>angesparte Vermögen schwankt nicht</i>, 
							ist im Durchschnitt jedoch <i>geringer</i>, da die Erträge gering sind.
							<br>
							<br>
						</li>
						
						<li>
							<i>Möglichkeit 2</i>: Sparen mit <b>Aktien</b>. 
							<br>
							Aktien sind Wertpapiere, mit denen Aktionäre Anteile an Unternehmen erwerben.
							<br>
							Aktien haben ein <i>höheres Risiko</i> und erzielen 
							<i>im Durchschnitt</i> über einen längeren Zeitraum 
							<i>höhere Erträge</i>. 
							Gleichzeitig können die Renditen und
							 der Gesamtwert des angesparten Vermögens <i>schwanken</i>.

							<br>
							<br>
						</li>
					</ul>

					<p>					
						In einer interaktiven Graphik versuchen wir, 
						die Abwägung zwischen <i>Chancen und Risiken</i> einer Anlage 
						am Kapitalmarkt zu verdeutlichen.
					</p>
					`,
					"en": `
					todo
				`,
				}
			}
		}

	}

	// page experiment guided tour dyn
	// {
	// 	page := q.AddPage()
	// 	page.GeneratorFuncName = "kneb202306guidedtour"
	// }

	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN0"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN1"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN2"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN3"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN4"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN5"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN6"
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306guidedtourN7"
	}

	// page experiment chart 0
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306simtool0"
	}

	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306simtool1"
	}
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")
		howSicher1to10(*qst.WrapPageT(page), "qc24_how_sicher", "qc24hs.")
	}

	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306simtool2"
	}
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")
		howSicher1to10(*qst.WrapPageT(page), "qc25_how_sicher", "qc25hs.")
	}
	{
		page := q.AddPage()
		page.GeneratorFuncName = "kneb202306simtool3"
	}
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")
		howSicher1to10(*qst.WrapPageT(page), "qc26_how_sicher", "qc26hs.")
	}

	// page chart 1
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_t1b"
		page.GeneratorFuncName = "kneb202306simtool4"
	}

	// page experiment +2
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Feedbackfragen zum Tool - 2",
			"en": "todo",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true

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
				"qf6_comprehensive",
				"qf7_helplful",
				"qf8_complex",
				"qf9_feelinginformed",
			}
			lblsQ := []trl.S{
				{
					"de": `Die graphische Darstellung in der interaktiven Graphik ist verständlich.`,
					"en": `todo`,
				},
				{
					"de": `Die graphische Darstellung in der interaktiven Graphik ist hilfreich.`,
					"en": `todo`,
				},
				{
					"de": `Die graphische Darstellung in der interaktiven Graphik ist kompliziert.`,
					"en": `todo`,
				},
				{
					"de": `Ich fühle mich durch die gezeigten Fälle in der interaktiven Graphik gut informiert.`,
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
				// distinct for last
				// if i == len(inputs)-1 {
				// 	lbls[0] = trl.S{
				// 		"de": "<small>ganz und gar uninformiert</small>     <div>0</div>",
				// 		"en": "<small>todo</small>     <div>0</div>",
				// 	}
				// 	lbls[10] = trl.S{
				// 		"de": "<small>ganz und gar informiert</small>       <div>10</div>",
				// 		"en": "<small>todo</small> <div>10</div>",
				// 	}
				// }

				gb := qst.NewGridBuilderRadiosWithValidator(
					columnTemplate11,
					lbls,
					[]string{inputs[i]},
					radioVals11,
					[]trl.S{{"de": ``, "en": ``}},
					"mustRadioGroup",
				)
				// gb.MainLabel = lblsQ[i].OutlineHid(fmt.Sprintf("%c)", rune(97+i)))
				gb.MainLabel = lblsQ[i].OutlineHid(fmt.Sprintf("F%v.", i+6)) // .Outline("F6., F7., ...  F9.")
				gr := page.AddGrid(gb)
				_ = gr
				// gr.BottomVSpacers = 2
				gr.Style = css.NewStylesResponsive(gr.Style)
				gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			}

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
					"de": `Haben Sie weitere Anmerkungen zu der Darstellung in der interaktiven Graphik?`,
					"en": `todo`,
				}.OutlineHid("F10.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qf10_tool_free"
				inp.MaxChars = 200
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

	}

	// page experiment +3a - conditional
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_t2a"

		page.Label = trl.S{
			"de": "Treatment: Giving advice - 1",
			"en": "todo",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}

		// page.SuppressInProgressbar = true
		page.Short = trl.S{
			"de": "Treatment:<br>Giving advice - 1",
			"en": "todo",
		}
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
					Welche Tipps würden Sie einem Freund oder einer Freundin geben, 
					wenn es um das Sparen für die Altersvorsorge geht? 
					Bitte beschreiben Sie Ihre Hinweise in einigen kurzen Sätzen oder Stichpunkten.

					`,
					"en": `todo`,
				}.OutlineHid("T1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Validator = "must"

				inp.Name = "qt1_advice"
				inp.MaxChars = 400
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

	}

	// page experiment +3b - conditional
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_t2a"

		page.Label = trl.S{
			"de": "Treatment: Giving advice - 2",
			"en": "todo",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}

		// page.SuppressInProgressbar = true
		page.Short = trl.S{
			"de": "Treatment:<br>Giving advice - 2",
			"en": "todo",
		}
		page.WidthMax("48rem")

		inputs := []string{
			"qt1b_a",
			"qt1b_b",
			"qt1b_c",
		}
		lblsQ := []trl.S{
			{
				"de": `
					Würden  Sie einem Freund oder einer Freundin empfehlen, am Kapitalmarkt zu investieren?
				`,
				"en": `todo`,
			},
			{
				"de": `
					Bitte denken Sie an einen Freund oder eine Freundin, die gerade ins Berufsleben startet.
					<br>
					Würden Sie ihm oder ihr empfehlen, sofort mit dem Sparen für das Alter anzufangen?
				
				`,
				"en": `todo`,
			},
			{
				"de": `
					Bitte denken Sie an einen Freund oder eine Freundin, die noch keine Erfahrung beim Thema Rente und Finanzen haben.
					<br>	
					Würden Sie einem Freund oder einer Freundin empfehlen, sich professionell zum Thema Rente und Finanzen beraten zu lassen?		
				`,
				"en": `todo`,
			},
		}

		for i := 0; i < len(inputs); i++ {

			lbls := labelsSelfKnowledge()
			lbls[0] = trl.S{
				"de": "<small>Auf keinen Fall empfehlen</small>     <div>0</div>",
				"en": "<small>todo</small>     <div>0</div>",
			}
			lbls[10] = trl.S{
				"de": "<small>Auf jeden Fall empfehlen</small>       <div>10</div>",
				"en": "<small>todo</small> <div>10</div>",
			}

			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				lbls,
				[]string{inputs[i]},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = lblsQ[i].OutlineHid(fmt.Sprintf("T1a%v.", i))
			gr := page.AddGrid(gb)
			_ = gr
			// gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page experiment +4
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Selbstvertrauen nach dem Tool",
			"en": "Confidence afterwards",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.Short = trl.S{
			"de": "Selbstvertrauen<br>nachher",
			"en": "Confidence<br>after",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
		{

			lbls := labelsSelfKnowledgeXX()
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				lbls,
				[]string{"qe2_confidence_after"},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
						<span style='font-size:120%'>
						Wie viel Vertrauen haben Sie in Ihre Fähigkeit, 
						gute finanzielle Entscheidungen zu treffen?
						</span>


						<small>
						Antworten Sie bitte anhand der folgenden Skala,
						wobei der Wert&nbsp;0 "Kein Vertrauen in die eigene Fähigkeit" bedeutet
						und der Wert&nbsp;10 "Hohes Vertrauen in die eigene Fähigkeit".
						</small>
					`,
				"en": `
						todo
					`,
			}.OutlineHid("E2.")
			gr := page.AddGrid(gb)
			gr.BottomVSpacers = 3
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Motivation nach Treatment",
			"en": "Motivation after treatment",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "knebSlightlyDistinctLabel"
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

		}

		inputs := []string{
			"q20_motivation",
		}

		for i := 0; i < len(inputs); i++ {

			lbls := labelsSelfKnowledge()
			lbls[0] = trl.S{
				"de": "<small>gar nicht motiviert</small>     <div>0</div>",
				"en": "<small>todo</small>     <div>0</div>",
			}
			lbls[10] = trl.S{
				"de": "<small>sehr stark motiviert</small>       <div>10</div>",
				"en": "<small>todo</small> <div>10</div>",
			}

			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate11,
				lbls,
				[]string{inputs[i]},
				radioVals11,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			// gb.MainLabel = lblsQ[i].OutlineHid(fmt.Sprintf("F%v.", i+6)) // .Outline("F6., F7., ...  F9.")
			gr := page.AddGrid(gb)
			_ = gr
			// gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page FL4
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
			gr.Cols = 6
			gr.BottomVSpacers = 3
			{
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"de": `
							<span style="font-size:110%">
							Im Folgenden haben wir weitere Quizfragen. 
							Bitte antworten Sie spontan. 
							</span>
	
							<br>
							<br>	
						

							Was ist die Hauptfunktion des Aktienmarktes?
						`,
						"en": `todo`,
					}.OutlineHid("FL4.")
					inp.ColSpan = gr.Cols
					inp.ColSpanLabel = 1
					inp.ColSpanControl = 0
				}

				for idx, label := range labelsStockMarketPurpose() {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Validator = "mustRadioGroup"

					rad.Name = "qfl4_stockmarketpurpose"
					rad.ValueRadio = fmt.Sprintf("%d", idx+1)

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}

	}

	// page FL5
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
					rad.Validator = "mustRadioGroup"

					rad.Name = "qfl5_assetclassvola"
					rad.ValueRadio = fmt.Sprintf("%d", idx+1)

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}

	}

	// page FL6
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

							Gleichzeitig besteht eine 10-prozentige Chance,
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
					rad.Validator = "mustRadioGroup"

					rad.Name = "qfl6_expectedvalue"
					rad.ValueRadio = fmt.Sprintf("%d", idx+1)

					rad.ColSpan = gr.Cols / 2
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label
					rad.ControlFirst()
				}
			}
		}
	}

	// FL confidence - new 2024-02
	{
		page := q.AddPage()

		page.Label = trl.S{
			"de": "Vertrauen in Fin Quiz",
			"en": "todo",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true

		page.WidthMax("48rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Im Laufe der Befragung haben wir Ihnen sechs Quizfragen zu finanziellem Wissen gestellt.
						<br>
						Wie viele Fragen haben Sie Ihrer Meinung nach richtig beantwortet?									
					`,
					"en": `todo`,
				}
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			inputs := []string{
				"qfl_confidence_",
			}
			lblsQ := []trl.S{
				{
					"de": `&nbsp;`,
					"en": `&nbsp;`,
				},
			}

			for i := 0; i < len(inputs); i++ {
				lbls := []trl.S{
					{
						"de": "keine",
						"en": "keine",
					},
					// 1-6
					{
						"de": " 1",
						"en": " 1",
					},
					{
						"de": " 2",
						"en": " 2",
					},
					{
						"de": " 3",
						"en": " 3",
					},
					{
						"de": " 4",
						"en": " 4",
					},
					{
						"de": " 5",
						"en": " 5",
					},
					{
						"de": " 6",
						"en": " 6",
					},

					{
						"de": "weiß<br>nicht",
						"en": "weiß<br>nicht",
					},
					{
						"de": "keine<br>Angabe",
						"en": "keine<br>Angabe",
					},
				}

				gb := qst.NewGridBuilderRadiosWithValidator(
					[]float32{

						0.2, 1.2,

						0, 1,
						0, 1,
						0, 1,
						0, 1,
						0, 1,
						0, 1,

						0.2, 1.2,
						0.0, 1.2,
					},
					lbls,
					[]string{inputs[i]},
					[]string{
						"keine", "1", "2", "3", "4", "5", "6",
						"dontknow", "noanswer",
					},
					[]trl.S{{"de": ``, "en": ``}},
					"mustRadioGroup",
				)
				gb.MainLabel = lblsQ[i].OutlineHid("FLConf")
				gr := page.AddGrid(gb)
				_ = gr
				// gr.BottomVSpacers = 2
			}

		}

	}

	// page experiment +5
	erfahrungMitFinanzenSplit1(&q, 0, 0)
	erfahrungMitFinanzenSplit2(&q, 0, 0)

	// page B1, B2, B3
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
				rad.Validator = "mustRadioGroup"

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
			gr := page.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Wie oft sprechen Sie mit Ihrem Finanzberater oder Ihrer Finanzberaterin?
					`,
					"en": `
						todo
					`,
				}.OutlineHid("B2.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, lbl := range labelsFrequency2() {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Validator = "mustRadioGroup"

				rad.Name = "qb2_frequency"
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

		//
		//
		// gr2
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate4,
				labelsFrequency(),
				[]string{"qb3_frequency"},
				radioVals5,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
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
	// page B4, B5
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

		// B5. was yesNo()
		// but got a third option 2024-01
		{

			lbl := trl.S{
				"de": `
						Würden Sie es vorziehen, 
						wenn eine andere Person finanzielle Entscheidungen für Sie trifft?
					`,
				"en": `todo`,
			}

			radioValues := []string{
				"yes",
				"no",
				"notapplicable",
			}

			labels := []trl.S{
				{
					"de": "&nbsp;&nbsp;ja",
					"en": "yes",
				},
				{
					"de": "nein",
					"en": "no",
				},
				{
					"de": "trifft nicht zu, ich treffe keine finanziellen Entscheidungen",
					"en": "no",
				},
			}

			gr := page.AddGroup()
			gr.Cols = 8 + 1
			gr.Cols = 6 + 1
			gr.WidthMax("40rem")
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = lbl.OutlineHid("B5.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}

			{
				// vertical spacer
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "", "en": ""}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Validator = "mustRadioGroup"
				rad.Name = "qb5_delegate"
				rad.ValueRadio = radioValues[idx]

				rad.ColSpan = gr.Cols
				rad.ColSpan = 2
				// if idx > 0 {
				// 	rad.ColSpan = gr.Cols - 2 - 1
				// }

				rad.ColSpanLabel = 1
				rad.ColSpanControl = 6

				rad.Label = label
				rad.ControlFirst()
			}

		}

		// B6 is now on separate, conditional page

	}

	//
	// page B6
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_b6_who_competent"

		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		//
		//
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate4,
				labelsWhoIsCompetent(),
				[]string{"qb6_whocompetent"},
				radioVals4,
				[]trl.S{{"de": ``, "en": ``}},
				"mustRadioGroup",
			)
			gb.MainLabel = trl.S{
				"de": `
					Wer könnte finanzielle Entscheidungen am besten für Sie treffen?
					<br>
					<br>
				`,
				"en": `
					todo
				`,
			}.OutlineHid("B6.")
			gr := page.AddGrid(gb)
			gr.WidthMax("36rem")
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
		}

	}

	// page H1
	{
		page := q.AddPage()
		page.NavigationCondition = "kneb_h1_who_responsibe"

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

	// page H2
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
			gr.BottomVSpacers = 1
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": `
					<span style="font-size:110%">
					Wir stellen Ihnen nun eine Frage zum <i>Finanzvermögen</i>: <br>
					Haben Sie (d.h. Ihr Haushalt) im Dezember   
						<span style="font-size 120%; color:#e22;" >2023</span> 
					eine der folgenden Vermögensarten besessen?
					</span>
				

					<small>
					Falls Sie nicht wissen, ob Ihr Partner diese Vermögensarten besitzt, 
					beantworten Sie die Fragen bitte <i>nur</i> für sich selbst.
					</small>
				`,
				"en": `todo`,
			}.OutlineHid("H2.")
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
		}

		// gr0
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
				<span style="font-size: 115%">Sparanlagen</span>
				<br>
				<small>
					(z.B. Sparbücher, Festgeldkonten,
					Tagesgeldkonten oder
					Sparverträge)
				</small>
				`,
				"en": `todo`,
			},
			"qh2_savingsaccount",
			"a)",
			true,
		)

		// gr1
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
				<span style="font-size: 115%">Bausparverträge</span>
				<br>
				<small>
					(die noch nicht in Darlehen
					umgewandelt wurden)
				</small>

				`,
				"en": `todo`,
			},
			"qh2_bausparen",
			"b)",
			true,
		)

		// gr2
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
				<span style="font-size: 115%">Festverzinsliche Wertpapiere</span>
				<small>
					(z. B. Spar- oder Pfandbriefe, Bundesschatzbriefe, Industrieanleihen oder Anteile an Rentenfonds)
				</small>`,
				"en": `todo`,
			},
			"qh2_bonds",
			"c)",
			true,
		)

		// gr3
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
				<span style="font-size: 115%">Aktien oder Aktienfonds und Immobilienfonds</span>
				<small>
					(auch Aktienanleihen, börsennotierte Fonds, ETFs, Mischfonds oder ähnliche Anlagen)
				</small>
				`,
				"en": `todo`,
			},
			"qh2_stocks_etf",
			"d)",
			true,
		)

		// gr4
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
				<span style="font-size: 115%">Sonstige Wertpapiere</span>
				<small>
					(z.B. Discountzertifikate, Hedgefonds, Filmfonds, Windenergiefonds und andere Finanzinnovationen)
				</small>
				`,
				"en": `todo`,
			},
			"qh2_other",
			"e)",
			true,
		)

		// gr5
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `<span style="font-size:115%">Gold</span>`,
				"en": `todo`,
			},
			"qh2_gold",
			"f)",
			true,
		)

		// gr6
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `<span style="font-size:115%">Kryptowährungen</span>`,
				"en": `todo`,
			},
			"qh2_crypto",
			"f)",
			true,
		)

	}

	// page H3
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
						Wie hoch <i style="font-size:110%">schätzen</i> Sie,
						ist das <i style="font-size:110%">monatlich</i> verfügbare 
						<i style="font-size:110%">Nettoeinkommen Ihres Haushalts</i>,
						also dasjenige Geld, das dem gesamten Haushalt nach Abzug
						von Steuern und Sozialversicherungsbeiträgen zur Deckung der Ausgaben
						 zur Verfügung steht?
					`,
					"en": `
						todo
					`,
				}.OutlineHid("H3.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0

			}

			ranges := []int{
				0,
				500,
				750,

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

			//
			insertDot := func(s string) string {

				s = strings.TrimSpace(s)

				ln := len(s)
				if ln > 3 {
					s = s[0:ln-3] + "." + s[ln-3:]
				}

				return s

			}

			for i := 0; i < len(ranges); i++ {

				rLow := ranges[i]
				rHig := -1
				if i < len(ranges)-1 {
					rHig = ranges[i+1]
				}

				rsLow := fmt.Sprintf("%10.0f", float64(rLow))
				rsHig := fmt.Sprintf("%10.0f", float64(rHig))
				// rsLow = strings.ReplaceAll(rsLow, ",", ".")
				// rsHig = strings.ReplaceAll(rsHig, ",", ".")

				rsLow = insertDot(rsLow)
				rsHig = insertDot(rsHig)

				opt := fmt.Sprintf("upto%d", rHig)
				lbl := trl.S{
					"de": fmt.Sprintf("%s€ bis unter %s€ monatlich", rsLow, rsHig),
					"en": fmt.Sprintf("%s€ to under  %s€ per month", rsLow, rsHig),
				}
				if i == 0 {
					lbl = trl.S{
						"de": fmt.Sprintf("unter %s€ monatlich", rsHig),
						"en": fmt.Sprintf("under %s€ per month", rsHig),
					}
				}
				if rHig == -1 {
					opt = "over15000"
					lbl = trl.S{
						"de": fmt.Sprintf("%s€ und mehr", rsLow),
						"en": fmt.Sprintf("%s€ and more", rsLow),
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
				rad.Validator = "mustRadioGroup"

				rad.Name = "qh3_income"
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
				inp.Validator = "must"

				inp.Name = "qh3_income"
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

	// page H4
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
						<span style="font-size:110%%">
						Nun sind wir an Ihrem Altersvorsorgevermögen interessiert:<br>
						
						Haben Sie (d.h. Ihr Haushalt) im Dezember <span style="font-size 120%%; color:#e22;">2023</span> 
						einen der folgenden privaten oder betrieblichen Altersvorsorgeverträge besessen?
					    </span>

						<small>
						Falls Sie nicht wissen, ob Ihr Partner diese Vermögensarten besitzt,
						beantworten Sie die Fragen bitte <i>nur</i> für sich selbst.
						</small>

						<!-- %v -->

					`,
						// december previous year - for 2023: 2022
						time.Now().Year()-1,
					),
					"en": `todo`,
				}.OutlineHid("H4.")
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
					<span style="font-size:115%">Private Lebensversicherungen</span>
					<small>
						(z.B. klassische und fondsgebundene Kapitallebensversicherungen,
						<i>nicht</i> reine Risikolebensversicherungen
						oder Direktversicherungen über den Arbeitgeber)
					</small>

				`,
				"en": `todo`,
			},
			"qh4a_lifeinsurance",
			"a)",
			true,
		)

		// gr 2
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					<span style="font-size:115%">Betriebliche Lebensversicherungen</span>
					<br>
					<small>
					(z. B. Direktversicherungen)
					</small>

				`,
				"en": `todo`,
			},
			"qh4b_directinsurance",
			"b)",
			true,
		)

		// gr 3
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					<span style="font-size:115%">Sonstige betriebliche Altersvorsorge</span>
					<small>
					(z. B. Betriebsrenten aus Pensions- oder Unterstützungskassen und betriebliche Direktzusagen sowie Zusatzversorgung im öffentlichen Dienst; auch aus früheren Beschäftigungsverhältnissen)
					</small>

				`,
				"en": `todo`,
			},
			"qh4c_otherpensions",
			"c)",
			true,
		)

		// gr 4
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					<span style="font-size:115%">Staatlich geförderte private Altersvorsorge ("Riester-Rente")</span>
					<small>
					(staatlich geförderte und zertifizierte Sparanlagen, auch "Rürup-" bzw. Basisrenten)
					</small>

				`,
				"en": `todo`,
			},
			"qh4d_otherpensions",
			"d)",
			true,
		)

		// gr 5
		yesNo(
			*qst.WrapPageT(page),
			trl.S{
				"de": `
					<span style="font-size:115%">Private Rentenversicherungen</span>
					<small>
					(z.B. private Rentenversicherungsverträge, die <u><i>nicht</i></u> staatlich gefördert werden bzw. abgeschlossen wurden, 
						bevor es solche Fördermöglichkeiten gab)
					</small>
				`,
				"en": `todo`,
			},
			"qh4e_privatepensions",
			"e)",
			true,
		)

	}

	// page finish 1
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Abschluss",
			"en": "todo",
		}
		page.Label = trl.S{
			"de": "",
			"en": "",
		}
		page.SuppressInProgressbar = true
		page.WidthMax("48rem")

		// gr0
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
					`,
					"en": `
						todo
					`,
				}
				inp.ColSpanLabel = 1
			}

			/*
				{
					inp := gr.AddInput()
					inp.Type = "dyn-textblock"
					inp.DynamicFunc = "ResponseStatistics"
					inp.ColSpan = 1
					inp.ColSpanControl = 0
					inp.ColSpanLabel = 1
				}
			*/

		}

		// gr1
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			// single column
			gr.Cols = 3
			var radioValues = []string{
				"weber",
				"zewexpertise",
				"verbraucherzentral",
				"nothanks",
			}

			// knebDownloadURL
			var labels = []trl.S{
				{
					"de": `Das E-Book "Genial einfach investieren" von Prof. Martin Weber `,
					"en": `todo`,
				},
				{
					"de": `Die ZEW-Kurzexpertise "Wie haben sich Coronakrise und Preissteigerungen auf die Altersvorsorge ausgewirkt?"`,
					"en": `todo`,
				},
				{
					"de": `Link zur Verbraucherzentrale  mit einer Übersicht zum Thema "Alles zur Geldanlage: Das müssen Sie dazu wissen"`,
					"en": `todo`,
				},
				{
					"de": `Nein, danke`,
					"en": `todo`,
				},
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"de": `
						Als Dankeschön bieten wir Ihnen die Möglichkeit, 
						eine dieser drei Quellen herunterzuladen.

						<small>
						(Download Link auf der nächsten Seite)
						</small>
					`,
					"en": `todo`,
				}.OutlineHid("Z1.")
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 0
			}
			for idx, label := range labels {
				rad := gr.AddInput()
				rad.Type = "radio"
				rad.Name = "qz1_download"
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
						Gibt es noch irgendetwas, was Sie uns zu diesem Fragebogen 
						oder Thema mitteilen möchten?
					`,
					"en": `
						todo
					`,
				}.OutlineHid("Z2.")
				inp.ColSpanLabel = 1
				inp.ColSpan = gr.Cols
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "qz2_comment"
				inp.MaxChars = 300
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}

		}

		//
		//
		// advance to last page "data saved"
		// {
		// 	gr := page.AddGroup()
		// 	gr.Style = css.NewStylesResponsive(gr.Style)
		// 	gr.Cols = 2
		// 	gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
		// 	// gr.Width = 80

		// 	{
		// 		inp := gr.AddInput()
		// 		inp.Type = "textblock"
		// 		inp.Label = trl.S{"de": "", "en": ""}
		// 		// inp.Label = trl.S{
		// 		// 	"de": "Durch Klicken erhalten Sie eine Zusammenfassung Ihrer Antworten",
		// 		// 	"en": "By clicking, you will receive a summary of your answers.",
		// 		// }
		// 		inp.ColSpan = 1
		// 		inp.ColSpanLabel = 1
		// 	}
		// 	{
		// 		inp := gr.AddInput()
		// 		inp.Type = "button"
		// 		inp.Name = "submitBtn"
		// 		// inp.Name = "finished"
		// 		inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since next page is appended below

		// 		// two more pages
		// 		inp.Label = cfg.Get().Mp["end"]
		// 		inp.Label = cfg.Get().Mp["finish_questionnaire"]
		// 		inp.ColSpan = 1
		// 		inp.ColSpanControl = 1
		// 		inp.AccessKey = "n"
		// 		inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
		// 		inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
		// 		// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
		// 	}
		// }

	}

	// page finish 2
	// report of results
	{
		page := q.AddPage()
		// page.NoNavigation = true
		page.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}

		// page.WidthMax("calc(100% - 1.2rem)")
		page.WidthMax("40rem")

		{
			gr := page.AddGroup()
			gr.Cols = 1
			/*
				{
					inp := gr.AddInput()
					inp.Type = "dyn-textblock"
					inp.DynamicFunc = "ResponseStatistics"
					inp.ColSpan = 1
					inp.ColSpanControl = 0
					inp.ColSpanLabel = 1
				}
			*/

			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "knebsDownloadURL"
				inp.ColSpan = gr.Cols
				inp.ColSpanLabel = 1
			}

			//
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "PermaLink"
				inp.DynamicFuncParamset = "hidden"
				inp.ColSpan = gr.Cols
				inp.ColSpanControl = 0
				inp.ColSpanLabel = 1
			}

			//
			/*
				{
					inp := gr.AddInput()
					inp.Type = "dyn-textblock"
					inp.DynamicFunc = "knebLinkBackToPanel"
					inp.DynamicFuncParamset = "success"
					inp.ColSpan = gr.Cols
					inp.ColSpanControl = 0
					inp.ColSpanLabel = 1
				}
			*/

			/*
				{
					inp := gr.AddInput()
					inp.Type = "dyn-textblock"
					inp.DynamicFunc = "LinkBack"
					inp.ColSpan = gr.Cols
					inp.ColSpanControl = 0
					inp.ColSpanLabel = 1
				}
			*/

			/*
				{
					inp := gr.AddInput()
					inp.Type = "dyn-textblock"
					inp.DynamicFunc = "RenderStaticContent"
					inp.DynamicFuncParamset = "./site-imprint.md"
					inp.ColSpan = 1
					inp.ColSpanLabel = 1
				}
			*/

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since next page is appended below
				inp.Label = trl.S{
					// "de": ` &nbsp;  &nbsp;  &nbsp; Zurück zu Ihrem Panel  &nbsp;  &nbsp;  &nbsp; `,
					"de": ` &nbsp;  &nbsp;  &nbsp; Weiter  &nbsp;  &nbsp;  &nbsp; `,
					"en": `todo`,
				}
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "center"
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "start"
			}

		}

	}

	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Forward to Panel - complete",
			"en": "Forward to Panel - complete",
		}
		page.RedirectFunc = "pageForwardKnebComplete"
		page.NoNavigation = true
	}

	// q.AddFinishButtonNextToLast()

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
