
let evtCounter = 0;

function countToMax(event) {

    evtCounter++;

    let q33aVal1 = 0, q33aVal2 = 0, q33aVal3 = 0;
    let q33bVal1 = 0, q33bVal2 = 0, q33bVal3 = 0;

    var elements = document.forms.frmMain.elements;
    for (var i = 0, element; element = elements[i++];) {

        if (element.type === "hidden") {
            continue;
        }

        if (element.type !== "radio") {
            continue;
        }

        if (element.name.substring(0, 4) == "q33a") {

            if (element.checked) {
                if (element.value == "1") {
                    q33aVal1++;
                }
                if (element.value == "2") {
                    q33aVal2++;
                }
                if (element.value == "3") {
                    q33aVal3++;
                }                
            }

        }

        if (element.name.substring(0, 4) == "q33b") {

            if (element.checked) {
                if (element.value == "1") {
                    q33bVal1++;
                }
                if (element.value == "2") {
                    q33bVal2++;
                }
                if (element.value == "3") {
                    q33bVal3++;
                }
            }

        }
    }

    let invalid = false;

    let tgtName = "undef";
    if (event && event.target) {
        tgtName = event.target.name;        
    }    

    if (q33aVal1 > 3 || q33aVal2 > 3 || q33aVal3 > 3) {
        console.log(`q33a: Maximal drei je Spalte \n  col1 ${q33aVal1} col2 ${q33aVal2}  col3 ${q33aVal3} - ${tgtName} ${evtCounter}`);
        invalid = true;
    }

    if (q33bVal1 > 3 || q33bVal2 > 3 || q33bVal3 > 3) {
        console.log(`q33b: Maximal drei je Spalte \n col1 ${q33bVal1} col2 ${q33bVal2}  col3 ${q33bVal3} - ${tgtName} ${evtCounter}`);
        invalid = true;
    }


    if (invalid) {
        try {
            if (event && event.target) {
                event.target.checked = false;
            }                
        } catch (error) {
            log.error(`cannot uncheck ${event.target}`,error);
        }
        alert("Maximal drei je Spalte");
    }

    // console.log(`q33a col1 ${q33aVal1} col2 ${q33aVal2}  col3 ${q33aVal3}  \nq33b col1 ${q33bVal1} col2 ${q33bVal2}  col3 ${q33bVal3}`);




    return true;
}


if (document.forms.frmMain.q4now) {
    // document.forms.frmMain.q4now.addEventListener('change',doStuff);
    // console.log("handlers assigned");
}

// addEventListener is cumulative
window.addEventListener("load", function (event) {

    countToMax();
    console.log("handler assigned to load");

    let radioCounter = 0;
    var elements = document.forms.frmMain.elements;
    for (var i = 0, element; element = elements[i++];) {
        if (element.type !== "radio") {
            continue;
        }
        radioCounter++;
        // document.forms.frmMain.q4now.addEventListener('change', showQ4a);
        element.addEventListener('change', countToMax );
        if (radioCounter < 5 ) {
            console.log(`handler assigned to el #${i}  - ${element.name}  val ${element.value}`);
        }
    }


    // first.focus();    
});