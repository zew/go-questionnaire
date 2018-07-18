// Package directlogin uses a single condensed value
// to create a valid login.
// Authentication is contained in the value too.
// Thus, no 'user database' required.
// This single value contains a user ID and a checksum on that user ID.
// User ID and checksum must fit together - preventing accidential data entry by another person.
// Param Length   determines the possible number of distinct logins. For example 30 power 3 => 27.000 logins.
// Param CheckSum determines fault sensitivity. For example 30 power 2 => 900 is the chance of accidentally correct login.
// This authentication strategy is relatively weak against brute force attacks.
// It should be applied only, where nothing is to be gained; in academic surveys for example.
// There is the issue of spamming the survey.
// Counter measure 1: Increase the checksum to 3; requiring 3,5 hours for every brute force hit.
// Counter measure 2: Backing off on the IP address - for failed logins only.
// The questionare must allow this authentication.
package directlogin

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/monoculum/formam"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
)

var ultraReadable = []byte("23456789ABCDEFGHKLMNPRSTUVWXYZ") // 30 chars - drop 5 and S ?
var salt = len(ultraReadable) / 2

// Todo: Recreate this map from time to time to free memory
var failBackOff = sync.Map{}

type directLoginT struct {
	Length   int    `json:"length"`    // Digits for ID
	CheckSum int    `json:"check_sum"` // Digits checksum
	Token    string `json:"token"`
	L        string `json:"-"` // The resulting login
}

// New returns a config for direct logins
func New(digitsID, digitsChecksum int) *directLoginT {
	return &directLoginT{Length: digitsID, CheckSum: digitsChecksum}
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

// CrossSum adds the index values of the digits
func (d *directLoginT) CrossSum() int {
	cs := 0
	for _, r := range d.L[:d.Length] {
		j := toDec(r)
		cs += j
	}
	return cs
}

// ToDecimalStr rewrites L6F3G to 17-4-13-1-14
// Separated by hyphen
func (d *directLoginT) ToDecimalStr() string {
	ret := ""
	for _, r := range d.L[:d.Length] {
		j := toDec(r)
		ret = fmt.Sprintf("%v-%v", ret, j)
	}
	ret = strings.Trim(ret, "-")
	return ret
}

// Decimal computes a*30*30 + b*30 + c*1
func (d *directLoginT) Decimal() int {

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

// Generate creates a login ID with checksum
func (d *directLoginT) Generate() {
	d.L = lgn.GeneratePwFromChars(ultraReadable, d.Length)
	div := 1
	for i := 0; i < d.CheckSum+1; i++ {
		div *= len(ultraReadable)
	}
	mod := (d.Decimal() + salt) % div
	d.L += toCode(mod)
}

// Takes a decimal number
// and encodes it as with the characters of ultraReadable
func toCode(i int) string {
	s := ""
	for {
		mod := i % len(ultraReadable)
		s = string(ultraReadable[mod]) + s
		i -= mod
		i /= len(ultraReadable)
		if i < len(ultraReadable) {
			break
		}
	}
	return s
}

func (d *directLoginT) Validate(string) bool {
	div := 1
	for i := 0; i < d.CheckSum+1; i++ {
		div *= len(ultraReadable)
	}
	mod := (d.Decimal() + salt) % div
	checkSum := toCode(mod)

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

	</form>

</body>
</html>
`
	d := New(3, 2)

	// err := r.ParseForm()

	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
	err := dec.Decode(r.Form, d)
	if err != nil {
		errMsg += fmt.Sprintf("Decoding error: %v\n", err)
	}

	if d.Length < 2 {
		errMsg += fmt.Sprintf("Number of chars for Login < 2 - %v\n", d.Length)
	}
	if d.CheckSum < 1 {
		errMsg += fmt.Sprintf("Number of digits for checksum < 1 - %v\n", d.CheckSum)
	}

	d.Generate()

	// log.Printf(util.IndentedDump(d))

	type dataT struct {
		SelfURL string
		Token   string

		ErrMsg string
		Cnt    string
		DL     directLoginT
	}
	data := dataT{
		SelfURL: r.URL.Path,
		Token:   lgn.FormToken(),
		ErrMsg:  errMsg,
		Cnt: fmt.Sprintf("%-24v \n%4v   \n%v \nValid: %v",
			d.L, d.ToDecimalStr(), d.Decimal(), d.Validate(d.L)),
		DL: *d,
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
