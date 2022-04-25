package tf

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/qst"
)

func LogAndRespond(w http.ResponseWriter, r *http.Request, s string, err error) {

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

// RetrieveFromLocal reads the JSONified questionnaires
// from app-bucket on the survey server
func RetrieveFromLocal(
	pth string,
	fetchAll string,

) ([]*qst.QuestionnaireT, error) {

	var qs []*qst.QuestionnaireT

	log.Printf("transferrer-endpoint: reading from dir  %v", pth)
	infos, err := cloudio.ReadDir(pth)
	// infos, err := dir.Readdir(-1)
	if err != nil {
		return qs, fmt.Errorf("Could not read directory; %w", err)
	}
	log.Printf("transferrer-endpoint: found %v files", len(*infos))

	for i, info := range *infos {
		if !info.IsDir {
			if i < 10 || i%50 == 0 {
				log.Printf("    iter %3v: Name: %v, Size: %v", i, info.Key, info.Size)
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
			s := fmt.Sprintf("iter %3v: No file %v found", i, pth)
			return qs, fmt.Errorf(s+" - %w", err)
		}

		/*
			this is no longer applicable for the pure data files
			err = q.Validate()
			if err != nil {
				LogAndRespond(w, r, fmt.Sprintf("iter %3v: Questionnaire validation caused error", i), nil)
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

		qs = append(qs, q)

	}

	return qs, nil

}

// PipeQStoResponse takes a slice of questionnaires
// and writess them into the http response
func PipeQStoResponse(
	w http.ResponseWriter,
	r *http.Request,
	qs []*qst.QuestionnaireT,
) {

	cntr := 0
	btsCtr := 0
	gz := gzip.NewWriter(w)
	defer gz.Close()

	gz.Write([]byte("["))
	for _, q := range qs {

		firstColLeftMostPrefix := " "
		bts, err := json.MarshalIndent(q, firstColLeftMostPrefix, "\t")
		if err != nil {
			LogAndRespond(w, r, "marshalling questionnaire failed", err)
			return
		}

		if cntr > 0 {
			gz.Write([]byte(","))
		}
		gzipBytes, err := gz.Write(bts)
		if err != nil {
			LogAndRespond(w, r, "gzipping questionnaire failed: %v", err)
			return
		}
		cntr++
		btsCtr += gzipBytes

	}
	gz.Write([]byte("]"))
	sz1 := fmt.Sprintf("%.3f MB", float64(btsCtr/(1<<10))/(1<<10))
	log.Printf("%v questionnaires to http response written - gzipped %v", cntr, sz1)

}
