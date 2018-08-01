// Package qst implements a four levels deep nested structure
// with input controls, groups, pages and questionaire;
// contains HTML rendering, page navigation,
// loading/saving from/to JSON file, consistence validation,
// multi-language support.
package qst

import (
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	"github.com/zew/go-questionaire/trl"

	"github.com/zew/go-questionaire/ctr"
)

// Special subtype of inputT; used for radiogroup
type radioT struct {
	HAlign horizontalAlignment `json:"hori_align,omitempty"` // label and description left/center/right of input, default left, similar setting for radioT but not for group
	Label  trl.S               `json:"label,omitempty"`
	Val    string              `json:"val,omitempty"` // Val is allowed to be nil; it then gets initialized to 1...n by Validate(). 0 indicates 'no entry'.
	// Notice the absence of Response;
}

func (i *inputT) AddRadio() *radioT {
	rad := &radioT{}
	i.Radios = append(i.Radios, rad)
	ret := i.Radios[len(i.Radios)-1]
	return ret
}

// Input represents a single form input element.
// There is one exception for multiple radios (radiogroup) with the same name but distinct values.
// Multiple checkboxes (checkboxgroup) with same name but distinct values are a dubious instrument. See comment to implementedType checkboxgroup.
type inputT struct {
	Name      string `json:"name,omitempty"`
	Type      string `json:"type,omitempty"`
	InputMode string `json:"input_mode,omitempty"` // set to i.e. " inputmode='numeric' "
	MaxChars  int    `json:"max_chars,omitempty"`  // Number of input chars, also used to computer width

	HAlignLabel   horizontalAlignment `json:"horizontal_align_label,omitempty"`   // description left/center/right of input, default left, similar setting for radioT but not for group
	HAlignControl horizontalAlignment `json:"horizontal_align_control,omitempty"` // label       left/center/right of input, default left, similar setting for radioT but not for group
	CSSLabel      string              `json:"css_label,omitempty"`
	CSSControl    string              `json:"css_control,omitempty"`
	Label         trl.S               `json:"label,omitempty"`
	Desc          trl.S               `json:"description,omitempty"`
	Suffix        trl.S               `json:"suffix,omitempty"`
	AccessKey     string              `json:"accesskey,omitempty"`

	// How many column slots of the overall layout should the control occupy?
	// The number adds up against group.Cols - determining newlines.
	// The number is used to compute the relative width (percentage).
	// If zero, a column width of one is assumend.
	ColSpanLabel   int `json:"col_span_label,omitempty"`
	ColSpanControl int `json:"col_span_control,omitempty"`

	Radios []*radioT `json:"radios,omitempty"` // This slice implements the radiogroup - and the senseless checkboxgroup

	Validator string `json:"validator,omitempty"` // i.e. inRange20 - any string from validators
	ErrMsg    trl.S  `json:"err_msg,omitempty"`

	Response      string  `json:"response,omitempty"`       // but also Value
	ResponseFloat float64 `json:"response_float,omitempty"` // also for integers

	DynamicFunc string `json:"dynamic_func,omitempty"` // Refers to dynFuncs, for type == 'dynamic'
}

// Returns an input filled in with globally enumerated label, decription etc.
func newInputUnused() inputT {
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
	case "dynamic":
		return fmt.Sprintf("<span class='go-quest-label %v'>%v</span>\n", i.CSSLabel, i.Label.Tr(langCode))

	case "button":
		lbl := fmt.Sprintf("<button type='submit' name='%v' value='%v' class='%v' accesskey='%v'><b>%v</b> %v</button>\n",
			i.Name, i.Response, i.CSSControl, i.AccessKey,
			i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode),
		)
		lbl = td(i.HAlignControl, colWidth(i.ColSpanControl, numCols), lbl)
		return lbl

	case "textblock":
		lbl := renderLabelDescription(i, langCode, numCols)
		return lbl

	case "radiogroup", "checkboxgroup":
		ctrl := ""
		innerType := "radio"
		if i.Type == "checkboxgroup" {
			innerType = "checkbox"
		}
		for _, rad := range i.Radios {
			one := ""
			checked := ""
			if i.Response == rad.Val {
				checked = "checked=\"checked\""
			}
			// one += fmt.Sprintf("Val %v", val)

			if rad.Label != nil && rad.HAlign == HLeft {
				one += fmt.Sprintf("<span class='go-quest-label vert-correct %v' >%v</span>\n", i.CSSLabel, rad.Label.Tr(langCode))
			}
			if rad.Label != nil && rad.HAlign == HCenter {
				one += fmt.Sprintf("<span class='go-quest-label vert-correct %v'>%v</span>\n", i.CSSLabel, rad.Label.Tr(langCode))
				one += vspacer
			}

			one += fmt.Sprintf("<input type='%v' name='%v' id='%v' title='%v %v' class='%v' value='%v' %v />\n",
				innerType, nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode), i.CSSControl,
				rad.Val, checked,
			)

			if rad.Label != nil && rad.HAlign == HRight {
				one += fmt.Sprintf("<span class='go-quest-label vert-correct %v'>%v</span>\n", i.CSSLabel, rad.Label.Tr(langCode))
			}
			one = td(rad.HAlign, colWidth(1, numCols), one)
			ctrl += one
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

		if i.Suffix.Set() {
			// compare suffix with no break for ordinary inputs
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.Suffix.TrSilent(langCode))
		}
		if i.ErrMsg.Set() {
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.ErrMsg.TrSilent(langCode)) // ugly layout  - but radiogroup and checkboxgroup won't have validation errors anyway
		}

		lbl := renderLabelDescription(i, langCode, numCols)
		return lbl + ctrl

	case "text", "textarea", "checkbox":
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
		if i.Type == "checkbox" || i.Type == "radio" {
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
		} else {
			ctrl += fmt.Sprintf("<input type='%v' %v name='%v' id='%v' title='%v %v' class='%v' style='%v' %v %v  value='%v' />\n",
				i.Type, i.InputMode,
				nm, nm, i.Label.TrSilent(langCode), i.Desc.TrSilent(langCode), i.CSSControl, width, maxChars, checked, val)
		}

		// The checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if i.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

		// Append suffix
		if i.Suffix.Set() {
			ctrl = strings.TrimSuffix(ctrl, "\n")
			sfx := fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.Suffix.TrSilent(langCode))
			// We want to prevent line-break of the '%' or '€' character.
			// inputs must be inline-block, for whitespace nowrap to work
			ctrl = fmt.Sprintf("<span style='white-space: nowrap;' >%v%v</span>\n", ctrl, sfx)
		}
		// Append error message
		if i.ErrMsg.Set() {
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.ErrMsg.TrSilent(langCode))
		}

		ctrl = td(i.HAlignControl, colWidth(i.ColSpanControl, numCols), ctrl)

		lbl := renderLabelDescription(i, langCode, numCols)
		return lbl + ctrl

	default:
		return fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, i.Type, implementedTypes)
	}

}

// A group consists of several input controls.
// It contains no response information.
// It can bundle checkbox or text inputs with *distinct* names.
// Whereas: radiogroup and checkboxgroup have the *same* name and a single response.
// A group is a layout unit with a configurable number of columns.
type groupT struct {
	// Name  string
	Label               trl.S `json:"label,omitempty"`
	Desc                trl.S `json:"description,omitempty"`
	GroupHeaderVSpacers int   `json:"bottom_half_rows,omitempty"` // number of half rows below the group header

	Vertical bool `json:"vertical,omitempty"` // groups vertically, not horizontally, not yet implemented

	OddRowsColoring bool `json:"odd_rows_coloring"` // color odd rows
	Width           int  `json:"width"`             // default is 100 percent

	// Number of vertical columns;
	// for horizontal *and* (not yet implemented) vertical layouts;
	//
	// Each label (if set) and each input occupy one columns.
	// inputT.ColSpanLabel and inputT.ColSpanControl may set this to more than 1.
	//
	// Cols determines the 'slot' width for these above settings using colWidth(colsElement, colsTotal)
	Cols int `json:"columns,omitempty"`

	Inputs []*inputT `json:"inputs,omitempty"`
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

// HTML renders a group of inputs to HTML
func (gr groupT) HTML(langCode string) string {

	b := bytes.Buffer{}

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

	for i := 0; i < gr.GroupHeaderVSpacers; i++ {
		b.WriteString(vspacer8)
	}

	// Rendering inputs
	// Adding up columns
	// Find out when a new row starts
	cols := 0 // cols counter
	rows := 0
	b.WriteString(gr.TableOpen(rows))
	for i, inp := range gr.Inputs {

		b.WriteString(inp.HTML(langCode, gr.Cols)) // rendering markup

		if gr.Cols > 0 {

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

			if inp.Type != "textblock" { // textblock has no control part
				if inp.ColSpanControl > 1 {
					cols += inp.ColSpanControl // larger input controls
				} else if len(inp.Radios) > 0 {
					cols += len(inp.Radios) // radiogroups, if no ColSpan is set
				} else {
					// nothing specified => input control occupies one column
					cols++
				}
			}

			// log.Printf("%12v %2v %2v", inp.Type, cols, cols%gr.Cols) // so far

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
	Section         trl.S `json:"section,omitempty"` // several pages have a section headline
	Label           trl.S `json:"label,omitempty"`
	Desc            trl.S `json:"description,omitempty"`
	Short           trl.S `json:"short,omitempty"`         // Short version of Label/Description - i.e. for progress bar
	NoNavigation    bool  `json:"no_navigation,omitempty"` // page will not show up in progress bar
	NavigationalNum int   `json:"navi_num"`                // the number in Navigation order; computed by q.Validate

	Width                 int `json:"width,omitempty"`                  // default is 100 percent
	AestheticCompensation int `json:"aesthetic_compensation,omitempty"` // default is zero percent; if controls do not reach the right border

	Finished time.Time `json:"finished,omitempty"` // truncated to second

	Groups []*groupT `json:"groups,omitempty"`
}

// AddGroup creates a new group
// and adds this group to the pages's groups
func (p *pageT) AddGroup() *groupT {
	g := &groupT{}
	p.Groups = append(p.Groups, g)
	ret := p.Groups[len(p.Groups)-1]
	return ret
}

// QuestionaireT contains pages with groups with inputs
type QuestionaireT struct {
	Survey      surveyT   `json:"survey,omitempty"`
	UserID      string    `json:"user_id,omitempty"`      // participant ID
	ClosingTime time.Time `json:"closing_time,omitempty"` // truncated to second
	RemoteIP    string    `json:"remote_ip,omitempty"`
	UserAgent   string    `json:"user_agent,omitempty"`
	Mobile      int       `json:"mobile,omitempty"` // 0 - no preference, 1 - desktop, 2 - mobile
	MD5         string    `json:"md_5,omitempty"`

	// LangCode and LangCodes are imposed from cfg.LangCodes via session."lang_code"
	LangCodes map[string]string `json:"lang_codes,omitempty"` // all possible lang codes - i.e. en, de
	LangCode  string            `json:"lang_code,omitempty"`  // default lang code - and current lang code - i.e. de

	CurrPage  int  `json:"curr_page,omitempty"`
	HasErrors bool `json:"has_errors,omitempty"` // If any response is faulty; set by ValidateReponseData

	Pages []*pageT `json:"pages,omitempty"`
}

// BasePath gives the 'root' for loading and saving questionaire JSON files.
func BasePath() string {
	return filepath.Join(".", "responses")
}

// FilePath1 returns the file system saving location of the questionaire.
// The waveID/userID fragment can optionally be submitted by an argument.
func (q *QuestionaireT) FilePath1(surveyAndWaveIDAndUserID ...string) string {
	pth := ""
	if len(surveyAndWaveIDAndUserID) > 0 {
		pth = filepath.Join(BasePath(), surveyAndWaveIDAndUserID[0])
	} else {
		pth = filepath.Join(BasePath(), q.Survey.Type, q.Survey.WaveID(), q.UserID)
	}

	if strings.HasSuffix(pth, ".json.json") {
		pth = strings.TrimSuffix(pth, ".json")
	}
	if !strings.HasSuffix(pth, ".json") {
		pth += ".json"
	}

	return pth
}

// AddPage creates a new page
// and adds this page to the questionaire's pages
func (q *QuestionaireT) AddPage() *pageT {
	cntr := ctr.Increment()
	p := &pageT{
		Label: trl.S{"en": fmt.Sprintf("PageLabel_%v", cntr), "de": fmt.Sprintf("Seitentitel_%v", cntr)},
		Desc:  trl.S{"en": "", "de": ""},
	}
	q.Pages = append(q.Pages, p)
	ret := q.Pages[len(q.Pages)-1]
	return ret
}

// SetLangCode tries to change the questionaire langCode if supported by langCodes.
func (q *QuestionaireT) SetLangCode(newCode string) error {
	if newCode != q.LangCode {
		oldCode := q.LangCode
		q.LangCode = newCode
		err := q.Validate()
		if err != nil {
			q.LangCode = oldCode
			return err
		}
		// sess.PutString("lang_code", q.LangCode)
	}
	return nil
}

// CurrentPageHTML is a comfort shortcut to PageHTML
func (q *QuestionaireT) CurrentPageHTML() (string, error) {
	return q.PageHTML(q.CurrPage)
}

// PageHTML generates HTML for a specific page of the questionaire
func (q *QuestionaireT) PageHTML(idx int) (string, error) {

	if q.CurrPage > len(q.Pages)-1 || q.CurrPage < 0 {
		s := fmt.Sprintf("You requested page %v out of %v. Page does not exist", idx, len(q.Pages)-1)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	p := q.Pages[idx]

	if _, ok := q.LangCodes[q.LangCode]; !ok || q.LangCode == "" {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	b := bytes.Buffer{}

	// set width less than 100 percent, for i.e. radios more closely together

	padding := p.AestheticCompensation
	width := fmt.Sprintf("<div style='width: %v%%; margin: 0 auto; padding-left: %v%%' >", p.Width, padding)
	b.WriteString(width)

	if p.Section != nil {
		b.WriteString(fmt.Sprintf("<span class='go-quest-page-section' >%v</span>", p.Section.Tr(q.LangCode)))
		if p.Label.Tr(q.LangCode) != "" {
			b.WriteString("<span class='go-quest-page-desc'> &nbsp; - &nbsp; </span>")
		}
	}

	b.WriteString(fmt.Sprintf("<span class='go-quest-page-header' >%v</span>", p.Label.Tr(q.LangCode)))
	b.WriteString(vspacer)
	b.WriteString(fmt.Sprintf("<p  class='go-quest-page-desc'>%v</p>", p.Desc.Tr(q.LangCode)))
	b.WriteString(vspacer16)

	for i := 0; i < len(p.Groups); i++ {
		b.WriteString(p.Groups[i].HTML(q.LangCode) + "\n")
		b.WriteString(vspacer16)
		if i < len(p.Groups)-1 { // no vertical distance at the end of groups
			b.WriteString(vspacer16)
			b.WriteString(vspacer16)
		}
	}

	b.WriteString("</div> <!-- width -->")

	return b.String(), nil
}

// next page to be shown in navigation
func (q *QuestionaireT) nextInNavi() (int, bool) {
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
func (q *QuestionaireT) prevInNavi() (int, bool) {
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
func (q *QuestionaireT) HasPrev() bool {
	_, ok := q.prevInNavi()
	return ok
}

// Prev returns index of the previous page
func (q *QuestionaireT) Prev() int {
	pg, _ := q.prevInNavi()
	return pg
}

// PrevNaviNum returns navigational number of the prev page
func (q *QuestionaireT) PrevNaviNum() string {
	pg, _ := q.prevInNavi()
	return fmt.Sprintf("%v", q.Pages[pg].NavigationalNum)
}

// HasNext if a next page exists
func (q *QuestionaireT) HasNext() bool {
	_, ok := q.nextInNavi()
	return ok
}

// Next returns index of the next page
func (q *QuestionaireT) Next() int {
	pg, _ := q.nextInNavi()
	return pg
}

// NextNaviNum returns navigational number of the next page
func (q *QuestionaireT) NextNaviNum() string {
	pg, _ := q.nextInNavi()
	return fmt.Sprintf("%v", q.Pages[pg].NavigationalNum)
}

// CurrPageInNavigation - is the current page
// shown in navigation; convenience func for templates
func (q *QuestionaireT) CurrPageInNavigation() bool {
	return !q.Pages[q.CurrPage].NoNavigation
}

// Compare compares page completion times and input responses.
// Compare stops with the first difference and returns an error.
func (q *QuestionaireT) Compare(v *QuestionaireT) (bool, error) {

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
			if qf.Sub(vf) > 20*time.Second || vf.Sub(qf) > 20*time.Second {
				return false, fmt.Errorf("Page %v: Comletion time too distinct: %v - %v", i1, vf, qf)
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
				if q.Pages[i1].Groups[i2].Inputs[i3].Response != v.Pages[i1].Groups[i2].Inputs[i3].Response {
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
