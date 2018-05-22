package qst

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"time"
)

var questPath = "."

// Load loads a questionaire from a JSON file.
func Load(fn string) (*QuestionaireT, error) {
	q := QuestionaireT{}

	bts, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Could not read file %v : %v", fn, err)
		return &q, err
	}

	err = json.Unmarshal(bts, &q)
	if err != nil {
		log.Fatalf("Unmarshal %v: %v", fn, err)
		return &q, err
	}

	return &q, nil
}

// Save a questionaire to JSON
func (q *QuestionaireT) Save(fn ...string) error {

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(q, firstColLeftMostPrefix, "\t")
	if err != nil {
		return err
	}

	saveDir := path.Dir(questPath)
	err = os.Chmod(saveDir, 0755)
	if err != nil {
		return err
	}

	questFile := path.Base(questPath)
	if len(fn) > 0 {
		questFile = fn[0]
	}

	pthOld := path.Join(saveDir, questFile)

	keepBackup := false
	if keepBackup {
		fileBackup := strings.Replace(questFile, ".json", fmt.Sprintf("_%v.json", time.Now().Unix()), 1)
		pthBackup := path.Join(saveDir, fileBackup)
		if questFile != "questionaire-example.json" {
			err = os.Rename(pthOld, pthBackup)
			if err != nil {
				return err
			}
		}
	}

	err = ioutil.WriteFile(pthOld, byts, 0644)
	if err != nil {
		return err
	}
	log.Printf("Saved questionaire file to %v", pthOld)
	return nil
}
