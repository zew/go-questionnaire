package main

import (
	"log"
	"net/http"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/tpl"

	"github.com/pkg/errors"
)

func mainH(w http.ResponseWriter, r *http.Request) {

	helper := func(err error, msg string) {
		err = errors.Wrap(err, msg)
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	sess := sessx.New(w, r)

	//
	// Load questionaire from file or from session
	var q = &qst.QuestionaireT{}
	ok, err := sess.EffectiveObj("questionaire", q)
	if err != nil {
		helper(err, "Reading questionaire from session caused error")
		return
	}
	if ok {
		log.Printf("Questionaire loaded from session; %v pages", len(q.Pages))
	} else {
		q, err = qst.Load("questionaire.json")
		if err != nil {
			helper(err, "Loading questionaire from file caused error")
			return
		}
		err = q.Validate()
		if err != nil {
			helper(err, "Questionaire validation caused error")
			return
		}
		log.Printf("Questionaire loaded from file; %v pages", len(q.Pages))
	}

	//
	// Change page logic
	prevPage, ok, err := sess.EffectiveInt("curr_page")
	if err != nil {
		helper(err, "Reading request parameter caused error")
		return
	}
	submit := sess.EffectiveStr("submit")
	log.Printf("submit is '%v'", submit)
	currPage := q.CurrPage
	if submit == "prev" {
		currPage = q.Prev()
	}
	if submit == "next" {
		currPage = q.Next()
	}
	q.CurrPage = currPage // Put current page into questionaire

	//
	// Put request values into questionaire
	if q.Pages[prevPage].Finished.IsZero() {
		q.Pages[prevPage].Finished = time.Now()
	}
	for i1 := 0; i1 < len(q.Pages[prevPage].Elements); i1++ {
		for i2 := range q.Pages[prevPage].Elements[i1].Members {

			nm := q.Pages[prevPage].Elements[i1].Members[i2].Name
			ok := sess.EffectiveIsSet(nm)
			if ok {
				val := sess.EffectiveStr(nm)
				log.Printf("(Page#%2v) Setting '%v' to '%v'", prevPage, nm, val)
				q.Pages[prevPage].Elements[i1].Members[i2].Response = val
			}

		}
	}

	//
	// Meta parameters
	if newCode, ok := sess.ReqParam("lang_code"); ok {
		oldCode := q.LangCode
		q.LangCode = newCode
		err := q.Validate()
		if err != nil {
			q.LangCode = oldCode
		}
	}

	//
	// Save questionaire into session
	err = sess.PutObject("questionaire", q)
	if err != nil {
		helper(err, "Putting questionaire into session caused error")
		return
	}

	content := ""
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tpl.Get(w, r).Execute(
		w,
		TplDataT{
			HtmlTitle:    cfg.Get().AppName,
			CntBefore:    content,
			CntAfter:     content,
			TemplateName: "ct01.html",
			Q:            q,
		},
	)
	if err != nil {
		helper(err, "Executing template caused error")
		return
	}

}
