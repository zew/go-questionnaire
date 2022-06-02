package biii

import "github.com/zew/go-questionnaire/pkg/trl"

var radioVals3 = []string{"1", "2", "3"}
var radioVals4 = []string{"1", "2", "3", "4"}
var radioVals5 = []string{"1", "2", "3", "4", "5"}
var radioVals6 = []string{"1", "2", "3", "4", "5", "6"}

var columnTemplate3 = []float32{
	0.9, 1,
	0.0, 1,
	0.0, 1,
}

var columnTemplate3a = []float32{
	3.6, 1,
	0.0, 1,
	0.0, 1,
}

var columnTemplate4 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0.4, 1, // no answer slightly apart
}
var columnTemplate5 = []float32{
	3.6, 1,
	0.0, 1,
	0.0, 1,
	0.0, 1,
	0.0, 1,
}

var columnTemplate6 = []float32{
	2, 1,
	0, 1,
	0, 1,
	0, 1,
	0, 1,
	0.4, 1,
}

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
	{"de": "1. Keine Armut"},
	{"de": "2. Kein Hunger"},
	{"de": "3. Gesundheit und Wohlergehen"},
	{"de": "4. Hochwertige Bildung"},
	{"de": "5. Geschlechtergleichheit"},
	{"de": "6. Sauberes Wasser und Sanitär&shy;ein&shy;richtungen"},
	{"de": "7.Bezahlbare und Saubere Energie"},
	{"de": "8. Menschenwürdige Arbeit und Wirtschaftswachstum"},
	{"de": "9. Industrie, Innovation und Infrastruktur"},
	{"de": "10. Weniger Ungleichheiten"},
	{"de": "11. Nachhaltige Städte und Gemeinden"},
	{"de": "12. Nachhaltiger Konsum und Produktion"},
	{"de": "13. Maßnahmen zum Klimaschutz"},
	{"de": "14. Leben unter Wasser"},
	{"de": "15. Leben an Land"},
	{"de": "16. Frieden, Gerechtigkeit und starke Institutionen"},
	{"de": "17. Partnerschaften zur Erreichung der Ziele"},
}

var q13inputNames = []string{
	"not",
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
	"14",
	"15",
	"16",
	"17",
}

var q25Columns = []trl.S{
	{"de": "Übertrifft meine Erwartung"},
	{"de": "Entspricht meiner Erwartung"},
	{"de": "Liegt unterhalb meiner Erwartung"},
}

var q25RadioVals = []string{"exceeds", "as_expected", "below"}

var oneToFiveEfficiency = []trl.S{
	{"de": "1<br>sehr effektiv"},
	{"de": "2<br><br>effektiv"},
	{"de": "3<br><br>teils/teils"},
	{"de": "4<br>eher ineffektiv"},
	{"de": "5<br><br>ineffektiv"},
}

var q26Labels = []trl.S{
	{"de": "Deutschland"},
	{"de": "international"},
}

var q26inputNames = []string{
	"germany",
	"international",
}

var q27columns = []trl.S{
	{"de": "1<br>sehr ein&shy;deutig"},
	{"de": "2<br><br>ein&shy;deutig"},
	{"de": "3<br><br>teils/teils"},
	{"de": "4<br>eher unklar"},
	{"de": "5<br>völlig unklar"},
}

var q30columns = []trl.S{
	{"de": "1<br>sehr großen"},
	{"de": "2<br><br>großen"},
	{"de": "3<br><br>geringen"},
	{"de": "4<br><br>keinen"},
	{"de": "5<br>keine Angabe"},
}

var q31columns = []trl.S{
	{"de": "1<br>äußerst relevant"},
	{"de": "2<br><br>relevant"},
	{"de": "3<br><br>teils/teils"},
	{"de": "4<br>eher irrelevant"},
	{"de": "5<br><br>irrelevant"},
}

var q32columns = []trl.S{
	{"de": "1<br>größter Fortschriftt"},
	{"de": "2<br><br>"},
	{"de": "3<br>geringster Fortschriftt"},
}

var q33aColumns = []trl.S{
	{"de": "1<br>größter Bedarf"},
	{"de": "2<br><br>"},
	{"de": "3<br>geringster Bedarf"},
}
var q33bColumns = []trl.S{
	{"de": "1<br>größtes Potenzial"},
	{"de": "2<br><br>"},
	{"de": "3<br>geringstes Potenzial"},
}

var q34columns = []trl.S{
	{"de": "1<br>höchste Entwicklung"},
	{"de": "2<br><br>"},
	{"de": "3<br>geringste Entwicklung"},
}
