package cpfmt

import (
	"fmt"
	"strings"

	qstif "github.com/zew/go-questionnaire/pkg/qst/compositeif"
	"github.com/zew/go-questionnaire/pkg/trl"
)

var columnTemplateLocal = []float32{
	4.0, 1,
	0.0, 1,
	0.0, 1,
	0.0, 1,
	0.0, 1,
	0.5, 1,
}

var radioVals6 = []string{"1", "2", "3", "4", "5", "6"}

func Special202212Q3(q qstif.Q, seq0to5, paramSetIdx int) (string, []string, error) {

	inpNames := []string{
		"qs3a_inf_narrative_a",
		"qs3a_inf_narrative_b",
		"qs3a_inf_narrative_c",
		"qs3a_inf_narrative_d",
		"qs3a_inf_narrative_e",
		"qs3a_inf_narrative_f",
	}

	rowLbls := []trl.S{
		{
			"de": `Eine Entspannung bei der Inflationsentwicklung, eine weniger restriktive Geldpolitik der EZB und nachlassende Rezessionsrisiken wirken sich
					<i>positiv</i>
					auf das Rendite-Risiko-Profil in 2023 aus.`,
			"en": `An easing in the development of inflation development, a less restrictive monetary stance by the ECB and diminishing recession risks have a
					<i>positive</i>
					impact on the return-risk-profile in 2023.`,
		},
		{
			"de": `Den DAX-Konzernen gelingt es auch weiterhin, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen unverändert bleiben oder sogar steigen, was sich
					<i>positiv</i>
					auf das Rendite-Risiko-Profil des DAX in 2023 auswirkt. `,
			"en": `DAX companies will continue to succeed in passing on their rising costs to their customers. The profit margins of DAX companies will therefore remain unchanged or even increase, which has a
					<i>positive</i>
					impact on the return-risk-profile of the DAX in 2023.`,
		},
		{
			"de": `Die Entwicklung der Inflation spielt für das Rendite-Risiko-Profil des DAX in 2023
					<i>keine Roll</i>e
					.`,
			"en": `The development of inflation does
					<i>not impact</i>
					the return-risk-profile of the DAX.`,
		},
		{
			"de": `	<i>Positive</i>
					und
					<i>negative</i>
					Effekte der Inflation gleichen sich aus. Die Entwicklung der Inflation ist daher insgesamt
					<i>neutral</i>
					für das Rendite-Risiko-Profil des DAX in 2023.`,
			"en": `
					<i>Positive</i>
					and
					<i>negative</i>
					effects of inflation cancel each other out. Overall, the development of inflation is
					<i>neutral</i>
					for the return-risk-profile of the DAX in 2023.`,
		},
		{
			"de": `Den DAX-Konzernen gelingt es nicht, ihre steigenden Kosten an ihre Kunden weiterzugeben. Die Gewinn-Margen der DAX-Konzerne werden deswegen fallen, was sich
					<i>negativ</i>
					auf das Rendite-Risiko-Profil des DAX in 2023 auswirkt.`,
			"en": `DAX companies will not to succeed in passing on their rising costs to their customers. The profit margins of DAX companies will therefore decrease, which has a
					<i>negative</i>
					impact on the return-risk-profile of the DAX in 2023.`,
		},
		{
			"de": `Anhaltend hohe Inflationsraten, weitere Zinserhöhungen durch die EZB und zunehmende Rezessionsrisiken wirken sich
					<i>negativ</i>
					auf das Rendite-Risiko-Profil des DAX in 2023 aus.
						`,
			"en": `Persistently high inflation rates, further interest rate hikes by the ECB and increasing recession risks will have a
					<i>negative</i>
					impact on the return-risk-profile of the DAX in 2023.`,
		},
	}

	if q.UserIDInt() < -100*1000 {
		s := ""
		inpNames = []string{}
		return s, inpNames, nil
	}

	sb := &strings.Builder{}

	tplContainer := `
	<div  class='pg07-grp03 grid-container '  >
	%v
	</div>
	`

	tplCell := `
		<div  class='pg07-grp03-inp09 grid-item-lvl-1'  >
			%v
			%v
		</div>
	`

	tplLabel := `
			<label  for="%v" >
				%v
			</label>
	`

	tplRadio := `
		<div  class='pg07-grp03-inp07-ctl grid-item-lvl-2' >
			<input type='radio'
				name='%v'
				id='%v_%v'
				value='%v'
				%v
			/>
		</div>
	`

	for i1 := 0; i1 < len(inpNames); i1++ {

		resp, err := q.ResponseByName(inpNames[i1])
		if err != nil {
			// generators.Create() and qst.Load() for new user
			//  are calling qst.Validate()
			//   => dynamic fields do not exist yet
		}
		if resp == "input-name-not-found" {
			resp = ""
		} else if resp != "" {
			//
		}

		for i2 := 0; i2 < 6; i2++ {

			checked := ""
			if resp == fmt.Sprintf("%v", i2+1) {
				checked = ` checked="checked" `
			}

			label := ""
			if i2 == 0 {
				label = fmt.Sprintf(
					tplLabel,
					// placeholder vals
					inpNames[i1],
					rowLbls[i1][q.GetLangCode()],
				)
			}
			radio := fmt.Sprintf(
				tplRadio,
				// placeholder vals
				inpNames[i1],       // input name
				inpNames[i1], i2+1, // id
				i2+1,    // radio val
				checked, // checked or not
			)

			fmt.Fprintf(
				sb,
				tplCell,
				label,
				radio,
			)

		}

	}

	// wrap container
	ret := fmt.Sprintf(
		tplContainer,
		sb.String(),
	)

	return ret, inpNames, nil

}
