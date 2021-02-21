package css

import (
	"fmt"
	"strings"
)

// CSSer can render to CSS styles
type CSSer interface {
	CSS() string
}

// Common styles
type Common struct {
	boxStyle  `json:"box_style,omitempty"`
	textStyle `json:"text_style,omitempty"`
	Other     string `json:"other,omitempty"`
}

// GridItem styles
type GridItem struct {
	gridItemStyle `json:"grid_item_style,omitempty"`
	boxStyle      `json:"box_style,omitempty"`
	textStyle     `json:"text_style,omitempty"`
	Other         string `json:"other,omitempty"`
}

// GridContainer styles
type GridContainer struct {
	gridContainerStyle `json:"grid_container_style,omitempty"`
	boxStyle           `json:"box_style,omitempty"`
	// do we need text styles ???
	// textStyle `json:"text_style,omitempty"`
	Other string `json:"other,omitempty"`
}

// CSS renders styles
func (cm Common) CSS() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "%v\n", cm.boxStyle)
	fmt.Fprintf(s, "%v\n", cm.textStyle)
	fmt.Fprintf(s, "%v\n", cm.Other)
	return s.String()
}

// CSS renders styles
func (gi GridItem) CSS() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "%v\n", gi.boxStyle)
	fmt.Fprintf(s, "%v\n", gi.gridItemStyle)
	fmt.Fprintf(s, "%v\n", gi.textStyle)
	fmt.Fprintf(s, "%v\n", gi.Other)
	return s.String()
}

// CSS renders styles
func (gc GridContainer) CSS() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "%v\n", gc.boxStyle)
	fmt.Fprintf(s, "%v\n", gc.gridContainerStyle)
	fmt.Fprintf(s, "%v\n", gc.Other)
	return s.String()
}

// Responsive contains styles for two different media
type Responsive struct {
	Desktop CSSer
	Mobile  CSSer
}

/* // GridItem holds the styles for a grid item
type GridItem struct {
	Self    Responsive
	Label   Responsive
	Control Responsive
}

// GridContainer holds the styles for a grid container
type GridContainer struct {
	Responsive
}
*/
