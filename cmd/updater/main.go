// Package updater makes a change to all questionaires in a given directory;
// can be applied to single origin json - as well as to filled out json files.
//     updater.exe -dir ../app-bucket/responses/mul.json
//     updater.exe -dir ../app-bucket/responses/mul/2019-02
//     updater.exe -dir ../app-bucket/responses/mul/2019-02/23121.json
//
// The saving to app-bucked via cloudio might fail
// and the changes might be saved to app-dir/app-bucket/...
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
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
			Long:       "directory",
			Short:      "dir",
			DefaultVal: "../../app-bucket/responses/downloaded/fmt/2021-04/11499.json",
			Desc:       "filename - or directory or to iterate",
		},
	)
	fl.Gen()
	dirSrc := fl.ByKey("dir").Val

	//
	isFile := false
	files, err := ioutil.ReadDir(dirSrc)
	if err != nil {
		log.Printf("Opening as directory failed: %v", err)
		f, err := os.OpenFile(dirSrc, 0, 0777)
		if err != nil {
			log.Printf("Opening as file      failed: %v", err)
			return
		}
		fi, err := f.Stat()
		if err != nil {
			log.Printf("Error obtaining file info: %v", err)
			return
		}
		log.Printf("Opening as file succeeded: %v", dirSrc)
		isFile = true
		files = append(files, fi)
	}

	for idx, f := range files {
		log.Printf("found file %3v: %v", idx, f.Name())
	}

	//
	cntrChanged := 0
	for i, f := range files {

		pSrc := path.Join(dirSrc, f.Name())
		if isFile {
			pSrc = filepath.Join(dirSrc)
		}
		log.Printf("%3v: opening file  %v", i, pSrc)

		pDst := path.Join(path.Dir(dirSrc), "updated", f.Name())

		bts, err := ioutil.ReadFile(pSrc)
		if err != nil {
			log.Printf("%3v: Error reading file %v: %v", i, pSrc, err)
			return
		}

		q := qst.QuestionnaireT{}
		err = json.Unmarshal(bts, &q)
		if err != nil {
			log.Printf("%3v: Error unmarshalling file %v: %v", i, pSrc, err)
			return
		}
		log.Printf("%3v: questionnaire %v - unmarshalled - %10.3f MB", i, pSrc, float64(len(bts)/(1<<10))/(1<<10))

		//
		//
		//
		// now we might perform various changes to the questionnaire
		// then saving the questionnaire; checksum via q.Save()

		var t1 time.Time
		q.ClosingTime = t1
		err = q.Save1(pDst)
		if err != nil {
			log.Printf("%3v: Error saving %v: %v", i, pSrc, err)
		}
		cntrChanged++

		if false {
			if q.Variations > 0 {
				log.Printf("%3v: questionnaire %v - correction needed %v", i, pSrc, q.Survey.Deadline)
				// q.Survey.Deadline = tInstead
				q.Variations = 0
				err := q.Save1(pDst)
				if err != nil {
					log.Printf("%3v: Error saving %v: %v", i, pSrc, err)
				}
				cntrChanged++
				log.Printf("%3v: questionnaire %v saved", i, pSrc)
			}

		}

		if false {
			search := q.Pages[1].Groups[8].Inputs[0].Label["fr"]
			old := "con-trainte"
			new := "contrainte"
			if strings.Contains(search, old) {
				replaced := strings.Replace(search, old, new, -1)
				q.Pages[1].Groups[8].Inputs[0].Label["fr"] = replaced
				err := q.Save1(pDst)
				if err != nil {
					log.Printf("%3v: Error saving %v: %v", i, pSrc, err)
				}
				cntrChanged++
				log.Printf("%3v: questionnaire %v - %v corrected to %v", i, pSrc, old, new)
			} else {
				log.Printf("%3v: questionnaire %v - correction not needed %v", i, pSrc, search)
			}
		}

	}
	log.Printf("================")
	log.Printf("Finish - %v changes", cntrChanged)

}
