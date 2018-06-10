package lgn

import (
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/sessx"
)

// LogoutH is a convenience handler to logout via http request
func LogoutH(w http.ResponseWriter, r *http.Request) error {
	sess := sessx.New(w, r)
	err := sess.Remove(w, "login")
	return err
}

// LoggedInCheck checks, whether as user is logged in,
// and checks whether he has the required roles
func LoggedInCheck(w http.ResponseWriter, r *http.Request, roles ...string) (l *LoginT, loggedIn bool, err error) {

	sess := sessx.New(w, r)

	l = &LoginT{}
	loggedIn, err = sess.EffectiveObj("login", l)
	if err != nil {
		return
	}

	if !loggedIn || l.User == "" {
		return
	}

	for _, role := range roles {
		if !l.HasRole(role) {
			err = fmt.Errorf("Logged in, but role %v is missing", role)
			return
		}
	}
	return

}

// ValidateAndLogin takes request value "username" and "password".
// Searches for matching user and stores that user
// into the session under key "login".
func ValidateAndLogin(w http.ResponseWriter, r *http.Request) error {

	sess := sessx.New(w, r)

	err := r.ParseForm()
	if err != nil {
		return err
	}

	// isSet ?  only POST
	if _, ok := r.PostForm["username"]; !ok {
		return nil // nothing to do
	}

	_, ok := r.PostForm["token"]
	if ok {
		err = ValidateFormToken(r.PostForm.Get("token"))
		if err != nil {
			return fmt.Errorf("Invalid request token: %v", err)
		}
	} else if !ok && r.Method == "POST" {
		return fmt.Errorf("Missing request token")
	}

	u := r.PostForm.Get("username")
	u = html.EscapeString(u)        // XSS prevention
	p := r.PostForm.Get("password") // unencrypted
	log.Printf("trying login1 '%v' '%v'  -  %v", u, p, r.URL)

	l, err := Get().FindAndCheck(u, p)
	if err != nil {
		time.Sleep(200 * time.Millisecond) // Brute force prevention
		return err
	}

	log.Printf("logged in as %v", l.User)

	err = sess.PutObject("login", l)
	if err != nil {
		return err
	}

	return nil
}

// ChangePassword takes values from request.Form
// and tries change the user's password.
// The result is updated in the session "login" type.
func ChangePassword(w http.ResponseWriter, r *http.Request) (string, error) {

	sess := sessx.New(w, r)

	l := LoginT{}
	ok, err := sess.EffectiveObj("login", &l)
	if err != nil {
		return "", fmt.Errorf("Could not get login from session.")
	}
	if !ok {
		return "", fmt.Errorf("You are not logged in.")
	}

	err = r.ParseForm()
	if err != nil {
		return "", fmt.Errorf("Could not parse form: %v", err)
	}

	// isSet ?  only POST
	if _, ok := r.PostForm["username"]; !ok {
		return "", nil // Nothing to do
	}

	{
		_, ok := r.PostForm["token"]
		if ok {
			err = ValidateFormToken(r.PostForm.Get("token"))
			if err != nil {
				return "", fmt.Errorf("Invalid request token: %v", err)
			}
		} else if !ok && r.Method == "POST" {
			return "", fmt.Errorf("Missing request token")
		}
	}

	u := r.PostForm.Get("username")
	o := r.PostForm.Get("oldpassword")
	n := r.PostForm.Get("newpassword")
	n2 := r.PostForm.Get("newpassword2")

	u = html.EscapeString(u) // XSS prevention

	if l.User != u {
		return "", fmt.Errorf("Changing passwd for user %v requires login as user %v.", u, u)
	}
	if u == "" {
		return "", fmt.Errorf("No username given.")
	}

	if _, ok := r.PostForm["newpassword"]; ok {
		if len(n) < 2 {
			return "", fmt.Errorf("New password too short.")
		}
	}

	// Optional confirmation of termsAndConditions; if request key not present - then we do not check
	if _, ok := r.PostForm["termsAndConditions"]; ok {
		log.Printf("tuc: %+v", r.PostForm["termsAndConditions"])
		vals := r.PostForm["termsAndConditions"]
		if len(vals) == 1 && vals[0] != "1" {
			return "", fmt.Errorf("You must agree to the terms and conditions.")
		} else if len(vals) == 2 {
			both := vals[0] != "1" && vals[1] != "1"
			if both {
				return "", fmt.Errorf("You must agree to the terms and conditions.")
			}
		}
	}

	log.Printf("Trying password change '%v' '%v' '%v'  -  %v", u, o, n, r.URL)

	passEncr := ComputeMD5Password(l.User, o, Get().Salt)

	logins := Get().Logins
	for idx, v := range logins {
		if v.User == u {
			log.Printf("Found user %v", u)
			match := false
			if v.IsInitPassword && v.PassInitial == o {
				match = true
				log.Printf("Matching init pw")
			} else if v.PassMd5 == passEncr {
				match = true
				log.Printf("Matching encr pw")
			}
			if match {

				if n != n2 {
					return "", fmt.Errorf("New passwords did not match.")
				}

				// Change user database (json file)
				newPassEncr := ComputeMD5Password(l.User, n, Get().Salt)
				logins[idx].PassInitial = ""
				logins[idx].IsInitPassword = false
				logins[idx].PassMd5 = newPassEncr
				err := Get().Save()
				if err != nil {
					return "", err
				}

				// Change user in session
				// (is this necessary?)
				l.PassInitial = ""
				l.IsInitPassword = false
				l.PassMd5 = newPassEncr
				err = sess.PutObject("login", l)
				if err != nil {
					return "", err
				}

				// SUCCESS
				return "Password changed successfully.", nil

			}
			return "", fmt.Errorf("Neither init nor encrypted password did match for user %v", u)
		}
	}
	return "", fmt.Errorf("Old password incorrect (or username not found).")
}

// LoginByHash takes request value "u" and hash "h";
// and *any number* of additional request parameters.
// It checks the hash against the values except "h";
// ordered by parameter name; plus loginsT.Salt.
// This way we are completely flexible; except for
// inadvertently added parameters.
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

	err := r.ParseForm()
	if err != nil {
		return false, err
	}

	if _, isSet := r.Form["u"]; !isSet { // Form has GET *and* POST values
		return false, nil
	}

	u := r.Form.Get("u")
	u = html.EscapeString(u) // XSS prevention
	h := r.Form.Get("h")     // hash

	l := &LoginT{}
	l.User = u
	l.Roles = map[string]string{}

	keys := []string{}
	for key := range r.Form {
		if key != "h" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	checkStr := ""
	for _, key := range keys {
		val := r.Form.Get(key)
		checkStr += val + "-"
		l.Roles[key] = val
	}
	checkStr += lgns.Salt
	log.Printf("trying hash login over keys %v: '%v' ", keys, checkStr)
	hExpected := Md5Str([]byte(checkStr))
	if hExpected != h {
		return false, fmt.Errorf("hash over check string unequal hash argument\n%v\n%v", hExpected, h)
	}

	// if false {
	// 	_ = sess.EffectiveStr("impossibleKey") // trigger ParsePostForm...
	// 	frm := struct {
	// 		SurveyID string `json:"survey_id,omitempty"`
	// 		WaveID   string `json:"wave_id,omitempty"`
	// 	}{}
	// 	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
	// 	err := dec.Decode(r.Form, &frm)
	// 	util.BubbleUp(err)
	// }

	log.Printf("logging in as %v", u)
	err = sess.PutObject("login", l)
	if err != nil {
		return false, err
	}

	return true, nil
}

// GenerateHashesH is a convenience func to lookup some hashes.
// See LoginByHash() for the construction of the check string.
func GenerateHashesH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	_, loggedIn, err := LoggedInCheck(w, r, "admin")
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	if !loggedIn {
		io.WriteString(w, "login required for this function")
		return
	}

	surveyID := r.URL.Query().Get("survey_id")
	if surveyID == "" {
		io.WriteString(w, "survey_id must be set as URL param")
		return
	}

	waveID := r.URL.Query().Get("wave_id")
	if waveID == "" {
		t := time.Now()
		if t.Day() > 20 {
			t = t.AddDate(0, 1, 0)
		}
		waveID = t.Format("2006-01")
	}

	for i := 99 * 1000; i > 99*1000-10; i-- {
		checkStr := fmt.Sprintf("%v-%v-%v-%v", surveyID, i, waveID, lgns.Salt)
		hsh := Md5Str([]byte(checkStr))
		url := fmt.Sprintf("%v?u=%v&survey_id=%v&wave_id=%v&h=%v", cfg.Pref(), i, surveyID, waveID, hsh)
		a := fmt.Sprintf("<a href='%v'>%v<a><br>", url, url)
		w.Write([]byte(a))
	}
	w.Write([]byte("Finish"))

}

// LoginPrimitiveH is a primitive handler for http form based login.
// It serves as pattern for an application specific login.
// It also serves as real handler for applications having only a few admin users.
func LoginPrimitiveH(w http.ResponseWriter, r *http.Request) {

	msg := ""
	err := ValidateAndLogin(w, r)
	if err != nil {
		msg += fmt.Sprintf("%v\n", err)
	}

	l, isLoggedIn, err := LoggedInCheck(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if isLoggedIn {
		s := fmt.Sprintf("Logged in as %v\n", l.User)
		log.Printf(s)
		msg += s
		// http.Redirect(w, r, cfg.Pref("/"), http.StatusFound)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	src := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Login</title>
</head>
<body>
	<form method="post" action="{{.SelfURL}}"  style="margin: 50px;"  >
		{{if  (len .Cnt) gt 0 }} <p style='white-space: pre; color:#E22'>{{.Cnt}}</p>{{end}}
		Login<br>
					<input name="token"    type="hidden"   value="{{.Token}}" />
		Username:	<input name="username" type="text"     value="{{.L.User}}"><br>
		Password:	<input name="password" type="password" /><br>
		<input type="submit" name="submitclassic" accesskey="l">
	</form>
</body>
</html>
`
	type dataT struct {
		SelfURL string
		Cnt     string
		Token   string
		L       *LoginT
	}
	data := dataT{
		SelfURL: r.URL.Path,
		Cnt:     msg,
		Token:   FormToken(),
		L:       l,
	}

	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error parsing login inline template: %v", err)))
	}

	err = tpl.Execute(w, data)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error executing login inline template: %v", err)))
	}

}

// ChangePasswordPrimitiveH is a primitive handler for http form based password change.
// It serves as pattern for an application specific password change.
// It also serves as real handler for applications having only a few admin users.
func ChangePasswordPrimitiveH(w http.ResponseWriter, r *http.Request) {

	msg, err := ChangePassword(w, r)
	if err != nil {
		msg += fmt.Sprintf("%v\n", err)
	}

	l, loggedIn, err := LoggedInCheck(w, r)
	if err != nil {
		msg += fmt.Sprintf("Logged in check error: %v\n", err)
	}
	if !loggedIn {
		msg += fmt.Sprintf("You are not logged in.\n")
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	src := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Login</title>
</head>
<body>
	<form method="post" action="{{.SelfURL}}"  style="margin: 50px;"  >
		{{if  (len .Cnt) gt 0 }} <p style='white-space: pre; color:#E22'>{{.Cnt}}</p>{{end}}
		Change password<br>
		                <input name="token"        type="hidden"   value="{{.Token}}" />
		Username: 		<input name="username"     type="text"     value="{{.L.User}}"><br>
		Old password: 	<input name="oldpassword"  type="password" value="" /><br>
		New password: 	<input name="newpassword"  type="password" value="" /><br>
		Repeat:			<input name="newpassword2" type="password" value="" /><br>
		<input type="submit" name="submitclassic" accesskey="l">
	</form>
</body>
</html>
`
	type dataT struct {
		SelfURL string
		Cnt     string
		Token   string
		L       *LoginT
	}
	data := dataT{
		SelfURL: r.URL.Path,
		Cnt:     msg,
		Token:   FormToken(),
		L:       l,
	}

	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error parsing changepassword inline template: %v", err)))
	}

	err = tpl.Execute(w, data)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Error executing changepassword inline template: %v", err)))
	}

}
