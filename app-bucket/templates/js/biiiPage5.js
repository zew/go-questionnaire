const names = [
    "q10a_pct",
    "q10b_pct",
    "q10c1_pct",
    "q10c2_pct",
    "q10d_pct",
    "q10e_pct",
];



function requireSiblingEntry(event) {

    // console.log(`name: ${event.target.name} - val is ${event.target.value}`)
    let inpSrc = event.target;    
    let ln = inpSrc.name.length;

    console.log(`inpSrc name is ${inpSrc.name} - length ${ln}`);

    siblingName = inpSrc.name.substr(0, ln-4);
    // console.log(`sibling name is ${siblingName}`);

    let sibling = document.forms.frmMain[siblingName];

    if (sibling.value === "") {
        alert("Bitte erst Euro-Betrag eintragen");
        inpSrc.value = "";
        return false;
    }

    return true;
}

function submitCheck(event) {
    console.log(`submit check start`);
    
    let sum = 0;
    for (const name of names) {
        let inp = document.forms.frmMain[name];
        // console.log(`${inp.name} - value has type `,typeof inp.value);
        // console.log(`${inp.name} - value has value `, Number(inp.value));

        sum += Number(inp.value);
        // console.log(`${inp.name} - val is -${Number(inp.value)}- - ${sum}`);
    }

    console.log(`sum is ${sum}`);

    if (sum != 0 || true) {

        if (sum < 99.999 || sum > 100.001) {
            alert("Ergibt nicht 100 Prozent");
            event.preventDefault();
            return false;        
        }

    }

    return true;


}

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    for (const inp of names) {
        document.forms.frmMain[inp].addEventListener('change', requireSiblingEntry);
    }

    document.forms.frmMain.addEventListener('submit', submitCheck);
    // submitCheck(null);

    console.log("handlers assigned");
});