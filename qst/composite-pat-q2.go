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
// return 2 are the input names, based on seq0to5;
func TimePreferenceSelf(q *QuestionnaireT, seq0to5, paramSetIdx int) (string, []string, error) {

	userID := 0
	if q != nil {
		userID = q.UserIDInt()
	}

	zeroTo15 := userID % 16

	// 	log.Printf(
	// 		`TimePreferenceSelf
	// userID  %5v - zeroTo15  %2v
	// paramSetIdx %v - seq0to0 %4v`,
	// 		userID, zeroTo15,
	// 		paramSetIdx, seq0to5,
	// 	)

	aOrB := "a"
	if paramSetIdx > 0 {
		aOrB = "b"
	}

	// log.Printf(`%v`, getQ2Labels(zeroTo15, aOrB))

	questionID := fmt.Sprintf("q2_seq%v__%02vof16", seq0to5+1, zeroTo15+1)
	questionID = fmt.Sprintf("q2_seq%v", seq0to5+1)

	return timePreferenceSelf(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		questionID,
		getQ2Labels(zeroTo15, aOrB),
	)
}

func timePreferenceSelf(q *QuestionnaireT, seq0to0 int, questionID string, rowLabels []string) (string, []string, error) {

	//
	//
	inputNames := []string{}
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("%v_row%v_rad", questionID, i+1)
		inputNames = append(inputNames, name)
	}
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("%v_row%v_chk", questionID, i+1)
		inputNames = append(inputNames, name)
	}

	if q == nil {
		// we are at static build time
		return "", inputNames, nil
	}

	//
	//
	inputValsOptiongroup := make([]string, 6)
	for row := 0; row < 3; row++ {
		inp := q.ByName(inputNames[row])
		if inp != nil {
			if inp.Response != "" && inp.Response != "0" {
				val, _ := strconv.Atoi(inp.Response) // can be 1 or 2
				inputValsOptiongroup[2*row+val-1] = " checked='checked' "
			}
		} else {
			log.Printf("timePref: did not find radio input %v", inputNames[row])
		}
	}

	//
	inputValsCheckbox := make([]string, 3)
	for row := 0; row < 3; row++ {
		inp := q.ByName(inputNames[row+3])
		if inp != nil {
			if inp.Response != "" && inp.Response != "0" {
				inputValsCheckbox[row] = " checked='checked' "
			}
		} else {
			log.Printf("timePref: did not find checkbox %v", inputNames[row])
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
        <td style="width: 15%%"  rowspan=2 >Von dieser Option abraten</td>
    </tr>
    <tr>
        <td  colspan="2" > <i>Bitte jeweils EIN Kreuz setzen</i> </td>
    </tr>

    <tr>
        <td colspan="4" class="betw"> &nbsp; </td>
    </tr>

    <tr>
		<td>%v</td>
		<td colspan="2">
			<input type="radio"    name="_row1_rad" value="1" %v >   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;oder&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
			<input type="radio"    name="_row1_rad" value="2" %v > 
		</td>
		<td>
			<input type="checkbox" name="_row1_chk" value="1" %v > 
			<input type="hidden"   name="_row1_chk" value="0"    >
		</td>
    </tr>

    <tr>
		<td>%v</td>
		<td colspan="2">
			<input type="radio"    name="_row2_rad" value="1" %v >   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;oder&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
			<input type="radio"    name="_row2_rad" value="2" %v > 
		</td>
		<td>
			<input type="checkbox" name="_row2_chk" value="1" %v > 
			<input type="hidden"   name="_row2_chk" value="0"    >
		</td>
    </tr>

    <tr>
		<td>%v</td>
		<td colspan="2">
			<input type="radio"    name="_row3_rad" value="1" %v >   &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;oder&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
			<input type="radio"    name="_row3_rad" value="2" %v > 
		</td>
		<td>
			<input type="checkbox" name="_row3_chk" value="1" %v > 
			<input type="hidden"   name="_row3_chk" value="0"    >
		</td>
    </tr>

</table>

</div>

	`,
		// seq0to0+1,
		rowLabels[0],
		inputValsOptiongroup[0], inputValsOptiongroup[1], inputValsCheckbox[0],
		rowLabels[1],
		inputValsOptiongroup[2], inputValsOptiongroup[3], inputValsCheckbox[1],
		rowLabels[2],
		inputValsOptiongroup[4], inputValsOptiongroup[5], inputValsCheckbox[2],
	)

	// prefix name=" with questionID
	rep := fmt.Sprintf(`name="%v`, questionID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}
