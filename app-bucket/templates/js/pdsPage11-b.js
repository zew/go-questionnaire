{

    // sum of new transactions - changes 
    // should be copied to three other inputs in the same column

    let triggers = [
        "ac1_tt1_q11a_numtransact_main",
        "ac1_tt2_q11a_numtransact_main",
        "ac1_tt3_q11a_numtransact_main",

        "ac2_tt1_q11a_numtransact_main",
        "ac2_tt2_q11a_numtransact_main",

        "ac3_tt1_q11a_numtransact_main",
        "ac3_tt2_q11a_numtransact_main",
    ];


    // destination rumps
    // safari cannot read it inside func, if declared let or const
    var destRumps = [
            "q11a_numtransact_floatingrate",
            "q11a_numtransact_esgdoc",
            "q11a_numtransact_esgratchet",

            "q11b_voltransact_main",
            "q11d_volbysegm_low",
            "q11d_volbysegm_mid",
            "q11d_volbysegm_upper",

            // ac2
            "q11d_volbysegm_core",
            "q11d_volbysegm_coreplus",
            "q11d_volbysegm_valueadd",
            "q11d_volbysegm_opportun",


            "q11e_volbyreg_uk",
            "q11e_volbyreg_france",
            "q11e_volbyreg_ger",
            "q11e_volbyreg_othereur",

            "q11f_volbysect_energy",
            "q11f_volbysect_materials",
            "q11f_volbysect_industrials",
            "q11f_volbysect_consumer_disc",
            "q11f_volbysect_consumer_stap",
            "q11f_volbysect_healthcare",
            "q11f_volbysect_financials",
            "q11f_volbysect_information_tech",
            "q11f_volbysect_communication_svc",
            "q11f_volbysect_utilities",
            "q11f_volbysect_real_estate",
            "q11f_volbysect_other",

            "q11f_volbysect_office", 
            "q11f_volbysect_retail", 
            "q11f_volbysect_hotel", 
            "q11f_volbysect_residential", 
            "q11f_volbysect_logistics", 
            "q11f_volbysect_other",

            "q11f_volbysect_transportation",
            "q11f_volbysect_power",
            "q11f_volbysect_renewables",
            "q11f_volbysect_utilities",
            "q11f_volbysect_telecoms",
            "q11f_volbysect_social",
            "q11f_volbysect_other",


        ];



    function myOnchange(evt) {
        let prefix = evt.srcElement.name.substring(0, 8); //      "ac2_tt2_" out of"ac2_tt2_q11a_numtransact_main",
        // console.log("myChange-b", evt.srcElement.name, evt.srcElement.value, prefix);
        for (let i0 = 0; i0 < destRumps.length; i0++) {
            const elID = destRumps[i0];
            let inpDst = document.getElementById(prefix + elID);
            if (inpDst) {
                if (evt.srcElement.value == 0 && !evt.srcElement.value == '') {
                    inpDst.disabled = true;
                } else {
                    inpDst.disabled = false;
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


    console.log("script complete - 'pdsPage11-b'");


}