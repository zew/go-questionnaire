package tpl

import (
	"html/template"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/lng"
	"github.com/zew/go-questionaire/sessx"
)

// TplDataT is a conduit for templates to access request, session and application data
// It is meant to be embedded/extended by various apps
type TplDataT struct {
	TplBundle *template.Template // A bundle of compiled templates, so we can can executeTemplate(TplBundle, name, data) without independently of the request
	TS        *StackT            // Stack of template names to pop from

	// Access to session and request values
	// Session also transmits the language via lang_code to main.html
	// => Session must be set
	// Session can also be used as scrapbook in subtemplates
	Sess *sessx.SessT
	L    *lgn.LoginT // Yes, we could retrieve it from the session but it is cumbersome in template lingo

	Cnt string // Alternative - just a string

	// After embedding, we would add the major app specific object.
	// For example Q   *qst.QuestionaireT
}

// Trls returns translated strings, for instance HtmlTitle
func (t TplDataT) Trls() lng.TrlsT {
	return cfg.Get().Trls
}
