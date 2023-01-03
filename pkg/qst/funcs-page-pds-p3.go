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
		"en": fmt.Sprintf("%v: &nbsp;&nbsp; Market Sentiment", ac.Lbl["en"]),
		"de": fmt.Sprintf("%v: &nbsp;&nbsp; Market Sentiment", ac.Lbl["de"]),
	}.Outline(fmt.Sprintf("%c3.", rn))
	page.Short = trl.S{
		"en": fmt.Sprintf("%v<br>sent.", ac.Short["en"]),
		"de": fmt.Sprintf("%v<br>sent.", ac.Short["de"]),
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

	page3Inputs := []string{
		"q31_financing_situation_pricing",
		"q32_deal_quality",
		"q33_deal_documentation",
		"q34_deal_amount",
	}
	page3Lbls := []trl.S{
		{
			"en": `Pricing`,
			"de": `Pricing`,
		},
		{
			"en": `Quality of documentation`,
			"de": `Quality of documentation`,
		},
		{
			"en": `Quality of deals`,
			"de": `Quality of deals`,
		},
		{
			"en": `Pipeline / Deal Flow`,
			"de": `Pipeline / Deal Flow`,
		},
	}

	page3LblsDescr := []trl.S{
		{
			"en": `Please assess the pricing of new loans.  (higher return=improved, lower return=declined) `,
			"de": `Please assess the pricing of new loans.  (higher return=improved, lower return=declined) `,
		},

		{
			"en": `Please assess the quality of the deal documentation for new loans. (stricter loan documentation=improved, looser loan documentation=declined)`,
			"de": `Please assess the quality of the deal documentation for new loans. (stricter loan documentation=improved, looser loan documentation=declined)`,
		},
		{
			"en": `Please assess the deal quality of new loans with respect to the risk-return profile.`,
			"de": `Please assess the deal quality of new loans with respect to the risk-return profile.`,
		},
		{
			"en": `Please assess the development of demand for credit. `,
			"de": `Please assess the development of demand for credit. `,
		},
	}

	page3LblsPrevAndPast := []trl.S{
		{
			"en": `How did the pricing (taking into account margins over the relevant reference rate and other return components like fees) for new deals you observed in the market change over the past quarter?`,
			"de": `How did the pricing (taking into account margins over the relevant reference rate and other return components like fees) for new deals you observed in the market change over the past quarter?`,
		},
		{
			"en": `How do you expect the pricing (taking into account margins over the relevant reference rate and other return components like fees) for new deals you observed in the market change in the coming quarter?`,
			"de": `How do you expect the pricing (taking into account margins over the relevant reference rate and other return components like fees) for new deals you observed in the market change in the coming quarter?`,
		},

		{
			"en": `How did the quality of loan documentation (taking into covenant quality, enforcement rights, etc.) for new deals you observed in the market change over the past quarter? `,
			"de": `How did the quality of loan documentation (taking into covenant quality, enforcement rights, etc.) for new deals you observed in the market change over the past quarter? `,
		},
		{
			"en": `How do you expect the quality of loan documentation (taking into covenant quality, enforcement rights, etc.) for new deals you observed in the market change in the coming quarter?`,
			"de": `How do you expect the quality of loan documentation (taking into covenant quality, enforcement rights, etc.) for new deals you observed in the market change in the coming quarter?`,
		},

		{
			"en": `How did the risk-return profile for new deals you observed in the market change over the past quarter?`,
			"de": `How did the risk-return profile for new deals you observed in the market change over the past quarter?`,
		},
		{
			"en": `How do you expect  the risk-return profile for new deals you observed in the market change in the coming quarter?`,
			"de": `How do you expect  the risk-return profile for new deals you observed in the market change in the coming quarter?`,
		},

		{
			"en": `Apart from normal seasonal fluctuations, how did the volume of deals you observed in the market change in the last quarter?`,
			"de": `Apart from normal seasonal fluctuations, how did the volume of deals you observed in the market change in the last quarter?`,
		},
		{
			"en": `Apart from normal seasonal fluctuations, how do you expect the volume of deals you observe in the market change in the coming quarter?`,
			"de": `Apart from normal seasonal fluctuations, how do you expect the volume of deals you observe in the market change in the coming quarter?`,
		},
	}
	_ = page3LblsPrevAndPast

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

		radiosLabelsTop(
			page,
			ac,
			inpName+"_past3m",
			tlbPrev3M,
			mCh4,
		)

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
			mCh4,
		)

	}

	return nil
}
