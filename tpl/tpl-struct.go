package tpl

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/trl"
)

// TplDataT is just a bunch of funcs;
// made available to templates,
// by adding an instance of TplDataT as "MoreFuncs" in Exec()
type TplDataT struct {
	// Q       *qst.QuestionnaireT // The major app specific object
	// Content string              // Alternative - just a string
}

// Trls returns translated strings, for instance HtmlTitle
func (t TplDataT) Trls() trl.Map {
	return cfg.Get().Mp
}

// DefaultLangCode returns the language code to be used,
// if no participant preference was set.
func (t TplDataT) DefaultLangCode() string {
	return cfg.Get().LangCodes[0]
}

// LanguageChooser renders unspecific languages
// taken from app instance config
func (t TplDataT) LanguageChooser(uri string, curr ...string) string {
	lcs := cfg.Get().LangCodes
	return t.LanguageChooserExplicit(lcs, uri, curr...)
}

// LanguageChooserExplicit renders a HTML language chooser
// If no current language is specified,
// then the default language is chosen.
// The URI is supposed to contain the app url prefix
func (t TplDataT) LanguageChooserExplicit(lcs []string, uri string, curr ...string) string {

	if lcs == nil {
		lcs = cfg.Get().LangCodes // necessary; member q might exist, but be nil
	}

	if len(lcs) == 1 {
		return "" // only one language exists - choosing in pointless
	}

	currCode := lcs[0]
	if len(curr) > 0 {
		currCode = curr[0]
	}

	s := []string{}
	for _, key := range lcs {
		keyCap := strings.Title(key)
		if key == currCode {
			s = append(s, fmt.Sprintf("<b           title='%v'>%v</b>\n", key, keyCap))
		} else {
			uriExt := uri + "?lang_code=" + key
			s = append(s, fmt.Sprintf("<a href='%v' title='%v'>%v</a>\n", uriExt, key, keyCap))
		}
	}
	return strings.Join(s, "  |  ")

}
