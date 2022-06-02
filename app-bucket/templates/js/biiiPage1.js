function showHideSubRadios(showHide) {

    const subRadiosClasses = 
    [
        "pg02-grp00-inp03", 
        "pg02-grp00-inp04", 
        "pg02-grp00-inp05", 
        "pg02-grp00-inp06", 
        "pg02-grp00-inp07"
    ];
    for (const subRadio of subRadiosClasses) {
        // let containers = document.querySelectorAll(querySelect); // would be OR selection
        let selector = subRadio + " grid-item-lvl-1";   // AND selection
        let containers = document.getElementsByClassName(selector);
        // console.log(`selecting ${selector} yields ${containers.length}`);
        if (containers.length > 0) {
            // console.log(`containers[${subRadio}] style is ${containers[0].style.display}`);
            if (showHide) {
                containers[0].style.display = ""; // not block
                console.log(`containers[${subRadio}] shown`);
            } else {
                containers[0].style.display = "none";
                console.log(`containers[${subRadio}] hidden`);
            }
        } else {
            console.error(`no elements found for ${selector}`);
        }
    }    

}

function showQ4a(event) {



    // let radio1Val = document.forms.frmMain.q04now.value;
    let radio1Checked = document.forms.frmMain.q04now.checked;
    // console.log("inp1: ", radio1Checked);

    showHideSubRadios(radio1Checked);

    return true;

}


if (document.forms.frmMain.q04now) {
    document.forms.frmMain.q04now.addEventListener('change',showQ4a);
    document.forms.frmMain.q04in_future.addEventListener('change', showQ4a);
    document.forms.frmMain.q04in_planning.addEventListener('change', showQ4a);
    document.forms.frmMain.q04no.addEventListener('change', showQ4a);
    console.log("handlers assigned");
}

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    showQ4a();
    const first = document.getElementById("q04now");
    first.focus();
});