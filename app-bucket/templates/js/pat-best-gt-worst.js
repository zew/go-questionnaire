function validateForm(event) {

   

    let nameInputs = [
        "q2_a",
        "q2_b",
        "q2_c",
    ];

    var allEmpty = true;
    let strInputs = ["", "", ""];
    let intInputs = [0, 0, 0];


    for (var i1 = 0, lenght1 = nameInputs.length; i1 < lenght1; i1++) {
        var inpName = nameInputs[i1];
        if (document.getElementById(inpName)) {
            strInputs[i1] = document.getElementById(inpName).value;
            if (strInputs[i1] != "") {
                allEmpty = false;
                intInputs[i1] = parseInt(strInputs[i1], 10);
            }
        }
    }

    console.log(strInputs);
    console.log(intInputs);
    // event.preventDefault(); // not only return false - but also preventDefault()
    // return false;

    var cond1 = intInputs[0] != 0 && intInputs[1] != 0 && intInputs[2] != 0;
    var cond2 = intInputs[0] < intInputs[1];
    var cond3 = intInputs[1] < intInputs[2];


    if (cond1 && (cond2 || cond3)) {
        var doContinue = window.confirm("{{.msg}}");
        if (doContinue) {
            return true;
        }
        event.preventDefault(); // not only return false - but also preventDefault()
        document.getElementById(nameInputs[0]).focus();
        return false;
    }

    return true;

}


var frm = document.forms.frmMain;
if (frm) {
    frm.addEventListener('submit', validateForm);
}
