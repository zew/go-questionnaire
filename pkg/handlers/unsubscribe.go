package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/pbberlin/dbg"
	"github.com/zew/go-questionnaire/pkg/cfg"
)

var mtxUnsubscribe = sync.Mutex{}

func emailHost() string {
	emailHorst := "hermes.zew-private.de:25" // developer machine - must be inside ZEW or ZEW VPN
	if cfg.Get().IsProduction {
		emailHorst = "hermes.zew.de:25" // from DMZ - does not work
	}
	return emailHorst
}

func adminEmail() []string {
	to := []string{"peter.buchmann@zew.de"}
	if cfg.Get().IsProduction {
		to = []string{"finanzmarkttest@zew.de", "peter.buchmann@zew.de"}
	}
	return to
}

type unsubscribeT struct {
	Project string `json:"project"`
	Task    string `json:"task"`
	Email   string `json:"email"`
	Date    string `json:"date"`

	Response template.HTML `json:"response"`
}

func (us unsubscribeT) String() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "Project:  %v\n", us.Project)
	fmt.Fprintf(b, "Task:     %v\n", us.Task)
	fmt.Fprintf(b, "Email:    %v\n", us.Email)
	fmt.Fprintf(b, "Date:     %v\n", us.Date)
	fmt.Fprintf(b, "Response: %v\n", us.Response)
	return b.String()
}
func (us unsubscribeT) CSVHeader() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", "project")
	fmt.Fprintf(b, "%v;", "task")
	fmt.Fprintf(b, "%v;", "email")
	fmt.Fprintf(b, "%v;", "date")
	fmt.Fprint(b, "\n")
	return b.String()
}
func (us unsubscribeT) CSVRow() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", us.Project)
	fmt.Fprintf(b, "%v;", us.Task)
	fmt.Fprintf(b, "%v;", us.Email)
	fmt.Fprintf(b, "%v;", us.Date)
	fmt.Fprint(b, "\n")
	return b.String()
}

// UnsubscribeH creates a generic CSV file containing requests to be removed from go-massmail tasks
func UnsubscribeH(w http.ResponseWriter, r *http.Request) {

	errMsg := ""

	//
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	src := `<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8" />
    <title>hash logins</title>
	<style>
	* {font-family: monospace;}
	</style>
</head>
<body>

<form method="get"   style='white-space:pre' >
    <b>Unsubscribe</b>
	Project     <input type="text"   name="project"   value="{{.Project}}"> <br>
	Task        <input type="text"   name="task"      value="{{.Task}}">    <br>
	Email       <input type="text"   name="email"     value="{{.Email}}">   <br>
	            <button id="submit" accesskey="s" autofocus>  <u>S</u>ubmit  </button>  <br>

	{{if gt (len .Response   ) 0 }} <p style="font-size:120%">{{.Response   }}</p>{{end}}
</form>
	
</body>
</html>
`

	fe := unsubscribeT{} // form entry

	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err := dec.Decode(&fe, r.Form)
	if err != nil {
		errMsg += fmt.Sprintf("Decoding error: %v\n", err)
	}

	fe.Date = time.Now().Format(time.DateTime)

	if fe.Project == "" {
		msg := "Project cannot be empty; 'fmt' or 'pds'"
		fe.Response += template.HTML(msg + "<br>")
	}
	if fe.Task == "" {
		msg := "Task cannot be empty; 'invitation' or 'reminder'"
		fe.Response += template.HTML(msg + "<br>")
	}
	if fe.Email == "" {
		msg := "Email cannot be empty"
		fe.Response += template.HTML(msg + "<br>")
	}

	dbg.Dump(fe)

	if fe.Response == "" {
		msg1, err := saveCSVRow(fe)
		if err != nil {
			fe.Response += template.HTML(err.Error() + "<br>")
		} else {
			fe.Response += template.HTML(msg1 + "<br>")
			msg2, err := sendAdminEmail(fe)
			if err != nil {
				fe.Response += template.HTML(err.Error() + "<br>")
			}
			fe.Response += template.HTML(msg2 + "<br>")
		}
	}

	//
	tpl := template.New("anyname.html")
	tpl, err = tpl.Parse(src)
	if err != nil {
		fmt.Fprintf(w, "Error parsing inline template: %v", err)
	}

	err = tpl.Execute(w, fe)
	if err != nil {
		fmt.Fprintf(w, "Error executing inline template: %v", err)
	}

}

func saveCSVRow(fe unsubscribeT) (string, error) {

	mtxUnsubscribe.Lock()
	defer mtxUnsubscribe.Unlock()

	fn := "unsubscribe.csv"

	fd, size := mustDir(fn)
	f, err := os.OpenFile(filepath.Join(fd, fn), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Sprintf("Datei %v konnte nicht geöffnet werden. Informieren Sie %v.", fn, adminEmail()), err
	}
	defer f.Close()
	if size < 10 {
		if _, err = f.WriteString(fe.CSVHeader()); err != nil {
			return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV header). Informieren Sie %v.", fn, adminEmail()), err
		}
	}
	if _, err = f.WriteString(fe.CSVRow()); err != nil {
		return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV row). Informieren Sie %v.", fn, adminEmail()), err
	}

	return "Ihre Daten wurden gespeichert", nil

}

func sendAdminEmail(fe unsubscribeT) (string, error) {

	mimeHTML := "MIME-version: 1.0;\r\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	b := &bytes.Buffer{}

	// headers start
	fmt.Fprintf(b, "To: %v \r\n", strings.Join(adminEmail(), ", ")) // "To: billy@microsoft.com, stevie@microsoft.com \r\n"
	fmt.Fprint(b, mimeHTML)
	fmt.Fprintf(b, "Subject: Unsubscribe request for %v-%v\r\n", fe.Project, fe.Task)
	fmt.Fprint(b, "\r\n")
	// headers stop

	// body
	fmt.Fprint(b, "<pre>\r\n")
	fmt.Fprint(b, fe.String())
	fmt.Fprint(b, "</pre>\r\n")
	fmt.Fprintf(b, "<p>Email sent %v</p>", time.Now().Format(time.RFC850))

	err := smtp.SendMail(
		emailHost(),
		nil,                          // smtp.Auth interface
		"unsubscribe@survey2.zew.de", // from
		adminEmail(),                 // twice - once here - and then again inside the body
		b.Bytes(),
	)
	if err != nil {
		return "Error sending email", err
	}

	return "Email to admin sent", nil

}
