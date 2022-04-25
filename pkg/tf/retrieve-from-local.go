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

// RetrieveFromLocal requests the JSONified questionnaires
// from the survey server endpoint; decompresses the GZIPPed
// response and parses the bytes into a slice of questionnaires
func RetrieveFromLocal(
	w http.ResponseWriter,
	r *http.Request,
	pth string,
	fetchAll string,
	// cfgRem *RemoteConnConfigT,
) []*qst.QuestionnaireT { // error,

	var qs []*qst.QuestionnaireT

	log.Printf("transferrer-endpoint-reading-directory %v", pth)
	infos, err := cloudio.ReadDir(pth)
	// infos, err := dir.Readdir(-1)
	if err != nil {
		LogAndRespond(w, r, "Could not read directory.", err)
		return qs
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
			LogAndRespond(w, r, fmt.Sprintf("iter %3v: No file %v found.", i, pth), err)
			return qs
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

		firstColLeftMostPrefix := " "
		byts, err := json.MarshalIndent(q, firstColLeftMostPrefix, "\t")
		if err != nil {
			LogAndRespond(w, r, "marshalling questionnaire failed", err)
			return qs
		}

		if cntr > 0 {
			gz.Write([]byte(","))
		}
		gzipBytes, err := gz.Write(byts)
		if err != nil {
			LogAndRespond(w, r, "gzipping questionnaire failed: %v", err)
			return qs
		}
		cntr++
		btsCtr += gzipBytes

	}
	gz.Write([]byte("]"))
	sz1 := fmt.Sprintf("%.3f MB", float64(btsCtr/(1<<10))/(1<<10))
	log.Printf("%v questionnaires to http response written - gzipped %v", cntr, sz1)

	return qs

}
