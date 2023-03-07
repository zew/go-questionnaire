package cpfmt

import (
	"fmt"
	"strings"

	qstif "github.com/zew/go-questionnaire/pkg/qstif"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func selectInput(lc, inpName string) string {

	if lc == "de" {
		return fmt.Sprintf(`
		<select name="%v" type="select">
			<option value="">bitte wählen</option>
			<option value="">stark positiv korreliert</option>
			<option value="">leicht positiv korreliert</option>
			<option value="">unkorreliert</option>
			<option value="">leicht negativ korreliert</option>
			<option value="">stark negativ korreliert</option>
			<option value="">keine Angabe</option>
        </select>	
		`,
			inpName,
		)

	}

	return fmt.Sprintf(`
		<select name="%v" type="select">
			<option value="">please select</option>
			<option value="">strongly positively correlated</option>
			<option value="" selected>slightly positively correlated</option>
			<option value="">uncorrelated</option>
			<option value="">slightly negatively correlated</option>
			<option value="">strongly negatively correlated</option>
			<option value="">no answer</option>
        </select>	
	`,
		inpName,
	)
}

func Special202303(q qstif.Q, seq0to5, paramSetIdx int, preflight bool) (string, []string, error) {

	inpNames := []string{
		"ass_stock_bondgovt",
		"ass_stock_bondcorp",
		"ass_stock_realestate",

		"ass_bondgovt_bondcorp",
		"ass_bondgovt_realestate",

		"ass_bond_realestateprv",
	}

	if preflight {
		return "", inpNames, nil
	}

	// if q.UserIDInt() < -100*1000 {
	// 	return "", []string{}, nil
	// }

	headerLbls := []trl.S{
		{
			"de": `Korrelation <br>zwischen Gesamtrenditen`,
			"en": `Correlation <br>between returns`,
		},
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
		{
			"de": `Immobilien           <br>Eurogebiet`,
			"en": `real estate          <br>(euro area)`,
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

			table{
				width: 98%;
				max-width: 75rem;
				min-width: 60rem;
				margin: 1.5rem auto;
				background-color: aliceblue;
				border: 1px solid blanchedalmond;
			}
			
		
			table, th, td {
				border: 1px solid black;
				border-collapse: collapse;
			}    

			td {
				text-align: center;
				vertical-align: middle;


				padding: 0.4rem 0.2rem;
				width:   18.5%;
			}

			td:first-child {
				width:   18%;
				text-align: right;
				padding-right: 0.7rem;
			}

			td:nth-child(2) {
				width:   12%;
			}


		</style>

	`)

	tplContainer := `
    <table border="border: 1px solid grey">
        <tr>    
            <td> %v </td>
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
		selectInput(lc, inpNames[0]),
		selectInput(lc, inpNames[1]),
		selectInput(lc, inpNames[2]),

		rowLbls[1].Tr(lc),
		selectInput(lc, inpNames[3]),
		selectInput(lc, inpNames[4]),

		rowLbls[2].Tr(lc),
		selectInput(lc, inpNames[5]),
	)

	return sb.String(), inpNames, nil

}
