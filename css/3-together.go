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
	StyleGridContainer `json:"grid_container_style,omitempty"`
	StyleBox           `json:"box_style,omitempty"`
	StyleGridItem      `json:"grid_item_style,omitempty"`
	StyleText          `json:"text_style,omitempty"`
}

// StylesResponsive contains styles for CSS grid container
type StylesResponsive struct {
	Desktop Styles `json:"desktop,omitempty"`
	Mobile  Styles `json:"mobile,omitempty"`

	// Classes []string `json:"classes,omitempty"` // static CSS classes
}

// NewStylesResponsive returns style struct
// if the arg is nil
func NewStylesResponsive(sr *StylesResponsive) *StylesResponsive {
	if sr != nil {
		return sr
	}
	return &StylesResponsive{
		Desktop: Styles{
			StyleBox: StyleBox{
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

	sr.Desktop.StyleGridContainer.AutoFlow = "row"
	sr.Desktop.StyleGridContainer.TemplateColumns = "minmax(4rem, 2fr) minmax(4rem, 2fr) minmax(4rem, 2fr)"

	sr.Mobile.StyleGridContainer.AutoFlow = "col"

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
	fmt.Fprint(w, notEmpty("\t/* box-style */\n", sr.Desktop.StyleBox.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* grid-container */\n", sr.Desktop.StyleGridContainer.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* grid-item */\n", sr.Desktop.StyleGridItem.CSS(), "\n"))
	fmt.Fprint(w, notEmpty("\t/* text-style */\n", sr.Desktop.StyleText.CSS(), "\n"))
	fmt.Fprint(w, "}\n")

	// mobile
	wMob := &strings.Builder{}
	fmt.Fprint(wMob, notEmpty("\t/* box-style */\n", sr.Mobile.StyleBox.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* grid-container */\n", sr.Mobile.StyleGridContainer.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* grid-item */\n", sr.Mobile.StyleGridItem.CSS(), "\n"))
	fmt.Fprint(wMob, notEmpty("\t/* text-style */\n", sr.Mobile.StyleText.CSS(), "\n"))
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
	sr.Desktop.StyleGridContainer.Combine(b.Desktop.StyleGridContainer)
	sr.Desktop.StyleBox.Combine(b.Desktop.StyleBox)
	sr.Desktop.StyleGridItem.Combine(b.Desktop.StyleGridItem)
	sr.Desktop.StyleText.Combine(b.Desktop.StyleText)

	sr.Mobile.StyleGridContainer.Combine(b.Mobile.StyleGridContainer)
	sr.Mobile.StyleBox.Combine(b.Mobile.StyleBox)
	sr.Mobile.StyleGridItem.Combine(b.Mobile.StyleGridItem)
	sr.Mobile.StyleText.Combine(b.Mobile.StyleText)
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
	sr.Desktop.StyleGridItem.JustifySelf = "center"
	sr.Desktop.StyleGridItem.AlignSelf = "center"
	sr.Desktop.StyleText.AlignHorizontal = "center"
	return sr
}

// ItemCenteredCA makes the input centered on main and cross axis (MCA)
func ItemCenteredCA(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.StyleGridItem.AlignSelf = "center"
	sr.Desktop.StyleText.AlignHorizontal = "center"
	return sr
}

// ItemStartCA aligns the item at the start on the cross-axis
func ItemStartCA(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.StyleGridItem.AlignSelf = "start"
	return sr
}

// ItemEndMA aligns the item at the end on the main-axis
func ItemEndMA(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.StyleGridItem.JustifySelf = "end"
	return sr
}

// TextStart makes the text content left aligned
func TextStart(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	// sr.Desktop.GridItemStyle.JustifySelf = "start"  // fails on multi line text
	sr.Desktop.StyleText.AlignHorizontal = "left"
	return sr
}

// TextEnd makes the text content right aligned
func TextEnd(sr *StylesResponsive) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	// sr.Desktop.GridItemStyle.JustifySelf = "start"  // fails on multi line text
	sr.Desktop.StyleText.AlignHorizontal = "right"
	return sr
}

// DesktopWidthMax limits width in desktop view
// for instance to 30rem
func DesktopWidthMax(sr *StylesResponsive, s string) *StylesResponsive {
	sr = NewStylesResponsive(sr)
	sr.Desktop.StyleBox.WidthMax = s
	sr.Mobile.StyleBox.WidthMax = "none"
	return sr
}
