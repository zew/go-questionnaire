package pat2

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

// Part1Intro renders
func Part1Intro(q *qst.QuestionnaireT) error {

	{
		page := q.AddPage()
		page.Label = trl.S{"de": ""}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "36rem") // 60

		//
		gr := page.AddGroup()
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Desc = trl.S{
				"de": `
				<h3>Teil 1</h3>

				<p>
				In diesem Teil der Studie treffen Sie acht Entscheidungen 
				(und beantworten einige Fragen). 
				Nach der Erhebung werden 10&nbsp;% aller Teilnehmer 
				zufällig ausgewählt. Von jedem ausgewählten Teilnehmer wird eine der acht Entscheidungen zufällig bestimmt und genau wie unten beschrieben umgesetzt (alle unten erwähnten Personen existieren wirklich und alle Auszahlungen werden wie beschrieben getätigt).				
				</p>

				<br>
				<br>
				`,
			}
		}

	}

	return nil
}
