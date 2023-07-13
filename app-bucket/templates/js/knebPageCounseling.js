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

function showB2(event) {
    let rad1 = document.forms.frmMain.qb1_pensionadvice5;
    if (!rad1) {
        console.error(`does not exist: document.forms.frmMain.qb1_pensionadvice`);        
    }
    // console.log(`showD7a called -${rad1.value}-`);
    let rad1Checked = rad1.checked;
    console.log(`inp1 checked ${rad1Checked}: `);
    showHideSubQuestions( !rad1Checked, "pg12-grp01");

    return true;
}

function showB7(event) {
    let rad1 = document.forms.frmMain.qb6_delegate5;
    let rad2 = document.forms.frmMain.qb6_delegate6;
    let rad3 = document.forms.frmMain.qb6_delegate7;
    let rad4 = document.forms.frmMain.qb6_delegate8;
    let rad5 = document.forms.frmMain.qb6_delegate9;
    let rad6 = document.forms.frmMain.qb6_delegate10;
    let rad7 = document.forms.frmMain.qb6_delegate11;

    // console.log(`showD7a called -${rad1.value}-`);
    let rad1Checked =  rad2.checked  || rad3.checked || rad4.checked || rad5.checked || rad6.checked|| rad7.checked ;
    console.log(`inp2 checked ${rad1Checked}: `);
    showHideSubQuestions( rad1Checked, "pg12-grp06");

    return true;
}


if (document.forms.frmMain.qb1_pensionadvice1) {
    document.forms.frmMain.qb1_pensionadvice1.addEventListener('change', showB2);
    document.forms.frmMain.qb1_pensionadvice2.addEventListener('change', showB2);
    document.forms.frmMain.qb1_pensionadvice3.addEventListener('change', showB2);
    document.forms.frmMain.qb1_pensionadvice4.addEventListener('change', showB2);
    document.forms.frmMain.qb1_pensionadvice5.addEventListener('change', showB2);
    console.log("handlers assigned 1");
}

if (document.forms.frmMain.qb6_delegate1) {
    document.forms.frmMain.qb6_delegate1.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate2.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate3.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate4.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate5.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate6.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate7.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate8.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate9.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate10.addEventListener('change', showB7);
    document.forms.frmMain.qb6_delegate11.addEventListener('change', showB7);
    console.log("handlers assigned 2");
}

// addEventListener is cumulative
window.addEventListener("load", function (event) {
    showB2();
    showB7();
    const first = document.getElementById("qb1_pensionadvice1");
    // first.focus();
});