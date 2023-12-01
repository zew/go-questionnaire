const elSum = document.getElementById("sum");


function inputsChanged(evt) {

    let sum = 0;

    for (let i = 0, inpS; inpS = inps[i++];) {

        let inp = document.getElementById(inpS);

        if(inp){
            // fine
        } else {
            console.error(`input named ${inpS} not found`);
        }

        if (inp.type === "hidden") {
            continue;
        }

        if (inp.type !== "number") {
            continue;
        }

        if (inp.value != inp.defaultValue) {
            // sum++;
        }

        if (inp.value > 0) {
            let vl = parseFloat(inp.value)
            sum += vl;
        }
    }


    console.log(`checking named ${evt.target.name}-${evt.target.type} - sum at ${sum}`);

    let sumRnd = Math.round(sum*10)/10;
    let sumStr = `${sumRnd}`;
    sumStr = sumStr.replace(".",",");
    elSum.value = sumStr;




    if (evt.type == "submit") {
        // alert("submit");
        const not100 = Math.abs(sum-100) > 0.2;
        const not0   = Math.abs(sum) > 0.2;

        // console.log(`not100 ${not100}, not0 ${not0}`);

        if(not100 && not0){
            alert("{{.msg2}}");
            evt.preventDefault();
            return false;
        }
    }

    // change event -
    if(false &&  sum>100){
        // alert("no more");
        alert("{{.msg1}}");
        // sum =  0;

        try {
            // evt.target.checked = false;
            // evt.target.value = 0;
            // sum--;
        } catch (err) {
            console.error(err);
        }
        // break;
    }


}


// addEventListener is cumulative
window.addEventListener("load", function (evt) {

    for (let i = 0, inpS; inpS = inps[i++];) {

        let inp = document.getElementById(inpS);

        if(inp){
            // fine
        } else {
            console.error(`input named ${inpS} not found`);
        }

        if (inp.type === "hidden") {
            continue;
        }

        if (inp.type !== "number") {
            continue;
        }

        // document.forms.frmMain.q03_no_investing[0].addEventListener('change', myOnchange);
        inp.addEventListener('change', inputsChanged);

        // console.log(`change listener assigned to ${inpS} - ${inp.type}`);

    }


    // init
    let inp01 = document.getElementById(inps[0]);
    const evtInit = new Event("change");
    inp01.dispatchEvent(evtInit);


    document.forms.frmMain.addEventListener('submit', inputsChanged);


});






var inps = {{.inps}};

console.log("inputs", inps);




