// Package handler contains metadata struct for http handlers;
// urls, titles, privileges and navigation menu info.
// It cannot be part of package handlers,
// since this would cause cyclical dependencies.
package handler

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

// Privilege encodes some
type Privilege int

const (
	// ReadOnly means no saving operation
	ReadOnly Privilege = iota
	// LoggedOut -> only show if no user is logged in
	LoggedOut
	// LoggedIn requires login
	LoggedIn
	// Admin has all rights
	Admin
	// Editor can save values
	Editor
	// ForwardCopyAllTabs can perform mass copy operations
	ForwardCopyAllTabs
)

// Info and Infos keep all the meta information
// about routes, such as Title, Description,
// Url(s), [GET,POST,PUT] and ShowInNavigation.
// They also contain the related handlerFunc.
// Last not least, they are injected into
// the template engine, so that URLs
// are dynamically fetched via {{ index (   byKey "landing-page" ).Urls  0  }}
type Info struct {
	//
	Keys    []string `json:"keys,omitempty"`    // identifier - does not change as URL or title - dynamically derived from the handlerfunc pointer - if not set, then autogenerated from title
	Urls    []string `json:"urls,omitempty"`    // first is idiomatic
	Methods []string `json:"methods,omitempty"` // GET POST ...

	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"` // Meta stuff

	Handler http.HandlerFunc `json:"-"`

	InNav    bool   `json:"in_nav,omitempty"`    // Show in Navigation
	NewCol   bool   `json:"new_col,omitempty"`   // Start new column afterwards
	ShortCut string `json:"short_cut,omitempty"` // accesskey attribute in main navi

	Allow map[Privilege]bool `json:"-"`
}

// InfosT models a collection of all app specific handlers
type InfosT []Info

var infos = InfosT{}

// Infos retunrs all handlers
func Infos() *InfosT {
	return &infos
}

// SetInfos saves arg as package variable infos
func SetInfos(arg InfosT) {
	infos = arg
}

// HasKey checks an info for some key.
// Usage in templates, where 'myKey' is known,
// to determine from a list of HIs, which one "isActive".
func (l *Info) HasKey(argKey string) bool {
	for _, key := range (*l).Keys {
		if argKey == key {
			return true
		}
	}
	return false
}

// ByKey retrieves a handlerinfo by key
func (l *InfosT) ByKey(argKey string) Info {
	for i := 0; i < len(*l); i++ {
		for _, key := range (*l)[i].Keys {
			if argKey == key {
				return (*l)[i]
			}
		}
	}
	log.Fatalf("unknown link key %v", argKey)
	return Info{}
}

// URLByKey directly returns canonical URL by key
func (l *InfosT) URLByKey(argKey string) string {
	ret := l.ByKey(argKey)
	return ret.Urls[0]
}

// ThisFunc returns a runtime func info
func ThisFunc() *runtime.Func {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc)
}

// ByRuntimeFunc retrieves a handlerinfo from *inside* a request handlerFunc:
//     handlerinfo.HIs.ByRuntimeFunc(util.ThisFunc())
//
// Handler funcs - similar to *this* in javaScript
// can now be retrieved with
//    myLink := links.Links.ByRuntimeFunc( util.ThisFunc() )
//
// We use
// 		reflect.ValueOf().Pointer() -
// and
// 		runtime.Caller() => program counter address => runtime.FuncForPc(pca).Entry()
// to compare the two program counter addresses.
//
// I am amazed that this works.
// Gives every handler access to *its* info object.
func (l *InfosT) ByRuntimeFunc(argFunc *runtime.Func) Info {

	for i := 0; i < len(*l); i++ {
		lpFunc := (*l)[i].Handler
		lpFuncType := reflect.ValueOf(lpFunc)
		if lpFuncType.Pointer() == argFunc.Entry() {
			return (*l)[i]
		}
	}
	log.Fatalf("unknown runtime func %+v", argFunc)
	return Info{}
}

// ByHandlerFunc returns the handlerinfo for a handler func
// from outside the request.
//
// 	  handlerFnc, canonPath := mux.Handler(req)
// 	  hi, err := handlerinfo.HIs.ByHandlerFunc(handlerFnc)
//
// Returns an error, if no handlerinfo exists.
func (l *InfosT) ByHandlerFunc(argFunc http.Handler) (Info, error) {

	hf, ok := argFunc.(http.HandlerFunc)
	if !ok {
		return Info{}, fmt.Errorf("Not a handlerfunc but %T\n\n %+v", argFunc, argFunc)
	}

	argFuncTypePointer := reflect.ValueOf(hf).Pointer()

	for i := 0; i < len(*l); i++ {
		lpFunc := (*l)[i].Handler
		lpFuncType := reflect.ValueOf(lpFunc)
		if lpFuncType.Pointer() == argFuncTypePointer {
			return (*l)[i], nil // found
		}
	}

	// Not found
	return Info{Allow: map[Privilege]bool{ReadOnly: true}}, fmt.Errorf("Unknown handler -%v- %T", argFunc, argFunc)
}

// MakeKeys generates keys, if none exists yet
func (l *InfosT) MakeKeys() {

	for i := 0; i < len(*l); i++ {
		if len((*l)[i].Keys) == 0 {
			key := strings.ToLower((*l)[i].Title)
			key = strings.Replace(key, " ", "-", -1)
			(*l)[i].Keys = append((*l)[i].Keys, key)
			// log.Printf("Url key %v", key)
		}
	}

	// Check uniqueness
	counter := map[string]int{}
	for _, link := range *l {
		for _, key := range link.Keys {
			counter[key]++
			if counter[key] > 1 {
				log.Fatalf("duplicate key %v", key)
			}
		}
	}

}
