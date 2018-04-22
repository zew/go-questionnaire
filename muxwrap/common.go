package muxwrap

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/zew/questionaire/sessx"
	"github.com/zew/util"
)

var lg = log.New(os.Stdout, "middleware ", log.Lshortfile|log.Ltime) // Logger with special prefix

//
// We call this once at request start
// to persist params into the session
func paramPersister(r *http.Request, sess *sessx.TSess) {

	// Excluding static files
	// from being middlewared.
	if util.StaticExtension(r) {
		return
	}

	keysToPersist := []string{"session-test-key", "request-test-key"}
	for _, key := range keysToPersist {
		if reqVal, ok := sess.RequestParamIsSet(key); ok {
			lg.Printf("\tsess key SET  %17v is %-16v", key, reqVal)
			sess.PutString(key, reqVal)
		}
	}

}

func microErrorPage(r *http.Request, w http.ResponseWriter, msg string) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	wr := func(s string, args ...interface{}) {
		_, err := w.Write([]byte(fmt.Sprintf(s, args...)))
		if err != nil {
			util.BubbleUp(err)
		}
	}

	wr("<div style='white-space: pre-wrap;'>\n")
	wr(strings.TrimSpace(msg))
	wr("\n")
	wr("</div>\n")
}
