package css

import (
	"fmt"
	"strings"
)

// BoxStyle styles
type BoxStyle struct {
	Display  string `json:"display,omitempty"`
	Position string `json:"position,omitempty"`
	ZIndex   int    `json:"z_index,omitempty"`
	Left     string `json:"left,omitempty"`
	Top      string `json:"top,omitempty"`
	Bottom   string `json:"bottom,omitempty"`
	Right    string `json:"right,omitempty"`

	Width     string `json:"width,omitempty"`
	WidthMin  string `json:"width_min,omitempty"`
	WidthMax  string `json:"width_max,omitempty"`
	Height    string `json:"height,omitempty"`
	HeightMin string `json:"height_min,omitempty"`
	HeightMax string `json:"height_max,omitempty"`

	Margin  string `json:"margin,omitempty"`  /*  vertical | horizontal  -  top | horizontal | bottom   -   top | right | bottom | left */
	Padding string `json:"padding,omitempty"` /*  vertical | horizontal  -  top | horizontal | bottom   -   top | right | bottom | left */
	Border  string `json:"border,omitempty"`
	Outline string `json:"outline,omitempty"`

	BackgroundColor string `json:"background_color,omitempty"`
	BorderRadius    string `json:"border_radius,omitempty"`
}

func boxStyleExample1() BoxStyle {
	return BoxStyle{
		Display:         "block",
		Position:        "relative",
		ZIndex:          103,
		Bottom:          "0.2rem",
		Right:           "0.3rem",
		WidthMax:        "920px",
		HeightMin:       "480px",
		Margin:          "0.2rem 2rem 0.3rem",
		Padding:         "0.2rem 2rem 0.3rem 3rem",
		Border:          "1px solid green",
		BackgroundColor: "rbga(255,244,255,0.5)",
		BorderRadius:    "0.2rem",
	}
}

func boxStyleExample1Want() string {
	return `	display: block;
	position: relative;
	z-index: 103;
	bottom: 0.2rem;
	right: 0.3rem;
	max-width: 920px;
	min-height: 480px;
	margin: 0.2rem 2rem 0.3rem;
	padding: 0.2rem 2rem 0.3rem 3rem;
	border: 1px solid green;
	background-color: rbga(255,244,255,0.5);
	border-radius: 0.2rem;
`
}

// CSS renders styles
func (bs BoxStyle) CSS() string {
	s := &strings.Builder{}
	if bs.Display != "" {
		fmt.Fprintf(s, "\tdisplay: %v;\n", bs.Display)
	}
	if bs.Position != "" {
		fmt.Fprintf(s, "\tposition: %v;\n", bs.Position)
	}
	if bs.ZIndex != 0 {
		fmt.Fprintf(s, "\tz-index: %v;\n", bs.ZIndex)
	}
	if bs.Left != "" {
		fmt.Fprintf(s, "\tleft: %v;\n", bs.Left)
	}
	if bs.Top != "" {
		fmt.Fprintf(s, "\ttop: %v;\n", bs.Top)
	}
	if bs.Bottom != "" {
		fmt.Fprintf(s, "\tbottom: %v;\n", bs.Bottom)
	}
	if bs.Right != "" {
		fmt.Fprintf(s, "\tright: %v;\n", bs.Right)
	}

	if bs.Width != "" {
		fmt.Fprintf(s, "\twidth: %v;\n", bs.Width)
	}
	if bs.WidthMin != "" {
		fmt.Fprintf(s, "\tmin-width: %v;\n", bs.WidthMin)
	}
	if bs.WidthMax != "" {
		fmt.Fprintf(s, "\tmax-width: %v;\n", bs.WidthMax)
	}
	if bs.Height != "" {
		fmt.Fprintf(s, "\theight: %v;\n", bs.Height)
	}
	if bs.HeightMin != "" {
		fmt.Fprintf(s, "\tmin-height: %v;\n", bs.HeightMin)
	}
	if bs.HeightMax != "" {
		fmt.Fprintf(s, "\tmax-height: %v;\n", bs.HeightMax)
	}

	if bs.Margin != "" {
		fmt.Fprintf(s, "\tmargin: %v;\n", bs.Margin)
	}
	if bs.Padding != "" {
		fmt.Fprintf(s, "\tpadding: %v;\n", bs.Padding)
	}
	if bs.Border != "" {
		fmt.Fprintf(s, "\tborder: %v;\n", bs.Border)
	}
	if bs.Outline != "" {
		fmt.Fprintf(s, "\toutline: %v;\n", bs.Outline)
	}

	if bs.BackgroundColor != "" {
		fmt.Fprintf(s, "\tbackground-color: %v;\n", bs.BackgroundColor)
	}
	if bs.BorderRadius != "" {
		fmt.Fprintf(s, "\tborder-radius: %v;\n", bs.BorderRadius)
	}

	return s.String()
}

// TextStyle styles
type TextStyle struct {
	FontFamily string `json:"font_family,omitempty"`
	FontSize   int    `json:"font_size,omitempty"` // percent
	Color      string `json:"color,omitempty"`

	HorizontalAlign string `json:"horizontal_align,omitempty"` // left, right, center, justify
	VerticalAlign   string `json:"vertical_align,omitempty"`   // baseline, bottom, top, middle

	LineHeight    int    `json:"line_height,omitempty"`
	LetterSpacing string `json:"letter_spacing,omitempty"`

	WhiteSpace string `json:"white_space,omitempty"`
}

func textStyleExample1() TextStyle {
	return TextStyle{
		FontFamily:      "'Segoe UI', Tahoma, Geneva, Verdana",
		FontSize:        120,
		Color:           "#CCC",
		HorizontalAlign: "justify",
		VerticalAlign:   "middle",
		LineHeight:      110,
		WhiteSpace:      "pre-wrap",
	}
}

func textStyleExample1Want() string {
	return `	font-family: 'Segoe UI', Tahoma, Geneva, Verdana;
	font-size: 120%;
	color: #CCC;
	align-text: justify;
	vertical-align: middle;
	line-height: 110%;
	white-space: pre-wrap;
`
}

// CSS renders styles
func (ts TextStyle) CSS() string {
	s := &strings.Builder{}
	if ts.FontFamily != "" {
		fmt.Fprintf(s, "\tfont-family: %v;\n", ts.FontFamily)
	}
	if ts.FontSize != 0 {
		fmt.Fprintf(s, "\tfont-size: %v%%;\n", ts.FontSize)
	}
	if ts.Color != "" {
		fmt.Fprintf(s, "\tcolor: %v;\n", ts.Color)
	}
	if ts.HorizontalAlign != "" {
		fmt.Fprintf(s, "\talign-text: %v;\n", ts.HorizontalAlign)
	}
	if ts.VerticalAlign != "" {
		fmt.Fprintf(s, "\tvertical-align: %v;\n", ts.VerticalAlign)
	}
	if ts.LineHeight != 0 {
		fmt.Fprintf(s, "\tline-height: %v%%;\n", ts.LineHeight)
	}
	if ts.LetterSpacing != "" {
		fmt.Fprintf(s, "\tletter-spacing: %v;\n", ts.LetterSpacing)
	}
	if ts.WhiteSpace != "" {
		fmt.Fprintf(s, "\twhite-space: %v;\n", ts.WhiteSpace)
	}
	return s.String()
}
