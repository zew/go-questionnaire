package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/xojoc/useragent"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/tpl"

	"github.com/pkg/errors"
)

// An extension with questionnaire
type tplDataExtT struct {
	tpl.TplDataT
	Q *qst.QuestionnaireT // The major app specific object
}

// Loading questionnaire.
// First from session.
// Then from file of previous session.
// Finally from template.
func loadQuestionnaire(w http.ResponseWriter, r *http.Request, l *lgn.LoginT) (*qst.QuestionnaireT, error) {

	sess := sessx.New(w, r)

	// from session
	var q = &qst.QuestionnaireT{}
	ok, err := sess.EffectiveObj("questionnaire", q)
	if err != nil {
		err = errors.Wrap(err, "Reading questionnaire from session caused error")
		return q, err
	}
	if ok {
		return q, nil
	}

	// from file
	log.Printf("Deriving from login: survey_id %v, wave_id %v, user_id %v", l.Attrs["survey_id"], l.Attrs["wave_id"], l.User)
	pthBase := path.Join(qst.BasePath(), l.Attrs["survey_id"]+".json")
	qBase, err := qst.Load1(pthBase)
	if err != nil {
		err = errors.Wrap(err, "Loading base questionnaire from template file caused error")
		return q, err
	}

	pth := l.QuestPath()
	log.Printf("Deriving path: %v", pth)
	qSplit, err := qst.Load1(pth) // previous session
	if err != nil {
		if !cloudio.IsNotExist(err) {
			return q, err
		}
		log.Printf("No previous user questionnaire file %v found. Using base file.", pth)
	} else {
		err = qBase.Join(qSplit)
		if err != nil {
			log.Printf("\tJoining base questionnaire with user data yielded error:    %v", err)
			return q, err
		}
	}

	q = qBase
	err = q.Validate()
	if err != nil {
		err = errors.Wrap(err, "Joined questionnaire validation error")
		return q, err
	}

	if q.Survey.Type != l.Attrs["survey_id"] {
		err = fmt.Errorf("Logged in for survey %v - but template is for %v", l.Attrs["survey_id"], q.Survey.Type)
		return q, err
	}
	if q.Survey.WaveID() != l.Attrs["wave_id"] {
		err = fmt.Errorf("Logged in for wave %v - but template is for %v", l.Attrs["wave_id"], q.Survey.WaveID())
		return q, err
	}

	log.Printf("Questionnaire loaded from file; %v pages", len(q.Pages))
	return q, nil

}

// You can provide
// 1.) an error
// 2.) an error with string to wrap around
// 3.) only a string - which is converted into an error
//
// Bad idea, because code lines of errors are lost.
func helper(w http.ResponseWriter, r *http.Request, err error, msgs ...string) {
	if len(msgs) > 0 {
		if err == nil {
			err = fmt.Errorf(msgs[0])
		} else {
			err = errors.Wrap(err, msgs[0])
		}
	}
	// log.Print(shorter) errorH does logging
	errorH(w, r, err.Error())
}

// LoginByHashID is an entry point for HashIDs
func LoginByHashID(w http.ResponseWriter, r *http.Request) {

	// Assuming https://mydomain.com/some/path/hash-id
	pth := r.URL.Path
	hashID := strings.ToUpper(path.Base(pth)) // last element of path contains hash-id

	err := r.ParseForm()
	if err != nil {
		log.Printf("Parse Form error: %v", err)
		fmt.Fprintf(w, "Parse Form error: %v", err)
		return
	}

	if r.Form.Get("h") != "" {
		log.Printf("There is already a Form value for h %v", r.Form.Get("h"))
		fmt.Fprintf(w, "There is already a Form value for h %v", r.Form.Get("h"))
		return
	}

	r.Form.Set("h", hashID)
	log.Printf("HashID put into h-param %v", r.Form.Get("h"))

	MainH(w, r)

}

// MainH loads and displays the questionnaire with page and lang_code
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
		sess.Remove(w, "questionnaire") // upon successful, possibly new login - remove previous questionnaire from session
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

	q, err := loadQuestionnaire(w, r, l)
	if err != nil {
		helper(w, r, err)
		return
	}
	q.UserID = l.User

	// Already finished?
	if !q.ClosingTime.IsZero() {
		s := cfg.Get().Mp["finished_by_participant"].All(q.ClosingTime.Format("02.01.2006 15:04"))
		helper(w, r, nil, s)
		return
	}

	// Deadline exceeded?
	if time.Now().After(q.Survey.Deadline) {
		s := cfg.Get().Mp["deadline_exceeded"].All(q.Survey.Deadline.Format("02.01.2006 15:04"))
		helper(w, r, nil, s)
		return
	}

	//
	// Meta parameters
	// =============

	//
	// Language code changed via URL parameter
	// => Save into questionnaire and session
	if newCode, ok := sess.ReqParam("lang_code"); ok {
		err := q.SetLangCode(newCode)
		if err != nil {
			log.Printf("Problem setting new lang_code '%v': %v", newCode, err)
		} else {
			sess.PutString("lang_code", q.LangCode)
			log.Printf("new quest lang_code set to '%v' - and saved to session", q.LangCode)
		}
	}

	// Questionnaire language code (still) not set
	// => Try to set questionnaire lang_code from session lang_code (from login)
	if q.LangCode == "" && sess.EffectiveIsSet("lang_code") {
		fromSess := sess.EffectiveStr("lang_code")
		err := q.SetLangCode(fromSess)
		if err != nil {
			log.Printf("Problem setting default lang_code '%v': %v", fromSess, err)
		} else {
			log.Printf("empty quest lang_code set to session value'%v'", fromSess)
		}
	}

	// Questionnaire language code (still) not set
	// => Try to set questionnaire to application default lang code
	if q.LangCode == "" {
		def := cfg.Get().LangCodes[0] // global default
		if len(q.LangCodesOrder) > 0 {
			def = q.LangCodesOrder[0] // questionnaire specific default
		}
		if langCodeUser, ok := l.Attrs["lang_code"]; ok {
			def = langCodeUser
		}
		err = q.SetLangCode(def)
		if err != nil {
			log.Printf("Problem setting default lang_code '%v': %v", def, err)
		}
		log.Printf("Empty quest lang_code set to %v", def)

	}
	// Sync *back* -
	// questionnaire lang_code => app lang_code
	if q.LangCode != "" {
		sess.PutString("lang_code", q.LangCode)
		log.Printf("empty lang_code set to userID lang_code or quest.Default '%v' - and saved to session", q.LangCode)
	}

	// Login attributes => questionaire attributes
	q.Attrs = l.Attrs

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
	if submit == "prev" {
		currPage = q.Prev()
	} else if submit == "next" {
		currPage = q.Next()
	} else {
		// Apart from "prev" and "next", submitBtn can also hold an explicit destination page
		explicit, ok, err := sess.EffectiveInt("submitBtn")
		if err != nil {
			// invalid page value, just dont use it
		}
		if ok && err == nil && explicit > -1 {
			log.Printf("curPage set explicitly by 'submitBtn' to %v", explicit)
			currPage = explicit
		}
	}
	// The progress bar uses "page" to submit an explicit destination page.
	// There are no conflicts of overriding submitBtn and page
	// since submitBtn has only a value if actually pressed.
	explicit, ok, err := sess.EffectiveInt("page")
	if err != nil {
		// invalid page value, just dont use it
	}
	if ok && err == nil && explicit > -1 {
		log.Printf("curPage set explicitly by param 'page' to %v", explicit)
		currPage = explicit
	}
	q.CurrPage = currPage // Put current page into questionnaire
	log.Printf("submitBtn was '%v' - new currPage is %v", submit, currPage)

	//
	// Put request values into questionnaire
	if q.Pages[prevPage].Finished.IsZero() {
		q.Pages[prevPage].Finished = time.Now().Truncate(time.Second)
	}
	for i1 := 0; i1 < len(q.Pages[prevPage].Groups); i1++ {
		for i2 := range q.Pages[prevPage].Groups[i1].Inputs {
			inp := q.Pages[prevPage].Groups[i1].Inputs[i2]
			if inp.IsLayout() {
				continue
			}
			ok := sess.EffectiveIsSet(inp.Name)
			if ok {
				val := sess.EffectiveStr(inp.Name)
				log.Printf("(Page#%2v) Setting %-24q to '%v'", prevPage, inp.Name, val)
				val = html.EscapeString(val) // XSS prevention
				q.Pages[prevPage].Groups[i1].Inputs[i2].Response = val
			}
		}
	}

	if sess.EffectiveStr("skip_validation") == "" {
		err = q.ValidateResponseData(prevPage, q.LangCode)
		if err != nil {
			q.CurrPage = prevPage // Prevent changing page, keep participant on page with errors
		}
	}

	if r.RemoteAddr != "" {
		q.RemoteIP = r.RemoteAddr
	}
	q.UserAgent = r.Header.Get("User-Agent")

	if ok := sess.EffectiveIsSet("finished"); ok {
		if sess.EffectiveStr("finished") == qst.ValSet {
			q.ClosingTime = time.Now().Truncate(time.Second)
		}
	}

	err = q.ComputeDynamicContent(q.CurrPage)
	if err != nil {
		log.Printf("ComputeDynamicContent computation for page %v caused error %v", prevPage, err)
	}

	mobile := computeMobile(w, r, q)

	//
	//
	// Save questionnaire into session
	err = sess.PutObject("questionnaire", q)
	if err != nil {
		helper(w, r, err, "Putting questionnaire into session caused error")
		return
	}

	q2, _ := q.Split()
	err = q2.Save1(l.QuestPath())
	if err != nil {
		helper(w, r, err, "Saving splitted repsonses to file caused error")
		return
	}

	//
	//
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplBundle := tpl.Get(w, r, "main_desktop.html")

	if mobile {
		tplBundle = tpl.Get(w, r, "main_mobile.html")
		q.Pages[q.CurrPage].Width = 100
		q.Pages[q.CurrPage].AestheticCompensation = 0
		for i := 0; i < len(q.Pages[q.CurrPage].Groups); i++ {
			q.Pages[q.CurrPage].Groups[i].Width = 100
		}
	}

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

func computeMobile(w http.ResponseWriter, r *http.Request, q *qst.QuestionnaireT) bool {

	sess := sessx.New(w, r)

	mobile := false

	// Automatic determination from the browser string
	if q.UserAgent != "" {
		ua := useragent.Parse(q.UserAgent)
		if ua == nil {
			log.Printf("useragent.Parse yiedled nil for '%v'", q.UserAgent)
		} else {
			log.Printf("%v on %v - V. %v - mobile or tablet: %v", ua.Name, ua.OS, ua.Version, ua.Mobile || ua.Tablet)
			if ua.Mobile || ua.Tablet {
				mobile = true
			}
		}
	}

	// Override by explicit url parameter
	if mP, ok := sess.ReqParam("mobile"); ok {
		if mP == "0" || mP == "false" {
			mobile = false
			q.Mobile = 0 // no user preference
		}
		if mP == "1" || mP == "true" {
			mobile = true
			q.Mobile = 1 // explicit mobile
		}
		if mP == "2" || mP == "desktop" {
			mobile = false
			q.Mobile = 2 // explicit desktop
		}
	}

	// log.Printf("Mobile = %v", q.Mobile)

	if q.Mobile == 1 {
		return true
	}
	if q.Mobile == 2 {
		return false
	}

	return mobile
}
