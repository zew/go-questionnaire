// Package handlers combines link data with http handler funcs.
package handlers

import (
	"net/http"
	"net/http/pprof"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/generators"
	"github.com/zew/go-questionnaire/handler"
	"github.com/zew/go-questionnaire/lgn"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/stream"
	"github.com/zew/go-questionnaire/tpl"
	"github.com/zew/go-questionnaire/wrap"
)

// RegisterHandlers cannot go into a global var or into an init() func,
// Since the handler funcs use the links slice
// and create circular dependencies.
func RegisterHandlers(mux *http.ServeMux) {

	infos := handler.InfosT{

		//
		//
		// generic app - generic admin
		{
			Urls:    []string{"/doc"},
			Handler: tpl.ServeDoc,
			Title:   "Documentation", // Serve content under app-bucket
			Keys:    []string{"doc-server", "documentation"},
		},
		{
			Urls: []string{"/login-primitive"},
			// 	Handler: LoginPrimitiveSiteH,
			Handler:  lgn.LoginPrimitiveH,
			Title:    "Login app",
			Keys:     []string{"login-primitive"},
			ShortCut: "l",
			Allow:    map[handler.Privilege]bool{handler.LoggedOut: true},
		},
		{
			Urls:    []string{"/change-password-primitive"},
			Title:   "Change password",
			Handler: lgn.ChangePasswordPrimitiveH,
			// 	Handler: ChangePasswordPrimitiveSiteH,
			Keys:  []string{"change-password-primitive"},
			Allow: map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/create-anonymous-id"},
			Title:   "Anonymous login",
			Handler: lgn.CreateAnonymousIDH,
			Keys:    []string{"create-anonymous-id"},
			Allow:   map[handler.Privilege]bool{handler.LoggedOut: true},
		},
		{
			Urls:    []string{"/logins/reload"},
			Handler: lgn.LoadH,
			Title:   "Logins reload",
			Keys:    []string{"logins-reload"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/logins/generate-password"},
			Title:   "Generate password",
			Handler: lgn.GeneratePasswordH,
			Keys:    []string{"logins-generate-password"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/config/reload"},
			Title:   "Config reload",
			Handler: ConfigReloadH, // must be in this package
			Keys:    []string{"config-reload"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/templates/reload"},
			Title:   "Templates reload",
			Handler: tpl.TemplatesPreparse,
			Keys:    []string{"templates-reload"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/slow-buffered"},
			Handler: stream.SlowBuffered,
			Title:   "Slow buffered",
			Keys:    []string{"slow-buffered"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/slow-hijacked"},
			Handler: stream.SlowHijacked,
			Title:   "Slow hijacked",
			Keys:    []string{"slow-hijacked"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/session-put"},
			Title:   "Session put",
			Handler: sessx.SessionPut,
			Keys:    []string{"session-put"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/session-get"},
			Title:   "Session get",
			Handler: sessx.SessionGet,
			Keys:    []string{"session-get"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/cloud-store-test"},
			Handler: TestCloudStore,
			Title:   "Cloud store test",
			Keys:    []string{"cloud-store-test"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/instance-info"},
			Handler: stream.InstanceInfo,
			Title:   "Instance info",
			Keys:    []string{"instance-info"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},

		// pprof
		{
			Urls:    []string{"/diag/pprof"},
			Title:   "PProf index",
			Handler: pprof.Index,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/allocs"},
			Title:   "PProf allocations",
			Handler: pprof.Handler("allocs").ServeHTTP,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/block"},
			Title:   "PProf block",
			Handler: pprof.Handler("block").ServeHTTP,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/cmdline"},
			Title:   "PProf cmdline",
			Handler: pprof.Cmdline,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/goroutine"},
			Title:   "PProf goroutine",
			Handler: pprof.Handler("goroutine").ServeHTTP,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/heap"},
			Title:   "PProf heap",
			Handler: pprof.Handler("heap").ServeHTTP,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/mutex"},
			Title:   "PProf mutex",
			Handler: pprof.Handler("mutex").ServeHTTP,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/profile"},
			Title:   "PProf profile",
			Handler: pprof.Profile,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/threadcreate"},
			Title:   "PProf thread",
			Handler: pprof.Handler("threadcreate").ServeHTTP,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/trace"},
			Title:   "PProf trace",
			Handler: pprof.Trace,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/diag/symbol"},
			Title:   "PProf symbol",
			Handler: pprof.Symbol,
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},

		// Application
		{
			Urls:     []string{"/", "/home"},
			Handler:  MainH,
			Title:    "Home",
			Keys:     []string{"main", "index"},
			ShortCut: "p",
		},
		{
			Urls:    []string{"/d"}, // 'd' for direct
			Title:   "Login by hash ID",
			Handler: LoginByHashID,
			Keys:    []string{"login-by-hash-id"},
		},
		{
			Urls:    []string{"/logout"},
			Title:   "Logout",
			Handler: LogoutSiteH,
			Keys:    []string{"logout"},
			Allow:   map[handler.Privilege]bool{handler.LoggedIn: true},
		},
		{
			Urls:    []string{"/fmreport-email", "/registrationfmr"}, // second without hyphen - avoid MS word escaping of URL
			Title:   "Registration FM Report email",
			Handler: RegistrationFMRH,
			Keys:    []string{"fmreport-email"},
		},
		{
			Urls:    []string{"/registrationfmt"}, // without hyphen - avoid MS word escaping of URL
			Title:   "Registration FMT",
			Handler: RegistrationFMTH,
			Keys:    []string{"registration-fmt"},
		},
		{
			Urls:    []string{"/doc/site-imprint.md"},
			Handler: tpl.ServeDoc,
			Title:   "Imprint",
			Keys:    []string{"imprint"},
		},

		// Application specific admin
		{
			Urls:    []string{"/generate-questionnaire-templates"},
			Handler: generators.GenerateQuestionnaireTemplates,
			Title:   "Generate Questionnaire Templates",
			Keys:    []string{"generate-questionnaire-templates"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/generate-landtag-variations"},
			Handler: generators.GenerateLandtagsVariations,
			Title:   "Generate Landtag Variations",
			Keys:    []string{"generate-landtag-variations"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/generate-hashes"},
			Handler: lgn.GenerateHashesH,
			Title:   "Generate Hashes",
			Keys:    []string{"generate-hashes"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/generate-hash-ids"},
			Handler: lgn.GenerateHashIDs,
			Title:   "Generate Hash IDs",
			Keys:    []string{"generate-hash-ids"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/reload-from-questionnaire-template"},
			Handler: lgn.ReloadH,
			Title:   "Reload from Questionnaire Template",
			Keys:    []string{"reload-from-questionnaire-template"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/shufflings-to-csv"},
			Handler: lgn.ShufflesToCSV,
			Title:   "Shufflings to CSV",
			Keys:    []string{"shufflings-to-csv"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
		{
			Urls:    []string{"/transferrer-endpoint"},
			Handler: TransferrerEndpointH,
			Title:   "Transferrer server",
			Keys:    []string{"transferrer-endpoint"},
			Allow:   map[handler.Privilege]bool{handler.Admin: true},
		},
	}

	infos.MakeKeys()

	handler.SetInfos(infos)

	// Registration with the mux
	for _, l := range *handler.Infos() {

		for _, url := range l.Urls {
			if l.Allow[handler.Admin] {
				mux.Handle(cfg.Pref(url), wrap.MustAdminHandler(l.Handler))
				mux.Handle(cfg.PrefTS(url), wrap.MustAdminHandler(l.Handler))
				continue
			}
			if l.Allow[handler.LoggedIn] {
				mux.Handle(cfg.Pref(url), wrap.MustLogin(l.Handler))
				mux.Handle(cfg.PrefTS(url), wrap.MustLogin(l.Handler))
				continue
			}
			mux.HandleFunc(cfg.Pref(url), l.Handler)
			mux.HandleFunc(cfg.PrefTS(url), l.Handler) // make sure /p1/p2/  also serves /p1/p2 - see config.Pref()
		}
	}
	// Registration with the mux - finished

}
