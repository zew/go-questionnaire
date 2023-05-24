package main

/*

These tests replaces the calling of sys_test.go
in package systemtest.

## Run tests

  Start the test from application root (i.e. /go-questionnaire)
	go test -v         .
	go test -v         ./...
	go test -v  -race  ./...


  For a particular package, start
	go test -v ./mypackage/... -test.v
	go test -v ./mypackage/...

  i.e. for systemtest
	go test -v ./cmd/server/...


## Coverage tests

<www.elastic.co/blog/code-coverage-for-your-golang-system-tests>
<https://www.alexedwards.net/blog/an-overview-of-go-tooling>

-covermode=count    => coverage records exact number, each statement is executed during tests.
-covermode=atomic   => same as avove, when t.Parallel() is used in any test

1.) Direct

  All main and all packages
	go test -v  ./...  -coverprofile=tmp-coverage.log -covermode=atomic
    go tool  cover  -html=tmp-coverage.log

  Note: "0% of statements" is for the last package;
	to see the coverage for main package,
	   scroll up,
	or omit -v,
	or replace ./... with .

  Restrict to pkg
  	go test  -v  ./...  -coverprofile=tmp-coverage.log -covermode=atomic  github.com/zew/go-questionnaire
  	go test  -v  ./...  -coverprofile=tmp-coverage.log -covermode=atomic  github.com/zew/util
	go tool  cover  -html=./tmp-coverage.log -o ./tmp-coverage.html


2.) With test executable

  a.) Create a compiled test executable 'go-questionnaire.exe'
	go test -c -cover -covermode=count -coverpkg ./...

  b.) We could create a _specific_ test executable:
	go test -c -cover -covermode=count -coverpkg ./...  -o go-questionnaire.test.exe

  c.) We could restrict by sub package:
	go test -c -cover -covermode=count -coverpkg  "github.com/zew/go-questionnaire/pkg/qst"
	go test -c -cover -covermode=count -coverpkg  "github.com/zew/go-questionnaire/systemtest"

3.) Now we can collect coverage info.
	go-questionnaire.test.exe  -test.v  -test.coverprofile tmp-coverage.log

4.) Convert to HTML
	go tool cover -html=./tmp-coverage.log -o ./tmp-coverage.html


*/
import (
	"log"
	"os"
	"path"
	"strings"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/tests/systemtest"
)

// Coverage is only started, when the test binary is run.
// Must call main() - otherwise coverage is not counted.
//
// We dont want to restrict execution by flag systemTest
// since the test should also run on gocover.io
func TestSystem(t *testing.T) {

	os.Setenv("GO_TEST_MODE", "true")
	defer os.Setenv("GO_TEST_MODE", "false")

	// for ./app-bucket and ./static relative paths to work,
	// working dir is ./cmd/server - changing to .
	os.Chdir(path.Join("..", ".."))

	go func() {
		main()
	}()
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second) // wait for the server to come up
		t.Logf("Waiting for the server to come up ... %v", i)
	}

	log.SetFlags(log.Lshortfile)

	tplDir := "responses"
	files, err := cloudio.ReadDir(path.Join(".", tplDir) + "/")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("iterating *.json files in ... /%v ", tplDir)
	for _, f := range *files {

		if f.IsDir {
			t.Logf("Skipping directory %v", f.Key)
			continue
		}
		if path.Ext(f.Key) != ".json" {
			t.Logf("Skipping non json file %v", f.Key)
			continue
		}

		// if strings.HasSuffix(f.Key, "pds-2023-01.json") {
		// 	t.Logf("Skipping file %v", f.Key)
		// 	continue
		// }
		if strings.HasSuffix(f.Key, "pds-2023-01-full-dynamic-content.json") {
			t.Logf("Skipping file %v", f.Key)
			continue
		}
		if strings.HasSuffix(f.Key, "pds-2023-04-full-dynamic-content.json") {
			t.Logf("Skipping file %v", f.Key)
			continue
		}

		t.Logf("\n\n\n")
		if false {
			t.Logf("\n\n\n%v", f.Key)
			continue
		}

		// t.Logf("Found questionnaire template %v", f.Key)
		qTpl, err := qst.Load1(f.Key) // qTpl merely stores some settings for later function calls to read
		if err != nil {
			t.Fatalf("Could not load %v: %v", f.Key, err)
		}
		log.Printf("No previous user questionnaire file %v found. Using base file.", f.Key)
		err = qTpl.Validate()
		if err != nil {
			t.Fatalf("Questionnaire validation caused error %v: %v", f.Key, err)
		}

		t.Logf("\tquestionnaire type - survey-id: %v - %v", qTpl.Survey.String(), f.Key)

		userName := "systemtest"
		surveyID := qTpl.Survey.Type
		waveID := qTpl.Survey.WaveID()

		loginURL := lgn.LoginURL(userName, surveyID, waveID, "")
		loginURL += "&override_closure=true"
		t.Logf("\tLoginURL: %v", loginURL)

		if false {
			// Deadline exceeded?
			if time.Now().After(qTpl.Survey.Deadline) {
				s := cfg.Get().Mp["deadline_exceeded"].All(qTpl.Survey.Deadline.Format("02.01.2006 15:04"))
				if len(s) > 100 {
					s = s[:100]
				}
				t.Logf("%v", s)
				t.Logf("Cannot test questionnaire that which are already closed: %v\n\n", qTpl.Survey.Type)
				continue
			}
		}

		if surveyID == "peu2018-or-special-survey-name" {
			//
		}

		// if surveyID != "pds" {
		// 	continue
		// }

		// if q.Survey.String() != "fmt-2021-08" {
		// 	continue
		// }

		// if surveyID != "fmt" {
		// 	continue
		// }

		// call with last arg "0" was for http user agend based differentiation of mobile or desktop rendering
		systemtest.SimulateLoad(t, qTpl, loginURL, "1")

	}

}
