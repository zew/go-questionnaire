function showHideSubQuestions(showHide, classOfSubQuestion) {

    let containers = document.getElementsByClassName(classOfSubQuestion);
    // console.log(`selecting ${classOfSubQuestion} yields ${containers.length}`);
    if (containers.length > 0) {
        // console.log(`containers[${subRadio}] style is ${containers[0].style.display}`);
        if (showHide) {
            containers[0].style.display = ""; // not block
            console.log(`containers[${classOfSubQuestion}] shown`);
        } else {
            containers[0].style.display = "none";
            console.log(`containers[${classOfSubQuestion}] hidden`);
        }
    } else {
        console.error(`no elements found for ${classOfSubQuestion}`);
    }

}

function showD7a(event) {

    {
        let rad1 = document.forms.frmMain.qd7_employmentnone;
        if (!rad1) {
            console.error(`does not exist: document.forms.frmMain.qd7_employmentnone`);        
        }
        // console.log(`showD7a called -${rad1.value}-`);
        let rad1Checked = rad1.checked;
        console.log(`inp1 checked ${rad1Checked}: `);
        showHideSubQuestions(rad1Checked, "pg02-grp04");
    }

    {
        let rad1 = document.forms.frmMain.qd7_employmentabove35hours;
        if (!rad1) {
            console.error(`does not exist: document.forms.frmMain.qd7_employmentabove35hours`);        
        }
        // console.log(`showD7a called -${rad1.value}-`);
        let rad1Checked = rad1.checked;
        console.log(`inp1 checked ${rad1Checked}: `);
        showHideSubQuestions(rad1Checked, "pg02-grp05");
    }


    return true;

}


if (document.forms.frmMain.qd7_employmentnone) {
    document.forms.frmMain.qd7_employmentabove35hours.addEventListener('change',showD7a);
    document.forms.frmMain.qd7_employmentbetween15and35hours.addEventListener('change',showD7a);
    document.forms.frmMain.qd7_employmentupto15hours.addEventListener('change',showD7a);
    document.forms.frmMain.qd7_employmentoccasionally.addEventListener('change',showD7a);
    document.forms.frmMain.qd7_employmentnone.addEventListener('change',showD7a);
    console.log("handlers assigned");
}

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    showD7a();
    const first = document.getElementById("qd7_employmentnone");
    // first.focus();
});