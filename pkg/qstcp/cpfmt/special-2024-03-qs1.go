package cpfmt

import (
	"fmt"
	"strings"

	qstif "github.com/zew/go-questionnaire/pkg/qstif"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func fourRadios(q qstif.Q, inpName, mod string) string {

	// tpl := `
	// <input type="radio" name="inpname_mod"   value="1" />
	// <input type="radio" name="inpname_mod"   value="2" />
	// <input type="radio" name="inpname_mod"   value="3" />
	// <input type="radio" name="inpname_mod"   value="4" />

	// `

	sb := &strings.Builder{}
	for i := 1; i <= 4; i++ {
		inm := fmt.Sprintf("%v_%v", inpName, mod) // input name multiplied
		checked := ""
		vl, err := q.ResponseByName(inm)
		if err == nil && fmt.Sprintf("%v", i) == vl {
			checked = "checked"
		}
		fmt.Fprintf(
			sb,
			"	<input type=\"radio\" name=\"%v\"   value=\"%v\"   %v />\n",
			inm,
			i,
			checked,
		)
	}

	return sb.String()

}

func Special202403QS1(q qstif.Q, seq0to5, paramSetIdx int, modePreflight bool) (string, []string, error) {

	inpNames := []string{
		"qs1_automotive",
		"qs1_industr", // industrials gets hyphenated
		"qs1_construction",
		"qs1_utilities",
	}
	mods := []string{
		"2030",
		"2040",
		"2050",
		"2050_after",
	}

	//
	//
	inpNamesMult := []string{} // input names multiplied
	for _, inp := range inpNames {
		for _, mod := range mods {
			inpNamesMult = append(inpNamesMult, fmt.Sprintf("%v_%v", inp, mod))
		}
		inpNamesMult = append(inpNamesMult, inp+"_noaw")
	}

	if modePreflight {
		return "", inpNamesMult, nil
	}

	rowLbls := []trl.S{
		{
			"de": `Fahrzeugbau`,
			"en": `Automotive`,
		},
		{
			"de": `Industrieunternehmen <ssmall>(Chemie, Pharma, Stahl, NE-Metalle, Elektro, Maschinenbau)</ssmall>`,
			"en": `Industrials <ssmall>(Chemicals, Pharma, Steel, Metal Products, Electronics, Machinery)</ssmall>`,
		},
		{
			"de": `Baugewerbe`,
			"en": `Construction`,
		},
		{
			"de": `Versorger  <ssmall>(e.g. Elektrizität, Gas, Wasser)</ssmall>`,
			"en": `Utilities  <ssmall>(e.g. electricity, gas, water)</ssmall>`,
		},
	}

	lc := q.GetLangCode()

	sb := &strings.Builder{}

	//
	//
	fmt.Fprint(sb, `
	<style>


	input[type=radio]:focus {
		box-shadow: none;
	}
	

    .tbl-1 {
        width: 99%;
        max-width: 75rem;
        min-width: 60rem;

        margin: 0.1rem auto;

        background-color: transparent;

 
    }

    table.tbl-1 td,
    dummy {

        margin:  0;
        padding: 0;

        border: 2px solid var(--clr-sec);

        border: 1px solid black;
        border-collapse: collapse;
        border: none;

    }

    table.tbl-1 td {
        text-align: center;
        vertical-align: middle;
        
        padding: 0.4rem 0.2rem;
        width: 21%;
    }
    table.tbl-1 tr:first-child td {
        vertical-align:  bottom;
    }

    


    /* first and last column */
    table.tbl-1 td:first-child {
        width: 11%;
        text-align: right;
        text-align: left;
        padding-right: 0.4rem;
    }

    table.tbl-1 td:last-child {
        width: 4%;
    }
    /* second column */
    table.tbl-1  td:nth-child(2) {
        /* background-color: chartreuse; */
    }


    /* contents */
    table.tbl-1 td .hdr{
        font-size: 85%;
        font-size: 92%;
    }


    table.tbl-1 td input[type=radio],
    table.tbl-1 td .hdr,
    dummy
    {
        display: inline-block;
        width: 22%;
        width: 20%;
        width: 18%;
        margin: 0;
        padding-left:  0.02rem;
        /* padding-right: 1.2rem; */
        border: 1px solid red;
        border: none;
    }

    table.tbl-1 td input:last-child,
    table.tbl-1 td .hdr:last-child
    {
        padding-right: 0.02rem;
    }

    


</style>

	`)

	tblStart := `
	<table class="tbl-1">

    <tr>
        <td> &nbsp; </td>
        <td> 2030 </td>
        <td> 2040 </td>
        <td> 2050 </td>
        <td> after 2050 </td>
        <td> no answer </td>
    </tr>
    <tr>
        <td> &nbsp; </td>
        <td> 
            <span class="hdr">--</span>
            <span class="hdr">-</span>
            <span class="hdr">+</span>
            <span class="hdr">++</span>
        </td>
        <td> 
            <span class="hdr">--</span>
            <span class="hdr">-</span>
            <span class="hdr">+</span>
            <span class="hdr">++</span>
        </td>
        <td> 
            <span class="hdr">--</span>
            <span class="hdr">-</span>
            <span class="hdr">+</span>
            <span class="hdr">++</span>
        </td>
        <td> 
            <span class="hdr">--</span>
            <span class="hdr">-</span>
            <span class="hdr">+</span>
            <span class="hdr">++</span>
        </td>
        <td> &nbsp; </td>
    </tr>

	`

	if lc == "de" {
		tblStart = strings.ReplaceAll(tblStart, "after", "nach")
		tblStart = strings.ReplaceAll(tblStart, "no answer", "keine Ang.")
	}

	fmt.Fprintf(
		sb,
		tblStart,
	)

	rowTpl := `
	<tr>
		<td> rowLabel </td>
		<td>
			<input type="radio" name="inpname_2030"   value="1" />
			<input type="radio" name="inpname_2030"   value="2" />
			<input type="radio" name="inpname_2030"   value="3" />
			<input type="radio" name="inpname_2030"   value="4" />
		</td>
		<td>
			<input type="radio" name="inpname_2040"   value="1" />
			<input type="radio" name="inpname_2040"   value="2" />
			<input type="radio" name="inpname_2040"   value="3" />
			<input type="radio" name="inpname_2040"   value="4" />
		</td>
		<td>
			<input type="radio" name="inpname_2050"   value="1" />
			<input type="radio" name="inpname_2050"   value="2" />
			<input type="radio" name="inpname_2050"   value="3" />
			<input type="radio" name="inpname_2050"   value="4" />
		</td>
		<td>
			<input type="radio" name="inpname_2050_after"   value="1" />
			<input type="radio" name="inpname_2050_after"   value="2" />
			<input type="radio" name="inpname_2050_after"   value="3" />
			<input type="radio" name="inpname_2050_after"   value="4" />
		</td>
		<td>
			<input type="checkbox" name="inpname_noaw">
		</td>
	</tr>
			
	`
	_ = rowTpl

	for rowIdx, inp := range inpNames {
		fmt.Fprint(sb, "<tr>\n")
		fmt.Fprintf(sb, "	<td> %v</td>\n", rowLbls[rowIdx].Tr(lc))
		for _, mod := range mods {
			fmt.Fprintf(sb, "	<td> %v</td>\n", fourRadios(q, inp, mod))
		}

		//
		inpNA := fmt.Sprintf("%v_noaw", inp)
		checked := ""
		vl, err := q.ResponseByName(inpNA)
		if err == nil && vl != "" {
			checked = "checked"
		}
		fmt.Fprintf(sb, "	<td><input type=\"checkbox\" name=\"%v_noaw\" %v>\n", inp, checked)
		fmt.Fprint(sb, "</tr>\n")
	}

	//
	// close table
	fmt.Fprint(
		sb,
		"</table>",
	)

	return sb.String(), inpNames, nil

}
