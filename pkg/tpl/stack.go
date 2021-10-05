package tpl

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/zew/util"
)

// Struct2Map converts any struct instance into map[string]interface{};
// previously we used extensions of struct type BaseTemplateData
// to guarantee certain keys in templates;
// recently we add these minimum keys in Exec();
// this func now allows usage of tpl.RenderStack() or tpl.Exec() with
// a struct instead of a 	map[string]interface{}
func Struct2Map(source interface{}) map[string]interface{} {

	v := reflect.ValueOf(source)
	t := reflect.TypeOf(source)

	if v.Type() != t {
		panic(fmt.Sprintf("%v should be equal to %v ?", v.Type(), t))
	}

	// dereference pointer to value
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// only accept struct
	if v.Kind() != reflect.Struct {
		// maybe our argument already *has* desired type map[string]interface{}
		if v.Kind() == reflect.Map {
			source2, ok := source.(map[string]interface{})
			if ok {
				// no need for any processing - return
				return source2
			}
		}
		log.Panicf("Struct2Map: source has type %T; could not json.Marshal %#v", source, source)
	}

	//
	dest := map[string]interface{}{}
	//

	if false {
		for i := 0; i < v.NumField(); i++ {
			ft := t.Field(i)
			dest[ft.Name] = v.Field(i).Interface()
		}
	}

	if true {
		for i := 0; i < v.NumField(); i++ {
			fv := v.Field(i)
			ft := t.Field(i)
			switch {
			// case ft.Anonymous:
			//     embedded *scalar* type -> treat as default
			//     stackoverflow.com/questions/24333494/
			case ft.Anonymous && fv.Kind() == reflect.Struct:
				// log.Printf("    recursing deeper for  %v", ft.Name)
				// dest["a_"+ft.Name] = v.Field(i).Interface()
				{
					source2 := v.Field(i).Interface()
					v := reflect.ValueOf(source2)
					t := reflect.TypeOf(source2)
					for i := 0; i < v.NumField(); i++ {
						fv := v.Field(i)
						ft := t.Field(i)
						if ft.Anonymous && fv.Kind() == reflect.Struct {
							// recurse deeper
						}
						dest[ft.Name] = v.Field(i).Interface()
					}
				}

			default:
				dest[ft.Name] = v.Field(i).Interface()
			}
		}
	}

	if false {
		log.Printf("source %T", source)
		for key, val := range dest {
			valStr := util.UpTo(fmt.Sprintf("%+v", val), 48)
			valStr = strings.ReplaceAll(valStr, "\n", " - ")
			log.Printf("  %-18v -> %-26T %v", key, val, valStr)
		}
	}

	return dest
}

type hieraT struct {
	chain   []string               // a chain of template names - from outermost to innermost
	binding map[string]interface{} // the data, with which the templates should be executed with
	req     *http.Request
}

func (hr *hieraT) step(tplName string) {
	wTemp := &strings.Builder{}
	// todo - tplName inside bundle
	Exec(wTemp, hr.req, hr.binding, tplName)
	hr.binding["Content"] = wTemp.String()
}

// RenderStack renders a chain of templates
func RenderStack(req *http.Request, w io.Writer, chain []string, binding map[string]interface{}) {

	if len(chain) < 1 {
		log.Panic("chain needs at least one template name")
	}

	hr := hieraT{}
	hr.req = req
	hr.chain = chain
	hr.binding = binding

	for i := 0; i < len(chain); i++ {
		idxDesc := len(hr.chain) - 1 - i
		hr.step(hr.chain[idxDesc])
	}

	fmt.Fprint(w, hr.binding["Content"])

}
