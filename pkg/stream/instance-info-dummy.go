// +build !appengine
//
// opposite would be
// +build appengie

package stream

import (
	"fmt"
	"net/http"
)

// InstanceInfoXXXX - the previous technique of having parallel
// implementations conditionally compiled
func InstanceInfoXXXX(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not on appengine")
}
