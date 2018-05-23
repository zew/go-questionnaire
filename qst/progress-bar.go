package qst

import (
	"bytes"
	"fmt"
)

const (
	htmlExample = `
		<ol class="progress progress--large">
			<li class="is-complete" data-step="1">
				Page 1
			</li>
			<li class="is-complete" data-step="2">
				Page 2
			</li>
			<li class="is-active" data-step="3">
				Page 3
			</li>
			<li data-step="4" class="progress__last">
				Confirm
			</li>
		</ol>
`
)

// ProgressBar generates a discrete position indicator
// https://fribly.com/2015/01/01/scalable-responsive-progress-bar/
func (q *QuestionaireT) ProgressBar() string {
	b := bytes.Buffer{}

	b.WriteString(fmt.Sprintf("<ol class='progress'>"))

	for idx, p := range q.Pages {
		liClass := ""
		if idx < q.CurrPage {
			liClass = "is-complete"
		} else if idx == q.CurrPage {
			liClass = "is-active"
		}
		b.WriteString(
			fmt.Sprintf("<li class='%v' data-step='%v'>%v</li> \n",
				liClass, idx+1, p.Label.Tr(q.LangCode),
			),
		)

	}
	b.WriteString(fmt.Sprintf("</ol>"))

	// b.WriteString(fmt.Sprintf(htmlExample))

	return b.String()
}
