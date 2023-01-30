package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage3AC1(q *QuestionnaireT, page *pageT) error {
	return pdsPage3(q, page, 0)
}
func pdsPage3AC2(q *QuestionnaireT, page *pageT) error {
	return pdsPage3(q, page, 1)
}
func pdsPage3AC3(q *QuestionnaireT, page *pageT) error {
	return pdsPage3(q, page, 2)
}

func pdsPage3(q *QuestionnaireT, page *pageT, acIdx int) error {

	ac := PDSAssetClasses[acIdx]
	ac = onlySelectedTranchTypes(q, ac)
	rn := rune(65 + acIdx) // ascii 65 is A; 97 is a

	page.Label = trl.S{
		"en": fmt.Sprintf("%v: &nbsp;&nbsp; Market assessment in %v", ac.Lbl["en"], q.Survey.Quarter(-1)),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp; Market assessment in %v", ac.Lbl["de"], q.Survey.Quarter(-1)),
	}.Outline(fmt.Sprintf("%c3.", rn))

	page.Short = trl.S{
		"en": ("<b> &nbsp;<br> &nbsp; </b> Market<br>&nbsp;&nbsp;Assessment"),
		"de": ("<b> &nbsp;<br> &nbsp; </b> Market<br>&nbsp;&nbsp;Assessment"),
	}

	page.CounterProgress = fmt.Sprintf("%c3", rn)

	page.WidthMax("52rem") // getting the nice "valley" alignment
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("42rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("38rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	lblIntro := []trl.S{
		{
			"en": `The following questions are related to the developments of the <i>overall European Corporate Direct Lending market</i>.`,
			"de": `The following questions are related to the developments of the <i>overall European Corporate Direct Lending market</i>.`,
		},
		{
			"en": `The following questions are aimed to capture your views on the developments of the <i>overall European Real Estate Debt market</i>.
			`,
			"de": `The following questions are aimed to capture your views on the developments of the <i>overall European Real Estate Debt market</i>.
			`,
		},
		{
			"en": `The following questions are aimed to capture your views on the developments of the <i>overall European Infrastructure Debt market</i>.
			`,
			"de": `The following questions are aimed to capture your views on the developments of the <i>overall European Infrastructure Debt market</i>.
			`,
		},
	}

	{
		gr := page.AddGroup()
		gr.BottomVSpacers = 1
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblIntro[acIdx]
		}

	}

	page3Inputs := []string{
		"q31_overall_environ",
		"q32_pricing",
		"q33_qual_documentation",
		"q34_pipeline",
	}
	page3Lbls := []trl.S{
		{
			"en": `Overall market environment`,
			"de": `Overall market environment`,
		},
		{
			"en": `Pricing`,
			"de": `Pricing`,
		},
		{
			"en": `Quality of documentation`,
			"de": `Quality of documentation`,
		},
		{
			"en": `Deal flow`,
			"de": `Deal flow`,
		},
	}

	page3LblsDescr := []trl.S{
		{
			"en": `Please assess the overall market environment over the past quarter.`,
			"de": `Please assess the overall market environment over the past quarter.`,
		},
		{
			"de": `Please assess the attractiveness of the pricing of new loans relative to their risk profile. `,
			"en": `Please assess the attractiveness of the pricing of new loans relative to their risk profile. `,
		},
		{
			"en": `Please assess the quality of deal documentation for new deals you observe in the market 
				(stricter loan documentation=good, looser loan documentation=bad). `,
			"de": `Please assess the quality of deal documentation for new deals you observe in the market 
				(stricter loan documentation=good, looser loan documentation=bad). `,
		},
		{
			"en": `Please assess the deal flow you observe in the market.`,
			"de": `Please assess the deal flow you observe in the market.`,
		},
	}

	page3LblsPrevAndPast := []trl.S{
		//
		{
			"en": `Over the past quarter, please provide your assessment of the overall market environment?`,
			"de": `Over the past quarter, please provide your assessment of the overall market environment?`,
		},
		{
			"en": `How do you expect the overall market environment to change [quarter-0]?`,
			"de": `How do you expect the overall market environment to change [quarter-0]?`,
		},

		//
		{
			"en": `How attractive is the pricing of new loans you observed in the market over the past quarter?`,
			"de": `How attractive is the pricing of new loans you observed in the market over the past quarter?`,
		},
		{
			"en": `How do you expect the attractiveness of pricing of new loans to change in [quarter-0]?`,
			"de": `How do you expect the attractiveness of pricing of new loans to change in [quarter-0]?`,
		},

		//
		{
			"en": `How do you assess the quality of loan documentation (e.g. covenant quality, enforcement rights,&nbsp;...) over the past quarter? `,
			"de": `How do you assess the quality of loan documentation (e.g. covenant quality, enforcement rights,&nbsp;...) over the past quarter? `,
		},
		{
			"en": `How do you expect the quality of loan documentation (e.g. covenant quality, enforcement rights, …) to change in [quarter-0]?`,
			"de": `How do you expect the quality of loan documentation (e.g. covenant quality, enforcement rights, …) to change in [quarter-0]?`,
		},

		//
		{
			"en": `How do you assess the deal flow in the market over the past quarter?`,
			"de": `How do you assess the deal flow in the market over the past quarter?`,
		},
		{
			"en": `How do you expect the deal flow to change in [quarter-0]?`,
			"de": `How do you expect the deal flow to change in [quarter-0]?`,
		},
	}

	for i := 0; i < len(page3Lbls); i++ {
		page3Lbls[i].Append90(page3LblsDescr[i])
	}

	for i := 0; i < len(page3Lbls); i++ {
		page3Lbls[i] = page3Lbls[i].Outline(fmt.Sprintf("3.%v", i+1))
	}

	for idx1, inpName := range page3Inputs {
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 0
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = page3Lbls[idx1].Bold()
				inp.Label = page3Lbls[idx1]

				inp.ColSpan = 1
				// inp.ColSpanLabel = 1
			}
		}

		tlbPrev3M := trl.S{
			"en": "<i>Last</i> 3&nbsp;months",
			"de": "<i>Last</i> 3&nbsp;months",
		}
		tlbPrev3M.Append90(page3LblsPrevAndPast[2*idx1+0])
		tlbPrev3M.Outline("a.)")

		if idx1 != 1 {
			radiosLabelsTop(
				page,
				ac,
				inpName+"_past3m",
				tlbPrev3M,
				mCh4Prev,
			)
		} else {
			radiosLabelsTop(
				page,
				ac,
				inpName+"_past3m",
				tlbPrev3M,
				mCh4PrevSpec,
			)

		}

		lblNext3M := trl.S{
			"en": "<i>Next</i> 3&nbsp;months",
			"de": "<i>Next</i> 3&nbsp;months",
		}
		lblNext3M.Append90(page3LblsPrevAndPast[2*idx1+1])
		lblNext3M.Outline("b.)")

		radiosLabelsTop(
			page,
			ac,
			inpName+"_next3m",
			lblNext3M,
			mCh4Next,
		)

	}

	return nil
}
