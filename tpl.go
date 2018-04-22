package main

import (
	"bytes"
	"html/template"
	"log"
	"net/http"

	"github.com/zew/questionaire/cfg"
)

type t struct {
	TemplateName string
	HtmlTitle    string
	Cnt          interface{}
	P            tPage
}

var staticTplFuncs = template.FuncMap{
	"toHtml": func(arg string) template.HTML { return template.HTML(arg) },
	"cfgVal": func(arg string) string { return cfg.Val(arg) },
}

func tpl(r *http.Request) *template.Template {

	var tpl *template.Template

	mp := template.FuncMap{}
	// Add static tpl funcs
	for key, fc := range staticTplFuncs {
		mp[key] = fc
	}
	// Add dynamic tpl funcs
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
	mp["sessionValue"] = func(name string) (ret string, err error) {
		session := sessionManager.Load(r)
		ret, err = session.GetString(name)
		if err != nil {
			log.Printf("sessionValue request %v %v", r.URL.String(), name)
			log.Printf("sessionValue erred:  %v", err)
			return
		}
		return
	}

	var err error
	tplBase := template.New("layout.html")
	tplBase = tplBase.Funcs(mp)

	tpl, err = tplBase.ParseFiles("./templates/layout.html")
	if err != nil {
		log.Fatal(err)
	}

	tpl, err = tpl.ParseGlob("./templates/*.html")
	if err != nil {
		log.Fatal(err)
	}

	return tpl

}
