// Package directlogin uses a single condensed value
// to create a valid login.
// Authentication is contained in the value too.
// Thus, no 'user database' required.
// This single value contains a user ID and a checksum on that user ID.
// User ID and checksum must fit together - preventing accidental data entry by another person.
// Param Length   determines the possible number of distinct logins. For example 30 power 3 => 27.000 logins.
// Param CheckSum determines fault sensitivity. For example 30 power 2 => 900 is the chance of accidentally correct login.
// This authentication strategy is relatively weak against brute force attacks.
// It should be applied only, where nothing is to be gained; in academic surveys for example.
// There is the issue of spamming the survey.
// Counter measure 1: Increase the checksum to 3; requiring 3,5 hours for every brute force hit.
// Counter measure 2: Backing off on failed logins - slowing down the brute forcer.
// The questionare must allow this authentication.
package directlogin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/sessx"
)

var ultraReadable = []byte("23456789ABCDEFGHKLMNPRSTUVWXYZ") // 30 chars - drop 5 and S ?
var salt = len(ultraReadable) / 2

// Todo: Recreate this map from time to time to free memory
var failBackOff = sync.Map{}

// DirectLoginT contains the ID part and the checksum part of a login
type DirectLoginT struct {
	Length   int    `json:"length"`    // Digits for ID
	CheckSum int    `json:"check_sum"` // Digits checksum
	L        string `json:"-"`         // The resulting login
}

type formEntryT struct {
	DirectLoginT
	Token string `json:"token"`
	Start int    `json:"start"`
	Stop  int    `json:"stop"`
}

// New returns a config for direct logins
func New(digitsID, digitsChecksum int) *DirectLoginT {
	return &DirectLoginT{Length: digitsID, CheckSum: digitsChecksum}
}

// Translates a single digit to decimal
// 2 => 0
// A => 8 etc
func toDec(r rune) int {
	ri := byte(r)
	for idx, rn := range ultraReadable {
		if ri == rn {
			return idx
		}
	}
	return 0
}

// Takes a decimal number
// and encodes it as with the characters of ultraReadable
// Inverse func of Decimal()
func toCode(i int) string {
	s := ""
	for {
		mod := i % len(ultraReadable)
		s = string(ultraReadable[mod]) + s
		i -= mod
		i /= len(ultraReadable)
		if i < len(ultraReadable) {
			if i != 0 {
				s = string(ultraReadable[i]) + s
			}
			break
		}
	}
	return s
}

func (d DirectLoginT) computeCheckSum() string {
	div := 1
	for i := 0; i < d.CheckSum; i++ {
		div *= len(ultraReadable)
	}
	mod := (d.Decimal() + salt) % div
	coded := toCode(mod)
	for len(coded) < d.CheckSum {
		coded = string(ultraReadable[0]) + coded
	}
	return coded
}

// DecimalHyphenated rewrites L6F3G to 17-4-13-1-14.
// Separated by hyphen.
// For debugging.
func (d *DirectLoginT) DecimalHyphenated() string {
	ret := ""
	for _, r := range d.L[:d.Length] {
		j := toDec(r)
		ret = fmt.Sprintf("%v-%v", ret, j)
	}
	ret = strings.Trim(ret, "-")
	return ret
}

// CrossSum adds the index values of the digits
// Unused
func (d *DirectLoginT) CrossSum() int {
	cs := 0
	for _, r := range d.L[:d.Length] {
		j := toDec(r)
		cs += j
	}
	return cs
}

// Decimal computes a*30*30 + b*30 + c*1
// Inverse func of toCode
// i.e. 22G => 16
func (d *DirectLoginT) Decimal() int {

	if d.Length > 12 {
		log.Printf("More than 12 digits: Decimal() will overflow. Truncating")
		return 0
	}

	dp := 0
	reversed := []rune{}
	for _, r := range d.L[:d.Length] {
		reversed = append([]rune{r}, reversed...)
	}

	base := 1
	for _, r := range reversed {
		j := toDec(r)
		dp += j * base
		// log.Printf("Adding %2v * %6v  => %6v", j, base, dp)
		base *= len(ultraReadable)
	}

	return dp
}

// GenerateRandom creates a random login ID with checksum
func (d *DirectLoginT) GenerateRandom() {
	d.L = lgn.GeneratePwFromChars(ultraReadable, d.Length)
	d.L += d.computeCheckSum()
}

// GenerateFromDec creates a login ID with checksum from the seed argument
func (d *DirectLoginT) GenerateFromDec(seed int) {
	d.L = toCode(seed)
	if len(d.L) > d.Length {
		overFlow := d.Length - len(d.L)
		d.L = d.L[overFlow:]
	}
	for len(d.L) < d.Length {
		d.L = string(ultraReadable[0]) + d.L // pad with leading 'zeroes'
	}
	d.L += d.computeCheckSum()
}

// Validate the login ID against the checksum
func (d *DirectLoginT) Validate() bool {
	checkSum := d.computeCheckSum()
	if checkSum == d.L[d.Length:] {
		return true
	}

	// Globally backing off on failed login attempts
	key := time.Now().Truncate(20 * time.Second)
	keyPrevious := key.Add(-20 * time.Second) // two slots - spanning 20 to 40 seconds
	ctr := 0
	{
		ifac, ok := failBackOff.Load(key)
		if ok {
			i, _ := ifac.(int)
			ctr += i
		}
	}
	{
		ifac, ok := failBackOff.Load(keyPrevious)
		if ok {
			i, _ := ifac.(int)
			ctr += i
		}
	}
	ctr++
	failBackOff.Store(key, ctr)

	bo := ctr * ctr // 1*1=1 ; 2*2=4 ; 5*5=25
	if bo > 15 {
		log.Printf("SPAM Suspicion: %v failed login attempts in 20-40secs", ctr)
		bo = 15
	}
	time.Sleep(time.Duration(bo) * time.Second) // Hesitate against brute force

	return false
}

// GenerateH is a convenience func to generate direct logins via http request.
// URL parameter len specifies the password length.
// URL parameter cs  specifies the checksum length.
func GenerateH(w http.ResponseWriter, r *http.Request) {

	if cfg.Get().IsProduction {
		_, loggedIn, err := lgn.LoggedInCheck(w, r, "admin")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !loggedIn {
			http.Error(w, "admin login required for this function", http.StatusInternalServerError)
			return
		}
	}

	errMsg := ""

	_, ok := r.PostForm["token"]
	if ok {
		err := lgn.ValidateFormToken(r.PostForm.Get("token"))
		if err != nil {
			errMsg += fmt.Sprintf("Invalid request token: %v\n", err)
		}
	} else if !ok && r.Method == "POST" {
		errMsg += fmt.Sprintf("Missing request token\n")
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	src := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>Direct Login</title>
</head>
<body>
	<form method="post" action="{{.SelfURL}}"  style="margin: 50px;"  >
		{{if  (len .ErrMsg) gt 0 }} <p style='white-space: pre; color:#E22'>{{.ErrMsg}}</p>{{end}}
		Create direct login<br>
		                <input name="token"       type="hidden"   value="{{.Token}}" />
		Digits Login: 	<input name="length"      type="text"     value="{{.DL.Length}}"><br>
		Digits Checksum:<input name="check_sum"   type="text"     value="{{.DL.CheckSum}}"><br>
		<input type="submit"   name="submitclassic" accesskey="s"><br>

		{{if  (len .Cnt   ) gt 0 }} <p style='white-space: pre; color:#222'>{{.Cnt   }}</p>{{end}}
		
		Start: 	<input name="start"      type="text"     value="{{.DL.Start}}"><br>
		Stop: 	<input name="stop"       type="text"     value="{{.DL.Stop}}" ><br>
		{{if  (len .Links  ) gt 0 }} <p style='                  color:#444'>{{.Links  }}</p>{{end}}
		{{if  (len .List   ) gt 0 }} <p style='white-space: pre; color:#444'>{{.List   }}</p>{{end}}

	</form>

</body>
</html>
`
	d := New(3, 2)
	fe := formEntryT{}
	fe.DirectLoginT = *d
	// err := r.ParseForm()

	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
	err := dec.Decode(r.Form, &fe)
	if err != nil {
		errMsg += fmt.Sprintf("Decoding error: %v\n", err)
	}
	d = &fe.DirectLoginT

	if d.Length < 2 {
		errMsg += fmt.Sprintf("Number of chars for Login < 2 - %v\n", d.Length)
	}
	if d.CheckSum < 1 {
		errMsg += fmt.Sprintf("Number of digits for checksum < 1 - %v\n", d.CheckSum)
	}

	d.GenerateRandom()

	di := New(d.Length, d.CheckSum)

	links := ""
	for i := fe.Start; i <= fe.Stop; i++ {
		di.GenerateFromDec(i)
		str := fmt.Sprintf(
			"<a href='%v%v' target='_blank' >%v</a><br>\n",
			cfg.PrefWTS("/direct"), di.L, di.Decimal(),
		)
		links += str
	}

	list := ""
	for i := fe.Start; i <= fe.Stop; i++ {
		di.GenerateFromDec(i)
		str := fmt.Sprintf(
			"%05v\t%v\t%v\t%v\tValid: %v\n",
			i, di.L, di.DecimalHyphenated(), di.Decimal(), di.Validate(),
		)
		list += str
	}

	// log.Printf(util.IndentedDump(d))

	type dataT struct {
		SelfURL string
		Token   string

		ErrMsg string
		Cnt    string
		DL     formEntryT
		Links  template.HTML
		List   string
	}
	data := dataT{
		SelfURL: r.URL.Path,
		Token:   lgn.FormToken(),
		ErrMsg:  errMsg,
		Cnt: fmt.Sprintf("%-24v \n%4v   \n%v \nValid: %v",
			d.L, d.DecimalHyphenated(), d.Decimal(), d.Validate()),
		DL:    fe,
		Links: template.HTML(links),
		List:  list,
	}

	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		fmt.Fprintf(w, "Error parsing inline template: %v", err)
	}

	err = tpl.Execute(w, data)
	if err != nil {
		fmt.Fprintf(w, "Error executing inline template: %v", err)
	}

}

// CheckFailed shows failed direct login attempts
func CheckFailed(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	f := func(key interface{}, value interface{}) bool {
		fmt.Fprintf(w, "%v  - Failed attempts %3v\n", key, value)
		return true
	}
	failBackOff.Range(f) // iterating the concurrency safe map

}

// ValidateAndLogin takes the *last* directory of the URL path as direct login
// The direct login parameters and the survey_id and wave_id are still hard coded.
// Login success forwards to the main handler
func ValidateAndLogin(w http.ResponseWriter, r *http.Request) {

	sess := sessx.New(w, r)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	p := r.URL.Path
	p = path.Base(p) // last element of path contains direct login

	dl := New(3, 2)
	dl.L = p
	good := dl.Validate()

	if !good {
		fmt.Fprintf(w, "Not a valid direct login %v\n", p)
		return
	}

	l := lgn.LoginT{}
	l.User = fmt.Sprintf("%v", dl.Decimal())
	l.Roles = map[string]string{}
	l.Roles["survey_id"] = "peu2018"
	l.Roles["wave_id"] = "2018-08"
	log.Printf("directly logged in as %v - ID %v", l.User, dl.Decimal())
	// fmt.Fprintf(w, "directly logged in as %v - ID %v\n", l.User, dl.Decimal())  // prevents redirect

	lgn.LogoutH(w, r) // remove all previous session info
	err := sess.PutObject("login", l)
	if err != nil {
		http.Error(w, "Error saving login to session", http.StatusInternalServerError)
		return
	}

	red := cfg.Pref("")
	if len(r.URL.RawQuery) > 0 {
		red += "?" + r.URL.RawQuery
	}
	http.Redirect(w, r, red, http.StatusSeeOther)

}
