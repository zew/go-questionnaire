function func01{{.inp_1}}(){

    let inp1 = document.forms.frmMain.xx_main.value;

    let inp2 = document.forms.frmMain.xx_low.value;
    let inp3 = document.forms.frmMain.xx_mid.value;
    let inp4 = document.forms.frmMain.xx_upper.value;

     inp1 = document.forms.frmMain.{{.inp_1}}.value;

     inp2 = document.forms.frmMain.{{.inp_2}}.value;
     inp3 = document.forms.frmMain.{{.inp_3}}.value;
     inp4 = document.forms.frmMain.{{.inp_4}}.value;

    // console.log("inp1-3: ", inp1, inp2, inp3);

    let i1 = 0
    if (inp1 != "") {
        let i1 = parseInt(inp1, 10);
    }
    let i2 = 0
    if (inp2 != "") {
        let i2 = parseInt(inp2, 10);
    }
    let i3 = 0
    if (inp3 != "") {
        let i3 = parseInt(inp3, 10);
    }
    let i4 = 0
    if (inp3 != "") {
        let i4 = parseInt(inp4, 10);
    }
    // console.log("inp1-3 integer: ", i1, i2, i3);

    let suspicious = false;
    
    // expectation between extremes?
    if (i1 != 0) {
        let sum1 = i2 + i3 + i4;
        if (sum1 !=  i1) {
            suspicious = true;
        }

    }

    return suspicious;

}

function func02{{.inp_1}}(event) {


    if (func01{{.inp_1}}()) {
        // alert("{{.msg}}");
        let doContinue = window.confirm("{{.msg}}");
        if (doContinue) {
            return true;
        }
        event.preventDefault(); // not only return false - but also preventDefault()
        return false;
    }

    return true;

}


let frm = document.forms.frmMain;
if (frm) {
    frm.addEventListener('submit', func02{{.inp_1}});
}
