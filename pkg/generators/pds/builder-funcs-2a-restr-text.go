package pds

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// config restricted text
type configRT struct {
	InputNameP2 string // second token of input name;  [numdeals, volbysegm, ...]
	Chars       int    // input characters max
	LblRow1     trl.S  //
	LblRow2     trl.S  // question in more detail
	Suffix      trl.S  // unit 'deals' or 'million €'

	SubNames    []string // suffixes
	SubLbls     map[string]string
	Placeholder trl.S //
}

var (
	rT1 = configRT{
		InputNameP2: "numdeals",
		Chars:       6,
		LblRow1: trl.S{
			"en": "Total number of new deals",
			"de": "Gesamtzahl neue Abschlüsse",
		},
		LblRow2: trl.S{
			"en": `Please state the number of deals closed in Q4 2022 by market segment: `,
			"de": `Please state the number of deals closed in Q4 2022 by market segment: `,
		},
		Suffix: trl.S{
			"en": "deals",
			"de": "Deals",
		},
		SubNames: []string{"low", "midupper", "other"},
		SubLbls: map[string]string{
			"low":      "Lower Mid-Market (0-15m € EBITDA)",
			"midupper": "Core- and Upper Mid-Market (>15m € EBITDA)",
			"other":    "Other",
		},
		Placeholder: trl.S{
			"en": "#",
			"de": "#",
		},
	}
	rT2 = configRT{
		InputNameP2: "volbysegm",
		Chars:       10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by segment",
			"de": "Gesamtvolumen neuer Abschlüsse nach Marktsegment",
		},
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by market segment: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in Q4 2022 nach Marktsegment: `,
		},
		Suffix: trl.S{
			// capitalizemytitle.com/how-to-abbreviate-million/
			// "en": "million €",
			"en": "MM €",
			"de": "Mio €",
		},
		SubNames: []string{"low", "mid", "upper"},
		SubLbls: map[string]string{
			"low":   "Lower Mid-Market (0-15m € EBITDA)",
			"mid":   "Core Mid-Market  (15-50m € EBITDA)",
			"upper": "Upper Mid-Market (>50m € EBITDA)",
		},
		Placeholder: trl.S{
			"en": "million Euro",
			"de": "Millionen Euro",
		},
	}
	rT3 = configRT{
		InputNameP2: "volbyreg",
		Chars:       10,
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by region: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in Q4 2022 nach Region: `,
		},
		Suffix: trl.S{
			// capitalizemytitle.com/how-to-abbreviate-million/
			// "en": "million €",
			"en": "MM €",
			"de": "Mio €",
		},
		SubNames: []string{"uk", "france", "dach", "benelux", "nordics", "southern_eu", "other"},
		SubLbls: map[string]string{
			"uk":          "UK",
			"france":      "France",
			"dach":        "DACH",
			"benelux":     "Benelux",
			"nordics":     "Nordics",
			"southern_eu": "Southern Europe",
			"other":       "Other",
		},
		Placeholder: trl.S{
			"en": "million Euro",
			"de": "Millionen Euro",
		},
	}
	rT4 = configRT{
		InputNameP2: "volbysect",
		Chars:       10,
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by region: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in Q4 2022 nach Region: `,
		},
		Suffix: trl.S{
			// capitalizemytitle.com/how-to-abbreviate-million/
			// "en": "million €",
			"en": "MM €",
			"de": "Mio €",
		},
		SubNames: []string{
			"energy",
			"materials",
			"industrials",
			"consumer_disc",
			"consumer_stap",
			"healthcare",
			"financials",
			"information_tech",
			"comunication_svc",
			"utilities",
			"real_estate",
			"other",
		},
		SubLbls: map[string]string{
			"energy":           "Energy",
			"materials":        "Materials",
			"industrials":      "Industrials",
			"consumer_disc":    "Consumer Discretionary",
			"consumer_stap":    "Consumer Staples",
			"healthcare":       "Health Care",
			"financials":       "Financials",
			"information_tech": "Information Technology",
			"comunication_svc": "Comunication Services",
			"utilities":        "Utilities",
			"real_estate":      "Real Estate",
			"other":            "Other",
		},
		Placeholder: trl.S{
			"en": "million Euro",
			"de": "Millionen Euro",
		},
	}
)

func restrictedText(
	page *qst.WrappedPageT,
	inpNameMain string,
	lbl trl.S,
	cf configRT,
) {

	// gr1
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v_main", inpNameMain, cf.InputNameP2)

			inp.Label = lbl
			inp.Suffix = cf.Suffix

			inp.MaxChars = 12
			inp.Step = 1
			inp.Min = 0
			inp.Max = 1000 * 1000
			// inp.Validator = "inRange100"

			inp.ColSpan = 1
			inp.ColSpanLabel = 3
			inp.ColSpanControl = 2
		}

	}

	// gr2
	{
		gr := page.AddGroup()
		gr.Cols = 2
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Margin = "0 0 0 2rem"
		gr.BottomVSpacers = 3

		{

			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 2
			inp.Label = cf.LblRow2

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Width = "60%"
			inp.Style.Mobile.StyleBox.Width = "96%"

		}

		for _, suffx := range cf.SubNames {

			inp := gr.AddInput()
			inp.Type = "number"
			inp.Name = fmt.Sprintf("%v_%v_%v", inpNameMain, cf.InputNameP2, suffx)

			inp.Label = trl.S{
				"en": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
				"de": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
				//
			}
			inp.Suffix = cf.Suffix

			inp.MaxChars = 10
			inp.Step = 1
			inp.Min = 0
			inp.Max = 1000 * 1000
			// inp.Validator = "inRange100"

			inp.ColSpan = 2
			inp.ColSpanLabel = 3
			inp.ColSpanControl = 2
		}

		inp := gr.AddInput()
		inp.ColSpanControl = 1
		inp.Type = "javascript-block"
		inp.Name = "restrictedTextSum"

		s1 := trl.S{
			"de": "Addiert sicht nicht auf. Trotzdem weiter?",
			"en": "Does not add up. Continue?",
		}
		inp.JSBlockTrls = map[string]trl.S{
			"msg": s1,
		}

		inp.JSBlockStrings = map[string]string{}
		//

		inpMain := fmt.Sprintf("%v_%v", inpNameMain, cf.InputNameP2)
		inp.JSBlockStrings["InpMain"] = inpMain

		summands := []string{}
		for _, suffx := range cf.SubNames {
			nm := fmt.Sprintf("%v_%v", inpMain, suffx)
			summands = append(summands, nm)
		}

		inp.JSBlockStrings["SummandNames"] = `"` + strings.Join(summands, `", "`) + `"`

	}

}
