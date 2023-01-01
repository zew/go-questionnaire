package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// config restricted text
type configRT struct {
	InputNameP2    string // second token of input name;  [numdeals, volbysegm, ...]
	Chars          int    // input characters max
	LblRow1        trl.S  //
	FirstRow100Pct bool   // first row of input is 100% disabled
	LblRow2        trl.S  // question in more detail
	Suffix         trl.S  // unit 'deals' or 'million €'

	GroupLeftIndent string

	SubNames    []string // suffixes
	SubLbls     map[string]string
	Placeholder trl.S //

	Min, Max, Step float64
}

var (
	// simple config
	rTSingleRowPercent = configRT{
		Chars: 4,
		LblRow1: trl.S{
			"de": "By segment",
			"en": "By segment",
		}.Outline("a)"),
		Min:         0,
		Max:         100,
		Step:        0.1,
		Suffix:      suffixPercent,
		Placeholder: placeHolderNum,
	}
	rTSingleRowMill = configRT{
		Chars: 10,
		LblRow1: trl.S{
			"en": "By region",
			"de": "By region",
		}.Outline("b)"),
		Min:         0,
		Max:         40000,
		Suffix:      suffixMillionEuro,
		Step:        0.1,
		Placeholder: placeHolderMillion,
	}

	// multi-row configs
	rT1 = configRT{
		InputNameP2: "q11a_numdeals",
		Chars:       6,
		LblRow1: trl.S{
			"en": "Total number of new deals",
			"de": "Gesamtzahl neue Abschlüsse",
		}.Outline("a.)"),
		GroupLeftIndent: outline2Indent,
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
			"low":      "Lower mid-market (0-15m € EBITDA)",
			"midupper": "Core- and Upper mid-market (>15m € EBITDA)",
			"other":    "Other",
		},
		Placeholder: placeHolderNum,
	}
	rT2 = configRT{
		InputNameP2: "q11c_volbysegm",
		Chars:       10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by segment",
			"de": "Gesamtvolumen neuer Abschlüsse nach Marktsegment",
		}.Outline("c.)"),
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by market segment: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in Q4 2022 nach Marktsegment: `,
		},
		Suffix:   suffixMillionEuro,
		Step:     0.1,
		SubNames: []string{"low", "mid", "upper"},
		SubLbls: map[string]string{
			"low":   "Lower mid-market (0-15m € EBITDA)",
			"mid":   "Core mid-market  (15-50m € EBITDA)",
			"upper": "Upper mid-market (>50m € EBITDA)",
		},
		Placeholder: placeHolderMillion,
	}
	rT3 = configRT{
		InputNameP2: "q11d_volbyreg",
		Chars:       10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by region",
			"de": "Gesamtvolumen neuer Abschlüsse nach Region",
		}.Outline("d.)"),
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by region: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in Q4 2022 nach Region: `,
		},
		Suffix: suffixMillionEuro,
		Step:   0.1,

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
		Placeholder: placeHolderMillion,
	}
	rT4 = configRT{
		InputNameP2: "q11e_volbysect",
		Chars:       10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by sector",
			"de": "Gesamtvolumen neuer Abschlüsse nach Sektor",
		}.Outline("e.)"),
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in Q4 2022 by sector: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in Q4 2022 nach Sektor: `,
		},
		Suffix: suffixMillionEuro,
		Step:   0.1,

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
			"consumer_disc":    "Consumer discretionary",
			"consumer_stap":    "Consumer staples",
			"healthcare":       "Health care",
			"financials":       "Financials",
			"information_tech": "Information technology",
			"comunication_svc": "Comunication services",
			"utilities":        "Utilities",
			"real_estate":      "Real estate",
			"other":            "Other",
		},
		Placeholder: placeHolderMillion,
	}

	//
	r221 = configRT{
		InputNameP2: "q22a_market_segment",
		Chars:       5,
		LblRow1: trl.S{
			"en": "Share of portfolio by segment (at fair market value)",
			"de": "Share of portfolio by segment (at fair market value)",
		}.Outline("a.)"),
		FirstRow100Pct: true,
		LblRow2: trl.S{
			"en": `Please enter percentages for each segment`,
			"de": `Please enter percentages for each segment`,
		},
		Suffix:   suffixPercent,
		SubNames: []string{"low", "core", "upper"},
		SubLbls: map[string]string{
			"low":   "Lower (5-15mn EBITDA)",
			"core":  "Core (15-50 mn EBITDA)",
			"upper": "Upper Mid Market (50-75 mn EBITDA)",
		},
		Placeholder: placeHolderNum,
	}

	r222 = configRT{
		InputNameP2: "q22b_region",
		Chars:       5,
		LblRow1: trl.S{
			"en": "Share of portfolio by region (at fair market value)",
			"de": "Share of portfolio by region (at fair market value)",
		}.Outline("b.)"),
		FirstRow100Pct: true,
		LblRow2: trl.S{
			"en": `Please enter percentages for each region`,
			"de": `Please enter percentages for each region`,
		},
		Suffix:   suffixPercent,
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
		Placeholder: placeHolderNum,
	}
)

func restrictedTextMultiCols(
	page *pageT,
	ac assetClass,
	cf configRT,
) {

	numCols := firstColLbl + float32(len(ac.TrancheTypes))
	idxLastCol := len(ac.TrancheTypes) - 1

	{
		gr := page.AddGroup()
		gr.Cols = numCols
		// gr.BottomVSpacers =

		// row0 - column headers, 1-4 tranche type names
		for idx1 := 0; idx1 < len(ac.TrancheTypes)+1; idx1++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			if idx1 == 0 {
				inp.ColSpan = firstColLbl
			}
			if idx1 > 0 {
				ttLbl := ac.TrancheTypes[idx1-1].Lbl
				inp.Label = ttLbl
			}
			inp.StyleLbl = styleHeaderCols1
		}

		// row1
		for idx1, trancheType := range ac.TrancheTypes {

			{
				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v_%v", ac.Prefix, trancheType.Prefix, cf.InputNameP2, "main")

				inp.MaxChars = cf.Chars
				inp.Step = 1
				if cf.Step != 0.0 {
					inp.Step = cf.Step
				}
				inp.Min = 0
				inp.Max = 1000 * 1000
				// inp.Validator = "inRange100"

				inp.Placeholder = cf.Placeholder

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				if idx1 == 0 {
					inp.Label = cf.LblRow1.Bold()
					inp.Label = cf.LblRow1
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1
				}
				if idx1 == idxLastCol {
					inp.Suffix = cf.Suffix
				}

				if cf.FirstRow100Pct {
					inp.Response = "100" // must parse to number
					inp.Disabled = true

				}

			}
		}

		// row2
		if !cf.LblRow2.Empty() {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols

			inp.Label = cf.LblRow2

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleBox.Width = "60%"
			inp.Style.Mobile.StyleBox.Width = "96%"
		}

		// rows 3,4...
		for _, suffx := range cf.SubNames {

			for idx2, trancheType := range ac.TrancheTypes {

				inp := gr.AddInput()
				inp.Type = "number"
				inp.Name = fmt.Sprintf("%v_%v_%v_%v", ac.Prefix, trancheType.Prefix, cf.InputNameP2, suffx)

				inp.MaxChars = cf.Chars
				inp.Step = 1
				if cf.Step != 0.0 {
					inp.Step = cf.Step
				}
				inp.Min = 0
				inp.Max = 1000 * 1000
				// inp.Validator = "inRange100"

				inp.Placeholder = cf.Placeholder

				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1

				if idx2 == 0 {
					inp.Label = trl.S{
						"en": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
						"de": fmt.Sprintf("- %v", cf.SubLbls[suffx]),
					}
					inp.ColSpan = firstColLbl + 1
					inp.ColSpanLabel = firstColLbl
					inp.ColSpanControl = 1

					inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
					inp.StyleLbl.Desktop.StyleBox.Margin = "0 0 0 1.2rem"
				}

				if idx2 == idxLastCol {
					inp.Suffix = cf.Suffix
				}

			}

		}

	} // /group

}
