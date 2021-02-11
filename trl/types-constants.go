package trl

// LangCodes for returning multiple translations.
// When no langCode is available, then the first entry rules.
// A call to All() returns explicitly all key-values.
// LangCodes will be initialized in cfg.Load().LangCodes; we prevent circular dependency
var LangCodes = []string{"de", "en"}

const noTrans = "multi lingual string not initialized."

// S stores a multi lingual string.
// Contains one value for each language code.
type S map[string]string

// Map - Translations Type
// Usage in templates
// 		{{.Trls.imprint.en                     }}  // directly accessing a specific translation; chaining the map keys
// 		{{.Trls.imprint.Tr       .Sess.LangCode}}  // using .Tr(langCode)
// 		{{.Trls.imprint.TrSilent .Sess.LangCode}}  //
type Map map[string]S

// MapSite for sub-site specific translations
type MapSite map[string]Map
