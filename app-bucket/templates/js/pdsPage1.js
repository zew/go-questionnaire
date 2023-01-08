const allCheckboxes = [
    "ac1_q03",
    "ac2_q03",
    "ac3_q03",

    "ac1_tt1_q031",
    "ac1_tt2_q031",
    "ac1_tt3_q031",

    "ac2_tt1_q031",
    "ac2_tt2_q031",

    "ac3_tt1_q031",
    "ac3_tt2_q031",
];



// for fast checking
function checkAll() {
    for (let i0 = 0; i0 < allCheckboxes.length; i0++) {
        let elID = allCheckboxes[i0];
        let checkBx = document.getElementById(elID);
        if (checkBx) {
            checkBx.checked = !checkBx.checked;
            const evt = new Event("change");
            checkBx.dispatchEvent(evt);
        }
    }
}




const someCheckboxes = [
    "ac1_q03",
    "ac2_q03",
    "ac3_q03",

    "ac1_tt1_q031",
    "ac1_tt2_q031",
    "ac1_tt3_q031",

    "ac2_tt1_q031",
    "ac2_tt2_q031",

    "ac3_tt1_q031",
    // "ac3_tt2_q031",
];
function checkSome() {
    for (let i0 = 0; i0 < someCheckboxes.length; i0++) {
        let elID = someCheckboxes[i0];
        let checkBx = document.getElementById(elID);
        if (checkBx) {
            checkBx.checked = !checkBx.checked;
            const evt = new Event("change");
            checkBx.dispatchEvent(evt);
        }
    }
}




// non global block
{
    let frm = document.forms.frmMain;

    const triggers = [
        "ac1_q03",
        "ac2_q03",
        "ac3_q03",
    ];

    // safari cannot read it inside func, if declared let or const
    var parentContainers = {
        "ac1_q03" : "pg01-grp02",
        "ac2_q03" : "pg01-grp03",
        "ac3_q03" : "pg01-grp04",
    };

    UNUSED_subInputs = [
        " ... inp00 should remain visible",
        "pg01-grp04-inp01",
        "pg01-grp04-inp02",
        "pg01-grp04-inp03",
        "pg01-grp04-inp04",
        "pg01-grp04-inp05",
        "pg01-grp04-inp06",
        "pg01-grp04-inp07",
        "..."
    ];

    function showHide(keyID, paramShowHide) {

        let parentContainer = parentContainers[keyID];

        // we start at index 1, for inp00 is the master control
        for (let i0 = 1; i0 < 25; i0++) {
            
            let i0Str = i0.toString();
            if (i0Str.length == 1) {
                i0Str = "0" + i0Str;
            }
            const className = parentContainer + "-inp" + i0Str;
            // console.log("classname", className) // for example pg01-grp04-inp11
            
            // getting div elements, having this classname...
            let containers = document.getElementsByClassName(className);
            if (containers.length > 0) {
                let container = containers[0];
                // overriding server side generated CSS class value for display
                // for i.e. pg01-grp02-inp04 { display: grid }
                // by setting an element style
                if (paramShowHide) {
                    container.style.display = "grid"; // not block
                    // console.log("  shown");
                } else {
                    container.style.display = "none";
                    // console.log("  hidden");
                }
            } else {
                // console.log("className ", className, " is the last plus one");
                break;
            }
        }


    }

    function myOnchange(evt) {
        // console.log("myChange", evt.srcElement.name, evt.srcElement.checked);
        showHide(evt.srcElement.id, evt.srcElement.checked);
        return true;
    }


    for (let i0 = 0; i0 < triggers.length; i0++) {
        const elID = triggers[i0];
        let checkBx = document.getElementById(elID);
        if (checkBx){
            // checkB[0].addEventListener('change', myChange);
            checkBx.addEventListener('change', myOnchange);
            console.log("onchange ", checkBx.name, "'pdsPage1'");
        }        
    }


    let initPage = (inst) => {
        // const evt = new Event("input");
        const evt = new Event("change");
        for (let i0 = 0; i0 < triggers.length; i0++) {
            const elID = triggers[i0];
            let checkBx = document.getElementById(elID);
            if (checkBx) {
                // checkB[0].addEventListener('change', myChange);
                checkBx.dispatchEvent(evt);
            }
        }
    }

    // init checkbox subgroups show/hide;
    window.addEventListener('load', initPage, false);


}
