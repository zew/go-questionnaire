const names = [
    "q0x_1",
    "q0x_2",
    "q0x_3",
    "q0x_4",
    "q0x_5",
    "q0x_6",
    "q0x_7",
    "q0x_8",
];


function checkAll(inpSrc) {

    let nonEmpty = 0;

    for (const name of names) {
        let inp = document.forms.frmMain[name]
        console.log(`   name: ${inp.name} - val is ${inp.value}`)
        if (inp.value !== "") {
            nonEmpty++;
        }
    }    
    console.log(`   non empty  ${nonEmpty}`)

    if (nonEmpty > 5) {
        alert("Maximal fünf");
        inpSrc.value = "";
    }

}

function checkOneToFive(event) {
    checkAll(event.target);
    // console.log(`name: ${event.target.name} - val is ${event.target.value}`)
    return true;
}


// addEventListener is cumulative
window.addEventListener("load", function (event) {
    for (const inp of names) {
        document.forms.frmMain[inp].addEventListener('change', checkOneToFive);
    }
    console.log("handlers assigned");
});