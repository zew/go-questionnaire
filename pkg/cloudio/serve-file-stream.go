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
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/stream"
)

// ServeFileStream writes a binary file into the response;
// it's inner body is similar to ReadFile;
// the URL for this handler must be exempted from session middleware and
// logAndRecover middleware - anything buffering the response.
// URL GET parameter hijack [true,false] governs the way of flushing;
// ordinary flushing requires middleware free unbuffered http.ResponseWriter
// and after each write a
//     flusher, ok := w.(http.Flusher) ...flusher.Flush()
// instead, we use
// 		w ... := stream.NewFlushable(w)
// with auto flushing;
// however, stream.NewFlushable() *hijacks* the connection;
// preempting middleware but also requiring explicit headers and
// 		fmt.Fprintf(w, "HTTP/1.1 200 OK\n")
//
// URL GET parameter file-size must contain the correct file size;
// URL GET parameter h (hask)  must contain a checksum;
func ServeFileStream(w http.ResponseWriter, req *http.Request) {

	// convenience
	logAndShow := func(f string, intf ...interface{}) {
		fmt.Fprintf(w, "HTTP/1.1 200 OK\n") // required for hijacked handling
		fmt.Fprintf(w, "\n")                // required for hijacked handling
		fmt.Fprintf(w, f+"<br>\n", intf...)
		log.Printf("\t"+f+"\n", intf...)
	}

	// streaming
	hijack := req.URL.Query().Get("hijack") == "true"
	if hijack {
		wHijacked, closer, err := stream.NewFlushable(w)
		defer closer()
		if err != nil {
			logAndShow("cloudio.ServeFileStream(): Error getting flushable writer: %v", err)
			return
		}
		w = wHijacked
	}

	pth := req.URL.Path
	pth = strings.TrimPrefix(pth, cfg.Pref())
	pth = strings.Trim(pth, "/")
	pth = strings.ReplaceAll(pth, "..", "") // prevent climbing up, such as ./app-bucket/../../../root/passwd
	// fpth := path.Join(".", "app-bucket", pth)
	fpth := path.Join(".", pth)
	fpth = strings.Replace(fpth, "download-stream", "download", -1)

	pthCheck := strings.TrimPrefix(pth, "download-stream")
	hshGot := req.URL.Query().Get("h")
	tNow := time.Now().Unix() / 100
	hshWnt1 := Md5Str(fmt.Sprintf("%v%v%v", pthCheck, tNow-1, cfg.Get().AppInstanceID))
	hshWnt2 := Md5Str(fmt.Sprintf("%v%v%v", pthCheck, tNow-0, cfg.Get().AppInstanceID))
	log.Printf("hash over\n\t%v\n\t%v\n\t%v\n", pthCheck, tNow, cfg.Get().AppInstanceID)
	if hshGot != hshWnt1 && hshGot != hshWnt2 {
		logAndShow("cloudio.ServeFileStream(): provided checksum invalid")
		return
	}

	ctx := context.Background()
	// var buck *blob.Bucket
	var errSec error

	// Bucket / directory
	buck, err := bucket()
	if err != nil {
		logAndShow("cloudio.ServeFileStream(): Error opening bucket: %v", err)
		return
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			logAndShow("cloudio.ServeFileStream(): Error closing bucket: %v", errSec)
		}
	}()

	// Reader of "filename" in bucket
	r, err := buck.NewReader(ctx, fpth, nil)
	if err != nil {
		if IsNotExist(err) {
			logAndShow("cloudio.ServeFileStream(): %v does not exist (%v)", fpth, err)
		} else {
			logAndShow("cloudio.ServeFileStream(): Error opening writer to bucket for %v: %v", fpth, err)
		}
		return
	}
	defer func() {
		errSec = r.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			logAndShow("cloudio.ServeFileStream(): Error closing writer to bucket: %v", errSec)
		}
	}()

	fileSize, err := strconv.ParseInt(req.URL.Query().Get("file-size"), 10, 64)
	if err != nil || fileSize == 0 {
		logAndShow("cloudio.ServeFileStream(): Error getting file size '%v' for %v: %v", fileSize, fpth, err)
		return
	}

	log.Printf("cloudio.ServeFileStream(): filename is %v, hijacked %v, fileSize %v kB", fpth, hijack, fileSize/1024)

	if hijack {
		fmt.Fprintf(w, "HTTP/1.1 200 OK\n")
	}
	if hijack {
		fmt.Fprintf(w, "Content-Length: %d\n", fileSize) // instead of fInfo.Size(); sending headers via w - instead of w.Header().Set(...)
	} else {
		w.Header().Set("Content-Length", fmt.Sprintf("%v", fileSize)) // instead of fInfo.Size()
	}
	m := mime.TypeByExtension(filepath.Ext(pth))
	if m != "" {
		if hijack {
			fmt.Fprintf(w, "Content-Type: %s\n", m) // sending headers via w - instead of w.Header().Set(...)
		} else {
			w.Header().Set("Content-Type", m)
		}
	}
	// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
	// w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*24))
	if hijack {
		fmt.Fprintf(w, "\n") // end of response headers
	}

	//
	// Reading bytes from reader from bucket
	bts, err := io.Copy(w, r) // most memory efficient
	if err != nil {
		logAndShow("cloudio.ServeFileStream(): Could not copy file stream into response writer %v: %v", fpth, err)
	}
	log.Printf("cloudio.ServeFileStream(): %v bytes written", bts)

}
