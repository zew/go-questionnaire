package qst

import "github.com/zew/go-questionnaire/trl"

// AddStaticDynamicGroups adds five groups
// the second remains static, the others are
// randomized based on
func AddStaticDynamicGroups(page pageT) {
	// DUMMMY  DUMMMY  DUMMMY
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.RandomizationGroup = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Name = "text03"
			inp.Desc = trl.S{"de": `sg1 - el1`}
		}
	}
	// DUMMMY  DUMMMY  DUMMMY
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.RandomizationGroup = 0
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Name = "text03"
			inp.Desc = trl.S{"de": `no shuffle`}
		}
	}
	// DUMMMY  DUMMMY  DUMMMY
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.RandomizationGroup = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Name = "text03"
			inp.Desc = trl.S{"de": `sg1 - el2`}
		}
	}
	// DUMMMY  DUMMMY  DUMMMY
	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.RandomizationGroup = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Name = "text03"
			inp.Desc = trl.S{"de": `sg1 - el3`}
		}
	}

}
