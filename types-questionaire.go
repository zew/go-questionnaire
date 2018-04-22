package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type tInput struct {
	Name string `json:"name,omitempty"`
	Desc string `json:"description,omitempty"`
	// Validator func() bool `json:"empty"`
	ErrMsg string `json:"err_msg,omitempty"`
}

type tPage struct {
	Name     string   `json:"name,omitempty"`
	Elements []tInput `json:"elements,omitempty"`
}

type tQuestionaire struct {
	Pages []tPage `json:"pages,omitempty"`
}

func newInput() tInput {
	t := tInput{
		Name: fmt.Sprintf("input_%v", cntr.increment()),
	}
	return t
}

func newPage() tPage {
	t := tPage{
		Name: fmt.Sprintf("page_%v", cntr.increment()),
	}
	return t
}

func (t tInput) HTML() string {
	str := fmt.Sprintf("<input type='text' name='%v' id='%v' label='%v' />\n", t.Name, t.Name, t.Desc)
	return str
}

func (t tInput) String() string {
	return t.HTML()
}

func (p tPage) HTML() string {
	str := fmt.Sprintf("<h3>%v</h3>", p.Name)
	// for inp := range p.Elements {
	// 	str += &inp.HTML()
	// }
	for i := 0; i < len(p.Elements); i++ {
		str += "<br><br>next element <br>\n"
		str += p.Elements[i].HTML() + "\n"
	}
	return str
}

func generateExample() *tQuestionaire {
	quest := tQuestionaire{}
	for i1 := 0; i1 < 2; i1++ {
		page := newPage()
		for i2 := 0; i2 < 2; i2++ {
			inp := newInput()
			page.Elements = append(page.Elements, inp)
		}
		quest.Pages = append(quest.Pages, page)
	}

	bts, err := json.MarshalIndent(quest, "", "  ")
	if err != nil {
		log.Fatalf("Unmarshal default cn: %v\n", err)
	}
	err = ioutil.WriteFile("example-questionaire.json", bts, 0644)
	if err != nil {
		log.Fatalf("Could not write example-quest.json: %v\n", err)
	}
	return &quest
}
