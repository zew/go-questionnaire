package main

import (
	"html/template"
	"net/url"

	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/tpl"
)

// TplDataT is a conduit for templates to access request, session and application data
type TplDataT struct {
	TplBundle *template.Template // A bundle of compiled templates, so we can can executeTemplate(TplBundle, name, data)
	TS        *tpl.StackT        // Stack of template names to pop from

	App  url.Values   // App data, for instance HtmlTitle
	Sess *sessx.SessT // Access to session and request values
	L    *lgn.LoginT  // Yes, we could retrieve it from the session but it is cumbersome in template lingo

	Q *qst.QuestionaireT // The major app specific object
}
