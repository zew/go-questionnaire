package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/zew/go-questionnaire/pkg/cfg"
)

var mtxUnsubscribe = sync.Mutex{}

const fileCSV = "unsubscribe.csv"

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
	Path    string `json:"path"`  // the original URL path
	Query   string `json:"query"` // the original URL raw query string

	Response template.HTML `json:"response"`
}

func (us unsubscribeT) String() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "Project:  %v\n", us.Project)
	fmt.Fprintf(b, "Task:     %v\n", us.Task)
	fmt.Fprintf(b, "Email:    %v\n", us.Email)
	fmt.Fprintf(b, "Date:     %v\n", us.Date)
	fmt.Fprintf(b, "Response: %v\n", us.Response)
	fmt.Fprintf(b, "Path:     %v\n", us.Path)
	fmt.Fprintf(b, "Query:    %v\n", us.Query)
	return b.String()
}
func (us unsubscribeT) CSVHeader() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", "project")
	fmt.Fprintf(b, "%v;", "task")
	fmt.Fprintf(b, "%v;", "email")
	fmt.Fprintf(b, "%v;", "date")
	fmt.Fprintf(b, "%v;", "path")
	fmt.Fprintf(b, "%v;", "query")
	fmt.Fprint(b, "\n")
	return b.String()
}
func (us unsubscribeT) CSVRow() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", us.Project)
	fmt.Fprintf(b, "%v;", us.Task)
	fmt.Fprintf(b, "%v;", us.Email)
	fmt.Fprintf(b, "%v;", us.Date)
	fmt.Fprintf(b, "%v;", us.Path)
	fmt.Fprintf(b, "%v;", us.Query)
	fmt.Fprint(b, "\n")
	return b.String()
}

func saveCSVRow(fe unsubscribeT) (string, error) {

	mtxUnsubscribe.Lock()
	defer mtxUnsubscribe.Unlock()

	fd, size := mustDir(fileCSV)
	f, err := os.OpenFile(filepath.Join(fd, fileCSV), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Sprintf("Could not open  %v. Email %v.", fileCSV, adminEmail()), err
	}
	defer f.Close()
	if size < 10 {
		if _, err = f.WriteString(fe.CSVHeader()); err != nil {
			return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV header). Informieren Sie %v.", fileCSV, adminEmail()), err
		}
	}
	if _, err = f.WriteString(fe.CSVRow()); err != nil {
		return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV row). Informieren Sie %v.", fileCSV, adminEmail()), err
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

	//
	// data from URL parameters
	//   email clients somehow garble those parameters
	fe := unsubscribeT{} // form entry
	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err := dec.Decode(&fe, r.Form)
	if err != nil {
		errMsg += fmt.Sprintf("Decoding error: %v\n", err)
	}

	//
	// data from URL path
	pths := strings.Split(r.URL.Path, "/") // path.SplitList(r.URL.Path) does not exist
	// pths[0...x] is "/.../.../unsubscribe/proj/tsk/someemail"
	// from the end...
	ln := len(pths) - 1
	// third last path element: project
	if pths[ln-2] != "" {
		fe.Project = pths[ln-2]
	}
	// second last path element: task
	if pths[ln-1] != "" {
		fe.Task = pths[ln-1]
	}
	// last path element: email
	if pths[ln-0] != "" {
		fe.Email = pths[ln-0]
		fe.Email = strings.ReplaceAll(fe.Email, "pct40", "@")
	}

	//
	// time.DateTime was introduced after go version 1.17 => codecov fails
	// fe.Date = time.Now().Format(time.DateTime)
	fe.Date = time.Now().Format("2006-01-02 15:04:05")

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

	// dbg.Dump(fe)

	// no errors
	if fe.Response == "" {
		// fe.Path = r.RequestURI
		fe.Path = r.URL.Path
		fe.Query = r.URL.RawQuery
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

// DownloadUnsubscribe responds CSV file;
// test using
// CURLOPT_SSL_VERIFYPEER  => --insecure
// curl -i --insecure https://localhost:8083/survey/unsubscribe-download -o tmp-unsubscribe-local.csv
// curl -i            https://survey2.zew.de/unsubscribe-download        -o tmp-unsubscribe-remote.csv
func DownloadUnsubscribe(w http.ResponseWriter, r *http.Request) {

	// morsels only => lame security
	cond1 := strings.Contains(r.RemoteAddr, "127.0.0.1")
	cond2 := strings.Contains(r.RemoteAddr, "[::1]")
	cond3 := strings.Contains(r.RemoteAddr, "193.196.")
	// cond3 := strings.Contains(r.RemoteAddr, "193.196.11")
	if cond1 || cond2 || cond3 {
		// proceed below
	} else {
		s := fmt.Sprintf("Only local or subnet access to %v. Email %v.", fileCSV, adminEmail())
		log.Print(s)
		http.Error(w, s, 400)
		return
	}

	//
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%v"`, fileCSV))
	fd, size := mustDir(fileCSV)
	w.Header().Add("Content-Length", fmt.Sprint(size))

	log.Printf("   %v - %v Bytes from --%v-- start", filepath.Join(fd, fileCSV), size, r.RemoteAddr)

	mtxUnsubscribe.Lock()
	defer mtxUnsubscribe.Unlock()
	bts, err := os.ReadFile(filepath.Join(fd, fileCSV))
	if err != nil {
		s := fmt.Sprintf("Could not open  %v. Email %v.", fileCSV, adminEmail())
		log.Print(s)
		http.Error(w, s, 400)
		return
	}
	// this fails; server closes network connection, before bts is written to w
	//    fmt.Fprint(w, bts)
	// we need use 	io.Copy(out, in)
	http.ServeContent(w, r, fileCSV, time.Now(), bytes.NewReader(bts))

	log.Printf("   %v - %v Bytes from --%v-- end", filepath.Join(fd, fileCSV), size, r.RemoteAddr)

}
