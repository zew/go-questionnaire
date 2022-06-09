const names = [
    "q33_paris",
    "q33_leisure",
    "q33_education",
    "q33_work",

    "q33_research",
    "q33_health",
    "q33_social_services",
    "q33_env_protection",

    "q33_oceans",
    "q33_wash",
    "q33_agriculture",
    "q33_energy",

    "q33_residential",
    "q33_it",
    "q33_production",
    "q33_urban_dev",

    "q33_financial_access",
    "q33_other",
];

const columnSuffixes = [
    "_need",
    "_pot",
];

function checkAll(inpSrc) {

    for (const suffix of columnSuffixes) {

        let nonEmpty = 0;

        for (const name of names) {
            let inp = document.forms.frmMain[name+suffix][0]; // checkboxes have an identially named hidden input
            let siblingInput = document.forms.frmMain[inp.name + "_addl"];
            let sibling = siblingInput.closest(".grid-item-lvl-2");
            // let sibling = siblingInput.closest(".grid-item-lvl-1");
            if (inp.value !== "" && inp.checked) {
                nonEmpty++;
                sibling.style.display = "";
            } else {
                sibling.style.display = "none";
            }
        }
        console.log(` col ${suffix}: non empty ${nonEmpty}`)

        if (nonEmpty > 5 && inpSrc) {
            alert("Maximal fünf");
            // inpSrc.value = "";
            inpSrc.checked = false;
        }
    }



}

function checkOneToFive(event) {
    checkAll(event.target);
    // console.log(`name: ${event.target.name} - val is ${event.target.value}`)
    return true;
}


// addEventListener is cumulative
window.addEventListener("load", function (event) {
    for (const suffix of columnSuffixes) {
        for (const name of names) {
            let nameComb = name + suffix;
            console.log(`name combined ${nameComb}`);
            let inp = document.forms.frmMain[nameComb][0]; // checkboxes have an identially named hidden input
            inp.addEventListener('change', checkOneToFive);
        }
    }
    checkAll(null);
    console.log("handlers assigned");
});