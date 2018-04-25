package cfg

import (
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"sync"

	"golang.org/x/crypto/md4"

	"github.com/zew/util"
)

type configT struct {
	sync.Mutex

	IsProduction bool `json:"is_production"` // true => templates are not recompiled

	AppName       string `json:"app_name"`
	UrlPathPrefix string `json:"urlpath_prefix"`

	Vals map[string]string `json:"vals"`

	Css map[string]string `json:"css"`

	HttpReadTimeOut  int `json:"http_read_time_out"`  // for large requests
	HttpWriteTimeOut int `json:"http_write_time_out"` // for *sending* large files over slow networks, i.e. ula's videos, set to 30 or 60 secs

	MaxDownloadRequestsPerMin int `json:"max_download_requests_per_min"`
}

var c = &configT{}

func Get() *configT {
	// I am not sure, whether we need locks here.
	// Since in Load(), we simply exchange one pointer by another at the end of loading
	// c.Lock()
	// defer c.Unlock()
	return c
}
func Val(s string) string {
	c.Lock()
	defer c.Unlock()
	return c.Vals[s]
}

func init() {
	Load()
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
	tempCfg := &configT{}
	err = decoder.Decode(tempCfg)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("\n%#s", util.IndentedDump(tempCfg))
	c = tempCfg
}

//

func LoadH(w http.ResponseWriter, r *http.Request) {
	Load()
}

func Md4Str(buf []byte) string {
	ctx := md4.New()
	ctx.Write(buf)
	byteSliceHash := ctx.Sum(nil)
	return hex.EncodeToString(byteSliceHash)
}
