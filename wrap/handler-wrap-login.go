package wrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/lgn"
)

type mustAdmin struct {
	inner http.Handler
}

// Admin returns a new http handler.
func Admin(innerHandler http.Handler) http.Handler {
	return &mustAdmin{
		inner: innerHandler,
	}
}

// AdminFunc returns a new http handler.
func AdminFunc(innerFunc http.HandlerFunc) http.Handler {
	return &mustAdmin{
		inner: innerFunc,
	}
}

// Implementing http.Handler interface
func (wr *mustAdmin) ServeHTTP(w http.ResponseWriter, req *http.Request) {

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
		if !l.HasRole("admin") {
			fmt.Fprintf(w, "Login found, but must have role 'admin'\n")
			return
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
