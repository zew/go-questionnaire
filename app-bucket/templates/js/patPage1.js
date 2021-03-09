function validateForm(event) {

    for (let i1 = 0; i1 < 3; i1++) {
        var checkName = "q1_seq1_r" + (i1 + 1)  // q1_seq1_r1
        var checks = document.getElementsByName(checkName);
        for (var i2 = 0, length = checks.length; i2 < length; i2++) {
            checks[i2].disabled = true;
        }
    }
    console.log("checks disabled");
    for (let i1 = 0; i1 < 1; i1++) {
        var radioName = "q1_seq" + (i1+1) + "_r"  // q1_seq1_r
        var radios = document.getElementsByName(radioName);
        for (var i2 = 0, length = radios.length; i2 < length; i2++) {
            radios[i2].disabled = true;
        }
    }
    console.log("radios disabled");

}



window.addEventListener('load', validateForm)

