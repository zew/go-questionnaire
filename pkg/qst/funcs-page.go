package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Notice, that these func are called
// on each request
type funcPageGeneratorT func(*QuestionnaireT, *pageT) error

var pageGens = ctr.New() // page generations

var funcPGs = map[string]funcPageGeneratorT{

	"pageFuncExample": pageFuncExample,

	"fmt202212": fmt202212,

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

func fmt202212(q *QuestionnaireT, page *pageT) error {

	page.Groups = nil

	var radioVals4 = []string{"1", "2", "3", "4"}

	//
	// gr1 - q4a
	{

		colLblsQ4 := []trl.S{
			{
				"de": "Meinen eigenen Analysen",
				"en": "My own analyses",
			},
			{
				"de": "Analysen von Experten/-innen aus meinem Unternehmen",
				"en": "Analyses by experts in my company",
			},
			{
				"de": "Analysen aus externen Quellen",
				"en": "Analyses from external sources",
			},

			{
				"de": "keine<br>Angabe",
				"en": "no answer",
			},
		}

		var columnTemplateLocal = []float32{
			4.0, 1,
			0.0, 1,
			0.0, 1,
			0.5, 1,
		}
		gb := NewGridBuilderRadios(
			columnTemplateLocal,
			colLblsQ4,
			[]string{
				"qs4a_growth",
				"qs4a_inf",
				"qs4a_dax",
			},
			radioVals4,
			[]trl.S{
				{
					"de": `Wirtschaftswachstum Deutschland`,
					"en": `GDP growth, Germany`,
				},
				{
					"de": `Inflation in Deutschland`,
					"en": `Inflation, Germany`,
				},
				{
					"de": `Entwicklung des DAX`,
					"en": `Developments of the DAX`,
				},
			},
		)

		gb.MainLabel = trl.S{
			"de": `
						Meine Einschätzungen mit Blick auf die folgenden Bereiche beruhen hauptsächlich auf
					`,
			"en": `
						My expectations with respect to the following areas are mainly based on
					`,
		}.Outline("4a.")

		gr := page.AddGrid(gb)
		_ = gr
	}

	//
	// gr2 - q4b
	{
		mainLbl4b := trl.S{
			"de": `Wie relevant sind die Prognosen der Bundesbank für Ihre eigenen Inflationsprognosen für Deutschland?`,
			"en": `How relevant are the inflation forecasts of Bundesbank for your own inflation forecasts for Germany?`,
		}.Outline("4b.")

		colLbls4b := []trl.S{
			{
				"de": "nicht relevant",
				"en": "not relevant",
			},
			{
				"de": "leicht relevant",
				"en": "slightly relevant",
			},
			{
				"de": "stark relevant",
				"en": "highly relevant",
			},

			{
				"de": "keine<br>Angabe",
				"en": "no answer",
			},
		}

		var columnTemplateLocal = []float32{
			5.0, 1,
			0.0, 1,
			0.0, 1,
			0.5, 1,
		}
		gb := NewGridBuilderRadios(
			columnTemplateLocal,
			colLbls4b,
			[]string{
				"qs4b_relevance",
			},
			radioVals4,
			[]trl.S{
				mainLbl4b,
			},
		)

		gr := page.AddGrid(gb)
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapColumn = "1.2rem"
		gr.BottomVSpacers = 3
		_ = gr
	}

	//
	// cutoff
	uid := q.UserIDInt()
	grp, ok := fmtRandomizationGroups[uid]

	if ok && grp < 7 {
		// show rest
	} else {
		return nil
	}

	{
		gr := page.AddGroup()
		gr.Cols = 1
		gr.BottomVSpacers = 2
		{
			inp := gr.AddInput()
			inp.Type = "textblock"
			inp.ColSpan = gr.Cols
			inp.ColSpanLabel = 1
			inp.Label = trl.S{
				"de": `
					Bundesbankpräsident Joachim Nagel äußert sich regelmäßig zum Inflationsausblick für Deutschland. Im November 2022 äußerte er sich folgendermaßen: "Auch im kommenden Jahr dürfte die Inflationsrate in Deutschland hoch bleiben. Ich halte es für wahrscheinlich, dass im Jahresdurchschnitt 2023 eine sieben vor dem Komma stehen wird".
						`,
				"en": `
					Bundesbank president Joachim Nagel regularly comments on the inflation outlook for Germany. In November 2022, he commented as follows: "The inflation rate in Germany is likely to remain high in the coming year. I believe it is likely that the annual average for 2023 will have a seven before the decimal point."
					`,
			}

		}
	}

	//
	// gr3 - q4c
	{

		colLbls4c := []trl.S{
			{
				"de": "ja",
				"en": "yes",
			},
			{
				"de": "nein",
				"en": "no",
			},
			{
				"de": "keine<br>Angabe",
				"en": "no answer",
			},
		}

		var columnTemplateLocal = []float32{
			5.0, 1,
			0.0, 1,
			0.5, 1,
		}

		lbl1 := trl.S{
			"de": `
					War Ihnen die Aussage von Bundesbankpräsident Joachim Nagel bereits bekannt?
						`,
			"en": `
					Were you aware of this statement by Bundesbank president Joachim Nagel?
					`,
		}.Outline("4c.")

		gb := NewGridBuilderRadios(
			columnTemplateLocal,
			colLbls4c,
			[]string{
				"qs4c_known",
			},
			radioVals4,
			[]trl.S{
				lbl1,
			},
		)

		// gb.MainLabel =

		gr := page.AddGrid(gb)
		gr.Style = css.NewStylesResponsive(gr.Style)
		gr.Style.Desktop.StyleGridContainer.GapColumn = "1.2rem"
	}

	return nil

}

var fmtRandomizationGroups = map[int]int{
	9990: 1,
	9991: 2,
	9992: 2,
	9993: 3,
	9995: 5,
	9996: 6,
	9997: 7,
	9998: 8,

	10001: 12,
	10003: 12,
	10008: 8,
	10009: 9,
	10012: 2,
	10013: 3,
	10014: 4,
	10015: 5,
	10017: 10,
	10021: 7,
	10023: 11,
	10025: 1,
	10027: 1,
	10033: 11,
	10034: 11,
	10035: 12,
	10040: 5,
	10041: 12,
	10042: 1,
	10051: 6,
	10052: 12,
	10054: 9,
	10056: 5,
	10057: 4,
	10058: 6,
	10060: 12,
	10062: 7,
	10063: 12,
	10070: 7,
	10071: 6,
	10072: 2,
	10073: 12,
	10078: 12,
	10079: 7,
	10080: 12,
	10081: 8,
	10084: 3,
	10085: 2,
	10086: 11,
	10089: 4,
	10090: 3,
	10095: 12,
	10096: 1,
	10098: 4,
	10103: 4,
	10105: 3,
	10109: 3,
	10112: 4,
	10113: 2,
	10115: 12,
	10117: 5,
	10118: 4,
	10127: 9,
	10128: 6,
	10129: 12,
	10131: 2,
	10133: 5,
	10134: 11,
	10138: 1,
	10140: 4,
	10143: 12,
	10146: 7,
	10147: 4,
	10150: 11,
	10154: 5,
	10160: 6,
	10161: 4,
	10162: 2,
	10163: 7,
	10165: 10,
	10167: 12,
	10168: 6,
	10172: 11,
	10173: 3,
	10175: 7,
	10178: 4,
	10179: 4,
	10180: 5,
	10185: 2,
	10197: 3,
	10199: 9,
	10205: 3,
	10209: 11,
	10210: 6,
	10218: 8,
	10224: 3,
	10228: 11,
	10231: 6,
	10235: 1,
	10247: 4,
	10256: 7,
	10260: 5,
	10261: 10,
	10262: 5,
	10263: 5,
	10267: 1,
	10268: 9,
	10274: 9,
	10278: 3,
	10281: 7,
	10286: 7,
	10296: 10,
	10297: 1,
	10305: 3,
	10307: 3,
	10310: 11,
	10311: 12,
	10315: 2,
	10316: 1,
	10319: 11,
	10321: 10,
	10322: 9,
	10330: 9,
	10337: 7,
	10339: 9,
	10343: 8,
	10344: 9,
	10345: 8,
	10349: 4,
	10356: 1,
	10361: 6,
	10362: 11,
	10363: 11,
	10364: 6,
	10366: 9,
	10367: 4,
	10369: 9,
	10372: 1,
	10374: 10,
	10376: 3,
	10377: 10,
	10381: 8,
	10385: 9,
	10387: 8,
	10388: 9,
	10391: 6,
	10405: 3,
	10415: 10,
	10418: 8,
	10420: 11,
	10421: 4,
	10794: 9,
	10802: 5,
	10806: 12,
	10807: 6,
	10812: 7,
	10813: 3,
	10816: 8,
	10821: 5,
	10825: 10,
	10826: 6,
	10828: 3,
	10829: 12,
	10830: 8,
	10834: 7,
	10837: 8,
	10839: 2,
	10842: 11,
	10844: 5,
	10947: 1,
	11241: 6,
	11242: 1,
	11246: 11,
	11251: 8,
	11270: 7,
	11272: 10,
	11275: 6,
	11345: 1,
	11389: 7,
	11396: 11,
	11398: 6,
	11400: 11,
	11403: 8,
	11425: 5,
	11426: 2,
	11435: 7,
	11440: 6,
	11445: 2,
	11452: 7,
	11453: 11,
	11465: 5,
	11466: 7,
	11475: 7,
	11482: 10,
	11490: 1,
	11495: 5,
	11497: 6,
	11499: 10,
	11500: 8,
	11504: 10,
	11506: 7,
	11514: 1,
	11534: 12,
	11544: 1,
	11546: 6,
	11547: 9,
	11548: 8,
	11556: 3,
	11558: 2,
	11559: 2,
	11563: 2,
	11565: 3,
	11568: 2,
	11569: 3,
	11572: 2,
	11573: 9,
	11574: 2,
	11575: 3,
	11579: 10,
	11588: 4,
	11589: 10,
	11593: 10,
	11594: 10,
	11596: 10,
	11598: 9,
	11599: 9,
	11600: 11,
	11601: 10,
	11602: 5,
	11603: 5,
	11604: 6,
	11605: 4,
	11606: 4,
	11607: 11,
	11608: 9,
	11614: 12,
	11616: 1,
	11617: 8,
	11618: 7,
	11619: 1,
	11620: 4,
	11621: 8,
	11622: 8,
	11623: 8,
	11625: 3,
	11627: 2,
	11628: 10,
	11630: 8,
	11631: 7,
	11632: 12,
	11633: 10,
	11634: 6,
	11635: 7,
	11636: 5,
	11637: 8,
	11638: 9,
	11639: 4,
	11640: 11,
	11642: 1,
	11643: 2,
	11644: 12,
	11645: 5,
	11646: 5,
	11648: 2,
	11649: 2,
	11650: 5,
	11651: 1,
	11652: 7,
	11653: 9,
	11654: 10,
	11655: 9,
	11657: 5,
	11658: 8,
	11659: 9,
	11660: 6,
	11661: 10,
	11662: 12,
	11665: 4,
	11668: 3,
	11669: 4,
	10007: 2,
	10016: 8,
	10219: 1,
	10220: 11,
	10412: 3,
	11348: 6,
	11477: 1,
	11478: 11,
	11610: 2,
}
