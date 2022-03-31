package fmt

import (
	"github.com/zew/go-questionnaire/pkg/qst"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func special202204(q *qst.QuestionnaireT) error {

	cond := false
	cond = cond || q.Survey.Year == 2022 && q.Survey.Month == 4
	if !cond {
		return nil
	}

	//
	//
	//
	{
		page := q.AddPage()
		page.Label = trl.S{
			"de": "Sonderfrage: Klimawandel",
			"en": "Special: Climate change",
		}
		page.Short = trl.S{
			"de": "Sonderfrage:<br>Klimawandel",
			"en": "Special:<br>Climate change",
		}
		page.WidthMax("36rem")
		page.WidthMax("48rem")

		// gr0
		{
			gb := qst.NewGridBuilderRadios(
				columnTemplate6,
				// labelsImproveDeteriorateSectoral(),
				labelsStronglyPositiveStronglyNegativeInfluence(),
				[]string{
					"climatech_sec_banks",
					"climatech_sec_insur",
					"climatech_sec_cars",
					"climatech_sec_chemi",
					"climatech_sec_steel",
					"climatech_sec_elect",
					"climatech_sec_mecha",
					"climatech_sec_consu",
					"climatech_sec_const",
					"climatech_sec_utili",
					"climatech_sec_servi",
					"climatech_sec_telec",
					"climatech_sec_infor",
				},
				radioVals6,
				rowLabelsSectors,
			)
			gb.MainLabel = trl.S{
				"de": `<b>3.</b> 
					<u>Die folgende Sonderfrage dient als Teil einer weiterführenden wissenschaftlichen Arbeit.</u>
					<br>
					Ein möglicher negativer Nebeneffekt der neuen geopolitischen Lage, 
					die von Russlands Krieg gegen die Ukraine herbeigeführt wurde, 
					ist, dass die Klimapolitik an Priorität verlieren könnte. 
					Was glauben Sie, wie wirkt sich der Klimawandel 
					im Allgemeinen auf die langfristigen Ertragschancen 
					von Unternehmen aus den folgenden Sektoren aus?
				`,
				"en": `<b>3.</b> 
					<u>The following special question serves as part of a continuing scientific work.</u>
					<br>

					 A possible negative side effect of the new geopolitical situation brought about by Russia's war against Ukraine is that climate policy may receive a lower priority. 
					 How will climate change in general affect the long-term earnings prospects of firms belonging to the following sectors?

				`,
			}
			gr := page.AddGrid(gb)
			gr.OddRowsColoring = true
		}

	}
	return nil

}
