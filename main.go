package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
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
		Name: fmt.Sprintf("NewInput_%v", cntr.increment()),
	}
	return t
}

func newPage() tPage {
	t := tPage{
		Name: fmt.Sprintf("NewPage_%v", cntr.increment()),
	}
	return t
}

func shortTime() string {
	return time.Now().Format("15:04:05") + "\t"
}

func (t *tInput) HTML() string {
	str := fmt.Sprintf("<input type='text' name='%v' id='%v' label='%v' />", t.Name, t.Name, t.Desc)
	return str
}

func (t tInput) String() string {
	return t.HTML()
}

func main() {

	log.SetFlags(log.Lshortfile)
	f, err := os.Create("questionaire.log")
	if err != nil {
		log.Fatalf("Could not open log file: %v\n", err)
	}
	log.SetOutput(f)

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
	err = ioutil.WriteFile("example-quest.json", bts, 0644)
	if err != nil {
		log.Fatalf("Could not write example-quest.json: %v\n", err)
	}

	//

}
