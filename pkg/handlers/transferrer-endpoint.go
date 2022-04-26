package handlers

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/tf"

	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/sessx"
)

// TransferrerEndpointH responds with finished questionnaires from the filesystem in JSON;
// preventing of huge filesizes, the response is gzipped;
// you need to be logged in with admin role;
// survey_id and wave_id must be set as URL params;
// only finished questionnaires are included (q.ClosingTime != zero);
// fetch_all=1 includes unfinished questionnaires;
func TransferrerEndpointH(w http.ResponseWriter, r *http.Request) {

	deadLine, ok := r.Context().Deadline()
	log.Printf("transferrer-endpoint-start; has deadline %v of %v", ok, deadLine)

	sess := sessx.New(w, r)

	l, isLoggedIn, err := lgn.LoggedInCheck(w, r)
	if err != nil {
		tf.LogAndRespond(w, r, "LoggedInCheck failed.", err)
		return
	}
	if !isLoggedIn {
		tf.LogAndRespond(w, r, "You are are not logged in.", nil)
		return
	}
	if !l.HasRole("admin") {
		tf.LogAndRespond(w, r, "Login succeeded, but must have role 'admin'", nil)
		return
	}

	fetchAll, _ := sess.ReqParam("fetch_all")

	surveyID, ok := sess.ReqParam("survey_id")
	if !ok {
		tf.LogAndRespond(w, r, "You need to specify a survey_id parameter.", nil)
		return
	}

	extraRoleRequired := fmt.Sprintf("%v-downloader", surveyID)
	if !l.HasRole(extraRoleRequired) {
		tf.LogAndRespond(w, r, fmt.Sprintf("Login succeeded, but must have role '%v'", extraRoleRequired), nil)
		return
	}

	waveID, ok := sess.ReqParam("wave_id")
	if !ok {
		tf.LogAndRespond(w, r, "You need to specify a wave_id parameter.", nil)
		return
	}

	pth := path.Join(qst.BasePath(), surveyID, waveID)

	//
	//
	qs, err := tf.RetrieveFromLocal(pth, fetchAll)
	if err != nil {
		tf.LogAndRespond(w, r, "Retrieval from local file system faild: %w.", err)
		return
	}

	//
	// GZIP mode - start
	format, _ := sess.ReqParam("format")
	if format != "CSV" {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Content-Encoding", "gzip")
		// w.Header().Set("Content-Length", fmt.Sprintf("%v", len(byts)))  // do not set, if response is gzipped !
		// w.Write(byts)
		tf.PipeQStoResponse(w, r, qs)
		return
	}
	// GZIP mode - end
	//

	//
	// CSV direct download mode - start
	//  direct download requires only a *minimal* config for the requested survey_id stored on the server
	// 	whereas client mode requires lots of configs for standalone operation.
	remoteCfgPath := path.Join("transferrer", fmt.Sprintf("%v-remote.json", surveyID))
	// instead of cfgRem := tf.LoadRemote()
	cfgRem := &tf.RemoteConnConfigT{}
	err = cloudio.ReadFileUnmarshal(remoteCfgPath, cfgRem)
	if err != nil {
		s := fmt.Sprintf("error reading transferrer config %v: %%v", remoteCfgPath)
		tf.LogAndRespond(w, r, s, err)
		return
	}
	// filling in survey name and wave ID from URL request
	cfgRem.SurveyType = surveyID
	cfgRem.WaveID = waveID

	saveQSFilesToDownloadDir := false
	csvPath, err := tf.ProcessQs(cfgRem, qs, saveQSFilesToDownloadDir)
	if err != nil {
		tf.LogAndRespond(w, r, "error processing questionnaires from remote: %v", err)
		return
	}
	log.Printf("CSV file saved under: %v", csvPath)

	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", "attachment; filename="+path.Base(csvPath))

	bts, err := cloudio.ReadFile(csvPath)
	if err != nil {
		tf.LogAndRespond(w, r, "error opening CSV: %v", err)
		return
	}
	err = cloudio.Delete(csvPath) // first delete the CSV, then serve the bytes
	if err != nil {
		tf.LogAndRespond(w, r, "error deleting CSV: %v", err)
	}

	w.Write(bts)
	// CSV direct download mode - start
	//

}
