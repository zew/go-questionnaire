package stream

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type slowRequestForm struct {
	Delay   int `json:"delay"    form:"min=0,max=1100,maxlength='3',size='3',suffix='seconds before write'"`
	Repeats int `json:"repeats"  form:"min=1,max=1100,maxlength='3',size='3'"`
	Chunks  int `json:"chunks"   form:"min=0,max=1024,maxlength='3',size='3',suffix='kB'"`
	// NoHijack bool `json:"no_hijack" form:""`
}

func env(w io.Writer) {

	fmt.Fprintf(w, "<pre>\n")

	gaeApplication := os.Getenv("GAE_APPLICATION")
	fmt.Fprintf(w, "GAE_APPLICATION: %v\n", gaeApplication)

	instanceID := os.Getenv("GAE_INSTANCE")
	fmt.Fprintf(w, "GAE_INSTANCE: %v\n", instanceID)

	gaeRuntime := os.Getenv("GAE_RUNTIME")
	fmt.Fprintf(w, "GAE_RUNTIME: %v\n", gaeRuntime)

	gaeVersion := os.Getenv("GAE_VERSION")
	fmt.Fprintf(w, "GAE_VERSION: %v\n", gaeVersion)

	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "</pre>")
}

// HijackTest is the original example
// from the docs https://golang.org/pkg/net/http/#example_Hijacker
func HijackTest(w http.ResponseWriter, r *http.Request) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer conn.Close()
	fmt.Fprint(bufrw, "Now we're speaking raw TCP. Say hi: ")
	err = bufrw.Flush()
	if err != nil {
		log.Printf("flushing error: %v", err)
	}
	s, err := bufrw.ReadString('\n')
	if err != nil {
		log.Printf("error reading string: %v", err)
		return
	}
	fmt.Fprintf(bufrw, "You said: %q\nBye.\n", s)
	err = bufrw.Flush()
	if err != nil {
		log.Printf("flushing error: %v", err)
	}
}
