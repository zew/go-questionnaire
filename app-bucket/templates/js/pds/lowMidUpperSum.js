function funcInner{{.inp_1}}(){

    // let inp1 = document.forms.frmMain.xx_main.value;
    let  inp1 = document.forms.frmMain.{{.inp_1}}.value;
    let  inp2 = document.forms.frmMain.{{.inp_2}}.value;
    let  inp3 = document.forms.frmMain.{{.inp_3}}.value;
    let  inp4 = document.forms.frmMain.{{.inp_4}}.value;

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
    let i4 = 0
    if (inp3 != "") {
        i4 = parseInt(inp4, 10);
    }

    // console.log("inp1-4 string:  ", inp1, inp2, inp3, inp4 );
    // console.log("inp1-4 integer: ", i1, i2 + i3 + i4, i2, i3, i4 );

    let suspicious = false;
    
    // parts adding up
    if (i1 != 0) {
        let sum1 = i2 + i3 + i4;
        if (sum1 !=  i1) {
            suspicious = true;
        }
    }

    return suspicious;

}

function funcOuter{{.inp_1}}(event) {

    if (funcInner{{.inp_1}}()) {
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

// non global block
{
    let frm = document.forms.frmMain;
    if (frm) {
        frm.addEventListener('submit', funcOuter{{.inp_1 }});
    }
    console.log("   funcOuter{{.inp_1 }} registered")
}
