package qst

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/zew/go-questionaire/trl"
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

	if q.Survey.Type == "" || !Mustaz09Underscore(q.Survey.Type) {
		s := fmt.Sprintf("WaveID must contain a SurveyID string consisting of lower case letters: %v", q.Survey.Type)
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

	navigationalNum := 0

	logEntries := 0

	// Check inputs
	// Set page and group width to 100
	// Set values for radiogroups
	for i1 := 0; i1 < len(q.Pages); i1++ {
		if q.Pages[i1].Width == 0 {
			q.Pages[i1].Width = 100
		}
		if !q.Pages[i1].NoNavigation {
			navigationalNum++
			q.Pages[i1].NavigationalNum = navigationalNum
		}
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			if q.Pages[i1].Groups[i2].Width == 0 {
				q.Pages[i1].Groups[i2].Width = 100
			}
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)
				inp := q.Pages[i1].Groups[i2].Inputs[i3]

				// Check input type
				if _, ok := implementedTypes[inp.Type]; !ok {
					return fmt.Errorf(s + fmt.Sprintf("Type '%v' is not in %v ", inp.Type, implementedTypes))
				}

				// Validator function exists
				if inp.Validator != "" {
					if _, ok := validators[inp.Validator]; !ok {
						return fmt.Errorf(s + fmt.Sprintf("Validator '%v' is not in %v ", inp.Validator, validators))
					}
				}

				// Helper: Add values 1...x for radiogroups
				for i4 := 0; i4 < len(inp.Radios); i4++ {
					if inp.Radios[i4].Val == "" {
						inp.Radios[i4].Val = fmt.Sprintf("%v", i4+1)
						logEntries++
						if logEntries < 10 {
							log.Printf(s + fmt.Sprintf("Value for %10v set to '%v'", inp.Radios[i4].Label, i4+1))
						}
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

				if q.Pages[i1].Groups[i2].Inputs[i3].IsLayout() {
					continue
				}

				if q.Pages[i1].Groups[i2].Inputs[i3].IsReserved() {
					return fmt.Errorf(s+"Name '%v' is reserved", nm)
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

	// err := q.ComputeDynamicContent()
	// if err != nil {
	// 	return err
	// }

	return nil
}

// ComputeDynamicContent computes statistics
func (q *QuestionaireT) ComputeDynamicContent(idx int) error {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if i1 != idx {
			continue
		}
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				if q.Pages[i1].Groups[i2].Inputs[i3].Type == "dynamic" {
					i := q.Pages[i1].Groups[i2].Inputs[i3]
					if _, ok := dynFuncs[i.DynamicFunc]; !ok {
						return fmt.Errorf("'%v' points to dynamic func '%v()' - which does not exist or is not registered", i.Name, i.DynamicFunc)
					}
					str, err := dynFuncs[i.DynamicFunc](q)
					if err != nil {
						return fmt.Errorf("'%v' points to dynamic func '%v()' - which returned error %v", i.Name, i.DynamicFunc, err)
					}
					q.Pages[i1].Groups[i2].Inputs[i3].Label = trl.S{q.LangCode: str}
					// log.Printf("'%v' points to dynamic func '%v()' - which returned '%v'", i.Name, i.DynamicFunc, str)
				}
			}
		}
	}
	return nil

}

// Hyphenize replaces "mittelfristig" with "mittel&shy;fristig"
// for all labels and descriptions
func (q *QuestionaireT) Hyphenize() {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				i := q.Pages[i1].Groups[i2].Inputs[i3]
				// s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)
				// log.Printf("Hyphenize: %v", s)
				for lc, v := range i.Label {
					v = trl.HyphenizeText(v)
					q.Pages[i1].Groups[i2].Inputs[i3].Label[lc] = v
				}
				for lc, v := range i.Desc {
					v := trl.HyphenizeText(v)
					q.Pages[i1].Groups[i2].Inputs[i3].Desc[lc] = v
				}
			}
		}
	}
}
