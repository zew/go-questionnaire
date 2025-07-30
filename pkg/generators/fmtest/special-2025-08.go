package fmtest

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202508(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2025 && q.Survey.Month == 8
	if !cond {
		return nil
	}

	page := q.AddPage()
	// pge.Section = trl.S{"de": "Sonderfrage", "en": "Special"}

	page.WidthMax("72rem")
	page.WidthMax("64rem")

	page.Label = trl.S{
		"de": "Sonderfrage: Bewertung des EU-US Handelsabkommens",
		"en": "Special question: New EU-USA trade agreement",
	}
	page.Short = trl.S{
		"de": "EU-US.Handels-<br>abkommen",
		"en": "EU-USA trade<br>agreement",
	}
	// page.WidthMax("42rem")

	//
	// gr2
	rowLabelsEconomicAreas := []trl.S{
		{
			"de": "Reales BIP-Wachstum in Deutschland",
			"en": "German real GDP growth",
		},
		{
			"de": "Deutsche Inflationsrate",
			"en": "German inflation rate",
		},
		{
			"de": "Reales BIP-Wachstum in der Eurozone",
			"en": "Euro area real GDP growth",
		},
		{
			"de": "Inflationsrate in der Eurozone",
			"en": "Euro area inflation rate",
		},
		{
			"de": "EZB Leitzinsen",
			"en": "ECB monetary policy rates",
		},
		{
			"de": "Reales BIP-Wachstum in den USA",
			"en": "US real GDP growth",
		},
		{
			"de": "Inflationsrate in den USA",
			"en": "US inflation rate",
		},
		{
			"de": "FED Leitzinsen",
			"en": "FED monetary policy rates",
		},
	}

	colTemplate, _, _ := colTemplateWithFreeRow()

	{
		gb := qst.NewGridBuilderRadios(
			colTemplate,
			labels202508(),
			[]string{
				"sq3_eu_us_agree_ger_w",
				"sq3_eu_us_agree_ger_pi",
				"sq3_eu_us_agree_ea_w",
				"sq3_eu_us_agree_ea_pi",
				"sq3_eu_us_agree_ecb_i",
				"sq3_eu_us_agree_us_w",
				"sq3_eu_us_agree_us_pi",
				"sq3_eu_us_agree_fed_i",
			},
			radioVals6,
			rowLabelsEconomicAreas,
		)

		gb.MainLabel = trl.S{
			"de": `
				Die Europäische Union und die Vereinigten Staaten haben am 27. Juli ein neues Handelsabkommen angekündigt. Wie erwarten Sie, dass sich dieses Abkommen in den kommenden 12 Monaten auf die folgenden Variablen auswirkt?
			`,
			"en": `
				The European Union and the United States announced a new trade agreement on July 27, 2025. How do you expect this agreement to affect the following variables over the next 12 months?
			`,
		}
		gr := page.AddGrid(gb)
		gr.BottomVSpacers = 3
	}

	return nil

}
