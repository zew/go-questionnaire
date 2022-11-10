function funcInner{{.InpMain}}(){

    // let inp1 = document.forms.frmMain.xx_main.value;
    // let totalInp = document.forms.frmMain.{{.InpMain}};
    let nameTotal = "{{.InpMain}}" + "_main";
    let totalInp = document.getElementById(nameTotal);
    if (totalInp) {
        // 
    } else {
        console.log(nameTotal + " does not exist");
        return;
    }

    let totalInpVal = totalInp.value;
    let totalInpInt = 0
    if (totalInpVal != "") {
        totalInpInt = parseInt(totalInpVal, 10);
        // alert(nameMain + " value " + iSumStr + "; " + iS);
    }

    // let summandNames = ["name1", "name2"];
    let summandNames = [{{.SummandNames}}];
    // let summandVals  = [1, 2];
    let summandValsStr  = [];
    let summandValsInt  = [];
    let sum = 0;


    for (let i1 = 0; i1 < summandNames.length; i1++) {
        const inpLp = document.getElementById( summandNames[i1] );
        summandValsStr.push(inpLp.value);        
        if (inpLp.value != "") {
            let iVal = parseInt(inpLp.value, 10);
            summandValsInt.push( iVal);
            sum += iVal;
        } else {
            summandValsInt.push(0);
        }
    }



    let suspicious = false;
    
    // parts adding up
    if (sum != 0 || totalInpInt != 0) {
        if (sum != totalInpInt) {

            console.log("total:    ", nameTotal, totalInpInt);
            console.log("summands str: ", summandValsStr);
            console.log("summands int: ", summandValsInt, " = " , sum);
            
            totalInp.focus();

            suspicious = true;
        }
    }

    return suspicious;

}

function funcOuter{{.InpMain}}(event) {

    if (funcInner{{.InpMain}}()) {
        // alert("{{.msg}}");
        let doContinue = window.confirm("{{.msg}} {{.InpMain}}");
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
        frm.addEventListener('submit', funcOuter{{.InpMain}});
    }
    console.log("   funcOuter{{.inp_1 }} registered")
}
