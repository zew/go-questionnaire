package lgn

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/form"
	hashids "github.com/speps/go-hashids"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/ctr"
)

type iDElements struct {
	MotherFirstNameFirstLetter string `json:"mother_first,omitempty"      form:"accesskey='m',maxlength='1',size='3',pattern='[A-Z]{1}',placeholder='A'"`
	FatherFirstNameFirstLetter string `json:"father_first,omitempty"      form:"accesskey='f',maxlength='1',size='3',pattern='[A-Z]{1}',placeholder='B'"`
	BirthdayDaySecondDigit     string `json:"birthday_second,omitempty"   form:"accesskey='b',maxlength='1',size='3',pattern='[0-9]{1}',placeholder='1'"`
	FirstNameLastLetter        string `json:"first_name_last,omitempty"   form:"accesskey='l',maxlength='1',size='3',pattern='[a-z]{1}',placeholder='l'"`
}

func (c *iDElements) String() string {
	if c.MotherFirstNameFirstLetter == "" ||
		c.FatherFirstNameFirstLetter == "" ||
		c.BirthdayDaySecondDigit == "" ||
		c.FirstNameLastLetter == "" {
		return ""
	}

	return c.MotherFirstNameFirstLetter + c.FatherFirstNameFirstLetter + c.BirthdayDaySecondDigit + c.FirstNameLastLetter
}

func (c *iDElements) Encode(w io.Writer) (int, string) {

	s := c.String()

	if len(s) < 4 {
		return 0, ""
	}

	concat := ""

	for _, char := range s {
		asciiVal := int(rune(char)) - 47
		concat = fmt.Sprintf("%02v%v", asciiVal, concat)
		log.Printf("Char %s  Rune %#3v  - asciiVal %3v - prod %16v", string(char), char, asciiVal, concat)
	}

	h, err := hashids.NewWithData(getGen())
	if err != nil {
		fmt.Fprintf(w, "Error creating hash ID: %v", err.Error())
	}

	numericID, err := strconv.Atoi(concat)
	if err != nil {
		fmt.Fprintf(w, "Error converting to int: %v", err.Error())
	}

	hashID, err := h.Encode([]int{numericID})
	if err != nil {
		fmt.Fprintf(w, "Error encoding hash ID: %v", err.Error())
	}

	return numericID, hashID

}

var hashIDCounter = ctr.New() //  not stable across app restarts

var doOnce sync.Once

// LoginWithoutID creates a hash ID
// and forwards to direct login /d - LoginByHashID
//
// the user ID is created from unix time
// 		plus some in-memory counter
//
// it's called permalink in subsequent logic
//
// it's related to CreateAnonymousID but the ID comes from an internal timestamp plus atomic counter,
//   it is not created from coarse personal attributes (such as first letter of father's name)
func LoginWithoutID(w http.ResponseWriter, r *http.Request) {

	// prevent *repeated* anonymous login
	l, isLoggedIn, err := LoggedInCheck(w, r)
	if err != nil {
		log.Printf("LoggedInCheck error inside LoginWithoutID: %v", err)
	}
	if isLoggedIn {
		if _, ok := l.Attrs["permalink"]; ok {
			log.Printf("LoginWithoutID - a login is already present and has the permalink attribute")
			// forward to existing questionnaire
			url := cfg.Pref("/")
			http.Redirect(w, r, url, http.StatusTemporaryRedirect)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			fmt.Fprint(w, "You are already logged in.")
			return
		}
	}

	doOnce.Do(func() {
		time.Sleep(1000 * time.Millisecond) // prevent identical ID conflicts during restart within *same* second; see below
	})

	// resolution is seconds
	// capped to seconds of 4 months instead of 100 years
	// for intra-second distinction: incremental counter
	// minimum size: 100.000 - to distinguish from pre-generated UIDs
	// max size 10 billion; is still acceptable
	//
	// 	doubles could occur during app restarts within the same second
	// 		in conjunction with multiple requests within the same second
	// 		=> doOnce() above prevents that
	ts := int(time.Now().Unix())
	last4Months := 3600 * 24 * 4 * 28 //                   9.676.800
	uid := ts % last4Months           //           0...    9.676.800
	uid += 100 * 1000                 //     100.000...    9.776.800
	uid *= 1000                       // 100.000.000...9.776.800.000

	incr := hashIDCounter.Increment()
	if incr > 999 {
		hashIDCounter.Reset()
		incr = 1
	}
	uid += int(incr)

	h, err := hashids.NewWithData(getGen())
	if err != nil {
		fmt.Fprintf(w, "Error creating hash ID: %v", err.Error())
	}

	hashID, err := h.Encode([]int{uid})
	if err != nil {
		fmt.Fprintf(w, "Error encoding hash ID: %v", err.Error())
	}

	// forward to LoginByHashID
	url := cfg.Pref(fmt.Sprintf("/d/%v--%v", cfg.Get().AnonymousSurveyID, hashID))

	if true {
		// http.Redirect(w, r, url, http.StatusFound)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}

	if true {
		link := ""
		// forward-link%vforward-link is used by decorators, to extract the forwarding URL
		link = fmt.Sprintf(`
			<a href='%v'>Start questionnaire</a>
			<!--forward-link%vforward-link-->
			`,
			url,
			url,
		)
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, link)
	}

}

// CreateAnonymousIDH has outer HTML scaffold - for more, see CreateAnonymousID
func CreateAnonymousIDH(w http.ResponseWriter, r *http.Request) {
	createAnonymousID(w, r, true)
}

// CreateAnonymousIDCoreH has *no* outer HTML scaffold - for more, see CreateAnonymousID
//    seems unused
func CreateAnonymousIDCoreH(w http.ResponseWriter, r *http.Request) {
	createAnonymousID(w, r, false)
}

// CreateAnonymousID creates a HashID for LoginByHashID
// from personal characteristics
// then forwards to LoginByHashID
func createAnonymousID(w http.ResponseWriter, r *http.Request, outerHTML bool) {

	/*
		Please note that the code is CASE sensitive.
		First letters should be in capitals, letters in the name should be small.

		This is needed if we collect more data on your financial choices and preferences
		and want to link the questionnaires using pseudonyms over time.
	*/
	msg := `We dont save <i>any</i> personal data. 
	<div style="height: 0.25em" ></div>
	Just enter some characters, so that you can relogin later.
	<div style="height: 0.25em" ></div>
	Note: Characters such as László should be used as Laszlo.
`
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parsing form %v", err)
		return
	}

	_, ok := r.PostForm["token"]
	if ok {
		err = ValidateFormToken(r.PostForm.Get("token"))
		if err != nil {
			fmt.Fprintf(w, "Invalid request token: %v", err)
		}
	}

	frm := iDElements{}

	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err = dec.Decode(&frm, r.Form)

	numericID, hashID := frm.Encode(w)

	link := ""
	if len(frm.String()) > 3 {
		url := cfg.Pref(fmt.Sprintf("/d/%v--%v", cfg.Get().AnonymousSurveyID, hashID))
		// forward-link%vforward-link is used by decorators, to extract the forwarding URL
		link = fmt.Sprintf(`
			<!--forward-link%vforward-link-->
			<a href='%v'>Start questionnaire</a>
			`,
			url,
			url,
		)
	}

	if outerHTML {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	}

	src := fmt.Sprintf(`


<h3>Anonymous login</h3>

<p style='awhite-space:pre-line;'>%v</p>

<div style='color: red;'>%v</div>

<div>

	<input name="token"    type="hidden"   value="%v" />


	<label for="mother_first" class="top">
	What is the first letter <br> 
	of your mother’s first name?  
	</label>
	<input  type="text"  id="mother_first"   name="mother_first"        value="%v"   maxlength='1' size='3' pattern='[A-Z]{1}' placeholder='A-Z' /> 
	<div class="postlabel" >e.g. <u>A</u>lice</div>

	<br>
	
	<label for="father_first" class="top">
	What is the first letter <br>
	of your father’s first name?  
	</label>
	<input  type="text"  id="father_first"   name="father_first"        value="%v"   maxlength='1' size='3' pattern='[A-Z]{1}' placeholder='A-Z' />
	<div class="postlabel" >e.g. <u>B</u>ob</div>

	<br>

	<label for="birthday_second" class="top">
	Last digit of your birthday?  
	</label>
	<input  type="text"  id="birthday_second"   name="birthday_second"   value="%v"   maxlength='2' size='3' pattern='[0-9]{1,2}' placeholder='0-9'
		inputmode="numeric"
	/>
	<div class="postlabel" >e.g. <u>30</u>. September or <u>7</u>. October</div>

	<br>

	<label for="first_name_last" class="top">
	What is the last letter 
	of your first name?           
	</label>
	<input  type="text"  id="first_name_last"   name="first_name_last"   value="%v"   maxlength='1' size='3' pattern='[a-z]{1}' placeholder='a-z' 
		autocapitalize=off
	/>
	<div class="postlabel" >e.g. Caro<u>l</u></div>

	<br>


	<button name="btnSubmit" accesskey="t" >Submi<u>t</u></button>

	<!--								   
		Your personal code is -%v-
		Numeric               -%v-
		Hash                  -%v-
	-->	

	%v

</div>        

<!--
<script> document.getElementById('btnSubmit').focus(); </script> 
-->
`,
		msg,
		err,
		// structform.HTML(frm),
		FormToken(),
		frm.MotherFirstNameFirstLetter,
		frm.FatherFirstNameFirstLetter,
		frm.BirthdayDaySecondDigit,
		frm.FirstNameLastLetter,
		frm.String(),
		numericID,
		hashID,
		link,
	)

	if outerHTML {
		src = OuterHTMLPost("Anonymous deterministic access ID", src)
		src = strings.ReplaceAll(src, "action=\"{{.SelfURL}}\"", "")
	}

	fmt.Fprint(w, src)

}
