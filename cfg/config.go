package cfg

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"sync"
	"time"

	"golang.org/x/crypto/md4"

	"github.com/zew/util"
)

type configT struct {
	sync.Mutex

	AppName       string `json:"app_name"`
	UrlPathPrefix string `json:"urlpath_prefix"`

	Vals map[string]string `json:"vals"`

	Css map[string]string `json:"css"`

	HttpReadTimeOut  int `json:"http_read_time_out"`  // for large requests
	HttpWriteTimeOut int `json:"http_write_time_out"` // for *sending* large files over slow networks, i.e. ula's videos, set to 30 or 60 secs

	MaxDownloadRequestsPerMin int `json:"max_download_requests_per_min"`
}

var c configT

func Get() *configT {
	c.Lock()
	defer c.Unlock()
	return &c
}
func Val(s string) string {
	c.Lock()
	defer c.Unlock()
	return c.Vals[s]
}

func init() {
	Load()
}

func Save() {
	c.Lock()
	defer c.Unlock()
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	firstColLeftMostPrefix := " "
	byts, err := json.MarshalIndent(c, firstColLeftMostPrefix, "\t")
	if err != nil {
		log.Fatal(err)
	}

	pthOld := path.Join(workDir, "config.json")
	pthNew := path.Join(workDir, fmt.Sprintf("config_%v.json", time.Now().Unix()))

	err = os.Rename(pthOld, pthNew)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(pthOld, byts, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Saved config file: %v", pthOld)

}
func Load() {
	c.Lock()
	defer c.Unlock()
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	pth := path.Join(workDir, "config.json")
	file, err := os.Open(pth)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("found config file: %v", pth)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n%#s", util.IndentedDump(c))
}

//

func LoadH(w http.ResponseWriter, r *http.Request) {
	Load()
}

func SaveH(w http.ResponseWriter, r *http.Request) {
	Save()
}

func Md4Str(buf []byte) string {
	ctx := md4.New()
	ctx.Write(buf)
	byteSliceHash := ctx.Sum(nil)
	return hex.EncodeToString(byteSliceHash)
}
