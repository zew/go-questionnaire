// Package qst implements a four levels deep nested structure
// with input controls, groups, pages and questionaire.
package qst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"time"

	"github.com/zew/go-questionaire/ctr"
)

var implementedTypes = map[string]interface{}{
	"text":     nil,
	"checkbox": nil, // A standalone checkbox - as a group, see below

	// radiogroup and checkboxgroup have the same input name
	"radiogroup":    nil, // A standalone radio makes no sense; only a radiogroup.
	"checkboxgroup": nil, // checkboxgroup has no *sensible* use case. There was an 'amenities' array in another app, with encodings: 4 for bath, 8 for balcony... They should better be designed as independent checkboxes bath and balcony. I cannot think of any useful 'additive flags', and those would have to be added and decoded server side. We keep the type, but untested.

	// Helpers
	"textblock": nil, // Only name and description are rendered
}

// checkbox inputs need standardized values for unchecked and checked
const valEmpty = "0"
const valSet = "1"
const vspacer = "<div class='go-quest-vspacer'></div>\n"

// Special subtype of inputT; used for radiogroup
type radioT struct {
	Right bool      `json:"right,omitempty"` // label and description right of input, default left, similar setting for radioT but not for group
	Label transMapT `json:"label,omitempty"`
	Val   string    `json:"val,omitempty"` // Val is allowed to be nil; it then gets initialized to 1...n by Validate(). 0 indicates 'no entry'.

	// Notice the absence of Response;

}

// Input represents a single form input element.
// There is one exception for multiple radios (radiogroup) with the same name but distinct values.
// Multiple checkboxes (checkboxgroup) with same name but distinct values are a dubious instrument. See comment to implementedType checkboxgroup.
type inputT struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	Radios []radioT `json:"radios,omitempty"` // This slice implements the radiogroup - and the senseless checkboxgroup

	Right   bool      `json:"right,omitempty"` // label and description right of input, default left, similar setting for radioT but not for group
	Label   transMapT `json:"label,omitempty"`
	Desc    transMapT `json:"description,omitempty"`
	ColSpan int       `json:"col_span,omitempty"` // How many table cells in overall layout should the control occupy, counts against group.Cols

	// Validator func() bool `json:"empty"`
	ErrMsg transMapT `json:"err_msg,omitempty"`

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

func renderLabelDescription(langCode string, right bool, lbl, desc transMapT) string {
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
	if right {
		ret = fmt.Sprintf(
			"<span class='go-quest-label-right'><b>%v</b> %v </span>\n",
			e1, e2,
		)
	} else {
		ret = fmt.Sprintf(
			"<span class='go-quest-label'     ><b>%v</b> %v  </span>\n",
			e1, e2,
		)
	}
	return ret
}

// Rendering one input to HTML
// func (i inputT) HTML(langCode string, namePrefix string) string {
func (i inputT) HTML(langCode string) string {

	nm := i.Name

	ctrl := ""
	switch i.Type {
	case "textblock":
		lbl := renderLabelDescription(langCode, i.Right, i.Label, i.Desc)
		return ctrl + lbl
	case "radiogroup", "checkboxgroup":

		innerType := "radio"
		if i.Type == "checkboxgroup" {
			innerType = "checkbox"
		}
		for _, rad := range i.Radios {
			checked := ""
			if i.Response == rad.Val {
				checked = "checked=\"checked\""
			}
			// ctrl += fmt.Sprintf("Val %v", val)

			if !rad.Right && rad.Label != nil {
				ctrl += fmt.Sprintf("<span class='go-quest-label-right'>%v</span>\n", rad.Label.Tr(langCode))
			}

			ctrl += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
				innerType, nm, nm, i.Label.Tr(langCode), i.Desc.Tr(langCode), rad.Val, checked)

			if rad.Right && rad.Label != nil {
				ctrl += fmt.Sprintf("<span class='go-quest-label'>%v</span>\n", rad.Label.Tr(langCode))
			}

		}
		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if innerType == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='%v' />\n",
				nm, nm, valEmpty)
		}

	case "text", "checkbox":

		val := i.Response

		checked := ""
		if i.Type == "checkbox" {
			if val == valSet {
				checked = "checked=\"checked\""
			}
			val = valSet
		}

		ctrl += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
			i.Type, nm, nm, i.Label.Tr(langCode), i.Desc.Tr(langCode), val, checked)

		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if i.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

	default:
		ctrl += fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, i.Type, implementedTypes)
	}

	lbl := renderLabelDescription(langCode, i.Right, i.Label, i.Desc)
	lbl = fmt.Sprintf("<label for='%v'>%v</label>\n", nm, lbl)
	return lbl + ctrl

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
	Cols     int  `json:"columns,omitempty"`  // number of vertical columns; for horizontal *and* vertical layouts

	Members []inputT
}

// Rendering a group of inputs to HTML
func (gr groupT) HTML(langCode string) string {

	b := bytes.Buffer{}

	b.WriteString("<div class='go-quest-group' >\n")

	lbl := renderLabelDescription(langCode, false, gr.Label, gr.Desc)

	b.WriteString(lbl)
	b.WriteString(vspacer)

	cols := 0 // cols counter
	for _, mem := range gr.Members {
		b.WriteString(mem.HTML(langCode))

		log.Printf("%12v %2v %2v", mem.Type, cols, len(mem.Radios))
		if gr.Cols > 0 {
			if cols > 0 && (cols+1)%gr.Cols == 0 || cols == len(gr.Members)-1 {
				b.WriteString(vspacer) // new row  - or end of group
			}
			if mem.ColSpan > 0 {
				cols += mem.ColSpan // larger input controls
			} else if len(mem.Radios) > 0 {
				cols += len(mem.Radios) // radiogroups, if no ColSpan is set
			} else {
				cols++ // default: every input occupies one column width
			}
		}
	}
	b.WriteString("</div>\n")

	return b.String()

}

// Type page is collection of tInputs and some meta data
type pageT struct {
	Label transMapT `json:"label,omitempty"`
	Desc  transMapT `json:"description,omitempty"`

	Elements []groupT  `json:"elements,omitempty"`
	Finished time.Time `json:"finished,omitempty"`
}

func newPage() pageT {
	t := pageT{
		Label: transMapT{"en": "Page Label", "de": fmt.Sprintf("Seitentitel_%v", ctr.Increment())},
		Desc:  transMapT{"en": "Page Description", "de": "Seitenbeschreibung"},
	}
	return t
}

// Type QuestionaireT contains pages with inputs
type QuestionaireT struct {
	Pages     []pageT           `json:"pages,omitempty"`
	LangCodes map[string]string `json:"lang_codes"` // all possible lang codes - i.e. en, de
	LangCode  string            `json:"lang_code"`  // default lang code - and current lang code - i.e. de

	CurrPage int `json:"curr_page,omitempty"`
}

// CurrentPageHTML is a comfort shortcut to PageHTML
func (q *QuestionaireT) CurrentPageHTML() (string, error) {
	return q.PageHTML(q.CurrPage)
}

// PageHTML generates HTML for a specific page of the questionaire
func (q *QuestionaireT) PageHTML(idx int) (string, error) {

	if q.CurrPage > len(q.Pages) || q.CurrPage < 0 {
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

	str := fmt.Sprintf("<h3>%v</h3>", p.Label.Tr(q.LangCode))
	str += fmt.Sprintf("<p>%v</p>", p.Desc.Tr(q.LangCode))
	for i := 0; i < len(p.Elements); i++ {
		str += p.Elements[i].HTML(q.LangCode) + "\n"
		str += vspacer
	}
	return str, nil
}

// Load loads a questionaire from a JSON file.
func Load(fn string) (*QuestionaireT, error) {
	q := QuestionaireT{}

	bts, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatalf("Could not read file %v : %v", fn, err)
		return &q, err
	}

	err = json.Unmarshal(bts, &q)
	if err != nil {
		log.Fatalf("Unmarshal %v: %v", fn, err)
		return &q, err
	}

	return &q, nil
}

func (q *QuestionaireT) HasPrev() bool {
	if q.CurrPage > 0 {
		return true
	}
	return false
}
func (q *QuestionaireT) Prev() int {
	if q.CurrPage > 0 {
		return q.CurrPage - 1
	}
	return 0
}

func (q *QuestionaireT) HasNext() bool {
	if q.CurrPage < len(q.Pages)-1 {
		return true
	}
	return false
}

func (q *QuestionaireT) Next() int {
	if q.CurrPage < len(q.Pages)-1 {
		return q.CurrPage + 1
	}
	return len(q.Pages)
}

var not09azHyphenUnderscore = regexp.MustCompile(`[^a-z0-9\_\-]+`)

// Example
func Mustaz09_(s string) bool {
	if not09azHyphenUnderscore.MatchString(s) {
		return false
	}
	return true
}