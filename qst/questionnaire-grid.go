package qst

import (
	"fmt"
	"io"
	"strings"
)

func gridItem(w io.Writer, className, cnt string) {

	fmt.Fprintf(w,
		"<div class='%v' >\n\t%v\n\t</div>\n",
		className,
		cnt,
	)

}

func gridContainer(w io.Writer, className, cnt string) {

	fmt.Fprintf(w,
		"<div class='%v' >\n\t%v\n\t</div>\n",
		className,
		cnt,
	)

}

func (gr *groupT) Header(w io.Writer, langCode string) {

	if gr.Label == nil && gr.Desc == nil ||
		gr.Label == nil && gr.Desc != nil { // no desc without label
		return
	}

	wLoc := &strings.Builder{}
	fmt.Fprintf(wLoc, "%v", gr.Label.Tr(langCode))
	if gr.Desc != nil {
		fmt.Fprintf(wLoc, "<div style='font-size: 90%%'>%v</div>", gr.Desc.Tr(langCode))
	}

	fmt.Fprintf(w,
		`
		<div
			class="grid-item group-header"
			style="
				font-size: 120%%;
				grid-column: 1 / span %v;
			"
		>
			%v
		</div>
`,
		gr.Cols,
		wLoc.String(),
	)

	fmt.Fprintf(w, vspacer)
	for i := 0; i < gr.HeaderBottomVSpacers; i++ {
		fmt.Fprintf(w, vspacer8)
	}

}

// GroupHTMLGrid renders a group of inputs to GroupHTMLGrid
func (q QuestionnaireT) GroupHTMLGrid(pageIdx, grpIdx int) string {

	wCSS := &strings.Builder{}
	gr := q.Pages[pageIdx].Groups[grpIdx]

	//
	//
	gr.Style.Desktop.BoxStyle.Display = "grid"
	if gr.Style.Desktop.GridContainerStyle.AutoFlow == "" {
		gr.Style.Desktop.GridContainerStyle.AutoFlow = "row"
	}
	if gr.Style.Desktop.GridContainerStyle.TemplateColumns == "" {
		gr.Style.Desktop.GridContainerStyle.TemplateColumns = strings.Repeat("1fr ", gr.Cols)
	}
	gridContainerClass := fmt.Sprintf("cls-pg%02v-grp%02v", pageIdx, grpIdx)
	fmt.Fprint(wCSS, gr.Style.CSS(gridContainerClass))

	//
	//
	wContInner := &strings.Builder{}
	gr.Header(wContInner, q.LangCode)

	for _, inp := range gr.Inputs {
		if inp.Type == "composit-scalar" {
			continue
		}
		if inp.Type == "composit" {
			continue
		}

		fmt.Fprint(wContInner, "<div class='grid-item'>\n")
		fmt.Fprint(wContInner, inp.HTML(q.LangCode, gr.Cols)) // rendering markup
		fmt.Fprint(wContInner, "\n")
		fmt.Fprint(wContInner, "<div>\n")
	}

	//
	//
	wContainer := &strings.Builder{}
	gridContainer(wContainer, gridContainerClass, wContInner.String())

	w := &strings.Builder{}
	fmt.Fprint(w, wCSS.String())
	fmt.Fprint(w, wContainer.String())
	return w.String()

}
