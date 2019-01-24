package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/sessx"
)

// ReloadH removes the existing questioniare from the session,
// reading it anew from the questionnaire template JSON file,
// allowing to start anew
func ReloadH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

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

	//

	userSurveyType := ""
	userWaveID := ""
	for attr, val := range l.Attrs {
		if attr == "survey_id" {
			userSurveyType = val
		}
		if attr == "wave_id" {
			userWaveID = val
		}
	}
	pth := filepath.Join(".", qst.BasePath(), userSurveyType, userWaveID, l.User) + ".json"
	err = os.Remove(pth)
	if err != nil {
		log.Printf("Error deleting questionnaire file: %v", err)
		// fmt.Fprintf(w, "Error deleting questionnaire file: %v", err)
	}
	log.Printf("removed quest file %v", pth)

	err = sess.Remove(w, "questionnaire")
	if err != nil {
		helper(w, r, err, "Error deleting questionnaire from session")
		return
	}
	log.Printf("removed quest session")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `
		<form method="POST" class="survey-edit-form" >
			<input type="text"   name="survey_id"           value="%v"   /> <br>
			<input type="text"   name="wave_id"             value="%v"   /> <br>
			<input type="text"   name="u"                   value="%v"   /> <br>
			<input type="text"   name="h"    size=40        value="%v"   /> <br>
			<input type="submit" name="submit" id="submit"  value="Submit" accesskey="s"  /> <br>
		</form>
		<script> document.getElementById('submit').focus(); </script>`,
		userSurveyType,
		userWaveID,
		l.User,
		r.Form.Get("h"),
	)

	fmt.Fprintf(w,
		"<a href='%v?u=%v&survey_id=%v&wave_id=%v&h=%v'  target='_blank'>Start questionnaire (again)<a> <br> ",
		cfg.PrefWTS(), l.User, userSurveyType, userWaveID, r.Form.Get("h"),
	)

}
