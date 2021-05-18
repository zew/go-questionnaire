package pat2

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Title adds title page
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

				<p>Dies ist eine Studie des Zentrums für Europäische Wirtschaftsforschung (ZEW) 
				in Mannheim sowie der Universitäten in Köln, Mannheim, Münster und Zürich. 
				
				Ihre Teilnahme ca. 15&nbsp;Minuten in Anspruch nehmen 
				und Sie unterstützen damit die Forschung.
				</p>

				<p> <b> Sie werden vollständig anonym bleiben. </b>
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
