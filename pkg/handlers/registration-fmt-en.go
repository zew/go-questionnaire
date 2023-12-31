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

type formRegistrationFMTEn struct {
	Vorname     string `json:"first_name"     form:"maxlength='40',size='25',label='Forename',xxnobreak='true'"`
	Nachname    string `json:"last_name"      form:"maxlength='40',size='30',label='Surname'"`
	Unternehmen string `json:"unternehmen"    form:"maxlength='40',size='40',label='Company',placeholder='your company or organization'"`
	Abteilung   string `json:"abteilung"      form:"maxlength='60',size='40',label='Department'"`
	Position    string `json:"position"       form:"maxlength='40',size='40',label='Position',suffix='Your current position'"`

	// Separator1 string `json:"separator1"      form:"subtype='separator',label=''"`

	Land string `json:"land"                  form:"maxlength='40',size='40',label='Country'"`
	// PLZ     string `json:"plz"             form:"maxlength='6',size='6',label='PLZ',xxnobreak='true'"`
	Ort     string `json:"ort"                form:"maxlength='120',size='40',label='City'"`
	Strasse string `json:"strasse"            form:"maxlength='120',size='40',label='Postal address',suffix=''"`

	// stackoverflow.com/questions/399078 - inside character classes escape ^-]\
	// the top level domain can be .info or longer
	Email   string `json:"email"              form:"maxlength='120',size='40',pattern='[a-zA-Z0-9\\.\\-_%+]+@[a-zA-Z0-9\\.\\-]+\\.[a-zA-Z]{0&comma;6}'"`
	Telefon string `json:"telefon"            form:"maxlength='120',size='40',label='Telefone'"`

	Separator2 string `json:"separator2"      form:"subtype='separator',label='replace_me_2'"`

	// Geschlecht  string `json:"geschlecht"     form:"subtype='select'"`
	Geburtsjahr string `json:"geburtsjahr"    form:"maxlength='5',size='5',label='Year of birth'"`
	Abschluss   string `json:"abschluss"      form:"maxlength='120',size='40',label='Highest qualification obtained',suffix='e.g. diploma'"`
	Studienfach string `json:"studienfach"    form:"maxlength='120',size='40',label='If applicable&comma; area of study',suffix='e.g. economics'"`
	Hochschule  string `json:"hochschule"     form:"maxlength='120',size='40',label='If applicable&comma; university',suffix='e.g. University of Mannheim'"`
	Einstieg    string `json:"einstieg"       form:"maxlength='5',size='5',label='Year of entry into employment  ',suffix='(year)'"`
	Leitung     string `json:"leitung"        form:"subtype='select',size='1',label='Management authority',suffix='number of co-workers'"`

	// Taetigkeiten
	Separator3           string `json:"separator3"                      form:"subtype='separator',label='replace_me_3'"`
	VWLAnalyse           string `json:"vwl_analyse"              form:"subtype='radiogroup',label='Economic analysis',label-style='min-width:240px;position: relative; left: -20px;'"`
	Wertpapierhandel     string `json:"wertpapierhandel"         form:"subtype='radiogroup',label='Securities trading',label-style='min-width:240px;position: relative; left: -20px;'"`
	Finanzierung         string `json:"finanzierung"             form:"subtype='radiogroup',label='Corporate finance',label-style='min-width:240px;position: relative; left: -20px;'"`
	Management           string `json:"management"               form:"subtype='radiogroup',label='Management',label-style='min-width:240px;position: relative; left: -20px;'"`
	Wertpapieranalyse    string `json:"wertpapieranalyse"        form:"subtype='radiogroup',label='Securities analysis',label-style='min-width:240px;position: relative; left: -20px;'"`
	Portfoliomanagement  string `json:"portfoliomanagement"      form:"subtype='radiogroup',label='Portfolio management',label-style='min-width:240px;position: relative; left: -20px;'"`
	Anlageberatung       string `json:"anlageberatung"           form:"subtype='radiogroup',label='Investment advisory services',label-style='min-width:240px;position: relative; left: -20px;'"`
	Vermoegensverwaltung string `json:"vermoegensverwaltung"     form:"subtype='radiogroup',label='Wealth management',label-style='min-width:240px;position: relative; left: -20px;'"`
	Risikomanagement     string `json:"risikomanagement"         form:"subtype='radiogroup',label='Risk management',label-style='min-width:240px;position: relative; left: -20px;'"`
	Sonstiges            string `json:"sonstiges"                form:"maxlength='40',size='40',label='Sonstiges',label-style='min-width:240px;position: relative; left: -20px;',suffix='sonstige Tätigkeiten'"`

	// VWLAnalyseHaupt           bool   `json:"vwl_analyse_haupt"               form:"label='Economic analysis',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// WertpapierhandelHaupt     bool   `json:"wertpapierhandel_haupt"          form:"label='Securities trading',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// FinanzierungHaupt         bool   `json:"finanzierung_haupt"              form:"label='Corporate finance',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// ManagementHaupt           bool   `json:"management_haupt"                form:"label='Management',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// WertpapieranalyseHaupt    bool   `json:"wertpapieranalyse_haupt"         form:"label='Securities analysis',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// PortfoliomanagementHaupt  bool   `json:"portfoliomanagement_haupt"       form:"label='Portfolio management',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// AnlageberatungHaupt       bool   `json:"anlageberatung_haupt"            form:"label='Investment advisory services',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// VermoegensverwaltungHaupt bool   `json:"vermoegensverwaltung_haupt"      form:"label='Wealth management',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// RisikomanagementHaupt     bool   `json:"risikomanagement_haupt"          form:"label='Risk management',label-style='min-width:240px;position: relative; left: -20px;',nobreak='true'"`
	// Sonstiges                 string `json:"Other activities"                form:"maxlength='40',size='40',label='Other activities',label-style='min-width:240px;position: relative; left: -20px;',suffix='other activities'"`

	Separator4 string `json:"separator4"   form:"subtype='separator',label=' &nbsp; '"`
	Terms      bool   `json:"terms"        form:"label='Data protection',suffix='replace_me_1'"`
}

func (frm formRegistrationFMTEn) Headline() string {
	return fmt.Sprintf("FMT Registration %v %v\r\n", frm.Vorname, frm.Nachname)
}

func (frm formRegistrationFMTEn) Validate() (map[string]string, bool) {
	errs := map[string]string{}
	g1 := frm.Vorname != ""
	if frm.Vorname == "" {
		errs["first_name"] = "Please enter your forename."
	}
	g2 := frm.Nachname != ""
	if frm.Nachname == "" {
		errs["last_name"] = "Please enter your surname."
	}
	g3 := frm.Unternehmen != ""
	if frm.Unternehmen == "" {
		errs["unternehmen"] = "Please enter your company/organization."
	}
	g4 := frm.Abteilung != ""
	if frm.Abteilung == "" {
		errs["abteilung"] = "Please enter your department."
	}
	g5 := frm.Position != ""
	if frm.Position == "" {
		errs["position"] = "Please enter your position."
	}
	g6a := frm.Land != ""
	if frm.Land == "" {
		errs["land"] = "Please enter your country."
	}
	g6b := frm.Ort != ""
	if frm.Ort == "" {
		errs["ort"] = "Please enter your city."
	}
	g6c := frm.Strasse != ""
	if frm.Strasse == "" {
		errs["strasse"] = "Please enter your postal address."
	}
	g7 := frm.Email != ""
	if frm.Email == "" {
		errs["email"] = "Please enter an email."
	}
	g8 := frm.Telefon != ""
	if frm.Telefon == "" {
		errs["telefon"] = "Please enter a phone number."
	}
	//

	g10 := yearValid(frm.Geburtsjahr)
	if !g10 {
		errs["geburtsjahr"] = "Please enter a reasonale year of birth - or leave the field empty."
	}
	g11 := yearValid(frm.Einstieg)
	if !g11 {
		errs["einstieg"] = "Please enter a reasonable year - or leave the field empty."
	}

	g20 := frm.Terms
	if !frm.Terms {
		errs["terms"] = "Please acknowledge our data protection policies."
	}
	fields := g1 && g2 && g3 && g4 && g5 && g6a && g6b && g6c && g7 && g8
	fields = fields && g10 && g11
	fields = fields && g20
	return errs, fields
}

// RegistrationFMTEnH shows a registraton form for the FMT
func RegistrationFMTEnH(w http.ResponseWriter, r *http.Request) {

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
	s2f.SetOptions("department", []string{"ub", "fm"}, []string{"UB", "FM"})
	s2f.SetOptions("geschlecht", []string{"", "male", "female", "diverse"}, []string{"Please choose", "male", "female", "diverse"})
	s2f.SetOptions("leitung", []string{"0", "<=10", "<=50", "<=100", "<=1000", ">1000"}, []string{"-", "up to 10", "up to 50", "up to 100", "up to 1000", "over 1000"})

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

	frm := formRegistrationFMTEn{}

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
			mtxFMT.Lock()
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
			fn := "registration-fmt-en.csv"
			fd, size := mustDir(fn)
			f, err := os.OpenFile(filepath.Join(fd, fn), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>%v could not be opened. Please inform peter.buchmann@zew.de.<br>%v</p>", fn, err)
				failureCSV = true
			}
			defer f.Close()
			if size < 10 {
				if _, err = f.WriteString(s2f.HeaderRow(frm, ";")); err != nil {
					fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Your data could not be saved to %v (header row). Informieren Sie peter.buchmann@zew.de.<br>%v</p>", fn, err)
					failureCSV = true
				}
			}
			if _, err = f.WriteString(s2f.CSVLine(frm, ";")); err != nil {
				fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Your data could not be saved to  %v. Please inform peter.buchmann@zew.de.<br>%v</p>", fn, err)
				failureCSV = true
			}

			if failureEmail && failureCSV {
				return
			}
			fmt.Fprintf(w, "<p style='color: red; font-size: 115%%;'>Your data was saved</p>")

		}
	}

	if !valid {
		w1 := &strings.Builder{}
		fmt.Fprint(w1, s2f.Form(frm))

		s2 := strings.ReplaceAll(w1.String(), "replace_me_1",
			`<div style="aamargin-top: 1.8em; max-width: 18rem; ">
			I acknowledge the  <a tabindex='-1' 
			href='https://www.zew.de/en/commitment-to-data-protection' target='_blank' >date protection terms</a> 
			</div>`,
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
					After your registration is complete we will send you every month the following information via email:

					<ul>
						<li style='margin-bottom: 0.4rem;'>
							Invitation to the new survey, including the current questionnaire.
						</li>
						<li style='margin-bottom: 0.4rem;'>
							Results of the survey (at publication date).
						</li>
					</ul>

				</label> 

				<label style="text-align: left; font-size: clamp(1.0rem, 0.86vw, 2.8rem); ">
					We would greatly appreciate it if you were able to provide additional information about yourself. 
					Your personal details will remain completely anonymous, 
					so that it will not be possible to trace this information back to yourself or your company. 
					These data will only be used for scientific purposes.
				</label> 
				

			</div>
		 `)

		s4 := strings.ReplaceAll(s3, "replace_me_3",
			`<div style='margin-left:1.6rem;margin-top:1.5rem;' >
				What are the main or occasional parts of your occupation?  <br>
				(multiple answers possible)<br>
				<div style='margin-top: 0.4rem; margin-left: 200px; '> &nbsp; &nbsp; &nbsp; Mainly     &nbsp;  &nbsp;  &nbsp; &nbsp; &nbsp; &nbsp;  Occasionally </div>
			</div>`)
		s5 := strings.ReplaceAll(
			s4,
			"<h3>Form registration fmten</h3>",
			fmt.Sprintf(`<h3>Registration for ZEW Financial Markets Survey <br>
			    - ZEW Indicator of Economic Sentiment  <br>
				&nbsp; <a href='%v' style='font-size: 70%%; font-weight: normal;' >German version</a>
			</h3>
			 `, cfg.Pref("/registrationfmtde")),
		)
		fmt.Fprint(w, s5)

	}

}
