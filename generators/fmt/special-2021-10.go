package fmt

import (
	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/trl"
)

var inputNamesGDPYears202110 = []string{
	"gdp_6m",
	"gdp_2022",
	"gdp_2023",
	"gdp_5yrs",
}

var rowLabelsGDPYears202110 = []trl.S{
	{
		"de": "Erwartetes BIP-Wachstum auf Sicht von 6&nbsp;Monaten",
		"en": "Expected GDP growth for next 6&nbsp;months",
	},
	{
		"de": "Erwartetes BIP-Wachstum für&nbsp;2022",
		"en": "Expected GDP growth for&nbsp;2022",
	},
	{
		"de": "Erwartetes BIP-Wachstum für&nbsp;2023",
		"en": "Expected GDP growth for&nbsp;2023",
	},
	{
		"de": "Erwartetes BIP-Wachstum für die nächsten 5&nbsp;Jahre",
		"en": "Expected GDP growth for the next 5&nbsp;years",
	},
}

var inputNamesDAXrestruct202110 = []string{
	"dax_restruct_6m",
	"dax_restruct_5yrs",
}

var rowLabelsDAXrestruct202110 = []trl.S{
	{
		"de": "Auf Sicht von 6&nbsp;Monaten",
		"en": "For the next 6&nbsp;months",
	},
	{
		"de": "Auf Sicht von 5&nbsp;Jahren",
		"en": "For the next 5&nbsp;years",
	},
}

func special202110(q *qst.QuestionnaireT) error {

	if q.Survey.Year != 2021 || q.Survey.Month != 10 {
		return nil
	}

	//
	//
	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfrage 1 zum Ausgang der Bundestagswahl",
			"en": "Special question 1: Elections in Germany",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Bundestagswahl",
			"en": "Special:<br>Elections",
		}
		page.Short = trl.S{
			"de": "Sonderfragen",
			"en": "Specials",
		}
		page.Style = css.DesktopWidthMaxForPages(page.Style, "46rem")

		//
		// gr1
		{
			var columnTemplateLocal = []float32{
				3.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				raiseDecrease6b(),
				inputNamesGDPYears202110,
				radioVals6,
				rowLabelsGDPYears202110,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style=''>
					Der Ausgang der Bundestagswahl hat sich 
					folgendermaßen auf meine kurz- und mittelfristigen 
					Prognosen des realen Bruttoinlandsproduktes (BIP) 
					ausgewirkt:
				</p>

				`,
				"en": `
				<p style=''>
					The results of the German federal elections
					have influenced my short to midterm outlook
					for the real GDP as follows:
				</p>
				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

		//
		// gr2
		{
			var columnTemplateLocal = []float32{
				3.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.0, 1,
				0.5, 1,
			}
			gb := qst.NewGridBuilderRadios(
				columnTemplateLocal,
				improvedDeteriorated6b(),
				inputNamesDAXrestruct202110,
				radioVals6,
				rowLabelsDAXrestruct202110,
			)

			gb.MainLabel = trl.S{
				"de": `
				<p style='' class='go-quest-page-header'>
					Sonderfrage 2 zur neuen Struktur des DAX
				</p>

				<p style=''>
					Seit dem 20. September 2021 gilt die neue 
					Zusammensetzung des DAX mit jetzt 40 statt 
					zuvor 30 Aktien. 
					Wie wirkt sich die Aufstockung 
					von 30 auf 40 Titeln auf Ihre Einschätzung 
					der zukünftigen Entwicklung des DAX aus? 
					<br>
					<br>
					Gegenüber dem DAX30 wird sich der DAX40 
					folgendermaßen entwickeln: 
				</p>

				`,
				"en": `

				<p style='' class='go-quest-page-header'>
					Special question 2: New structure of DAX stock index
				</p>

				<p style=''>
					Since September 20th the DAX stock index
					has a new composition.
					It now contains 40 instead of 
					30 most capitalized stocks.
					How does the increased scope affect your
					assessment of the future DAX development?
					<br>
					<br>
					Compared to the DAX30, the DAX40
					will develop as follows:
				</p>

				`,
			}

			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	}

	return nil

}
