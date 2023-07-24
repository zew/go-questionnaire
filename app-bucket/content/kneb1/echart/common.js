var myChart;

function refresh() {

    dataObject.resetData()

    // setOption or resize
    myChart.resize();

    if (true) {
        myChart.setOption({
            series: [
                {
                    data: dataObject.computeData(),
                },
                {
                    data: dataObject.computeData(),
                },
                {
                    data: dataObject.computeData(),
                },
            ]
        });                
    }
}


// UI funcs
function nextStep() {
    let dta = dataObject.computeData();
    refresh();
    // console.log("next step complete", myChart, dta)
    return false;
}


function forever() {
    let dta = dataObject.computeData();
    setInterval(() => {
        refresh();
    }, 200);
    return false;
}

// "Sparbetrag" increase and decrease - onclick event handler
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



    let sbChange = (evt) => {
        let src = evt.srcElement;

        sb = src.value;

        sbInpBG.value = src.value;

        refresh();

        console.log(`sbChange ${sb}`)
    }

    sbInp.onchange = sbChange


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

            refresh();
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
    console.log(`found ${knobs.length} knobs`)
    knobs.forEach(assignEvents);




    // 
    let chartDom = document.getElementById('chart_container');
    // console.log(chartDom);
    myChart = echarts.init(chartDom);
    
    optEchart && myChart.setOption(optEchart);
    console.log(`echart config and creation complete`)
    


    console.log(`page init complete`)
}

// init checkbox subgroups show/hide;
window.addEventListener('load', initPage, false);



// console.log(`common.js loaded`)


