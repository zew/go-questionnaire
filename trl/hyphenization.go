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
	"Arbeits&shy;markt&shy;öko&shy;no&shy;mie",
	"as&shy;so&shy;zi&shy;ie&shy;rt",
	"Aus&shy;fall&shy;risiken",
	"aus&shy;ge&shy;schlos&shy;sen",
	"Bil&shy;dung", // 	"Bil|dung",
	"blei&shy;ben", // blei|ben
	"Da&shy;ten",   // Da|ten
	"da&shyfür",
	"da&shyge&shygen",
	"Deutsch&shy;land",
	"Do&shy;zen&shy;tIn",
	"Dok&shy;to&shy;ran&shy;din",
	"Ein&shy;fluss", // Ein|fluss
	"Ent&shy;schei&shy;dungs&shy;pro&shy;zesse",
	"Ent&shy;wick&shy;lungs&shy;öko&shy;no&shy;mie",
	"er&shy;stre&shy;bens&shy;wert",
	"Erwartungs&shy;wert",
	"Euro&shy;raum",
	"Ex&shy;port&shy;markt",  // "Ex|port|markt
	"Ex&shy;port&shy;märkte", // "Ex|port|märkte
	"Fi&shy;nanz&shy;wis&shy;sen&shy;schaft",
	"Finanz&shy;wirt&shy;schaft",
	"ge&shy;samt",                // ge|samt
	"Geld&shy;po&shy;li&shy;tik", //
	"Groß&shy;unter&shy;nehmen",
	"Handels&shy;konflikte", //
	"Immob&shy;ilien&shy;kredite",
	"In&shy;dus&shy;trie&shy;öko&shy;no&shy;mie",
	"Junior&shy;pro&shy;fes&shy;sorIn",
	"kom&shy;plett",
	"Kon&shy;junk&shy;tur&shy;da&shy;ten ", // Kon|junk|tur
	"Kon&shy;sum&shy;enten&shy;kredite",
	"Kredit&shy;angebot",
	"Kredit&shy;nach&shy;frage",
	"lang&shy;fristig",
	"Ma&shy;kro&shy;öko&shy;no&shy;mie",
	"Mi&shy;kro&shy;öko&shy;no&shy;mie",
	"mittel&shy;fristig", // mit|tel|fris|tig - reduced
	"ne&shy;ga&shy;tiv",  // ne|ga|tiv
	"neu&shy;tral",       // neu|tral
	"nied&shy;rig",       // nied|rig
	"nor&shy;mal",        // nor|mal
	"Öko&shy;no&shy;mie",
	"Po&shy;li&shy;tik", // Po|li|tik
	"po&shy;si&shy;tiv", // po|si|tiv
	"Prä&shy;ferenz&shy;konstellation",
	"Pro&shy;fes&shy;sorIn",
	"Re&shy;finanz&shy;ierung",
	"Re&shy;gie&shy;rung",                   // 	"Re|gie|rung",
	"Re&shy;gie&shy;rungs&shy;bil&shy;dung", // Re|gie|rungs|bil|dung
	"Regierungs&shy;bildung",
	"Regu&shy;lierung",
	"ri&shy;si&shy;ko&shy;be&shy;reit",
	"Ri&shy;si&shy;ko&shy;ver&shy;mei&shy;den",
	"Risiko&shy;trag&shy;fähig&shy;keit",
	"Roh&shy;stoff&shy;preise",
	"schlech&shy;teste",
	"si&shy;cher",
	"sin&shy;ken",  // sin|ken
	"stei&shy;gen", // stei|gen
	"Stif&shy;tung",
	"Te&shy;le&shy;kom&shy;mu&shy;ni&shy;ka&shy;ti&shy;on", // Te|le|kom|mu|ni|ka|ti|on
	"über&shy;haupt",
	"Um&shy;welt&shy;öko&shy;no&shy;mie",
	"un&shy;ent&shy;schieden",
	"Un&shy;ter&shy;neh&shy;men", // Un|ter|neh|men
	"un&shy;wich&shy;tig",
	"ver&shy;än&shy;dern",         // "ver|än|dern",
	"Ver&shy;bes&shy;se&shy;rung", // Ver|bes|se|rung
	"ver&shy;bes&shy;sern",        // "ver|bes|sern",
	"ver&shy;füg&shy;bar",
	"Ver&shy;schlech&shy;te&shy;rung", // Ver|schlech|te|rung
	"ver&shy;schlech&shy;tern",        // "ver|schlech|tern",
	"Wechsel&shy;kurse",               //
	"Welt&shy;wirt&shy;schaft",
	"Wett&shy;be&shy;werbs&shy;sit&shy;uation",
	"wich&shy;tig",
	"wirt&shy;schaft",          // wirt|schaft
	"wirt&shy;schaft&shy;lich", // wirt|schaft|lich
	"Wirt&shy;schafts&shy;politik",

	// english
	"ac&shy;counts",
	"ad&shy;min&shy;is&shy;tra&shy;tion",
	"an&shy;swer",
	"ap&shy;pli&shy;ca&shy;ble",
	"as&shy;sis&shy;tant",
	"as&shy;so&shy;ci&shy;ate",
	"av&shy;er&shy;age",
	"busi&shy;ness",
	"can&shy;di&shy;date",
	"com&shy;plete&shy;ly",
	"Comp&shy;etitive",
	"Cons&shy;umer",
	"de&shy;crease",
	"de&shy;te&shy;ri&shy;o&shy;rate",
	"de&shy;vel&shy;op&shy;ment",
	"Dis&shy;agree",
	"eco&shy;nom&shy;ic",
	"eco&shy;nom&shy;ics",
	"econ&shy;o&shy;my",
	"en&shy;vi&shy;ron&shy;ment",
	"en&shy;vi&shy;ron&shy;mental",
	"enter&shy;prises",
	"environ&shy;ment",
	"im&shy;prove",
	"in&shy;crease",
	"in&shy;dif&shy;fer&shy;ent",
	"in&shy;dus&shy;tri&shy;al",
	"in&shy;flu&shy;ence",
	"in&shy;sti&shy;tute",
	"in&shy;sti&shy;tutes",
	"in&shy;ter&shy;na&shy;tion&shy;al",
	"in&shy;vest&shy;ment",
	"mac&shy;ro&shy;eco&shy;nom&shy;ics",
	"mi&shy;cro&shy;eco&shy;nom&shy;ics",
	"na&shy;tion&shy;al&shy;i&shy;ty",
	"neg&shy;ative",
	"nei&shy;ther",
	"op&shy;ti&shy;mis&shy;tic",
	"pes&shy;si&shy;mis&shy;tic",
	"pos&shy;itive",
	"pro&shy;fes&shy;sor",
	"Re&shy;financing",
	"re&shy;main",
	"re&shy;search",
	"Reg&shy;ulation",
	"sav&shy;ings",
	"Small+&shy;medium",
	"strong&shy;ly",
	"strong&shy;ly",
	"un&shy;changed",
	"Un&shy;decided",
	"will&shy;ing",

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
