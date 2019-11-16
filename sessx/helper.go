package sessx

// ReqParam searches for the effective value
// of the *request*, not in session.
// First among the POST fields.
// Not  among the URL "path" parameters.
// Then among the URL GET parameters.
//
// It checks, whether any of the above had the param
// key set to *empty* string.
// 
// Second return value is 'is set'
func (sess *SessT) ReqParam(key string, defaultVal ...string) (string, bool) {

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
		return gets.Get(key), true // if there are multiple GET params, this returns the *first* one
	}

	return p, false

}

// ReqURI is a template helper.
// The return value contains app url prefix.
func (sess *SessT) ReqURI() string {
	uri := sess.r.URL.Path
	return uri
}
