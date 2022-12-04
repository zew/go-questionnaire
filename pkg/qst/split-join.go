package qst

import (
	"fmt"
)

// Split creates a copy of q containing only the user responses
// and q metadata
func (q *QuestionnaireT) Split() (*QuestionnaireT, error) {
	q2 := *q
	q2.Pages = nil
	for i1 := 0; i1 < len(q.Pages); i1++ {
		p2 := q2.AddPage()
		p2.Finished = q.Pages[i1].Finished
		p2.Label = q.Pages[i1].Label                         // for debugging
		p2.GeneratorFuncName = q.Pages[i1].GeneratorFuncName // essential for joining later
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			if q.Pages[i1].Groups[i2].ID == "footer" {
				continue
			}
			gr := p2.AddGroup()
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				// log.Printf("Added p%02v  gr%02v  inp%02v", i1, i2, i3)
				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				inp2 := gr.AddInput()
				if inp.IsLayout() {
					inp2.Name = inp.Name
					inp2.Type = inp.Type
					// for debugging
					if inp2.Type == "dyn-textblock" || inp2.Type == "dyn-composite" {
						inp2.Label = inp.Label.Left(40)
					}
					continue
				}
				inp2.ErrMsg = inp.ErrMsg
				inp2.Name = inp.Name
				inp2.Response = inp.Response
				inp2.Type = inp.Type
			}
		}
	}
	return &q2, nil
}

// Join adds user input from q2 onto q
func (q *QuestionnaireT) Join(q2 *QuestionnaireT) error {

	// preflight: check identical structures of both questionnaires
	if len(q.Pages) != len(q2.Pages) {
		return fmt.Errorf("qBase has %v pages - q2 %v", len(q.Pages), len(q2.Pages))
	}
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			if q.Pages[i1].GeneratorFuncName != "" {
				// see DynamicPages()
				continue
			}
			if len(q.Pages[i1].Groups) != len(q2.Pages[i1].Groups) {
				return fmt.Errorf("qBase page %v has %v groups - q2 %v", i1, len(q.Pages[i1].Groups), len(q2.Pages[i1].Groups))
			}
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				// log.Printf("adding p%02v  gr%02v  inp%02v", i1, i2, i3)
				if len(q.Pages[i1].Groups[i2].Inputs) != len(q2.Pages[i1].Groups[i2].Inputs) {
					return fmt.Errorf("qBase page %v group %v has %v inputs - q2 %v", i1, i2, len(q.Pages[i1].Groups[i2].Inputs), len(q2.Pages[i1].Groups[i2].Inputs))
				}
				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				inp2 := q2.Pages[i1].Groups[i2].Inputs[i3]
				if inp.Name != inp2.Name {
					return fmt.Errorf("qBase page %v group %v inp %v has name %v - q2 %v", i1, i2, i3, inp.Name, inp2.Name)
				}
				if inp.Type != inp2.Type {
					return fmt.Errorf("qBase page %v group %v inp %v has type %v - q2 %v", i1, i2, i3, inp.Type, inp2.Type)
				}
			}
		}
	}

	q.UserID = q2.UserID
	q.ClosingTime = q2.ClosingTime
	q.RemoteIP = q2.RemoteIP
	q.UserAgent = q2.UserAgent
	q.LangCode = q2.LangCode
	q.CurrPage = q2.CurrPage
	q.HasErrors = q2.HasErrors
	q.VersionEffective = q2.VersionEffective

	attrs := map[string]string{}
	for k, v := range q2.Attrs {
		attrs[k] = v
	}
	q.Attrs = attrs

	//
	// transfer input values
	for i1 := 0; i1 < len(q.Pages); i1++ {
		q.Pages[i1].Finished = q2.Pages[i1].Finished
		// log.Printf("\tSetting q.Pages[%v].Finished to %v", i1, q2.Pages[i1].Finished)
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				// log.Printf("adding p%02v  gr%02v  inp%02v", i1, i2, i3)
				inp := q.Pages[i1].Groups[i2].Inputs[i3]
				if inp.IsLayout() {
					continue
				}
				inp2 := q2.Pages[i1].Groups[i2].Inputs[i3]
				q.Pages[i1].Groups[i2].Inputs[i3].ErrMsg = inp2.ErrMsg
				q.Pages[i1].Groups[i2].Inputs[i3].Response = inp2.Response
			}
		}
	}

	return nil
}
