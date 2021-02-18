package qst

// Compposite inputs combine challenging HTML and
// multiple inputs in complicated ways
//
// A composite func returns dynamic HTML with session values inserted from the questionnaire
// A composite func also returns the input *names* for json generation of the questionnaire template
//
// Matching is required for
//      returned input names
//      input names in HTML form
//      input names to query session values from *q argument
//
//
// Parameters
//   dynamic questionnaire - filled with response values
//   sequence  idx  -  usually a visible page sequence
//   param set idx  -  statically determined - from a slice of param sets
//
//
// Returns
//   rendered HTML of the group
//   slice of input names
//   error
//
type compositFuncT func(*QuestionnaireT, int, int) (string, []string, error)

var compositeFuncs = map[string]compositFuncT{
	"PoliticalFoundationsPretext": PoliticalFoundationsPretext, // belongs to pat
	"PoliticalFoundations":        PoliticalFoundations,        // belongs to pat
	"TimePreferenceSelf":          TimePreferenceSelf,          // belongs to pat
	"GroupPreferences":            GroupPreferences,            // belongs to pat
}
