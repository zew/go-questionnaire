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
			"de": `Wir sind Wissenschaftler des Zentrums für Europäische Wirtschaftsforschung (ZEW) in Mannheim 
				sowie der Universitäten in Köln, Mannheim, Münster und Zürich. 
				Ihre Teilnahme an unserer Umfrage, die aus drei Teilen besteht, 
				wird nur wenige Minuten in Anspruch nehmen. 
				Sie unterstützen damit die Wissenschaft.
				Es gibt keine richtigen oder falschen Antworten. 
				Bitte entscheiden Sie immer gemäß Ihren persönlichen Ansichten. 
				Teilweise werden Ihre Entscheidungen andere Personen betreffen. 
				Bitte beachten Sie, dass wir eine dieser Entscheidungen zufällig auswählen
				 und genauso umsetzen werden, wie hier beschrieben. 
				 Sie werden also mit Ihren Antworten tatsächlich auf andere Personen 
				 Einfluss ausüben. 
				 <b>Gleichzeitig werden alle Ihre Antworten vollständig anonym bleiben</b>.
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
		// page.Section = trl.S{"de": "TEIL 1"}
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
			gr.RandomizationGroup = 2 - 1
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
			gr.RandomizationGroup = 2 - 1
			gr.BottomVSpacers = 1
			gr.Cols = 6
			gr.Width = 90
			gr.Label = trl.S{"de": "Frage [groupID]"}
			gr.Desc = trl.S{"de": "&nbsp;"}
		}

		/* 		page := q.AddPage()
		   		page.Label = trl.S{"de": "TEIL 1 - Sektion B"}
		   		page.Label = trl.S{"de": "Sektion B"}
		   		page.Width = 65
		*/
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
         <b>0</b> Euro in <b>1</b> Monat und<br>
        <b>15</b> Euro in <b>7</b> Monaten
    </div>
</div>


<div class="b1">
    <div class="b2">
        Option B
    </div>

    <div class="b3">
        <b>3</b> Euro in <b>1</b> Monat und<br>
        <b>7</b> Euro in <b>7</b> Monaten
    </div>
</div>


<div class="b1">
    <div class="b2">
        Option C
    </div>

    <div class="b3">
        <b>4</b> Euro in <b>1</b> Monat und<br>
        <b>1</b> Euro in <b>7</b> Monaten
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

	//  page 2
	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "TEIL 2"}
		page.Label = trl.S{"de": "TEIL 2a"}
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
				inp.Desc = trl.S{"de": `Jede der sechs Bundestagsparteien ist mit einer parteinahen Stiftung verbunden.<sup>2</sup>  
				In diesem Teil der Studie geht es um eine Spende in Höhe von 30 EURO, 
				welche <i>eine</i> dieser Stiftungen erhalten wird. 
				Wir haben fünf deutsche Staatsangehörige befragt, 
				welche der Stiftungen die Spende erhalten soll. 
				Sie entscheiden nun darüber, 
				auf welche Weise die Wünsche der fünf Befragten 
				in eine finale Entscheidung zusammengefasst werden sollen.`}
			}
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				// inp.Label = trl.S{"de": "Dummy<br>"}
				inp.Desc = trl.S{"de": `Dafür zeigt Ihnen jede der nächsten sechs Fragen die möglichen Wünsche von fünf Befragten. 
				Die Wünsche in <i>einer</i> dieser Fragen stammen von einer Gruppe, 
				die wir tatsächlich bereits befragt haben. 
				Wir werden die Entscheidung ausführen, die Sie in dieser Frage treffen. 
				Allerdings wissen Sie nicht, welche der Fragen das ist. 
				Die fünf Befragten wurden aus einer Stichprobe gezogen, 
				in welcher es jeweils gleich viele Personen gab, 
				die politisch rechts, mittig, oder links stehen.
				
				<br>
				<br>
				
				<b>Entscheidung 1</b>
				<br>
				<br>

				Die Stiftungen sind anonymisiert; die Auswahl von drei Stiftungen 
				aus den sechs Stiftungen erfolgte zufällig. 
				Sie können nun entscheiden, ob Sie beispielsweise eher eine Kompromisslösung 
				oder eher eine kontroverse Mehrheitslösung implementieren, 
				aber sie können nicht angeben, ob Sie eine bestimmte Stiftung bevorzugen. 

				<br>
				<br>
				Die Wünsche der fünf Befragten werden wie folgt grafisch dargestellt. 


				`}
			}
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "checkbox"
			// 	inp.Name = "checktest"
			// }
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "radiogroup"
			// 	inp.Name = "radiotest"
			// 	rd1 := inp.AddRadio()
			// 	rd2 := inp.AddRadio()
			// 	_, _ = rd1, rd2
			// }
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			gr.RandomizationGroup = 1 - 1
			{
				inp := gr.AddInput()
				inp.Type = "composit"
				inp.DynamicFunc = "PoliticalFoundations__0"
			}
			_, inputNames, _ := qst.PoliticalFoundations(nil, -1, 0)
			for _, inpName := range inputNames {
				inp := gr.AddInput()
				inp.Type = "composit-scalar"
				inp.Name = inpName
			}

		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				inp.Desc = trl.S{"de": `
				
				In diesem Fall mögen zwei Befragte Stiftung A am meisten, drei weitere mögen Stiftung C am meisten, und niemand der fünf mag Stiftung B am meisten. Dafür mögen zwei Staatsangehörige Stiftung B am zweitmeisten, und niemand mag Stiftung C am zweitmeisten. Und so weiter.  				
				<br>
				<br>
				
				Entscheiden Sie bitte nun, welche Stiftung die Spende erhalten soll. Dafür setzen Sie bitte ein Kreuz in die Spalte „Auswahl“ bei jener Alternative, von der Sie denken, dass sie die Wünsche der Befragten am besten reflektiert. Falls sie eine zweite Alternative genauso gut finden wie die ausgewählte Alternative, setzen Sie bitte ein Kreuz in die Spalte „Genauso gut“. 
				<br>

				<span style='display: inline-block; margin: 2rem;'>
				<i>
					Es gibt keine richtigen oder falschen Antworten; teilen Sie uns bitte einfach mit, welche Stiftung angesichts der Wünsche der Befragten die Spende erhalten soll. 
				</i>
				</span>


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
					"de": `<sup>2)</sup> 
					Die Stiftungen und zugehörige Parteien sind: Friedrich-Ebert-Stiftung (SPD, 1954), 
					Konrad-Adenauer-Stiftung (CDU, 1955), 
					Friedrich-Naumann-Stiftung für die Freiheit (FDP, 1958), 
					Rosa-Luxemburg-Stiftung (Die Linke, 1990), 
					Heinrich-Böll-Stiftung (Grüne, 1996), 
					Desiderius-Erasmus-Stiftung (AfD, 2017) 
					(Liste geordnet nach Gründungsjahr).
					`,
				}
			}

		}

	}

	{

		page := q.AddPage()
		// page.Section = trl.S{"de": "TEIL 2"}
		page.Label = trl.S{"de": "TEIL 2b"}
		page.Width = 60

		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Name = "text01"
				inp.Desc = trl.S{"de": `
				
				Bitte entscheiden Sie nun auch für die folgenden Fälle, welche Stiftung die Spende erhalten soll. Eine der Entscheidungen 1 – 6 zeigt die echten Wünsche der fünf Befragten; wir werden die entsprechende Entscheidung ausführen. 

				`}
			}
		}

		// loop over matrix questions
		for i := 1; i < 6; i++ {
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 3
				gr.RandomizationGroup = 1 - 1

				{
					inp := gr.AddInput()
					inp.Type = "composit"
					inp.DynamicFunc = fmt.Sprintf("PoliticalFoundations__%v", i)
				}
				_, inputNames, _ := qst.PoliticalFoundations(nil, -1, i)
				for _, inpName := range inputNames {
					inp := gr.AddInput()
					inp.Type = "composit-scalar"
					inp.Name = inpName
				}
			}
		}

		if false {
			// DUMMMY  DUMMMY  DUMMMY
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.RandomizationGroup = 2
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Name = "text03"
					inp.Desc = trl.S{"de": `sg1 - el1`}
				}
			}
			// DUMMMY  DUMMMY  DUMMMY
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.RandomizationGroup = 0
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Name = "text03"
					inp.Desc = trl.S{"de": `no shuffle`}
				}
			}
			// DUMMMY  DUMMMY  DUMMMY
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.RandomizationGroup = 2
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Name = "text03"
					inp.Desc = trl.S{"de": `sg1 - el2`}
				}
			}
			// DUMMMY  DUMMMY  DUMMMY
			{
				gr := page.AddGroup()
				gr.Cols = 1
				gr.RandomizationGroup = 2
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Name = "text03"
					inp.Desc = trl.S{"de": `sg1 - el3`}
				}
			}
		}

	}

	// page 3
	{
		page := q.AddPage()
		// page.Section = trl.S{"de": "TEIL 3"}
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
