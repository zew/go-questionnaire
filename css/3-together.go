package css

import (
	"fmt"
	"strings"
)

// CSSerSimple renders just styles - without a class enclosure
type CSSerSimple interface {
	CSS() string
}

// CSSer renders styles and CSS class enclosure
type CSSer interface {
	CSS(className string) string
}

/*



 */

// Core styles - neither CSS container nor CSS item
type Core struct {
	BoxStyle  `json:"box_style,omitempty"`
	TextStyle `json:"text_style,omitempty"`
}

// GridItem styles
type GridItem struct {
	GridItemStyle `json:"grid_item_style,omitempty"`
	BoxStyle      `json:"box_style,omitempty"`
	TextStyle     `json:"text_style,omitempty"`
}

// GridContainer styles
type GridContainer struct {
	GridContainerStyle `json:"grid_container_style,omitempty"`
	BoxStyle           `json:"box_style,omitempty"`
}

// CoreResponsive contains styles for labels or controls
type CoreResponsive struct {
	Desktop Core `json:"desktop,omitempty"`
	Mobile  Core `json:"mobile,omitempty"`
}

// NewCore returns style struct
func NewCore() *CoreResponsive {
	return &CoreResponsive{
		Desktop: Core{},
		Mobile:  Core{},
	}
}

// GridContainerResponsive contains styles for CSS grid container
type GridContainerResponsive struct {
	Desktop GridContainer `json:"desktop,omitempty"`
	Mobile  GridContainer `json:"mobile,omitempty"`
}

// NewGridContainer returns style struct
func NewGridContainer() *GridContainerResponsive {
	return &GridContainerResponsive{
		Desktop: GridContainer{},
		Mobile:  GridContainer{},
	}
}

// GridContainerResponsiveExample to test
func GridContainerResponsiveExample() *GridContainerResponsive {
	grSt := NewGridContainer()

	grSt.Desktop.GridContainerStyle.AutoFlow = "row"
	grSt.Desktop.GridContainerStyle.TemplateColumns = "minmax(4rem, 2fr) minmax(4rem, 2fr) minmax(4rem, 2fr)"

	grSt.Mobile.GridContainerStyle.AutoFlow = "col"
	return grSt
}

func gridContainerResponsiveExampleWant(className string) string {
	return fmt.Sprintf(`<style>
.%v {
	/* grid-container */
	grid-template-columns: minmax(4rem, 2fr) minmax(4rem, 2fr) minmax(4rem, 2fr);

}
@media screen and (max-width: 800px) {
.%v {
	/* grid-container */
	grid-auto-flow: row;

}
}
</style>

`, className, className)
}

func notEmpty(prefix, s, suffix string) string {
	if s != "" {
		return fmt.Sprintf("%v%v%v", prefix, s, suffix)
	}
	return ""
}

// CSS renders styles
func (gcr GridContainerResponsive) CSS(className string) string {
	s := &strings.Builder{}

	fmt.Fprintf(s, "<style>\n")

	fmt.Fprintf(s, ".%v {\n", className)
	fmt.Fprint(s, notEmpty("\t/* box-style */\n", gcr.Desktop.BoxStyle.CSS(), "\n"))
	fmt.Fprint(s, notEmpty("\t/* grid-container */\n", gcr.Desktop.GridContainerStyle.CSS(), "\n"))
	fmt.Fprintf(s, "}\n")

	fmt.Fprintf(s, "@media screen and (max-width: 800px) {\n")
	fmt.Fprintf(s, ".%v {\n", className)
	fmt.Fprint(s, notEmpty("\t/* box-style */\n", gcr.Mobile.BoxStyle.CSS(), "\n"))
	fmt.Fprint(s, notEmpty("\t/* grid-container */\n", gcr.Mobile.GridContainerStyle.CSS(), "\n"))
	fmt.Fprintf(s, "}\n")
	fmt.Fprintf(s, "}\n")

	fmt.Fprintf(s, "</style>\n\n")
	return s.String()
}

// GridItemResponsive contains styles for CSS grid item
type GridItemResponsive struct {
	Desktop GridItem `json:"desktop,omitempty"`
	Mobile  GridItem `json:"mobile,omitempty"`
}

// NewGridItem returns style struct
func NewGridItem() *GridItemResponsive {
	return &GridItemResponsive{
		Desktop: GridItem{},
		Mobile:  GridItem{},
	}
}
