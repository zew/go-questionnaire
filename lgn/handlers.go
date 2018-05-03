package lgn

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zew/go-questionaire/sessx"
)

func LogoutH(w http.ResponseWriter, r *http.Request) error {
	sess := sessx.New(w, r)
	err := sess.Remove(w, "login")
	return err
}

// Takes request value "username" and "password".
// Searches for matching user and stores that user
// into the session under key "login".
func LoginCheckH(w http.ResponseWriter, r *http.Request) error {

	sess := sessx.New(w, r)

	err := r.ParseForm()
	if err != nil {
		return err
	}

	if _, ok := r.PostForm["username"]; !ok {
		return nil
	}

	u := r.PostForm.Get("username")
	p := r.PostForm.Get("password") // unencrypted
	log.Printf("trying login1 '%v' '%v'  -  %v", u, p, r.URL)

	l, err := Get().FindAndCheck(u, p)
	if err != nil {
		return err
	}

	log.Printf("logged in as %v", l.User)

	err = sess.PutObject("login", l)
	if err != nil {
		return err
	}

	return nil
}

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

	u := r.PostForm.Get("username")
	o := r.PostForm.Get("oldpassword")
	n := r.PostForm.Get("newpassword")
	n2 := r.PostForm.Get("newpassword2")
	// termsAndConditions := r.PostForm.Get("termsAndConditions")

	if _, ok := r.PostForm["username"]; !ok {
		return "", nil // Nothing to do
	}

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

			} else {
				return "", fmt.Errorf("Neither init nor encrypted password did match for user %v", u)
			}
			return "", loginNotFound
		}
	}
	return "", fmt.Errorf("Old password incorrect (or username not found).")
}
