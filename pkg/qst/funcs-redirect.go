package qst

import (
	"fmt"
	"log"
	"net/http"
)

// this replaces the forwarding based on ErrorForward

// funcs are called before the page is navigated to
type funcRedirectT func(*QuestionnaireT, *pageT, http.ResponseWriter, *http.Request) error

var redirectFuncs = map[string]funcRedirectT{
	"pageFuncExample":          pageForwardExample,
	"pageForwardKnebScreenout": pageForwardKnebScreenout,
	"pageForwardKnebComplete":  pageForwardKnebComplete,
}

// RedirectFuncExec replaces the ErrorForward design
func (p *pageT) RedirectFuncExec(q *QuestionnaireT, w http.ResponseWriter, r *http.Request) error {
	if p.RedirectFunc != "" {
		if fw, ok := redirectFuncs[p.RedirectFunc]; ok {
			log.Printf("RedirectFuncExec %v", p.RedirectFunc)
			return fw(q, p, w, r)
		} else {
			return fmt.Errorf("%v not found in redirectFuncs", p.RedirectFunc)
		}
	}
	return nil
}

func pageForwardExample(q *QuestionnaireT, page *pageT, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func pageForwardKnebScreenout(q *QuestionnaireT, page *pageT, w http.ResponseWriter, r *http.Request) error {
	return pageForwardKneb(q, page, w, r, "screenout")
}
func pageForwardKnebComplete(q *QuestionnaireT, page *pageT, w http.ResponseWriter, r *http.Request) error {
	return pageForwardKneb(q, page, w, r, "complete")
}

func pageForwardKneb(q *QuestionnaireT, page *pageT, w http.ResponseWriter, r *http.Request, paramSet string) error {

	gimStatus := "complete"
	oppStatus := "1"

	if paramSet == "screenout" {
		gimStatus = "screenout"
		oppStatus = "2"
	}

	mailLink := fmt.Sprintf(
		`<a href="mailto:Caroline.Knebel@zew.de?subject=Umfrage Finanzentscheidungen - UID %v&body=Backlink zum Panel nicht möglich." 
		   >Frau Knebel</a>`,
		q.UserIDInt(),
	)

	url := ""
	if panelUID, ok := q.Attrs["i_survey"]; ok {
		url = fmt.Sprintf(
			`https://www.gimpulse.com/?m=6006&return=%v&i_survey=%v`,
			gimStatus,
			panelUID,
		)
	}

	if panUID, ok := q.Attrs["respBack"]; ok {
		url = fmt.Sprintf(
			`https://www.opensurvey.com/survey/1579439651/1704195870?respBack=%v&statusBack=%v`,
			panUID,
			oppStatus,
		)
	}

	if url != "" {
		// http.StatusTemporaryRedirect
		http.Redirect(w, r, url, http.StatusFound)
		log.Printf("Redirected to external %v", url)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(
			w,
			`
				Sie sollten automatisch zum 
				<a _target='blank' href='%v' >Panel-Anbieter</a> 
				zurückgeleitet werden. 
				<br>
				<br>
				Notfalls klicken Sie auf den obigen Link.
				<br>
				<br>
				Wenn auch das nichts hilft, kontaktieren Sie bitte %v.
				<br>
			`,
			url,
			mailLink,
		)

		return nil

	}

	fmt.Fprintf(
		w,
		`
			Keine Panel Benutzer-ID vorhanden. 
			Falls Sie von GIMpulse oder Talk Online / Open Panel gekommen sind, 
			kontaktieren Sie bitte %v.<br>
			<br>
		`,
		mailLink,
	)

	return nil
	// return fmt.Errorf(`no panel ID found`)

}
