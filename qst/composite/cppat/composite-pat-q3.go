package cppat

import (
	"fmt"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	qstif "github.com/zew/go-questionnaire/qst/compositeif"
)

// GroupPreferences creates a HTML table with three columns
// based on userIDInt() - 8 versions - via paramSetIdx + dataQ3;
// seq0to5 is the numbering;
// see composite.go for more.
func GroupPreferences(q qstif.Q, seq0to5, paramSetIdx int) (string, []string, error) {

	zeroTo15 := q.Version()

	aOrB := "a"
	if paramSetIdx > 0 {
		aOrB = "b"
	}

	// log.Printf(`%v`, getQ2Labels(zeroTo15, aOrB))

	// questionID := fmt.Sprintf("q2_seq%v__%02vof16", seq0to5+1, zeroTo15+1)
	questionID := fmt.Sprintf("q2_seq%v", seq0to5+1)

	return groupPreferences(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		questionID,
		getQ3Labels(zeroTo15, aOrB),
	)
}

// GroupPreferencesPOP3 - yields TimePreference - getQ2Labels() - not getQ3Labels
func GroupPreferencesPOP3(q qstif.Q, seq0to5, paramSetIdx int) (string, []string, error) {

	zeroTo15 := q.Version()

	aOrB := "a"
	if paramSetIdx > 0 {
		aOrB = "b"
	}

	questionID := fmt.Sprintf("q2_seq%v", seq0to5+1)

	return groupPreferences(
		q,
		seq0to5, // visible question seq 1...6 on the questionnaire
		questionID,
		getQ2Labels(zeroTo15, aOrB),
	)
}

func groupPreferences(q qstif.Q, seq0to0 int, questionID string, rowLabels []string) (string, []string, error) {

	//
	inputNames := []string{}

	if q == nil {
		// we are at static build time
		return "", inputNames, nil
	}

	s := fmt.Sprintf(`


<div id="t03">




<div class="b1">
    <div class="b2">Option A</div>
    <div class="b3">%v</div>
</div>


<div class="b1">
    <div class="b2">Option B</div>
	<div class="b3">%v</div>
</div>


<div class="b1">
    <div class="b2">Option C</div>
	<div class="b3">%v</div>
</div>

</div>

	`,
		// seq0to0+1,
		rowLabels[0],
		rowLabels[1],
		rowLabels[2],
	)

	// prefix name=" with questionID
	// is this even necessary here?
	rep := fmt.Sprintf(`name="%v`, questionID)
	s = strings.ReplaceAll(s, `name="`, rep)

	s = strings.ReplaceAll(s, "/survey/", cfg.PrefTS())

	return s, inputNames, nil

}
