package patbelief

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"

	qstif "github.com/zew/go-questionnaire/pkg/qstif"
)

type DraggableItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func shuffle(seed int, items []DraggableItem) []DraggableItem {
	rand.New(rand.NewSource(int64(seed)))
	shuffled := make([]DraggableItem, len(items))
	copy(shuffled, items)

	for i := len(shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	}

	return shuffled
}

func groupName(g DraggableItem) string {
	return fmt.Sprintf("group_%s", g.Value)
}

var items = [5]DraggableItem{{
	Name:  "one",
	Value: "1",
}, {
	Name:  "two",
	Value: "2",
}, {
	Name:  "three",
	Value: "3",
}, {
	Name:  "four",
	Value: "4",
}, {
	Name:  "five",
	Value: "5",
}}

var choices = [6]DraggableItem{{
	Name:  "A = Hans Böckler<br>B = Bund der Steuerzahler<br>C = Ludwig Erhard<br>",
	Value: "hbl",
}, {
	Name:  "A = Hans Böckler<br>B = Ludwig Erhard<br>C = Bund der Steuerzahler<br>",
	Value: "hlb",
}, {
	Name:  "A = Bund der Steuerzahler<br>B = Hans Böckler<br>C = Ludwig Erhard<br>",
	Value: "bhl",
}, {
	Name:  "A = Bund der Steuerzahler<br>B = Ludwig Erhard<br>C = Hans Böckler<br>",
	Value: "blh",
}, {
	Name:  "A = Ludwig Erhard<br>B = Hans Böckler<br>C = Bund der Steuerzahler<br>",
	Value: "lhb",
}, {
	Name:  "A = Ludwig Erhard<br>B = Bund der Steuerzahler<br>C = Hans Böckler<br>",
	Value: "lbh",
}}

func sortItemsByOrder(items []DraggableItem, order []string) []DraggableItem {
	orderIndex := make(map[string]int)
	for i, val := range order {
		orderIndex[val] = i
	}

	sort.Slice(items, func(i, j int) bool {
		return orderIndex[groupName(items[i])] < orderIndex[groupName(items[j])]
	})

	return items
}

func choiceNames(items []DraggableItem) []string {
	var names []string
	for i := 0; i < len(items); i++ {
		names = append(names, groupName(items[i]))
	}
	return names
}

func PatPoliticalBeliefs(q qstif.Q, seq0to5, paramSetIdx int, preflight bool) (string, []string, error) {
	findResponse := func(name string) string {
		value, err := q.ResponseByName(name)
		if err != nil {
			log.Printf("could not find input name %v", err)
			return ""
		}
		return value
	}
	var orderName = "group-order"
	currentOrder := findResponse(orderName)

	var choiceOrder []string
	var userId = q.UserIDInt()

	if currentOrder == "" {
		var shuffled = shuffle(userId, choices[:])
		choiceOrder = choiceNames(shuffled)
	} else {
		choiceOrder = strings.Split(currentOrder, ",")
	}

	var hiddenInputs []string
	for _, choice := range choices {
		name := groupName(choice)
		value := findResponse(name)
		hiddenInputs = append(hiddenInputs, fmt.Sprintf("<input type='text' name='%s' id='%s' value='%s'/>", name, name, value))
	}

	hiddenInputs = append(hiddenInputs, fmt.Sprintf(`
		<input type='text' name='%s' value='%s' />`,
		orderName, strings.Join(choiceOrder, ",")))

	var shuffledChoices = sortItemsByOrder(choices[:], choiceOrder)

	choicesHtml := make([]string, len(shuffledChoices))
	for i, choice := range shuffledChoices {
		choicesHtml[i] = fmt.Sprintf(`
			<div class='droppable-box'>
				<div class='droppable-header'>%s</div>
				<div data-droppable data-target='%s' data-id='%v'></div>
			</div>`,
			choice.Name,
			groupName(choices[i]), i)
	}

	itemsHtml := make([]string, len(items))
	for i, choice := range items {
		itemsHtml[i] = fmt.Sprintf("<div class='item' id='item-%v' draggable='true' data-value='%s'>%s</div>", i, choice.Value, choice.Name)
	}

	var html = fmt.Sprintf(`
		<div class='hidden-group'>%s</div>
		<div class='droppable-container'>%s</div>
		<div id='draggable-list'>%s</div>`,
		strings.Join(hiddenInputs, ""),
		strings.Join(choicesHtml, ""),
		strings.Join(itemsHtml, ""))

	var inputNames = choiceNames(choices[:])
	inputNames = append(inputNames, orderName)
	return html, inputNames, nil
}
