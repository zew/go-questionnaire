package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/zew/go-questionnaire/qst"

	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/util"
)

// TransferrerEndpointH responds with finished questionnaires from the filesystem.
func TransferrerEndpointH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")

	// w.Header().Set("Content-Length", fmt.Sprintf("%v", len(byts)))  // do not set, if response is gzipped !
	// w.Write(byts)

	dl, ok := r.Context().Deadline()
	log.Printf("transferrer-endpoint-start; has deadline %v of %v", ok, dl)

	sess := sessx.New(w, r)

	l, isLoggedIn, err := lgn.LoggedInCheck(w, r)
	if err != nil {
		helper(w, r, err, "LoggedInCheck failed.")
		return
	}
	if !isLoggedIn {
		helper(w, r, nil, "You are are not logged in.")
		return
	}
	if !l.HasRole("admin") {
		helper(w, r, nil, "Login succeeded, but must have role 'admin'")
		return
	}

	surveyID, ok := sess.ReqParam("survey_id")
	if !ok {
		helper(w, r, nil, "You need to specify a survey_id parameter.")
		return
	}
	waveID, ok := sess.ReqParam("wave_id")
	if !ok {
		helper(w, r, nil, "You need to specify a wave_id parameter.")
		return
	}
	pth := filepath.Join(qst.BasePath(), surveyID, waveID)
	dir, err := util.Directory(pth)
	if err != nil {
		helper(w, r, err, "Your wave_id value pointed to a non existing directory.")
		return
	}
	defer dir.Close()

	fetchAll, _ := sess.ReqParam("fetch_all")

	log.Printf("transferrer-endpoint-reading-directory %v", dir.Name())
	infos, err := dir.Readdir(-1)
	if err != nil {
		helper(w, r, err, "Could not read directory.")
		return
	}

	// qs := []*qst.QuestionnaireT{}
	cntr := 0
	btsCtr := 0
	gz := gzip.NewWriter(w)
	defer gz.Close()

	gz.Write([]byte("["))
	for i, info := range infos {
		if info.Mode().IsRegular() {
			if i < 10 || i%50 == 0 {
				log.Printf("iter %3v: Name: %v, Size: %v", i, info.Name(), info.Size())
			}
		}
		pth := filepath.Join(qst.BasePath(), surveyID, waveID, info.Name())
		// var q = &qst.QuestionnaireT{}
		q, err := qst.Load1(pth)
		if err != nil {
			helper(w, r, err, fmt.Sprintf("iter %3v: No file %v found.", i, pth))
		}
		err = q.Validate()
		if err != nil {
			helper(w, r, err, fmt.Sprintf("iter %3v: Questionnaire validation caused error", i))
		}

		if q.ClosingTime.IsZero() && fetchAll == "" {
			log.Printf("%v unfinished yet; %v", info.Name(), q.ClosingTime)
			if time.Now().Before(q.Survey.Deadline) {
				log.Printf("%v not yet past global deadline => skipping", info.Name())
				continue
			}
		}

		firstColLeftMostPrefix := " "
		byts, err := json.MarshalIndent(q, firstColLeftMostPrefix, "\t")
		if err != nil {
			helper(w, r, fmt.Errorf("marshalling questionnaire failed: %v", err))
			return
		}

		if cntr > 0 {
			gz.Write([]byte(","))
		}
		gzipBytes, err := gz.Write(byts)
		if err != nil {
			helper(w, r, fmt.Errorf("gzipping questionnaire failed: %v", err))
			return
		}
		cntr++
		btsCtr += gzipBytes

	}
	gz.Write([]byte("]"))
	sz1 := fmt.Sprintf("%.3f MB", float64(btsCtr/(1<<10))/(1<<10))
	log.Printf("%v questionnaires to http response written - gzipped %v", cntr, sz1)

}
