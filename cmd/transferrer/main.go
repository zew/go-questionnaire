// see package tf
package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/zew/go-questionnaire/pkg/tf"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	//
	// we must change to main app dir,
	// so that referring to ./app-bucket works
	err := os.Chdir("../..")
	if err != nil {
		log.Fatalf("Error - cannot 'cd' to main app dir: %v", err)
	}

	cfgRem := tf.ConfigsThree()

	qs, err := tf.RetrieveFromRemote(cfgRem)
	if err != nil {
		log.Printf("error retrieving questionnaires from remote: %v", err)
		return
	}

	csvPath, err := tf.ProcessQs(cfgRem, qs)
	if err != nil {
		log.Printf("error processing questionnaires from remote: %v", err)
		return
	}
	log.Printf("CSV file saved under: %v", csvPath)

}
