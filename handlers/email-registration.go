package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/go-playground/form"
	"github.com/pbberlin/struc2frm"
	"github.com/zew/util"
)

type reportForm struct {
	// stackoverflow.com/questions/399078 - inside character classes escape ^-]\
	Email       string `json:"email"        form:"maxlength='42',size='28',pattern='[a-zA-Z0-9\\.\\-_%+]+@[a-zA-Z0-9\\.\\-]+\\.[a-zA-Z]{2&comma;18}'"`
	Firstname   string `json:"vorname"      form:"maxlength='42',size='28',suffix=''"`
	Lastname    string `json:"nachname"     form:"maxlength='42',size='28',suffix=''"`
	Affiliation string `json:"affiliation"  form:"maxlength='42',size='28',suffix='',placeholder='Ihre Organisation'"`
	Terms       bool   `json:"terms"        form:"suffix='replace_me'"`
}

func (rp *reportForm) CSVLine() string {
	return fmt.Sprintf("%v;%v;%v;%v;\n", rp.Email, rp.Firstname, rp.Lastname, rp.Affiliation)
}

var mtx = sync.Mutex{}

// FMReportFormH shows a registraton form for FMT report
func FMReportFormH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// preprocessing request form
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "cannot parse form: %v<br>\n <pre>%v</pre>", err, util.IndentedDump(r.Form))
		return
	}
	dec := form.NewDecoder()
	dec.SetTagName("json")
	frm := &reportForm{}
	err = dec.Decode(frm, r.Form)
	if err != nil {
		fmt.Fprintf(w, "cannot decode request into form: %v<br>\n <pre>%v</pre>", err, util.IndentedDump(r.Form))
		return
	}

	//
	fmt.Fprintf(w, "<h3>Finanzmarktreport per E-Mail</h3>")
	fmt.Fprintf(w, `<p>Sehr geehrte Damen und Herren,<br><br>
	bitte tragen Sie die E-Mail-Adresse ein, <br>
	unter der Sie künftig den Finanzmarktreport im PDF-Format erhalten möchten.
	
	</p>`)

	if r.PostForm.Get("btnSubmit") != "" {
		if frm.Email == "" {
			fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Email darf nicht leer sein.</p>")
		}
		if !frm.Terms {
			fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Bitte Einverständnis mit Datenschutz ankreuzen.</p>")
		}

		if frm.Email != "" && frm.Terms {
			mtx.Lock()
			defer mtx.Unlock()

			f, err := os.OpenFile("fmr.csv", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>fmr.csv konnte nicht geöffnet werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", err)
				return
			}
			defer f.Close()

			if _, err = f.WriteString(frm.CSVLine()); err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten konnten nicht nach fmr.csv gespeichert werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", err)
				return
			}
			fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten wurden gespeichert</p>")

		}
	}

	w1 := &bytes.Buffer{}
	s2f := struc2frm.New()
	s2f.Indent = 170
	s2f.CSS = strings.ReplaceAll(s2f.CSS, "max-width: 40px;", "max-width: 220px;")

	fmt.Fprint(w1, s2f.HTML(*frm))

	s2 := strings.ReplaceAll(w1.String(), "replace_me", `Ich erkläre mich mit den <a tabindex='-1' href='https://www.zew.de/de/datenschutz' target='_blank' >Datenschutzbestimmungen</a> einverstanden`)

	s2 = strings.ReplaceAll(s2, ">Email", ">E-Mail")
	s2 = strings.ReplaceAll(s2, ">Vorname", ">Vorname (optional)")
	s2 = strings.ReplaceAll(s2, ">Nachname", ">Nachname (optional)")
	s2 = strings.ReplaceAll(s2, ">Affiliation", ">Affiliation (optional)")
	s2 = strings.ReplaceAll(s2, ">Terms", ">Datenschutzbedingungen")
	s2 = strings.ReplaceAll(s2, "<b>S</b>ubmit", "<b>S</b>peichern")

	fmt.Fprint(w, `<style>
		body {
			font-family: -apple-system,BlinkMacSystemFont, Segoe UI, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol;
		}
	
	</style>`)

	fmt.Fprint(w, s2)

}
