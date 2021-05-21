package qst

// CompositeFuncT inputs combine challenging HTML and
// multiple inputs in complicated ways
//
// Compare dynFuncT, validatorT
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
type CompositeFuncT func(*QuestionnaireT, int, int) (string, []string, error)

// CompositeFuncs is a lookup map
var CompositeFuncs = map[string]CompositeFuncT{
	"PoliticalFoundationsPretext":            PoliticalFoundationsPretext,            // belongs to pat
	"PoliticalFoundations":                   PoliticalFoundations,                   //   ...
	"PoliticalFoundationsStatic":             PoliticalFoundationsStatic,             //   ... no input
	"PoliticalFoundationsComprehensionCheck": PoliticalFoundationsComprehensionCheck, //   ... no input
	"TimePreferenceSelf":                     TimePreferenceSelf,                     // belongs to pat
	"TimePreferenceSelfStatic":               TimePreferenceSelfStatic,               //   ... disabled
	"GroupPreferences":                       GroupPreferences,                       // belongs to pat
	"GroupPreferencesPOP3":                   GroupPreferencesPOP3,                   //
}
