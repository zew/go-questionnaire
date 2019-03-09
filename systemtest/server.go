package systemtest

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/bootstrap"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/handlers"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/muxwrap"
	"github.com/zew/go-questionnaire/sessx"
)

// StartTestServer starts a server almost like main().
//
// Coverage tests must be run from app root.
// and they must call main().
//
// Thus, if we use the actual app server,
// then why do we need a test server?
//
// See app root main_test for more details.
func StartTestServer(t *testing.T, doChDirUp bool) {

	// For database files, static files and templates relative paths to work,
	// as if running from main app dir:
	if doChDirUp {
		os.Chdir("..")
	}

	wd, _ := os.Getwd()
	t.Logf("test directory one up: %v ; should be app main dir", wd)

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	os.Setenv("GO_TEST_MODE", "true")
	defer os.Setenv("GO_TEST_MODE", "false")
	bootstrap.Config()

	//
	// Start the server
	{
		mux1 := http.NewServeMux()
		mux1.HandleFunc(cfg.Pref("/login-primitive"), lgn.LoginPrimitiveH)
		mux1.HandleFunc(cfg.Pref("/"), handlers.MainH)
		mux1.HandleFunc(cfg.PrefWTS("/"), handlers.MainH)
		mux2 := muxwrap.NewHandlerMiddleware(mux1)
		sessx.Mgr().Lifetime = 2 * time.Hour // default is 24 hours
		sessx.Mgr().Cookie.Persist = false
		mux3 := sessx.Mgr().LoadAndSave(mux2)

		IPPort := fmt.Sprintf("%v:%v", cfg.Get().BindHost, cfg.Get().BindSocket)
		t.Logf("starting http server at %v ...", IPPort)

		chSuccess := make(chan error)
		bootFunc := func(ch chan error) {
			err := http.ListenAndServe(IPPort, mux3)
			ch <- err
		}

		go bootFunc(chSuccess)

		select {
		case errBoot := <-chSuccess:
			if errBoot != nil {
				t.Fatalf("\nCould not start test server. \nLive system running? \nError %v", errBoot)
				return
			}
		case <-time.After(1100 * time.Millisecond):
			t.Logf("Test server came up without error")
		}
		// time.Sleep(1100 * time.Millisecond) // wait for application to come up
	}

}
