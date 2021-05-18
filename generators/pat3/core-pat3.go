package pat3

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// POP3Part1Intro renders
func POP3Part1Intro(q *qst.QuestionnaireT) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	{
		gr := page.AddGroup()
		gr.Cols = 1
		// gr.BottomVSpacers = 2

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<p>
					Bisher konnten Sie die Optionen festgelegen, 
					die einer Person (nicht) zur Verfügung stehen sollen, 
					die an einer zukünftigen Studie teilnimmt. 
					
					Wir haben drei verschiedenen Gruppen von Studienteilnehmern 
					die gleichen Fragen wie Ihnen gestellt. 
					Die drei Gruppen sind die Folgenden:  				
				</p>

				<p style='padding-bottom: 0; padding-top: 0.5rem'>
					Es geht in diesem Teil wieder um die drei Gruppen aus dem letzten Teil:
				</p>

				<ul>
					<li>
						Gruppe 1: Eine repräsentative Gruppe deutscher Bürger.
					</li>
					<li>
						Gruppe 2: Eine repräsentative Gruppe deutscher Land- und Bundestagspolitiker.
					</li>
					<li>
						Gruppe 3: Eine Gruppe deutscher Bürger, die <i>keine Politiker</i> sind, 
						die aber die <i>gleichen demographischen Eigenschaften wie Politiker</i> haben. 
						
						Das heißt, Gruppe 3 besteht z. B. zu 70&nbsp;% aus Männern, 
						nur 3&nbsp;% der Mitglieder sind unter 30&nbsp;Jahre alt, 
						87&nbsp;% der Mitglieder haben einen Hochschulabschluss 
						und nur 17&nbsp;% sind alleinstehend.
					</li>
				<ul>

				<p>
					Wir haben aus jeder der drei Gruppen jeweils eine Person zufällig für Sie ausgewählt. 
					Falls dieser zweite Teil der Studie umgesetzt wird, 
					wird nicht Ihre eigene Entscheidung bestimmen, 
					welche Optionen dem zukünftigen Teilnehmer zur Verfügung stehen. 
					Stattdessen wird eine Entscheidung eines der drei ausgewählten 
					Gruppenmitglieder ausschlaggebend dafür sein, welche Optionen das sein werden.					
				</p>

				

				<p>
					Sie können aber entscheiden, welcher Gruppe (also Gruppe 1, 2 oder 3) 
					die Person angehören soll, die bestimmt, 
					welche Optionen dem zukünftigen Teilnehmer zur Verfügung stehen.				
				</p>

				`,
			}
		}

	}
	return nil

}

func POP3Part1Decision34(q *qst.QuestionnaireT, decisionNumber int, inpName string) error {

	keyVals := []string{
		"pol_gr1:Ein Politiker aus Gruppe 1 <br>(deutsche Land- und Bundestagspolitiker)",
		"cit_gr2:Ein Bürger aus Gruppe 2    <br>(repräsentativer deutscher Bürger)",
		"cit_gr3:Ein Bürger aus Gruppe 3    <br>(deutsche Bürger mit gleichen demographischen Eigenschaften wie die Politiker)",
	}

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	//
	if decisionNumber == 3 {
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = -1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": fmt.Sprintf(`
			<p><b> Entscheidung %v </b> </p>
			<p style='padding-bottom:0'>
				In dieser Entscheidung kann das Gruppenmitglied 
				folgende Optionen verfügbar machen:
			</p>
			`, decisionNumber)}
		}
	}

	//
	if decisionNumber == 4 {
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = -1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": fmt.Sprintf(`
			<p><b> Entscheidung %v </b> </p>
			<p style='padding-bottom:0'>
				Wer soll die Verfügbarkeit der Optionen festlegen, 
				falls diese stattdessen wie folgt lauten? 
				(Sowohl die Zeitpunkte der Auszahlung 
				als auch die Beträge sind anders als in der vorherigen Entscheidung.)
			</p>
			`, decisionNumber)}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "dyn-composite"
			inp.ColSpanControl = 12
			if decisionNumber == 3 {
				inp.DynamicFunc = "GroupPreferences__0__0"
			}
			if decisionNumber == 4 {
				inp.DynamicFunc = "GroupPreferences__1__1"
			}
		}
	}

	//
	//
	gr := page.AddGroup()
	gr.Cols = 1
	gr.BottomVSpacers = 0
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = 1
		inp.Desc = trl.S{"de": `
			<p>
				Wer soll die verfügbaren Optionen für den zukünftigen 
				Teilnehmer festlegen?
			</p>
			`}
	}
	for _, kv := range keyVals {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			// if decisionNumber > -4 && idx == 2 {
			// 	gr.BottomVSpacers = 2
			// }
			gr.RandomizationGroup = 1
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q1"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.Label = lbl
			rad.ControlFirst()
			rad.ControlTop()
		}
	}

	//
	//
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
			<p>
				<br>
				Wer soll die verfügbaren Optionen 
				für den zukünftigen Teilnehmer möglichst <i>nicht</i> festlegen? 
			</p>
			`}
		}
	}
	for _, kv := range keyVals {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 2
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q2"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.Label = lbl
			rad.ControlFirst()
			rad.ControlTop()
		}
	}

	//
	//

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if decisionNumber == 3 {
				inp.Desc = trl.S{"de": `
			<p style="font-size:86%">
				<br>
				<b>Erläuterung:</b>
				Falls dieser Teil umgesetzt wird, 
				werden zufällig zwei der drei Gruppen ausgewählt, 
				die tatsächlich festlegen können, welche Auswahloptionen 
				dem zukünftigen Studienteilnehmer zur Verfügung stehen. 
				Die dritte Gruppe wird die Entscheidung definitiv nicht treffen. 
				Von den zwei Gruppen, die die Entscheidung treffen können, 
				wird jene die Entscheidung treffen, die Sie 
				gemäß Ihrer Antworten auf die letzten Fragen als besser erachten.				
			</p>
			`}
			} else {
				inp.Desc = trl.S{"de": `<p></p>`}
			}
		}
	}

	return nil
}

// POP3Part2Intro renders
func POP3Part2Intro(q *qst.QuestionnaireT) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	{
		gr := page.AddGroup()
		gr.Cols = 1
		// gr.BottomVSpacers = 2

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
			<p>
				In diesem Teil der Studie beantworten Sie sechs Fragen. 
				Nach der Erhebung werden 10&nbsp;% aller Teilnehmer zufällig ausgewählt. 
				Jeder ausgewählte Teilnehmer 
				wird in Abhängigkeit der Genauigkeit seiner Antworten 
				eine Bonuszahlung von bis zu 50&nbsp;Norstat Coins erhalten (Wert: 5&nbsp;Euro).

				<br>
				<br>

				Es geht in diesem Teil wieder um die drei Gruppen aus dem letzten Teil:
			</p>


			<ul>
				<li>
					Gruppe 1: Eine repräsentative Gruppe deutscher Bürger.
				</li>
				<li>
					Gruppe 2: Eine repräsentative Gruppe deutscher Land- und Bundestagspolitiker.
				</li>
				<li>
					Gruppe 3: Eine Gruppe deutscher Bürger, die <i>keine Politiker</i> sind, 
					die aber die <i>gleichen demographischen Eigenschaften wie Politiker</i> haben. 
					
					Das heißt, Gruppe 3 besteht z. B. zu 70&nbsp;% aus Männern, 
					nur 3&nbsp;% der Mitglieder sind unter 30&nbsp;Jahre alt, 
					87&nbsp;% der Mitglieder haben einen Hochschulabschluss 
					und nur 17&nbsp;% sind alleinstehend.
				</li>
			<ul>

			<p>
				Wir bitten Sie zu schätzen, 
				welche Optionen die Mitglieder dieser Gruppen (nicht) verfügbar gemacht haben. 
			</p>


			<p>
				Falls Sie eine Bonuszahlung erhalten, 
				werden wir eine der sechs Fragen zufällig auswählen 
				und Ihre Schätzung mit den echten Entscheidungen 
				der Gruppenmitglieder abgleichen. 
				Ihre Bonuszahlung ist umso höher, 
				je genauer Ihre Einschätzung ist. 
				Bitte überlegen Sie sich Ihre Antworten daher sehr genau!
			</p>


			<p style="font-size:86%">
				<br>
				<b>Erläuterung:</b>
				Falls Sie in der ausgewählten Frage eine 100&nbsp;% richtige 
				Antwort geben, werden Sie 50&nbsp;Norstat&nbsp;coins erhalten. 
				
				Für jede Person, die Sie bei Ihren folgenden Schätzungen z
				u viel oder zu wenig angeben, werden Sie 2.5&nbsp;Norstat&nbsp;coins 
				verlieren. 
				
				Falls beispielsweise alle 10&nbsp;Gruppenmitglieder Stiftung&nbsp;C 
				wählten, Sie aber angeben, 
				dass 5&nbsp;Gruppenmitglieder Stiftung&nbsp;B wählen, 
				und weitere 5 Stiftung&nbsp;C wählten, 
				dann haben Sie für Stiftung&nbsp;C fünf Gruppenmitglieder 
				zu wenig angegeben, 
				und für Stiftung&nbsp;B fünf zu viel. 
				Entsprechend wird Ihre Bezahlung 
				auf 50-2.5× 5 -2.5× 5=25 Norstat&nbsp;coins gesenkt.
			</p>




			`,
			}
		}

	}
	return nil

}

func POP3Part2Questions123and456(q *qst.QuestionnaireT, start int) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 0

		if start == 1 {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
			
			<p style='padding-bottom: 0.12rem;'><b>
			Fragen %v. - %v.
			</b></p>				
			
			<p>
			Schätzen Sie im Folgenden für diese Entscheidungssituation:
			</p>

			`, start+0, start+2),
			}
		}
		if start == 4 {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
			
			<p style='padding-bottom: 0.12rem;'><b>
			Fragen %v. - %v.
			</b></p>

			<p style='padding-bottom: 0.2rem;'>
			Was glauben Sie, wie haben die Mitglieder der drei Gruppen entschieden, 
			falls die Optionen stattdessen wie folgt gegeben waren? 
			(Sowohl die Zeitpunkte der Auszahlung 
			als auch die Beträge sind anders als in den vorherigen Fragen.)
			</p>
			`, start+0, start+2),
			}
		}

		//
		//
		if start == 1 {
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "TimePreferenceSelfStatic__0__0"
			}
		}
		if start == 4 {
			gr := page.AddGroup()
			gr.Cols = 12
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				inp.DynamicFunc = "TimePreferenceSelfStatic__1__1"
			}
		}

		part2ThreeQuestions(q, start)

		if start == 4 {
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": `
					<br>
					<p>
					<b>
						Dies ist das Ende dieser Studie. 
						Wir bedanken uns ganz herzlich für Ihre Teilnahme. 
						Falls Sie zu den zufällig ausgewählten 10% gehören, 
						werden Sie Ihre Bonuszahlung wie versprochen in den nächsten Tagen erhalten. 
					</b>
					</p>
				`}
			}
		}

	}

	return nil
}

func part2ThreeQuestions(q *qst.QuestionnaireT, blockStart int) error {

	pageIdx := len(q.Pages) - 1
	page := q.EditPage(pageIdx)

	page.ValidationFuncName = "pat3-add-to-10"
	page.ValidationFuncMsg = trl.S{"de": "Wollen Sie wirklich weiterfahren, ohne dass sich Ihre Eintraege auf 10 summieren?"}

	keyVals1 := []string{
		"Wie viele von 10 zufällig ausgewählten Politikern aus Gruppe 1 haben jeweils die Optionen A, B oder C verfügbar gemacht?",
		"Wie viele von 10 zufällig ausgewählten Bürgern aus Gruppe 2 haben jeweils die Optionen A, B oder C verfügbar gemacht?",
		"Wie viele von 10 zufällig ausgewählten Bürgern aus Gruppe 3 (gleiche demographische Eigenschaften wie die Politiker) haben jeweils die Optionen A, B oder C verfügbar gemacht?",
	}
	keyVals1a := []string{
		"Politikern",
		"Bürgern",
		"Bürgern aus Gruppe 3 (gleiche demographische Eigenschaften wie die Politiker)",
	}

	keyVals2 := []string{
		"von 10&nbsp;%v haben <b>alle Optionen verfügbar</b> gemacht",
		"von 10&nbsp;%v haben <b>nur Optionen A und B verfügbar</b> gemacht",
		"von 10&nbsp;%v haben <b>nur Option A verfügbar</b> gemacht",
		"von 10&nbsp;%v haben <b>andere Optionen verfügbar</b> gemacht <br>(z. B. nur Option B oder nur Option C)",
	}

	inpName := "pop3_part2"

	for idx1, kv := range keyVals1 {

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{"de": fmt.Sprintf(
					`
				<p><b>Frage %v.</b></p>
				<p>%v</p>
				`,
					blockStart+idx1,
					kv,
				)}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "dyn-composite"
				inp.ColSpanControl = 1
				if blockStart == 1 {
					inp.DynamicFunc = "GroupPreferences__0__0"
				}
				if blockStart == 4 {
					inp.DynamicFunc = "GroupPreferences__1__1"
				}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			for idx2, kv := range keyVals2 {
				lbl := fmt.Sprintf(kv, keyVals1a[idx1])
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = inpName + fmt.Sprintf("_q%v_%v", blockStart+idx1, idx2+1)
				inp.MaxChars = 3
				inp.Min = 0
				inp.Max = 10
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 14
				inp.Label = trl.S{"de": lbl}
				inp.Validator = "inRange10"

				inp.ControlFirst()
				inp.ControlTop()
			}

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1
				inp.Desc = trl.S{
					"de": `
					<p style='font-size:90%'>
					Ihre Antworten müssen sich auf 10 summieren.	
					</p>
					`,
				}
				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
				inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "center"
			}
		}
	}

	return nil
}
