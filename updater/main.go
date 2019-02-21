// Package updater makes a change to all questionaires in a given directory;
// can be applied to single origin json - as well as to filled out json files.
//     updater.exe -dir ../responses/mul.json
//     updater.exe -dir ../responses/mul/2019-02
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/qst"
	"github.com/zew/util"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	// The actual config for *this* app:
	fl := util.NewFlags()
	fl.Add(
		util.FlagT{
			Long:  "directory",
			Short: "dir",
			// DefaultVal: "../responses/mul.json",
			DefaultVal: "../responses/mul/2019-02/",
			Desc:       "filename - or directory or to iterate",
		},
	)
	fl.Gen()
	dir := fl.ByKey("dir").Val

	//
	isFile := false
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Printf("Opening as directory failed: %v", err)
		f, err := os.OpenFile(dir, 0, 0777)
		if err != nil {
			log.Printf("Opening as file      failed: %v", err)
			return
		}
		fi, err := f.Stat()
		if err != nil {
			log.Printf("Error obtaining file info: %v", err)
			return
		}
		log.Printf("Opening as file succeeded: %v", dir)
		isFile = true
		files = append(files, fi)
	}

	//
	cntrChanged := 0
	for i, f := range files {

		pth := filepath.Join(dir, f.Name())
		if isFile {
			pth = filepath.Join(dir)
		}
		log.Printf("%3v: opening file  %v", i, pth)

		bts, err := ioutil.ReadFile(pth)
		if err != nil {
			log.Printf("%3v: Error reading file %v: %v", i, pth, err)
			return
		}

		q := qst.QuestionnaireT{}
		err = json.Unmarshal(bts, &q)
		if err != nil {
			log.Printf("%3v: Error unmarshalling file %v: %v", i, pth, err)
			return
		}
		log.Printf("%3v: questionnaire %v - unmarshalled - %10.3f MB", i, pth, float64(len(bts)/(1<<10))/(1<<10))

		// changing and saving the questionnaire; checksum via q.Save()
		//
		// loc := time.FixedZone("UTC", 1*3600)
		// tToBeCorrected := time.Date(2018, 10, 31, 24, 59, 0, 0, loc)  // "deadline": "2018-10-31T23:59:00Z"
		// if q.Survey.Deadline.Equal(tToBeCorrected) {
		if false && q.Variations > 0 {
			log.Printf("%3v: questionnaire %v - correction needed %v", i, pth, q.Survey.Deadline)
			// q.Survey.Deadline = tInstead
			q.Variations = 0
			err := q.Save1(pth)
			if err != nil {
				log.Printf("%3v: Error saving %v: %v", i, pth, err)
			}
			cntrChanged++
			log.Printf("%3v: questionnaire %v saved", i, pth)
		}

		search := q.Pages[1].Groups[8].Desc["fr"]
		old := "con-trainte"
		new := "contrainte"
		if strings.Contains(search, old) {
			replaced := strings.Replace(search, old, new, -1)
			q.Pages[1].Groups[8].Desc["fr"] = replaced
			err := q.Save1(pth)
			if err != nil {
				log.Printf("%3v: Error saving %v: %v", i, pth, err)
			}
			cntrChanged++
			log.Printf("%3v: questionnaire %v - %v corrected to %v", i, pth, old, new)
		} else {
			log.Printf("%3v: questionnaire %v - correction not needed %v", i, pth, search)
		}

	}
	log.Printf("================")
	log.Printf("Finish - %v changes", cntrChanged)

}
