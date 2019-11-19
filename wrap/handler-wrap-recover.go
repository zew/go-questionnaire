// Package wrap inserts some common logic,
// before calling the actual handler func.
package wrap

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/logx"
	"github.com/zew/util"
)

// The wrapper has two characteristics
// 1.) It implements http.Handler
// 2.) It *could* store the previous mux/handler.
// http ListenAndServe and ListenAndServeTLS each dont
// need a mux; they only need http handler.
type logAndRecover struct {
	h http.Handler
}

// LogAndRecover takes a single handler
// and returns a new wrapped http handler.
//
// Magically, we can pass in an *entire mux* -
// having *all* its routes wrapped into recover
func LogAndRecover(innerHandler http.Handler) http.Handler {
	return &logAndRecover{
		h: innerHandler,
	}
}

//
// 1. Logging each request
// 2. Recover from panic
// 3. Execute inner handler
// And implementing http.Handler interface
func (m *logAndRecover) ServeHTTP(w http.ResponseWriter, rNew *http.Request) {

	// Global logging stuff
	shortened := fmt.Sprintf("%v?%v", rNew.URL.Path, rNew.URL.RawQuery)
	// lg.Printf("------------------------------------------")
	if !strings.HasSuffix(rNew.URL.Path, "favicon.ico") {
		if !util.StaticExtension(rNew) {
			lg.Printf("%-60v | referr %v", shortened, util.UrlBeautify(rNew.Referer()))
		}
	}

	// Limit POST request body size.
	// Beware: Restricts file upload size.
	// There is no default restriction - only 10 MB *memory* limit - rest goes to hard disk - stackoverflow.com/questions/28282370/
	maxPostSize := cfg.Get().MaxPostSize
	if maxPostSize > 0 {
		rNew.Body = http.MaxBytesReader(w, rNew.Body, maxPostSize)
	}

	// Skip remaining stuff for static files.
	// Maybe better dynamically read all dirs under ./static
	// and trap allow all requests to those dirs here.
	if util.StaticExtension(rNew) {
		m.h.ServeHTTP(w, rNew)
		return
	}

	// Global session stuff
	sess := sessx.New(w, rNew)
	sess.PutString("session-test-key", "session-test-value")
	paramPersister(rNew, &sess) // before creating first state

	/*
		Filippo Valsorda
		blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/

		Additional constraint on each request;
		more restrictive than ReadTimeout/ReadHeaderTimeout + WriteTimeout
		Similar to http.TimeoutHandler()

		Reason: We have to globally increase WriteTimeout
		due to transferrer-endpoint and download requests;
		here we restrain normal requests again
	*/
	if cfg.Get().TimeOutUsual > 0 {
		perReqTimeout := time.Duration(cfg.Get().TimeOutUsual) * time.Second // reduce from 60 secs to 10 secs
		apply := true
		for _, s := range cfg.Get().TimeOutExceptions {
			if strings.Contains(rNew.URL.Path, s) {
				apply = false
			}
		}
		if apply {
			log.Printf("DE-creasing req timeout to %v for %v", perReqTimeout, rNew.URL.Path)
			ctx, perReqCancel := context.WithTimeout(rNew.Context(), time.Duration(perReqTimeout*time.Second))
			defer perReqCancel()
			rNew = rNew.WithContext(ctx)
		}
	}

	//
	// Wrapping the handler into a "global panic catcher"
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				msg1 := fmt.Sprintf("Panic in http handler %v:\n%v\n", shortened, rec)
				lg.Print(msg1)
				logx.SPrintStackTrace(1, 10)
				msg1 += "<span style='font-size:80%;'>If panic not triggered by logx.Fatal:</span>\n"
				msg1 += "Direct stacktrace."
				msg1 += logx.SPrintStackTrace(1, 10)
				microErrorPage(w, rNew, msg1)
			}
		}()
		m.h.ServeHTTP(w, rNew)
	}()
}

//
// UseLogRecover is an alternative way to create the same middle ware
func UseLogRecover(inner http.Handler, aParam int) http.Handler {
	// Possible stuff outside the closure
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), ctxKey("param"), aParam)
		inner.ServeHTTP(w, r.WithContext(ctx))
	})
}

type ctxKey string
