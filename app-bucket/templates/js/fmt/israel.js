
function check(evt) {

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

        if (inp.type !== "checkbox") {
            continue;
        }

        if (inp.checked) {
            sum++;
        }
    }


    console.log(`checking named ${evt.target.name}-${evt.target.type} - sum at ${sum}`);

    if(sum>3){
        // alert("no more");
        alert("{{.msg}}");
        sum =  0;

        try {
            evt.target.checked = false;
            sum--;
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

        if (inp.type !== "checkbox") {
            continue;
        }

        // document.forms.frmMain.q03_no_investing[0].addEventListener('change', myOnchange);
        inp.addEventListener('change', check);

        // console.log(`change listener assigned to ${inpS} - ${inp.type}`);

    }
});






var inps = {{.inps}};

console.log("inputs", inps);


