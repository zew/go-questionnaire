package fmtest

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func randomizedVerticalRadiosWithFree(
	page *qst.WrappedPageT,
	inputStem string,
	rowLbls []trl.S,
	randomGroup int,
	lastWithFree bool,
) {

	for i := 0; i < len(rowLbls); i++ {

		gr := page.AddGroup()
		gr.Cols = 6
		gr.RandomizationGroup = randomGroup

		isLast := i == len(rowLbls)-1

		gr.BottomVSpacers = 1

		if lastWithFree && isLast {
			gr.RandomizationGroup = 0
			gr.BottomVSpacers = 3
		}

		{

			inp1 := gr.AddInput()
			inp1.Type = "radio"
			inp1.Name = inputStem
			inp1.ValueRadio = fmt.Sprintf("%v", i+1)
			inp1.ColSpan = gr.Cols
			inp1.ColSpanLabel = 1
			inp1.ColSpanControl = 12
			inp1.Label = rowLbls[i]

			inp1.ControlFirst()

			if lastWithFree && isLast {

				inp1.ColSpan = 2
				inp1.ColSpanLabel = 2.4
				inp1.ColSpanControl = 7.7

				//
				inp2 := gr.AddInput()
				inp2.Type = "text"
				inp2.Name = inputStem + "_free"
				inp2.MaxChars = 100

				inp2.ColSpan = gr.Cols - inp1.ColSpan
				inp2.ColSpanLabel = 0
				inp2.ColSpanControl = 1
			}

		}

		{
			inp := gr.AddInput()
			inp.ColSpanControl = 1
			inp.Type = "javascript-block"
			inp.Name = "radio-xor-text"
			s1 := trl.S{
				"de": "unused",
				"en": "unused",
			}
			inp.JSBlockTrls = map[string]trl.S{
				"msg": s1,
			}
			inp.JSBlockStrings = map[string]string{
				"inp1":    inputStem,
				"inp2":    inputStem + "_free",
				"radioOn": inputStem + "6",
			}
		}
	}

	//
	// since the groups above can be randomized,
	// we cannot give a vertical spacer at the "end"
	// =>  explicit vertical spacer
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.Label = trl.S{
				"de": `&nbsp;`,
				"en": `&nbsp;`,
			}
		}
	}

}

func randomizedMatrixWithFree(
	page *qst.WrappedPageT,
	colLabels []trl.S,
	inputStem string,
	rowLbls []trl.S,
	randGroup int,
	// showFree bool,
	lblFree trl.S,
) {

	// colTemplate, colsRowFree, styleRowFree := colTemplateWithFreeRow()

	colTemplateStr := "7fr       1fr 1fr 1fr 1fr 1fr   1.4fr"
	styleRowFree := "  7fr       1fr 1fr 1fr 1fr 1fr   1.4fr"

	//
	{
		gr := page.AddGroup()
		gr.Cols = 7
		gr.BottomVSpacers = 0

		// equal to below
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		gr.Style.Desktop.StyleText.FontSize = 90

		// col labels slightly upwards
		gr.Style.Desktop.StyleBox.Margin = "0 0 0.5ch 0 "

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = trl.S{
				"de": "&nbsp;",
				"en": "&nbsp;",
			}
		}
		for i := 0; i < len(colLabels); i++ {
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = colLabels[i]
			inp.LabelCenter()
			inp.LabelBottom()
		}
	}

	//
	//
	for i1 := 0; i1 < len(rowLbls); i1++ {
		gr := page.AddGroup()
		firstCol := float32(1)
		gr.Cols = firstCol + 6
		gr.RandomizationGroup = randGroup
		gr.BottomVSpacers = 0
		if i1 == (len(rowLbls) - 1) {
			gr.BottomVSpacers = 2 // bad, because of shuffling
			gr.BottomVSpacers = 0
		}

		// equal to above
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.Display = "grid"
		// gr.Style.Desktop.StyleGridContainer.TemplateColumns = "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		// gr.Style.Mobile.StyleGridContainer.TemplateColumns =  "7fr 1fr 1fr 1fr 1fr 1fr 1.4fr"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = colTemplateStr
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = colTemplateStr

		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		// distinct
		gr.Style.Desktop.StyleBox.Margin = "0 0 0.6rem" // bottom margin

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = firstCol
			inp.ColSpanLabel = 1
			inp.ColSpanControl = 0
			inp.Label = rowLbls[i1]
		}
		for i2 := 0; i2 < 6; i2++ {
			{
				inp := gr.AddInput()
				inp.Type = "radio"
				inp.Name = fmt.Sprintf("%v_%v", inputStem, i1+1)
				inp.ValueRadio = fmt.Sprintf("%v", i2+1)
				inp.ColSpan = 1
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}
	}

	//
	//
	//
	//
	// row free input
	if lblFree == nil {

		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 3

	} else {

		gr := page.AddGroup()
		gr.Cols = 7

		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleBox.Display = "grid"
		gr.Style.Desktop.StyleGridContainer.TemplateColumns = styleRowFree
		gr.Style.Mobile.StyleGridContainer.TemplateColumns = styleRowFree
		gr.Style.Desktop.StyleGridContainer.GapColumn = "0.8rem"
		gr.Style.Mobile.StyleGridContainer.GapColumn = "0.2rem"

		gr.BottomVSpacers = 3

		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = fmt.Sprintf("%v_free", inputStem)
			// inp.MaxChars = 17
			inp.MaxChars = 25
			inp.ColSpan = 1
			inp.ColSpanLabel = 2.4
			inp.ColSpanLabel = 0.9
			inp.ColSpanControl = 4
			inp.Label = lblFree
		}

		//
		for idx := 0; idx < len(colLabels); idx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = fmt.Sprintf("%v_free_val", inputStem)
			rad.ValueRadio = fmt.Sprint(idx + 1)
			rad.ColSpan = 1

			rad.ColSpanLabel = 0
			rad.ColSpanControl = 1
		}

	}

}
