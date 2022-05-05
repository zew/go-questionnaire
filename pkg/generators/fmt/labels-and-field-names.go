package fmt

import "github.com/zew/go-questionnaire/pkg/trl"

var radioVals4 = []string{"1", "2", "3", "4"}
var radioVals6 = []string{"1", "2", "3", "4", "5", "6"}
var columnTemplate4 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0.4, 1, // no answer slightly apart
}
var columnTemplate6 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

var rowLabelsEuroGerUSGlob1 = []trl.S{
	{
		"de": "Euroraum",
		"en": "Euro area",
	},
	{
		"de": "Deutschland",
		"en": "Germany",
	},
	{
		"de": "USA",
		"en": "US",
	},
	{
		"de": "China",
		"en": "China",
	},
}
var rowLabelsEuroGerUSGlob2 = []trl.S{
	{
		"de": "Euroraum",
		"en": "Euro area",
	},
	{
		"de": "USA",
		"en": "US",
	},
	{
		"de": "China",
		"en": "China",
	},
}

var rowLabelsSectors = []trl.S{
	{
		"de": "Banken",
		"en": "Banks",
	},
	{
		"de": "Versicherungen",
		"en": "Insurance",
	},
	{
		"de": "Fahrzeugbau",
		"en": "Automotive",
	},
	{
		"de": "Chemie, Pharma",
		"en": "Chemicals, Pharma",
	},
	{
		"de": "Stahl/NE-Metalle",
		"en": "Steel/Metal Products",
	},
	{
		"de": "Elektronik",
		"en": "Electronics",
	},
	{
		"de": "Maschinen&shy;bau",
		"en": "Machinery",
	},
	// row 2
	{
		"de": "Konsum, Handel",
		"en": "Private Consumption / Retail Sales",
	},
	{
		"de": "Baugewerbe",
		"en": "Construction",
	},
	{
		"de": "Versorger",
		"en": "Utilities",
	},
	{
		"de": "Dienstleister",
		"en": "Services",
	},
	{
		"de": "Telekommunikation",
		"en": "Tele&shy;communications",
	},
	{
		"de": "Inform.-Technologien",
		"en": "Inform.-Technologies",
	},
}
