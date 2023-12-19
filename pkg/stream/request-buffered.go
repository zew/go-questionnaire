package stream

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/pbberlin/dbg"
	"github.com/pbberlin/struc2frm"
)

// SlowBuffered - standard request with timeouts
func SlowBuffered(w http.ResponseWriter, r *http.Request) {

	// convenience
	logAndShow := func(f string, intf ...interface{}) {
		fmt.Fprintf(w, f+"<br>\n", intf...)
		log.Printf("\t"+f+"\n", intf...)
	}

	// preprocessing request form
	err := r.ParseForm()
	if err != nil {
		logAndShow("cannot parse form: %v<br>\n <pre>%v</pre>", err, dbg.Dump2String(r.Form))
		return
	}
	dec := form.NewDecoder()
	dec.SetTagName("json")
	frm := &slowRequestForm{Repeats: 4}
	err = dec.Decode(frm, r.Form)
	if err != nil {
		logAndShow("cannot decode form: %v<br>\n <pre>%v</pre>", err, dbg.Dump2String(r.Form))
		return
	}

	for i := 0; i < frm.Delay; i++ {
		log.Printf("preliminary delay ... %2d secs", i) // dont write any response bytes yet
		time.Sleep(time.Second)
	}
	log.Printf("preliminary delay ... %2d secs", frm.Delay)

	//
	// first response writes
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	//
	fmt.Fprintf(w, "<h3>Slow requests - fully buffered</h3>")
	fmt.Fprint(w, struc2frm.New().Form(*frm))

	env(w)

	if len(r.Form) == 0 {
		log.Printf("No parameters set; return")
		return
	}

	// slow body
	intrval := 1
	for i := 0; i < frm.Repeats; i++ {
		time.Sleep(time.Duration(intrval) * time.Second)
		logAndShow("response body writing ... %2d secs", (i+1)*intrval)
		if frm.Chunks > 0 {
			fmt.Fprint(w, strings.Repeat(" -", 512*frm.Chunks))
			fmt.Fprint(w, "<br>\n")
		}
		flusher, ok := w.(http.Flusher)
		if ok {
			flusher.Flush()
			fmt.Fprint(w, "flushed<br>\n")
		} else {
			fmt.Fprint(w, "NO flusher<br>\n")
		}
	}

	logAndShow("end of slow request test\n\n")

}
