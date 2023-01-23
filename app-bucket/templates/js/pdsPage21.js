{

    // triggers == 0 should disable the entire column

    // pdsPage11-b.js only disables input[number]
    //  this also disables input[range]


    let triggers = [
        "ac1_tt1_q2a_vol_realized_loans",
        "ac1_tt2_q2a_vol_realized_loans",
        "ac1_tt3_q2a_vol_realized_loans",

        "ac2_tt1_q2a_vol_realized_loans",
        "ac2_tt2_q2a_vol_realized_loans",

        "ac3_tt1_q2a_vol_realized_loans",
        "ac3_tt2_q2a_vol_realized_loans",
    ];


    // destination rumps
    // safari cannot read it inside func, if declared let or const
    var destRumps = [
        "q2b_time_to_maturity",
        "q2b_time_to_maturity",
        "q2b_time_to_maturity",

        "q2c_gross_irr",
        "q2c_gross_irr",
        "q2c_gross_irr",

        "q2d_gross_moic",
        "q2d_gross_moic",
        "q2d_gross_moic",
    ];



    function myOnchange(evt) {
        let prefix = evt.srcElement.name.substring(0, 8); //      "ac2_tt2_" out of"ac2_tt2_q11a_numtransact_main",
        // console.log("myChange-21", evt.srcElement.name, evt.srcElement.value, prefix);
        for (let i0 = 0; i0 < destRumps.length; i0++) {
            const elID = destRumps[i0];
            let inpDst = document.getElementById(prefix + elID);
            if (inpDst) {
                if (inpDst.type != "range") {
                    // console.log("myChange-21--2 found", prefix + elID);
                    if (evt.srcElement.value == 0 && !evt.srcElement.value == '') {
                        inpDst.disabled = true;
                    } else {
                        inpDst.disabled = false;
                    }
                } else {
                    let noAnsw = document.getElementById(prefix + elID + "_noanswer");
                    if (noAnsw) {
                        // console.log("myChange-21--3 found slider", prefix + elID);
                        if (evt.srcElement.value == 0 && !evt.srcElement.value == '') {
                            noAnsw.checked = true;
                            const evt = new Event("input");
                            noAnsw.dispatchEvent(evt);
                        } else {
                            if (false) {
                                noAnsw.checked = false;
                                // this is not perfect yet;
                                // we should distinguish depending on [prefix + elID]_hidd 
                                // has a value or not.
                                inpDst.classList.remove("noanswer");
                                inpDst.dataset.dirty = "";                                
                            }

                        }
                    }
                }
            }
        }
        return true;
    }


    for (let i0 = 0; i0 < triggers.length; i0++) {
        const elID = triggers[i0];
        let inp = document.getElementById(elID);
        if (inp) {
            // checkB[0].addEventListener('change', myChange);
            inp.addEventListener('change', myOnchange);
            inp.addEventListener('input',  myOnchange);
            // console.log("onchange ", inp.name, "'pdsPage11-b'");
        }
    }


    let initPage = (inst) => {
        // const evt = new Event("input"); // cursor keys
        const evt = new Event("change");
        for (let i0 = 0; i0 < triggers.length; i0++) {
            const elID = triggers[i0];
            let inp = document.getElementById(elID);
            if (inp) {
                // checkB[0].addEventListener('change', myChange);
                inp.dispatchEvent(evt);
            }
        }
    }

    // init checkbox subgroups show/hide;
    window.addEventListener('load', initPage, false);


    console.log("script complete - 'pdsPage21'");


}