package qst

import (
	"bytes"
	"fmt"

	"github.com/zew/go-questionaire/cfg"
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
			// onclick and style added - to make hyperlinks to the pages
			fmt.Sprintf(`
					<li 
						onclick="location.href='%v?page=%v';" style="cursor:pointer"  
						class="%v" data-step="%v">
						%v
					</li> 
				`,
				cfg.Pref(""), idx, liClass, idx+1, p.Label.Tr(q.LangCode),
			),
		)

	}
	b.WriteString(fmt.Sprintf("</ol>"))

	// b.WriteString(fmt.Sprintf(htmlExample))

	return b.String()
}
