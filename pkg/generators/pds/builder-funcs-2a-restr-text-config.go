package pds

import (
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
		LblRow1: trl.S{
			"en": "Total volume of new deals by region",
			"de": "Gesamtvolumen neuer Abschlüsse nach Region",
		},
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
		LblRow1: trl.S{
			"en": "Total volume of new deals by sector",
			"de": "Gesamtvolumen neuer Abschlüsse nach Sektor",
		},
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

	rT5 = configRT{
		InputNameP2: "share_loans_default",
		Chars:       3,
		LblRow1: trl.S{
			"en": `lbl1`,
			"de": `lbl1`,
		},
		LblRow2: trl.S{},
		Suffix: trl.S{
			"en": "%",
			"de": "%",
		},
		SubNames: []string{},
		SubLbls:  map[string]string{},
		Placeholder: trl.S{
			"en": "#",
			"de": "#",
		},
	}
)
