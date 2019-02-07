package qst

import (
	"fmt"
	"log"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/trl"
)

type dynFuncT func(*QuestionnaireT) (string, error)

var dynFuncs = map[string]dynFuncT{
	"RepsonseStatistics": RepsonseStatistics,
	"PersonalLink":       PersonalLink,
	"HasEuroQuestion":    ResponseTextHasEuro,
}

// Statistics returns the percentage of
// answers responded to.
// It is helper to RepsonseStatistics().
func (q *QuestionnaireT) Statistics() (int, int, float64) {
	responses := 0
	inputs := 0
	for _, p := range q.Pages {
		for _, gr := range p.Groups {
			for _, i := range gr.Inputs {
				if i.IsLayout() {
					continue
				}
				if i.Type == "textarea" {
					continue
				}
				inputs++
				if i.Response != "" && i.Response != "0" {
					responses++
				}
			}
		}

	}
	return responses, inputs, 100 * float64(responses) / float64(inputs)
}

// RepsonseStatistics returns the percentage of
// answers responded to.
func RepsonseStatistics(q *QuestionnaireT) (string, error) {

	responses, inputs, pct := q.Statistics()
	ct := q.Survey.Deadline
	// ct = ct.Truncate(time.Hour)
	cts := ct.Format("02.01.2006 15:04")
	nextDay := q.Survey.Deadline.Add(24 * time.Hour)
	nextDayS := nextDay.Format("02.01.2006")

	s1 := fmt.Sprintf(cfg.Get().Mp["percentage_answered"].Tr(q.LangCode), responses, inputs, pct)
	s2 := fmt.Sprintf(cfg.Get().Mp["survey_ending"].Tr(q.LangCode), cts, nextDayS)
	ret := s1 + s2
	// log.Print("RepsonseStatistics: " + ret)
	return ret, nil
}

// PersonalLink returns the entry link
func PersonalLink(q *QuestionnaireT) (string, error) {
	closed := q.FinishedEntirely()
	ret := ""
	if closed {
		ret = cfg.Get().Mp["finished_by_participant"].Tr(q.LangCode)
		ret = fmt.Sprintf(ret, q.ClosingTime.Format("02.01.2006 15:04"))
	} else {
		ret = cfg.Get().Mp["review_by_personal_link"].Tr(q.LangCode)
	}
	log.Printf("PersonalLink: closed is %v", closed)
	return ret, nil
}

// ResponseTextHasEuro yields texts => want to keep € - want to have €
func ResponseTextHasEuro(q *QuestionnaireT) (string, error) {

	attr1 := q.Attrs["euro-member"]
	attr2 := q.Attrs["country"]

	hl := trl.S{
		"de": "Wirtschaftlicher Nutzen des Euro<br>",
		"en": "Economic benefits of euro<br>",
		"fr": "Avantages économiques de l'euro<br>",
		"it": "Benefici economici dell'Euro<br>",
	}
	desc := ""
	ret := ""

	if attr1 == "yes" {
		s1 := trl.S{
			"de": fmt.Sprintf("Den Euro in %v als die offizielle Währung zu haben, ist wirtschaftlich vorteilhaft.", attr2),
			"en": fmt.Sprintf("Having the euro in %v as the official currency is economically beneficial.", attr2),
			"fr": fmt.Sprintf("Avoir l'euro dans %v comme monnaie officielle est économiquement avantageux.", attr2),
			"it": fmt.Sprintf("Avere l'Euro come valuta ufficiale nel %v è economicamente vantaggioso.", attr2),
		}
		desc = s1[q.LangCode]

	} else {
		s1 := trl.S{
			"de": fmt.Sprintf("Den Euro in %v als offizielle Währung einzuführen, wäre wirtschaftlich vorteilhaft. ", attr2),
			"en": fmt.Sprintf("Introducing the euro in %v as the official currency would be economically beneficial.", attr2),
			"fr": fmt.Sprintf("L'introduction de l'euro dans %v en tant que monnaie officielle serait économiquement avantageuse.", attr2),
			"it": fmt.Sprintf("Introdurre l'Euro come valuta ufficiale nel %v sarebbe economicamente vantaggioso.", attr2),
		}
		desc = s1[q.LangCode]
	}

	ret = fmt.Sprintf("<b> %v </b> %v", hl[q.LangCode], desc)

	return ret, nil

}
