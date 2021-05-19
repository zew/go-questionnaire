package pat

import (
	"fmt"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// VariableElements is bag of options to questionnaire modules
type VariableElements struct {
	// Part 1 and 2
	NumberingQuestions int
	NumberingSections  int

	AllMandatory    bool
	ZumSchlussOrNun bool

	// Part 2
	ZumXtenTeil string
}

// PersonalQuestions1 - numbered 5-7
func PersonalQuestions1(q *qst.QuestionnaireT, vE VariableElements) error {

	grStPage78 := css.NewStylesResponsive(nil)
	grStPage78.Desktop.StyleGridContainer.GapRow = "0.1rem"
	grStPage78.Desktop.StyleGridContainer.GapColumn = "0.01rem"

	validatorRadio := ""
	if vE.AllMandatory {
		validatorRadio = "mustRadioGroup"
	}

	zumSchlussOrNun := "Zum Schluss bitten wir Sie, drei Fragen über sich selbst zu beantworten: "
	if vE.ZumSchlussOrNun {
		zumSchlussOrNun = "Nun bitten wir Sie, einige Fragen über sich selbst zu beantworten: "
	}

	// page 8
	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Short = trl.S{"de": "Eigene Einstellung 2"}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "30rem")

		// gr1
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate7,
				labelsOneToSeven2,
				[]string{"q5"},
				radioVals7,
				[]trl.S{},
				validatorRadio,
			)
			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
					<p>
					<b>%v</b>

					<br>
					<br>
					<b>Frage %v.</b>
					Sind Sie im Vergleich zu anderen im Allgemeinen bereit, 
					heute auf etwas zu verzichten, 
					um in der Zukunft davon zu profitieren, 
					oder sind Sie im Vergleich zu anderen dazu nicht bereit? 

					</p>

				`, zumSchlussOrNun,
					vE.NumberingQuestions+0,
				),
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
		}

		// gr2
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate7,
				labelsOneToSeven3,
				[]string{"q6"},
				radioVals7,
				[]trl.S{},
				validatorRadio,
			)
			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
					</p>
					<b>Frage %v.</b>
					Wie schätzen Sie sich persönlich ein? 
					Sind Sie im Allgemeinen ein risikobereiter Mensch 
					oder versuchen Sie, Risiken zu vermeiden?
					</p>
				`, vE.NumberingQuestions+1),
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
		}

		// gr3
		{
			gb := qst.NewGridBuilderRadiosWithValidator(
				columnTemplate7,
				labelsOneToSeven2,
				[]string{"q7"},
				radioVals7,
				[]trl.S{},
				validatorRadio,
			)
			gb.MainLabel = trl.S{
				"de": fmt.Sprintf(`
					<p>
					<b>Frage&nbsp;%v.</b>
					Wie schätzen Sie Ihre Bereitschaft ein, mit anderen zu teilen, 
					ohne dafür eine Gegenleistung zu erwarten?
					</p>
				`, vE.NumberingQuestions+2),
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
			gr.Style = grStPage78
			gr.BottomVSpacers = 4
		}

	}

	return nil
}
