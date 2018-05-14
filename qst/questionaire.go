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

// For all multi lingual strings.
// Contains one value for each language code.
type transMapT map[string]string // type translation map

// Tr translates
func (t transMapT) Tr(langCode string) string {
	if val, ok := t[langCode]; ok {
		return val
	}
	if val, ok := t["en"]; ok {
		return val
	}
	for _, val := range t {
		return val
	}
	if t == nil {
		return "Missing translation. Map not initialized."
	}
	return "Missing translation."
}

var implementedTypes = map[string]interface{}{
	"text":     nil,
	"checkbox": nil, // A standalone checkbox - as a group, see below

	// radiogroup and checkboxgroup have the same input name
	"radiogroup":    nil, // A standalone radio makes no sense; only a radiogroup
	"checkboxgroup": nil, // checkboxgroup has no use case. For example OR flags such as 4 - bath, 8 - balcony. They should better be designed as independent checkboxes bath and balcony
}

// For radio and for checkbox inputs
const valEmpty = "0"
const valSet = "1"
const vspacer = "<div class='go-quest-vspacer'></div>\n"

// Special for radiogroup
type radioT struct {
	Val   string    `json:"val,omitempty"`
	Label transMapT `json:"label,omitempty"`
	Right bool      `json:"right,omitempty"` // label and description right of input, default left
}

type inputT struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	Radios []radioT `json:"radios,omitempty"`

	Label transMapT `json:"label,omitempty"`
	Desc  transMapT `json:"description,omitempty"`
	Right bool      `json:"right,omitempty"` // label and description right of input, default left
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

// Rendering one input to HTML
// func (i inputT) HTML(langCode string, namePrefix string) string {
func (i inputT) HTML(langCode string, cols int) string {

	nm := i.Name

	ctrl := ""
	switch i.Type {
	case "radiogroup", "checkboxgroup":

		innerType := "radio"
		if i.Type == "checkboxgroup" {
			innerType = "checkbox"
		}
		for i2, rad := range i.Radios {
			checked := ""
			if i.Response == rad.Val {
				checked = "checked=\"checked\""
			}
			// ctrl += fmt.Sprintf("Val %v", val)

			if !rad.Right {
				ctrl += fmt.Sprintf("<span class='go-quest-label-right'>%v</span>\n", rad.Label.Tr(langCode))
			}

			ctrl += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
				innerType, nm, nm, i.Label.Tr(langCode), i.Desc.Tr(langCode), rad.Val, checked)

			if rad.Right {
				ctrl += fmt.Sprintf("<span class='go-quest-label'>%v</span>\n", rad.Label.Tr(langCode))
			}

			if cols > 0 && (i2+1)%cols == 0 && i2 > 0 || i2 == len(i.Radios)-1 {
				ctrl += vspacer
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

	if i.Right {
		lbl := fmt.Sprintf("<span class='go-quest-label-right'><label for='%v'>  <b>%v</b> %v </label></span>\n", nm, i.Label.Tr(langCode), i.Desc.Tr(langCode))
		return ctrl + lbl
	}
	lbl := fmt.Sprintf("<span class='go-quest-label'><label for='%v'>  <b>%v</b> %v </label></span>\n", nm, i.Label.Tr(langCode), i.Desc.Tr(langCode))
	return lbl + ctrl

}

// A group consists of several input controls
// It can bundle checkbox or text inputs with *distinct* names.
// Whereas: radiogroup and checkboxgroup have the *same* name.
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

	lbl := fmt.Sprintf("<span class='go-quest-group-header'><b>%v</b> %v</span> \n", gr.Label.Tr(langCode), gr.Desc.Tr(langCode))
	b.WriteString(lbl)
	b.WriteString(vspacer)

	for i, mem := range gr.Members {
		b.WriteString(mem.HTML(langCode, gr.Cols))
		if gr.Cols > 0 && (i+1)%gr.Cols == 0 && i > 0 {
			b.WriteString(vspacer)
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
