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
var radioVals5 = []string{"1", "2", "3", "4", "5"}
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

/* var headersGoodBad = []trl.S{
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
*/

var oneToFiveVolume = []trl.S{
	{"de": "1<br>höchstes<br>Vol."},
	{"de": "2"},
	{"de": "3"},
	{"de": "4"},
	{"de": "5<br>niedrigstes<br>Vol."},
	// {"de": "keine<br>Angabe"},
}

var q12Labels = []trl.S{
	{"de": "Nicht Anwendbar (N/A)"},
	{"de": "Paris-Aligned oder Net Zero"},
	{"de": "Kultur und Freizeit (Kultur, Kunst, Sport, sonstige Freizeitgestaltung und soziale Vereine)"},
	{"de": "Bildung (Grundschule, Sekundarschule, Hochschule, Sonstiges)"},
	{"de": "Erwerbstätigkeit"},
	{"de": "Forschung"},
	{"de": "Gesundheit (Krankenhäuser, Rehabilitation, Pflegeheime, psychische Gesundheit / Krisenintervention)"},
	{"de": "Soziale Dienste (Notfall, Hilfe, Einkommens&shy;unterstützung / Unterhalt)"},
	{"de": "Umweltschutz (Forstwirtschaft, Land, Abfall, Luft, biologische Vielfalt und Ökosysteme, Meere und Küstengebiete)"},
	{"de": "WASH (Wasser, Sanitärversorgung und Hygiene)"},
	{"de": "Landwirtschaft"},
	{"de": "Energie (Zugang zu Energie, erneuerbare Energie)"},
	{"de": "Wohnen"},
	{"de": "IT / Technologien"},
	{"de": "Fertigung / Produktion"},
	{"de": "Stadterneuerung / Territoriale Entwicklung"},
	{"de": "Finanzielle Eingliederung und Zugang zu Finanzmitteln (d.h. Mikro&shy;finanzierung, Mikro&shy;versicherungen, Finanz Bildungs&shy;dienst&shy;leistungen, Bankwesen)"},
}

var q12inputNames = []string{
	"na",
	"paris_align",
	"culture_sports",
	"education",
	"work",
	"research",
	"health",
	"social_service",
	"environment",
	"sanitary",
	"agriculture",
	"energy",
	"residential",
	"technology",
	"prodution",
	"urban_dev",
	"microfinance",
}

var oneToFiveImportance = []trl.S{
	{"de": "1<br>höchste<br>Wichtigkeit"},
	{"de": "2"},
	{"de": "3"},
	{"de": "4"},
	{"de": "5<br>niedrigste<br>Wichtigkeit"},
	// {"de": "keine<br>Angabe"},
}

var q13Labels = []trl.S{
	{"de": "Wir berücksichtigen die SDGs in unserer Strategie nicht"},
	{"de": "Keine Armut"},
	{"de": "Kein Hunger"},
	{"de": "Gesundheit und Wohlergehen"},
	{"de": "Hochwertige Bildung"},
	{"de": "Geschlechtergleichheit"},
	{"de": "Sauberes Wasser und Sanitäreinrichtungen"},
	{"de": "Bezahlbare und Saubere Energie"},
	{"de": "Menschenwürdige Arbeit und Wirtschaftswachstum"},
	{"de": "Industrie, Innovation und Infrastruktur"},
	{"de": "Weniger Ungleichheiten"},
	{"de": "Nachhaltige Städte und Gemeinden"},
	{"de": "Nachhaltiger Konsum und Produktion"},
	{"de": "Maßnahmen zum Klimaschutz"},
	{"de": "Leben unter Wasser"},
	{"de": "Leben an Land"},
	{"de": "Frieden, Gerechtigkeit und starke Institutionen"},
	{"de": "Partnerschaften zur Erreichung der Ziele"},
}

var q13inputNames = []string{
	"1",
	"2",
	"3",
	"4",
	"5",
	"6",
	"7",
	"8",
	"9",
	"10",
	"11",
	"12",
	"13",
}

var q25Columns = []trl.S{
	{"de": "Übertrifft meine Erwartung"},
	{"de": "Entspricht meiner Erwartung"},
	{"de": "Liegt unterhalb meiner Erwartung"},
}

var q25RadioVals = []string{"exceeds", "as_expected", "below"}

var oneToFiveEfficiency = []trl.S{
	{"de": "1<br>sehr effektiv"},
	{"de": "2<br<br>effektiv"},
	{"de": "3<br><br>teils/teils"},
	{"de": "4<br>eher ineffektiv"},
	{"de": "5<br><brineffektiv"},
}

var q26Labels = []trl.S{
	{"de": "Deutschland"},
	{"de": "international"},
}

var q26inputNames = []string{
	"germany",
	"international",
}
