const frmM = document.forms.frmMain;
const inpLeave = document.forms.frmMain.leave_survey[0];

let checked = document.forms.frmMain.leave_survey[0].checked;

let  param1 = "{{.xxx1}}"


async function submitFrmMainNoReload() {
    try{
        const frm = document.forms["frmMain"];
        if (!frm) {
            console.error("form frmMain not found");
            return;
        }
        const url    = frm.action;
        const method = (frm.method || "POST").toUpperCase();
        const data   = new FormData(frm);
        // fetch prevents browser navigation/page reload
        const rsp = await fetch(url, {
            method:       method,
            body:         data,
        });
        if (!rsp.ok) {
            console.error("form submit failed", rsp.status);
            return;
        }
        const txt = await rsp.text();
        console.log("values uploaded");
    } catch (exc) {
        handleExc(exc, `submitFrmMainNoReload() `);
    }
}      


let getUserInput = evt => {

    console.log(`evt.type ${evt.type}`, { evt });

    if (inpLeave.checked) {
        // alert("{{.msg}}");

        console.log(` leave survey is ${inpLeave.checked} `)

        let doContinue = window.confirm("{{.msg}}");
        if (doContinue) {
            submitFrmMainNoReload()
            return true;
        }

        // not only return false - but also preventDefault()
        if (evt.preventDefault) {
            evt.preventDefault();
        }
        return false;
    }


    return true;


}


inpLeave.addEventListener('input', getUserInput);
console.log(` event  ${inpLeave.id} `)





// register on form
if (frmM) {
    // frmM.addEventListener('submit', validateRow);
}

