package qst

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
)

// GroupPreferences creates a HTML table with three columns
// based on userIDInt() - 8 versions - via paramSetIdx + dataQ3;
// seq0to5 is the numbering;
// see composite.go for more.
func GroupPreferences(q *QuestionnaireT, seq0to5, paramSetIdx int) (string, []string, error) {

	userID := 0
	if q != nil {
		userID = q.UserIDInt()
	}

	zeroTo15 := userID % 16

	aOrB := "a"
	if paramSetIdx > 0 {
		aOrB = "b"
	}

	// log.Printf(`%v`, getQ2Labels(zeroTo15, aOrB))

	questionID := fmt.Sprintf("q2_seq%v__%02vof16", seq0to5+1, zeroTo15+1)
	questionID = fmt.Sprintf("q2_seq%v", seq0to5+1)

	return groupPreferences(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		questionID,
		getQ3Labels(zeroTo15, aOrB),
	)
}

func groupPreferences(q *QuestionnaireT, seq0to0 int, questionID string, rowLabels []string) (string, []string, error) {

	//
	inputNames := []string{}

	if q == nil {
		// we are at static build time
		return "", inputNames, nil
	}

	s := fmt.Sprintf(`


<div id="t03">


<style>
    #t03 .b1 {
        display: inline-block;
        margin: 0.7rem;
        width: 10.4rem;
        border: 1px solid grey;
    }
	#t03 .b2, 
	#t03 .b3 {
        padding: 0.2rem;
    }
    #t03 .b2 {
        border-bottom: 1px solid grey;
    }

</style>


<div class="b1">
    <div class="b2">
        Option A
    </div>
    <div class="b3">
        %v
    </div>
</div>


<div class="b1">
    <div class="b2">
        Option B
    </div>
	<div class="b3">
		%v
    </div>
</div>


<div class="b1">
    <div class="b2">
        Option C
    </div>
	<div class="b3">
		%v
    </div>
</div>

</div>

	`,
		// seq0to0+1,
		rowLabels[0],
		rowLabels[1],
		rowLabels[2],
	)

	// prefix name=" with questionID
	rep := fmt.Sprintf(`name="%v`, questionID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}
