package qst

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/zew/go-questionnaire/cfg"
)

// server side validation for an inputT
//
// Compare CompositeFuncT, dynFuncT
type validatorT func(*QuestionnaireT, *inputT) error

var validators = map[string]validatorT{}

func init() {

	functionBase := func(q *QuestionnaireT, inp *inputT, limit float64) error {

		arg := strings.TrimSpace(inp.Response)

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

	validators["inRange10"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 10) }
	validators["inRange20"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 20) }
	validators["inRange100"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 100) }
	validators["inRange1000"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 1000) }
	validators["inRange10000"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 10*1000) }
	validators["inRange50000"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 50*1000) }
	validators["inRange1Mio"] = func(q *QuestionnaireT, inp *inputT) error { return functionBase(q, inp, 1*1000*1000) }

	validators["must"] = func(q *QuestionnaireT, inp *inputT) error {
		arg := strings.TrimSpace(inp.Response)
		if arg == "" {
			return fmt.Errorf(cfg.Get().Mp["must_not_empty"].Tr(q.LangCode))
		}
		return nil
	}

	//
	validators["mustRadioGroup"] = func(q *QuestionnaireT, inp *inputT) error {
		arg := strings.TrimSpace(inp.Response)
		if arg == "0" || arg == "" {
			return fmt.Errorf(cfg.Get().Mp["must_one_option"].Tr(q.LangCode))
		}
		return nil
	}

	// party affiliation - pat1-4
	validators["otherParty"] = func(q *QuestionnaireT, inp *inputT) error {

		arg := strings.TrimSpace(inp.Response)

		inpMaster := q.ByName("q14")
		if inpMaster == nil {
			return nil
		}

		// q14 == "other"?
		if inpMaster.Response == "other" && arg == "" {
			return fmt.Errorf("Bitte andere Partei eintragen.")
		}
		return nil
	}

	validators["part2_qx_q123"] = func(q *QuestionnaireT, inp *inputT) error {

		core := strings.TrimPrefix(inp.Name, "part2_q")
		core = core[0:1]

		neighbors := []string{}
		for i := 1; i < 4; i++ {
			neighbors = append(neighbors, fmt.Sprintf("part2_q%v_q%v", core, i))
		}

		sum := 0

		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			summand, _ := strconv.Atoi(nb.Response)
			sum += summand
		}

		if sum != 10 {
			return fmt.Errorf("Summe muss 10 ergeben.")
		}

		return nil
	}
	validators["pat3_q4ab_opt123"] = func(q *QuestionnaireT, inp *inputT) error {

		core := strings.TrimPrefix(inp.Name, "q4")
		core = core[0:1]

		neighbors := []string{}
		for i := 1; i < 4; i++ {
			neighbors = append(neighbors, fmt.Sprintf("q4%v_opt%v", core, i))
		}

		// log.Printf("all three neigbors: %v", neighbors)

		sum := 0

		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			summand, _ := strconv.Atoi(nb.Response)
			sum += summand
		}

		if sum != 10 {
			return fmt.Errorf("Summe muss 10 ergeben.")
		}

		return nil
	}

	// pop2_part2 - IntroBUndEntscheidung78
	//
	// also
	// pop3_part2
	validators["preventInversion"] = func(q *QuestionnaireT, inp *inputT) error {

		// dec7_q1, dec7_q2
		// dec8_q1, dec8_q2
		//
		// dec3_q2, dec3_q2
		// dec4_q2, dec4_q2
		core := strings.TrimPrefix(inp.Name, "dec")
		core = core[0:1]

		neighbors := []string{}
		for i := 1; i < 3; i++ {
			neighbors = append(neighbors, fmt.Sprintf("dec%v_q%v", core, i))
		}
		// log.Printf("all three neigbors: %v", neighbors)

		val := "init"
		equal := false

		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			if val != "init" && nb.Response != "" && nb.Response == val {
				equal = true
				// log.Printf("neigbors: %v - value %v -  val %v - equal %v", nb.Name, nb.Response, val, equal)
			}
			if nb.Response != "" {
				val = nb.Response
			}
		}

		if equal {
			return fmt.Errorf("Schliesst sich aus.")
		}

		return nil
	}

	// pop3_part3
	validators["preventInversion2"] = func(q *QuestionnaireT, inp *inputT) error {

		// dec3_q1pol_gr1
		// dec3_q1
		core := strings.TrimPrefix(inp.Name, "dec")
		core = core[0:1]

		neighbors := []string{}
		for i := 1; i < 3; i++ {
			neighbors = append(neighbors, fmt.Sprintf("dec%v_q%v", core, i))
		}
		// log.Printf("all three neigbors: %v", neighbors)

		val := "init"
		equal := false

		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			if val != "init" && nb.Response != "" && nb.Response == val {
				equal = true
				// log.Printf("neigbors: %v - value %v -  val %v - equal %v", nb.Name, nb.Response, val, equal)
			}
			if nb.Response != "" {
				val = nb.Response
			}
		}

		if equal {
			return fmt.Errorf("Schliesst sich aus.")
		}

		return nil
	}

	// pop3_part2_q6_3
	validators["pop3_part2_q123456_1234"] = func(q *QuestionnaireT, inp *inputT) error {

		core := strings.TrimPrefix(inp.Name, "pop3_part2_q")
		core = core[0:1]

		neighbors := []string{}
		for i := 1; i < 5; i++ {
			neighbors = append(neighbors, fmt.Sprintf("pop3_part2_q%v_%v", core, i))
		}

		// log.Printf("all three neigbors: %v", neighbors)

		sum := 0

		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			summand, _ := strconv.Atoi(nb.Response)
			sum += summand
		}

		if sum != 10 {
			return fmt.Errorf("Summe muss 10 ergeben.")
		}

		return nil
	}

	// pat - for pop1-3
	validators["oneOfPrefixQ20"] = func(q *QuestionnaireT, inp *inputT) error {

		fields := []string{
			"q20_inactive",
			"q20_votesometimes",
			"q20_voteregular",
			"q20_petitions",
			"q20_communal",
			"q20_state",
			"q20_federal",
		}
		atLeastOne := false

		for _, fld := range fields {
			fd := q.ByName(fld)
			if fd.Response == "1" {
				atLeastOne = true
			}
		}

		if !atLeastOne {
			return fmt.Errorf("Wenigstens eine Angabe.")
		}

		return nil

	}

	// pat - for pop3
	validators["patMustOneAvailabe"] = func(q *QuestionnaireT, inp *inputT) error {

		/*
		   q2_seq1_row1_rad - vals 1 or 2
		   q2_seq1_row2_rad - vals 1 or 2
		   q2_seq1_row3_rad - vals 1 or 2

		   q2_seq2_row1_rad
		   q2_seq2_row2_rad
		   q2_seq2_row3_rad

		*/

		core := strings.TrimPrefix(inp.Name, "q2_seq")
		core = core[0:1]

		neighbors := []string{}
		for i := 1; i < 4; i++ {
			neighbors = append(neighbors, fmt.Sprintf("q2_seq%v_row%v_rad", core, i))
		}

		allSet := true

		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			// summand, _ := strconv.Atoi(nb.Response)
			if nb.Response != "1" && nb.Response != "2" {
				allSet = false
			}
		}
		if !allSet {
			return fmt.Errorf("Bitte für alle drei Zeilen eine Option setzen.")
		}

		allUnavailable := 0
		for _, neighbor := range neighbors {
			nb := q.ByName(neighbor)
			summand, _ := strconv.Atoi(nb.Response)
			allUnavailable += summand
		}
		if allUnavailable > 5 {
			return fmt.Errorf("Mindestens eine Option muss verfügbar sein.")
		}

		return nil
	}

	// q17
	validators["citizenshipyes"] = func(q *QuestionnaireT, inp *inputT) error {
		if inp.Response != "" && inp.Response != "citizenshipyes" {
			err1 := ErrorForward{markDownPath: "must-german-citizen.md"}
			err := errors.Wrap(err1, "Dt. Staatsbürger erforderl")
			return err
		}
		return nil
	}

	validators["comprehensionPOP2"] = func(q *QuestionnaireT, inp *inputT) error {

		erroneous := false
		empty := false
		neighbors := map[string]string{
			"q_found_compr_a": "est_2",
			"q_found_compr_b": "est_c",
		}
		for neighbor, solution := range neighbors {
			nb := q.ByName(neighbor)
			// summand, _ := strconv.Atoi(nb.Response)
			if nb.Response != "" && strings.TrimSpace(nb.Response) != solution {
				erroneous = true
			}
			if nb.Response == "" {
				empty = true
			}
		}

		if empty {
			err := errors.New("Bitte beide Fragen beantworten.")
			return err
		}

		if erroneous {
			vlStr, _ := q.Attrs["comprehensionPOP2"]
			vl, _ := strconv.Atoi(vlStr)
			vl++
			q.Attrs["comprehensionPOP2"] = fmt.Sprint(vl)

			if vl%2 == 1 {
				err := fmt.Errorf(`
					<div  class='comprehension-error'>
						Mindestens eine Ihrer beiden Antworten ist falsch.
						Lesen Sie genau Anleitung u Fragen. 
						<br>
						<span  style='font-size: 110%%;'>Versuch %v von 3  </span>
					</div>`,
					vl/2+1,
				)

				if vl > 3 {
					err1 := ErrorForward{
						// Quality-Redirect
						// ... ErrorForward{markDownPath: "must-german-citizen.md"}
						markDownPath: "https://webs.norstatsurveys.com/z/Quality",
					}
					err = errors.Wrap(err1, err.Error())
				}

				return err
			}

		} else {
			q.Attrs["comprehensionPOP2"] = "0" // reset
		}

		return nil
	}

	//
	//
	//
	validators["comprehensionPOP3"] = func(q *QuestionnaireT, inp *inputT) error {

		erroneous := false
		empty := false
		neighbors := map[string]string{
			"q_tpref_compr_a": "3",
			"q_tpref_compr_b": "7",
		}
		for neighbor, solution := range neighbors {
			nb := q.ByName(neighbor)
			// summand, _ := strconv.Atoi(nb.Response)
			if nb.Response != "" && strings.TrimSpace(nb.Response) != solution {
				erroneous = true
			}
			if nb.Response == "" {
				empty = true
			}
		}

		if empty {
			err := errors.New("Bitte beide Fragen beantworten.")
			return err
		}

		if erroneous {
			vlStr, _ := q.Attrs["comprehensionPOP3"]
			vl, _ := strconv.Atoi(vlStr)
			vl++
			q.Attrs["comprehensionPOP3"] = fmt.Sprint(vl)

			if vl%2 == 1 {
				err := fmt.Errorf(`
					<div  class='comprehension-error'>
						Mindestens eine Ihrer beiden Antworten ist falsch.
						Lesen Sie genau Anleitung u Fragen. 
						<br>
						<span  style='font-size: 110%%;'>Versuch %v von 3  </span>
					</div>`,
					vl/2+1,
				)

				if vl > 3 {
					err1 := ErrorForward{
						// Quality-Redirect
						// ... ErrorForward{markDownPath: "must-german-citizen.md"}
						markDownPath: "https://webs.norstatsurveys.com/z/Quality",
					}
					err = errors.Wrap(err1, err.Error())
				}

				return err
			}

		} else {
			q.Attrs["comprehensionPOP3"] = "0" // reset
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
func (q *QuestionnaireT) ValidateResponseData(pageNum int, langCode string) (last error, forward *ErrorForward) {

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if i1 != pageNum {
			continue
		}

		// pre process error proxies
		//   collect a list of error proxies accessible by Param
		//   (re-) set their error messages to ""
		errorProxies := map[string]*inputT{} // per page, not global
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				if inp.Type == "dyn-textblock" && inp.DynamicFunc == "ErrorProxy" {
					q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = ""
					errorProxies[inp.Param] = inp
				}
			}
		}

		//
		// main run
		//    executing validator funcs
		//    storing results in inp.ErrMsg
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				// s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)
				inp := q.Pages[i1].Groups[i2].Inputs[i3]

				// Validator function exists
				if inp.Validator != "" {

					// Reset previous errors
					q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = ""

					valiKeys := strings.Split(inp.Validator, ";")
					for _, valiKey := range valiKeys {
						if valiFunc, ok := validators[strings.TrimSpace(valiKey)]; ok {
							err := valiFunc(q, inp)
							// log.Printf("%-10v %-20s  %-12s  %v", inp.Name, valiKey, inp.Response, err)
							if err != nil {
								last = err
								q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = err.Error()

								checkForSpecialErr := &ErrorForward{}
								// log.Printf("\t\terr as *ErrorForward %v", errors.As(err, checkForSpecialErr))
								if errors.As(err, checkForSpecialErr) {
									forward = checkForSpecialErr
								}
							}
						}
					}
				}

			}
		}

		// post process error proxies
		//    for all inputs having an error message
		//      for those having an error proxy
		// 		  error proxy takes the error message; input error message gets deleted
		if len(errorProxies) > 0 {
			for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
				for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
					inp := q.Pages[i1].Groups[i2].Inputs[i3]
					if inp.ErrMsg != "" {
						if inp.Type != "dyn-textblock" && inp.DynamicFunc != "ErrorProxy" {

							for startsWith, errProx := range errorProxies {
								if strings.HasPrefix(inp.Name, startsWith) {
									errProx.ErrMsg = inp.ErrMsg
									q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = ""
									// log.Printf("Re-assigning ErrMsg from %v to ErrorProxy: %v", inp.Name, errProx.ErrMsg)
								}
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

// ErrorForward contains a markdown page or page id
// to jump to
type ErrorForward struct {
	markDownPath string
	// pageIdx int
}

// Error implements the errors.Error interface
func (ef ErrorForward) Error() string {
	return ""
}

// MarkDownPath returns the path to redirect to;
//   it should be renamed - since it might also contain absolute URLs
func (ef ErrorForward) MarkDownPath() string {
	return ef.markDownPath
}
