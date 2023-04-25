const frmM = document.forms.frmMain;


// ["inf2023", "inf2024", "inf2025"]
let  yrs = {{.Yrs}}
let ivls = {{.Ivls}}


// let inputNamesMatrix = {
//     "inf2023": ["inf2023_under4", "inf2023_between4and6", "inf2023_between6and8","inf2023_between8and10", "inf2023_above10"],
//     "inf2024": ["inf2024_under4", "inf2024_between4and6", "inf2024_between6and8","inf2024_between8and10", "inf2024_above10"],
//     "inf2025": ["inf2025_under4", "inf2025_between4and6", "inf2025_between6and8","inf2025_between8and10", "inf2025_above10"],
// };
var inputNamesMatrix = {}
var inps = []

yrs.forEach(yr => {
    inputNamesMatrix[yr] = []
    ivls.forEach(ivl => {
        const nm = `${yr}_${ivl}`;
        inputNamesMatrix[yr].push(nm)
        inps.push(frmM[nm])
    });
});




const globRowKeys = Object.keys(inputNamesMatrix);

// console.log(inputNamesMatrix)
// console.log(inps)
// console.log(globRowKeys)



let validateRow = evt => {


    let srcName = evt.srcElement.name;
    // neighbours - identified by prefix inf2023_..., inf2024_... 
    let rowKey = srcName.substring(0, 7); 

    const inpNames = inputNamesMatrix[rowKey]; // siblings on row

    let inpsRow = inpNames.map(nme => {
        return frmM[nme]; // input objects on row
    });

    let filled = false;  // at least one field is filled ?
    let vals = inpsRow.map(inp => {
        let val = 0;
        if (inp.value != "") {
            val = parseInt(inp.value, 10);
            filled = true;
        }
        return val;
    });
    // console.log(`rowKey ${rowKey} - vals ${vals}`);

    if (!filled) {
        return true;  // no validation for completely empty row
    }


    let sum = vals.reduce((x, y) => x + y);
    console.log(`rowKey ${rowKey} - vals ${vals} - sum ${sum}`);


    if (sum > 0) {
        if (sum != 100) {


            // console.log(`evt.type ${evt.type}`, { evt });
            let doAsk = false;
            if (evt.type == "blur") {
                doAsk = false; // ask on blur
                // doAsk = true; // ask on blur
                // console.log(`blur ${evt.srcElement.name}`, {evt});
            } else {
                // always ask on submit
                doAsk = true;
            }
            

            if (doAsk) {
                // alert("{{.msg}}");
                let doContinue = window.confirm("{{.msg}}");
                if (doContinue) {
                    return true;
                }

                // not only return false - but also preventDefault()
                if (evt.preventDefault) {
                    evt.preventDefault(); 
                }

                try {
                    // let inpFoc = document.getElementById(inputNamesMatrix["inf2023"][0] );
                    let inpFoc = inpsRow[0];
                    if (inpFoc) {
                        inpFoc.focus();
                    }                    
                } catch (error) {
                    
                }
                return false;
            }

        }

    }

    return true;

}




// register on each input
inps.forEach( inp => {
    if (inp) {
        // console.log(`listener added - ${inp.name}`);
        inp.addEventListener('blur', validateRow);
    } else {
        console.log(`cannot find - ${inp}`);
    }
});



function validateAllRows(evt) {

    // validateRow for every row
    let rowStates = globRowKeys.map( key => {
        let fakeInp = document.createElement("input");
        fakeInp.name = inputNamesMatrix[key][0];
        let fakeEvt = {
            type: "submit",
            srcElement: fakeInp,
        }
        return validateRow( fakeEvt );
    });
    
    
    // combine row validations;
    // any non-true => state = false
    let state = rowStates.reduce(  (x, y) => x && y   ); 
    console.log(`rowStates ${rowStates} => state ${state}`);

    if (!state) {
        evt.preventDefault();
    }
    return state;

}


// register on form
if (frmM) {
    frmM.addEventListener('submit', validateAllRows);
}

