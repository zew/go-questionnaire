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

	"github.com/pkg/errors"
	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/cloudio"
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
	Name  string `json:"name,omitempty"`
	Type  string `json:"type,omitempty"`  // see implementedTypes
	Param string `json:"param,omitempty"` // for dyn-text - name of parameter set

	MaxChars    int     `json:"max_chars,omitempty"`  // input chars; => SIZE for input, MAXLENGTH for textarea, text; also used for width
	Step        float64 `json:"step,omitempty"`       // for number input:  stepping interval, i.e. 2 or 0.1
	Min         float64 `json:"min,omitempty"`        //      ~
	Max         float64 `json:"max,omitempty"`        //      ~
	OnInvalid   trl.S   `json:"on_invalid,omitempty"` // message for javascript error messages on HTML5 invalid state - compare ErrMsg
	Placeholder trl.S   `json:"placeholder,omitempty"`

	Label     trl.S  `json:"label,omitempty"`
	Desc      trl.S  `json:"description,omitempty"`
	Suffix    trl.S  `json:"suffix,omitempty"` // only for short units - such as € or % - for longer text use label.Style...Order = 2
	Tooltip   trl.S  `json:"tooltip,omitempty"`
	AccessKey string `json:"accesskey,omitempty"`

	/*Colspan determines, how many column slots of the group column layout
	the input occupies.
	Default value is assumed to be 1.
	Increase it manually in the generator function.
	ColSpanLabel and ColSpanControl do *not* influence Colspan.
	*/
	ColSpan float32 `json:"col_span,omitempty"`

	// ColSpanLabel/-Control work only as proportion of Colspan
	ColSpanLabel   float32 `json:"col_span_label,omitempty"`
	ColSpanControl float32 `json:"col_span_control,omitempty"`

	Radios []*radioT  `json:"radios,omitempty"`    // This slice implements the radiogroup - and the senseless checkboxgroup
	DD     *DropdownT `json:"drop_down,omitempty"` // As pointer to prevent JSON cluttering

	Validator string `json:"validator,omitempty"` // i.e. any key from map of validators, i.e. "must;inRange20"
	ErrMsg    string `json:"err_msg,omitempty"`   // a key for coreTranslations, compare OnInvalid

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
	ID             string `json:"-"`                         // is only a helper, but must be exported in order to be stored in session
	BottomVSpacers int    `json:"bottom_vspacers,omitempty"` // number of rows below the group, addGroup() initializes to 3

	// Number of vertical columns;
	// for horizontal *and* (not yet implemented) vertical layouts;
	//
	// Each label (if set) and each input occupy columns according to
	// inputT.ColSpanLabel and inputT.ColSpanControl.
	Cols            float32 `json:"columns,omitempty"`
	OddRowsColoring bool    `json:"odd_rows_coloring,omitempty"` // color odd rows

	Inputs []*inputT `json:"inputs,omitempty"`

	// > 0 => group belongs to a set of groups,
	// reordering / re-shuffled according to questionnaireT.ShufflingsMax
	RandomizationGroup int `json:"randomization_group,omitempty"`

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
	pageIdx, grpIdx int, compFuncNameWithParamSet string) (CompositeFuncT, int, int) {

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

	Style *css.StylesResponsive `json:"style,omitempty"`

	// *not* a marker for questionnaire finished entirely;
	// see q.ClosingTime instead;
	// truncated to second
	Finished time.Time `json:"finished,omitempty"`

	Groups []*groupT `json:"groups,omitempty"`

	ValidationFuncName string `json:"validation_func_name,omitempty"` // javascript validation func name
	ValidationFuncMsg  trl.S  `json:"validation_func_msg,omitempty"`
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
	Survey surveyT           `json:"survey,omitempty"`
	UserID string            `json:"user_id,omitempty"`    // participant ID, decimal, but string, i.E. 1011
	Attrs  map[string]string `json:"user_attrs,omitempty"` // i.e. user country or euro-member - taken from lgn.LoginT
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

	// primitive permutation mechanism
	// deterministically reordering / reshuffling a set of groups
	// based on
	//      * user id
	// 		* page idx
	// 		* groupT.RandomizationGroup
	ShufflingsMax int `json:"shufflings_max,omitempty"`

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
func (q *QuestionnaireT) EditPage(idx int) *pageT {
	if idx < 0 || idx > len(q.Pages)-1 {
		log.Panicf("EditPage(): %v pages - valid indexes are 0...%v", len(q.Pages), len(q.Pages)-1)
	}
	return q.Pages[idx]
}

/*
// RemoveGroup from page
func (q *QuestionnaireT) RemoveGroup(pageIdx, groupIdx int) {

	if pageIdx < 0 || pageIdx > len(q.Pages)-1 {
		log.Panicf(
			"AddPageAfter(): %v pages - valid indexes are 0...%v",
			len(q.Pages), len(q.Pages)-1)
	}

	if groupIdx < 0 || groupIdx > len(q.Pages[pageIdx].Groups)-1 {
		log.Panicf(
			"AddPageAfter(): %v pages - valid indexes are 0...%v",
			len(q.Pages[pageIdx].Groups),
			len(q.Pages[pageIdx].Groups)-1,
		)
	}

	copy(
		q.Pages[pageIdx].Groups[groupIdx+0:],
		q.Pages[pageIdx].Groups[groupIdx+1:],
	) // shift one slot to the left

	q.Pages[pageIdx].Groups = q.Pages[pageIdx].Groups[:len(q.Pages[pageIdx].Groups)-1]

}
*/

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

var debugShuffling = false

// RandomizeOrder creates a shuffled ordering of groups marked by .RandomizationGroup;
// static groups with RandomizationGroup==0 remain on fixed order position;
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

	if debugShuffling {
		log.Printf(
			"max sg idx %v \nshufflingGroups %v",
			maxSg,
			util.IndentedDump(shufflingGroups),
		)
	}

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
				// q.ShufflingsMax instead of len(shufflingGroups[sg])
				// order = order[0:len(shufflingGroups[sg])]
				sh := shuffler.New(q.UserID, q.ShufflingsMax, len(shufflingGroups[sg]))
				order := sh.Slice(pageIdx) // cannot add sg to conform to ShufflesToCSV()
				if debugShuffling {
					log.Printf("%v - seq %16s in order %16s - iter %v", sg, fmt.Sprint(shufflingGroups[sg]), fmt.Sprint(order), pageIdx+sg)
				}
				for i := 0; i < len(shufflingGroups[sg]); i++ {
					offset := shufflingGroups[sg][i] // i.e. [1, 9]
					i2 := order[i]
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

// PageHTML generates HTML for a specific page of the questionnaire
func (q *QuestionnaireT) PageHTML(pageIdx int) (string, error) {

	if q.CurrPage > len(q.Pages)-1 || q.CurrPage < 0 {
		s := fmt.Sprintf("You requested page %v out of %v. Page does not exist", pageIdx, len(q.Pages)-1)
		log.Printf(s)
		return s, fmt.Errorf(s)
	}

	page := q.Pages[pageIdx]

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

	page.Style = css.PageMarginsAuto(page.Style)
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

	compositCntr := -1
	nonCompositCntr := -1
	for loopIdx, grpIdx := range grpOrder {
		if page.Groups[grpIdx].HasComposit() {
			compositCntr++
			compFuncNameWithParamSet := page.Groups[grpIdx].Inputs[0].DynamicFunc
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
			grpHTML := q.GroupHTMLGridBased(pageIdx, grpIdx)

			if strings.Contains(grpHTML, "[groupID]") {
				nonCompositCntr++
				grpHTML = strings.Replace(grpHTML, "[groupID]", fmt.Sprintf("%v", nonCompositCntr+1), -1)
			}
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

		tplFile := page.ValidationFuncName + ".js"

		mp := map[string]string{}
		mp["msg"] = page.ValidationFuncMsg.Tr(q.LangCode)
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
	"\r\n", " - ",
	"\n", " - ",
	":", " - ",
	";", " - ",
	// ",", " - ",
)

// no comma, no colon, no semicolon - only dot or hyphen
// no newline
var englishTextAndNumbersOnly = regexp.MustCompile(`[^a-zA-Z0-9\.\_\- ]+`)
var severalSpaces = regexp.MustCompile(`[ ]+`)

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

// EnglishTextAndNumbersOnly replaces all other UTF characters by space
func EnglishTextAndNumbersOnly(s string) string {

	// sBefore := s

	s = germanUmlaute.Replace(s)
	s = separators.Replace(s)

	s = englishTextAndNumbersOnly.ReplaceAllString(s, " ")
	s = severalSpaces.ReplaceAllString(s, " ")
	s = strings.TrimSpace(s)

	// if sBefore != s {
	// 	log.Printf("\n%v\n%v", sBefore, s)
	// }

	return s
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
		return nil, errors.Wrap(err, msg)
	}

	w := &strings.Builder{}
	fmt.Fprintf(w, "<script>\n")
	fmt.Fprintf(w, string(cnts))
	fmt.Fprintf(w, "console.log('JS tpl %v successfully added')", pth)
	fmt.Fprintf(w, "</script>\n")

	base := template.New(tName)
	tDerived, err := base.Parse(w.String())
	if err != nil {
		msg := fmt.Sprintf("parsing failed for %v: %v", pth, err)
		return nil, errors.Wrap(err, msg)
	}

	return tDerived, nil
}
