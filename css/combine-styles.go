package css

import (
	"log"
	"reflect"
)

/*
	the same reflect methods to
	copy one struct instance over the other

	only the func argument types differ

	for usage in

	func (sr *StylesResponsive) Combine(b StylesResponsive) {
		...
	}

*/

// Combine adds a to sr
func (sr *GridContainerStyle) Combine(b GridContainerStyle) {

	valOfA := reflect.ValueOf(sr)
	valOfA = valOfA.Elem() // dereference

	valOfB := reflect.ValueOf(b)

	typeOfB := valOfB.Type()
	if valOfB.Kind().String() != "struct" {
		//
	}

	for i := 0; i < valOfB.NumField(); i++ {

		fnB := typeOfB.Field(i).Name
		// log.Printf("Fieldname %v", fnB)

		vlB := valOfB.Field(i).Interface()

		// string
		if vlBStr, ok := vlB.(string); ok {
			if vlBStr != "" {
				vlA := valOfA.Field(i).Interface()
				if vlAStr, ok := vlA.(string); ok {
					if vlAStr == "" { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set string %v", fnB)
						} else {
							log.Printf("Cannot set string %v", fnB)
						}
					} else {
						// log.Printf("Base for string %v not empty %v", fnB, vlAStr)
					}
				}
			}
		}

		// int
		if vlBInt, ok := vlB.(int); ok {
			if vlBInt != 0 {
				vlA := valOfA.Field(i).Interface()
				if vlAInt, ok := vlA.(int); ok {
					if vlAInt == 0 { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set int %v", fnB)
						} else {
							log.Printf("Cannot set int %v", fnB)
						}
					} else {
						// log.Printf("Base for int %v not empty %v", fnB, vlAInt)
					}
				}
			}
		}

	}
}

// Combine adds a to sr
func (sr *BoxStyle) Combine(b BoxStyle) {

	valOfA := reflect.ValueOf(sr)
	valOfA = valOfA.Elem() // dereference

	valOfB := reflect.ValueOf(b)

	typeOfB := valOfB.Type()
	if valOfB.Kind().String() != "struct" {
		//
	}

	for i := 0; i < valOfB.NumField(); i++ {

		fnB := typeOfB.Field(i).Name
		// log.Printf("Fieldname %v", fnB)

		vlB := valOfB.Field(i).Interface()

		// string
		if vlBStr, ok := vlB.(string); ok {
			if vlBStr != "" {
				vlA := valOfA.Field(i).Interface()
				if vlAStr, ok := vlA.(string); ok {
					if vlAStr == "" { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set string %v", fnB)
						} else {
							log.Printf("Cannot set string %v", fnB)
						}
					} else {
						// log.Printf("Base for string %v not empty %v", fnB, vlAStr)
					}
				}
			}
		}

		// int
		if vlBInt, ok := vlB.(int); ok {
			if vlBInt != 0 {
				vlA := valOfA.Field(i).Interface()
				if vlAInt, ok := vlA.(int); ok {
					if vlAInt == 0 { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set int %v", fnB)
						} else {
							log.Printf("Cannot set int %v", fnB)
						}
					} else {
						// log.Printf("Base for int %v not empty %v", fnB, vlAInt)
					}
				}
			}
		}

	}
}

// Combine adds a to sr
func (sr *GridItemStyle) Combine(b GridItemStyle) {

	valOfA := reflect.ValueOf(sr)
	valOfA = valOfA.Elem() // dereference

	valOfB := reflect.ValueOf(b)

	typeOfB := valOfB.Type()
	if valOfB.Kind().String() != "struct" {
		//
	}

	for i := 0; i < valOfB.NumField(); i++ {

		fnB := typeOfB.Field(i).Name
		// log.Printf("Fieldname %v", fnB)

		vlB := valOfB.Field(i).Interface()

		// string
		if vlBStr, ok := vlB.(string); ok {
			if vlBStr != "" {
				vlA := valOfA.Field(i).Interface()
				if vlAStr, ok := vlA.(string); ok {
					if vlAStr == "" { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set string %v", fnB)
						} else {
							log.Printf("Cannot set string %v", fnB)
						}
					} else {
						// log.Printf("Base for string %v not empty %v", fnB, vlAStr)
					}
				}
			}
		}

		// int
		if vlBInt, ok := vlB.(int); ok {
			if vlBInt != 0 {
				vlA := valOfA.Field(i).Interface()
				if vlAInt, ok := vlA.(int); ok {
					if vlAInt == 0 { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set int %v", fnB)
						} else {
							log.Printf("Cannot set int %v", fnB)
						}
					} else {
						// log.Printf("Base for int %v not empty %v", fnB, vlAInt)
					}
				}
			}
		}

	}
}

// Combine adds a to sr
func (sr *TextStyle) Combine(b TextStyle) {

	valOfA := reflect.ValueOf(sr)
	valOfA = valOfA.Elem() // dereference

	valOfB := reflect.ValueOf(b)

	typeOfB := valOfB.Type()
	if valOfB.Kind().String() != "struct" {
		//
	}

	for i := 0; i < valOfB.NumField(); i++ {

		fnB := typeOfB.Field(i).Name
		// log.Printf("Fieldname %v", fnB)

		vlB := valOfB.Field(i).Interface()

		// string
		if vlBStr, ok := vlB.(string); ok {
			if vlBStr != "" {
				vlA := valOfA.Field(i).Interface()
				if vlAStr, ok := vlA.(string); ok {
					if vlAStr == "" { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set string %v", fnB)
						} else {
							log.Printf("Cannot set string %v", fnB)
						}
					} else {
						// log.Printf("Base for string %v not empty %v", fnB, vlAStr)
					}
				}
			}
		}

		// int
		if vlBInt, ok := vlB.(int); ok {
			if vlBInt != 0 {
				vlA := valOfA.Field(i).Interface()
				if vlAInt, ok := vlA.(int); ok {
					if vlAInt == 0 { // base is empty
						if valOfA.FieldByName(fnB).CanSet() {
							valOfA.FieldByName(fnB).Set(valOfB.Field(i))
							// log.Printf("Set int %v", fnB)
						} else {
							log.Printf("Cannot set int %v", fnB)
						}
					} else {
						// log.Printf("Base for int %v not empty %v", fnB, vlAInt)
					}
				}
			}
		}

	}
}
