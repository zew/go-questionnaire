// const pageID = "pg13-grp01";
const pageID = "{{.pageID}}-grp01";

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

    let rad1 = document.forms.frmMain.qb1_pensionadvice1;
    let rad2 = document.forms.frmMain.qb1_pensionadvice2;
    let rad3 = document.forms.frmMain.qb1_pensionadvice3;
    let rad4 = document.forms.frmMain.qb1_pensionadvice4;
    let rad5 = document.forms.frmMain.qb1_pensionadvice5;

    if (!rad5) {
        console.error(`does not exist: document.forms.frmMain.qb1_pensionadvice`);        
    }
    // console.log(`showD7a called -${rad1.value}-`);
    
    let effective = !rad5.checked;

    // default: none checked
    if (!rad1.checked && !rad2.checked && !rad3.checked && !rad4.checked && !rad5.checked) {
        effective = false;
    }
    
    console.log(`effective checked ${effective}: `);

    showHideSubQuestions( effective, pageID);

    return true;
}




// addEventListener is cumulative
window.addEventListener("load", function (event) {

    if (document.forms.frmMain.qb1_pensionadvice1) {
        document.forms.frmMain.qb1_pensionadvice1.addEventListener('change', showB2);
        document.forms.frmMain.qb1_pensionadvice2.addEventListener('change', showB2);
        document.forms.frmMain.qb1_pensionadvice3.addEventListener('change', showB2);
        document.forms.frmMain.qb1_pensionadvice4.addEventListener('change', showB2);
        document.forms.frmMain.qb1_pensionadvice5.addEventListener('change', showB2);
        console.log("handlers assigned 1");
    }
    
    showB2();

    const first = document.getElementById("qb1_pensionadvice1");
    // first.focus();
});