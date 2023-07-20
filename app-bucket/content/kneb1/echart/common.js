
// UI funcs
function nextStep() {
    myChart.setOption({
        dataset: getData(),
        // series: {
        //     data: makeRandomData()
        // }
    });
    return false;
}


function forever() {
    setInterval(() => {
        myChart.setOption({
            dataset: getData(),
            // series: {
            //     data: makeRandomData()
            // }
        });
    }, 200);
    return false;
}



let initPage = (inst) => {
    // const evt = new Event("input");
    const evt = new Event("change");
    // let checkBx = document.getElementById(elID);
    // checkBx.dispatchEvent(evt);

    // if (frm) {
    //     frm.addEventListener('submit', validateForm);
    // }

    var slider   = document.getElementById("sliderInner");

    var safe     = document.getElementsByName("share_safe")[0];
    var risky    = document.getElementsByName("share_risky")[0];

    var safeBG   = document.getElementById("share_safe_bg");
    var riskyBG  = document.getElementById("share_risky_bg");

    // init
    // if (safeBG && safeBG.value != ""  && safeBG.value != 0) {
    if (safeBG && safeBG.value != "" ) {
        safe.value  = safeBG.value;
        risky.value = riskyBG.value;
        slider.value = risky.value;
    }

    try {
        safe.value = 100 - slider.value;
        risky.value = slider.value;
    } catch (error) {

    }



    // update
    let funcUpdate = function () {

        try {
            safe.value = 100 - this.value;
            risky.value = this.value;

            // console.log(`safe.value = ${safe.value}`)

            if(safeBG){
                safeBG.value = safe.value;
                // console.log(`safeBG.value = ${safeBG.value}`)
            } else {
                console.log(`safeBG undefined`)
            }
            if(riskyBG){
                riskyBG.value = risky.value;
            }

        } catch (error) {

        }

    } // end of update func


    // slider.oninput = funcUpdate


    console.log(`page init complete`)
}

// init checkbox subgroups show/hide;
window.addEventListener('load', initPage, false);



// console.log(`common.js loaded`)