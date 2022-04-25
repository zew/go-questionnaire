package handlers

import (
	"log"
	"net/http"
	"path"

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

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")

	// w.Header().Set("Content-Length", fmt.Sprintf("%v", len(byts)))  // do not set, if response is gzipped !
	// w.Write(byts)

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
	waveID, ok := sess.ReqParam("wave_id")
	if !ok {
		tf.LogAndRespond(w, r, "You need to specify a wave_id parameter.", nil)
		return
	}
	pth := path.Join(qst.BasePath(), surveyID, waveID)

	tf.RetrieveFromLocal(w, r, pth, fetchAll)

}
