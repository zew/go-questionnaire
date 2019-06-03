package lgn

import (
	"fmt"
	"log"
	"net/http"

	"github.com/monoculum/formam"
	hashids "github.com/speps/go-hashids"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/util"
)

var gen *hashids.HashIDData

// Every trial with init() construction failed
// because Get().Salt is not yet initialized
func getGen() *hashids.HashIDData {
	if gen == nil {
		gen = hashids.NewData()
		gen.MinLength = 6
		gen.Alphabet = "23456789ABCDEFGHKLMNPRSTUVWXYZ" // ultraReadable; 30 chars - drop 5 and S ?
		gen.Salt = Get().Salt
	}
	return gen

}

// HashIDDecodeFirst turns a string into a slice of integers
// and returns the first integer
func HashIDDecodeFirst(encoded string) int {

	h, err := hashids.NewWithData(getGen())
	decoded, err := h.DecodeWithError(encoded)
	if err != nil {
		log.Printf("Could not decode %v", encoded)
	}
	if len(decoded) > 0 {
		return decoded[0]
	}
	return -1

}

// GenerateHashIDs encodes integer IDs into a kind of base64 encoded string.
func GenerateHashIDs(w http.ResponseWriter, r *http.Request) {

	if cfg.Get().IsProduction {
		l, isLoggedIn, err := LoggedInCheck(w, r)
		if err != nil {
			fmt.Fprintf(w, "Login error %v", err)
			return
		}
		if !isLoggedIn {
			fmt.Fprintf(w, "Not logged in")
			return
		}
		if !l.HasRole("admin") {
			fmt.Fprintf(w, "admin login required")
			return
		}
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	msg := `GenerateHashIDs encodes user IDs into encoded strings;
See https://hashids.org for background;	
the alphabet is strongly reduced for readability (far less than base64);
encoded strings are usable for direct login by the string alone,
if matching DirectLoginRanges are configured in package cfg.
The encoded strings avalanche nicely. But the algorithm is *no* MD5 hashing.

Security:
With six characters we get roughly 730 million combinations (30^^6).

If we allow DirectLoginRangesT for UserIDs between 0 and 5000,
then there are 5000 possible combinations out of 730 million to a valid direct login.

Odds for brute force are 1 in 150.000 - two tries per second - one hit per day.

Keep in mind that the avalanching is not perfect.

That could still be acceptable in small surveys with trusted participants,
where comfort really matters. Where there is little incentive to cheat and 
little to gain from getting someone elses responses.

But keep an eye on the application log.
	
`
	w.Write([]byte(msg))

	type formEntryT struct {
		Start int `json:"start"`
		Stop  int `json:"stop"`
	}
	fe := formEntryT{}
	dec := formam.NewDecoder(&formam.DecoderOptions{TagName: "json"})
	err := dec.Decode(r.Form, &fe)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		fmt.Fprintf(w, "%v\n", util.IndentedDump(fe))
		return
	}

	if fe.Start == 0 || fe.Stop == 0 {
		fe.Start = 1000
		fe.Stop = 1010
	}

	for i := fe.Start + 0; i < fe.Stop; i++ {

		h, err := hashids.NewWithData(getGen())
		if err != nil {
			w.Write([]byte(err.Error()))
		}
		encoded, err := h.Encode([]int{i})
		if err != nil {
			w.Write([]byte(err.Error()))
		}

		// encode multiple numbers
		if false {
			encodedLong, _ := h.Encode([]int{i, i + 1, i + 2, i + 3, i + 4, i + 5})
			fmt.Fprintf(w, "%v - %v\n", i, encodedLong)
		}

		fmt.Fprintf(w, "%v\t%v\t%v\n", i, encoded, HashIDDecodeFirst(encoded))
	}

}
