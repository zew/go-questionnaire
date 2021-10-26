package fmt

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202111(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 11 {
		return nil
	}

	lblStyleRight := css.NewStylesResponsive(nil)
	lblStyleRight.Desktop.StyleText.AlignHorizontal = "right"
	lblStyleRight.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
	lblStyleRight.Mobile.StyleBox.Padding = " 0 0.2rem 0 0"

	/*
		SELECT
			frage_kurz,
			GROUP_CONCAT( DISTINCT antwort ORDER BY antwort) aw,
			count(*) anz
		FROM sonderfragen_ger
		WHERE survey_id = 202005
		GROUP BY frage_kurz
	*/

	page := q.AddPage()
	// page.Section = trl.S{"de": "Sonderfrage", "en": "Special"}
	page.Label = trl.S{
		"de": "Sonderfrage: Inflation, Prognosetreiber und Geldpolitik",
		"en": "Special: Inflation, forecast drivers and monetary policy",
	}
	page.Short = trl.S{
		"de": "Inflation,<br>Geldpolitik",
		"en": "Inflation,<br>Mon. Policy",
	}
	page.Style = css.DesktopWidthMaxForPages(page.Style, "48rem")

	{
		gr := page.AddGroup()
		gr.Cols = 9
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Width = "70%"
		gr.Style.Mobile.StyleBox.Width = "100%"

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 9
			// inp.ColSpanLabel = 12
			inp.Label = trl.S{
				"de": `
					<p>
						Angesichts der Inflationsentwicklung der letzten
						 Monate möchten wir Ihre Einschätzungen 
						 zu den Ursachen und der weiteren Entwicklung 
						 der Inflation im Eurogebiet noch detaillierter 
						 als üblicherweise erfragen
					</p>
				`,
				"en": `
					<p>
						todo english
					</p>
				`,
			}
		}

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 9
			// inp.ColSpanLabel = 12
			inp.Label = trl.S{
				"de": `<b>1.</b> Punktprognose der <b>jährlichen Inflationsrate im Euroraum</b>
				<br>
				Anstieg des HICP von Jan bis Dez; Erwartungswert
				`,
				"en": `<b>1.</b> Forecast <b>yearly inflation rate in the Euro area</b>
				<br>
				HICP  increase from Jan to Dec; expected value
				`,
			}
		}

		for idx := range []int{0, 1, 2} {

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("ppjinf_jp%v", idx) //"p1_y1"
			inp.Min = -10
			inp.Max = +20
			inp.Validator = "inRange20"
			inp.MaxChars = 5
			inp.Step = 0.01
			inp.Label = trl.S{
				"de": q.Survey.YearStr(idx),
				"en": q.Survey.YearStr(idx),
			}
			inp.Suffix = trl.S{
				"de": "%",
				"en": "pct",
			}

			inp.ColSpan = 3
			inp.ColSpanLabel = 2
			inp.ColSpanControl = 2

			inp.StyleLbl = lblStyleRight
		}

	}

	return nil
}
