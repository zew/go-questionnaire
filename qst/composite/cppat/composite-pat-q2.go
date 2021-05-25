package cppat

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	qstif "github.com/zew/go-questionnaire/qst/compositeif"
)

// TimePreferenceSelf creates
// a HTML table with six option and three checkbox inputs;
// based on userIDInt() - 8 versions - via paramSetIdx + dataQ2;
// seq0to5 is the numbering;
// see composite.go for more.
func TimePreferenceSelf(q qstif.Q, seq0to5, paramSetIdx int) (string, []string, error) {

	zeroTo15 := q.Version()

	aOrB := "a"
	if paramSetIdx > 0 {
		aOrB = "b"
	}

	questionID := fmt.Sprintf("q2_seq%v", seq0to5+1)

	return timePreferenceSelf(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		questionID,
		getQ2Labels(zeroTo15, aOrB),
	)
}

// TimePreferenceSelfStatic similar to TimePreferenceSelf;
// but inputs are disabled
func TimePreferenceSelfStatic(q qstif.Q, seq0to5, paramSetIdx int) (string, []string, error) {

	s, inputs, err := TimePreferenceSelf(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		paramSetIdx,
	)

	s = strings.ReplaceAll(s, "<input ", "<input disabled ")
	s = strings.ReplaceAll(s, `checked='checked'`, " ")

	/*
		https://www.regextester.com/
		https://stackoverflow.com/questions/37106834/golang-multiline-regex-not-working
		https://github.com/google/re2/wiki/Syntax

		(?is)  is setting flags to insensitive and . to matching newlines
		(.*?)  the question mark is for non greedy

		var re = regexp.MustCompile(`(?is)<input(.*?)>`)
		s = re.ReplaceAllString(s, "<!-- input replaced -->")
	*/

	return s, inputs, err
}

func timePreferenceSelf(q qstif.Q, seq0to0 int, questionID string, rowLabels []string) (string, []string, error) {

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

	//
	//
	inputValsOptiongroup := make([]string, 6)
	for row := 0; row < 3; row++ {
		resp, err := q.ResponseByName(inputNames[row])
		if err != nil {
			// generators.Create() and qst.Load() for new user
			//  are calling qst.Validate()
			//   => dynamic fields do not exist yet
		} else {
			if resp != "" && resp != "0" {
				val, _ := strconv.Atoi(resp) // can be 1 or 2
				inputValsOptiongroup[2*row+val-1] = " checked='checked' "
			}
		}
	}

	//
	//
	inputValsCheckbox := make([]string, 3)
	for row := 0; row < 3; row++ {
		resp, err := q.ResponseByName(inputNames[row+3])
		if err != nil {
			// generators.Create() and qst.Load() for new user
			//  are calling qst.Validate()
			//   => dynamic fields do not exist yet
		} else {
			if resp != "" && resp != "0" {
				inputValsCheckbox[row] = " checked='checked' "
			}
		}
	}

	consolidatedErrMsg := ""
	for _, inpName := range inputNames {
		em, err := q.ErrByName(inpName)
		if err == nil && em != "" {
			consolidatedErrMsg = fmt.Sprintf(`
			<div class="    error   error-block-input " style="">
			  %v
			</div>
			`, em)
		}
	}

	//
	//
	s := fmt.Sprintf(`


<div id="t02">

<div class="vspacer-08"> &nbsp; </div>

%v

<table>
    <tr>
        <td style="width: 37%%"  rowspan=2 >Optionen</td>
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
		consolidatedErrMsg,
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
