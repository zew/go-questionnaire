"strict mode";

// sbInp, sbInpBG - declared and initialized in echart-config.mjs
// myChart, dataObject too


let pageLoaded = (inst) => {

    const evt = new Event("change");

    let sbChange = (evt) => {
        let src = evt.srcElement;
        sb = src.value;
        refresh(myChart, dataObject);
        console.log(`sbChange ${sb}`)

        if (false) {
            simHist.push({ "sb_ch": sb });
            simHistBG.value = JSON.stringify(simHist);
        }
        // console.log(`simHistBG.value ${simHistBG.value}`)
    }

    sbInp.onchange = sbChange


    //
    let chartDom = document.getElementById('chart_container');
    myChart = echarts.init(chartDom);

    optEchart && myChart.setOption(optEchart);
    console.log(`echart config and creation complete`)



    console.log(`pageLoaded() complete`)
}


window.addEventListener('load', pageLoaded, false);


