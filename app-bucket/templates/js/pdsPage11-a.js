{

    // triggers == 0 should disable the entire column
 

    let triggers = [
        "ac1_tt1_q11b_voltransact_main",
        "ac1_tt2_q11b_voltransact_main",
        "ac1_tt3_q11b_voltransact_main",

        "ac2_tt1_q11b_voltransact_main",
        "ac2_tt2_q11b_voltransact_main",

        "ac3_tt1_q11b_voltransact_main",
        "ac3_tt2_q11b_voltransact_main",

    ];

    let dests = {

        "ac1_tt1_q11b_voltransact_main": ["ac1_tt1_q11d_volbysegm_main", "ac1_tt1_q11e_volbyreg_main","ac1_tt1_q11f_volbysect_main"],
        "ac1_tt2_q11b_voltransact_main": ["ac1_tt2_q11d_volbysegm_main", "ac1_tt2_q11e_volbyreg_main","ac1_tt2_q11f_volbysect_main"],
        "ac1_tt3_q11b_voltransact_main": ["ac1_tt3_q11d_volbysegm_main", "ac1_tt3_q11e_volbyreg_main","ac1_tt3_q11f_volbysect_main"],

        "ac2_tt1_q11b_voltransact_main": ["ac2_tt1_q11d_volbysegm_main", "ac2_tt1_q11e_volbyreg_main", "ac2_tt1_q11f_volbysect_main"],
        "ac2_tt2_q11b_voltransact_main": ["ac2_tt2_q11d_volbysegm_main", "ac2_tt2_q11e_volbyreg_main", "ac2_tt2_q11f_volbysect_main"],

        "ac3_tt1_q11b_voltransact_main": ["ac3_tt1_q11d_volbysegm_main", "ac3_tt1_q11e_volbyreg_main", "ac3_tt1_q11f_volbysect_main"],
        "ac3_tt2_q11b_voltransact_main": ["ac3_tt2_q11d_volbysegm_main", "ac3_tt2_q11e_volbyreg_main", "ac3_tt2_q11f_volbysect_main"],

    };



    function myOnchange(evt) {
        // console.log("myChange-a", evt.srcElement.name, evt.srcElement.value);
        for (let i0 = 0; i0 < dests[evt.srcElement.id].length; i0++) {
            const elID = dests[evt.srcElement.id][i0];
            let inpDst = document.getElementById(elID);
            if (inpDst) {
                // checkB[0].addEventListener('change', myChange);
                inpDst.value = evt.srcElement.value;
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
            // console.log("onchange ", inp.name, "'pdsPage11-a'");
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

    // console.log("script complete - 'pdsPage11-a'");



}