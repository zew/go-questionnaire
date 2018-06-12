// Package muxwrap inserts some common logic,
// before calling the actual handler func.
package muxwrap

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/logx"
	"github.com/zew/util"
)

// The wrapper has two characteristics
// 1.) It implements http.Handler
// 2.) It stores the previous mux/handler.
// http ListenAndServe and ListenAndServeTLS each dont
// need a mux; they only need http handler.
type handlerWrapper struct {
	h http.Handler
}

// NewHandlerMiddleware returns a new http handlerfunc.
func NewHandlerMiddleware(innerHandler http.Handler) http.Handler {
	m := &handlerWrapper{
		h: innerHandler,
	}
	return m
}

// Implementing http.Handler interface
func (m *handlerWrapper) ServeHTTP(w http.ResponseWriter, rNew *http.Request) {

	// Global logging stuff
	shortened := fmt.Sprintf("%v?%v", rNew.URL.Path, rNew.URL.RawQuery)
	// lg.Printf("------------------------------------------")
	if !strings.HasSuffix(rNew.URL.Path, "favicon.ico") {
		lg.Printf("%-60v | referr %v", shortened, util.UrlBeautify(rNew.Referer()))
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

	//
	// Access rights

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
				microErrorPage(rNew, w, msg1)
			}
		}()
		m.h.ServeHTTP(w, rNew)
	}()

}

//
// Alternative way to create the same middle ware would be:
func (m *handlerWrapper) Use(next http.Handler, anotherParam int) http.Handler {
	// Possible stuff outside the closure
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), ctxKey("anotherParam"), anotherParam)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type ctxKey string
