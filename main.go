// +build !appengine

package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/http/pprof"
	"os"
	"time"

	"github.com/zew/go-questionnaire/bootstrap"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/generators"
	"github.com/zew/go-questionnaire/handlers"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/go-questionnaire/wrap"
	"golang.org/x/crypto/acme/autocert"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
	log.SetFlags(log.Lshortfile | log.Ltime)

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
	mux1.HandleFunc(cfg.PrefTS("/css/"), tpl.ServeDynCSS)
	// markdown files in /doc
	tpl.CreateAndRegisterHandlerForDocs(mux1)

	serveIcon := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
		// but does not help
		w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*24))
		bts, _ := ioutil.ReadFile("./static/img/ui/favicon.ico")
		fmt.Fprint(w, bts)
	}
	mux1.HandleFunc("/favicon.ico", serveIcon)
	if cfg.Pref() != "" {
		mux1.HandleFunc(cfg.Pref("favicon.ico"), serveIcon)
		mux1.HandleFunc(cfg.PrefTS("favicon.ico"), serveIcon)
	}

	// Login primitives for everybody
	mux1.HandleFunc(cfg.Pref("/login-primitive"), lgn.LoginPrimitiveH)
	mux1.HandleFunc(cfg.Pref("/change-password-primitive"), lgn.ChangePasswordPrimitiveH)
	mux1.HandleFunc(cfg.PrefTS("/d"), handlers.LoginByHashID) // 'd' for direct
	// Administrative handlers - common
	mux1.Handle(cfg.Pref("/session-put"), wrap.AdminFunc(sessx.SessionPut))
	mux1.Handle(cfg.Pref("/session-get"), wrap.AdminFunc(sessx.SessionGet))
	mux1.Handle(cfg.Pref("/config-reload"), wrap.AdminFunc(handlers.ConfigReloadH))
	mux1.Handle(cfg.Pref("/templates-reload"), wrap.AdminFunc(tpl.ParseH))
	// Workflow - logins for survey
	mux1.Handle(cfg.Pref("/generate-questionnaire-templates"), wrap.AdminFunc(generators.SurveyGenerate))
	mux1.Handle(cfg.Pref("/generate-hashes"), wrap.AdminFunc(lgn.GenerateHashesH))
	mux1.Handle(cfg.Pref("/generate-hash-ids"), wrap.AdminFunc(lgn.GenerateHashIDs))
	mux1.Handle(cfg.Pref("/reload-from-questionnaire-template"), wrap.AdminFunc(lgn.ReloadH))
	mux1.Handle(cfg.PrefTS("/reload-from-questionnaire-template"), wrap.AdminFunc(lgn.ReloadH))
	mux1.Handle(cfg.Pref("/shufflings-to-csv"), wrap.AdminFunc(lgn.ShufflesToCSV))
	// Rare login funcs
	mux1.Handle(cfg.Pref("/logins-save"), wrap.AdminFunc(lgn.SaveH))
	mux1.Handle(cfg.Pref("/logins-reload"), wrap.AdminFunc(lgn.LoadH))
	mux1.Handle(cfg.Pref("/generate-password"), wrap.AdminFunc(lgn.GeneratePasswordH))
	mux1.HandleFunc(cfg.Pref("/create-anonymous-id"), lgn.CreateAnonymousID)
	// PProf stuff
	mux1.Handle(cfg.Pref("/diag/pprof"), wrap.AdminFunc(pprof.Index))
	mux1.Handle(cfg.Pref("/diag/allocs"), wrap.AdminFunc(pprof.Handler("allocs").ServeHTTP))
	mux1.Handle(cfg.Pref("/diag/block"), wrap.AdminFunc(pprof.Handler("block").ServeHTTP))
	mux1.Handle(cfg.Pref("/diag/cmdline"), wrap.AdminFunc(pprof.Cmdline))
	mux1.Handle(cfg.Pref("/diag/goroutine"), wrap.AdminFunc(pprof.Handler("goroutine").ServeHTTP))
	mux1.Handle(cfg.Pref("/diag/heap"), wrap.AdminFunc(pprof.Handler("heap").ServeHTTP))
	mux1.Handle(cfg.Pref("/diag/mutex"), wrap.AdminFunc(pprof.Handler("mutex").ServeHTTP))
	mux1.Handle(cfg.Pref("/diag/profile"), wrap.AdminFunc(pprof.Profile))
	mux1.Handle(cfg.Pref("/diag/threadcreate"), wrap.AdminFunc(pprof.Handler("threadcreate").ServeHTTP))
	mux1.Handle(cfg.Pref("/diag/trace"), wrap.AdminFunc(pprof.Trace))
	mux1.Handle(cfg.Pref("/diag/symbol"), wrap.AdminFunc(pprof.Symbol))

	//
	// App specific
	mux1.HandleFunc("/", handlers.MainH)
	if cfg.Pref() != "" {
		mux1.HandleFunc(cfg.Pref("/"), handlers.MainH)
		mux1.HandleFunc(cfg.PrefTS("/"), handlers.MainH)
	}
	mux1.HandleFunc(cfg.Pref("/transferrer-endpoint"), handlers.TransferrerEndpointH)

	/*
		Adding session management.

		The order is counter-intuitive.
		We want requests be handled like this:

		mux3         - establishes sessions
		 	mux2     - recovery, logging
		  		mux1 - call the actual handler

		 To achieve this, we must *reversely* wrap
		 mux1 first in mux2, then in mux3

	*/
	mux2 := wrap.LogAndRecover(mux1)
	// => Wrap in mux2 in session manager
	mux3 := sessx.Mgr().LoadAndSave(mux2)

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
