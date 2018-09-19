package main

/*

These tests replaces the calling of sys_test.go
in package systemtest.

Start the test from application root (i.e. /go-questionaire) with
	go test ./...  -v

For a particular package, start
	go test ./mypackage/... -test.v
	go test ./mypackage/...


## Coverage tests

Lots of hoops, if we want coverage tests, i.e. at gocover.io
An introduction is at www.elastic.co/blog/code-coverage-for-your-golang-system-tests

Note: This file is necessary for go-questionaire.test.exe binary to be generated.


1.) This leads to coverage: 0% of statements.
    go test ./... -coverprofile=coverage.log


2a.) We have to run the compiled test executable. Create it:
	go test -c -cover -covermode=count -coverpkg ./...

2b.) We could create a specific executable:
	go test -c -cover -covermode=count -coverpkg ./...  -o go-questionaire1.test.exe

2c.) We could restrict by sub package:
	go test -c -cover -covermode=count -coverpkg  "github.com/zew/go-questionaire/qst"
	go test -c -cover -covermode=count -coverpkg  "github.com/zew/go-questionaire/systemtest"

3.) Now we can collect coverage info.
	go-questionaire.test.exe  -test.v -test.coverprofile coverage.log

4.) Convert to HTML
	go tool cover -html=./coverage.log -o ./coverage.html


*/
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/systemtest"
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
			t.Fatalf("Questionaire validation caused error %v: %v", pth, err)
		}

		userName := "systemtest"
		surveyID := q.Survey.Type
		waveID := q.Survey.WaveID()

		t.Logf("\tquesionnaire type - survey-id: %v %v", surveyID, waveID)

		checkStr := fmt.Sprintf("%v-%v-%v-%v", surveyID, userName, waveID, lgn.Get().Salt)
		hsh := lgn.Md5Str([]byte(checkStr))
		loginURL := fmt.Sprintf("%v?u=%v&survey_id=%v&wave_id=%v&h=%v", cfg.PrefWTS(), userName, surveyID, waveID, hsh)
		t.Logf("\tLoginURL: %v", loginURL)

		removeSystemtestJSON(t)
		systemtest.SimulateLoad(t, q, loginURL, "0")

		removeSystemtestJSON(t)
		systemtest.SimulateLoad(t, q, loginURL, "1")
		// if surveyID == "peu2018" {
		// 	systemtest.SimulateLoad(t, q, loginURL)
		// }
	}

}

func removeSystemtestJSON(t *testing.T) {
	err := filepath.Walk(filepath.Join(".", "responses"), func(path string, f os.FileInfo, err error) error {
		base := filepath.Base(path)
		if base == "systemtest.json" {
			log.Printf("Removing %v", path)
			os.Remove(path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}
