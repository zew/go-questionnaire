package cpfmt

import (
	"fmt"
	"log"
	"strings"

	qstif "github.com/zew/go-questionnaire/pkg/qstif"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func selectInp02(langCode, inpName, inpValue string) string {

	ret := ""

	if langCode == "de" {
		ret = fmt.Sprintf(`
		<select name="%v" type="select">
			<option value="" >bitte wählen</option>
			<option value="1">stark positiv korreliert</option>
			<option value="2">leicht positiv korreliert</option>
			<option value="3">unkorreliert</option>
			<option value="4">leicht negativ korreliert</option>
			<option value="5">stark negativ korreliert</option>
			<option value="6">keine Angabe</option>
        </select>	
		`,
			inpName,
		)

	} else {
		ret = fmt.Sprintf(`
		<select name="%v" type="select">
			<option value="" >please select</option>
			<option value="1">strongly positively correlated</option>
			<option value="2">slightly positively correlated</option>
			<option value="3">uncorrelated</option>
			<option value="4">slightly negatively correlated</option>
			<option value="5">strongly negatively correlated</option>
			<option value="6">no answer</option>
        </select>	
		`,
			inpName,
		)
	}

	anchor := fmt.Sprintf(`value="%v"`, inpValue)
	ret = strings.ReplaceAll(ret, anchor, anchor+" selected")

	return ret
}

func Special202403(q qstif.Q, seq0to5, paramSetIdx int, preflight bool) (string, []string, error) {

	inpNames := []string{
		"qs1_automotive",
		"qs1_industrials",
		"qs1_construction",
		"qs1_utilities",
	}

	if preflight {
		return "", inpNames, nil
	}

	inpValues := make([]string, 0, len(inpNames))
	for _, n := range inpNames {
		v, err := q.ResponseByName(n)
		if err != nil {
			log.Printf("could not find input name %v: %v", n, err)
		}
		inpValues = append(inpValues, v)
	}

	// if q.UserIDInt() < -100*1000 {
	// 	return "", []string{}, nil
	// }

	headerLbls := []trl.S{
		{
			"de": `2030`,
			"en": `2030`,
		},
		{
			"de": `2040`,
			"en": `2040`,
		},
		{
			"de": `2050`,
			"en": `2050`,
		},
		{
			"de": `after 2050`,
			"en": `nach  2050`,
		},
		{
			"de": `2030`,
			"en": `2030`,
		},
	}

	rowLbls := []trl.S{
		{
			"de": `Aktien               <br>Eurogebiet`,
			"en": `stocks               <br>(euro area)`,
		},
		{
			"de": `Staatsanleihen       <br>Eurogebiet`,
			"en": `sovereign bonds      <br>(euro area)`,
		},
		{
			"de": `Unternehmensanleihen <br>Eurogebiet`,
			"en": `corporate bonds      <br>(euro area)`,
		},
	}

	lc := q.GetLangCode()

	sb := &strings.Builder{}

	fmt.Fprint(sb, `
		<style>

			table.table-special-202303 {
				width:     99%;
				max-width: 75rem;
				min-width: 60rem;

				margin: 0.1rem auto;

				background-color: transparent;
				/* border: 1px solid blanchedalmond; */

			}


			table.table-special-202303,
			table.table-special-202303 td,
			dummy {
				border-collapse: collapse;
				border: 1px solid black;
				border: none;
			}

			table.table-special-202303 td,
			dummy {
				border: 2px solid var(--clr-sec);
			}


			table.table-special-202303 td {
				text-align:     center;
				vertical-align: middle;

				padding: 0.4rem 0.2rem;
				width: 18.5%;
			}

			table.table-special-202303 td:first-child {
				width: 18%;
				text-align: right;
				padding-right: 0.7rem;
			}

			table.table-special-202303 td:nth-child(2) {
				width: 12%;
			}




			table.table-special-202303 tr:first-child td {
				background-color: var(--clr-sec-lgt1);
			}
			table.table-special-202303 td:first-child {
				background-color: var(--clr-sec-lgt1);
			}


		</style>

	`)

	tplContainer := `
    <table border="border: 1px solid grey" class="table-special-202303" >
        <tr>    
            <td> <span style="font-size: 85%%">  %v  </span> </td>
            <td> %v </td>
            <td> %v </td>
            <td> %v </td>
            <td> %v </td>
        </tr>

        <tr>    
            <td>%v</td>

            <td> - </td>
            <td> %v </td>
            <td> %v </td>
            <td> %v </td>

        </tr>

        <tr>    
            <td>%v</td>

            <td> - </td>
            <td> - </td>
            <td> %v </td>
            <td> %v </td>
        </tr>

        <tr>    
            <td>%v</td>

            <td> - </td>
            <td> - </td>
            <td> - </td>
            <td> %v </td>

        </tr>


    </table>

	`

	fmt.Fprintf(

		sb,
		tplContainer,

		headerLbls[0].Tr(lc),
		headerLbls[1].Tr(lc),
		headerLbls[2].Tr(lc),
		headerLbls[3].Tr(lc),
		headerLbls[4].Tr(lc),

		rowLbls[0].Tr(lc),
		selectInp02(lc, inpNames[0], inpValues[0]),
		selectInp02(lc, inpNames[1], inpValues[1]),
		selectInp02(lc, inpNames[2], inpValues[2]),

		rowLbls[1].Tr(lc),
		selectInp02(lc, inpNames[3], inpValues[3]),
		selectInp02(lc, inpNames[4], inpValues[4]),

		rowLbls[2].Tr(lc),
		selectInp02(lc, inpNames[5], inpValues[5]),
	)

	return sb.String(), inpNames, nil

}
