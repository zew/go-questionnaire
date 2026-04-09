package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func rowsBasisPoints(
	page *qst.WrappedPageT,
	main trl.S,
	inps []string, // input base for each row
	colHeaders []trl.S,
	rowLabels1 []trl.S,
) {

	gr := page.AddGroup()
	gr.BottomVSpacers = 3
	gr.Cols = 2
	gr.ColWidths("2.1fr  4.9fr")

	// gr.WidthMax("33rem")

	// main label
	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols
		inp.Label = main
	}

	// first row
	for _, lbl := range colHeaders {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.ColSpan = gr.Cols / float32(len(colHeaders))
		inp.Label = lbl.Bold()
		inp.Style = css.NewStylesResponsive(inp.Style)
		inp.Style.Desktop.Margin = "1.5ch 0 0 10rem "
		inp.Style.Mobile.Margin = "1ch  0    0 0"
	}

	//
	//
	// second to fourth row: inputs
	for i, row := range rowLabels1 {

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.Label = row
			inp.Style = css.ItemEndMA(inp.Style)
			inp.LabelPadRight()

		}

		//
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("rexp_ecb_%v", inps[i])
			inp.Suffix = trl.S{"de": "Basispunkte", "en": "basis points"}
			inp.ColSpan = 1
			inp.ColSpanControl = 3
			inp.Min = -200
			inp.Max = 200
			inp.Step = 1
			inp.MaxChars = 5
			// inp.Style = css.ItemCenteredMCA(inp.Style)
			// inp.Style = css.ItemEndCA(inp.Style)

		}

	}

}

func special202604(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2026 && q.Survey.Month == 4
	if !cond {
		return nil
	}

	page := q.AddPage()
	page.Label = trl.S{
		"de": "Zinserwartungen",
		"en": "Interest rate expectations",
	}
	page.Label = trl.S{
		"de": "",
		"en": "",
	}
	page.Short = trl.S{
		"de": "Sonderfragen:<br>Zinserwartungen",
		"en": "Special questions:<br>Interest rates",
	}
	page.SuppressInProgressbar = true

	page.WidthMax("55rem")

	{

		inps := []string{
			"26_04",
			"26_06",
			"26_07",
			"26_09",
			"26_10",
			"26_12",
		}

		lblMain := trl.S{
			"de": `
				Wir möchten Sie zu Ihren Erwartungen bezüglich zukünftiger Zinsentscheidungen der Europäischen Zentralbank (EZB) befragen. 
				<br>
				<br>
				Derzeit beträgt der Einlagezinssatz 2,00%.
				<br>
				<br>
				Bitte geben Sie für jedes der unten aufgeführten Treffen des EZB-Rats an, 
				um wie viele Basispunkte sich der <b>Einlagezinssatz</b> gegenüber der jeweils vorherigen Sitzung Ihrer Einschätzung nach ändern wird.
				<br>
				<br>
				<small>
					Hinweis: Bitte verwenden Sie positive Werte für Zinserhöhungen und negative Werte für Zinssenkungen. Wenn Sie keine Änderung erwarten, tragen Sie bitte „0“ ein.
				</small>

			`,
			"en": `
				We would like to ask about your expectations regarding future interest rate decisions of the European Central Bank (ECB).
				<br>
				<br>
				The current deposit facility rate is 2.00%.
				<br>
				<br>
				Please indicate for each of the ECB Governing Council meetings listed below, by how many basis points you expect the <b>deposit facility rate</b> to change relative to the previous meeting.
				<br>
				<br>
				<small>
				Note: Please use positive values for rate increases and negative values for rate cuts. If you expect no change, please enter “0”.
				</small>
			
			`,
		}.Outline("4.")

		headers := []trl.S{
			{
				"de": `Änderung des Einlagezinssatzes`,
				"en": `Change in the deposit facility rate`,
			},
		}

		rowLbls := []trl.S{
			{
				"de": "30.&nbsp;April&nbsp;2026",
				"en": "30&nbsp;April&nbsp;2026",
			},
			{
				"de": "11.&nbsp;Juni&nbsp;2026",
				"en": "11&nbsp;June&nbsp;2026",
			},
			{
				"de": "23.&nbsp;Juli&nbsp;2026",
				"en": "23&nbsp;July&nbsp;2026",
			},
			{
				"de": "10.&nbsp;September&nbsp;2026",
				"en": "10&nbsp;September&nbsp;2026",
			},
			{
				"de": "29.&nbsp;Oktober&nbsp;2026",
				"en": "29&nbsp;October&nbsp;2026",
			},
			{
				"de": "17.&nbsp;Dezember&nbsp;2026",
				"en": "17&nbsp;December&nbsp;2026",
			},
		}

		rowsBasisPoints(
			qst.WrapPageT(page),
			lblMain,
			inps,
			headers,
			rowLbls,
		)
	}

	return nil
}
