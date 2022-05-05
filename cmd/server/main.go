package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/zew/go-questionnaire/pkg/bootstrap"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/handlers"
	"github.com/zew/go-questionnaire/pkg/sessx"
	"github.com/zew/go-questionnaire/pkg/tpl"
	"github.com/zew/go-questionnaire/pkg/wrap"
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

	/*
		Http server and mux

		The order of the muxes is counter-intuitive:

		mux4             - static requests, streaming - no session
			mux3         - establishes sessions
				mux2     - recovery, logging
					mux1 - dynamic requests with session

		To achieve this, we must *reversely*
		wrap mux1 with mux2
		wrap mux2 with mux3
		wrap mux3 with mux4

	*/
	tpl.NewDocServer("/doc/")  // before RegisterHandlers() - needs session to differentiate files by language setting and survey name
	mux1 := http.NewServeMux() // base router
	handlers.RegisterHandlers(mux1)
	// dbg.Dump(handler.Tree())
	// handler.Tree().HTML(io.Discard)

	if cfg.Pref() != "" {
		mux1.HandleFunc("/", handlers.MainH) // also register for document root
	}

	mux2 := wrap.LogAndRecover(mux1)      // wrap mux1 into logger/recoverer
	mux3 := sessx.Mgr().LoadAndSave(mux2) // wrap mux2 into session manager

	mux4 := http.NewServeMux() // top router for non-middlewared handlers
	mux4.Handle("/", mux3)
	if cfg.Pref() != "" {
		mux4.Handle(cfg.Pref("/"), mux3) // wrap mux3 into mux4
		mux4.Handle(cfg.PrefTS("/"), mux3)
	}
	// mux4.HandleFunc(cfg.Pref("/slow-buffered"), stream.SlowBuffered)
	// mux4.HandleFunc(cfg.Pref("/slow-hijacked"), stream.SlowHijacked)

	// Static file serving to the base router.
	// Static requests will also trigger the middleware funcs below.
	staticDirs := []string{"/img", "/js"}
	for _, v := range staticDirs {
		mux4.HandleFunc(cfg.Pref(v), tpl.StaticDownloadH)
		mux4.HandleFunc(cfg.PrefTS(v), tpl.StaticDownloadH)
		log.Printf("static service %-20v => /static/[stripped:%v]%v", cfg.Pref(v), cfg.Get().AppMnemonic, v)
	}
	// Extra handler for dynamic css - served from templates
	mux4.HandleFunc(cfg.PrefTS("/css/"), tpl.ServeDynCSS)

	serveIcon := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/x-icon")
		// andrewlock.net/adding-cache-control-headers-to-static-files-in-asp-net.core/
		// but does not help
		w.Header().Set("Cache-Control", fmt.Sprintf("public,max-age=%d", 60*60*24))
		bts, _ := os.ReadFile("./static/img/ui/favicon.ico")
		fmt.Fprint(w, bts)
	}
	mux4.HandleFunc("/favicon.ico", serveIcon)
	if cfg.Pref() != "" {
		mux4.HandleFunc(cfg.Pref("favicon.ico"), serveIcon)
		mux4.HandleFunc(cfg.PrefTS("favicon.ico"), serveIcon)
	}

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
			// upper limits for individual request timeouts - see wrap.LogAndRecover()
			ReadTimeout:       time.Duration(cfg.Get().ReadTimeOut) * time.Second,
			ReadHeaderTimeout: time.Duration(cfg.Get().ReadHeaderTimeOut) * time.Second,
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

		// Valsorda config from 2018  - compatible to sparkassen
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
		if false {
			// Recommendation from 2022 - causes sparkassen to get "no common cipher suit"
			// Sparkassen using www.f-i.de report trouble
			tlsCfg = &tls.Config{
				MinVersion:               tls.VersionTLS12,
				CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
				PreferServerCipherSuites: true,
				CipherSuites: []uint16{
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
					tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
					tls.TLS_RSA_WITH_AES_256_CBC_SHA,
				},
			}

		}
		if !cfg.Get().TLS13 {
			// Best disabled, as they don't provide Forward Secrecy,
			// but might be necessary for some clients, i.e. Internet Explorer 11
			tlsCfg.CipherSuites = append(tlsCfg.CipherSuites, tls.TLS_RSA_WITH_AES_256_GCM_SHA384)
			tlsCfg.CipherSuites = append(tlsCfg.CipherSuites, tls.TLS_RSA_WITH_AES_128_GCM_SHA256)
		}

		// err = http.ListenAndServeTLS(IPPort, "server.pem", "server.key", mux4)
		srv := &http.Server{
			// upper limits for individual request timeouts - see wrap.LogAndRecover()
			ReadTimeout:       time.Duration(cfg.Get().ReadTimeOut) * time.Second,
			ReadHeaderTimeout: time.Duration(cfg.Get().ReadHeaderTimeOut) * time.Second,
			WriteTimeout:      time.Duration(cfg.Get().WriteTimeOut) * time.Second,
			IdleTimeout:       120 * time.Second,
			Addr:              IPPort,
			TLSConfig:         tlsCfg,
			TLSNextProto:      make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
			Handler:           mux4,
		}
		if cfg.Get().LetsEncrypt {
			tlsCfg.GetCertificate = certManager.GetCertificate
			// "", "" => empty key and cert files; key+cert come from Let's Encrypt
			log.Fatal(srv.ListenAndServeTLS("", ""))
		} else {
			// chrome://flags/#allow-insecure-localhost
			pthPem := path.Join("static", "certs", "server.pem")
			pthKey := path.Join("static", "certs", "server.key")
			log.Fatal(srv.ListenAndServeTLS(pthPem, pthKey))

		}
	} else {
		log.Fatal(http.ListenAndServe(IPPort, mux4))
	}

}
