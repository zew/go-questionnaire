function validateForm(event) {

    let nameInputs = [
        "q4a_opt1",
        "q4a_opt2",
        "q4a_opt3",

        "q4b_opt1",
        "q4b_opt2",
        "q4b_opt3",
    ];

    var allEmpty = true;
    let strInputs = ["", "", "", "", "", ""];
    let intInputs = [0, 0, 0, 0, 0, 0];


    for (var i1 = 0, lenght1 = nameInputs.length; i1 < lenght1; i1++) {
        var inpName = nameInputs[i1];
        strInputs[i1] = document.getElementById(inpName).value;
        if (strInputs[i1] != "") {
            allEmpty = false;
            intInputs[i1] = parseInt(strInputs[i1], 10);
        }
    }

    console.log(strInputs);
    console.log(intInputs);
    // event.preventDefault(); // not only return false - but also preventDefault()
    // return false;



    var sum1 = intInputs[0] + intInputs[1] + intInputs[2] ;
    var sum2 = intInputs[3] + intInputs[4] + intInputs[5];

    console.log("sum1", sum1, "sum2", sum2);

    var cond1 = sum1 > 0 && sum1 != 10;
    var cond2 = sum2 > 0 && sum2 != 10;

    // if (allEmpty || cond1 || cond2) {
    if (cond1 || cond2) {
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
