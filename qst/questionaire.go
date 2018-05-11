package qst

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/zew/go-questionaire/ctr"
)

type transMapT map[string]string // type translation map

// Translate
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
	return "missing_translation"
}

// a standalone radio makes no sense; only radiogroup
var implementedTypes = map[string]interface{}{"checkbox": nil, "text": nil, "radiogroup": nil, "checkboxgroup": nil}

const valEmpty = "0"
const valSet = "1"

type inputT struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`

	Vals []string `json:"vals,omitempty"` // only for radiogroup or checkboxgroup

	Title transMapT `json:"title,omitempty"`
	Desc  transMapT `json:"description,omitempty"`
	// Validator func() bool `json:"empty"`
	ErrMsg transMapT `json:"err_msg,omitempty"`

	// These are only useful a part of wave-data
	Response      string  `json:"response,omitempty"`
	ResponseFloat float64 `json:"response_float,omitempty"` // also for integers
}

// Returns an input filled in with globally enumerated title, label etc.
func newInput() inputT {
	id := ctr.Increment()
	t := inputT{
		Name:  fmt.Sprintf("input_%v", id),
		Type:  "text",
		Title: transMapT{"en": fmt.Sprintf("Title %v", id), "de": fmt.Sprintf("Titel %v", id)},
		Desc:  transMapT{"en": "Description", "de": "Beschreibung"},
	}
	return t
}

// Rendering one input to HTML
// func (i inputT) HTML(langCode string, namePrefix string) string {
func (i inputT) HTML(langCode string, cols int) string {

	nm := i.Name

	lbl := fmt.Sprintf("<label for='%v'>  <b>%v</b> %v </label>\n", nm, i.Title.Tr(langCode), i.Desc.Tr(langCode))
	ctrl := ""
	switch i.Type {
	case "radiogroup", "checkboxgroup":

		innerType := "radio"
		if i.Type == "checkboxgroup" {
			innerType = "checkbox"
		}
		for i2, val := range i.Vals {
			_ = i2
			checked := ""
			if i.Response == val {
				checked = "checked=\"checked\""
			}
			ctrl += fmt.Sprintf("Val %v", val)
			ctrl += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
				innerType, nm, nm, i.Title.Tr(langCode), i.Desc.Tr(langCode), val, checked)

			if cols > 0 && (i2+1)%cols == 0 && i2 > 0 || i2 == len(i.Vals)-1 {
				ctrl += "<br>\n"
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
			i.Type, nm, nm, i.Title.Tr(langCode), i.Desc.Tr(langCode), val, checked)

		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if i.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

	default:
		ctrl += fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, i.Type, implementedTypes)
	}

	switch i.Type {
	case "text":
		return lbl + ctrl
	case "checkbox", "radio":
		return ctrl + lbl
	default:
		return lbl + ctrl
	}

	return lbl + ctrl
}

type groupT struct {
	// Name  string    `json:"name,omitempty"`
	Title transMapT `json:"title,omitempty"`
	Desc  transMapT `json:"description,omitempty"`

	Vertical bool `json:"vertical,omitempty"` // groups vertically, not horizontally
	Cols     int  `json:"columns,omitempty"`  // number of vertical columns; for horizontal *and* vertical layouts

	Members []inputT
}

// Rendering a group of inputs to HTML
func (gr groupT) HTML(langCode string) string {

	b := bytes.Buffer{}

	b.WriteString("<div style='border: 1px solid #2c2;'>\n")

	lbl := fmt.Sprintf("group:  <b>%v</b> %v <br>\n", gr.Title.Tr(langCode), gr.Desc.Tr(langCode))
	b.WriteString(lbl)

	for i, mem := range gr.Members {
		b.WriteString(mem.HTML(langCode, gr.Cols))
		if gr.Cols > 0 && (i+1)%gr.Cols == 0 && i > 0 {
			b.WriteString("<br>\n")
		}
	}
	b.WriteString("</div>\n")

	return b.String()

}

// Type page is collection of tInputs and some meta data
type pageT struct {
	Title transMapT `json:"title,omitempty"`
	Desc  transMapT `json:"description,omitempty"`

	Elements []groupT  `json:"elements,omitempty"`
	Finished time.Time `json:"finished,omitempty"`
}

func newPage() pageT {
	t := pageT{
		Title: transMapT{"en": "Page Title", "de": "Seitentitel"},
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

// Validate checks whether essential elements of the questionaire exist.
func (q *QuestionaireT) Validate() error {

	if q.LangCode == "" {
		s := fmt.Sprintf("Language code is empty. Must be one of %v", q.LangCodes)
		log.Printf(s)
		return fmt.Errorf(s)
	}
	if _, ok := q.LangCodes[q.LangCode]; !ok {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Printf(s)
		return fmt.Errorf(s)
	}

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Elements); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Elements[i2].Members); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)

				inp := q.Pages[i1].Elements[i2].Members[i3]
				if _, ok := implementedTypes[inp.Type]; !ok {
					return fmt.Errorf(s + fmt.Sprintf("Type %v is not in %v ", inp.Type, implementedTypes))
				}

			}
		}
	}

	names := map[string]int{}
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Elements); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Elements[i2].Members); i3++ {

				s := fmt.Sprintf("Page %v - Group %v - Input %v: ", i1, i2, i3)

				// grp := q.Pages[i1].Elements[i2].Name
				nm := q.Pages[i1].Elements[i2].Members[i3].Name
				// tp := q.Pages[i1].Elements[i2].Members[i3].Type

				if nm == "" {
					return fmt.Errorf(s+"Name %v is empty", nm)
				}

				if not09azHyphenUnderscore.MatchString(nm) {
					return fmt.Errorf(s+"Name %v must consist of [a-z0-9_-]", nm)
				}

				names[nm]++

			}
		}
	}

	for k, v := range names {
		if v > 1 {
			s := fmt.Sprintf("Page element '%v' is not unique  (%v)", k, v)
			log.Printf(s)
			return fmt.Errorf(s)
		}
		if k != strings.ToLower(k) {
			s := fmt.Sprintf("Page element '%v' is not lower case  (%v)", k, v)
			log.Printf(s)
			return fmt.Errorf(s)
		}
	}

	return nil
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

	str := fmt.Sprintf("<h3>%v</h3>", p.Title.Tr(q.LangCode))
	str += fmt.Sprintf("<p>%v</p>", p.Desc.Tr(q.LangCode))
	for i := 0; i < len(p.Elements); i++ {
		str += p.Elements[i].HTML(q.LangCode) + "\n"
		str += "<br>\n"
	}
	return str, nil
}

var not09azHyphenUnderscore = regexp.MustCompile(`[^a-z0-9\_\-]+`)

// Example
func Mustaz09_(s string) bool {
	if not09azHyphenUnderscore.MatchString(s) {
		return false
	}
	return true
}
