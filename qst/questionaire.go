// Package qst implements a four levels deep nested structure
// with input controls, groups, pages and questionaire;
// contains HTML rendering, page navigation,
// loading/saving from/to JSON file, consistence validation,
// multi-language support.
package qst

import (
	"bytes"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/zew/go-questionaire/cfg"
	"github.com/zew/go-questionaire/ctr"
)

// Special subtype of inputT; used for radiogroup
type radioT struct {
	HAlign horizontalAlignment `json:"horizontal_align,omitempty"` // label and description left/center/right of input, default left, similar setting for radioT but not for group
	Label  transMapT           `json:"label,omitempty"`
	Val    string              `json:"val,omitempty"` // Val is allowed to be nil; it then gets initialized to 1...n by Validate(). 0 indicates 'no entry'.
	// Notice the absence of Response;
}

// Input represents a single form input element.
// There is one exception for multiple radios (radiogroup) with the same name but distinct values.
// Multiple checkboxes (checkboxgroup) with same name but distinct values are a dubious instrument. See comment to implementedType checkboxgroup.
type inputT struct {
	Name          string              `json:"name,omitempty"`
	Type          string              `json:"type,omitempty"`
	HAlignControl horizontalAlignment `json:"horizontal_align_control,omitempty"` // label       left/center/right of input, default left, similar setting for radioT but not for group
	HAlignLabel   horizontalAlignment `json:"horizontal_align_label,omitempty"`   // description left/center/right of input, default left, similar setting for radioT but not for group
	Label         transMapT           `json:"label,omitempty"`
	Desc          transMapT           `json:"description,omitempty"`
	ColSpan       int                 `json:"col_span,omitempty"` // How many table cells in overall layout should the control occupy, counts against group.Cols

	Radios []radioT `json:"radios,omitempty"` // This slice implements the radiogroup - and the senseless checkboxgroup

	Validator string    `json:"validator,omitempty"` // i.e. inRange20 - any string from validators
	ErrMsg    transMapT `json:"err_msg,omitempty"`

	// These are only useful a part of wave-data
	Response      string  `json:"response,omitempty"`
	ResponseFloat float64 `json:"response_float,omitempty"` // also for integers
}

// Returns an input filled in with globally enumerated label, decription etc.
func newInput() inputT {
	id := ctr.Increment()
	t := inputT{
		Name:  fmt.Sprintf("input_%v", id),
		Type:  "text",
		Label: transMapT{"en": fmt.Sprintf("Label %v", id), "de": fmt.Sprintf("Titel %v", id)},
		Desc:  transMapT{"en": "Description", "de": "Beschreibung"},
	}
	return t
}

// Argument numCols computes precise width in percent
func renderLabelDescription(langCode string, hAlign horizontalAlignment, lbl, desc transMapT, css string, numCols int) string {
	ret := ""
	if lbl == nil && desc == nil {
		return ret
	}
	e1 := lbl.Tr(langCode)
	if lbl == nil {
		e1 = "" // Suppress "Translation map not initialized." here
	}
	e2 := desc.Tr(langCode)
	if desc == nil {
		e2 = "" // Suppress "Translation map not initialized." here
	}
	ret = fmt.Sprintf(
		"<span class='go-quest-label %v' ><b>%v</b> %v </span>\n", css, e1, e2,
	)
	ret = fmt.Sprintf("<span class='go-quest-cell-%v'  style='%v'>%v</span>\n", hAlign, colWidth(numCols), ret)
	return ret
}

// Rendering one input to HTML
// func (i inputT) HTML(langCode string, namePrefix string) string {
func (i inputT) HTML(langCode string, numCols int) string {

	nm := i.Name

	switch i.Type {
	case "textblock":
		lbl := renderLabelDescription(langCode, i.HAlignLabel, i.Label, i.Desc, "", numCols)
		return lbl

	case "radiogroup", "checkboxgroup":
		ctrl := ""
		innerType := "radio"
		if i.Type == "checkboxgroup" {
			innerType = "checkbox"
		}
		for _, rad := range i.Radios {
			one := ""
			checked := ""
			if i.Response == rad.Val {
				checked = "checked=\"checked\""
			}
			// one += fmt.Sprintf("Val %v", val)

			if rad.Label != nil && rad.HAlign == HLeft {
				one += fmt.Sprintf("<span class='go-quest-label'>%v</span>\n", rad.Label.Tr(langCode))
			}
			if rad.Label != nil && rad.HAlign == HCenter {
				one += fmt.Sprintf("<span class='go-quest-label'>%v</span>\n", rad.Label.Tr(langCode))
				one += vspacer
			}

			one += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
				innerType, nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode), rad.Val, checked)

			if rad.Label != nil && rad.HAlign == HRight {
				one += fmt.Sprintf("<span class='go-quest-label'>%v</span>\n", rad.Label.Tr(langCode))
			}
			one = fmt.Sprintf("<span class='go-quest-cell-%v' style='%v'>%v</span>\n", rad.HAlign, colWidth(numCols), one)
			ctrl += one
		}
		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if innerType == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='%v' />\n",
				nm, nm, valEmpty)
		}

		ctrl += i.ErrMsg.TrSilent(langCode) // ugly layout  - but radiogroup and checkboxgroup won't have validation errors anyway

		lbl := renderLabelDescription(langCode, i.HAlignLabel, i.Label, i.Desc, "", numCols)
		// lbl = fmt.Sprintf("<label for='%v'>%v</label>\n", nm, lbl)
		return lbl + ctrl

	case "text", "checkbox":
		ctrl := ""
		val := i.Response

		checked := ""
		if i.Type == "checkbox" {
			if val == valSet {
				checked = "checked=\"checked\""
			}
			val = valSet
		}
		ctrl += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
			i.Type, nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode), val, checked)

		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if i.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

		ctrl += i.ErrMsg.TrSilent(langCode)

		ctrl = fmt.Sprintf("<span class='go-quest-cell-%v' style='%v'>%v</span>\n", i.HAlignControl, colWidth(numCols), ctrl)

		lbl := renderLabelDescription(langCode, i.HAlignLabel, i.Label, i.Desc, "", numCols)
		lbl = fmt.Sprintf("<label for='%v'>%v</label>\n", nm, lbl)
		return lbl + ctrl

	default:
		return fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, i.Type, implementedTypes)
	}

}

// A group consists of several input controls.
// It contains no response information.
// It can bundle checkbox or text inputs with *distinct* names.
// Whereas: radiogroup and checkboxgroup have the *same* name and a single response.
// A group is a layout unit with a configurable number of columns.
type groupT struct {
	// Name  string
	Label transMapT `json:"label,omitempty"`
	Desc  transMapT `json:"description,omitempty"`

	Vertical bool `json:"vertical,omitempty"` // groups vertically, not horizontally
	Cols     int  `json:"columns,omitempty"`  // number of vertical columns; for horizontal *and* vertical layouts; you must count the labels too

	Inputs []inputT `json:"inputs,omitempty"`
}

// Rendering a group of inputs to HTML
func (gr groupT) HTML(langCode string) string {

	b := bytes.Buffer{}

	b.WriteString("<div class='go-quest-group' >\n")

	lbl := renderLabelDescription(langCode, HLeft, gr.Label, gr.Desc, "go-quest-group-header", gr.Cols)

	b.WriteString(lbl)
	b.WriteString(vspacer)

	cols := 0 // cols counter
	for i, mem := range gr.Inputs {
		b.WriteString(mem.HTML(langCode, gr.Cols))

		if gr.Cols > 0 {
			if mem.ColSpan > 0 {
				cols += mem.ColSpan // larger input controls
			} else if len(mem.Radios) > 0 {
				cols += len(mem.Radios) // radiogroups, if no ColSpan is set
				if mem.Label != nil || mem.Desc != nil {
					cols++ // plus one for the leading label of the input
				}
			} else {
				cols++ // default: every input occupies one column width
			}

			// log.Printf("%12v %2v %2v", mem.Type, cols, cols%gr.Cols)

			if (cols+0)%gr.Cols == 0 || i == len(gr.Inputs)-1 {
				b.WriteString(vspacer) // end of row  - or end of group
			}

		}
	}
	b.WriteString("</div>\n")
	return b.String()

}

// Type page contains groups with inputs
type pageT struct {
	Label transMapT `json:"label,omitempty"`
	Desc  transMapT `json:"description,omitempty"`

	Groups []groupT `json:"groups,omitempty"`

	Finished time.Time `json:"finished,omitempty"`
}

func newPage() pageT {
	t := pageT{
		Label: transMapT{"en": "Page Label", "de": fmt.Sprintf("Seitentitel_%v", ctr.Increment())},
		Desc:  transMapT{"de": "", "en": ""},
	}
	return t
}

// QuestionaireT contains pages with groups with inputs
type QuestionaireT struct {
	Pages     []pageT           `json:"pages,omitempty"`
	LangCodes map[string]string `json:"lang_codes"` // all possible lang codes - i.e. en, de
	LangCode  string            `json:"lang_code"`  // default lang code - and current lang code - i.e. de

	CurrPage int `json:"curr_page,omitempty"`
}

// LanguageChooser renders a HTML language chooser
func (q *QuestionaireT) LanguageChooser() string {
	s := []string{}
	for key, lang := range q.LangCodes {
		keyCap := strings.Title(key)
		if q.LangCode == "en" {
			keyCap = key
		}
		if key == q.LangCode {
			s = append(s, fmt.Sprintf("<b           title='%v'>%v</b>\n", lang, keyCap))
		} else {
			uri := cfg.Pref("/") + "?lang_code=" + key
			s = append(s, fmt.Sprintf("<a href='%v' title='%v'>%v</a>\n", uri, lang, keyCap))
		}
	}
	return strings.Join(s, "  |  ")
}

// CurrentPageHTML is a comfort shortcut to PageHTML
func (q *QuestionaireT) CurrentPageHTML() (string, error) {
	return q.PageHTML(q.CurrPage)
}

// PageHTML generates HTML for a specific page of the questionaire
func (q *QuestionaireT) PageHTML(idx int) (string, error) {

	if q.CurrPage > len(q.Pages)-1 || q.CurrPage < 0 {
		s := fmt.Sprintf("You requested page %v out of %v. Page does not exist", idx, len(q.Pages)-1)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	p := q.Pages[idx]

	if _, ok := q.LangCodes[q.LangCode]; !ok || q.LangCode == "" {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	b := bytes.Buffer{}

	b.WriteString(fmt.Sprintf("<h3 class='go-quest-page-header' >%v</h3>", p.Label.Tr(q.LangCode)))
	b.WriteString(vspacer)
	b.WriteString(fmt.Sprintf("<p  class='go-quest-page-desc'>%v</p>", p.Desc.Tr(q.LangCode)))
	b.WriteString(vspacer16)

	for i := 0; i < len(p.Groups); i++ {
		b.WriteString(p.Groups[i].HTML(q.LangCode) + "\n")
		b.WriteString(vspacer16)
		if i < len(p.Groups)-1 { // no vertical distance at the end of groups
			b.WriteString(vspacer16)
			b.WriteString(vspacer16)
		}
	}
	return b.String(), nil
}

// HasPrev if a previous page exists
func (q *QuestionaireT) HasPrev() bool {
	if q.CurrPage > 0 {
		return true
	}
	return false
}

// Prev returns number of the previous page
func (q *QuestionaireT) Prev() int {
	if q.CurrPage > 0 {
		return q.CurrPage - 1
	}
	return 0
}

// HasNext if a next page exists
func (q *QuestionaireT) HasNext() bool {
	if q.CurrPage < len(q.Pages)-1 {
		return true
	}
	return false
}

// Next returns number of the next page
func (q *QuestionaireT) Next() int {
	if q.CurrPage < len(q.Pages)-1 {
		return q.CurrPage + 1
	}
	return len(q.Pages)
}
