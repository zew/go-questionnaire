package qst

import "fmt"

var implementedTypes = map[string]interface{}{
	"text":     nil,
	"textarea": nil,
	"checkbox": nil, // A standalone checkbox - as a group, see below

	// radiogroup and checkboxgroup have the same input name
	"radiogroup":    nil, // A standalone radio makes no sense; only a radiogroup.
	"checkboxgroup": nil, // checkboxgroup has no *sensible* use case. There was an 'amenities' array in another app, with encodings: 4 for bath, 8 for balcony... They should better be designed as independent checkboxes bath and balcony. I cannot think of any useful 'additive flags', and those would have to be added and decoded server side. We keep the type, but untested.

	// Helpers
	"textblock": nil, // Only name, label and description are rendered
	"button":    nil, // Only name, label and description are rendered
}

const (
	// Checkbox inputs need standardized values for unchecked and checked
	// ValEmpty is returned, if the checkbox was unchecked
	valEmpty = "0"
	// ValSet is returned, if the checkbox was checked
	ValSet    = "1"
	vspacer   = "<div class='go-quest-vspacer'> &nbsp; </div>\n"
	vspacer16 = "<div class='go-quest-vspacer-16'> &nbsp; </div>\n"
)

type horizontalAlignment int

const (
	// HLeft encodes left horizontal alignment
	HLeft = horizontalAlignment(0)
	// HCenter encodes centered horizontal alignment
	HCenter = horizontalAlignment(1)
	// HRight encodes right horizontal alignment
	HRight = horizontalAlignment(2)
)

// String converts the value to a CSS compliant string
func (h horizontalAlignment) String() string {
	switch h {
	case horizontalAlignment(0):
		return "left"
	case horizontalAlignment(1):
		return "center"
	case horizontalAlignment(2):
		return "right"
	}
	return "left"
}

// On colsTotal == 0  division by zero case:
// We return no CSS.
// 		=> No width restriction - elements grow horizontally as much as needed
func colWidth(colsElement, colsTotal int) string {
	css := ""
	if colsTotal < 1 { // Prevent any division by zero
		return css
	}

	if colsElement == 0 {
		colsElement = 1
	}

	fract := float32(colsElement) * float32(97.5) / float32(colsTotal)
	if fract > 100.0 {
		fract = 100
	}
	fractStr := fmt.Sprintf("%4.1f", fract)
	css = fmt.Sprintf("width: %v%%;", fractStr)
	return css
}
