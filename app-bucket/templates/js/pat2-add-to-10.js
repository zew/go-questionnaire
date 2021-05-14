function validateForm(event) {

    /*   
    let nameInputs = [
        "q4a_opt1",
        "q4a_opt2",
        "q4a_opt3",

        "q4b_opt1",
        "q4b_opt2",
        "q4b_opt3",
    ]; 
    */

    let nameInputs = [
        "part2_q1_q1",
        "part2_q1_q2",
        "part2_q1_q3",

        "part2_q2_q1",
        "part2_q2_q2",
        "part2_q2_q3",

        "part2_q3_q1",
        "part2_q3_q2",
        "part2_q3_q3",

        "part2_q4_q1",
        "part2_q4_q2",
        "part2_q4_q3",

        "part2_q5_q1",
        "part2_q5_q2",
        "part2_q5_q3",

        "part2_q6_q1",
        "part2_q6_q2",
        "part2_q6_q3",
    ];

    var allEmpty = true;
    let strInputs = ["", "", "",   "", "", "",   "", "", "",     "", "", "",   "", "", "",   "", "", ""];
    let intInputs = [0, 0, 0,   0, 0, 0,   0, 0, 0,     0, 0, 0,   0, 0, 0,   0, 0, 0   ];


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



    var sum1 =  intInputs[0] +  intInputs[1] +  intInputs[2];
    var sum2 =  intInputs[3] +  intInputs[4] +  intInputs[5];
    var sum3 =  intInputs[6] +  intInputs[7] +  intInputs[8];
    
    var sum4 =  intInputs[9] + intInputs[10] + intInputs[11];
    var sum5 = intInputs[12] + intInputs[13] + intInputs[14];
    var sum6 = intInputs[15] + intInputs[16] + intInputs[17];

    console.log("sum1", sum1, "sum2", sum2);

    var cond1 = sum1 > 0 && sum1 != 10;
    var cond2 = sum2 > 0 && sum2 != 10;
    var cond3 = sum3 > 0 && sum3 != 10;

    var cond4 = sum4 > 0 && sum4 != 10;
    var cond5 = sum5 > 0 && sum5 != 10;
    var cond6 = sum6 > 0 && sum6 != 10;

    // if (allEmpty || cond1 || cond2) {
    if (cond1 || cond2 || cond3 || cond4 || cond5 || cond6) {
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
