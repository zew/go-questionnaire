package lgn

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pbberlin/dbg"
	"github.com/pbberlin/struc2frm"
	"github.com/zew/go-questionnaire/pkg/lgn/shuffler"
)

type entryForm struct {
	Group01     string `json:"group01,omitempty"    form:"subtype='fieldset',label='Range of User IDs'"`
	UserIDStart int    `json:"start"`
	UserIDStop  int    `json:"stop"`

	Group02              string `json:"group02,omitempty"    form:"subtype='fieldset',label='Global settings'"`
	ShufflingVariations  int    `json:"variations"           form:"label='Number distinct sufflings before repeat'" `                    // q.ShufflingVariations
	ShufflingRepetitions int    `json:"repetitions"          form:"label='Number of shuffling operations',suffix='just keep default 3'"` // q.ShufflingRepetitions

	Group03           string `json:"group03,omitempty"       form:"subtype='fieldset',label='Randomization group'"`
	RandomizationSeed int    `json:"randomization_seed"      form:"suffix='>0 - distinguish multiple shufflings on same page'"`
	MaxElements       int    `json:"max_elements"            form:"suffix='number of elements to shuffle / randomize'" `
}

func (frm entryForm) Validate() (map[string]string, bool) {
	errs := map[string]string{}
	g1 := frm.UserIDStart == 0
	if !g1 {
		errs["Start"] = "Missing Start"
	}
	return errs, g1
}

// ShufflesToCSV computes random but deterministic shufflings for usage outside the app
func ShufflesToCSV(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	fmt.Fprintf(w, "<h3>Builtin shufflings aka randomization of groups on pages</h3>")

	// getting a converter
	s2f := struc2frm.New()   // or clone existing one
	s2f.ShowHeadline = false // set options
	s2f.Indent = 280
	s2f.SetOptions("department", []string{"ub", "fm"}, []string{"UB", "FM"})

	// init values - non-multiple
	frm := entryForm{
		UserIDStart: 10000,
		UserIDStop:  10100,

		ShufflingVariations:  8,
		ShufflingRepetitions: 3,

		RandomizationSeed: 1,
		MaxElements:       3,
	}

	// pulling in values from http request
	populated, err := struc2frm.Decode(req, &frm)
	if populated && err != nil {
		s2f.AddError("global", fmt.Sprintf("cannot decode form: %v<br>\n <pre>%v</pre>", err, dbg.Dump2String(req.Form)))
		log.Printf("cannot decode form: %v<br>\n <pre>%v</pre>", err, dbg.Dump2String(req.Form))
	}

	// init values - multiple
	if !populated {
		// n.a.
	}

	errs, valid := frm.Validate()

	if populated {
		if !valid {
			s2f.AddErrors(errs) // add errors only for a populated form
		} else {
			// further processing with valid form data
		}
	}

	if !valid {
		// render to HTML for user input / error correction
		fmt.Fprint(w, s2f.Form(frm))
	}

	//
	// Keep this conform with
	// (q *QuestionnaireT) RandomizeOrder()
	fmt.Fprintf(w, "<pre>")
	fmt.Fprintf(w, "%5v\t%v\t%v\n", "userID", "class", "[...]")
	for userID := frm.UserIDStart; userID <= frm.UserIDStop; userID++ {
		seedSufflngs := userID + frm.RandomizationSeed
		sh := shuffler.New(seedSufflngs, frm.ShufflingVariations, frm.MaxElements)
		order := sh.Slice(frm.ShufflingRepetitions)
		class := seedSufflngs % frm.ShufflingVariations
		fmt.Fprintf(w, "%5v\t%v\t%v\n", userID, class, order)
	}
	fmt.Fprintf(w, "</pre>")

}
