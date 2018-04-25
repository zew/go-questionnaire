package tpl

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"sync"

	"github.com/alexedwards/scs"

	"github.com/zew/questionaire/cfg"
	"github.com/zew/questionaire/sessx"
)

var staticTplFuncs = template.FuncMap{
	"toHtml": func(arg string) template.HTML { return template.HTML(arg) },
	"cfgVal": func(arg string) string { return cfg.Val(arg) },
	"addint": func(i1, i2 int) int { return i1 + i2 },
	// dummies, to make parsing work
	"executeTemplate": func(name string, data interface{}) (ret template.HTML, err error) { return },
	"sess":            func() (sess sessx.TSess, err error) { return },
}

func StaticFuncMap() template.FuncMap {
	return staticTplFuncs
}

type baseTplT struct {
	IsParsed bool
	sync.Mutex
	*template.Template
}

var bt baseTplT

func (bt *baseTplT) Parse() *template.Template {

	if !bt.IsParsed {

		bt.Lock()
		defer bt.Unlock()

		mp := template.FuncMap{}
		// Add static tpl funcs
		for key, fc := range staticTplFuncs {
			mp[key] = fc
		}

		var err error
		tplBase := template.New("layout.html")
		tplBase = tplBase.Funcs(mp)

		bt.Template, err = tplBase.ParseFiles("./templates/layout.html")
		if err != nil {
			log.Fatal(err)
		}

		bt.Template, err = bt.Template.ParseGlob("./templates/*.html")
		if err != nil {
			log.Fatal(err)
		}

		bt.IsParsed = true
	}

	return bt.Template
}

func Get(w http.ResponseWriter, r *http.Request, sessionManager *scs.Manager) *template.Template {

	if !cfg.Get().IsProduction {
		bt.IsParsed = false
	}

	tpl := bt.Parse()
	tpl, err := tpl.Clone()
	if err != nil {
		log.Fatal(err)
	}

	mp := template.FuncMap{}
	// Add static tpl funcs
	for key, fc := range staticTplFuncs {
		mp[key] = fc
	}
	// Add dynamic tpl funcs
	// Request specific template var
	mp["executeTemplate"] = func(name string, data interface{}) (ret template.HTML, err error) {
		buf := bytes.NewBuffer([]byte{})
		err = tpl.ExecuteTemplate(buf, name, data)
		if err != nil {
			log.Printf("callTemplate erred: %v", err)
			return
		}
		ret = template.HTML(buf.String())
		return
	}
	// response writer + request specific closure
	// Use sess.EffectiveParam("name") => sess.EffectiveParam "session-test-key"
	mp["sess"] = func() (sess *sessx.TSess, err error) {
		sessVal := sessx.New(w, r, sessionManager)
		sess = &sessVal
		return
	}

	tpl = tpl.Funcs(mp)
	return tpl

}
