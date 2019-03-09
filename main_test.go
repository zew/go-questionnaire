package main

/*

These tests replaces the calling of sys_test.go
in package systemtest.

Start the test from application root (i.e. /go-questionnaire) with
	go test ./...    -v
	go test .        -v

For a particular package, start
	go test ./mypackage/... -test.v
	go test ./mypackage/...


## Coverage tests

Lots of hoops, if we want coverage tests, i.e. at gocover.io
An introduction is at www.elastic.co/blog/code-coverage-for-your-golang-system-tests

Note: This file is necessary for go-questionnaire.test.exe binary to be generated.


1.) This leads to coverage: 0% of statements.
    go test ./... -coverprofile=coverage.log


2a.) We have to run the compiled test executable. Create it:
	go test -c -cover -covermode=count -coverpkg ./...

2b.) We could create a specific executable:
	go test -c -cover -covermode=count -coverpkg ./...  -o go-questionnaire1.test.exe

2c.) We could restrict by sub package:
	go test -c -cover -covermode=count -coverpkg  "github.com/zew/go-questionnaire/qst"
	go test -c -cover -covermode=count -coverpkg  "github.com/zew/go-questionnaire/systemtest"

3.) Now we can collect coverage info.
	go-questionnaire.test.exe  -test.v -test.coverprofile coverage.log

4.) Convert to HTML
	go tool cover -html=./coverage.log -o ./coverage.html


*/
import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/go-questionnaire/systemtest"
)

// Coverage is only started, when the test binary is run.
// Must call main() - otherwise coverage is not counted.
//
// We dont want to restrict execution by flag systemTest
// since the test should also run on gocover.io
func TestSystem(t *testing.T) {
	os.Setenv("GO_TEST_MODE", "true")
	defer os.Setenv("GO_TEST_MODE", "false")
	go func() {
		main()
	}()
	for i := 0; i < 6; i++ {
		time.Sleep(time.Second) // wait for the server to come up
		t.Logf("Waiting for the server to come up ... %v", i)
	}

	tplDir := "responses"

	files, err := ioutil.ReadDir(filepath.Join(".", tplDir))
	if err != nil {
		t.Fatal(err)
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		t.Logf("Found quesionnaire template %v", f.Name())

		pth := filepath.Join(".", tplDir, f.Name())
		q, err := qst.Load1(pth)
		if err != nil {
			t.Fatalf("Could not load %v: %v", pth, err)
		}
		err = q.Validate()
		if err != nil {
			t.Fatalf("Questionnaire validation caused error %v: %v", pth, err)
		}

		userName := "systemtest"
		surveyID := q.Survey.Type
		waveID := q.Survey.WaveID()
		t.Logf("\tquesionnaire type - survey-id: %v %v", surveyID, waveID)

		loginURL := lgn.LoginURL(userName, surveyID, waveID)
		t.Logf("\tLoginURL: %v", loginURL)

		// Deadline exceeded?
		if time.Now().After(q.Survey.Deadline) {
			s := cfg.Get().Mp["deadline_exceeded"].All(q.Survey.Deadline.Format("02.01.2006 15:04"))
			if len(s) > 100 {
				s = s[:100]
			}
			t.Logf("%v", s)
			t.Logf("Cannot test questionnaire that which are already closed: %v\n\n", q.Survey.Type)
			continue
		}

		systemtest.SimulateLoad(t, q, loginURL, "0")

		systemtest.SimulateLoad(t, q, loginURL, "1")
		// if surveyID == "peu2018" {
		// 	systemtest.SimulateLoad(t, q, loginURL)
		// }
	}

}
