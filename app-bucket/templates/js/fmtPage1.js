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

function mustAddUp100(event) {

    var inp1 = document.forms.frmMain.y_probgood.value;
    var inp2 = document.forms.frmMain.y_probnormal.value;
    var inp3 = document.forms.frmMain.y_probbad.value;
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

    var sum = i1 + i2 + i3;
    console.log("sum",sum);


    if (sum > 0) {        
        if (sum != 100 ) {

            // console.log("event.type", event.type);
            var doAsk = false;
            if (event.type=="blur") {
                if (inp1 != "" && inp2 != "" && inp3 != "") {
                    // doAsk = true;
                    console.log("show error msg");
                }
            } else {
                // submit
                doAsk = true;
            }

            if (doAsk) {
                // alert("{{.msg}}");
                var doContinue = window.confirm("{{.msg}}");
                if (doContinue) {
                    return true;
                }
                event.preventDefault(); // not only return false - but also preventDefault()
                return false;
            }
        }
    }

    return true;

}


var frm = document.forms.frmMain;
if (frm) {
    frm.addEventListener('submit', mustAddUp100);    
}

if (document.forms.frmMain.y_probgood) {
    document.forms.frmMain.y_probgood.addEventListener('blur',mustAddUp100);
}
if (document.forms.frmMain.y_probnormal) {
    document.forms.frmMain.y_probnormal.addEventListener('blur',mustAddUp100);
}
if (document.forms.frmMain.y_probbad) {
    document.forms.frmMain.y_probbad.addEventListener('blur',mustAddUp100);
}



