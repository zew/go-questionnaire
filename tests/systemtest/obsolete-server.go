package systemtest

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"testing"
	"time"

	"github.com/zew/go-questionnaire/pkg/bootstrap"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/handlers"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/sessx"
	"github.com/zew/go-questionnaire/pkg/wrap"
)

// This is obsolete - see ../main_test.go for details
//
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

	// for ./app-bucket and ./static relative paths to work,
	// working dir is ./cmd/server - changing to .
	if doChDirUp {
		os.Chdir(path.Join("..", ".."))
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
		mux1.HandleFunc(cfg.PrefTS("/"), handlers.MainH)
		mux2 := wrap.LogAndRecover(mux1)

		sessx.Mgr().Lifetime = time.Duration(cfg.Get().SessionTimeout) * time.Hour // default is 24 hours
		// sessx.Mgr().Secure(true)            // true breaks session persistence in excel-db - but not in go-countdown - it leads to sesson breakdown on iphone safari mobile, maybe because appengine is http with TLS outside
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
