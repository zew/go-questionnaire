

// non global block
{
    let frm = document.forms.frmMain;

    const triggers = [
        "q03_ac1_corplending",
        "q03_ac2_realestate",
        "q03_ac3_infrastruct",
    ];

    const parentContainers = {
        "q03_ac1_corplending" : "pg01-grp02",
        "q03_ac2_realestate"  : "pg01-grp03",
        "q03_ac3_infrastruct" : "pg01-grp04",
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
        for (let i0 = 1; i0 < 20; i0++) {
            
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
                if (paramShowHide) {
                    container.style.display = ""; // not block
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
        let checkB = document.getElementById(elID);
        if (checkB){
            // checkB[0].addEventListener('change', myChange);
            checkB.addEventListener('change', myOnchange);
            console.log("onchange ", checkB.name, "'pdsPage1'");
        }        
    }


    let initPage = (inst) => {
        // const evt = new Event("input");
        const evt = new Event("change");
        for (let i0 = 0; i0 < triggers.length; i0++) {
            const elID = triggers[i0];
            let checkB = document.getElementById(elID);
            if (checkB) {
                // checkB[0].addEventListener('change', myChange);
                checkB.dispatchEvent(evt);
            }
        }
    }

    // init checkbox subgroups show/hide;
    window.addEventListener('load', initPage, false);


}




