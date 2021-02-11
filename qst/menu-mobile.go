package qst

import (
	"bytes"
	"fmt"

	"github.com/zew/go-questionnaire/cfg"
)

// MenuMobile renders the progress bar and the
// footer links such as imprint and language chooser.
// Compare ProgressBar()
func (q *QuestionnaireT) MenuMobile() string {

	b := bytes.Buffer{}

	b.WriteString(`<input type="hidden" name="page" value="-1" >`)

	b.WriteString("<ul class='navigation'>")

	lbl1 := q.Survey.Name.TrSilent(q.LangCode)
	lbl2 := q.Survey.WaveIDPretty()

	fmt.Fprintf(&b, `
		<li style='
				font-size: 120%%; 
				padding-left:   44px; 
				padding-top:    5px; 
				padding-bottom: 8px; 
				aaborder-bottom: 2px solid #666; 
			'
		>
			<!-- space for the menu button -->
			%v - %v
		</li>
`, lbl1, lbl2)

	for idx, p := range q.Pages {

		if p.NoNavigation {
			continue
		}

		eff := p.Label.Tr(q.LangCode)
		if p.Short != nil { // short label dedicated to menu
			eff = p.Short.Tr(q.LangCode)
		}

		liClass := ""
		if idx < q.CurrPage {
			liClass = "is-complete"
		} else if idx == q.CurrPage {
			liClass = "is-active"
		}

		//
		// Make hyperlinks to the pages.
		// See ProgressBar() comment for Java Script intricacies.
		onclick := fmt.Sprintf(` onclick="document.forms.frmMain.page.value='%v';document.forms.frmMain.submit();" `, idx)
		pointr := " style='cursor:pointer' "
		if cfg.Get().AllowSkipForward == false && idx > q.CurrPage {
			onclick = ""
			pointr = ""
		}

		fmt.Fprintf(&b, `
			<li 
				%v %v
				class=" %v" data-step="%v">
				<span style='display: inline-block; line-height: 95%%;'>
					%v %v
				<span>
			</li> 
			`,
			onclick, pointr,
			liClass, p.NavigationalNum,
			// sect, eff,
			eff, "",
		)

	}

	imp := fmt.Sprintf("<a href='%v'>%v</a>",
		cfg.Pref("/doc/site-imprint.md"),
		cfg.Get().Mp["imprint"].Tr(q.LangCode),
	)

	fmt.Fprintf(&b, `
		<li class="imprint" >
			%v
		</li> 
	`, imp)

	fmt.Fprintf(&b, "</ul>")

	return b.String()
}
