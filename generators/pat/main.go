package pat

import (
	"fmt"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/ctr"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Create creates an minimal example questionnaire with a few pages and inputs.
// It is saved to disk as an example.
func Create(params []qst.ParamT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	// qst.RadioVali = "mustRadioGroup"
	qst.RadioVali = ""
	qst.CSSLabelHeader = ""
	qst.CSSLabelRow = ""

	q := qst.QuestionnaireT{}
	q.Survey = qst.NewSurvey("pat")
	q.Survey.Params = params
	q.LangCodes = []string{"de"} // governs default language code

	q.Survey.Org = trl.S{"de": "ZEW"}
	q.Survey.Name = trl.S{"de": "Paternalismus Umfrage"}
	q.Variations = 0
	q.Variations = 4

	userID := q.UserIDInt()

	// page 0
	{
		page := q.AddPage()
		page.Width = 50
		page.Label = trl.S{"de": "&nbsp;"}
		page.Label = trl.S{"de": ""}
		page.NoNavigation = true
		gr := page.AddGroup()
		gr.Cols = 1
		gr.Label = trl.S{
			"de": "HERZLICH WILLKOMMEN ZU UNSERER STUDIE UND VIELEN DANK FÜR IHRE TEILNAHME!<br><br>",
		}
		gr.Desc = trl.S{
			"de": `
<p>Dies ist eine Studie des Zentrums für €päische Wirtschaftsforschung (ZEW) in Mannheim sowie der Universitäten in Köln, Mannheim, Münster und Zürich. Ihre Teilnahme wird nur wenige Minuten in Anspruch nehmen und Sie unterstützen damit die Forschung zu Entscheidungsprozessen in der Politik.
</p>

<p>In dieser Studie treffen Sie acht Entscheidungen und beantworten sieben Fragen. Nach der Erhebung werden 10 % aller Teilnehmer zufällig ausgewählt. Von jedem ausgewählten Teilnehmer wird eine der acht Entscheidungen zufällig bestimmt und genau wie unten beschrieben umgesetzt (alle unten erwähnten Personen existieren wirklich und alle Auszahlungen werden wie beschrieben getätigt).
</p>

<p>In dieser Umfrage gibt es keine richtigen oder falschen Antworten. Bitte entscheiden Sie daher immer gemäß Ihren persönlichen Ansichten. Sie werden dabei vollständig anonym bleiben.
</p>

				 <br>
				 <br>
				`,
		}

		{
			inp := gr.AddInput()
			inp.Type = "button"
			inp.Name = "submitBtn"
			inp.Response = "1"
			inp.Label = trl.S{"de": "Weiter"}
			inp.AccessKey = "n"
		}

	}

	// page 1
	{
		page := q.AddPage()
		page.Label = trl.S{"de": "TEIL 1a"}
		page.Short = trl.S{"de": "TEIL 1a"}
		page.Width = 60

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{"de": `
				<p><b>
				Im Folgenden geht es um eine Spende von 30 €, die <i>eine</i> dieser drei Stiftungen erhalten soll:
				</b></p>

				<br>

				<style>
					table.drei-stiftungen td {
						text-align: center;
						text-align: center;

					}
				</style>

				<table class="drei-stiftungen">
				<tr>
					<td style="width: 33%">Politisch links</td>
					<td style="width: 33%">Politische Mitte</td>
					<td style="width: 33%">Politisch konservativ</td>
				<tr>
				<tr>
					<td style="vertical-align: top;" ><b>Hans-Böckler-Stiftung</b></td>
					<td style="vertical-align: top;" ><b>Bund der Steuerzahler Deutschland e.V.</b></td>
					<td style="vertical-align: top;" ><b>Ludwig-Erhard-Stiftung</b></td>
				<tr>
				</table>

				<p>
				Für jede Ihrer ersten sechs Entscheidungen zeigen wir Ihnen die Präferenzen fünf deutscher Staatsangehöriger darüber, welche der drei Stiftungen die Spende erhalten soll.  Sie entscheiden, wie die Präferenzen der fünf Personen in eine gemeinsame Entscheidung zusammengefasst werden.
				</p>

				<p>
				Die Präferenzen stammen von fünf Personen, die an einer Vorstudie teilgenommen haben<sup>1)</sup>.  Diese fünf Personen wurden aus einer Stichprobe gezogen, in der sich gleich viele Personen politisch links, in der Mitte oder als konservativ einordnen. Jede Person wurde einzeln befragt, welche der drei Stiftungen sie als am besten, mittel, und am schlechtesten erachtet. Den Personen wurde mitgeteilt, dass ihre Präferenzen zusammen mit den Präferenzen von vier anderen Personen an einen zukünftigen Teilnehmer der Studie gegeben werden, der die Präferenzen in eine Entscheidung zusammenfasst.
				</p>
				`}

			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{"de": fmt.Sprintf(`
				In <b>Entscheidung 1</b> wird Stiftung A von zwei Personen 
				
				(<img src='%v' style='display: inline-block; height: 1.0rem;
					position: relative; top: 0.2rem; left: 0.1rem;
				'> ) 
				
				mittel eingestuft und von drei weiteren am schlechtesten. Stiftung B wird von drei Personen am besten eingestuft und von zweien am schlechtesten, und so weiter.
				`, cfg.Pref("/img/person.png")),
				}
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			gr.RandomizationGroup = 1 - 1
			{
				inp := gr.AddInput()
				inp.Type = "composit"
				inp.DynamicFunc = "PoliticalFoundations__0__0"
			}
			_, inputNames, _ := qst.PoliticalFoundations(nil, 0, 0, userID)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "composit-scalar"
				inp.Name = inpName + "_page0"
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				inp.Desc = trl.S{"de": `
				Die Stiftungen wurden anonymisiert und in eine zufällige Reihenfolge gebracht, so dass Sie nicht wissen, um welche Stiftung es sich bei den Stiftungen A, B und C handelt. Sie entscheiden also nicht darüber, welche Stiftung die 30 € erhält. Stattdessen entscheiden Sie, wie die Präferenzen der Gruppenmitglieder in eine Entscheidung zusammengefasst werden und ob Sie beispielsweise eher eine Kompromisslösung oder eher eine Mehrheitslösung für Ihre Gruppe bevorzugen.
				`}
			}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text03"
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
	{

		page := q.AddPage()
		page.Label = trl.S{"de": "TEIL 1b"}
		page.Short = trl.S{"de": "TEIL 1b"}
		page.Width = 60

		{
			gr := page.AddGroup()
			gr.Cols = 2
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.ColSpanLabel = 2
				inp.Type = "textblock"
				inp.Name = "text01"
				inp.Desc = trl.S{"de": `
				Entscheiden Sie im Folgenden, an welche Stiftung das Geld gehen soll. Setzen Sie dazu bei der entsprechenden Stiftung ein Kreuz in der Spalte „Auswahl“. Falls Sie eine zweite oder dritte Alternative als genauso gut empfinden, setzen Sie ein Kreuz in der Spalte „Gleich gut“. Berücksichtigen Sie die dargestellten Präferenzen der Gruppenmitglieder. 
				`}
			}
		}

		// loop over matrix questions
		for i := 0; i < 6; i++ {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 2
				gr.RandomizationGroup = 1 - 1

				{
					inp := gr.AddInput()
					inp.Type = "composit"
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundations__%v__%v", i, i)
				}
				_, inputNames, _ := qst.PoliticalFoundations(nil, i, i, userID)
				for _, inpName := range inputNames {
					inp := gr.AddInput()
					inp.Type = "composit-scalar"
					inp.Name = inpName
				}
			}
		}

	}

	// page 3
	{
		page := q.AddPage()
		// p.Section = trl.S{"de": "TEIL 1"}
		page.Label = trl.S{"de": "TEIL 1"}
		page.Width = 60

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{
					"de": `In Teil 1 sind sie einem/r deutschen Staatsangehörigen
					im Alter von 18 bis 30 zugeordnet
					der-/die ebenfalls an dieser Studie teilnimmt.
					Wir bitten Sie zu entscheiden, welche Entscheidungsoptionen dieser Person
					zur Verfügung stehen. Die Person wird frei wählen können zwischen allen Optionen,
					die Sie zur Verfügung stellen;
					die anderen Optionen kann die Person nicht wählen.
					In der letzten Spalte können Sie darüber hinaus von Optionen abraten.`,
				}
			}
		}

		{
			gr := page.AddRadioMatrixGroup(labelsVerfuegbarNicht, q1Names, q1Labels, 3)
			gr.RandomizationGroup = 2 - 2
			gr.BottomVSpacers = 1
			gr.Cols = 6
			gr.Width = 90

			gr.Label = trl.S{
				"de": "Frage [groupID]<br>",
			}
			gr.Desc = trl.S{
				"de": `Wir werden eine Geldzahlung an die Person auslösen.
				Die Person kann zwischen den von Ihnen verfügbar gemachten Optionen
				für Zahlungen an zwei unterschiedlichen Zeitpunkten
				(sofort und 6 Monate nach der Studie) wählen.`,
			}
		}

		{
			gr := page.AddRadioMatrixGroup(labelsVerfuegbarNicht, q2Names, q2Labels, 3)
			gr.RandomizationGroup = 2 - 2
			gr.BottomVSpacers = 1
			gr.Cols = 6
			gr.Width = 90
			gr.Label = trl.S{"de": "Frage [groupID]"}
			gr.Desc = trl.S{"de": "&nbsp;"}
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text04"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{
					"de": `In den nächsten zwei Entscheidungen bitten wir Sie,
					vorherzusagen, wie sich die Staatsangehörigen entscheiden werden:`,
				}
			}
		}

		{

			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text05"
				inp.Label = trl.S{"de": "Frage 3<br>"}
				inp.Desc = trl.S{
					"de": `

						Stellen Sie sich eine repräsentative Gruppe von 10 Personen vor,
						die sich frei zwischen den drei folgenden Optionen entscheiden.

<br>

<style>
    .b1 {
        display: inline-block;
        margin: 0.7rem;
        width: 10.4rem;
        border: 1px solid grey;
    }
    .b2, .b3 {
        padding: 0.2rem;
    }
    .b2 {
        border-bottom: 1px solid grey;
    }

</style>


<div class="b1">
    <div class="b2">
        Option A
    </div>
    <div class="b3">
         <b>0</b> € in <b>1</b> Monat und<br>
        <b>15</b> € in <b>7</b> Monaten
    </div>
</div>


<div class="b1">
    <div class="b2">
        Option B
    </div>

    <div class="b3">
        <b>3</b> € in <b>1</b> Monat und<br>
        <b>7</b> € in <b>7</b> Monaten
    </div>
</div>


<div class="b1">
    <div class="b2">
        Option C
    </div>

    <div class="b3">
        <b>4</b> € in <b>1</b> Monat und<br>
        <b>1</b> € in <b>7</b> Monaten
    </div>
</div>






					`,
				}
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 4
			gr.Width = 55
			gr.BottomVSpacers = 2

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q3_a"
				inp.MaxChars = 2
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				inp.Desc = trl.S{"de": "Wie viele wählen Option A? Ihre Antwort:"}
				inp.Suffix = trl.S{"de": " &nbsp; von 10"}
				inp.Validator = "inRange20"
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q3_b"
				inp.MaxChars = 2
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				inp.Desc = trl.S{"de": "Wie viele wählen Option B? Ihre Antwort:"}
				inp.Suffix = trl.S{"de": " &nbsp; von 10"}
				inp.Validator = "inRange20"
			}
			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = "q3_c"
				inp.MaxChars = 2
				inp.ColSpanLabel = 3
				inp.ColSpanControl = 1
				inp.Desc = trl.S{"de": "Wie viele wählen Option C? Ihre Antwort:"}
				inp.Suffix = trl.S{"de": " &nbsp; von 10"}
				inp.Validator = "inRange20"
			}
		}

		{
			gr := page.AddRadioGroupVertical("q4", q4Labels, 1)
			gr.Cols = 1
			gr.Width = 90
			gr.Label = trl.S{
				"de": "Frage 4<br>",
			}
			gr.Desc = trl.S{
				"de": `Sollten alle Erwerbstätigen in Deutschland verpflichtend
				einen gewissen Anteil Teil Ihres Arbeitseinkommens
				für die private Altersvorsorge sparen,
				und falls ja, wieviel?`,
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text03"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{
					"de": `<sup>1)</sup>
					Falls Sie hier ankreuzen werden Studienteilnehmende folgende Botschaft sehen:
					“Ein gewählter Volksvertreter oder eine gewählte Volksvertreterin der
					oder die an dieser Studie teilgenommen hat, rät Ihnen davon ab,
					diese Alternative zu wählen”.
					`,
				}
			}
		}

	}

	// page 4
	{
		page := q.AddPage()
		// p.Section = trl.S{"de": "TEIL 3"}
		page.Label = trl.S{"de": "TEIL 3"}
		page.Width = 55

		{
			gr := page.AddRadioMatrixGroupNoLabels(labelsOneToTen1, []string{"q11_time_pref"})
			gr.RandomizationGroup = 1 - 1
			gr.BottomVSpacers = 2
			gr.Cols = 11
			gr.Width = 100

			gr.Label = trl.S{
				"de": "Frage [groupID]<br>",
			}
			gr.Desc = trl.S{
				"de": `
				Sind Sie im Vergleich zu Anderen im Allgemeinen bereit,
				heute auf etwas zu verzichten,
				um in der Zukunft davon zu profitieren,
				oder sind Sie im Vergleich zu Anderen dazu nicht bereit?
				Bitte klicken Sie ein Kästchen auf der Skala an.
				`,
			}
		}

		{
			gr := page.AddRadioMatrixGroupNoLabels(labelsOneToTen2, []string{"q12_risk_pref"})
			gr.RandomizationGroup = 1 - 1
			gr.BottomVSpacers = 2
			gr.Cols = 11
			gr.Width = 100

			gr.Label = trl.S{
				"de": "Frage [groupID]<br>",
			}
			gr.Desc = trl.S{
				"de": `
				Wie schätzen Sie sich persönlich ein?
				Sind Sie im Allgemeinen ein risikobereiter Mensch oder versuchen Sie,
				Risiken zu vermeiden?
				Bitte klicken Sie ein Kästchen auf der Skala an.
				`,
			}
		}

		{
			gr := page.AddRadioMatrixGroupNoLabels(labelsOneToTen3, []string{"q13_sharing"})
			gr.RandomizationGroup = 1 - 1
			gr.BottomVSpacers = 3
			gr.Cols = 11
			gr.Width = 100

			gr.Label = trl.S{
				"de": "Frage [groupID]<br>",
			}
			gr.Desc = trl.S{
				"de": `
				Wie schätzen Sie Ihre Bereitschaft mit anderen zu teilen,
				ohne dafür eine Gegenleistung zu erwarten?
				Bitte klicken Sie ein Kästchen auf der Skala an.
				`,
			}
		}

		//
		// explicit button to finish page, which is outsite navigation
		{
			gr := page.AddGroup()
			gr.BottomVSpacers = 2
			gr.Cols = 1
			gr.Width = 100

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "finished"
				inp.Name = "submitBtn"
				inp.CSSControl = "special-line-height-higher"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = trl.S{"de": "", "en": ""}
				inp.Desc = cfg.Get().Mp["end"]
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.HAlignControl = qst.HCenter
				inp.HAlignControl = qst.HLeft
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
		p := q.AddPage()
		p.Label = cfg.Get().Mp["end"]
		p.NoNavigation = true
		{
			// Only one group => shuffling is no problem
			gr := p.AddGroup()
			gr.Cols = 1

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.CSSLabel = "special-line-height-higher"
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
				inp.CSSLabel = "special-line-height-higher"
				inp.Desc = trl.S{
					"de": "<span style='font-size: 100%;'>Ihre Eingaben wurden gespeichert.</span>",
					"en": "<span style='font-size: 100%;'>Your entries have been saved.</span>",
					"es": "<span style='font-size: 100%;'>Sus entradas se han guardado.</span>",
					"fr": "<span style='font-size: 100%;'>Vos réponses ont été sauvegardées.</span>",
					"it": "<span style='font-size: 100%;'>Le Sue risposte sono state salvate.</span>",
					"pl": "<span style='font-size: 100%;'>Twoje wpisy zostały zapisane.</span>",
				}
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
