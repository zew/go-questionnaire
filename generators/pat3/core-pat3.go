package pat3

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/generators/pat2"
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
		gr.BottomVSpacers = 0

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<p>
					<b style='font-size: 110%'> Teil 2 </b> <br>
				
					Bisher konnten Sie die Optionen festgelegen, 
					die einer Person (nicht) zur Verfügung stehen sollen, 
					die an einer zukünftigen Studie teilnimmt.  
				</p>
				
				<p>
					Nun entscheiden Sie, an wen Sie diese Entscheidung delegieren möchten. 
					Die von Ihnen ausgewählte Person wird dann entscheiden, 
					welche Optionen dem deutschen Staatsangehörigen 
					zur Verfügung stehen werden. 
				</p>

				<p>
					Sie können Ihre Entscheidung an Personen aus den folgenden 
					drei Gruppen delegieren:
				</p>
				`,
			}
		}
	}
	for idx, txt := range pat2.PartIGroupsLong {
		gr := page.AddGroup()
		gr.RandomizationGroup = 1
		gr.RandomizationSeed = 1
		gr.BottomVSpacers = 0

		gr.Cols = 1
		{
			txt = fmt.Sprintf(txt, pat2.GroupIDs[idx])
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
				<ul>
				<li>
				%v
				</li>
				</ul>
				
				`, txt),
			}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<br>
				<p style="font-size:86%">
				<b>Erläuterung</b>: 
				
				<i>
				Wir haben den Mitgliedern dieser drei Gruppen 
				die gleichen Fragen wie Ihnen gestellt 
				und sie haben ihre Entscheidungen bereits gefällt. 
				
				Wir haben aus jeder der drei Gruppen jeweils eine Person zufällig 
				für Sie ausgewählt. 
				
				Falls dieser Teil der Studie umgesetzt wird, 
				wird die Entscheidung dieser Person bestimmen, 
				welche Optionen deutschen Staatsangehörigen 
				zur Verfügung stehen werden. 				
				</i>


				</p>

				

				`,
			}
		}

	}
	return nil

}

// POP3Part1Decision34 - part 1 of 2
func POP3Part1Decision34(q *qst.QuestionnaireT, decisionNumber int, inpName string) error {

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

				An wen möchten Sie die Entscheidung delegieren, 
				wenn die drei Optionen wie folgt gegeben sind?



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

				An wen möchten Sie die Entscheidung delegieren, 
				wenn die drei Optionen wie folgt gegeben sind?

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
				inp.DynamicFunc = "GroupPreferencesPOP3__0__0"
			}
			if decisionNumber == 4 {
				inp.DynamicFunc = "GroupPreferencesPOP3__1__1"
			}
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
			<p style="margin-bottom: 0.45rem">
				Wer soll entscheiden,
				welche Optionen 
				 dem*der zukünftigen Teilnehmer*in 
				zur Verfügung stehen werden?
			</p>
			`}
		}

		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.ColSpanControl = 1
			inp.DynamicFunc = "ErrorProxy"
			inp.Param = inpName + "_q1"
		}
	}

	for idx, kv := range pat2.PartIGroupsShort {
		{
			gr := page.AddGroup()
			gr.RandomizationGroup = 1
			gr.RandomizationSeed = 1

			gr.Cols = 1
			gr.BottomVSpacers = 1
			// if decisionNumber > -4 && idx == 2 {
			// 	gr.BottomVSpacers = 2
			// }
			gr.RandomizationGroup = 1
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			val = fmt.Sprintf(val, pat2.GroupIDs[idx])
			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q1"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.Label = lbl
			rad.ControlFirst()
			rad.ControlTop()

			if idx == 0 {
				rad.Validator = "mustRadioGroup;preventInversion"
			}
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
			<p style="margin-bottom: 0.45rem">
				<br>
				Wer soll die verfügbaren Optionen 
				für den*die zukünftige*n Teilnehmer*in 
				möglichst <i><b>nicht</b></i> festlegen? 
			</p>
			`}
		}

		{
			inp := gr.AddInput()
			inp.Type = "dyn-textblock"
			inp.ColSpanControl = 1
			inp.DynamicFunc = "ErrorProxy"
			inp.Param = inpName + "_q2"
		}
	}

	for idx, kv := range pat2.PartIGroupsShort {
		{
			gr := page.AddGroup()
			gr.RandomizationGroup = 2
			gr.RandomizationSeed = 1

			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 2
			sp := strings.Split(kv, ":")
			key := sp[0]
			val := sp[1]
			val = fmt.Sprintf(val, pat2.GroupIDs[idx])
			lbl := trl.S{"de": val}

			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = inpName + "_q2"
			rad.ValueRadio = key
			rad.ColSpan = 1
			rad.Label = lbl
			rad.ControlFirst()
			rad.ControlTop()

			if idx == 0 {
				rad.Validator = "mustRadioGroup;preventInversion"
			}
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
		gr.BottomVSpacers = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
			<p>

				<b style='font-size: 110%'> Teil 3 </b> <br>

				In diesem Teil der Studie beantworten Sie sechs Fragen. 
				Nach der Erhebung werden 10&nbsp;% aller Teilnehmer*innen zufällig ausgewählt. 

				Jede*r ausgewählte Teilnehmer*in 
				wird in Abhängigkeit der Genauigkeit 
				seiner*ihrer Antworten 
				eine Bonuszahlung von bis zu 50&nbsp;Norstat Coins erhalten 
				(Wert: 5&nbsp;Euro).

				<br>
				<br>

				Es geht in diesem Teil wieder um die drei Gruppen,
				an welche Sie im vorherigen Teil Ihre Entscheidung delegieren konnten:
			</p>



			`,
			}
		}
	}

	for idx, txt := range pat2.PartIGroupsLong {
		gr := page.AddGroup()
		// gr.RandomizationGroup = 1
		// gr.RandomizationSeed = 1
		gr.BottomVSpacers = 0

		gr.Cols = 1
		{
			txt = fmt.Sprintf(txt, pat2.GroupIDs[idx])
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": fmt.Sprintf(`
				<ul>
				<li>
				%v
				</li>
				</ul>
				
				`, txt),
			}
		}
	}

	//
	{
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `

			<br>
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
				Antwort geben, 
				erhalten Sie 50&nbsp;Norstat&nbsp;coins. 
				
				Für jede Person, die Sie bei Ihren Schätzungen 
				zu viel oder zu wenig angeben, 
				verlieren Sie 2,5&nbsp;Norstat&nbsp;coins. 
				
				Falls beispielsweise alle 10 Gruppenmitglieder 
				nur Option&nbsp;A verfügbar gemacht haben, 
				Sie aber angeben, dass 5 Gruppenmitglieder 
				nur Option&nbsp;A verfügbar gemacht haben, 
				und weitere 5 alle Optionen verfügbar gemacht haben, 
				dann haben Sie für den Fall des Verfügbarmachens 
				aller Optionen fünf Gruppenmitglieder zu wenig angegeben, 
				und für den Fall des Verfügbarmachens von 
				Option&nbsp;A (alleine) fünf zu viel. 
				
				Entsprechend wird Ihre Bezahlung auf 
				50 - 2,5*5 -2,5*5 = 25&nbsp;Norstat&nbsp;coins gesenkt.
			</p>




			`,
			}
		}

	}
	return nil

}

// POP3Part2Questions123and456 - part 2 of 2
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

		part3ThreeQuestions(q, start)

	}

	return nil
}

func part3ThreeQuestions(q *qst.QuestionnaireT, blockStart int) error {

	pageIdx := len(q.Pages) - 1
	page := q.EditPage(pageIdx)

	// page.ValidationFuncName = "pat3-add-to-10"
	// page.ValidationFuncMsg = trl.S{"de": "Wollen Sie wirklich weiterfahren, ohne dass sich Ihre Eintraege auf 10 summieren?"}

	variousOptionsMadeAvailablePerm1 := []string{
		"1;;von 10&nbsp;%v haben <b>nur Option A         verfügbar</b> gemacht",
		"2;;von 10&nbsp;%v haben <b>nur Optionen A und B verfügbar</b> gemacht",
		"3;;von 10&nbsp;%v haben <b>alle Optionen        verfügbar</b> gemacht",
		"4;;von 10&nbsp;%v haben <b>andere Optionen      verfügbar</b> gemacht (z. B. nur Option B oder nur Option C)",
	}
	variousOptionsMadeAvailablePerm2 := []string{
		"3;;von 10&nbsp;%v haben <b>alle Optionen        verfügbar</b> gemacht",
		"2;;von 10&nbsp;%v haben <b>nur Optionen A und B verfügbar</b> gemacht",
		"1;;von 10&nbsp;%v haben <b>nur Option A         verfügbar</b> gemacht",
		"4;;von 10&nbsp;%v haben <b>andere Optionen      verfügbar</b> gemacht (z. B. nur Option B oder nur Option C)",
	}

	inpName := "pop3_part2"

	for idx1, groupName := range pat2.Pat3Part2 {

		groupName = fmt.Sprintf(groupName, pat2.GroupIDs[idx1])

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.ColSpan = 1

				lbl := fmt.Sprintf(`
					Wie viele von 10 zufällig ausgewählten %v 
					haben jeweils die Optionen A, B oder C verfügbar gemacht?
				`, groupName)

				inp.Desc = trl.S{"de": fmt.Sprintf(
					`
				<p><b>Frage %v:</b>  &nbsp; %v </p>
				<p>%v</p>
				`,
					blockStart+idx1,
					pat2.PartIGroupsShortNominativ[idx1],
					lbl,
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
					inp.DynamicFunc = "GroupPreferencesPOP3__0__0"
				}
				if blockStart == 4 {
					inp.DynamicFunc = "GroupPreferencesPOP3__1__1"
				}
			}
		}

		for idx2 := 0; idx2 < len(variousOptionsMadeAvailablePerm1); idx2++ {
			// for idx2, avaiilable := range variousOptionsMadeAvailablePerm1 {

			avaiilable := variousOptionsMadeAvailablePerm1[idx2]
			if idx1 == 1 || idx2 == 1 {
				avaiilable = variousOptionsMadeAvailablePerm2[idx2]
			}

			avaiilableSl := strings.Split(avaiilable, ";;")

			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			// gr.RandomizationSeed = 1
			// gr.RandomizationGroup = 1 + idx1
			{

				// groupName := pat2.Pat3Part2[idx1]
				lbl := fmt.Sprintf(avaiilableSl[1], groupName)
				// lbl = avaiilableSl[1] + groupName
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = inpName + fmt.Sprintf("_q%v_%v", blockStart+idx1, avaiilableSl[0])
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

				if idx2 == 0 {
					inp.Validator += ";pop3_part2_q123456_1234"
				}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
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

// POP3Part2Questions78
func POP3Part2Questions78(q *qst.QuestionnaireT) error {

	page := q.AddPage()
	page.Label = trl.S{"de": ""}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Desc = trl.S{"de": `
			<p>
			Unabhängig von Ihren bisherigen Entscheidungen 
			können Sie abschließend entscheiden, 
			ob der*die zukünftige*n Studienteilnehmer*in, 
			der*die Ihnen zugeordnet ist, 
			einen zusätzlichen Bonus für seine*ihre 
			Studienteilnahme erhält 
			oder ob ihm*ihr Geld von seinem*ihrem Entgelt 
			für die Teilnahme abgezogen wird. 
			
			Sie entscheiden, ob wir die entsprechende 
			Bonuszahlung oder den Abzug 
			in der zukünftigen Studie umsetzen. 
			
			Eine der nächsten zwei Entscheidungen 
			wird zufällig ausgewählt und umgesetzt.
			
			</p>
			`}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 12
		gr.BottomVSpacers = 2

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
				<p style="margin-bottom: 0.55rem">
					Ich möchte, dass der zusätzliche Pauschalbetrag 
					von 0.50&nbsp;€ 
					zum Entgelt der Person hinzugefügt wird.
				</p>
				`,
			}
		}

		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "pop3_agio"
			rad.ValueRadio = "yes"
			rad.ColSpan = 4
			rad.ColSpanLabel = 4
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Ja&nbsp;",
				"en": "Yes",
			}
			rad.LabelRight()
			rad.Validator = "mustRadioGroup"
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "pop3_agio"
			rad.ValueRadio = "no"
			rad.ColSpan = 4
			rad.ColSpanLabel = 4
			rad.ColSpanControl = 1
			rad.LabelRight()
			rad.Label = trl.S{
				"de": "Nein",
				"en": "No",
			}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 12

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 12
			inp.Label = trl.S{
				"de": `
				<p style="margin-bottom: 0.7rem">
					Ich möchte, dass der Pauschalbetrag 
					von 0.50&nbsp;€ 
					vom Entgelt der Person abgezogen wird.
				</p>
				`,
			}
		}

		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "pop3_disagio"
			rad.ValueRadio = "yes"
			rad.ColSpan = 4
			rad.ColSpanLabel = 4
			rad.ColSpanControl = 1
			rad.Label = trl.S{
				"de": "Ja&nbsp;",
				"en": "Yes",
			}
			rad.LabelRight()
			rad.Validator = "mustRadioGroup"
		}
		{
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = "pop3_disagio"
			rad.ValueRadio = "no"
			rad.ColSpan = 4
			rad.ColSpanLabel = 4
			rad.ColSpanControl = 1
			rad.LabelRight()
			rad.Label = trl.S{
				"de": "Nein",
				"en": "No",
			}
		}
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1

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
	return nil
}
