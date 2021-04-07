// Package stream contains convenience funcs to hijack a connection
// and flush each write to the client; test handlers simulate long requests;
// some information for app engine deployments is provided
package stream

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

// WriterFlushable is a http response writer
// for usage just like w http.ResponseWriter;
// use makeFlushable to convert from ordinary http.ResponseWriter;
// see NewFlushable() for more details
type WriterFlushable struct {
	bufWr       bufio.ReadWriter
	respWr      http.ResponseWriter
	NoAutoFlush bool
}

// WriteHeader implements the http.ResponseWriter interface
func (w WriterFlushable) WriteHeader(statuscode int) {
	log.Printf("headers in hijacked reponses must be written directly via bufio.ReadWriter")
	w.respWr.WriteHeader(statuscode)
}

// Header implements the http.ResponseWriter interface
func (w WriterFlushable) Header() http.Header {
	log.Printf("headers in hijacked reponses must be written directly via bufio.ReadWriter. This returns an empty map.")
	return w.respWr.Header()
}

// Header implements the http.ResponseWriter interface
func (w WriterFlushable) Write(bts []byte) (int, error) {
	if !w.NoAutoFlush {
		err := w.bufWr.Flush()
		if err != nil {
			log.Printf("WriterFlushable().Write()...Flush(): %v", err)
		}
	}
	return w.bufWr.Write(bts)
}

// Flush is integrated into Write;
// use this only with NoAutoFlush.
func (w WriterFlushable) Flush() {
	err := w.bufWr.Flush()
	if err != nil {
		log.Printf("WriterFlushable()...Flush(): %v", err)
	}
}

/*NewFlushable returns a flushable *auto flushing* response writer;
for this, it must hijack the request.
It returns a function that must be called at the end of the response,
to close the hijacked connection, signalling end of the response to the http client.
On errors, the old response writer is returned

Ordinary flushing would go like this:

	w1, ok := w.(http.Flusher)
	if ok {
		w1.Flush()
		log.Printf("\t\tflushed")
	}

Ordinary flushing is prevented from any middleware writer,
which might be buffered and or does not implement the flusher interface.
But even then, the flushing would be required inside of io.Copy after each
chunk;
hence this response writer implements *auto flushing*;
incurring the price of hijacking the request and writing
headers directly
	fmt.Fprintf(w, "HTTP/1.1 200 OK\n")

Second parameter is a func to call at the end of the request
*/
func NewFlushable(w http.ResponseWriter) (http.ResponseWriter, func(), error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		return w, func() {}, fmt.Errorf("webserver doesn't support hijacking")
	}
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		return w, func() {}, fmt.Errorf("could not hijack: %v", err)
	}
	f := func() {
		err := bufrw.Flush()
		if err != nil {
			log.Printf("WriterFlushable().'Close()'...Flush(): %v", err)
		}
		err = conn.Close()
		if err != nil {
			log.Printf("WriterFlushable().conn.Close(): %v", err)
		}
		log.Printf("Flushed and closed")
	}
	wReturn := WriterFlushable{bufWr: *bufrw, respWr: w}
	return wReturn, f, nil
}
