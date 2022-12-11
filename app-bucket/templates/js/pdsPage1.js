function showHide(keyName, showHide) {

    const subRadiosClasses = {
        "q03_ac1_corplending": [
            "pg01-grp04-inp01",
            "pg01-grp04-inp02",
            "pg01-grp04-inp03",
            "pg01-grp04-inp04",
            "pg01-grp04-inp05",
            "pg01-grp04-inp06",
            "pg01-grp04-inp07",
        ],
        "q03_ac2_realestate":[
            "pg01-grp04-inp09",
            "pg01-grp04-inp10",
            "pg01-grp04-inp11",
            "pg01-grp04-inp12",
            "pg01-grp04-inp13",
        ],
        "q03_ac3_infrastruct":[
            "pg01-grp04-inp15",
            "pg01-grp04-inp16",
            "pg01-grp04-inp17",
            "pg01-grp04-inp18",
            "pg01-grp04-inp19",
        ],
    } 
    ;

    let arrSel = subRadiosClasses[keyName];

    for (let index = 0; index < arrSel.length; index++) {
        const className = arrSel[index];
        // console.log("classname", className)
        let containers = document.getElementsByClassName(className);
        if (containers.length > 0) {
            let container = containers[0];
            if (showHide) {
                container.style.display = ""; // not block
                // console.log("  shown");
            } else {
                container.style.display = "none";
                // console.log("  hidden");
            }
        } else {
            console.error("no elements found for ", className);
        }
    }


}

function myChange(evt) {
    // console.log("myChange", evt.srcElement.name, evt.srcElement.checked);
    showHide(evt.srcElement.name, evt.srcElement.checked);
    return true;
}


// non global block
{
    let frm = document.forms.frmMain;

    if (frm) {
        if (frm.q03_ac1_corplending[0]) {
            frm.q03_ac1_corplending[0].addEventListener('change',myChange);
            console.log("handler 'pdsPage1' assigned");
        }
        if (frm.q03_ac2_realestate[0]) {
            frm.q03_ac2_realestate[0].addEventListener('change', myChange);
            console.log("handler 'pdsPage1' assigned");
        }
        if (frm.q03_ac3_infrastruct[0]) {
            frm.q03_ac3_infrastruct[0].addEventListener('change', myChange);
            console.log("handler 'pdsPage1' assigned");
        }
    }

    let înitPage = (inst) => {
        // const evt = new Event("input");
        const evt = new Event("change");
        frm.q03_ac1_corplending[0].dispatchEvent(evt);
        frm.q03_ac2_realestate[0].dispatchEvent(evt);
        frm.q03_ac3_infrastruct[0].dispatchEvent(evt);
    }

    // init slider;
    // triggers oninput above
    window.addEventListener('load', înitPage, false);


}




