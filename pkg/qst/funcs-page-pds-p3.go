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

	page.Label = trl.S{
		"en": "3. Index Questions",
		"de": "3. Index Questions",
	}
	page.Short = trl.S{
		"en": "Indizes",
		"de": "Indizes",
	}
	page.CounterProgress = "3"

	page.WidthMax("52rem") // getting the nice "valley" alignment
	if len(ac.TrancheTypes) == 2 {
		page.WidthMax("36rem")
	}
	if len(ac.TrancheTypes) == 1 {
		page.WidthMax("24rem")
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
			"en": `Financing situation/pricing <br>
						<span class=font-size-90> Assess the development of expected returns  </span>`,
			"de": `Financing situation/pricing <br>
						<span class=font-size-90> Assess the development of expected returns  </span>`,
		},
		{
			"en": `Assess the change in deal quality with respect to the risk return profile <br>
						<span class=font-size-90> Assess the change in deal quality with respect to the risk return profile </span>`,
			"de": `Assess the change in deal quality with respect to the risk return profile <br>
						<span class=font-size-90> Assess the change in deal quality with respect to the risk return profile </span>`,
		},
		{
			"en": `Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) <br>
						<span class=font-size-90>  Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) </span>`,
			"de": `Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) <br>
						<span class=font-size-90>  Assess the quality of deal documentation (covenant strength, enforcement rights, etc.) </span>`,
		},
		{
			"en": `Do you observe more deals, same amount of deals or less deals <br>
						<span class=font-size-90> Do you observe more deals, same amount of deals or less deals  </span>`,
			"de": `Do you observe more deals, same amount of deals or less deals <br>
						<span class=font-size-90> Do you observe more deals, same amount of deals or less deals  </span>`,
		},
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

		radiosLabelsTop(
			page,
			ac,
			inpName+"_past3m",
			trl.S{
				"en": "<i>Last</i> 3&nbsp;months",
				"de": "<i>Last</i> 3&nbsp;months",
			}.Outline("a.)"),
			mCh4,
		)
		radiosLabelsTop(
			page,
			ac,
			inpName+"_next3m",
			trl.S{
				"en": "<i>Next</i> 3&nbsp;months",
				"de": "<i>Next</i> 3&nbsp;months",
			}.Outline("b.)"),
			mCh4,
		)

	}

	return nil
}
