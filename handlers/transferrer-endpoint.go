package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/zew/go-questionaire/qst"

	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/util"
)

// TransferrerEndpointH responds with finished questionaires from the filesystem.
func TransferrerEndpointH(w http.ResponseWriter, r *http.Request) {

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

	infos, err := dir.Readdir(-1)
	if err != nil {
		helper(w, r, err, "Could not read directory.")
		return
	}
	qs := []*qst.QuestionaireT{}
	for i, info := range infos {
		if info.Mode().IsRegular() {
			log.Printf("Name: %v, Size: %v", info.Name(), info.Size())
		}
		pth := filepath.Join(qst.BasePath(), surveyID, waveID, info.Name())
		// var q = &qst.QuestionaireT{}
		q, err := qst.Load1(pth)
		if err != nil {
			helper(w, r, err, fmt.Sprintf("iter %3v: No file %v found.", i, pth))
		}
		err = q.Validate()
		if err != nil {
			helper(w, r, err, fmt.Sprintf("iter %3v: Questionaire validation caused error", i))
		}

		if q.ClosingTime.IsZero() {
			log.Printf("%v unfinished yet; %v", info.Name(), q.ClosingTime)
			if time.Now().Before(q.Survey.Deadline) {
				log.Printf("%v not yet past global deadline => skipping", info.Name())
				continue
			}
		}
		qs = append(qs, q)
	}

	log.Printf("%3v questionaires ready for fetching home", len(qs))

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(qs, firstColLeftMostPrefix, "\t")
	if err != nil {
		helper(w, r, fmt.Errorf("Marschalling questionair failed: %v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(byts)))
	w.Write(byts)

}
