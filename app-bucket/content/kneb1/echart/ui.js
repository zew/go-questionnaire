"strict mode";

var myChart;



// UI funcs
function nextStep() {
    let dta = dataObject.computeData();
    refresh(myChart, dataObject);
    // console.log("next step complete", myChart, dta)
    return false;
}


function forever() {
    let dta = dataObject.computeData();
    setInterval(() => {
        refresh(myChart, dataObject);
    }, 200);
    return false;
}



let pageLoaded = (inst) => {


    // const evt = new Event("input");
    const evt = new Event("change");
    // let checkBx = document.getElementById(elID);
    // checkBx.dispatchEvent(evt);

    // if (frm) {
    //     frm.addEventListener('submit', validateForm);
    // }


    let sbChange = (evt) => {
        let src = evt.srcElement;

        sb = src.value;
        sbInpBG.value = src.value;

        refresh(myChart, dataObject);

        console.log(`sbChange ${sb}`)
    }

    sbInp.onchange = sbChange


    // "Sparbetrag" increase and decrease - onclick event handler
    let fcSpin = (evt,upOrDown) => {
        // console.log(`upOrDown = ${upOrDown}`);
        let inp = document.getElementById("sparbetrag")
        if (inp) {
            // console.log(`upOrDown = ${upOrDown}, val = ${inp.value}`)
            if (upOrDown==='up') {
                inp.value =  parseInt(inp.value) + 10;
            } else if (upOrDown==='down') {
                inp.value =  parseInt(inp.value) - 10;
            }
            console.log(`upOrDown = ${upOrDown}, val = ${inp.value}`)

            evt.preventDefault();

            sb = inp.value;
            sbInpBG.value = inp.value;
            refresh(myChart, dataObject);

        }
        // sbChange(evt)
    }

    let fcSpinUp = (evt) => {
        fcSpin(evt, "up")
    }
    let fcSpinDown = (evt) => {
        fcSpin(evt, "down")
    }

    let elUp   = document.getElementById("elUp");
    elUp.onclick = fcSpinUp
    let elDown = document.getElementById("elDown");
    elDown.onclick = fcSpinDown


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

            refresh(myChart, dataObject);
            console.log(`knobClick new val ${riskyBG.value}`)

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

    let assignEvents = function(kn) {
        kn.onclick = knobClick
        kn.onkeyup = knobKey
        // console.log("test", kn);
    }
    // console.log(`found ${knobs.length} knobs`)
    knobs.forEach(assignEvents);


    // set init knob visual
    knobs.forEach(el => {
        //   console.log(`${el.innerHTML} - ${riskyBG.value}`)
        if (el.innerHTML.includes(`${riskyBG.value}`)) {
            el.classList.add("knob-inverse")
        }
    });

    let arrayFVs = dataObject.FVs()

    let elFV = document.getElementById('elFV');
    if (elFV) {
        elFV.innerHTML = knebelFormat(arrayFVs[1])
    } else {
        console.error(`did not find elFV`)
    }

    let elFV05 = document.getElementById('elFV05');
    if (elFV05) {
        elFV05.innerHTML = knebelFormat(arrayFVs[0])
    } else {
        console.error(`did not find elFV95`)
    }

    let elFV95 = document.getElementById('elFV95');
    if (elFV95) {
        elFV95.innerHTML = knebelFormat(arrayFVs[2])
    } else {
        console.error(`did not find elFV95`)
    }




    //
    let chartDom = document.getElementById('chart_container');
    myChart = echarts.init(chartDom);

    optEchart && myChart.setOption(optEchart);
    console.log(`echart config and creation complete`)


    let elReset = document.getElementById('elReset');
    if (elReset) {
        elReset.onclick = (evt) => {
            sb = 100
            if (sbInp) {
                sbInp.value = sb;
            }
            if (sbInpBG) {
                sbInpBG.value = sb;
            }
            safeBG.value = 50
            riskyBG.value = 100 - safeBG.value
            // set knob visual
            knobs.forEach(knobReset);
            knobs.forEach( el => {
                if (el.innerHTML.includes(`${riskyBG.value}`)) {
                    el.classList.add("knob-inverse")
                }
            });

            refresh(myChart, dataObject);


            console.log(`reset clicked`)
            evt.preventDefault();
        }
    } else {
        console.error(`did not find elReset`)
    }


    // addMarkPoint(myChart);


    console.log(`page loaded complete`)
}

// init checkbox subgroups show/hide;
window.addEventListener('load', pageLoaded, false);



// console.log(`ui.js for echart loaded`)


