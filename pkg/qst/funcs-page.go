package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/trl"
)

/*
	Example usage
	{
		page := q.AddPage()
		page.GeneratorFuncName = "fmt202212"
		page.NavigationCondition = naviCondition
	}

*/

// Notice, that these func are called
// on each request
type funcPageGeneratorT func(*QuestionnaireT, *pageT) error

var pageGens = ctr.New() // page generations

var funcPGs = map[string]funcPageGeneratorT{

	"pageFuncExample": pageFuncExample,

	"fmt202212": fmt202212,
	"fmt202312": fmt202312,
	"fmt202402": fmt202402,
	"fmt202405": fmt202405,

	"pdsPage11-ac1": pdsPage11AC1,
	"pdsPage11-ac2": pdsPage11AC2,
	"pdsPage11-ac3": pdsPage11AC3,

	"pdsPage12-ac1": pdsPage12AC1,
	"pdsPage12-ac2": pdsPage12AC2,
	"pdsPage12-ac3": pdsPage12AC3,

	"pdsPage21-ac1": pdsPage21AC1,
	"pdsPage21-ac2": pdsPage21AC2,
	"pdsPage21-ac3": pdsPage21AC3,

	"pdsPage23-ac1": pdsPage23AC1,
	"pdsPage23-ac2": pdsPage23AC2,
	"pdsPage23-ac3": pdsPage23AC3,

	"pdsPage3-ac1": pdsPage3AC1,
	"pdsPage3-ac2": pdsPage3AC2,
	"pdsPage3-ac3": pdsPage3AC3,

	"pdsPage4-ac1": pdsPage4AC1,
	"pdsPage4-ac2": pdsPage4AC2,
	"pdsPage4-ac3": pdsPage4AC3,

	//
	"pdsPage3-acx": pdsPage3ACX,
	"pdsPage4-acx": pdsPage4ACX,

	"kneb202306guidedtourN0": kneb202306guidedtourN0,
	"kneb202306guidedtourN1": kneb202306guidedtourN1,
	"kneb202306guidedtourN2": kneb202306guidedtourN2,
	"kneb202306guidedtourN3": kneb202306guidedtourN3,
	"kneb202306guidedtourN4": kneb202306guidedtourN4,
	"kneb202306guidedtourN5": kneb202306guidedtourN5,
	"kneb202306guidedtourN6": kneb202306guidedtourN6,
	"kneb202306guidedtourN7": kneb202306guidedtourN7,

	"kneb202306simtool0": kneb202306simtool0, // first instance - exercise
	"kneb202306simtool1": kneb202306simtool1, //
	"kneb202306simtool2": kneb202306simtool2, //
	"kneb202306simtool3": kneb202306simtool3, //
	"kneb202306simtool4": kneb202306simtool4, // second instance - input values

	"fmt202511Pg2": fmt202511Pg2,
	"fmt202511Pg3": fmt202511Pg3,
}

func pageFuncExample(q *QuestionnaireT, page *pageT) error {

	gn := pageGens.Increment()

	// modify/overwrite page-data
	page.WidthMax("42rem")
	page.NoNavigation = false
	page.SuppressInProgressbar = true

	page.Label = trl.S{
		"de": fmt.Sprintf("Dynamic page example %v", gn),
		"en": fmt.Sprintf("Dynamic page example %v", gn),
	}
	page.Short = trl.S{
		"de": fmt.Sprintf("Dynamic<br>page %v", gn),
		"en": fmt.Sprintf("Dynamic<br>page %v", gn),
	}

	// dynamically recreate the groups
	page.Groups = nil

	// gr1
	{

		lblMain := trl.S{
			"en": fmt.Sprintf(`lbl main %v - lbl main lbl main lbl main`, gn),
			"de": fmt.Sprintf(`lbl main %v - lbl main lbl main lbl main`, gn),
		}

		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 1
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.Label = lblMain
			inp.ColSpan = 1
			inp.ColSpanLabel = 1
		}
		{
			inp := gr.AddInput()
			inp.Type = "text"
			inp.Name = fmt.Sprintf("text%v", gn)
			inp.Label = trl.S{
				"en": "label input",
				"de": "label input",
			}
			inp.ColSpan = 1
			inp.ColSpan = 1
			inp.ColSpanControl = 1
			inp.MaxChars = 40
		}

	}

	return nil

}
