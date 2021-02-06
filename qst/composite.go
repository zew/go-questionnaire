package qst

// compposite inputs combine challenging HTML and
// multiple inputs in complicated ways
//
//
//
// parameters
//   dynamic questionnaire - filled with response values
//   param set idx  -  statically determined - from a slice of param sets
//   sequence  idx  -  usually a visible page sequence
//   userID         -  for custom randomization
//
// returns
//   rendered HTML of the group
//   slice of input names
//   error
type compositFuncT func(*QuestionnaireT, int, int, int) (string, []string, error)

var compositeFuncs = map[string]compositFuncT{
	"PoliticalFoundations": PoliticalFoundations,
	"TimePreferenceSelf":   TimePreferenceSelf,
}
