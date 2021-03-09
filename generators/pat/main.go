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
	q.Version = 2

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Paternalismus Umfrage"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}
	q.Survey.Name = trl.S{"de": "Politische Entscheidungsprozesse"}
	q.Survey.Name = trl.S{"de": "Entscheidungsprozesse in der Politik"}
	q.Variations = 4
	q.Variations = 0

	// page 0
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.NoNavigation = true
		page.Style = css.DesktopWidthMax(page.Style, "36rem") // 60

		//
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{
				"de": "HERZLICH WILLKOMMEN UND VIELEN DANK FÜR IHRE TEILNAHME!<br><br>",
			}
			inp.Desc = trl.S{
				"de": `
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

				<p>In dieser Umfrage gibt es keine richtigen oder falschen Antworten. 
				Bitte entscheiden Sie daher immer gemäß Ihren persönlichen Ansichten. 
				Sie werden dabei vollständig anonym bleiben.
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

	// page 1
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Stiftungen 1"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem") // 60

		page.ValidationFuncName = "patPage1"
		page.ValidationFuncMsg = trl.S{
			"de": "no message",
			// "en": "Does not add up. Really continue?",
		}

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
				// inp.Label = trl.S{"de": "input label"}
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
				Für jede Ihrer ersten sechs Entscheidungen zeigen wir Ihnen die Präferenzen fünf deutscher Staatsangehöriger darüber, 
				welche der drei Stiftungen die Spende erhalten soll. 
				Sie entscheiden, wie die Präferenzen der fünf Personen in <i>eine</i> gemeinsame Entscheidung zusammengefasst werden.
				</p>

				<p>
				Die Präferenzen stammen von fünf Personen, die an einer Vorstudie teilgenommen haben<sup>1)</sup>. 
				Diese fünf Personen wurden aus einer Stichprobe gezogen, in der sich gleich viele Personen politisch links, 
				in der Mitte oder als konservativ einordnen. 
				Jede Person wurde einzeln befragt, welche der drei Stiftungen sie als am besten, mittel, und am schlechtesten erachtet. 
				</p>
				
				<p>
				Den Personen wurde mitgeteilt, dass ihre Präferenzen 
				zusammen mit den Präferenzen von vier anderen Personen 
				an einen zukünftigen Teilnehmer der Studie gegeben werden, 
				der die Präferenzen in eine Entscheidung zusammenfasst. 
				<b>Dieser zukünftige Teilnehmer sind Sie.</b>
				</p>
				
				`}
			}
		}

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 1 - 1
			// q1-pretext
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PoliticalFoundationsPretext__0__0"
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 1 - 1

			// q1a
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "PoliticalFoundations__0__0"
			}
			_, inputNames, _ := qst.PoliticalFoundations(nil, 0, 0)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "dyn-composite-scalar"
				inp.Name = inpName + "_page0"
			}

		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = trl.S{"de": `
				<p>
				Die Stiftungen wurden anonymisiert und in eine zufällige Reihenfolge gebracht, so dass Sie nicht wissen, 
				um welche Stiftung es sich bei den Stiftungen A, B und C handelt. 
				Sie entscheiden also nicht darüber, welche Stiftung die 30&nbsp;€ erhält. 
				Stattdessen entscheiden Sie, wie die Präferenzen der Gruppenmitglieder in <i>eine</i> Entscheidung zusammengefasst werden 
				und ob Sie beispielsweise eher eine Kompromisslösung oder eher eine Mehrheitslösung für Ihre Gruppe bevorzugen.
				</p>
				`}
			}
		}

		// gr4
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{
					"de": `
				<span style="display: inline-block; font-size:87%; line-height: 120%;">

				<sup>1)</sup>
				Nur in einer der sechs Entscheidungen stammen die Präferenzen von den Personen, die wir befragt haben, und nur diese Entscheidung kann umgesetzt werden. In den anderen Entscheidungen wurden die Präferenzen von uns zusammengestellt. Da Sie nicht wissen, in welcher Entscheidung die Präferenzen von den Befragten stammen, sollten Sie in allen sechs Entscheidungen so entscheiden, als seien die jeweiligen Präferenzen von der echten Gruppe.
				</span>
				`,
				}
			}
		}

	}

	// page 2
	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Short = trl.S{"de": "Stiftungen 2"}
	page.Style = css.DesktopWidthMax(page.Style, "36rem") // 60

	page.ValidationFuncName = "patPage2"
	page.ValidationFuncMsg = trl.S{
		"de": "Sie haben nicht alle drei Entscheidungen bewertet. Wirklich weiter?",
		// "en": "Does not add up. Really continue?",
	}

	// gr0
	{
		gr := page.AddGroup()
		gr.Cols = 2
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.ColSpan = 2
			inp.ColSpanLabel = 2
			inp.Type = "textblock"
			inp.Desc = trl.S{"de": `
			<p>
			Entscheiden Sie im Folgenden, an welche Stiftung das Geld gehen soll. 
			<b>Setzen Sie dazu bei der entsprechenden Stiftung ein Kreuz in der Spalte „Auswahl“.</b>
			Falls Sie eine zweite oder dritte Alternative als genauso gut empfinden, 
			setzen Sie ein Kreuz in der Spalte „Gleich gut“. 
			Berücksichtigen Sie die dargestellten Präferenzen der Gruppen&shy;mitglieder.
			</p>
			`}
		}

		// loop over matrix questions
		for i := 0; i < 3; i++ {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 2
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
		page.Style = css.DesktopWidthMax(page.Style, "36rem") // 60

		page.ValidationFuncName = "patPage3"
		page.ValidationFuncMsg = trl.S{
			"de": "Sie haben nicht alle drei Entscheidungen bewertet. Wirklich weiter?",
			// "en": "Does not add up. Really continue?",
		}

		// loop over matrix questions
		for i := 3; i < 6; i++ {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 2
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
		page.Short = trl.S{"de": "Auswertung"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem") // 60

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": "Frage 3<br>"}
				inp.Desc = trl.S{
					"de": `

					<p>
					<b>Frage 1.</b> Schätzen Sie bitte: Was wäre eine zufällig ausgewählte Person aus unserer Vorstudie maximal bereit zu zahlen, damit eine Spende von 30&nbsp;€ an die Stiftung überwiesen wird, die diese Person als am besten/mittel/am schlechtesten erachtet?
					<i>(Wenn Sie meinen, die Person würde dafür bezahlen, dass die Stiftung die 30&nbsp;€ nicht erhält, schreiben Sie bitte ein Minuszeichen vor den jeweiligen Betrag.)</i>
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
				inp.Desc = trl.S{"de": "Beste Stiftung"}
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
				inp.Desc = trl.S{"de": "Mittlere Stiftung"}
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
				inp.Desc = trl.S{"de": "Schlechteste Stiftung"}
				inp.Suffix = trl.S{"de": "€"}
				inp.Validator = "inRange1000"
			}
		}

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": "Frage 3<br>"}
				inp.Desc = trl.S{
					"de": `
					<p>
					<b>Im Folgenden legen Sie fest, welche Optionen ein zukünftiger Studienteilnehmer wählen kann:</b>
					</p>

					<div class="vspacer-16"> &nbsp; </div>

					<p>
					<b>Entscheidung 7.</b><br>
					Sie sind einem deutschen Staatsangehörigen zugeordnet, der an einer zukünftigen Studie teilnehmen wird 
					und verschiedene Optionen für Geldauszahlungen an unterschiedlichen Zeitpunkten hat.
					
					<b>Er erhält in dieser Studie genau eine der unten beschriebenen Optionen, 
					die ihm tatsächlich an den genannten Zeitpunkten ausgezahlt wird.</b>
					</p>

					<p>
					Den Personen wurde mitgeteilt, dass ihre Präferenzen 
					zusammen mit den Präferenzen von vier anderen Personen 
					an einen zukünftigen Teilnehmer der Studie gegeben werden, 
					der die Präferenzen in <i>eine</i> Entscheidung zusammenfasst. 
					<b>Dieser zukünftige Teilnehmer sind Sie.</b>
					</p>


					<p>
					Sie können nun entscheiden, welche der drei Optionen der Person (nicht) zur Verfügung stehen sollen, 
					indem Sie in jeder Zeile ein Kreuz entweder bei „Verfügbar” oder bei „Nicht verfügbar” setzen. 
					<b>Es muss am Ende mindestens eine Option „Verfügbar“ sein.</b> 
					Folgen Sie dabei einfach Ihren Vorstellungen – es gibt keine richtigen oder falschen Antworten. 
					Die nicht verfügbaren Optionen werden der Person nicht als Auswahloptionen angezeigt. 
					Falls mehr als eine Option verfügbar ist, kann die die Person aus diesen Optionen wählen.
					</p>

					<p>
					Bei verfügbar gemachten Optionen können Sie zusätzlich „Von dieser Option abraten“ ankreuzen. In diesem Fall erhält die Person die Botschaft: „Ein früherer Teilnehmer dieser Studie rät Ihnen davon ab, diese Option zu wählen”.
					</p>

					`,
				}
			}
		}

		// gr3
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

		// gr4
		{

			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": "Frage 3<br>"}
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

		// gr5
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

	// page 5
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Gruppen-<br>präferenzen"}
		page.Style = css.DesktopWidthMax(page.Style, "36rem") // 60

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.Label = trl.S{"de": "Frage 2<br>"}
				inp.Desc = trl.S{
					"de": `
					<p>
					<b>Frage 2. </b>
					Schätzen Sie bitte: Wie viele Mitglieder einer Gruppe von 10 zufällig ausgewählten Personen wählen jeweils die folgenden Optionen A, B und C.
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

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 24 - 0
			gr.BottomVSpacers = 2
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
				inp.Desc = trl.S{"de": "Wie viele wählen Option A? Ihre Antwort:"}
				inp.Desc = trl.S{"de": " "}
				inp.Suffix = trl.S{"de": "von 10<br>wählen Option A"}
				inp.Validator = "inRange10"
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
				inp.Desc = trl.S{"de": "Wie viele wählen Option B? Ihre Antwort:"}
				inp.Desc = trl.S{"de": " "}
				inp.Suffix = trl.S{"de": "von 10<br>wählen Option B"}
				inp.Validator = "inRange10"
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
				inp.Desc = trl.S{"de": "Wie viele wählen Option C? Ihre Antwort:"}
				inp.Desc = trl.S{"de": " "}
				inp.Suffix = trl.S{"de": "von 10<br>wählen Option C"}
				inp.Validator = "inRange10"
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
				// inp.Label = trl.S{"de": "Frage 3<br>"}
				inp.Desc = trl.S{
					"de": `
					<p>
					<b>Frage 3. </b>
					Und wie lautet Ihre Schätzung für die folgenden drei Optionen?
					<i>(Ihre Antworten müssen sich auf 10 summieren. Die Zeitpunkte und Beträge sind anders als in Frage 2.)</i>
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
			gr.BottomVSpacers = 2

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
				inp.Desc = trl.S{"de": "Wie viele wählen Option A? Ihre Antwort:"}
				inp.Desc = trl.S{"de": " "}
				inp.Suffix = trl.S{"de": "von 10<br>wählen Option A"}
				inp.Validator = "inRange10"
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
				inp.Desc = trl.S{"de": "Wie viele wählen Option B? Ihre Antwort:"}
				inp.Desc = trl.S{"de": " "}
				inp.Suffix = trl.S{"de": "von 10<br>wählen Option B"}
				inp.Validator = "inRange10"
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
				inp.Desc = trl.S{"de": "Wie viele wählen Option C? Ihre Antwort:"}
				inp.Desc = trl.S{"de": " "}
				inp.Suffix = trl.S{"de": "von 10<br>wählen Option C"}
				inp.Validator = "inRange10"
			}
		}

	}

	// page 6
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Eigene Einstellung"}
		page.Style = css.DesktopWidthMax(page.Style, "30rem")

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
				<i>„Alle Erwerbstätigen in Deutschland sollten verpflichtend einen gewissen Teil ihres Arbeitseinkommens 
				im Rahmen einer privaten Altersvorsorge sparen, 
				um eine Rentenhöhe zu erreichen, die über dem Rentenanspruch 
				aus der gesetzlichen Rentenversicherung liegt.</i>“
				</p>
				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		// gr2
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
					<b>Zum Schluss bitten wir Sie drei Fragen über sich selbst zu beantworten:</b>

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
		}

		// gr3
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
		}

		// gr4
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
					Wie schätzen Sie Ihre Bereitschaft ein mit anderen zu teilen, 
					ohne dafür eine Gegenleistung zu erwarten?
					</p>
				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// explicit button to finish page, which is outsite navigation
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			gr.Cols = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				// inp.CSSControl = "special-line-height-higher"
				inp.Desc = trl.S{"de": "Vielen Dank für das Ausfüllen dieser Umfrage! "}
				inp.Desc = trl.S{"de": "  "} // next page
			}

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "finished"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
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
		page.Style = css.DesktopWidthMax(page.Style, "30rem")

		{
			// Only one group => shuffling is no problem
			gr := page.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = cfg.Get().Mp["thanks_for_participation"]
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Desc = cfg.Get().Mp["entries_saved"]
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
