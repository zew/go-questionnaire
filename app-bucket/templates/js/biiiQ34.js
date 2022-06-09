const names = [
    "q34_private_equity",
    "q34_public_equity",
    "q34_private_external",
    "q34_public_external",

    "q34_real_estate",
    "q34_money_deposits",
    "q34_soc",
    "q34_green_bonds",

    "q34_emerging_markets",
    "q34_microfinance",
    "q34_commodities",
    "q34_slb",

    "q34_hybrid",
    "q34_other",

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
        let siblingInput = document.forms.frmMain[inp.name + "_addl"];
        // let sibling = siblingInput.closest(".grid-item-lvl-2");
        let sibling = siblingInput.closest(".grid-item-lvl-1");
        // console.log(inp);
        // console.log(`   name: ${inp.name} - val is ${inp.value}`)
        if (inp.value !== "" && inp.checked) {
            nonEmpty++;

            sibling.style.display = "";


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
        } else {
            sibling.style.display = "none";
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