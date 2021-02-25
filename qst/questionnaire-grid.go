package qst

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"

	"github.com/zew/go-questionnaire/css"
)

func divWrap(w io.Writer, className, style, content string) {

	if style != "" {
		styles := strings.Split(style, ";")
		style = strings.Join(styles, ";\n\t\t")
		style = fmt.Sprintf("%v%v%v", "\n\t\t", style, "\n\t\t")
	}

	fmt.Fprintf(w,
		"<div class='%v' style='%v' >\n%v\n</div> <!-- /%v -->\n",
		className,
		style,
		content,
		className,
	)

}

func (inp inputT) labelDescription(w io.Writer, langCode string) {

	if !inp.Label.Set() && !inp.Desc.Set() {
		return
	}

	// classes are only for font-size, font-weight
	// inline-block styles are applied in outer wrapper
	wInner := &strings.Builder{}
	if !inp.Label.Empty() {
		fmt.Fprintf(wInner, " <span class='input-label-text' >%v</span>", inp.Label.Tr(langCode))
	}
	if !inp.Desc.Empty() {
		fmt.Fprintf(wInner, " <span class='input-description-text' >%v</span>", inp.Desc.Tr(langCode))
	}

	wOuter := &strings.Builder{}
	if !inp.IsLayout() && inp.Name != "" {
		fmt.Fprintf(wOuter, "<label for='%v'>%v</label>\n", inp.Name, wInner.String())
		fmt.Fprint(w, wOuter.String())
	} else {
		fmt.Fprint(w, wInner.String())
	}

}

//
//
// GroupHTMLGrid renders a group of inputs to GroupHTMLGrid
func (q QuestionnaireT) GroupHTMLGrid(pageIdx, grpIdx int) string {

	wCSS := &strings.Builder{}
	gr := q.Pages[pageIdx].Groups[grpIdx]

	//
	//
	if gr.Style == nil {
		gr.Style = css.NewStylesResponsive()
	}
	gr.Style.Desktop.BoxStyle.Display = "grid"
	if gr.Style.Desktop.GridContainerStyle.AutoFlow == "" {
		gr.Style.Desktop.GridContainerStyle.AutoFlow = "row"
	}
	if gr.Style.Desktop.GridContainerStyle.TemplateColumns == "" {
		gr.Style.Desktop.GridContainerStyle.TemplateColumns = strings.Repeat("1fr ", gr.Cols)
	}
	gridContainerClass := fmt.Sprintf("pg%02v-grp%02v", pageIdx, grpIdx)
	fmt.Fprint(wCSS, gr.Style.CSS(gridContainerClass))

	//
	//
	wInner := &strings.Builder{} // inside the container

	for inpIdx, inp := range gr.Inputs {
		if inp.Type == "composit-scalar" {
			continue
		}
		if inp.Type == "composit" {
			continue
		}

		if inp.Style == nil {
			inp.Style = css.NewStylesResponsive()
		}
		if inp.Style.Desktop.GridItemStyle.Col == "" {
			// input wrapper is item      to group
			inp.Style.Desktop.GridItemStyle.Col = fmt.Sprintf("auto / span %v", inp.ColSpanLabel+inp.ColSpanControl)

			// input wrapper is container to label and control
			inp.Style.Desktop.BoxStyle.Display = "grid"
			if inp.Style.Desktop.GridContainerStyle.AutoFlow == "" {
				inp.Style.Desktop.GridContainerStyle.AutoFlow = "row"
			}
			inp.Style.Desktop.GridContainerStyle.TemplateColumns = strings.Repeat("1fr ", inp.ColSpanLabel+inp.ColSpanControl)
		}
		gridItemClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v", pageIdx, grpIdx, inpIdx)
		fmt.Fprint(wCSS, inp.Style.CSS(gridItemClass))

		wInp := &strings.Builder{} // label and control of input

		{

			if inp.ColSpanLabel > 0 {
				wLbl := &strings.Builder{}
				lblStyle := css.NewStylesResponsive()
				lblStyle.Desktop.GridItemStyle.Col = fmt.Sprintf("auto / span %v", inp.ColSpanLabel)
				lblStyle.Desktop.GridItemStyle.JustifySelf = "end"
				lblStyle.Desktop.GridItemStyle.AlignSelf = "center"
				lblClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v-lbl", pageIdx, grpIdx, inpIdx)
				fmt.Fprint(wCSS, lblStyle.CSS(lblClass))
				inp.labelDescription(wLbl, q.LangCode)
				divWrap(wInp, lblClass+" grid-item-lvl-2", "", wLbl.String())
			}

			if inp.ColSpanControl > 0 {
				wCtl := &strings.Builder{}
				ctlStyle := css.NewStylesResponsive()
				ctlStyle.Desktop.GridItemStyle.Col = fmt.Sprintf("auto / span %v", inp.ColSpanControl)
				ctlClass := fmt.Sprintf("pg%02v-grp%02v-inp%02v-ctl", pageIdx, grpIdx, inpIdx)
				fmt.Fprint(wCSS, ctlStyle.CSS(ctlClass))
				fmt.Fprint(wCtl, q.InputHTMLGrid(pageIdx, grpIdx, inpIdx))
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
func (q QuestionnaireT) InputHTMLGrid(pageIdx, grpIdx, inpIdx int) string {

	gr := q.Pages[pageIdx].Groups[grpIdx]
	inp := *q.Pages[pageIdx].Groups[grpIdx].Inputs[inpIdx]
	nm := inp.Name

	switch inp.Type {

	case "textblock":
		return ""

	case "button":
		lbl := fmt.Sprintf("<button type='submit' name='%v' value='%v' class='%v' accesskey='%v'><b>%v</b> %v</button>\n",
			inp.Name, inp.Response, inp.CSSControl, inp.AccessKey,
			inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
		)
		lbl = td(inp.HAlignControl, colWidth(inp.ColSpanControl, gr.Cols), lbl)
		return lbl

	case "text", "number", "hidden", "textarea", "checkbox", "dropdown":

		ctrl := ""
		val := inp.Response

		checked := ""
		if inp.Type == "checkbox" {
			if val == ValSet {
				checked = "checked=\"checked\""
			}
			val = ValSet
		}

		if inp.Type == "textarea" {
			width := ""
			colsRows := fmt.Sprintf(" cols='%v' rows='1' ", inp.MaxChars+1)
			if inp.MaxChars > 80 {
				colsRows = fmt.Sprintf(" cols='80' rows='%v' ", inp.MaxChars/80+1)
				// width = fmt.Sprintf("width: %vem;", int(float64(80)*1.05))
				width = "width: 98%;"
			}
			ctrl += fmt.Sprintf("<textarea        name='%v' id='%v' title='%v %v' class='%v' style='%v' MAXLENGTH='%v' %v>%v</textarea>\n",
				nm, nm, inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode), inp.CSSControl, width, inp.MaxChars, colsRows, val)

		} else if inp.Type == "dropdown" {

			// i.DD = &DropdownT{}
			inp.DD.SetName(inp.Name)
			inp.DD.LC = q.LangCode
			inp.DD.SetTitle(inp.Label.TrSilent(q.LangCode) + " " + inp.Desc.TrSilent(q.LangCode))
			inp.DD.Select(inp.Response)
			inp.DD.SetAttr("class", inp.CSSControl)

			sort.Sort(inp.DD)

			ctrl += inp.DD.RenderStr()

		} else {
			// input
			inputMode := ""
			if inp.Type == "number" {
				if inp.Step != 0 {
					if inp.Step >= 1 {
						inputMode = fmt.Sprintf(" step='%.0f'  ", inp.Step)
					} else {
						prec := int(math.Log10(1 / inp.Step))
						f := fmt.Sprintf(" step='%%.%vf'  ", prec)
						inputMode = fmt.Sprintf(f, inp.Step)
					}
				}
			}
			ctrl += fmt.Sprintf(
				`<input type='%v'  %v  name='%v' id='%v' title='%v %v' 
				class='%v' style='width:%vrem'  SIZE='%v' MAXLENGTH=%v MIN='%v' MAX='%v'  %v  value='%v' />
				`,
				inp.Type, inputMode,
				nm, nm, inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
				inp.CSSControl, fmt.Sprintf("%.2f", float32(inp.MaxChars)*0.65), inp.MaxChars, inp.MaxChars, inp.Min, inp.Max, checked, val)
		}

		// the checkbox "empty catcher" must follow *after* the actual checkbox input,
		// since http.Form.Get() fetches the first value.
		if inp.Type == "checkbox" {
			ctrl += fmt.Sprintf("<input type='hidden' name='%v' id='%v_hidd' value='0' />\n", nm, nm)
		}

		// append suffix
		ctrl = appendSuffix(ctrl, &inp, q.LangCode)

		// append error message
		if inp.ErrMsg.Set() {
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", inp.CSSLabel, inp.ErrMsg.TrSilent(q.LangCode))
		}

		ctrl = td(inp.HAlignControl, colWidth(inp.ColSpanControl, gr.Cols), ctrl)

		return ctrl

	case "radiogroup", "checkboxgroup":
		ctrl := ""
		innerType := "radio"
		if inp.Type == "checkboxgroup" {
			innerType = "checkbox"
		}

		for radIdx, rad := range inp.Radios {
			one := ""
			checked := ""
			if inp.Response == rad.Val {
				checked = "checked=\"checked\""
			}

			radio := fmt.Sprintf(
				// 2021-01 - id must be unique
				"<input type='%v' name='%v' id-disabled='%v' title='%v %v' class='%v' value='%v' %v />\n",
				innerType, nm, nm, inp.Label.TrSilent(q.LangCode), inp.Desc.TrSilent(q.LangCode),
				inp.CSSControl,
				rad.Val, checked,
			)

			lbl := ""
			if !rad.Label.Set() {
				one = radio
			} else {
				if rad.HAlign == HLeft {
					lbl = fmt.Sprintf(
						"<span class=' vert-correct-left-right' >%v</span>\n",
						rad.Label.Tr(q.LangCode),
					)
					one = nobreakGlue(lbl, "&nbsp;", radio)
				}
				if rad.HAlign == HCenter {
					// no i.CSSLabel to prevent left margins/paddings to uncenter
					lbl = fmt.Sprintf(
						"<span class=' vert-correct-center '>%v</span>\n",
						rad.Label.Tr(q.LangCode),
					)
					lbl += vspacer0
					one = lbl + radio
				}

				if rad.HAlign == HRight {
					lbl = fmt.Sprintf(
						"<span class=' vert-correct-left-right'>%v</span>\n",
						rad.Label.Tr(q.LangCode),
					)
					one = nobreakGlue(radio, "&nbsp;", lbl)
				}
			}

			cssNoLeft := inp.CSSLabel
			if rad.HAlign == HCenter {
				cssNoLeft = strings.Replace(inp.CSSLabel, "special-input-left-padding", "", -1)
			}
			one = fmt.Sprintf("<span class='go-quest-label %v'>%v</span>\n", cssNoLeft, one)

			cellAlign := rad.HAlign
			if rad.HAlign == HRight {
				cellAlign = HLeft
			}

			if rad.Cols == 0 {
				one = td(cellAlign, colWidth(1, gr.Cols), one)
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

				if (rad.Col+0)%rad.Cols == rad.Cols-1 || radIdx == len(inp.Radios)-1 {

					ctrl += tableClose
				}
				if (rad.Col+0)%rad.Cols == rad.Cols-1 && radIdx < len(inp.Radios)-1 {
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

		// append suffix
		if inp.Suffix.Set() {
			// compare appendSuffix() forcing no wrap for ordinary inputs
			// ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", i.CSSLabel, i.Suffix.TrSilent(q.LangCode))
			ctrl += fmt.Sprintf("<span class='postlabel %v' >%v</span>\n", inp.CSSLabel, inp.Suffix.TrSilent(q.LangCode))
		}

		if inp.ErrMsg.Set() {
			ctrl += fmt.Sprintf("<span class='go-quest-label %v' >%v</span>\n", inp.CSSLabel, inp.ErrMsg.TrSilent(q.LangCode)) // ugly layout  - but radiogroup and checkboxgroup won't have validation errors anyway
		}

		return ctrl

	case "dynamic":
		return fmt.Sprintf("<span class='go-quest-label %v'>%v</span>\n", inp.CSSLabel, inp.Label.Tr(q.LangCode))

	case "composit", "composit-scalar":
		// rendered at group level -  rendered by composit
		return ""

	default:
		return fmt.Sprintf("input %v: unknown type '%v'  - allowed are %v\n", nm, inp.Type, implementedTypes)
	}

}
