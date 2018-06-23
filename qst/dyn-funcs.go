package qst

import (
	"fmt"
	"log"
	"time"

	"github.com/zew/go-questionaire/cfg"
)

type dynFuncT func(*QuestionaireT) (string, error)

var dynFuncs = map[string]dynFuncT{
	"RepsonseStatistics": RepsonseStatistics,
	"PersonalLink":       PersonalLink,
}

// Statistics returns the percentage of
// answers responded to
func (q *QuestionaireT) Statistics() (int, int, float64) {
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
// answers responded to
func RepsonseStatistics(q *QuestionaireT) (string, error) {

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
func PersonalLink(q *QuestionaireT) (string, error) {

	closed := false
	for _, p := range q.Pages {
		for _, gr := range p.Groups {
			for _, inp := range gr.Inputs {
				if inp.Name == "finished" {
					if inp.Response == ValSet {
						closed = true
					}
				}
			}

		}
	}

	ret := ""
	if closed {
		ret = cfg.Get().Mp["finished_by_user"].Tr(q.LangCode)
		ret = fmt.Sprintf(ret, q.ClosingTime.Format("02.01.2006 15:04"))
	} else {
		ret = cfg.Get().Mp["review_by_personal_link"].Tr(q.LangCode)
	}
	log.Printf("PersonalLink: closed is %v", closed)
	return ret, nil
}
