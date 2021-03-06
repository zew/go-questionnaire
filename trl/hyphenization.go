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
	"An&shy;gabe",
	"Aus&shy;fall&shy;risiken",
	"aus&shy;ge&shy;schlos&shy;sen",
	"Bil&shy;dung", // 	"Bil|dung",
	"blei&shy;ben", // blei|ben
	"Da&shy;ten",   // Da|ten
	"da&shyfür",
	"da&shyge&shygen",
	"Deutsch&shy;land",
	"Ein&shy;fluss", // Ein|fluss
	"er&shy;stre&shy;bens&shy;wert",
	"Erwartungs&shy;wert",
	"Euro&shy;raum",
	"Ex&shy;port&shy;markt",      // "Ex|port|markt
	"Ex&shy;port&shy;märkte",     // "Ex|port|märkte
	"Geld&shy;po&shy;li&shy;tik", //
	"ge&shy;samt",                // ge|samt
	"Groß&shy;unter&shy;nehmen",
	"Immob&shy;ilien&shy;kredite",
	"kom&shy;plett",
	"Kon&shy;junk&shy;tur&shy;da&shy;ten ", // Kon|junk|tur
	"Kon&shy;sum&shy;enten&shy;kredite",
	"Kredit&shy;angebot",
	"Kredit&shy;nach&shy;frage",
	"lang&shy;fristig",
	"mittel&shy;fristig", // mit|tel|fris|tig - reduced
	"ne&shy;ga&shy;tiv",  // ne|ga|tiv
	"neu&shy;tral",       // neu|tral
	"nied&shy;rig",       // nied|rig
	"nor&shy;mal",        // nor|mal
	"Po&shy;li&shy;tik",  // Po|li|tik
	"po&shy;si&shy;tiv",  // po|si|tiv
	"Re&shy;finanz&shy;ierung",
	"Re&shy;gie&shy;rung",                   // 	"Re|gie|rung",
	"Re&shy;gie&shy;rungs&shy;bil&shy;dung", // Re|gie|rungs|bil|dung
	"Regierungs&shy;bildung",
	"Regu&shy;lierung",
	"Risiko&shy;trag&shy;fähig&shy;keit",
	"Roh&shy;stoff&shy;preise",
	"schlech&shy;teste",
	"si&shy;cher",
	"sin&shy;ken",  // sin|ken
	"stei&shy;gen", // stei|gen
	"über&shy;haupt",
	"un&shy;ent&shy;schieden",
	"un&shy;wich&shy;tig",
	"Un&shy;ter&shy;neh&shy;men",      // Un|ter|neh|men
	"ver&shy;än&shy;dern",             // "ver|än|dern",
	"ver&shy;bes&shy;sern",            // "ver|bes|sern",
	"ver&shy;schlech&shy;tern",        // "ver|schlech|tern",
	"Ver&shy;bes&shy;se&shy;rung",     // Ver|bes|se|rung
	"Ver&shy;schlech&shy;te&shy;rung", // Ver|schlech|te|rung
	"Wechsel&shy;kurse",               //
	"wirt&shy;schaft&shy;lich",        // wirt|schaft|lich
	"wirt&shy;schaft",                 // wirt|schaft
	"Welt&shy;wirt&shy;schaft",
	"Wett&shy;be&shy;werbs&shy;sit&shy;uation",
	"wich&shy;tig",

	// german euref
	"Öko&shy;no&shy;mie",
	"Mi&shy;kro&shy;öko&shy;no&shy;mie",
	"Ma&shy;kro&shy;öko&shy;no&shy;mie",
	"Fi&shy;nanz&shy;wis&shy;sen&shy;schaft",
	"Arbeits&shy;markt&shy;öko&shy;no&shy;mie",
	"Finanz&shy;wirt&shy;schaft",
	"Wirt&shy;schafts&shy;politik",
	"Ent&shy;wick&shy;lungs&shy;öko&shy;no&shy;mie",
	"Um&shy;welt&shy;öko&shy;no&shy;mie",
	"In&shy;dus&shy;trie&shy;öko&shy;no&shy;mie",
	"Pro&shy;fes&shy;sorIn",
	"Junior&shy;pro&shy;fes&shy;sorIn",
	"as&shy;so&shy;zi&shy;ie&shy;rt",
	"Do&shy;zen&shy;tIn",
	"Dok&shy;to&shy;ran&shy;din",

	// english
	"Small+&shy;medium",
	"enter&shy;prises",
	"in&shy;crease",
	"de&shy;crease",
	"un&shy;changed",
	"in&shy;fluence",
	"strong&shy;ly",
	"pos&shy;itive",
	"neg&shy;ative",
	"Re&shy;financing",
	"Comp&shy;etitive",
	"environ&shy;ment",
	"Cons&shy;umer",
	"Reg&shy;ulation",
	"Dis&shy;agree",
	"Un&shy;decided",
	// english euref:
	"na&shy;tion&shy;al&shy;i&shy;ty",
	"mi&shy;cro&shy;eco&shy;nom&shy;ics",
	"mac&shy;ro&shy;eco&shy;nom&shy;ics",
	"eco&shy;nom&shy;ics",
	"econ&shy;o&shy;my",
	"eco&shy;nom&shy;ic",
	"de&shy;vel&shy;op&shy;ment",
	"in&shy;ter&shy;na&shy;tion&shy;al",
	"en&shy;vi&shy;ron&shy;ment",
	"en&shy;vi&shy;ron&shy;mental",
	"in&shy;dus&shy;tri&shy;al",
	"busi&shy;ness",
	"ad&shy;min&shy;is&shy;tra&shy;tion",
	"pro&shy;fes&shy;sor",
	"as&shy;sis&shy;tant",
	"as&shy;so&shy;ci&shy;ate",
	"re&shy;search",
	"in&shy;sti&shy;tute",
	"in&shy;sti&shy;tutes",
	"can&shy;di&shy;date",
	// spanish
	"in­&shy;deciso",
	"acuer&shy;do",
	"desa&shy;cuer&shy;do",
	// french
	"acc&shy;ord",
	"In&shy;diff&shy;érent",
	// italian
	"favo&shy;revole",
	"Favo&shy;revole",
	"in&shy;deciso",
	"In&shy;deciso",

	// financial literacy
	"ap&shy;pli&shy;ca&shy;ble",
	"in&shy;vest&shy;ment",
	"sav&shy;ings",
	"ac&shy;counts",
	"av&shy;er&shy;age",
	"an&shy;swer",
	"com&shy;plete&shy;ly",
	"strong&shy;ly",
	"op&shy;ti&shy;mis&shy;tic",
	"pes&shy;si&shy;mis&shy;tic",
	"in&shy;dif&shy;fer&shy;ent",
	"will&shy;ing",
	"nei&shy;ther",

	// paternalism
	"Ent&shy;schei&shy;dungs&shy;pro&shy;zesse",
	// "Ver&shy;füg&shy;bar",
	"ver&shy;füg&shy;bar",
	"Ri&shy;si&shy;ko&shy;ver&shy;mei&shy;den",
	// "Ri&shy;si&shy;ko&shy;be&shy;reit",
	"ri&shy;si&shy;ko&shy;be&shy;reit",
	"Prä&shy;ferenz&shy;konstellation",
	"Stif&shy;tung",
}

// hyphm is a map with words and their hyphenized form as value.
// hyphm is filled during app initialization from hyph above.
var hyphm = map[string]string{}

func init() {
	for _, v := range hyph {
		if len(v) < 1 {
			continue
		}
		// lower case word
		v1 := strings.ToLower(v)
		v1a := strings.Replace(v1, "&shy;", "", -1)
		hyphm[v1a] = v1
		// capitalize first UTF8 rune
		v2 := strings.ToUpper(string([]rune(v)[:1]))
		v2 += string([]rune(v)[1:]) // remainder
		v2a := strings.Replace(v2, "&shy;", "", -1)
		hyphm[v2a] = v2

		// logx.Printf("%-20v %v", v1a, v1)
		// logx.Printf("%-20v %v", v2a, v2)
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

// HyphenizeText replaces "mittelfristig" with "mittel&shy;fristig"
// Hyphenization is done _once_ during creation of the questionare JSON template.
//
//
// We replace longer keys first,
// to prevent erratic results for example from
//
// desa&shy;cuer&shy;do
//         acuer&shy;do
//
func HyphenizeText(s string) string {
	for _, k := range byLen {
		v := hyphm[k]
		s = strings.Replace(s, k, v, -1)
	}
	return s
}
