package trl

import (
	"log"
	"sort"
	"strings"
)

//
//
// Hyphenization
// =================
//
// hyph is a slice _with_ hyphenized words.
// During application initialization we use it to fill hyphm above.
var hyph = []string{

	// german
	"An|gabe",
	"Arbeits|markt|öko|no|mie",
	"as|so|zi|ie|rt",
	"Aus|fall|risiken",
	"aus|ge|schlos|sen",
	"ab|wer|ten",  // ab|wer|ten
	"auf|wer|ten", // auf|wer|ten
	"Bil|dung",    // 	"Bil|dung",
	"blei|ben",    // blei|ben
	"Da|ten",      // Da|ten
	"da|für",
	"da|ge|gen",
	"Deutsch|land",
	"Do|zen|tIn",
	"Dok|to|ran|din",
	"Ein|fluss", // Ein|fluss
	"Ent|schei|dungs|pro|zesse",
	"Ent|wick|lungs|öko|no|mie",
	"er|stre|bens|wert",
	"Erwartungs|wert",
	"Euro|raum",
	"Ex|port|markt",  // "Ex|port|markt
	"Ex|port|märkte", // "Ex|port|märkte
	"Fi|nanz|wis|sen|schaft",
	"Finanz|wirt|schaft",
	"ge|samt",        // ge|samt
	"Geld|po|li|tik", //
	"Groß|unter|nehmen",
	"Handels|konflikte", //
	"Immob|ilien|kredite",
	"In|dus|trie|öko|no|mie",
	"Junior|pro|fes|sorIn",
	"kom|plett",
	"Kon|junk|tur|da|ten ", // Kon|junk|tur
	"Kon|sum|enten|kredite",
	"Kredit|angebot",
	"Kredit|nach|frage",
	"lang|fristig",
	"Ma|kro|öko|no|mie",
	"Mi|kro|öko|no|mie",
	"mittel|fristig", // mit|tel|fris|tig - reduced
	"ne|ga|tiv",      // ne|ga|tiv
	"neu|tral",       // neu|tral
	"nied|rig",       // nied|rig
	"nor|mal",        // nor|mal
	"Öko|no|mie",
	"Po|li|tik", // Po|li|tik
	"po|si|tiv", // po|si|tiv
	"Prä|ferenz|konstellation",
	"Pro|fes|sorIn",
	"Re|finanz|ierung",
	"Re|gie|rung",           // 	"Re|gie|rung",
	"Re|gie|rungs|bil|dung", // Re|gie|rungs|bil|dung
	"Regierungs|bildung",
	"Regu|lierung",
	"ri|si|ko|be|reit",
	"Ri|si|ko|ver|mei|den",
	"Risiko|trag|fähig|keit",
	"Roh|stoff|preise",
	"schlech|teste",
	"si|cher",
	"sin|ken",  // sin|ken
	"stei|gen", // stei|gen
	"Stif|tung",
	"Te|le|kom|mu|ni|ka|ti|on", // Te|le|kom|mu|ni|ka|ti|on
	"über|haupt",
	"Um|welt|öko|no|mie",
	"un|ent|schieden",
	"Un|ter|neh|men", // Un|ter|neh|men
	"un|wich|tig",
	"ver|än|dern",     // "ver|än|dern",
	"Ver|bes|se|rung", // Ver|bes|se|rung
	"ver|bes|sern",    // "ver|bes|sern",
	"ver|füg|bar",
	"Ver|schlech|te|rung", // Ver|schlech|te|rung
	"ver|schlech|tern",    // "ver|schlech|tern",
	"Wechsel|kurse",       //
	"Welt|wirt|schaft",
	"Wett|be|werbs|sit|uation",
	"wich|tig",
	"wirt|schaft",      // wirt|schaft
	"wirt|schaft|lich", // wirt|schaft|lich
	"Wirt|schafts|politik",

	"Kompromiss|lösung",
	"Mehrheits|lösung",
	"mitt|le|re", // "mittlere"
	"Grup|pe",
	"Grup|pen",
	"Mit|glied",
	"Mit|glieder",
	"Grup|pen|mit|glieder", // Gruppenmitglieder

	// english
	"ac|counts",
	"ad|min|is|tra|tion",
	"an|swer",
	"ap|pli|ca|ble",
	"ap|​pre|​ci|​ate", // ap·​pre·​ci·​ate
	"as|sis|tant",
	"as|so|ci|ate",
	"av|er|age",
	"busi|ness",
	"can|di|date",
	"com|plete|ly",
	"Comp|etitive",
	"Cons|umer",
	"de|crease",
	"de|​pre|​ci|​ate", // de·​pre·​ci·​ate
	"de|te|ri|o|rate",
	"de|vel|op|ment",
	"Dis|agree",
	"eco|nom|ic",
	"eco|nom|ics",
	"econ|o|my",
	"en|vi|ron|ment",
	"en|vi|ron|mental",
	"enter|prises",
	"environ|ment",
	"im|prove",
	"in|crease",
	"in|dif|fer|ent",
	"in|dus|tri|al",
	"in|flu|ence",
	"in|sti|tute",
	"in|sti|tutes",
	"in|ter|na|tion|al",
	"in|vest|ment",
	"mac|ro|eco|nom|ics",
	"mi|cro|eco|nom|ics",
	"na|tion|al|i|ty",
	"neg|ative",
	"nei|ther",
	"op|ti|mis|tic",
	"pes|si|mis|tic",
	"pos|itive",
	"pro|fes|sor",
	"Re|financing",
	"re|main",
	"re|search",
	"Reg|ulation",
	"sav|ings",
	"Small+|medium",
	"strong|ly",
	"strong|ly",
	"un|changed",
	"Un|decided",
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
//
// We replace longer keys first,
// to prevent erratic results for example from
//
// desa|cuer|do
//         acuer|do
//
func HyphenizeText(s string) string {
	s1 := s

	for _, k := range byLen {
		v := hyphm[k]
		s = strings.Replace(s, k, v, -1)
	}
	if s1 == "Mittlere Stiftung" {
		log.Printf("hyphenate\n%v\n%v", s1, s)
	}
	return s
}
