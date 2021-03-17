function demo(event) {
    if (confirm("Press a button!")) {
        txt = "You pressed OK!";
        console.log(txt);
        return true;
    } else {
        txt = "You pressed Cancel!";
        console.log(txt);
        event.preventDefault(); // not only return false - but also preventDefault()
        return false;
    }
}

function validateForm(event) {

    var inp1 = document.forms.frmMain.dax_erw.value;
    var inp2 = document.forms.frmMain.dax_min.value;
    var inp3 = document.forms.frmMain.dax_max.value;
    // console.log("inp1-3: ", inp1, inp2, inp3);

    var i1 = 0
    if (inp1 != "") {
        var i1 = parseInt(inp1, 10);
    }
    var i2 = 0
    if (inp2 != "") {
        var i2 = parseInt(inp2, 10);
    }
    var i3 = 0
    if (inp3 != "") {
        var i3 = parseInt(inp3, 10);
    }
    // console.log("inp1-3 integer: ", i1, i2, i3);

    var suspicious = false;
    if (i1 != 0) {
        if (i2 != 0 && i1 < i2) {
            suspicious = true;
        }

        if (i3 != 0 && i1 > i3 ) {
            suspicious = true;
        }
    }

    // min > maix
    if (i2 != 0 && i3 != 0 && i2 > i3) {
        suspicious = true;
    }


    if (suspicious) {
        // alert("{{.msg}}");
        var doContinue = window.confirm("{{.msg}}");
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
