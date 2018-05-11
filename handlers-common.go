package main

import "github.com/zew/go-questionaire/qst"

// Template Data
type TplDataT struct {
	TemplateName string
	HtmlTitle    string
	CntBefore    interface{}
	CntAfter     interface{}
	Q            *qst.QuestionaireT
}
