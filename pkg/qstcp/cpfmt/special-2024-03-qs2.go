package cpfmt

import (
	"fmt"
	"strings"

	qstif "github.com/zew/go-questionnaire/pkg/qstif"
	"github.com/zew/go-questionnaire/pkg/trl"
)

func oneRadioBlock(q qstif.Q, inpName, mod string) string {
	return ""

}

func Special202403QS2(q qstif.Q, seq0to5, paramSetIdx int, modePreflight bool) (string, []string, error) {

	inpNames := []string{
		"qs2_automotive",
		"qs2_industr", // industrials gets hyphenated
		"qs2_construction",
		"qs2_utilities",
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

	fmt.Fprint(sb, `
	<style>
		.tbl-2a {
			width: 99%;
			max-width: 75rem;
			min-width: 60rem;
			margin: 0.1rem auto;
			background-color: transparent;
			margin-top: 1.2rem;
		}

		table.tbl-2a td {
			margin:  0;
			padding: 0;
			border-collapse: collapse;
			border: none;
		}

		table.tbl-2a td {
			vertical-align: top;
		}
		table.tbl-2a td ul {
			margin: 0;
			margin-right: 3rem;
			padding: 0;
		}


	</style>
	`)

	pre := `
	<table class="tbl-2a">

    <tr>
        <td>
            &nbsp;
        </td>
        <td>
            Possible economic benefits
        </td>
        <td>
            Possible economic costs
        </td>
    </tr>
    <tr>
        <td>
            &nbsp;
        </td>
        <td>
            <ul>
                <li>Save transition costs (costs related to the price of CO2, fines, stranded assets, etc.)</li>
                <li>Reputational benefits with customers, employees etc.</li>
                <li>Continued access to capital</li>
                <li>Access to government subsidies</li>
                <li>Transition leads to product innovation and thereby better products</li>
                <li>Transition leads to process innovations and thereby more efficient processes and/or cheaper inputs </li>
                <li>Better protection against physical risks </li>
            </ul>
        </td>
        <td>
            <ul>
                <li>Transition leads to less efficient processes and/or more expensive inputs</li>
                <li>The quality of transition/climate-neutral products is lower than quality of current products</li>
                <li>Less reliability of clean energy sources</li>
            </ul>
        </td>
    </tr>

	</table>	
	
	`

	if lc == "de" {
		pre = `

		<table class="tbl-2a">

		<tr>
			<td>
				&nbsp;
			</td>
			<td>
				Mögliche Vorteile
			</td>
			<td>
				Mögliche Nachteile
			</td>
		</tr>
		<tr>
			<td>
				&nbsp;
			</td>
			<td>
				<ul>
					<li>Einsparung von Übergangskosten (Kosten im Zusammenhang mit dem CO2-Preis, Geldbußen, gestrandete Vermögenswerte usw.)</li>
					<li>Reputationsvorteile bei Kunden, Mitarbeitern usw.</li>
					<li>Zugang zu Kapital und Krediten</li>
					<li>Staatliche Subventionen</li>
					<li>Der Übergang führt zu Produktinnovationen und damit zu besseren Produkten</li>
					<li>Der Übergang führt zu Prozessinnovationen und damit zu effizienteren Prozessen und/oder niedrigeren Inputkosten </li>
					<li>Verbesserter Schutz gegen physische Klimarisiken </li>
				</ul>
			</td>
			<td>
				<ul>
					<li>Der Übergang führt zu weniger effizienten Prozessen und/oder teureren Inputkosten</li>
					<li>Die Qualität von Übergangsprodukten/klimaneutralen Produkten ist geringer als die Qualität der derzeitigen Produkte</li>
					<li>Geringere Zuverlässigkeit von sauberen Energiequellen</li>
				</ul>
			</td>
		</tr>
	
	</table>		
		
		`
	}

	fmt.Fprint(sb, pre)

	//
	//
	fmt.Fprint(sb, `
	<style>


	

    .tbl-2 {
        width: 99%;
        max-width: 75rem;
        min-width: 60rem;

        margin: 0.1rem auto;

        background-color: transparent;

 
    }

    table.tbl-2 td,
    dummy {

        margin:  0;
        padding: 0;

        border: 2px solid var(--clr-sec);

        border: 1px solid black;
        border-collapse: collapse;
        border: none;

    }

    table.tbl-2 td {
        text-align: center;
        vertical-align: middle;
        
        padding: 0.4rem 0.2rem;
        width: 21%;
    }
    table.tbl-2 tr:first-child td {
        vertical-align:  bottom;
    }

    


    /* first and last column */
    table.tbl-2 td:first-child {
        width: 11%;
        text-align: right;
        text-align: left;
        padding-right: 0.4rem;
    }

    table.tbl-2 td:last-child {
        width: 4%;
    }
    /* second column */
    table.tbl-2  td:nth-child(2) {
        /* background-color: chartreuse; */
    }


    /* contents */
    table.tbl-2 td .hdr{
        font-size: 85%;
        font-size: 92%;
    }


    table.tbl-2 td input[type=radio],
    table.tbl-2 td .hdr,
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

    table.tbl-2 td input:last-child,
    table.tbl-2 td .hdr:last-child
    {
        padding-right: 0.02rem;
    }

    


</style>

	`)

	tblStart := `
	<table class="tbl-2">

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
			fmt.Fprintf(sb, "	<td> %v</td>\n", oneRadioBlock(q, inp, mod))
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
