package lgn

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/go-playground/form"
	hashids "github.com/speps/go-hashids"
	"github.com/zew/go-questionnaire/cfg"
)

type iDElements struct {
	MotherFirstNameFirstLetter string `json:"mother_first,omitempty"      form:"accesskey='m',maxlength='1',size='3',pattern='[A-Z]{1}',placeholder='A'"`
	FatherFirstNameFirstLetter string `json:"father_first,omitempty"      form:"accesskey='f',maxlength='1',size='3',pattern='[A-Z]{1}',placeholder='B'"`
	BirthdayDaySecondDigit     string `json:"birthday_second,omitempty"   form:"accesskey='b',maxlength='1',size='3',pattern='[0-9]{1}',placeholder='1'"`
	FirstNameLastLetter        string `json:"first_name_last,omitempty"   form:"accesskey='l',maxlength='1',size='3',pattern='[a-z]{1}',placeholder='l'"`
}

func (c *iDElements) String() string {
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

// CreateAnonymousID creates a non-personal ID from personal characteristics
func CreateAnonymousID(w http.ResponseWriter, r *http.Request) {

	/*
			Please note that the code is CASE sensitive.
		First letters should be in capitals, letters in the name should be small.

	*/
	msg := `	This is needed if we collect more data on your financial choices and preferences 
	and want to link the questionnaires using pseudonyms over time.

	Names such as László should be used as Laszlo.

	It is important that you construct the code in such a way 
	that you can reconstruct it each time we conduct a survey.
`
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "Error parsing form %v", err)
		return
	}

	frm := iDElements{}

	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err = dec.Decode(&frm, r.Form)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	numericID, hashID := frm.Encode(w)

	url := ""
	if len(frm.String()) > 3 {
		url = fmt.Sprintf(
			"<a href='%v'>Start questionnaire</a>",
			cfg.Pref(fmt.Sprintf("/d/%v--%v", cfg.Get().AnonymousSurveyID, hashID)),
		)
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
	<title>Anonymous access ID</title>
	<meta name="viewport" content="width=device-width, initial-scale=1.0"/>

	<style>
	* {
		font-family: monospace;
		font-size: 11px;
	}
	::placeholder { 
		color:    #ccc;
		opacity:  1;
	}
	</style>
</head>
<body>


<div style='margin-left: 10px;'>

	<div style='color: red;'>%v</div>
	
	<h3>Anonymous deterministic access ID</h3>
	
	<div style='white-space:pre-line;'>%v</div>

</div>
	
	
  <form method="POST" class="survey-edit-form"  style='white-space:pre; font-size: 11px;'>

  What is the first letter 
  of your mother’s first name?   <input type="text"   name="mother_first"         value="%v"   maxlength='1' size='3' pattern='[A-Z]{1}' placeholder='A-Z' /> e.g. <u>A</u>lice  <br>
  What is the first letter 
  of your father’s first name?   <input type="text"   name="father_first"         value="%v"   maxlength='1' size='3' pattern='[A-Z]{1}' placeholder='A-Z' /> e.g. <u>B</u>ob  <br>
  Last digit of your birthday?   <input type="text"   name="birthday_second"      value="%v"   maxlength='1' size='3' pattern='[0-9]{1}' placeholder='0-9'/> e.g. 3<u>0</u>. September or <u>7</u>. October,   <br>
  What is the last letter 
  of your first name?            <input type="text"   name="first_name_last"      value="%v"   maxlength='1' size='3' pattern='[a-z]{1}' placeholder='a-z' /> e.g. Caro<u>l</u>  <br>
                                 <input type="submit" name="btnSubmit" id="btnSubmit"  value="Submit" accesskey="s" style='padding: 8px 28px;' />
<!--								   
    Your personal code is -%v-
    Numeric               -%v-
	Hash                  -%v-
-->	
  %v
</form>        
<script> document.getElementById('btnSubmit').focus(); </script> 


</body>
</html>`,
		err,
		msg,
		// structform.HTML(frm),
		frm.MotherFirstNameFirstLetter,
		frm.FatherFirstNameFirstLetter,
		frm.BirthdayDaySecondDigit,
		frm.FirstNameLastLetter,
		frm.String(),
		numericID,
		hashID,
		url,
	)

}
