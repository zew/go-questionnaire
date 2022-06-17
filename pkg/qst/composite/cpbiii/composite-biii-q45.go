package cpbiii

import qstif "github.com/zew/go-questionnaire/pkg/qst/compositeif"

//
func QuestForOrg(q qstif.Q, seq0to5, paramSetIdx int) (string, []string, error) {

	s := `
		<label  for="q46" >
			Um Ihre Daten wissenschaftlich optimal 
			<a href='https://de.wikipedia.org/wiki/Lineare_Paneldatenmodelle' target="_blank" >auswerten</a> 
			zu können, wäre es hilfreich, wenn Sie uns  
			den Domänennamen ihrer Organisation nennen würden.
			<br>
		</label>
		<input type="text"  name="q46"  value="" %v
			placeholder="bspw. deutsche-bank.de oder dkb.de"
			maxlength="40"
		> 
	`
	inputNames := []string{"q46"}

	if q.UserIDInt() < (100 * 1000) {
		s = ""
		inputNames = []string{}
	}

	return s, inputNames, nil

}
