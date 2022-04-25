package tf

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/pbberlin/flags"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/lgn"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/util"
)

func getClient() *http.Client {
	client := util.HttpClient()
	client = &http.Client{}
	log.Printf("client timeout is %v", client.Timeout)
	return client
}

// RemoteConnConfigT is on top of
// of the ordinary config
type RemoteConnConfigT struct {
	RemoteHost    string
	BindSocket    string
	URLPathPrefix string

	AdminLogin string // Some admin account of the remote machine
	Pass       string

	SurveyType string
	WaveID     string // special value "current" is evaluated to current year

	DownloadDir string
	MinUserID   int // constrain range of UserIDs being processed, exclude test user data entry
	MaxUserID   int // see MinUserID
}

// Example returns a minimal configuration, to be extended or adapted
func Example() RemoteConnConfigT {
	r := RemoteConnConfigT{}
	r.RemoteHost = "https://www.peu2018.eu"
	r.RemoteHost = "https://financial-literacy-test.appspot.com"
	r.RemoteHost = "https://survey2.zew.de"

	r.BindSocket = "443"
	r.URLPathPrefix = "survey"
	r.URLPathPrefix = ""

	r.AdminLogin = "transferrer"
	r.Pass = "Spark!sh32"

	r.SurveyType = "fmt"
	r.SurveyType = "flit"
	r.SurveyType = "lt2020"

	r.WaveID = qst.NewSurvey(r.SurveyType).WaveID()
	r.WaveID = "2020-05"

	r.DownloadDir = "responses/downloaded"

	return r
}

// LoadRemote reads from an io.Reader
// to avoid cyclical deps.
func LoadRemote(r io.Reader) *RemoteConnConfigT {

	log.Printf("Loading remote config...")

	decoder := json.NewDecoder(r)
	tempCfg := RemoteConnConfigT{}
	err := decoder.Decode(&tempCfg)
	if err != nil {
		log.Fatalf("error decoding into RemoteConnConfigT: %v", err)
	}

	if tempCfg.WaveID == "current" {
		tNow := time.Now()
		tempCfg.WaveID = fmt.Sprintf("%v%02d", tNow.Year(), int(tNow.Month()))
	}

	return &tempCfg
}

// ConfigsMainApp loads main app config and main app logins
func ConfigsMainApp() {

	// we need config and logins
	// for main app at least initialized
	{
		//
		// We take a config;
		// save it to file and then activate it from file.
		cf := &cfg.ConfigT{}
		cf.AppName = "Transferrer for Go Questionnaire - http client"
		cf.AppMnemonic = "tf"
		cf.LangCodes = []string{"en"}
		cf.Loc = time.FixedZone("UTC", 1*3600) // cf.Loc is needed below
		// cf.URLPathPrefix is needed for cfg.Pref() properly working
		// It is set later from transferrer config

		pthAutogen := path.Join("/transferrer", "config-autogen.json")
		cloudio.MarshalWriteFile(&cf, pthAutogen)
		cfg.CfgPath = pthAutogen

		fileName := cfg.CfgPath
		r, bucketClose, err := cloudio.Open(fileName)
		if err != nil {
			log.Fatalf("Error opening writer to %v: %v", fileName, err)
		}
		defer func() {
			if r != nil {
				err := r.Close()
				if err != nil {
					log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
				}
			}
		}()
		defer func() {
			err := bucketClose()
			if err != nil {
				log.Printf("Error closing bucket of writer to %v: %v", fileName, err)
			}
		}()
		log.Printf("Opened reader to cloud config %v", fileName)
		cfg.Load(r)
	}

	//
	//
	// logins data is directly read from file;
	// it only contains the remote salt
	// required to create form request tokens
	lgn.LgnsPath = path.Join("/transferrer", "logins-remote-salt.json")
	{
		fileName := lgn.LgnsPath
		r, bucketClose, err := cloudio.Open(fileName)
		if err != nil {
			log.Fatalf("Error opening writer to %v: %v", fileName, err)
		}
		defer func() {
			if r != nil {
				err := r.Close()
				if err != nil {
					log.Printf("Error closing writer to bucket to %v: %v", fileName, err)
				}
			}
		}()
		defer func() {
			err := bucketClose()
			if err != nil {
				log.Printf("Error closing bucket of writer to %v: %v", fileName, err)
			}
		}()
		log.Printf("Opened reader to cloud config %v", fileName)
		lgn.Load(r)

		cloudio.MarshalWriteFile(lgn.Example(), path.Join("/transferrer", "logins-example.json"))

	}

}

// ConfigTransferrer loads the transferrer config;
// using loadRemote();
// TransferrerEndpointH() uses another method
func ConfigTransferrer() *RemoteConnConfigT {

	fl := flags.New()
	fl.Add(
		flags.FlagT{
			Long:       "remote_file",
			Short:      "rmt",
			DefaultVal: path.Join("/transferrer", "remote.json"),
			Desc:       "JSON file containing connection to remote host",
		},
	)
	fl.Gen()
	var cfgRem RemoteConnConfigT
	cfgRem = Example()
	cloudio.MarshalWriteFile(&cfgRem, path.Join("/transferrer", "example-remote.json"))
	{
		rmt := fl.ByKey("rmt").Val
		fileName := rmt
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
		log.Printf("Opened reader to cloud config %v", fileName)
		cfgRem = *(LoadRemote(r))
	}

	// make cfg.Pref() work properly:
	cfg.Get().URLPathPrefix = cfgRem.URLPathPrefix

	return &cfgRem
}
