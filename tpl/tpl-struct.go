package tpl

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/trl"
)

// TplDataT is a conduit for templates to access request, session and application data
// It is meant to be embedded/extended by various apps
type TplDataT struct {
	TplBundle *template.Template // A bundle of compiled templates, so we can can executeTemplate(TplBundle, name, data) without independently of the request
	// TS        *StackT            // Stack of template names to pop from

	// Access to session and request values
	// Session also transmits the language via lang_code to main.html
	// => Session must be set
	// Session can also be used as scrapbook in subtemplates
	Sess *sessx.SessT
	L    *lgn.LoginT // Yes, we could retrieve it from the session but it is cumbersome in template lingo

	Cnt string // Alternative - just a string

	// After embedding, we would add the major app specific object.
	// For example Q   *qst.QuestionnaireT
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
