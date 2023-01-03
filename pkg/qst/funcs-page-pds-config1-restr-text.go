package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

// config restricted text
type configRT struct {
	InputNameP2      string // second token of input name;  [numdeals, volbysegm, ...]
	SuppressSumField bool   // dont show InputNameP2; only show Subnames...
	Chars            int    // input characters max
	LblRow1          trl.S  //
	FirstRow100Pct   bool   // first row of input is 100% disabled
	LblRow2          trl.S  // question in more detail
	Suffix           trl.S  // unit 'deals' or 'million €'

	GroupLeftIndent string

	SubNames    []string // suffixes
	SubLbls     map[string]string
	Placeholder trl.S //

	Min, Max, Step float64
}

var (

	// multi-row configs
	rT1 = configRT{
		InputNameP2: "q11a_numtransact",
		Chars:       6,
		LblRow1: trl.S{
			"en": "Total number of transactions",
			"de": "Gesamtzahl neue Abschlüsse",
		}.Outline("a.)"),
		// GroupLeftIndent: outline2Indent,
		LblRow2: trl.S{
			"en": fmt.Sprint(
				// `Please state the number of deals closed in Q4 2022 by market segment: `,
				`Please state the number of transactions closed in [quarter] [year] `,
			),
			"de": fmt.Sprint(
				`Please state the number of transactions closed in [quarter] [year] `,
			),
		},
		Suffix: trl.S{
			"en": "deals",
			"de": "Deals",
			// "en": "transactions",
			// "de": "Transaktionen",
		},

		SubNames: []string{"floatingrate", "esgdoc", "esgratchet"},
		SubLbls: map[string]string{
			"floatingrate": "Thereof with floating interest rate",
			"esgdoc":       "Thereof with explicit ESG targets in the credit documentation",
			"esgratchet":   "Thereof with ESG ratchets",
		},
		Placeholder: placeHolderNum,
	}

	rT1b = configRT{
		InputNameP2: "q11b_voltransact",
		Chars:       10,
		LblRow1: trl.S{
			"en": fmt.Sprint(
				`Please state the volume (in mn €) of transactions closed in [quarter] [year]`,
			),
			"de": fmt.Sprint(
				`Please state the volume (in mn €) of transactions closed in [quarter] [year]`,
			),
		},
		Suffix:      suffixMillionEuro,
		Placeholder: placeHolderMillion,
	}

	rT2 = configRT{
		InputNameP2:      "q11d_volbysegm",
		SuppressSumField: true,
		Chars:            10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by segment",
			"de": "Gesamtvolumen neuer Abschlüsse nach Marktsegment",
		}.Outline("d.)"),
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in [quarter] [year] by market segment: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in [quarter] [year] nach Marktsegment: `,
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
	rT2RealEstate = configRT{}

	rT3 = configRT{
		InputNameP2:      "q11e_volbyreg",
		SuppressSumField: true,
		Chars:            10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by region",
			"de": "Gesamtvolumen neuer Abschlüsse nach Region",
		}.Outline("e.)"),
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in [quarter] [year] by region: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in [quarter] [year] nach Region: `,
		},
		Suffix: suffixMillionEuro,
		Step:   0.1,

		// SubNames: []string{"uk", "france", "dach", "benelux", "nordics", "southern_eu", "other"},
		// SubLbls: map[string]string{
		// 	"uk":          "UK",
		// 	"france":      "France",
		// 	"dach":        "DACH",
		// 	"benelux":     "Benelux",
		// 	"nordics":     "Nordics",
		// 	"southern_eu": "Southern Europe",
		// 	"other":       "Other",
		// },
		SubNames: []string{"uk", "france", "ger", "othereur"},
		SubLbls: map[string]string{
			"uk":       "UK",
			"france":   "France",
			"ger":      "Germany",
			"othereur": "Rest of Europe",
		},
		Placeholder: placeHolderMillion,
	}
	rT4 = configRT{
		InputNameP2:      "q11f_volbysect",
		SuppressSumField: true,
		Chars:            10,
		LblRow1: trl.S{
			"en": "Total volume of new deals by sector",
			"de": "Gesamtvolumen neuer Abschlüsse nach Sektor",
		}.Outline("f.)"),
		LblRow2: trl.S{
			"en": `Please state the volume (in million Euro) of deals closed in [quarter] [year] by sector: `,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in [quarter] [year] nach Sektor: `,
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
	rT4RealEstate = configRT{}
	rT4Infrastruc = configRT{}

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
		Suffix: suffixPercent,

		// SubNames: []string{"uk", "france", "dach", "benelux", "nordics", "southern_eu", "other"},
		// SubLbls: map[string]string{
		// 	"uk":          "UK",
		// 	"france":      "France",
		// 	"dach":        "DACH",
		// 	"benelux":     "Benelux",
		// 	"nordics":     "Nordics",
		// 	"southern_eu": "Southern Europe",
		// 	"other":       "Other",
		// },
		SubNames: []string{"uk", "france", "ger", "othereur"},
		SubLbls: map[string]string{
			"uk":       "UK",
			"france":   "France",
			"ger":      "Germany",
			"othereur": "Rest of Europe",
		},

		Placeholder: placeHolderNum,
	}

	//
	//
	// single row configs for page 3
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
	// single row configs for page 3
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
)

func init() {

	rT2RealEstate = rT2
	rT2RealEstate.SubNames = []string{"core", "coreplus", "valueadd", "opportun"}
	rT2RealEstate.SubLbls = map[string]string{
		"core":     "Core",
		"coreplus": "Core+",
		"valueadd": "Value add",
		"opportun": "Opportunistic",
	}

	rT4RealEstate = rT4
	rT4RealEstate.SubNames = []string{"office", "retail", "hotel", "residential", "logistics", "other"}
	rT4RealEstate.SubLbls = map[string]string{
		"office":      "Office",
		"retail":      "Retail",
		"hotel":       "Hospitality",
		"residential": "Residential",
		"logistics":   "Logistics",
		"other":       "Other",
	}

	rT4Infrastruc = rT4
	rT4Infrastruc.SubNames = []string{
		"transportation",
		"power",
		"renewables",
		"utilities",
		"telecoms",
		"social",
		"other"}
	rT4Infrastruc.SubLbls = map[string]string{
		"transportation": "Transportation",
		"power":          "Power",
		"renewables":     "Renewables",
		"utilities":      "Utilities",
		"telecoms":       "Telecoms",
		"social":         "Social",
		"other":          "Other",
	}

}
