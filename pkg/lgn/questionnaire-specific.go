package lgn

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/sessx"
)

// LoginByHash first checks for direct login;
// extra short and preconfigured in config.json;
// last part of the path - moved into h;
//
// LoginByHash then takes request values "u" and hash "h" - and not any password;
// it checks the hash against the values except "h";
// any other request parameters are sorted and included into the hashing check;
// extended by loginsT.Salt.
// reduced  by those in 'exempted';
//
// request values "h" without an "u" are considered direct login attempts,
// where the hash actually represents a user ID.
//
// On success, function creates a logged in user
// out of nothing by the name of u.
// This user gets assigned the params/values as
// role-names/values.
//
// On wrong hashes, it returns the difference as error.
// Be careful not to show the error to the end user.
func LoginByHash(w http.ResponseWriter, r *http.Request) (bool, error) {

	sess := sessx.New(w, r)

	// extra treatments for FMS participants who escape the & in the URL query string
	if strings.Contains(r.URL.RawQuery, "&amp;") {
		r.URL.RawQuery = strings.ReplaceAll(r.URL.RawQuery, "&amp;", "&") // potentially activating ampersands within parameters and values
		log.Printf("Query string cleansed from &amp; %v", r.URL.RawQuery)
		r.Form = nil // as it has already been parsed using the broken query string
	}

	err := r.ParseForm()
	if err != nil {
		return false, err
	}

	u := r.Form.Get("u")
	u = html.EscapeString(u) // XSS prevention
	h := r.Form.Get("h")     // hash

	if u == "" && h == "" {
		return false, nil
	}

	// First - try case direct login
	if _, isSet := r.Form["u"]; !isSet { // Note: r.Form[key] contains GET *and* POST values
		if _, isSet := r.Form["h"]; isSet {
			// => userId is not set - but hash is set

			//
			surveyID := ""
			permaLink := ""
			parts := strings.Split(h, "--") // h coming from anonymous id
			if len(parts) > 1 {
				surveyID = strings.ToLower(parts[0])
				permaLink = h
				h = parts[1]
			}

			strUserID := fmt.Sprint(HashIDDecodeFirst(h))
			userID, err := strconv.Atoi(strUserID)
			if err != nil {
				// "This can merely indicate, that the hash was not intended encoding a userID
				log.Printf("Hash yielded userID - no hash ID could be extracted: %v from %v", err, h)
			}

			if userID > 0 {
				log.Printf("Trying anonymous login - surveyID | hashID | userID - %v | %v | %v", surveyID, h, userID)
				for _, dlr := range cfg.Get().DirectLoginRanges {
					// log.Printf("  Checking dlr %v - %4v <=  %4v <=  %4v", dlr.SurveyID, dlr.Start, userID, dlr.Stop)

					// matching the ranges
					// ---------------------
					// either non-empty survey ID matches - regardless of userID
					// or         empty survey ID - but userID in range
					if (surveyID != "" && surveyID == dlr.SurveyID) ||
						userID >= dlr.Start && userID <= dlr.Stop {

						log.Printf(
							"Matching survey %v - or direct login range %v <=  %v <=  %v",
							dlr.SurveyID, dlr.Start, userID, dlr.Stop,
						)
						l := LoginT{}
						l.User = strUserID
						l.Provider = "anonymous/direct"
						l.IsInitPassword = false // roles become effective only for non init passwords
						l.Roles = map[string]string{}
						l.Attrs = map[string]string{}
						l.Attrs["survey_id"] = dlr.SurveyID
						if permaLink != "" {
							l.Attrs["permalink"] = permaLink
						}

						l.Attrs["wave_id"] = dlr.WaveID
						for pk, pv := range dlr.Profile {
							l.Attrs[pk] = pv
						}

						for _, param := range preservedIntoAttrs {
							val := r.FormValue(param)
							if val != "" {
								l.Attrs[param] = val
							}
						}
						// todo - attrs from profiles? see below

						if sess.EffectiveStr("override_closure") == "true" {
							sess.PutString("override_closure", "true")
						}

						sess.PutObject("login", l)
						log.Printf("login saved to session as %T from loginByHash", l)
						return true, nil

					}
				}
			}
		}
		return false, nil
	}

	// Second - try login user by param u and hash h

	l := LoginT{}
	l.User = u
	l.IsInitPassword = false // roles become effective only for non init passwords
	l.Roles = map[string]string{}
	l.Attrs = map[string]string{}

	chkKeys := []string{}
	for key := range r.Form {
		if _, ok := exempted[key]; ok {
			continue
		}
		chkKeys = append(chkKeys, key)
	}

	sort.Strings(chkKeys)
	checkStr := ""
	for _, key := range chkKeys {
		val := r.Form.Get(key)
		checkStr += val + "-"
	}
	checkStr += lgns.Salt
	log.Printf("trying hash login over chkKeys %v-salt: '%v' ", strings.Join(chkKeys, "-"), checkStr)
	hExpected := Md5Str([]byte(checkStr))
	if hExpected != h {
		return false, fmt.Errorf("hash over check string unequal hash argument\n%v\n%v", hExpected, h)
	}
	l.Provider = "hash"

	//
	// set attributes from configured profiles
	// &sid=fmt&p=1 => config.json -> fmt1
	// &sid=fmt&p=2 => config.json -> fmt2
	for key := range r.Form {
		// same key - multiple values
		// attrs=country:Sweden&attrs=height:176
		if val, ok := userAttrs[key]; ok {
			if key == "attrs" {
				// obsolete
			} else if key == "p" {
				profileKey := r.Form.Get("sid") + r.Form.Get(key)
				prof, ok := cfg.Get().Profiles[profileKey]
				if !ok {
					log.Printf("Did not find profile %v", profileKey)
					continue
				}
				for pk, pv := range prof {
					log.Printf("\tprofile to attr key-val  %-20v - %v", pk, pv)
					l.Attrs[pk] = pv
				}
			} else {
				l.Attrs[val] = r.Form.Get(key)
			}
		}
	}

	if sess.EffectiveStr("override_closure") == "true" {
		sess.PutString("override_closure", "true")
	}

	log.Printf("logging in %v - attrs %v ; %T", u, l.Attrs, l)
	sess.PutObject("login", l)
	log.Printf("login saved to session as %T from loginByHash", l)

	return true, nil
}

// ReloadH removes the existing questioniare from the session,
// reading it anew from the questionnaire template JSON file,
// allowing to start anew
func ReloadH(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	log.Printf("reset start")

	_, err := LoginByHash(w, r)
	if err != nil {
		log.Printf("Login by hash error 1: %v", err)
		// Don't show the revealing original error
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoginByHash error."
		log.Print(s)
		w.Write([]byte(s))
		return
	}
	l, isLoggedIn, err := LoggedInCheck(w, r)
	if err != nil {
		log.Printf("Login by hash error 2: %v", err)
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "LoggedInCheck error."
		log.Print(s)
		w.Write([]byte(s))
		return
	}
	if !isLoggedIn {
		log.Printf("Login by hash error 3: %v", "not logged in")
		s := cfg.Get().Mp["login_by_hash_failed"].All()
		s += "You are not logged in."
		log.Print(s)
		w.Write([]byte(s))
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	msg := ""
	if r.Method == "POST" {
		l.DeleteFiles()
		sess.Remove(r.Context(), "questionnaire")
		log.Printf("removed quest session")
		msg = "User files deleted. Questionnaire deleted from session."
	} else {
		msg = "Not a POST request. No delete action taken."
	}

	if sess.EffectiveStr("skip_validation") != "" {
		sess.PutString("skip_validation", "true")
	} else {
		sess.Remove(r.Context(), "skip_validation")
	}

	if sess.EffectiveStr("override_closure") != "" {
		sess.PutString("override_closure", "true")
	} else {
		sess.Remove(r.Context(), "override_closure")
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
	  <head>
		<meta charset="utf-8" />
		<title>Reset entries</title>
		<style>
		* {font-family: monospace;}
		</style>
	  </head>
      <body style='white-space:pre'>
        <b>%v</b>
        <form method="POST" class="survey-edit-form" >
            User ID          <input type="text"   name="u"                   value="%v"   /> <br>
            Survey ID        <input type="text"   name="sid"                 value="%v"   /> <br>
            Wave ID          <input type="text"   name="wid"                 value="%v"   /> <br>
            User profile #   <input type="text"   name="p"                   value="%v"   /> country name, currency etc.<br>
            Hash             <input type="text"   name="h"    size=40        value="%v"   /> <br>
            Lang code        <input type="text"   name="lang_code"  size=6   value="%v"   /> 'en', 'de' ...<br>
            Page             <input type="number" name="page"                value="%v"  min=0 max=88 /> zero-indexed <br>
            Mobile           <input type="text"   name="mobile"              value="%v"   /> 0-auto, 1-mobile, 2-desktop <br>
            Skip validation  <input type="text"   name="skip_validation"     value="%v"   /> <br>
            Override closure <input type="text"   name="override_closure"     value="%v"   /> <br>
            %v
                             <input type="submit" name="submit" id="submit"  value="Submit" accesskey="s"  /> <br>
		</form>        
		<script> document.getElementById('submit').focus(); </script>  `,
		msg,
		l.User,
		l.Attrs["survey_id"],
		// l.Attrs["variant"], is not part of the checksum - and is derived from userID - thus no need to have the parameter
		l.Attrs["wave_id"],
		relForm.Get("p"),
		relForm.Get("h"),
		relForm.Get("lang_code"),
		relForm.Get("page"),
		relForm.Get("mobile"),
		relForm.Get("skip_validation"),
		relForm.Get("override_closure"),
		attrsStr,
	)

	queryString := Query(
		relForm.Get("u"), relForm.Get("sid"), relForm.Get("wid"), relForm.Get("p"), relForm.Get("h"),
	)
	if relForm.Get("lang_code") != "" {
		queryString += "&lang_code=" + relForm.Get("lang_code")
	}
	if relForm.Get("page") != "" {
		queryString += "&page=" + relForm.Get("page")
	}
	if relForm.Get("mobile") != "" {
		queryString += "&mobile=" + relForm.Get("mobile")
	}
	for _, attr := range relForm["attrs"] {
		if attr != "" {
			queryString += "&attrs=" + attr
		}
	}

	url := fmt.Sprintf("%v?%v", cfg.PrefTS(), queryString)

	fmt.Fprintf(w, "<a href='%v'  target='_blank'>Start questionnaire (again)<a> <br>\n", url)
	if r.Method == "POST" {
		fmt.Fprintf(w,
			`
		<SCRIPT language="JavaScript1.2">
			//var win = window.open('%s','qst','menubar=1,resizable=1,width=350,height=250,target=q');
			var win = window.open('%s', 'qst');
			win.focus();
			console.log('window opened')
		</SCRIPT>`,
			url,
			url,
		)
	}

	fmt.Fprint(w, "\t</body>\n</html>")

}

// basePath gives the 'root' for loading and saving questionnaire JSON files.
// Duplicate of qst.basePath - cyclic dependencies
func basePath() string {
	return path.Join(".", "responses")
}

// QuestPath returns the path to the JSON *user* questionnaire,
// Similar to qst.QuestionnaireT.FilePath1()
// See also userAttrs{}.
// Param suffix is appended to
func (l *LoginT) QuestPath(suffixFN ...string) string {

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

	if userSurveyType == "" || userWaveID == "" {
		log.Printf("Error constructing path for user questionnaire file; userSurveyType or userWaveID is empty: %v - %v", userSurveyType, userWaveID)
	}

	fn := l.User
	if len(suffixFN) > 0 {
		fn = fmt.Sprintf("%v_%v", fn, suffixFN[0])
	}

	pth := path.Join(".", basePath(), userSurveyType, userWaveID, fn) + ".json"
	return pth
}

// DeleteFiles deletes all JSON files
func (l *LoginT) DeleteFiles() {

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

	if userSurveyType == "" || userWaveID == "" {
		log.Printf("Error deleting questionnaire file;  userSurveyType or userWaveID is empty: %v - %v", userSurveyType, userWaveID)
		return
	}

	// pth1 := path.Join(".", BasePath(), userSurveyType, userWaveID, l.User) + "_joined.json"
	// pth2 := path.Join(".", BasePath(), userSurveyType, userWaveID, l.User) + "_split.json"
	pth3 := path.Join(".", basePath(), userSurveyType, userWaveID, l.User) + ".json"
	// pth4 := path.Join(".", BasePath(), userSurveyType, userWaveID, l.User) + ".json.attrs"

	pths := []string{pth3}

	for _, pth := range pths {
		err := cloudio.Delete(pth)
		if err != nil {
			if !cloudio.IsNotExist(err) {
				log.Printf("Error deleting questionnaire file: %v", err)
			}
		} else {
			log.Printf("removed quest file %v", pth)
		}
	}

}
