package detect

import (
	"net/http"
	"regexp"
)

// Mobile Safari
// Mobile/14E304 Safari
var regC = regexp.MustCompile(`Mobile[/]{0,1}[0-9A-Za-z]{0,9} Safari`)

// Mobile ... Firefox
var regD = regexp.MustCompile(`Mobile.{0,}Firefox`)

// IsMobile just answers yes or no
func IsMobile(r *http.Request) bool {
	ua := r.Header.Get("User-Agent")
	if regC.MatchString(ua) {
		// log.Printf("UA1 %s", ua)
		return true
	}
	if regD.MatchString(ua) {
		// log.Printf("UA2 %s", ua)
		return true
	}
	return false
}
