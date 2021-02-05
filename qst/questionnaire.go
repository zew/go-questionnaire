// Package qst implements a four levels deep nested structure
// with input controls, groups, pages and questionnaire;
// contains HTML rendering, page navigation,
// loading/saving from/to JSON file, consistence validation,
// multi-language support.
package qst

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"path"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/lgn/shuffler"
	"github.com/zew/go-questionnaire/sessx"
	"github.com/zew/go-questionnaire/trl"
	"github.com/zew/util"

	"github.com/zew/go-questionnaire/ctr"
)

// No line wrapping between element 1 and 2
//
//  But line wrapping *inside* each of them.
//
//  el1 and el2 must be inline-block, for whitespace nowrap to work.
func nobreakGlue(el1, glue, el2 string) string {

	if el1 == "" || el2 == "" {
		return el1 + el2
	}

	reduction := 82 // 	el2 is overflowing :(
	reduction = 95  // 2020-05 - relaxed

	el1 = strings.TrimSpace(el1) // includes \n
	el2 = strings.TrimSpace(el2)

	el1 = fmt.Sprintf(
		"<span style='white-space: normal; display: inline-block; vertical-align: top;' >%v</span>",
		el1,
	)
	el2 = fmt.Sprintf(
		"<span style='white-space: normal; display: inline-block; vertical-align: top; ' >%v</span>",
		el2,
	)
	ret := fmt.Sprintf(
		"<span style='white-space: nowrap; display: inline-block; width: %v%%;'>%v%v%v</span>\n",
		reduction,
		el1, glue, el2,
	)
	return ret
}

// no wrap between input and suffix
func appendSuffix(ctrl string, i *inputT, langCode string) string {

	if !i.Suffix.Set() {
		return ctrl
	}

	ctrl = strings.TrimSuffix(ctrl, "\n")
	// We want to prevent line-break of the '%' or '€' suffix character.
	// inputs must be inline-block, for whitespace nowrap to work.
	// At the same time: suffix-inner enables wrapping for the suffix itself
	sfx := fmt.Sprintf("<span class='go-quest-label %v  suffix-inner' >%v</span>\n", i.CSSLabel, i.Suffix.TrSilent(langCode))
	ctrl = fmt.Sprintf("<span class='suffix-nowrap' >%v%v</span>\n", ctrl, sfx)

	return ctrl
}

// Special subtype of inputT; used for radiogroup
type radioT struct {
	HAlign horizontalAlignment `json:"hori_align,omitempty"` // label and description left/center/right of input, default left, similar setting for radioT but not for group
	Label  trl.S               `json:"label,omitempty"`
	Val    string              `json:"val,omitempty"`     // Val is allowed to be nil; it then gets initialized to 1...n by Validate(). 0 indicates 'no entry'.
	Col    int                 `json:"column,omitempty"`  // col x of cols
	Cols   int                 `json:"columns,omitempty"` //
	// field 'response' is absent, it is added dynamically;
}

func (i *inputT) AddRadio() *radioT {
	rad := &radioT{}
	i.Radios = append(i.Radios, rad)
	ret := i.Radios[len(i.Radios)-1]
	return ret
}

// Input represents a single form input element.
// There is one exception for multiple radios (radiogroup) with the same name but distinct values.
// Multiple checkboxes (checkboxgroup) with same name but distinct values are a dubious instrument.
// See comment to implementedType checkboxgroup.
type inputT struct {
	Name     string  `json:"name,omitempty"`
	Type     string  `json:"type,omitempty"`      // see implementedTypes
	MaxChars int     `json:"max_chars,omitempty"` // Number of input chars, also to compute width
	Step     float64 `json:"step,omitempty"`      // stepping interval for number input

	Label     trl.S  `json:"label,omitempty"`
	Desc      trl.S  `json:"description,omitempty"`
	Suffix    trl.S  `json:"suffix,omitempty"`
	AccessKey string `json:"accesskey,omitempty"`

	HAlignLabel   horizontalAlignment `json:"horizontal_align_label,omitempty"`   // description left/center/right of input, default left, similar setting for radioT but not for group
	HAlignControl horizontalAlignment `json:"horizontal_align_control,omitempty"` // label       left/center/right of input, default left, similar setting for radioT but not for group

	// extra styling - a CSS class must exist
	CSSLabel   string `json:"css_label,omitempty"`   // vertical margins, line-height, indent - usually for the entire label+input
	CSSControl string `json:"css_control,omitempty"` // usually only for the input element's inner style

	// How many column slots of the overall layout should the control occupy?
	// The number adds up against group.Cols - determining newlines.
	// The number is used to compute the relative width (percentage).
	// If zero, a column width of one is assumend.
	ColSpanLabel   int `json:"col_span_label,omitempty"`
	ColSpanControl int `json:"col_span_control,omitempty"`

	Radios []*radioT  `json:"radios,omitempty"`    // This slice implements the radiogroup - and the senseless checkboxgroup
	DD     *DropdownT `json:"drop_down,omitempty"` // As pointer to prevent JSON cluttering

	Validator string `json:"validator,omitempty"` // i.e. any key from validators, i.e. "must;inRange20"
	ErrMsg    trl.S  `json:"err_msg,omitempty"`

	Response string `json:"response,omitempty"` // also contains the Value of options and checkboxes
	//  ResponseFloat float64  - floats and integers are stored as strings in Response
	DynamicFunc string `json:"dynamic_func,omitempty"` // compositFunc for type == 'composit' OR dynFunc for type == 'dynamic'
}

// NewInput returns an input filled in with globally enumerated label, decription etc.
func NewInput() inputT {
	cntr := ctr.Increment()
	t := inputT{
		Name:  fmt.Sprintf("input_%v", cntr),
		Type:  "text",
		Label: trl.S{"en": fmt.Sprintf("Label %v", cntr), "de": fmt.Sprintf("Titel %v", cntr)},
		Desc:  trl.S{"en": "Description", "de": "Beschreibung"},
	}
	return t
}

func renderLabelDescription(i inputT, langCode string, numCols int) string {
	return renderLabelDescription2(i, langCode, i.Name, i.HAlignLabel,
		i.Label, i.Desc, i.CSSLabel, i.ColSpanLabel, numCols)
}

// renderLabelDescription wraps lbl+desc into a <span> of class 'go-quest-cell' or td-cell.
// A percent width is dynamically computed from colsLabel / numCols.
// Argument numCols is the total number of cols per row.
// It is used to compute the precise width in percent
func renderLabelDescription2(i inputT, langCode string, nm string, hAlign horizontalAlignment,
	lbl, desc trl.S, css string, colsLabel, numCols int) string {
	ret := ""
	if lbl == nil && desc == nil {
		return ret
	}
	e1 := lbl.Tr(langCode)
	if lbl == nil {
		e1 = "" // Suppress "Translation map not initialized." here
	}
	e2 := desc.Tr(langCode)
	if desc == nil {
		e2 = "" // Suppress "Translation map not initialized." here
	}

	ret = fmt.Sprintf(
		"<span class='go-quest-label %v'><b>%v</b> %v </span>\n", css, e1, e2,
	)

	if nm != "" && !i.IsLayout() {
		ret = fmt.Sprintf("<label for='%v'>%v</label>\n", nm, ret)
	}

	ret = td(hAlign, colWidth(colsLabel, numCols), ret)
	return ret
}

// IsLayout returns whether the input type is merely ornamental
func (i inputT) IsLayout() bool {
	if i.Type == "textblock" {
		return true
	}
	if i.Type == "button" {
		return true
	}
	if i.Type == "dynamic" {
		return true
	}
	return false
}

// IsReserved returns whether the input name is reserved the survey engine
func (i inputT) IsReserved() bool {
	if i.Name == "page" {
		return true
	}
	if i.Name == "lang_code" {
		return true
	}
	return false
}

// Rendering one input to HTML
// func (i inputT) HTML(langCode string, namePrefix string) string {
func (i inputT) HTML(langCode string, numCols int) string {

	nm := i.Name

	switch i.Type {

	case "textblock":
		lbl := renderLabelDescription(i, langCode, numCols)
		return lbl

	case "button":
		lbl := fmt.Sprintf("<button type='submit' name='%v' value='%v' class='%v' accesskey='%v'><b>%v</b> %v</button>\n",
			i.Name, i.Response, i.CSSControl, i.AccessKey,
			i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode),
		)
		lbl = td(i.HAlignControl, colWidth(i.ColSpanControl, numCols), lbl)
		return lbl

	case "text", "number", "hidden", "textarea", "checkbox", "dropdown":

		ctrl := ""
		val := i.Response

		checked := ""
		if i.Type == "checkbox" {
			if val == ValSet {
				checked = "checked=\"checked\""
			}
			val = ValSet
		}

		width := fmt.Sprintf("width: %vem;", int(float64(i.MaxChars)*1.05))
		// width = "width: 98%;"
		if i.Type == "checkbox" || i.Type == "radio" || i.Type == "dropdown" {
			width = ""
		}
		maxChars := ""
		if i.MaxChars > 0 {
			maxChars = fmt.Sprintf(" MAXLENGTH='%v' ", i.MaxChars) // the right attribute for input and textarea
		}

		if i.Type == "textarea" {
			colsRows := fmt.Sprintf(" cols='%v' rows='1' ", i.MaxChars+1)
			if i.MaxChars > 80 {
				colsRows = fmt.Sprintf(" cols='80' rows='%v' ", i.MaxChars/80+1)
				// width = fmt.Sprintf("width: %vem;", int(float64(80)*1.05))
				width = "width: 98%;"
			}
			ctrl += fmt.Sprintf("<textarea        name='%v' id='%v' title='%v %v' class='%v' style='%v' %v %v>%v</textarea>\n",
				nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode), i.CSSControl, width, maxChars, colsRows, val)

		} else if i.Type == "dropdown" {

			// i.DD = &DropdownT{}
			i.DD.SetName(i.Name)
			i.DD.LC = langCode
			i.DD.SetTitle(i.Label.TrSilent(langCode) + " " + i.Desc.TrSilent(langCode))
			i.DD.Select(i.Response)
			i.DD.SetAttr("style", width)
			i.DD.SetAttr("class", i.CSSControl)

			sort.Sort(i.DD)

			ctrl += i.DD.RenderStr()

		} else {
			// input
			inputMode := ""
			if i.Type == "number" {
				if i.Step != 0 {
					if i.Step >= 1 {
						inputMode = fmt.Sprintf(" step='%.0f'  ", i.Step)
					} else {
						prec := int(math.Log10(1 / i.Step))
						f := fmt.Sprintf(" step='%%.%vf'  ", prec)
						inputMode = fmt.Sprintf(f, i.Step)
					}
				}
			}
			ctrl += fmt.Sprintf("<input type='%v'  %v  name='%v' id='%v' title='%v %v' class='%v' style='%v' %v %v  value='%v' />\n",
				i.Type, inputMode,
				nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode), i.CSSControl, width, maxChars, checked, val)
		}

		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if i.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

		// Append suffix
		ctrl = appendSuffix(ctrl, &i, langCode)

		// Append error message
		if i.ErrMsg.Set() {
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.ErrMsg.TrSilent(langCode))
		}

		ctrl = td(i.HAlignControl, colWidth(i.ColSpanControl, numCols), ctrl)

		lbl := renderLabelDescription(i, langCode, numCols)
		return lbl + ctrl

	case "radiogroup", "checkboxgroup":
		ctrl := ""
		innerType := "radio"
		if i.Type == "checkboxgroup" {
			innerType = "checkbox"
		}

		for radIdx, rad := range i.Radios {
			one := ""
			checked := ""
			if i.Response == rad.Val {
				checked = "checked=\"checked\""
			}

			radio := fmt.Sprintf(
				// 2021-01 - id must be unique
				"<input type='%v' name='%v' id-disabled='%v' title='%v %v' class='%v' value='%v' %v />\n",
				innerType, nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode),
				i.CSSControl,
				rad.Val, checked,
			)

			lbl := ""
			if !rad.Label.Set() {
				one = radio
			} else {
				if rad.HAlign == HLeft {
					lbl = fmt.Sprintf(
						"<span class=' vert-correct-left-right' >%v</span>\n",
						rad.Label.Tr(langCode),
					)
					one = nobreakGlue(lbl, "&nbsp;", radio)
				}
				if rad.HAlign == HCenter {
					// no i.CSSLabel to prevent left margins/paddings to uncenter
					lbl = fmt.Sprintf(
						"<span class=' vert-correct-center '>%v</span>\n",
						rad.Label.Tr(langCode),
					)
					lbl += vspacer
					one = lbl + radio
				}

				if rad.HAlign == HRight {
					lbl = fmt.Sprintf(
						"<span class=' vert-correct-left-right'>%v</span>\n",
						rad.Label.Tr(langCode),
					)
					one = nobreakGlue(radio, "&nbsp;", lbl)
				}
			}

			cssNoLeft := i.CSSLabel
			if rad.HAlign == HCenter {
				cssNoLeft = strings.Replace(i.CSSLabel, "special-input-left-padding", "", -1)
			}
			one = fmt.Sprintf("<span class='go-quest-label %v'>%v</span>\n", cssNoLeft, one)

			cellAlign := rad.HAlign
			if rad.HAlign == HRight {
				cellAlign = HLeft
			}

			if rad.Cols == 0 {
				one = td(cellAlign, colWidth(1, numCols), one)
				ctrl += one
			} else {

				/* Explanation by example:
				   Cols = 3, Col = [0, 1, 2, 3, 4, 5, 6], Col%Cols = [0, 1, 2, 0, 1, 2, 0]
				   closing/opening should happen  after      Col%Cols == Cols-1 == 2
				           opening should happen *before*    Col == 0
				   closing         should happen  after      Col == len(Cols)-1


					Cols = 2, Col = [0, 1, 2, 3, 4, 5, 6], Col%Cols = [0, 1, 0, 1, 0, 1, 0]

				*/

				if rad.Col == 0 {
					tOpen := tableOpen
					width := fmt.Sprintf(" style='width: %v%%;' >", 100) // table width 100 percent - better would be group.Width
					tOpen = strings.Replace(tOpen, ">", width, -1)
					ctrl += tOpen
				}

				one = td(cellAlign, colWidth(1, rad.Cols), one)
				ctrl += one

				if (rad.Col+0)%rad.Cols == rad.Cols-1 || radIdx == len(i.Radios)-1 {

					ctrl += tableClose
				}
				if (rad.Col+0)%rad.Cols == rad.Cols-1 && radIdx < len(i.Radios)-1 {
					tOpen := tableOpen
					width := fmt.Sprintf(" style='width: %v%%;' >", 100) // table width 100 percent - better would be group.Width
					tOpen = strings.Replace(tOpen, ">", width, -1)
					ctrl += tOpen
				}

			}

		}
		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since golang http.Form.Get() fetches the *first* value.
		//
		// The radio "empty catcher" becomes necessary,
		// if no radio was selected by the participant;
		// but a "must..." validation rule is registered
		if innerType == "radio" || innerType == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='%v' />\n",
				nm, nm, valEmpty)
		}

		// Append suffix
		if i.Suffix.Set() {
			// compare appendSuffix() forcing no wrap for ordinary inputs
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.Suffix.TrSilent(langCode))
		}

		if i.ErrMsg.Set() {
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.ErrMsg.TrSilent(langCode)) // ugly layout  - but radiogroup and checkboxgroup won't have validation errors anyway
		}

		lbl := renderLabelDescription(i, langCode, numCols)
		return lbl + ctrl

	case "dynamic":
		return fmt.Sprintf("<span class='go-quest-label %v'>%v</span>\n", i.CSSLabel, i.Label.Tr(langCode))

	case "composit", "composit-scalar":
		// rendered at group level -  rendered by composit
		return ""

	default:
		return fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, i.Type, implementedTypes)
	}

}

// A group consists of several input controls;
// it contains no response information;
// a group is a layout unit with a configurable number of columns.
type groupT struct {
	// Name  string
	Label trl.S `json:"label,omitempty"`
	Desc  trl.S `json:"description,omitempty"`

	// Vertical space control:
	HeaderBottomVSpacers int `json:"header_bottom_vspacers,omitempty"` // number of half rows below the group header
	BottomVSpacers       int `json:"bottom_vspacers,omitempty"`        // number of rows below the group, initialized to 3

	Vertical bool `json:"vertical,omitempty"` // groups vertically, not horizontally

	OddRowsColoring bool `json:"odd_rows_coloring,omitempty"` // color odd rows
	Width           int  `json:"width,omitempty"`             // default is 100 percent

	// Number of vertical columns;
	// for horizontal *and* (not yet implemented) vertical layouts;
	//
	// Each label (if set) and each input occupy one columns.
	// inputT.ColSpanLabel and inputT.ColSpanControl may set this to more than 1.
	//
	// Cols determines the 'slot' width for these above settings using colWidth(colsElement, colsTotal)
	Cols int `json:"columns,omitempty"`

	Inputs             []*inputT `json:"inputs,omitempty"`
	RandomizationGroup int       `json:"randomization_group,omitempty"` // > 0 => group can be repositioned for randomization
}

// AddInput creates a new input
// and adds this input to the group's inputs
func (gr *groupT) AddInput() *inputT {
	i := &inputT{}
	gr.Inputs = append(gr.Inputs, i)
	ret := gr.Inputs[len(gr.Inputs)-1]
	return ret
}

// TableOpen creates a table markup with various CSS parameters
func (gr *groupT) TableOpen(rows int) string {
	to := tableOpen

	if gr.OddRowsColoring {
		to = strings.Replace(to, "class='main-table' ", "class='main-table bordered'  ", -1) // enable bordering as a whole
	}
	if rows%2 == 1 && gr.OddRowsColoring {
		to = strings.Replace(to, "bordered", "bordered alternate-row-color", -1) // grew background for odd row
	}

	// set width less than 100 percent, for i.e. radios more closely together
	width := fmt.Sprintf(" style='width: %v%%;' >", gr.Width)
	to = strings.Replace(to, ">", width, -1)

	return to
}

// HasComposit - group contains composit element?
func (gr groupT) HasComposit() bool {
	hasComposit := false
	for _, inp := range gr.Inputs {
		if inp.Type == "composit" {
			hasComposit = true
			break
		}
	}
	if hasComposit {
		for _, inp := range gr.Inputs {
			if inp.Type != "composit" && inp.Type != "composit-scalar" {
				log.Panicf("group contains a input type 'composit' - but *other* inputs too")
			}
		}
	}
	return hasComposit
}

func validateComposit(
	pageIdx, grpIdx int, compFuncNameWithParamSet string) (compositFuncT, int, int) {

	splt := strings.Split(compFuncNameWithParamSet, "__")
	if len(splt) != 3 {
		log.Panicf(
			`page %v group %v: 
			composite func name %v 
			must consist of func name '__' param set index '__' sequence idx`,
			pageIdx,
			grpIdx,
			compFuncNameWithParamSet,
		)
	}

	compFuncName := splt[0]
	cF, ok := compositeFuncs[compFuncName]
	if !ok {
		log.Panicf(
			`page %v group %v: 
			composite func name %v does not exist`,
			pageIdx,
			grpIdx,
			compFuncName,
		)
	}

	paramSetIdx, err := strconv.Atoi(splt[1])
	if err != nil {
		log.Panicf(
			`page %v group %v: 
			second part of composite func name %v 
			could not be parsed into int
			%v`,
			pageIdx,
			grpIdx,
			compFuncNameWithParamSet,
			err,
		)
	}
	seqIdx, err := strconv.Atoi(splt[2])
	if err != nil {
		log.Panicf(
			`page %v group %v: 
			third part of composite func name %v 
			could not be parsed into int
			%v`,
			pageIdx,
			grpIdx,
			compFuncNameWithParamSet,
			err,
		)
	}

	return cF, paramSetIdx, seqIdx

}

// HTML renders a group of inputs to HTML
func (gr groupT) HTML(langCode string) string {

	b := &bytes.Buffer{}

	if gr.Width == 0 {
		gr.Width = 100
	}
	b.WriteString(fmt.Sprintf("<div class='go-quest-group' style='width:%v%%;'  cols='%v'>\n", gr.Width, gr.Cols)) // cols is just for debugging
	i := inputT{Type: "textblock"}
	i.HAlignLabel = HLeft
	i.Label = gr.Label
	i.Desc = gr.Desc
	i.CSSLabel = "go-quest-group-header"
	i.ColSpanLabel = gr.Cols
	lbl := renderLabelDescription(i, langCode, gr.Cols)
	// lbl := renderLabelDescription(inputT{Type: "textblock"},	langCode, "", HLeft, gr.Label, gr.Desc, "go-quest-group-header", gr.Cols, gr.Cols)

	b.WriteString(lbl)
	b.WriteString(vspacer)

	b.WriteString("</div>\n")

	for i := 0; i < gr.HeaderBottomVSpacers; i++ {
		b.WriteString(vspacer8)
	}

	// Rendering inputs
	// Adding up columns
	// Find out when a new row starts
	cols := 0 // cols counter
	rows := 0
	b.WriteString(gr.TableOpen(rows))
	for i, inp := range gr.Inputs {

		if inp.Type == "composit-scalar" {
			continue
		}
		if inp.Type == "composit" {
			continue
		}

		fmt.Fprint(b, "<td>")
		fmt.Fprint(b, inp.HTML(langCode, gr.Cols)) // rendering markup
		fmt.Fprint(b, "</td>")

		if gr.Cols > 0 {

			// incrementing label columns
			if inp.Type != "button" { // button has label *inside of it*
				if inp.ColSpanLabel > 1 {
					cols += inp.ColSpanLabel // wider labels
				} else {
					// nothing specified
					if inp.Label != nil || inp.Desc != nil {
						// if a label is set, it occupies one column
						cols++
					}
				}
			}

			// incrementing control columns
			if inp.Type != "textblock" { // textblock has no control part
				if inp.ColSpanControl > 1 {
					cols += inp.ColSpanControl // larger input controls
				} else if len(inp.Radios) > 0 {
					if inp.Radios[0].Cols > 0 {
						cols += inp.Radios[0].Cols
					} else {
						cols += len(inp.Radios) // radiogroups, if no ColSpan is set
					}

				} else {
					// nothing specified => input control occupies one column
					cols++
				}
			}

			// log.Printf("%12v %2v %2v %2v => %3v %2v", inp.Type,
			// 	inp.ColSpanLabel, inp.ColSpanControl, inp.ColSpanLabel+inp.ColSpanControl,
			// 	cols, cols%gr.Cols,
			// ) // dump columns filled so far

			// end of row  - or end of group
			if (cols+0)%gr.Cols == 0 || i == len(gr.Inputs)-1 {
				b.WriteString(tableClose)
			}
			if (cols+0)%gr.Cols == 0 && i < len(gr.Inputs)-1 {
				rows++
				b.WriteString(gr.TableOpen(rows))
			}

		}
	}

	// b.WriteString(tableClose) // this was double of code above

	return b.String()

}

// Type page contains groups with inputs
type pageT struct {
	Section         trl.S `json:"section,omitempty"` // Several pages have a section headline, showing up mobile navigation menu
	Label           trl.S `json:"label,omitempty"`
	Desc            trl.S `json:"description,omitempty"`
	Short           trl.S `json:"short,omitempty"`         // Short version of Label/Description - i.e. for progress bar, replaces Label/Desc in progressbar
	NoNavigation    bool  `json:"no_navigation,omitempty"` // Page will not show up in progress bar
	NavigationalNum int   `json:"navi_num"`                // The number in Navigation order; based on NoNavigation; computed by q.Validate

	Width                 int `json:"width,omitempty"`                  // default is 100 percent
	AestheticCompensation int `json:"aesthetic_compensation,omitempty"` // default is zero percent; if controls do not reach the right border

	Finished time.Time `json:"finished,omitempty"` // truncated to second; *not* a marker for finished entirely - for that we use q.FinishedEntirely

	Groups []*groupT `json:"groups,omitempty"`
}

// AddGroup creates a new group
// and adds this group to the pages's groups
func (p *pageT) AddGroup() *groupT {
	g := &groupT{}
	g.BottomVSpacers = 3
	p.Groups = append(p.Groups, g)
	ret := p.Groups[len(p.Groups)-1]
	return ret
}

// QuestionnaireT contains pages with groups with inputs
type QuestionnaireT struct {
	Survey      surveyT           `json:"survey,omitempty"`
	UserID      string            `json:"user_id,omitempty"`      // participant ID, decimal, but string, i.E. 1011
	Attrs       map[string]string `json:"user_attrs,omitempty"`   // i.e. user country or euro-member - taken from lgn.LoginT
	ClosingTime time.Time         `json:"closing_time,omitempty"` // truncated to second
	RemoteIP    string            `json:"remote_ip,omitempty"`
	UserAgent   string            `json:"user_agent,omitempty"`
	Mobile      int               `json:"mobile,omitempty"` // 0 - no preference, 1 - desktop, 2 - mobile
	MD5         string            `json:"md_5,omitempty"`

	LangCodes []string `json:"lang_codes,omitempty"` // default, order and availability - [en, de, ...] or [de, en, ...]
	LangCode  string   `json:"lang_code,omitempty"`  // current lang code - i.e. 'de' - session key lang_code

	CurrPage  int  `json:"curr_page,omitempty"`
	HasErrors bool `json:"has_errors,omitempty"` // If any response is faulty; set by ValidateReponseData

	Variations int `json:"variations,omitempty"` //  Deterministically shuffle groups based on user id into ... variations.
	MaxGroups  int `json:"max_groups,omitempty"` //  Max number of groups - computed during initialization.

	Pages []*pageT `json:"pages,omitempty"`
}

// We need to register all types who are saved into a session
func init() {
	gob.Register(QuestionnaireT{})
}

// FromSession loads a graph from session;
// second return value contains 'is set'.
func FromSession(w io.Writer, r *http.Request) (*QuestionnaireT, bool, error) {

	sess := sessx.New(w, r)
	key := "questionnaire"

	qstIntf, ok := sess.EffectiveObj(key)
	if !ok {
		log.Printf("key %v for QuestionnaireT{} is not in session", key)
		return nil, false, nil
	}

	q, ok := qstIntf.(QuestionnaireT)
	if !ok {
		return nil, false, fmt.Errorf("key %v for QuestionnaireT{} does not point to qst.QuestionnaireT - but to %T", key, qstIntf)
	}

	return &q, true, nil
}

// BasePath gives the 'root' for loading and saving questionnaire JSON files.
func BasePath() string {
	return path.Join(".", "responses")
}

// FinishedEntirely does not go for the
// page.Finished timestamps, but for
// an explicit input called 'finished'
func (q *QuestionnaireT) FinishedEntirely() (closed bool) {
	for _, p := range q.Pages {
		for _, gr := range p.Groups {
			for _, inp := range gr.Inputs {
				if inp.Name == "finished" {
					if inp.Response == ValSet {
						closed = true
						return
					}
				}
			}
		}
	}
	return
}

// FilePath1 returns the location of the questionnaire file.
// Similar to lgn.LoginT.QuestPath()
func (q *QuestionnaireT) FilePath1() string {
	pth := path.Join(BasePath(), q.Survey.Type, q.Survey.WaveID(), q.UserID)
	if strings.HasSuffix(pth, ".json.json") {
		pth = strings.TrimSuffix(pth, ".json")
	}
	if !strings.HasSuffix(pth, ".json") {
		pth += ".json"
	}
	return pth
}

// AddPage creates a new page
// and adds this page to the questionnaire's pages
func (q *QuestionnaireT) AddPage() *pageT {
	cntr := ctr.Increment()
	p := &pageT{
		Label: trl.S{"en": fmt.Sprintf("PageLabel_%v", cntr), "de": fmt.Sprintf("Seitentitel_%v", cntr)},
		Desc:  trl.S{"en": "", "de": ""},
	}
	q.Pages = append(q.Pages, p)
	ret := q.Pages[len(q.Pages)-1]
	return ret
}

// SetLangCode tries to change the questionnaire langCode if supported by langCodes.
func (q *QuestionnaireT) SetLangCode(newCode string) error {
	if newCode != q.LangCode {
		found := false
		for _, lc := range q.LangCodes {
			if newCode == lc {
				found = true
				break
			}
		}
		if !found {
			err := fmt.Errorf("Language code '%v' is not supported in %v", newCode, q.LangCodes)
			log.Print(err)
			return err
		}
		q.LangCode = newCode
	}
	return nil
}

// CurrentPageHTML is a comfort shortcut to PageHTML
func (q *QuestionnaireT) CurrentPageHTML() (string, error) {
	return q.PageHTML(q.CurrPage)
}

// shufflingGroupsT is a helper for RandomizeOrder()
type shufflingGroupsT struct {
	Orig     int // orig pos
	Shuffled int // shuffled pos - new pos

	Group int // shuffling group

	Start int // shuffling group start idx    - across gaps
	Idx   int // sequence in shuffling group  - across gaps - dense 0,1...6,7

	// seqStart int // shuffling group start idx - continuous chunk
	// seqIdx   int // index in shuffling group  - continuous chunk
}

// String representation for dump
func (sg shufflingGroupsT) String() string {
	return fmt.Sprintf("orig %02v -> shuff %02v - G%v strt%v seq%v", sg.Orig, sg.Shuffled, sg.Group, sg.Start, sg.Idx)
}

// RandomizeOrder creates a shuffled ordering of groups marked by .RandomizationGroup ;
// static groups with RandomizationGroup==0 remain on fixed order position ;
// others get a randomized position
func (q *QuestionnaireT) RandomizeOrder(pageIdx int) []int {

	p := q.Pages[pageIdx]

	// helper - separating groups by their RandomizationGroup value - with positional indexes
	shufflingGroups := map[int][]int{}
	maxSg := 0
	for i := 0; i < len(p.Groups); i++ {
		sg := p.Groups[i].RandomizationGroup
		shufflingGroups[sg] = append(shufflingGroups[sg], i)
		if sg > maxSg {
			maxSg = sg
		}
	}
	if len(shufflingGroups) == 1 && maxSg == 0 {
		return shufflingGroups[0]
	}

	// helper to construct the sequence across gaps within each shuffling group
	shufflingGroupsCntr := map[int]int{}
	for sg := range shufflingGroups {
		shufflingGroupsCntr[sg] = 0
	}

	log.Printf(
		"max sg idx %v \nshufflingGroups %v",
		maxSg,
		util.IndentedDump(shufflingGroups),
	)

	//
	// compute the main array
	sgs := make([]shufflingGroupsT, len(p.Groups))
	for i := 0; i < len(p.Groups); i++ {

		sg := p.Groups[i].RandomizationGroup
		sgs[i].Orig = i
		sgs[i].Group = sg

		sgs[i].Start = shufflingGroups[sg][0]
		sgs[i].Idx = shufflingGroupsCntr[sg]
		shufflingGroupsCntr[sg]++
	}

	//
	// randomize
	for i := 0; i < len(sgs); i++ {
		if sgs[i].Group == 0 {
			sgs[i].Shuffled = sgs[i].Orig
		} else {
			if sgs[i].Idx == 0 {
				sg := sgs[i].Group
				// this must conform with ShufflesToCSV()
				// q.MaxGroups instead of len(shufflingGroups[sg])
				// order = order[0:len(shufflingGroups[sg])]
				sh := shuffler.New(q.UserID, q.Variations, len(shufflingGroups[sg]))
				order := sh.Slice(pageIdx) // cannot add sg to conform to ShufflesToCSV()
				log.Printf("%v - seq %16s in order %16s - iter %v", sg, fmt.Sprint(shufflingGroups[sg]), fmt.Sprint(order), pageIdx+sg)
				for i := 0; i < len(shufflingGroups[sg]); i++ {
					offset := shufflingGroups[sg][i] // i.e. [1, 9]
					i2 := order[i]
					sgs[offset].Shuffled = shufflingGroups[sg][i2]
				}

			}

		}
	}
	for i := 0; i < len(sgs); i++ {
		log.Printf("lp%02v  %v", i, sgs[i])
	}

	// extract the new order - with randomized elements
	shuffled := make([]int, len(p.Groups))
	for i := 0; i < len(p.Groups); i++ {
		shuffled[i] = sgs[i].Shuffled
	}

	log.Printf("=> shuffled %v", shuffled)

	return shuffled

}

// PageHTML generates HTML for a specific page of the questionnaire
func (q *QuestionnaireT) PageHTML(pageIdx int) (string, error) {

	if q.CurrPage > len(q.Pages)-1 || q.CurrPage < 0 {
		s := fmt.Sprintf("You requested page %v out of %v. Page does not exist", pageIdx, len(q.Pages)-1)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	p := q.Pages[pageIdx]

	found := false
	for _, lc := range q.LangCodes {
		if q.LangCode == lc {
			found = true
			break
		}
	}
	if !found {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	b := bytes.Buffer{}

	// set width less than 100 percent, for i.e. radios more closely together

	padding := p.AestheticCompensation
	width := fmt.Sprintf("<div class='page-margins'  style='width: %v%%; margin: 0 auto; padding-left: %v%%' >", p.Width, padding)
	b.WriteString(width)

	if p.Section != nil {
		b.WriteString(fmt.Sprintf("<span class='go-quest-page-section' >%v</span>", p.Section.Tr(q.LangCode)))
		if p.Label.Tr(q.LangCode) != "" {
			b.WriteString("<span class='go-quest-page-desc'> &nbsp; - &nbsp; </span>")
		}
	}

	b.WriteString(fmt.Sprintf("<span class='go-quest-page-header' >%v</span>", p.Label.Tr(q.LangCode)))
	if p.Desc.Tr(q.LangCode) != "" {
		b.WriteString(vspacer)
		b.WriteString(fmt.Sprintf("<p  class='go-quest-page-desc'>%v</p>", p.Desc.Tr(q.LangCode)))
	}
	b.WriteString(vspacer16)

	grpOrder := q.RandomizeOrder(pageIdx)
	compositCntr := -1
	nonCompositCntr := -1
	for loopIdx, grpIdx := range grpOrder {
		if p.Groups[grpIdx].HasComposit() {
			compositCntr++
			compFuncNameWithParamSet := p.Groups[grpIdx].Inputs[0].DynamicFunc
			cF, paramSetIdx, seqIdx := validateComposit(pageIdx, grpIdx, compFuncNameWithParamSet)
			grpHTML, _, err := cF(q, paramSetIdx, seqIdx, q.UserIDInt())
			if err != nil {
				b.WriteString(fmt.Sprintf("composite func error %v \n", err))
			} else {
				b.WriteString(grpHTML + "\n")
			}
		} else {
			grpHTML := p.Groups[grpIdx].HTML(q.LangCode)
			if strings.Contains(grpHTML, "[groupID]") {
				nonCompositCntr++
				grpHTML = strings.Replace(grpHTML, "[groupID]", fmt.Sprintf("%v", nonCompositCntr+1), -1)
			}
			b.WriteString(grpHTML + "\n")
		}

		// vertical distance at the end of groups
		if loopIdx < len(p.Groups)-1 {
			for i2 := 0; i2 < p.Groups[grpIdx].BottomVSpacers; i2++ {
				b.WriteString(vspacer16)
			}
		} else {
			b.WriteString(vspacer16)
		}
	}

	b.WriteString("</div> <!-- width -->")

	ret := b.String()

	// Inject user data into HTML text
	// i.e. [attr-country] => Latvia
	for k, v := range q.Attrs {
		k1 := fmt.Sprintf("[attr-%v]", strings.ToLower(k))
		ret = strings.Replace(ret, k1, v, -1)
	}

	return ret, nil
}

// next page to be shown in navigation
func (q *QuestionnaireT) nextInNavi() (int, bool) {
	// Find next page in navigation
	for i := q.CurrPage + 1; i < len(q.Pages); i++ {
		if !q.Pages[i].NoNavigation {
			return i, true
		}
	}
	// Fallback: Last page in navigation
	for i := len(q.Pages) - 1; i >= 0; i-- {
		if !q.Pages[i].NoNavigation {
			return i, false
		}
	}
	return len(q.Pages) - 1, false
}

// prev page to be shown in navigation
func (q *QuestionnaireT) prevInNavi() (int, bool) {
	// Find prev page in navigation
	for i := q.CurrPage - 1; i >= 0; i-- {
		if !q.Pages[i].NoNavigation {
			return i, true
		}
	}
	// Fallback: First page in navigation
	for i := 0; i < len(q.Pages); i++ {
		if !q.Pages[i].NoNavigation {
			return i, false
		}
	}
	return 0, false
}

// HasPrev if a previous page exists
func (q *QuestionnaireT) HasPrev() bool {
	_, ok := q.prevInNavi()
	return ok
}

// Prev returns index of the previous page
func (q *QuestionnaireT) Prev() int {
	pg, _ := q.prevInNavi()
	return pg
}

// PrevNaviNum returns navigational number of the prev page
func (q *QuestionnaireT) PrevNaviNum() string {
	pg, _ := q.prevInNavi()
	return fmt.Sprintf("%v", q.Pages[pg].NavigationalNum)
}

// HasNext if a next page exists
func (q *QuestionnaireT) HasNext() bool {
	_, ok := q.nextInNavi()
	return ok
}

// Next returns index of the next page
func (q *QuestionnaireT) Next() int {
	pg, _ := q.nextInNavi()
	return pg
}

// NextNaviNum returns navigational number of the next page
func (q *QuestionnaireT) NextNaviNum() string {
	pg, _ := q.nextInNavi()
	return fmt.Sprintf("%v", q.Pages[pg].NavigationalNum)
}

// CurrPageInNavigation - is the current page
// shown in navigation; convenience func for templates
func (q *QuestionnaireT) CurrPageInNavigation() bool {
	return !q.Pages[q.CurrPage].NoNavigation
}

// Compare compares page completion times and input responses.
// Compare stops with the first difference and returns an error.
func (q *QuestionnaireT) Compare(v *QuestionnaireT, lenient bool) (bool, error) {

	if len(q.Pages) != len(v.Pages) {
		return false, fmt.Errorf("Unequal numbers of pages: %v - %v", len(q.Pages), len(v.Pages))
	}

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if len(q.Pages[i1].Groups) != len(q.Pages[i1].Groups) {
			return false, fmt.Errorf("Page %v: Unequal numbers of groups: %v - %v", i1, len(q.Pages[i1].Groups), v.Pages[i1].Groups)
		}
		if i1 < len(q.Pages)-1 { // No completion time comparison for last page
			qf := q.Pages[i1].Finished
			vf := v.Pages[i1].Finished
			if qf.Sub(vf) > 30*time.Second || vf.Sub(qf) > 30*time.Second {
				return false, fmt.Errorf("Page %v: Completion time too distinct: %v - %v", i1, qf, vf)
			}
		}

		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			if len(q.Pages[i1].Groups[i2].Inputs) != len(v.Pages[i1].Groups[i2].Inputs) {
				return false, fmt.Errorf("Page %v: Group %v: Unequal numbers of groups: %v - %v", i1, i2, len(q.Pages[i1].Groups[i2].Inputs), len(v.Pages[i1].Groups[i2].Inputs))
			}
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				if q.Pages[i1].Groups[i2].Inputs[i3].IsLayout() {
					continue
				}
				qr := q.Pages[i1].Groups[i2].Inputs[i3].Response
				vr := v.Pages[i1].Groups[i2].Inputs[i3].Response
				if lenient && (qr == "" && vr == "0" || qr == "0" && vr == "") {
					qr = "0"
					vr = "0"
				}
				if qr != vr {
					return false, fmt.Errorf(
						"Page %v: Group %v: Input %v %v: '%v' != '%v'",
						i1, i2, i3,
						q.Pages[i1].Groups[i2].Inputs[i3].Name,
						q.Pages[i1].Groups[i2].Inputs[i3].Response,
						v.Pages[i1].Groups[i2].Inputs[i3].Response,
					)
				}
			}
		}
	}
	return true, nil
}

// KeysValues returns all pages finish times; keys and values in defined order.
// Empty values are also returned.
// Major purpose is CSV export across several questionnaires.
func (q *QuestionnaireT) KeysValues() (finishes, keys, vals []string) {
	for i1 := 0; i1 < len(q.Pages); i1++ {
		if q.Pages[i1].Finished.IsZero() {
			finishes = append(finishes, "not_saved")
		} else {
			finishes = append(finishes, q.Pages[i1].Finished.Format("02.01.06 15:04:05"))
		}
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				if q.Pages[i1].Groups[i2].Inputs[i3].IsLayout() {
					continue
				}
				keys = append(keys, q.Pages[i1].Groups[i2].Inputs[i3].Name)
				vals = append(vals, q.Pages[i1].Groups[i2].Inputs[i3].Response)
			}
		}
	}
	return
}

// UserIDInt retrieves the userID as int
func (q *QuestionnaireT) UserIDInt() int {
	userID, err := strconv.Atoi(q.UserID)
	if err != nil {
		if q.UserID == "" {
			return 0
		}
		log.Panicf(
			`questionnaire user ID %v
			could not be parsed into integer
			%v`,
			q.UserID, err,
		)
	}
	return userID
}

// ByName retrieves an input element by name
func (q *QuestionnaireT) ByName(n string) *inputT {
	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {
				if q.Pages[i1].Groups[i2].Inputs[i3].IsLayout() {
					continue
				}

				if q.Pages[i1].Groups[i2].Inputs[i3].Name == n {
					return q.Pages[i1].Groups[i2].Inputs[i3]
				}

			}
		}
	}
	return nil
}
