package cloudio

import (
	"log"

	"cloud.google.com/go/storage"
)

// ReadDir preparation
var beforeList = func(as func(interface{}) bool) error {
	var q *storage.Query
	if as(&q) { // access storage.Query via q here.
		// log.Printf("beforeFunc(): delim - pref - versions: %v %v %#v", q.Delimiter, q.Prefix, q.Versions)
	} else {
		log.Printf("beforeFunc(): no response to %T", as)
	}
	return nil
}
