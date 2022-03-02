// Package cloudio is a wrapper for godoc.org/gocloud.dev/blob
// emulating os.WriteFile and os.ReadFile.
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
//
// ServeFileBulk() serves files to via http
// ServeFileStream() serves large files as stream
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

	"cloud.google.com/go/storage"
	"gocloud.dev/blob"
	_ "gocloud.dev/blob/fileblob" // local file system
	_ "gocloud.dev/blob/gcsblob"
)

var appsID string // Google app engine ID

// if executable is run in ./cmd/server
// var exeToAppRoot = path.Join("..", "..")
// if executable is in app root then
var exeToAppRoot = "."

func init() {
	appsID = os.Getenv("GAE_APPLICATION")
	if len(appsID) > 2 {
		// chopping of g~ or h~ ...
		tokens := strings.Split(appsID, "~")
		if len(tokens) > 1 {
			appsID = tokens[1]
		}
	}
}

func prepareLocalDir() error {
	// bucketDir := filepath.Join(".", "app-bucket")
	bucketDir := filepath.Join(exeToAppRoot, "app-bucket")
	if err := os.MkdirAll(filepath.Join(".", bucketDir), 0750); err != nil {
		return err // path error - "already exists" is not reported
	}
	return nil
}

func bucketLocal() (*blob.Bucket, error) {
	err := prepareLocalDir()
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	storageDriverURL := fmt.Sprintf("file:///%s", filepath.Join(wd, exeToAppRoot, "app-bucket")+"/") // relative directory not working on travis - but on appengine and windows
	// storageDriverURL := cfg.Get().StorageDriverURL
	bucket, err := blob.OpenBucket(ctx, storageDriverURL)
	if err != nil {
		return nil, fmt.Errorf("could not open local bucket for %v: %v", storageDriverURL, err)
	}
	return bucket, nil
}

func bucketGoogle() (*blob.Bucket, error) {
	ctx := context.Background()
	if len(appsID) < 2 {
		envCreds := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if envCreds == "" {
			log.Print("SET GOOGLE_APPLICATION_CREDENTIALS=~/.ssh/google-cloud-[appname]-creds.json")
		}
	}
	// storageDriverURL := cfg.Get().StorageDriverURL
	storageDriverURL := fmt.Sprintf("gs://%s.appspot.com", appsID)
	bucket, err := blob.OpenBucket(ctx, storageDriverURL)
	if err != nil {
		return nil, fmt.Errorf("could not open google bucket for %v: %v", storageDriverURL, err)
	}
	return bucket, nil
}

func bucket() (*blob.Bucket, error) {
	if appsID != "" {
		return bucketGoogle()
	}
	return bucketLocal()
}

// Attrs retrieves the attributes from a path;
// though file size seems the only use case.
func Attrs(fpth string) (attrs *blob.Attributes, err error) {

	ctx := context.Background()
	var errSec error
	attrs = &blob.Attributes{}

	// Bucket / directory
	buck, err := bucket()
	if err != nil {
		log.Printf("cloudio.Stream(): Error opening bucket: %v", err)
		return
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = combiErr{err, errSec}
			log.Printf("cloudio.Stream(): Error closing bucket: %v", errSec)
		}
	}()

	attrs, err = buck.Attributes(ctx, fpth)
	if err != nil {
		log.Printf("cloudio.Stream(): Error reading attrs for %v: %v", fpth, err)
		return
	}

	return

}

// WriteFile is the cousin of os.WriteFile.
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
			err = combiErr{err, errSec}
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
			err = combiErr{err, errSec}
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

// ReadFile is the cousin of os.ReadFile
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
			err = combiErr{err, errSec}
			log.Printf("Error closing bucket: %v", errSec)
		}
	}()

	// Reader of "filename" in bucket
	r, err := buck.NewReader(ctx, fileName, nil)
	if err != nil {
		if IsNotExist(err) {
			err = CreateNotExist(err)
		} else {
			log.Printf("Error opening writer to file in bucket: %v", err)
		}
		return
	}
	defer func() {
		errSec = r.Close()
		if errSec != nil {
			err = combiErr{err, errSec}
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

Example

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

	// workDir, err := os.Getwd()
	// if err != nil {
	// 	return
	// }
	workDir := "." // same as above

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

// MarshalWriteFile is like WriteFile - but marshals to JSON first;
// allocates contents of intf into []byte slice
func MarshalWriteFile(intf interface{}, fileName string) error {

	firstColLeftMostPrefix := " "
	bts, err := json.MarshalIndent(intf, firstColLeftMostPrefix, "\t")
	if err != nil {
		log.Printf("Could not marshal file %v - %v", fileName, err)
		return err
	}
	err = WriteFile(fileName, bytes.NewReader(bts), 0644)
	if err != nil {
		log.Printf("Could not save file %v - %v", fileName, err)
		return err
	}
	// log.Printf("File saved: %v ", fileName)
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
			err = combiErr{err, errSec}
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
	// log.Printf("File %v successfully deleted from bucket", fileName)
	return nil
}

//
//
// ReadDir preparation
var beforeList = func(as func(interface{}) bool) error {
	var q *storage.Query
	if as(&q) { // access storage.Query via q here.
		// log.Printf("beforeFunc(): delim - pref - versions: %v %v %#v", q.Delimiter, q.Prefix, q.Versions)
	} else {
		log.Printf("beforeFunc(): no response to %T", as)
	}
	return nil
}

var list func(context.Context, *blob.Bucket, string, int, int, *[]*blob.ListObject) //

func init() {
	/*
		list() lists files in buck starting with prefix.
		list() recurses into "directories" delimited by "/",
		Strongly consistent, but not guaranteed to work on all services.

		Usage
		ret := &[]*blob.ListObject{}
		list(ctx, buck, "", 0,0, ret)
	*/
	list = func(ctx context.Context, buck *blob.Bucket, prefix string, indent, maxIndent int, results *[]*blob.ListObject) {
		iter := buck.List(
			&blob.ListOptions{
				Delimiter:  "/",
				Prefix:     prefix,
				BeforeList: beforeList,
			})
		for {
			obj, err := iter.Next(ctx)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error iterating with Next(): %v", err)
				break
			}
			if obj.Key == prefix && obj.IsDir {
				continue // skip the directory itself -  occurs only on cloud - not on filesystem
			}
			// fmt.Printf("%s%s\n", strings.Repeat("    ", indent), obj.Key)
			*results = append(*results, obj)
			if obj.IsDir && indent < maxIndent {
				list(ctx, buck, obj.Key, indent+1, maxIndent, results)
			}
		}
	}
}

// ReadDir is similar to os.ReadDir but returns ListObjects instead of FileInfos;
// prefix is the path to search into;
// returned keys will nevertheless consist of  path . path.Separator . fileName;
// under windows we might have to
//     o.Key = strings.ReplaceAll(o.Key, "\\", "/")
//
// On windows,   `prefix` directory itself is not returned
// On appengine, `prefix` directory itself is returned as well
func ReadDir(prefix string) (*[]*blob.ListObject, error) {
	ctx := context.Background()
	var buck *blob.Bucket
	var errSec error

	// ret := []os.FileInfo{}
	ret := &[]*blob.ListObject{}

	// necessary, but entails returning the prefix dir itself
	// see: skip the directory itself
	if !strings.HasSuffix(prefix, "/") {
		prefix += "/"
	}

	buck, err := bucket()
	if err != nil {
		log.Printf("Error opening bucket for file deletion: %v", err)
		return ret, err
	}
	defer func() {
		errSec = buck.Close()
		if errSec != nil {
			err = combiErr{err, errSec}
			log.Printf("Error closing bucket for file deletion: %v", errSec)
		}
	}()

	list(ctx, buck, prefix, 0, 0, ret)
	return ret, nil

}
