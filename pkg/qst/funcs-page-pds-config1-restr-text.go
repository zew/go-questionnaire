package qst

import (
	"github.com/zew/go-questionnaire/pkg/trl"
)

// config restricted text
type configRT struct {
	InputNameP2       string // second token of input name;  [numdeals, volbysegm, ...]
	SuppressSumField  bool   // dont show InputNameP2; only show Subnames...
	AddendsLighterSum bool   // addends not _equal_ but _lighter-than_ sum
	Chars             int    // input characters max
	LblRow1           trl.S  //
	FirstRowDisabled  bool   // first row of input is disabled
	FirstRow100Pct    bool   // first row of input is disabled and set to 100%
	LblRow2           trl.S  // question in more detail
	Suffix            trl.S  // unit 'deals' or 'million €'

	GroupLeftIndent string

	SubNames    []string // suffixes
	SubLbls     map[string]string
	Placeholder trl.S //

	Min, Max, Step float64
}

var (

	// multi-row configs
	rT11a = configRT{
		InputNameP2:       "q11a_numtransact",
		AddendsLighterSum: true,
		Chars:             6,
		LblRow1: trl.S{
			"en": "Total number of transactions",
			"de": "Gesamtzahl neue Abschlüsse",
		}.Outline("a.)"),

		LblRow2: trl.S{
			"en": "Thereof with...",
			"de": "Thereof with...",
		},
		Suffix:   suffixNumDeals,
		SubNames: []string{"floatingrate", "esgdoc", "esgratchet"},
		SubLbls: map[string]string{
			"floatingrate": `
				...floating interest rate
					<span class=font-size-90-block style='margin-left: 0.6rem; margin-top: 0.3rem;' >
					Please state the number of transactions with floating interest rate
					</span>
				`,
			"esgdoc": `
				...explicit ESG targets in the credit documentation
					<span class=font-size-90-block style='margin-left: 0.6rem; margin-top: 0.3rem;' >
					Please state the number of transactions with explicit ESG targets in the credit documentation.
					</span>
				`,
			"esgratchet": `
				...ESG ratchets
					<span class=font-size-90-block style='margin-left: 0.6rem; margin-top: 0.3rem;' >
					Pleases state the number of transactions with ESG ratchets.
					</span>
				`,
		},
		Placeholder: placeHolderNum,
	}

	rT11b = configRT{
		InputNameP2: "q11b_voltransact",
		Chars:       10,
		// LblRow1: see init()
		Suffix:      suffixMillionEuro,
		Step:        0.1,
		Placeholder: placeHolderMillion,
	}

	rT11dCorpLend = configRT{
		InputNameP2: "q11d_volbysegm",
		// SuppressSumField: true,
		Chars: 10,
		// LblRow1: see init()
		FirstRowDisabled: true,
		Suffix:           suffixMillionEuro,
		Step:             0.1,
		SubNames:         []string{"low", "mid", "upper"},
		SubLbls: map[string]string{
			"low":   "Lower mid-market (0-15m € EBITDA)",
			"mid":   "Core mid-market  (15-50m € EBITDA)",
			"upper": "Upper mid-market (>50m € EBITDA)",
		},
		Placeholder: placeHolderMillion,
	}
	rT11dRealEstate = configRT{}

	rT11e = configRT{
		InputNameP2: "q11e_volbyreg",
		// SuppressSumField: true,
		Chars: 10,
		LblRow1: trl.S{
			"en": "Region",
			"de": "Gesamtvolumen neuer Abschlüsse nach Region",
		}.Outline("e.)"),
		FirstRowDisabled: true,
		LblRow2: trl.S{
			"en": `Please state the volume (in mn EUR) of transactions closed in [quarter-1] by region`,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in [quarter-1] nach Region: `,
		},
		Suffix:   suffixMillionEuro,
		Step:     0.1,
		SubNames: []string{"uk", "france", "ger", "othereur"},
		SubLbls: map[string]string{
			"uk":       "UK",
			"france":   "France",
			"ger":      "Germany",
			"othereur": "Rest of Europe",
		},
		Placeholder: placeHolderMillion,
	}
	rT11fCorpLend = configRT{
		InputNameP2: "q11f_volbysect",
		// SuppressSumField: true,
		Chars: 10,
		LblRow1: trl.S{
			"en": "Sector",
			"de": "Gesamtvolumen neuer Abschlüsse nach Sektor",
		}.Outline("f.)"),
		FirstRowDisabled: true,
		LblRow2: trl.S{
			"en": `Please state the volume (in mn EUR) of transactions closed in [quarter-1] by sector`,
			"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in [quarter-1] nach Sektor: `,
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
			"communication_svc",
			"utilities",
			"real_estate",
			"other",
		},
		SubLbls: map[string]string{
			"energy":            "Energy",
			"materials":         "Materials",
			"industrials":       "Industrials",
			"consumer_disc":     "Consumer discretionary",
			"consumer_stap":     "Consumer staples",
			"healthcare":        "Health care",
			"financials":        "Financials",
			"information_tech":  "Information technology",
			"communication_svc": "Communication services",
			"utilities":         "Utilities",
			"real_estate":       "Real estate",
			"other":             "Other",
		},
		Placeholder: placeHolderMillion,
	}
	rT11fRealEstate = configRT{}
	rT11fInfrastruc = configRT{}

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
		Max:         100 * 1000 * 1000,
		Suffix:      suffixMillionEuro,
		Step:        0.1,
		Placeholder: placeHolderMillion,
	}

	rTSingleRowNumber = configRT{
		Chars:       7,
		Min:         0,
		Max:         1000 * 1000,
		Step:        1,
		Suffix:      suffixNumDeals,
		Placeholder: placeHolderNum,
	}
)

func init() {
	//
	lblB := trl.S{
		"en": "Total transaction volume (in mn EUR)",
		"de": "Total transaction volume (in mn EUR)",
	}
	lblB2 := trl.S{
		"en": `Please state the volume (in mn EUR) of transactions closed in [quarter-1]`,
		"de": `Please state the volume (in mn EUR) of transactions closed in [quarter-1]`,
	}
	lblB.Append90(lblB2)
	lblB.Outline("b.)")
	rT11b.LblRow1 = lblB

	//
	lbld := trl.S{
		"en": "Market segment",
		"de": "Gesamtvolumen neuer Abschlüsse nach Marktsegment",
	}
	lblD2 := trl.S{
		"en": `Please state the volume (in mn EUR) of transactions closed in [quarter-1] by market segment`,
		"de": `Bitte nennen Sie das Volumen (in Millionen Euro) von Abschlüssen in [quarter-1] nach Marktsegment: `,
	}
	// lbld.Append90(lblD2)
	rT11dCorpLend.LblRow2 = lblD2
	lbld.Outline("d.)")
	rT11dCorpLend.LblRow1 = lbld

	rT11dRealEstate = rT11dCorpLend
	rT11dRealEstate.SubNames = []string{"core", "coreplus", "valueadd", "opportun"}
	rT11dRealEstate.SubLbls = map[string]string{
		"core":     "Core",
		"coreplus": "Core+",
		"valueadd": "Value add",
		"opportun": "Opportunistic",
	}

	rT11fRealEstate = rT11fCorpLend
	rT11fRealEstate.SubNames = []string{"office", "retail", "hotel", "residential", "logistics", "other"}
	rT11fRealEstate.SubLbls = map[string]string{
		"office":      "Office",
		"retail":      "Retail",
		"hotel":       "Hospitality",
		"residential": "Residential",
		"logistics":   "Logistics",
		"other":       "Other",
	}

	rT11fInfrastruc = rT11fCorpLend
	rT11fInfrastruc.SubNames = []string{
		"transportation",
		"power",
		"renewables",
		"utilities",
		"telecoms",
		"social",
		"other",
	}
	rT11fInfrastruc.SubLbls = map[string]string{
		"transportation": "Transportation",
		"power":          "Power",
		"renewables":     "Renewables",
		"utilities":      "Utilities",
		"telecoms":       "Telecoms",
		"social":         "Social",
		"other":          "Other",
	}

}
