// Package lgn implements an internal login database;
// users are stored in a JSON file;
// contains convenience handlers for user retrieval and password change;
// contains login by hashed URL and login by hash ID;
// contains profiles - groups of attributes for users.
// Filename must be given as command line argument
// or environment variable.
// Access to the logins data is threadsafe.
package lgn

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/pbberlin/dbg"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/sessx"
)

var errFoundButWrongPassword = fmt.Errorf("User found but wrong password")
var errLoginNotFound = fmt.Errorf("Login not found")

// Exempted URL params are not hashed for login check.
// They can be freely added to the login-by-hash URL to modify app state.
// Compare package wrap paramPersister.
var exempted = map[string]interface{}{
	// general app control
	"page":      nil,
	"submit":    nil,
	"mobile":    nil,
	"lang_code": nil,
	"v":         nil, // the version - if we want to set it via direct link, compare q.Version() and 'version-from-login-url'
	// the hash itself
	"h": nil,
	// "attrs": nil, // user attributes at login time - must be hashed to prevent tampering

	"skip_validation":        nil, // saved to session only in LoadH(), used in home()
	"override_closure":       nil, // saved to session only in LoginByHash(), used in home() and systemtest
	"redirected_console_log": nil, //
	// we dont use wrap.paramPersister, because its too broad
}

// userAttrs contains URL params which we want to be saved into user attributes.
// The are saved during login into
// - qst.QuestionaireT.Attrs
// - lgn.LoginT.Attrs
// They serve as a property bag session.
// Key is the short form - from the URL. Val is the long form to be saved as login attrs
// LoginT methods Query(), LoginURL() and partly QuestPath() tie into this logic.
var userAttrs = map[string]string{
	"sid": "survey_id",
	"wid": "wave_id",
	"p":   "profile", // user profile, replaces attrs
	"v":   "version", // version set via URL
	// "a":   "attrs",   // general purpose - can occur several times - key:value - or a profile id
}

/*
LoginT must be exported, *not* because we need to pass a type to sessx.GetObject

	loginIntf, ok := sess.EffectiveObj(key)
	if !ok {
		// log.Printf("key %v for LoginT{} is not in session", key)
		return &LoginT{}, false, nil
	}
	l, ok := loginIntf.(LoginT)
	if !ok {
		return &LoginT{}, false, fmt.Errorf("key %v for LoginT{} does not point to lgn.LoginT{} - but to %T", key, loginIntf)
	}

but because we need to declare variables of this type

	type TplDataT struct {
		...
		L      lgn.LoginT
		...
	}
*/
type LoginT struct {
	User     string            `json:"user"`
	Email    string            `json:"email"`
	Group    string            `json:"-"`     // Derived from email domain - or LDAP org
	Provider string            `json:"-"`     // twitter, facebook, ... or hash, anonymous/direct, JSON, LDAP
	Roles    map[string]string `json:"roles"` // i.e. admin: true, can only be set via JSON config; therefore safe
	Attrs    map[string]string `json:"attrs"` // i.e. country: Poland, gender: female, height: 188, a few keys can set via URL params, these are unsafe.

	PassInitial    string `json:"pass_initial"`       // For first login - unencrypted - grants restricted access to change password only
	IsInitPassword bool   `json:"is_init_password"`   // Indicates authentication against PassInitial
	PassMd5        string `json:"pass_md5,omitempty"` // Encrypted password, created from login, permanent password, salt
}

// We need to register all types who are saved into a session
func init() {
	gob.Register(LoginT{})
}

// FromSession loads a login from session;
// second return value contains 'is set'.
func FromSession(w io.Writer, r *http.Request) (*LoginT, bool, error) {

	sess := sessx.New(w, r)
	key := "login"

	loginIntf, ok := sess.EffectiveObj(key)
	if !ok {
		// log.Printf("key %v for LoginT{} is not in session", key)
		return &LoginT{}, false, nil
	}

	l, ok := loginIntf.(LoginT)
	if !ok {
		return &LoginT{}, false, fmt.Errorf("key %v for LoginT{} does not point to lgn.LoginT{} - but to %T", key, loginIntf)
	}

	return &l, true, nil
}

// Query returns a query fragment,
// using the expected param names u, sid, wid, h
// See also userAttrs{}
func Query(userName, surveyID, waveID, profile string, optHash ...string) string {

	// Needs to be in alphabetic order of keys
	// p-sid-u-wid  then salt
	checkStr := fmt.Sprintf("%v-%v-%v-%v-%v", profile, surveyID, userName, waveID, Get().Salt)
	hsh := ""
	if len(optHash) == 0 {
		hsh = Md5Str([]byte(checkStr))
	} else {
		hsh = optHash[0]
	}

	loginURL := fmt.Sprintf("u=%v&sid=%v&wid=%v&p=%v&h=%v", userName, surveyID, waveID, profile, hsh)
	return loginURL

}

// LoginURL returns a URL plus a query fragment,
func LoginURL(userName, surveyID, waveID, profile string, optHash ...string) string {
	loginURL := fmt.Sprintf("%v?%v", cfg.PrefTS(), Query(userName, surveyID, waveID, profile, optHash...))
	return loginURL
}

// HasRole checks the login for a particular role, for instance "admin"
func (l *LoginT) HasRole(role string) bool {
	if l.IsInitPassword {
		return false // no admin functions with an init password
	}
	_, ok := l.Roles[role]
	return ok
}

// ComputeMD5Password is deliberately not a method
func ComputeMD5Password(u, p, salt string) string {
	hashBase := u + p + salt
	return Md5Str([]byte(hashBase))
}

// SetInitPW sets an init password
func (l *LoginT) SetInitPW(salt string) {
	if l.IsInitPassword && l.PassInitial == "" {
		l.PassInitial = GeneratePassword(8)
		log.Printf("\tNew pw is %v", l.PassInitial)
		hashBase := l.User + l.PassInitial + salt
		pfx := cfg.Get().URLPathPrefix
		if pfx == "taxkit" || pfx == "eta" {
			hashBase = l.PassInitial + salt
		}
		l.PassMd5 = Md5Str([]byte(hashBase))
	}
}

type loginsT struct {
	// sync.Mutex
	Salt   string   `json:"salt"`
	Logins []LoginT `json:"logins"`
}

// LgnsPath is obtained by ENV variable or command line flag in main package.
// Being set from the main package.
// Holds the relative path and filename to look for; could be ".lgn/logins.json".
// Relative to the app main dir.
var LgnsPath = path.Join(".", "logins.json")

var lgns *loginsT // package variable 'singleton' - needs to be an allocated struct - to hold pointer receiver-re-assignment

// Get provides access to the logins data.
// It is essential to return a pointer,
// otherwise the unlocking of the returned struct does not work.
func Get() *loginsT {
	// Same as cfg.Get().
	// No lock needed here.
	// Since in load(), we simply exchange one pointer by another at the end of loading.
	// logins.Lock()
	// defer logins.Unlock()
	return lgns
}

// AddTestLogin adds a systemtest login.
// This func is only called by test funcs.
func AddTestLogin() {
	systest := LoginT{
		User:  "systemtest",
		Email: "delete this user in production environment",
		Roles: map[string]string{
			"survey_id": "fmt",
		},
		PassInitial:    "systemtest",
		IsInitPassword: true,
	}
	lgns.Logins = append(lgns.Logins, systest)
}

// Load reads from a JSON file.
// No method to loginsT, no pointer receiver;
// We could only *copy*:  *c = *newCfg
func Load(r io.Reader) {

	decoder := json.NewDecoder(r)
	tmpLogins := loginsT{} // Important, to avoid inconsistent reads from other goroutines
	err := decoder.Decode(&tmpLogins)
	if err != nil {
		log.Fatal(err)
	}

	if len(tmpLogins.Salt) < 5 {
		log.Fatal("Your logins config must contain a salt of at least five characters.")
	}

	log.Printf("Decode from JSON successful. Found %v logins", len(tmpLogins.Logins))

	// Initiallize MD5 hashes from passwords
	// explicitly set.
	for i := 0; i < len(tmpLogins.Logins); i++ {
		tmpLogins.Logins[i].SetInitPW(tmpLogins.Salt)
	}

	// Compute group - i.e. domain
	for i := 0; i < len(tmpLogins.Logins); i++ {
		els := strings.Split(tmpLogins.Logins[i].Email, "@")
		group := els[0]
		if len(els) > 1 {
			group = els[1]
		}
		tmpLogins.Logins[i].Group = group
		tmpLogins.Logins[i].Provider = "JSON"
	}

	log.Printf("\n%s", dbg.Dump2String(tmpLogins))
	lgns = &tmpLogins // replace pointer in one go - should be threadsafe
}

// FindAndCheck takes a username and password
// and scans for matching users in the internal JSON database.
// If optPw is given, a check for matching password is also made
func (l *loginsT) FindAndCheck(u string, optPw ...string) (LoginT, error) {

	u = strings.ToLower(u)
	u = strings.TrimSpace(u)

	checkPassword := false
	passUnencr := ""
	passEncr := ""
	if len(optPw) > 0 {
		checkPassword = true
		passUnencr = optPw[0]
		passEncr = ComputeMD5Password(u, passUnencr, l.Salt)
	}

	for idx := 0; idx < len(l.Logins); idx++ {
		if u == strings.ToLower(l.Logins[idx].User) {
			// log.Printf("found user %v", dbg.Dump2String(l.Logins[idx]))
			if checkPassword {
				pw := l.Logins[idx].PassMd5
				if passEncr == pw {
					return l.Logins[idx], nil
				}
				if l.Logins[idx].IsInitPassword {
					if l.Logins[idx].PassInitial == passUnencr {
						return l.Logins[idx], nil
					}
				}
				return LoginT{}, errFoundButWrongPassword

			}
			return l.Logins[idx], nil
		}
	}
	return LoginT{}, errLoginNotFound
}

// IsFound checks the error argument, if it says
// user found - but wrong password.
func IsFound(err error) bool {
	if err == errFoundButWrongPassword {
		return true
	}
	return false
}

// LoadH is a convenience func to reload logins via http request.
// It reloads logins from json file
// and checks for a specific login
func LoadH(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fileName := LgnsPath
	r, bucketClose, err := cloudio.Open(fileName)
	if err != nil {
		log.Panicf("Error opening writer to %v: %v", fileName, err)
	}
	defer func() {
		err := r.Close()
		if err != nil {
			log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
		}
	}()
	defer func() {
		err := bucketClose()
		if err != nil {
			log.Printf("Error closing bucket of writer to %v: %v", fileName, err)
		}
	}()
	log.Printf("Opened reader to cloud config %v", fileName)
	Load(r)

	fmt.Fprint(w, "Login json file reloaded successfully. \n\n")
	fmt.Fprint(w, "Check for specific user with ?u=[loginname] \n\n")

	err = req.ParseForm()
	if err != nil {
		log.Printf("Error parsing form: %v", err)
		fmt.Fprintf(w, "Error parsing form: %v", err)
	}
	u := req.Form.Get("u")
	l, err := Get().FindAndCheck(u)
	if err != nil {
		str := fmt.Sprintf("%q not found: %v \n", u, err)
		fmt.Fprint(w, str)
		return
	}

	l.PassMd5 = "xxxx"
	l.PassInitial = "xxxx"
	str := fmt.Sprintf("Found %v => %v \n", u, dbg.Dump2String(l))
	fmt.Fprint(w, str)
}

// SaveH is a convenience func to save logins file via http request.
func SaveH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	err := cloudio.MarshalWriteFile(lgns, LgnsPath)
	if err != nil {
		fmt.Fprintf(w, "error writing logins file: %v", err)
		return
	}
	fmt.Fprint(w, "logins saved")
}

// GeneratePasswordH is a convenience func to generate passwords via http request.
// URL parameter len specifies the password length.
func GeneratePasswordH(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parsing form: %v \n\n", err)
	}
	sl := r.Form.Get("len")

	l, _ := strconv.Atoi(sl)
	if l < 3 {
		fmt.Fprint(w, "Specify number of chars with ?len=xx \n\n")
		return
	}
	fmt.Fprintf(w, "len %v => %v \n", l, GeneratePassword(l))
}

// stackoverflow.com/questions/55556
var fullRange = []byte("!#%+23456789:=?@ABCDEFGHJKLMNPRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var stdChars = []byte("!+23456789ABCDEFGHKLMNPRSTUVWXYZabcdefghkmnoprstuvwxyz") // remove ugly or dubious chars

// GeneratePassword creates a password of requested length
func GeneratePassword(length int) string {
	return GeneratePwFromChars(stdChars, length)
}

// GeneratePwFromChars uses chars to create a password of requested length
func GeneratePwFromChars(chars []byte, length int) string {

	numChars := byte(len(chars))      // number of possible chars, i.e. 60, 64
	modulus := 256 % len(chars)       // i.e. 256 % 60 =  16 ;  256 % 128 =   0 ; 256 % 127 = 2
	maxReadIdx := byte(255 - modulus) // i.e. 255 - 16 = 239 ;  255 -   0 = 255 ; 255 -   2 = 253
	// Erroneous edge case: byte(256-0) == 0
	log.Printf("Generate %v-char password - len(chars) %v -  modulus to 255 is %v => maxReadIdx is %v", length, len(chars), modulus, maxReadIdx)

	pw := make([]byte, length)
	rBB := make([]byte, length+(length/4)) // random bytes buffer; there is no reason why it has to by 1.25 * length

	i := 0
	fuse := 0
	for {
		fuse++
		if fuse > 5*1000 {
			panic("something wrong with password generation 1")
		}
		// Fetch random bytes into random bytes buffer
		if _, err := io.ReadFull(rand.Reader, rBB); err != nil {
			panic(err)
		}
		for _, c := range rBB {
			if c > maxReadIdx { // i.e. skip 240...255 since unequal draws from
				continue
			}
			pw[i] = chars[c%numChars] // 88 % 60 => 28
			// log.Printf("\t adding %2vth letter: %3v => %3v", i, c, c%numChars)
			i++
			if i == length {
				return string(pw)
			}
		}
	}
	// panic("something wrong with password generation 2")  // golint says "unreachable"
}

// Md5Str computes the MD5 hash of a byte slice;
// MD5 consists 16 bytes of number.
// We encode these 16 bytes into a Base64 encoded string suitable for URLs.
// Such Base64-for-URL encoded MD5 hash consist of 23 characters.
//
// Now we want to reduce that length to get short login URLs.
// The safe line length for emails is 70 character.
// stackoverflow.com/questions/11794698 suggests maximum size of 70 or 76 characters
//
// We can simply cut off some of the 23 characters,
// since MD5 has good avalanching properties.
//
// For questionnaires without large monetary rewards,
// a hash length of 5 => 64^^5 = 1.073.741.824 combinations is
// sufficient to discourage brute force attacks.
//
// Our URL now has 64 characters:
// https://survey2.zew.de/?u=1000&sid=fmt&wid=2019-06&h=57I7U&p=12
func Md5Str(buf []byte) string {
	hasher := sha256.New()
	_, err := hasher.Write(buf)
	if err != nil {
		log.Printf("lgn.Md5Str() could not write %v: %v", buf, err)
	}
	hshBytes := hasher.Sum(nil)
	// return hex.EncodeToString(hshBytes)  // old

	// Convert 3x 8bit source bytes into 4 bytes
	// Changes must be harmonized with individualbericht_ergebnis.php
	// return base64.RawURLEncoding.EncodeToString(hshBytes)[:5] // direct login
	// return base64.URLEncoding.EncodeToString(hshBytes) //        trailing equal signs
	return base64.RawURLEncoding.EncodeToString(hshBytes) //     no trailing equal signs
}

// Example writes a single login to file, to be extended or adapted
func Example() *loginsT {
	ex := &loginsT{
		Salt: "your salt here",
		Logins: []LoginT{
			{
				User:           "myUser",
				Email:          "myUser@example.com",
				Roles:          map[string]string{"admin": "yes"},
				Attrs:          map[string]string{"country": "Sweden", "height": "174"},
				PassInitial:    "Keep empty - have it set during startup - then call /logins-save",
				IsInitPassword: true,
			},
		},
	}
	return ex
}
