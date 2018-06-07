// Package handlers is not in main,
// because the systemtests must access the handler funcs.
package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/tpl"

	"github.com/pkg/errors"
)

// An extension with questionaire
type tplDataExtT struct {
	tpl.TplDataT
	Q *qst.QuestionaireT // The major app specific object
}

// Loading questionaire.
// First from session.
// Then from file of previous session.
// Finally from template.
func loadQuestionaire(w http.ResponseWriter, r *http.Request, userSurveyID, userWaveID, userID string) (*qst.QuestionaireT, error) {

	sess := sessx.New(w, r)

	log.Printf("Deriving from the login: survey_id %v, wave_id %v, user_id: %v", userSurveyID, userWaveID, userID)

	// from session
	var q = &qst.QuestionaireT{}
	ok, err := sess.EffectiveObj("questionaire", q)
	if err != nil {
		err = errors.Wrap(err, "Reading questionaire from session caused error")
		return q, err
	}
	if ok {
		log.Printf("Questionaire loaded from session; %v pages", len(q.Pages))
		return q, nil
	}

	// from file
	pth := q.FilePath1(filepath.Join(userSurveyID, userWaveID, userID))
	log.Printf("Deriving path: %v", pth)
	q, err = qst.Load1(pth) // previous session
	if err != nil {
		log.Printf("No previous file %v found. Loading new questionaire from file.", pth)
		q, err = qst.Load1(q.FilePath1(userSurveyID)) // new from template
	}
	if err != nil {
		err = errors.Wrap(err, "Loading questionaire from file caused error")
		return q, err
	}
	err = q.Validate()
	if err != nil {
		err = errors.Wrap(err, "Questionaire validation caused error")
		return q, err
	}

	if q.WaveID.SurveyID != userSurveyID {
		err = fmt.Errorf("Logged in for survey %v - but template is for %v", userSurveyID, q.WaveID.SurveyID)
		return q, err
	}
	if q.WaveID.String() != userWaveID {
		err = fmt.Errorf("Logged in for wave %v - but template is for %v", userWaveID, q.WaveID.String())
		return q, err
	}

	log.Printf("Questionaire loaded from file; %v pages", len(q.Pages))
	return q, nil

}

// ReloadH removes the existing questioniare from the session,
// allowing to start anew
func ReloadH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	var q = &qst.QuestionaireT{}
	ok, err := sess.EffectiveObj("questionaire", q)
	if err != nil {
		helper(w, r, err, "Error retrieving questionaire from session")
		return
	}

	if ok {
		err := os.Remove(q.FilePath1())
		if err != nil {
			helper(w, r, err, "Error deleting questionaire file")
			return
		}
	}
	sess.Remove(w, "questionaire")

}

// You can provite
// 1.) an error
// 2.) an error with string to wrap around
// 3.) only a string - which is converted into an error
func helper(w http.ResponseWriter, r *http.Request, err error, msgs ...string) {
	if len(msgs) > 0 {
		if err == nil {
			err = fmt.Errorf(msgs[0])
		} else {
			err = errors.Wrap(err, msgs[0])
		}
	}
	log.Print(err)
	errorH(w, r, err.Error())
}

// MainH loads and displays the questionaire with page and lang_code
func MainH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	ok, err := lgn.LoginByHash(w, r)
	if err != nil {
		log.Printf("Login by hash error 1: %v", err)
		// Don't show the revealing original error
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoginByHash error."
		helper(w, r, nil, s)
		return
	}
	if ok && err == nil {
		sess.Remove(w, "questionaire") // upon successful, possibly new login - remove previous questionaire from session
	}

	l, isLoggedIn, err := lgn.LoggedInCheck(w, r)
	if err != nil {
		log.Printf("Login by hash error 2: %v", err)
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoggedInCheck error."
		helper(w, r, err, s)
		return
	}
	if !isLoggedIn {
		log.Printf("Login by hash error 3: %v", "not logged in")
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "You are not logged in."
		helper(w, r, nil, s)
		return
	}

	userSurveyID := ""
	userWaveID := ""
	for role, val := range l.Roles {
		if role == "survey_id" {
			userSurveyID = val
		}
		if role == "wave_id" {
			userWaveID = val
		}
	}

	token, ok := sess.ReqParam("token")
	if ok {
		err = lgn.ValidateFormToken(token)
		if err != nil {
			helper(w, r, err)
			return
		}
	} else if !ok && r.Method == "POST" {
		helper(w, r, nil, "Missing request token")
		return
	}

	q, err := loadQuestionaire(w, r, userSurveyID, userWaveID, l.User)
	if err != nil {
		helper(w, r, err)
		return
	}
	q.UserID = l.User

	// Already finished?
	if !q.ClosingTime.IsZero() {
		s := cfg.Get().Mp["finished_by_user"].All()
		s = fmt.Sprintf(s, q.ClosingTime.Format("02.01.2006 15:04"), q.ClosingTime.Format("02.01.2006 15:04"))
		helper(w, r, nil, s)
		return
	}

	// Deadline exceeded?
	if time.Now().After(q.WaveID.Deadline) {
		s := cfg.Get().Mp["deadline_exceeded"].All()
		s = fmt.Sprintf(s, q.WaveID.Deadline.Format("02.01.2006 15:04"), q.WaveID.Deadline.Format("02.01.2006 15:04"))
		helper(w, r, nil, s)
		return
	}

	//
	// Meta parameters
	//
	// Language code changed via URL parameter
	if newCode, ok := sess.ReqParam("lang_code"); ok {
		err := q.SetLangCode(newCode)
		if err != nil {
			log.Printf("Problem setting new lang_code '%v': %v", newCode, err)
		} else {
			sess.PutString("lang_code", q.LangCode)
			log.Printf("new quest lang_code set to '%v' - and saved to session", q.LangCode)
		}
	}
	// Language code not set
	// => try to set questionaire to application default lang code
	if !sess.EffectiveIsSet("lang_code") {
		def := cfg.Get().LangCodes[0]
		err := q.SetLangCode(def)
		if err != nil {
			log.Printf("Problem setting default lang_code '%v': %v", def, err)
		} else {
			log.Printf("quest lang_code set to default '%v'", def)
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
		q.Pages[prevPage].Finished = time.Now().Truncate(time.Second)
	}
	for i1 := 0; i1 < len(q.Pages[prevPage].Groups); i1++ {
		for i2 := range q.Pages[prevPage].Groups[i1].Inputs {
			nm := q.Pages[prevPage].Groups[i1].Inputs[i2].Name
			ok := sess.EffectiveIsSet(nm)
			if ok {
				val := sess.EffectiveStr(nm)
				log.Printf("(Page#%2v) Setting '%v' to '%v'", prevPage, nm, val)
				val = html.EscapeString(val) // XSS prevention
				q.Pages[prevPage].Groups[i1].Inputs[i2].Response = val
			}
		}
	}
	err = q.ValidateReponseData(prevPage, q.LangCode)
	if err != nil {
		q.CurrPage = prevPage // Prevent changing page, keep user on page with errors
	}
	if r.RemoteAddr != "" {
		q.RemoteIP = r.RemoteAddr
	}
	if ok := sess.EffectiveIsSet("finished"); ok {
		if sess.EffectiveStr("finished") == qst.ValSet {
			q.ClosingTime = time.Now().Truncate(time.Second)
		}
	}

	//
	//
	// Save questionaire into session
	err = sess.PutObject("questionaire", q)
	if err != nil {
		helper(w, r, err, "Putting questionaire into session caused error")
		return
	}

	//
	// Save questionaire to file
	pth := q.FilePath1()
	err = os.MkdirAll(filepath.Dir(pth), 0755)
	if err != nil {
		s := fmt.Sprintf("Could not create path %v", filepath.Dir(pth))
		helper(w, r, err, s)
		return
	}

	err = q.Save1(pth)
	if err != nil {
		helper(w, r, err, "Putting questionaire into session caused error")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tplBundle := tpl.Get(w, r, "main.html")
	ts := &tpl.StackT{"quest.html"}

	d := tplDataExtT{
		Q: q,
	}
	d.TplDataT = tpl.TplDataT{
		TplBundle: tplBundle,
		TS:        ts,
		Sess:      &sess,
	}

	err = tplBundle.Execute(w, d)
	if err != nil {
		helper(w, r, err, "Executing template caused error")
		return
	}

}
