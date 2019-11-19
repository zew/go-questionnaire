// Package sessx reads effective parameter values
// from GET, POST and SESSION.
// It also reads consolidated request params (GET, POST).
// Environment variable REDIS_SESSION_STORE switches to redis storage.
package sessx

import (
	"context"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/alexedwards/scs/redisstore"
	"github.com/gomodule/redigo/redis"

	"github.com/alexedwards/scs/v2"
	"github.com/alexedwards/scs/v2/memstore"
)

var sessionManager = scs.New()

func init() {

	//
	addr := os.Getenv("REDIS_SESSION_STORE")
	if addr != "" {
		pool := &redis.Pool{
			MaxIdle: 10,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", addr)
			},
		}
		sessionManager.Store = redisstore.New(pool)
		log.Printf("session storage: redis")

		d := net.Dialer{Timeout: 4 * time.Second}
		_, err := d.Dial("tcp", addr)
		if err != nil {
			log.Printf("could not reach redis server %v - %v", addr, err)
			sessionManager.Store = memstore.New() // fallback to memory
		}
	}

	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 2 * time.Hour

}

// Mgr exposes the session manager
func Mgr() *scs.SessionManager {
	return sessionManager
}

// SessT enhances the alexedwards/scs session.
type SessT struct {
	*scs.SessionManager
	ctx context.Context
	// w   io.Writer // no longer required
	r *http.Request
}

// New returns a new enhanced session variable.
func New(w io.Writer, r *http.Request) SessT {
	return SessT{
		SessionManager: sessionManager, // I cannot see the problem with this linter msg here :(
		ctx:            r.Context(),
		// w:              w,
		r: r,
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
	exists := sess.SessionManager.Exists(sess.ctx, key)
	if exists {
		return true
	}

	return false

}

// EffectiveStr returns the corresponding value from request or session.
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
	exists := sess.SessionManager.Exists(sess.ctx, key)
	if exists {
		p = sess.SessionManager.GetString(sess.ctx, key)
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
func (sess *SessT) EffectiveObj(key string) (obj interface{}, ok bool) {
	ok = sess.EffectiveIsSet(key)
	if !ok {
		return
	}
	obj = sess.SessionManager.Get(sess.ctx, key)
	return
}

// PutString stores a string into the session.
// Almost identical to PutObject.
func (sess *SessT) PutString(key, val string) {
	sess.SessionManager.Put(sess.ctx, key, val)
}

// PutObject stores an object into the session.
// Almost identical to PutString.
// val can be pointer or value.
func (sess *SessT) PutObject(key string, val interface{}) {
	sess.SessionManager.Put(sess.ctx, key, val)
}

//
//
//
type testObject struct {
	Name  string
	Birth time.Time
}

func (to testObject) String() string {
	return fmt.Sprintf("Name %v, Birth %v", to.Name, to.Birth.Format(time.RFC822))
}

func init() {
	gob.Register(testObject{})
}

// SessionPut - request handler for session diagnosis via http
func SessionPut(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	sess := New(w, r)

	sess.PutString("key1", "valStr1")
	fmt.Fprint(w, "session[key1] set to valStr1\n")

	testObj := testObject{"Horst", time.Date(1968, 6, 15, 10, 33, 0, 0, time.UTC)}
	sess.PutObject("key2", &testObj)
	fmt.Fprintf(w, "session[key2] set to %s\n", testObj)

}

// SessionGet - request handler for session diagnosis via http
func SessionGet(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	sess := New(w, r)

	{
		val := sess.EffectiveStr("key1")
		cnt := fmt.Sprintf("key1   is %v\n", val)
		fmt.Fprint(w, cnt)
	}

	{
		val := sess.EffectiveStr("reqKey")
		cnt := fmt.Sprintf("reqKey is %v\n", val)
		fmt.Fprint(w, cnt)
	}

	{
		testObjIntf, ok := sess.EffectiveObj("key2")
		cnt := fmt.Sprintf("key2   is %v - %s - %T %#v\n", ok, testObjIntf, testObjIntf, testObjIntf)
		fmt.Fprint(w, cnt)

		testObj, ok := testObjIntf.(testObject)
		cnt = fmt.Sprintf("conversio %v - %s - %T\n", ok, testObj, testObj)
		fmt.Fprint(w, cnt)
	}

	fmt.Fprint(w, "\n\n")
	keys := sess.Keys(sess.ctx)
	for _, key := range keys {
		dis := fmt.Sprintf("key %-12v is set", key)
		// Beware - since the vals are typed differently
		func() {
			if rec := recover(); rec != nil {
				fmt.Fprintf(w, "Error: %v", rec)
			}
			val, _ := sess.EffectiveObj(key)
			objToStr := fmt.Sprintf("%#v", val)
			if len(objToStr) > 140 {
				objToStr = objToStr[:140] + "..."
			}
			dis += fmt.Sprintf("; type %-18T - %v\n", val, objToStr)
		}()
		fmt.Fprint(w, dis)
	}
}
