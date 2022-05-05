package biii

import "github.com/zew/go-questionnaire/pkg/trl"

var radioVals3 = []string{"1", "2", "3"}

var radioValsQ1 = []string{"investor", "assetmgr", "passiveparticipant"}
var roleOrFunctionQ1 = []trl.S{
	{"de": "Investor<br>(asset owner)"},
	{"de": "Vermögensverwalter<br>(asset manager)"},
	{"de": "Ein anderer (passiver) Marktteilnehmer (z.B. Berater, ...)"},
}

var radioVals4 = []string{"1", "2", "3", "4"}
var radioVals6 = []string{"1", "2", "3", "4", "5", "6"}

var columnTemplate3 = []float32{
	0.9, 1,
	0.0, 1,
	0.0, 1,
}
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

var headersGoodBad = []trl.S{
	{
		"de": "gut",
		"en": "good",
	},
	{
		"de": "normal",
		"en": "normal",
	},
	{
		"de": "schlecht",
		"en": "bad",
	},
	{
		"de": "keine<br>Angabe",
		"en": "no estimate",
	},
}
