package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/util"
)

// TestCloudStore checks if the cloud store is readable and writable;
// Implementation in package cloudio causes circular dependencies
func TestCloudStore(w1 http.ResponseWriter, r *http.Request) {
	w1.Header().Set("Content-Type", "text/html; charset=utf-8")

	w := &bytes.Buffer{}

	// fetch all env variables
	for _, element := range os.Environ() {
		keyVal := strings.Split(element, "=")
		if len(keyVal) == 2 {
			fmt.Fprintf(w, "%30s => %s <br>\n", keyVal[0], keyVal[1])
		} else {
			fmt.Fprintf(w, "%30v <br>\n", keyVal)
		}
	}

	// CloudStore()
	bts, err := cloudio.ReadFile("cloudio-test-file.json")
	if err != nil {
		fmt.Fprintf(w, "Error reading from bucket: %v<br>\n", err)
	} else {
		fmt.Fprintf(w, "Reading from bucket SUCCESS<br>\n")
	}

	s := string(bts)
	s += " - " + time.Now().Format(time.RFC1123)

	err = cloudio.WriteFile("cloudio-test-file.json", strings.NewReader(s), 0644)
	if err != nil {
		fmt.Fprintf(w, "Error writing to bucket: %v<br>\n", err)
	} else {
		fmt.Fprintf(w, "Writing to bucket SUCCESS<br>\n")
	}

	fmt.Fprintf(w, "Config is <pre>%v</pre>", util.IndentedDump(cfg.Get()))

	tpl.ExecContent(w1, r, w.String(), "main-desktop.html")

}
