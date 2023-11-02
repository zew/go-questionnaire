package fmt

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202311(q *qst.QuestionnaireT, page *qst.WrappedPageT) error {

	cond := false
	cond = cond || q.Survey.Year == 2023 && q.Survey.Month == 11
	if !cond {
		return nil
	}

	lbls := []trl.S{
		{
			"de": "Konjunkturentwicklung im Eurogebiet",
			"en": "Development of GDP in the euro area",
		},
		{
			"de": "Entwicklung der Löhne im Eurogebiet",
			"en": "Development of wages in the euro area",
		},
		{
			"de": "Entwicklung der Energiepreise",
			"en": "Development of energy prices",
		},
		{
			"de": "Entwicklung der Rohstoffpreise (ohne Energiepreise)",
			"en": "Development of prices for raw materials (except energy) ",
		},
		{
			"de": "Veränderung der Wechselkurse (relativ zum Euro)",
			"en": "Changes in exchange rates (relative to the euro)",
		},
		{
			"de": "Geldpolitik der EZB",
			"en": "Monetary policy of the ECB",
		},
		{
			"de": "Internationale Handelskonflikte",
			"en": "International trade conflicts",
		},
		{
			"de": "Internationale Lieferengpässe",
			"en": "International supply bottlenecks",
		},
		{
			"de": "Grüne Transformation",
			"en": "Green transformation",
		},
		{
			"de": "Krieg in der Ukraine",
			"en": "War in Ukraine",
		},
	}

	inps := []string{
		"israel_bus_cycle",
		"israel_wages",
		"israel_energy_prices",
		"israel_commodity_prices",
		"israel_exch_rates",
		"israel_mp_ecb",
		"israel_trade_conflicts",
		"israel_supply_shortages",
		"israel_green_trafo",
		"israel_war_ukraine",
	}

	lbl := trl.S{
		"de": fmt.Sprintf(` 
				Wählen Sie die 
				<i>drei Bereiche</i> 
				aus, auf die der 
				<i>Israel-Konflikt</i> 
				Ihrer Meinung nach auf Sicht von 
				<u>sechs Monaten</u> 
				den stärksten Einfluss haben wird 
				(egal ob positiv oder negativ):
			`,
		),
		"en": fmt.Sprintf(`
				In your opinion, what are the 
				<i>three factors</i> 
				most strongly affected by the 
				<i>conflict in Israel</i> 
				over the next 
				<u>six months</u> 
				(either positive or negative):
			`),
	}.Outline("3.")

	// gr := page.AddGrid(gb)
	gr := page.AddGroup()

	col1 := float32(2)
	col2 := float32(1)
	col3 := float32(3)

	gr.Cols = col1 + col2 + col3
	gr.BottomVSpacers = 4

	gr.Style = css.NewStylesResponsive(gr.Style)
	// gr.Style.Mobile.StyleGridContainer.TemplateColumns = "1fr 1fr 1fr 1fr 1fr 1fr"
	//
	gr.Style.Mobile.StyleGridContainer.TemplateColumns = "5fr 1fr 1fr 0.4fr 0.4fr 0.4fr"

	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = lbl
		inp.ColSpan = gr.Cols
		inp.ColSpanLabel = 1
	}

	for i, inpName := range inps {
		{
			inp := gr.AddInput()
			inp.Type = "checkbox"
			inp.Name = inpName
			inp.MaxChars = 15

			inp.ColSpan = col1 + col2
			inp.ColSpanLabel = col1
			inp.ColSpanControl = col2

			inp.Label = lbls[i]
		}
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			// inp.Label =
			inp.ColSpan = col3
			inp.ColSpanLabel = 1
		}

	}

	//
	// row free input
	{
		inp := gr.AddInput()
		inp.Type = "text"
		inp.Name = "israel_free_label"
		inp.MaxChars = 15
		inp.ColSpan = col1
		inp.ColSpanLabel = 2.4
		inp.ColSpanControl = 4
		inp.Label = trl.S{
			"de": "Andere",
			"en": "Other",
		}
	}
	{
		inp := gr.AddInput()
		inp.Type = "checkbox"
		inp.Name = "israel_free"
		inp.ColSpan = col2
	}

	{
		inp := gr.AddInput()
		inp.Type = "textblock"
		// inp.Label =
		inp.ColSpan = col3
		inp.ColSpanLabel = 1
	}

	//

	{
		inp := gr.AddInput()
		inp.ColSpanControl = 1
		inp.Type = "javascript-block"
		inp.Name = "israel"

		s1 := trl.S{
			"de": "Bitte maximal drei.",
			"en": "Please choose at most three.",
		}
		inp.JSBlockTrls = map[string]trl.S{
			"msg": s1,
		}

		inp.JSBlockStrings = map[string]string{}

		ivls := []string{} // intervals
		for _, name := range inps {
			ivl := fmt.Sprintf("\"%v\"", name)
			ivls = append(ivls, ivl)
		}
		ivl := fmt.Sprintf("\"%v\"", "israel_free")
		ivls = append(ivls, ivl)

		inp.JSBlockStrings["inps"] = "[" + strings.Join(ivls, ", ") + "]"

	}

	return nil
}
