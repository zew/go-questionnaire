const names = [
    "q12_paris_align",
    "q12_culture_sports",
    "q12_education",
    "q12_work",
    "q12_research",
    "q12_health",
    "q12_social_service",
    "q12_environment_land",
    "q12_environment_sea",
    "q12_sanitary",
    "q12_agriculture",
    "q12_energy",
    "q12_residential",
    "q12_technology",
    "q12_prodution",
    "q12_urban_dev",
    "q12_microfinance",
    "q12_other",
    "q12_na",
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