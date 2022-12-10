package pds

import (
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
		Step:        1,
		Suffix:      suffixMillionEuro,
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
		Suffix:   suffixMillionEuro,
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
		InputNameP2: "221market_segment",
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
		InputNameP2: "222region",
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
