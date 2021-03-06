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
	A, B string
	C    int
}

// NewStylesResponsive returns style struct
// if the arg is nil
func NewStylesResponsive(sr *StylesResponsive) *StylesResponsive {
	if sr != nil {
		return sr
	}
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
	sr := NewStylesResponsive(nil)

	sr.Desktop.GridContainerStyle.AutoFlow = "row"
	sr.Desktop.GridContainerStyle.TemplateColumns = "minmax(4rem, 2fr) minmax(4rem, 2fr) minmax(4rem, 2fr)"

	sr.Mobile.GridContainerStyle.AutoFlow = "col"

	sr.A = "prop a"
	sr.B = "prop b"
	sr.C = 17
	return sr
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
func (sr StylesResponsive) CSS(className string) string {
	w := &strings.Builder{}

	// desktop
	fmt.Fprintf(w, ".%v {\n", className)
	fmt.Fprint(w, notEmpty("\t/* box-style */\n", sr.Desktop.BoxStyle.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* grid-container */\n", sr.Desktop.GridContainerStyle.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* grid-item */\n", sr.Desktop.GridItemStyle.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* text-style */\n", sr.Desktop.TextStyle.CSS(), "\n"))
	fmt.Fprint(w, "}\n")

	// mobile
	wMob := &strings.Builder{}
	fmt.Fprint(wMob, notEmpty("\t/* box-style */\n", sr.Mobile.BoxStyle.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* grid-container */\n", sr.Mobile.GridContainerStyle.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* grid-item */\n", sr.Mobile.GridItemStyle.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* text-style */\n", sr.Mobile.TextStyle.CSS(), "\n"))
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

// Combine copies b over sr - if sr has no value
func (sr *StylesResponsive) Combine(b StylesResponsive) {
	sr.Desktop.GridContainerStyle.Combine(b.Desktop.GridContainerStyle)
	sr.Desktop.BoxStyle.Combine(b.Desktop.BoxStyle)
	sr.Desktop.GridItemStyle.Combine(b.Desktop.GridItemStyle)
	sr.Desktop.TextStyle.Combine(b.Desktop.TextStyle)

	sr.Mobile.GridContainerStyle.Combine(b.Mobile.GridContainerStyle)
	sr.Mobile.BoxStyle.Combine(b.Mobile.BoxStyle)
	sr.Mobile.GridItemStyle.Combine(b.Mobile.GridItemStyle)
	sr.Mobile.TextStyle.Combine(b.Mobile.TextStyle)
}

// // Alloc makes sure, sr is not nil
// func (sr *StylesResponsive) Alloc() *StylesResponsive {
// 	if sr == nil {
// 		return NewStylesResponsive()
// 	}
// 	return sr
// }

// ItemCenteredMCA makes the input centered on main and cross axis (MCA)
func ItemCenteredMCA(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.GridItemStyle.JustifySelf = "center"
	sr.Desktop.GridItemStyle.AlignSelf = "center"
	sr.Desktop.TextStyle.AlignHorizontal = "center"
	return sr
}

// ItemStartCA aligns the item at the start on the cross-axis
func ItemStartCA(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.GridItemStyle.AlignSelf = "start"
	return sr
}

// ItemEndMA aligns the item at the end on the main-axis
func ItemEndMA(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.GridItemStyle.JustifySelf = "end"
	return sr
}

// TextStart makes the text content left aligned
func TextStart(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	// sr.Desktop.GridItemStyle.JustifySelf = "start"  // fails on multi line text
	sr.Desktop.TextStyle.AlignHorizontal = "left"
	return sr
}

// TextEnd makes the text content right aligned
func TextEnd(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	// sr.Desktop.GridItemStyle.JustifySelf = "start"  // fails on multi line text
	sr.Desktop.TextStyle.AlignHorizontal = "right"
	return sr
}
