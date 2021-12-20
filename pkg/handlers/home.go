package handlers

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/detect"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/sessx"
	"github.com/zew/go-questionnaire/pkg/tpl"

	"github.com/pkg/errors"
)

// Loading questionnaire.
// First from session.
// Then from file of previous session.
// Finally from template.
func loadQuestionnaire(w http.ResponseWriter, r *http.Request, l *lgn.LoginT) (*qst.QuestionnaireT, error) {

	if r.Form.Get("reload") != "" {
		// forcing reload from file -
		// to regain page.width and group.width != 100
		sess := sessx.New(w, r)
		sess.Remove(r.Context(), "questionnaire")
		log.Printf("template will be reloaded from file")
	}

	q, ok, err := qst.FromSession(w, r)
	if err != nil {
		err = errors.Wrap(err, "Reading questionnaire from session caused error")
		return q, err
	}
	if ok {
		return q, nil
	}

	// from file
	log.Printf("Deriving from login: survey_id %v, wave_id %v, variant %v, user_id %v", l.Attrs["survey_id"], l.Attrs["survey_variant"], l.Attrs["wave_id"], l.User)
	fnCore := l.Attrs["survey_id"] + "-" + l.Attrs["wave_id"]
	pthBase := path.Join(qst.BasePath(), fnCore+".json")
	if l.Attrs["survey_variant"] != "" {
		pthBase = path.Join(qst.BasePath(), fnCore+"-"+l.Attrs["survey_variant"]+".json")
	}
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
		// is not exist...
		qBase.UserID = l.User
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

	// since 2021-10 the base file contains the wave id;
	// thus following two checks are much less important
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
		log.Printf("parse form error: %v", err)
		fmt.Fprintf(w, "parse form error: %v", err)
		return
	}

	if r.Form.Get("h") != "" {
		log.Printf("there is already a Form value for h %v", r.Form.Get("h"))
		fmt.Fprintf(w, "there is already a Form value for h %v", r.Form.Get("h"))
		return
	}

	r.Form.Set("h", hashID)
	log.Printf("hashID put into h-param %v", r.Form.Get("h"))

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
		sess.Remove(r.Context(), "questionnaire") // upon successful, possibly new login - remove previous questionnaire from session
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

	// Already finished?
	closed := !q.ClosingTime.IsZero()
	if closed {
		if sess.EffectiveStr("override_closure") == "true" {
			//
		} else {
			s := cfg.Get().Mp["finished_by_participant"].All(q.ClosingTime.Format("02.01.2006 15:04"))
			helper(w, r, nil, s)
			return
		}
	}

	// Deadline exceeded?
	if time.Now().After(q.Survey.Deadline) {
		if sess.EffectiveStr("override_closure") == "true" {
			//
		} else {
			s := cfg.Get().Mp["deadline_exceeded"].All(q.Survey.Deadline.Format("02.01.2006 15:04"))
			helper(w, r, nil, s)
			return
		}
	}

	//
	// Meta parameters
	// =============

	// lang_code of questionnaire - defaults
	if q.LangCode == "" {
		lc := l.Attrs["lang_code"] // from login / login profile
		if lc != "" {
			err := q.SetLangCode(lc)
			if err != nil {
				log.Printf("Problem setting default lang_code from login attr '%v': %v", l.Attrs["lang_code"], err)
			}
			log.Printf("quest lang_code set to default from login/profile  %v", lc)
		}
		if q.LangCode == "" {
			if len(q.LangCodes) > 0 {
				lc = q.LangCodes[0] // questionnaire specific default
			}
			if lc == "" {
				lc = cfg.Get().LangCodes[0] // global default
			}
			err = q.SetLangCode(lc)
			if err != nil {
				log.Printf("Problem setting default lang_code from q.LangCodesOrder / cfg.LangCodes '%v': %v", lc, err)
			}
			log.Printf("quest lang_code set to default from global/quest %v", lc)
		}
	}

	// lang_code of URL GET...
	lcReq, okReq := sess.ReqParam("lang_code")
	if okReq {
		// ... dominates session
		// lcSess := sess.EffectiveStr("lang_code")
		lcSess := sess.GetString(r.Context(), "lang_code") // EffectiveStr is dominated by the URL GET value
		if lcReq != lcSess {
			sess.PutString("lang_code", lcReq)
			log.Printf("REQ lang_code '%v' synced back to session", lcReq)
		}
		// ... dominates questionnaire
		if lcReq != q.LangCode {
			err := q.SetLangCode(lcReq)
			if err != nil {
				log.Printf("error setting quest lang_code '%v' from URL GET or session: %v", lcReq, err)
			} else {
				log.Printf("setting quest lang_code '%v' from URL GET", lcReq)
			}
		}
	} else {
		// questionnaire dominates session
		lcSess := sess.EffectiveStr("lang_code")
		if q.LangCode != lcSess {
			sess.PutString("lang_code", q.LangCode)
			log.Printf("Quest lang_code '%v' synced back to session", q.LangCode)
		}
	}

	// Add login attributes to questionaire attributes
	//    q.Attrs might already contain other values
	if q.Attrs == nil {
		q.Attrs = map[string]string{}
	}
	for k, v := range l.Attrs {
		q.Attrs[k] = v
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
	savedFields := map[string]string{} // prevent repetitions for multiple radios with same name
	for i1 := 0; i1 < len(q.Pages[prevPage].Groups); i1++ {
		for i2 := range q.Pages[prevPage].Groups[i1].Inputs {
			inp := q.Pages[prevPage].Groups[i1].Inputs[i2]
			if inp.IsLayout() {
				continue
			}
			// log.Printf("checking for %v", inp.Name)
			// amazingly, this works for scattered radio inputs as well
			ok := sess.EffectiveIsSet(inp.Name)
			if ok {
				val := sess.EffectiveStr(inp.Name)
				savedFields[inp.Name] = val
				val = html.EscapeString(val) // XSS prevention
				q.Pages[prevPage].Groups[i1].Inputs[i2].Response = val
			}
		}
	}
	for inpName, val := range savedFields {
		log.Printf("(Page#%2v) Setting %-24q to '%v'", prevPage, inpName, val)
	}

	if sess.EffectiveStr("skip_validation") == "" && r.Method == "POST" {
		var forward *qst.ErrorForward
		err, forward = q.ValidateResponseData(prevPage, q.LangCode)
		if err != nil {
			if submit != "prev" { // effectively allow going back - but not going forth
				q.CurrPage = prevPage // Prevent changing page, keep participant on page with errors
			} else {
				q.HasErrors = false
			}
		}

		if forward != nil {
			if strings.HasPrefix(forward.MarkDownPath(), "https://") {
				// previously http.StatusTemporaryRedirect - caused rejection from norstat
				http.Redirect(w, r, forward.MarkDownPath(), http.StatusFound)
				log.Printf("Redirected to external %v", forward.MarkDownPath())
				// for hdrKey, hdrVal := range w.Header() {
				// 	log.Printf("\t%v\t%v", hdrKey, hdrVal)
				// }
			} else {
				core, _ := tpl.SiteCore(q.Survey.Type)
				relURL := path.Join("/doc/", core, q.LangCode, forward.MarkDownPath())
				relURL = cfg.Pref(relURL)
				http.Redirect(w, r, relURL, http.StatusTemporaryRedirect)
				log.Printf("Redirected to markdown %v", relURL)
				// tpl.RenderStaticContent(w, forward.MarkDownPath(), core, q.LangCode)
			}
			return
		}
	}

	if r.RemoteAddr != "" {
		q.RemoteIP = r.RemoteAddr
	}
	q.UserAgent = r.Header.Get("User-Agent")

	if ok := sess.EffectiveIsSet("finished"); ok {
		if sess.EffectiveStr("finished") == qst.Finished {
			q.ClosingTime = time.Now().Truncate(time.Second)
		}
	}

	q.EnumeratePages()

	err = q.ComputeDynamicContent(q.CurrPage)
	if err != nil {
		log.Printf("ComputeDynamicContent computation for page %v caused error %v", prevPage, err)
	}

	//
	//
	// Save questionnaire into session
	sess.PutObject("questionnaire", q)

	q2, _ := q.Split()
	err = q2.Save1(l.QuestPath())
	if err != nil {
		helper(w, r, err, "Saving splitted responses to file caused error")
		return
	}

	//
	//
	htmlTitle := fmt.Sprintf(
		"%v %v",
		cfg.Get().MpSite[q.Survey.Type]["app_org"].TrSilent(q.LangCode),
		cfg.Get().MpSite[q.Survey.Type]["app_label"].TrSilent(q.LangCode),
	)

	//
	mp := map[string]interface{}{
		"LangCode":  q.LangCode, // default would be cfg.Get().LangCodes[0]
		"Site":      q.Survey.Type,
		"HTMLTitle": htmlTitle,
		"LogoTitle": q.Survey.TemplateLogoText(q.LangCode),
		"Q":         q,
		"Content":   "",
	}

	// mobile := computeMobile(w, r, q)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	w1 := &strings.Builder{}
	tpl.Exec(w1, r, mp, "quest.html")

	mp["Content"] = w1.String()
	// tpl.RenderStack(r, w, []string{"layout.html"}, mp)

	tpl.Exec(w, r, mp, "layout.html")

}

func computeMobile(w http.ResponseWriter, r *http.Request, q *qst.QuestionnaireT) bool {

	sess := sessx.New(w, r)

	// determination from the browser string
	mobile := false
	if detect.IsMobile(r) {
		mobile = true
	}

	qMobile := 0 // q.Mobile int  -  `json:"mobile,omitempty"` // 0 - no preference, 1 - desktop, 2 - mobile

	// override by explicit url parameter
	if mP, ok := sess.ReqParam("mobile"); ok {
		if mP == "0" || mP == "false" {
			mobile = false
			qMobile = 0 // no user preference
		}
		if mP == "1" || mP == "true" {
			mobile = true
			qMobile = 1 // explicit mobile
		}
		if mP == "2" || mP == "desktop" {
			mobile = false
			qMobile = 2 // explicit desktop
		}
	}

	// log.Printf("Mobile = %v", qMobile)

	if qMobile == 1 {
		return true
	}
	if qMobile == 2 {
		return false
	}

	return mobile
}
