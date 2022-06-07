const names = [
    "q12_paris_align",
    "q12_culture_sports",
    "q12_education",
    "q12_work",
    "q12_research",
    "q12_health",
    "q12_social_service",
    "q12_environment",
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
        let inp = document.forms.frmMain[name]
        // console.log(`   name: ${inp.name} - val is ${inp.value}`)
        if (inp.value !== "") {
            nonEmpty++;

            let key = inp.value;
            try {
                key = key.trim();
            } catch (error) {
                // not critical
            }
            taken[key]++;
            if (taken[key] > 1) {
                alert(`${inp.value} schon gewählt.`);
                inpSrc.value = "";
                return;
            }
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