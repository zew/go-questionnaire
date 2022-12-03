package qst

import (
	"fmt"

	"github.com/zew/go-questionnaire/pkg/css"
	"github.com/zew/go-questionnaire/pkg/ctr"
	"github.com/zew/go-questionnaire/pkg/trl"
)

type funcPageGeneratorT func(*QuestionnaireT, *pageT) error

var pageGens = ctr.New() // page generations

var funcPGs = map[string]funcPageGeneratorT{
	"pds01":     pds01,
	"fmt202212": fmt202212,
}

func KeysValuesPageCollect(q *QuestionnaireT, page *pageT) map[string]string {

	// finishes, ks, vs := q.KeysValues(true)
	// _, _, _ = finishes, ks, vs
	ret := map[string]string{}

	cleanse := false
	for i2 := 0; i2 < len(page.Groups); i2++ {
		for i3 := 0; i3 < len(page.Groups[i2].Inputs); i3++ {
			if page.Groups[i2].Inputs[i3].IsLayout() {
				continue
			}
			// keys = append(keys, page.Groups[i2].Inputs[i3].Name)
			val := page.Groups[i2].Inputs[i3].Response
			if cleanse {
				if page.Groups[i2].Inputs[i3].Type == "number" {
					val = DelocalizeNumber(val)
				}
				val = EnglishTextAndNumbersOnly(val)
			}
			key := page.Groups[i2].Inputs[i3].Name
			ret[key] = val
			// vals = append(vals, val)
		}
	}

	return ret
}
func KeysValuesPageApply(q *QuestionnaireT, page *pageT, kv map[string]string) {
	for i2 := 0; i2 < len(page.Groups); i2++ {
		for i3 := 0; i3 < len(page.Groups[i2].Inputs); i3++ {
			if page.Groups[i2].Inputs[i3].IsLayout() {
				continue
			}
			key := page.Groups[i2].Inputs[i3].Name
			page.Groups[i2].Inputs[i3].Response = kv[key]
		}
	}
}

func pds01(q *QuestionnaireT, page *pageT) error {

	kv := KeysValuesPageCollect(q, page)
	defer KeysValuesPageApply(q, page, kv)

	page.NoNavigation = false

	gn := pageGens.Increment()
	lblMain := trl.S{
		"en": fmt.Sprintf(`lbl main %v - lbl main lbl main lbl main`, gn),
		"de": fmt.Sprintf(`lbl main %v - lbl main lbl main lbl main`, gn),
	}

	// if page.Finished.IsZero() {
	page.Groups = nil
	if true {

		// gr1
		{
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

	}

	return nil

}

func fmt202212(q *QuestionnaireT, page *pageT) error {

	kv := KeysValuesPageCollect(q, page)
	defer KeysValuesPageApply(q, page, kv)

	page.Groups = nil
	if page.Finished.IsZero() || true {

		page.NoNavigation = false
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

		uid := q.UserIDInt()
		grp, ok := fmtRandomizationGroups[uid]

		if ok && grp < 7 {
			// show rest
		} else {
			return nil
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
			gr.BottomVSpacers = 1
			_ = gr
		}

		{
			gr := page.AddGroup()
			gr.Cols = 1
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

	}

	return nil

}
