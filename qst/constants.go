package qst

import "fmt"

var implementedTypes = map[string]interface{}{
	"text":     nil,
	"checkbox": nil, // A standalone checkbox - as a group, see below

	// radiogroup and checkboxgroup have the same input name
	"radiogroup":    nil, // A standalone radio makes no sense; only a radiogroup.
	"checkboxgroup": nil, // checkboxgroup has no *sensible* use case. There was an 'amenities' array in another app, with encodings: 4 for bath, 8 for balcony... They should better be designed as independent checkboxes bath and balcony. I cannot think of any useful 'additive flags', and those would have to be added and decoded server side. We keep the type, but untested.

	// Helpers
	"textblock": nil, // Only name and description are rendered
}

// checkbox inputs need standardized values for unchecked and checked
const valEmpty = "0"
const valSet = "1"
const vspacer = "<div class='go-quest-vspacer'> &nbsp; </div>\n"
const vspacer16 = "<div class='go-quest-vspacer-16'> &nbsp; </div>\n"

type horizontalAlignment int

const (
	HLeft   = horizontalAlignment(0)
	HCenter = horizontalAlignment(1)
	HRight  = horizontalAlignment(2)
)

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
// => No width restriction - elements grow horizontally as much as needed
func colWidth(colsElement, colsTotal int) string {
	css := ""
	if colsTotal > 0 {

		if colsElement == 0 {
			colsElement = 1
		}

		fract := float32(colsElement) * float32(97.5) / float32(colsTotal)
		fractStr := fmt.Sprintf("%4.1f", fract)
		css = fmt.Sprintf("width: %v%%;", fractStr)
	}
	return css
}
