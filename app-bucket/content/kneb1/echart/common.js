
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

function fcSpin(upOrDown){
    let inp = document.getElementById("sparbetrag")
    if (inp) {
        console.log(`upOrDown = ${upOrDown}, val = ${inp.value}`)
        if (upOrDown==='up') {
            inp.value =  parseInt(inp.value) + 10;
        }
        if (upOrDown==='down') {
            inp.value =  parseInt(inp.value) - 10;
        }            
        console.log(`upOrDown = ${upOrDown}, val = ${inp.value}`)
    }
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
        // slider.value = risky.value;
    }

    try {
        safe.value = 100 - slider.value;
        risky.value = slider.value;
    } catch (error) {

    }

    let knobs = [...document.getElementsByClassName("knob")];

    let knobReset = kn => kn.classList.remove("knob-inverse")
    

    let knobClick = (evt) => {
        try {
            let src = evt.srcElement;
            let inner = src.innerHTML;
            inner = inner.replace("&nbsp;%","");
            let val = parseInt(inner)

            safe.value = 100 - val
            risky.value = val;

            if(safeBG){
                safeBG.value = safe.value;
                riskyBG.value = risky.value;
                // console.log(`safeBG.value = ${safeBG.value}`)
            } else {
                console.error(`safeBG undefined`)
            }

            knobs.forEach(knobReset);
            src.classList.add("knob-inverse")


        } catch (err) {
            console.error(`knob click error`, err)
        }
    }
    let assignEvent = function(kn) {
        kn.onclick = knobClick
        // console.log("test", kn);
    }
    console.log(`found ${knobs.length} knobs`)
    knobs.forEach(assignEvent);



    console.log(`page init complete`)
}

// init checkbox subgroups show/hide;
window.addEventListener('load', initPage, false);



// console.log(`common.js loaded`)