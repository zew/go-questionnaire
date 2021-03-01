package qst

import (
	"fmt"
	"log"

	"github.com/zew/go-questionnaire/css"
	"github.com/zew/go-questionnaire/trl"
)

/*RadioMatrix to generate tables of inputs
 */
type RadioMatrix struct {
	InpNames []string // for each row an input name - i.e. y_deu, y_usa, y_glob,

	Labels1stRow []trl.S // good, normal, bad, no answer
	InpCols      int     // if Labels1stRow are nil
	RadioValues  []int   // corresponding to Labels1stRow/InpCols; i.e. no answer (no entry) 0 - good 1 - normal 2 - bad 3 - (explicit) no answer 4
	//
	Labels           map[int][]trl.S // by col; i.e. first col is Euro area, Germany, US, Global economy; second col is empty
	ColSpansOfLabels map[int]int     // by col - span for the labels - controls have span 1 - applies to each row

	// boolean helpers
	has1stRowLabel, has1stColLabel bool
}

// ExampleRadioMatrix demonstrates usage
func ExampleRadioMatrix() *RadioMatrix {

	rm := RadioMatrix{}

	rm.InpNames = []string{
		"y0_ez",
		"y0_deu",
		"y0_usa",
		"y0_glob",
	}

	//
	//
	rm.Labels1stRow = []trl.S{
		{
			"de": "gut",
			"en": "good",
		},
		{
			"de": "normal",
			"en": "normal",
		},
		{
			"de": "schlecht",
			"en": "bad",
		},
		{
			"de": "keine<br>Angabe",
			"en": "no answer",
		},
	}

	rm.InpCols = len(rm.Labels1stRow)

	// rm.RadioValues = []int{1, 2, 3, 4}
	rm.RadioValues = make([]int, 0, rm.InpCols)
	for i := 0; i < rm.InpCols; i++ {
		rm.RadioValues = append(rm.RadioValues, i+1)
	}

	//
	//
	rm.Labels = map[int][]trl.S{
		0: []trl.S{
			{
				"de": "Euroraum",
				"en": "Euro area",
			},
			{
				"de": "Deutschland",
				"en": "Germany",
			},
			{
				"de": "USA",
				"en": "US",
			},
			{
				"de": "Weltwirtschaft",
				"en": "Global economy",
			},
		},
		3: []trl.S{
			{
				"de": "--",
				"en": "--",
			},
			{
				"de": "--",
				"en": "--",
			},
			{
				"de": "--",
				"en": "--",
			},
			{
				"de": "--",
				"en": "--",
			},
		},
	}
	rm.ColSpansOfLabels = map[int]int{
		0: 2,
		3: 2,
	}

	delete(rm.Labels, 3)
	delete(rm.ColSpansOfLabels, 3)

	return &rm
}

func (rm *RadioMatrix) checkConsistency() error {

	rm.has1stRowLabel = false
	if len(rm.Labels1stRow) > 0 {
		rm.has1stRowLabel = true
	}

	//
	if rm.has1stRowLabel {
		if rm.InpCols != len(rm.Labels1stRow) {
			return fmt.Errorf("AddRadioMatrixGroup(): labels for 1st row exist, want %v, got %v", rm.InpCols, len(rm.Labels1stRow))
		}
	}

	//
	for col, colLabels := range rm.Labels {
		if len(rm.InpNames) != len(colLabels) {
			return fmt.Errorf("AddRadioMatrixGroup(): labels for col %v exist, want %v, got %v", col, len(rm.InpNames), len(colLabels))
		}

		if _, ok := rm.ColSpansOfLabels[col]; !ok {
			return fmt.Errorf("AddRadioMatrixGroup(): labels for col %v exist, but no colspan setting %v", col, rm.ColSpansOfLabels)
		}

	}

	if rm.ColSpansOfLabels[0] > 0 {
		rm.has1stColLabel = true
	}

	// log.Printf("has1stRowLabel %v - has1stColLabel %v", rm.has1stRowLabel, rm.has1stColLabel)

	return nil
}

// AddRadioMatrix2 adds a group to page
func (p *pageT) AddRadioMatrix2(rm *RadioMatrix) *groupT {

	err := rm.checkConsistency()
	if err != nil {
		log.Panic(err)
	}

	log.Printf("has1stRowLabel %v - has1stColLabel %v", rm.has1stRowLabel, rm.has1stColLabel)

	//
	gr := p.AddGroup()
	gr.Cols = rm.InpCols
	log.Printf("gr.Cols %v - rm.InpCols %v - rm.ColSpansOfLabels %+v", gr.Cols, rm.InpCols, rm.ColSpansOfLabels)
	for _, colSp := range rm.ColSpansOfLabels {
		gr.Cols += colSp
		log.Printf("gr.Cols %v - colSp %v", gr.Cols, colSp)
	}
	log.Printf("gr.Cols %v", gr.Cols)

	//
	// Header row - first column
	// (empty cell in top-left)
	if rm.has1stRowLabel && rm.has1stColLabel {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Label = trl.S{
			"en": " &nbsp; ",
			"de": " &nbsp; ",
		}
		inp.ColSpanLabel = rm.ColSpansOfLabels[0]
		inp.CSSLabel = CSSLabelHeader // apply even if its empty
	}

	//
	// Header row - next columns
	for _, lbl := range rm.Labels1stRow {
		inp := gr.AddInput()
		inp.Type = "textblock"
		inp.Style = css.NewStylesResponsive()
		inp.Style.Desktop.GridItemStyle.JustifySelf = "center"
		inp.Style.Desktop.GridItemStyle.AlignSelf = "center"
		inp.Style.Desktop.TextStyle.AlignHorizontal = "center"
		inp.StyleLbl = inp.Style
		inp.Desc = lbl // for instance trl.S{"de": "gut", "en": "good"}
		inp.CSSLabel = CSSLabelHeader
	}

	//
	// Input rows
	for rowIdx, name := range rm.InpNames {

		for colIdx := 0; colIdx < rm.InpCols; colIdx++ {
			rad := gr.AddInput()
			rad.Type = "radio"
			rad.Name = name // "y_euro"
			rad.ValueRadio = fmt.Sprint(rm.RadioValues[colIdx])

			if lblsByRow, hasLbl := rm.Labels[colIdx]; hasLbl {
				rad.Label = lblsByRow[rowIdx]
				rad.ColSpanLabel = rm.ColSpansOfLabels[colIdx]
			} else {
				rad.ColSpanLabel = 0
			}

			rad.ColSpanControl = 1

		}

	}

	return gr

}
