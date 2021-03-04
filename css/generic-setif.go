package css

import (
	"log"
	"reflect"
)

// SetIf sets a property only if it was null value ("", 0)
func (gcs *GridContainerStyle) SetIf(fn string, v reflect.Value) {

	valOfA := reflect.ValueOf(gcs)
	valOfA = valOfA.Elem() // dereference
	typeOfA := valOfA.Type()

	if _, ok := typeOfA.FieldByName(fn); ok { // fn exists
		fld := valOfA.FieldByName(fn)
		if fld.CanSet() {
			fld.Set(v)
			// log.Printf("Set string %v", fn)
		} else {
			log.Printf("Cannot set string %v", fn)
		}
	}

}
