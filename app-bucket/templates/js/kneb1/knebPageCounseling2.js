// const pageID = "pg13-grp01";
const pageID = "{{.pageID}}-grp03";

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


function showB7(event) {
    let rad7 = document.forms.frmMain.qb6_delegate7;
    let rad8 = document.forms.frmMain.qb6_delegate8;
    let rad9 = document.forms.frmMain.qb6_delegate9;
    let rad10 = document.forms.frmMain.qb6_delegate10;
    let rad11 = document.forms.frmMain.qb6_delegate11;

    // console.log(`showD7a called -${rad1.value}-`);
    let effective =  rad7.checked || rad8.checked || rad9.checked || rad10.checked|| rad11.checked ;
    console.log(`effective checked? ${effective}: `);
    showHideSubQuestions( effective, pageID);

    return true;
}



// addEventListener is cumulative
window.addEventListener("load", function (event) {

    
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

    showB7();
   
});