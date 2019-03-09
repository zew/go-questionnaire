// Package sessx reads effective parameter values
// from GET, POST and SESSION.
// It also reads consolidated request params (GET, POST).
// Alex Edwards has changed the API three times now.
package sessx

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.etcd.io/bbolt"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/boltstore"
	"github.com/alexedwards/scs/memstore"
)

var sessionManager = scs.NewSession()

func init() {

	//
	var store1 scs.Store
	if false {
		var db *bbolt.DB
		store1 = boltstore.NewWithCleanupInterval(db, 20*time.Second)
	}

	store2 := memstore.New()
	//
	// defer store2.StopCleanup()
	sessionManager.Store = store1
	sessionManager.Store = store2

	if false {
		// These are examples
		// Values should be set in main()
		sessionManager.Lifetime = 3 * time.Hour
		sessionManager.IdleTimeout = 20 * time.Minute
		sessionManager.Cookie.HttpOnly = true
		sessionManager.Cookie.Domain = "example.com" // default is ""
		sessionManager.Cookie.Path = "/example/"     // default is "/"
		sessionManager.Cookie.Persist = true
		sessionManager.Cookie.SameSite = http.SameSiteStrictMode
		// sessionManager.Cookie.Secure = true
	}

}

// Mgr exposes the session manager
func Mgr() *scs.Session {
	return sessionManager
}

// SessT enhances the alexedwards/scs session.
type SessT struct {
	scs.Session
	w http.ResponseWriter
	r *http.Request
}

// New returns a new enhanced session variable.
func New(w http.ResponseWriter, r *http.Request) SessT {
	// sess := sessionManager.Load(r)
	return SessT{
		w:       w,
		r:       r,
		Session: *sessionManager,
	}
}

// Clear removes all keys
func (sess *SessT) Clear() {
	sess.Session.Destroy(sess.r.Context())
}

// Remove removes a specific key
func (sess *SessT) Remove(key string) {
	sess.Session.Remove(sess.r.Context(), key)
}

// EffectiveIsSet checks, whether a key is set.
// First inside the current request via RequestParamIsSet()
// Then inside the session.
//
// RequestParamIsSet returns the param value as string.
// But EffectiveIsSet refers to different types in session:
// integers, floats or objects.
//
//
// If ParamPersisterMiddleWare is in action,
// then a few designated session params are always set.
func (sess *SessT) EffectiveIsSet(key string) bool {

	_, ok := sess.ReqParam(key)
	if ok {
		return true
	}

	// Session was set, but with empty string?
	exists := sess.Exists(sess.r.Context(), key)
	// log.Printf("Checking   key  %-14v - exists %v", key, exists)
	if exists {
		return true
	}
	return false

}

// EffectiveStr returns the corresponding value from request or session .
// It returns the zero value "", regardless whether the key was not set at all,
// or whether key was set to value "".
func (sess *SessT) EffectiveStr(key string, defaultVal ...string) string {

	// Request
	p, ok := sess.ReqParam(key, defaultVal...)
	if ok {
		return p
	}

	// Session
	// Session was set, but with empty string?
	exists := sess.Exists(sess.r.Context(), key)
	if exists {
		p := sess.GetString(sess.r.Context(), key)
		return p
	}

	// default
	def := ""
	if len(defaultVal) > 0 {
		def = defaultVal[0]
	}
	return def
}

// EffectiveInt is a wrapper around EffectiveStr
// with subsequent parsing into an int
func (sess *SessT) EffectiveInt(key string, defaultVal ...int) (int, bool, error) {

	ok := sess.EffectiveIsSet(key)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], false, nil
		}
		return 0, false, nil
	}

	s := sess.EffectiveStr(key)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0], true, nil
		}
		return 0, true, nil
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, true, err
	}

	return i, true, err

}

// EffectiveFloat is a wrapper around EffectiveStr
// with subsequent parsing into float
func (sess *SessT) EffectiveFloat(key string, defaultVal ...float64) (float64, bool, error) {

	ok := sess.EffectiveIsSet(key)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], false, nil
		}
		return 0.0, false, nil
	}

	s := sess.EffectiveStr(key)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0], true, nil
		}
		return 0.0, true, nil
	}

	fl, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, true, err
	}

	return fl, true, nil

}

// EffectiveObj helps to retrieve an compound variable from the session.
func (sess *SessT) EffectiveObj(key string) (interface{}, bool, error) {
	ok := sess.EffectiveIsSet(key)
	if !ok {
		return nil, false, nil
	}
	obj := sess.Session.Get(sess.r.Context(), key)
	// log.Printf("Object from session: %v %T ", key, obj)
	return obj, true, nil
}

// PutString stores a string into the session.
func (sess *SessT) PutString(key, val string) error {
	sess.Session.Put(sess.r.Context(), key, val)
	return nil
}

// PutObject stores an object into the session.
// Retrieval via sess.get(...)
// Super tricky caveat:
// Same request -      pointers are retrieved as pointers.
// Succinct requests - pointers are retrieved de-referenced.
func (sess *SessT) PutObject(key string, val interface{}) error {
	sess.Session.Put(sess.r.Context(), key, val)
	// log.Printf("Object into session: %v %T ", key, val)
	return nil
}

// SessionPut is a convenience request handler for diagnosis via http
func SessionPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	sess := New(w, r)
	sess.PutString("session-test-key", "session-test-value")
	w.Write([]byte("session[session-test-key] set to session-test-value"))
}

// SessionGet is a convenience request handler for diagnosis via http
func SessionGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	sess := New(w, r)
	val1 := sess.EffectiveStr("session-test-key")
	cnt1 := fmt.Sprintf("session-test-key is %v\n", val1)
	w.Write([]byte(cnt1))
	val2 := sess.EffectiveStr("request-test-key")
	cnt2 := fmt.Sprintf("request-test-key is %v\n", val2)
	w.Write([]byte(cnt2))

	w.Write([]byte("\n\n"))
	keys := sess.Keys(sess.r.Context())
	for _, key := range keys {
		dis := fmt.Sprintf("key %20v is set", key)
		// Beware - since the vals are typed differently
		func() {
			if rec := recover(); rec != nil {
				w.Write([]byte(fmt.Sprintf("Error: %v", rec)))
			}
			val := sess.EffectiveStr(key)
			if len(val) > 80 {
				val = val[:80]
			}
			dis += fmt.Sprintf("; val is %v\n\n", val)
		}()
		w.Write([]byte(dis))
	}
}
