// +build !appengine

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/zew/go-questionaire/bootstrap"
	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/lgn"
	"github.com/zew/go-questionaire/muxwrap"
	"github.com/zew/go-questionaire/qst"
	"github.com/zew/go-questionaire/sessx"
	"github.com/zew/go-questionaire/tpl"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	bootstrap.Config()
	cfg.Example()
	lgn.Example()

	//
	qst.GenerateExample2()

	//
	//
	// Http server
	mux1 := http.NewServeMux() // base router

	// Static file serving to the base router.
	// Static requests will also trigger the middleware funcs below.
	staticDirs := []string{"/img", "/js"}
	for _, v := range staticDirs {
		mux1.HandleFunc(cfg.Pref(v), tpl.StaticDownloadH)
		mux1.HandleFunc(cfg.PrefWTS(v), tpl.StaticDownloadH)
		log.Printf("static service %-20v => /static/[stripped:%v]%v", cfg.Pref(v), cfg.Get().AppMnemonic, v)
	}
	serveIcon := func(w http.ResponseWriter, r *http.Request) {
		bts, _ := ioutil.ReadFile("./static/img/ui/favicon.ico")
		w.Write(bts)
	}
	mux1.HandleFunc("/favicon.ico", serveIcon)
	mux1.HandleFunc(cfg.Pref("favicon.ico"), serveIcon)
	mux1.HandleFunc(cfg.PrefWTS("favicon.ico"), serveIcon)

	//
	// Extra handler for dynamic css
	mux1.HandleFunc(cfg.PrefWTS("/css/"), tpl.ServeDynCss)

	//
	// Administrative handlers
	mux1.HandleFunc(cfg.Pref("/session-put"), sessx.SessionPut)
	mux1.HandleFunc(cfg.Pref("/session-get"), sessx.SessionGet)
	mux1.HandleFunc(cfg.Pref("/config-reload"), cfg.LoadH)
	mux1.HandleFunc(cfg.Pref("/login-primitive"), lgn.LoginPrimitiveH)
	mux1.HandleFunc(cfg.Pref("/logins-save"), lgn.SaveH)
	mux1.HandleFunc(cfg.Pref("/logins-reload"), lgn.LoadH)
	mux1.HandleFunc(cfg.Pref("/generate-password"), lgn.GeneratePasswordH)
	mux1.HandleFunc(cfg.Pref("/generate-hashes"), lgn.GenerateHashesH)
	mux1.HandleFunc(cfg.Pref("/templates-reload"), tpl.ParseH)

	//
	// Standard handlers
	tpl.CreateAndRegisterHandlerForDocs(mux1)
	mux1.HandleFunc("/", mainH)
	mux1.HandleFunc(cfg.Pref("/"), mainH)
	mux1.HandleFunc(cfg.PrefWTS("/"), mainH)
	mux1.HandleFunc(cfg.Pref("/reload-from-file"), reloadH)

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
	mux2 := muxwrap.NewHandlerMiddleware(mux1)
	// => Wrap in mux2 in session manager
	sessx.Mgr().Secure(true)            // true breaks session persistence in excel-db - but not in go-countdown
	sessx.Mgr().Lifetime(2 * time.Hour) // default is 24 hours
	sessx.Mgr().Persist(false)
	mux3 := sessx.Mgr().Use(mux2)

	//
	// Prepare web server launch
	IpPort := fmt.Sprintf("%v:%v", cfg.Get().BindHost, cfg.Get().BindSocket)
	log.Printf("starting http server at %v ... (Forward from %v)", IpPort, cfg.Get().BindSocketFallbackHttp)

	//
	if cfg.Get().Tls {
		fallbackSrv := &http.Server{
			ReadTimeout: time.Duration(cfg.Get().HttpReadTimeOut) * time.Second,
			// ReadHeaderTimeout:  120 * time.Second,  // individual request can control body timeout
			WriteTimeout: time.Duration(cfg.Get().HttpWriteTimeOut) * time.Second,
			IdleTimeout:  120 * time.Second,
			Addr:         fmt.Sprintf(":%v", cfg.Get().BindSocketFallbackHttp),
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
			},
		}
		if !cfg.Get().Tls13 {
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients, i.e. Internet Explorer 11
			tlsCfg.CipherSuites = append(tlsCfg.CipherSuites, tls.TLS_RSA_WITH_AES_256_GCM_SHA384)
			tlsCfg.CipherSuites = append(tlsCfg.CipherSuites, tls.TLS_RSA_WITH_AES_128_GCM_SHA256)
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
		log.Fatal(srv.ListenAndServeTLS("server.pem", "server.key"))
	} else {
		log.Fatal(http.ListenAndServe(IpPort, mux3))
	}

}
