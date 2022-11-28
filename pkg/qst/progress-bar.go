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

	progressItems := []int{}
	for idx := range q.Pages {
		if !q.IsInNavigation(idx) {
			continue
		}
		if q.Pages[idx].SuppressInProgressbar {
			continue
		}
		progressItems = append(progressItems, idx)
	}
	progressItems = append(progressItems, 1000*1000)
	// log.Printf("progressItems %+v", progressItems)

	pbActive := 100
	boundLower := 0
	for i, boundUpper := range progressItems {
		// log.Printf("checking q.CurrPage %v is between [%v,%v] => activePBItem %v", q.CurrPage, boundLower, boundUpper, pbActive)
		if q.CurrPage >= boundLower && q.CurrPage < boundUpper {
			pbActive = i - 1 // -1 because we iterate over the max bounds
			// log.Printf("    q.CurrPage %v is between [%v,%v] => activePBItem %v - progressItems %+v", q.CurrPage, boundLower, boundUpper, pbActive, progressItems)
			break
		}
		boundLower = boundUpper
	}

	pbCurr := -1 // progress bar item number
	for idx, p := range q.Pages {

		if !q.IsInNavigation(idx) {
			continue
		}
		if q.Pages[idx].SuppressInProgressbar {
			continue
		}
		pbCurr++
		completeOrActive := ""

		if pbCurr < pbActive {
			completeOrActive = "is-complete"
		} else if pbCurr == pbActive {
			completeOrActive = "is-active"
		}
		// else default:  is-in-future

		shortLbl := p.Short.TrSilent(q.LangCode)

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

		counterInCircle := fmt.Sprintf("%v", pbCurr+1)
		if q.Pages[idx].CounterProgress != "" {
			counterInCircle = q.Pages[idx].CounterProgress
		}
		if q.Pages[idx].CounterProgress == "-" {
			counterInCircle = ""
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
				completeOrActive,
				counterInCircle,
				shortLbl,
			),
		)

	}
	b.WriteString(fmt.Sprintf("</ol>"))

	return b.String()
}
