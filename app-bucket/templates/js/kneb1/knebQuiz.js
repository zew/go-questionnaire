/*
    two funcs to make entry into the two number inputs of the knebel "Quiz" pages
    more convenient.

    Regular keyboard entry of the dot key (".") is suppressed.

    Pasting a number from clipboard, such as "48.800 " is cleansed of dots and trimmed
    and then entered

*/

// const pageID = "{.pageID}-grp03";


const cleanseKeypress = evt => {
    if (evt.charCode == 46){
        console.log(`no dots '.'`);
        evt.preventDefault();
    }
}

// https://stackoverflow.com/questions/2176861
const cleansePaste = evt => {
    let dataClip, dataPasted;

    // get pasted data - clipboard API
    dataClip = evt.clipboardData || window.clipboardData;
    dataPasted = dataClip.getData('Text');

    dataPasted = dataPasted.replace(".", "");
    dataPasted = dataPasted.trim();

    console.log(`cleansed pasted '${dataPasted}'`);
    evt.srcElement.value = dataPasted;

    // stop data actually being pasted
    // at the *end* of method  =>  if above fails, propagation will *continue*
    evt.stopPropagation();
    evt.preventDefault();
}



// addEventListener is cumulative
window.addEventListener("load", function (event) {

    inps = [
        "qc24_nf_return",
        "qc26_area_nf",

        "qc24_ff_return",
        "qc26_area_ff",
    ]

    try {
        inps.forEach(inpID => {

            const inp = document.getElementById(inpID);
            if (inp) {
                inp.addEventListener('keypress', cleanseKeypress);
                inp.addEventListener('paste',    cleansePaste);
                console.log(`event handlers for cleansing dots ${inpID} success`)
            }

        });

    } catch (err) {
        console.error("attaching event handlers to cleanse dots", err)
    }
    console.log("paste and dot-replacer installed ")



});