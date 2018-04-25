package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/zew/questionaire/cfg"
	"github.com/zew/questionaire/muxwrap"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/cookiestore" // encrypted cookie
	"github.com/alexedwards/scs/stores/memstore"    // fast, but sessions do not survive server restart
)

var sessionManager1 = scs.NewManager(cookiestore.New([]byte("u46IpCV9y5Vl332168vODJEhgOY8m9JVE4")))
var sessionManager2 = scs.NewManager(memstore.New(2 * time.Hour))
var sessionManager = sessionManager2

func prefix(a ...string) string {
	pref := cfg.Val("urlPrefix")
	if pref == "" {
		ret := path.Join(a...)
		if !strings.HasSuffix(ret, "/") {
			return ret + "/"
		}
		return ret
	}

	ret := path.Join(a...)
	ret = path.Join(pref, ret)
	return ret + "/"
}

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	//
	generateExample()

	//
	//
	// Http server
	mux1 := http.NewServeMux() // base router
	staticDirs := []string{"/img", "/js"}
	for _, v := range staticDirs {
		mux1.HandleFunc(prefix(v), staticDownloadH)
		log.Printf("static service %-20v => /static/[stripped:%v]%v", prefix(v), cfg.Val("appMnemonic"), v)
	}
	serveIcon := func(w http.ResponseWriter, r *http.Request) {
		bts, _ := ioutil.ReadFile("./static/img/ui/favicon.ico")
		w.Write(bts)
	}
	mux1.HandleFunc("/favicon.ico", serveIcon)
	mux1.HandleFunc(prefix("favicon.ico"), serveIcon)

	//
	// Extra handler for dynamic css
	mux1.HandleFunc(prefix("/css/design.css"), serveCss)

	//
	// Standard handlers
	mux1.HandleFunc(prefix("/"), mainH)
	mux1.HandleFunc(prefix("/config-reload"), cfg.LoadH)
	mux1.HandleFunc(prefix("/session-put"), sessionPut)
	mux1.HandleFunc(prefix("/session-get"), sessionGet)

	//
	// Session manager and session management.
	// The order is counter-intuitive.
	// We want requests be handled like this:
	//
	//  mux3         - establishes sessions
	// 		mux2     - logging - param persisting
	//  		mux1 - call special handler
	//
	// To achieve this, we must *reversely* wrap
	// mux1 first in mux2, then in mux3
	//
	// => Wrap the base router into an unconditional middleware
	mux2 := muxwrap.NewHandlerMiddleware(mux1, sessionManager)
	// => Wrap in mux2 in session manager
	sessionManager.Secure(true)            // true breaks session persistence in exceldb - but not in fmtdownload
	sessionManager.Lifetime(2 * time.Hour) // default is 24 hours
	sessionManager.Persist(false)
	mux3 := sessionManager.Use(mux2)

	//
	// Prepare web server launch
	IpPort := fmt.Sprintf("%v:%v", cfg.Val("BindHost"), cfg.Val("BindSocket"))
	log.Printf("starting http server at %v ... (Forward from %v)", IpPort, cfg.Val("BindSocketFallbackHttp"))

	//
	if cfg.Val("Tls") != "" {
		fallbackSrv := &http.Server{
			ReadTimeout: time.Duration(cfg.Get().HttpReadTimeOut) * time.Second,
			// ReadHeaderTimeout:  120 * time.Second,  // individual request can control body timeout
			WriteTimeout: time.Duration(cfg.Get().HttpWriteTimeOut) * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         cfg.Val("BindSocketFallbackHttp"),
			Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Connection", "close")
				url := "https://" + req.Host + req.URL.String()
				http.Redirect(w, req, url, http.StatusMovedPermanently)
			}),
		}
		go func() { log.Fatal(fallbackSrv.ListenAndServe()) }()

		//
		tlsCfg := &tls.Config{
			// Causes servers to use Go's default ciphersuite preferences,
			// which are tuned to avoid attacks. Does nothing on clients.
			PreferServerCipherSuites: true,
			// Only use curves which have assembly implementations
			CurvePreferences: []tls.CurveID{
				tls.CurveP256,
				tls.X25519, // Go 1.8 only
			},
			MinVersion: tls.VersionTLS12,
			CipherSuites: []uint16{
				tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				// Best disabled, as they don't provide Forward Secrecy, but might be necessary for some clients
				// tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
				// tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			},
		}

		// err = http.ListenAndServeTLS(IpPort, "server.pem", "server.key", mux3)
		srv := &http.Server{
			ReadTimeout:  time.Duration(cfg.Get().HttpReadTimeOut) * time.Second,
			WriteTimeout: time.Duration(cfg.Get().HttpWriteTimeOut) * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         IpPort,
			TLSConfig:    tlsCfg,
			TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			Handler:      mux3,
		}

		// Checking the modulus
		// openssl x509 -noout -modulus -in fmtdownload.pem
		// openssl rsa -check -noout -modulus -in fmtdownload.key
		log.Fatal(srv.ListenAndServeTLS("fmtdownload.pem", "fmtdownload.key"))
	} else {
		log.Fatal(http.ListenAndServe(IpPort, mux3))
	}

}
