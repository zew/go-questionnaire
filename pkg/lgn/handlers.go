package lgn

import (
	"bytes"
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-playground/form"
	"github.com/pbberlin/dbg"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/sessx"
)

// OuterHTMLPost wraps parameter content into a HTML 5 scaffold
// and a form tag
func OuterHTMLPost(htmlTitle, content string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>%v</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
</head>
<body>
	<form method="post" action="{{.SelfURL}}"  style="margin: 50px;"  >
	%v
	</form>
</body>
</html>
`, htmlTitle, content)
}

// LogoutH is a convenience handler to logout via http request
func LogoutH(w http.ResponseWriter, r *http.Request) error {
	sess := sessx.New(w, r)
	return sess.Clear(r.Context())
}

// LoggedInCheck checks, whether as user is logged in,
// and checks whether he has the required roles;
// l is always initialized and never nil
func LoggedInCheck(w io.Writer, r *http.Request, roles ...string) (l *LoginT, loggedIn bool, err error) {

	l, loggedIn, err = FromSession(w, r)
	if err != nil {
		return
	}
	if !loggedIn || l.User == "" {
		loggedIn = false
		return
	}

	for _, role := range roles {
		if !l.HasRole(role) {
			err = fmt.Errorf("Logged in, but role %v is missing", role)

			// returning loggedIn == true, but error != nil
			return
		}
	}
	return

}

// ValidateAndLogin takes request values "username" and "password".
// Searches for matching user in the internal JSON database and stores that user
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

	err = LogoutH(w, r) // remove all previous session info
	if err != nil {
		log.Printf("logout error %v", err)
	}
	log.Printf("logged in as %v %T", l.User, l)

	sess.PutObject("login", l)
	log.Printf("login saved to session as %T by validate (username+password)", &l)

	return nil
}

// ChangePassword takes values from request.Form
// and tries change the user's password.
// The result is updated in the session "login" type.
func ChangePassword(w http.ResponseWriter, r *http.Request) (string, error) {

	l, ok, err := FromSession(w, r)
	if err != nil {
		return "", fmt.Errorf("Could not get login from session; %v", err)
	}
	if !ok {
		return "", fmt.Errorf("You are not logged in")
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
		return "", fmt.Errorf("changing passwd for user %v requires login as user %v", u, u)
	}
	if u == "" {
		return "", fmt.Errorf("no username given")
	}

	if _, ok := r.PostForm["newpassword"]; ok {
		if len(n) < 2 {
			return "", fmt.Errorf("new password too short")
		}
	}

	// Optional confirmation of termsAndConditions; if request key not present - then we do not check
	if _, ok := r.PostForm["termsAndConditions"]; ok {
		log.Printf("tuc: %+v", r.PostForm["termsAndConditions"])
		vals := r.PostForm["termsAndConditions"]
		if len(vals) == 1 && vals[0] != "1" {
			return "", fmt.Errorf("you must agree to the terms and conditions")
		} else if len(vals) == 2 {
			both := vals[0] != "1" && vals[1] != "1"
			if both {
				return "", fmt.Errorf("you must agree to the terms and conditions")
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
					return "", fmt.Errorf("new passwords did not match")
				}

				// Change user database (json file)
				newPassEncr := ComputeMD5Password(l.User, n, Get().Salt)
				logins[idx].PassInitial = ""
				logins[idx].IsInitPassword = false
				logins[idx].PassMd5 = newPassEncr
				err = cloudio.MarshalWriteFile(Get(), LgnsPath)
				if err != nil {
					return "", fmt.Errorf("could not save logins file: %v", err)
				}

				// Change user in session
				// (is this necessary?)
				l.PassInitial = ""
				l.IsInitPassword = false
				l.PassMd5 = newPassEncr
				sess := sessx.New(w, r)
				sess.PutObject("login", *l)

				// SUCCESS
				return "Password changed successfully.", nil

			}
			return "", fmt.Errorf("Neither init nor encrypted password did match for user %v", u)
		}
	}
	return "", fmt.Errorf("old password incorrect (or username not found)")
}

// GenerateHashesH is a admin UI to create login hashes for specific survey and user profile.
// See LoginByHash() for the construction of the check string.
func GenerateHashesH(w http.ResponseWriter, r *http.Request) {

	errMsg := ""

	_, ok := r.PostForm["token"]
	if ok {
		err := ValidateFormToken(r.PostForm.Get("token"))
		if err != nil {
			errMsg += fmt.Sprintf("Invalid request token: %v\n", err)
		}
	} else if !ok && r.Method == "POST" {
		errMsg += fmt.Sprintf("Missing request token\n")
	}

	//
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	src := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>hash logins</title>
	<style>
	* {font-family: monospace;}
	</style>
</head>
<body>

{{if gt (len .ErrMsg) 0 }} <p style='white-space: pre; color:#E22'>{{.ErrMsg}}</p>{{end}}

<form method="post" action="{{.SelfURL}}"  style='white-space:pre' >
    <b>Create hash logins</b>
		<input name="token"       type="hidden"   value="{{.Token}}" />
        Host            <input name="host"        type="text"     value="{{.Host}}"><br>
        
		Survey
            Survey ID   <input name="survey_id"   type="text"     value="{{.SurveyID}}"><br>
            Wave ID     <input name="wave_id"     type="text"     value="{{.WaveID}}" ><br>

        User ID
            Start       <input name="start"       type="text"     value="{{.Start}}"><br>
            Stop        <input name="stop"        type="text"     value="{{.Stop}}" ><br>

        User profile<br>
            Profile     <input name="p"           type="text"     value="{{.Profile}}"><br>

                        <input type="submit" name="submitclassic" id="submit" value="submit" accesskey="s"><br>
        <script> document.getElementById('submit').focus(); </script>
        {{if gt (len .Links  ) 0 }} <p style='                  color:#444'>{{.Links  }}</p>{{end}}
        {{if gt (len .List   ) 0 }} <p style='white-space: pre; color:#444'>{{.List   }}</p>{{end}}
</form>
	
</body>
</html>
`

	type formEntryT struct {
		SelfURL string `json:"self_url"`
		Token   string `json:"token"`

		ErrMsg string `json:"err_msg"`

		Host string `json:"host,omitempty"`

		SurveyID        string `json:"survey_id"`
		WaveID          string `json:"wave_id"`
		Start           int    `json:"start"`
		Stop            int    `json:"stop"`
		Profile         string `json:"p"`
		OverrideClosure string `json:"override"`

		LangCode string   `json:"lang_code"`
		Attrs    []string `json:"attrs"`

		SubmitClassic string `json:"submitclassic"`

		Links template.HTML `json:"links"`
		List  string        `json:"list"`
	}
	fe := formEntryT{}
	fe.Token = FormToken()
	for len(fe.Attrs) < 3+1 {
		fe.Attrs = append(fe.Attrs, "")
	}

	//
	// dec := for--mam.NewDecoder(&for--mam.DecoderOptions{TagName: "json"})
	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err := dec.Decode(&fe, r.Form)
	if err != nil {
		errMsg += fmt.Sprintf("Decoding error: %v\n", err)
	}
	dbg.Dump(fe)

	if fe.SurveyID == "" {
		fmt.Fprint(w, "survey_id must be set as URL param")
		return
	}

	if fe.WaveID == "" {
		t := time.Now()
		if t.Day() > 20 {
			t = t.AddDate(0, 1, 0)
		}
		fe.WaveID = t.Format("2006-01")
	}

	if fe.Start == 0 {
		fe.Start = 10000
		fe.Start = 9990 // 10000 are real FMS users
	}

	if fe.Stop == 0 {
		fe.Stop = 10020
		fe.Stop = 10000
	}

	b1 := &bytes.Buffer{}
	b2 := &bytes.Buffer{}
	for i := fe.Start; i < fe.Stop; i++ {

		queryString := Query(fmt.Sprintf("%v", i), fe.SurveyID, fe.WaveID, fe.Profile)
		if fe.LangCode != "" {
			queryString += "&lang_code=" + fe.LangCode
		}
		for _, attr := range fe.Attrs {
			if attr != "" {
				queryString += "&attrs=" + attr
			}
		}
		if fe.OverrideClosure != "" {
			queryString += "&override_closure=true"
		}

		pfx := cfg.PrefTS()
		if fe.Host == "https://survey2.zew.de" {
			pfx = ""
		}

		url := fmt.Sprintf("%v?%v", fe.Host+pfx, queryString)

		fmt.Fprintf(b1, "<a href='%v'  target='_blank' >login as user %4v<a> \t ", url, i)
		fmt.Fprintf(b2, "%4v\t\t%v\n", i, url)

		fmt.Fprint(b1, " &nbsp; &nbsp; &nbsp; &nbsp; ")

		url2 := fmt.Sprintf("%v?&%v", fe.Host+cfg.PrefTS("reload-from-questionnaire-template"), queryString)

		fmt.Fprintf(b1, "<a href='%v'  target='_blank' >reload from template<a>", url2)
		fmt.Fprint(b1, "<br>")
	}

	fe.Links = template.HTML(b1.String()) // causes gosec error
	fe.List = b2.String()

	fe.ErrMsg = errMsg

	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		fmt.Fprintf(w, "Error parsing inline template: %v", err)
	}

	err = tpl.Execute(w, fe)
	if err != nil {
		fmt.Fprintf(w, "Error executing inline template: %v", err)
	}

}

// LoginPrimitiveH has outer HTML scaffold - for more, see loginPrimitive
func LoginPrimitiveH(w http.ResponseWriter, r *http.Request) {
	loginPrimitive(w, r, true)
}

// LoginPrimitiveCoreH has *no* outer HTML scaffold - for more, see loginPrimitive
func LoginPrimitiveCoreH(w http.ResponseWriter, r *http.Request) {
	loginPrimitive(w, r, false)
}

// loginPrimitive is a primitive handler for http form based login by username and password.
// It serves as pattern for an application specific login.
// It also serves as real handler for applications having only a few admin users.
func loginPrimitive(w http.ResponseWriter, r *http.Request, outerHTML bool) {

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

	if outerHTML {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	src := `
		{{if gt (len .Content) 0 }} <p style='white-space: pre; color:#E22'>{{.Content}}</p>{{end}}
		<h3>Application login</h3>
		<br>
		<input name="token"    type="hidden"   value="{{.Token}}" />

		<label for="username" class="top">Username</label>
		<input id="username" name="username" type="text"     value="{{.L.User}}"><br>

		<label for="password" class="top">Password</label>
		<input id="password" name="password" type="password" /><br>

		<div style="height: 0.6rem"> </div>
		<button name="btnSubmit" accesskey="t" >Submi<u>t</u></button>
	`

	if outerHTML {
		src = OuterHTMLPost("Login", src)
	}

	type dataT struct {
		SelfURL string
		Content string
		Token   string
		L       *LoginT
	}
	data := dataT{
		SelfURL: r.URL.Path,
		Content: msg,
		Token:   FormToken(),
		L:       l,
	}

	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		fmt.Fprintf(w, "Error parsing login inline template: %v", err)
	}

	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error executing login inline template: %v", err)
	}

}

// ChangePasswordPrimitiveH has outer HTML scaffold - for more, see ChangePasswordPrimitive
func ChangePasswordPrimitiveH(w http.ResponseWriter, r *http.Request) {
	changePasswordPrimitive(w, r, true)
}

// ChangePasswordPrimitiveCoreH has *no* outer HTML scaffold - for more, see ChangePasswordPrimitive
func ChangePasswordPrimitiveCoreH(w http.ResponseWriter, r *http.Request) {
	changePasswordPrimitive(w, r, false)
}

// changePasswordPrimitive is a primitive handler for http form based password change.
// It serves as pattern for an application specific password change.
// It also serves as real handler for applications having only a few admin users.
func changePasswordPrimitive(w http.ResponseWriter, r *http.Request, outerHTML bool) {

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

	if outerHTML {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	src := `
		{{if gt (len .Content) 0 }} <p style='white-space: pre; color:#E22'>{{.Content}}</p>{{end}}

		<h3>Change password</h3>
		<br>
		
		<input name="token"       type="hidden"   value="{{.Token}}" />

		<label for="username"     class="top">Username</label>
		<input id="username"      name="username"     type="text"     value="{{.L.User}}"><br>

		<label for="oldpassword"  class="top">Old password</label>
		<input id="oldpassword"   name="oldpassword"  type="password" value="" /><br>

		<label for="newpassword"  class="top">New password</label>
		<input id="newpassword"   name="newpassword"  type="password" value="" /><br>

		<label for="newpassword2" class="top">Repeat</label>
		<input id="newpassword2"  name="newpassword2" type="password" value="" /><br>

		<div style="height: 0.6rem"> </div>
		<button name="btnSubmit" accesskey="t" >Submi<u>t</u></button>
	`
	if outerHTML {
		src = OuterHTMLPost("Change Password", src)
	}

	type dataT struct {
		SelfURL string
		Content string
		Token   string
		L       *LoginT
	}
	data := dataT{
		SelfURL: r.URL.Path,
		Content: msg,
		Token:   FormToken(),
		L:       l,
	}

	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		fmt.Fprintf(w, "Error parsing changepassword inline template: %v", err)
	}

	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error executing changepassword inline template: %v", err)
	}

}
