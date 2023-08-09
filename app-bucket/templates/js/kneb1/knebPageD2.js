function showHideSubQuestions(showHide, classOfSubQuestion) {

    let containers = document.getElementsByClassName(classOfSubQuestion);
    // console.log(`selecting ${classOfSubQuestion} yields ${containers.length}`);
    if (containers.length > 0) {
        // console.log(`containers[${subRadio}] style is ${containers[0].style.display}`);
        if (showHide) {
            containers[0].style.display = ""; // not block
            containers[0].style.marginBottom = "3rem"
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

    let triggeringIDs1 = ['qd7_employmentabove35hours', 'qd7_employmentbetween15and35hours']
    let any1 = false
    triggeringIDs1.forEach(triggeringID => {
        let radio =  document.getElementById(triggeringID);
        if (!radio) {
            console.error(`does not exist: ${triggeringID}`);        
        }
        // console.log(`showD7a called -${rad1.value}-`);
        let radioChecked = radio.checked;
        console.log(`inp1 checked ${radioChecked}: `);
        any1 = any1 || radioChecked
    });
    showHideSubQuestions(any1, "{{.pageID}}-grp03");


    let triggeringIDs2 = ['qd7_employmentupto15hours','qd7_employmentoccasionally', 'qd7_employmentnone']
    let any2 = false
    triggeringIDs2.forEach(triggeringID => {
        let radio =  document.getElementById(triggeringID);
        if (!radio) {
            console.error(`does not exist: ${triggeringID}`);        
        }
        // console.log(`showD7a called -${rad1.value}-`);
        let radioChecked = radio.checked;
        console.log(`inp1 checked ${radioChecked}: `);
        any2 = any2 || radioChecked
    });
    showHideSubQuestions(any2, "{{.pageID}}-grp02");



    // {
    //     let rad1 = document.forms.frmMain.qd7_employmentabove35hours;
    //     if (!rad1) {
    //         console.error(`does not exist: document.forms.frmMain.qd7_employmentabove35hours`);        
    //     }
    //     // console.log(`showD7a called -${rad1.value}-`);
    //     let rad1Checked = rad1.checked;
    //     console.log(`inp1 checked ${rad1Checked}: `);
    //     showHideSubQuestions(rad1Checked, "{{.pageID}}-grp03");
    // }


    return true;

}



// addEventListener is cumulative
window.addEventListener("load", function (event) {

    if (document.forms.frmMain.qd7_employmentnone) {
   
        document.forms.frmMain.qd7_employmentabove35hours.addEventListener('change',showD7a);
        document.forms.frmMain.qd7_employmentbetween15and35hours.addEventListener('change',showD7a);
   
        document.forms.frmMain.qd7_employmentupto15hours.addEventListener('change',showD7a);
        document.forms.frmMain.qd7_employmentoccasionally.addEventListener('change',showD7a);
        document.forms.frmMain.qd7_employmentnone.addEventListener('change',showD7a);
   
        console.log("handlers assigned");
    }
    

    showD7a();
    const first = document.getElementById("qd7_employmentnone");
    // first.focus();
});