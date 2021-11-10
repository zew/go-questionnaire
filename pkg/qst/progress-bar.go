package qst

import (
	"fmt"
	"strings"
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
//
// It should be clickable and jump to the indicated page.
// This means, we have to submit the form and submit the destination page.
// Compare MenuMobile()
func (q *QuestionnaireT) ProgressBar() string {

	b := &strings.Builder{}
	// b.WriteString(fmt.Sprintf(htmlExample))

	b.WriteString("\n")
	b.WriteString(fmt.Sprintf("\t\t\t\t<ol class='progress'>\n"))
	b.WriteString(fmt.Sprintf("\t\t\t\t\t<input type='hidden' name='page' value='-1' >\n"))

	for idx, p := range q.Pages {

		if !q.isNavigation(idx) {
			continue
		}

		completeOrActive := ""
		if idx < q.CurrPage {
			completeOrActive = "is-complete"
		} else if idx == q.CurrPage {
			completeOrActive = "is-active"
		}

		eff := p.Short.TrSilent(q.LangCode)

		/*
			<li> elements become hyperlinks to the questionnaire pages

			This would not post current page form data:
					location.href='%v?page=%v
			Thus:
					this.form.page.value='%v';
					this.form.submit();  // does not work for <li> element
					document.forms.frmMain.submit() // instead
			Debug with
					console.log('document.forms.frmMain.page.value: ',document.forms.frmMain.page.value);
		*/
		onclick := fmt.Sprintf(` onclick="document.forms.frmMain.page.value='%v';document.forms.frmMain.submit();" `, idx)
		pointr := " style='cursor:pointer' "
		if q.PreventSkipForward && idx > q.CurrPage {
			onclick = ""
			pointr = ""
		}

		b.WriteString(
			fmt.Sprintf(`
					<li 
						%v %v
						class="%v" data-step="%v">
						<span  class="progress-bar-label" >
							%v
						<span>
					</li> 
				`,
				onclick, pointr,
				completeOrActive, p.NavigationalNum,
				eff,
			),
		)

	}
	b.WriteString(fmt.Sprintf("</ol>"))

	return b.String()
}
