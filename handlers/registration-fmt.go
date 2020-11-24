package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/pbberlin/struc2frm"
	"github.com/zew/util"
)

type formRegistrationFMT struct {
	Vorname     string `json:"first_name"     form:"maxlength='40',size='25',label='Vorname',xxnobreak='true'"`
	Nachname    string `json:"last_name"      form:"maxlength='40',size='30',label='Name'"`
	Unternehmen string `json:"unternehmen"    form:"maxlength='40',size='40',placeholder='Ihr Unternehmen oder Organisation'"`
	Abteilung   string `json:"abteilung"      form:"maxlength='40',size='40',"`
	Position    string `json:"position"       form:"maxlength='40',size='40',suffix='Bezeichnung Ihrer aktuellen Position'"`
	Anschrift   string `json:"anschrift"      form:"maxlength='40',size='40',label='Postanschrift'"`
	// stackoverflow.com/questions/399078 - inside character classes escape ^-]\
	Email   string `json:"email"              form:"maxlength='40',size='40',pattern='[a-zA-Z0-9\\.\\-_%+]+@[a-zA-Z0-9\\.\\-]+\\.[a-zA-Z]{0&comma;2}'"`
	Telefon string `json:"telefon"            form:"maxlength='40',size='40',label='Telefon'"`

	Separator1 string `json:"separator1"      form:"subtype='separator'"`

	Geschlecht  string `json:"geschlecht"     form:"subtype='select'"`
	Geburtsjahr int    `json:"geburtsjahr"    form:"maxlength='5',size='5',min='0',max='2010'"`
	Abschluss   string `json:"abschluss"      form:"maxlength='40',size='40',label='Höchster Abschluss',suffix='z.B. Diplom'"`
	Studienfach string `json:"studienfach"    form:"maxlength='40',size='40',label='Ggf. Studienfach',suffix='z.B. VWL'"`
	Hochschule  string `json:"hochschule"     form:"maxlength='40',size='40',label='Ggf. Hochschule',suffix='z.B. Uni Mannheim'"`
	Einstieg    int    `json:"einstieg"       form:"maxlength='5',size='5',min='0',max='2010',label='Einstieg ins Berufsleben',suffix='Jahr'"`
	Leitung     int    `json:"Leitung"        form:"maxlength='6',size='6',min='0',max='1000000',label='Leitungsbefugnis',suffix='über Anzahl Mitarbeiter'"`

	Terms bool `json:"terms"        form:"label='Datenschutz',suffix='replace_me'"`
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
	g6 := frm.Anschrift != ""
	if frm.Anschrift == "" {
		errs["anschrift"] = "Bitte geben Sie Ihre Anschrift an."
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

	g12 := frm.Terms
	if !frm.Terms {
		errs["terms"] = "Bitte nehmen Sie Kenntnis von unserer Datenschutz-Richtlinie."
	}
	fields := g1 && g2 && g3 && g4 && g5 && g6 && g7 && g8
	fields = fields && g12
	return errs, fields
}

var mtxFMT = sync.Mutex{}

// RegistrationFMTH shows a registraton form for the FMT
func RegistrationFMTH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//
	fmt.Fprintf(w, "<h3>Registrierung von Finanzmarkttest-Teilnehmern</h3>")

	s2f := struc2frm.New()
	s2f.Indent = 170
	s2f.CSS = strings.ReplaceAll(s2f.CSS, "max-width: 40px;", "max-width: 220px;")
	s2f.CSS = strings.ReplaceAll(s2f.CSS, "div.struc2frm {", " * { font-family: BlinkMacSystemFont, Segoe UI, Helvetica, Arial, sans-serif, Apple Color Emoji, Segoe UI Emoji, Segoe UI Symbol; }  div.struc2frm {")
	s2f.SetOptions("department", []string{"ub", "fm"}, []string{"UB", "FM"})
	s2f.SetOptions("geschlecht", []string{"", "male", "female"}, []string{"Bitte auswählen", "Männlich", "Weiblich"})

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

			f, err := os.OpenFile("registration-fmt.csv", os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>registration-fmt.csv konnte nicht geöffnet werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", err)
				return
			}
			defer f.Close()
			if _, err = f.WriteString(s2f.CSVLine(frm)); err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten konnten nicht nach fmr.csv gespeichert werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", err)
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
		s2 := strings.ReplaceAll(w1.String(), "replace_me",
			`<div style="margin-top: 1.8em;">
			Ich erkläre mich mit den <a tabindex='-1' 
			href='https://www.zew.de/de/datenschutz' target='_blank' >Datenschutzbestimmungen</a> 
			einverstanden</div>`,
		)
		s3 := strings.ReplaceAll(s2, "<div class='separator'></div>",
			`
			<div style="margin-left:200px; max-width: 500px; margin-bottom: 1.4rem; font-size: 85%;" >
			<p> 1. Im ZEW Finanzmarktreport liefern wir Ihnen monatlich eine detaillierte Darstellung 
				der aktuellsten Umfrageergebnisse. <br>
				Den ZEW Finanzmarktreport sowie den monatlichen Finanzmarkttest-Fragebogen 
				schicken wir Ihnen grundsätzlich per Email zu. 
			</p> 
			
			<p> 2. Wir würden uns sehr darüber freuen, wenn Sie zusätzliche Angaben zu Ihrer Person machen könnten. <br>
				Ihre Daten bleiben anonym, so dass keine Rückschlüsse auf Ihre Person oder Ihr Unternehmen möglich sind.
			</p> 
			</div>
		 `)
		fmt.Fprint(w, s3)

	}

}
