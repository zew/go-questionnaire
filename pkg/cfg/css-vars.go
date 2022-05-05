package cfg

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
)

/*
	cssVar contains CSS variables -
	which go from config JSON file into CSS blocks of HTML templates.
	For instance
	--nav-height:            8vh;
	--clr-pri:               rgba(  0, 105, 180, 0.999);    slash* primary color *slash
*/
/* */
type cssVar struct {
	Key   string `json:"key,omitempty"`
	Val   string `json:"val,omitempty"`
	Desc  string `json:"desc,omitempty"`
	IsURL bool   `json:"is_url,omitempty"` // prepend val with prefix and clad into url() => content: url(/taxkit/img/ui/icon-forschung-zew-prim.svg

	// if val is empty - following fields for color become relevant
	Colorname string  `json:"color_name,omitempty"` // alternative to RGB-Alpha: 'white' or 'cyan' - replaces R/G/B - has no alpha value
	R         uint8   `json:"r,omitempty"`
	G         uint8   `json:"g,omitempty"`
	B         uint8   `json:"b,omitempty"`
	Alpha     float32 `json:"alpha,omitempty"`
}

func addComment(s, desc string) string {
	if desc == "" {
		return s + "\n"
	}
	repeat := ""
	if 56 > len(s) {
		repeat = strings.Repeat(" ", 56-len(s))
	}
	return fmt.Sprintf("%s %s /* %s */\n", s, repeat, desc)
}

// Plain template data - for instance logo-text
func (c cssVar) Plain() string {
	return c.Val
}

// for HTML header <meta name="theme-color"...  - ColorHex not needed
func (c cssVar) RGBA() string {
	if c.Alpha == 0 {
		return fmt.Sprintf("rgb(%3v, %3v, %3v)", c.R, c.G, c.B)
	}
	return fmt.Sprintf("rgba(%3v, %3v, %3v, %5.3f)", c.R, c.G, c.B, c.Alpha)
}

// for HTML header <meta name="theme-color"...  - but RGBA works as well
func (c cssVar) ColorHex() string {
	return fmt.Sprintf("#%X%X%X", c.R, c.G, c.B)
}

// HTML rennders a single CSS var into HTML
// i.e. --clr-pri:               rgba(  0, 105, 180, 0.999);    /* primary color */
func (c cssVar) HTML() string {

	// key-val
	if c.Val != "" {
		repeat := ""
		if 20 > len(c.Key) {
			repeat = strings.Repeat(" ", 20-len(c.Key))
		}
		if c.IsURL && c.Val != "none" {
			c.Val = fmt.Sprintf("url(%v)", Pref(c.Val))
		}
		s := fmt.Sprintf("\t\t--%v: %v %v;", c.Key, repeat, c.Val)
		return addComment(s, c.Desc)
	}

	// color
	clr := ""
	if c.Colorname != "" {
		clr = c.Colorname
	} else {
		if c.Alpha == 0.00 {
			clr = fmt.Sprintf("rgb(%3v, %3v, %3v)", c.R, c.G, c.B)
		} else {
			clr = fmt.Sprintf("rgba(%3v, %3v, %3v, %5.3f)", c.R, c.G, c.B, c.Alpha)
		}
	}

	repeat := ""
	if 16 > len(c.Key) {
		repeat = strings.Repeat(" ", 16-len(c.Key))
	}
	//                    --clr-pri-vis:    rgba(  0,  71, 122, 0.999);  /* slightly darker */
	s := fmt.Sprintf("\t\t--clr-%s: %s %s;", c.Key, repeat, clr)
	return addComment(s, c.Desc)

}

type cssVars []cssVar

// HTML prints all CSS vars into as list for the template header
func (cs cssVars) HTML() template.CSS {
	b := &bytes.Buffer{}
	fmt.Fprint(b, "\n")
	fmt.Fprint(b, "\t\t/* CSS vars from config - start */\n")
	for _, c := range cs {
		fmt.Fprintf(b, "%v", c.HTML())
	}
	fmt.Fprint(b, "\t\t/* CSS vars from config - stop */\n")
	return template.CSS(b.String())
}

// ByKey returns a single CSS var by key
func (cs cssVars) ByKey(k string) cssVar {
	for _, c := range cs {
		if c.Key == k {
			return c
		}
	}
	return cssVar{Key: "css-var-not-found", Val: k}
}

// Stack combines / merges addenum into base
func Stack(base, addenum cssVars) cssVars {

	// copy of base
	ret := make(cssVars, 0, len(base))
	for _, c := range base {
		ret = append(ret, c)
	}

	// clobber addenum over base
	for i1, cadd := range addenum {
		found := false
		for i2, c := range ret {
			if cadd.Key == c.Key {
				// log.Printf("  overwr %-10v => %10v\n", c.Key, addenum[i1])
				ret[i2] = addenum[i1]
				found = true
				break
			}
		}
		if !found {
			// not contained in base => not overwrite - but append
			// log.Printf("  appending \n\t\t%+v", cadd)
			ret = append(ret, cadd)
		}
	}
	return ret
}
