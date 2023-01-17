package qst

import (
	"fmt"
	"log"
	"strings"
)

func pdsSpecialDisableColumns(q *QuestionnaireT, page *pageT, pageIdx, acIdx int) {

	if q.Survey.Type != "pds" {
		return
	}

	cond2 := page.CounterProgress == "A2" || page.CounterProgress == "B2" || page.CounterProgress == "C2"
	cond1 := page.CounterProgress == "page12"

	if cond1 || cond2 {

		/*
			pages 2- 6   => acIdx = 0
			pages 7-11   => acIdx = 1
			pages12-16   => acIdx = 2

			but we just run it for all combinations
		*/

		for _, ac := range PDSAssetClasses {

			for _, tt := range ac.TrancheTypes {
				pfx := fmt.Sprintf("%v_%v", ac.Prefix, tt.Prefix)
				name := fmt.Sprintf("%v_q11a_numtransact_main", pfx)
				inp := q.ByName(name)
				if inp == nil { // page not initialized
					// log.Printf("   DA: name %v not initialized", name)
					continue
				}
				val := inp.Response

				if val != "" {
					log.Printf(
						"   DA: name %v -> val %q - %q",
						name, val, page.CounterProgress,
					)
				}

				if val == "0" {
					for i1 := 0; i1 < len(page.Groups); i1++ {
						for i2 := 0; i2 < len(page.Groups[i1].Inputs); i2++ {
							if strings.HasPrefix(page.Groups[i1].Inputs[i2].Name, pfx) {
								page.Groups[i1].Inputs[i2].Disabled = true
								if page.Groups[i1].Inputs[i2].Type == "range" {
									page.Groups[i1].Inputs[i2].Response = "no answ."
								}
							}
						}
					}
				}

			}
		}

	}

}
