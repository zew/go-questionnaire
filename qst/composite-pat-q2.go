package qst

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

// TimePreferenceSelf creates
// a HTML table with three option and three checkbox inputs;
// renderSeq is the numbering
// TimePreferenceSelfParamsT.ID for rendering the numbering;
// TimePreferenceSelfParamsT.Ppls for rendering icons of peoples to certain positions;
// return 1 is the HTML code
// return 2 are the input names, based on seq0to0;
func TimePreferenceSelf(q *QuestionnaireT, paramSetIdx, seq0to0, userID int) (string, []string, error) {

	zeroTo15 := userID % 16

	log.Printf(
		`TimePreferenceSelf
userID  %5v - zeroTo15  %2v
paramSetIdx %v - seq0to0 %4v`,
		userID, zeroTo15,
		paramSetIdx, seq0to0,
	)

	aOrB := "a"
	if paramSetIdx > 0 {
		aOrB = "b"
	}

	log.Printf(`%v`, getQ2Labels(zeroTo15, aOrB))

	return timePreferenceSelf(
		q,
		seq0to0,                         // visible question seq 1...6 on the questionnaire
		fmt.Sprintf("q2_%v_", zeroTo15), // questionID - fourPermutationsOf6x3x3[oneOfFour][oneOfSix] -
		getQ2Labels(zeroTo15, aOrB),
	)
}

func timePreferenceSelf(q *QuestionnaireT, seq0to0 int, questionID string, rowLabels []string) (string, []string, error) {

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

	s := fmt.Sprintf(`


<div id="t02">


<table>
    <tr>
        <td style="width: 35%%"  rowspan=2 >Optionen</td>
        <td style="width: 18%%"            ><b>Verfügbar</b></td>
        <td style="width: 18%%"            ><b>Nicht verfügbar</b></td>
        <td style="width: 14%%"  rowspan=2 >Von dieser Option abraten</td>
    </tr>
    <tr>
        <td  colspan="2" > <i>Bitte jeweils EIN Kreuz setzen</i> </td>
    </tr>

    <tr>
        <td colspan="4" class="betw"> &nbsp; </td>
    </tr>

    <tr>
        <td>%v</td>
        <td colspan="2">%v</td>
        <td>%v</td>
    </tr>

    <tr>
        <td>%v</td>
        <td colspan="2">%v</td>
        <td>%v</td>
    </tr>

    <tr>
        <td>%v</td>
        <td colspan="2">%v</td>
        <td>%v</td>
    </tr>

</table>

</div>

	`,
		// seq0to0+1,
		rowLabels[0],
		inputValsOptiongroup[0], inputValsCheckbox[0],
		rowLabels[1],
		inputValsOptiongroup[1], inputValsCheckbox[1],
		rowLabels[2],
		inputValsOptiongroup[2], inputValsCheckbox[2],
	)

	// prefix name=" with questionID
	rep := fmt.Sprintf(`name="dec%v`, questionID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}
