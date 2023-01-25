package qst

import (
	"fmt"
	"io"
	"log"
	"math"
	"path"
	"strconv"
	"strings"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func wrap(w io.Writer, tagName, forVal, className, style, content string) {

	forKeyVal := ""
	if forVal != "" {
		forKeyVal = fmt.Sprintf("for='%v'", forVal)

	}

	if style != "" {
		styles := strings.Split(style, ";")
		style = strings.Join(styles, ";\n\t\t")
		style = fmt.Sprintf("%v%v%v", "\n\t\t", style, "\n\t\t")
	}

	fmt.Fprintf(w,
		"<%v %v class='%v' style='%v' >\n%v\n</%v>\n",
		tagName,
		forKeyVal,
		className,
		style,
		content,
		tagName,
	)

	// fmt.Fprintf(w, "<!-- /%v -->\n", className)

}
func divWrap(w io.Writer, className, style, content string) {
	wrap(w, "div", "", className, style, content)
}

func (inp inputT) labelDescription(w io.Writer, langCode string) {

	if inp.Label.Empty() && inp.Desc.Empty() {
		return
	}

	if inp.Type == "label-as-input" {
		return
	}

	if inp.IsHidden() {
		return
	}

	if inp.IsLayout() {
		if !inp.Label.Empty() {
			fmt.Fprint(w, inp.Label.Tr(langCode))
		}
		if !inp.Desc.Empty() {
			fmt.Fprint(w, inp.Desc.Tr(langCode))
		}
		return
	}

	// classes are only for font-size, font-weight
	// inline-block styles are applied in outer wrapper
	if !inp.Label.Empty() {
		// fmt.Fprintf(w, " <span class='input-label-text'       >%v</span>", inp.Label.Tr(langCode))
		fmt.Fprintf(w, "%v", inp.Label.Tr(langCode))
	}
	if !inp.Desc.Empty() {
		fmt.Fprintf(w, " <span class='input-description-text' >%v</span>", inp.Desc.Tr(langCode))
	}

}

// shortSuffix appends suffix to the input - using &nbsp;
// for longer text, reverse order of label and control;
// sadly, &nbsp; has no effect;
// we have to use white-space: nowrap; on the grid-item
func (inp *inputT) shortSuffix(ctrl string, langCode string) string {

	if inp.Suffix.Empty() {
		return ctrl
	}

	ctrl = strings.Trim(ctrl, "\r\n\t ")
	ctrl = fmt.Sprintf("%v%v", ctrl, inp.Suffix.TrSilent(langCode)) // &nbsp; is broken anyway

	return ctrl
}

/*
	Methods for input style; overall, label and control

	There are *generic* style methods
		css.ItemStartCA()
		css.TextCenter()
*/

// ControlFirst puts label behind input element;
// for radio and checkbox inputs;
// using CSS grid styles
func (inp *inputT) ControlFirst() {

	inp.StyleLbl = css.ItemEndMA(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "start"
	inp.StyleLbl.Desktop.StyleGridItem.Order = 2

	if inp.ColSpanControl == 0 && inp.ColSpanLabel == 0 {
		inp.ColSpanControl = 8
		inp.ColSpanLabel = 1
	}
	// inp.ColSpanControl, inp.ColSpanLabel = inp.ColSpanLabel, inp.ColSpanControl
}

// ControlTop puts the control vertically at top;
// default would be vertically centered
func (inp *inputT) ControlTop() {
	inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
	inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "start"
}

// ControlTopNudge adds a small nudge to ControlTop
func (inp *inputT) ControlTopNudge() {
	inp.ControlTop()
	// a tiny nudge
	inp.StyleCtl.Desktop.StyleBox.Margin = "0.25rem 0 0 0"
}

// ControlBottom puts the control vertically at bottom;
// default would be vertically centered
func (inp *inputT) ControlBottom() {
	inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
	inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "end"
}

// LabelRight aligns the label right;
// but also the text right;
// different from ControlFirst()
func (inp *inputT) LabelRight() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "end"
	inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "right"
}

// LabelCenter - label *and* text horizontally centered;
// default is 'start/left';
// vertical remains 'center'
func (inp *inputT) LabelCenter() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "center"
	inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "center"
}

// LabelBottom - vertically bottomed;
// default is 'center'
func (inp *inputT) LabelBottom() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.AlignSelf = "end"
}

// LabelTop - vertically top;
// default is 'center'
func (inp *inputT) LabelTop() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.AlignSelf = "start"
}

// ControlCenter - input horizontally centered;
// default is 'start';
// vertical remains 'center'
func (inp *inputT) ControlCenter() {
	inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
	inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "center"
}

// LabelPadRight puts a padding right on the label
// to prevent touching of the control
func (inp *inputT) LabelPadRight() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
}

// VerticalLabel - ?
func (inp *inputT) VerticalLabel() {
	inp.LabelVertical()
}

// LabelVertical - label box and text vertically centered - horizontally bottomed
func (inp *inputT) LabelVertical() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "center"
	inp.StyleLbl.Desktop.StyleGridItem.AlignSelf = "end"
	inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "center"
}

// Vertical order of control and label instead of horizontal
func (inp *inputT) Vertical() {
	inp.Style = css.NewStylesResponsive(inp.Style)
	inp.Style.Desktop.StyleGridContainer.AutoFlow = "column"
	rows := 1
	inp.Style.Desktop.StyleGridContainer.TemplateRows = strings.Repeat("1fr ", rows)
	inp.Style.Desktop.StyleGridContainer.TemplateColumns = " " // empty string
}

// DisplayNone setting the input div to display:none.
// Use case: Inputs be shown or hidden based on Javascript.
// Notice that input wrapper div has a dynamic class with display=grid.
func (inp *inputT) DisplayNone() {
	inp.Style = css.NewStylesResponsive(inp.Style)
	inp.Style.Desktop.StyleBox.Display = "none"
}

// appendTooltip appends an explanation
func (inp *inputT) appendTooltip(w io.Writer, langCode string) {

	if inp.Tooltip.Empty() {
		return
	}
	fmt.Fprintf(w, "<span class='question-mark-tooltip' title='%v'>&#10068;</span>", inp.Tooltip.TrSilent(langCode))

}

// GroupHTMLGridBased renders a group of inputs to grid based HTML
func (q QuestionnaireT) GroupHTMLGridBased(pageIdx, grpIdx int) string {

	wCSS := &strings.Builder{}
	gr := q.Pages[pageIdx].Groups[grpIdx]

	//
	//
	// compare (gr *groupT) Vertical
	gr.Style = css.NewStylesResponsive(gr.Style)
	if gr.Style.Desktop.StyleBox.Display == "" {
		gr.Style.Desktop.StyleBox.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.AutoFlow = "row"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = strings.Repeat("1fr ", int(gr.Cols))
	}

	// if gr.Style.Desktop.StyleGridContainer.GapColumn == "" {
	// // this skews column width when nesting grids
	// // => disabled 2022-06
	// gr.Style.Desktop.StyleGridContainer.GapColumn = "0.4rem"
	// }
	if gr.Style.Desktop.StyleGridContainer.GapRow == "" {
		gr.Style.Desktop.StyleGridContainer.GapRow = "0.8rem"
	}
	gridContainerClass := fmt.Sprintf("pg%02v-grp%02v", pageIdx, grpIdx)
	fmt.Fprint(wCSS, gr.Style.CSS(gridContainerClass))

	//
	wInner := &strings.Builder{} // inside the group grid container
	for inpIdx, inp := range gr.Inputs {
		if inp.Type == "dyn-composite-scalar" {
			continue
		}
		if inp.Type == "dyn-composite" {
			continue
		}

		// no grid cells
		if inp.IsHidden() {
			fmt.Fprint(wInner, q.InputHTMLGrid(pageIdx, grpIdx, inpIdx, q.LangCode))
			continue
		}

		inp.Style = css.NewStylesResponsive(inp.Style)

		//
		// 1.) input div is item      to group
		if inp.ColSpan == 0 {
			// CSS Col property effectively defaults to 1
		} else {
			inp.Style.Desktop.StyleGridItem.Col = fmt.Sprintf("auto / span %v", inp.ColSpan)
		}

		//
		// 2.) input div is container to label and control
		if inp.Style.Desktop.StyleBox.Display == "" {
			inp.Style.Desktop.StyleBox.Display = "grid"
		}

		// flow / main axis =>  row
		if inp.Style.Desktop.StyleGridContainer.AutoFlow == "" {
			inp.Style.Desktop.StyleGridContainer.AutoFlow = "row"
			if inp.ColSpanLabel > 0.2 && inp.ColSpanControl > 0.2 {
				inp.Style.Desktop.StyleGridContainer.TemplateColumns =
					fmt.Sprintf("%4.1ffr  %4.1ffr", inp.ColSpanLabel, inp.ColSpanControl)
			} else {
				inp.Style.Desktop.StyleGridContainer.TemplateColumns = "1fr"
			}
		}

		// flow / main axis =>  column
		if inp.Style.Desktop.StyleGridContainer.AutoFlow == "column" {
			inp.Style.Desktop.StyleGridContainer.TemplateColumns = "none"   // unset
			inp.Style.Desktop.StyleGridContainer.TemplateRows = "0.9fr 1fr" // must be more than one row in order to work
		}

		gridItemClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v", pageIdx, grpIdx, inpIdx)
		fmt.Fprint(wCSS, inp.Style.CSS(gridItemClass))

		wInp := &strings.Builder{} // label and control of input

		if inp.ErrMsg != "" {
			// log.Printf("Input %v with value %v has server validation error %v", inp.Name, inp.Response, inp.ErrMsg)
			oldStyle := false
			if oldStyle {
				stl := fmt.Sprintf("grid-column: auto / span %v", inp.ColSpan)
				divWrap(wInp, " error error-block-input", stl, inp.ErrMsg)
			} else {
				wErr := &strings.Builder{}
				divWrap(wErr, " popup-invalid-content-grid-item   error   error-block-input ", "", inp.ErrMsg)
				divWrap(wInp, " popup-invalid-anchor-grid-item", "", wErr.String())
			}
		}

		{
			if inp.ColSpanLabel > 0.2 {
				wLbl := &strings.Builder{}

				inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)

				if inp.StyleLbl.Desktop.StyleGridItem.AlignSelf == "" {
					inp.StyleLbl.Desktop.StyleGridItem.AlignSelf = "center"
				}

				// styleLbl.Desktop.GridItemStyle.Order = 2  // put label behind control

				lblClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v-lbl", pageIdx, grpIdx, inpIdx)
				fmt.Fprint(wCSS, inp.StyleLbl.CSS(lblClass))

				inp.labelDescription(wLbl, q.LangCode)
				inp.appendTooltip(wLbl, q.LangCode)

				if inp.IsLayout() {
					divWrap(wInp, lblClass+" grid-item-lvl-2", "", wLbl.String())
				} else {
					wrap(wInp, "label", inp.Name, lblClass+" grid-item-lvl-2", "", wLbl.String())
				}
			}

			if inp.ColSpanControl > 0 {
				wCtl := &strings.Builder{}
				if inp.StyleCtl == nil {
					inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
					inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "center"
					if !inp.IsLayout() {
						inp.StyleCtl.Desktop.StyleText.WhiteSpace = "nowrap" // prevent suffix from being wrapped
					}
				}
				if inp.Type == "radio" || inp.Type == "checkbox" {
					inp.StyleCtl = css.ItemCenteredMCA(inp.StyleCtl)
				}

				ctlClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v-ctl", pageIdx, grpIdx, inpIdx)
				fmt.Fprint(wCSS, inp.StyleCtl.CSS(ctlClass))

				fmt.Fprint(wCtl, q.InputHTMLGrid(pageIdx, grpIdx, inpIdx, q.LangCode))
				divWrap(wInp, ctlClass+" grid-item-lvl-2", "", wCtl.String())

			}

		}

		//
		divWrap(wInner, gridItemClass+" grid-item-lvl-1", "", wInp.String())
	}

	//
	//
	wContainer := &strings.Builder{}
	divWrap(wContainer, gridContainerClass+" grid-container "+gr.Class, "", wInner.String())

	w := &strings.Builder{}
	fmt.Fprint(w, css.StyleTag(wCSS.String()))
	fmt.Fprint(w, wContainer.String())
	return w.String()

}

// InputHTMLGrid renders an input to HTML
func (q QuestionnaireT) InputHTMLGrid(pageIdx, grpIdx, inpIdx int, langCode string) string {

	// gr := q.Pages[pageIdx].Groups[grpIdx]
	inp := *q.Pages[pageIdx].Groups[grpIdx].Inputs[inpIdx]
	nm := inp.Name
	ctrl := ""

	switch inp.Type {

	case "textblock":
		// no op

	case "javascript-block":
		sb := &strings.Builder{}
		q.RenderJS(
			sb,
			path.Join(q.Survey.Type, inp.Name),
			inp.JSBlockTrls,
			inp.JSBlockStrings,
		)
		ctrl = sb.String()

	case "button":
		ctrl = fmt.Sprintf(
			"<button type='submit' name='%v' value='%v' accesskey='%v'><b>%v</b> %v</button>\n",
			inp.Name, inp.Response, inp.AccessKey,
			inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
		)

	case "textarea":
		width := ""
		colsRows := fmt.Sprintf(" cols='%v' rows='1' ", inp.MaxChars+1)
		if inp.MaxChars > 80 {
			colsRows = fmt.Sprintf(" cols='80' rows='%v' ", inp.MaxChars/80+1)
			// width = fmt.Sprintf("width: %vem;", int(float64(80)*1.05))
			width = "width: 98%;"
		}
		ctrl += fmt.Sprintf("<textarea        name='%v' id='%v' title='%v %v' style='%v' maxlength='%v' %v  autocomplete='off' >%v</textarea>\n",
			nm, nm, inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode), width, inp.MaxChars, colsRows, inp.Response)

	case "dropdown":
		inp.DD.SetName(inp.Name)
		inp.DD.LC = q.LangCode
		inp.DD.Select(inp.Response)
		// inp.DD.SetAttr("class", inp.CSSControl)

		ctrl += inp.DD.RenderStr()

	case "label-as-input":
		if !inp.Label.Empty() {
			ctrl += fmt.Sprintf("<span data='label-as-input'>%v</span> ", inp.Label.Tr(q.LangCode))
		}
	case "range":

		// HTML spec demands the default value for a range input to be max-min/2;
		// thus we have _no_ chance to find out about the init state
		//
		// we have two remedies for measuring default

		// remedy 1: A data attribute 'data-dirty';
		// javascript onload-input can check for this
		// 		src.dataset.dirty === "false"
		// unused
		dirty := ""
		if inp.Response == "" {
			dirty = " data-dirty='false' "
		}

		// remedy 2: An "empty catcher" hidden input;
		// with checkboxes, the empty catcher must _succeed_ corresponding input,
		// since http.Form.Get() fetches the first value.

		// in the range case, this is not true; instead it _precedes_ the corresponding input;
		// is it because input-wrapper is display=flex and column-reverse?
		if inp.Response == "" {
			ctrl += fmt.Sprintf(
				`<input type='hidden' name='%v' id='%v_hidd' value='' />`,
				nm, nm,
			)
		}

		var rangeCfg rangeConf
		rangeCfg.New(&inp)

		// range-input wrapper
		// display flex, column-reverse
		ctrl += fmt.Sprintf("<div class='input-wrapper-signature-%v'>", rangeCfg.CSSType)

		classCSS := []string{}
		if inp.Response == "" {
			classCSS = append(classCSS, "hidethumb")
		} else if inp.Response == "no answ." {
			classCSS = append(classCSS, "noanswer", "hidethumb")
		}

		dispVal := ""
		if inp.Response == "" {
			dispVal = " -- " // init
		} else if inp.Response == "no answ." {
			dispVal = "no answ." // also adapt in pdsRange.js display.value
		} else {
			floatFormatter := "%.0f"
			if inp.Step < 1.0 {
				floatFormatter = "%.1f"
			}
			if inp.Step <= 0.5 {
				floatFormatter = "%.2f"
			}
			if inp.Step <= 0.05 {
				floatFormatter = "%.3f"
			}
			// fmStr := fmt.Sprintf("%v - %v", floatFormatter, floatFormatter)
			resp, err := strconv.ParseFloat(inp.Response, 64)
			if err != nil {
				log.Printf("inp %v - cannot convert inp.Repsonse %v to float", inp.Name, inp.Response)
			}
			lower := fmt.Sprintf(floatFormatter, resp)
			upper := fmt.Sprintf(floatFormatter, resp+inp.Step)
			// log.Printf("float formatter for %v, %v \n\t\t%v", resp, resp+inp.Step, fmStr)
			lower = humanizeRangeDisplay(lower)
			upper = humanizeRangeDisplay(upper)
			dispVal = fmt.Sprintf("%v – <%v", lower, upper)

			if resp == rangeCfg.UpperLastRegular {
				dispVal = fmt.Sprintf("%v", lower)
			}

			if resp <= rangeCfg.LowerThreshold && rangeCfg.LowerDisplay != "" {
				dispVal = rangeCfg.LowerDisplay

			}

			if resp >= rangeCfg.UpperThreshold && rangeCfg.UpperDisplay != "" {
				dispVal = rangeCfg.UpperDisplay
			}
		}

		//

		ctrl += fmt.Sprintf(
			`<input type='%v'  
				name='%v' id='%v' 
				min='%v' max='%v' step='%v' 
				value='%v'
				oninput='pdsRangeInput(this)' 
				onclick='pdsRangeClick(this)' 
				onfocus='pdsRangeClick(this)' 
				ontouchstart='pdsRangeClick(this)' 
				%v
				class='%v'
				%v
			/>
			`,
			inp.Type,
			nm, nm,
			inp.Min, inp.Max, inp.Step,
			inp.Response,
			dirty,
			strings.Join(classCSS, " "),
			rangeCfg.lowerUpperAttrs(),
		)

		// '500 - 510' is the max content, defined by CSS, approx 3.9rem
		maxWidthDisplay := ""
		if inp.Max < 100 && inp.Step >= 1 {
			maxWidthDisplay = ` style="max-width:3.5rem" `
		}

		// value indicator [rangename]_display
		display := `<input 
			type='text' 
			name='%v_display' 
			id='%v_display' 
			class='range-display'
			size='%v' 
			maxlength='%v' 
			%v
			value='%v'
			disabled='true'
		>` // no space at the end - it follows the [unit]

		display = fmt.Sprintf(
			display,
			inp.Name,
			inp.Name,
			inp.MaxChars,
			inp.MaxChars,
			maxWidthDisplay,
			dispVal,
		)

		checked := ""
		if inp.Response == "no answ." {
			checked = " checked=true "
		}

		// no answer radio [rangename]_noanswer
		noAnswer := `<input 
			type='radio' 
			name='%v_noanswer' 
			id='%v_noanswer' 
			class='range-noanswer'
			value='true'
			%v
			title='no answer'
			oninput='pdsRangeRadioInput(this)' 
			tabindex='-1'
		>`
		noAnswer = fmt.Sprintf(
			noAnswer,
			inp.Name,
			inp.Name,
			checked,
		)

		// we render a second label, serves as anchor
		// a.) for the display
		// b.) for the ticks
		//
		// the label containing the description text is in another grid item above
		//
		// this is a _trailing_ sibling to the input[range];
		// display:flex and column-reverse design will render it above
		//
		//
		//
		ctrl += fmt.Sprintf(`
			<label for="%v" >
				%v%v
				<div class="labels" aria-hidden="true" 
				>%v</div>
			</label>	
			`,
			inp.Name,
			display, inp.Suffix[q.LangCode],
			inp.rangeLabels(rangeCfg),
		)

		ctrl += noAnswer

		ctrl += `</div>` // /input-wrapper

		inp.Suffix = trl.S{} // delete - since range writes its own suffix

		if false {
			// render JS;
			// we fire an input event at the range-input on window load
			sb := &strings.Builder{}
			inp.JSBlockStrings = map[string]string{
				"inputName": inp.Name,
			}
			q.RenderJS(
				sb,
				path.Join(q.Survey.Type, "rangeAuto"),
				inp.JSBlockTrls,
				inp.JSBlockStrings,
			)
			ctrl += sb.String()
		}

	case "text", "number", "hidden", "checkbox", "radio":
		rspvl := inp.Response

		checked := ""
		if inp.Type == "checkbox" {
			if rspvl == ValSet {
				checked = "checked=\"checked\""
			}
			rspvl = ValSet
		}
		if inp.Type == "radio" {
			if rspvl == inp.ValueRadio {
				checked = "checked=\"checked\""
			}
			rspvl = inp.ValueRadio
		}

		width := fmt.Sprintf("width:%.2frem", float32(inp.MaxChars)*0.65)
		if inp.Type == "checkbox" || inp.Type == "radio" {
			width = ""
		}

		stepping := ""
		excelFormat := ""
		if inp.Type == "number" {
			if inp.Step != 0 {
				if inp.Step >= 1 {
					stepping = fmt.Sprintf(" step='%.0f'  ", inp.Step)
					excelFormat = strings.Repeat("#", int(inp.Step))
				} else {
					prec := int(math.Log10(1 / inp.Step))
					f := fmt.Sprintf(" step='%%.%vf'  ", prec)
					stepping = fmt.Sprintf(f, inp.Step)
					excelFormat = "#." + strings.Repeat("0", prec)
				}
			}
		}

		placeHolder := ""
		if inp.Placeholder.TrSilent(langCode) != "" {
			placeHolder = fmt.Sprintf("placeholder='%v'", inp.Placeholder.TrSilent(langCode))
		}
		if placeHolder == "" && inp.Type == "number" {
			placeHolder = fmt.Sprintf("placeholder='%v'", excelFormat)
		}

		//
		// stackoverflow.com/questions/19122886
		// raise invalid message - clear it
		onInvalid := ""
		if inp.OnInvalid.TrSilent(langCode) != "" {
			// onInvalid = fmt.Sprintf("oninvalid='setCustomValidity(\"%v\")' oninput='setCustomValidity(\"\")'", inp.OnInvalid.TrSilent(langCode))
			onInvalid = fmt.Sprintf("data-validation_msg='%v'", inp.OnInvalid.TrSilent(langCode))
		} else {
			if inp.Type == "number" {
				if inp.Min != 0 || inp.Max != 0 {
					txt := fmt.Sprintf(cfg.Get().Mp["entry_range"].TrSilent(langCode), inp.Min, inp.Max)
					if inp.Step != 0 && inp.Step != 1 {
						txt += "; "
						txt += fmt.Sprintf(cfg.Get().Mp["entry_stepping"].TrSilent(langCode), inp.Step)
					}
					// onInvalid = fmt.Sprintf("oninvalid='setCustomValidity(\"%v\")' oninput='setCustomValidity(\"\")'", txt)
					onInvalid = fmt.Sprintf("data-validation_msg='%v'", txt)
				}
			}
		}

		autocomplete := ""
		if inp.Type == "text" {
			autocomplete = "autocomplete='off'"
		}

		disabled := " "
		if inp.Disabled {
			disabled = "disabled=true "

		}

		// 2022-12: rendering the title as tooltip dropped - too long - too many dangerous characters
		// 		inside the input template:   title='%v %v'
		// 		placeholder values:          inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),

		// 2023-01: for PCAG number input: lang='de_DE' does not help
		ctrl += fmt.Sprintf(
			`<input type='%v'  %v  
				name='%v' id='%v' 
				style='%v'  
				size='%v' maxlength=%v min='%v' max='%v' %v %v %v value='%v' %v %v />
			`,
			inp.Type, stepping,
			nm, fmt.Sprintf("%v%v", nm, inp.ValueRadio),
			width,
			inp.MaxChars, inp.MaxChars, inp.Min, inp.Max,
			placeHolder, onInvalid, autocomplete,
			rspvl, checked, disabled,
		)

		// the checkbox "empty catcher" must follow _after_ the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if inp.Type == "checkbox" {
			ctrl += fmt.Sprintf(
				"<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

	case "dyn-textblock":
		ctrl = fmt.Sprintf("<span>%v</span>\n", inp.Label.Tr(q.LangCode))

	case "dyn-composite", "dyn-composite-scalar":
		// no op
		// rendered at group level -  rendered by composite

	default:
		ctrl = fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, inp.Type, implementedTypes)
	}

	//
	// common

	// append suffix
	ctrl = inp.shortSuffix(ctrl, q.LangCode)

	// error rendering moved to GroupHTMLGridBased

	return ctrl

}

/*
rangeLabels renders HTML for ticks and tick-labels;
based on an encoding in DynamicFuncParamset.

[CSS-class -- [step:label];[step:label];...]
*/
func (inp *inputT) rangeLabels(rc rangeConf) string {

	// if strings.Contains(inp.DynamicFuncParamset, "1.5") {
	// 	log.Printf("   tickS %v", inp.DynamicFuncParamset)
	// 	log.Printf("     xs %+v", xs)
	// 	log.Printf("   lbls %+v", lbls)
	// }

	// prelimin
	ws1 := []float64{} // widths delta, first element 0, adding up to 1.000
	{
		xsPrelim := make([]float64, len(rc.xs))
		copy(xsPrelim, rc.xs)

		ws0 := []float64{} // widths based on zero
		ctr := 0.0
		for stp := inp.Min; stp <= inp.Max; stp += inp.Step {
			// correcting the cumulative rounding errors of x += inp.Step - instead using _one_ multiplication
			stp = inp.Min + ctr*inp.Step
			ctr++

			// check if current step x has a label assigned xs[0]
			if len(xsPrelim) > 0 && stp == xsPrelim[0] {
				xsPrelim = xsPrelim[1:] // chop off leading ticks
				ws0 = append(ws0, stp-inp.Min)
			}
		}
		// if strings.Contains(inp.DynamicFuncParamset, "1.5") {
		// 	log.Printf("   ws0  %+v", ws0)
		// }

		for i := 0; i < len(ws0); i++ {
			ws0[i] = ws0[i] / (inp.Max - inp.Min)
		}
		// log.Printf("   w2: %+v", ws0)

		// deltas
		// we need a new slice, because looking back one step
		// otherwise points to values delta-ficated
		sum := 0.0
		for i := 0; i < len(ws0); i++ {
			delta := ws0[i]
			if i > 0 {
				delta = ws0[i] - ws0[i-1]
			}
			ws1 = append(ws1, delta)
			sum += delta
		}
		// log.Printf("   w3: %+v - sum %v", ws1, sum)

		ws1 = ws1[1:] // chop off the leading zero

	}

	// if strings.Contains(inp.DynamicFuncParamset, "1.5") {
	// 	log.Printf("   ws1  %+v", ws1)
	// }

	/*
		core is the html string builder
		<span class=lbl-anchor>
			  <span class=lbl>optional label</span>
		</span>
		<span class=tick  width=11.1% >distance block - with left|right border making the tick</span>

	*/
	core := &strings.Builder{}

	itr1 := -2
	itr2 := -1

	ctr := 0.0
	noTick := 0

	for stp := inp.Min; stp <= inp.Max; stp += inp.Step {

		// correcting the cumulative rounding errors of x += inp.Step - instead using _one_ multiplication
		stp = inp.Min + ctr*inp.Step
		ctr++

		// check if current step should have a tick in xs[0]
		if len(rc.xs) > 0 && stp == rc.xs[0] {

			itr1++
			itr1++

			itr2++

			lbl := rc.lbls[0]
			if strings.HasSuffix(lbl, "nt") {
				lbl = lbl[:len(lbl)-2]
				noTick = 1
			} else {
				noTick = 0
			}
			// label - width 0
			fmt.Fprintf(
				core,
				`<span class='lbl-anchor'><span class='lbl'>%v</span></span>`,
				lbl,
			)

			// a width block with class 'tick'
			// 'first' has borders left+right; all others have border-right
			if len(ws1) > 0 {

				pct := 100 * ws1[0]
				pct = math.Floor(pct*100) / 100 // round *down* two digits after decimal separator

				tickClass := " class='tick' "
				if itr2 == 0 {
					tickClass = " class='tick first' "
				}
				if noTick == 1 {
					tickClass = " class='tick suppress-right' "
					if itr2 == 0 {
						tickClass = " class='tick first suppress-left' "
					}
				}

				fmt.Fprintf(
					core,
					`<span %v style="width: %v%%" ></span>`,
					tickClass,
					fmt.Sprintf("%6.2f", pct), // percentage width
				)

				rc.xs = rc.xs[1:]     // chop off leading ticks
				rc.lbls = rc.lbls[1:] // ... and labels
				ws1 = ws1[1:]
			}

		}

	}

	// log.Printf("process end %v", inp.DynamicFuncParamset)

	return core.String()

}

// for range inputs, we have a display panel;
// we want to chop off decimal digits like follows
//
//	1.25 - 1.5
//	1.5  - 1.75
//	1.75 - 2
func humanizeRangeDisplay(s string) string {

	if !strings.Contains(s, ".") {
		return s
	}

	for len(s) > 0 {
		// fmt.Print(s[len(s)-1:], "\n")
		if s[len(s)-1:] == "0" {
			s = s[:len(s)-1]
		} else if s[len(s)-1:] == "." {
			s = s[:len(s)-1]
			break
		} else {
			break
		}
	}
	return s
}
