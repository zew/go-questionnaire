package handlers

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pbberlin/dbg"
	"github.com/pbberlin/struc2frm"
	"github.com/zew/go-questionnaire/pkg/cfg"
)

type formRegistrationFMTDe struct {
	Vorname     string `json:"first_name"     form:"maxlength='40',size='25',label='Vorname',xxnobreak='true'"`
	Nachname    string `json:"last_name"      form:"maxlength='40',size='30',label='Name'"`
	Unternehmen string `json:"unternehmen"    form:"maxlength='40',size='40',label='Unternehmen',placeholder='Ihr Unternehmen oder Organisation'"`
	Abteilung   string `json:"abteilung"      form:"maxlength='60',size='40',label='Abteilung'"`
	Position    string `json:"position"       form:"maxlength='40',size='40',label='Position',suffix='Bezeichnung Ihrer aktuellen Position'"`

	// Separator1 string `json:"separator1"      form:"subtype='separator',label=''"`

	PLZ     string `json:"plz"                form:"maxlength='6',size='6',label='PLZ',xxnobreak='true'"`
	Ort     string `json:"ort"                form:"maxlength='120',size='40',label='Ort'"`
	Strasse string `json:"strasse"            form:"maxlength='120',size='40',label='Strasse',suffix='mit Hausnummer'"`
	// stackoverflow.com/questions/399078 - inside character classes escape ^-]\
	// the top level domain can be .info or longer
	Email   string `json:"email"              form:"maxlength='120',size='40',pattern='[a-zA-Z0-9\\.\\-_%+]+@[a-zA-Z0-9\\.\\-]+\\.[a-zA-Z]{0&comma;6}'"`
	Telefon string `json:"telefon"            form:"maxlength='120',size='40',label='Telefon'"`

	Separator2 string `json:"separator2"      form:"subtype='separator',label='replace_me_2'"`

	Geschlecht  string `json:"geschlecht"     form:"subtype='select'"`
	Geburtsjahr string `json:"geburtsjahr"    form:"maxlength='5',size='5',label='Geburtsjahr'"`
	Abschluss   string `json:"abschluss"      form:"maxlength='120',size='40',label='Höchster Abschluss',suffix='z.B. Diplom'"`
	Studienfach string `json:"studienfach"    form:"maxlength='120',size='40',label='Ggf. Studienfach',suffix='z.B. VWL'"`
	Hochschule  string `json:"hochschule"     form:"maxlength='120',size='40',label='Ggf. Hochschule',suffix='z.B. Uni Mannheim'"`
	Einstieg    string `json:"einstieg"       form:"maxlength='5',size='5',label='Einstieg ins Berufsleben',suffix='(Jahr)'"`
	Leitung     string `json:"leitung"        form:"subtype='select',size='1',label='Leitungsbefugnis über',suffix='Mitarbeiter'"`

	// Taetigkeiten
	Separator3           string `json:"separator3"               form:"subtype='separator',label='replace_me_3'"`
	VWLAnalyse           string `json:"vwl_analyse"              form:"subtype='radiogroup',label='Volkswirtschaftl. Analyse',label-style='min-width:240px;position: relative; left: -20px;'"`
	Wertpapierhandel     string `json:"wertpapierhandel"         form:"subtype='radiogroup',label='Wertpapierhandel',label-style='min-width:240px;position: relative; left: -20px;'"`
	Finanzierung         string `json:"finanzierung"             form:"subtype='radiogroup',label='Finanzierung',label-style='min-width:240px;position: relative; left: -20px;'"`
	Management           string `json:"management"               form:"subtype='radiogroup',label='Management',label-style='min-width:240px;position: relative; left: -20px;'"`
	Wertpapieranalyse    string `json:"wertpapieranalyse"        form:"subtype='radiogroup',label='Wertpapieranalyse',label-style='min-width:240px;position: relative; left: -20px;'"`
	Portfoliomanagement  string `json:"portfoliomanagement"      form:"subtype='radiogroup',label='Fonds-/Portfoliomanagmt.',label-style='min-width:240px;position: relative; left: -20px;'"`
	Anlageberatung       string `json:"anlageberatung"           form:"subtype='radiogroup',label='Anlageberatung',label-style='min-width:240px;position: relative; left: -20px;'"`
	Vermoegensverwaltung string `json:"vermoegensverwaltung"     form:"subtype='radiogroup',label='Vermögensverwaltung',label-style='min-width:240px;position: relative; left: -20px;'"`
	Risikomanagement     string `json:"risikomanagement"         form:"subtype='radiogroup',label='Risikomanagement',label-style='min-width:240px;position: relative; left: -20px;'"`
	Sonstiges            string `json:"sonstiges"                form:"maxlength='40',size='40',label='Sonstiges',label-style='min-width:240px;position: relative; left: -20px;',suffix='sonstige Tätigkeiten'"`

	Separator4 string `json:"separator4"   form:"subtype='separator',label=' &nbsp; '"`
	Terms      bool   `json:"terms"        form:"label='Datenschutz',suffix='replace_me_1'"`
}

func (frm formRegistrationFMTDe) Headline() string {
	return fmt.Sprintf("FMT Registration %v %v\r\n", frm.Vorname, frm.Nachname)
}

func (frm formRegistrationFMTDe) Validate() (map[string]string, bool) {
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

// RegistrationFMTDeH shows a registraton form for the FMT
func RegistrationFMTDeH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	s2f := struc2frm.New()
	s2f.ShowHeadline = true
	s2f.Indent = 180
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
	s2f.CSS += ` div.struc2frm input[type=radio] { margin-left: 0.8rem ; margin-right: 4.8rem } `
	s2f.SetOptions("geschlecht", []string{"", "male", "female", "diverse"}, []string{"Bitte auswählen", "Männlich", "Weiblich", "Divers"})
	s2f.SetOptions("leitung", []string{"0", "<=10", "<=50", "<=100", "<=1000", ">1000"}, []string{"-", "bis 10", "bis 50", "bis 100", "bis 1000", "über 1000"})

	depth := []string{
		"vwl_analyse",
		"wertpapierhandel",
		"finanzierung",
		"management",
		"wertpapieranalyse",
		"portfoliomanagement",
		"anlageberatung",
		"vermoegensverwaltung",
		"risikomanagement",
	}
	for _, s := range depth {
		s2f.SetOptions(s, []string{"primary", "secondary"}, []string{"", ""})
	}

	frm := formRegistrationFMTDe{}

	// pulling in values from http request
	populated, err := struc2frm.Decode(r, &frm)
	if populated && err != nil {
		s2f.AddError("global", fmt.Sprintf("cannot decode form: %v<br>\n <pre>%v</pre>", err, dbg.Dump2String(r.Form)))
		log.Printf("cannot decode form: %v<br>\n <pre>%v</pre>", err, dbg.Dump2String(r.Form))
	}

	// init values - multiple
	if !populated {
		// if frm.Unternehmen == "" {
		// 	frm.Unternehmen = ""
		// }
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
			log.Printf("trying to acquire mtxFMT lock")
			mtxFMT.Lock()
			log.Printf("mtxFMT lock acquired")
			defer mtxFMT.Unlock()

			var failureEmail, failureCSV bool

			mimeHTML := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
			body := &bytes.Buffer{}
			// headers
			fmt.Fprintf(body, "To: %v \r\n", strings.Join(adminEmail(), ", ")) // "To: billy@microsoft.com, stevie@microsoft.com \r\n"
			fmt.Fprintf(body, mimeHTML)
			fmt.Fprintf(body, frm.Headline())
			// ending of headers
			fmt.Fprint(body, "\r\n")
			s2f.CardViewOptions.SkipEmpty = true
			fmt.Fprint(body, s2f.Card(frm))
			fmt.Fprintf(body, "<p>Form sent %v</p>", time.Now().Format(time.RFC850))

			err = isPortOpen( emailHost(), 4000*time.Second)
			if err != nil {
				log.Print(w, fmt.Sprintf(" Error connecting to %v: %v", emailHost(), err))
				failureEmail = true
			} else {
				err = smtp.SendMail(
					emailHost(),
					nil,                               // smtp.Auth interface
					"Registration-FMT@survey2.zew.de", // from
					adminEmail(),                      // twice - once here - and then again inside the body
					body.Bytes(),
				)
				if err != nil {
					// fmt.Fprint(w, fmt.Sprintf("Error sending email: %v <br>\n", err))
					log.Print(w, fmt.Sprintf(" Error sending email: %v", err))
					failureEmail = true
				}
			}

			err = smtp.SendMail(
				emailHost(),
				nil,                               // smtp.Auth interface
				"Registration-FMT@survey2.zew.de", // from
				adminEmail(),                      // twice - once here - and then again inside the body
				body.Bytes(),
			)
			if err != nil {
				// fmt.Fprint(w, fmt.Sprintf("Error sending email: %v <br>\n", err))
				log.Print(w, fmt.Sprintf(" Error sending email: %v", err))
				failureEmail = true
			}
			fn := "registration-fmt-de.csv"
			fd, size := mustDir(fn)
			f, err := os.OpenFile(filepath.Join(fd, fn), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>%v konnte nicht geöffnet werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", fn, err)
				failureCSV = true
			}
			defer f.Close()
			if size < 10 {
				if _, err = f.WriteString(s2f.HeaderRow(frm, ";")); err != nil {
					fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten konnten nicht nach %v gespeichert werde (header row). Informieren Sie peter.buchmann@zew.de.<br>%v</p>", fn, err)
					failureCSV = true
				}
			}
			if _, err = f.WriteString(s2f.CSVLine(frm, ";")); err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Ihre Daten konnten nicht nach %v gespeichert werden. Informieren Sie peter.buchmann@zew.de.<br>%v</p>", fn, err)
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
		fmt.Fprint(w1, s2f.Form(frm))

		s2 := strings.ReplaceAll(w1.String(), "replace_me_1",
			`<div style="aamargin-top: 1.8em; max-width: 18rem; ">
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

				<label style="text-align: left; font-size: clamp(1.0rem, 0.86vw, 2.8rem); white-space: normal; ">
					Nach ihrer erfolgreichen Anmeldung erhalten Sie monatlich folgende Dokumente per E-Mail:

					<ul>
						<li style='margin-bottom: 0.4rem;'>
							Den ZEW-Finanzmarkttest-Fragebogen, jeweils zum Umfragebeginn
						</li>
						<li style='margin-bottom: 0.4rem;'>
							Die Umfrageergebnisse, jeweils zum Veröffentlichungszeitpunkt
						</li>
						<li style='margin-bottom: 0.4rem;'>
							Den ZEW-Finanzmarktteport, in dem die Ergebnisse detailliert analysiert werden,
							einige Tage nach Veröffentlichung der jeweils neusten Umfrageergebnisse.
						</li>
					</ul>

				</label>

				<label style="text-align: left; font-size: clamp(1.0rem, 0.86vw, 2.8rem); ">
					Wir würden uns freuen, wenn Sie uns noch zusätzliche Angaben zu Ihrer Person machen könnten.
					Diese Informationen werden <u><b>ausschließlich für wissenschaftliche Zwecke</b></u> verwendet.
					Alle Informationen, die Sie uns mit dieser Anmeldung geben, bleiben anonym,
					so dass keine Rückschlüsse auf Ihre Person oder Ihr Unternehmen möglich sind.
				</label>


			</div>
		 `)

		s4 := strings.ReplaceAll(s3, "replace_me_3",
			`<div style='margin-left:1.6rem;margin-top:1.5rem;' >
				Welche Tätigkeiten führen Sie beruflich aus? <br>
				(Mehrfachantwort möglich)<br>
				<div style='margin-top: 0.4rem; margin-left: 200px; '>Haupttätigkeit   &nbsp; &nbsp;  Gelegentliche Tätigkeit</div>
			</div>`)
		s5 := strings.ReplaceAll(
			s4,
			"<h3>Form registration fmtde</h3>",
			fmt.Sprintf(`<h3>Registrierung für Teilnahme am ZEW Index /
				ZEW Finanzmarkttest <br>
				&nbsp; <a href='%v' style='font-size: 70%%; font-weight: normal;' >Englische Version</a>
			</h3>
			 `, cfg.Pref("/registrationfmten")),
		)
		fmt.Fprint(w, s5)

	}

}
