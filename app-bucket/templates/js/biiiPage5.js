const names = [
    "q10a_pct",
    "q10b_pct",
    "q10c1_pct",
    "q10c2_pct",
    "q10d_pct",
    "q10e_pct",
];



function checkOneToFive(event) {

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


// addEventListener is cumulative
window.addEventListener("load", function (event) {
    for (const inp of names) {
        document.forms.frmMain[inp].addEventListener('change', checkOneToFive);
    }
    console.log("handlers assigned");
});