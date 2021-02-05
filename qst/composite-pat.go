package qst

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

type preferences3x3T struct {
	ID   int     // Frage-ID - not the sequence
	Ppls [][]int // three cols - three rows  =>  0...5 person images
}

var fourPermutationsOf6x3x3 = make([][]preferences3x3T, 4)

var reshuffle6basedOn16 = [][]int{
	{1, 2, 3, 4, 5, 6},
	{1, 2, 3, 4, 5, 6},
	{1, 2, 3, 4, 5, 6},
	{1, 2, 3, 4, 5, 6},

	{6, 5, 4, 3, 2, 1},
	{6, 5, 4, 3, 2, 1},
	{6, 5, 4, 3, 2, 1},
	{6, 5, 4, 3, 2, 1},

	{6, 1, 5, 2, 4, 3},
	{6, 1, 5, 2, 4, 3},
	{6, 1, 5, 2, 4, 3},
	{6, 1, 5, 2, 4, 3},

	{3, 4, 2, 5, 1, 6},
	{3, 4, 2, 5, 1, 6},
	{3, 4, 2, 5, 1, 6},
	{3, 4, 2, 5, 1, 6},
}

// PoliticalFoundations creates
// a HTML table with three option and three checkbox inputs;
// renderSeq is the numbering
// politicalFoundationsParamsT.ID for rendering the numbering;
// politicalFoundationsParamsT.Ppls for rendering icons of peoples to certain positions;
// return 1 is the HTML code
// return 2 are the input names, based on seq0to5;
func PoliticalFoundations(q *QuestionnaireT, userID, seq0to5, paramSetIdx int) (string, []string, error) {

	zeroTo15 := userID % 16

	oneOfSix := reshuffle6basedOn16[zeroTo15][seq0to5] - 1

	oneOfFour := zeroTo15 % 4

	return politicalFoundations(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		fmt.Sprintf("%v_%v", oneOfFour+1, oneOfSix+1), // questionID - fourPermutationsOf6x3x3[oneOfFour][oneOfSix] -
		fourPermutationsOf6x3x3[oneOfFour][oneOfSix].Ppls,
	)
}

func politicalFoundations(q *QuestionnaireT, seq0to5 int, questionID string, ppls [][]int) (string, []string, error) {

	//
	//
	inputNames := []string{}
	name := fmt.Sprintf("dec%v_r", questionID)
	inputNames = append(inputNames, name)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("dec%v_r%v", questionID, i+1)
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
			log.Printf(" sequenceID %v - inputNames[0] %v - inp.Response %v - idx %v", questionID, inputNames[0], inp.Response, idx)
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
		seq0to5+1,
		imgs[ppls[0][0]], imgs[ppls[0][1]], imgs[ppls[0][2]],
		inputValsOptiongroup[0], inputValsCheckbox[0],
		imgs[ppls[1][0]], imgs[ppls[1][1]], imgs[ppls[1][2]],
		inputValsOptiongroup[1], inputValsCheckbox[1],
		imgs[ppls[2][0]], imgs[ppls[2][1]], imgs[ppls[2][2]],
		inputValsOptiongroup[2], inputValsCheckbox[2],
	)

	rep := fmt.Sprintf(`name="dec%v`, questionID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}
