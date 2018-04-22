// Package sessx reads effective parameter values
// from get, post and session.
// It also reads consolidated request params (GET, POST).
package sessx

import (
	"log"
	"net/http"
	"strconv"

	"github.com/zew/util"

	"github.com/alexedwards/scs"
)

type TSess struct {
	scs.Session
	w http.ResponseWriter
	r *http.Request
}

func New(w http.ResponseWriter, r *http.Request, mgr *scs.Manager) TSess {
	sess := mgr.Load(r)
	return TSess{
		w:       w,
		r:       r,
		Session: *sess,
	}
}

// EffectiveParamInt is a wrapper around EffectiveParam
// with subsequent parsing into an int
func (sess *TSess) EffectiveParamInt(key string, defaultVal ...int) (int, bool, error) {

	s, ok := sess.EffectiveParamIsSet(key)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], false, nil
		} else {
			return 0, false, nil
		}
	}

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

// EffectiveParamFloat is a wrapper around EffectiveParam
// with subsequent parsing into float
func (sess *TSess) EffectiveParamFloat(key string, defaultVal ...float64) (float64, bool, error) {

	s, ok := sess.EffectiveParamIsSet(key)
	if !ok {
		if len(defaultVal) > 0 {
			return defaultVal[0], false, nil
		} else {
			return 0.0, false, nil
		}
	}

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

// EffectiveParamIsSet searches for the effective value
// of the request.
// First among the POST fields.
// Then among the URL "path" parameters.
// Then among the URL GET parameters.
//
// It checks, whether whether any of the above had the param
// key set to *empty* string.
func (sess *TSess) RequestParamIsSet(key string, defaultVal ...string) (string, bool) {

	p := ""

	// Which to call: r.ParseForm() or r.ParseMultipartForm(1024*1024)
	// https://blog.saush.com/2015/03/18/html-forms-and-go/
	_ = sess.r.PostFormValue("impossibleKey") // hopefully causing the right parsing

	// POST Param overrides GET param
	posts := sess.r.PostForm
	if _, ok := posts[key]; ok {
		return posts.Get(key), true
	}

	// Path Param
	// [deleted]

	// URL Get Param
	gets := sess.r.URL.Query()
	if _, ok := gets[key]; ok {
		return gets.Get(key), true
	}

	return p, false

}

// EffectiveParamIsSet searches for the effective value.
// First among inside the current request via RequestParamIsSet()
// Then inside the session.
//
// It checks, whether session had the param
// key set to *empty* string.
//
// If ParamPersisterMiddleWare is in action,
// then all params are condensed to session level.
// But we usually dont want to persist *all* params,
// and we dont want to register ParamPersisterMiddleWare
// for all routes.
func (sess *TSess) EffectiveParamIsSet(key string, defaultVal ...string) (string, bool) {

	p, ok := sess.RequestParamIsSet(key, defaultVal...)
	if ok {
		return p, true
	}

	// Session
	p, err := sess.GetString(key)
	util.CheckErr(err)
	if p != "" {
		return p, true
	}

	// Session was set, but with empty string?
	exists, err := sess.Exists(key)
	util.CheckErr(err)
	if exists {
		return p, true
	}

	// default
	def := ""
	if len(defaultVal) > 0 {
		def = defaultVal[0]
	}
	// logx.Debugf(r,"!Found & returns the def: %s", def)
	return def, false

}

// Returns zero value, regardless whether the param was set or not.
func (sess *TSess) EffectiveParam(key string, defaultVal ...string) string {
	ret, _ := sess.EffectiveParamIsSet(key, defaultVal...)
	return ret
}

// Returns zero value, regardless whether the param was set or not.
func (sess *TSess) PutString(key, val string) {
	err := sess.Session.PutString(sess.w, key, val)
	if err != nil {
		log.Fatalf("Put session key session-key => session-value failed: %v", err)
	}
	// log.Printf("PutString: %v => %v", key, val)
}
