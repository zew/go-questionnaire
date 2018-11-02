package trl

import (
	"sort"
	"strings"
)

// Map - Translations Type
// Usage in templates
// 		{{.Trls.imprint.en                     }}  // directly accessing a specific translation; chaining the map keys
// 		{{.Trls.imprint.Tr       .Sess.LangCode}}  // using .Tr(langCode)
// 		{{.Trls.imprint.TrSilent .Sess.LangCode}}  //
type Map map[string]S

//
//
// Hyphenization
// =================
//
// hyph is a slice _with_ hyphenized words.
// During application initialization we use it to fill hyphm above.
var hyph = []string{

	"ver&shy;schlechtern",
	// "saison&shy;bereinigt",

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
// hyphm is filled during app initialization from hyph below.
var hyphm = map[string]string{}

func init() {
	for _, v := range hyph {
		hyphm[strings.Replace(v, "&shy;", "", -1)] = v
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
	return len(s[i]) > len(s[j]) // descending
}

func init() {
	for k := range hyphm {
		byLen = append(byLen, k)
	}
	sort.Sort(byLen)
	// log.Printf(util.IndentedDump(byLen))
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
