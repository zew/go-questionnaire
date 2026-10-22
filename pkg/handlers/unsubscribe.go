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

const csvPost = "unsubscribe-post.csv"
const csvGet = "unsubscribe-get.csv"

func emailHost() string {
	return cfg.Get().SMTPHost.HostNamePort
}

func adminEmail() []string {
	return cfg.Get().AdminEmails
}

func restoreObfuscation(s string) string {
	s = strings.ReplaceAll(s, "mmpp", "&")
	s = strings.ReplaceAll(s, "qquu", "=")
	s = strings.ReplaceAll(s, "ddtt", ".")
	s = strings.ReplaceAll(s, "aatt", "@")
	s = strings.ReplaceAll(s, "hhyy", "-")
	s = strings.ReplaceAll(s, "pct40", "@")
	return strings.TrimSpace(s)
}

type unsubscribeT struct {
	Project string `json:"project"`
	Task    string `json:"task"`
	All     bool   `json:"all"` // all tasks
	Email   string `json:"email"`
	Date    string `json:"date"`
	Path    string `json:"path"`  // the original URL path
	Query   string `json:"query"` // the original URL raw query string

	EmailPath  string `json:"-"` // email extracted from URL path
	EmailQuery string `json:"-"` // email extracted from URL GET param

	Response template.HTML `json:"response"`
}

func (us unsubscribeT) CSVHeader() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", "project")
	fmt.Fprintf(b, "%v;", "task")
	// fmt.Fprintf(b, "%v;", "all")
	fmt.Fprintf(b, "%v;", "email")
	fmt.Fprintf(b, "%v;", "date")
	fmt.Fprintf(b, "%v;", "path")
	fmt.Fprintf(b, "%v;", "query") // remove trailing ;  at next release
	fmt.Fprint(b, "\n")
	return b.String()
}
func (us unsubscribeT) CSVRow() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", us.Project)
	fmt.Fprintf(b, "%v;", us.Task)
	// fmt.Fprintf(b, "%v;", us.All)
	fmt.Fprintf(b, "%v;", us.Email)
	fmt.Fprintf(b, "%v;", us.Date)
	fmt.Fprintf(b, "%v;", us.Path)
	fmt.Fprintf(b, "%v;", us.Query) // remove trailing ;  at next release
	fmt.Fprint(b, "\n")
	return b.String()
}

func (us unsubscribeT) CSVHeaderGet() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", "project")
	fmt.Fprintf(b, "%v;", "task")
	fmt.Fprintf(b, "%v;", "email")
	fmt.Fprintf(b, "%v;", "email_path")
	fmt.Fprintf(b, "%v;", "email_query")
	fmt.Fprintf(b, "%v;", "date")
	fmt.Fprintf(b, "%v;", "path")
	fmt.Fprintf(b, "%v;", "query")
	fmt.Fprint(b, "\n")
	return b.String()
}

func (us unsubscribeT) CSVRowGet() string {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "%v;", us.Project)
	fmt.Fprintf(b, "%v;", us.Task)
	fmt.Fprintf(b, "%v;", us.Email)
	fmt.Fprintf(b, "%v;", us.EmailPath)
	fmt.Fprintf(b, "%v;", us.EmailQuery)
	fmt.Fprintf(b, "%v;", us.Date)
	fmt.Fprintf(b, "%v;", us.Path)
	fmt.Fprintf(b, "%v;", us.Query)
	fmt.Fprint(b, "\n")
	return b.String()
}

func (us unsubscribeT) String() string {
	// prevent repetition...
	b := &bytes.Buffer{}
	hds := strings.Split(strings.TrimSpace(us.CSVHeader()), ";")
	els := strings.Split(strings.TrimSpace(us.CSVRow()), ";")
	for i := 0; i < len(els); i++ {
		fmt.Fprintf(b, "%-20v: %v\n", hds[i], els[i])
	}
	return b.String()
}

// saving a row to the CSV POST file
func saveRowPost(fe unsubscribeT) (string, error) {

	mtxUnsubscribe.Lock()
	defer mtxUnsubscribe.Unlock()

	fd, size := mustDir(csvPost)
	f, err := os.OpenFile(filepath.Join(fd, csvPost), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Sprintf("Could not open  %v. Email %v.", csvPost, adminEmail()), err
	}
	defer f.Close()
	if size < 10 {
		if _, err = f.WriteString(fe.CSVHeader()); err != nil {
			return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV header). Informieren Sie %v.", csvPost, adminEmail()), err
		}
	}
	if _, err = f.WriteString(fe.CSVRow()); err != nil {
		return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV row). Informieren Sie %v.", csvPost, adminEmail()), err
	}

	return "Ihre Daten wurden gespeichert", nil

}

// saving a row to the CSV GET file
func saveRowGet(fe unsubscribeT) (string, error) {

	mtxUnsubscribe.Lock()
	defer mtxUnsubscribe.Unlock()

	fd, size := mustDir(csvGet)
	f, err := os.OpenFile(filepath.Join(fd, csvGet), os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Sprintf("Could not open  %v. Email %v.", csvGet, adminEmail()), err
	}
	defer f.Close()
	if size < 10 {
		if _, err = f.WriteString(fe.CSVHeaderGet()); err != nil {
			return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV header). Informieren Sie %v.", csvGet, adminEmail()), err
		}
	}
	if _, err = f.WriteString(fe.CSVRowGet()); err != nil {
		return fmt.Sprintf("Ihre Daten konnten nicht nach %v gespeichert werden (CSV row). Informieren Sie %v.", csvGet, adminEmail()), err
	}

	return "GET request logged", nil

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

<form method="post"  style='white-space:pre' >
    <b>Unsubscribe</b>
	Project     <input type="text"   name="project"   value="{{.Project}}"> <br>
	Task        <input type="text"   name="task"      value="{{.Task}}">    <br>
	Email       <input type="text"   name="email"     value="{{.Email}}">   <br>

	<br>
	All tasks    <input type="checkbox"  name="all"   value="true"  >      <br>

	            <button id="submit" accesskey="s" autofocus>  <u>S</u>ubmit  </button>  <br>

	{{if gt (len .Response   ) 0 }} <p style="font-size:120%">{{.Response   }}</p>{{end}}
</form>
	
</body>
</html>
`

	fe := unsubscribeT{} // form entry

	//
	// parse form first
	err := r.ParseForm()
	if err != nil {
		errMsg += fmt.Sprintf("ParseForm error: %v\n", err)
	}
	fe.EmailQuery = restoreObfuscation(r.URL.Query().Get("email"))

	//
	// data from URL path
	var pathProject, pathTask, pathEmail string
	dirs := strings.Split(r.URL.Path, "/") // path.SplitList(r.URL.Path) does not exist
	// pths[0...x] is "/.../.../unsubscribe/proj/tsk/someemail/dummy"
	if len(dirs) > 0 {
		dirs = dirs[:len(dirs)-1] // cut off the last element - "dummy"
	}
	// from the end...
	if len(dirs) > 3 {
		ln := len(dirs) - 1
		// third last path element: project
		if dirs[ln-2] != "" {
			pathProject = restoreObfuscation(dirs[ln-2])
		}
		// second last path element: task
		if dirs[ln-1] != "" {
			pathTask = restoreObfuscation(dirs[ln-1])
		}
		// last path element: email
		if dirs[ln-0] != "" {
			pathEmail = restoreObfuscation(dirs[ln-0])
		}
	}
	fe.EmailPath = pathEmail

	//
	// data from URL parameters or POST body
	//   email clients somehow garble those parameters
	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	if r.Method == http.MethodPost {
		// on POST: set URL path values as defaults first, then decode PostForm so user edits take precedence
		fe.Project = pathProject
		fe.Task = pathTask
		fe.Email = pathEmail
		err = dec.Decode(&fe, r.PostForm)
		if err != nil {
			errMsg += fmt.Sprintf("Decoding error: %v\n", err)
		}
	} else {
		// on GET: decode query parameters first, and do not overwrite non-empty values with URL path
		err = dec.Decode(&fe, r.Form)
		if err != nil {
			errMsg += fmt.Sprintf("Decoding error: %v\n", err)
		}
		if fe.Project == "" {
			fe.Project = pathProject
		}
		if fe.Task == "" {
			fe.Task = pathTask
		}
		if fe.Email == "" {
			fe.Email = pathEmail
		}
	}

	// all supersedes task
	if fe.All {
		fe.Task = "all"
	}

	//
	// time.DateTime was introduced after go version 1.17 => codecov fails
	// fe.Date = time.Now().Format(time.DateTime)
	fe.Date = time.Now().Format("2006-01-02 15:04:05")
	fe.Path = r.URL.Path
	fe.Query = r.URL.RawQuery

	if r.Method != http.MethodPost {

		// log all GET requests into separate CSV file and return 200 OK with confirmation form
		_, _ = saveRowGet(fe)
		if fe.Project != "" && fe.Task != "" && fe.Email != "" && fe.Email != "your-email" {
			fe.Response += template.HTML("Please click Submit to confirm your unsubscribe request.<br>")
		}
		w.WriteHeader(http.StatusOK)

	} else {

		if fe.Project == "" {
			msg := "Project cannot be empty; 'fmt' or 'pds'"
			fe.Response += template.HTML(msg + "<br>")
		}
		if fe.Task == "" {
			msg := "Task cannot be empty; 'invitation' or 'reminder'"
			fe.Response += template.HTML(msg + "<br>")
		}
		if fe.Email == "" || fe.Email == "your-email" {
			msg := "Email cannot be empty"
			fe.Response += template.HTML(msg + "<br>")
		}

		// dbg.Dump(fe)

		// no errors
		if fe.Response == "" {
			msg1, err := saveRowPost(fe)
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
		s := fmt.Sprintf("Only local or subnet access to %v. Email %v.", csvPost, adminEmail())
		log.Print(s)
		http.Error(w, s, 400)
		return
	}

	//
	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Add("Content-Disposition", fmt.Sprintf(`attachment; filename="%v"`, csvPost))
	fd, size := mustDir(csvPost)
	w.Header().Add("Content-Length", fmt.Sprint(size))

	log.Printf("   %v - %v Bytes from --%v-- start", filepath.Join(fd, csvPost), size, r.RemoteAddr)

	mtxUnsubscribe.Lock()
	defer mtxUnsubscribe.Unlock()
	bts, err := os.ReadFile(filepath.Join(fd, csvPost))
	if err != nil {
		s := fmt.Sprintf("Could not open  %v. Email %v.", csvPost, adminEmail())
		log.Print(s)
		http.Error(w, s, 400)
		return
	}
	// this fails; server closes network connection, before bts is written to w
	//    fmt.Fprint(w, bts)
	// we need use 	io.Copy(out, in)
	http.ServeContent(w, r, csvPost, time.Now(), bytes.NewReader(bts))

	log.Printf("   %v - %v Bytes from --%v-- end", filepath.Join(fd, csvPost), size, r.RemoteAddr)

}
