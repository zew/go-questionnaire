const frmM = document.forms.frmMain;

const inputNamesMatrix = {
    "inf2022": ["inf2022_under1", "inf2022_between1and2", "inf2022_between2and3", "inf2022_above3"],
    "inf2023": ["inf2023_under1", "inf2023_between1and2", "inf2023_between2and3", "inf2023_above3"],
    "inf2024": ["inf2024_under1", "inf2024_between1and2", "inf2024_between2and3", "inf2024_above3"],
};

const globRowKeys = Object.keys(inputNamesMatrix);



let validateRowVals = evt => {


    let srcName = evt.srcElement.name;
    let rowKey = srcName.substring(0, 7); // neighbours            

    const inpNames = inputNamesMatrix[rowKey]; // siblings on row

    let inps = inpNames.map(nme => {
        return frmM[nme]; // input objects on row
    });

    let filled = false;  // at least one field is filled ?
    let vals = inps.map(inp => {
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
                return false;
            }

        }

    }

    return true;

}




// register on each input
globRowKeys.forEach( key => {
    let inps = inputNamesMatrix[key].map(nme => {
        return frmM[nme];
    });
    inps.forEach( inp => {
        // console.log(`listener added - ${inp.name}`);
        inp.addEventListener('blur', validateRowVals);
    });
});


function validateRowValsAll(evt) {

    // validateRowVals for every row
    let rowStates = globRowKeys.map( key => {
        let fakeInp = document.createElement("input");
        fakeInp.name = inputNamesMatrix[key][0];
        let fakeEvt = {
            type: "submit",
            srcElement: fakeInp,
        }
        return validateRowVals( fakeEvt );
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
    frmM.addEventListener('submit', validateRowValsAll);
}

