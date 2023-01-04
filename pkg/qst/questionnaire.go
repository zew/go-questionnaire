// Package qst implements a four levels deep nested structure
// with input controls, groups, pages and questionnaire;
// contains HTML rendering, page navigation,
// loading/saving from/to JSON file, consistence validation,
// multi-language support.
package qst

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/pbberlin/dbg"
	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/cloudio"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/lgn/shuffler"
	"github.com/zew/go-questionnaire/pkg/sessx"
	"github.com/zew/go-questionnaire/pkg/trl"

	"github.com/zew/go-questionnaire/pkg/ctr"
)

// No line wrapping between element 1 and 2
//
//	But line wrapping *inside* each of them.
//
//	el1 and el2 must be inline-block, for whitespace nowrap to work.
func nobreakGlue(el1, glue, el2 string) string {

	if el1 == "" || el2 == "" {
		return el1 + el2
	}

	// reduction := 82 // 	el2 is overflowing :(
	reduction := 95 // 2020-05 - relaxed

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

// Input represents a single form input element.
// There is one exception for multiple radios (radiogroup) with the same name but distinct values.
// Multiple checkboxes (checkboxgroup) with same name but distinct values are a dubious instrument.
// See comment to implementedType checkboxgroup.
type inputT struct {
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"` // see implementedTypes

	MaxChars    int     `json:"max_chars,omitempty"`  // input chars; => SIZE for input, MAXLENGTH for textarea, text; also used for width
	Step        float64 `json:"step,omitempty"`       // for number input:  stepping interval, i.e. 2 or 0.1
	Min         float64 `json:"min,omitempty"`        //      ~
	Max         float64 `json:"max,omitempty"`        //      ~
	Disabled    bool    `json:"disabled,omitempty"`   // simple the HTML input attribute
	OnInvalid   trl.S   `json:"on_invalid,omitempty"` // message for javascript error messages on HTML5 invalid state - compare ErrMsg
	Placeholder trl.S   `json:"placeholder,omitempty"`

	Label     trl.S  `json:"label,omitempty"`
	Desc      trl.S  `json:"description,omitempty"`
	Suffix    trl.S  `json:"suffix,omitempty"` // only for short units - such as € or % - for longer text use label.Style...Order = 2
	Tooltip   trl.S  `json:"tooltip,omitempty"`
	AccessKey string `json:"accesskey,omitempty"`

	/* Colspan determines, how many column slots of the group column layout
	the input occupies.
	Default value is assumed to be 1.
	Increase it manually in the generator function.
	ColSpanLabel and ColSpanControl do *not* influence Colspan.
	*/
	ColSpan float32 `json:"col_span,omitempty"`

	// ColSpanLabel/-Control work only as proportion of Colspan
	ColSpanLabel   float32 `json:"col_span_label,omitempty"`
	ColSpanControl float32 `json:"col_span_control,omitempty"`

	DD *DropdownT `json:"drop_down,omitempty"` // As pointer to prevent JSON cluttering

	Validator string `json:"validator,omitempty"` // i.e. any key from map of validators, i.e. "must;inRange20"
	// key to coreTranslations, content comes from Validator(Response), compare OnInvalid
	// for radio inputs, see ErrorProxy
	ErrMsg string `json:"err_msg,omitempty"`

	// Response - input.value - numbers are stored as strings too - also contains the value of options and checkboxes
	Response   string `json:"response,omitempty"`
	ValueRadio string `json:"value_radio,omitempty"` // for type = radio

	// depending if
	// 		inp.Type == "dyn-composite"
	// 		inp.Type == "dyn-textblock"
	// then
	// 		=> lookup in CompositeFuncs in funcs-composite.go
	// 		=> lookup in       dynFuncs in funcs-dynamic.go
	//       if "dyn-composite" =>   first arg paramSetIdx, second arg seqIdx, for example TimePreferenceSelfComprehensionCheck__0__0
	// 		 if "dyn-textblock" =>   param in inp.DynamicFuncParamset
	DynamicFunc         string `json:"dynamic_func,omitempty"`
	DynamicFuncParamset string `json:"dynamic_func_paramset,omitempty"` // for "dyn-textblock" - name of parameter set

	Style    *css.StylesResponsive `json:"style,omitempty"` // pointer, to avoid empty JSON blocks
	StyleLbl *css.StylesResponsive `json:"style_label,omitempty"`
	StyleCtl *css.StylesResponsive `json:"style_control,omitempty"`

	JSBlockStrings map[string]string `json:"js_block_strings,omitempty"`
	JSBlockTrls    map[string]trl.S  `json:"js_block_translations,omitempty"`
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

// IsLayout returns whether the input type is merely ornamental
// and has no return values
func (inp inputT) IsLayout() bool {
	if inp.Type == "textblock" {
		return true
	}
	if inp.Type == "button" { // we dont care
		return true
	}
	if inp.Type == "label-as-input" {
		return true
	}
	if inp.Type == "dyn-textblock" {
		return true
	}
	if inp.Type == "dyn-composite" { // inputs are in "dyn-composite-scalar"
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

// IsHidden types having neither visible ctrl nor label part;
// thus no grid-cells are rendered
func (inp inputT) IsHidden() bool {
	if inp.Type == "hidden" {
		return true
	}
	if inp.Type == "javascript-block" {
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

// Signature returns the first token from DynamicFuncParamset;
// usually this is [1,2,3]
func (inp inputT) Signature() string {
	if inp.DynamicFuncParamset != "" {
		parts := strings.Split(inp.DynamicFuncParamset, "--")
		return fmt.Sprintf("signature-%v", parts[0])
	}
	return fmt.Sprintf("signature-%v-%v-%v", inp.Min, inp.Max, inp.Step)
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
	ID             string `json:"-"`                         // is only a helper, but must be exported in order to be stored in session
	BottomVSpacers int    `json:"bottom_vspacers,omitempty"` // number of rows below the group, addGroup() initializes to 3

	// Number of vertical columns;
	// for horizontal *and* (not yet implemented) vertical layouts;
	//
	// Each label (if set) and each input occupy columns according to
	// inputT.ColSpanLabel and inputT.ColSpanControl.
	Cols float32 `json:"columns,omitempty"`

	Inputs []*inputT `json:"inputs,omitempty"`

	// reordering / re-shuffled according to questionnaireT.ShufflingVariations
	// > 0 => group belongs to a set of groups,
	// there can be multiple groups on one page
	RandomizationGroup int `json:"randomization_group,omitempty"`
	// but all have the same shuffling
	RandomizationSeed int `json:"randomization_seed,omitempty"`

	Style *css.StylesResponsive `json:"style,omitempty"` // pointer, to avoid empty JSON blocks
	Class string                `json:"class,omitempty"` // additional explicit CSS class; for example   .group-class-1 > .grid-item-lvl-1 {...}
}

// AddInput creates a new input
// and adds this input to the group's inputs
func (gr *groupT) AddInput() *inputT {
	inp := &inputT{}
	gr.Inputs = append(gr.Inputs, inp)
	ret := gr.Inputs[len(gr.Inputs)-1]
	return ret
}

// addInputArg adds an existing input to the group
func (gr *groupT) addInputArg(inp *inputT) {
	gr.Inputs = append(gr.Inputs, inp)
}

// emptyTextblock creates a new input;
// an empty input *is* rendered a empty cell
func emptyTextblock() *inputT {
	inp := &inputT{}
	inp.Type = "textblock"
	inp.Label = trl.S{
		"en": " &nbsp; ",
		"de": " &nbsp; ",
	}
	return inp
}

// addEmptyTextblock creates a new input
// and adds this input to the group's inputs
func (gr *groupT) addEmptyTextblock() *inputT {
	inp := emptyTextblock()
	gr.Inputs = append(gr.Inputs, inp)
	ret := gr.Inputs[len(gr.Inputs)-1]
	return ret
}

/*
	Methods for group style;

	There are *generic* style methods
		css.ItemStartCA()
		css.TextCenter()
*/

// Vertical changes CSS grid style to vertical;
// compare default case GroupHTMLGridBased()
func (gr *groupT) Vertical(argRows ...int) {

	rows := 1
	if len(argRows) > 0 {
		rows = argRows[0]
	}

	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleBox.Display = "grid"
	gr.Style.Desktop.StyleGridContainer.AutoFlow = "column"
	// gr.Style.Desktop.StyleGridContainer.TemplateColumns = " " // empty string
	gr.Style.Desktop.StyleGridContainer.TemplateRows = strings.Repeat("1fr ", rows)

}

// WidthMax limits width in desktop view
// for instance to 30rem;
// mobile view: no limitation
//
//	compare page.WidthMax
func (gr *groupT) WidthMax(s string) {
	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleBox.WidthMax = s
	gr.Style.Mobile.StyleBox.WidthMax = "none" // => 100% of page - page has margins; replaced desktop max-width
}

// ColWidth custom col width - not equal for each
func (gr *groupT) ColWidths(colWidths string) {

	gr.Style = css.NewStylesResponsive(gr.Style)
	gr.Style.Desktop.StyleBox.Display = "grid"
	gr.Style.Desktop.StyleGridContainer.AutoFlow = "row"
	// gr.Style.Desktop.StyleGridContainer.TemplateColumns = "1.6fr    2.7fr 3.1fr 3.1fr 2.4fr    2.4fr  1.4fr"
	gr.Style.Desktop.StyleGridContainer.TemplateColumns = colWidths

}

// HasComposit - group contains composit element?
//
//	if yes, we retrieve its input names
func (q *QuestionnaireT) HasComposit(pgIdx, grIdx int) ([]string, bool, error) {

	gr := q.Pages[pgIdx].Groups[grIdx]

	hasComposit := false
	for _, inp := range gr.Inputs {
		if inp.Type == "dyn-composite" {
			if inp.DynamicFunc == "" {
				log.Panicf(`
					page %v group %v 
					contains a input type 'dyn-composite' - but inp.DynamicFunc is empty`,
					pgIdx,
					grIdx,
				)
			}
			hasComposit = true
			break
		}
	}

	if !hasComposit {
		return nil, false, nil
	}

	// further checks
	if hasComposit {
		for _, inp := range gr.Inputs {
			if inp.Type != "dyn-composite" && inp.Type != "dyn-composite-scalar" {
				log.Panicf(`
					page %v group %v 
					contains a input type 'dyn-composite' - but *normal* input types too`,
					pgIdx,
					grIdx,
				)
			}
		}
	}

	compFuncNameWithParamSet := gr.Inputs[0].DynamicFunc
	cF, seqIdx, paramSetIdx := parseComposite(compFuncNameWithParamSet)

	// execute in preflight mode; only return the input names
	_, inpNames, err := cF(q, seqIdx, paramSetIdx, true)

	return inpNames, hasComposit, err
}

// parseComposite returns the func, the sequence idx, the param set idx
func parseComposite(compFuncNameWithParamSet string) (CompositeFuncT, int, int) {

	splt := strings.Split(compFuncNameWithParamSet, "__")
	if len(splt) != 3 {
		log.Panicf(
			`composite func name %q 
			must consist of func name '__' param set index '__' sequence idx`,
			compFuncNameWithParamSet,
		)
	}

	compFuncName := splt[0]
	cF, ok := CompositeFuncs[compFuncName]
	if !ok {
		log.Panicf(
			`composite func name %q out of %q does not exist`,
			compFuncName,
			compFuncNameWithParamSet,
		)
	}

	seqIdx, err := strconv.Atoi(splt[1])
	if err != nil {
		log.Panicf(
			`second part of composite func name %q 
			could not be parsed into int
			%v`,
			compFuncNameWithParamSet,
			err,
		)
	}

	paramSetIdx, err := strconv.Atoi(splt[2])
	if err != nil {
		log.Panicf(
			`third part of composite func name %q 
			could not be parsed into int
			%v`,
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
	Section trl.S `json:"section,omitempty"`     // extra strong before label in content - summary headline for multiple pages
	Label   trl.S `json:"label,omitempty"`       // headline, set to "" to prevent rendering
	Desc    trl.S `json:"description,omitempty"` // abstract

	CounterProgress string `json:"counter_progress,omitempty"` // number shown in progress bar bullet; "-" overrides the natural counter - navigationSequenceNum
	Short           trl.S  `json:"short,omitempty"`            // sort version of section/label/description - in progress bar and navigation menu

	// Navi control stuff
	//
	// SuppressProgressbar
	// 		suppresses display *of* progress bar,
	//
	// NoNavigation
	// 		suppresses button "back to page X" and "continue to page Y"
	// 		suppresses display *in* progress bar,
	// 		for introduction pages and closing pages, or other standalone pages
	// 			back/continue buttons must be explicitly programmed, and set to next()/prev() or some page index
	//
	// NavigationCondition
	// 		provides additional, dynamic conditions
	// 		for exclusion of a page from navigation
	//
	// Both, NoNavigation and NavigationCondition,
	// are evaluated in func IsInNavigation()
	//
	SuppressProgressbar bool   `json:"suppress_progressbar,omitempty"`
	NoNavigation        bool   `json:"no_navigation,omitempty"`
	NavigationCondition string `json:"navigation_condition,omitempty"` // see NoNavigation

	// SuppressInProgressbar is a weak form of NoNavigation
	// 		because navigation buttons are still shown in body
	SuppressInProgressbar bool `json:"suppress_in_progressbar,omitempty"`

	navigationSequenceNum int // page number in navigation order; dynamically computed in MainH()

	Style *css.StylesResponsive `json:"style,omitempty"`

	// *not* a marker for questionnaire finished entirely;
	// see q.ClosingTime instead;
	// truncated to second
	Finished time.Time `json:"finished,omitempty"`

	Groups []*groupT `json:"groups,omitempty"`

	// we want to migrate those to inputs of type javascript-block
	ValidationFuncName string `json:"validation_func_name,omitempty"` // file name containing javascript validation func template, can be comma separated such as "pdsRange,pdsPage1"
	// there should be several strings - key/value - not just one "msg"
	ValidationFuncMsg trl.S `json:"validation_func_msg,omitempty"`

	// Page body is created dynamically.
	// Static testing becomes meaningless.
	GeneratorFuncName string `json:"generator_func_name,omitempty"`
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

// WidthMax limits width in desktop view;
// horizontal centering by default via WidthDefault();
// for instance to 30rem;
// mobile view: no limitation
// compare groupT.WidthMax
func (p *pageT) WidthMax(s string) {
	p.Style = css.NewStylesResponsive(p.Style)
	p.Style.Desktop.StyleBox.WidthMax = s
	p.Style.Mobile.StyleBox.WidthMax = "calc(100% - 1.2rem)" // 0.6rem margin-left and -right in mobile view
}

// WidthDefault is called for every page - setting auto margins
func (p *pageT) WidthDefault() {
	p.Style = css.NewStylesResponsive(p.Style)
	if p.Style.Desktop.StyleBox.Margin == "" && p.Style.Mobile.StyleBox.Margin == "" {
		p.Style.Desktop.StyleBox.Margin = "1.2rem auto 0 auto"
		p.Style.Mobile.StyleBox.Margin = "0.8rem auto 0 auto"
	}
}

// QuestionnaireT contains pages with groups with inputs
type QuestionnaireT struct {
	Survey SurveyT `json:"survey,omitempty"`
	UserID string  `json:"user_id,omitempty"` // participant ID, decimal, but string, i.E. 1011

	// Attrs are user specific key-value pairs -
	//    i.e. user country or euro-member
	// 		whereas surveyT.Params are specific to the survey + wave
	// Attrs survey ID and wave ID and lang code come from login;
	// additional Attrs come from cfg.Profiles - mediated by login `p` parameter
	//
	//    key 'survey_variant' loads distinct questionnaire templates
	//
	//  Attrs are dynamically replaced
	//    attr-country
	Attrs map[string]string `json:"user_attrs,omitempty"`

	// if any response key "finished" equals qst.Finished
	// this is set to time.Now() - truncated to second
	// it is the marker for preventing any more edits
	ClosingTime time.Time `json:"closing_time,omitempty"`
	RemoteIP    string    `json:"remote_ip,omitempty"`
	UserAgent   string    `json:"user_agent,omitempty"`
	MD5         string    `json:"md_5,omitempty"`

	LangCodes []string `json:"lang_codes,omitempty"` // default, order and availability - [en, de, ...] or [de, en, ...]
	LangCode  string   `json:"lang_code,omitempty"`  // current lang code - i.e. 'de' - session key lang_code

	CurrPage  int  `json:"curr_page,omitempty"`
	HasErrors bool `json:"has_errors,omitempty"` // If any response is faulty; set by ValidateReponseData

	// ShufflingVariations indicated how many different reshufflings occur;
	// until repetition; primitive permutation mechanism;
	// deterministically reordering / reshuffling a set of groups
	// based on
	//      * user id
	// 		* groupT.RandomizationSeed
	//
	// 	page idx is no longer relevant
	//  groupT.RandomizationGroup is only to distinguish multiple groups per page
	ShufflingVariations  int `json:"shuffling_variations,omitempty"`
	ShufflingRepetitions int `json:"shuffling_repetitions,omitempty"` // if equals 0, then defaults to three; usually you dont have to touch this value

	// PreventSkipForward - skipping back always possible,
	// skipping forward is preventable
	PreventSkipForward bool `json:"allow_skip_forward"`
	// PostponeNavigationButtons - how many seconds delay, before the navigation appears
	PostponeNavigationButtons int `json:"postpone_navigation_buttons,omitempty"`

	// Previously we had SurveyT.Variant; now all questionnaire variations
	// should be captured with a distinct VersionEffective.
	// Notice the difference to tpl.SiteCore();
	// tpl.SiteCore() yields a common *style* identifier
	// for completely different questionnaires
	VersionMax       int    `json:"version_max,omitempty"`    // total number of versions - usually permutations
	AssignVersion    string `json:"assign_version,omitempty"` // default is UserID modulo - other value is "round-robin"
	VersionEffective int    `json:"version_effective"`        // result of q.Version()

	MaxGroups int `json:"max_groups,omitempty"` //  Max number of groups - a helper value - computed during questionnaire creation - previously used for shuffing of groups.

	Pages []*pageT `json:"pages,omitempty"`

	// Version int `json:"version,omitempty"` // 0 - rendering as HTML table   -   1 - rendering as CSS Grid
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
// Duplicate of lgn.basePath - cyclic dependencies
func BasePath() string {
	return path.Join(".", "responses")
}

// unusedFinishedEntirely does not go for the
// page.Finished timestamps, but for
// an explicit input called 'finished'
//
// use !q.ClosingTime.IsZero instead
func (q *QuestionnaireT) unusedFinishedEntirely() (closed bool) {
	for _, p := range q.Pages {
		for _, gr := range p.Groups {
			for _, inp := range gr.Inputs {
				if inp.Name == "finished" {
					if inp.Response == Finished {
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

// AddPageAfter creates a new page
// and adds this page to the questionnaire's pages
func (q *QuestionnaireT) AddPageAfter(idx int) *pageT {

	if idx < 0 || idx > len(q.Pages)-1 {
		log.Panicf("AddPageAfter(): %v pages - valid indexes are 0...%v", len(q.Pages), len(q.Pages)-1)
	}

	cntr := ctr.Increment()
	p := &pageT{
		Label: trl.S{"en": fmt.Sprintf("PageLabel_%v", cntr), "de": fmt.Sprintf("Seitentitel_%v", cntr)},
		Desc:  trl.S{"en": "", "de": ""},
	}

	q.Pages = append(q.Pages, nil)         // make room
	copy(q.Pages[idx+2:], q.Pages[idx+1:]) // shift one slot to the right
	q.Pages[idx+1] = p                     //

	return q.Pages[idx+1]
}

// EditPage returns page X
// compare WrapPage
func (q *QuestionnaireT) EditPage(idx int) *pageT {
	if idx < 0 || idx > len(q.Pages)-1 {
		log.Panicf("EditPage(): %v pages - valid indexes are 0...%v", len(q.Pages), len(q.Pages)-1)
	}
	return q.Pages[idx]
}

// AddFinishButtonNextToLast from page
//
// Adding explicit button to finish page, which is outsite navigation.
// Button is added to the next-to-last page.
// Call this method at the end of page insertions.
func (q *QuestionnaireT) AddFinishButtonNextToLast() {

	ln := len(q.Pages)

	if ln < 2 {
		log.Panicf(
			"AddFinishButton(): At least 2 pages needed, before we can add a 'finish' button; has %v",
			len(q.Pages))
	}

	page := q.Pages[ln-2]

	{
		gr := page.AddGroup()
		gr.BottomVSpacers = 2
		gr.Cols = 2

		{
			inp := gr.AddInput()
			inp.Type = "button"
			inp.Name = "finished"
			inp.Name = "submitBtn"
			inp.Response = fmt.Sprintf("%v", len(q.Pages)-1) // we assume, all pages were already added
			// inp.Label = cfg.Get().Mp["end"]
			inp.Label = cfg.Get().Mp["finish_questionnaire"]
			inp.ColSpan = 2
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 1
			inp.AccessKey = "n"

			inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
			inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
		}
	}

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
			err := fmt.Errorf("language code '%v' is not supported in %v", newCode, q.LangCodes)
			log.Print(err)
			return err
		}
		q.LangCode = newCode
	}
	return nil
}

// PrevPage computation;
// q contains currPage from *last* request;
// we remember this, because we want to store request values *there*
func (q *QuestionnaireT) PrevPage() (prevPage int) {
	prevPage = q.CurrPage
	if prevPage > len(q.Pages)-1 || prevPage < 0 {
		q.CurrPage = 0
		prevPage = 0
	}
	return prevPage
}

// FindNewPage determines the new page
func (q *QuestionnaireT) FindNewPage(sess *sessx.SessT) {

	prevPage := q.PrevPage()
	currPage := prevPage // Default assumption: we are still on prev page - unless there is some modification:
	submit := sess.EffectiveStr("submitBtn")
	if submit == "prev" {
		currPage = q.Prev()
	} else if submit == "next" {
		currPage = q.Next()
	} else {
		// Apart from "prev" and "next", submitBtn can also hold an explicit destination page
		explicit, ok, err := sess.EffectiveInt("submitBtn")
		if err != nil {
			// invalid page value, just dont use it
		}
		if ok && err == nil && explicit > -1 {
			log.Printf("curPage set explicitly by 'submitBtn' to %v", explicit)
			currPage = explicit
		}
	}
	// The progress bar uses "page" to submit an explicit destination page.
	// There are no conflicts of overriding submitBtn and page
	// since submitBtn has only a value if actually pressed.
	explicit, ok, err := sess.EffectiveInt("page")
	if err != nil {
		// invalid page value, just dont use it
	}
	if ok && err == nil && explicit > -1 {
		log.Printf("curPage set explicitly by param 'page' to %v", explicit)
		currPage = explicit
	}
	q.CurrPage = currPage // Put current page into questionnaire
	log.Printf("submitBtn was '%v' - new currPage is %v", submit, currPage)

}

// CurrentPageHTML is a comfort shortcut to PageHTML
func (q *QuestionnaireT) CurrentPageHTML() (string, error) {
	return q.PageHTML(q.CurrPage)
}

// GetLangCode -
func (q *QuestionnaireT) GetLangCode() string {
	return q.LangCode
}

// shufflingGroupsT is a helper for RandomizeOrder()
type shufflingGroupsT struct {
	Orig     int // orig pos
	Shuffled int // shuffled pos - new pos

	GroupID           int // shuffling group
	RandomizationSeed int

	Start int // shuffling group start idx    - across gaps
	Idx   int // sequence in shuffling group  - across gaps - dense 0,1...6,7

	// seqStart int // shuffling group start idx - continuous chunk
	// seqIdx   int // index in shuffling group  - continuous chunk
}

// String representation for dump
func (sg shufflingGroupsT) String() string {
	return fmt.Sprintf("orig %02v -> shuff %02v - G%v Sd%v strt%02v seq%v", sg.Orig, sg.Shuffled, sg.GroupID, sg.RandomizationSeed, sg.Start, sg.Idx)
}

// var debugShuffling = true
var debugShuffling = false

// RandomizeOrder creates a shuffled ordering of groups
// determined by UserID and .RandomizationGroup;
// groups with RandomizationGroup==0 retain their position
// on fixed order position;
// others get a randomized position
func (q *QuestionnaireT) RandomizeOrder(pageIdx int) []int {

	p := q.Pages[pageIdx]

	// helper - separating groups by their RandomizationGroup value -
	// with positional indexes
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

	if debugShuffling {
		log.Printf(
			"max sg idx %v \nshufflingGroups %v",
			maxSg,
			dbg.Dump2String(shufflingGroups),
		)
	}

	//
	// compute the main array
	sgs := make([]shufflingGroupsT, len(p.Groups))
	for i := 0; i < len(p.Groups); i++ {
		sg := p.Groups[i].RandomizationGroup
		sgs[i].Orig = i
		sgs[i].GroupID = sg
		sgs[i].RandomizationSeed = p.Groups[i].RandomizationSeed
		sgs[i].Start = shufflingGroups[sg][0]
		sgs[i].Idx = shufflingGroupsCntr[sg]
		shufflingGroupsCntr[sg]++
	}

	//
	// randomize
	for i := 0; i < len(sgs); i++ {
		if sgs[i].GroupID == 0 {
			sgs[i].Shuffled = sgs[i].Orig
		} else {
			if sgs[i].Idx == 0 {
				sg := sgs[i].GroupID
				// this must conform with ShufflesToCSV()
				seedShufflngs := q.UserIDInt() + sgs[i].RandomizationSeed

				if q.Survey.Type == "fmt" && q.Survey.Year == 2022 && q.Survey.Month == 12 {
					seedShufflngs = fmtRandomizationGroups[q.UserIDInt()] + sgs[i].RandomizationSeed
					if debugShuffling {
						log.Printf("fmt-2022-12 seedShufflngs for user %v is %v", q.UserIDInt(), seedShufflngs)
					}
				}

				sh := shuffler.New(seedShufflngs, q.ShufflingVariations, len(shufflingGroups[sg]))
				newOrder := sh.Slice(q.ShufflingRepetitions) // adding sg breaks compatibility to ShufflesToCSV()
				if debugShuffling {
					log.Printf("%v - seq %16s in order %16s - iter %v", sg, fmt.Sprint(shufflingGroups[sg]), fmt.Sprint(newOrder), pageIdx+sg)
				}
				for i := 0; i < len(shufflingGroups[sg]); i++ {
					offset := shufflingGroups[sg][i] // i.e. [1, 9]
					i2 := newOrder[i]
					sgs[offset].Shuffled = shufflingGroups[sg][i2]
				}
			}
		}
	}

	if debugShuffling {
		for i := 0; i < len(sgs); i++ {
			log.Printf("lp%02v  %v", i, sgs[i])
		}
	}

	// extract the new order - with randomized elements
	shuffled := make([]int, len(p.Groups))
	for i := 0; i < len(p.Groups); i++ {
		shuffled[i] = sgs[i].Shuffled
	}

	if debugShuffling {
		log.Printf("=> shuffled %v", shuffled)
	}

	return shuffled

}

// RenderJS prints the contents of a JavaScript template into the HTML response;
// used for page level JavaScript blocks;
// used also for inputs of type 'javascript-block'
func (q *QuestionnaireT) RenderJS(
	w io.Writer,
	fileName string,
	translations map[string]trl.S,
	jsStrings map[string]string,
) {

	tplFile := fileName + ".js"

	mp := map[string]interface{}{}

	for k, v := range translations {
		mp[k] = v.Tr(q.LangCode)
	}
	for k, v := range jsStrings {
		mp[k] = template.JS(v)
	}

	t, err := ParseJavaScript(tplFile)
	if err != nil {
		log.Printf("Error parsing tpl %v: %v", tplFile, err)
	} else {
		err := t.Execute(w, mp)
		if err != nil {
			log.Printf("Error executing tpl %v: %v", tplFile, err)
		}
	}

}

// PageHTML generates HTML for a specific page of the questionnaire
func (q *QuestionnaireT) PageHTML(pageIdx int) (string, error) {

	if q.CurrPage > len(q.Pages)-1 || q.CurrPage < 0 {
		s := fmt.Sprintf("You requested page %v out of %v. Page does not exist", pageIdx, len(q.Pages)-1)
		log.Print(s)
		return s, fmt.Errorf(s)
	}

	page := q.Pages[pageIdx]

	kv := q.DynamicPageValues()
	err := q.DynamicPages()
	if err != nil {
		err = fmt.Errorf("dyn page creation in PageHTML() q: %w", err)
		return err.Error(), err
	}
	q.DynamicPagesApplyValues(kv)

	found := false
	for _, lc := range q.LangCodes {
		if q.LangCode == lc {
			found = true
			break
		}
	}
	if !found {
		s := fmt.Sprintf("Language code '%v' is not supported in %v", q.LangCode, q.LangCodes)
		log.Print(s)
		return s, fmt.Errorf(s)
	}

	// adding a group containing the previous/next buttons
	if !page.NoNavigation {

		var footer *groupT

		if len(page.Groups) == 0 { // stupid edge case
			footer = page.AddGroup()
			footer.ID = "footer"
			footer.BottomVSpacers = 0
		}
		if page.Groups[len(page.Groups)-1].ID == "footer" {
			footer = page.Groups[len(page.Groups)-1]
			footer.Inputs = nil
		} else {
			footer = page.AddGroup()
			footer.ID = "footer"
			footer.BottomVSpacers = 0
		}
		footer.Cols = 2

		lblNext := cfg.Get().Mp["page"]
		lblNext = cfg.Get().Mp["continue_to_page_x"]
		cloneNext := lblNext.Pad(2)
		cloneNext = cloneNext.Fill(q.NextNaviNum())

		lblPrev := cfg.Get().Mp["previous"]
		lblPrev = cfg.Get().Mp["back_to_page_x"]
		clonePrev := lblPrev.Pad(1)
		clonePrev = clonePrev.Fill(q.PrevNaviNum())

		if q.HasNext() {
			inp := footer.AddInput()
			inp.Type = "button"
			inp.Name = "submitBtn"
			inp.Response = "next"
			inp.Label = cloneNext
			inp.AccessKey = "n"
			inp.ColSpanControl = 1

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleGridItem.Order = 2

			inp.StyleCtl = css.ItemEndMA(inp.StyleCtl)
			inp.StyleCtl.Desktop.StyleBox.Position = "relative"
			inp.StyleCtl.Desktop.StyleBox.Left = "3rem"
			inp.StyleCtl.Mobile.StyleBox.Left = "0"
		} else {
			inp := footer.addEmptyTextblock()
			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleGridItem.Order = 2
		}

		if q.HasPrev() {
			inp := footer.AddInput()
			inp.Type = "button"
			inp.Name = "submitBtn"
			inp.Response = "prev"
			inp.Label = clonePrev
			inp.AccessKey = "p"
			inp.ColSpanControl = 1

			inp.Style = css.NewStylesResponsive(inp.Style)
			inp.Style.Desktop.StyleGridItem.AlignSelf = "end" // smaller font-size

			inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
			// inp.StyleCtl = css.ItemEndMA(inp.StyleCtl)
			inp.StyleCtl.Desktop.StyleBox.Position = "relative"
			inp.StyleCtl.Desktop.StyleBox.Left = "-2.5rem"
			inp.StyleCtl.Mobile.StyleBox.Left = "0"
			inp.StyleCtl.Mobile.StyleText.FontSize = 85

		} else {
			footer.addEmptyTextblock()
		}
	}

	w := &strings.Builder{}

	//
	//
	if q.PostponeNavigationButtons > 0 {

		s := `
		<style>
			button[type="submit"], 
			button[accesskey="n"] 
			{
				animation:  %vms ease-in-out 1ms 1 nameAppear;
			}
		</style>
		`
		fmt.Fprintf(w, s, 1000*q.PostponeNavigationButtons)

	}

	page.WidthDefault()
	pageClass := fmt.Sprintf("pg%02v", pageIdx)
	fmt.Fprint(w, css.StyleTag(page.Style.CSS(pageClass)))

	// i.e. smaller - for i.e. radios more closely together
	width := fmt.Sprintf("<div class='%v' >\n", pageClass)
	fmt.Fprint(w, width)

	if q.HasErrors {
		fmt.Fprintf(w,
			`<p class="error" id="page-error" >%v</p>`,
			cfg.Get().Mp["correct_errors"].Tr(q.LangCode),
		)
	}

	hasHeader := false

	if page.Section != nil {
		fmt.Fprintf(w, "<span class='go-quest-page-section' >%v</span>", page.Section.Tr(q.LangCode))
		if page.Label.Tr(q.LangCode) != "" {
			fmt.Fprint(w, "<span class='go-quest-page-desc'> &nbsp; - &nbsp; </span>")
		}
		hasHeader = true
	}
	if page.Label.Tr(q.LangCode) != "" {
		fmt.Fprintf(w, "<span class='go-quest-page-header' >%v</span>", page.Label.Tr(q.LangCode))
		hasHeader = true
	}
	if page.Desc.Tr(q.LangCode) != "" {
		fmt.Fprint(w, vspacer0)
		fmt.Fprintf(w, "<p  class='go-quest-page-desc'>%v</p>", page.Desc.Tr(q.LangCode))
		hasHeader = true
	}

	if hasHeader {
		fmt.Fprint(w, vspacer16)
	}

	grpOrder := q.RandomizeOrder(pageIdx)

	page.ConsolidateRadioErrors(grpOrder)

	compositCntr := -1    // group counter - per page
	nonCompositCntr := -1 // group counter - per page
	for loopIdx, grpIdx := range grpOrder {

		if _, ok, _ := q.HasComposit(pageIdx, grpIdx); ok {
			compositCntr++
			compFuncNameWithParamSet := page.Groups[grpIdx].Inputs[0].DynamicFunc
			cF, seqIdx, paramSetIdx := parseComposite(compFuncNameWithParamSet)
			grpHTML, _, err := cF(q, seqIdx, paramSetIdx, false) // QuestionnaireT must comply to qstif.Q
			if err != nil {
				fmt.Fprintf(w, "composite func error %v \n", err)
			} else {
				// grpHTML also contains HTML and CSS stuff - which could be hyphenized too
				grpHTML = trl.HyphenizeText(grpHTML)
				fmt.Fprint(w, grpHTML+"\n")
			}
		} else {
			grpHTML := q.GroupHTMLGridBased(pageIdx, grpIdx)

			// dynamic numbering - based on group sequence per page after shuffling
			if strings.Contains(grpHTML, "[groupID]") {
				nonCompositCntr++
				grpHTML = strings.Replace(grpHTML, "[groupID]", fmt.Sprintf("%v", nonCompositCntr+1), -1)
			}

			// dynamic question numbering - based on NavigationCondition, IsInNavigation()
			// todo

			fmt.Fprint(w, grpHTML+"\n")
		}

		// vertical distance at the end of groups
		if loopIdx < len(page.Groups)-1 {
			for i2 := 0; i2 < page.Groups[grpIdx].BottomVSpacers; i2++ {
				fmt.Fprint(w, vspacer16)
			}
		} else {
			fmt.Fprint(w, vspacer16)
		}

	}

	fmt.Fprintf(w, "</div> <!-- /%v -->\n\n", pageClass)

	//
	//
	if page.ValidationFuncName != "" {

		funcNames := strings.Split(page.ValidationFuncName, ",")
		for _, funcName := range funcNames {
			q.RenderJS(
				w,
				funcName,
				map[string]trl.S{"msg": page.ValidationFuncMsg},
				map[string]string{},
			)
		}

	}

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

// SuppressProgressBar - dont show progress bar in current page;
// while back/forth buttons may still be rendered
// convenience func for templates;
func (q *QuestionnaireT) SuppressProgressBar() bool {
	return q.Pages[q.CurrPage].SuppressProgressbar
}

// IsInNavigation checks whether pageIdx is suitable
// as next or previous page
// and whether it should show up in progress bar
func (q *QuestionnaireT) IsInNavigation(pageIdx int) bool {

	if pageIdx < 0 || pageIdx > len(q.Pages)-1 {
		return false
	}

	if q.Pages[pageIdx].NoNavigation {
		return false
	}

	if fc, ok := naviFuncs[q.Pages[pageIdx].NavigationCondition]; ok {
		return fc(q, pageIdx)
	}

	return true
}

// EnumeratePages allocates a sequence number
// based on IsInNavigation()
func (q *QuestionnaireT) EnumeratePages() {
	pageCntr := 0
	for i1 := 0; i1 < len(q.Pages); i1++ {
		if q.IsInNavigation(i1) {
			pageCntr++
			q.Pages[i1].navigationSequenceNum = pageCntr
		}
	}
}

// next page to be shown in navigation
func (q *QuestionnaireT) nextInNavi() (int, bool) {
	// Find next page in navigation
	for i := q.CurrPage + 1; i < len(q.Pages); i++ {
		if q.IsInNavigation(i) {
			return i, true
		}
	}
	// Fallback: Last page in navigation
	for i := len(q.Pages) - 1; i >= 0; i-- {
		if q.IsInNavigation(i) {
			return i, false
		}
	}
	return len(q.Pages) - 1, false
}

// prev page to be shown in navigation
func (q *QuestionnaireT) prevInNavi() (int, bool) {
	// Find prev page in navigation
	for i := q.CurrPage - 1; i >= 0; i-- {
		if q.IsInNavigation(i) {
			return i, true
		}
	}
	// Fallback: First page in navigation
	for i := 0; i < len(q.Pages); i++ {
		if q.IsInNavigation(i) {
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
	return fmt.Sprintf("%v", q.Pages[pg].navigationSequenceNum)
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
	return fmt.Sprintf("%v", q.Pages[pg].navigationSequenceNum)
}

// Compare compares page completion times and input responses.
// Compare stops with the first difference and returns an error.
func (q *QuestionnaireT) Compare(v *QuestionnaireT, lenient bool) (bool, error) {

	if len(q.Pages) != len(v.Pages) {
		return false, fmt.Errorf("unequal numbers of pages: %v - %v", len(q.Pages), len(v.Pages))
	}

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if len(q.Pages[i1].Groups) != len(v.Pages[i1].Groups) {
			return false, fmt.Errorf("page %v: Unequal numbers of groups: %v - %v", i1, len(q.Pages[i1].Groups), len(v.Pages[i1].Groups))
		}
		if i1 < len(q.Pages)-1 { // No completion time comparison for last page
			qf := q.Pages[i1].Finished
			vf := v.Pages[i1].Finished
			if qf.Sub(vf) > 30*time.Second || vf.Sub(qf) > 30*time.Second {
				return false, fmt.Errorf("page %v: Completion time too distinct: %v - %v", i1, qf, vf)
			}
		}

		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			if len(q.Pages[i1].Groups[i2].Inputs) != len(v.Pages[i1].Groups[i2].Inputs) {
				return false, fmt.Errorf("page %v: Group %v: Unequal numbers of inputs: %v - %v", i1, i2, len(q.Pages[i1].Groups[i2].Inputs), len(v.Pages[i1].Groups[i2].Inputs))
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
						"page %v: Group %v: Input %v %v: '%v' != '%v'",
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

// DelocalizeNumber removes localized number formatting;
// HTML5 input type number does not de-localize values;
// 123,456.78 is allowed - and so is 123.456,78;
// at least we have at most *one* dot and *one* comma;
// https://www.ryadel.com/en/html-input-type-number-with-localized-decimal-values-jquery/
func DelocalizeNumber(s string) string {

	if !strings.Contains(s, ",") {
		// no comma; all is fine
		return s
	}

	// contains comma, but contains no dot
	if !strings.Contains(s, ".") {
		// => just comma instead of dot
		s = strings.ReplaceAll(s, ",", ".")
		return s
	}

	// contains comma *and* dot
	//
	// 123,456.78
	if strings.Index(s, ".") > strings.Index(s, ",") {
		// => remove  the thousands separator - 123456.78
		s = strings.ReplaceAll(s, ",", "")
		return s
	}
	//
	// 123.456,78
	if strings.Index(s, ",") > strings.Index(s, ".") {
		// => remove  the thousands separator - 123456,78
		s = strings.ReplaceAll(s, ".", "")
		// => replace the decimal   separator - 123456.78
		s = strings.ReplaceAll(s, ",", ".")
		return s
	}

	return s
}

var germanUmlaute = strings.NewReplacer(
	"ä", "ae",
	"ö", "oe",
	"ü", "ue",
	"Ä", "ae",
	"Ö", "oe",
	"Ü", "ue",
	"ß", "ss",
)

var separators = strings.NewReplacer(
	"\r\n", " -- ",
	"\n", " -- ",
	":", " - ",
	";", " - ",
	// ",", " - ",
)

// no comma, no colon, no semicolon - only dot or hyphen
// no newline
var englishTextAndNumbersOnly = regexp.MustCompile(`[^a-zA-Z0-9\.\_\-\+ ]+`)
var severalSpaces = regexp.MustCompile(`[ ]+`)

// EnglishTextAndNumbersOnly replaces all other UTF characters by space
func EnglishTextAndNumbersOnly(s string) string {

	// sBefore := s

	s = strings.TrimSpace(s)
	s = separators.Replace(s)

	s = germanUmlaute.Replace(s)
	s = englishTextAndNumbersOnly.ReplaceAllString(s, " ")
	s = severalSpaces.ReplaceAllString(s, " ")

	s = strings.TrimSpace(s)

	// if sBefore != s {
	// 	log.Printf("\n%v\n%v", sBefore, s)
	// }

	return s
}

// CleanseUserAgent is stricter than EnglishTextAndNumbersOnly
func CleanseUserAgent(s string) string {
	s = EnglishTextAndNumbersOnly(s)
	s = strings.ReplaceAll(s, "+", "plus") // additional removal
	return s
}

// compare codebase openingSpan
var openingDiv = regexp.MustCompile(`<div.*?>`)
var openingP = regexp.MustCompile(`<p.*?>`)

// LabelCleanse removes some common HTML stuff;
// argument q is not yet used
func (q *QuestionnaireT) LabelCleanse(s string) string {

	s = openingDiv.ReplaceAllString(s, " ")
	s = openingP.ReplaceAllString(s, " ")
	s = strings.ReplaceAll(s, "</div>", " ")
	s = strings.ReplaceAll(s, "</p>", " ")

	s = strings.ReplaceAll(s, "&#931;", " sum ") // Σ - greek sum symbol
	s = strings.ReplaceAll(s, "&shy;", "")
	s = strings.ReplaceAll(s, "&nbsp;", " ")
	s = strings.ReplaceAll(s, "<br>", " ")

	s = strings.ReplaceAll(s, "<b>", " ")
	s = strings.ReplaceAll(s, "</b>", " ")
	s = strings.ReplaceAll(s, "<bx>", " ")
	s = strings.ReplaceAll(s, "</bx>", " ")

	s = EnglishTextAndNumbersOnly(s)

	s = strings.TrimPrefix(s, "-- ")
	s = strings.TrimSuffix(s, " --")

	return s
}

// see unit test for LabelIsOutline()
var outlineNumbering1 = regexp.MustCompile(`^[0-9]+[a-z\.)]*[\s]+`)
var outlineNumbering2 = regexp.MustCompile(`^[a-zA-Z][\.)]*[\s]+`)

// LabelIsOutline - if s starts with some outline
// see unit test
func (q *QuestionnaireT) LabelIsOutline(s string) bool {
	return outlineNumbering1.MatchString(s) || outlineNumbering2.MatchString(s)
}

// cleanseIdentical takes the labeling for multiple radio inputs;
// radio input labels are concatenated from bottom to top;
// resulting in multiple identical labels;
// i.e. three times
//
//	Your growth estimate? Ger        O  O  O
//
// and three more times
//
//	Your growth estimate? Ger   US   O  O  O
//
// or they might differ only in a suffix
//
//	Are you Alice?        Yes O   No O
//
// use cleanseIdentical for the first case
// use cleansePrefixes for the second case
func cleanseIdentical(ss []string) []string {

	counted := map[string]int{}
	ret := []string{}
	for _, s := range ss {
		counted[s]++
		if counted[s] < 2 {
			ret = append(ret, s)
		}
	}

	return ret

}

// cleansePrefixes - see cleanseIdentical()
// for documentation
func cleansePrefixes(ss []string) []string {

	ret := []string{}
	for _, s := range ss {
		stripped := ""
		for i := len(ss) - 1; i > -1; i-- { // reversely
			pref := ss[i]
			if s != pref && strings.HasPrefix(s, pref) {

				stripped = strings.TrimPrefix(s, pref)

				stripped = strings.TrimSpace(stripped)
				stripped = strings.TrimPrefix(stripped, "-- ")
				stripped = strings.TrimSuffix(stripped, " --")

				// log.Printf("stripped off\n\t%q  \n\t%q  \n\t%q", s, pref, stripped)
				break
			}
		}
		if stripped == "" {
			ret = append(ret, s)
		} else {
			ret = append(ret, stripped)
		}
	}

	return ret

}

// LabelsByInputNames extracts the label texts for each input;
// starting from the input to the top;
// since there is a lot of summarized labeling for multiple
// inputs, we have no clear relationship;
// for each input, we traverse up on the page
// collecting all text until we hit a piece of text
// which is an outline number.
//
// If no limiting outline number is found, text is concatenated
// all the way up to the start of the page.
//
// functions cleanseIdentical(...) and cleansePrefixes(...)
// are used to clear out redundancies; see documentation.
func (q *QuestionnaireT) LabelsByInputNames() (lblsByNames map[string]string, keys, lbls []string) {

	lblsByNames = map[string]string{} // init return

	// helpers
	keysByPage := make([][]string, len(q.Pages))
	lblsByPage := make([][]string, len(q.Pages))
	labelsPerRadio := map[string][]string{}

	for i1 := 0; i1 < len(q.Pages); i1++ {
		for i2 := 0; i2 < len(q.Pages[i1].Groups); i2++ {
			for i3 := 0; i3 < len(q.Pages[i1].Groups[i2].Inputs); i3++ {

				inp := q.Pages[i1].Groups[i2].Inputs[i3]

				if inp.IsLayout() {
					continue
				}

				// going up/back until we have found a label
				lbl := ""
			nestedLoop:
				for grUp := i2; grUp > -1; grUp-- {

					countDownInputsFrom := i3
					if grUp != i2 {
						countDownInputsFrom = len(q.Pages[i1].Groups[grUp].Inputs) - 1
					}

					for inpUp := countDownInputsFrom; inpUp > -1; inpUp-- {
						lb := q.Pages[i1].Groups[grUp].Inputs[inpUp].Label.TrSilent("en")
						lb = q.LabelCleanse(lb)
						if lb != "" {
							if lbl != "" {
								// slow, create a string buffer someday:
								lbl = lb + " -- " + lbl
							} else {
								lbl = lb
							}
						}
						if q.LabelIsOutline(lb) {
							// log.Printf("\t\t\tfound lb at gr%02v.inp%02v: '%v'", grUp, inpUp, q.LabelCleanse(lb))
							break nestedLoop
						}
					}
				}

				lblsByPage[i1] = append(lblsByPage[i1], lbl)
				keysByPage[i1] = append(keysByPage[i1], inp.Name)

				// special treatment for radio inputs - who occur several times:
				//   collect their labels
				if inp.Type == "radio" {
					if lbl != "" {
						if labelsPerRadio[inp.Name] == nil {
							labelsPerRadio[inp.Name] = []string{}
						}
						labelsPerRadio[inp.Name] = append(labelsPerRadio[inp.Name], lbl)
					}
				}

			}
		}

	}

	// cleanse repeating radios
	for pageIdx := 0; pageIdx < len(keysByPage); pageIdx++ {
		for inpIdx, inpName := range keysByPage[pageIdx] {
			if lbls, ok := labelsPerRadio[inpName]; ok {
				lbls = cleanseIdentical(lbls)
				lbls = cleansePrefixes(lbls)
				lblsByPage[pageIdx][inpIdx] = strings.Join(lbls, " -- ")
			}
		}
	}

	// cleanse repeating prefixes
	for pageIdx := 0; pageIdx < len(lblsByPage); pageIdx++ {
		lblsByPage[pageIdx] = cleansePrefixes(lblsByPage[pageIdx])
	}

	for pageIdx := 0; pageIdx < len(keysByPage); pageIdx++ {
		for inpIdx, inpName := range keysByPage[pageIdx] {
			keys = append(keys, inpName)
			lbls = append(keys, lblsByPage[pageIdx][inpIdx])
			lblsByNames[inpName] = lblsByPage[pageIdx][inpIdx]
		}
	}

	return
}

// KeysValues returns all pages finish times; keys and values in defined order.
// Empty values are also returned.
// Major purpose is CSV export across several questionnaires.
func (q *QuestionnaireT) KeysValues(cleanse bool) (finishes, keys, vals []string) {
	// log.Printf("Collecting keys+vals for %v", q.UserID)
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
				val := q.Pages[i1].Groups[i2].Inputs[i3].Response
				if cleanse {
					if q.Pages[i1].Groups[i2].Inputs[i3].Type == "number" {
						val = DelocalizeNumber(val)
					}
					val = EnglishTextAndNumbersOnly(val)
				}
				vals = append(vals, val)
			}
		}
	}
	// log.Printf("%v", keys)
	// log.Printf("%v", vals)
	return
}

// UserIDInt retrieves the userID as int
func (q *QuestionnaireT) UserIDInt() int {

	if q.UserID == "systemtest" {
		return -3216
	}

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

var ctrLogin = ctr.New()

// Version retrieves the questionnaire's version;
// default - version depends on user ID
// 'round-robin' - version depends on login order
func (q *QuestionnaireT) Version() int {

	if q == nil {
		return 0 // questionnaire creation
	}

	if q.VersionMax > 0 && q.VersionEffective < 0 {
		if strings.ToLower(q.Survey.Type) == "pat" && q.UserIDInt() > 80000 && q.UserIDInt() < 81000 {
			q.VersionEffective = int(ctrLogin.Increment()) % q.VersionMax
			// log.Printf("Assign version based on central counter: %v", q.VersionEffective)
		} else if q.AssignVersion == "round-robin" {
			q.VersionEffective = int(ctrLogin.Increment()) % q.VersionMax
		} else {
			q.VersionEffective = q.UserIDInt() % q.VersionMax
			if q.UserIDInt() == -3216 {
				q.VersionEffective = 0 // damn hack for user systemtest
			}
			// log.Printf("Assign version based on user id: %v", q.VersionEffective)
		}
	}

	return q.VersionEffective
}

// ByName retrieves an input element by name.
// Returns nil if the input element was not found.
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

// ResponseByName implements qstif.Q
func (q *QuestionnaireT) ResponseByName(n string) (string, error) {
	inp := q.ByName(n)
	if inp == nil {
		return "input-name-not-found", fmt.Errorf("input named %v not fount", n)
	}
	return inp.Response, nil
}

// ErrByName returns the error for an input name;
// implements qstif.Q
func (q *QuestionnaireT) ErrByName(n string) (string, error) {
	inp := q.ByName(n)
	if inp == nil {
		return "", fmt.Errorf("input named %v not fount", n)
	}
	return inp.ErrMsg, nil
}

// ParseJavaScript loads js file
// and embeds its contents into a <script> tag;
// so that parsing the template does not do harmful escaping;
//
// would better belong into package tpl - but circular dependencies;
func ParseJavaScript(tName string) (*template.Template, error) {

	pth := path.Join(".", "templates", "js", tName) // not filepath; cloudio always has forward slash
	cnts, err := cloudio.ReadFile(pth)
	if err != nil {
		msg := fmt.Sprintf("cannot open template %v: %v", pth, err)
		return nil, fmt.Errorf(msg+" %w", err)
	}

	w := &strings.Builder{}
	fmt.Fprintf(w, "<script>\n")
	fmt.Fprint(w, string(cnts))
	// put the console message into the script itself, if needed
	// fmt.Fprintf(w, "console.log('JS tpl %v successfully added')", pth)
	fmt.Fprintf(w, "</script>\n")

	base := template.New(tName)
	tDerived, err := base.Parse(w.String())
	if err != nil {
		msg := fmt.Sprintf("parsing failed for %v: %v", pth, err)
		return nil, fmt.Errorf(msg+" %w", err)
	}

	return tDerived, nil
}

func (q *QuestionnaireT) DynamicPageValues() map[string]string {

	ret := map[string]string{}

	for i1 := 0; i1 < len(q.Pages); i1++ {
		if q.Pages[i1].GeneratorFuncName == "" {
			continue
		}
		// log.Printf("\t\tpage %vi1 is dynamic...", i1)
		page := q.Pages[i1]
		cleanse := false
		for i2 := 0; i2 < len(page.Groups); i2++ {
			for i3 := 0; i3 < len(page.Groups[i2].Inputs); i3++ {
				if page.Groups[i2].Inputs[i3].IsLayout() {
					continue
				}
				// keys = append(keys, page.Groups[i2].Inputs[i3].Name)
				val := page.Groups[i2].Inputs[i3].Response
				if cleanse {
					if page.Groups[i2].Inputs[i3].Type == "number" {
						val = DelocalizeNumber(val)
					}
					val = EnglishTextAndNumbersOnly(val)
				}
				key := page.Groups[i2].Inputs[i3].Name
				ret[key] = val
				// log.Printf("\t\t\tkey %v - val %v", key, val)
			}
		}
	}

	return ret
}
func (q *QuestionnaireT) DynamicPagesApplyValues(kv map[string]string) {
	for i1 := 0; i1 < len(q.Pages); i1++ {
		if q.Pages[i1].GeneratorFuncName == "" {
			continue
		}
		page := q.Pages[i1]
		for i2 := 0; i2 < len(page.Groups); i2++ {
			for i3 := 0; i3 < len(page.Groups[i2].Inputs); i3++ {
				if page.Groups[i2].Inputs[i3].IsLayout() {
					continue
				}
				key := page.Groups[i2].Inputs[i3].Name
				page.Groups[i2].Inputs[i3].Response = kv[key]
			}
		}
	}
}

// DynamicPages dynamically re-creates groups and inputs
// based on conditions like q.UserID or values from other pages;
//
// there are three issues integrating this method with the Join() method:
//
//   - the structure depends on user input from other pages - we must call Join() first
//   - Join() will see distinct structures for groups/inputs between base (empty) and split (previously entered data) - we coded an exception
//   - DynamicPages() will now create re-groups and _empty_ inputs
//   - We now pull _previous_ dynamic page values (qSplit.DynamicPageValues) and apply them to the joined questionnaire
//
// We have to do a similar thing in PageHTML()
// because conditional values in other forms may have changed
func (q *QuestionnaireT) DynamicPages() error {

	dynPagesCreated := false
	for i1 := 0; i1 < len(q.Pages); i1++ {
		// dynamic page creation
		if q.Pages[i1].GeneratorFuncName != "" {
			err := funcPGs[q.Pages[i1].GeneratorFuncName](q, q.Pages[i1])
			dynPagesCreated = true
			if err != nil {
				s := fmt.Sprintf("Page %v; GeneratorFuncName %v returned error %v",
					i1,
					q.Pages[i1].GeneratorFuncName,
					err,
				)
				log.Print(s)
				return fmt.Errorf(s)
			}
			// log.Printf("dyn page#%02v - generator %q created %v groups", i1, q.Pages[i1].GeneratorFuncName, len(q.Pages[i1].Groups))
		}
	}
	if dynPagesCreated {
		// again
		q.Hyphenize()
		q.ComputeMaxGroups()
		q.SetColspans()
		// but not q.Validate
	}

	return nil

}
