package wrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/lgn"
)

type mustLogin struct {
	inner     http.Handler
	mustAdmin bool
}

// MustLogin takes a handler *func* and returns a wrapped around http handler.
func MustLogin(innerFunc http.HandlerFunc) http.Handler {
	return &mustLogin{
		inner:     innerFunc,
		mustAdmin: false,
	}
}

// MustAdmin takes a handler *func* and returns a wrapped around http handler.
func MustAdmin(innerFunc http.HandlerFunc) http.Handler {
	return &mustLogin{
		inner:     innerFunc,
		mustAdmin: true,
	}
}

// MustAdminHandler takes a handler and returns a wrapped around http handler.
func MustAdminHandler(innerHandler http.Handler) http.Handler {
	return &mustLogin{
		inner:     innerHandler,
		mustAdmin: true,
	}
}

// MustAdmin upgrades requirements from
// just a logged in user to one with admin rights.
// Unused, since we dont expose mustLogin
func (wr *mustLogin) MustAdmin(must bool) {
	wr.mustAdmin = must
}

// Implementing http.Handler interface
func (wr *mustLogin) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	var (
		l          *lgn.LoginT
		isLoggedIn bool
		err        error
	)

	//
	// Check access rights
	if cfg.Get().IsProduction {
		l, isLoggedIn, err = lgn.LoggedInCheck(w, req)
		if err != nil {
			fmt.Fprintf(w, "Login error %v\n", err)
			return
		}
		if !isLoggedIn {
			fmt.Fprintf(w, "Not logged in\n")
			return
		}
		if wr.mustAdmin {
			if !l.HasRole("admin") {
				fmt.Fprintf(w, "Login found, but must have role 'admin' - and no init password\n")
				return
			}
		}
	}

	//
	ctx := context.WithValue(req.Context(), ctxKey("login_middleware"), l) // unused
	wr.inner.ServeHTTP(w, req.WithContext(ctx))

}

//
// UseAdmin is an alternative way to create the same middle ware.
func UseAdmin(inner http.Handler, aParam int) http.Handler {
	// Possible stuff outside the closure
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(context.Background(), ctxKey("param"), aParam)
		inner.ServeHTTP(w, r.WithContext(ctx))
	})
}
