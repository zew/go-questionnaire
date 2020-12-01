package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pbberlin/struc2frm"
	"github.com/zew/util"
)

type formRegistrationFMT struct {
	Vorname     string `json:"first_name"     form:"maxlength='40',size='25',label='Vorname',xxnobreak='true'"`
	Nachname    string `json:"last_name"      form:"maxlength='40',size='30',label='Name'"`
	Unternehmen string `json:"unternehmen"    form:"maxlength='40',size='40',placeholder='Ihr Unternehmen oder Organisation'"`
	Abteilung   string `json:"abteilung"      form:"maxlength='40',size='40',"`
	Position    string `json:"position"       form:"maxlength='40',size='40',suffix='Bezeichnung Ihrer aktuellen Position'"`

	Separator1 string `json:"separator1"      form:"subtype='separator',label=''"`

	PLZ     string `json:"plz"                form:"maxlength='6',size='6',label='PLZ',xxnobreak='true'"`
	Ort     string `json:"ort"                form:"maxlength='40',size='40'"`
	Strasse string `json:"strasse"            form:"maxlength='40',size='40',suffix='mit Hausnummer'"`
	// stackoverflow.com/questions/399078 - inside character classes escape ^-]\
	Email   string `json:"email"              form:"maxlength='40',size='40',pattern='[a-zA-Z0-9\\.\\-_%+]+@[a-zA-Z0-9\\.\\-]+\\.[a-zA-Z]{0&comma;2}'"`
	Telefon string `json:"telefon"            form:"maxlength='40',size='40',label='Telefon'"`

	Separator2 string `json:"separator2"      form:"subtype='separator',label='replace_me_2'"`

	Geschlecht  string `json:"geschlecht"     form:"subtype='select'"`
	Geburtsjahr string `json:"geburtsjahr"    form:"maxlength='5',size='5'"`
	Abschluss   string `json:"abschluss"      form:"maxlength='40',size='40',label='Höchster Abschluss',suffix='z.B. Diplom'"`
	Studienfach string `json:"studienfach"    form:"maxlength='40',size='40',label='Ggf. Studienfach',suffix='z.B. VWL'"`
	Hochschule  string `json:"hochschule"     form:"maxlength='40',size='40',label='Ggf. Hochschule',suffix='z.B. Uni Mannheim'"`
	Einstieg    string `json:"einstieg"       form:"maxlength='5',size='5',label='Einstieg ins Berufsleben',suffix='(Jahr)'"`
	Leitung     string `json:"leitung"        form:"subtype='select',size='1',label='Leitungsbefugnis über',suffix='Mitarbeiter'"`

	Terms bool `json:"terms"        form:"label='Datenschutz',suffix='replace_me_1'"`
}

// yearValid - either empty or within 1930 and 2050
func yearValid(s string) bool {
	if s == "" {
		return true // no number is ok
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return false // not a number => not ok
	}
	if i < 1930 {
		return false
	}
	if i > 2050 {
		return false
	}
	return true
}

func (frm formRegistrationFMT) Validate() (map[string]string, bool) {
	errs := map[string]string{}
	g1 := frm.Vorname != ""
	if frm.Vorname == "" {
		errs["first_name"] = "Bitte geben Sie Ihren Vornamen an."
	}
	g2 := frm.Nachname != ""
	if frm.Nachname == "" {
		errs["last_name"] = "Bitte geben Sie Ihren Nachnamen an."
	}
	g3 := frm.Unternehmen != ""
	if frm.Unternehmen == "" {
		errs["unternehmen"] = "Bitte geben Sie Ihr Unternehmen an."
	}
	g4 := frm.Abteilung != ""
	if frm.Abteilung == "" {
		errs["abteilung"] = "Bitte geben Sie Ihre Abteilung an."
	}
	g5 := frm.Position != ""
	if frm.Position == "" {
		errs["position"] = "Bitte geben Sie Ihre Position an."
	}
	g6a := frm.PLZ != ""
	if frm.PLZ == "" {
		errs["plz"] = "Bitte geben Sie Ihre PLZ an."
	}
	g6b := frm.Ort != ""
	if frm.Ort == "" {
		errs["ort"] = "Bitte geben Sie Ihren Ort an."
	}
	g6c := frm.Strasse != ""
	if frm.Strasse == "" {
		errs["strasse"] = "Bitte geben Sie Ihre Strasse an."
	}
	g7 := frm.Email != ""
	if frm.Email == "" {
		errs["email"] = "Bitte geben Sie eine Email an."
	}
	g8 := frm.Telefon != ""
	if frm.Telefon == "" {
		errs["telefon"] = "Bitte geben Sie eine Telefonnummer an."
	}
	//

	g10 := yearValid(frm.Geburtsjahr)
	if !g10 {
		errs["geburtsjahr"] = "Bitte geben Sie ein sinnvolles Geburtsjahr - oder gar nichts - ein."
	}
	g11 := yearValid(frm.Einstieg)
	if !g11 {
		errs["einstieg"] = "Bitte geben Sie ein sinnvolles Einstiegsjahr - oder gar nichts - ein."
	}

	g20 := frm.Terms
	if !frm.Terms {
		errs["terms"] = "Bitte nehmen Sie Kenntnis von unserer Datenschutz-Richtlinie."
	}
	fields := g1 && g2 && g3 && g4 && g5 && g6a && g6b && g6c && g7 && g8
	fields = fields && g10 && g11
	fields = fields && g20
	return errs, fields
}

var mtxFMT = sync.Mutex{}

// RegistrationFMTH shows a registraton form for the FMT
func RegistrationFMTH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s2f := struc2frm.New()
	s2f.ShowHeadline = true
	s2f.Indent = 170
	s2f.CSS = strings.ReplaceAll(
		s2f.CSS,
		"max-width: 40px;",
		"max-width: 220px;",
	)
	s2f.CSS += ` 
	* { 
		font-family: BlinkMacSystemFont, Segoe UI, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol; 
	}  `
	s2f.CSS += ` div.struc2frm span.postlabel { font-size: 80%; } `
	s2f.SetOptions("department", []string{"ub", "fm"}, []string{"UB", "FM"})
	s2f.SetOptions("geschlecht", []string{"", "male", "female"}, []string{"Bitte auswählen", "Männlich", "Weiblich"})
	s2f.SetOptions("leitung", []string{"", "<=10", "<=50", "<=100", "<=1000", ">1000"}, []string{" ", "bis 10", "bis 50", "bis 100", "bis 1000", "über 1000"})

	frm := formRegistrationFMT{}

	// pulling in values from http request
	populated, err := struc2frm.Decode(r, &frm)
	if populated && err != nil {
		s2f.AddError("global", fmt.Sprintf("cannot decode form: %v<br>\n <pre>%v</pre>", err, util.IndentedDump(r.Form)))
		log.Printf("cannot decode form: %v<br>\n <pre>%v</pre>", err, util.IndentedDump(r.Form))
	}

	// init values - multiple
	if !populated {
		if frm.Unternehmen == "" {
			frm.Unternehmen = ""
		}
	}

	errs, valid := frm.Validate()

	if populated {
		if !valid {
			s2f.AddErrors(errs) // add errors only for a populated form
			// render to HTML for user input / error correction
			// fmt.Fprint(w, s2f.Form(frm))
		} else {
			//
			// further processing with valid form data
			mtxFMT.Lock()
			defer mtxFMT.Unlock()

			var failureEmail, failureCSV bool

			to := []string{"finanzmarkttest@zew.de", "peter.buchmann@zew.de"}
			to = []string{"finanzmarkttest@zew.de"}
			mimeHTML := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
			body := &bytes.Buffer{}
			// headers
			fmt.Fprintf(body, "To: %v \r\n", strings.Join(to, ", ")) // "To: billy@microsoft.com, stevie@microsoft.com \r\n"
			fmt.Fprintf(body, mimeHTML)
			fmt.Fprintf(body, "FMT Registration %v %v\r\n", frm.Vorname, frm.Nachname)
			// ending of headers
			fmt.Fprint(body, "\r\n")
			s2f.CardViewOptions.SkipEmpty = true
			fmt.Fprint(body, s2f.Card(frm))
			fmt.Fprintf(body, "<p>Form sent %v</p>", time.Now().Format(time.RFC850))
			err = smtp.SendMail(
				// "hermes.zew-private.de:25", // from intern
				"hermes.zew.de:25",                // from DMZ - does not work
				nil,                               // smtp.Auth interface
				"Registration-FMT@survey2.zew.de", // from
				to,                                // twice - once here - and then again inside the body
				body.Bytes(),
			)
			if err != nil {
				// fmt.Fprint(w, fmt.Sprintf("Error sending email: %v <br>\n", err))
				log.Print(w, fmt.Sprintf(" Error sending email: %v", err))
				failureEmail = true
			}

			f, err := os.OpenFile("registration-fmt.csv", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>registration-fmt.csv konnte nicht geöffnet werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", err)
				failureCSV = true
			}
			defer f.Close()
			if _, err = f.WriteString(s2f.CSVLine(frm, ";")); err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten konnten nicht nach fmr.csv gespeichert werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", err)
				failureCSV = true
			}

			if failureEmail && failureCSV {
				return
			}
			fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten wurden gespeichert</p>")

		}
	}

	if !valid {
		w1 := &strings.Builder{}
		// fmt.Fprintf(w1,
		// 	`<p>REGISTRIERUNG VON FINANZMARKTTEST-TEILNEHMERN
		// 	</p>`)
		fmt.Fprint(w1, s2f.Form(frm))
		s2 := strings.ReplaceAll(w1.String(), "replace_me_1",
			`<div style="margin-top: 1.8em;">
			Ich erkläre mich mit den <a tabindex='-1' 
			href='https://www.zew.de/de/datenschutz' target='_blank' >Datenschutzbestimmungen</a> 
			einverstanden</div>`,
		)

		s3 := strings.ReplaceAll(s2, "replace_me_2",
			`
			<div style="
				margin:0.2rem  3rem;  
				margin-top:    1.4rem; 
				margin-bottom: 1.4rem; 
				max-width: 49rem;
				" 			
			>

				<label style="text-align: left; font-size: clamp(0.7rem, 0.86vw, 2.8rem); ">
					Wir werden Sie jeden Monat direkt nach der Umfrage über die aktuellen Ergebnisse per Email informieren. 
					Außerdem erhalten Sie von uns einige Tage später den ZEW-Finanzmarktreport mit detaillierten Analysen der Ergebnisse. 
					Den neuen Fragebogen senden wir Ihnen jeweils bei Umfragebeginn an Ihre Email-Adresse.
				</label> 

				<label style="text-align: left; font-size: clamp(0.7rem, 0.86vw, 2.8rem); ">
					Wir würden uns freuen, wenn Sie uns mit dieser Anmeldung noch 
					zusätzliche Angaben zu Ihrer Person machen könnten. 
					Wir werden diese Informationen in einigen wissenschaftlichen 
					Analysen zur Erwartungsbildung verwenden.
				</label> 
				
				<label style="text-align: left; font-size: clamp(0.7rem, 0.86vw, 2.8rem); ">
					Alle Informationen, die Sie uns mit dieser Anmeldung geben, 
					bleiben selbstverständlich anonym, 
					so dass keine Rückschlüsse auf Ihre Person oder Ihr Unternehmen möglich sind.
				</label> 

			</div>
		 `)

		s4 := strings.ReplaceAll(s3, "Form registration fmt", "Registrierung von Finanzmarkttest-Teilnehmern")
		fmt.Fprint(w, s4)

	}

}
