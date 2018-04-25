package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type ttransMap map[string]string // type translation map

func (t ttransMap) Tr(langCode string) string {
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

type tInput struct {
	Name  string    `json:"name,omitempty"`
	Type  string    `json:"type,omitempty"`
	Title ttransMap `json:"title,omitempty"`
	Desc  ttransMap `json:"description,omitempty"`
	// Validator func() bool `json:"empty"`
	ErrMsg ttransMap `json:"err_msg,omitempty"`

	// These are only useful a part of wave-data
	Response      string  `json:"response,omitempty"`
	ResponseFloat float64 `json:"response_float,omitempty"` // also for integers
}

func newInput() tInput {
	id := cntr.increment()
	t := tInput{
		Name:  fmt.Sprintf("input_%v", id),
		Type:  "text",
		Title: ttransMap{"en": fmt.Sprintf("Title %v", id), "de": fmt.Sprintf("Titel %v", id)},
		Desc:  ttransMap{"en": "Description", "de": "Beschreibung"},
	}
	return t
}

func (i tInput) HTML(langCode string) string {
	lbl := fmt.Sprintf("<label for='%v'>  <b>%v</b> %v </label>\n", i.Name, i.Title.Tr(langCode), i.Desc.Tr(langCode))
	ctrl := ""
	switch i.Type {
	case "text", "checkbox", "radio":
		checked := ""
		if i.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", i.Name, i.Name)
			if i.Response == "1" {
				checked = "checked"
			}
			i.Response = "1"
		}
		ctrl += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' value='%v' %v />\n",
			i.Type, i.Name, i.Name, i.Title.Tr(langCode), i.Desc.Tr(langCode), i.Response, checked)
	default:
		ctrl += fmt.Sprintf("input %v: unknown type '%v'\n", i.Name, i.Type)
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

type tPage struct {
	Title ttransMap `json:"title,omitempty"`
	Desc  ttransMap `json:"description,omitempty"`

	Elements []tInput  `json:"elements,omitempty"`
	Finished time.Time `json:"finished,omitempty"`
}

func newPage() tPage {
	t := tPage{
		Title: ttransMap{"en": "Page Title", "de": "Seitentitel"},
		Desc:  ttransMap{"en": "Page Description", "de": "Seitenbeschreibung"},
	}
	return t
}

type tQuestionaire struct {
	Pages     []tPage           `json:"pages,omitempty"`
	LangCodes map[string]string `json:"lang_codes"`        // all possible lang codes - i.e. en, de
	LangCode  string            `json:"lang_code_default"` // default lang code - i.e. de

	CurrPage int `json:"curr_page,omitempty"`
}

func (q *tQuestionaire) Validate() error {

	if q.LangCode == "" {
		s := fmt.Sprintf("Language code must be one of %v", q.LangCodes)
		log.Printf(s)
		return fmt.Errorf(s)
	}
	if _, ok := q.LangCodes[q.LangCode]; !ok {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Printf(s)
		return fmt.Errorf(s)
	}

	return nil
}

func LoadQuestionaire(fn string) (*tQuestionaire, error) {
	q := tQuestionaire{}

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

func (q *tQuestionaire) HasPrev() bool {
	if q.CurrPage > 0 {
		return true
	}
	return false
}
func (q *tQuestionaire) Prev() int {
	if q.CurrPage > 0 {
		return q.CurrPage - 1
	}
	return 0
}

func (q *tQuestionaire) HasNext() bool {
	if q.CurrPage < len(q.Pages)-1 {
		return true
	}
	return false
}

func (q *tQuestionaire) Next() int {
	if q.CurrPage < len(q.Pages)-1 {
		return q.CurrPage + 1
	}
	return len(q.Pages)
}

func (q *tQuestionaire) CurrentPageHTML() (string, error) {
	return q.PageHTML(q.CurrPage)
}
func (q *tQuestionaire) PageHTML(idx int) (string, error) {

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

func generateExample() *tQuestionaire {

	quest := tQuestionaire{}

	quest.LangCodes = map[string]string{"de": "Deutsch", "en": "English"}
	quest.LangCode = "de"

	for i1 := 0; i1 < 3; i1++ {
		page := newPage()
		for i2 := 0; i2 < i1+1; i2++ {
			inp := newInput()
			if i2 == 1 {
				inp.Type = "checkbox"
			}
			if i2 == 2 {
				inp.Type = "radio"
			}
			page.Elements = append(page.Elements, inp)
		}
		quest.Pages = append(quest.Pages, page)
	}

	bts, err := json.MarshalIndent(quest, "", "  ")
	if err != nil {
		log.Fatalf("Marshal default questionaire: %v", err)
	}
	err = ioutil.WriteFile("example-questionaire.json", bts, 0644)
	if err != nil {
		log.Fatalf("Could not write example-quest.json: %v", err)
	}
	return &quest
}
