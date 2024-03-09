package trl

import (
	"log"
	"sort"
	"strings"
)

// Hyphenization
// =================
//
// hyph is a slice _with_ hyphenized words.
// During application initialization we use it to fill hyphm above.
var hyph = []string{

	// www.duden.de/rechtschreibung/Abweichung

	// reduced hyphenization
	"All|tags|si|tu|a|ti|on",
	"All|tags|si|tu|a|ti|on|en",
	"Aktien|markt",
	"Bildungs|abschluss",
	"unter|bewertet", // 	"be|wer|tet",
	"über|bewertet",
	"Über|wa|chung",
	"Beitrags|prämien",
	"Einlagen|sicherungs|fonds",
	"Entschädigungs|fall ",
	"Hoch|schul|abschluss",
	"Konjunktur|entwicklung",
	"Kredit|geber",
	"Kredit|nehmer",
	"Risiko|komponente",
	"Maschinen|bau",
	"Prozent|punkt",
	"Transformations|projekt",

	"Ober|gren|ze",
	"Kon|fi|denz|in|ter|vall",
	"Un|ter|gren|ze",

	// full german
	"de|mo|gra|fisch",
	"de|mo|gra|phisch",

	//
	// german
	"ab|wer|ten",
	"Ab|wei|chung",
	"Ana|ly|se",
	"Ana|ly|sen",
	"All|tag",  // www.duden.de/rechtschreibung/Alltag
	"An|ga|be", // www.duden.de/rechtschreibung/Angabe
	"Ant|wort", // www.duden.de/rechtschreibung/Antwort
	"Arbeits|markt|öko|no|mie",
	"as|so|zi|ie|rt",
	"auf|wer|ten",
	"Aus|fall|risiken",
	"aus|ge|schlos|sen",
	"be|wer|ten",
	"be|wer|tet",
	"Bil|dung",
	"be|ant|wor|ten",
	"blei|ben",
	"da|für",
	"da|ge|gen",
	"Da|ten",
	"deut|lich",
	"Deutsch|land",
	"Do|zen|tIn",
	"Dok|to|ran|din",
	"ei|ge|ne", // https://www.duden.de/rechtschreibung/eigener
	"ef|fek|tiv",
	"Ein|fluss",
	"emp|feh|len",
	"ein|zu|ge|hen",
	"Ent|schei|dungs|pro|zesse",
	"Ent|wick|lungs|öko|no|mie",
	"er|stre|bens|wert",
	"er|hal|ten",
	"er|heb|lich",
	"er|höht",
	"Erwartungs|wert",
	"Euro|raum",
	"Euro|gebiet",
	"Ex|port|markt",
	"Ex|port|märkte",
	"Fi|nanz|wis|sen|schaft",
	"Finanz|wirt|schaft",
	"For|schung",
	"ge|samt",
	"ge|ra|ten", // www.duden.de/rechtschreibung/geraten_gelingen_hingelangen
	"ge|eig|net",
	"Geld|po|li|tik",
	"Groß|unter|nehmen",
	"grö|ßer",
	"Grup|pe",
	"Grup|pen",
	"Grup|pen|mit|glieder",
	"Handels|konflikte",
	"höchs|te",
	"hö|her",
	"Immob|ilien|kredite",
	"In|fla|ti|on",
	"in|for|miert", // www.duden.de/rechtschreibung/uninformiert
	"in|ef|fek|tiv",
	"In|dus|trie|öko|no|mie",
	"ir|re|le|vant",
	"Junior|pro|fes|sorIn",
	"kei|ne", // hard; but allowed https://www.duden.de/rechtschreibung/keine
	"kom|plett",
	"Kompromiss|lösung",
	"Kon|junk|tur|da|ten ",
	"Kon|sum|enten|kredite",
	"Kredit|angebot",
	"Kredit|nach|frage",
	"Kom|po|nen|te",
	"lang|fristig",
	"Ma|kro|öko|no|mie",
	"Mehrheits|lösung",
	"Mi|kro|öko|no|mie",
	"Mit|glied",
	"Mit|glieder",
	"mitt|le|re",
	"mittel|fristig",
	"ne|ga|tiv",
	"neu|tral",
	"nied|rig",
	"nor|mal",
	"Öko|no|mie",
	"Po|li|tik",
	"po|si|tiv",
	"po|ten|zi|ell",
	"Prä|ferenz|konstellation",
	"Pro|fes|sorIn",
	"Pro|jekt",
	"Pro|gno|se",
	"Pro|zent",
	"re|du|zie|ren",
	"re|du|ziert",
	"Re|finanz|ierung",
	"Re|gie|rung",
	"Re|gie|rungs|bil|dung",
	"Regierungs|bildung",
	"Regu|lierung",
	"re|le|vant",
	"Ri|si|ko",
	"ri|si|ko|be|reit",
	"Ri|si|ko|ver|mei|den",
	"Risiko|trag|fähig|keit",
	"Roh|stoff|preise",
	"schlech|teste",
	"schlech|ter", // ugly but necessary
	"Schul|den|brem|se",
	"si|cher",
	"sin|ken",
	"stei|gen",
	"stim|me", // daring, as in "stimme voll zu"
	"Stif|tung",
	"Si|tu|a|ti|on", // www.duden.de/rechtschreibung/Situation
	"Stan|dard|abweichung ",
	"Te|le|kom|mu|ni|ka|ti|on",
	"Trans|for|ma|ti|on",
	"über|haupt",
	"Um|welt|öko|no|mie",
	"un|ent|schieden",
	"un|in|for|miert", // www.duden.de/rechtschreibung/uninformiert
	"un|ter",
	"Un|ter|neh|men",
	"un|ver|än|dert",
	"un|wich|tig",
	"ver|än|dern",
	"ver|än|dert",
	"Ver|bes|se|rung",
	"ver|bes|sern",
	"ver|bes|sert",
	"ver|füg|bar",
	"Ver|trau|en", // https://www.duden.de/rechtschreibung/Vertrauen
	"Ver|schlech|te|rung",
	"ver|schlech|tern",
	"ver|schlech|tert",
	"ver|än|dern",
	"ver|än|dert",

	"Vermögens|verwalter ",
	"Markt|teil|nehmer ",

	"Wechsel|kurse",
	"Welt|wirt|schaft",
	"Wett|be|werbs|sit|uation",
	"wich|tig|keit",
	"wich|tig",
	"wirt|schaft",
	"wirt|schaft|lich",
	"Wirt|schafts|politik",
	"Wirt|schafts|for|schung",
	"Wirt|schafts|for|schung",
	"zwi|schen",
	"zu|tref|fen",

	// https://www.merriam-webster.com/dictionary/estimate
	// english
	"ac|counts",
	"ad|min|is|tra|tion",
	"an|swer",
	"ap|pli|ca|ble",
	"ap|pre|ci|ate", // appreciate
	"as|sis|tant",
	"as|so|ci|ate",
	"av|er|age",
	"be|low",
	"be|tween",
	"busi|ness",
	"can|di|date",

	"chem|i|cal",
	"con|fi|dence",
	"com|plete|ly",
	"com|mu|ni|ca|tion",
	"tele|com|mu|ni|ca|tion",
	"comp|etitive",
	"cons|umer",
	"con|stant",
	"de|cid|ed",
	"de|crease",
	"de|cline",
	"de|pre|ci|ate", // "depreciate"
	"de|te|ri|o|rate",
	"de|vel|op|ment",
	"dis|agree",
	"dis|cre|tion|ary",
	"doc|u|men|ta|tion",
	"eco|nom|ic",
	"eco|nom|ics",
	"econ|o|my",
	"en|vi|ron|ment",
	"en|vi|ron|mental",
	"es|ti|mate",
	"enter|prises",
	"environ|ment",
	"im|prove",
	"Ger|ma|ny",
	"im|prove|ment",
	"in|crease",
	"in|dif|fer|ent",
	"in|dus|tri|al",
	"in|for|ma|tion",
	"in|flu|ence",
	"in|sti|tute",
	"in|sti|tutes",
	"in|ter|val",
	"in|ter|na|tion|al",
	"in|vest|ment",
	"man|age|ment",
	"mac|ro|eco|nom|ics",
	"mez|za|nine",
	"mi|cro|eco|nom|ics",
	"na|tion|al|i|ty",
	"neg|ative",
	"nei|ther",
	"op|ti|mis|tic",
	"pes|si|mis|tic",
	"per|cent",
	"phar|ma|ceu|ti|cal",
	"pos|itive",
	"po|ten|tial",
	"pro|fes|sor",
	"ques|tion|naire ",
	"re|duce",
	"re|duced",
	"Re|financing",
	"re|main",
	"re|search",
	"Reg|ulation",
	"sav|ings",
	"ser|vice",
	"sig|nif|i|cant",
	"sig|nif|i|cant|ly",
	"slight|ly",
	"Small+|medium",
	"strong|ly",
	"strong|ly",
	"sub|or|di|nat|ed",
	"tech|nol|o|gy",
	"trans|ac|tion",
	"un|changed",
	"un|de|cid|ed",
	"unit|tranche",
	"wors|en",
	"will|ing",

	// spanish
	"in|deciso",
	"acuer|do",
	"desa|cuer|do",

	// french
	"acc|ord",
	"In|diff|érent",

	// italian
	"favo|revole",
	"in|deciso",
}

// hyphm is a map with words and their hyphenized form as value.
// hyphm is filled during app initialization from hyph above.
var hyphm = map[string]string{}

// todo: Differentiation by language?
func init() {
	cntr := -1
	for _, v := range hyph {
		cntr++
		if len(v) < 1 {
			continue
		}

		if strings.Contains(v, "­") {
			log.Fatalf("hyph contains strange char in %v", v)
		}

		v = strings.Replace(v, "­", "", -1) // strange invisible char - pasted from Duden

		key := strings.Replace(v, "|", "", -1)
		val := strings.Replace(v, "|", "&shy;", -1)

		keyLow := strings.ToLower(key)
		valLow := strings.ToLower(val)

		keyHig := strings.ToUpper(string([]rune(key)[:1])) // capitalize first UTF8 rune
		keyHig += string([]rune(key)[1:])                  // concat remainder

		valHig := strings.ToUpper(string([]rune(val)[:1])) //    ~
		valHig += string([]rune(val)[1:])                  //    ~

		hyphm[keyLow] = valLow
		hyphm[keyHig] = valHig

		/*
			if cntr < 90 && cntr > 80 {
				log.Printf("hyphm:  %-24v %v", keyLow, valLow)
				log.Printf("hyphm:  %-24v %v", keyHig, valHig)
			}
		*/
	}
}

// The last step, is to order our hyphenization map
// by key length.
// Otherwise we get stochastic results.
type byLength []string

var byLen = byLength{}

func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	if len(s[i]) == len(s[j]) {
		return s[i] > s[j] // alphanumerical
	}
	return len(s[i]) > len(s[j]) // descending
}

// Init - step 2 - sorting by length
func init() {
	for k := range hyphm {
		byLen = append(byLen, k)
	}
	sort.Sort(byLen)

	if false {
		for _, k := range byLen {
			log.Printf("%-20v %v", k, hyphm[k])
		}
	}
}

// HyphenizeText replaces "mittelfristig" with "mittel|fristig"
// Hyphenization is done _once_ during creation of the questionare JSON template.
//
// We replace longer keys first,
// to prevent erratic results for example from
//
// desa|cuer|do
func HyphenizeText(s string) string {

	// we have to prevent hyphenization inside URL paths
	if false && strings.Contains(s, "/") {
		return s
	}

	for _, k := range byLen {
		v := hyphm[k]
		s = strings.Replace(s, k, v, -1)
	}
	// if s1 == "Hochschulabschluss" {
	// 	log.Printf("hyphenate\n%v\n%v", s1, s)
	// }
	return s
}
