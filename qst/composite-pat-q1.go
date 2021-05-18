package qst

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

type preferences3x3T struct {
	ID   int     // unused - Frage-ID - not the sequence - unused
	Ppls [][]int // three cols - three rows  =>  0...5 person images
}

var fourPermutationsOf6x3x3 = make([][]preferences3x3T, 4)

var reshuffle6basedOn16 = [][]int{
	{1, 2, 3, 4, 5, 6}, // Version 1
	{1, 2, 3, 4, 5, 6}, // Version 2
	{1, 2, 3, 4, 5, 6}, // Version 3
	{1, 2, 3, 4, 5, 6}, // Version 4

	{6, 5, 4, 3, 2, 1}, // Version 5
	{6, 5, 4, 3, 2, 1}, // Version 6
	{6, 5, 4, 3, 2, 1}, // Version 7
	{6, 5, 4, 3, 2, 1}, // Version 8

	{6, 1, 5, 2, 4, 3}, // Version 9
	{6, 1, 5, 2, 4, 3}, // Version 10
	{6, 1, 5, 2, 4, 3}, // Version 11
	{6, 1, 5, 2, 4, 3}, // Version 12

	{3, 4, 2, 5, 1, 6}, // Version 13
	{3, 4, 2, 5, 1, 6}, // Version 14
	{3, 4, 2, 5, 1, 6}, // Version 15
	{3, 4, 2, 5, 1, 6}, // Version 16
}

// population - three pairs
var populationByVersion = [][]int{
	{1, 4},
	{1, 4},
	{1, 4},
	{1, 4},

	{2, 6},
	{2, 6},
	{2, 6},
	{2, 6},

	{2, 5},
	{2, 5},
	{2, 5},
	{2, 5},

	{2, 5},
	{2, 5},
	{2, 5},
	{2, 5},
}

// PoliticalFoundations creates
// a HTML table with three option and three checkbox inputs;
// seq0to5 is the numbering;
// based on userIDInt() - 4 versions / 4 permutations - via fourPermutationsOf6x3x3 + reshuffle6basedOn16;
// see composite.go for more.
func PoliticalFoundations(q *QuestionnaireT, seq0to5, paramSetIdx int) (string, []string, error) {

	zeroTo15 := q.Version()

	oneOfSix := reshuffle6basedOn16[zeroTo15][seq0to5] - 1 // display order => reshuffled questions order

	oneOfFour := zeroTo15 % 4 // table rows permutation

	// 	log.Printf(
	// 		`PoliticalFoundations
	// userID  %4v - zeroTo15  %2v
	// seq0to5 %4v - oneOfFour [0...3] %2v  - oneOfSix [0...5] %2v`,
	// 		userID, zeroTo15,
	// 		seq0to5, oneOfFour, oneOfSix,
	// 	)

	// log.Printf(`%v`, fourPermutationsOf6x3x3[oneOfFour][oneOfSix].Ppls)

	// questionID := fmt.Sprintf("q1_seq%v__%vof6_%vof4", seq0to5+1, oneOfSix+1, oneOfFour+1)
	questionID := fmt.Sprintf("q1_seq%v", seq0to5+1)

	return politicalFoundations(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		questionID,
		fourPermutationsOf6x3x3[oneOfFour][oneOfSix].Ppls,
	)
}

func PoliticalFoundationsStaticSub(q *QuestionnaireT, seq0to1, paramSetIdx int) (string, []string, error) {

	zeroTo15 := q.Version()
	threeDistinctPairs := populationByVersion[zeroTo15][seq0to1] - 1

	oneOfFour := zeroTo15 % 4 // table rows permutation

	return politicalFoundations(
		q,
		seq0to1,        // visible question seq 1...6 on the questionnaire
		"noQuestIDinQ", // not an input field in q
		fourPermutationsOf6x3x3[oneOfFour][threeDistinctPairs].Ppls,
	)
}

func politicalFoundations(q *QuestionnaireT, seq0to5 int, questionID string, ppls [][]int) (string, []string, error) {

	//
	inputNames := []string{}
	name := fmt.Sprintf("%v_r", questionID)
	inputNames = append(inputNames, name)
	for i := 0; i < 3; i++ {
		name := fmt.Sprintf("%v_r%v", questionID, i+1)
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
			// log.Printf(" sequenceID %v - inputNames[0] %v - inp.Response %v - idx %v", questionID, inputNames[0], inp.Response, idx)
			inputValsOptiongroup[idx] = " checked='checked' "
		}
	} else {
		log.Printf("poliFoundations: did not find radio input %v", inputNames[0])
	}

	inputValsCheckbox := make([]string, 3)
	for i := 1; i < len(inputNames); i++ {
		inp := q.ByName(inputNames[i])
		if inp != nil {
			if inp.Response != "" && inp.Response != "0" {
				inputValsCheckbox[i-1] = " checked='checked' "
			}
		} else {
			log.Printf("poliFoundations: did not find checkbox %v", inputNames[i])
		}
	}

	//
	//
	img := `<img src="/survey/img/pat/person.png">`

	imgs := []string{
		"",
		img,
		img + img,
		img + img + img,
		img + img + img + img,
		img + img + img + img + img,
		img + img + img + img + img + img,
	}

	explain := ""
	if seq0to5 == 0 {
		explain, _, _ = PoliticalFoundationsPretext(q, seq0to5, 0)
	}

	s := fmt.Sprintf(`

	
<p>
	<b>Entscheidung %v.</b>
</p>

<p>
	Welche Stiftung soll die 30&nbsp;€ bei folgender Präferenzkonstellation erhalten?<br>

	<span style="font-size: 88%%;">
		(Bitte ein Kreuz in der Spalte „Auswahl“, 
		und ggf. weitere Kreuze in der Spalte „Gleich gut“)
	</span>
</p>

%v
<div class="vspacer-08"> &nbsp; </div>
<div class="vspacer-08"> &nbsp; </div>


<div id="t01">


<table>
    <tr>
        <td style="width: 18%%;" > &nbsp; </td>
        <td style="width: 20%%;" >Am besten</td>
        <td style="width: 20%%;" >Mittel</td>
		<!-- https://www.wortbedeutung.info/schlecht/  -->
        <td style="width: 20%%;" >Am schlech&shy;tes&shy;ten</td>
        <td data-ia=1 style="width: 11%%;" >Aus&shy;wahl</td>
        <td data-ia=1 style="width: 11%%;" >Gleich gut</td>
    </tr>

    <tr>
        <td colspan="6" class="betw"> &nbsp; </td>
    </tr>

    <tr>
        <td>Stiftung A</td>
        <td> %v </td>
        <td> %v </td>
        <td> %v </td>
		<td data-ia=1> 
			<input type="radio"    name="_r"  value="1" %v > 
		</td>
		<td data-ia=1>
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
		<td data-ia=1> 
			<input type="radio"    name="_r"  value="2" %v > 
		</td>
		<td data-ia=1>  
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
		<td data-ia=1> 
			<input type="radio"    name="_r"  value="3" %v > 
		</td>
		<td data-ia=1>
			<input type="checkbox" name="_r3" value="1" %v  > 
			<input type="hidden"   name="_r3" value="0" >
		 </td>
    </tr>


</table>

 <!-- </span> /go-quest-label -->


</div>
	`,
		seq0to5+1,
		explain,
		imgs[ppls[0][0]], imgs[ppls[0][1]], imgs[ppls[0][2]],
		inputValsOptiongroup[0], inputValsCheckbox[0],

		imgs[ppls[1][0]], imgs[ppls[1][1]], imgs[ppls[1][2]],
		inputValsOptiongroup[1], inputValsCheckbox[1],

		imgs[ppls[2][0]], imgs[ppls[2][1]], imgs[ppls[2][2]],
		inputValsOptiongroup[2], inputValsCheckbox[2],
	)

	// prefix name=" with questionID
	rep := fmt.Sprintf(`name="%v`, questionID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}

/*
	https://www.regextester.com/
	https://stackoverflow.com/questions/37106834/golang-multiline-regex-not-working
	https://github.com/google/re2/wiki/Syntax

    (?is)  is setting flags to insensitive and . to matching newlines
	(.*?)  the question mark is for non greedy
*/
var re = regexp.MustCompile(`(?is)<td data-ia=1(.*?)<\/td>`)

var cols6to4 = strings.NewReplacer(
	`<td colspan="6" class="betw"> &nbsp; </td>`,
	`<td colspan="4" class="betw"> &nbsp; </td>`,
)

// PoliticalFoundationsStatic - like PoliticalFoundations() but filtering out
// the input columns
func PoliticalFoundationsStatic(q *QuestionnaireT, seq0to5, paramSetIdx int) (string, []string, error) {

	ret, _, err := PoliticalFoundationsStaticSub(q, seq0to5, paramSetIdx)

	completeDeletionOfCols56 := true
	if completeDeletionOfCols56 {
		ret = cols6to4.Replace(ret)
		fillin := ""
		ret = re.ReplaceAllString(ret, fillin)
	} else {
		fillin := "<td>&nbsp;</td>"
		ret = re.ReplaceAllString(ret, fillin)
	}

	//
	sep := `<div id="t01">`
	rets := strings.Split(ret, sep)
	if len(rets) != 2 {
		msg := fmt.Sprintf("Splitting by <pre>%v </pre> failed: Changes in politicalFoundations()?", sep)
		return msg, nil, fmt.Errorf(msg)
	}

	rets[1] = sep + rets[1]

	return rets[1], nil, err
}
