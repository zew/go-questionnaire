package qst

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"github.com/zew/go-questionnaire/cfg"
	"github.com/zew/go-questionnaire/css"
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

// ShortSuffix appends suffix to the input - using &nbsp;
// for longer text, reverse order of label and control;
// sadly, &nbsp; has no effect;
// we have to use white-space: nowrap; on the grid-item
func (inp *inputT) ShortSuffix(ctrl string, langCode string) string {

	if inp.Suffix.Empty() {
		return ctrl
	}

	ctrl = strings.Trim(ctrl, "\r\n\t ")
	ctrl = fmt.Sprintf("%v%v", ctrl, inp.Suffix.TrSilent(langCode)) // &nbsp; is broken anyway

	return ctrl
}

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

// ControlTop puts the control vertically on top;
// default would be vertically centered
func (inp *inputT) ControlTop() {
	inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
	inp.StyleCtl.Desktop.StyleGridItem.AlignSelf = "start"
}

// LabelRight aligns the label right;
// but also the text right
func (inp *inputT) LabelRight() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleGridItem.JustifySelf = "end"
	inp.StyleLbl.Desktop.StyleText.AlignHorizontal = "right"
}

// LabelPadRight puts a padding right on the label
// to prevent touching of the control
func (inp *inputT) LabelPadRight() {
	inp.StyleLbl = css.NewStylesResponsive(inp.StyleLbl)
	inp.StyleLbl.Desktop.StyleBox.Padding = "0 1.0rem 0 0"
}

// appendTooltip appends an explanation
func (inp *inputT) appendTooltip(w io.Writer, langCode string) {

	if inp.Tooltip.Empty() {
		return
	}
	fmt.Fprintf(w, "<span class='question-mark-tooltip' title='%v'>&#10068;</span>", inp.Tooltip.TrSilent(langCode))

}

//
//
// GroupHTMLGridBased renders a group of inputs to GroupHTMLGridBased
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

	if gr.Style.Desktop.StyleGridContainer.GapColumn == "" {
		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.4rem"
	}
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
	divWrap(wContainer, gridContainerClass+" grid-container", "", wInner.String())

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
		// i.DD = &DropdownT{}
		inp.DD.SetName(inp.Name)
		inp.DD.LC = q.LangCode
		inp.DD.SetTitle(inp.Label.TrSilent(q.LangCode) + " " + inp.Desc.TrSilent(q.LangCode))
		inp.DD.Select(inp.Response)
		// inp.DD.SetAttr("class", inp.CSSControl)
		sort.Sort(inp.DD)

		ctrl += inp.DD.RenderStr()

	case "label-as-input":
		if !inp.Label.Empty() {
			ctrl += fmt.Sprintf("<span data='label-as-input'>%v</span> ", inp.Label.Tr(q.LangCode))
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

		ctrl += fmt.Sprintf(
			`<input type='%v'  %v  
				name='%v' id='%v' title='%v %v' 
				style='%v'  
				size='%v' maxlength=%v min='%v' max='%v' %v %v %v value='%v' %v />
			`,
			inp.Type, stepping,
			nm, fmt.Sprintf("%v%v", nm, inp.ValueRadio), inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
			width,
			inp.MaxChars, inp.MaxChars, inp.Min, inp.Max,
			placeHolder, onInvalid, autocomplete,
			rspvl, checked,
		)

		// the checkbox "empty catcher" must follow *after* the actual checkbox input,
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
	ctrl = inp.ShortSuffix(ctrl, q.LangCode)

	// error rendering moved to GroupHTMLGridBased

	return ctrl

}
