package cloudio

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zew/go-questionnaire/cfg"
	"gocloud.dev/blob"
)

/*ServeFileBulk writes a binary file into the response;
it's inner body is similar to ReadFile;
files > 128 kB are offloaded to ServeFileStream()
at hugely reduced memory consumption.
As ServeFileStream() can not read session data,
since the session middleware brings a buffered writer,
we need to authorize the hand off with a hash
*/
func ServeFileBulk(w http.ResponseWriter, req *http.Request) {

	// convenience
	logAndShow := func(f string, intf ...interface{}) {
		fmt.Fprintf(w, f+"<br>\n", intf...)
		log.Printf("\t"+f+"\n", intf...)
	}

	pth := req.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref())
	pth = strings.Trim(pth, "/")
	pth = strings.ReplaceAll(pth, "..", "") // prevent climbing up, such as ./app-bucket/../../../root/passwd
	fpth := path.Join(".", "app-bucket", pth)
	fpth = path.Join(".", pth)

	// log.Printf("cloudio.Stream(): initiating download %v", fpth)

	ctx := context.Background()
	var buck *blob.Bucket
	var errSec error

	// Bucket / directory
	buck, err := bucket()
	if err != nil {
		logAndShow("cloudio.ServeFileBulk(): Error opening bucket: %v", err)
		return
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			logAndShow("cloudio.ServeFileBulk(): Error closing bucket: %v", errSec)
		}
	}()

	// Reader of "filename" in bucket
	r, err := buck.NewReader(ctx, fpth, nil)
	if err != nil {
		if IsNotExist(err) {
			logAndShow("cloudio.ServeFileBulk(): %v does not exist (%v)", fpth, err)
		} else {
			logAndShow("cloudio.ServeFileBulk(): Error opening writer to bucket for %v: %v", fpth, err)
		}
		return
	}
	defer func() {
		errSec = r.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			logAndShow("cloudio.ServeFileBulk(): Error closing writer to bucket: %v", errSec)
		}
	}()

	attrs, err := buck.Attributes(ctx, fpth)
	if err != nil {
		logAndShow("cloudio.ServeFileBulk(): Error reading attrs for %v: %v", fpth, err)
		return
	}
	// handing off to streamer
	if attrs.Size > 128*1024 {
		log.Printf("handing off to ServeFileStream")
		pathReduced := req.URL.Path
		pathReduced = strings.TrimPrefix(pathReduced, "/")
		pathReduced = strings.TrimPrefix(pathReduced, cfg.Get().URLPathPrefix)
		pathReduced = strings.TrimPrefix(pathReduced, "/")
		pathReduced = strings.TrimPrefix(pathReduced, "download")

		// checksum
		tNow := time.Now().Unix() / 100
		hsh := Md5Str(fmt.Sprintf("%v%v%v", pathReduced, tNow, cfg.Get().AppInstanceID))
		log.Printf("hash over\n\t%v\n\t%v\n\t%v\n", pathReduced, tNow, cfg.Get().AppInstanceID)
		url := cfg.Pref(
			fmt.Sprintf("/download-stream/%s?file-size=%d&hijack=true&h=%v",
				pathReduced, attrs.Size, hsh,
			),
		)
		http.Redirect(w, req, url, http.StatusFound)
		return
	}

	w.Header().Set("Content-Length", fmt.Sprintf("%v", attrs.Size)) // instead of fInfo.Size()
	m := mime.TypeByExtension(filepath.Ext(pth))
	if m != "" {
		w.Header().Set("Content-Type", m)
	}
	// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
	// w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*24))

	// Reading bytes from reader from  bucket
	_, err = io.Copy(w, r) // most memory efficient
	if err != nil {
		logAndShow("cloudio.ServeFileBulk(): Could not copy file stream into response writer %v: %v", fpth, err)
	}

}
