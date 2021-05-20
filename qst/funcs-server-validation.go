package qst

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

// server side validation for an inputT
//
// Compare CompositeFuncT, dynFuncT
type validatorT func(*QuestionnaireT, string) error

var validators = map[string]validatorT{}

func init() {

	functionBase := func(q *QuestionnaireT, arg string, limit float64) error {

		arg = strings.TrimSpace(arg)

		// non-empty is in separate validator 'must'
		if arg == "" {
			return nil
		}

		// comma => dot
		if strings.Contains(arg, ",") && !strings.Contains(arg, ".") {
			arg = strings.Replace(arg, ",", ".", -1)
		}

		// comma and dot; i.e. 100.000,00 or 100,000.00
		if strings.Contains(arg, ",") && strings.Contains(arg, ".") {
			arg = strings.Replace(arg, ",", ".", -1) // map everything to dot
		}

		// 100.000.00 => 100000.00
		if occs := strings.Count(arg, "."); occs > 1 {
			arg = strings.Replace(arg, ".", "", occs-1) // replace every dot but the last
		}

		fl, err := strconv.ParseFloat(arg, 64)

		// log.Printf("Checking %6v against %6v => %6v %v", arg, limit, fl, err)
		if err != nil {
			// ParseFloat yields ugly error messages
			// strconv.ParseFloat: parsing "3 3" invalid syntax
			return fmt.Errorf(cfg.Get().Mp["not_a_number"].Tr(q.LangCode), arg)
		}
		// Understandable in every language
		if fl > limit {
			log.Printf("%.2f > max %.0f", fl, limit)
			return fmt.Errorf(cfg.Get().Mp["too_big"].Tr(q.LangCode), limit)
		}
		if fl < -limit {
			log.Printf("%.2f < min %.0f", fl, -limit)
			return fmt.Errorf(cfg.Get().Mp["too_small"].Tr(q.LangCode), -limit)
		}
		return nil
	}

	validators["inRange10"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 10) }
	validators["inRange20"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 20) }
	validators["inRange100"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 100) }
	validators["inRange1000"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 1000) }
	validators["inRange10000"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 10*1000) }
	validators["inRange50000"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 50*1000) }
	validators["inRange1Mio"] = func(q *QuestionnaireT, arg string) error { return functionBase(q, arg, 1*1000*1000) }

	validators["mustRadioGroup"] = func(q *QuestionnaireT, arg string) error {
		if arg == "0" || arg == "" {
			return fmt.Errorf(cfg.Get().Mp["must_one_option"].Tr(q.LangCode))
		}
		return nil
	}
	validators["must"] = func(q *QuestionnaireT, arg string) error {
		if strings.TrimSpace(arg) == "" {
			return fmt.Errorf(cfg.Get().Mp["must_not_empty"].Tr(q.LangCode))
		}
		return nil
	}
	validators["otherParty"] = func(q *QuestionnaireT, arg string) error {
		inp := q.ByName("q14")
		if inp == nil {
			return nil
		}
		// q14 == "other"?
		if inp.Response == "other" && arg == "" {
			return fmt.Errorf("Bitte andere Partei eintragen.")
		}
		return nil
	}

}

// ConsolidateRadioErrors removes repeating error messages from radio inputs
func (page *pageT) ConsolidateRadioErrors(grpOrder []int) {

	wasRadio := ""

	for _, grpIdx := range grpOrder {
		// for i2 := 0; i2 < len(page.Groups); i2++ {
		for i3 := 0; i3 < len(page.Groups[grpIdx].Inputs); i3++ {
			name := page.Groups[grpIdx].Inputs[i3].Name
			isRadio := page.Groups[grpIdx].Inputs[i3].Type == "radio"
			hasMsg := page.Groups[grpIdx].Inputs[i3].ErrMsg != ""
			if isRadio && hasMsg {
				if wasRadio == "" || (wasRadio != "" && wasRadio != name) {
					wasRadio = name
				} else {
					page.Groups[grpIdx].Inputs[i3].ErrMsg = ""
				}
			}
		}
	}

}

// ValidateResponseData applies all input validation rules on the responses.
// Restricted by page, since validation errors are handled page-wise.
func (q *QuestionnaireT) ValidateResponseData(pageNum int, langCode string) (last error) {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if i1 != pageNum {
			continue
		}
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				// s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)

				// Check input type
				inp := q.Pages[i1].Groups[i2].Inputs[i3]

				// Validator function exists
				if inp.Validator != "" {
					valiKeys := strings.Split(inp.Validator, ";")
					for _, valiKey := range valiKeys {
						if valiFunc, ok := validators[strings.TrimSpace(valiKey)]; ok {
							err := valiFunc(q, inp.Response)
							// log.Printf("Validating %22s  -%s-  %v", inp.Name, inp.Response, err)
							if err != nil {
								last = err
								q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = err.Error()
							} else {
								// Reset previous errors
								q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = ""
							}
						}
					}
				}

			}
		}

		// grpOrder := q.RandomizeOrder(pageNum)
		// q.Pages[i1].ConsolidateRadioErrors(grpOrder)

	}

	if last != nil {
		q.HasErrors = true
	} else {
		q.HasErrors = false
	}

	return
}

// DumpErrors logs all ErrMsgs from the questionnaire
func (q *QuestionnaireT) DumpErrors() {
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				if inp.ErrMsg != "" {
					log.Print(inp.ErrMsg)
				}
			}
		}
	}
}
