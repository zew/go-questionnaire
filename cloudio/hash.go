package cloudio

import (
	"crypto/sha256"
	"encoding/base64"
	"log"
)

// Md5Str gives MD5 of s
func Md5Str(s string) string {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(s))
	if err != nil {
		log.Printf("error writing %v to sha256 hasher: %v", s, err)
	}
	hshBytes := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(hshBytes) // no trailing equal signs
}
