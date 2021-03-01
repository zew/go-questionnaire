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

// Styles groups all possible styles;
// there is no point in separating different sets
type Styles struct {
	GridContainerStyle `json:"grid_container_style,omitempty"`
	BoxStyle           `json:"box_style,omitempty"`
	GridItemStyle      `json:"grid_item_style,omitempty"`
	TextStyle          `json:"text_style,omitempty"`
}

// StylesResponsive contains styles for CSS grid container
type StylesResponsive struct {
	Desktop Styles `json:"desktop,omitempty"`
	Mobile  Styles `json:"mobile,omitempty"`
	// Classes []string `json:"classes,omitempty"` // static CSS classes
}

// NewStylesResponsive returns style struct
func NewStylesResponsive() *StylesResponsive {
	return &StylesResponsive{
		Desktop: Styles{
			BoxStyle: BoxStyle{
				// HeightMin: "0", // default auto - stackoverflow.com/questions/43311943/
				// WidthMin:  "0",
			},
		},
		Mobile: Styles{},
	}
}

// stylesResponsiveExample to test
func stylesResponsiveExample() *StylesResponsive {
	grSt := NewStylesResponsive()

	grSt.Desktop.GridContainerStyle.AutoFlow = "row"
	grSt.Desktop.GridContainerStyle.TemplateColumns = "minmax(4rem, 2fr) minmax(4rem, 2fr) minmax(4rem, 2fr)"

	grSt.Mobile.GridContainerStyle.AutoFlow = "col"
	return grSt
}

func stylesResponsiveExampleWant(className string) string {
	return fmt.Sprintf(`.%v {
	/* grid-container */
	grid-auto-flow: row;
	grid-template-columns: minmax(4rem, 2fr) minmax(4rem, 2fr) minmax(4rem, 2fr);

}
@media screen and (max-width: 800px) {
.%v {
	/* grid-container */
	grid-auto-flow: col;

}
}

`, className, className)
}

func notEmpty(prefix, s, suffix string) string {
	if s != "" {
		return fmt.Sprintf("%v%v%v", prefix, s, suffix)
	}
	return ""
}

// StyleTag wraps s into a style tag
func StyleTag(content string) string {
	s := &strings.Builder{}
	fmt.Fprint(s, "<style>\n")
	fmt.Fprint(s, content)
	fmt.Fprint(s, "</style>\n\n")
	return s.String()
}

// CSS renders styles
func (gcr StylesResponsive) CSS(className string) string {
	w := &strings.Builder{}

	// desktop
	fmt.Fprintf(w, ".%v {\n", className)
	fmt.Fprint(w, notEmpty("\t/* box-style */\n", gcr.Desktop.BoxStyle.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* grid-container */\n", gcr.Desktop.GridContainerStyle.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* grid-item */\n", gcr.Desktop.GridItemStyle.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* text-style */\n", gcr.Desktop.TextStyle.CSS(), "\n"))
	fmt.Fprint(w, "}\n")

	// mobile
	wMob := &strings.Builder{}
	fmt.Fprint(wMob, notEmpty("\t/* box-style */\n", gcr.Mobile.BoxStyle.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* grid-container */\n", gcr.Mobile.GridContainerStyle.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* grid-item */\n", gcr.Mobile.GridItemStyle.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* text-style */\n", gcr.Mobile.TextStyle.CSS(), "\n"))
	if wMob.Len() > 0 {
		fmt.Fprint(w, "@media screen and (max-width: 800px) {\n")
		fmt.Fprintf(w, ".%v {\n", className)
		fmt.Fprint(w, wMob.String())
		fmt.Fprint(w, "}\n")
		fmt.Fprint(w, "}\n")
	}

	fmt.Fprint(w, "\n")

	return w.String()
}
