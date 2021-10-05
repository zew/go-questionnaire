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
func (sr *StyleGridContainer) Combine(b StyleGridContainer) {

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
func (sr *StyleBox) Combine(b StyleBox) {

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
func (sr *StyleGridItem) Combine(b StyleGridItem) {

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
func (sr *StyleText) Combine(b StyleText) {

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
