// Package qst implements a four levels deep nested structure
// with input controls, groups, pages and questionnaire;
// contains HTML rendering, page navigation,
// loading/saving from/to JSON file, consistence validation,
// multi-language support.
package qst

import (
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/zew/go-questionnaire/css"
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
func appendSuffix(ctrl string, inp *inputT, langCode string) string {

	if inp.Suffix.Empty() {
		return ctrl
	}

	ctrl = strings.TrimSuffix(ctrl, "\n")
	// We want to prevent line-break between input and suffix with '%' or '€'.
	// inputs must be inline-block, for whitespace nowrap to work.
	// At the same time: suffix-inner enables wrapping for the suffix itself
	sfx := fmt.Sprintf("<span class=' %v  postlabel suffix-inner' >%v</span>\n", inp.CSSLabel, inp.Suffix.TrSilent(langCode))
	ctrl = fmt.Sprintf("<span class='suffix-nowrap' >%v%v</span>\n", ctrl, sfx)

	return ctrl
}

// Special subtype of inputT; used for radiogroup
type radioT struct {
	HAlign horizontalAlignment `json:"hori_align,omitempty"` // label and description left/center/right of input, default left, similar setting for radioT but not for group
	Label  trl.S               `json:"label,omitempty"`
	Val    string              `json:"val,omitempty"`     // Val is allowed to be nil; it then gets initialized to 1...n by Validate(). 0 indicates 'no entry'.
	Col    float32             `json:"column,omitempty"`  // col x of cols
	Cols   float32             `json:"columns,omitempty"` //
	// field 'response' is absent, it is added dynamically;
}

func (inp *inputT) AddRadio() *radioT {
	rad := &radioT{}
	inp.Radios = append(inp.Radios, rad)
	ret := inp.Radios[len(inp.Radios)-1]
	return ret
}

// Input represents a single form input element.
// There is one exception for multiple radios (radiogroup) with the same name but distinct values.
// Multiple checkboxes (checkboxgroup) with same name but distinct values are a dubious instrument.
// See comment to implementedType checkboxgroup.
type inputT struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"` // see implementedTypes

	MaxChars int     `json:"max_chars,omitempty"` // input chars; => SIZE for input, MAXLENGTH for textarea, text; also used for width
	Step     float64 `json:"step,omitempty"`      // for number input:  stepping interval
	Min      float64 `json:"min,omitempty"`       //      ~
	Max      float64 `json:"max,omitempty"`       //      ~

	Label     trl.S  `json:"label,omitempty"`
	Desc      trl.S  `json:"description,omitempty"`
	Suffix    trl.S  `json:"suffix,omitempty"` // only for short units - such as € or % - for longer text use label.Style...Order = 2
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
	ColSpanLabel   float32 `json:"col_span_label,omitempty"`
	ColSpanControl float32 `json:"col_span_control,omitempty"`

	Radios []*radioT  `json:"radios,omitempty"`    // This slice implements the radiogroup - and the senseless checkboxgroup
	DD     *DropdownT `json:"drop_down,omitempty"` // As pointer to prevent JSON cluttering

	Validator string `json:"validator,omitempty"` // i.e. any key from validators, i.e. "must;inRange20"
	ErrMsg    trl.S  `json:"err_msg,omitempty"`

	// ResponseFloat float64  - floats and integers are stored as strings in Response
	// also contains the Value of options and checkboxes
	Response string `json:"response,omitempty"`

	ValueRadio string `json:"value_radio,omitempty"` // for type = radio

	/* compositFunc == 'composit' OR dynFunc == 'dynamic'
	'composit' =>    first arg paramSetIdx, second arg seqIdx */
	DynamicFunc string `json:"dynamic_func,omitempty"`

	Style    *css.StylesResponsive `json:"style,omitempty"` // pointer, to avoid empty JSON blocks
	StyleLbl *css.StylesResponsive `json:"style_label,omitempty"`
	StyleCtl *css.StylesResponsive `json:"style_control,omitempty"`
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

func renderLabelDescription(i inputT, langCode string, numCols float32) string {
	return renderLabelDescription2(i, langCode, i.Name, i.HAlignLabel,
		i.Label, i.Desc, i.CSSLabel, i.ColSpanLabel, numCols)
}

// renderLabelDescription wraps lbl+desc into a <span> of class 'go-quest-cell' or td-cell.
// A percent width is dynamically computed from colsLabel / numCols.
// Argument numCols is the total number of cols per row.
// It is used to compute the precise width in percent
func renderLabelDescription2(i inputT, langCode string, name string, hAlign horizontalAlignment,
	lbl, desc trl.S, css string, colsLabel float32, numCols float32) string {
	ret := ""
	if lbl == nil && desc == nil {
		return ret
	}
	e1 := lbl.Tr(langCode)
	if lbl == nil {
		e1 = "" // suppress "Translation map not initialized." here
	}
	e2 := desc.Tr(langCode)
	if desc == nil {
		e2 = "" // suppress "Translation map not initialized." here
	}

	// pure text or layout
	ret = fmt.Sprintf(
		"<span class='%v'><b>%v</b> %v </span>\n",
		css, e1, e2,
	)

	if name != "" && !i.IsLayout() {
		ret = fmt.Sprintf(
			"<label for='%v' class='%v' ><b>%v</b> %v </label>\n",
			name, css, e1, e2,
		)
	}

	ret = td(hAlign, colWidth(colsLabel, numCols), ret)
	return ret
}

// IsLayout returns whether the input type is merely ornamental
// and has no return values
func (inp inputT) IsLayout() bool {
	if inp.Type == "textblock" {
		return true
	}
	if inp.Type == "dyn-textblock" {
		return true
	}
	if inp.Type == "dyn-composite" { // inputs are in "dyn-composite-scalar"
		return true
	}
	if inp.Type == "button" { // we dont care
		return true
	}
	return false
}

// IsControlOnly types having no ctrl part but only a label
func (inp inputT) IsControlOnly() bool {
	if inp.Type == "button" {
		return true
	}
	return false
}

// IsLabelOnly types having no ctrl part but only a label
func (inp inputT) IsLabelOnly() bool {
	if inp.Type == "textblock" {
		return true
	}
	if inp.Type == "dyn-textblock" {
		return true
	}
	return false
}

// IsHidden types having neither visible ctrl nor label part
func (inp inputT) IsHidden() bool {
	if inp.Type == "hidden" {
		return true
	}
	if inp.Type == "dyn-composite-scalar" {
		return true
	}
	return false
}

// IsReserved returns whether the input name is reserved by the survey engine
func (inp inputT) IsReserved() bool {
	if inp.Name == "page" {
		return true
	}
	if inp.Name == "lang_code" {
		return true
	}
	return false
}

// Rendering one input to HTML
// func (inp inputT) HTML(langCode string, namePrefix string) string {
func (inp inputT) HTML(langCode string, numCols float32) string {
	return "Version < 2 no longer implemented"
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
	BottomVSpacers       int `json:"bottom_vspacers,omitempty"`        // number of rows below the group, addGroup() initializes to 3

	Vertical bool `json:"vertical,omitempty"` // groups vertically, not horizontally

	OddRowsColoring bool `json:"odd_rows_coloring,omitempty"` // color odd rows

	// Number of vertical columns;
	// for horizontal *and* (not yet implemented) vertical layouts;
	//
	// Each label (if set) and each input occupy one columns.
	// inputT.ColSpanLabel and inputT.ColSpanControl may set this to more than 1.
	//
	// Cols determines the 'slot' width for these above settings using colWidth(colsElement, colsTotal)
	Cols float32 `json:"columns,omitempty"`

	Inputs             []*inputT `json:"inputs,omitempty"`
	RandomizationGroup int       `json:"randomization_group,omitempty"` // > 0 => group can be repositioned for randomization

	Style *css.StylesResponsive `json:"style,omitempty"` // pointer, to avoid empty JSON blocks
}

// AddInput creates a new input
// and adds this input to the group's inputs
func (gr *groupT) AddInput() *inputT {
	inp := &inputT{}
	gr.Inputs = append(gr.Inputs, inp)
	ret := gr.Inputs[len(gr.Inputs)-1]
	return ret
}

// addInputArg adds arg to group
func (gr *groupT) addInputArg(inp *inputT) {
	gr.Inputs = append(gr.Inputs, inp)
}

// InputEmpty creates a new input;
// an empty input *is* rendered a empty cell
func InputEmpty() *inputT {
	inp := &inputT{}
	inp.Type = "textblock"
	inp.Label = trl.S{
		"en": " &nbsp; ",
		"de": " &nbsp; ",
	}
	return inp
}

// addInputEmpty creates a new input
// and adds this input to the group's inputs
func (gr *groupT) addInputEmpty() *inputT {
	inp := InputEmpty()
	gr.Inputs = append(gr.Inputs, inp)
	ret := gr.Inputs[len(gr.Inputs)-1]
	return ret
}

// HasComposit - group contains composit element?
func (gr groupT) HasComposit() bool {
	hasComposit := false
	for _, inp := range gr.Inputs {
		if inp.Type == "dyn-composite" {
			hasComposit = true
			break
		}
	}
	if hasComposit {
		for _, inp := range gr.Inputs {
			if inp.Type != "dyn-composite" && inp.Type != "dyn-composite-scalar" {
				log.Panicf("group contains a input type 'composit' - but *other* inputs too")
			}
		}
	}
	return hasComposit
}

// returns the func, the sequence idx, the param set idx
func validateComposite(
	pageIdx, grpIdx int, compFuncNameWithParamSet string) (CompositFuncT, int, int) {

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
	cF, ok := CompositeFuncs[compFuncName]
	if !ok {
		log.Panicf(
			`page %v group %v: 
			composite func name %v does not exist`,
			pageIdx,
			grpIdx,
			compFuncName,
		)
	}

	seqIdx, err := strconv.Atoi(splt[1])
	if err != nil {
		log.Panicf(
			`page %v group %v: 
			third part of composite func name %v 
			could not be parsed into int
			%v`,
			seqIdx,
			grpIdx,
			compFuncNameWithParamSet,
			err,
		)
	}

	paramSetIdx, err := strconv.Atoi(splt[2])
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

	return cF, seqIdx, paramSetIdx

}

// GroupHTMLTableBased renders a group of inputs to GroupHTMLTableBased
func (q QuestionnaireT) GroupHTMLTableBased(pageIdx, grpIdx int) string {
	return "GroupHTMLTableBased no longer implemented"
}

// Type page contains groups with inputs
type pageT struct {
	Section         trl.S `json:"section,omitempty"`       // extra strong before label in content - summary headline for multiple pages
	Label           trl.S `json:"label,omitempty"`         // headline, set to "" to prevent rendering
	Desc            trl.S `json:"description,omitempty"`   // abstract
	Short           trl.S `json:"short,omitempty"`         // sort version of section/label/description - in progress bar and navigation menu
	NoNavigation    bool  `json:"no_navigation,omitempty"` // Page will not show up in progress bar
	NavigationalNum int   `json:"navi_num"`                // The number in Navigation order; based on NoNavigation; computed by q.Validate

	Width int `json:"width,omitempty"` // default is 100 percent
	// AestheticCompensation int `json:"aesthetic_compensation,omitempty"` // default is zero percent; if controls do not reach the right border

	Finished time.Time `json:"finished,omitempty"` // truncated to second; *not* a marker for finished entirely - for that we use q.FinishedEntirely

	Groups []*groupT `json:"groups,omitempty"`

	ValidationFuncName string `json:"validation_func_name,omitempty"` // javascript validation func name
	ValidationFunc     string `json:"validation_func,omitempty"`      // javascript validation func implementation
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

	Version int `json:"version,omitempty"` // 0 - rendering as HTML table   -   1 - rendering as CSS Grid
}

// registering all types, being saved into a session
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

	w := &strings.Builder{}

	// i.e. smaller - for i.e. radios more closely together
	// todo: change to Style
	width := fmt.Sprintf("<div class='page-margins' style='margin: 0 auto; margin-top: 0.6rem; width: %v%%'  >\n", p.Width)
	fmt.Fprint(w, width)

	hasHeader := false

	if p.Section != nil {
		fmt.Fprintf(w, "<span class='go-quest-page-section' >%v</span>", p.Section.Tr(q.LangCode))
		if p.Label.Tr(q.LangCode) != "" {
			fmt.Fprint(w, "<span class='go-quest-page-desc'> &nbsp; - &nbsp; </span>")
		}
		hasHeader = true
	}
	if p.Label.Tr(q.LangCode) != "" {
		fmt.Fprintf(w, "<span class='go-quest-page-header' >%v</span>", p.Label.Tr(q.LangCode))
		hasHeader = true
	}
	if p.Desc.Tr(q.LangCode) != "" {
		fmt.Fprint(w, vspacer0)
		fmt.Fprintf(w, "<p  class='go-quest-page-desc'>%v</p>", p.Desc.Tr(q.LangCode))
		hasHeader = true
	}

	if hasHeader {
		fmt.Fprint(w, vspacer16)
	}

	grpOrder := q.RandomizeOrder(pageIdx)
	compositCntr := -1
	nonCompositCntr := -1
	for loopIdx, grpIdx := range grpOrder {
		if p.Groups[grpIdx].HasComposit() {
			compositCntr++
			compFuncNameWithParamSet := p.Groups[grpIdx].Inputs[0].DynamicFunc
			cF, seqIdx, paramSetIdx := validateComposite(pageIdx, grpIdx, compFuncNameWithParamSet)
			grpHTML, _, err := cF(q, seqIdx, paramSetIdx)
			if err != nil {
				fmt.Fprintf(w, "composite func error %v \n", err)
			} else {
				// grpHTML also contains HTML and CSS stuff - which could be hyphenized too
				grpHTML = trl.HyphenizeText(grpHTML)
				fmt.Fprint(w, grpHTML+"\n")
			}
		} else {
			grpHTML := ""
			if q.Version > 1 {
				grpHTML = q.GroupHTMLGridBased(pageIdx, grpIdx)
			} else {
				grpHTML = q.GroupHTMLTableBased(pageIdx, grpIdx)
			}

			if strings.Contains(grpHTML, "[groupID]") {
				nonCompositCntr++
				grpHTML = strings.Replace(grpHTML, "[groupID]", fmt.Sprintf("%v", nonCompositCntr+1), -1)
			}
			fmt.Fprint(w, grpHTML+"\n")
		}

		// vertical distance at the end of groups
		if loopIdx < len(p.Groups)-1 {
			for i2 := 0; i2 < p.Groups[grpIdx].BottomVSpacers; i2++ {
				fmt.Fprint(w, vspacer16)
			}
		} else {
			fmt.Fprint(w, vspacer16)
		}
	}

	fmt.Fprint(w, "</div> <!-- width -->")

	ret := w.String()

	// inject user data into HTML text
	// i.e. [attr-country] => Latvia
	for k, v := range q.Attrs {
		k1 := fmt.Sprintf("[attr-%v]", strings.ToLower(k))
		ret = strings.Replace(ret, k1, v, -1)
	}

	if strings.Contains(ret, "(MISSING)") {
		log.Printf("PageHTML() returns (MISSING). Reason:  Printf(w, fmt.Sprintf('xxx ... %% ... '))  -  remove suffix 'f' from outer call.")
		// ret = strings.ReplaceAll(ret, "(MISSING)", "")
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
