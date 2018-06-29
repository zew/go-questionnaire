package lgn

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

var tokenSaltNotWorking = GeneratePassword(22) // not interoperational between multiple instances of go-questionaire, transferrer, generator

func tok(hourOffset int) string {
	hasher := md5.New()
	io.WriteString(hasher, lgns.Salt)
	t := time.Now()
	if hourOffset != 0 {
		t = t.Add(time.Duration(hourOffset) * time.Hour)
	}
	// log.Printf("token time: %v", t.Format("02.01.2006 15"))
	io.WriteString(hasher, t.Format("02.01.2006 15"))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// FormToken returns a form token.
// User independent.
/*
	Should we add the user name into the hashed base?
*/
func FormToken() string {
	return tok(0)
}

// ValidateFormToken checks tokens
// from previous two hours - and from next hour
func ValidateFormToken(arg string) error {
	for i := 0; i > -3; i-- {
		if arg == tok(i) {
			return nil
		}
	}
	if arg == tok(1) {
		return nil
	}
	return fmt.Errorf("The form token was not issued within the last two hours. \nPlease re-login.")
}
