package pat2

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// TitlePat23 adds title page
func TitlePat23(q *qst.QuestionnaireT) error {

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

				<p>
				Dies ist eine Studie des Zentrums für Europäische Wirtschaftsforschung (ZEW) 
				in Mannheim sowie der Universitäten Köln, Mannheim, Münster und Zürich. 

				<br>
				<br>
				
				Ihre Teilnahme wird ca. 15&nbsp;Minuten in Anspruch nehmen 
				und Sie unterstützen damit die Forschung 
				zu Entscheidungsprozessen in der Politik.
				</p>

				<p>
				Sie werden diese Umfrage nur abschliessen können, wenn Sie einen Verständnistest richtig beantworten. Bei mehrmaligen falschen Antworten wird die Umfrage automatisch terminiert. Bitte lesen Sie die Anleitungen daher sehr genau. 
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

	return nil
}
