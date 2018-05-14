package qst

import (
	"fmt"
	"log"
	"strings"
)

// Validate checks whether essential elements of the questionaire exist.
func (q *QuestionaireT) Validate() error {

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

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Elements); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Elements[i2].Members); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)

				inp := q.Pages[i1].Elements[i2].Members[i3]
				if _, ok := implementedTypes[inp.Type]; !ok {
					return fmt.Errorf(s + fmt.Sprintf("Type %v is not in %v ", inp.Type, implementedTypes))
				}

				for i4 := 0; i4 < len(inp.Radios); i4++ {
					if inp.Radios[i4].Val == "" {
						inp.Radios[i4].Val = fmt.Sprintf("%v", i4+1)
						log.Printf(s + fmt.Sprintf("Value for %v set to %v", inp.Radios[i4].Label, i4+1))
					}
				}

			}
		}
	}

	names := map[string]int{}
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Elements); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Elements[i2].Members); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)

				// grp := q.Pages[i1].Elements[i2].Name
				nm := q.Pages[i1].Elements[i2].Members[i3].Name
				// tp := q.Pages[i1].Elements[i2].Members[i3].Type

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
