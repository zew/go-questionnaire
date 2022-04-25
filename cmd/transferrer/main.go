// see package tf
package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/qst"
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

	{
		qs, cfgRem, err := tf.RetrieveFromServer()
		if err != nil {
			log.Printf("Error retrieving questionnaires from remote: %v", err)
			return
		}

		dirFull := path.Join(cfgRem.DownloadDir, cfgRem.SurveyType, cfgRem.WaveID)
		dirEmpty := path.Join(dirFull, "empty")

		//
		//
		//
		maxPages := 0
		for _, q := range qs {
			if maxPages < len(q.Pages) {
				maxPages = len(q.Pages)
			}
		}

		keysByQ := [][]string{} // per questionnaire
		valsByQ := [][]string{} // per questionnaire

		// CSV header stuff:
		staticCols := []string{ // across all questionnaires
			"user_id",
			"lang_code",
			"closing_time",
			"status",
			"remote_ip",
			"user_agent",
			"version",
			"version_max",
		}
		for iPg := 0; iPg < maxPages; iPg++ {
			staticCols = append(staticCols, fmt.Sprintf("page_%v", iPg+1))
		}

		nonEmpty := 0
		empty := 0

		//
		//
		// Process questionnaires
		for i, q := range qs {

			// log.Printf("  ")

			serverSideMD5 := q.MD5

			pthFull := path.Join(dirFull, q.UserID+".json")
			err := q.Save1(pthFull)
			if err != nil {
				log.Printf("%3v: Error saving %v: %v", i, pthFull, err)
				continue
			}

			//
			if q.MD5 != serverSideMD5 {
				// log.Printf("%3v: MD5 does not match: %v\nwnt %v\ngot %v", i, pth2, md5BeforeSave, q.MD5)
				log.Printf("%3v: Server side and new client side MD5 hashes do not match %v - %v", i, q.Survey.String(), pthFull)
			}

			if cfgRem.MinUserID != 0 {
				if q.UserIDInt() < cfgRem.MinUserID {
					log.Printf("%3v: Skipping UserID %v < %v", i, q.UserID, cfgRem.MinUserID)
					continue
				}
			}

			if cfgRem.MaxUserID != 0 {
				if q.UserIDInt() > cfgRem.MaxUserID {
					log.Printf("%3v: Skipping UserID %v > %v", i, q.UserID, cfgRem.MaxUserID)
					continue
				}
			}

			//
			//
			// Delete empty questionnaires and save them elsewhere
			//
			// previous runs: remove their empty files
			pthEmpty := path.Join(dirEmpty, q.UserID+".json")
			err = cloudio.Delete(pthEmpty)
			if err != nil && !cloudio.IsNotExist(err) {
				log.Printf("%3v: Error removing previously empty %v - %v", i, pthEmpty, err)
			}
			// current run: move empty to dir empty
			realEntries, _, _ := q.Statistics()
			if realEntries == 0 {
				log.Printf("%3v: %v. No answers, moving to %v.", i, pthFull, "empty")
				err = cloudio.Delete(pthFull)
				if err != nil && !cloudio.IsNotExist(err) {
					log.Printf("%3v: Error removing empty %v - %v", i, pthFull, err)
				}
				err := q.Save1(pthEmpty)
				if err != nil {
					log.Printf("%3v: Error saving  to empty %v: %v", i, pthEmpty, err)
				}
				empty++
				continue
			}

			nonEmpty++

			// Prepare columns...
			finishes, ks, vs := q.KeysValues(true)

			ks = append(staticCols, ks...)
			keysByQ = append(keysByQ, ks)

			formattedClosingTime := ""
			status := "0"
			if qs[i].ClosingTime.IsZero() {
				for i2 := len(qs[i].Pages) - 1; i2 > -1; i2-- {
					if !qs[i].Pages[i2].Finished.IsZero() {
						formattedClosingTime = fmt.Sprintf("%v", qs[i].Pages[i2].Finished.Unix())
						status = "1"
						break
					}
				}
			} else {
				formattedClosingTime = fmt.Sprintf("%v", qs[i].ClosingTime.Unix())
				status = "2"
			}

			// equivalent staticCols...
			prepend := []string{
				qs[i].UserID,         // user_id
				qs[i].LangCode,       // lang_code
				formattedClosingTime, // closing_time
				status,               // status
				q.RemoteIP,           // remote_ip
				qst.EnglishTextAndNumbersOnly(q.UserAgent), // user_agent
				fmt.Sprint(q.VersionEffective),             // version effective
				fmt.Sprint(q.VersionMax),                   // number of versions
			}
			for iPg := 0; iPg < maxPages; iPg++ {
				if iPg < len(finishes) {
					prepend = append(prepend, finishes[iPg])
				} else {
					prepend = append(prepend, "n.a.") // response had less than max pages - not finishing time
				}
			}
			vs = append(prepend, vs...)
			valsByQ = append(valsByQ, vs)

		} // forr questionnaires

		allKeysSuperset := Superset(keysByQ)

		allKeysSSMap := map[string]int{}
		for idx, v := range allKeysSuperset {
			allKeysSSMap[v] = idx
		}
		valsBySuperset := [][]string{}

		for colIdx, colName := range allKeysSuperset {
			log.Printf("\tcol %2v  %v", colIdx, colName)
		}

		// Collect values...
		for i1 := 0; i1 < len(valsByQ); i1++ {
			keys := keysByQ[i1]
			vals := valsByQ[i1]
			valsBySuperset = append(valsBySuperset, make([]string, len(allKeysSuperset)))
			for i2 := 0; i2 < len(keys); i2++ {
				v := vals[i2]
				k := keys[i2]
				pos := allKeysSSMap[k]
				valsBySuperset[i1][pos] = v
			}
		}

		// Data into CSV matrix...
		var wtr = new(bytes.Buffer)
		csvWtr := csv.NewWriter(wtr)
		csvWtr.Comma = ';'
		if err := csvWtr.Write(allKeysSuperset); err != nil {
			log.Printf("error writing header line to csv: %v", err)
		}
		for _, record := range valsBySuperset {
			if err := csvWtr.Write(record); err != nil {
				log.Printf("error writing record to csv: %v", err)
			}
		}

		// Write any buffered data to the underlying writer (standard output).
		csvWtr.Flush()
		if err := csvWtr.Error(); err != nil {
			log.Printf("error flushing csv to response writer: %v", err)
		}

		fn := path.Join(cfgRem.DownloadDir, fmt.Sprintf("%v-%v.csv", cfgRem.SurveyType, cfgRem.WaveID))
		err = cloudio.WriteFile(fn, wtr, 0644)
		if err != nil {
			log.Printf("Could not write file %v: %v", fn, err)
		}

		//
		// Labels into separate CSV file
		if len(qs) > 0 {

			nams := []string{} // input names
			lbls := []string{} // input labels

			fnCore := cfgRem.SurveyType + "-" + cfgRem.WaveID
			pthBase := path.Join(qst.BasePath(), fnCore+".json")
			qBase, err := qst.Load1(pthBase)
			if err != nil {
				log.Printf("Loading base questionnaire error %v", err)
			}

			// enclosing every cell value in double quotes allows to include newlines
			// excelWindowsNewline is the inside cell newlince character for Excel under Windows
			// excel newline for windows - inside cells
			const excelNL = string(rune(int32(10)))

			// copy(staticLabels, staticCols)
			byNames, _, _ := qBase.LabelsByInputNames()
			for _, name := range allKeysSuperset {
				nams = append(nams, name)
				if lbl, ok := byNames[name]; ok {
					if !strings.HasPrefix(lbl, excelNL) {
						lbl += excelNL
					}
					lbl = "\"" + strings.ReplaceAll(lbl, " -- ", excelNL) + "\""
					lbls = append(lbls, lbl)
				} else {
					lbls = append(lbls, name)
				}
			}

			buf := &bytes.Buffer{}
			buf.WriteString(strings.Join(nams, ";"))
			buf.WriteString("\n")
			buf.WriteString(strings.Join(lbls, ";"))

			fnLabels := strings.ReplaceAll(fn, ".csv", "-labels.csv")
			err = cloudio.WriteFile(fnLabels, buf, 0644)
			if err != nil {
				log.Printf("writing file failed: %v - error %v", fnLabels, err)
			}

		}

		log.Printf(
			"\n\nRegular finish. %v questionnaire(s) processed\n%v non empty - %v empty\nresults in %v\n\n", len(qs),
			nonEmpty, empty, fn,
		)

	}

}
