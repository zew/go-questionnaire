package handlers

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/qst"

	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
)

func logAndRespond(w http.ResponseWriter, r *http.Request, s string, err error) {

	fmt.Fprintf(w, "%v<br>\n", s)
	if err != nil {
		fmt.Fprintf(w, "%v<br>\n", err)
	}

	log.Printf("%v", s)
	if err != nil {
		log.Printf("    %v", err)
	}

	if r != nil && r.Response != nil {
		r.Response.StatusCode = 401
		r.Response.Status = fmt.Sprintf("%v %v", s, err)
	}
}

// TransferrerEndpointH responds with finished questionnaires from the filesystem in JSON;
// preventing of huge filesizes, the response is gzipped;
// you need to be logged in with admin role;
// survey_id and wave_id must be set as URL params;
// only finished questionnaires are included (q.ClosingTime != zero);
// fetch_all=1 includes unfinished questionnaires;
func TransferrerEndpointH(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")

	// w.Header().Set("Content-Length", fmt.Sprintf("%v", len(byts)))  // do not set, if response is gzipped !
	// w.Write(byts)

	deadLine, ok := r.Context().Deadline()
	log.Printf("transferrer-endpoint-start; has deadline %v of %v", ok, deadLine)

	sess := sessx.New(w, r)

	l, isLoggedIn, err := lgn.LoggedInCheck(w, r)
	if err != nil {
		logAndRespond(w, r, "LoggedInCheck failed.", err)
		return
	}
	if !isLoggedIn {
		logAndRespond(w, r, "You are are not logged in.", nil)
		return
	}
	if !l.HasRole("admin") {
		logAndRespond(w, r, "Login succeeded, but must have role 'admin'", nil)
		return
	}

	fetchAll, _ := sess.ReqParam("fetch_all")

	surveyID, ok := sess.ReqParam("survey_id")
	if !ok {
		logAndRespond(w, r, "You need to specify a survey_id parameter.", nil)
		return
	}
	waveID, ok := sess.ReqParam("wave_id")
	if !ok {
		logAndRespond(w, r, "You need to specify a wave_id parameter.", nil)
		return
	}
	pth := path.Join(qst.BasePath(), surveyID, waveID)

	log.Printf("transferrer-endpoint-reading-directory %v", pth)
	infos, err := cloudio.ReadDir(pth)
	// infos, err := dir.Readdir(-1)
	if err != nil {
		logAndRespond(w, r, "Could not read directory.", err)
		return
	}

	cntr := 0
	btsCtr := 0
	gz := gzip.NewWriter(w)
	defer gz.Close()

	gz.Write([]byte("["))
	for i, info := range *infos {
		if !info.IsDir {
			if i < 10 || i%50 == 0 {
				log.Printf("iter %3v: Name: %v, Size: %v", i, info.Key, info.Size)
			}
		}

		if strings.HasSuffix(info.Key, ".inspection") {
			continue
		}

		// pth := path.Join(qst.BasePath(), surveyID, waveID, info.Key)
		pth := info.Key
		// var q = &qst.QuestionnaireT{}
		q, err := qst.Load1(pth)
		if err != nil {
			logAndRespond(w, r, fmt.Sprintf("iter %3v: No file %v found.", i, pth), err)
			return
		}

		/*
			this is no longer applicable for the pure data files
			err = q.Validate()
			if err != nil {
				logAndRespond(w, r, fmt.Sprintf("iter %3v: Questionnaire validation caused error", i), nil)
			}
		*/

		// user questionnaire unfinished
		if q.ClosingTime.IsZero() && fetchAll == "" {
			log.Printf("%v unfinished yet; %v", info.Key, q.ClosingTime)
			if time.Now().Before(q.Survey.Deadline) {
				log.Printf("%v not yet past global deadline => skipping", info.Key)
				continue
			}
		}

		firstColLeftMostPrefix := " "
		byts, err := json.MarshalIndent(q, firstColLeftMostPrefix, "\t")
		if err != nil {
			logAndRespond(w, r, "marshalling questionnaire failed", err)
			return
		}

		if cntr > 0 {
			gz.Write([]byte(","))
		}
		gzipBytes, err := gz.Write(byts)
		if err != nil {
			logAndRespond(w, r, "gzipping questionnaire failed: %v", err)
			return
		}
		cntr++
		btsCtr += gzipBytes

	}
	gz.Write([]byte("]"))
	sz1 := fmt.Sprintf("%.3f MB", float64(btsCtr/(1<<10))/(1<<10))
	log.Printf("%v questionnaires to http response written - gzipped %v", cntr, sz1)

}
