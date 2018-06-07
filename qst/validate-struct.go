package qst

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var not09azHyphenUnderscore = regexp.MustCompile(`[^a-z0-9\_\-]+`)

// Mustaz09Underscore tests strings for a-z, 0-9, _
func Mustaz09Underscore(s string) bool {
	if not09azHyphenUnderscore.MatchString(s) {
		return false
	}
	return true
}

// Validate checks whether essential elements of the questionaire exist.
func (q *QuestionaireT) Validate() error {

	if q.WaveID.SurveyID == "" || !Mustaz09Underscore(q.WaveID.SurveyID) {
		s := fmt.Sprintf("WaveID must contain a SurveyID string consisting of lower case letters: %v", q.WaveID.SurveyID)
		log.Printf(s)
		return fmt.Errorf(s)
	}
	if q.LangCode == "" {
		s := fmt.Sprintf("Language code is empty. Must be one of %v", q.LangCodes)
		log.Printf(s)
		return fmt.Errorf(s)
	}
	if _, ok := q.LangCodes[q.LangCode]; !ok {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Printf(s)
		return fmt.Errorf(s)
	}

	// Check inputs
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)
				inp := q.Pages[i1].Groups[i2].Inputs[i3]

				// Check input type
				if _, ok := implementedTypes[inp.Type]; !ok {
					return fmt.Errorf(s + fmt.Sprintf("Type %v is not in %v ", inp.Type, implementedTypes))
				}

				// Validator function exists
				if inp.Validator != "" {
					if _, ok := validators[inp.Validator]; !ok {
						return fmt.Errorf(s + fmt.Sprintf("Type %v is not in %v ", inp.Type, implementedTypes))
					}
				}

				// Helper: Add values 1...x for radiogroups
				for i4 := 0; i4 < len(inp.Radios); i4++ {
					if inp.Radios[i4].Val == "" {
						inp.Radios[i4].Val = fmt.Sprintf("%v", i4+1)
						log.Printf(s + fmt.Sprintf("Value for %10v set to '%v'", inp.Radios[i4].Label, i4+1))
					}
				}

			}
		}
	}

	// Make sure, input names are unique
	names := map[string]int{}
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)

				// grp := q.Pages[i1].Elements[i2].Name
				nm := q.Pages[i1].Groups[i2].Inputs[i3].Name
				tp := q.Pages[i1].Groups[i2].Inputs[i3].Type

				if tp == "textblock" {
					continue
				}

				if nm == "" {
					return fmt.Errorf(s+"Name %v is empty", nm)
				}

				if not09azHyphenUnderscore.MatchString(nm) {
					return fmt.Errorf(s+"Name %v must consist of [a-z0-9_-]", nm)
				}

				names[nm]++

			}
		}
	}

	for k, v := range names {
		if v > 1 {
			s := fmt.Sprintf("Page element '%v' is not unique  (%v)", k, v)
			log.Printf(s)
			return fmt.Errorf(s)
		}
		if k != strings.ToLower(k) {
			s := fmt.Sprintf("Page element '%v' is not lower case  (%v)", k, v)
			log.Printf(s)
			return fmt.Errorf(s)
		}
	}

	return nil
}
