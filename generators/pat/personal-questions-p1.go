package pat

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

func PersonalQuestions1(q *qst.QuestionnaireT) error {

	grStPage78 := css.NewStylesResponsive(nil)
	grStPage78.Desktop.StyleGridContainer.GapRow = "0.1rem"
	grStPage78.Desktop.StyleGridContainer.GapColumn = "0.01rem"

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
					<b>Frage&nbsp;7.</b>
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

	}

	return nil
}
