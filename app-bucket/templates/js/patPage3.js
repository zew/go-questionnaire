function validateForm(event) {

    let strInputs = ["","",""];
    let intInputs = [0,0,0];


    for (let i1 = 0; i1 < 3; i1++) {
        var radioName = "q1_seq" + (i1+1+3) + "_r"  // q1_seq1_r
        var radios1 = document.getElementsByName(radioName);
        for (var i2 = 0, length = radios1.length; i2 < length; i2++) {
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

  
    var sum = intInputs[0] + intInputs[1] + intInputs[2];
    console.log("sum",sum);

    if (sum > 0) {        
        if (intInputs[0] < 1 || intInputs[1] < 1 || intInputs[2] < 1) {
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
