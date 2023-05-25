package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/trl"
)

func pdsPage3X(q *QuestionnaireT, page *pageT) error {

	// ac := PDSAssetClasses[acIdx]
	// ac = onlySelectedTranchTypes(q, ac)
	ac := PDSAssetClassGlob
	ac = onlySelectedTranchTypes2(q, ac)

	page.Label = trl.S{
		"en": fmt.Sprintf(" European Private Debt Markets in %v", q.Survey.Quarter(0)),
		"de": fmt.Sprintf(" European Private Debt Markets in %v", q.Survey.Quarter(0)),
	}

	page.Label.Outline("1.")

	page.Short = trl.S{
		"en": (" Private <br>   Debt Markets 1"),
		"de": (" Private <br>   Debt Markets 1"),
	}

	// page.CounterProgress = fmt.Sprintf("%c3", rn)
	page.CounterProgress = "1"

	page.WidthMax("52rem")
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("42rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("32rem")
	}

	// dynamically recreate the groups
	page.Groups = nil

	lblIntro := []trl.S{
		{
			"en": `<i>The following questions relate to the developments of overall European private debt markets.</i>`,
			"de": `<i>The following questions relate to the developments of overall European private debt markets.</i>`,
		},
	}

	{
		gr := page.AddGroup()
		gr.BottomVSpacers = 1
		gr.Cols = 1

		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblIntro[0]
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
			"en": `The overall market environment`,
			"de": `The overall market environment`,
		},
		{
			"en": `The pricing of new loans`,
			"de": `The pricing of new loans`,
		},
		{
			"en": `Documentation quality`,
			"de": `Documentation quality`,
		},
		{
			"en": `Deal flow`,
			"de": `Deal flow`,
		},
	}

	page3LblsDescr := []trl.S{
		{
			"en": `How do you assess the overall market environment during the previous quarter?`,
			"de": `How do you assess the overall market environment during the previous quarter?`,
		},
		{
			"de": `Relative to the risk profile, how attractive was the pricing of new loans you observed in the market during the previous quarter?`,
			"en": `Relative to the risk profile, how attractive was the pricing of new loans you observed in the market during the previous quarter?`,
		},
		{
			"en": `How do you assess the quality of deal documentation (e.g. covenant quality, enforcement rights, etc.) for new deals you observed in the market during the previous quarter (stricter loan documentation=good, looser loan documentation=bad)?`,
			"de": `How do you assess the quality of deal documentation (e.g. covenant quality, enforcement rights, etc.) for new deals you observed in the market during the previous quarter (stricter loan documentation=good, looser loan documentation=bad)?`,
		},
		{
			"en": `How do you assess the deal flow you observed in the market during the previous quarter?`,
			"de": `How do you assess the deal flow you observed in the market during the previous quarter?`,
		},
	}

	page3LblsPrevAndPast := []trl.S{
		//
		{
			"en": `How do you assess the overall market environment during the previous quarter?`,
			"de": `How do you assess the overall market environment during the previous quarter?`,
		},
		{
			"en": `Compared to the last quarter, how do you expect the overall market environment to change in the current quarter?`,
			"de": `Compared to the last quarter, how do you expect the overall market environment to change in the current quarter?`,
		},

		//
		//
		{
			"en": `Relative to the risk profile, how attractive was the pricing of new loans you observed in the market during the previous quarter?`,
			"de": `Relative to the risk profile, how attractive was the pricing of new loans you observed in the market during the previous quarter?`,
		},
		{
			"en": `Compared to the last quarter, how do you expect the attractiveness of the pricing of new loans (relative to the risk profile) to change in the current quarter?`,
			"de": `Compared to the last quarter, how do you expect the attractiveness of the pricing of new loans (relative to the risk profile) to change in the current quarter?`,
		},

		//
		//
		{
			"en": `How do you assess the quality of deal documentation (e.g. covenant quality, enforcement rights, etc.) for new deals you observed in the market during the previous quarter (stricter loan documentation=good, looser loan documentation=bad)`,
			"de": `How do you assess the quality of deal documentation (e.g. covenant quality, enforcement rights, etc.) for new deals you observed in the market during the previous quarter (stricter loan documentation=good, looser loan documentation=bad)`,
		},
		{
			"en": `Compared to the last quarter, how do you expect the quality of loan documentation (e.g. covenant quality, enforcement rights, etc.) to change in the current quarter?`,
			"de": `Compared to the last quarter, how do you expect the quality of loan documentation (e.g. covenant quality, enforcement rights, etc.) to change in the current quarter?`,
		},

		//
		{
			"en": `How do you assess the deal flow you observed in the market during the previous quarter?`,
			"de": `How do you assess the deal flow you observed in the market during the previous quarter?`,
		},
		{
			"en": `Compared to the last quarter, how do you expect the deal flow to change in the current quarter?`,
			"de": `Compared to the last quarter, how do you expect the deal flow to change in the current quarter?`,
		},
	}

	// for i := 0; i < len(page3Lbls); i++ {
	// 	page3Lbls[i].Append90(page3LblsDescr[i])
	// }
	_ = page3LblsDescr

	for i := 0; i < len(page3Lbls); i++ {
		// page3Lbls[i] = page3Lbls[i].Outline(fmt.Sprintf("3.%v", i+1))
		page3Lbls[i] = page3Lbls[i].Outline(fmt.Sprintf("1.%v", i+1))
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
			"en": fmt.Sprintf(" Past quarter (%v)", q.Survey.Quarter(-1)),
			"de": fmt.Sprintf(" Past quarter (%v)", q.Survey.Quarter(-1)),
		}
		tlbPrev3M.Append90(page3LblsPrevAndPast[2*idx1+0])
		tlbPrev3M.Outline("a.)")

		radiosLabelsTop(
			page,
			ac,
			inpName+"_past3m",
			tlbPrev3M,
			mCh4Prev,
		)

		lblNext3M := trl.S{
			"en": fmt.Sprintf(" Current quarter (%v)", q.Survey.Quarter(-0)),
			"de": fmt.Sprintf(" Current quarter (%v)", q.Survey.Quarter(-0)),
		}
		lblNext3M.Append90(page3LblsPrevAndPast[2*idx1+1])
		lblNext3M.Outline("b.)")

		if idx1 != 0 {
			radiosLabelsTop(
				page,
				ac,
				inpName+"_next3m",
				lblNext3M,
				mCh4Next,
			)
		} else {
			radiosLabelsTop(
				page,
				ac,
				inpName+"_next3m",
				lblNext3M,
				mCh4Next2,
			)
		}

	}

	{
		// just for vertical spacer
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = trl.S{}
			inp.ColSpan = 1
		}
	}

	return nil
}
