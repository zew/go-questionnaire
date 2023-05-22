package pds

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/cfg"
	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

// Create questionnaire
func Create(s qst.SurveyT) (*qst.QuestionnaireT, error) {

	ctr.Reset()

	q := qst.QuestionnaireT{}
	q.Survey = s
	q.LangCodes = []string{"en"} // governs default language code
	// q.LangCode = "en"

	q.Survey.Org = trl.S{
		"en": "ZEW",
		"de": "ZEW",
	}
	q.Survey.Name = trl.S{
		"en": "Private Debt Survey",
		"de": "Private Debt Survey",
	}
	// q.Variations = 1

	// page0
	{
		page := q.AddPage()
		page.ValidationFuncName = ""

		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true

		page.Label = trl.S{
			"en": "Dear Madam / Sir,",
			"de": "Sehr geehrter Damen und Herren",
		}
		// page.Short = trl.S{
		// 	"en": "Greeting",
		// 	"de": "Begrüßung",
		// }

		page.WidthMax("42rem")

		// gr0
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 3
			{
				inp := gr.AddInput()
				inp.Type = "dyn-textblock"
				inp.DynamicFunc = "RenderStaticContent"
				inp.DynamicFuncParamset = "./welcome-page.html"
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}
		}

	}

	// page0a - consent 1
	{
		page := q.AddPage()
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		// page.Label = trl.S{
		// 	"en": "Dear Madam / Sir,",
		// 	"de": "Sehr geehrter Damen und Herren",
		// }

		consent(
			qst.WrapPageT(page),
			1,
		)

	}

	// page1 - asset classes
	{
		page := q.AddPage()
		// page.SuppressInProgressbar = true

		page.SuppressProgressbar = true

		page.ValidationFuncName = "pdsPage1"

		page.Label = trl.S{
			"en": "Identification and asset classes",
			"de": "Identification and asset classes",
		}
		page.Short = trl.S{
			"en": "Asset Class Selection,<br>Tranches",
			"de": "Asset Class Selection,<br>Tranches",
		}
		page.CounterProgress = "-"
		// https://www.fileformat.info/info/charset/UTF-8/list.htm?start=2048
		page.CounterProgress = "௵"
		page.CounterProgress = "᎒" // e18e92

		// https://utf8-icons.com/white-square-containing-black-small-square-9635
		page.CounterProgress = "&#9632;" // black square; https://utf8-icons.com/black-square-9632

		page.WidthMax("42rem")

		// gr1
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 1
			{
				inp := gr.AddInput()
				inp.Type = "text"
				inp.Name = "q01_identification"
				inp.MaxChars = 32
				inp.Placeholder = trl.S{
					"en": "name of manager",
					"de": "Name Manager",
				}
				inp.Label = trl.S{
					"en": "Identification",
					"de": "Identifikation",
				}
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
				inp.ColSpanControl = 2
			}
		}

		/*
			if false {
				// gr2
				radiosSingleRow(
					qst.WrapPageT(page),
					"q02_teamsize",
					lblMain,
					mCh2,
				)
			}
		*/

		//
		// gr2
		/*
			<span style='font-size: 80%;'>
			 &nbsp;&nbsp;&nbsp;&nbsp;
			<a href='#' onclick='checkSome();' >For testing: Check some</a>
			</span>

			<span style='font-size: 80%;'>
			 &nbsp;&nbsp;&nbsp;&nbsp;
			<a href='#' onclick='checkAll();' >Check all</a>
			</span>

		*/
		{
			lblMain := trl.S{
				"en": `Which asset classes do you invest in?

			<span style='font-size: 80%;'>
			 &nbsp;&nbsp;&nbsp;&nbsp;
			<a href='#' onclick='checkSome();' >For testing: Check some</a>
			</span>

			<span style='font-size: 80%;'>
			 &nbsp;&nbsp;&nbsp;&nbsp;
			<a href='#' onclick='checkAll();' >Check all</a>
			</span>


					`,
				"de": `Wählen Sie Ihre Assetklassen.
				`,
			}
			checkBoxCascade(
				qst.WrapPageT(page),
				lblMain,
			)

		}

	}

	/*
		for i := 0; i < 3; i++ {

			naviCondition := fmt.Sprintf("pds_ac%v", i+1)

			// page11
			{
				page := q.AddPage()
				page.GeneratorFuncName = fmt.Sprintf("pdsPage11-ac%v", i+1)
				page.NavigationCondition = naviCondition
			}
			// page12
			{
				page := q.AddPage()
				page.GeneratorFuncName = fmt.Sprintf("pdsPage12-ac%v", i+1)
				page.NavigationCondition = naviCondition
			}
			// page21
			{
				page := q.AddPage()
				page.GeneratorFuncName = fmt.Sprintf("pdsPage21-ac%v", i+1)
				page.NavigationCondition = naviCondition
			}
			// // page23
			// {
			// 	page := q.AddPage()
			// 	page.GeneratorFuncName = fmt.Sprintf("pdsPage23-ac%v", i+1)
			// 	page.NavigationCondition = naviCondition
			// }
			// page3
			{
				page := q.AddPage()
				page.GeneratorFuncName = fmt.Sprintf("pdsPage3-ac%v", i+1)
				page.NavigationCondition = naviCondition
			}
			// page4
			{
				page := q.AddPage()
				page.GeneratorFuncName = fmt.Sprintf("pdsPage4-ac%v", i+1)
				page.NavigationCondition = naviCondition
			}

		}

	*/

	// page3
	{
		page := q.AddPage()
		page.GeneratorFuncName = fmt.Sprintf("pdsPage3-ac%v", "x")
		// page.NavigationCondition = naviCondition
	}
	// page4
	{
		page := q.AddPage()
		page.GeneratorFuncName = fmt.Sprintf("pdsPage4-ac%v", "x")
		// page.NavigationCondition = naviCondition
	}

	// page6 - finish
	{
		page := q.AddPage()
		page.Label = trl.S{
			"en": "Consent",
			"de": "Abschluss<br><br>",
		}
		page.Short = trl.S{
			"en": "Finish",
			"de": "DSGVO",
		}
		page.SuppressInProgressbar = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		consent(
			qst.WrapPageT(page),
			2,
		)

		/*
			// gr1
			{
				labels := []trl.S{
					{
						"en": `
						I hereby agree that the answers I have given as part of the ZEW Private Debt Survey will be passed on to Prime Capital AG in non-anonymized form. No personal data is passed on to Prime Capital AG, only the company name and the answers given in the survey. Prime Capital AG will not pass on the data received and will only process this data within the scope of the business activities of the Investment Advisory & Solutions division, in particular for the purposes of database enrichment in the context of manager selection or for the purpose of deriving capital market assumptions.
						`,
					},

					{
						"en": `
						I do <i>not</i> agree with my data being forwarded to Prime Capital AG in non-anonymized form.
						`,
					},
				}
				radioValues := []string{
					"datasharing_yes",
					// "datasharing_anonymous",
					"datasharing_not",
				}

				gr := page.AddGroup()
				gr.Cols = 1
				gr.BottomVSpacers = 2
				{
					inp := gr.AddInput()
					inp.Type = "textblock"
					inp.Label = trl.S{
						"en": `
						Declaration of consent to forward answers in non-anonymized form to Prime Capital&nbsp;AG:
						<br> <!-- vertical space for the must error message -->
						<br> <!-- vertical space for the must error message -->

					`,
					}
					inp.ColSpan = gr.Cols
				}

				for idx, label := range labels {
					rad := gr.AddInput()
					rad.Type = "radio"
					rad.Name = "q62_sharing"
					rad.ValueRadio = radioValues[idx]

					rad.ColSpan = 1
					rad.ColSpanLabel = 1
					rad.ColSpanControl = 6

					rad.Label = label

					rad.ControlFirst()
					rad.ControlTop()

					rad.Validator = "mustRadioGroup"

				}
			}
		*/

		// gr2
		{
			gr := page.AddGroup()
			gr.Cols = 1
			gr.BottomVSpacers = 2
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Style.Desktop.StyleGridContainer.GapRow = "0.2rem"
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{"de": "Kommentar zur Umfrage: ", "en": "Comment on the survey: "}
				inp.Label = trl.S{
					"de": "Wollen Sie uns noch etwas mitteilen?",
					"en": "Any remarks or advice for us?",
				}
				inp.ColSpanLabel = 1
			}
			{
				inp := gr.AddInput()
				inp.Type = "textarea"
				inp.Name = "q63_comment"
				inp.MaxChars = 300
				inp.ColSpanLabel = 0
				inp.ColSpanControl = 1
			}
		}

		// gr3
		{
			gr := page.AddGroup()
			gr.Style = css.NewStylesResponsive(gr.Style)
			gr.Cols = 2
			gr.Style.Desktop.StyleGridContainer.TemplateColumns = "3fr 1fr"
			// gr.Width = 80

			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = cfg.Get().Mp["finish_save_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanLabel = 1
			}

			{
				inp := gr.AddInput()
				inp.Type = "button"
				inp.Name = "submitBtn"
				inp.Response = fmt.Sprintf("%v", len(q.Pages)-1+1) // +1 since one page is appended below
				inp.Label = cfg.Get().Mp["end"]
				inp.Label = cfg.Get().Mp["finish_questionnaire"]
				inp.ColSpan = 1
				inp.ColSpanControl = 1
				inp.AccessKey = "n"
				inp.StyleCtl = css.NewStylesResponsive(inp.StyleCtl)
				inp.StyleCtl.Desktop.StyleGridItem.JustifySelf = "end"
				// inp.StyleCtl.Desktop.StyleBox.WidthMin = "8rem" // does not help with button
			}
		}

		// pge.ExampleSixColumnsLabelRight()

	}

	//
	//
	// Report of results
	{
		page := q.AddPage()
		page.NoNavigation = true
		page.SuppressProgressbar = true
		page.WidthMax("40rem")

		page.Label = trl.S{
			"de": "Ihre Eingaben sind gespeichert.",
			"en": "Your entries have been saved.",
		}
		{
			gr := page.AddGroup()
			gr.Cols = 1
			{
				inp := gr.AddInput()
				inp.Type = "textblock"
				inp.Label = trl.S{
					"en": `&nbsp;`,
					"de": `&nbsp;`,
				}
			}
			// {
			// 	inp := gr.AddInput()
			// 	inp.Type = "dyn-textblock"
			// 	inp.DynamicFunc = "RepsonseStatistics"
			// }
		}
	}

	q.Hyphenize()
	q.ComputeMaxGroups()
	q.SetColspans()

	if err := (&q).TranslationCompleteness(); err != nil {
		return &q, err
	}
	if err := (&q).Validate(); err != nil {
		return &q, err
	}
	return &q, nil
}
