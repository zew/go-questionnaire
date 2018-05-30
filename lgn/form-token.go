package lgn

/*
	Should we add the user name into the hashed base?
*/

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"time"
)

var tokenSalt = GeneratePassword(22) // we could renew this every x hours

func tok(hourOffset int) string {
	hasher := md5.New()
	io.WriteString(hasher, tokenSalt)
	t := time.Now()
	if hourOffset != 0 {
		t.Add(time.Duration(hourOffset) * time.Hour)
	}
	io.WriteString(hasher, t.Format("02.01.2006 15"))
	hash := hasher.Sum(nil)
	return hex.EncodeToString(hash)
}

// FormToken returns a form token.
// User independent.
func FormToken() string {
	return tok(0)
}

// ValidateFormToken checks tokens
// from previous hour to next two hours
func ValidateFormToken(arg string) error {
	if arg == tok(0) {
		return nil
	}
	if arg == tok(1) {
		return nil
	}
	if arg == tok(2) {
		return nil
	}
	if arg == tok(-1) {
		return nil
	}
	return fmt.Errorf("The form token was not issued within the last two hours. \nPlease re-login.")
}
