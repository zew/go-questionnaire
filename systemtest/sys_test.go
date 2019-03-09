package systemtest

import (
	"log"
	"os"
	"testing"

	"github.com/zew/go-questionnaire/qst"
)

// This is obsolete - see ../main_test.go for details
//
// Run with
//   go test ./systemtest/... -test.v
// or
//   go test ./systemtest/...
func Test_1(t *testing.T) {
	// This is run from main_test... on dir higher
	if true {
		return
	}

	//
	if os.Getenv("DATASOURCE1") == "travis" {
		log.Printf("On travis, tests are executed in the app dir main_test.go")
	} else {
		// log.Printf("On gocover.io, tests in the app dir are ignored")
		StartTestServer(t, true)
		q := &qst.QuestionnaireT{}
		SimulateLoad(t, q, "loginURL", "0")
	}
}
