function validateForm(event) {


    let nameInputs = [
        "dec7_q1",
        "dec7_q2",

        "dec8_q1",
        "dec8_q2",
    ];

    var allEmpty = true;
    let strInputs = ["", "",   "", ""];
    let intInputs = [ 0,  0,    0,  0];


    for (var i1 = 0, length1 = nameInputs.length; i1 < length1; i1++) {
        // var radioName = "q1_seq" + (i1 + 1) + "_r"  // q1_seq1_r
        var radioName = nameInputs[i1];
        var radios1 = document.getElementsByName(radioName);
        for (var i2 = 0, length2 = radios1.length; i2 < length2; i2++) {
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


    var sum1 = intInputs[0] != "" && intInputs[1] != "";
    var sum2 = intInputs[2] != "" && intInputs[2] != "";
    console.log(`sum1 ${sum1}, sum2 ${sum2}, `);

    if (sum1 || sum2 ) {
        if (sum1 && strInputs[0] == strInputs[1] || sum2 && strInputs[2] == strInputs[3]) {
            // alert("{{.msg}}");
            var doContinue = window.confirm("{{.msg}}");
            if (doContinue) {
                return true;
            }
            event.preventDefault(); // not only return false - but also preventDefault()
            return false;
        }
    }

    return true;

}


var frm = document.forms.frmMain;
if (frm) {
    frm.addEventListener('submit', validateForm);
}
