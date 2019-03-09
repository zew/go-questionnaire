package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
)

// ReloadH removes the existing questioniare from the session,
// reading it anew from the questionnaire template JSON file,
// allowing to start anew
func ReloadH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	log.Printf("reset start")

	_, err := lgn.LoginByHash(w, r)
	if err != nil {
		log.Printf("Login by hash error 1: %v", err)
		// Don't show the revealing original error
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoginByHash error."
		helper(w, r, nil, s)
		return
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	msg := ""
	if r.Method == "POST" {
		l.DeleteFiles()
		sess.Remove("questionnaire")
		log.Printf("removed quest session")
		msg = "User files deleted. Questionnaire deleted from session."
	} else {
		msg = "Not a POST request. No delete action taken."
	}

	relForm := r.Form // relevant Form
	if len(r.PostForm) > 5 {
		relForm = r.PostForm
	}

	attrsStr := ""
	for _, val := range relForm["attrs"] {
		if val != "" {
			attrsStr += fmt.Sprintf("<input type=\"text\" name=\"attrs\" value=\"%v\" /> <br>\n", val)
		}
	}

	fmt.Fprintf(w, `
	<html>
		<head></head>
		<body>
			<b>%v<b>
			<form method="POST" class="survey-edit-form" >
					<input type="text"   name="sid"                 value="%v"   /> <br>
					<input type="text"   name="wid"                 value="%v"   /> <br>
					<input type="text"   name="u"                   value="%v"   /> <br>
					<input type="text"   name="h"    size=40        value="%v"   /> <br>
		lang code	<input type="text"   name="lang_code"  size=6   value="%v"   /> <br>
					%v
					<input type="submit" name="submit" id="submit"  value="Submit" accesskey="s"  /> <br>
			</form>
			<script> document.getElementById('submit').focus(); </script>
				
		</body>
	</html>

		`,
		msg,
		l.Attrs["survey_id"],
		l.Attrs["wave_id"],
		l.User,
		relForm.Get("h"),
		relForm.Get("lang_code"),
		attrsStr,
	)

	queryString := lgn.LoginURL(
		relForm.Get("u"), relForm.Get("sid"), relForm.Get("wid"), relForm.Get("h"),
	)
	if relForm.Get("lang_code") != "" {
		queryString += "&lang_code=" + relForm.Get("lang_code")
	}
	for _, attr := range relForm["attrs"] {
		if attr != "" {
			queryString += "&attrs=" + attr
		}
	}

	url := fmt.Sprintf("%v?%v", cfg.PrefWTS(), queryString)

	fmt.Fprintf(w, "<a href='%v'  target='_blank'>Start questionnaire (again)<a> <br> ", url)

}
