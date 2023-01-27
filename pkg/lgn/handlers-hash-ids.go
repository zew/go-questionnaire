package lgn

import (
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/go-playground/form"
	"github.com/pbberlin/dbg"
	hashids "github.com/speps/go-hashids"
	"github.com/zew/go-questionnaire/pkg/cfg"
)

var gen *hashids.HashIDData

// init() is impossible
// because lgn.Get().Salt is not yet initialized
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
	if err != nil {
		log.Printf("Could not creae HashID from %v; error %v", getGen(), err)
	}
	decoded, err := h.DecodeWithError(encoded)
	if err != nil {
		log.Printf("Could not decode %v; error %v", encoded, err)
	}
	if len(decoded) > 0 {
		return decoded[0]
	}
	return -1

}

// GenerateHashIDs encodes integer IDs into a kind of base64 encoded string.
func GenerateHashIDs(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	msg := `<pre>GenerateHashIDs encodes user IDs into encoded strings;
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


</pre>	
<span style='font-family: monospace;' >
`
	fmt.Fprint(w, msg)

	type formEntryT struct {
		Host  string `json:"host,omitempty"`
		Start int    `json:"start"`
		Stop  int    `json:"stop"`
	}
	fe := formEntryT{}
	dec := form.NewDecoder()
	dec.SetTagName("json") // recognizes and ignores ,omitempty
	err := dec.Decode(&fe, r.Form)
	if err != nil {
		fmt.Fprintf(w, "%v\n", err)
		fmt.Fprintf(w, "%v\n", dbg.Dump2String(fe))
		return
	}

	if fe.Start == 0 || fe.Stop == 0 {
		fe.Start = 10000 + 0
		fe.Stop = 10000 + 10
	}

	for i := fe.Start + 0; i < fe.Stop; i++ {

		h, err := hashids.NewWithData(getGen())
		if err != nil {
			fmt.Fprintf(w, "Error creating hash ID: %v", err.Error())
		}

		encoded, err := h.Encode([]int{i})
		if err != nil {
			fmt.Fprintf(w, "Error encoding hash ID: %v", err.Error())
		}

		// encode multiple numbers
		if false {
			encodedLong, _ := h.Encode([]int{i, i + 1, i + 2, i + 3, i + 4, i + 5})
			fmt.Fprintf(w, "%v - %v\n", i, encodedLong)
		}

		pathPrefixed := cfg.Pref("d/" + encoded)
		if fe.Host == "https://survey2.zew.de" || fe.Host == "https://private-debt-survey.zew.de" {
			pathPrefixed = path.Join("/", "d", encoded) // this host hast no prefix
		}

		fmt.Fprintf(w,
			"%v\t%v\t<a href='%v' target=_blank>%v</a>   <br>\n",
			i,
			HashIDDecodeFirst(encoded),
			fe.Host+pathPrefixed,
			encoded,
		)
	}

	fmt.Fprint(w, "</span>")
}
