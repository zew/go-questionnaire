package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
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
	// Change page logic
	prevPage, ok, err := sess.EffectiveInt("curr_page")
	if err != nil {
		helper(w, err, "Reading request parameter caused error")
		return
	}
	if !ok {
		prevPage = 0
	}

	submit := sess.EffectiveStr("submit")
	log.Printf("submit is '%v'", submit)
	currPage := q.CurrPage
	if currPage > len(q.Pages)-1 || currPage < 0 {
		q.CurrPage = 0
		currPage = 0
	}

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

	appData := url.Values{}
	appData.Set("HtmlTitle", cfg.Get().AppName)

	err = tplBundle.Execute(
		w,
		TplDataT{
			TplBundle: tplBundle,
			TS:        ts,

			App:  appData,
			Sess: &sess,

			Q: q,
		},
	)
	if err != nil {
		helper(w, err, "Executing template caused error")
		return
	}

}
