
// UI funcs
function nextStep() {
    let dta = getData();
    myChart.setOption({
        // dataset: dta,
        series: [
            {
              data: dta,
            },
            {
              data: dta,
            }
        ],
    });
    console.log("next step complete", myChart, dta)
    return false;
}


function forever() {
    let dta = getData();
    setInterval(() => {
        myChart.setOption({
            // dataset: dta,
            series: [
                {
                    data: dta,
                },
                {
                    data: dta,
                }
            ],    
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

    // var safe     = document.getElementsByName("share_safe")[0];
    // var risky    = document.getElementsByName("share_risky")[0];

    var safeBG   = document.getElementById("share_safe_bg");
    var riskyBG  = document.getElementById("share_risky_bg");

    // init
    // if (safeBG && safeBG.value != ""  && safeBG.value != 0) {
    if (safeBG && safeBG.value != "" ) {
        // safe.value  = safeBG.value;
        // risky.value = riskyBG.value;
    }

    // try {
    //     safe.value = 100 - slider.value;
    //     risky.value = slider.value;
    // } catch (error) {

    // }

    let knobs = [...document.getElementsByClassName("knob")];

    let knobReset = kn => kn.classList.remove("knob-inverse")

    let knobClick = (evt) => {
        try {
            let src = evt.srcElement;
            let inner = src.innerHTML;
            inner = inner.replace("&nbsp;%","");
            let val = parseInt(inner)

            safeBG.value = 100 - val
            riskyBG.value = val;


            knobs.forEach(knobReset);
            src.classList.add("knob-inverse")


        } catch (err) {
            console.error(`knob click error`, err)
        }
    }

    let knobKey = (evt) => {
        if (evt.code !== "Tab") {
            // consume evt - so it doesn't get handled twice - unless user moves focus
            evt.preventDefault();
        }
        if (evt.code === "Space" ||  evt.code === "Enter") {
            knobClick(evt)
        }
    }

    let assignEvent = function(kn) {
        kn.onclick = knobClick
        kn.onkeyup = knobKey
        // console.log("test", kn);
    }
    console.log(`found ${knobs.length} knobs`)
    knobs.forEach(assignEvent);



    console.log(`page init complete`)
}

// init checkbox subgroups show/hide;
window.addEventListener('load', initPage, false);



// console.log(`common.js loaded`)