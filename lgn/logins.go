// Package lgn implements an internal login database;
// users are stored in a JSON file;
// contains convenience handlers for user retrieval and password change.
// Filename must be given as command line argument
// or environment variable.
// Access to the logins data is made in threadsafe manner.
package lgn

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"golang.org/x/crypto/md4"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/util"
)

var foundButWrongPassword = fmt.Errorf("User found but wrong password")
var loginNotFound = fmt.Errorf("Login not found")

// Type LoginT must be exported, *not* because we need to pass a type to sessx.GetObject
// 		l := lgn.LoginT{}
// 		ok, err := sess.EffectiveObj("login", &l)
// but because we need to declare variables of this type
// 		type TplDataT struct {
// 			...
// 			L      lgn.LoginT
// 			...
// 		}

type LoginT struct {
	User  string            `json:"user"`
	Email string            `json:"email"`
	Group string            `json:"empty"` // Derived from emai domain - or LDAP org
	Roles map[string]string `json:"roles"` // i.e. admin: true , gender: female, height: 188

	PassInitial    string `json:"pass_initial"`       // For first login - unencrypted - grants restricted access to change password only
	IsInitPassword bool   `json:"is_init_password"`   // Indicates authentication against PassInitial
	PassMd5        string `json:"pass_md5,omitempty"` // Encrypted password, created from login, permanent password, salt
}

func (l *LoginT) HasRole(role string) bool {
	_, ok := l.Roles[role]
	return ok
}

// Deliberately not a method
func ComputeMD5Password(u, p, salt string) string {
	hashBase := u + p + salt
	pfx := cfg.Get().UrlPathPrefix
	if pfx == "taxkit" || pfx == "eta" {
		hashBase = p + salt
	}
	return Md5Str([]byte(hashBase))
}

// Set an init password
func (l *LoginT) SetInitPW(salt string) {
	if l.IsInitPassword && l.PassInitial == "" {
		l.PassInitial = GeneratePassword(8)
		log.Printf("\tNew pw is %v", l.PassInitial)
		hashBase := l.User + l.PassInitial + salt
		pfx := cfg.Get().UrlPathPrefix
		if pfx == "taxkit" || pfx == "eta" {
			hashBase = l.PassInitial + salt
		}
		l.PassMd5 = Md5Str([]byte(hashBase))
	}
}

type loginsT struct {
	sync.Mutex
	Salt   string   `json:"salt"`
	Logins []LoginT `json:"logins"`
}

// Obtained by ENV variable or command line flag in main package.
// Being set from the main package.
// Holds the relative path and filename to look for; could be ".lgn/logins.json".
// Relative to the app main dir.
var LgnsPath = path.Join(".", "logins.json")

var lgns *loginsT // package variable 'singleton' - needs to be an allocated struct - to hold pointer receiver-re-assignment

// It is essential to return a poiner,
// otherwise the unlocking of the returned struct does not work.
func Get() *loginsT {
	// Same as cfg.Get().
	// No lock needed here.
	// Since in load(), we simply exchange one pointer by another at the end of loading.
	// logins.Lock()
	// defer logins.Unlock()
	return lgns
}

// No method to loginsT, no pointer receiver;
// We could only *copy*:  *c = *newCfg
func Load() {
	// l.Lock()
	// defer l.Unlock()

	file, err := util.LoadConfigFile(LgnsPath)
	if err != nil {
		log.Fatalf("Could not load logins file: %v", err)
	}
	log.Printf("Found logins file: %v", LgnsPath)
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Error closing logins file: %v", err)
		}
		log.Printf("Closed logins file: %v", LgnsPath)
	}()

	decoder := json.NewDecoder(file)
	tmpLogins := loginsT{} // Important, to avoid inconsistent reads from other goroutines
	err = decoder.Decode(&tmpLogins)
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
		if len(els) > 0 {
			group = els[1]
		}
		tmpLogins.Logins[i].Group = group
	}

	log.Printf("\n%#s", util.IndentedDump(tmpLogins))
	lgns = &tmpLogins // replace pointer in one go - should be threadsafe
}

func (l *loginsT) Save(fn ...string) error {
	l.Lock()
	defer l.Unlock()

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(l, firstColLeftMostPrefix, "\t")
	if err != nil {
		return err
	}

	saveDir := path.Dir(LgnsPath)
	err = os.Chmod(saveDir, 0755)
	if err != nil {
		return err
	}

	loginsFile := path.Base(LgnsPath)
	if len(fn) > 0 {
		loginsFile = fn[0]
	}

	pthOld := path.Join(saveDir, loginsFile)
	fileBackup := strings.Replace(loginsFile, ".json", fmt.Sprintf("_%v.json", time.Now().Unix()), 1)
	pthNew := path.Join(saveDir, fileBackup)

	if loginsFile != "logins-example.json" {
		err = os.Rename(pthOld, pthNew)
		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(pthOld, byts, 0644)
	if err != nil {
		return err
	}

	log.Printf("Saved logins file to %v", pthNew)
	return nil
}

// Method FindAndCheck takes a username and password
// and scans for matching users.
// If optPw is given, a check for matching password is also made
func (l loginsT) FindAndCheck(u string, optPw ...string) (LoginT, error) {

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
			// log.Printf("found user %v", util.IndentedDump(l.Logins[idx]))
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
				return LoginT{}, foundButWrongPassword

			} else {
				return l.Logins[idx], nil
			}
		}
	}
	return LoginT{}, loginNotFound
}

func IsFound(err error) bool {
	if err == foundButWrongPassword {
		return true
	}
	return false
}

// Reload logins json file
// and check for a specific login
func LoadH(w http.ResponseWriter, r *http.Request) {

	_, loggedIn, err := LoggedInCheck(w, r, "admin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !loggedIn {
		http.Error(w, "admin login required for this function", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	Load()

	w.Write([]byte("Login json file reloaded successfully. \n\n"))
	w.Write([]byte("Check for specific user with ?u=[loginname] \n\n"))

	r.ParseForm()
	u := r.Form.Get("u")
	l, err := Get().FindAndCheck(u)
	if err != nil {
		str := fmt.Sprintf("%q not found: %v \n", u, err)
		w.Write([]byte(str))
		return
	}

	l.PassMd5 = "xxxx"
	l.PassInitial = "xxxx"
	str := fmt.Sprintf("Found %v => %v \n", u, util.IndentedDump(l))
	w.Write([]byte(str))

}

func SaveH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	_, loggedIn, err := LoggedInCheck(w, r, "admin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !loggedIn {
		http.Error(w, "login required for this function", http.StatusInternalServerError)
		return
	}

	err = lgns.Save()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("logins saved"))
}

func GeneratePasswordH(w http.ResponseWriter, r *http.Request) {

	_, loggedIn, err := LoggedInCheck(w, r, "admin")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !loggedIn {
		http.Error(w, "admin login required for this function", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Write([]byte("Specify number of chars with ?len=xx \n\n"))

	r.ParseForm()
	sl := r.Form.Get("len")

	if sl != "" {
		l, _ := strconv.Atoi(sl)
		if l > 2 {
			pw := GeneratePassword(l)
			str := fmt.Sprintf("len %v => %v \n", l, pw)
			w.Write([]byte(str))
		}
	}
}

// stackoverflow.com/questions/55556
var stdCharsXX = []byte("!#%+23456789:=?@ABCDEFGHJKLMNPRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var stdChars = []byte("!+23456789ABCDEFGHJKLMNPRSTUVWXYZabcdefghijkmnopqrstuvwxyz") // remove some ugly special chars

func GeneratePassword(length int) string {

	numChars := byte(len(stdChars))   // number of possible chars, i.e. 60, 64
	modulus := 256 % len(stdChars)    // i.e. 256 % 60 =  16 ;  256 % 128 =   0 ; 256 % 127 = 2
	maxReadIdx := byte(255 - modulus) // i.e. 255 - 16 = 239 ;  255 -   0 = 255 ; 255 -   2 = 253
	// Erroneous edge case: byte(256-0) == 0
	log.Printf("Generate password - len(stdChars) %v -  modulus to 255 is %v => maxReadIdx is %v", len(stdChars), modulus, maxReadIdx)

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
			pw[i] = stdChars[c%numChars] // 88 % 60 => 28
			// log.Printf("\t adding %2vth letter: %3v => %3v", i, c, c%numChars)
			i++
			if i == length {
				return string(pw)
			}
		}
	}
	panic("something wrong with password generation 2")
}

// Function Md5Str computes the md5 hash of a byte slice.
func Md5Str(buf []byte) string {
	pfx := cfg.Get().UrlPathPrefix
	if pfx == "taxkit" || pfx == "eta" {
		// It is md4 (FOUR) by mistake -
		// but since the application was already deplyoed, we cannot correct
		hasher := md4.New()
		hasher.Write(buf)
		hash := hasher.Sum(nil)
		return hex.EncodeToString(hash)
	} else {
		hasher := md5.New()
		hasher.Write(buf)
		hash := hasher.Sum(nil)
		return hex.EncodeToString(hash)
	}
}

func Example() {
	ex := &loginsT{
		Salt: "your salt here",
		Logins: []LoginT{
			LoginT{
				User:           "myUser",
				Email:          "myUser@example.com",
				Roles:          map[string]string{"admin": "yes"},
				PassInitial:    "Keep emtpy - have it set during startup - then call /logins-save",
				IsInitPassword: true,
			},
		},
	}
	ex.Save("logins-example.json")
}
