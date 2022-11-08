function funcInner{{.inp_1}}(){

    // let inp1 = document.forms.frmMain.xx_main.value;

    let  iSum = document.forms.frmMain.{{.inp_1}}.value;
    let  inp1 = document.forms.frmMain.{{.inp_2}}.value;
    let  inp2 = document.forms.frmMain.{{.inp_3}}.value;
    let  inp3 = document.forms.frmMain.{{.inp_4}}.value;

    let iS = 0
    if (iSum != "") {
        iS = parseInt(iSum, 10);
    }
    let i1 = 0
    if (inp1 != "") {
        i1 = parseInt(inp1, 10);
    }
    let i2 = 0
    if (inp2 != "") {
        i2 = parseInt(inp2, 10);
    }
    let i3 = 0
    if (inp3 != "") {
        i3 = parseInt(inp3, 10);
    }

    console.log("inp1-4 string:  ", iSum, inp1, inp2, inp3, " --" + iSum + inp1 + inp2 + inp3+ "-- " );
    console.log("inp1-4 integer: ", iS , " = ", i1 + i2 + i3, " ; summanden: ", i1, i2, i3 );

    let suspicious = false;
    
    // parts adding up
    if (iS != 0) {
        let sum1 = i1 + i2 + i3;
        if (sum1 !=  iS) {
            suspicious = true;
        }
    }

    return suspicious;

}

function funcOuter{{.inp_1}}(event) {

    if (funcInner{{.inp_1}}()) {
        // alert("{{.msg}}");
        let doContinue = window.confirm("{{.msg}} {{.inp_1}}");
        if (doContinue) {
            return true;
        }
        event.preventDefault(); // not only return false - but also preventDefault()
        return false;
    }

    return true;

}

// non global block
{
    let frm = document.forms.frmMain;
    if (frm) {
        frm.addEventListener('submit', funcOuter{{.inp_1}});
    }
    console.log("   funcOuter{{.inp_1 }} registered")
}
