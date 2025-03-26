package cloudio

import (
	"log"

	"cloud.google.com/go/storage"
	"gocloud.dev/blob"
)

func getBlobOpts(pfx string) *blob.ListOptions {
	return &blob.ListOptions{
		Delimiter:  "/",
		Prefix:     pfx,
		BeforeList: beforeList,
	}
}

// 2023-12: I tried to disable this code, to remove the huge package cloud.google.com/go/storage
//    but it remains as indirect dependeny

// ReadDir preparation
var beforeList = func(as func(interface{}) bool) error {

	// pbu 2025-03 - temporary disabled for go get -u ./... to complete
	var q *storage.Query
	if as(&q) { // access storage.Query via q here.
		// log.Printf("beforeFunc(): delim - pref - versions: %v %v %#v", q.Delimiter, q.Prefix, q.Versions)
	} else {
		log.Printf("beforeFunc(): no response to %T", as)
	}

	return nil
}
