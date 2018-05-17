package qst

// For all multi lingual strings.
// Contains one value for each language code.
type transMapT map[string]string // type translation map

// Tr translates
func (t transMapT) Tr(langCode string) string {
	if val, ok := t[langCode]; ok {
		return val
	}
	if val, ok := t["en"]; ok {
		return val
	}
	for _, val := range t {
		return val
	}
	if t == nil {
		return "Translation map not initialized."
	}
	return "Translation map not initialized."
}

// Default "stringer" implementation
func (t transMapT) String() string {
	return t.Tr("en")
}
