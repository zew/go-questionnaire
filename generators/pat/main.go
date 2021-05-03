package pat

import (
	"fmt"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var radioVals7 = []string{"1", "2", "3", "4", "5", "6", "7"}
var columnTemplate7 = []float32{
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
}

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat")
	q.Survey.Params = params
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}

	q.VersionMax = 16
	q.AssignVersion = "round-robin"
	q.VersionEffective = -2 // must be re-set at the end - after validate

	// page 0
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.NoNavigation = true
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		//
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				HERZLICH WILLKOMMEN UND VIELEN DANK FÜR IHRE TEILNAHME!<br><br>

				<p>Dies ist eine Studie des Zentrums für Europäische Wirtschaftsforschung (ZEW) in Mannheim 
				sowie der Universitäten in Köln, Mannheim, Münster und Zürich. 
				Ihre Teilnahme wird nur wenige Minuten in Anspruch nehmen 
				und Sie unterstützen damit die Forschung zu Entscheidungsprozessen in der Politik.
				</p>

				<p>In dieser Studie treffen Sie acht Entscheidungen und beantworten sieben Fragen. 
				Nach der Erhebung werden 10&nbsp;% aller Teilnehmer zufällig ausgewählt. 
				Von jedem ausgewählten Teilnehmer wird eine der acht Entscheidungen zufällig bestimmt 
				und genau wie im Folgenden beschrieben umgesetzt 
				(alle erwähnten Personen existieren wirklich und alle Auszahlungen werden wie beschrieben getätigt).
				</p>

				<p>
				In dieser Umfrage gibt es keine richtigen oder falschen Antworten. 
				Bitte entscheiden Sie daher immer gemäß Ihrer persönlichen Ansichten. 
				Ihre Antworten werden dabei streng vertraulich behandelt.
				</p>

				<br>
				<br>
				`,
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.ColSpanControl = 1
			inp.DynamicFunc = "PatLogos"
		}
		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "button"
			inp.Name = "submitBtn"
			inp.Response = "1"
			inp.Label = trl.S{"de": "Weiter"}
			inp.StyleCtl = css.ItemEndMA(inp.StyleCtl)
			inp.AccessKey = "n"

		}

	}

	// erster Teil

	// page 1
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Stiftungen 1"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 6
			gr.BottomVSpacers = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 6
				inp.ColSpanLabel = 6
				inp.Desc = trl.S{"de": `
				<p><b>
				Im Folgenden geht es um eine Spende von 30&nbsp;€, die <i>eine</i> dieser drei Stiftungen erhalten soll:
				</b></p>

				<br>

				<!-- beware of hyphenization of css class names  -->
				<style>
					table.xxx01 td {
						text-align: center;
						text-align: center;
					}
				</style>

				<table class="xxx01">
				<tr>
					<td style="width: 33.2%">Politisch links</td>
					<td style="width: 33.2%">Politische Mitte</td>
					<td style="width: 33.2%">Politisch konservativ</td>
				<tr>
				<tr>
					<td style="vertical-align: top;" ><b>Hans-Böckler-Stiftung</b></td>
					<td style="vertical-align: top;" ><b>Bund der Steuerzahler Deutschland e.V.</b></td>
					<td style="vertical-align: top;" ><b>Ludwig-Erhard-Stiftung</b></td>
				<tr>
				</table>

				<div class='vspacer-08'> &nbsp; </div>

				<p>
					Fünf deutsche Staatsangehörige haben an einer Vorstudie teilgenommen. 
					
					Jede dieser fünf Personen hat in der Vorstudie angegeben, welche der drei Stiftungen sie als am besten, mittel und am schlechtesten erachtet.
				</p>
				<p>
					Wir sind nun daran interessiert, wie Sie die fünf individuellen Präferenzen in eine Gruppenentscheidung zusammenfassen, an welche Stiftung die 30&nbsp;€ gehen sollen. Bevorzugen Sie beispielsweise eher eine Kompromisslösung oder eher eine Mehrheitslösung? Ihre eigene Meinung über die Stiftungen soll dabei keine Rolle spielen. Deshalb sind die Stiftungen im Folgenden als Stiftung A, B und C anonymisiert.
				</p>
				<p>
					Sie werden insgesamt sechs Entscheidungen treffen, wie die Präferenzen der Gruppe zusammengefasst werden sollen. Eine der sechs Entscheidungen stellt die echten Präferenzen der Gruppenmitglieder aus der Vorstudie dar und kann daher zufällig ausgewählt und tatsächlich umgesetzt werden. 
					
					Da Sie nicht wissen, welche Entscheidung die echten Präferenzen darstellt, nehmen Sie bitte in allen Fällen an, dass die jeweilige Entscheidung tatsächlich umgesetzt wird.
				</p>
				
				`}
			}
		}
	}

	// page 2
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Stiftungen 2"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		page.ValidationFuncName = "patPage2"
		page.ValidationFuncMsg = trl.S{
			"de": "Wollen Sie wirklich weitergehen oder wollen Sie Ihre bisherigen Antworten vervollständigen?",
			// "en": "Does not add up. Really continue?",
		}

		// gr-1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 1 - 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{
					"de": `Um zu entscheiden, an welche Stiftung das Geld gehen soll, 
					<b>setzen Sie bitte bei der entsprechenden Stiftung <i>ein</i> Kreuz 
					in der Spalte „Auswahl“.</b>
					 Falls Sie eine zweite oder dritte Alternative als genauso gut empfinden, setzen Sie ein Kreuz in der Spalte „Gleich gut“.`,
				}
			}
		}

		// loop over matrix questions
		for i := 0; i < 3; i++ {

			// if i == 0 {
			// 	// explanation after first question
			// 	gr := page.AddGroup()
			// 	gr.Cols = 1
			// 	gr.BottomVSpacers = 3
			// 	gr.RandomizationGroup = 1 - 1
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-composite"
			// 	inp.ColSpanControl = 1
			// 	inp.DynamicFunc = "PoliticalFoundationsPretext__0__0"
			// }

			{
				gr := page.AddGroup()
				gr.Cols = 1
				if i == 0 {
					gr.BottomVSpacers = 1 // because explanation
					gr.BottomVSpacers = 3
				} else {
					gr.BottomVSpacers = 3
				}
				gr.RandomizationGroup = 1 - 1

				// q1b
				{
					inp := gr.AddInput()
					inp.Type = "dyn-composite"
					inp.ColSpanControl = 1
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundations__%v__%v", i, i)
				}
				_, inputNames, _ := qst.PoliticalFoundations(nil, i, i)
				for _, inpName := range inputNames {
					inp := gr.AddInput()
					inp.Type = "dyn-composite-scalar"
					inp.Name = inpName
				}

			}
		}

	}

	// page 3
	{

		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Stiftungen 3"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		page.ValidationFuncName = "patPage3"
		page.ValidationFuncMsg = trl.S{
			"de": "Wollen Sie wirklich weitergehen oder wollen Sie Ihre bisherigen Antworten vervollständigen?",
			// "en": "Does not add up. Really continue?",
		}

		// loop over matrix questions
		for i := 3; i < 6; i++ {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 3
				gr.RandomizationGroup = 1 - 1

				// q1b
				{
					inp := gr.AddInput()
					inp.Type = "dyn-composite"
					inp.ColSpanControl = 1
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundations__%v__%v", i, i)
				}
				_, inputNames, _ := qst.PoliticalFoundations(nil, i, i)
				for _, inpName := range inputNames {
					inp := gr.AddInput()
					inp.Type = "dyn-composite-scalar"
					inp.Name = inpName
				}
			}
		}

	}

	// page 4
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Auswertung 1"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `

					<p>
					<b>Frage 1.</b> Schätzen Sie bitte: Was wäre eine zufällig ausgewählte Person 
					aus unserer Vorstudie maximal bereit zu zahlen, 
					damit eine Spende von 30&nbsp;€ an die Stiftung überwiesen wird, 
					die diese Person als am besten/mittel/am schlechtesten erachtet?
					<i>(Wenn Sie meinen, die Person würde dafür bezahlen, 
						dass die Stiftung die 30&nbsp;€ <i>nicht</i> erhält, 
						schreiben Sie bitte ein Minuszeichen vor den entsprechenden Betrag.)</i>
					</p>

					<p>
					<!--
					Beste Stiftung:_______	Mittlere Stiftung:_______	Schlechteste Stiftung:_______
					</p>
					-->


					`,
				}
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 15
			gr.BottomVSpacers = 2

			// q2
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q2_a"
				inp.MaxChars = 3
				inp.Min = -999
				inp.Max = 999
				inp.ColSpan = 5
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 2
				inp.Label = trl.S{"de": "Beste Stiftung"}
				inp.Suffix = trl.S{"de": "€"}
				inp.Validator = "inRange1000"
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q2_b"
				inp.MaxChars = 3
				inp.Min = -999
				inp.Max = 999
				inp.ColSpan = 5
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 2
				inp.Label = trl.S{"de": "Mittlere Stiftung"}
				inp.Suffix = trl.S{"de": "€"}
				inp.Validator = "inRange1000"
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q2_c"
				inp.MaxChars = 3
				inp.Min = -999
				inp.Max = 999
				inp.ColSpan = 5
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 2
				inp.Label = trl.S{"de": "Schlechteste Stiftung"}
				inp.Suffix = trl.S{"de": "€"}
				inp.Validator = "inRange1000"
			}
		}

	}

	// zweiter Teil

	// page 5
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Auswertung 2"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		page.ValidationFuncName = "patPage5"
		page.ValidationFuncMsg = trl.S{
			"de": "Wollen Sie wirklich weitergehen oder wollen Sie Ihre bisherigen Antworten vervollständigen?",
			// "en": "Does not add up. Really continue?",
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
					<p>
						<b>
						Nun kommen wir zum zweiten Teil unserer Studie. 
						</b>
						In diesem Teil treffen Sie jeweils Entscheidungen für einen deutschen Staatsangehörigen, der Ihnen zugeordnet ist und der an einer zukünftigen Studie teilnehmen wird. Diese Person wird in der Studie entscheiden, wie ihr das Entgelt für die Studienteilnahme ausbezahlt wird. Je eher diese Person bereit ist, auf ihr Geld zu warten, desto mehr Geld wird ihr insgesamt bezahlt.						
					</p>
					<p>
						Wir bitten Sie zu entscheiden, wie geduldig oder wie ungeduldig die Person wählen kann. 
						
						Dazu bestimmen Sie für jede von drei Optionen, ob die jeweilige Option der Person zur Verfügung stehen soll oder nicht. 
						
						Falls Sie mehrere Optionen verfügbar machen, kann die Person aus diesen wählen. Mindestens eine Option muss „Verfügbar“ sein.
					</p>

					<p style="font-size: 87%;">
						<i>
						Details: Die nicht verfügbaren Optionen werden der Person nicht als Auswahloptionen angezeigt. Bei verfügbar gemachten Optionen können Sie zusätzlich „Von dieser Option abraten“ ankreuzen. In diesem Fall erhält die Person die Botschaft: „Ein früherer Teilnehmer dieser Studie rät Ihnen davon ab, diese Option zu wählen”.
						</i>
					</p>
					<br>


					<p>
						<b>Entscheidung 7. </b><br>
						Welche Optionen sollen der Person (nicht) zur Verfügung stehen, falls die Optionen wie folgt lauten?

					</p>


					`,
				}
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 2

			// q3a
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "TimePreferenceSelf__0__0"
			}
			_, inputNames, _ := qst.TimePreferenceSelf(nil, 0, 0)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "dyn-composite-scalar"
				inp.Name = inpName
			}
		}

		// gr3
		{

			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `

					<p>
					<b>Entscheidung 8. </b><br>
					Welche Optionen sollen der Person (nicht) zur Verfügung stehen, falls die Optionen wie folgt lauten?
					<i>(Beachten Sie: Sowohl die Zeitpunkte der Auszahlung als auch die Beträge sind anders als in der vorherigen Entscheidung.)</i>
					</p>

					`,
				}
			}
		}

		// gr4
		{
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 2

			// q3b
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "TimePreferenceSelf__1__1"
			}
			_, inputNames, _ := qst.TimePreferenceSelf(nil, 1, 1)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "dyn-composite-scalar"
				inp.Name = inpName
			}
		}

	}

	// page 6
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Gruppen-<br>präferenzen"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		page.ValidationFuncName = "patPage6"
		page.ValidationFuncMsg = trl.S{"de": "Wollen Sie wirklich weiterfahren, ohne dass sich Ihre Eintraege auf 10 summieren?"}

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
					<p>
					<b>Frage 2. </b>
						Schätzen Sie bitte: Wie viele Mitglieder einer Gruppe von 10 zufällig ausgewählten Personen, die an einer solchen Studie teilnehmen, wählen jeweils die folgenden Optionen A, B und C, 
					<b>
						wenn sie sich jeweils für genau eine der drei Optionen entscheiden müssen?					
					</b>

					<br>
					<i>(Ihre Antworten müssen sich auf 10 summieren.)</i>
					</p>
					`,
				}
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 0

			// q4a
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 12
				inp.DynamicFunc = "GroupPreferences__0__0"
			}
			_, inputNames, _ := qst.GroupPreferences(nil, 0, 0)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Name = inpName
			}
		}

		inpSt1 := css.NewStylesResponsive(nil)
		inpSt1.Desktop.StyleGridItem.Order = 2

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 24 - 0
			gr.BottomVSpacers = 3
			// q4a
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q4a_opt1"
				inp.MaxChars = 2
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 8
				inp.ColSpanLabel = 3 - 2
				inp.ColSpanControl = 9
				inp.Label = trl.S{"de": "von 10 wählen Option&nbsp;A"}
				// inp.Label = trl.S{"de": " "}
				// inp.Suffix = trl.S{"de": "von 10<br>wählen Option&nbsp;A"}
				inp.Validator = "inRange10"
				inp.StyleLbl = inpSt1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q4a_opt2"
				inp.MaxChars = 2
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 8
				inp.ColSpanLabel = 3 - 2
				inp.ColSpanControl = 9
				inp.Label = trl.S{"de": "von 10 wählen Option&nbsp;B"}
				// inp.Label = trl.S{"de": " "}
				// inp.Suffix = trl.S{"de": "von 10<br>wählen Option&nbsp;B"}
				inp.Validator = "inRange10"
				inp.StyleLbl = inpSt1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q4a_opt3"
				inp.MaxChars = 2
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 8
				inp.ColSpanLabel = 3 - 2
				inp.ColSpanControl = 9
				inp.Label = trl.S{"de": "von 10 wählen Option&nbsp;C"}
				// inp.Label = trl.S{"de": " "}
				// inp.Suffix = trl.S{"de": "von 10<br>wählen Option&nbsp;C"}
				inp.Validator = "inRange10"
				inp.StyleLbl = inpSt1
			}
		}

		//
		//
		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{
					"de": `
					<p>
					<b>Frage 3. </b>
					Und wie lautet Ihre Schätzung für die folgenden drei Optionen?
					<br>
					<i>(Ihre Antworten müssen sich auf 10 summieren.)
					</i>
					<br>
					<i>
						Bitte beachten Sie, dass die Zeitpunkte und Beträge anders sind als in Frage 2.
					</i>

					</p>
					`,
				}
			}
		}

		// gr4
		{
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 0

			// q4b
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 12
				inp.DynamicFunc = "GroupPreferences__1__1"
			}
			_, inputNames, _ := qst.GroupPreferences(nil, 1, 1)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "dyn-composite-scalar"
				inp.Name = inpName
			}
		}

		// gr5
		{
			gr := page.AddGroup()
			gr.Cols = 24 - 0
			gr.BottomVSpacers = 3

			// q4b
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q4b_opt1"
				inp.MaxChars = 2
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 8
				inp.ColSpanLabel = 3 - 2
				inp.ColSpanControl = 9
				inp.Label = trl.S{"de": "von 10 wählen Option&nbsp;A"}
				// inp.Label = trl.S{"de": " "}
				// inp.Suffix = trl.S{"de": "von 10<br>wählen Option&nbsp;A"}
				inp.Validator = "inRange10"
				inp.StyleLbl = inpSt1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q4b_opt2"
				inp.MaxChars = 2
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 8
				inp.ColSpanLabel = 3 - 2
				inp.ColSpanControl = 9
				inp.Label = trl.S{"de": "von 10 wählen Option&nbsp;B"}
				// inp.Label = trl.S{"de": " "}
				// inp.Suffix = trl.S{"de": "von 10<br>wählen Option&nbsp;B"}
				inp.Validator = "inRange10"
				inp.StyleLbl = inpSt1
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q4b_opt3"
				inp.MaxChars = 2
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 8
				inp.ColSpanLabel = 3 - 2
				inp.ColSpanControl = 9
				inp.Label = trl.S{"de": "von 10 wählen Option&nbsp;C"}
				// inp.Label = trl.S{"de": " "}
				// inp.Suffix = trl.S{"de": "von 10<br>wählen Option&nbsp;C"}
				inp.Validator = "inRange10"
				inp.StyleLbl = inpSt1
			}
		}

	}

	grStPage78 := css.NewStylesResponsive(nil)
	grStPage78.Desktop.StyleGridContainer.GapRow = "0.1rem"
	grStPage78.Desktop.StyleGridContainer.GapColumn = "0.01rem"

	// page 7
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Eigene Einstellung 1"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "30rem")

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsOneToSeven1,
				[]string{"q4"},
				radioVals7,
				[]trl.S{},
			)
			gb.MainLabel = trl.S{
				"de": `
				<p>
				<b>Frage 4.</b> 
				Wie sehr stimmen Sie der folgenden Aussage zu: 
				<i>„Alle Erwerbstätigen in Deutschland sollten verpflichtend 
				einen gewissen Teil ihres Arbeitseinkommens 
				im Rahmen einer privaten Altersvorsorge sparen, 
				um eine Rentenhöhe zu erreichen, die über dem Rentenanspruch 
				aus der gesetzlichen Rentenversicherung liegt.</i>“
				</p>
				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
			gr.BottomVSpacers = 4
		}

	}

	// page 8
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Eigene Einstellung 2"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "30rem")

		// gr1
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsOneToSeven2,
				[]string{"q5"},
				radioVals7,
				[]trl.S{},
			)
			gb.MainLabel = trl.S{
				"de": `
					<p>
					<b>Zum Schluss bitten wir Sie, drei Fragen über sich selbst zu beantworten:</b>

					<br>
					<br>
					<b>Frage 5.</b>
					Sind Sie im Vergleich zu anderen im Allgemeinen bereit, 
					heute auf etwas zu verzichten, 
					um in der Zukunft davon zu profitieren, 
					oder sind Sie im Vergleich zu anderen dazu nicht bereit? 

					</p>

				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsOneToSeven3,
				[]string{"q6"},
				radioVals7,
				[]trl.S{},
			)
			gb.MainLabel = trl.S{
				"de": `
					</p>
					<b>Frage 6.</b>
					Wie schätzen Sie sich persönlich ein? 
					Sind Sie im Allgemeinen ein risikobereiter Mensch 
					oder versuchen Sie, Risiken zu vermeiden?
					</p>

				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
		}

		// gr3
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate7,
				labelsOneToSeven2,
				[]string{"q7"},
				radioVals7,
				[]trl.S{},
			)
			gb.MainLabel = trl.S{
				"de": `
					<p>
					<b>Frage 7.</b>
					Wie schätzen Sie Ihre Bereitschaft ein, mit anderen zu teilen, 
					ohne dafür eine Gegenleistung zu erwarten?
					</p>
				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
			gr.BottomVSpacers = 4
		}

		//
		// explicit button to finish page, which is outsite navigation
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			gr.Cols = 2

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "finished"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 2
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"

				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
			}
		}

	}

	//
	//
	// page end
	// End page is a copy of page finish
	// without "End" button
	// without navigation
	{
		page := q.AddPage()
		page.Label = cfg.Get().Mp["end"]
		page.NoNavigation = true
		page.Style = css.DesktopWidthMaxForPages(page.Style, "30rem")

		{
			// Only one group => shuffling is no problem
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = cfg.Get().Mp["entries_saved"]
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				// inp.ColSpanLabel = 2
				inp.Desc = trl.S{"de": "Vielen Dank für das Ausfüllen dieser Umfrage! "}
			}

		}

	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	if err := q.TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := q.Validate(); err != nil {
		return &q, err
	}

	q.VersionEffective = -2

	return &q, nil

}
