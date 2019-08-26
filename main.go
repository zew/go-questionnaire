// +build !appengine

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"

	"github.com/zew/go-questionnaire/bootstrap"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/generators"
	"github.com/zew/go-questionnaire/handlers"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/muxwrap"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/tpl"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)

	bootstrap.Config()

	if os.Getenv("GO_TEST_MODE") == "true" {
		cfg.SwitchToTestConfig()
	}

	//
	//
	// Http server
	mux1 := http.NewServeMux() // base router

	// Static file serving to the base router.
	// Static requests will also trigger the middleware funcs below.
	staticDirs := []string{"/img", "/js"}
	for _, v := range staticDirs {
		mux1.HandleFunc(cfg.Pref(v), tpl.StaticDownloadH)
		mux1.HandleFunc(cfg.PrefTS(v), tpl.StaticDownloadH)
		log.Printf("static service %-20v => /static/[stripped:%v]%v", cfg.Pref(v), cfg.Get().AppMnemonic, v)
	}
	// Extra handler for dynamic css - served from templates
	mux1.HandleFunc(cfg.PrefTS("/css/"), tpl.ServeDynCss)
	// markdown files in /doc
	tpl.CreateAndRegisterHandlerForDocs(mux1)

	serveIcon := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		bts, _ := ioutil.ReadFile("./static/img/ui/favicon.ico")
		fmt.Fprint(w, bts)
	}
	mux1.HandleFunc("/favicon.ico", serveIcon)
	if cfg.Pref() != "" {
		mux1.HandleFunc(cfg.Pref("favicon.ico"), serveIcon)
		mux1.HandleFunc(cfg.PrefTS("favicon.ico"), serveIcon)
	}

	//
	// Administrative handlers - common
	mux1.HandleFunc(cfg.Pref("/session-put"), sessx.SessionPut)
	mux1.HandleFunc(cfg.Pref("/session-get"), sessx.SessionGet)
	mux1.HandleFunc(cfg.Pref("/config-reload"), handlers.ConfigReloadH)
	mux1.HandleFunc(cfg.Pref("/templates-reload"), tpl.ParseH)
	// Login primitives
	mux1.HandleFunc(cfg.Pref("/login-primitive"), lgn.LoginPrimitiveH)
	mux1.HandleFunc(cfg.Pref("/change-password-primitive"), lgn.ChangePasswordPrimitiveH)
	mux1.HandleFunc(cfg.PrefTS("/d"), handlers.LoginByHashID) // 'd' for direct
	// Workflow - logins for survey
	mux1.HandleFunc(cfg.Pref("/generate-questionnaire-templates"), generators.SurveyGenerate)
	mux1.HandleFunc(cfg.Pref("/generate-hashes"), lgn.GenerateHashesH)
	mux1.HandleFunc(cfg.Pref("/generate-hash-ids"), lgn.GenerateHashIDs)
	mux1.HandleFunc(cfg.Pref("/reload-from-questionnaire-template"), lgn.ReloadH)
	mux1.HandleFunc(cfg.PrefTS("/reload-from-questionnaire-template"), lgn.ReloadH)
	mux1.HandleFunc(cfg.Pref("/shufflings-to-csv"), lgn.ShufflesToCSV)
	// Rare login funcs
	mux1.HandleFunc(cfg.Pref("/logins-save"), lgn.SaveH)
	mux1.HandleFunc(cfg.Pref("/logins-reload"), lgn.LoadH)
	mux1.HandleFunc(cfg.Pref("/generate-password"), lgn.GeneratePasswordH)
	mux1.HandleFunc(cfg.Pref("/create-anonymous-id"), lgn.CreateAnonymousID)

	//
	// App specific
	mux1.HandleFunc("/", handlers.MainH)
	if cfg.Pref() != "" {
		mux1.HandleFunc(cfg.Pref("/"), handlers.MainH)
		mux1.HandleFunc(cfg.PrefTS("/"), handlers.MainH)
	}
	mux1.HandleFunc(cfg.Pref("/transferrer-endpoint"), handlers.TransferrerEndpointH)

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
	//
	// Special port handling for google appengine
	port := os.Getenv("PORT")
	if port == "" {
		port = fmt.Sprintf("%v", cfg.Get().BindSocket)
		log.Printf("No env variable PORT - defaulting cfg val %s", port)
	}

	IPPort := fmt.Sprintf("%v:%v", cfg.Get().BindHost, port)
	log.Printf("starting http server at %v ... (Forward from %v)", IPPort, cfg.Get().BindSocketFallbackHTTP)
	log.Printf("==========================")
	log.Printf("  ")

	//
	if cfg.Get().TLS {
		// stackoverflow.com/questions/37321760
		certManager := autocert.Manager{
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(cfg.Get().HostName), // Your domain here
			Cache:      autocert.DirCache("certs"),                 // Folder for storing certificates
		}

		fallbackSrv := &http.Server{
			ReadTimeout:       time.Duration(cfg.Get().ReadTimeOut) * time.Second,
			ReadHeaderTimeout: time.Duration(cfg.Get().ReadHeaderTimeOut) * time.Second, // individual request cannot control body timeout
			WriteTimeout:      time.Duration(cfg.Get().WriteTimeOut) * time.Second,
			IdleTimeout:       120 * time.Second,
			Addr:              fmt.Sprintf(":%v", cfg.Get().BindSocketFallbackHTTP),
			Handler: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Set("Connection", "close")
				url := "https://" + req.Host + req.URL.String()
				http.Redirect(w, req, url, http.StatusMovedPermanently)
			}),
		}
		if cfg.Get().LetsEncrypt {
			fallbackSrv.Handler = certManager.HTTPHandler(nil)
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
		if !cfg.Get().TLS13 {
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients, i.e. Internet Explorer 11
			tlsCfg.CipherSuites = append(tlsCfg.CipherSuites, tls.TLS_RSA_WITH_AES_256_GCM_SHA384)
			tlsCfg.CipherSuites = append(tlsCfg.CipherSuites, tls.TLS_RSA_WITH_AES_128_GCM_SHA256)
		}
		if cfg.Get().LetsEncrypt {
			tlsCfg.GetCertificate = certManager.GetCertificate
		}

		// err = http.ListenAndServeTLS(IPPort, "server.pem", "server.key", mux3)
		srv := &http.Server{
			ReadTimeout:       time.Duration(cfg.Get().ReadTimeOut) * time.Second,
			ReadHeaderTimeout: time.Duration(cfg.Get().ReadHeaderTimeOut) * time.Second, // individual request cannot control body timeout
			WriteTimeout:      time.Duration(cfg.Get().WriteTimeOut) * time.Second,
			IdleTimeout:       120 * time.Second,
			Addr:              IPPort,
			TLSConfig:         tlsCfg,
			TLSNextProto:      make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			Handler:           mux3,
		}
		if cfg.Get().LetsEncrypt {
			log.Fatal(srv.ListenAndServeTLS("", "")) // "", "" => empty key and cert files; key+cert come from Let's Encrypt
		} else {
			log.Fatal(srv.ListenAndServeTLS("server.pem", "server.key"))
		}
	} else {
		log.Fatal(http.ListenAndServe(IPPort, mux3))
	}

}
