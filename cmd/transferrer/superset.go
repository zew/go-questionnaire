package main

import (
	"io"
	"log"
	"os"
)

var lgS = log.New(os.Stdout, "", log.Llongfile)

func init() {
	lgS.SetOutput(io.Discard) // enable/disable logging in this file
}

// Superset returns the union of all keys;
// the sort order is kept;
// new or sparse keys are inserted according to their last known predecessor;
// otherwise at the end.
// Purpose: Create CSV columns in intuitive order.
func Superset(keys [][]string) (superset []string) {
	if len(keys) > 0 {
		superset = make([]string, 0, len(keys[0]))
	}

	has := map[string]interface{}{}
	for i := 0; i < len(keys); i++ {
		for i2 := 0; i2 < len(keys[i]); i2++ {
			k := keys[i][i2]
			if _, ok := has[k]; ok {
				continue
			}
			lgS.Printf("%v of %v not found in %v", k, keys[i], superset)

			// Insertion position for non-existing key k?
			// We take the predecessors of k in descending order.
			// For each predecessor, we search in superset for
			// a similar element.
			insAft := len(superset) - 1 // default: Insert at the end
			for i3 := i2 - 1; i3 >= 0; i3-- {
				similar := keys[i][i3]
				lgS.Printf("\tTesting %v", similar)
				if _, ok := has[similar]; ok {
					for i4, val := range superset {
						if val == similar {
							insAft = i4
						}
					}
					lgS.Printf("\tLast matching element of %v with %v is %v at %v",
						keys[i], superset, similar, insAft)
					break
				}
			}
			temp := make([]string, 0, len(superset)+1)
			temp = append(temp, superset[:insAft+1]...)
			temp = append(temp, k)
			temp = append(temp, superset[insAft+1:]...)
			superset = temp
			has[k] = nil
			lgS.Printf("New superset is %+v", superset)
		}
	}
	return
}
