function validateForm(event) {

    let nameInputs = [
        "q2_seq1_row1_rad",
        "q2_seq1_row2_rad",
        "q2_seq1_row3_rad",

        "q2_seq2_row1_rad",
        "q2_seq2_row2_rad",
        "q2_seq2_row3_rad",
    ];

    let strInputs = ["", "", "", "", "", ""];
    let intInputs = [0, 0, 0, 0, 0, 0];


    for (var i1 = 0, lenght1 = nameInputs.length; i1 < lenght1; i1++) {
        var radioName = nameInputs[i1];
        var radios1 = document.getElementsByName(radioName);
        for (var i2 = 0, lenght2 = radios1.length; i2 < lenght2; i2++) {
            if (radios1[i2].checked) {
                strInputs[i1] = radios1[i2].value;
                if (strInputs[i1] != "") {
                    intInputs[i1] = parseInt(strInputs[i1], 10);
                }
                break;
            }
        }
    }

    console.log(strInputs);
    console.log(intInputs);
    // event.preventDefault(); // not only return false - but also preventDefault()
    // return false;


    var sum =  intInputs[0] + intInputs[1] + intInputs[2] + intInputs[3] + intInputs[4] + intInputs[5];
    var sum1 = intInputs[0] + intInputs[1] + intInputs[2] ;
    var sum2 = intInputs[3] + intInputs[4] + intInputs[5];
    console.log(`sum ${sum} - sum1 ${sum1} - sum2 ${sum2}`,);

    if (sum > 0) {
        if (intInputs[0] == 0 || intInputs[1] == 0 || intInputs[3] == 0 || intInputs[3] == 0 || intInputs[4] == 0 || intInputs[5] == 0 ) {
            // alert("{{.msg}}");
            var doContinue = window.confirm("{{.msg}}");
            if (doContinue) {
                return true;
            }
            event.preventDefault(); // not only return false - but also preventDefault()
            return false;
        }
    }

    if (sum1 == 6 || sum2 == 6) {
        // alert("{{.msg}}");
        var doContinue = window.confirm("Wollen Sie wirklich alle Optionen auf 'nicht verfügbar setzen' oder wollen Sie Ihre Antworten noch verändern?");
        if (doContinue) {
            return true;
        }
        event.preventDefault(); // not only return false - but also preventDefault()
        return false;
    }

    return true;

}


var frm = document.forms.frmMain;
if (frm) {
    frm.addEventListener('submit', validateForm);
}
