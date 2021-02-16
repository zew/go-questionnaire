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
	"kom&shy;plett",
	"schlech&shy;teste",
	"ver&shy;schlechtern",
	"An&shy;gabe",
	"Konjunktur&shy;daten ",
	"Ausfall&shy;risiken",
	"Risiko&shy;trag&shy;fähig&shy;keit",
	"Re&shy;finanz&shy;ierung",
	"Wett&shy;be&shy;werbs&shy;sit&shy;uation",
	"Regu&shy;lierung",
	"Groß&shy;unter&shy;nehmen",
	"Immob&shy;ilien&shy;kredite",
	"Kon&shy;sum&shy;enten&shy;kredite",
	"Regierungs&shy;bildung",
	"Kredit&shy;nach&shy;frage",
	"Kredit&shy;angebot",
	"mittel&shy;fristig",
	"lang&shy;fristig",
	"Deutsch&shy;land",
	"Welt&shy;wirtschaft",
	"un&shy;ent&shy;schieden",
	"über&shy;haupt",
	"er&shy;stre&shy;bens&shy;wert",
	"aus&shy;ge&shy;schlos&shy;sen",
	"si&shy;cher",
	"un&shy;wich&shy;tig",
	"wich&shy;tig",
	"da&shyfür",
	"da&shyge&shygen",
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
