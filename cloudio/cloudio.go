// Package cloudio is a wrapper for godoc.org/gocloud.dev/blob
// emulating ioutil.WriteFile and ioutil.ReadFile.
//
// It is zero config; either saving to local ./app-bucket/
// or to appenginge bucket <appID>, depending on environment variables.
//
// The zero configuration is important, cause we load the *actual* configuration file
// with this package, and want to avoid circular trouble or bootstrap hell.
//
// MarshalWriteFile and ReadFileUnmarshal incorporate
// JSON serialization and deserialization.
//
// Open() is similar to file.Open
//    r, err := file.Open("name")
// but deviates in that is also returns a bucket closer func.
//
// OpenAny() is just a wrapper arond Open() searching in various subdirectories.
package cloudio

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob" // local file system
	_ "gocloud.dev/blob/gcsblob"
)

// IsNotExist function for cloudio
func IsNotExist(err error) bool {
	if os.IsNotExist(err) || strings.Contains(err.Error(), "code=NotFound") {
		return true
	}
	return false
}

func prepareLocalDir() error {
	// bucketDir := cfg.Get().StorageDriverURL
	// bucketDir = strings.TrimPrefix(bucketDir, "file:///")
	bucketDir := filepath.Join(".", "app-bucket")
	if err := os.MkdirAll(filepath.Join(".", bucketDir), 0755); err != nil {
		return err // MkdirAll does not report "already exists" as error
	}
	return nil
}

func bucketLocal() (*blob.Bucket, error) {
	err := prepareLocalDir()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	// storageDriverURL := cfg.Get().StorageDriverURL
	storageDriverURL := fmt.Sprintf("file:///%s", filepath.Join(".", "app-bucket"))
	bucket, err := blob.OpenBucket(ctx, storageDriverURL)
	if err != nil {
		return nil, fmt.Errorf("could not open local bucket for %v: %v", storageDriverURL, err)
	}
	return bucket, nil
}

func bucketGoogle() (*blob.Bucket, error) {
	ctx := context.Background()
	envCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	if envCreds == "" {
		log.Print("SET GOOGLE_APPLICATION_CREDENTIALS=~/.ssh/google-cloud-[appname]-creds.json")
	}
	appID := os.Getenv("GAE_APPLICATION")
	if strings.HasPrefix(appID, "e~") {
		appID = strings.TrimPrefix(appID, "e~")
	}
	if strings.HasPrefix(appID, "h~") {
		appID = strings.TrimPrefix(appID, "h~")
	}
	// storageDriverURL := cfg.Get().StorageDriverURL
	storageDriverURL := fmt.Sprintf("gs://%s.appspot.com", appID)
	bucket, err := blob.OpenBucket(ctx, storageDriverURL)
	if err != nil {
		return nil, fmt.Errorf("could not open google bucket for %v: %v", storageDriverURL, err)
	}
	return bucket, nil
}

func bucket() (*blob.Bucket, error) {
	// se := cfg.Get().StorageEngine
	// if se == "gcsblob" {
	// } else if se == "fileblob" || se == "memblob" {
	// }
	// return nil, fmt.Errorf("Unknown storage engine: %v", se)
	appID := os.Getenv("GAE_APPLICATION")
	if appID != "" {
		return bucketGoogle()
	}
	return bucketLocal()
}

// WriteFile is the cousin of ioutil.WriteFile.
// Permissions of parameter perm are not implemented.
//
// No memory allocation. intf is streamed into blob.
func WriteFile(fileName string, r io.Reader, perm os.FileMode) (err error) {

	ctx := context.Background()
	var buck *blob.Bucket
	var errSec error

	// Bucket / directory
	buck, err = bucket()
	if err != nil {
		log.Printf("Error opening bucket: %v", err)
		return
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			log.Printf("Error closing bucket: %v", errSec)
		}
	}()

	// Writer to "filename" in bucket
	w, err := buck.NewWriter(ctx, fileName, nil)
	if err != nil {
		log.Printf("Error opening writer to bucket: %v", err)
		return
	}
	defer func() {
		errSec = w.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			log.Printf("Error closing writer to bucket: %v", errSec)
		}
	}()

	// Writing bytes to writer to bucket
	_, err = io.Copy(w, r) // most memory efficient
	if err != nil {
		log.Printf("Error writing to bucket %v: %v", fileName, err)
	}

	return
}

// ReadFile is the cousin of ioutil.ReadFile
// Memory allocation for the file contents.
//
// Open() is another version returning a reader
// but also an io.ReadCloser and an io.Closer - for bucket and reader.
func ReadFile(fileName string) (bts []byte, err error) {

	ctx := context.Background()
	var buck *blob.Bucket
	var errSec error

	// Bucket / directory
	buck, err = bucket()
	if err != nil {
		log.Printf("Error opening bucket: %v", err)
		return
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			log.Printf("Error closing bucket: %v", errSec)
		}
	}()

	// Reader of "filename" in bucket
	r, err := buck.NewReader(ctx, fileName, nil)
	if err != nil {
		log.Printf("Error opening writer to bucket: %v", err)
		return
	}
	defer func() {
		errSec = r.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			log.Printf("Error closing writer to bucket: %v", errSec)
		}
	}()

	// Reading bytes from reader from  bucket
	buf := &bytes.Buffer{}
	_, err = io.Copy(buf, r) // most memory efficient
	if err != nil {
		log.Printf("Error reading from reader from bucket: %v", errSec)
	}
	bts = buf.Bytes()
	return

}

/*
Open a blob/file fileName in default bucket.
Returns ReadCloser for the blob/file and bucketCloser() for the underlying bucket

No memory allocation

    fileName := "config.json"
    r, bucketClose, err := cloudio.Open(fileName)
    if err != nil {
        log.Fatalf("Error opening writer to %v: %v", fileName, err)
    }
    defer func() {
        err := r.Close()
        if err != nil {
            log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
        }
    }()
    defer func() {
        err := bucketClose()
        if err != nil {
            log.Printf("Error closing bucket of writer to %v: %v", fileName, err)
        }
    }()

*/
func Open(fileName string) (r io.ReadCloser, bucketClose func() error, err error) {

	ctx := context.Background()
	var buck *blob.Bucket

	// Bucket / directory
	buck, err = bucket()
	if err != nil {
		log.Printf("Error opening bucket: %v", err)
		return
	}
	bucketClose = func() error {
		err := buck.Close()
		if err != nil {
			log.Printf("Error closing bucket: %v", err)
		}
		return err
	}

	// Reader of "fileName" in bucket
	r, err = buck.NewReader(ctx, fileName, nil)
	if err != nil {
		log.Printf("Error opening writer to bucket: %v", err)
		return
	}
	return
}

// OpenAny tries Open(fileName) for any subdirs and then the appdir "."
func OpenAny(fileName string, optSubdirs ...string) (r io.ReadCloser, bucketClose func() error, err error) {

	workDir, err := os.Getwd()
	if err != nil {
		return
	}
	workDir = "." // same as above

	log.Printf("work dir: %v, subdirs: %v, fileName: %v", workDir, optSubdirs, fileName)

	paths := []string{}
	for _, subDir := range optSubdirs {
		if len(subDir) > 0 {
			paths = append(paths, path.Join(workDir, subDir, fileName))
		}
	}
	paths = append(paths, path.Join(workDir, fileName))

	//
	for _, v := range paths {
		r, bucketClose, err = Open(v)
		if err != nil && IsNotExist(err) {
			log.Printf("Not found in  %v", v)
			continue
		} else if err != nil {
			err = fmt.Errorf("Error searching for file: %v", err)
			return
		}
		log.Printf("Found file in: %v", v)
		break
	}
	return
}

// MarshalWriteFile is like WriteFile - but marshals to JSON first
// allocates contents of intf into []byte slice
func MarshalWriteFile(intf interface{}, fileName string) error {

	firstColLeftMostPrefix := " "
	bts, err := json.MarshalIndent(intf, firstColLeftMostPrefix, "\t")
	if err != nil {
		log.Printf("Could not marshal example config %v", err)
		return err
	}
	err = WriteFile(fileName, bytes.NewReader(bts), 0644)
	if err != nil {
		log.Printf("Could not save example config: %v", err)
		return err
	}
	log.Printf("File %v saved", fileName)
	return nil

}

// ReadFileUnmarshal is like ReadFile - but unmarshals from JSON afterwards;
// allocates contents of file into []byte slice
func ReadFileUnmarshal(fileName string, intf interface{}) (err error) {

	bts, err := ReadFile(fileName)
	if err != nil {
		log.Printf("Cannot read from file: %v -  %v\n", fileName, err)
		return
	}

	err = json.Unmarshal(bts, intf)
	if err != nil {
		if len(bts) > 300 {
			bts = bts[0:300]
		}
		log.Printf("Cannot unmarshal %v - %v: \n%v\n", fileName, string(bts), err)
		return
	}
	return

}

// Delete a file
func Delete(fileName string) error {
	ctx := context.Background()
	var buck *blob.Bucket
	var errSec error

	// Bucket / directory
	buck, err := bucket()
	if err != nil {
		log.Printf("Error opening bucket for file deletion: %v", err)
		return err
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = errors.Wrap(err, errSec.Error())
			log.Printf("Error closing bucket for file deletion: %v", errSec)
		}
	}()
	err = buck.Delete(ctx, fileName)
	if err != nil {
		if IsNotExist(err) {
			return nil
		}
		log.Printf("Error deleting %v from bucket: %v", fileName, err)
		return err
	}
	log.Printf("File %v successfully deleted from bucket", fileName)
	return nil
}
