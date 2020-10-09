package qst

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

type politicalFoundationsParamsT struct {
	ID   int
	Ppls [][]int
}

var politicalFoundationsParams = []politicalFoundationsParamsT{
	{
		ID: 0, // Frage 1
		Ppls: [][]int{
			{2, 3, 0},
			{0, 2, 3},
			{3, 0, 2},
		},
	},
	{
		ID: 1, // Frage 2
		Ppls: [][]int{
			{3, 0, 2},
			{0, 3, 2},
			{2, 2, 1},
		},
	},
	{
		ID: 2, // Frage 3
		Ppls: [][]int{
			{2, 0, 3},
			{0, 5, 0},
			{3, 0, 2},
		},
	},
	{
		ID: 3, // Frage 4
		Ppls: [][]int{
			{1, 0, 4},
			{1, 4, 0},
			{3, 1, 2},
		},
	},
	{
		ID: 4, // Frage 5
		Ppls: [][]int{
			{3, 0, 2},
			{2, 1, 2},
			{0, 4, 1},
		},
	},
	{
		ID: 5, // Frage 6
		Ppls: [][]int{
			{1, 0, 4},
			{4, 0, 1},
			{0, 5, 0},
		},
	},
}

// PoliticalFoundations creates
// a HTML table with three option and three checkbox inputs;
// renderSeq is the numbering
// politicalFoundationsParamsT.ID for rendering the numbering;
// politicalFoundationsParamsT.Ppls for rendering icons of peoples to certain positions;
// return 1 is the HTML code
// return 2 are the input names, based in paramSetIdx;
func PoliticalFoundations(q *QuestionnaireT, renderSeq int, paramSetIdx int) (string, []string, error) {
	return politicalFoundations(
		q,
		renderSeq,
		politicalFoundationsParams[paramSetIdx].ID,
		politicalFoundationsParams[paramSetIdx].Ppls,
	)
}

func politicalFoundations(q *QuestionnaireT, renderSeq, sequenceID int, ppls [][]int) (string, []string, error) {

	//
	//
	inputNames := []string{}
	name := fmt.Sprintf("dec%v_r", sequenceID)
	inputNames = append(inputNames, name)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("dec%v_r%v", sequenceID, i+1)
		inputNames = append(inputNames, name)
	}

	if q == nil {
		// we are at static build time
		return "", inputNames, nil
	}

	//
	//
	inputValsOptiongroup := make([]string, 3)
	inp := q.ByName(inputNames[0])
	if inp != nil {
		if inp.Response != "" && inp.Response != "0" {
			idx, _ := strconv.Atoi(inp.Response)
			idx--
			log.Printf(" sequenceID %v - inputNames[0] %v - inp.Response %v - idx %v", sequenceID, inputNames[0], inp.Response, idx)
			inputValsOptiongroup[idx] = " checked='checked' "
		}
	}

	inputValsCheckbox := make([]string, 3)
	for i := 1; i < len(inputNames); i++ {
		inp := q.ByName(inputNames[i])
		if inp != nil {
			if inp.Response != "" && inp.Response != "0" {
				inputValsCheckbox[i-1] = " checked='checked' "
			}
		}
	}

	//
	//
	img := `<img src="/survey/img/person.png">`

	imgs := []string{
		"",
		img,
		img + img,
		img + img + img,
		img + img + img + img,
		img + img + img + img + img,
		img + img + img + img + img + img,
	}

	s := fmt.Sprintf(`
<div id="t01">


<table>
    <tr>
        <td>Entscheidung&nbsp;%v</td>
        <td>Beste</td>
        <td>Mittel</td>
        <td>Schlechteste</td>
        <td>Auswahl</td>
        <td>Genauso gut</td>
    </tr>

    <tr>
        <td colspan="6" class="betw"> &nbsp; </td>
    </tr>

    <tr>
        <td>Stiftung A</td>
        <td> %v </td>
        <td> %v </td>
        <td> %v </td>
        <td> <input type="radio"    name="_r"  value="1" %v > </td>
		<td>
			<input type="checkbox" name="_r1" value="1"  %v > 
			<input type="hidden"   name="_r1" value="0" >
		</td>
    </tr>

    <tr>
        <td colspan="6" class="betw"> &nbsp; </td>
    </tr>

    <tr>
        <td>Stiftung B</td>
        <td> %v </td>
        <td> %v </td>
        <td> %v </td>
        <td> <input type="radio"    name="_r"  value="2" %v > </td>
		<td>  
			<input type="checkbox" name="_r2" value="1"  %v > 
			<input type="hidden"   name="_r2" value="0" >
		</td>
    </tr>

    <tr>
        <td colspan="6" class="betw"> &nbsp; </td>
    </tr>

    <tr>
        <td>Stiftung C</td>
        <td> %v </td>
        <td> %v </td>
        <td> %v </td>
        <td> <input type="radio"    name="_r"  value="3" %v > </td>
		<td>
			<input type="checkbox" name="_r3" value="1" %v  > 
			<input type="hidden"   name="_r3" value="0" >
		 </td>
    </tr>


</table>

</div>
	`,
		renderSeq+1,
		imgs[ppls[0][0]], imgs[ppls[0][1]], imgs[ppls[0][2]],
		inputValsOptiongroup[0], inputValsCheckbox[0],
		imgs[ppls[1][0]], imgs[ppls[1][1]], imgs[ppls[1][2]],
		inputValsOptiongroup[1], inputValsCheckbox[1],
		imgs[ppls[2][0]], imgs[ppls[2][1]], imgs[ppls[2][2]],
		inputValsOptiongroup[2], inputValsCheckbox[2],
	)

	rep := fmt.Sprintf(`name="dec%v`, sequenceID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}
