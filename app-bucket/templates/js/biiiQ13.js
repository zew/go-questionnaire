const names = [
    "q13_not",
    "q13_1",
    "q13_2",
    "q13_3",
    "q13_4",
    "q13_5",
    "q13_6",
    "q13_7",
    "q13_8",
    "q13_9",
    "q13_10",
    "q13_11",
    "q13_12",
    "q13_13",
    "q13_14",
    "q13_15",
    "q13_16",
    "q13_17",
];


function checkAll(inpSrc) {

    let nonEmpty = 0;

    let taken = {
        "1":0,
        "2":0,
        "3":0,
        "4":0,
        "5":0,
    } 

    for (const name of names) {
        let inp = document.forms.frmMain[name][0]; // checkboxes have an identially named hidden input
        if (inp.value !== "" && inp.checked) {
            nonEmpty++;

            let key = inp.value;
            try {
                key = key.trim();
            } catch (error) {
                // not critical
            }
            taken[key]++;
            // if (taken[key] > 1) {
            //     alert(`${inp.value} schon gewählt.`);
            //     inpSrc.value = "";
            //     return;
            // }
        }
    }    
    console.log(`   non empty  ${nonEmpty}`)

    if (nonEmpty > 5 && inpSrc) {
        alert("Maximal fünf");
        // inpSrc.value = "";
        inpSrc.checked = false;
    }

}

function checkOneToFive(event) {
    checkAll(event.target);
    // console.log(`name: ${event.target.name} - val is ${event.target.value}`)
    return true;
}


// addEventListener is cumulative
window.addEventListener("load", function (event) {
    for (const name of names) {
        // console.log(`inp name ${name}`);
        let inp = document.forms.frmMain[name][0]; // checkboxes have an identially named hidden input
        // console.log(inp);
        // console.log(`inp name ${inp.Name}`);
        inp.addEventListener('change', checkOneToFive);
    }
    checkAll(null);    
    console.log("handlers assigned");
});