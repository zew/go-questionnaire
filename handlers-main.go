package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/tpl"

	"github.com/pkg/errors"
)

func helper(w http.ResponseWriter, err error, msgs ...string) {
	if len(msgs) > 0 {
		err = errors.Wrap(err, msgs[0])
	}
	log.Print(err)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func loadQuestionaire(w http.ResponseWriter, r *http.Request) (*qst.QuestionaireT, error) {

	sess := sessx.New(w, r)

	//
	// Load questionaire from file or from session
	var q = &qst.QuestionaireT{}
	ok, err := sess.EffectiveObj("questionaire", q)
	if err != nil {
		err = errors.Wrap(err, "Reading questionaire from session caused error")
		return q, err
	}
	if ok {
		log.Printf("Questionaire loaded from session; %v pages", len(q.Pages))
	} else {
		q, err = qst.Load("questionaire.json")
		if err != nil {
			err = errors.Wrap(err, "Loading questionaire from file caused error")
			return q, err
		}
		err = q.Validate()
		if err != nil {
			err = errors.Wrap(err, "Questionaire validation caused error")
			return q, err
		}
		log.Printf("Questionaire loaded from file; %v pages", len(q.Pages))
	}
	return q, nil

}

func reloadH(w http.ResponseWriter, r *http.Request) {
	sess := sessx.New(w, r)
	sess.Remove(w, "questionaire")
}

func mainH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	q, err := loadQuestionaire(w, r)
	if err != nil {
		helper(w, err)
		return
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
	// Page logic
	//
	// contains currPage from last request
	// remember, because we want to store request values *there*
	prevPage := q.CurrPage
	if prevPage > len(q.Pages)-1 || prevPage < 0 {
		q.CurrPage = 0
		prevPage = 0
	}
	currPage := prevPage // Default assumption: we are still on prev page - unless there is some modification:
	submit := sess.EffectiveStr("submitBtn")
	log.Printf("submitBtn is '%v'", submit)
	if submit == "prev" {
		currPage = q.Prev()
	}
	if submit == "next" {
		currPage = q.Next()
	}
	explicit, ok, err := sess.EffectiveInt("page")
	if err != nil {
		// invalid page value, just dont use it
	}
	if ok && err == nil && explicit > -1 {
		log.Printf("curPage set explicitly to %v", explicit)
		currPage = explicit
	}
	q.CurrPage = currPage // Put current page into questionaire

	//
	// Put request values into questionaire
	if q.Pages[prevPage].Finished.IsZero() {
		q.Pages[prevPage].Finished = time.Now()
	}
	for i1 := 0; i1 < len(q.Pages[prevPage].Groups); i1++ {
		for i2 := range q.Pages[prevPage].Groups[i1].Inputs {
			nm := q.Pages[prevPage].Groups[i1].Inputs[i2].Name
			ok := sess.EffectiveIsSet(nm)
			if ok {
				val := sess.EffectiveStr(nm)
				log.Printf("(Page#%2v) Setting '%v' to '%v'", prevPage, nm, val)
				q.Pages[prevPage].Groups[i1].Inputs[i2].Response = val
			}
		}
	}
	err = q.ValidateReponseData(prevPage, q.LangCode)
	if err != nil {
		q.CurrPage = prevPage // Prevent changing page, keep user on page with errors
	}

	//
	//
	// Save questionaire into session
	err = sess.PutObject("questionaire", q)
	if err != nil {
		helper(w, err, "Putting questionaire into session caused error")
		return
	}

	//
	// Save questionaire to file
	err = q.Save(fmt.Sprintf("tmp_sess_%02d", time.Now().Minute()/10))
	if err != nil {
		helper(w, err, "Putting questionaire into session caused error")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tplBundle := tpl.Get(w, r, "main.html")
	ts := &tpl.StackT{"main.html", "content1.html"}
	ts = &tpl.StackT{"content1.html"}

	err = tplBundle.Execute(
		w,
		TplDataT{
			TplBundle: tplBundle,
			TS:        ts,

			Trls: cfg.Get().Trls,
			Sess: &sess,

			Q: q,
		},
	)
	if err != nil {
		helper(w, err, "Executing template caused error")
		return
	}

}
