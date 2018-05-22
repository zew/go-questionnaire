// Package sessx reads effective parameter values
// from get, post and session.
// It also reads consolidated request params (GET, POST).
package sessx

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/zew/util"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/cookiestore" // encrypted cookie
	"github.com/alexedwards/scs/stores/memstore"    // fast, but sessions do not survive server restart
)

var key = "rSRYHsYd2di3PWTDp3fhMTdZCwE5Ne8TxX!" // lgn.GeneratePassword(35)
var sessionManager1 = scs.NewManager(cookiestore.New([]byte(key)))
var sessionManager2 = scs.NewManager(memstore.New(2 * time.Hour))
var sessionManager = sessionManager2

func Mgr() *scs.Manager {
	return sessionManager
}

type SessT struct {
	scs.Session
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request) SessT {
	sess := sessionManager.Load(r)
	return SessT{
		w:       w,
		r:       r,
		Session: *sess,
	}
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
	exists, err := sess.Exists(key)
	util.CheckErr(err)
	if exists {
		return true
	}

	return false

}

// Returns zero value, regardless whether the param was set or not.
func (sess *SessT) EffectiveStr(key string, defaultVal ...string) string {

	// Request
	p, ok := sess.ReqParam(key, defaultVal...)
	if ok {
		return p
	}

	// Session
	// Session was set, but with empty string?
	exists, err := sess.Exists(key)
	util.CheckErr(err)
	if exists {
		p, err := sess.GetString(key)
		util.CheckErr(err)
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
		} else {
			return 0, false, nil
		}
	}

	s := sess.EffectiveStr(key)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0], true, nil
		} else {
			return 0, true, nil
		}
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
		} else {
			return 0.0, false, nil
		}
	}

	s := sess.EffectiveStr(key)
	if s == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0], true, nil
		} else {
			return 0.0, true, nil
		}
	}

	fl, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0, true, err
	}

	return fl, true, nil

}

// Param obj should be pointer
func (sess *SessT) EffectiveObj(key string, obj interface{}) (bool, error) {
	ok := sess.EffectiveIsSet(key)
	if !ok {
		return false, nil
	}
	err := sess.GetObject(key, obj)
	if err != nil {
		return false, err
	}
	return true, nil
}

// Returns zero value, regardless whether the param was set or not.
func (sess *SessT) PutString(key, val string) error {
	err := sess.Session.PutString(sess.w, key, val)
	if err != nil {
		log.Printf("Put string for session session-key %v failed: %v", key, err)
	}
	return err
}

func (sess *SessT) PutObject(key string, val interface{}) error {
	err := sess.Session.PutObject(sess.w, key, val)
	if err != nil {
		log.Printf("Put object for session session-key %v failed: %v", key, err)
	}
	return err
}

//
// Some request handlers for diagnosis
func SessionPut(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	sess := New(w, r)
	sess.PutString("session-test-key", "session-test-value")
	w.Write([]byte("session[session-test-key] set to session-test-value"))
}

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
	keys, _ := sess.Keys()
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
